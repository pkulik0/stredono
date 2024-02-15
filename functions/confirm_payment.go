package functions

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkulik0/stredono/pb"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	donationsPubsubTopic = "donations"
)

func ConfirmPayment(w http.ResponseWriter, r *http.Request) {
	CloudMiddleware(CloudConfig{
		Firestore: true,
		Pubsub:    true,
	}, confirmPayment)(w, r)
}

func confirmPayment(w http.ResponseWriter, r *http.Request) {
	donationId := r.URL.Query().Get("id")
	if donationId == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	firestoreClient, ok := GetFirestore(r.Context())
	if !ok {
		log.Error("Failed to get firestore client")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	docRef := firestoreClient.Collection("donations").Doc(donationId)
	doc, err := docRef.Get(r.Context())
	if err != nil {
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}

	donateReq := pb.SendDonateRequest{}
	if err := doc.DataTo(&donateReq); err != nil {
		log.Errorf("Failed to unmarshal data: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	if donateReq.Status != pb.DonateStatus_PAYMENT_PENDING {
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}
	donateReq.Status = pb.DonateStatus_PAYMENT_SUCCESS

	_, err = docRef.Set(r.Context(), &donateReq)
	if err != nil {
		log.Errorf("Failed to update donation: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	pubsubClient, ok := GetPubsub(r.Context())
	topic := pubsubClient.Topic(donationsPubsubTopic)
	topic.PublishSettings.NumGoroutines = 1
	defer topic.Stop()

	data, err := proto.Marshal(&donateReq)
	if err != nil {
		log.Errorf("Failed to marshal data: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	result := topic.Publish(r.Context(), &pubsub.Message{
		Data: data,
		Attributes: map[string]string{
			"recipientId": donateReq.RecipientId,
		},
	})
	id, err := result.Get(r.Context())
	if err != nil {
		log.Errorf("Failed to publish message: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	_, _ = fmt.Fprintf(w, "Published message with ID: %s", id)
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
}
