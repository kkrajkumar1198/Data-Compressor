package initializers

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type storageConnection struct {
	Client *storage.Client
}

var (
	client *storageConnection
	once   sync.Once
)

// GetGCSClient gets singleton object for Google Storage
func GetGCSClient(ctx context.Context) (*storage.Client, error) {
	var clientErr error
	once.Do(func() {
		storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("gcp-keys.json"))
		if err != nil {
			clientErr = fmt.Errorf("failed to create GCS client error: %s", err.Error())
		} else {
			client = &storageConnection{
				Client: storageClient,
			}
		}
	})
	return client.Client, clientErr
}
