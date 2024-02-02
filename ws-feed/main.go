package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type FeedService struct {
	ctx      context.Context
	cancel   context.CancelFunc
	client   *pubsub.Client
	upgrader websocket.Upgrader
}

func NewFeedService() (*FeedService, error) {
	ctx, cancel := context.WithCancel(context.Background())
	client, err := pubsub.NewClient(ctx, "stredono-5ccdd")
	if err != nil {
		return nil, err
	}

	log.Debug("Connected to pubsub")

	return &FeedService{
		ctx:    ctx,
		cancel: cancel,
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

func (s *FeedService) handleError(conn *websocket.Conn, err error) {
	if websocket.IsCloseError(err, websocket.CloseGoingAway) {
		log.Debugf("Connection closed by %s", conn)
	}
	s.cancel()
	log.Errorf("Connection %s error: %s", conn.RemoteAddr(), err)
}

func (s *FeedService) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	conn.SetReadLimit(readLimit)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		log.Debugf("Received pong from %s", conn.RemoteAddr())
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	sub := s.client.Subscription("donations-XYZ")

	go func() {
		log.Debugf("Starting ping ticker for %s", conn.RemoteAddr())
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(pongWait)); err != nil {
					s.handleError(conn, err)
					break
				}
				log.Debugf("Sent ping to %s", conn.RemoteAddr())
			case <-s.ctx.Done():
				break
			}
		}
		log.Debugf("Stopped ping ticker for %s", conn.RemoteAddr())
	}()

	go func() {
		log.Debugf("Starting read loop for %s", conn.RemoteAddr())
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				s.handleError(conn, err)
				break
			}
		}
		log.Debugf("Stopped read loop for %s", conn.RemoteAddr())
	}()

	log.Debugf("%s started listening to subscription: %s", conn.RemoteAddr(), sub.String())
	for {
		err := sub.Receive(s.ctx, func(ctx context.Context, message *pubsub.Message) {
			log.Debugf("Received message: %s for %s", message.ID, conn.RemoteAddr())
			conn.WriteMessage(websocket.BinaryMessage, message.Data)
			message.Ack()
		})
		if err != nil {
			s.handleError(conn, err)
			break
		}
	}

	log.Debugf("Unsubscribed from topic: %s", sub.String())
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
