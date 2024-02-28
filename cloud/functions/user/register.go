package user

import (
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
	ctx, err := providers.NewContext(r, &providers.Config{
		DocDb: true,
		Auth:  true,
	})
	if err != nil {
		http.Error(w, platform.ServerErrorMessage, http.StatusInternalServerError)
		return
	}

	register(ctx, w, r)
}

const (
	defaultUsernameLength = 16
	methodEmailLink       = "emailLink"
	methodTwitch          = "oidc.twitch"
	googleKeysUrl         = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"
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
	case methodEmailLink:
		break
	case methodTwitch: // add new providers here
		err := handleOauthRegistration(ctx, user, claims)
		if err != nil {
			return err
		}
	default:
		log.Errorf("Invalid sign in method | %s", claims.SignInMethod)
		return errorInvalidMethod
	}

	db, ok := ctx.GetDocDb()
	if !ok {
		return errorContextValue
	}
	_, err = db.Collection("users").Doc(user.Uid).Create(ctx.Ctx, user)
	return err
}

func handleOauthRegistration(ctx *providers.Context, user *pb.User, claims *userClaims) error {
	db, ok := ctx.GetDocDb()
	if !ok {
		return errorContextValue
	}

	_, err := db.Collection("tokens").Doc(claims.Uid).Set(ctx.Ctx, map[string]interface{}{
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
