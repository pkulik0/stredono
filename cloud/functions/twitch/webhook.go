package twitch

// https://dev.twitch.tv/docs/eventsub/

import (
	"encoding/json"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	eventsubMessageHeaderType      = "Twitch-Eventsub-Message-Type"
	eventsubMessageHeaderId        = "Twitch-Eventsub-Message-Id"
	eventsubMessageHeaderTimestamp = "Twitch-Eventsub-Message-Timestamp"

	eventsubMessageTypeWebhookCallback = "webhook_callback_verification"
	eventsubMessageTypeNotification    = "notification"
	eventsubMessageTypeRevocation      = "revocation"
)

func WebhookEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		RealtimeDb:    true,
		SecretManager: true,
		PubSub:        true,
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

	isVerified := helix.VerifyEventSubNotification(string(eventsubSecret), r.Header, string(body))
	if !isVerified {
		log.Errorf("Failed to verify signature")
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	notification := &eventsubNotification{}
	if err := json.Unmarshal(body, &notification); err != nil {
		log.Errorf("Failed to unmarshal body | %s", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}
	notification.fill(r.Header)

	switch notification.Type {
	case eventsubMessageTypeWebhookCallback:
		_, err = w.Write([]byte(notification.Challenge))
		if err != nil {
			log.Errorf("Failed to write response | %s", err)
			return
		}
		log.Infof("Webhook verification successful: %v", notification.Subscription)
	case eventsubMessageTypeNotification:
		err = handleEvent(ctx, notification)
		if err != nil {
			log.Errorf("Failed to handle event | %s \n %v", err, notification)
			http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
			return
		}
	case eventsubMessageTypeRevocation:
		log.Infof("Revoked subscription: %v", notification.Subscription)
		if _, err := w.Write([]byte("OK")); err != nil {
			log.Errorf("Failed to write response | %s", err)
		}
	default:
		log.Errorf("Unknown message type: %s", notification.Type)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
	}
}

func handleEvent(ctx *providers.Context, notification *eventsubNotification) error {
	bytes, err := json.Marshal(notification.Event)
	if err != nil {
		return err
	}

	pubsubClient, ok := ctx.GetPubSub()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	topic := pubsubClient.Topic("twitch-eventsub")
	defer topic.Close()

	eventType := notification.Subscription.Type
	publishedId, err := topic.Publish(ctx.Ctx, &modules.PubSubMessage{
		Data: bytes,
		Attributes: map[string]string{
			"twitchUid": notification.Subscription.Condition.BroadcasterUserID,
			"type":      eventType,
			"id":        notification.Id,
			"timestamp": notification.Timestamp,
		},
	})
	if err != nil {
		return err
	}

	log.Infof("Published %s (id %s): %+v", eventType, publishedId, notification.Event)
	return nil
}

type eventsubNotification struct {
	Subscription helix.EventSubSubscription `json:"subscription"`
	Event        map[string]interface{}     `json:"event"`
	Challenge    string                     `json:"challenge"`
	Id           string
	Timestamp    string
	Type         string
}

func (e *eventsubNotification) fill(headers http.Header) {
	e.Id = headers.Get(eventsubMessageHeaderId)
	e.Timestamp = headers.Get(eventsubMessageHeaderTimestamp)
	e.Type = headers.Get(eventsubMessageHeaderType)
}
