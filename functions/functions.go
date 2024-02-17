package functions

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("OnRegister", OnRegister)

	functions.HTTP("SendDonate", DonateSend)
	functions.HTTP("ConfirmPayment", Confirm)
	functions.HTTP("GetListeners", GetListeners)

	functions.HTTP("ConnectTwitch", Connect)
	functions.HTTP("ConnectTwitchCallback", ConnectCallback)
	functions.HTTP("GetTwitchData", GetData)
	functions.HTTP("WebhookTwitch", Webhook)
	functions.HTTP("CreateSubscription", CreateSubscription)
}
