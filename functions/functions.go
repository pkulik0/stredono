package functions

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/pkulik0/stredono/functions/donations"
	"github.com/pkulik0/stredono/functions/twitch"
)

func init() {
	functions.HTTP("sendDonate", donations.SendDonate)
	functions.HTTP("confirmPayment", donations.ConfirmPayment)

	functions.HTTP("getListeners", getListeners)

	functions.HTTP("connectTwitch", twitch.Connect)
	functions.HTTP("connectTwitchCallback", twitch.ConnectCallback)
	functions.HTTP("getTwitchData", twitch.GetData)
	functions.HTTP("webhookTwitch", twitch.Webhook)
}
