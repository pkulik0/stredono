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

func ReturnError(w http.ResponseWriter, code int, errorMessage string) {
	log.Errorf("Error: %s", errorMessage)

	errRes := pb.ErrorResponse{
		Error: errorMessage,
	}
	data, err := proto.Marshal(&errRes)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)

	return
}

func SendDonate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		ReturnError(w, http.StatusBadRequest, "Failed to read request")
		return
	}

	req := pb.SendDonateRequest{}
	if err := proto.Unmarshal(body, &req); err != nil {
		ReturnError(w, http.StatusBadRequest, "Failed to unmarshal request")
		return
	}

	redirectUrl := "http://google.com"
	sdRes := pb.SendDonateResponse{
		RedirectUrl: redirectUrl,
	}

	data, err := proto.Marshal(&sdRes)
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, "Failed to marshal response")
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write(data)
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, "Failed to write response")
		return
	}
	log.Infof("Redirecting to: %s", redirectUrl)
}

func ConfirmPayment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "stredono-5ccdd")
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, "Failed to connect to pubsub")
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
		Amount:    205.58,
		Currency:  "EUR",
		Message:   "aha aha2 aha3 aha4 aha5 aha6 aha7 aha8 aha9 aha10 aha11 aha12 aha13 aha14 aha15 aha16 aha17 aha18 aha19 aha21 aha22 aha23 aha24 aha25 aha26 aha27 aha28 aha29 aha30 aha31 aha32 aha33 aha34 aha35 aha36 aha37 aha38 aha39 aha40 aha41 oeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuaeouaoeueoauoeauaoeuae",
	}
	data, err := proto.Marshal(&sdReq)
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, "Failed to marshal message")
		return
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(ctx)
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, "Failed to publish message")
		return
	}

	fmt.Fprintf(w, "Published message with ID: %s", id)
}
