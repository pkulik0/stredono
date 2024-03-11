package handlers

import (
	"fmt"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/functions/twitch/eventsub"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func onUserUpdate(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubUserUpdateEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received user update: %v", eventData)

	return nil
}

func onUserAuthorizationGrant(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubUserAuthenticationGrantEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("User authorization granted: %v", eventData)

	helixClient, err := providers.GetHelixAppClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to get helix client | %v", err)
	}

	transport, err := providers.GetHelixTransport(ctx)
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

func onUserAuthorizationRevoke(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubUserAuthenticationRevokeEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("User authorization revoked: %v", eventData)

	return nil
}
