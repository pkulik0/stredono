package modules

import (
	"github.com/nicklaw5/helix/v2"
)

type HelixClient interface {
	SetUserAccessToken(token string)
	RequestAppAccessToken(scopes []string) (*helix.AppAccessTokenResponse, error)
	SetAppAccessToken(token string)
	GetUsers(params *helix.UsersParams) (*helix.UsersResponse, error)
	CreateEventSubSubscription(payload *helix.EventSubSubscription) (*helix.EventSubSubscriptionsResponse, error)
	RemoveEventSubSubscription(id string) (*helix.RemoveEventSubSubscriptionParamsResponse, error)
	GetEventSubSubscriptions(params *helix.EventSubSubscriptionsParams) (*helix.EventSubSubscriptionsResponse, error)
}
