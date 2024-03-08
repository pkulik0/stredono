package adapters

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/pkulik0/stredono/cloud/platform/modules"
)

type GcpPubSubClient struct {
	Client *pubsub.Client
}

func (c *GcpPubSubClient) Topic(name string) modules.PubSubTopic {
	topic := c.Client.Topic(name)
	topic.PublishSettings.NumGoroutines = 1 // Given the nature of cloud functions
	return &GcpPubSubTopic{topic: topic}
}

func (c *GcpPubSubClient) Close() error {
	return c.Client.Close()
}

type GcpPubSubTopic struct {
	topic *pubsub.Topic
}

func (t *GcpPubSubTopic) Close() {
	t.topic.Stop()
}

func platformToGcpMessage(msg *modules.PubSubMessage) *pubsub.Message {
	return &pubsub.Message{
		Data:       msg.Data,
		Attributes: msg.Attributes,
	}
}

func (t *GcpPubSubTopic) Publish(ctx context.Context, msg *modules.PubSubMessage) (string, error) {
	result := t.topic.Publish(ctx, platformToGcpMessage(msg))
	id, err := result.Get(ctx)
	if err != nil {
		return "", err
	}
	return id, nil
}
