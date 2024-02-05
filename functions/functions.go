package functions

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	functions.HTTP("sendDonate", sendDonate)
	functions.HTTP("confirmPayment", confirmPayment)
	functions.HTTP("getListeners", getListeners)
	functions.HTTP("onRegister", onRegister)
}

const projectID = "stredono-5ccdd"

func getFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}
	return app.Firestore(ctx)
}

func returnError(w http.ResponseWriter, r *http.Request, code int, errorMessage string) {
	log.Errorf("Error (%d): %s \nRequest: %+v", code, errorMessage, r)
	http.Error(w, errorMessage, code)
}
