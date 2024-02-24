package providers

const (
	authCtxKey       = "auth"
	docDbCtxKey      = "docDb"
	realtimeDbCtxKey = "realtimeDb"
	storageCtxKey    = "storage"
	pubsubCtxKey     = "pubsub"
	messagingCtxKey  = "messaging"
	secretsCtxKey    = "secrets"
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
}
