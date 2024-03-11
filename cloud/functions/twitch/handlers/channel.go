package handlers

import (
	"fmt"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/functions"
	"github.com/pkulik0/stredono/cloud/functions/twitch/eventsub"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func onChannelBan(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelBanEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received ban: %v", eventData)

	return nil
}

func onChannelUnban(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelUnbanEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received unban: %v", eventData)

	return nil
}

func onChannelCheer(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelCheerEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received cheer: %v", eventData)

	msgId, err := functions.PublishProto(ctx, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_Cheer{
			Cheer: &pb.Event_CheerPayload{
				Message: eventData.Message,
				Amount:  int32(eventData.Bits),
			},
		},
	}, "events", platform.ProviderTwitch)
	if err != nil {
		return err
	}

	log.Printf("Published cheer event with msg id %s", msgId)
	return nil
}

func onChannelFollow(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelFollowEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received follow: %v", eventData)

	msgId, err := functions.PublishProto(ctx, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_Follow{
			Follow: &pb.Event_FollowPayload{},
		},
	}, "events", platform.ProviderTwitch)
	if err != nil {
		return err
	}

	log.Printf("Published follow event with msg id %s", msgId)
	return nil
}

func onChannelModeratorRemove(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubModeratorRemoveEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received channel moderator remove: %v", eventData)

	if eventData.UserID != platform.TwitchUid {
		return nil
	}

	keyManager, ok := ctx.GetKeyManager()
	if !ok {
		return platform.ErrorMissingContextValue
	}
	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	token := &pb.TokenEntry{}
	if err := rtdb.NewRef("Users").Child(platform.ProviderTwitch).Child(eventData.BroadcasterUserID).Get(ctx.Ctx, token); err != nil {
		return err
	}

	decryptedToken, err := keyManager.Decrypt(ctx.Ctx, platform.EncryptionKey, token.EncryptedToken)
	if err != nil {
		return err
	}
	userToken := &pb.Token{}
	if err := proto.Unmarshal(decryptedToken, userToken); err != nil {
		return err
	}

	helixClient, err := providers.GetHelixClient(ctx) // TODO: change client
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

func onChannelRaid(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelRaidEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received raid: %v", eventData)

	msgId, err := functions.PublishProto(ctx, &pb.Event{
		Channel:  eventData.ToBroadcasterUserID,
		Username: eventData.FromBroadcasterUserName,
		Payload: &pb.Event_Raid{
			Raid: &pb.Event_RaidPayload{
				Viewers: int32(eventData.Viewers),
			},
		},
	}, "events", platform.ProviderTwitch)
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

func onChannelSubscription(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelSubscribeEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received sub: %v", eventData)

	msgId, err := functions.PublishProto(ctx, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_Sub{
			Sub: &pb.Event_SubPayload{
				Message: "",
				Tier:    subTierToEnum(eventData.Tier),
			},
		},
	}, "events", platform.ProviderTwitch)
	if err != nil {
		return err
	}

	log.Printf("Published sub event with msg id %s", msgId)
	return nil
}

func onChannelSubscriptionGift(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelSubscriptionGiftEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received sub gift: %v", eventData)

	msgId, err := functions.PublishProto(ctx, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_SubGift{
			SubGift: &pb.Event_SubGiftPayload{
				Tier:  subTierToEnum(eventData.Tier),
				Count: int32(eventData.Total),
				Total: int32(eventData.CumulativeTotal),
			},
		},
	}, "events", platform.ProviderTwitch)
	if err != nil {
		return err
	}

	log.Printf("Published sub gift event with msg id %s", msgId)
	return nil
}

func onChannelSubscriptionMessage(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelSubscriptionMessageEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received sub message: %v", eventData)

	msgId, err := functions.PublishProto(ctx, &pb.Event{
		Channel:  eventData.BroadcasterUserID,
		Username: eventData.UserName,
		Payload: &pb.Event_Sub{
			Sub: &pb.Event_SubPayload{
				Message: eventData.Message.Text,
				Tier:    subTierToEnum(eventData.Tier),
			},
		},
	}, "events", platform.ProviderTwitch)
	if err != nil {
		return err
	}

	log.Printf("Published sub message event with msg id %s", msgId)
	return nil
}

func onChannelUpdate(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelUpdateEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received channel update: %v", eventData)

	return nil
}

func onChannelAdBreakBegin(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubAdBreakBeginEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received ad break begin: %v", eventData)

	return nil
}
