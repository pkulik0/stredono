package eventsub

import (
	"fmt"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func createEventsubSubscription(client modules.HelixClient, eventType string, version string, condition helix.EventSubCondition,
	transport *helix.EventSubTransport) error {
	res, err := client.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:      eventType,
		Version:   version,
		Condition: condition,
		Transport: *transport,
	})
	if err != nil {
		log.Printf("Failed to create eventsub subscription | %v", err)
		return err
	}

	if res.StatusCode != http.StatusAccepted && res.StatusCode != http.StatusConflict {
		return fmt.Errorf("failed to create eventsub subscription %v", res)
	}

	log.Printf("Eventsub subscription \"%s\" created", eventType)

	return nil
}

func InitEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		SecretManager: true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	initEventsub(ctx, w, r)
}

func handleInit(ctx *providers.Context) error {
	helixClient, err := providers.GetHelixAppClient(ctx)
	if err != nil {
		return err
	}

	transport, err := providers.GetHelixTransport(ctx)
	if err != nil {
		return err
	}
	condition := helix.EventSubCondition{
		ClientID: platform.TwitchClientId,
	}

	if err = createEventsubSubscription(helixClient, "user.authorization.grant", "1", condition, transport); err != nil {
		return err
	}
	if err = createEventsubSubscription(helixClient, "user.authorization.revoke", "1", condition, transport); err != nil {
		return err
	}

	return nil
}

func initEventsub(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	if err := handleInit(ctx); err != nil {
		log.Printf("Failed to initialize eventsub | %v", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Failed to write response | %v", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
	}
}
