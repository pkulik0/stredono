package functions

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

const (
	ProjectID     = "stredono-5ccdd"
	DatabaseUrl   = "https://stredono.europe-west1.firebasedatabase.app"
	GcSecretsPath = "projects/621885503876/secrets/"
)

func init() {
	functions.HTTP("onRegister", OnRegister)

	functions.HTTP("sendDonate", SendDonate)
	functions.HTTP("confirmPayment", ConfirmPayment)
	functions.HTTP("getListeners", GetListeners)

	functions.HTTP("connectTwitch", Connect)
	functions.HTTP("connectTwitchCallback", ConnectCallback)
	functions.HTTP("getTwitchData", GetData)
	functions.HTTP("webhookTwitch", Webhook)
	functions.HTTP("createSubscription", CreateSubscription)
}
