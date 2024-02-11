package main

import (
	"context"
	log "github.com/sirupsen/logrus"
)

const (
	botCollection = "bot"
	roomsDoc      = "rooms"
)

func (b *Bot) listenForRooms() {
	ctx := context.Background()

	snap := b.firestoreClient.Collection(botCollection).Doc(roomsDoc).Snapshots(ctx)
	defer snap.Stop()

	for {
		doc, err := snap.Next()
		if err != nil {
			log.Errorf("failed to get next snapshot | %s", err)
		}

		for uid, channelI := range doc.Data() {
			channel, ok := channelI.(string)
			if !ok {
				log.Errorf("failed to parse channel | %s", channelI)
				continue
			}
			if _, ok = b.rooms[uid]; ok {
				continue
			}

			log.Printf("Joining channel %s", channel)
			b.rooms[uid] = channel
			b.client.Join(channel)
		}

		for uid, channel := range b.rooms {
			if _, ok := doc.Data()[uid]; ok {
				continue
			}
			log.Printf("Leaving channel %s", channel)
			delete(b.rooms, uid)
			b.client.Depart(channel)
		}

		select {
		case <-b.stopCh:
			return
		default:
		}
	}
}
