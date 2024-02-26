package tips

import (
	"context"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	"net/http"
)

func ConfirmEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.CreateContext(r.Context(), &providers.Config{
		DocDb:  true,
		PubSub: true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	r = r.WithContext(ctx)

	confirm(w, r)
}

func handleConfirmation(ctx context.Context, tipId string) error {
	db, ok := providers.GetDocDb(ctx)
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

	_, err = docRef.Set(ctx, &tip, modules.DbOpts{})
	if err != nil {
		return err
	}

	return nil
}

func confirm(w http.ResponseWriter, r *http.Request) {
	tipId := r.URL.Query().Get("id")
	if tipId == "" {
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	if err := handleConfirmation(r.Context(), tipId); err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
}
