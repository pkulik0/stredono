package providers

import (
	"context"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

func getTwitchOauth2Config(ctx context.Context) (*oauth2.Config, error) {
	secretClient, ok := GetSecretManager(ctx)
	if !ok {
		return nil, platform.ErrorMissingContextValue
	}

	clientSecret, err := secretClient.GetSecret(ctx, modules.TwitchClientSecretName, "latest")
	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     modules.TwitchClientId,
		ClientSecret: string(clientSecret),
		RedirectURL:  modules.TwitchRedirectUrl,
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
