module github.com/openshift/oadp-operator

go 1.16

require (
	github.com/go-logr/logr v0.4.0
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/hashicorp/go-hclog v0.8.0 // indirect
	github.com/hashicorp/go-plugin v1.0.1 // indirect
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/noobaa/noobaa-operator/v5/pkg/apis/noobaa/v1alpha1 v0.0.0-00010101000000-000000000000
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
	github.com/openshift/api v3.9.1-0.20190924102528-32369d4db2ad+incompatible
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/vmware-tanzu/velero v1.6.1-0.20210806003158-ed5809b7fc22 // TODO: Update this to a pinned version
	golang.org/x/tools v0.1.5 // indirect
	k8s.io/api v0.22.1
	k8s.io/apiextensions-apiserver v0.21.3
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	k8s.io/utils v0.0.0-20210722164352-7f3ee0f31471
	sigs.k8s.io/controller-runtime v0.9.6
)

replace github.com/noobaa/noobaa-operator/v5/pkg/apis/noobaa/v1alpha1 => ./pkg/apis/noobaa/v1alpha1
