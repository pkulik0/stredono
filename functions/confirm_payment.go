package functions

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkulik0/stredono/pb"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func confirmPayment(w http.ResponseWriter, r *http.Request) {
	donationId := r.URL.Query().Get("id")
	if donationId == "" {
		returnError(w, r, http.StatusBadRequest, "Missing id")
		return
	}

	ctx := context.Background()
	firestoreClient, err := getFirestoreClient(ctx)
	if err != nil {
		returnError(w, r, http.StatusInternalServerError, "Firestore connection failed")
		return
	}
	defer firestoreClient.Close()

	docRef := firestoreClient.Collection("donations").Doc(donationId)
	doc, err := docRef.Get(ctx)
	if err != nil {
		returnError(w, r, http.StatusInternalServerError, "Failed to get donation")
		return
	}

	donateReq := pb.SendDonateRequest{}
	if err := doc.DataTo(&donateReq); err != nil {
		returnError(w, r, http.StatusInternalServerError, "Failed to parse donation")
		return
	}
	if donateReq.Status != pb.DonateStatus_PAYMENT_PENDING {
		returnError(w, r, http.StatusBadRequest, "Invalid status")
		return
	}
	donateReq.Status = pb.DonateStatus_PAYMENT_SUCCESS

	_, err = docRef.Set(ctx, &donateReq)
	if err != nil {
		returnError(w, r, http.StatusInternalServerError, "Failed to update donation")
		return
	}

	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		returnError(w, r, http.StatusInternalServerError, "PubSub connection failed")
		return
	}
	defer pubsubClient.Close()

	topic := pubsubClient.Topic("donations")
	topic.PublishSettings.NumGoroutines = 1
	defer topic.Stop()

	data, err := proto.Marshal(&donateReq)
	if err != nil {
		returnError(w, r, http.StatusInternalServerError, "Failed to marshal message")
		return
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: data,
		Attributes: map[string]string{
			"recipientId": donateReq.RecipientId,
		},
	})
	id, err := result.Get(ctx)
	if err != nil {
		returnError(w, r, http.StatusInternalServerError, "Failed to publish message")
		return
	}

	_, _ = fmt.Fprintf(w, "Published message with ID: %s", id)
	if err != nil {
		log.Errorf("Failed to write response: %s", err)
		return
	}
}
