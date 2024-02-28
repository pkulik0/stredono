package modules

import "context"

type PubSubMessage struct {
	Data       []byte
	Attributes map[string]string
}

type PubSubTopic interface {
	Close()
	Publish(ctx context.Context, msg *PubSubMessage)
}

type PubSubClient interface {
	Topic(name string) PubSubTopic
	Close() error
}
