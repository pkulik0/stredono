package adapters

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"github.com/pkulik0/stredono/cloud/platform"
)

type GcpSecretManager struct {
	Client *secretmanager.Client
}

func (gsm *GcpSecretManager) GetSecret(ctx context.Context, name string, version string) ([]byte, error) {
	fullName := "project/" + platform.ProjectNumber + "/secrets/" + name + "/versions/" + version
	eventsubSecretResponse, err := gsm.Client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: fullName,
	})
	if err != nil {
		return nil, err
	}
	return eventsubSecretResponse.Payload.Data, nil
}
