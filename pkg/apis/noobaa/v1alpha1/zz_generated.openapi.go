// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	// spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStore":       schema_pkg_apis_noobaa_v1alpha1_BackingStore(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreSpec":   schema_pkg_apis_noobaa_v1alpha1_BackingStoreSpec(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreStatus": schema_pkg_apis_noobaa_v1alpha1_BackingStoreStatus(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClass":        schema_pkg_apis_noobaa_v1alpha1_BucketClass(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassSpec":    schema_pkg_apis_noobaa_v1alpha1_BucketClassSpec(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassStatus":  schema_pkg_apis_noobaa_v1alpha1_BucketClassStatus(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.EndpointsSpec":      schema_pkg_apis_noobaa_v1alpha1_EndpointsSpec(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaa":             schema_pkg_apis_noobaa_v1alpha1_NooBaa(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaSpec":         schema_pkg_apis_noobaa_v1alpha1_NooBaaSpec(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaStatus":       schema_pkg_apis_noobaa_v1alpha1_NooBaaStatus(ref),
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BackingStore(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BackingStore is the Schema for the backingstores API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object metadata.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of the noobaa BackingStore.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Most recently observed status of the noobaa BackingStore.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BackingStoreSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BackingStoreSpec defines the desired state of BackingStore",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type is an enum of supported types",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"awsS3": {
						SchemaProps: spec.SchemaProps{
							Description: "AWSS3Spec specifies a backing store of type aws-s3",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AWSS3Spec"),
						},
					},
					"s3Compatible": {
						SchemaProps: spec.SchemaProps{
							Description: "S3Compatible specifies a backing store of type s3-compatible",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.S3CompatibleSpec"),
						},
					},
					"azureBlob": {
						SchemaProps: spec.SchemaProps{
							Description: "AzureBlob specifies a backing store of type azure-blob",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AzureBlobSpec"),
						},
					},
					"googleCloudStorage": {
						SchemaProps: spec.SchemaProps{
							Description: "GoogleCloudStorage specifies a backing store of type google-cloud-storage",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.GoogleCloudStorageSpec"),
						},
					},
					"pvPool": {
						SchemaProps: spec.SchemaProps{
							Description: "PVPool specifies a backing store of type pv-pool",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.PVPoolSpec"),
						},
					},
				},
				Required: []string{"type"},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AWSS3Spec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AzureBlobSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.GoogleCloudStorageSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.PVPoolSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.S3CompatibleSpec"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BackingStoreStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BackingStoreStatus defines the observed state of BackingStore",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is a simple, high-level summary of where the backing store is in its lifecycle",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "set",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a list of conditions related to operator reconciliation",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"relatedObjects": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "RelatedObjects is a list of objects related to this operator.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
					"mode": {
						SchemaProps: spec.SchemaProps{
							Description: "Mode specifies the updating mode of a BackingStore",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreMode"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreMode", "github.com/openshift/custom-resource-status/conditions/v1.Condition", "k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BucketClass(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BucketClass is the Schema for the bucketclasses API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object metadata.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of the noobaa BucketClass.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Most recently observed status of the noobaa BackingStore.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BucketClassSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BucketClassSpec defines the desired state of BucketClass",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"placementPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "PlacementPolicy specifies the placement policy for the bucket class",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.PlacementPolicy"),
						},
					},
				},
				Required: []string{"placementPolicy"},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.PlacementPolicy"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BucketClassStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BucketClassStatus defines the observed state of BucketClass",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is a simple, high-level summary of where the System is in its lifecycle",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "set",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a list of conditions related to operator reconciliation",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"relatedObjects": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "RelatedObjects is a list of objects related to this operator.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
					"mode": {
						SchemaProps: spec.SchemaProps{
							Description: "Mode is a simple, high-level summary of where the System is in its lifecycle",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/custom-resource-status/conditions/v1.Condition", "k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_EndpointsSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EndpointsSpec defines the desired state of noobaa endpoint deployment",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"minCount": {
						SchemaProps: spec.SchemaProps{
							Description: "MinCount, the number of endpoint instances (pods) to be used as the lower bound when autoscaling",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"maxCount": {
						SchemaProps: spec.SchemaProps{
							Description: "MaxCount, the number of endpoint instances (pods) to be used as the upper bound when autoscaling",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"additionalVirtualHosts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "AdditionalVirtualHosts (optional) provide a list of additional hostnames (on top of the buildin names defined by the cluster: service name, elb name, route name) to be used as virtual hosts by the the endpoints in the endpoint deployment",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Resources (optional) overrides the default resource requirements for every endpoint pod",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_NooBaa(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NooBaa is the Schema for the NooBaas API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object metadata.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of the noobaa system.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Most recently observed status of the noobaa system.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_NooBaaSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NooBaaSpec defines the desired state of System",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image (optional) overrides the default image for the server container",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dbImage": {
						SchemaProps: spec.SchemaProps{
							Description: "DBImage (optional) overrides the default image for the db container",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"coreResources": {
						SchemaProps: spec.SchemaProps{
							Description: "CoreResources (optional) overrides the default resource requirements for the server container",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"dbResources": {
						SchemaProps: spec.SchemaProps{
							Description: "DBResources (optional) overrides the default resource requirements for the db container",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"dbVolumeResources": {
						SchemaProps: spec.SchemaProps{
							Description: "DBVolumeResources (optional) overrides the default PVC resource requirements for the database volume. For the time being this field is immutable and can only be set on system creation. This is because volume size updates are only supported for increasing the size, and only if the storage class specifies `allowVolumeExpansion: true`,",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"dbStorageClass": {
						SchemaProps: spec.SchemaProps{
							Description: "DBStorageClass (optional) overrides the default cluster StorageClass for the database volume. For the time being this field is immutable and can only be set on system creation. This affects where the system stores its database which contains system config, buckets, objects meta-data and mapping file parts to storage locations.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pvPoolDefaultStorageClass": {
						SchemaProps: spec.SchemaProps{
							Description: "PVPoolDefaultStorageClass (optional) overrides the default cluster StorageClass for the pv-pool volumes. This affects where the system stores data chunks (encrypted). Updates to this field will only affect new pv-pools, but updates to existing pools are not supported by the operator.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tolerations": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Tolerations (optional) passed through to noobaa's pods",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"affinity": {
						SchemaProps: spec.SchemaProps{
							Description: "Affinity (optional) passed through to noobaa's pods",
							Ref:         ref("k8s.io/api/core/v1.Affinity"),
						},
					},
					"imagePullSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "ImagePullSecret (optional) sets a pull secret for the system image",
							Ref:         ref("k8s.io/api/core/v1.LocalObjectReference"),
						},
					},
					"endpoints": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoints (optional) sets configuration info for the noobaa endpoint deployment.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.EndpointsSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.EndpointsSpec", "k8s.io/api/core/v1.Affinity", "k8s.io/api/core/v1.LocalObjectReference", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_NooBaaStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NooBaaStatus defines the observed state of System",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "ObservedGeneration is the most recent generation observed for this noobaa system. It corresponds to the CR generation, which is updated on mutation by the API Server.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is a simple, high-level summary of where the System is in its lifecycle",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "set",
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a list of conditions related to operator reconciliation",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"relatedObjects": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "RelatedObjects is a list of objects related to this operator.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
					"actualImage": {
						SchemaProps: spec.SchemaProps{
							Description: "ActualImage is set to report which image the operator is using",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accounts": {
						SchemaProps: spec.SchemaProps{
							Description: "Accounts reports accounts info for the admin account",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AccountsStatus"),
						},
					},
					"services": {
						SchemaProps: spec.SchemaProps{
							Description: "Services reports addresses for the services",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.ServicesStatus"),
						},
					},
					"endpoints": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoints reports the actual number of endpoints in the endpoint deployment and the virtual hosts list used recognized by the endpoints",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.EndpointsStatus"),
						},
					},
					"readme": {
						SchemaProps: spec.SchemaProps{
							Description: "Readme is a user readable string with explanations on the system",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AccountsStatus", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.EndpointsStatus", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.ServicesStatus", "github.com/openshift/custom-resource-status/conditions/v1.Condition", "k8s.io/api/core/v1.ObjectReference"},
	}
}
