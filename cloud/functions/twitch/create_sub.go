package twitch

import (
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const webhookUrl = "http://localhost:8080/twitchWebhook"

func CreateSubEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		DocDb: true,
		Auth:  true,
		Helix: true,
	})
	if err != nil {
		log.Errorf("Failed to create context | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	createSub(ctx, w, r)
}

func createSub(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	client, ok := ctx.GetHelix()
	if !ok {
		log.Errorf("Failed to get helix client")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err := client.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    "channel.follow",
		Version: "1",
		Condition: helix.EventSubCondition{
			BroadcasterUserID: "123",
		},
		Transport: helix.EventSubTransport{
			Method:   "webhook",
			Callback: webhookUrl,
			Secret:   "secret",
		},
	})
	if err != nil {
		log.Errorf("Failed to create subscription | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Errorf("Failed to write response | %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
}
