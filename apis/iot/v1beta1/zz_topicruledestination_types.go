// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type TopicRuleDestinationInitParameters struct {

	// Whether or not to enable the destination. Default: true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Configuration of the virtual private cloud (VPC) connection. For more info, see the AWS documentation.
	VPCConfiguration []VPCConfigurationInitParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type TopicRuleDestinationObservation struct {

	// The ARN of the topic rule destination
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Whether or not to enable the destination. Default: true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration of the virtual private cloud (VPC) connection. For more info, see the AWS documentation.
	VPCConfiguration []VPCConfigurationObservation `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type TopicRuleDestinationParameters struct {

	// Whether or not to enable the destination. Default: true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration of the virtual private cloud (VPC) connection. For more info, see the AWS documentation.
	// +kubebuilder:validation:Optional
	VPCConfiguration []VPCConfigurationParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type VPCConfigurationInitParameters struct {
}

type VPCConfigurationObservation struct {

	// The ARN of a role that has permission to create and attach to elastic network interfaces (ENIs).
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The security groups of the VPC destination.
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// The subnet IDs of the VPC destination.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The ID of the VPC.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCConfigurationParameters struct {

	// The ARN of a role that has permission to create and attach to elastic network interfaces (ENIs).
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroups.
	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// The security groups of the VPC destination.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The subnet IDs of the VPC destination.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// The ID of the VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// TopicRuleDestinationSpec defines the desired state of TopicRuleDestination
type TopicRuleDestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicRuleDestinationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TopicRuleDestinationInitParameters `json:"initProvider,omitempty"`
}

// TopicRuleDestinationStatus defines the observed state of TopicRuleDestination.
type TopicRuleDestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicRuleDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TopicRuleDestination is the Schema for the TopicRuleDestinations API. Creates and manages an AWS IoT topic rule destination
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TopicRuleDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcConfiguration) || (has(self.initProvider) && has(self.initProvider.vpcConfiguration))",message="spec.forProvider.vpcConfiguration is a required parameter"
	Spec   TopicRuleDestinationSpec   `json:"spec"`
	Status TopicRuleDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicRuleDestinationList contains a list of TopicRuleDestinations
type TopicRuleDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicRuleDestination `json:"items"`
}

// Repository type metadata.
var (
	TopicRuleDestination_Kind             = "TopicRuleDestination"
	TopicRuleDestination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicRuleDestination_Kind}.String()
	TopicRuleDestination_KindAPIVersion   = TopicRuleDestination_Kind + "." + CRDGroupVersion.String()
	TopicRuleDestination_GroupVersionKind = CRDGroupVersion.WithKind(TopicRuleDestination_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicRuleDestination{}, &TopicRuleDestinationList{})
}
