package functions

import (
	"fmt"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	"google.golang.org/protobuf/proto"
)

type EventMessageData struct {
	Message EventPubSubMessage
}

type EventPubSubMessage struct {
	Data        []byte            `json:"data"`
	Attributes  map[string]string `json:"attributes"`
	MessageId   string            `json:"messageId"`
	PublishTime string            `json:"publishTime"`
	OrderingKey string            `json:"orderingKey"`
}

func PublishProto(ctx *providers.Context, message proto.Message, topicName, provider string) (string, error) {
	pubsubClient, ok := ctx.GetPubSub()
	if !ok {
		return "", platform.ErrorMissingContextValue
	}

	topic := pubsubClient.Topic(topicName)
	defer topic.Close()

	data, err := proto.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("failed to marshal event | %v", err)
	}

	msgId, err := topic.Publish(ctx.Ctx, &modules.PubSubMessage{
		Data: data,
		Attributes: map[string]string{
			"provider": provider,
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to publish message | %v", err)
	}

	return msgId, nil
}
