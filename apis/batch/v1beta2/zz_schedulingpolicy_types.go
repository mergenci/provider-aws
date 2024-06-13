// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FairSharePolicyInitParameters struct {

	// A value used to reserve some of the available maximum vCPU for fair share identifiers that have not yet been used. For more information, see FairsharePolicy.
	ComputeReservation *float64 `json:"computeReservation,omitempty" tf:"compute_reservation,omitempty"`

	ShareDecaySeconds *float64 `json:"shareDecaySeconds,omitempty" tf:"share_decay_seconds,omitempty"`

	// One or more share distribution blocks which define the weights for the fair share identifiers for the fair share policy. For more information, see FairsharePolicy. The share_distribution block is documented below.
	ShareDistribution []ShareDistributionInitParameters `json:"shareDistribution,omitempty" tf:"share_distribution,omitempty"`
}

type FairSharePolicyObservation struct {

	// A value used to reserve some of the available maximum vCPU for fair share identifiers that have not yet been used. For more information, see FairsharePolicy.
	ComputeReservation *float64 `json:"computeReservation,omitempty" tf:"compute_reservation,omitempty"`

	ShareDecaySeconds *float64 `json:"shareDecaySeconds,omitempty" tf:"share_decay_seconds,omitempty"`

	// One or more share distribution blocks which define the weights for the fair share identifiers for the fair share policy. For more information, see FairsharePolicy. The share_distribution block is documented below.
	ShareDistribution []ShareDistributionObservation `json:"shareDistribution,omitempty" tf:"share_distribution,omitempty"`
}

type FairSharePolicyParameters struct {

	// A value used to reserve some of the available maximum vCPU for fair share identifiers that have not yet been used. For more information, see FairsharePolicy.
	// +kubebuilder:validation:Optional
	ComputeReservation *float64 `json:"computeReservation,omitempty" tf:"compute_reservation,omitempty"`

	// +kubebuilder:validation:Optional
	ShareDecaySeconds *float64 `json:"shareDecaySeconds,omitempty" tf:"share_decay_seconds,omitempty"`

	// One or more share distribution blocks which define the weights for the fair share identifiers for the fair share policy. For more information, see FairsharePolicy. The share_distribution block is documented below.
	// +kubebuilder:validation:Optional
	ShareDistribution []ShareDistributionParameters `json:"shareDistribution,omitempty" tf:"share_distribution,omitempty"`
}

type SchedulingPolicyInitParameters struct {
	FairSharePolicy *FairSharePolicyInitParameters `json:"fairSharePolicy,omitempty" tf:"fair_share_policy,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SchedulingPolicyObservation struct {

	// The Amazon Resource Name of the scheduling policy.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	FairSharePolicy *FairSharePolicyObservation `json:"fairSharePolicy,omitempty" tf:"fair_share_policy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SchedulingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	FairSharePolicy *FairSharePolicyParameters `json:"fairSharePolicy,omitempty" tf:"fair_share_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ShareDistributionInitParameters struct {

	// A fair share identifier or fair share identifier prefix. For more information, see ShareAttributes.
	ShareIdentifier *string `json:"shareIdentifier,omitempty" tf:"share_identifier,omitempty"`

	// The weight factor for the fair share identifier. For more information, see ShareAttributes.
	WeightFactor *float64 `json:"weightFactor,omitempty" tf:"weight_factor,omitempty"`
}

type ShareDistributionObservation struct {

	// A fair share identifier or fair share identifier prefix. For more information, see ShareAttributes.
	ShareIdentifier *string `json:"shareIdentifier,omitempty" tf:"share_identifier,omitempty"`

	// The weight factor for the fair share identifier. For more information, see ShareAttributes.
	WeightFactor *float64 `json:"weightFactor,omitempty" tf:"weight_factor,omitempty"`
}

type ShareDistributionParameters struct {

	// A fair share identifier or fair share identifier prefix. For more information, see ShareAttributes.
	// +kubebuilder:validation:Optional
	ShareIdentifier *string `json:"shareIdentifier" tf:"share_identifier,omitempty"`

	// The weight factor for the fair share identifier. For more information, see ShareAttributes.
	// +kubebuilder:validation:Optional
	WeightFactor *float64 `json:"weightFactor,omitempty" tf:"weight_factor,omitempty"`
}

// SchedulingPolicySpec defines the desired state of SchedulingPolicy
type SchedulingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SchedulingPolicyParameters `json:"forProvider"`
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
	InitProvider SchedulingPolicyInitParameters `json:"initProvider,omitempty"`
}

// SchedulingPolicyStatus defines the observed state of SchedulingPolicy.
type SchedulingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SchedulingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SchedulingPolicy is the Schema for the SchedulingPolicys API. Provides a Batch Scheduling Policy resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SchedulingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchedulingPolicySpec   `json:"spec"`
	Status            SchedulingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchedulingPolicyList contains a list of SchedulingPolicys
type SchedulingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SchedulingPolicy `json:"items"`
}

// Repository type metadata.
var (
	SchedulingPolicy_Kind             = "SchedulingPolicy"
	SchedulingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SchedulingPolicy_Kind}.String()
	SchedulingPolicy_KindAPIVersion   = SchedulingPolicy_Kind + "." + CRDGroupVersion.String()
	SchedulingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SchedulingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SchedulingPolicy{}, &SchedulingPolicyList{})
}