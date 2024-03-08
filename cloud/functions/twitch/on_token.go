package twitch

import (
	"context"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/googleapis/google-cloudevents-go/cloud/firestoredata"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"
)

func eventDocumentToToken(document *firestoredata.Document) (*pb.Token, error) {
	fields := document.GetFields()

	accessToken, ok := fields["AccessToken"]
	if !ok {
		return nil, fmt.Errorf("missing access token")
	}

	refreshToken, ok := fields["RefreshToken"]
	if !ok {
		return nil, fmt.Errorf("missing refresh token")
	}

	providerUid, ok := fields["ProviderUid"]
	if !ok {
		return nil, fmt.Errorf("missing provider uid")
	}

	token := &pb.Token{
		AccessToken:  accessToken.GetStringValue(),
		RefreshToken: refreshToken.GetStringValue(),
		ProviderUid:  providerUid.GetStringValue(),
	}

	return token, nil
}

// Event name - version
var eventsubSubs = map[string]string{
	"channel.update":               "2",
	"channel.follow":               "2",
	"channel.chat.message":         "1",
	"channel.subscribe":            "1",
	"channel.subscription.message": "1",
	"channel.subscription.gift":    "1",
	"channel.cheer":                "1",
	"channel.raid":                 "1",
	"channel.ban":                  "1",
	"channel.unban":                "1",
	"channel.moderator.add":        "1",
	"channel.moderator.remove":     "1",
	"stream.online":                "1",
	"stream.offline":               "1",
	"user.authorization.revoke":    "1",
	"user.update":                  "1",
}

func OnTokenEntrypoint(ctx context.Context, e event.Event) error {
	iCtx, err := providers.NewContextEvent(ctx, &providers.Config{
		SecretManager: true,
	})
	if err != nil {
		log.Printf("Failed to create context | %v", err)
		return fmt.Errorf("failed to create context | %v", err)
	}

	eventData := &firestoredata.DocumentEventData{}
	if err := proto.Unmarshal(e.Data(), eventData); err != nil {
		log.Printf("Failed to unmarshal data | %v", err)
		return fmt.Errorf("failed to unmarshal data | %v", err)
	}

	log.Printf("Received event: %v", eventData)

	source := e.Source()
	log.Printf("Received event from: %v", source) // TODO: what happens here? debug

	doc := eventData.GetValue()
	if doc == nil {
		log.Printf("Token removed - skipping")
		return nil
	}
	docName := doc.GetName()
	uid := docName[strings.LastIndex(docName, "/")+1:]
	log.Printf("Received token for user: %v", uid)

	tokenData, err := eventDocumentToToken(doc)
	if err != nil {
		log.Printf("Failed to get token from document | %v", err)
		return fmt.Errorf("failed to get token from document | %v", err)
	}

	if updateMask := eventData.GetUpdateMask(); updateMask != nil {
		oldDoc := eventData.GetOldValue()
		oldTokenData, err := eventDocumentToToken(oldDoc)
		if err != nil {
			log.Printf("Failed to get old token from document | %v", err)
			return fmt.Errorf("failed to get old token from document | %v", err)
		}

		if tokenData.RefreshToken == oldTokenData.RefreshToken {
			log.Printf("Token refresh did not change - skipping")
			return nil
		}

		log.Printf("Token refresh changed - updating")
	}

	helixClient, err := providers.GetHelixAppClient(iCtx)
	if err != nil {
		log.Printf("Failed to get helix client | %v", err)
		return fmt.Errorf("failed to get helix client | %v", err)
	}

	secretManager, ok := iCtx.GetSecretManager()
	if !ok {
		log.Printf("No secret manager")
		return platform.ErrorMissingContextValue
	}
	webhookSecret, err := secretManager.GetSecret(ctx, "twitch-eventsub-secret", "latest")
	if err != nil {
		log.Printf("Failed to get webhook secret | %v", err)
		return fmt.Errorf("failed to get webhook secret | %v", err)
	}

	transport := helix.EventSubTransport{
		Method:   "webhook",
		Callback: "https://europe-west1-stredono-6394ee11.cloudfunctions.net/TwitchWebhook",
		Secret:   string(webhookSecret),
	}
	condition := helix.EventSubCondition{
		BroadcasterUserID:   tokenData.ProviderUid,
		ToBroadcasterUserID: tokenData.ProviderUid,
		ModeratorUserID:     tokenData.ProviderUid,
		UserID:              "1033918710",
		ClientID:            platform.TwitchClientId,
	}

	createdTypes := make([]string, 0)
	for subType, subVersion := range eventsubSubs {
		res, err := helixClient.CreateEventSubSubscription(&helix.EventSubSubscription{
			Type:      subType,
			Version:   subVersion,
			Condition: condition,
			Transport: transport,
		})
		if err != nil {
			log.Printf("Failed to create eventsub subscription | %v", err)
			return err
		}

		// Either new or already exists (token refresh for somebody that already has subscriptions)
		if res.StatusCode != http.StatusAccepted && res.StatusCode != http.StatusConflict {
			log.Printf("Failed to create eventsub subscription %s | %v", subType, res)
			return fmt.Errorf("failed to create eventsub subscription %s | %v", subType, res)
		}

		log.Printf("Eventsub subscription created: %+v", res)
		createdTypes = append(createdTypes, subType)
	}

	log.Printf("User %s subscribed to eventsub types: %v", uid, createdTypes)
	return nil
}
