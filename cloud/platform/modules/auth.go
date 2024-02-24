package modules

import "context"

type Auth interface {
	VerifyToken(ctx context.Context, token string) (Token, error)
}

type Token interface {
	UserId() string
}
