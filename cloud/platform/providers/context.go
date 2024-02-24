package providers

import (
	"cloud.google.com/go/pubsub"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/adapters"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

func CreateContext(ctx context.Context, config *Config) (context.Context, error) {
	conf := &firebase.Config{ProjectID: platform.ProjectId, DatabaseURL: platform.FirebaseDatabaseUrl}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}

	if config.Auth {
		client, err := app.Auth(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, authCtxKey, &adapters.FirebaseAuth{Client: client})
	}

	if config.DocDb {
		client, err := app.Firestore(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, docDbCtxKey, &adapters.FirestoreDatabase{Client: client})
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
		client, err := pubsub.NewClient(ctx, platform.ProjectId)
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

	if config.SecretManager {
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

func CreateHelixContext(ctx context.Context, r *http.Request) (context.Context, error) {
	token, _, ok := GetAuthToken(ctx, r)
	if !ok {
		return nil, platform.ErrorMissingAuthToken
	}
	uid := token.UserId()

	db, ok := GetDocDb(ctx)
	if !ok {
		return nil, platform.ErrorMissingContextValue
	}

	tokensDoc, err := db.Collection("tokens").Doc(uid).Get(ctx)
	if err != nil {
		return nil, err
	}

	var tokensData map[string]interface{}
	if err := tokensDoc.DataTo(&tokensData); err != nil {
		return nil, err
	}

	twitchTokenData, ok := tokensData["twitch"].(map[string]interface{})
	if !ok {
		return nil, platform.ErrorMissingAuthToken
	}

	oauthConfig, err := getTwitchOauth2Config(ctx)
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
		_, err = db.Collection("tokens").Doc(uid).Set(ctx, map[string]interface{}{
			"twitch": oauthToken,
		}, &modules.DbOpts{MergeAll: true})
		if err != nil {
			return nil, err
		}
		log.Infof("Updated token for user %s", uid)
	}

	client, err := helix.NewClient(&helix.Options{
		ClientID: modules.TwitchClientId,
	})
	if err != nil {
		return nil, err
	}
	client.SetUserAccessToken(oauthToken.AccessToken)

	return context.WithValue(ctx, helixCtxKey, client), nil
}
