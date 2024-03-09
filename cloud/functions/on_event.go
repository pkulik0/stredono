package functions

import (
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/pkulik0/stredono/cloud/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func OnEventEntrypoint(ctx context.Context, e event.Event) error {
	var msg EventMessageData
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("failed to convert data | %v", err)
	}

	eventData := &pb.Event{}
	if err := proto.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}

	log.Printf("Received event: %v", eventData)

	switch eventData.Payload.(type) {
	case *pb.Event_Tip:
		log.Printf("Received tip: %v", eventData.GetTip())
	case *pb.Event_Follow:
		log.Printf("Received follow: %v", eventData.GetFollow())
	case *pb.Event_Cheer:
		log.Printf("Received review: %v", eventData.GetCheer())
	case *pb.Event_Sub:
		log.Printf("Received sub: %v", eventData.GetSub())
	case *pb.Event_SubGift:
		log.Printf("Received sub gift: %v", eventData.GetSubGift())
	case *pb.Event_Raid:
		log.Printf("Received raid: %v", eventData.GetRaid())
	case *pb.Event_ChatMessage:
		log.Printf("Received chat message: %v", eventData.GetChatMessage())
	default:
		return fmt.Errorf("unknown event type")
	}

	return nil
}
