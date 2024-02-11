package functions

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"net/http"
)

const (
	ProjectID   = "stredono-5ccdd"
	DatabaseUrl = "https://stredono.europe-west1.firebasedatabase.app"
)

func init() {
	functions.HTTP("sendDonate", sendDonate)
	functions.HTTP("confirmPayment", confirmPayment)
	functions.HTTP("getListeners", getListeners)
	functions.HTTP("onRegister", onRegister)
	functions.HTTP("connectTwitch", connectTwitch)
	functions.HTTP("connectTwitchCallback", connectTwitchCallback)
}

func setupCors(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", authHeader)
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return false
}
