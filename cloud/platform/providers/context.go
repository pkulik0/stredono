package providers

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/pubsub"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/nicklaw5/helix/v2"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/adapters"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"google.golang.org/protobuf/proto"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Auth          bool
	DocDb         bool
	RealtimeDb    bool
	Storage       bool
	PubSub        bool
	SecretManager bool
	KeyManager    bool
	Proxy         bool
	TextToSpeech  bool
}

type Context struct {
	Ctx           context.Context
	DocDb         modules.DocDb
	Auth          modules.Auth
	Storage       modules.Storage
	RealtimeDb    modules.RealtimeDb
	PubSub        modules.PubSubClient
	TTSPlus       modules.TTS
	TTSBasic      modules.TTS
	SecretManager modules.SecretManager
	KeyManager    modules.KeyManager
}

func (c *Context) GetDocDb() (db modules.DocDb, ok bool) {
	return c.DocDb, c.DocDb != nil
}

func (c *Context) GetRealtimeDb() (db modules.RealtimeDb, ok bool) {
	return c.RealtimeDb, c.RealtimeDb != nil
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

func (c *Context) GetKeyManager() (keyManager modules.KeyManager, ok bool) {
	return c.KeyManager, c.KeyManager != nil
}

const (
	SecretClientSecret = "twitch-client-secret"
	SecretBotToken     = "twitch-bot-token"
)

func GetHelixAppClient(ctx *Context) (modules.HelixClient, error) {
	secretManager, ok := ctx.GetSecretManager()
	if !ok {
		return nil, platform.ErrorMissingContextValue
	}

	clientSecret, err := secretManager.GetSecret(ctx.Ctx, SecretClientSecret, "latest")
	if err != nil {
		return nil, err
	}

	client, err := helix.NewClient(&helix.Options{
		RedirectURI:  "http://localhost:8080/TwitchBotInit", // used only for the bot account (others go through firebase auth)
		ClientID:     platform.TwitchClientId,
		ClientSecret: string(clientSecret),
	})
	if err != nil {
		return nil, err
	}

	// idk why the helix pkg requires scopes, the client credential grant flow doesn't need them
	appToken, err := client.RequestAppAccessToken([]string{})
	if err != nil {
		return nil, err
	}
	client.SetAppAccessToken(appToken.Data.AccessToken)

	return client, nil
}

func GetHelixBotClient(ctx *Context) (modules.HelixClient, error) {
	secretManager, ok := ctx.GetSecretManager()
	if !ok {
		return nil, platform.ErrorMissingContextValue
	}

	clientSecret, err := secretManager.GetSecret(ctx.Ctx, SecretClientSecret, "latest")
	if err != nil {
		return nil, err
	}

	client, err := helix.NewClient(&helix.Options{
		ClientID:     platform.TwitchClientId,
		ClientSecret: string(clientSecret),
	})
	if err != nil {
		return nil, err
	}

	tokenBytes, err := secretManager.GetSecret(ctx.Ctx, SecretBotToken, "latest")
	if err != nil {
		return nil, err
	}

	token := &pb.Token{}
	if err := proto.Unmarshal(tokenBytes, token); err != nil {
		return nil, err
	}

	client.SetRefreshToken(token.RefreshToken)
	client.SetUserAccessToken(token.AccessToken)

	return client, nil
}

func GetHelixClient(ctx *Context) (modules.HelixClient, error) {
	secretManager, ok := ctx.GetSecretManager()
	if !ok {
		return nil, platform.ErrorMissingContextValue
	}

	clientSecret, err := secretManager.GetSecret(ctx.Ctx, SecretClientSecret, "latest")
	if err != nil {
		return nil, err
	}

	client, err := helix.NewClient(&helix.Options{
		ClientID:     platform.TwitchClientId,
		ClientSecret: string(clientSecret),
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetHelixTransport(ctx *Context) (*helix.EventSubTransport, error) {
	secretManager, ok := ctx.GetSecretManager()
	if !ok {
		return nil, platform.ErrorMissingContextValue
	}

	webhookSecret, err := secretManager.GetSecret(ctx.Ctx, "twitch-eventsub-secret", "latest")
	if err != nil {
		return nil, err
	}

	return &helix.EventSubTransport{
		Method:   "webhook",
		Callback: platform.FunctionsUrl + "/TwitchWebhook",
		Secret:   string(webhookSecret),
	}, nil
}

func NewContextEvent(ctx context.Context, config *Config) (*Context, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", "http://localhost", nil)
	if err != nil {
		return nil, err
	}
	return NewContext(r, config)
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

	if config.KeyManager {
		client, err := kms.NewKeyManagementClient(outCtx.Ctx)
		if err != nil {
			return nil, err
		}
		outCtx.KeyManager = &adapters.KeyManagerGoogle{Client: client}
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

	if config.RealtimeDb {
		client, err := app.Database(outCtx.Ctx)
		if err != nil {
			return nil, err
		}
		outCtx.RealtimeDb = &adapters.FirebaseRealtimeDbAdapter{Client: client}
	}

	return outCtx, nil
}

func (c *Context) Close() error {
	if c.PubSub != nil {
		if err := c.PubSub.Close(); err != nil {
			return err
		}
	}

	if c.SecretManager != nil {
		if err := c.SecretManager.Close(); err != nil {
			return err
		}
	}

	return nil
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
