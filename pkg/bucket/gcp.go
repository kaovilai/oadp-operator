package bucket

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/openshift/oadp-operator/api/v1alpha1"
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

func (g gcpBucketClient) Create() (bool, error) {
	sc, err := g.getClient()
	if err != nil {
		return false, err
	}
	if g.bucket.Name == "" {
		return false, fmt.Errorf("bucket name is empty")
	}
	if g.bucket.Spec.ProjectID == "" {
		return false, fmt.Errorf("project id is empty")
	}
	// Create bucket 🪣
	err = sc.Bucket(g.bucket.Spec.Name).Create(context.Background(), g.bucket.Spec.ProjectID, nil)
	return err == nil, err
}

func (g gcpBucketClient) Exists() (bool, error) {
	sc, err := g.getClient()
	if err != nil {
		return false, err
	}
	_, err = sc.Bucket(g.bucket.Spec.Name).Attrs(context.Background())
	if err == storage.ErrBucketNotExist {
		return false, nil
	}
	return true, err
}

func (g gcpBucketClient) Delete() (bool, error) {
	sc, err := g.getClient()
	if err != nil {
		return false, err
	}
	err = sc.Bucket(g.bucket.Spec.Name).Delete(context.Background())
	return err == nil, err
}

func (g gcpBucketClient) ForceCredentialRefresh() error {
	//TODO: implement
	return nil
}

func (g gcpBucketClient) getClient() (*storage.Client, error) {
	ctx := context.Background()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return storageClient, nil
}
