package functions

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func onRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request", http.StatusBadRequest)
		return
	}

	log.Infof("Received request: %v", r)
	log.Infof("Body: %s", body)
	_, err = fmt.Fprintf(w, "%+v", body)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	return
}
