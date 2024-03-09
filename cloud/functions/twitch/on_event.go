package twitch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/functions"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"net/http"
)

const EventSubTypeUserAuthorizationGrant = "user.authorization.grant" // not in the helix package

var eventHandlers = map[string]func(ctx context.Context, msg functions.EventMessageData) error{
	"channel.ad_break.begin":                     onChannelAdBreakBegin,
	helix.EventSubTypeChannelUpdate:              onChannelUpdate,
	helix.EventSubTypeChannelFollow:              onChannelFollow,
	helix.EventSubTypeChannelChatMessage:         onChannelChatMessage,
	helix.EventSubTypeChannelSubscription:        onChannelSubscription,
	helix.EventSubTypeChannelSubscriptionMessage: onChannelSubscriptionMessage,
	helix.EventSubTypeChannelSubscriptionGift:    onChannelSubscriptionGift,
	helix.EventSubTypeChannelCheer:               onChannelCheer,
	helix.EventSubTypeChannelRaid:                onChannelRaid,
	helix.EventSubTypeChannelBan:                 onChannelBan,
	helix.EventSubTypeChannelUnban:               onChannelUnban,
	helix.EventSubTypeModeratorRemove:            onChannelModeratorRemove,
	helix.EventSubTypeStreamOnline:               onStreamOnline,
	helix.EventSubTypeStreamOffline:              onStreamOffline,
	helix.EventSubTypeUserUpdate:                 onUserUpdate,
	helix.EventSubTypeUserAuthorizationRevoke:    onUserAuthorizationRevoke,
	EventSubTypeUserAuthorizationGrant:           onUserAuthorizationGrant,
}

var eventsubSubs = map[string]string{
	"channel.ad_break.begin":                     "1",
	helix.EventSubTypeChannelUpdate:              "2",
	helix.EventSubTypeChannelFollow:              "2",
	helix.EventSubTypeChannelChatMessage:         "1",
	helix.EventSubTypeChannelSubscription:        "1",
	helix.EventSubTypeChannelSubscriptionMessage: "1",
	helix.EventSubTypeChannelSubscriptionGift:    "1",
	helix.EventSubTypeChannelCheer:               "1",
	helix.EventSubTypeChannelRaid:                "1",
	helix.EventSubTypeChannelBan:                 "1",
	helix.EventSubTypeChannelUnban:               "1",
	helix.EventSubTypeModeratorRemove:            "1",
	helix.EventSubTypeStreamOnline:               "1",
	helix.EventSubTypeStreamOffline:              "1",
	helix.EventSubTypeUserUpdate:                 "1",
}

func OnEventEntrypoint(ctx context.Context, e event.Event) error {
	var msg functions.EventMessageData
	if err := e.DataAs(&msg); err != nil {
		log.Printf("Failed to read event data | %v", err)
		return fmt.Errorf("failed to read event data | %v", err)
	}

	eventType, ok := msg.Message.Attributes["type"]
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

func onUserUpdate(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubUserUpdateEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received user update: %v", eventData)

	return nil
}

// the helix pkg doesn't have an alias for grant
type eventSubUserAuthorizationGrant = helix.EventSubUserAuthenticationRevokeEvent

func onUserAuthorizationGrant(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &eventSubUserAuthorizationGrant{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("User authorization granted: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
		PubSub:        true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	helixClient, err := providers.GetHelixAppClient(provCtx)
	if err != nil {
		return fmt.Errorf("failed to get helix client | %v", err)
	}

	transport, err := getHelixTransport(provCtx)
	if err != nil {
		return fmt.Errorf("failed to get helix transport | %v", err)
	}
	condition := helix.EventSubCondition{
		BroadcasterUserID:   eventData.UserID,
		ToBroadcasterUserID: eventData.UserID,
		ModeratorUserID:     eventData.UserID,
		UserID:              platform.TwitchUid,
	}

	for subType, subVersion := range eventsubSubs {
		res, err := helixClient.CreateEventSubSubscription(&helix.EventSubSubscription{
			Type:      subType,
			Version:   subVersion,
			Condition: condition,
			Transport: *transport,
		})
		if err != nil {
			return err
		}

		if res.StatusCode != http.StatusAccepted && res.StatusCode != http.StatusConflict {
			return fmt.Errorf("failed to create eventsub subscription %s | %v", subType, res)
		}
	}

	log.Printf("Created eventsub subscriptions for user %s", eventData.UserID)
	return nil
}

func onUserAuthorizationRevoke(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubUserAuthenticationRevokeEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("User authorization revoked: %v", eventData)

	return nil
}

func onChannelBan(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelBanEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received ban: %v", eventData)

	return nil
}

func onChannelUnban(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelUnbanEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received unban: %v", eventData)

	return nil
}

func onChannelChatMessage(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelChatMessageEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received chat message: %+v", eventData)

	if eventData.ChatterUserID == platform.TwitchUid {
		return nil
	}

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	helixClient, err := providers.GetHelixAppClient(provCtx)
	if err != nil {
		return fmt.Errorf("failed to get helix client | %v", err)
	}

	res, err := helixClient.SendChatMessage(&helix.SendChatMessageParams{
		SenderID:             platform.TwitchUid,
		Message:              eventData.Message.Text,
		ReplyParentMessageID: eventData.MessageID,
		BroadcasterID:        eventData.BroadcasterUserID,
	})
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send chat message | %v", res)
	}

	return nil
}

func publishEvent(ctx *providers.Context, msg functions.EventMessageData, event *pb.Event) (string, error) {
	event.Id = msg.Message.MessageId

	pubsubClient, ok := ctx.GetPubSub()
	if !ok {
		return "", platform.ErrorMissingContextValue
	}

	topic := pubsubClient.Topic("events")
	defer topic.Close()

	data, err := proto.Marshal(event)
	if err != nil {
		return "", fmt.Errorf("failed to marshal event | %v", err)
	}

	msgId, err := topic.Publish(ctx.Ctx, &modules.PubSubMessage{
		Data:       data,
		Attributes: msg.Message.Attributes,
	})
	if err != nil {
		return "", fmt.Errorf("failed to publish message | %v", err)
	}

	return msgId, nil
}

func onChannelCheer(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelCheerEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received cheer: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
		PubSub:        true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	msgId, err := publishEvent(provCtx, msg, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_Cheer{
			Cheer: &pb.Event_CheerPayload{
				Message: eventData.Message,
				Amount:  int32(eventData.Bits),
			},
		},
	})
	if err != nil {
		return err
	}

	log.Printf("Published cheer event with msg id %s", msgId)
	return nil
}

func onChannelFollow(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelFollowEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received follow: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
		PubSub:        true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	msgId, err := publishEvent(provCtx, msg, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_Follow{
			Follow: &pb.Event_FollowPayload{},
		},
	})
	if err != nil {
		return err
	}

	log.Printf("Published follow event with msg id %s", msgId)
	return nil
}

func onChannelModeratorRemove(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubModeratorRemoveEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received channel moderator remove: %v", eventData)

	if eventData.UserID != platform.TwitchUid {
		return nil
	}

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
		PubSub:        true,
		KeyManager:    true,
		DocDb:         true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	keyManager, ok := provCtx.GetKeyManager()
	if !ok {
		return platform.ErrorMissingContextValue
	}
	db, ok := provCtx.GetDocDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	tokenEntry, err := db.Collection("twitch-tokens").Doc(eventData.BroadcasterUserID).Get(provCtx.Ctx)
	if err != nil {
		return err
	}
	token := &pb.TokenEntry{}
	if err := tokenEntry.DataTo(token); err != nil {
		return err
	}
	decryptedToken, err := keyManager.Decrypt(provCtx.Ctx, platform.EncryptionKey, token.EncryptedToken)
	if err != nil {
		return err
	}
	userToken := &pb.Token{}
	if err := proto.Unmarshal(decryptedToken, userToken); err != nil {
		return err
	}

	helixClient, err := providers.GetHelixClient(provCtx) // TODO: change client
	if err != nil {
		return fmt.Errorf("failed to get helix client | %v", err)
	}
	helixClient.SetUserAccessToken(userToken.AccessToken)
	helixClient.SetRefreshToken(userToken.RefreshToken)

	res, err := helixClient.AddChannelModerator(&helix.AddChannelModeratorParams{
		BroadcasterID: eventData.BroadcasterUserID,
		UserID:        platform.TwitchUid,
	})
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to remod user | %v", res)
	}

	log.Printf("add mod response: %v", res)

	log.Printf("Remodded stredono bot (%s) in channel %s (%s)", platform.TwitchUid, eventData.BroadcasterUserName,
		eventData.BroadcasterUserID)

	return nil
}

func onChannelRaid(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelRaidEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received raid: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
		PubSub:        true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	msgId, err := publishEvent(provCtx, msg, &pb.Event{
		Channel:  eventData.ToBroadcasterUserID,
		Username: eventData.FromBroadcasterUserName,
		Payload: &pb.Event_Raid{
			Raid: &pb.Event_RaidPayload{
				Viewers: int32(eventData.Viewers),
			},
		},
	})
	if err != nil {
		return err
	}

	log.Printf("Published raid event with msg id %s", msgId)
	return nil
}

func subTierToEnum(tier string) pb.SubTier {
	switch tier {
	case "1000":
		return pb.SubTier_TIER_1
	case "2000":
		return pb.SubTier_TIER_2
	case "3000":
		return pb.SubTier_TIER_3
	default:
		return pb.SubTier_TIER_UNKNOWN
	}
}

func onChannelSubscription(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelSubscribeEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received sub: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
		PubSub:        true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	msgId, err := publishEvent(provCtx, msg, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_Sub{
			Sub: &pb.Event_SubPayload{
				Message: "",
				Tier:    subTierToEnum(eventData.Tier),
			},
		},
	})
	if err != nil {
		return err
	}

	log.Printf("Published sub event with msg id %s", msgId)
	return nil
}

func onChannelSubscriptionGift(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelSubscriptionGiftEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received sub gift: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
		PubSub:        true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	msgId, err := publishEvent(provCtx, msg, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_SubGift{
			SubGift: &pb.Event_SubGiftPayload{
				Tier:  subTierToEnum(eventData.Tier),
				Count: int32(eventData.Total),
				Total: int32(eventData.CumulativeTotal),
			},
		},
	})
	if err != nil {
		return err
	}

	log.Printf("Published sub gift event with msg id %s", msgId)
	return nil
}

func onChannelSubscriptionMessage(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelSubscriptionMessageEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received sub message: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
		PubSub:        true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	msgId, err := publishEvent(provCtx, msg, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_Sub{
			Sub: &pb.Event_SubPayload{
				Message: eventData.Message.Text,
				Tier:    subTierToEnum(eventData.Tier),
			},
		},
	})
	if err != nil {
		return err
	}

	log.Printf("Published sub message event with msg id %s", msgId)
	return nil
}

func onChannelUpdate(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubChannelUpdateEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received channel update: %v", eventData)

	return nil
}

func onStreamOnline(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubStreamOnlineEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received stream online: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	helixClient, err := providers.GetHelixBotClient(provCtx)
	if err != nil {
		log.Printf("Failed to get helix client | %v", err)
		return fmt.Errorf("failed to get helix client | %v", err)
	}

	res, err := helixClient.SendChatAnnouncement(&helix.SendChatAnnouncementParams{
		BroadcasterID: eventData.BroadcasterUserID,
		ModeratorID:   platform.TwitchUid,
		Message:       "Stream is live!",
	})
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send chat announcement | %v", res)
	}

	log.Printf("Sent chat announcement about stream going live to %s (%s)", eventData.BroadcasterUserName, eventData.BroadcasterUserID)
	return nil
}

func onStreamOffline(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &helix.EventSubStreamOfflineEvent{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received stream offline: %v", eventData)

	provCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	helixClient, err := providers.GetHelixBotClient(provCtx)
	if err != nil {
		log.Printf("Failed to get helix client | %v", err)
		return fmt.Errorf("failed to get helix client | %v", err)
	}

	res, err := helixClient.SendChatAnnouncement(&helix.SendChatAnnouncementParams{
		BroadcasterID: eventData.BroadcasterUserID,
		ModeratorID:   platform.TwitchUid,
		Message:       "See you next time!",
	})
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send chat announcement | %v", res)
	}

	log.Printf("Sent chat announcement about stream going offline to %s (%s)", eventData.BroadcasterUserName, eventData.BroadcasterUserID)
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

func onChannelAdBreakBegin(ctx context.Context, msg functions.EventMessageData) error {
	eventData := &eventAdBreakBegin{}
	if err := json.Unmarshal(msg.Message.Data, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received ad break begin: %v", eventData)

	return nil
}
