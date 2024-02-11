package twitch

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/functions/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

const twitchContextKey = "twitchHelix"

func TwitchMiddleware(next util.HandlerFunc) util.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, ok := util.GetAuthToken(r.Context())
		if !ok {
			log.Error("Failed to get auth token")
			http.Error(w, util.ServerErrorMessage, http.StatusInternalServerError)
			return
		}

		firestoreClient, ok := util.GetFirestore(r.Context())
		if !ok {
			log.Error("Failed to get firestore client")
			http.Error(w, util.ServerErrorMessage, http.StatusInternalServerError)
			return
		}

		tokensDoc, err := firestoreClient.Collection("tokens").Doc(token.UID).Get(r.Context())
		if err != nil {
			log.Errorf("Failed to get token: %s", err)
			http.Error(w, util.ServerErrorMessage, http.StatusInternalServerError)
			return
		}

		var tokensData map[string]interface{}
		if err := tokensDoc.DataTo(&tokensData); err != nil {
			log.Errorf("Failed to unmarshal token data: %s", err)
			http.Error(w, util.ServerErrorMessage, http.StatusInternalServerError)
			return
		}

		twitchTokenData, ok := tokensData["twitch"].(map[string]interface{})
		if !ok {
			http.Error(w, "Missing twitch token", http.StatusBadRequest)
			return
		}

		oauthConfig, err := getOauthConfig(r)
		if err != nil {
			log.Errorf("Failed to get twitch oauth config: %s", err)
			http.Error(w, util.ServerErrorMessage, http.StatusInternalServerError)
			return
		}

		oauthToken := &oauth2.Token{
			AccessToken:  twitchTokenData["AccessToken"].(string),
			RefreshToken: twitchTokenData["RefreshToken"].(string),
			Expiry:       twitchTokenData["Expiry"].(time.Time),
			TokenType:    twitchTokenData["TokenType"].(string),
		}
		oldAccessToken := oauthToken.AccessToken

		tokenSource := oauthConfig.TokenSource(r.Context(), oauthToken)
		oauthToken, err = tokenSource.Token()
		if err != nil {
			log.Errorf("Failed to refresh token: %s", err)
			http.Error(w, util.ServerErrorMessage, http.StatusInternalServerError)
			return
		}

		if oldAccessToken != oauthToken.AccessToken {
			_, err = firestoreClient.Collection("tokens").Doc(token.UID).Set(r.Context(), map[string]interface{}{
				"twitch": oauthToken,
			}, firestore.MergeAll)
			if err != nil {
				log.Errorf("Failed to save token: %s", err)
				http.Error(w, util.ServerErrorMessage, http.StatusInternalServerError)
				return
			}
			log.Infof("Updated token for user %s", token.UID)
		}

		client, err := helix.NewClient(&helix.Options{
			ClientID: twitchClientId,
		})
		if err != nil {
			log.Errorf("Failed to create helix client: %s", err)
			http.Error(w, util.ServerErrorMessage, http.StatusInternalServerError)
			return
		}
		client.SetUserAccessToken(oauthToken.AccessToken)

		ctx := context.WithValue(r.Context(), twitchContextKey, client)
		next(w, r.WithContext(ctx))
	}
}

func GetTwitchClient(ctx context.Context) (*helix.Client, bool) {
	client, ok := ctx.Value(twitchContextKey).(*helix.Client)
	return client, ok
}
