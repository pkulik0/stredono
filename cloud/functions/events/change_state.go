package events

import (
	"context"
	"fmt"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func ChangeStateEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		DocDb:      true,
		RealtimeDb: true,
		Auth:       true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	changeState(ctx, w, r)
}

func changeState(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		log.Errorf("Missing rtdb")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	db, ok := ctx.GetDocDb()
	if !ok {
		log.Errorf("Missing doc db")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	var uid string

	overlayKey := r.URL.Query().Get("key")
	if overlayKey != "" {
		if err := rtdb.NewRef("Users").Child("Overlay").Child(overlayKey).Get(ctx.Ctx, &uid); err != nil {
			log.Printf("Failed to get user id for key %s | %v", overlayKey, err)
			http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
			return
		}
		if uid == "" {
			log.Printf("User id not found for key %s", overlayKey)
			http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
			return
		}
		log.Printf("Key %s belongs to user %s", overlayKey, uid)
	} else {
		uid = r.URL.Query().Get("uid")
		if uid == "" {
			log.Printf("Missing user id")
			http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
			return
		}

		token, ok := ctx.GetAuthToken(r)
		if !ok {
			http.Error(w, platform.UnauthorizedMessage, http.StatusUnauthorized)
			return
		}

		if token.UserId() != uid {
			var moderators []string
			if err := rtdb.NewRef("Moderators").Child("From").Child(uid).Get(ctx.Ctx, &moderators); err != nil {
				log.Printf("Failed to get moderators for user %s | %v", uid, err)
				http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
				return
			}

			isMod := false
			for _, mod := range moderators {
				if mod == token.UserId() {
					isMod = true
					break
				}
			}
			if !isMod {
				log.Printf("User %s is not allowed to modify events for user %s", token.UserId(), uid)
				http.Error(w, platform.UnauthorizedMessage, http.StatusUnauthorized)
				return
			}
		}
		log.Printf("User %s is allowed to modify events for user %s", token.UserId(), uid)
	}

	action := r.URL.Query().Get("action")
	if action == "" {
		log.Printf("Missing action")
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	eventId := r.URL.Query().Get("event")
	if eventId == "" {
		log.Printf("User %s - %s", uid, action)

		settingsRef := rtdb.NewRef("Data").Child(uid).Child("Settings").Child("Events")
		value := true
		switch action {
		case "unmute":
			value = false
			fallthrough
		case "mute":
			if err := settingsRef.Child("IsMuted").Set(ctx.Ctx, value); err != nil {
				log.Printf("Failed to mute | %v", err)
				http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
				return
			}
		case "resume":
			value = false
			fallthrough
		case "pause":
			if err := settingsRef.Child("IsPaused").Set(ctx.Ctx, value); err != nil {
				log.Printf("Failed to pause | %v", err)
				http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
				return
			}
		case "rerun":
			timeMinStr := r.URL.Query().Get("minutes")
			if timeMinStr == "" {
				log.Printf("Missing rerun time")
				http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
				return
			}
			timeMin, err := strconv.ParseInt(timeMinStr, 10, 64)
			if err != nil {
				log.Printf("Failed to parse rerun time | %v", err)
				http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
				return
			}
			if timeMin > 60 || timeMin < 1 {
				log.Printf("Rerun time invalid")
				http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
				return
			}

			dur := time.Duration(timeMin) * time.Minute
			rerunTimestamp := time.Now().Add(-dur).Unix()

			q := db.Collection("events").Where("Uid", "==", uid).Where("WasShown", "==", true).Where("Timestamp", ">", rerunTimestamp)
			err = db.RunTransaction(ctx.Ctx, func(ctx context.Context, tx modules.Transaction) error {
				iter := tx.Documents(q)
				defer iter.Stop()

				docs, err := iter.All()
				if err != nil {
					return err
				}

				rerunIDs := make([]string, 0)

				for _, snap := range docs {
					event := &pb.Event{}
					if err := snap.DataTo(event); err != nil {
						return err
					}

					if !event.WasShown {
						continue
					}
					event.WasShown = false

					if err := tx.Set(snap.Ref(), event, modules.DbOpts{}); err != nil {
						return err
					}

					rerunIDs = append(rerunIDs, event.ID)
				}

				if len(rerunIDs) > 0 {
					log.Printf("Rerun %d events: %v", len(rerunIDs), rerunIDs)
				}

				return nil
			})
			if err != nil {
				log.Printf("Failed to rerun event | %v", err)
				http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
				return
			}
			log.Printf("Rerun last %d minutes for user %s", timeMin, uid)
			return
		default:
			log.Printf("Unknown action %s", action)
			http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
			return
		}
		return
	}

	hasUserErr := false
	q := db.Collection("events").Where("ID", "==", eventId).Where("Uid", "==", uid).Limit(1)
	err := db.RunTransaction(ctx.Ctx, func(ctx context.Context, tx modules.Transaction) error {
		snap, err := tx.Documents(q).Next()
		if err != nil {
			return err
		}

		event := &pb.Event{}
		if err := snap.DataTo(event); err != nil {
			return err
		}

		switch action {
		case "cancel":
			fallthrough
		case "confirm":
			if event.WasShown {
				return nil
			}
			event.WasShown = true
		case "approve":
			if event.IsApproved {
				return nil
			}
			event.IsApproved = true
		case "rerun":
			if !event.WasShown {
				return nil
			}
			event.WasShown = false
		default:
			hasUserErr = true
			return fmt.Errorf("unknown action %s", action)
		}

		return tx.Set(snap.Ref(), event, modules.DbOpts{})
	})
	if err != nil {
		log.Printf("Failed to confirm event | %v", err)
		if hasUserErr {
			http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		} else {
			http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		}
		return
	}
	log.Printf("User %s - event %s - %s", uid, eventId, action)

	if _, err = w.Write([]byte("OK")); err != nil {
		log.Printf("Failed to write response | %v", err)
	}
}
