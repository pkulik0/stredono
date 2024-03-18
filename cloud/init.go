package cloud

import (
	cloudfunc "github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/pkulik0/stredono/cloud/functions/events"
	"github.com/pkulik0/stredono/cloud/functions/tips"
	"github.com/pkulik0/stredono/cloud/functions/tts"
	"github.com/pkulik0/stredono/cloud/functions/twitch"
	"github.com/pkulik0/stredono/cloud/functions/twitch/eventsub"
	"github.com/pkulik0/stredono/cloud/functions/user"
	"github.com/pkulik0/stredono/cloud/platform"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func setupEmulators() error {
	if strings.ToLower(os.Getenv("IS_LOCAL")) != "true" {
		return nil
	}

	if err := os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", "localhost:30501"); err != nil {
		return err
	}

	if err := os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:30502"); err != nil {
		return err
	}

	if err := os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:30505"); err != nil {
		return err
	}

	if err := os.Setenv("STORAGE_EMULATOR_HOST", "localhost:30506"); err != nil {
		return err
	}

	log.Info("RunninObject.values(data)g in local mode. Emulators are set up.")
	return nil
}

func init() {
	log.SetLevel(log.DebugLevel)

	if err := setupEmulators(); err != nil {
		log.Fatal(err)
	}

	cloudfunc.CloudEvent("OnEvent", events.OnEventEntrypoint)
	cloudfunc.HTTP("EventChangeState", platform.CorsMiddleware(events.ChangeStateEntrypoint))

	cloudfunc.HTTP("UserRegister", platform.CorsMiddleware(user.RegisterEntrypoint))
	cloudfunc.HTTP("UserEdit", platform.CorsMiddleware(user.EditEntrypoint))
	cloudfunc.HTTP("UserRegenerateKey", platform.CorsMiddleware(user.RegenerateKeyEntrypoint))
	cloudfunc.HTTP("UserModeratorChange", platform.CorsMiddleware(user.ModeratorChangeEntrypoint))

	cloudfunc.HTTP("TipSend", platform.CorsMiddleware(tips.SendEntrypoint))
	cloudfunc.HTTP("TipConfirm", platform.CorsMiddleware(tips.ConfirmEntrypoint))

	cloudfunc.HTTP("TwitchWebhook", platform.CorsMiddleware(twitch.WebhookEntrypoint))
	cloudfunc.HTTP("TwitchEventsubList", platform.CorsMiddleware(eventsub.ListEntrypoint))

	// Below are rarely needed endpoints
	cloudfunc.HTTP("TwitchEventsubInit", platform.CorsMiddleware(eventsub.InitEntrypoint))

	cloudfunc.HTTP("SpeechUpdate", platform.CorsMiddleware(tts.UpdateEntrypoint))
	cloudfunc.HTTP("ChatBotInit", platform.CorsMiddleware(twitch.InitEntrypoint))
}
