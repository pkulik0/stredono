package adapters

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"fmt"
	"github.com/pkulik0/stredono/cloud/platform"
)

type GcpSecretManager struct {
	Client *secretmanager.Client
}

func (gsm *GcpSecretManager) GetSecret(ctx context.Context, name string, version string) ([]byte, error) {
	response, err := gsm.Client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/%s", platform.ProjectNumber, name, version),
	})
	if err != nil {
		return nil, err
	}
	return response.Payload.Data, nil
}

func (gsm *GcpSecretManager) Close() error {
	return gsm.Client.Close()
}
