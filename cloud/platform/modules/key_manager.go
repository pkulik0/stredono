package modules

import "context"

type KeyManager interface {
	Encrypt(ctx context.Context, keyName string, data []byte) ([]byte, error)
	Decrypt(ctx context.Context, keyName string, data []byte) ([]byte, error)
}
