package functions

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkulik0/stredono/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"strings"
	"time"
)

func validateNewDonate(req *pb.SendDonateRequest) error {
	if req.Id != "" {
		return fmt.Errorf("id should not be set")
	}
	if req.Sender == "" {
		return fmt.Errorf("invalid sender: %s", req.Sender)
	}
	if req.RecipientId == "" {
		return fmt.Errorf("invalid recipient id: %s", req.RecipientId)
	}
	if req.Status != pb.DonateStatus_INITIATED {
		return fmt.Errorf("invalid status: %s", req.Status)
	}
	if req.Amount < 0 {
		return fmt.Errorf("invalid amount: %d", req.Amount)
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

func Send(w http.ResponseWriter, r *http.Request) {
	CorsMiddleware(CloudMiddleware(CloudConfig{
		Firestore: true,
	}, send))(w, r)
}

func send(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Failed to read request: %s", err)
		http.Error(w, "Failed to read request", http.StatusBadRequest)
		return
	}

	req := &pb.SendDonateRequest{}
	if err := proto.Unmarshal(body, req); err != nil {
		log.Errorf("Failed to unmarshal request: %s", err)
		http.Error(w, "Failed to unmarshal request", http.StatusBadRequest)
		return
	}

	if err := validateNewDonate(req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	req.Id = strings.ReplaceAll(uuid.New().String(), "-", "")
	req.Status = pb.DonateStatus_PAYMENT_PENDING
	req.Timestamp = time.Now().Unix()

	ctx := r.Context()
	firestoreClient, ok := GetFirestore(ctx)
	if !ok {
		log.Error("Failed to get firestore client")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	doc, _, err := firestoreClient.Collection("donations").Add(ctx, req) // TODO: check
	if err != nil {
		log.Errorf("Failed to save donation: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	donateId := doc.ID
	log.Infof("Saved donation with id: %s", donateId)

	redirectUrl := "http://google.com"
	sdRes := pb.SendDonateResponse{
		RedirectUrl: redirectUrl,
	}

	data, err := proto.Marshal(&sdRes)
	if err != nil {
		log.Errorf("Failed to marshal response: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write(data)
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}
