package platform

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	realtime "firebase.google.com/go/v4/db"
	"firebase.google.com/go/v4/messaging"
	"firebase.google.com/go/v4/storage"
	"github.com/nicklaw5/helix"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
	"time"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

const (
	authCtxKey       = "auth"
	tokenCtxKey      = "token"
	docDbCtxKey      = "docDb"
	realtimeDbCtxKey = "realtimeDb"
	storageCtxKey    = "storage"
	pubsubCtxKey     = "pubsub"
	messagingCtxKey  = "messaging"
	secretsCtxKey    = "secrets"
	helixCtxKey      = "twitchHelix"
)

func CorsMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

type AuthConfig struct {
	Client      bool
	VerifyToken bool
}

type CloudConfig struct {
	Auth           AuthConfig
	DocDb          bool
	RealtimeDb     bool
	Storage        bool
	PubSub         bool
	Messaging      bool
	SecretsManager bool
}

func createContext(r *http.Request, config *CloudConfig) (context.Context, error) {
	ctx := r.Context()

	conf := &firebase.Config{ProjectID: ProjectId, DatabaseURL: FirebaseDatabaseUrl}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}

	if config.Auth.Client || config.Auth.VerifyToken {
		client, err := app.Auth(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, authCtxKey, client)
	}

	if config.Auth.Client || config.Auth.VerifyToken {
		client, err := app.Auth(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, authCtxKey, client)

		if config.Auth.VerifyToken {
			token := r.Header.Get("Authorization")
			if token == "" {
				return nil, ErrorInvalidAuthHeader
			}

			prefix := "Bearer "
			if strings.HasPrefix(token, prefix) {
				token = strings.TrimPrefix(token, prefix)
			} else {
				return nil, ErrorInvalidAuthHeader
			}

			fbToken, err := client.VerifyIDToken(r.Context(), token)
			if err != nil {
				return nil, err
			}
			ctx = context.WithValue(ctx, tokenCtxKey, fbToken)
		}
	}

	if config.DocDb {
		client, err := app.Firestore(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, docDbCtxKey, client)
		defer func(client *firestore.Client) {
			if err := client.Close(); err != nil {
				log.Errorf("failed to close firestore client | %s", err)
			}
		}(client)
	}

	if config.Storage {
		client, err := app.Storage(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, storageCtxKey, client)
	}

	if config.Messaging {
		client, err := app.Messaging(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, messagingCtxKey, client)
	}

	if config.RealtimeDb {
		client, err := app.Database(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, realtimeDbCtxKey, client)
	}

	if config.PubSub {
		client, err := pubsub.NewClient(r.Context(), ProjectId)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, pubsubCtxKey, client)
		defer func(client *pubsub.Client) {
			if err := client.Close(); err != nil {
				log.Errorf("failed to close pubsub client | %s", err)
			}
		}(client)
	}

	if config.SecretsManager {
		client, err := secretmanager.NewClient(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, secretsCtxKey, client)
		defer func(client *secretmanager.Client) {
			if err := client.Close(); err != nil {
				log.Errorf("failed to close secrets manager client | %s", err)
			}
		}(client)
	}

	return ctx, nil
}

func CloudMiddleware(config *CloudConfig, next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := createContext(r, config)
		if err != nil {
			log.Errorf("failed to create context | %s", err)
			http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
			return
		}
		next(w, r.WithContext(ctx))
	}
}

func createHelixContext(ctx context.Context) (context.Context, error) {
	token, ok := GetAuthToken(ctx)
	if !ok {
		return nil, ErrorMissingContextValue
	}

	db, ok := GetDocDb(ctx)
	if !ok {
		return nil, ErrorMissingContextValue
	}

	tokensDoc, err := db.Collection("tokens").Doc(token.UID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var tokensData map[string]interface{}
	if err := tokensDoc.DataTo(&tokensData); err != nil {
		return nil, err
	}

	twitchTokenData, ok := tokensData["twitch"].(map[string]interface{})
	if !ok {
		return nil, ErrorMissingAuthToken
	}

	oauthConfig, err := GetTwitchOauth2Config(ctx)
	if err != nil {
		return nil, err
	}

	oauthToken := &oauth2.Token{
		AccessToken:  twitchTokenData["AccessToken"].(string),
		RefreshToken: twitchTokenData["RefreshToken"].(string),
		Expiry:       twitchTokenData["Expiry"].(time.Time),
		TokenType:    twitchTokenData["TokenType"].(string),
	}
	oldAccessToken := oauthToken.AccessToken

	tokenSource := oauthConfig.TokenSource(ctx, oauthToken)
	oauthToken, err = tokenSource.Token()
	if err != nil {
		return nil, err
	}

	if oldAccessToken != oauthToken.AccessToken {
		_, err = db.Collection("tokens").Doc(token.UID).Set(ctx, map[string]interface{}{
			"twitch": oauthToken,
		}, &DbOpts{MergeAll: true})
		if err != nil {
			return nil, err
		}
		log.Infof("Updated token for user %s", token.UID)
	}

	client, err := helix.NewClient(&helix.Options{
		ClientID: twitchClientId,
	})
	if err != nil {
		return nil, err
	}
	client.SetUserAccessToken(oauthToken.AccessToken)

	return context.WithValue(ctx, helixCtxKey, client), nil
}

func HelixMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := createHelixContext(r.Context())
		if err != nil {
			log.Errorf("failed to create helix context | %s", err)
			http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
			return
		}
		next(w, r.WithContext(ctx))
	}
}

func GetHelix(ctx context.Context) (HelixClient, bool) {
	client, ok := ctx.Value(helixCtxKey).(HelixClient)
	return client, ok
}

func GetDocDb(ctx context.Context) (NoSqlDb, bool) {
	db, ok := ctx.Value(docDbCtxKey).(NoSqlDb)
	return db, ok
}

func GetSecretManager(ctx context.Context) (SecretManager, bool) {
	client, ok := ctx.Value(secretsCtxKey).(SecretManager)
	return client, ok
}

func GetPubsub(ctx context.Context) (PubSubClient, bool) {
	client, ok := ctx.Value(pubsubCtxKey).(PubSubClient)
	return client, ok
}

// --- Not wrapped ---

func GetAuth(ctx context.Context) (*auth.Client, bool) {
	authClient, ok := ctx.Value(authCtxKey).(*auth.Client)
	return authClient, ok
}

func GetAuthToken(ctx context.Context) (*auth.Token, bool) {
	token, ok := ctx.Value(tokenCtxKey).(*auth.Token)
	return token, ok
}

func GetRealtimeDb(ctx context.Context) (*realtime.Client, bool) {
	db, ok := ctx.Value(realtimeDbCtxKey).(*realtime.Client)
	return db, ok
}

func GetStorage(ctx context.Context) (*storage.Client, bool) {
	client, ok := ctx.Value(storageCtxKey).(*storage.Client)
	return client, ok
}

func GetMessaging(ctx context.Context) (*messaging.Client, bool) {
	client, ok := ctx.Value(messagingCtxKey).(*messaging.Client)
	return client, ok
}
