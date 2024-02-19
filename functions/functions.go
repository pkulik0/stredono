package functions

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

const GcSecretsPath = "project/" + ProjectNumber + "/secrets/"

func init() {
	functions.HTTP("OnRegister", OnRegister)

	functions.HTTP("SendTip", SendTip)
	functions.HTTP("ConfirmPayment", ConfirmPayment)
	functions.HTTP("GetListeners", GetListeners)

	functions.HTTP("ConnectTwitch", Connect)
	functions.HTTP("ConnectTwitchCallback", ConnectCallback)
	functions.HTTP("GetTwitchData", GetData)
	functions.HTTP("WebhookTwitch", Webhook)
	functions.HTTP("CreateSubscription", CreateSubscription)
}
