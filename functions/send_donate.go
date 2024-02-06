package functions

import (
	"context"
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

func saveDonateToFirestore(donateReq *pb.SendDonateRequest) (string, error) {
	ctx := context.Background()
	client, err := getFirestoreClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	doc, _, err := client.Collection("donations").Add(ctx, donateReq)
	if err != nil {
		return "", err
	}
	return doc.ID, nil
}

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

func sendDonate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		returnError(w, r, http.StatusBadRequest, "Failed to read request")
		return
	}

	req := pb.SendDonateRequest{}
	if err := proto.Unmarshal(body, &req); err != nil {
		returnError(w, r, http.StatusBadRequest, "Failed to parse request")
		return
	}

	if err := validateNewDonate(&req); err != nil {
		returnError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	req.Id = strings.ReplaceAll(uuid.New().String(), "-", "")
	req.Status = pb.DonateStatus_PAYMENT_PENDING
	req.Timestamp = time.Now().Unix()

	donateId, err := saveDonateToFirestore(&req)
	if err != nil {
		returnError(w, r, http.StatusInternalServerError, "Failed to save donation")
		return
	}
	log.Infof("Saved donation with id: %s", donateId)

	redirectUrl := "http://google.com"
	sdRes := pb.SendDonateResponse{
		RedirectUrl: redirectUrl,
	}

	data, err := proto.Marshal(&sdRes)
	if err != nil {
		returnError(w, r, http.StatusInternalServerError, "Invalid response")
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write(data)
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}
