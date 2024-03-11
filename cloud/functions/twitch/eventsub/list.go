package eventsub

import (
	"encoding/json"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ListEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		SecretManager: true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	list(ctx, w, r)
}

func list(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	helixClient, err := providers.GetHelixAppClient(ctx)
	if err != nil {
		log.Printf("Failed to get Helix client | %v", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	subs := make([]helix.EventSubSubscription, 0)
	pageCursor := ""
	for {
		resp, err := helixClient.GetEventSubSubscriptions(&helix.EventSubSubscriptionsParams{
			After: pageCursor,
		})
		if err != nil {
			log.Printf("Failed to get eventsub subscriptions | %v", err)
			http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
			return
		}

		subs = append(subs, resp.Data.EventSubSubscriptions...)
		if resp.Data.Pagination.Cursor == "" {
			break
		}
		pageCursor = resp.Data.Pagination.Cursor
	}

	data, err := json.Marshal(subs)
	if err != nil {
		log.Printf("Failed to marshal subscriptions | %v", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to write response | %v", err)
		return
	}
}
