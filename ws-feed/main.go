package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/pkulik0/stredono/pb"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type FeedService struct {
	ctx      context.Context
	client   *pubsub.Client
	upgrader websocket.Upgrader
}

func NewFeedService() (*FeedService, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "stredono-5ccdd")
	if err != nil {
		return nil, err
	}

	log.Debug("Connected to pubsub")

	return &FeedService{
		ctx:    ctx,
		client: client,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}, nil
}

func (s *FeedService) Close() {
	s.client.Close()
}

func (s *FeedService) Serve(addr string) error {
	http.HandleFunc("/ws", s.wsHandler)
	return http.ListenAndServe(addr, nil)
}

const (
	pingPeriod = 50 * time.Second
	pongWait   = 60 * time.Second
	readLimit  = 0
)

func (s *FeedService) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	conn.SetReadLimit(readLimit)

	if err := conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Errorf("Failed to set read deadline: %v", err)
		return
	}
	conn.SetPongHandler(func(string) error {
		if err := conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			return err
		}
		return nil
	})

	sub := s.client.Subscription("donations-XYZ")

	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	go func() {
		log.Debugf("Starting ping ticker for %s", conn.RemoteAddr())
		defer log.Debugf("Stopped ping ticker for %s", conn.RemoteAddr())

		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(pongWait)); err != nil {
					cancel()
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		log.Debugf("Starting read loop for %s", conn.RemoteAddr())
		defer log.Debugf("Stopped read loop for %s", conn.RemoteAddr())

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				cancel()
				return
			}
		}
	}()

	log.Debugf("%s started listening to subscription: %s", conn.RemoteAddr(), sub.String())
	defer log.Debugf("%s unsubscribed from: %s", conn.RemoteAddr(), sub.String())

	log.Debugf("Waiting for message for %s", conn.RemoteAddr())
	err = sub.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
		log.Debugf("Received message: %s for %s", message.ID, conn.RemoteAddr())

		var sdReq pb.SendDonateRequest
		if err := proto.Unmarshal(message.Data, &sdReq); err != nil {
			log.Errorf("Failed to unmarshal message: %v", err)
			return
		}
		log.Debugf("Unmarshaled message: %s", sdReq.String())

		if err := conn.WriteMessage(websocket.BinaryMessage, message.Data); err != nil {
			log.Errorf("Failed to write message to %s: %v", conn.RemoteAddr(), err)
			return
		}
		message.Ack()
	})
	if err != nil {
		log.Errorf("Failed to receive message: %v", err)
	}
}

func main() {
	log.SetLevel(log.DebugLevel)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
		log.Infof("Defaulting to addr: %s", addr)
	}

	fs, err := NewFeedService()
	if err != nil {
		log.Fatalf("Failed to create feed service: %v", err)
	}
	defer fs.Close()

	log.Infof("Starting server on %s", addr)
	if err := fs.Serve(addr); err != nil {
		log.Errorf("Failed to start server: %v", err)
	}
}
