package functions

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

const listenersRegex = ".*%s.*"
const noMoreItemsErr = "no more items in iterator"

func getListeners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	uid := r.URL.Query().Get("uid")
	if uid == "" {
		returnError(w, r, http.StatusBadRequest, "Missing uid")
		return
	}

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		returnError(w, r, http.StatusInternalServerError, "Failed to create client")
		return
	}

	subscriptions := client.Subscriptions(ctx)
	listenerCount := 0
	for {
		sub, err := subscriptions.Next()
		if err == nil {
			matched, err := regexp.Match(fmt.Sprintf(listenersRegex, uid), []byte(sub.ID()))
			if err != nil {
				returnError(w, r, http.StatusInternalServerError, "Failed to match")
				return
			}
			if matched {
				listenerCount++
			}
		} else if err.Error() != noMoreItemsErr {
			returnError(w, r, http.StatusInternalServerError, "Failed to iterate")
			return
		} else {
			break
		}
	}
	log.Infof("Found %d listeners for %s", listenerCount, uid)

	_, err = w.Write([]byte(fmt.Sprintf("%d", listenerCount)))
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}
