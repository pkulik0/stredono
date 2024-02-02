package functions

import (
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/golang/protobuf/proto"
	"github.com/pkulik0/stredono/functions/pb"
	"io"
	"net/http"
)

func init() {
	functions.HTTP("SendDonate", SendDonate)
	functions.HTTP("ConfirmPayment", ConfirmPayment)
}

func SendDonate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request", http.StatusBadRequest)
		return
	}

	req := pb.SendDonateRequest{}
	if err := proto.Unmarshal(body, &req); err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s: %s donated %s %s: %s", req.Recipient, req.Sender, req.Amount, req.Currency, req.Message)
}

func ConfirmPayment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Confirmed?"))
}
