# Guide: Backup and Restore Stateful Applications on OpenShift using OADP and ODF

## Table of Contents
- [Guide: Backup and Restore Stateful Applications on OpenShift using OADP and ODF](#guide-backup-and-restore-stateful-applications-on-openshift-using-oadp-and-odf)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Installing OpenShift Data Foundation Operator](#installing-openshift-data-foundation-operator)
    - [Creating StorageSystem](#creating-storagesystem)
    - [Verify OpenShift Data Foundation Operator installation](#verify-openshift-data-foundation-operator-installation)
    - [Creating Object Bucket Claim](#creating-object-bucket-claim)
    - [Gathering information from Object Bucket](#gathering-information-from-object-bucket)
  - [Setup OpenShift Virtualization](#setup-openshift-virtualization)
  - [Creating VolumeSnapshotClass with Velero CSI label for Hostpath Provisioner](#creating-volumesnapshotclass-with-velero-csi-label-for-hostpath-provisioner)
  - [Install Windows 10 VM on OpenShift Virtualization](#install-windows-10-vm-on-openshift-virtualization)
  - [Creating DataProtectionApplication](#creating-dataprotectionapplication)

## Prerequisites
- Terminal environment
  - Your terminal has the following commands
    - [oc](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.9/html/cli_tools/openshift-cli-oc) binary
    - [git](https://git-scm.com/downloads) binary
    - velero
      - Set alias to use command from cluster (preferred)
        - `alias velero='oc -n openshift-adp exec deployment/velero -c velero -it -- ./velero'`
      - [Download velero from Github Release](https://velero.io/docs/v1.8/basic-install/#option-2-github-release)
  - Alternatively enter prepared environment in your terminal with `docker run -it ghcr.io/kaovilai/oadp-cli:v1.0.1 bash`
    - source can be found at https://github.com/kaovilai/oadp-cli
- [Authenticate as Cluster Admin inside your environment](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.9/html/cli_tools/openshift-cli-oc#cli-logging-in_cli-developer-commands) of an OpenShift 4.9 Cluster.
- Your cluster meets the minimum requirement for [OpenShift Data Foundation](https://access.redhat.com/documentation/en-us/red_hat_openshift_data_foundation/4.9/html/planning_your_deployment/infrastructure-requirements_rhodf#minimum-deployment-resource-requirements) in Internal Mode deployment
  - 3 worker nodes, each with at least:
    - 8 logical CPU
    - 24 GiB memory
    - 1+ storage devices
- You have [OpenShift Container Platform 4.10 installed](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.10/html/installing/ocp-installation-overview)
- You have [installed OpenShift API for Data Protection Operator (OADP)](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.10/html-single/backup_and_restore/index#installing-and-configuring-oadp)

## Installing OpenShift Data Foundation Operator
We will be using OpenShift Data Foundation to simplify application deployment across cloud providers which will be covered in the next section.

1. Open the OpenShift Web Console by navigating to the url below, make sure you are in Administrator view, not Developer.

   ```sh
   oc get route console -n openshift-console -ojsonpath="{.spec.host}"
   ```
   Authenticate with your credentials if necessary.

2. Navigate to *OperatorHub*, search for and install **OpenShift Data Foundation**

   ![OpenShift Data Foundation Installation](odfInstall.png)

### Creating StorageSystem

![OpenShift Data Foundation Installation finished](ODFfinishedInstall.png)

1. Click *Create StorageSystem* button after the install is completed (turns blue).
2. [Go to Product Documentation for Red Hat OpenShift Data Foundation 4.9](https://access.redhat.com/documentation/en-us/red_hat_openshift_data_foundation/4.9)
    1. Filter *Category* by **Deploying**
    2. Open deployment documentation your cloud provider.
    3. Follow *Creating an OpenShift Data Foundation cluster* instructions.

### Verify OpenShift Data Foundation Operator installation

You can validate the successful deployment of OpenShift Data Foundationn cluster following *Verifying OpenShift Data Foundation deployment* in the previous deployment documentation or with the following command:
```sh
oc get storagecluster -n openshift-storage ocs-storagecluster -o jsonpath='{.status.phase}{"\n"}'
```
And for the Multi-Cluster Gateway (MCG):
```sh
oc get noobaa -n openshift-storage noobaa -o jsonpath='{.status.phase}{"\n"}'
```
### Creating Object Bucket Claim
Object Bucket Claim creates a persistent storage bucket for Velero to store backed up kubernetes manifests.

1. Navigate to *Storage* > *Object Bucket Claim* and click *Create Object Bucket Claim*
   ![](ObjectBucketClaimCreate.png)
   Note the Project you are currently in. You can create a new Project or leave as *default*

2. set the following values:
   - ObjectBucketClaim Name:  `oadp-bucket`
   - StorageClass: `openshift-storage.noobaa.io`
   - BucketClass: `noobaa-default-bucket-class`

   ![](ObjectBucketClaimFields.png)

3. Click *Create*

   ![](ObjectBucketClaimReady.png)
   When the *Status* is *Bound*, the bucket is ready.

### Gathering information from Object Bucket
1. Gathering bucket name and host 
   - Using OpenShift CLI:
      - Get *bucket name*
        ```
        oc get configmap oadp-bucket -n default -o jsonpath='{.data.BUCKET_NAME}{"\n"}'
        ```
      - Get *bucket host*
        ```
        oc get configmap oadp-bucket -n default -o jsonpath='{.data.BUCKET_HOST}{"\n"}'
        ```
   - Using OpenShift Web Console:

     1. Click on Object Bucket *obc-default-oadp-bucket* and select YAML view

        ![](obc-default-oadp-bucket.png)
        Take note of the following information which may differ from the guide:
          - `.spec.endpoint.bucketName`. Seen in my screenshot as `oadp-bucket-c21e8d02-4d0b-4d19-a295-cecbf247f51f`
          - `.spec.endpoint.bucketHost`: Seen in my screenshot as `s3.openshift-storage.svc`

2. Gather oadp-bucket secret
   - Using OpenShift CLI:
      1. Get *AWS_ACCESS_KEY*
        ```
        oc get secret oadp-bucket -n default -o jsonpath='{.data.AWS_ACCESS_KEY_ID}{"\n"}' | base64 -d
        ```
      2. Get *AWS_SECRET_ACCESS_KEY*
        ```
        oc get secret oadp-bucket -n default -o jsonpath='{.data.AWS_SECRET_ACCESS_KEY}{"\n"}' | base64 -d
        ```
   - Using OpenShift Web Console
     1. Navigate to *Storage* > *Object Bucket Claim* > *oadp-bucket*. Ensure you are in the same *Project* used to create *oadp-bucket*.
     2. Click on oadp-secret in the bottom left to view bucket secrets
     3. Click Reveal values to see the bucket secret values. Copy data from *AWS_ACCESS_KEY_ID* and *AWS_SECRET_ACCESS_KEY* and save it as we'll need it later when installing the OADP Operator.
   
   Note: regardless of the cloud provider, the secret field names seen here may contain *AWS_\**.
3. Now you should have the following information:
   - *bucket name*
   - *bucket host*
   - *AWS_ACCESS_KEY_ID*
   - *AWS_SECRET_ACCESS_KEY*

## Setup OpenShift Virtualization
- [Follow OpenShift Virtualization install guide](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.10/html/virtualization/installing)
  - Ensure HyperConverged is created (Step 4.4.1.7)
    - `oc get hco -n openshift-cnv` should give you the hyperconverged object.
      - ```sh
        NAME                      AGE
        kubevirt-hyperconverged   1d
        ```
## Creating VolumeSnapshotClass with Velero CSI label for Hostpath Provisioner
VolumeSnapshotClass provides a way to describe the "classes" of storage when provisioning a volume snapshot.

Setting a `DeletionPolicy` of `Retain` on the VolumeSnapshotClass will preserve the volume snapshot in the storage system for the lifetime of the Velero backup and will prevent the deletion of the volume snapshot, in the storage system, in the event of a disaster where the namespace with the VolumeSnapshot object may be lost.

The Velero CSI plugin, to backup CSI backed PVCs, will choose the VolumeSnapshotClass in the cluster that has the same driver name and also has the `velero.io/csi-volumesnapshot-class: "true"` label set on it.
- Create VolumeSnapshotClass
  ```yaml
  apiVersion: snapshot.storage.k8s.io/v1
  kind: VolumeSnapshotClass
  metadata:
    name: hpp-snapclass
    labels:
      velero.io/csi-volumesnapshot-class: "true"
  driver: kubevirt.io/hostpath-provisioner
  deletionPolicy: Retain
  ```
## Install Windows 10 VM on OpenShift Virtualization

## Creating DataProtectionApplication


Without emulation virt-launcher
has 
  nodeSelector:
    hyperv.node.kubevirt.io/synictimer: 'true'
    hyperv.node.kubevirt.io/reset: 'true'
    kubevirt.io/schedulable: 'true'
    hyperv.node.kubevirt.io/runtime: 'true'
    hyperv.node.kubevirt.io/vpindex: 'true'
    hyperv.node.kubevirt.io/frequencies: 'true'
    hyperv.node.kubevirt.io/reenlightenment: 'true'
    hyperv.node.kubevirt.io/synic: 'true'
    hyperv.node.kubevirt.io/ipi: 'true'
    hyperv.node.kubevirt.io/tlbflush: 'true'