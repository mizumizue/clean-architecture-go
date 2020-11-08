package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"os"
)

func withInitializedClient(ctx context.Context, fn func(fs *firestore.Client) error) error {
	fs, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		return err
	}
	return fn(fs)
}
