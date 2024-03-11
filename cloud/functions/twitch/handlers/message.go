package handlers

import (
	"fmt"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/functions/twitch/eventsub"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/chat"
	"github.com/pkulik0/stredono/cloud/platform/providers"
)

func onChannelChatMessage(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubChannelChatMessageEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}

	isMod := false
	isSub := false
	isVip := false
	isOwner := eventData.BroadcasterUserID == eventData.ChatterUserID
	if !isOwner {
		for _, badge := range eventData.Badges {
			if badge.SetID == "moderator" {
				isMod = true
			}
			if badge.SetID == "subscriber" {
				isSub = true
			}
			if badge.SetID == "vip" {
				isVip = true
			}
		}
	} else {
		isSub = true // Just in case the owner wants to use some function that checks this
	}

	role := pb.Role_NORMAL
	if isOwner {
		role = pb.Role_OWNER
	} else if isMod {
		role = pb.Role_MODERATOR
	} else if isVip {
		role = pb.Role_VIP
	}

	return chat.HandleMessage(ctx, &pb.ChatMessage{
		ID:                 eventData.MessageID,
		ChatID:             eventData.BroadcasterUserID,
		ChatName:           eventData.BroadcasterUserName,
		SenderID:           eventData.ChatterUserID,
		SenderName:         eventData.ChatterUserName,
		SenderRole:         role,
		SenderIsSubscriber: isSub,
		Message:            eventData.Message.Text,
		Timestamp:          notification.Timestamp.Unix(),
	}, platform.ProviderTwitch)
}
