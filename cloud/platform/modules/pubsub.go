package modules

import "context"

type PubSubMessage struct {
	Data       []byte
	Attributes map[string]string
}

type PubSubTopic interface {
	Stop()
	Publish(ctx context.Context, msg *PubSubMessage)
}

type PubSubClient interface {
	Topic(name string) PubSubTopic
	Stop() error
}
