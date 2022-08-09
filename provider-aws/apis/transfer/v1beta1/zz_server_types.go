/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EndpointDetailsObservation struct {
}

type EndpointDetailsParameters struct {

	// +kubebuilder:validation:Optional
	AddressAllocationIds []*string `json:"addressAllocationIds,omitempty" tf:"address_allocation_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	VPCEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type OnUploadObservation struct {
}

type OnUploadParameters struct {

	// +kubebuilder:validation:Required
	ExecutionRole *string `json:"executionRole" tf:"execution_role,omitempty"`

	// +kubebuilder:validation:Required
	WorkflowID *string `json:"workflowId" tf:"workflow_id,omitempty"`
}

type ServerObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	HostKeyFingerprint *string `json:"hostKeyFingerprint,omitempty" tf:"host_key_fingerprint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ServerParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/acm/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateRef *v1.Reference `json:"certificateRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CertificateSelector *v1.Selector `json:"certificateSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointDetails []EndpointDetailsParameters `json:"endpointDetails,omitempty" tf:"endpoint_details,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// +kubebuilder:validation:Optional
	Function *string `json:"function,omitempty" tf:"function,omitempty"`

	// +kubebuilder:validation:Optional
	HostKeySecretRef *v1.SecretKeySelector `json:"hostKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IdentityProviderType *string `json:"identityProviderType,omitempty" tf:"identity_provider_type,omitempty"`

	// +kubebuilder:validation:Optional
	InvocationRole *string `json:"invocationRole,omitempty" tf:"invocation_role,omitempty"`

	// +kubebuilder:validation:Optional
	LoggingRole *string `json:"loggingRole,omitempty" tf:"logging_role,omitempty"`

	// +kubebuilder:validation:Optional
	PostAuthenticationLoginBannerSecretRef *v1.SecretKeySelector `json:"postAuthenticationLoginBannerSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PreAuthenticationLoginBannerSecretRef *v1.SecretKeySelector `json:"preAuthenticationLoginBannerSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityPolicyName *string `json:"securityPolicyName,omitempty" tf:"security_policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// +kubebuilder:validation:Optional
	WorkflowDetails []WorkflowDetailsParameters `json:"workflowDetails,omitempty" tf:"workflow_details,omitempty"`
}

type WorkflowDetailsObservation struct {
}

type WorkflowDetailsParameters struct {

	// +kubebuilder:validation:Optional
	OnUpload []OnUploadParameters `json:"onUpload,omitempty" tf:"on_upload,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}