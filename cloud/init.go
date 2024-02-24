package cloud

import (
	cloudfunc "github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/pkulik0/stredono/cloud/functions/tips"
	"github.com/pkulik0/stredono/cloud/functions/twitch"
	"github.com/pkulik0/stredono/cloud/functions/user"
	"github.com/pkulik0/stredono/cloud/platform"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)

	cloudfunc.HTTP("UserRegister", platform.CorsMiddleware(user.RegisterEntrypoint))
	cloudfunc.HTTP("UserEdit", platform.CorsMiddleware(user.EditEntrypoint))

	cloudfunc.HTTP("TipSend", platform.CorsMiddleware(tips.SendEntrypoint))
	cloudfunc.HTTP("TipConfirm", platform.CorsMiddleware(tips.ConfirmEntrypoint))

	cloudfunc.HTTP("TwitchGetData", platform.CorsMiddleware(twitch.GetDataEntrypoint))
	cloudfunc.HTTP("TwitchWebhook", platform.CorsMiddleware(twitch.WebhookEntrypoint))
	cloudfunc.HTTP("TwitchCreateSub", platform.CorsMiddleware(twitch.CreateSubEntrypoint))
}
