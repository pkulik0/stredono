package providers

import (
	"context"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"net/http"
	"strings"
)

func GetHelix(ctx context.Context) (modules.HelixClient, bool) {
	client, ok := ctx.Value(helixCtxKey).(modules.HelixClient)
	return client, ok
}

func GetDocDb(ctx context.Context) (modules.NoSqlDb, bool) {
	db, ok := ctx.Value(docDbCtxKey).(modules.NoSqlDb)
	return db, ok
}

func GetSecretManager(ctx context.Context) (modules.SecretManager, bool) {
	client, ok := ctx.Value(secretsCtxKey).(modules.SecretManager)
	return client, ok
}

func GetPubsub(ctx context.Context) (modules.PubSubClient, bool) {
	client, ok := ctx.Value(pubsubCtxKey).(modules.PubSubClient)
	return client, ok
}

func GetAuth(ctx context.Context) (modules.Auth, bool) {
	authClient, ok := ctx.Value(authCtxKey).(modules.Auth)
	return authClient, ok
}

func GetAuthToken(ctx context.Context, r *http.Request) (modules.Token, modules.Auth, bool) {
	client, ok := GetAuth(ctx)
	if !ok {
		return nil, nil, false
	}

	const authHeaderPrefix = "Bearer "
	header := r.Header.Get("Authorization")
	if !strings.HasPrefix(header, authHeaderPrefix) {
		return nil, nil, false
	}

	token, err := client.VerifyToken(ctx, strings.TrimPrefix(header, authHeaderPrefix))
	if err != nil {
		return nil, nil, false
	}
	return token, client, true
}

// --- Not wrapped ---

//func GetRealtimeDb(ctx context.Context) (*realtime.Client, bool) {
//	db, ok := ctx.Value(RealtimeDbCtxKey).(*realtime.Client)
//	return db, ok
//}
//
//func GetStorage(ctx context.Context) (*storage.Client, bool) {
//	client, ok := ctx.Value(StorageCtxKey).(*storage.Client)
//	return client, ok
//}
//
//func GetMessaging(ctx context.Context) (*messaging.Client, bool) {
//	client, ok := ctx.Value(MessagingCtxKey).(*messaging.Client)
//	return client, ok
//}
