package tips

import (
	"context"
	"errors"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ConfirmEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		DocDb:  true,
		PubSub: true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	confirm(ctx, w, r)
}

func handleConfirmation(ctx *providers.Context, tipId string) error {
	db, ok := ctx.GetDocDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	docRef := db.Collection("tips").Doc(tipId)
	return db.RunTransaction(ctx.Ctx, func(ctx context.Context, tx modules.Transaction) error {
		doc, err := tx.Get(docRef)
		if err != nil {
			return err
		}

		tip := pb.Tip{}
		if err := doc.DataTo(&tip); err != nil {
			return err
		}

		if tip.Status != pb.TipStatus_PAYMENT_PENDING {
			return platform.ErrorInvalidStatus
		}

		tip.Status = pb.TipStatus_PAYMENT_SUCCESS
		return tx.Set(docRef, &tip, modules.DbOpts{})
	})
}

func confirm(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	tipId := r.URL.Query().Get("id")
	if tipId == "" {
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	if err := handleConfirmation(ctx, tipId); err != nil {
		if errors.Is(err, platform.ErrorInvalidStatus) {
			http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		} else {
			http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		}
		return
	}

	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Errorf("Failed to write response | %s", err)
	}
}
