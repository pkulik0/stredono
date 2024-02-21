package cloud

import (
	"context"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"net/http"
)

func TipConfirm(w http.ResponseWriter, r *http.Request) {
	platform.CloudMiddleware(&platform.CloudConfig{
		DocDb:  true,
		PubSub: true,
	}, tipConfirm)(w, r)
}

func handleTipConfirmation(ctx context.Context, tipId string) error {
	db, ok := platform.GetDocDb(ctx)
	if !ok {
		return platform.ErrorMissingContextValue
	}

	docRef := db.Collection("tips").Doc(tipId)
	doc, err := docRef.Get(ctx)
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

	_, err = docRef.Set(ctx, &tip, &platform.DbOpts{})
	if err != nil {
		return err
	}

	return nil
}

func tipConfirm(w http.ResponseWriter, r *http.Request) {
	tipId := r.URL.Query().Get("id")
	if tipId == "" {
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	if err := handleTipConfirmation(r.Context(), tipId); err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
}
