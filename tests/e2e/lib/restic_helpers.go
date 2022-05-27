package lib

import (
	"context"
	"log"
	"reflect"

	oadpv1alpha1 "github.com/openshift/oadp-operator/api/v1alpha1"
	"github.com/openshift/oadp-operator/controllers"
	"github.com/openshift/oadp-operator/pkg/common"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func HasCorrectNumResticPods(namespace string) wait.ConditionFunc {
	return func() (bool, error) {
		client, err := setUpClient()
		if err != nil {
			return false, err
		}
		resticOptions := metav1.ListOptions{
			FieldSelector: "metadata.name=restic",
		}
		resticDaemeonSet, err := client.AppsV1().DaemonSets(namespace).List(context.TODO(), resticOptions)
		if err != nil {
			return false, err
		}
		var numScheduled int32
		var numDesired int32

		for _, daemonSetInfo := range (*resticDaemeonSet).Items {
			numScheduled = daemonSetInfo.Status.CurrentNumberScheduled
			numDesired = daemonSetInfo.Status.DesiredNumberScheduled
		}
		// check correct num of Restic pods are initialized
		if numScheduled != 0 && numDesired != 0 {
			if numScheduled == numDesired {
				return true, nil
			}
		}
		if numDesired == 0 {
			return true, nil
		}
		return false, err
	}
}

func AreResticDaemonsetUpdatedAndReady(namespace string, desiredSpec corev1.PodSpec) wait.ConditionFunc {
	log.Printf("Checking if Restic daemonset is ready...")
	return func() (bool, error) {
		rds, err := GetResticDaemonSet(namespace, "restic")
		if err != nil {
			return false, err
		}
		if rds.Status.UpdatedNumberScheduled != rds.Status.DesiredNumberScheduled ||
			rds.Status.NumberUnavailable != 0 {
			log.Printf("Restic daemonset is not ready with condition: %v", rds.Status.Conditions)
			return false, nil
		}
		clientset, err := setUpClient()
		if err != nil {
			return false, err
		}
		// at this point we know that Restic daemonset is ready, now we check if it's ready with desired spec.
		labelSelectors := matchLabelsMapToStringArray(rds.Spec.Selector.MatchLabels)
		var podList *corev1.PodList
		if podList, err = clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
			LabelSelector: labelSelectors[0],
		}); err != nil {
			return false, err
		}
		for _, pod := range podList.Items {
			// if pod is owned by Restic daemonset
			for _, owner := range pod.GetOwnerReferences() {
				if owner.UID == rds.UID && !reflect.DeepEqual(pod.Spec, desiredSpec) {
					log.Printf("Pod %s is not ready with desired spec", pod.Name)
					return false, nil
				}
			}
		}

		return true, nil
	}
}

func DoesDaemonSetExists(namespace string, resticName string) wait.ConditionFunc {
	log.Printf("Checking if restic daemonset exists...")
	return func() (bool, error) {
		_, err := GetResticDaemonSet(namespace, resticName)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func GetResticDaemonSet(namespace, resticName string) (*appsv1.DaemonSet, error) {
	clientset, err := setUpClient()
	if err != nil {
		return nil, err
	}
	return clientset.AppsV1().DaemonSets(namespace).Get(context.Background(), resticName, metav1.GetOptions{})
}

// keep for now
func IsResticDaemonsetDeleted(namespace string) wait.ConditionFunc {
	log.Printf("Checking if Restic daemonset has been deleted...")
	return func() (bool, error) {
		_, err := GetResticDaemonSet(namespace, "restic")
		if apierrors.IsNotFound(err) {
			return true, nil
		}
		return false, err
	}
}

func (v *DpaCustomResource) DisableRestic(namespace string, instanceName string) error {
	err := v.SetClient()
	if err != nil {
		return err
	}
	dpa := &oadpv1alpha1.DataProtectionApplication{}
	err = v.Client.Get(context.Background(), client.ObjectKey{
		Namespace: v.Namespace,
		Name:      v.Name,
	}, dpa)
	if err != nil {
		return err
	}
	dpa.Spec.Configuration.Restic.Enable = pointer.Bool(false)

	err = v.Client.Update(context.Background(), dpa)
	if err != nil {
		return err
	}
	log.Printf("spec 'enable_restic' has been updated to false")
	return nil
}

func (v *DpaCustomResource) EnableResticNodeSelector(namespace string, s3Bucket string, credSecretRef string, instanceName string) error {
	err := v.SetClient()
	if err != nil {
		return err
	}
	dpa := &oadpv1alpha1.DataProtectionApplication{}
	err = v.Client.Get(context.Background(), client.ObjectKey{
		Namespace: v.Namespace,
		Name:      v.Name,
	}, dpa)
	if err != nil {
		return err
	}
	nodeSelector := map[string]string{"foo": "bar"}
	dpa.Spec.Configuration.Restic.PodConfig.NodeSelector = nodeSelector

	err = v.Client.Update(context.Background(), dpa)
	if err != nil {
		return err
	}
	log.Printf("spec 'restic_node_selector' has been updated")
	return nil
}

func ResticDaemonSetHasNodeSelector(namespace, key, value string) wait.ConditionFunc {
	return func() (bool, error) {
		client, err := setUpClient()
		if err != nil {
			return false, nil
		}
		ds, err := client.AppsV1().DaemonSets(namespace).Get(context.TODO(), "restic", metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		// verify daemonset has nodeSelector "foo": "bar"
		selector := ds.Spec.Template.Spec.NodeSelector[key]

		if selector == value {
			return true, nil
		}
		return false, err
	}
}

func GetResticDaemonsetList(namespace string) (*appsv1.DaemonSetList, error) {
	client, err := setUpClient()
	if err != nil {
		return nil, err
	}
	registryListOptions := metav1.ListOptions{
		LabelSelector: "component=velero",
	}
	// get pods in the oadp-operator-e2e namespace with label selector
	deploymentList, err := client.AppsV1().DaemonSets(namespace).List(context.TODO(), registryListOptions)
	if err != nil {
		return nil, err
	}
	return deploymentList, nil
}
var mountPropagationToHostContainer = corev1.MountPropagationHostToContainer
func ResticPodSpec(namespace string, opts ...PodSpecOption) (corev1.PodSpec, error) {
	spec := resticPodSpec()
	for _, opt := range opts {
		err := opt(&spec)
		if err != nil {
			return corev1.PodSpec{}, err
		}
	}
	return spec, nil
}

func resticPodSpec() corev1.PodSpec {
	return corev1.PodSpec{
		NodeSelector:       dpa.Spec.Configuration.Restic.PodConfig.NodeSelector,
		ServiceAccountName: common.Velero,
		SecurityContext: &corev1.PodSecurityContext{
			RunAsUser:          pointer.Int64(0),
			SupplementalGroups: dpa.Spec.Configuration.Restic.SupplementalGroups,
		},
		Volumes: []corev1.Volume{
			// Cloud Provider volumes are dynamically added in the for loop below
			{
				Name: controllers.HostPods,
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: "/var/lib/kubelet/pods",
					},
				},
			},
			{
				Name: "scratch",
				VolumeSource: corev1.VolumeSource{
					EmptyDir: &corev1.EmptyDirVolumeSource{},
				},
			},
			{
				Name: "certs",
				VolumeSource: corev1.VolumeSource{
					EmptyDir: &corev1.EmptyDirVolumeSource{},
				},
			},
		},
		Tolerations: dpa.Spec.Configuration.Restic.PodConfig.Tolerations,
		Containers: []corev1.Container{
			{
				Name: common.Restic,
				SecurityContext: &corev1.SecurityContext{
					Privileged: pointer.Bool(true),
				},
				Image:           common.VeleroImage,
				ImagePullPolicy: corev1.PullAlways,
				Command: []string{
					"/velero",
				},
				Args: []string{
					"restic",
					"server",
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:             "host-pods",
						MountPath:        "/host_pods",
						MountPropagation: &mountPropagationToHostContainer,
					},
					{
						Name:      "scratch",
						MountPath: "/scratch",
					},
					{
						Name:      "certs",
						MountPath: "/etc/ssl/certs",
					},
				},
				Resources: corev1.ResourceRequirements{
					Limits: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse("1"),
						corev1.ResourceMemory: resource.MustParse("512Mi"),
					},
					Requests: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse("500m"),
						corev1.ResourceMemory: resource.MustParse("128Mi"),
					},
				},
				Env: []corev1.EnvVar{
					{
						Name: "NODE_NAME",
						ValueFrom: &corev1.EnvVarSource{
							FieldRef: &corev1.ObjectFieldSelector{
								FieldPath: "spec.nodeName",
							},
						},
					},
					{
						Name: "VELERO_NAMESPACE",
						ValueFrom: &corev1.EnvVarSource{
							FieldRef: &corev1.ObjectFieldSelector{
								FieldPath: "metadata.namespace",
							},
						},
					},
					{
						Name:  "VELERO_SCRATCH_DIR",
						Value: "/scratch",
					},
				},
			},
		},
	}
}