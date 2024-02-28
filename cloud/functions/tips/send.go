package tips

import (
	"fmt"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"time"
)

func SendEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		DocDb: true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	send(ctx, w, r)
}

func validateTip(tip *pb.Tip) error {
	if tip.SenderId == "" {
		return fmt.Errorf("invalid sender id (%s)", tip.SenderId)
	}
	if tip.DisplayName == "" {
		return fmt.Errorf("invalid display name (%s)", tip.DisplayName)
	}
	if tip.RecipientId == "" {
		return fmt.Errorf("invalid recipient id (%s)", tip.RecipientId)
	}
	if tip.Status != pb.TipStatus_INITIATED {
		return fmt.Errorf("invalid status (%s)", tip.Status)
	}
	if tip.Amount < 0 {
		return fmt.Errorf("invalid amount (%f)", tip.Amount)
	}
	if tip.Currency <= pb.Currency_UNKNOWN || tip.Currency > pb.Currency_PLN { // Bits can only be sent by internal services
		return fmt.Errorf("invalid currency (%s)", tip.Currency)
	}
	if tip.Email == "" {
		return fmt.Errorf("invalid email (%s)", tip.Email)
	}
	if tip.Timestamp != 0 {
		return fmt.Errorf("timestamp should not be set")
	}
	return nil
}

func handleSend(ctx *providers.Context, tip *pb.Tip) (string, error) {
	db, ok := ctx.GetDocDb()
	if !ok {
		return "", platform.ErrorMissingContextValue
	}

	tip.Status = pb.TipStatus_PAYMENT_PENDING
	tip.Timestamp = time.Now().Unix()

	result, err := db.Collection("tips").Add(ctx.Ctx, tip)
	if err != nil {
		return "", err
	}

	return result.Doc.Id(), nil
}

func send(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Failed to read request: %s", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	tip := &pb.Tip{}
	if err := proto.Unmarshal(body, tip); err != nil {
		log.Errorf("Failed to unmarshal request: %s", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	if err := validateTip(tip); err != nil {
		log.Errorf("Invalid tip: %s", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	id, err := handleSend(ctx, tip)
	if err != nil {
		log.Errorf("Failed to handle new tip: %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	log.Infof("New tip added: %s", id) // TODO: use it

	redirectUrl := "http://google.com"
	_, err = w.Write([]byte(redirectUrl))
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}
