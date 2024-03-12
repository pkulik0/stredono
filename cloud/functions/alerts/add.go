package alerts

import (
	"github.com/google/uuid"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

func AddEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		RealtimeDb: true,
		Auth:       true,
	})
	if err != nil {
		log.Error("failed to create context | ", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	add(ctx, w, r)
}

func validateAlert(alert *pb.Alert) error {
	// TODO: validation
	log.Printf("Validating alert: %+v", alert)
	return nil
}

func add(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		log.Error("missing realtime db")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	token, ok := ctx.GetAuthToken(r)
	if !ok {
		http.Error(w, platform.UnauthorizedMessage, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("failed to read request body | ", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	alert := &pb.Alert{}
	if err := proto.Unmarshal(body, alert); err != nil {
		log.Error("failed to unmarshal request | ", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	if err := validateAlert(alert); err != nil {
		log.Error("failed to validate alert | ", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		log.Error("failed to generate alert id | ", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	alert.ID = id.String()

	err = rtdb.NewRef("Data").Child(token.UserId()).Child("Settings").Child("Alerts").Transaction(ctx.Ctx, func(node modules.TransactionNode) (interface{}, error) {
		var alerts []*pb.Alert
		if err := node.Unmarshal(&alerts); err != nil {
			return nil, err
		}

		return append(alerts, alert), nil
	})
	if err != nil {
		log.Error("failed to run transaction | ", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		log.Error("failed to write response | ", err)
		return
	}
}
