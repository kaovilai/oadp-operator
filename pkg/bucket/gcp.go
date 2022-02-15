package bucket

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/openshift/oadp-operator/api/v1alpha1"
	"google.golang.org/api/option"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/**
In order to create GCP Bucket, we need to create a GCP project.
We expect user to create a GCP project and provide the project ID in config.

*/

type gcpBucketClient struct {
	bucket v1alpha1.CloudStorage
	client client.Client
}

// Return true if bucket got created
func (g gcpBucketClient) Create() (bool, error) {
	sc, err := g.getClient()
	if err != nil {
		return false, err
	}
	defer sc.Close()
	if g.bucket.Name == "" {
		return false, fmt.Errorf("bucket name is empty")
	}
	// Don't check for empty project ID. Defer to API defaults.
	// if g.bucket.Spec.ProjectID == "" {
	// 	return false, fmt.Errorf("project id is empty")
	// }
	// Create bucket 🪣
	err = sc.Bucket(g.bucket.Spec.Name).Create(context.Background(), g.bucket.Spec.ProjectID,
		&storage.BucketAttrs{
			Location: g.bucket.Spec.Region,
			// PublicAccessPrevention: ,
			// StorageClass: ,
			// VersioningEnabled: ,
			Labels: g.bucket.Spec.Tags,
		})
	return err == nil, err
}

// Retusn true if bucket exists
// Return false if bucket does not exist
func (g gcpBucketClient) Exists() (bool, error) {
	sc, err := g.getClient()
	if err != nil {
		return false, err
	}
	defer sc.Close()
	_, err = sc.Bucket(g.bucket.Spec.Name).Attrs(context.Background())
	return err != storage.ErrBucketNotExist, err
}

// Returns true if bucket is deleted
// Returns false if bucket is not deleted
func (g gcpBucketClient) Delete() (bool, error) {
	sc, err := g.getClient()
	if err != nil {
		return false, err
	}
	defer sc.Close()
	err = sc.Bucket(g.bucket.Spec.Name).Delete(context.Background())
	return err == nil, err
}

func (g gcpBucketClient) ForceCredentialRefresh() error {
	//TODO: implement
	return nil
}

// helper function to get GCP Storage client
func (g gcpBucketClient) getClient() (*storage.Client, error) {
	ctx := context.Background()
	cred, err := getCredentialFromCloudStorageSecret(g.client, g.bucket)
	if err != nil {
		return nil, err
	}
	//debug
	credstring := string(cred)
	fmt.Printf("credential: %s\n", credstring)
	storageClient, err := storage.NewClient(ctx, option.WithAPIKey(string(cred)))
	return storageClient, err
}
