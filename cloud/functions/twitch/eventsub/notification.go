package eventsub

import (
	"fmt"
	"github.com/nicklaw5/helix/v2"
	"net/http"
	"time"
)

const (
	eventsubMessageHeaderType      = "Twitch-Eventsub-Message-Type"
	eventsubMessageHeaderId        = "Twitch-Eventsub-Message-Id"
	eventsubMessageHeaderTimestamp = "Twitch-Eventsub-Message-Timestamp"
)

type Notification struct {
	Subscription helix.EventSubSubscription `json:"subscription"`
	Event        interface{}                `json:"event"`
	Challenge    string                     `json:"challenge"`
	ID           string
	Timestamp    time.Time
	Type         string
}

func (e *Notification) Fill(headers http.Header) error {
	e.ID = headers.Get(eventsubMessageHeaderId)
	e.Type = headers.Get(eventsubMessageHeaderType)

	timeStr := headers.Get(eventsubMessageHeaderTimestamp)
	if timeStr == "" {
		return fmt.Errorf("missing timestamp")
	}
	timestamp, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp | %v", err)
	}
	e.Timestamp = timestamp

	return nil
}
