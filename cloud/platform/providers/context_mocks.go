package providers

import (
	"context"
	"github.com/pkulik0/stredono/cloud/platform/mocks"
	"testing"
)

func CreateContextMock(ctx context.Context, config *Config, t *testing.T) (context.Context, error) {
	if config.Auth {
		client := mocks.NewMockAuth(t)
		ctx = context.WithValue(ctx, authCtxKey, client)
	}

	if config.DocDb {
		client := mocks.NewMockNoSqlDb(t)
		ctx = context.WithValue(ctx, docDbCtxKey, client)
	}

	//if config.Storage {
	//	client := mocks.NewMockStorage(t)
	//	ctx = context.WithValue(ctx, StorageCtxKey, client)
	//}
	//
	//if config.Messaging {
	//	client := mocks.NewMockMessaging(t)
	//	ctx = context.WithValue(ctx, MessagingCtxKey, client)
	//}
	//
	//if config.RealtimeDb {
	//	client := mocks.NewMockRealtimeDb(t)
	//	ctx = context.WithValue(ctx, RealtimeDbCtxKey, client)
	//}

	if config.PubSub {
		client := mocks.NewMockPubSubClient(t)
		ctx = context.WithValue(ctx, pubsubCtxKey, client)
	}

	if config.SecretManager {
		client := mocks.NewMockSecretManager(t)
		ctx = context.WithValue(ctx, secretsCtxKey, client)
	}

	return ctx, nil
}

func CreateHelixContextMock(ctx context.Context, t *testing.T) (context.Context, error) {
	client := mocks.NewMockHelixClient(t)
	return context.WithValue(ctx, helixCtxKey, client), nil
}
