package modules

import (
	"github.com/nicklaw5/helix/v2"
)

type HelixClient interface {
	SetUserAccessToken(token string)
	SetRefreshToken(token string)
	RequestAppAccessToken(scopes []string) (*helix.AppAccessTokenResponse, error)
	SetAppAccessToken(token string)
	GetUsers(params *helix.UsersParams) (*helix.UsersResponse, error)
	CreateEventSubSubscription(payload *helix.EventSubSubscription) (*helix.EventSubSubscriptionsResponse, error)
	RemoveEventSubSubscription(id string) (*helix.RemoveEventSubSubscriptionParamsResponse, error)
	GetEventSubSubscriptions(params *helix.EventSubSubscriptionsParams) (*helix.EventSubSubscriptionsResponse, error)
	SendChatMessage(params *helix.SendChatMessageParams) (*helix.ChatMessageResponse, error)
	SendChatAnnouncement(params *helix.SendChatAnnouncementParams) (*helix.SendChatAnnouncementResponse, error)
	GetAuthorizationURL(params *helix.AuthorizationURLParams) string
	RequestUserAccessToken(code string) (*helix.UserAccessTokenResponse, error)
	AddChannelModerator(params *helix.AddChannelModeratorParams) (*helix.AddChannelModeratorResponse, error)
	GetModerators(params *helix.GetModeratorsParams) (*helix.ModeratorsResponse, error)
	EditChannelInformation(params *helix.EditChannelInformationParams) (*helix.EditChannelInformationResponse, error)
}
