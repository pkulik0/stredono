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

	if err := handleEvent(provCtx, eventData); err != nil {
		log.Printf("Failed to handle event | %v", err)
		return fmt.Errorf("failed to handle event | %v", err)
	}
	return nil
}

func fillTemplate(event *pb.Event, eventSettings *pb.EventSettings) (string, error) {
	template := eventSettings.MessageTemplate
	template = strings.ReplaceAll(template, "{value}", event.Data["Value"])
	template = strings.ReplaceAll(template, "{user}", event.SenderName)

	switch event.Type {
	case pb.EventType_TIP:
		currency, ok := event.Data["Currency"]
		if !ok {
			return "", fmt.Errorf("missing currency")
		}
		template = strings.ReplaceAll(template, "{currency}", currency)
	case pb.EventType_SUB_GIFT:
		total, ok := event.Data["Total"]
		if !ok {
			return "", fmt.Errorf("missing total")
		}
		template = strings.ReplaceAll(template, "{total}", total)
	case pb.EventType_SUB:
	case pb.EventType_CHEER:
	case pb.EventType_CHAT_TTS:
	case pb.EventType_RAID:
	case pb.EventType_FOLLOW:
	default:
		return "", fmt.Errorf("unknown event type: %v", event.Type)
	}

	return template, nil
}

func handleTTS(ctx *providers.Context, event *pb.Event, eventSettings *pb.EventSettings, ttsSettings *pb.TTSSettings, text,
	uid string) error {
	if !eventSettings.EnableTTS {
		return nil
	}

	minValue := eventSettings.MinimumValue
	if eventSettings.MinimumForTTS != nil {
		minValue = *eventSettings.MinimumForTTS
	}

	valueStr := event.Data["Value"]
	switch event.Type {
	case pb.EventType_TIP:
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return err
		}
		if value < float64(minValue) {
			return nil
		}
	case pb.EventType_CHEER:
		fallthrough
	case pb.EventType_SUB:
		fallthrough
	case pb.EventType_SUB_GIFT:
		fallthrough
	case pb.EventType_RAID:
		value, err := strconv.ParseInt(valueStr, 10, 32)
		if err != nil {
			return err
		}
		if int32(value) < minValue {
			return nil
		}
	case pb.EventType_CHAT_TTS:
	case pb.EventType_FOLLOW: // Do nothing
	default:
		return fmt.Errorf("unknown event type: %v", event.Type)
	}

	path, err := providers.GenerateSpeech(ctx, &pb.TTSRequest{
		ID:       event.ID,
		Uid:      uid,
		Text:     text,
		Settings: ttsSettings,
	})
	if err != nil {
		return err
	}

	event.TTSUrl = path
	return nil
}

func handleEvent(ctx *providers.Context, event *pb.Event) error {
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

	eventsSettings := &pb.EventsSettings{}
	eventsRef := rtdb.NewRef("Data").Child(event.Uid).Child("Settings").Child("Events")
	if err := eventsRef.Get(ctx.Ctx, &eventsSettings); err != nil {
		return err
	}

	eventSettings, ok := eventsSettings.Event[event.Type.String()]
	if !ok {
		return fmt.Errorf("event not found: %s", event.Type.String())
	}

	header, err := fillTemplate(event, eventSettings)
	if err != nil {
		return err
	}

	text := fmt.Sprintf("%s. %s", header, event.Data["Message"])
	if err := handleTTS(ctx, event, eventSettings, eventsSettings.TTS, text, event.Uid); err != nil {
		log.Errorf("Failed to handle TTS | %v", err)
		// Don't return so even if TTS fails, the event is still shown without TTS
		// TODO: report issue
	}

	event.IsApproved = !eventsSettings.RequireApproval || event.Type == pb.EventType_CHAT_TTS

	return addEventToQueue(ctx, event)
}

func addEventToQueue(ctx *providers.Context, event *pb.Event) error {
	db, ok := ctx.GetDocDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	if _, err := db.Collection("Events").Add(ctx.Ctx, event); err != nil {
		return err
	}
	return nil
}
