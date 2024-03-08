package twitch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/functions"
	log "github.com/sirupsen/logrus"
)

var eventHandlers = map[string]func(ctx context.Context, msg functions.MessagePublishedData) error{
	"channel.ad_break.begin":       onChannelAdBreakBegin,
	"channel.update":               onChannelUpdate,
	"channel.chat.message":         onChannelChatMessage,
	"channel.cheer":                onChannelCheer,
	"channel.follow":               onChannelFollow,
	"channel.raid":                 onChannelRaid,
	"channel.subscribe":            onChannelSubscription,
	"channel.subscription.gift":    onChannelSubscriptionGift,
	"channel.subscription.message": onChannelSubscriptionMessage,
	"channel.moderator.add":        onChannelModeratorAdd,
	"channel.moderator.remove":     onChannelModeratorRemove,
	"channel.ban":                  onChannelBan,
	"channel.unban":                onChannelUnban,
	"stream.online":                onStreamOnline,
	"stream.offline":               onStreamOffline,
	"user.authorization.revoke":    onUserAuthorizationRevoke,
	"user.update":                  onUserUpdate,
}

func OnEventEntrypoint(ctx context.Context, e event.Event) error {
	var msg functions.MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		log.Printf("Failed to read event data | %v", err)
		return fmt.Errorf("failed to read event data | %v", err)
	}

	eventType, ok := msg.Message.Attributes["eventType"]
	if !ok {
		log.Printf("Missing eventType in message attributes")
		return fmt.Errorf("missing eventType in message attributes")
	}

	handler, ok := eventHandlers[eventType]
	if !ok {
		log.Printf("No handler for event type %s", eventType)
		return fmt.Errorf("no handler for event type %s", eventType)
	}

	err := handler(ctx, msg)
	if err != nil {
		log.Printf("Failed to handle event | %v", err)
	}
	return err
}

func onUserUpdate(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubUserUpdateEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received user update: %v", eventData)

	return nil
}

func onUserAuthorizationRevoke(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubUserAuthenticationRevokeEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("User authorization revoked: %v", eventData)

	return nil
}

func onChannelBan(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelBanEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received ban: %v", eventData)

	return nil
}

func onChannelUnban(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelUnbanEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received unban: %v", eventData)

	return nil
}

func onChannelChatMessage(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelChatMessageEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received chat message: %v", eventData)

	return nil
}

func onChannelCheer(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelCheerEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received cheer: %v", eventData)

	return nil
}

func onChannelFollow(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelFollowEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received follow: %v", eventData)

	return nil
}

func onChannelModeratorAdd(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubModeratorAddEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received channel moderator add: %v", eventData)

	return nil
}

func onChannelModeratorRemove(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubModeratorRemoveEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received channel moderator remove: %v", eventData)

	return nil
}

func onChannelRaid(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelRaidEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received raid: %v", eventData)

	return nil
}

func onChannelSubscription(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelSubscribeEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received sub: %v", eventData)

	return nil
}

func onChannelSubscriptionGift(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelSubscriptionGiftEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received sub gift: %v", eventData)

	return nil
}

func onChannelSubscriptionMessage(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelSubscriptionMessageEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received sub message: %v", eventData)

	return nil
}

func onChannelUpdate(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubChannelUpdateEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received channel update: %v", eventData)

	return nil
}

func onStreamOnline(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubStreamOnlineEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received stream online: %v", eventData)

	return nil
}

func onStreamOffline(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &helix.EventSubStreamOfflineEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received stream offline: %v", eventData)

	return nil
}

// channel.ad_break.begin isn't in the helix package
type eventAdBreakBegin struct {
	BroadcasterId    string `json:"broadcaster_user_id"`
	BroadcasterLogin string `json:"broadcaster_user_login"`
	BroadcasterName  string `json:"broadcaster_user_name"`
	RequesterId      string `json:"requester_id"`
	RequesterLogin   string `json:"requester_login"`
	RequesterName    string `json:"requester_name"`
	Duration         int    `json:"duration"`
	IsAutomatic      bool   `json:"is_automatic"`
	StartedAt        string `json:"started_at"`
}

func onChannelAdBreakBegin(ctx context.Context, msg functions.MessagePublishedData) error {
	eventData := &eventAdBreakBegin{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received ad break begin: %v", eventData)

	return nil
}
