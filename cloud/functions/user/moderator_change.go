package user

import (
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

func ModeratorChangeEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		Auth:       true,
		RealtimeDb: true,
	})
	if err != nil {
		log.Errorf("failed to create context | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	moderatorChange(ctx, w, r)
}

func moderatorChange(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	log.Infof("received request | %s", r.URL.String())

	uid := r.URL.Query().Get("uid")
	if matched, err := regexp.Match(`^[a-zA-Z0-9]{28}$`, []byte(uid)); err != nil || !matched {
		log.Errorf("received request with invalid uid | %s", uid)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	action := r.URL.Query().Get("action")
	if action != "add" && action != "remove" {
		log.Errorf("received request with invalid action | %s", action)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		log.Errorf("failed to get realtime db")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	token, ok := ctx.GetAuthToken(r)
	if !ok {
		log.Errorf("received request without token")
		http.Error(w, platform.UnauthorizedMessage, http.StatusUnauthorized)
		return
	}

	// TODO: check if the given uid actually exists

	err := rtdb.NewRef("Moderators").Child("From").Child(token.UserId()).Transaction(ctx.Ctx, func(node modules.TransactionNode) (interface{},
		error) {
		var moderators []string
		if err := node.Unmarshal(&moderators); err != nil {
			return nil, err
		}

		switch action {
		case "add":
			for _, m := range moderators {
				if m == uid {
					log.Infof("%s already a moderator of %s", uid, token.UserId())
					return moderators, nil
				}
			}

			log.Infof("added %s as a moderator of %s", uid, token.UserId())
			moderators = append(moderators, uid)
		case "remove":
			for i, m := range moderators {
				if m == uid {
					moderators = append(moderators[:i], moderators[i+1:]...)
					log.Infof("removed %s from moderators of %s", uid, token.UserId())
					break
				}
			}
		}

		return moderators, nil
	})
	if err != nil {
		log.Errorf("failed to add moderator (step 1) | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	err = rtdb.NewRef("Moderators").Child("To").Child(uid).Transaction(ctx.Ctx, func(node modules.TransactionNode) (interface{}, error) {
		var moderatedUsers []string
		if err := node.Unmarshal(&moderatedUsers); err != nil {
			return nil, err
		}

		switch action {
		case "add":
			for _, m := range moderatedUsers {
				if m == token.UserId() {
					log.Infof("step 2 - %s already a moderator of %s", uid, token.UserId())
					return moderatedUsers, nil
				}
			}
			log.Infof("step 2 - added %s as a moderator of %s", uid, token.UserId())
			moderatedUsers = append(moderatedUsers, token.UserId())
		case "remove":
			for i, m := range moderatedUsers {
				if m == token.UserId() {
					moderatedUsers = append(moderatedUsers[:i], moderatedUsers[i+1:]...)
					log.Infof("step 2 - removed %s from moderators of %s", uid, token.UserId())
					break
				}
			}
		}

		return moderatedUsers, nil
	})
	if err != nil {
		log.Errorf("failed to add moderator (step 2) | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Errorf("failed to write response | %s", err)
		return
	}
}
