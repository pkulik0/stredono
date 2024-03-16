package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func RegisterEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.NewContext(r, &providers.Config{
		DocDb:         true,
		Auth:          true,
		SecretManager: true,
		KeyManager:    true,
		RealtimeDb:    true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	register(ctx, w, r)
}

const (
	methodEmailLink = "emailLink"
	methodTwitch    = "oidc.twitch"
	googleKeysUrl   = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"
)

var (
	errorInvalidToken    = errors.New("invalid token")
	errorContextValue    = errors.New("missing context value")
	errorMissingClaims   = errors.New("missing claims")
	errorInvalidAudience = errors.New("invalid audience")
	errorInvalidMethod   = errors.New("invalid sign in method")
)

type userClaims struct {
	Aud               string `json:"aud"`
	Uid               string `json:"sub"`
	SignInMethod      string `json:"sign_in_method"`
	OauthAccessToken  string `json:"oauth_access_token"`
	OauthRefreshToken string `json:"oauth_refresh_token"`
}

func (c *userClaims) Valid() error {
	log.Printf("Validating claims | %v", c)
	matched, err := regexp.Match(`^https://onregister-[a-z0-9-.]+\.run\.app$`, []byte(c.Aud))
	if err != nil {
		return err
	}
	if !matched {
		return errorInvalidAudience
	}
	if c.Uid == "" {
		return errorMissingClaims
	}
	if c.SignInMethod == "" {
		return errorMissingClaims
	}
	return nil
}

func handleRegistration(ctx *providers.Context, claims *userClaims) error {
	db, ok := ctx.GetDocDb()
	if !ok {
		return errorContextValue
	}
	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		return errorContextValue
	}

	user := &pb.User{
		Username:    "",
		DisplayName: "",
		Uid:         claims.Uid,
		Url:         "",
		Description: "",
	}

	switch claims.SignInMethod {
	case methodEmailLink:
		break
	case methodTwitch: // add new providers here
		err := handleOauthRegistration(ctx, claims, user)
		if err != nil {
			return err
		}
	default:
		log.Errorf("Invalid sign in method | %s", claims.SignInMethod)
		return errorInvalidMethod
	}

	if _, err := db.Collection("users").Doc(user.Uid).Create(ctx.Ctx, user); err != nil {
		return err
	}

	keyUuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	overlayKey := strings.ReplaceAll(keyUuid.String(), "-", "")

	if err := rtdb.NewRef("Data").Child(user.Uid).Set(ctx.Ctx, &pb.UserData{
		Settings: &pb.UserSettings{
			TTS: &pb.TTSSettings{
				VoiceIdPlus:  "onwK4e9ZLuTAKqWW03F9",
				VoiceIdBasic: "pl-PL-Wavenet-C",
				UsePlus:      true,
			},
			Events: &pb.EventsSettings{
				RequireApproval: false,
				MinDisplayTime:  10,
				Tip: &pb.TipSettings{
					Template:  "{user} tipped {value} {currency}",
					MinAmount: 1,
				},
				Cheer: &pb.CheerSettings{
					Template:  "{user} cheered {value} bits",
					MinAmount: 100,
				},
				Follow: &pb.FollowSettings{
					Template:  "{user} followed",
					IsEnabled: false,
				},
				Sub: &pb.SubSettings{
					Template:  "{user} subscribed. It's their {value} month",
					MinMonths: 1,
				},
				SubGift: &pb.SubGiftSettings{
					Template: "{user} gifted {value} subs",
					MinCount: 1,
				},
				Raid: &pb.RaidSettings{
					Template:   "{user} raided with {value} viewers",
					MinViewers: 10,
				},
				ChatTTS: &pb.ChatTTSSettings{
					Template:  "{user} said:",
					IsEnabled: true,
				},
			},
			Alerts: make([]*pb.Alert, 0),
			Overlay: &pb.OverlaySettings{
				Key: overlayKey,
			},
		},
		Media: &pb.MediaRequest{
			IsEnabled: true,
			Settings: &pb.MediaRequestSettings{
				MinRole:         pb.Role_MODERATOR,
				MinViews:        100,
				MinLikes:        10,
				RequireApproval: false,
			},
			Queue: make([]*pb.MediaRequest_QueueItem, 0),
		},
		Commands: make(map[string]string),
	}); err != nil {
		return err
	}

	if err := rtdb.NewRef("Users").Child("Overlay").Child(overlayKey).Set(ctx.Ctx, user.Uid); err != nil {
		return err
	}

	return nil
}

func handleOauthRegistration(ctx *providers.Context, claims *userClaims, user *pb.User) error {
	keyManager, ok := ctx.GetKeyManager()
	if !ok {
		return errorContextValue
	}
	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		return errorContextValue
	}

	client, err := providers.GetHelixClient(ctx)
	if err != nil {
		log.Printf("Failed to get helix client | %v", err)
		return fmt.Errorf("failed to get helix client | %v", err)
	}
	client.SetUserAccessToken(claims.OauthAccessToken)

	users, err := client.GetUsers(&helix.UsersParams{})
	if err != nil {
		log.Printf("Failed to get users | %v", err)
		return err
	}
	if len(users.Data.Users) == 0 {
		return fmt.Errorf("no twitch user returned from api")
	}
	twitchUser := users.Data.Users[0]

	user.DisplayName = twitchUser.DisplayName
	user.PictureUrl = twitchUser.ProfileImageURL
	user.Description = twitchUser.Description

	token := &pb.Token{
		AccessToken:  claims.OauthAccessToken,
		RefreshToken: claims.OauthRefreshToken,
	}
	tokenBytes, err := proto.Marshal(token)
	if err != nil {
		return err
	}

	encryptedToken, err := keyManager.Encrypt(ctx.Ctx, platform.EncryptionKey, tokenBytes)
	if err != nil {
		return err
	}

	if err := rtdb.NewRef("Users").Child(platform.ProviderTwitch).Child(twitchUser.ID).Set(ctx.Ctx, &pb.TokenEntry{
		Uid:            user.Uid,
		EncryptedToken: encryptedToken,
	}); err != nil {
		return err
	}

	if twitchUser.ID == platform.TwitchUid {
		return nil
	}

	modsRes, err := client.GetModerators(&helix.GetModeratorsParams{
		BroadcasterID: twitchUser.ID,
		UserIDs:       []string{platform.TwitchUid},
	})
	if err != nil {
		log.Printf("Failed to get moderators | %v", err)
		return err
	}
	if len(modsRes.Data.Moderators) > 0 {
		log.Printf("Bot account (%s) is already a moderator of %s", platform.TwitchUid, twitchUser.ID)
		return nil
	}

	addRes, err := client.AddChannelModerator(&helix.AddChannelModeratorParams{
		BroadcasterID: twitchUser.ID,
		UserID:        platform.TwitchUid,
	})
	if err != nil {
		log.Printf("Failed to add channel moderator | %v \n %v", err, addRes)
		return err
	}
	if addRes.StatusCode != http.StatusNoContent {
		log.Printf("Failed to add channel moderator | %v", addRes)
		return fmt.Errorf("failed to add channel moderator %v", addRes.StatusCode)
	}

	log.Printf("Added bot account (%s) as a moderator of %s", platform.TwitchUid, twitchUser.ID)
	return nil
}

func getGoogleSigningKeys() (map[string]string, error) {
	res, err := http.Get(googleKeysUrl)
	if err != nil {
		return nil, err
	}

	var keys map[string]string
	if err = json.NewDecoder(res.Body).Decode(&keys); err != nil {
		return nil, err
	}
	return keys, nil
}

type onRegisterPayload struct {
	JwtToken string `json:"jwt"`
}

type onRegisterRequest struct {
	Data onRegisterPayload `json:"data"`
}

func requestToClaims(r *http.Request, signingKeys map[string]string) (*userClaims, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	requestBody := onRegisterRequest{}
	if err := json.Unmarshal(body, &requestBody); err != nil {
		return nil, err
	}

	var token *jwt.Token
	for _, key := range signingKeys {
		token, err = jwt.ParseWithClaims(requestBody.Data.JwtToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err == nil {
			break
		}
	}
	if token == nil {
		return nil, errorInvalidToken
	}

	log.Println("requestToClaims", token.Claims)

	return token.Claims.(*userClaims), nil
}

func register(ctx *providers.Context, w http.ResponseWriter, r *http.Request) {
	signingKeys, err := getGoogleSigningKeys()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	claims, err := requestToClaims(r, signingKeys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handleRegistration(ctx, claims); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := w.Write([]byte("{\"status\":\"OK\"}")); err != nil {
		log.Errorf("Failed to write response | %v", err)
	}
}
