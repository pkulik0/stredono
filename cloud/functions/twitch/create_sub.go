package twitch

import (
	"context"
	"errors"
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	"net/http"
)

const webhookUrl = "http://localhost:8080/twitchWebhook"

func CreateSubEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.CreateContext(r.Context(), &providers.Config{
		DocDb: true,
		Auth:  true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	ctx, err = providers.CreateHelixContext(ctx, r)
	if err != nil {
		if errors.Is(err, platform.ErrorMissingContextValue) {
			http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
			return
		}
		http.Error(w, platform.UnauthorizedMessage, http.StatusUnauthorized)
		return
	}

	createSub(w, r)
}

func handleCreateSub(ctx context.Context) error {
	client, ok := providers.GetHelix(ctx)
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

func createSub(w http.ResponseWriter, r *http.Request) {
	if err := handleCreateSub(r.Context()); err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
}
