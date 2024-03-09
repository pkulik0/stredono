package adapters

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"hash/crc32"
)

type KeyManagerGoogle struct {
	Client *kms.KeyManagementClient
}

func crc32c(data []byte) uint32 {
	t := crc32.MakeTable(crc32.Castagnoli)
	return crc32.Checksum(data, t)
}

func (k *KeyManagerGoogle) Encrypt(ctx context.Context, keyName string, data []byte) ([]byte, error) {
	req := &kmspb.EncryptRequest{
		Name:            keyName,
		Plaintext:       data,
		PlaintextCrc32C: wrapperspb.Int64(int64(crc32c(data))),
	}

	resp, err := k.Client.Encrypt(ctx, req)
	if err != nil {
		return nil, err
	}

	if int64(crc32c(resp.Ciphertext)) != resp.CiphertextCrc32C.GetValue() {
		return nil, fmt.Errorf("ciphertext response failed CRC32C check")
	}

	return resp.Ciphertext, nil
}

func (k *KeyManagerGoogle) Decrypt(ctx context.Context, keyName string, data []byte) ([]byte, error) {
	req := &kmspb.DecryptRequest{
		Name:             keyName,
		Ciphertext:       data,
		CiphertextCrc32C: wrapperspb.Int64(int64(crc32c(data))),
	}

	resp, err := k.Client.Decrypt(ctx, req)
	if err != nil {
		return nil, err
	}

	if int64(crc32c(resp.Plaintext)) != resp.PlaintextCrc32C.GetValue() {
		return nil, fmt.Errorf("plaintext response failed CRC32C check")
	}

	return resp.Plaintext, nil
}
