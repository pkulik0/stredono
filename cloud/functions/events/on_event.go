package events

import (
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/pkulik0/stredono/cloud/functions"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
)

func OnEventEntrypoint(ctx context.Context, e event.Event) error {
	var msg functions.EventMessageData
	if err := e.DataAs(&msg); err != nil {
		log.Printf("Failed to convert data | %v", err)
		return fmt.Errorf("failed to convert data | %v", err)
	}

	eventData := &pb.Event{}
	if err := proto.Unmarshal(msg.Message.Data, eventData); err != nil {
		log.Printf("Failed to unmarshal data | %v", err)
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}

	log.Printf("Received event: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		RealtimeDb:    true,
		Storage:       true,
		TextToSpeech:  true,
		Proxy:         true,
		DocDb:         true,
		SecretManager: true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	if err := onEvent(provCtx, eventData); err != nil {
		log.Printf("Failed to handle event | %v", err)
		return fmt.Errorf("failed to handle event | %v", err)
	}
	return nil
}

func onEvent(ctx *providers.Context, event *pb.Event) error {
	if event.Uid == "" {
		if event.Provider == "" || event.ProviderID == "" {
			return fmt.Errorf("missing provider or provider id")
		}

		uid, err := providers.ProviderIdToUid(ctx, event.Provider, event.ProviderID)
		if err != nil {
			return err
		}
		event.Uid = uid
	}

	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	settings := &pb.EventsSettings{}
	eventsRef := rtdb.NewRef("Data").Child(event.Uid).Child("Settings").Child("Events")
	if err := eventsRef.Get(ctx.Ctx, &settings); err != nil {
		return err
	}

	checkIntValue := func(valueStr string, minValue int32) bool {
		value, err := strconv.ParseInt(valueStr, 10, 32)
		if err != nil {
			return false
		}
		return int32(value) >= minValue
	}

	valueStr := event.Data["Value"]
	text := ""

	switch event.Type {
	case pb.EventType_TIP:
		minValue := settings.Tip.MinAmount
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return err
		}
		if value < minValue {
			return nil
		}

		currency, ok := event.Data["Currency"]
		if !ok {
			return fmt.Errorf("missing currency")
		}

		text = settings.Tip.Template
		text = strings.ReplaceAll(text, "{currency}", currency)
	case pb.EventType_CHEER:
		minValue := settings.Cheer.MinAmount
		if !checkIntValue(valueStr, minValue) {
			return nil
		}
		text = settings.Cheer.Template
	case pb.EventType_SUB:
		minValue := settings.Sub.MinMonths
		if !checkIntValue(valueStr, minValue) {
			return nil
		}
		text = settings.Sub.Template
	case pb.EventType_SUB_GIFT:
		minValue := settings.SubGift.MinCount
		if !checkIntValue(valueStr, minValue) {
			return nil
		}
		text = settings.SubGift.Template
		total, ok := event.Data["Total"]
		if !ok {
			return fmt.Errorf("missing total")
		}
		text = strings.ReplaceAll(text, "{total}", total)
	case pb.EventType_RAID:
		minValue := settings.Raid.MinViewers
		if !checkIntValue(valueStr, minValue) {
			return nil
		}
		text = settings.Raid.Template
	case pb.EventType_FOLLOW:
		if !settings.Follow.IsEnabled {
			return nil
		}
		text = settings.Follow.Template
	case pb.EventType_CHAT_TTS:
		if !settings.ChatTTS.IsEnabled {
			return nil
		}
		text = settings.ChatTTS.Template
	default:
		return fmt.Errorf("unknown event type: %v", event.Type)
	}

	text = strings.ReplaceAll(text, "{value}", valueStr)
	text = strings.ReplaceAll(text, "{user}", event.SenderName)
	text += ". " + event.Data["Message"]

	path, err := providers.GenerateSpeech(ctx, &pb.TTSRequest{
		ID:   event.ID,
		Uid:  event.Uid,
		Text: text,
	})
	if err != nil { // still allow the event to be published but without tts
		log.Errorf("Failed to generate speech | %v", err)
		path = ""
	}
	event.TTSUrl = path

	event.IsApproved = !settings.RequireApproval || event.Type == pb.EventType_CHAT_TTS // only by mods, so it's fine to auto-approve

	db, ok := ctx.GetDocDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	if _, err := db.Collection("events").Add(ctx.Ctx, event); err != nil {
		return err
	}

	return nil
}
