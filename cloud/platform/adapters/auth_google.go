package adapters

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/pkulik0/stredono/cloud/platform/modules"
)

type FirebaseAuth struct {
	Client *auth.Client
}

func (fa *FirebaseAuth) UpdateUser(ctx context.Context, uid string, data *modules.UserToUpdate) error {
	fData := &auth.UserToUpdate{}

	if data.Email != nil {
		fData = fData.Email(*data.Email)
	}
	if data.IsEmailVerified != nil {
		fData = fData.EmailVerified(*data.IsEmailVerified)
	}

	_, err := fa.Client.UpdateUser(ctx, uid, fData)
	return err
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
