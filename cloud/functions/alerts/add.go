package alerts

import (
	"context"
	"errors"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strings"
)

func AddEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.CreateContext(r.Context(), &providers.Config{
		DocDb: true,
		Auth:  true,
	})
	if err != nil {
		log.Error(err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	r = r.WithContext(ctx)

	add(w, r)
}

func validateAlert(style *pb.Alert) error {
	return nil
}

func add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	docDb, ok := providers.GetDocDb(ctx)
	if !ok {
		log.Error("failed to get docDb client")
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	token, _, ok := providers.GetAuthToken(ctx, r)
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

	alertsDoc := docDb.Collection("alerts").Doc(token.UserId())
	err = docDb.RunTransaction(ctx, func(ctx context.Context, tx modules.Transaction) error {
		usersAlerts := &pb.UsersAlerts{}
		log.Info("getting alerts doc: ", alertsDoc)
		snap, err := tx.Get(alertsDoc)
		if err != nil {
			if !strings.Contains(err.Error(), "NotFound") {
				return err
			}
			usersAlerts.Alerts = make([]*pb.Alert, 0)
		} else {
			if err := snap.DataTo(&usersAlerts); err != nil {
				return err
			}
		}

		if err := validateAlert(alert); err != nil {
			return errors.New("invalid alert: " + err.Error())
		}

		log.Info("adding alert: ", alert)
		usersAlerts.Alerts = append(usersAlerts.Alerts, alert)
		if err = tx.Set(alertsDoc, usersAlerts, modules.DbOpts{}); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "invalid alert") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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
