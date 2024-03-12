package events

import (
	"context"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ConfirmEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		DocDb:      true,
		RealtimeDb: true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	confirm(ctx, w, r)
}

func confirm(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	overlayKey := r.URL.Query().Get("key")
	if overlayKey == "" {
		log.Printf("Missing overlay key")
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	eventId := r.URL.Query().Get("event")
	if eventId == "" {
		log.Printf("Missing event id")
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		log.Printf("Missing realtime db")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	db, ok := ctx.GetDocDb()
	if !ok {
		log.Printf("Missing doc db")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	var uid string
	if err := rtdb.NewRef("Users").Child("Overlay").Child(overlayKey).Get(ctx.Ctx, &uid); err != nil {
		log.Printf("Failed to get user id | %v", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	if uid == "" {
		log.Printf("User not found for key %s", overlayKey)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}
	log.Printf("Key %s belongs to user %s", overlayKey, uid)

	q := db.Collection("Events").Where("ID", "==", eventId).Where("Uid", "==", uid).Limit(1)
	if err := db.RunTransaction(ctx.Ctx, func(ctx context.Context, tx modules.Transaction) error {
		snap, err := tx.Documents(q).Next()
		if err != nil {
			return err
		}

		event := &pb.Event{}
		if err := snap.DataTo(event); err != nil {
			return err
		}

		if event.WasShown {
			return nil
		}

		event.WasShown = true
		return tx.Set(snap.Ref(), event, modules.DbOpts{})
	}); err != nil {
		log.Printf("Failed to confirm event | %v", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	log.Printf("Event %s confirmed", eventId)

	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Printf("Failed to write response | %v", err)
		return
	}
}
