package util

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/storage"
	"context"
	"firebase.google.com/go/messaging"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	realtime "firebase.google.com/go/v4/db"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	ServerErrorMessage = "Internal server error"
	BadRequestMessage  = "Invalid request"
	invalidStatusError = "Invalid status"
	unauthorizedError  = "Unauthorized"

	authHeader        = "Authorization"
	contextTypeHeader = "Content-Type"
	tokenPrefix       = "Bearer "

	authContextKey      = "fbAuth"
	tokenContextKey     = "fbToken"
	appContextKey       = "fbApp"
	firestoreContextKey = "fbFirestore"
	storageContextKey   = "fbStorage"
	messagingContextKey = "fbMessaging"
	realtimeContextKey  = "fbRealtime"
	pubsubContextKey    = "gcPubsub"
	secretsContextKey   = "gcSecrets"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type AuthConfig struct {
	Client bool
	Token  bool
}

type CloudConfig struct {
	Auth      AuthConfig
	Firestore bool
	Storage   bool
	Messaging bool
	Realtime  bool
	Pubsub    bool
	Secrets   bool
}

func CorsMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", authHeader+", "+contextTypeHeader)
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

func CloudMiddleware(config CloudConfig, next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		conf := &firebase.Config{ProjectID: ProjectID, DatabaseURL: DatabaseUrl}
		app, err := firebase.NewApp(ctx, conf)

		if err != nil {
			log.Errorf("failed to get firebase app | %s", err)
			http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
			return
		}
		ctx = context.WithValue(ctx, appContextKey, app)

		if config.Auth.Client || config.Auth.Token {
			client, err := app.Auth(ctx)
			if err != nil {
				log.Errorf("failed to get firebase auth | %s", err)
				http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, authContextKey, client)

			if config.Auth.Token {
				token := r.Header.Get(authHeader)
				if token == "" {
					log.Errorf("missing authorization header")
					http.Error(w, unauthorizedError, http.StatusUnauthorized)
					return
				}

				if strings.HasPrefix(token, tokenPrefix) {
					token = strings.TrimPrefix(token, tokenPrefix)
				} else {
					log.Errorf("invalid token format")
					http.Error(w, unauthorizedError, http.StatusUnauthorized)
					return
				}

				fbToken, err := client.VerifyIDToken(r.Context(), token)
				if err != nil {
					log.Errorf("failed to verify token | %s", err)
					http.Error(w, unauthorizedError, http.StatusUnauthorized)
					return
				}
				ctx = context.WithValue(ctx, tokenContextKey, fbToken)
			}
		}

		if config.Firestore {
			client, err := app.Firestore(ctx)
			if err != nil {
				log.Errorf("failed to get firestore client | %s", err)
				http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, firestoreContextKey, client)
			defer func(client *firestore.Client) {
				if err := client.Close(); err != nil {
					log.Errorf("failed to close firestore client | %s", err)
				}
			}(client)
		}

		if config.Storage {
			client, err := app.Storage(ctx)
			if err != nil {
				log.Errorf("failed to get storage client | %s", err)
				http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, storageContextKey, client)
		}

		if config.Messaging {
			client, err := app.Messaging(ctx)
			if err != nil {
				log.Errorf("failed to get messaging client | %s", err)
				http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, messagingContextKey, client)
		}

		if config.Realtime {
			client, err := app.Database(ctx)
			if err != nil {
				log.Errorf("failed to get realtime database client | %s", err)
				http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, realtimeContextKey, client)
		}

		if config.Pubsub {
			client, err := pubsub.NewClient(r.Context(), ProjectID)
			if err != nil {
				log.Errorf("failed to get pubsub client | %s", err)
				http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, pubsubContextKey, client)
			defer func(client *pubsub.Client) {
				if err := client.Close(); err != nil {
					log.Errorf("failed to close pubsub client | %s", err)
				}
			}(client)
		}

		if config.Secrets {
			client, err := secretmanager.NewClient(ctx)
			if err != nil {
				log.Errorf("failed to get secrets manager client | %s", err)
				http.Error(w, ServerErrorMessage, http.StatusInternalServerError)
				return
			}
			ctx = context.WithValue(ctx, secretsContextKey, client)
			defer func(client *secretmanager.Client) {
				if err := client.Close(); err != nil {
					log.Errorf("failed to close secrets manager client | %s", err)
				}
			}(client)
		}

		next(w, r.WithContext(ctx))
	}
}

func GetFirebaseAuth(ctx context.Context) (*auth.Client, bool) {
	authClient, ok := ctx.Value(authContextKey).(*auth.Client)
	return authClient, ok
}

func GetFirestore(ctx context.Context) (*firestore.Client, bool) {
	client, ok := ctx.Value(firestoreContextKey).(*firestore.Client)
	return client, ok
}

func GetFirebaseApp(ctx context.Context) (*firebase.App, bool) {
	app, ok := ctx.Value(appContextKey).(*firebase.App)
	return app, ok
}

func GetFirebaseRealtimeDb(ctx context.Context) (*realtime.Client, bool) {
	db, ok := ctx.Value(realtimeContextKey).(*realtime.Client)
	return db, ok
}

func GetFirebaseStorage(ctx context.Context) (*storage.Client, bool) {
	client, ok := ctx.Value(storageContextKey).(*storage.Client)
	return client, ok
}

func GetFirebaseMessaging(ctx context.Context) (*messaging.Client, bool) {
	client, ok := ctx.Value(messagingContextKey).(*messaging.Client)
	return client, ok
}

func GetPubsub(ctx context.Context) (*pubsub.Client, bool) {
	client, ok := ctx.Value(pubsubContextKey).(*pubsub.Client)
	return client, ok
}

func GetAuthToken(ctx context.Context) (*auth.Token, bool) {
	token, ok := ctx.Value(tokenContextKey).(*auth.Token)
	return token, ok
}

func GetSecretsManager(ctx context.Context) (*secretmanager.Client, bool) {
	client, ok := ctx.Value(secretsContextKey).(*secretmanager.Client)
	return client, ok
}
