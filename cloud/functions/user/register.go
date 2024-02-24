package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func RegisterEntrypoint(w http.ResponseWriter, r *http.Request) {
	ctx, err := providers.CreateContext(r.Context(), &providers.Config{
		DocDb: true,
		Auth:  true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}
	r = r.WithContext(ctx)

	register(w, r)
}

const (
	defaultUsernameLength    = 16
	identityProviderPassword = "password"
	identityProviderTwitch   = "oidc.twitch"
	googleKeysUrl            = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"
)

var (
	errorInvalidToken    = errors.New("invalid token")
	errorContextValue    = errors.New("missing context value")
	errorMissingClaims   = errors.New("missing claims")
	errorInvalidAudience = errors.New("invalid audience")
	errorInvalidProvider = errors.New("invalid provider")
)

type userClaims struct {
	Aud          string `json:"aud"`
	Uid          string `json:"sub"`
	ProviderData struct {
		ProviderId string `json:"provider_id"`
	} `json:"provider_data"`
}

func (c *userClaims) Valid() error {
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
	if c.ProviderData.ProviderId != identityProviderPassword && c.ProviderData.ProviderId != identityProviderTwitch {
		return errorInvalidProvider
	}
	return nil
}

type userClaimsPassword struct {
	userClaims
	Email string `json:"email"`
}

func (c *userClaimsPassword) Valid() error {
	if err := c.userClaims.Valid(); err != nil {
		return err
	}
	if c.Email == "" {
		return errorMissingClaims
	}
	return nil
}

type userClaimsOauth struct {
	userClaims
	SignInAttributes struct {
		PreferredUsername string `json:"preferred_username"`
	} `json:"sign_in_attributes"`
	OauthAccessToken  string `json:"oauth_access_token"`
	OauthRefreshToken string `json:"oauth_refresh_token"`
}

func (c *userClaimsOauth) Valid() error {
	if err := c.userClaims.Valid(); err != nil {
		return err
	}
	if c.SignInAttributes.PreferredUsername == "" {
		return errorMissingClaims
	}
	if c.OauthAccessToken == "" {
		return errorMissingClaims
	}
	if c.OauthRefreshToken == "" {
		return errorMissingClaims
	}
	return nil
}

func handleRegistration(ctx context.Context, claims jwt.Claims) error {
	registerClaims, ok := claims.(*userClaims)
	if !ok {
		return errorMissingClaims
	}
	if err := registerClaims.Valid(); err != nil {
		return err
	}

	randomUsernameUuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	randomUsername := strings.ReplaceAll(randomUsernameUuid.String(), "-", "")[:defaultUsernameLength]

	user := &pb.User{
		Username:      randomUsername,
		DisplayName:   "???",
		Uid:           registerClaims.Uid,
		Url:           "",
		Description:   "Default description.", // TODO: Change to something better, based on locale
		MinimumAmount: 1,
		Currency:      pb.Currency_PLN,
		Alerts:        make([]*pb.Alert, 0),
	}

	provider := registerClaims.ProviderData.ProviderId
	switch provider {
	case identityProviderPassword:
		claimsPassword, ok := claims.(*userClaimsPassword)
		if !ok {
			return errorMissingClaims
		}
		if err := claimsPassword.Valid(); err != nil {
			return err
		}
		user.DisplayName = strings.Split(claimsPassword.Email, "@")[0]
	case identityProviderTwitch: // add new providers here
		claimsOauth, ok := claims.(*userClaimsOauth)
		if !ok {
			return errorMissingClaims
		}
		if err := claimsOauth.Valid(); err != nil {
			return err
		}
		err := handleOauthRegistration(ctx, user, claimsOauth, provider)
		if err != nil {
			return err
		}
	default:
		return errorInvalidProvider
	}

	db, ok := providers.GetDocDb(ctx)
	if !ok {
		return errorContextValue
	}
	_, err = db.Collection("users").Doc(user.Uid).Create(ctx, user)
	return err
}

func handleOauthRegistration(ctx context.Context, user *pb.User, claims *userClaimsOauth, provider string) error {
	db, ok := providers.GetDocDb(ctx)
	if !ok {
		return errorContextValue
	}

	_, err := db.Collection("tokens").Doc(claims.Uid).Set(ctx, map[string]interface{}{
		provider: &pb.Token{
			Access:  claims.OauthAccessToken,
			Refresh: claims.OauthRefreshToken,
		},
	}, &modules.DbOpts{MergeAll: true})
	if err != nil {
		return err
	}

	user.DisplayName = claims.SignInAttributes.PreferredUsername
	return nil
}

func getGoogleSigningKeys() (map[string]string, error) {
	googleKeys, err := http.Get(googleKeysUrl)
	if err != nil {
		return nil, err
	}

	var keys map[string]string
	if err = json.NewDecoder(googleKeys.Body).Decode(&keys); err != nil {
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

func requestToClaims(r *http.Request, signingKeys map[string]string) (jwt.Claims, error) {
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
		token, err = jwt.Parse(requestBody.Data.JwtToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err == nil {
			break
		}
	}
	if token == nil {
		return nil, errorInvalidToken
	}

	return token.Claims, nil
}

func register(w http.ResponseWriter, r *http.Request) {
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

	err = handleRegistration(r.Context(), claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := w.Write([]byte("OK")); err != nil {
		log.Errorf("Failed to write response | %v", err)
	}
}
