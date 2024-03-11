package chat

import (
	"fmt"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func Send(ctx *providers.Context, msg *pb.BotMessage, provider string) error {
	switch provider {
	case "twitch":
		if err := sendTwitch(ctx, msg); err != nil {
			log.Printf("Failed to handle chat message | %v", err)
			return err
		}
	default:
		log.Printf("Unknown provider: %s", provider)
		return fmt.Errorf("unknown provider: %s", provider)
	}

	return nil
}

func sendTwitch(ctx *providers.Context, msg *pb.BotMessage) error {
	log.Printf("Received bot message: %+v", msg)
	client, err := providers.GetHelixBotClient(ctx)
	if err != nil {
		return err
	}

	switch msg.Data.(type) {
	case *pb.BotMessage_Normal:
		res, err := client.SendChatMessage(&helix.SendChatMessageParams{
			BroadcasterID:        msg.ChatID,
			SenderID:             platform.TwitchUid,
			Message:              msg.Message,
			ReplyParentMessageID: msg.GetNormal().ReplyMessageID,
		})
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to send chat message | %v", res.StatusCode)
		}
	case *pb.BotMessage_Announcement:
		colorValue := int32(msg.GetAnnouncement().Color.Number())
		colorName := pb.BotMessage_AnnouncementData_AnnouncementColor_name[colorValue]

		res, err := client.SendChatAnnouncement(&helix.SendChatAnnouncementParams{
			BroadcasterID: msg.ChatID,
			ModeratorID:   platform.TwitchUid,
			Message:       msg.Message,
			Color:         strings.ToLower(colorName),
		})
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to send chat announcement | %v", res.StatusCode)
		}
	default:
		return fmt.Errorf("unknown message type %T", msg.Data)
	}
	return nil
}
