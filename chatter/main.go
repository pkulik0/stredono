package main

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	bot, err := NewBot(ctx)
	if err != nil {
		log.Fatalf("failed to create bot | %s", err)
	}

	log.Println("Connecting to Twitch")
	if err := bot.Connect(); err != nil {
		log.Fatalf("bot failed to connect | %s", err)
	}
	defer bot.Close()

	<-bot.stopCh
}
