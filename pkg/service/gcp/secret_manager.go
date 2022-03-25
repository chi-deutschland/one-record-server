package gcp

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"fmt"
	"github.com/chi-deutschland/one-record-server/pkg/service"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type SecretManagerService struct {
}

func (s *SecretManagerService) GetSecret(
	secret,
	version,
	projectID string,
) (
	[]byte,
	error,
) {

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	defer client.Close()

	name := fmt.Sprintf(
		"projects/%s/secrets/%s/versions/%s",
		projectID,
		secret,
		version)

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %v", err)
	}

	return result.Payload.Data, nil
}

func NewSecretManagerService() (*SecretManagerService, error) {
	var s SecretManagerService
	return &s, nil
}

var _ service.SecretManager = (*SecretManagerService)(nil)
