package lib

import (
	"context"
	"fmt"
	"os"

	snapshotv1beta1api "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1beta1"
	snapshotv1beta1client "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	velero "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	pkgbackup "github.com/vmware-tanzu/velero/pkg/backup"
	"github.com/vmware-tanzu/velero/pkg/cmd"
	"github.com/vmware-tanzu/velero/pkg/cmd/util/output"
	"github.com/vmware-tanzu/velero/pkg/features"
	veleroClientset "github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned"
	"github.com/vmware-tanzu/velero/pkg/label"
	"github.com/vmware-tanzu/velero/pkg/restic"
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
func DescribeBackup(ocClient client.Client, backup velero.Backup) string {
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
		fmt.Fprintf(os.Stderr, "error getting DeleteBackupRequests for backup %s: %v\n", backup.Name, err)
	}

	opts := label.NewListOptionsForBackup(backup.Name)
	podVolumeBackupList, err := veleroClient.VeleroV1().PodVolumeBackups(backup.Namespace).List(context.TODO(), opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting PodVolumeBackups for backup %s: %v\n", backup.Name, err)
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
			fmt.Fprintf(os.Stderr, "error getting VolumeSnapshotContent objects for backup %s: %v\n", backup.Name, err)
		}
	}

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
		fmt.Fprintf(os.Stderr, "error getting PodVolumeRestores for restore %s: %v\n", restore.Name, err)
	}

	return output.DescribeRestore(context.Background(), ocClient, &restore, podvolumeRestoreList.Items, details, veleroClient, insecureSkipTLSVerify, caCertFile)
}