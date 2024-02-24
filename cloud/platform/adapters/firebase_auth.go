package adapters

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/pkulik0/stredono/cloud/platform/modules"
)

type FirebaseAuth struct {
	Client *auth.Client
}

func (fa *FirebaseAuth) VerifyToken(ctx context.Context, tokenStr string) (modules.Token, error) {
	token, err := fa.Client.VerifyIDToken(ctx, tokenStr)
	return &FirebaseToken{Token: token}, err
}

type FirebaseToken struct {
	Token *auth.Token
}

func (ft *FirebaseToken) UserId() string {
	return ft.Token.UID
}
