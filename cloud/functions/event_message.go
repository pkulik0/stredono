package functions

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
