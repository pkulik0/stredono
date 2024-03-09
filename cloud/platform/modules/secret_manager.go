package modules

import "context"

type SecretManager interface {
	GetSecret(ctx context.Context, name string, version string) ([]byte, error)
	SetSecret(ctx context.Context, name string, secret []byte) error
	Close() error
}
