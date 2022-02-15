package bucket

// import (
// 	"context"
// 	"fmt"

// 	"github.com/openshift/oadp-operator/api/v1alpha1"
// 	"sigs.k8s.io/controller-runtime/pkg/client"
// )

// func getDataProtectionApplication(namespace string, c client.Client) (*v1alpha1.DataProtectionApplication, error) {
// 	dpaList := &v1alpha1.DataProtectionApplicationList{}
// 	err := c.List(context.TODO(), dpaList, client.InNamespace(namespace))
// 	if err != nil {
// 		//cannot retrieve dpa list
// 		return nil, err
// 	}
// 	if len (dpaList.Items) > 1 {
// 		//cannot create bucket with more than one dpa
// 		return nil, fmt.Errorf("cannot create bucket with more than one dpa")
// 	} else if len (dpaList.Items) == 0 {
// 		//cannot create bucket without dpa
// 		return nil, fmt.Errorf("cannot create bucket without dpa")
// 	}
// 	return &dpaList.Items[0], nil
// }

// func getConfigForBucket(bucket v1alpha1.CloudStorage, client client.Client) (*map[string]string, error) {
// 	dpa, err := getDataProtectionApplication(bucket.Namespace, client)
// 	if err != nil {
// 		// error getting dpa
// 		// cannot create bucket without dpa
// 		return nil, err
// 	}

// 	for _, cloudStorageLocationFromList := range dpa.Spec.BackupLocations {
// 		if cloudStorageLocationFromList.Velero != nil {
// 			//Velero not nil, CloudStorage won't be processed 🌧
// 			continue
// 		}
// 		if cloudStorageLocationFromList.CloudStorage != nil && cloudStorageLocationFromList.CloudStorage.CloudStorageRef.Name == bucket.Name{
// 			// CloudStorageLocation is referring to this CloudStorage ☁️=☁️
// 			// get Config from CloudStorageLocation
// 			return &cloudStorageLocationFromList.CloudStorage.Config, nil
// 		//do something with cloudStorageLocationFromList
// 		}
// 	}
// 	return nil, fmt.Errorf("cannot find cloudStorageLocation for bucket")
// }
