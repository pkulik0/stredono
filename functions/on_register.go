package functions

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func OnRegister(w http.ResponseWriter, r *http.Request) {
	CloudMiddleware(CloudConfig{
		Firestore: true,
	}, onRegister)(w, r)
}

func onRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Failed to read body: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	log.Infof("%s", string(body))

	log.Infof("\nRequest: \n%+v", r)

	w.Write([]byte("OK"))
}
