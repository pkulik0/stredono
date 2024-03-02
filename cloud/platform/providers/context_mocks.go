package providers

import (
	"github.com/pkulik0/stredono/cloud/platform/mocks"
	"testing"
)

func CreateContextMock(config *Config, t *testing.T) *Context {
	c := &Context{}

	if config.Auth {
		c.Auth = mocks.NewMockAuth(t)
	}

	if config.DocDb {
		c.DocDb = mocks.NewMockDocDb(t)
	}

	if config.PubSub {
		c.PubSub = mocks.NewMockPubSubClient(t)
	}

	if config.SecretManager {
		c.SecretManager = mocks.NewMockSecretManager(t)
	}

	if config.TextToSpeech {
		c.TTSPlus = mocks.NewMockTTS(t)
		c.TTSBasic = mocks.NewMockTTS(t)
	}

	if config.Storage {
		c.Storage = mocks.NewMockStorage(t)
	}

	if config.Helix {
		c.Helix = mocks.NewMockHelixClient(t)
	}

	//if config.Messaging {
	//	c.Messaging = mocks.NewMockMessaging(t)
	//}
	//
	//if config.RealtimeDb {
	//	c.RealtimeDb = mocks.NewMockRealtimeDb(t)
	//}

	return c
}
