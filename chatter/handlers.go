package main

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func (b *Bot) handlePrivMsg() {
	startTime := time.Now()
	var count int64
	count = 0

	for {
		select {
		case msg := <-b.privMsgCh:
			count++
			msgsPerMinute := float64(count) / time.Since(startTime).Minutes()
			if count%10 == 0 {
				log.Infof("count: %d | pm: %f | time: %s", count, msgsPerMinute, time.Since(startTime))
			}
			_ = msg
		case <-b.stopCh:
			return
		}
	}
}

func (b *Bot) handleWhisperMsg() {
	for {
		select {
		case msg := <-b.whisperMsgCh:
			_ = msg
		case <-b.stopCh:
			return
		}
	}
}

func (b *Bot) handleNoticeMsg() {
	for {
		select {
		case msg := <-b.noticeMsgCh:
			_ = msg
		case <-b.stopCh:
			return
		}
	}
}

func (b *Bot) handleUserNoticeMsg() {
	for {
		select {
		case msg := <-b.userNoticeMsgCh:
			_ = msg
		case <-b.stopCh:
			return
		}
	}
}

func (b *Bot) handleUserStateMsg() {
	for {
		select {
		case msg := <-b.userStateMsgCh:
			_ = msg
		case <-b.stopCh:
			return
		}
	}
}

func (b *Bot) handleGlobalUserStateMsg() {
	for {
		select {
		case msg := <-b.globalUserStateMsgCh:
			_ = msg
		case <-b.stopCh:
			return
		}
	}
}

func (b *Bot) handleNamesMsg() {
	for {
		select {
		case msg := <-b.namesMsgCh:
			_ = msg
			//func() {
			//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			//	defer cancel()
			//
			//	path := fmt.Sprintf("online/%s/%d", msg.Channel, time.Now().Unix())
			//
			//	if err := b.realtimeDb.NewRef(path).Set(ctx, msg.Users); err != nil {
			//		log.Errorf("failed to set names message | %s", err)
			//	}
			//}()
		case <-b.stopCh:
			return
		}
	}
}
