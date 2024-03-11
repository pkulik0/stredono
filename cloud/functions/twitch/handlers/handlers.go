package handlers

import (
	"encoding/json"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/functions/twitch/eventsub"
	"github.com/pkulik0/stredono/cloud/platform/providers"
)

var TypeToHandler = map[string]func(ctx *providers.Context, notification *eventsub.Notification) error{
	helix.EventSubTypeChannelAdBreakBegin:        onChannelAdBreakBegin,
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

	helix.EventSubTypeUserAuthorizationRevoke: onUserAuthorizationRevoke,
	helix.EventSubTypeUserAuthorizationGrant:  onUserAuthorizationGrant,
}

var eventsubSubs = map[string]string{
	helix.EventSubTypeChannelAdBreakBegin:        "1",
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

func notificationToEvent[T any](notification *eventsub.Notification, v *T) error {
	data, err := json.Marshal(notification.Event)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
