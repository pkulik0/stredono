package functions

import (
	"github.com/nicklaw5/helix"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const webhookUrl = "http://localhost:8080/twitchWebhook"

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	CorsMiddleware(CloudMiddleware(CloudConfig{
		Firestore: true,
		Auth: AuthConfig{
			Client: true,
			Token:  true,
		},
	}, HelixMiddleware(createSubscription)))(w, r)
}

func createSubscription(w http.ResponseWriter, r *http.Request) {
	_, ok := GetAuthToken(r.Context())
	if !ok {
		log.Error("Failed to get auth token")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	client, ok := GetHelixClient(r.Context())
	if !ok {
		log.Error("Failed to get twitch client")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
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
		log.Errorf("Failed to create subscription: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

}
