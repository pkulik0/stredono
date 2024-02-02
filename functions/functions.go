package functions

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/golang/protobuf/proto"
	"github.com/pkulik0/stredono/pb"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func init() {
	functions.HTTP("SendDonate", SendDonate)
	functions.HTTP("ConfirmPayment", ConfirmPayment)
}

func SendDonate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request", http.StatusBadRequest)
		return
	}

	req := pb.SendDonateRequest{}
	if err := proto.Unmarshal(body, &req); err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s: %s donated %s %s: %s", req.Recipient, req.Sender, req.Amount, req.Currency, req.Message)
}

func ConfirmPayment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "stredono-5ccdd")
	if err != nil {
		http.Error(w, "Failed to connect to pubsub", http.StatusInternalServerError)
		log.Errorf("Failed to connect to pubsub: %v", err)
		return
	}
	defer client.Close()

	topic := client.Topic("donations")
	topic.PublishSettings.NumGoroutines = 1
	defer topic.Stop()

	sdReq := pb.SendDonateRequest{
		Email:     "abc@email.com",
		Recipient: "stredono",
		Sender:    "pkulik0",
		Amount:    "205.58",
		Currency:  "EUR",
		Message:   "aha aha2 aha3 aha4 aha5 aha6 aha7 aha8 aha9 aha10 aha11 aha12 aha13 aha14 aha15 aha16 aha17 aha18 aha19 aha21 aha22 aha23 aha24 aha25 aha26 aha27 aha28 aha29 aha30 aha31 aha32 aha33 aha34 aha35 aha36 aha37 aha38 aha39 aha40 aha41 oeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuae",
	}
	data, err := proto.Marshal(&sdReq)
	if err != nil {
		http.Error(w, "Failed to marshal message", http.StatusInternalServerError)
		log.Errorf("Failed to marshal message: %v", err)
		return
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(ctx)
	if err != nil {
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		log.Errorf("Failed to publish message: %v", err)
		return
	}

	fmt.Fprintf(w, "Published message with ID: %s", id)
}
