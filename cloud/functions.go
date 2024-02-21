package cloud

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)

	functions.HTTP("OnRegister", OnRegister)

	functions.HTTP("TipSend", TipSend)
	functions.HTTP("TipConfirm", TipConfirm)

	functions.HTTP("GetTwitchData", TwitchGetData)
	functions.HTTP("WebhookTwitch", TwitchWebhook)
	functions.HTTP("CreateSubscription", CreateSubscription)

	log.Println("Cloud functions initialized")
}
