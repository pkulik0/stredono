package functions

// https://dev.twitch.tv/docs/eventsub/

import (
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	eventsubSecretName  = GcSecretsPath + "twitch-eventsub-secret/versions/1"
	eventsubPubsubTopic = "twitch-eventsub"

	eventsubIdHeader                  = "Twitch-Eventsub-Message-Id"
	eventsubRetryHeader               = "Twitch-Eventsub-Message-Retry"
	eventsubMessageTypeHeader         = "Twitch-Eventsub-Message-Type"
	eventsubMessageSignatureHeader    = "Twitch-Eventsub-Message-Signature"
	eventsubMessageTimestampHeader    = "Twitch-Eventsub-Message-Timestamp"
	eventsubSubscriptionTypeHeader    = "Twitch-Eventsub-Subscription-Type"
	eventsubSubscriptionVersionHeader = "Twitch-Eventsub-Subscription-Version"

	eventsubMessageTypeWebhookCallback = "webhook_callback_verification"
	eventsubMessageTypeNotification    = "notification"
	eventsubMessageTypeRevocation      = "revocation"

	// Not all event types are supported
	// https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/
	eventTypeChannelUpdate                             = "channel.update"
	eventTypeChannelFollow                             = "channel.follow"
	eventTypeChannelAdBreakBegin                       = "channel.ad_break.begin"
	eventTypeChannelChatMessage                        = "channel.chat.message"
	eventTypeChannelChatNotif                          = "channel.chat.notification"
	eventTypeChannelChatSettings                       = "channel.chat_settings.update"
	eventTypeChannelSubscription                       = "channel.subscribe"
	eventTypeChannelSubscriptionEnd                    = "channel.subscription.end"
	eventTypeChannelSubscriptionGift                   = "channel.subscription.gift"
	eventTypeChannelSubscriptionMessage                = "channel.subscription.message"
	eventTypeChannelCheer                              = "channel.cheer"
	eventTypeChannelRaid                               = "channel.raid"
	eventTypeChannelModeratorAdd                       = "channel.moderator.add"
	eventTypeChannelModeratorRemove                    = "channel.moderator.remove"
	eventTypeChannelPointsCustomRewardAdd              = "channel.channel_points_custom_reward.add"
	eventTypeChannelPointsCustomRewardUpdate           = "channel.channel_points_custom_reward.update"
	eventTypeChannelPointsCustomRewardRemove           = "channel.channel_points_custom_reward.remove"
	eventTypeChannelPointsCustomRewardRedemptionAdd    = "channel.channel_points_custom_reward_redemption.add"
	eventTypeChannelPointsCustomRewardRedemptionUpdate = "channel.channel_points_custom_reward_redemption.update"
	eventTypeStreamOnline                              = "stream.online"
	eventTypeStreamOffline                             = "stream.offline"
	eventTypeUserAuthorizationGrant                    = "user.authorization.grant"
	eventTypeUserAuthorizationRevoke                   = "user.authorization.revoke"
	eventTypeUserUpdate                                = "user.update"

	messageFragmentTypeText      = "text"
	messageFragmentTypeCheermote = "cheermote"
	messageFragmentTypeEmote     = "emote"
	messageFragmentTypeMention   = "mention"

	chatNotifSub              = "sub"
	chatNotifResub            = "resub"
	chatNotifSubGift          = "subgift"
	chatNotifCommunitySubGift = "community_sub_gift"
	chatNotifGiftPaidUpgrade  = "gift_paid_upgrade"
	chatNotifPrimePaidUpgrade = "prime_paid_upgrade"
	chatNotifRaid             = "raid"
	chatNotifUnraid           = "unraid"
	chatNotifPayItForward     = "pay_it_forward"
	chatNotifAnnouncement     = "announcement"
	chatNotifBitsBadgeTier    = "bits_badge_tier"
	chatNotifCharityDonation  = "charity_donation"

	streamOnlineTypeLive       = "live"
	streamOnlineTypePlaylist   = "playlist"
	streamOnlineTypeWatchParty = "watch_party"
	streamOnlineTypePremiere   = "premiere"
	streamOnlineTypeReRun      = "rerun"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	CorsMiddleware(CloudMiddleware(CloudConfig{
		Realtime: true,
		Secrets:  true,
		Pubsub:   true,
	}, webhook))(w, r)
}

func calculateHmacSignature(secret string, headers *eventsubHeaders, body string) string {
	h := hmac.New(sha256.New, []byte(secret))
	message := headers.Id + headers.Timestamp + body
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func webhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Failed to read request | %s", err)
		http.Error(w, "Failed to read request", http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Errorf("Failed to close request body | %s", err)
		}
	}(r.Body)
	headers := newEventsubHeaders(r)

	secretsClient, ok := GetSecretsManager(r.Context())
	if !ok {
		log.Error("Failed to get secrets client")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	eventsubSecretResponse, err := secretsClient.AccessSecretVersion(r.Context(), &secretmanagerpb.AccessSecretVersionRequest{
		Name: eventsubSecretName,
	})
	if err != nil {
		log.Errorf("Failed to get eventsub secret | %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	eventsubSecret := string(eventsubSecretResponse.Payload.Data)

	expectedSig := "sha256=" + calculateHmacSignature(eventsubSecret, &headers, string(body))
	if !hmac.Equal([]byte(expectedSig), []byte(headers.Signature)) {
		log.Errorf("Failed to validate signature")
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
		return
	}

	var notification eventsubNotification
	if err := json.Unmarshal(body, &notification); err != nil {
		log.Errorf("Failed to unmarshal request | %s", err)
		http.Error(w, "Failed to unmarshal request", http.StatusBadRequest)
		return
	}

	switch headers.Type {
	case eventsubMessageTypeWebhookCallback:
		_, err = w.Write([]byte(notification.Challenge))
		if err != nil {
			log.Errorf("Failed to write response | %s", err)
			return
		}
	case eventsubMessageTypeNotification:
		handleEvent(w, r, &notification)
	case eventsubMessageTypeRevocation:
		log.Printf("Revoked subscription: %s", notification.Subscription.Id)
		// TODO: handle revocation
	default:
		log.Errorf("unknown message type: %s", headers.Type)
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
	}
}

type eventsubHeaders struct {
	Id           string
	Retry        string
	Type         string
	Signature    string
	Timestamp    string
	Subscription struct {
		Type    string
		Version string
	}
}

func newEventsubHeaders(r *http.Request) eventsubHeaders {
	return eventsubHeaders{
		Id:        r.Header.Get(eventsubIdHeader),
		Retry:     r.Header.Get(eventsubRetryHeader),
		Type:      r.Header.Get(eventsubMessageTypeHeader),
		Signature: r.Header.Get(eventsubMessageSignatureHeader),
		Timestamp: r.Header.Get(eventsubMessageTimestampHeader),
		Subscription: struct {
			Type    string
			Version string
		}{
			Type:    r.Header.Get(eventsubSubscriptionTypeHeader),
			Version: r.Header.Get(eventsubSubscriptionVersionHeader),
		},
	}
}

type eventsubSubscription struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Version   string `json:"version"`
	Cost      int    `json:"cost"`
	Condition struct {
		BroadcasterUserId string `json:"broadcaster_user_id"`
		ModeratorUserId   string `json:"moderator_user_id"`
	} `json:"condition"`
	Transport struct {
		Method   string `json:"method"`
		Callback string `json:"callback"`
		Secret   string `json:"secret"`
	}
	CreatedAt string `json:"created_at"`
}

type eventsubNotification struct {
	Subscription eventsubSubscription   `json:"subscription"`
	Event        map[string]interface{} `json:"event"`
	Challenge    string                 `json:"challenge"`
}

func unmarshalOrError[T any](w http.ResponseWriter, notif *eventsubNotification, out *T) bool {
	jsonData, err := json.Marshal(notif.Event)
	if err != nil {
		log.Errorf("Failed to marshal event | %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return false
	}

	if err := json.Unmarshal(jsonData, out); err != nil {
		log.Errorf("Failed to unmarshal event | %s", err)
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return false
	}

	return true
}

func handleEvent(w http.ResponseWriter, r *http.Request, notification *eventsubNotification) {
	pubsubClient, ok := GetPubsub(r.Context())
	if !ok {
		log.Error("Failed to get pubsub client")
		http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	topic := pubsubClient.Topic(eventsubPubsubTopic)
	topic.PublishSettings.NumGoroutines = 1
	defer topic.Stop()

	publishEvent := func(event interface{}, uid string, kind string) bool {
		data, err := json.Marshal(event)
		if err != nil {
			log.Errorf("Failed to marshal event | %s", err)
			http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
			return false
		}
		topic.Publish(r.Context(), &pubsub.Message{
			Data: data,
			Attributes: map[string]string{
				"uid":   uid,
				"event": kind,
			},
		})
		return true
	}

	switch notification.Subscription.Type {
	case eventTypeChannelUpdate:
		var event eventChannelUpdate
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.update: %+v", event)
	case eventTypeChannelFollow:
		var event eventFollow
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.follow: %+v", event)
		if !publishEvent(event, event.BroadcasterId, "follow") {
			return
		}
	case eventTypeChannelAdBreakBegin:
		var event eventAdBreakBegin
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.ad_break.begin: %+v", event)
	case eventTypeChannelChatMessage:
		var event eventChatMessage
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.chat.message: %+v", event)
	case eventTypeChannelChatNotif:
		var event eventChatNotification
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.chat.notification: %+v", event)
	case eventTypeChannelChatSettings:
		var event eventChatSettingsUpdate
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.chat_settings.update: %+v", event)
	case eventTypeChannelSubscription:
		var event eventChannelSubscription
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.subscribe: %+v", event)
	case eventTypeChannelSubscriptionEnd:
		var event eventChannelSubscription
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.subscription.end: %+v", event)
	case eventTypeChannelSubscriptionGift:
		var event eventChannelSubscriptionGift
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.subscription.gift: %+v", event)
	case eventTypeChannelSubscriptionMessage:
		var event eventChannelSubscriptionMessage
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.subscription.message: %+v", event)
	case eventTypeChannelCheer:
		var event eventChannelCheer
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.cheer: %+v", event)
	case eventTypeChannelRaid:
		var event eventChannelRaid
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.raid: %+v", event)
	case eventTypeChannelModeratorAdd:
		var event eventModeratorChange
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.moderator.add: %+v", event)
	case eventTypeChannelModeratorRemove:
		var event eventModeratorChange
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.moderator.remove: %+v", event)
	case eventTypeChannelPointsCustomRewardAdd:
		var event eventChannelPointsCustomReward
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.channel_points_custom_reward.add: %+v", event)
	case eventTypeChannelPointsCustomRewardUpdate:
		var event eventChannelPointsCustomReward
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.channel_points_custom_reward.update: %+v", event)
	case eventTypeChannelPointsCustomRewardRemove:
		var event eventChannelPointsCustomReward
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.channel_points_custom_reward.remove: %+v", event)
	case eventTypeChannelPointsCustomRewardRedemptionAdd:
		var event eventChannelCustomRewardRedemption
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.channel_points_custom_reward_redemption.add: %+v", event)
	case eventTypeChannelPointsCustomRewardRedemptionUpdate:
		var event eventChannelCustomRewardRedemption
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("channel.channel_points_custom_reward_redemption.update: %+v", event)
	case eventTypeStreamOnline:
		var event eventStreamOnline
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("stream.online: %+v", event)
	case eventTypeStreamOffline:
		var event eventStreamOffline
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("stream.offline: %+v", event)
	case eventTypeUserAuthorizationGrant:
		var event eventUserAuthorization
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("user.authorization.grant: %+v", event)
	case eventTypeUserAuthorizationRevoke:
		var event eventUserAuthorization
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("user.authorization.revoke: %+v", event)
	case eventTypeUserUpdate:
		var event eventUserUpdate
		if !unmarshalOrError(w, notification, &event) {
			return
		}
		log.Printf("user.update: %+v", event)
	default:
		log.Errorf("unknown subscription received: %+v", notification.Subscription)
		http.Error(w, BadRequestMessage, http.StatusBadRequest)
	}
}

type broadcasterData struct {
	BroadcasterId    string `json:"broadcaster_user_id"`
	BroadcasterLogin string `json:"broadcaster_user_login"`
	BroadcasterName  string `json:"broadcaster_user_name"`
}

type userData struct {
	UserId    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

type chatterData struct {
	ChatterId    string `json:"chatter_id"`
	ChatterLogin string `json:"chatter_login"`
	ChatterName  string `json:"chatter_name"`
}

// channel.update
type eventChannelUpdate struct {
	broadcasterData
	Title                       string   `json:"title"`
	Language                    string   `json:"language"`
	CategoryId                  string   `json:"category_id"`
	CategoryName                string   `json:"category_name"`
	ContentClassificationLabels []string `json:"content_classification_labels"`
}

// channel.follow
type eventFollow struct {
	broadcasterData
	userData
	FollowedAt string `json:"followed_at"`
}

// channel.ad_break.begin
type eventAdBreakBegin struct {
	broadcasterData
	RequesterId    string `json:"requester_id"`
	RequesterLogin string `json:"requester_login"`
	RequesterName  string `json:"requester_name"`
	Duration       int    `json:"duration"`
	IsAutomatic    bool   `json:"is_automatic"`
	StartedAt      string `json:"started_at"`
}

type badgeData struct {
	Id    string `json:"id"`
	SetId string `json:"set_id"`
	Info  string `json:"info"`
}

type emoteData struct {
	Id         string `json:"id"`
	EmoteSetId string `json:"emote_set_id"`
	OwnerId    string `json:"owner_id"`
	Format     string `json:"format"`
}

type cheermoteData struct {
	Prefix string `json:"prefix"`
	Bits   int    `json:"bits"`
	Tier   int    `json:"tier"`
}

type fragmentsData struct {
	Type      string         `json:"type"`
	Text      string         `json:"text"`
	Cheermote *cheermoteData `json:"cheermote"`
	Emote     *emoteData     `json:"emote"`
	Mention   *userData      `json:"mention"`
}

type replyData struct {
	ParentMessageId   string `json:"parent_message_id"`
	ParentMessageBody string `json:"parent_message_body"`
	ParentUserId      string `json:"parent_user_id"`
	ParentUserLogin   string `json:"parent_user_login"`
	ParentUserName    string `json:"parent_user_name"`
	ThreadMessageId   string `json:"thread_message_id"`
	ThreadUserId      string `json:"thread_user_id"`
	ThreadUserLogin   string `json:"thread_user_login"`
	ThreadUserName    string `json:"thread_user_name"`
}

type messageData struct {
	Text      string        `json:"text"`
	Fragments fragmentsData `json:"fragments"`
}

type cheerData struct {
	Bits int `json:"bits"`
}

// channel.chat.message
type eventChatMessage struct {
	broadcasterData
	chatterData
	MessageId      string      `json:"message_id"`
	Message        messageData `json:"message"`
	Color          string      `json:"color"`
	Badges         []badgeData `json:"badges"`
	MessageType    string      `json:"message_type"`
	Cheer          *cheerData  `json:"cheer"`
	Reply          *replyData  `json:"reply"`
	CustomRewardId *string     `json:"channel_points_custom_reward_id"`
}

type subData struct {
	SubTier        string `json:"sub_tier"`
	IsPrime        bool   `json:"is_prime"`
	DurationMonths int    `json:"duration_months"`
}

type resubData struct {
	SubTier           string  `json:"sub_tier"`
	IsPrime           *bool   `json:"is_prime"`
	DurationMonths    int     `json:"duration_months"`
	CumulativeMonths  int     `json:"cumulative_months"`
	StreakMonths      int     `json:"streak_months"`
	IsGift            bool    `json:"is_gift"`
	GifterIsAnonymous *bool   `json:"gifter_is_anonymous"`
	GifterUserId      *string `json:"gifter_user_id"`
	GifterUserName    *string `json:"gifter_user_name"`
	GifterUserLogin   *string `json:"gifter_user_login"`
}

type subGiftData struct {
	DurationMonths     int     `json:"duration_months"`
	CumulativeTotal    *int    `json:"cumulative_total"`
	RecipientUserId    string  `json:"recipient_user_id"`
	RecipientUserName  string  `json:"recipient_user_name"`
	RecipientUserLogin string  `json:"recipient_user_login"`
	SubTier            string  `json:"sub_tier"`
	CommunityGiftId    *string `json:"community_gift_id"`
}

type communitySubGiftData struct {
	Id              string `json:"id"`
	Total           int    `json:"total"`
	SubTier         string `json:"sub_tier"`
	CumulativeTotal *int   `json:"cumulative_total"`
}

type raidData struct {
	userData
	ViewerCount     int    `json:"viewer_count"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type payItForwardData struct {
	GifterIsAnonymous bool   `json:"gifter_is_anonymous"`
	GifterUserId      string `json:"gifter_user_id"`
	GifterUserName    string `json:"gifter_user_name"`
	GifterUserLogin   string `json:"gifter_user_login"`
}

type announcementData struct {
	Color string `json:"color"`
}

type charityDonationData struct {
	CharityName string `json:"charity_name"`
	Amount      struct {
		Value        int    `json:"value"`
		DecimalPlace int    `json:"decimal_place"`
		Currency     string `json:"currency"`
	}
}

type bitsBadgeTierData struct {
	Tier string `json:"tier"`
}

// channel.chat.notification
type eventChatNotification struct {
	broadcasterData
	chatterData
	ChatterIsAnonymous bool                  `json:"chatter_is_anonymous"`
	Color              string                `json:"color"`
	Badges             []badgeData           `json:"badges"`
	SystemMessage      string                `json:"system_message"`
	MessageId          string                `json:"message_id"`
	Message            messageData           `json:"message"`
	NoticeType         string                `json:"notice_type"`
	Sub                *subData              `json:"sub"`
	Resub              *resubData            `json:"resub"`
	SubGift            *subGiftData          `json:"sub_gift"`
	CommunitySubGift   *communitySubGiftData `json:"community_sub_gift"`
	GiftPaidUpgrade    *userData             `json:"gift_paid_upgrade"`
	PrimePaidUpgrade   *userData             `json:"prime_paid_upgrade"`
	Raid               *raidData             `json:"raid"`
	Unraid             interface{}           `json:"unraid"` // No data
	PayItForward       *payItForwardData     `json:"pay_it_forward"`
	Announcement       *announcementData     `json:"announcement"`
	CharityDonation    *charityDonationData  `json:"charity_donation"`
	BitsBadgeTier      *bitsBadgeTierData    `json:"bits_badge_tier"`
}

// channel.chat_settings.update
type eventChatSettingsUpdate struct {
	broadcasterData
	IsEmoteMode          bool `json:"emote_mode"`
	IsFollowerMode       bool `json:"follower_mode"`
	FollowerModeDuration int  `json:"follower_mode_duration_minutes"`
	IsSlowMode           bool `json:"slow_mode"`
	SlowModeDuration     int  `json:"slow_mode_wait_time_seconds"`
	IsSubMode            bool `json:"subscriber_mode"`
	IsUniqueMode         bool `json:"unique_chat_mode"`
}

// channel.subscribe
// channel.subscription.end
type eventChannelSubscription struct {
	broadcasterData
	userData
	IsGift bool   `json:"is_gift"`
	Tier   string `json:"tier"`
}

// channel.subscription.gift
type eventChannelSubscriptionGift struct {
	eventChannelSubscription
	CumulativeMonths int  `json:"cumulative_months"`
	IsAnonymous      bool `json:"is_anonymous"`
}

// channel.subscription.message
type eventChannelSubscriptionMessage struct {
	eventChannelSubscription
	Message          string `json:"message"`
	CumulativeMonths int    `json:"cumulative_months"`
	StreakMonths     int    `json:"streak_months"`
	DurationMonths   int    `json:"duration_months"`
}

// channel.cheer
type eventChannelCheer struct {
	broadcasterData
	userData
	IsAnonymous bool   `json:"is_anonymous"`
	Bits        int    `json:"bits"`
	Message     string `json:"message"`
}

// channel.raid
type eventChannelRaid struct {
	FromBroadcasterId    string `json:"from_broadcaster_user_id"`
	FromBroadcasterLogin string `json:"from_broadcaster_user_login"`
	FromBroadcasterName  string `json:"from_broadcaster_user_name"`
	ToBroadcasterId      string `json:"to_broadcaster_user_id"`
	ToBroadcasterLogin   string `json:"to_broadcaster_user_login"`
	ToBroadcasterName    string `json:"to_broadcaster_user_name"`
	Viewers              int    `json:"viewer_count"`
}

// channel.moderator.add
// channel.moderator.remove
type eventModeratorChange struct {
	broadcasterData
	userData
}

type maxPerStreamData struct {
	IsEnabled bool `json:"is_enabled"`
	Value     int  `json:"value"`
}

type imageData struct {
	Url1x string `json:"url_1x"`
	Url2x string `json:"url_2x"`
	Url4x string `json:"url_4x"`
}

type cooldownData struct {
	IsEnabled bool `json:"is_enabled"`
	Seconds   int  `json:"seconds"`
}

// channel.channel_points_custom_reward.add
// channel.channel_points_custom_reward.update
// channel.channel_points_custom_reward.remove
type eventChannelPointsCustomReward struct {
	broadcasterData
	Id                               string           `json:"id"`
	IsEnabled                        bool             `json:"is_enabled"`
	IsPaused                         bool             `json:"is_paused"`
	IsInStock                        bool             `json:"is_in_stock"`
	Title                            string           `json:"title"`
	Cost                             int              `json:"cost"`
	Prompt                           string           `json:"prompt"`
	IsUserInputRequired              bool             `json:"is_user_input_required"`
	ShouldRedemptionSkipRequestQueue bool             `json:"should_redemptions_skip_request_queue"`
	MaxPerStream                     maxPerStreamData `json:"max_per_stream"`
	MaxPerUserPerStream              maxPerStreamData `json:"max_per_user_per_stream"`
	BackgroundColor                  string           `json:"background_color"`
	Image                            *imageData       `json:"image"`
	DefaultImage                     imageData        `json:"default_image"`
	GlobalCooldown                   cooldownData     `json:"global_cooldown"`
	CooldownExpiresAt                *string          `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream *int             `json:"redemptions_redeemed_current_stream"`
}

type rewardData struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Cost   int    `json:"cost"`
	Prompt string `json:"prompt"`
}

// channel.channel_points_custom_reward_redemption.add
// channel.channel_points_custom_reward_redemption.update
type eventChannelCustomRewardRedemption struct {
	broadcasterData
	userData
	Id         string     `json:"id"`
	UserInput  string     `json:"user_input"`
	Status     string     `json:"status"`
	Reward     rewardData `json:"reward"`
	RedemeedAt string     `json:"redeemed_at"`
}

// stream.online
type eventStreamOnline struct {
	broadcasterData
	Id        string `json:"id"`
	Type      string `json:"type"`
	StartedAt string `json:"started_at"`
}

// stream.offline
type eventStreamOffline struct {
	broadcasterData
}

// user.authorization.grant
// user.authorization.revoke
type eventUserAuthorization struct {
	userData
	ClientId string `json:"client_id"`
}

// user.update
type eventUserUpdate struct {
	userData
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"email_verified"`
	Description     string `json:"description"`
}
