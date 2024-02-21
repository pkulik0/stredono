package cloud

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strings"
	"time"
)

func validateIncomingTip(req *pb.Tip) error {
	if req.Id != "" {
		return fmt.Errorf("id should not be set")
	}
	if req.Sender == "" {
		return fmt.Errorf("invalid sender: %s", req.Sender)
	}
	if req.RecipientId == "" {
		return fmt.Errorf("invalid recipient id: %s", req.RecipientId)
	}
	if req.Status != pb.TipStatus_INITIATED {
		return fmt.Errorf("invalid status: %s", req.Status)
	}
	if req.Amount < 0 {
		return fmt.Errorf("invalid amount: %f", req.Amount)
	}
	if req.Currency == "" {
		return fmt.Errorf("invalid currency: %s", req.Currency)
	}
	if req.Email == "" {
		return fmt.Errorf("invalid email: %s", req.Email)
	}
	if req.Timestamp != 0 {
		return fmt.Errorf("timestamp should not be set")
	}
	return nil
}

func TipSend(w http.ResponseWriter, r *http.Request) {
	platform.CorsMiddleware(platform.CloudMiddleware(&platform.CloudConfig{
		DocDb: true,
	}, tipSend))(w, r)
}

func handleTipSend(ctx context.Context, req *pb.Tip) error {
	db, ok := platform.GetDocDb(ctx)
	if !ok {
		return platform.ErrorMissingContextValue
	}

	req.Id = strings.ReplaceAll(uuid.New().String(), "-", "")
	req.Status = pb.TipStatus_PAYMENT_PENDING
	req.Timestamp = time.Now().Unix()

	_, err := db.Collection("tips").Doc(req.Id).Create(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func tipSend(w http.ResponseWriter, r *http.Request) {
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

	if err := validateIncomingTip(tip); err != nil {
		log.Errorf("Invalid tip: %s", err)
		http.Error(w, platform.BadRequestMessage, http.StatusBadRequest)
		return
	}

	if err := handleTipSend(r.Context(), tip); err != nil {
		log.Errorf("Failed to handle new tip: %s", err)
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	redirectUrl := "http://google.com"
	_, err = w.Write([]byte(redirectUrl))
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}
