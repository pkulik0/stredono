package twitch

// https://dev.twitch.tv/docs/eventsub/

import (
	"encoding/json"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/functions/twitch/eventsub"
	"github.com/pkulik0/stredono/cloud/functions/twitch/handlers"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	eventsubMessageTypeWebhookCallback = "webhook_callback_verification"
	eventsubMessageTypeNotification    = "notification"
	eventsubMessageTypeRevocation      = "revocation"
)

func WebhookEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		SecretManager: true,
		PubSub:        true,
		KeyManager:    true,
		RealtimeDb:    true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	defer func(ctx *providers.Context) {
		if err := ctx.Close(); err != nil {
			log.Errorf("Failed to close context | %s", err)
		}
	}(ctx)

	webhook(ctx, w, r)
}

func webhook(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Failed to read request | %s", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	secretManager, ok := ctx.GetSecretManager()
	if !ok {
		log.Errorf("Missing secret manager")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	eventsubSecret, err := secretManager.GetSecret(ctx.Ctx, "twitch-eventsub-secret", "latest")
	if err != nil {
		log.Errorf("Failed to get secret | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	if isVerified := helix.VerifyEventSubNotification(string(eventsubSecret), r.Header, string(body)); !isVerified {
		log.Errorf("Failed to verify signature")
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	notification := &eventsub.Notification{}
	if err := json.Unmarshal(body, &notification); err != nil {
		log.Errorf("Failed to unmarshal body | %s", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}
	if err := notification.Fill(r.Header); err != nil {
		log.Errorf("Failed to fill notification | %s", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	switch notification.Type {
	case eventsubMessageTypeWebhookCallback:
		_, err = w.Write([]byte(notification.Challenge))
		if err != nil {
			log.Errorf("Failed to write response | %s", err)
			return
		}
		log.Infof("Webhook verification successful: %v", notification.Subscription)
	case eventsubMessageTypeNotification:
		eventType := notification.Subscription.Type
		handler, ok := handlers.TypeToHandler[eventType]
		if !ok {
			log.Errorf("No handler for event type: %s", eventType)
			http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
			return
		}

		if err := handler(ctx, notification); err != nil {
			log.Errorf("Failed to handle notification | %s", err)
		}
		// Always respond with 200 OK to acknowledge the notification, even if there was an error on our side
		if _, err := w.Write([]byte("OK")); err != nil {
			log.Errorf("Failed to write response | %s", err)
		}
	case eventsubMessageTypeRevocation:
		log.Infof("Revoked subscription: %v", notification.Subscription)
		if _, err := w.Write([]byte("OK")); err != nil {
			log.Errorf("Failed to write response | %s", err)
		}
	default:
		log.Errorf("Unknown notification type received: %s", notification.Type)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
	}
}
