package initializers

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	gcsClient     *storage.Client
	gcsClientErr  error
	gcsClientOnce sync.Once
)

const (
	// GCSBucket name
	GCSBucket = "zocket-source-imgs"
	// ProjectID Google Project ID name
	ProjectID = "zocket-400012"
	Delimitor = "_"
)

// GetGCSClient returns a singleton instance of the Google Cloud Storage client.
func GetGCSClient() (*storage.Client, error) {

	gcsClientOnce.Do(func() {

		ctx := context.Background()
		CredentialFile := "gcp-keys.json"
		gcsClient, gcsClientErr = storage.NewClient(ctx, option.WithCredentialsFile(CredentialFile))
		if gcsClientErr != nil {
			log.Printf("Failed to create GCS client: %v", gcsClientErr)
			return
		}

	})

	return gcsClient, gcsClientErr
}
