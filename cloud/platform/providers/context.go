package providers

import (
	"cloud.google.com/go/pubsub"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/nicklaw5/helix"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/adapters"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	authCtxKey       = "auth"
	docDbCtxKey      = "docDb"
	realtimeDbCtxKey = "realtimeDb"
	storageCtxKey    = "storage"
	pubsubCtxKey     = "pubsub"
	messagingCtxKey  = "messaging"
	secretsCtxKey    = "secrets"
	ttsCtxKey        = "tts"
	helixCtxKey      = "twitchHelix"
)

type Config struct {
	Auth          bool
	DocDb         bool
	RealtimeDb    bool
	Storage       bool
	PubSub        bool
	Messaging     bool
	SecretManager bool
	Proxy         bool
	TextToSpeech  bool
	Helix         bool
}

type Context struct {
	Ctx           context.Context
	DocDb         modules.DocDb
	Auth          modules.Auth
	Storage       modules.Storage
	PubSub        modules.PubSubClient
	TTSPlus       modules.TTS
	TTSBasic      modules.TTS
	SecretManager modules.SecretManager
	Helix         modules.HelixClient
}

func (c *Context) WithValue(key, value interface{}) *Context {
	c.Ctx = context.WithValue(c.Ctx, key, value)
	return c
}

func (c *Context) GetDocDb() (db modules.DocDb, ok bool) {
	return c.DocDb, c.DocDb != nil
}

func (c *Context) GetAuth() (auth modules.Auth, ok bool) {
	return c.Auth, c.Auth != nil
}

func (c *Context) GetStorage() (storage modules.Storage, ok bool) {
	return c.Storage, c.Storage != nil
}

func (c *Context) GetPubSub() (pubsub modules.PubSubClient, ok bool) {
	return c.PubSub, c.PubSub != nil
}

func (c *Context) GetTTSPlus() (tts modules.TTS, ok bool) {
	return c.TTSPlus, c.TTSPlus != nil
}

func (c *Context) GetTTSBasic() (tts modules.TTS, ok bool) {
	return c.TTSBasic, c.TTSBasic != nil
}

func (c *Context) GetSecretManager() (secretManager modules.SecretManager, ok bool) {
	return c.SecretManager, c.SecretManager != nil
}

func (c *Context) GetHelix() (helix modules.HelixClient, ok bool) {
	return c.Helix, c.Helix != nil
}

func NewContext(r *http.Request, config *Config) (*Context, error) {
	outCtx := &Context{Ctx: r.Context()}

	conf := &firebase.Config{ProjectID: platform.ProjectId, DatabaseURL: platform.FirebaseDatabaseUrl, StorageBucket: platform.FirebaseStorageBucket}
	app, err := firebase.NewApp(outCtx.Ctx, conf)
	if err != nil {
		return nil, err
	}

	if config.Auth {
		client, err := app.Auth(outCtx.Ctx)
		if err != nil {
			return nil, err
		}
		outCtx.Auth = &adapters.FirebaseAuth{Client: client}
	}

	if config.DocDb {
		client, err := app.Firestore(outCtx.Ctx)
		if err != nil {
			return nil, err
		}
		outCtx.DocDb = &adapters.FirestoreDatabase{Client: client}
	}

	if config.PubSub {
		client, err := pubsub.NewClient(outCtx.Ctx, platform.ProjectId)
		if err != nil {
			return nil, err
		}
		outCtx.PubSub = &adapters.GcpPubSubClient{Client: client}
	}

	if config.SecretManager {
		client, err := secretmanager.NewClient(outCtx.Ctx)
		if err != nil {
			return nil, err
		}
		outCtx.SecretManager = &adapters.GcpSecretManager{Client: client}
	}

	if config.Proxy {
		if outCtx.SecretManager == nil {
			return nil, platform.ErrorMissingModuleDep
		}

		proxyAddr, err := outCtx.SecretManager.GetSecret(outCtx.Ctx, "proxy", "latest")
		if err != nil {
			return nil, err
		}

		if err := os.Setenv("HTTPS_PROXY", string(proxyAddr)); err != nil {
			return nil, err
		}
	}

	if config.TextToSpeech {
		client, err := texttospeech.NewClient(outCtx.Ctx)
		if err != nil {
			return nil, err
		}
		outCtx.TTSBasic = &adapters.GoogleTTS{Client: client}

		if outCtx.DocDb == nil {
			return nil, platform.ErrorMissingModuleDep
		}
		if outCtx.SecretManager == nil {
			return nil, platform.ErrorMissingModuleDep
		}

		outCtx.TTSPlus, err = adapters.NewElevenLabs(outCtx.DocDb, "elevenlabs-keys")
		if err != nil {
			return nil, err
		}
	}

	if config.Storage {
		client, err := app.Storage(outCtx.Ctx)
		if err != nil {
			return nil, err
		}
		outCtx.Storage = &adapters.FirebaseStorage{Client: client}
	}

	if config.Helix {
		token, ok := outCtx.GetAuthToken(r)
		if !ok {
			return nil, platform.ErrorMissingAuthToken
		}
		uid := token.UserId()

		tokensDoc, err := outCtx.DocDb.Collection("tokens").Doc(uid).Get(outCtx.Ctx)
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

		oauthConfig, err := outCtx.getTwitchOauth2Config()
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

		tokenSource := oauthConfig.TokenSource(outCtx.Ctx, oauthToken)
		oauthToken, err = tokenSource.Token()
		if err != nil {
			return nil, err
		}

		if oldAccessToken != oauthToken.AccessToken {
			_, err = outCtx.DocDb.Collection("tokens").Doc(uid).Set(outCtx.Ctx, map[string]interface{}{
				"twitch": oauthToken,
			}, modules.DbOpts{MergeAll: true})
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

		outCtx.Helix = client
	}

	//if config.Messaging {
	//	client, err := app.Messaging(ctx)
	//	if err != nil {
	//		return nil, err
	//	}
	//	ctx = context.WithValue(ctx, messagingCtxKey, client)
	//}
	//
	//if config.RealtimeDb {
	//	client, err := app.Database(ctx)
	//	if err != nil {
	//		return nil, err
	//	}
	//	ctx = context.WithValue(ctx, realtimeDbCtxKey, client)
	//}

	return outCtx, nil
}

func (c *Context) GetAuthToken(r *http.Request) (modules.Token, bool) {
	const authHeaderPrefix = "Bearer "
	header := r.Header.Get("Authorization")
	if !strings.HasPrefix(header, authHeaderPrefix) {
		return nil, false
	}

	token, err := c.Auth.VerifyToken(c.Ctx, strings.TrimPrefix(header, authHeaderPrefix))
	if err != nil {
		return nil, false
	}
	return token, true
}
