package adapters

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/pkulik0/stredono/cloud/platform"
)

type GcpPubSubClient struct {
	client *pubsub.Client
}

func (c *GcpPubSubClient) Topic(name string) platform.PubSubTopic {
	topic := c.client.Topic(name)
	topic.PublishSettings.NumGoroutines = 1 // Given the nature of cloud functions
	return &GcpPubSubTopic{topic: topic}
}

func (c *GcpPubSubClient) Stop() error {
	return c.client.Close()
}

type GcpPubSubTopic struct {
	topic *pubsub.Topic
}

func (t *GcpPubSubTopic) Stop() {
	t.topic.Stop()
}

func platformToGcpMessage(msg *platform.PubSubMessage) *pubsub.Message {
	return &pubsub.Message{
		Data:       msg.Data,
		Attributes: msg.Attributes,
	}
}

func (t *GcpPubSubTopic) Publish(ctx context.Context, msg *platform.PubSubMessage) {
	_ = t.topic.Publish(ctx, platformToGcpMessage(msg))
}
