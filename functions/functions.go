package functions

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/golang/protobuf/proto"
	"github.com/pkulik0/stredono/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"net/http"
)

func init() {
	functions.HTTP("SendDonate", SendDonate)
	functions.HTTP("ConfirmPayment", ConfirmPayment)
	functions.HTTP("OnRegister", OnRegister)
}

func OnRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request", http.StatusBadRequest)
		return
	}

	log.Infof("Received request: %v", r)
	log.Infof("Body: %s", body)
	_, err = fmt.Fprintf(w, "%+v", body)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	return
}

func ReturnError(w http.ResponseWriter, r *http.Request, code int, errorMessage string) {
	log.Errorf("Error (%d): %s \nRequest: %+v", code, errorMessage, r)
	http.Error(w, errorMessage, code)
}

const projectID = "stredono-5ccdd"

func GetFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}
	return app.Firestore(ctx)
}

func SaveDonateToDb(donateReq *pb.SendDonateRequest) (string, error) {
	ctx := context.Background()
	client, err := GetFirestoreClient(ctx)
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

func ValidateNewDonate(req *pb.SendDonateRequest) error {
	if req.Sender == "" {
		return fmt.Errorf("invalid sender: %s", req.Sender)
	}
	if req.Recipient == "" {
		return fmt.Errorf("invalid recipient: %s", req.Recipient)
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
	if req.Created != nil {
		return fmt.Errorf("timestamp should not be set")
	}
	return nil
}

func SendDonate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		ReturnError(w, r, http.StatusBadRequest, "Failed to read request")
		return
	}

	req := pb.SendDonateRequest{}
	if err := proto.Unmarshal(body, &req); err != nil {
		ReturnError(w, r, http.StatusBadRequest, "Failed to parse request")
		return
	}

	if err := ValidateNewDonate(&req); err != nil {
		ReturnError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	req.Status = pb.DonateStatus_PAYMENT_PENDING
	req.Created = timestamppb.Now()

	donateId, err := SaveDonateToDb(&req)
	if err != nil {
		ReturnError(w, r, http.StatusInternalServerError, "Failed to save donation")
		return
	}
	log.Infof("Saved donation with id: %s", donateId)

	redirectUrl := "http://google.com"
	sdRes := pb.SendDonateResponse{
		RedirectUrl: redirectUrl,
	}

	data, err := proto.Marshal(&sdRes)
	if err != nil {
		ReturnError(w, r, http.StatusInternalServerError, "Invalid response")
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write(data)
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}

func ConfirmPayment(w http.ResponseWriter, r *http.Request) {
	donationId := r.URL.Query().Get("id")
	if donationId == "" {
		ReturnError(w, r, http.StatusBadRequest, "Missing id")
		return
	}

	ctx := context.Background()
	firestoreClient, err := GetFirestoreClient(ctx)
	if err != nil {
		ReturnError(w, r, http.StatusInternalServerError, "Firestore connection failed")
		return
	}
	defer firestoreClient.Close()

	docRef := firestoreClient.Collection("donations").Doc(donationId)
	doc, err := docRef.Get(ctx)
	if err != nil {
		ReturnError(w, r, http.StatusInternalServerError, "Failed to get donation")
		return
	}

	donateReq := pb.SendDonateRequest{}
	if err := doc.DataTo(&donateReq); err != nil {
		ReturnError(w, r, http.StatusInternalServerError, "Failed to parse donation")
		return
	}
	if donateReq.Status != pb.DonateStatus_PAYMENT_PENDING {
		ReturnError(w, r, http.StatusBadRequest, "Invalid status")
		return
	}
	donateReq.Status = pb.DonateStatus_PAYMENT_SUCCESS

	_, err = docRef.Set(ctx, &donateReq)
	if err != nil {
		ReturnError(w, r, http.StatusInternalServerError, "Failed to update donation")
		return
	}

	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		ReturnError(w, r, http.StatusInternalServerError, "PubSub connection failed")
		return
	}
	defer pubsubClient.Close()

	topic := pubsubClient.Topic("donations")
	topic.PublishSettings.NumGoroutines = 1
	defer topic.Stop()

	data, err := proto.Marshal(&donateReq)
	if err != nil {
		ReturnError(w, r, http.StatusInternalServerError, "Failed to marshal message")
		return
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(ctx)
	if err != nil {
		ReturnError(w, r, http.StatusInternalServerError, "Failed to publish message")
		return
	}

	fmt.Fprintf(w, "Published message with ID: %s", id)
}
