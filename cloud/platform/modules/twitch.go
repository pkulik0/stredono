package modules

import (
	"github.com/nicklaw5/helix"
)

const (
	TwitchClientId         = "t1kl0vkt6hv06bi4ah4691hi8fexso"
	TwitchClientSecretName = "twitch-client-secret"
	TwitchRedirectUrl      = "http://localhost:8080/connectTwitchCallback"
)

type HelixClient interface {
	CreateEventSubSubscription(payload *helix.EventSubSubscription) (*helix.EventSubSubscriptionsResponse, error)
	GetUsers(params *helix.UsersParams) (*helix.UsersResponse, error)
}
