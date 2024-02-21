package platform

import (
	"context"
	"github.com/nicklaw5/helix"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

const (
	twitchClientId         = "t1kl0vkt6hv06bi4ah4691hi8fexso"
	twitchClientSecretName = "twitch-client-secret"
	twitchRedirectUrl      = "http://localhost:8080/connectTwitchCallback"
)

type HelixClient interface {
	CreateEventSubSubscription(payload *helix.EventSubSubscription) (*helix.EventSubSubscriptionsResponse, error)
	GetUsers(params *helix.UsersParams) (*helix.UsersResponse, error)
}

func GetTwitchOauth2Config(ctx context.Context) (*oauth2.Config, error) {
	secretClient, ok := GetSecretManager(ctx)
	if !ok {
		return nil, ErrorMissingContextValue
	}

	clientSecret, err := secretClient.GetSecret(ctx, twitchClientSecretName, "latest")
	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     twitchClientId,
		ClientSecret: string(clientSecret),
		RedirectURL:  twitchRedirectUrl,
		Scopes: []string{
			"user:read:email",
			"moderator:read:followers", // channel.follow
			"channel:read:subscriptions",
			"channel:read:redemptions",
			"bits:read",
			"channel:manage:ads",
			"channel:read:ads", // channel.ad_break_begin
			"channel:manage:broadcast",
			"channel:edit:commercial",
			"channel:read:hype_train",
			"channel:read:goals",
			"channel:read:vips",
			"user:read:broadcast",
			"user:read:chat",
		},
		Endpoint: twitch.Endpoint,
	}, nil
}
