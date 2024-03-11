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

func onStreamOnline(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubStreamOnlineEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}

	helixClient, err := providers.GetHelixBotClient(ctx)
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

func onStreamOffline(ctx *providers.Context, notification *eventsub.Notification) error {
	eventData := &helix.EventSubStreamOfflineEvent{}
	if err := notificationToEvent(notification, eventData); err != nil {
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}
	log.Printf("Received stream offline: %v", eventData)

	helixClient, err := providers.GetHelixBotClient(ctx)
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
