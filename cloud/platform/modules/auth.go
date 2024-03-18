package modules

import "context"

type UserToUpdate struct {
	Email           *string
	IsEmailVerified *bool
}

type Auth interface {
	VerifyToken(ctx context.Context, token string) (Token, error)
	UpdateUser(ctx context.Context, uid string, data *UserToUpdate) error
}

type Token interface {
	UserId() string
}
