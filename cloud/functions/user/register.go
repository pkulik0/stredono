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
	Aud               string `json:"aud"`
	Uid               string `json:"sub"`
	SignInMethod      string `json:"sign_in_method"`
	OauthAccessToken  string `json:"oauth_access_token"`
	OauthRefreshToken string `json:"oauth_refresh_token"`
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
	if c.SignInMethod == "" {
		return errorMissingClaims
	}
	return nil
}

func handleRegistration(ctx context.Context, claims *userClaims) error {
	log.Println("handleRegistration", claims)

	randomUsernameUuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	randomUsername := strings.ReplaceAll(randomUsernameUuid.String(), "-", "")[:defaultUsernameLength]

	user := &pb.User{
		Username:     randomUsername,
		DisplayName:  "Unnamed",
		Uid:          claims.Uid,
		Url:          "",
		Description:  "Default description.", // TODO: Change to something better, based on locale
		MinAmount:    1,
		MinAuthLevel: pb.AuthLevel_NONE,
		Currency:     pb.Currency_PLN,
	}

	switch claims.SignInMethod {
	case identityProviderPassword:
		break
	case identityProviderTwitch: // add new providers here
		err := handleOauthRegistration(ctx, user, claims)
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

func handleOauthRegistration(ctx context.Context, user *pb.User, claims *userClaims) error {
	db, ok := providers.GetDocDb(ctx)
	if !ok {
		return errorContextValue
	}

	_, err := db.Collection("tokens").Doc(claims.Uid).Set(ctx, map[string]interface{}{
		claims.SignInMethod: &pb.Token{
			Access:  claims.OauthAccessToken,
			Refresh: claims.OauthRefreshToken,
		},
	}, modules.DbOpts{MergeAll: true})
	if err != nil {
		return err
	}

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

	return token.Claims.(*userClaims), nil
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

	if _, err := w.Write([]byte("{\"status\":\"OK\"}")); err != nil {
		log.Errorf("Failed to write response | %v", err)
	}
}
