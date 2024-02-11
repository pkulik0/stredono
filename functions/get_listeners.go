package functions

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

const listenersRegex = ".*%s.*"
const noMoreItemsErr = "no more items in iterator"

func getListeners(w http.ResponseWriter, r *http.Request) {
	Middleware(MiddlewareConfig{
		Pubsub: true,
	}, getListenersInternal)(w, r)
}

func getListenersInternal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	uid := r.URL.Query().Get("uid")
	if uid == "" {
		http.Error(w, "Missing uid", http.StatusBadRequest)
		return
	}

	client, ok := GetPubsub(r.Context())
	if !ok {
		log.Errorf("Failed to get pubsub client")
		http.Error(w, internalServerError, http.StatusInternalServerError)
		return
	}

	subscriptions := client.Subscriptions(r.Context())
	listenerCount := 0
	for {
		sub, err := subscriptions.Next()
		if err == nil {
			matched, err := regexp.Match(fmt.Sprintf(listenersRegex, uid), []byte(sub.ID()))
			if err != nil {
				log.Errorf("Failed to match regex: %s", err)
				http.Error(w, internalServerError, http.StatusInternalServerError)
				return
			}
			if matched {
				listenerCount++
			}
		} else if err.Error() != noMoreItemsErr {
			log.Errorf("Failed to get next subscription: %s", err)
			http.Error(w, internalServerError, http.StatusInternalServerError)
			return
		} else {
			break
		}
	}
	log.Infof("Found %d listeners for %s", listenerCount, uid)

	_, err := w.Write([]byte(fmt.Sprintf("%d", listenerCount)))
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}
