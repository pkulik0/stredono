package chat

import (
	"fmt"
	"github.com/pkulik0/stredono/cloud/functions"
	"github.com/pkulik0/stredono/cloud/pb"
	"github.com/pkulik0/stredono/cloud/platform"
	"github.com/pkulik0/stredono/cloud/platform/modules"
	"github.com/pkulik0/stredono/cloud/platform/providers"
	log "github.com/sirupsen/logrus"
	"strings"
)

func isSuperUser(msg *pb.ChatMessage) bool {
	return msg.SenderRole == pb.Role_OWNER || msg.SenderRole == pb.Role_MODERATOR
}

func HandleMessage(ctx *providers.Context, msg *pb.ChatMessage, provider string) error {
	log.Printf("Received chat message: %+v", msg)

	if isCmd := strings.HasPrefix(msg.Message, "!"); !isCmd {
		log.Printf("Not a command")
		return nil
	}

	uid, err := providers.ProviderIdToUid(ctx, provider, msg.ChatID)
	if err != nil {
		return err
	}

	parts := strings.Split(msg.Message, " ")
	cmd := parts[0][1:]
	args := parts[1:]

	switch cmd {
	case "tts":
		return handleCommandTTS(ctx, msg, provider, args)
	case "command":
		fallthrough
	case "cmd":
		return handleCommandCmd(ctx, msg, uid, args)
	case "vr":
		fallthrough
	case "sr":
		fallthrough
	case "mr":
		return handleCommandMediaRequest(ctx, msg, provider, uid, args)
	}

	return handleCustomCommand(ctx, msg, provider, uid, cmd)
}

func validateMediaUrl(mediaUrl string) bool {
	// TODO: check domain etc
	return true
}

const (
	maxMediaQueueLength = 50
)

func handleCommandMediaRequest(ctx *providers.Context, msg *pb.ChatMessage, provider, uid string, args []string) error {
	mrData := &pb.MediaRequest{
		Queue:    make([]*pb.MediaRequest_QueueItem, 0),
		Settings: &pb.MediaRequestSettings{},
	}
	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}
	mediaRef := rtdb.NewRef("Data").Child(uid).Child("Media")
	if err := mediaRef.Get(ctx.Ctx, mrData); err != nil {
		return fmt.Errorf("failed to get media request data | %v", err)
	}

	if len(args) == 0 {
		// return URL to media request page
		return nil
	}

	if mrData.Settings.MinRole.Number() > msg.SenderRole.Number() {
		log.Printf("User role too low for media request")
		return nil
	}

	if len(args) > 1 {
		// Invalid format
		return nil
	}
	arg := args[0]

	if !mrData.IsEnabled && arg != "enable" {
		log.Printf("Media request disabled")
		return nil
	}

	switch arg {
	case "enable":
		if isSuperUser(msg) {
			return mediaRef.Child("IsEnabled").Set(ctx.Ctx, true)
		}
		return nil
	case "disable":
		if isSuperUser(msg) {
			return mediaRef.Child("IsEnabled").Set(ctx.Ctx, false)
		}
		fallthrough // Also pause when disabled
	case "pause":
		if isSuperUser(msg) {
			// stop media right away
			return mediaRef.Child("IsPlaying").Set(ctx.Ctx, false)
		}
		return nil
	case "play":
		if isSuperUser(msg) {
			return mediaRef.Child("IsPlaying").Set(ctx.Ctx, true)
		}
		return nil
	case "skip":
		// skip media
		return mediaRef.Child("CurrentQueueIndex").Transaction(ctx.Ctx, func(node modules.TransactionNode) (interface{}, error) {
			if err := node.Unmarshal(&mrData.CurrentQueueIndex); err != nil {
				return nil, fmt.Errorf("failed to unmarshal current queue index | %v", err)
			}
			mrData.CurrentQueueIndex++
			return mrData.CurrentQueueIndex, nil
		})
	}

	url := arg // just for clarity
	if !validateMediaUrl(url) {
		// Invalid URL
		log.Printf("Invalid media URL: %s", url)
		return nil
	}

	return mediaRef.Transaction(ctx.Ctx, func(node modules.TransactionNode) (interface{}, error) {
		if err := node.Unmarshal(&mrData); err != nil {
			return nil, fmt.Errorf("failed to unmarshal mr data | %v", err)
		}

		if len(mrData.Queue)-int(mrData.CurrentQueueIndex) > maxMediaQueueLength {
			// Queue full
			log.Printf("Queue full")
			return mrData, nil
		}

		newItem := &pb.MediaRequest_QueueItem{
			URL:               url,
			RequesterID:       msg.SenderID,
			RequesterName:     msg.SenderName,
			RequesterProvider: provider,
			Timestamp:         msg.Timestamp,
			IsApproved:        !mrData.Settings.RequireApproval || isSuperUser(msg),
			Progress:          0,
		}
		mrData.Queue = append(mrData.Queue, newItem)

		// Trim old items
		if mrData.CurrentQueueIndex > maxMediaQueueLength {
			mrData.CurrentQueueIndex = mrData.CurrentQueueIndex - (maxMediaQueueLength / 2) // Leave half to still be able to go back
			mrData.Queue = mrData.Queue[mrData.CurrentQueueIndex:]
		}

		return mrData, nil
	})
}

func handleCustomCommand(ctx *providers.Context, msg *pb.ChatMessage, provider, uid, cmd string) error {
	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	commands := make(map[string]string)
	if err := rtdb.NewRef("Data").Child(uid).Child("Commands").Get(ctx.Ctx, &commands); err != nil {
		return fmt.Errorf("failed to get commands | %v", err)
	}

	foundValue := ""
	for customCmd, value := range commands {
		if cmd != customCmd {
			continue
		}

		foundValue = value
	}
	if foundValue == "" {
		// Command not found
		log.Printf("Command not found: %s", cmd)
		return nil
	}

	return Send(ctx, &pb.BotMessage{
		ChatID:  msg.ChatID,
		Message: foundValue,
		Data: &pb.BotMessage_Normal{
			Normal: &pb.BotMessage_NormalData{
				ReplyMessageID: msg.ID,
			},
		},
	}, provider)
}

func handleCommandTTS(ctx *providers.Context, msg *pb.ChatMessage, provider string, args []string) error {
	if !isSuperUser(msg) {
		log.Printf("TTS not allowed")
		return nil
	}

	message := strings.Join(args, " ")
	if message == "" {
		log.Printf("Empty TTS message")
		return nil
	}

	msgId, err := functions.PublishProto(ctx, &pb.Event{
		ID:         msg.ID,
		Type:       pb.EventType_CHAT_TTS,
		Provider:   provider,
		ProviderID: msg.ChatID,
		SenderName: msg.SenderName,
		SenderID:   msg.SenderID,
		Timestamp:  msg.Timestamp,
		Data: map[string]string{
			"Message": message,
		},
	}, "events")
	if err != nil {
		return err
	}
	log.Printf("Published chat message event with msg id %s", msgId)
	return nil
}

func handleCommandCmd(ctx *providers.Context, msg *pb.ChatMessage, uid string, args []string) error {
	if !isSuperUser(msg) {
		log.Printf("Not a super user")
		return nil
	}
	if len(args) == 0 {
		// !cmd
		// TODO: Send URL to commands page
		return nil
	}
	if len(args) < 2 {
		// Invalid format
		return nil
	}

	action := args[0]
	command := args[1]
	args = args[2:]

	rtdb, ok := ctx.GetRealtimeDb()
	if !ok {
		return platform.ErrorMissingContextValue
	}

	cmdRef := rtdb.NewRef("Data").Child(uid).Child("Commands").Child(command)
	switch action {
	case "add":
		// !commands add <command> <value>
		fallthrough
	case "set":
		fallthrough
	case "edit":
		if len(args) < 1 {
			// Invalid format
			log.Printf("Invalid add command format")
			return nil
		}
		value := strings.Join(args, " ")

		if err := cmdRef.Set(ctx.Ctx, value); err != nil {
			return err
		}
		log.Printf("Added command: %s, %s", command, value)
	case "remove":
		// !commands remove <command>
		if len(args) != 0 {
			// Invalid format
			log.Printf("Invalid remove command format")
			return nil
		}
		if err := cmdRef.Delete(ctx.Ctx); err != nil {
			return err
		}
		log.Printf("Removed command: %s", command)
	default:
		// Invalid arg
		log.Printf("Invalid action: %s", action)
	}
	return nil
}
