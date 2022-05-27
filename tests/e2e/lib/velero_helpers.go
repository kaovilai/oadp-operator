package lib

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	snapshotv1beta1api "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1beta1"
	snapshotv1beta1client "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	"github.com/openshift/oadp-operator/pkg/common"
	velero "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	pkgbackup "github.com/vmware-tanzu/velero/pkg/backup"
	"github.com/vmware-tanzu/velero/pkg/cmd"
	"github.com/vmware-tanzu/velero/pkg/cmd/util/downloadrequest"
	"github.com/vmware-tanzu/velero/pkg/cmd/util/output"
	"github.com/vmware-tanzu/velero/pkg/features"
	veleroClientset "github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned"
	"github.com/vmware-tanzu/velero/pkg/label"
	"github.com/vmware-tanzu/velero/pkg/restic"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetVeleroClient() (veleroClientset.Interface, error) {
	if vc, err := veleroClientset.NewForConfig(getKubeConfig()); err == nil {
		return vc, nil
	} else {
		return nil, err
	}
}

// https://github.com/vmware-tanzu/velero/blob/11bfe82342c9f54c63f40d3e97313ce763b446f2/pkg/cmd/cli/backup/describe.go#L77-L111
func DescribeBackup(ocClient client.Client, backup velero.Backup) (backupDescription string) {
	err := ocClient.Get(context.Background(), client.ObjectKey{
		Namespace: backup.Namespace,
		Name:      backup.Name,
	}, &backup)
	if err != nil {
		return "could not get provided backup: " + err.Error()
	}
	veleroClient, err := GetVeleroClient()
	if err != nil {
		return err.Error()
	}
	details := true
	insecureSkipTLSVerify := true
	caCertFile := ""

	deleteRequestListOptions := pkgbackup.NewDeleteBackupRequestListOptions(backup.Name, string(backup.UID))
	deleteRequestList, err := veleroClient.VeleroV1().DeleteBackupRequests(backup.Namespace).List(context.TODO(), deleteRequestListOptions)
	if err != nil {
		log.Printf("error getting DeleteBackupRequests for backup %s: %v\n", backup.Name, err)
	}

	opts := label.NewListOptionsForBackup(backup.Name)
	podVolumeBackupList, err := veleroClient.VeleroV1().PodVolumeBackups(backup.Namespace).List(context.TODO(), opts)
	if err != nil {
		log.Printf("error getting PodVolumeBackups for backup %s: %v\n", backup.Name, err)
	}

	var csiClient *snapshotv1beta1client.Clientset
	// declare vscList up here since it may be empty and we'll pass the empty Items field into DescribeBackup
	vscList := new(snapshotv1beta1api.VolumeSnapshotContentList)
	if features.IsEnabled(velero.CSIFeatureFlag) {
		clientConfig := getKubeConfig()

		csiClient, err = snapshotv1beta1client.NewForConfig(clientConfig)
		cmd.CheckError(err)

		vscList, err = csiClient.SnapshotV1beta1().VolumeSnapshotContents().List(context.TODO(), opts)
		if err != nil {
			log.Printf("error getting VolumeSnapshotContent objects for backup %s: %v\n", backup.Name, err)
		}
	}
	// output.DescribeBackup is a helper function from velero CLI that attempts to download logs for a backup.
	// if a backup failed, this function may panic. Recover from the panic and return string of backup object
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in DescribeBackup: %v\n", r)
			log.Print("returning backup object instead")
			backupDescription = fmt.Sprint(backup)
		}
	}()
	return output.DescribeBackup(context.Background(), ocClient, &backup, deleteRequestList.Items, podVolumeBackupList.Items, vscList.Items, details, veleroClient, insecureSkipTLSVerify, caCertFile)
}

// https://github.com/vmware-tanzu/velero/blob/11bfe82342c9f54c63f40d3e97313ce763b446f2/pkg/cmd/cli/restore/describe.go#L72-L78
func DescribeRestore(ocClient client.Client, restore velero.Restore) string {
	err := ocClient.Get(context.Background(), client.ObjectKey{
		Namespace: restore.Namespace,
		Name:      restore.Name,
	}, &restore)
	if err != nil {
		return "could not get provided backup: " + err.Error()
	}
	veleroClient, err := GetVeleroClient()
	if err != nil {
		return err.Error()
	}
	details := true
	insecureSkipTLSVerify := true
	caCertFile := ""
	opts := restic.NewPodVolumeRestoreListOptions(restore.Name)
	podvolumeRestoreList, err := veleroClient.VeleroV1().PodVolumeRestores(restore.Namespace).List(context.TODO(), opts)
	if err != nil {
		log.Printf("error getting PodVolumeRestores for restore %s: %v\n", restore.Name, err)
	}

	return output.DescribeRestore(context.Background(), ocClient, &restore, podvolumeRestoreList.Items, details, veleroClient, insecureSkipTLSVerify, caCertFile)
}

func BackupLogs(ocClient client.Client, backup velero.Backup) (backupLogs string) {
	insecureSkipTLSVerify := true
	caCertFile := ""
	// new io.Writer that store the logs in a string
	logs := &bytes.Buffer{}
	// new io.Writer that store the logs in a string
	// if a backup failed, this function may panic. Recover from the panic and return container logs
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in BackupLogs: %v\n", r)
			log.Print("returning container logs instead")
			var err error
			backupLogs, err = getVeleroContainerLogs(backup.Namespace)
			if err != nil {
				log.Printf("error getting container logs: %v\n", err)
			}
		}
	}()
	downloadrequest.Stream(context.Background(), ocClient, backup.Namespace, backup.Name, velero.DownloadTargetKindBackupLog, logs, time.Minute, insecureSkipTLSVerify, caCertFile)

	return logs.String()
}

func RestoreLogs(ocClient client.Client, restore velero.Restore) (restoreLogs string) {
	insecureSkipTLSVerify := true
	caCertFile := ""
	// new io.Writer that store the logs in a string
	logs := &bytes.Buffer{}
	// new io.Writer that store the logs in a string
	// if a backup failed, this function may panic. Recover from the panic and return container logs
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in BackupLogs: %v\n", r)
			log.Print("returning container logs instead")
			var err error
			restoreLogs, err = getVeleroContainerLogs(restore.Namespace)
			if err != nil {
				log.Printf("error getting container logs: %v\n", err)
			}
		}
	}()
	downloadrequest.Stream(context.Background(), ocClient, restore.Namespace, restore.Name, velero.DownloadTargetKindRestoreLog, logs, time.Minute, insecureSkipTLSVerify, caCertFile)

	return logs.String()
}

func BackupErrorLogs(ocClient client.Client, backup velero.Backup) []string {
	bl := BackupLogs(ocClient, backup)
	errorRegex, err := regexp.Compile("error|Error")
	if err != nil {
		return []string{"could not compile regex: ", err.Error()}
	}
	logLines := []string{}
	for _, line := range strings.Split(bl, "\n") {
		if errorRegex.MatchString(line) {
			logLines = append(logLines, line)
		}
	}
	return logLines
}

func RestoreErrorLogs(ocClient client.Client, restore velero.Restore) []string {
	rl := RestoreLogs(ocClient, restore)
	errorRegex, err := regexp.Compile("error|Error")
	if err != nil {
		return []string{"could not compile regex: ", err.Error()}
	}
	logLines := []string{}
	for _, line := range strings.Split(rl, "\n") {
		if errorRegex.MatchString(line) {
			logLines = append(logLines, line)
		}
	}
	return logLines
}

func BackupStorageLocationIsAvailable(ocClient client.Client, bslName, namespace string) (bool, error) {
	var bsl velero.BackupStorageLocation
	err := ocClient.Get(context.Background(), client.ObjectKey{
		Namespace: namespace,
		Name:      bslName,
	}, &bsl)
	if err != nil {
		return false, err
	}
	return bsl.Status.Phase == velero.BackupStorageLocationPhaseAvailable, nil
}

func GetBackupStorageLocation(ocClient client.Client, bslName, namespace string) (velero.BackupStorageLocation, error) {
	var bsl velero.BackupStorageLocation
	err := ocClient.Get(context.Background(), client.ObjectKey{
		Namespace: namespace,
		Name:      bslName,
	}, &bsl)
	if err != nil {
		return bsl, err
	}
	return bsl, nil
}

func CreateBackupStorageLocation(ocClient client.Client, backupStorageLocation velero.BackupStorageLocation) error {
	veleroClient, err := GetVeleroClient()
	if err != nil {
		return err
	}
	_, err = veleroClient.VeleroV1().BackupStorageLocations(backupStorageLocation.Namespace).Create(context.TODO(), &backupStorageLocation, metav1.CreateOptions{})
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			return nil
		}
		return err
	}
	return nil
}

func DeleteBackupStorageLocation(ocClient client.Client, backupStorageLocation velero.BackupStorageLocation) error {
	veleroClient, err := GetVeleroClient()
	if err != nil {
		return err
	}
	err = veleroClient.VeleroV1().BackupStorageLocations(backupStorageLocation.Namespace).Delete(context.TODO(), backupStorageLocation.Name, metav1.DeleteOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return err
	}
	return nil
}

func GetVeleroDeployment(namespace string) (*appsv1.Deployment, error) {
	client, err := setUpClient()
	if err != nil {
		return nil, err
	}
	// get pods in the oadp-operator-e2e namespace with label selector
	deployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), "velero", metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

func VeleroPodSpec(namespace string, opts ...PodSpecOption) (corev1.PodSpec, error) {
	podSpec := veleroPodSpec()
	for _, opt := range opts {
		err := opt(&podSpec)
		if err != nil {
			return corev1.PodSpec{}, err
		}
	}
	return podSpec, nil
}

func veleroPodSpec() corev1.PodSpec{
	return corev1.PodSpec{
							RestartPolicy:      corev1.RestartPolicyAlways,
							ServiceAccountName: common.Velero,
							Containers: []corev1.Container{
								{
									Name:            common.Velero,
									Image:           common.VeleroImage,
									ImagePullPolicy: corev1.PullAlways,
									Ports: []corev1.ContainerPort{
										{
											Name:          "metrics",
											ContainerPort: 8085,
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
									Command: []string{"/velero"},
									Args: []string{
										"server",
										"--restic-timeout=1h",
									},
									VolumeMounts: []corev1.VolumeMount{
										{
											Name:      "plugins",
											MountPath: "/plugins",
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
									Env: []corev1.EnvVar{
										{
											Name:  common.VeleroScratchDirEnvKey,
											Value: "/scratch",
										},
										{
											Name: common.VeleroNamespaceEnvKey,
											ValueFrom: &corev1.EnvVarSource{
												FieldRef: &corev1.ObjectFieldSelector{
													FieldPath: "metadata.namespace",
												},
											},
										},
										{
											Name:  common.LDLibraryPathEnvKey,
											Value: "/plugins",
										},
									},
								},
							},
							Volumes: []corev1.Volume{
								{
									Name: "plugins",
									VolumeSource: corev1.VolumeSource{
										EmptyDir: &corev1.EmptyDirVolumeSource{},
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
							InitContainers: []corev1.Container{},
						}
					}