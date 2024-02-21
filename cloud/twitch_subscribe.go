package cloud

import (
	"context"
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/cloud/platform"
	"net/http"
)

const webhookUrl = "http://localhost:8080/twitchWebhook"

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	platform.CorsMiddleware(platform.CloudMiddleware(&platform.CloudConfig{
		DocDb: true,
		Auth: platform.AuthConfig{
			Client:      true,
			VerifyToken: true,
		},
	}, platform.HelixMiddleware(createSubscription)))(w, r)
}

func handleCreateSubscription(ctx context.Context) error {
	_, ok := platform.GetAuthToken(ctx)
	if !ok {
		return platform.ErrorMissingContextValue
	}

	client, ok := platform.GetHelix(ctx)
	if !ok {
		return platform.ErrorMissingContextValue
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
	return err
}

func createSubscription(w http.ResponseWriter, r *http.Request) {
	if err := handleCreateSubscription(r.Context()); err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
}
