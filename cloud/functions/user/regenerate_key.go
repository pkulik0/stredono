package user

import (
	"github.com/google/uuid"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func RegenerateKeyEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		Auth:       true,
		RealtimeDb: true,
	})
	if err != nil {
		log.Errorf("failed to create context | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	regenerateKey(ctx, w, r)
}

func regenerateKey(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
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

	keyUuid, err := uuid.NewUUID()
	if err != nil {
		log.Errorf("failed to generate uuid | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	newKey := strings.ReplaceAll(keyUuid.String(), "-", "")

	refData := rtdb.NewRef("Data").Child(token.UserId()).Child("OverlayKey")
	err = refData.Transaction(ctx.Ctx, func(node modules.TransactionNode) (interface{}, error) {
		var readKey string
		if err := node.Unmarshal(&readKey); err != nil {
			return nil, err
		}

		keysRef := rtdb.NewRef("Users").Child("Overlay")
		if err := keysRef.Child(readKey).Delete(ctx.Ctx); err != nil {
			return nil, err
		}

		if err := keysRef.Child(newKey).Set(ctx.Ctx, token.UserId()); err != nil {
			return nil, err
		}

		return newKey, nil
	})
	if err != nil {
		log.Errorf("failed to regenerate key | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Errorf("failed to write response | %s", err)
		return
	}
}
