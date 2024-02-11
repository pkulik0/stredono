package main

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	realtime "firebase.google.com/go/v4/db"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/pkulik0/stredono/functions"
	log "github.com/sirupsen/logrus"
)

type roomMap map[string]string

type Bot struct {
	client *twitch.Client
	rooms  roomMap
	stopCh chan struct{}

	app             *firebase.App
	firestoreClient *firestore.Client
	realtimeDb      *realtime.Client

	privMsgCh            chan twitch.PrivateMessage
	whisperMsgCh         chan twitch.WhisperMessage
	noticeMsgCh          chan twitch.NoticeMessage
	userNoticeMsgCh      chan twitch.UserNoticeMessage
	userStateMsgCh       chan twitch.UserStateMessage
	globalUserStateMsgCh chan twitch.GlobalUserStateMessage
	namesMsgCh           chan twitch.NamesMessage
}

const bufferLarge = 100
const bufferSmall = 10

func NewBot(ctx context.Context) (*Bot, error) {
	app, err := functions.GetFirebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	realtimeDb, err := app.Database(ctx)
	if err != nil {
		return nil, err
	}

	client := twitch.NewAnonymousClient()

	bot := &Bot{
		client: client,
		rooms:  make(roomMap),
		stopCh: make(chan struct{}),

		app:             app,
		firestoreClient: firestoreClient,
		realtimeDb:      realtimeDb,

		privMsgCh:            make(chan twitch.PrivateMessage, bufferLarge),
		whisperMsgCh:         make(chan twitch.WhisperMessage, bufferSmall),
		noticeMsgCh:          make(chan twitch.NoticeMessage, bufferLarge),
		userNoticeMsgCh:      make(chan twitch.UserNoticeMessage, bufferSmall),
		userStateMsgCh:       make(chan twitch.UserStateMessage, bufferSmall),
		globalUserStateMsgCh: make(chan twitch.GlobalUserStateMessage, bufferSmall),
		namesMsgCh:           make(chan twitch.NamesMessage, bufferSmall),
	}

	client.Capabilities = []string{twitch.CommandsCapability, twitch.TagsCapability, twitch.MembershipCapability}

	client.OnConnect(func() {
		log.Println("Connected to Twitch")
	})

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		select {
		case bot.privMsgCh <- message:
		default:
			log.Warnf("failed to send private message to channel | %s", message)
		}
	})

	client.OnWhisperMessage(func(message twitch.WhisperMessage) {
		select {
		case bot.whisperMsgCh <- message:
		default:
			log.Warnf("failed to send whisper message to channel | %s", message)
		}
	})

	client.OnNoticeMessage(func(message twitch.NoticeMessage) {
		select {
		case bot.noticeMsgCh <- message:
		default:
			log.Warnf("failed to send notice message to channel | %s", message)
		}
	})

	client.OnUserNoticeMessage(func(message twitch.UserNoticeMessage) {
		select {
		case bot.userNoticeMsgCh <- message:
		default:
			log.Warnf("failed to send user notice message to channel | %s", message)
		}
	})

	client.OnUserStateMessage(func(message twitch.UserStateMessage) {
		select {
		case bot.userStateMsgCh <- message:
		default:
			log.Warnf("failed to send user state message to channel | %s", message)
		}
	})

	client.OnGlobalUserStateMessage(func(message twitch.GlobalUserStateMessage) {
		select {
		case bot.globalUserStateMsgCh <- message:
		default:
			log.Warnf("failed to send global user state message to channel | %s", message)
		}
	})

	client.OnNamesMessage(func(message twitch.NamesMessage) {
		select {
		case bot.namesMsgCh <- message:
		default:
			log.Warnf("failed to send names message to channel | %s", message)
		}
	})

	return bot, nil
}

func (b *Bot) startGoroutines() {
	go b.listenForRooms()

	go b.handlePrivMsg()
	go b.handleWhisperMsg()
	go b.handleNoticeMsg()
	go b.handleUserNoticeMsg()
	go b.handleUserStateMsg()
	go b.handleGlobalUserStateMsg()
	go b.handleNamesMsg()
	log.Println("Goroutines started")
}

func (b *Bot) Connect() error {
	b.startGoroutines()
	return b.client.Connect()
}

func (b *Bot) Close() error {
	if err := b.client.Disconnect(); err != nil {
		return err
	}
	close(b.stopCh)
	return nil
}
