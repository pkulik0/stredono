package modules

import "context"

type SecretManager interface {
	GetSecret(ctx context.Context, name string, version string) ([]byte, error)
}
