package controllers

import (
	// uncomment to use
	// "github.com/noobaa/noobaa-operator/v2/pkg/apis/noobaa/v1alpha1"
	"context"

	"github.com/go-logr/logr"
	nbv1 "github.com/noobaa/noobaa-operator/v5/pkg/apis/noobaa/v1alpha1"
	oadpv1alpha1 "github.com/openshift/oadp-operator/api/v1alpha1"
	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	apiVersion = "noobaa.io/v1alpha1"
	nooBaaKind = "NooBaa"
)

var (
	expected_crd_names = []string{
		"noobaas.noobaa.io",
		"bucketclasses.noobaa.io",
		"backingstores.noobaa.io",
	}
	noobaa_labels = map[string]string{
		"app": "noobaa",
	}
	deletePropagationForeground = metav1.DeletePropagationForeground
)

// func noobaaCrdObjectMeta() metav1.ObjectMeta {
// 	return metav1.ObjectMeta{
// 			Name: ,
// 		}
// }

func (r *VeleroReconciler) getNoobaaObjectMeta() metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:      "noobaa",
		Namespace: r.NamespacedName.Namespace,
	}
}

func (r *VeleroReconciler) getNoobaaBSL() velerov1.BackupStorageLocation {
	return velerov1.BackupStorageLocation{
		ObjectMeta: r.getNoobaaObjectMeta(),
		TypeMeta: metav1.TypeMeta{
			Kind:       "BackupStorageLocation",
			APIVersion: "velero.io/v1",
		},
	}
}

func (r *VeleroReconciler) getOadpOwnedNoobaaSystem() nbv1.NooBaa {
	return nbv1.NooBaa{
		ObjectMeta: r.getNoobaaObjectMeta(),
		TypeMeta: metav1.TypeMeta{
			Kind:       nooBaaKind,
			APIVersion: apiVersion,
		},
	}
}

func (r *VeleroReconciler) ReconcileNoobaa(log logr.Logger) (bool, error) {
	velero := oadpv1alpha1.Velero{}
	deletePropagationBackground := metav1.DeletePropagationForeground
	oadpOwnedNoobaaSystemSelector := nbv1.NooBaa{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: r.NamespacedName.Namespace,
		},
		TypeMeta: metav1.TypeMeta{
			Kind: "NooBaa",
		},
	}
	if !velero.Spec.Noobaa {
		// if not noobaa
		// Deleting existing NooBaa bsl storage
		noobaaBsl := r.getNoobaaBSL()
		err := r.Delete(context.Background(), &noobaaBsl, &client.DeleteOptions{
			PropagationPolicy: &deletePropagationBackground,
		})
		if err != nil {
			r.EventRecorder.Event(&noobaaBsl, corev1.EventTypeWarning, "Error deleting NooBaa BSL", "")
		}
		//Discover existing NooBaa systems owned by OADP operator
		var thisObj *client.Object
		r.Get(r.Context, client.ObjectKeyFromObject(&oadpOwnedNoobaaSystemSelector), *thisObj)
		for _, expected_crd_name := range expected_crd_names {
			existingCrdSelector := oadpOwnedNoobaaSystemSelector
			existingCrdSelector.Name = expected_crd_name
			existingCrdSelector.Labels = noobaa_labels
			var existingCrd *nbv1.NooBaa
			err := r.Get(r.Context, r.NamespacedName, existingCrd)
			if err == nil {
				// Deleting NooBaa Systems owned by OADP operator
				//No errors, therefor existingCrd has the object to delete
				//Delete Object\
				err := r.DeleteAllOf(context.Background(), existingCrd, &client.DeleteAllOfOptions{
					DeleteOptions: client.DeleteOptions{
						PropagationPolicy: &deletePropagationForeground,
					},
				})
				if err != nil {
					// report error
				}
			}
		}

		return true, nil
	}
	if velero.Spec.Noobaa {
		// Check if NooBaa CRDs exist
		// Validate that NooBaa CRDs exist
		// if NooBaa CRD exist: Discover existing NooBaa System or create new
		// Waiting for NooBaa system to become ready
		//   # discover routes and secrets
		// Create Noobaa s3 cloud-credentials secret
		//   # NooBaa Bucket + noobaa bsl creation

	}

	// op, err := controllerutil.CreateOrUpdate(r.Context, r.Client, ds, func() error {

	// })

	// if err != nil {
	// 	return false, err
	// }

	// if op == controllerutil.OperationResultCreated || op == controllerutil.OperationResultUpdated {
	// 	// Trigger event to indicate restic was created or updated
	// 	r.EventRecorder.Event(ds,
	// 		v1.EventTypeNormal,
	// 		"ResticDaemonsetReconciled",
	// 		fmt.Sprintf("performed %s on restic deployment %s/%s", op, ds.Namespace, ds.Name),
	// 	)
	// }

	return true, nil
}
