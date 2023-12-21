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

type EvaluateOnExitInitParameters struct {

	// Specifies the action to take if all of the specified conditions are met. The values are not case sensitive. Valid values: RETRY, EXIT.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// A glob pattern to match against the decimal representation of the exit code returned for a job.
	OnExitCode *string `json:"onExitCode,omitempty" tf:"on_exit_code,omitempty"`

	// A glob pattern to match against the reason returned for a job.
	OnReason *string `json:"onReason,omitempty" tf:"on_reason,omitempty"`

	// A glob pattern to match against the status reason returned for a job.
	OnStatusReason *string `json:"onStatusReason,omitempty" tf:"on_status_reason,omitempty"`
}

type EvaluateOnExitObservation struct {

	// Specifies the action to take if all of the specified conditions are met. The values are not case sensitive. Valid values: RETRY, EXIT.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// A glob pattern to match against the decimal representation of the exit code returned for a job.
	OnExitCode *string `json:"onExitCode,omitempty" tf:"on_exit_code,omitempty"`

	// A glob pattern to match against the reason returned for a job.
	OnReason *string `json:"onReason,omitempty" tf:"on_reason,omitempty"`

	// A glob pattern to match against the status reason returned for a job.
	OnStatusReason *string `json:"onStatusReason,omitempty" tf:"on_status_reason,omitempty"`
}

type EvaluateOnExitParameters struct {

	// Specifies the action to take if all of the specified conditions are met. The values are not case sensitive. Valid values: RETRY, EXIT.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// A glob pattern to match against the decimal representation of the exit code returned for a job.
	// +kubebuilder:validation:Optional
	OnExitCode *string `json:"onExitCode,omitempty" tf:"on_exit_code,omitempty"`

	// A glob pattern to match against the reason returned for a job.
	// +kubebuilder:validation:Optional
	OnReason *string `json:"onReason,omitempty" tf:"on_reason,omitempty"`

	// A glob pattern to match against the status reason returned for a job.
	// +kubebuilder:validation:Optional
	OnStatusReason *string `json:"onStatusReason,omitempty" tf:"on_status_reason,omitempty"`
}

type JobDefinitionInitParameters struct {

	// A valid container properties
	// provided as a single valid JSON document. This parameter is required if the type parameter is container.
	ContainerProperties *string `json:"containerProperties,omitempty" tf:"container_properties,omitempty"`

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A valid node properties
	// provided as a single valid JSON document. This parameter is required if the type parameter is multinode.
	NodeProperties *string `json:"nodeProperties,omitempty" tf:"node_properties,omitempty"`

	// Specifies the parameter substitution placeholders to set in the job definition.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The platform capabilities required by the job definition. If no value is specified, it defaults to EC2. To run the job on Fargate resources, specify FARGATE.
	// +listType=set
	PlatformCapabilities []*string `json:"platformCapabilities,omitempty" tf:"platform_capabilities,omitempty"`

	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is false.
	PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of retry_strategy is 1.  Defined below.
	RetryStrategy []RetryStrategyInitParameters `json:"retryStrategy,omitempty" tf:"retry_strategy,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of timeout is 1. Defined below.
	Timeout []TimeoutInitParameters `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of job definition. Must be container or multinode.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type JobDefinitionObservation struct {

	// The Amazon Resource Name of the job definition.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A valid container properties
	// provided as a single valid JSON document. This parameter is required if the type parameter is container.
	ContainerProperties *string `json:"containerProperties,omitempty" tf:"container_properties,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A valid node properties
	// provided as a single valid JSON document. This parameter is required if the type parameter is multinode.
	NodeProperties *string `json:"nodeProperties,omitempty" tf:"node_properties,omitempty"`

	// Specifies the parameter substitution placeholders to set in the job definition.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The platform capabilities required by the job definition. If no value is specified, it defaults to EC2. To run the job on Fargate resources, specify FARGATE.
	// +listType=set
	PlatformCapabilities []*string `json:"platformCapabilities,omitempty" tf:"platform_capabilities,omitempty"`

	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is false.
	PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of retry_strategy is 1.  Defined below.
	RetryStrategy []RetryStrategyObservation `json:"retryStrategy,omitempty" tf:"retry_strategy,omitempty"`

	// The revision of the job definition.
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of timeout is 1. Defined below.
	Timeout []TimeoutObservation `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of job definition. Must be container or multinode.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type JobDefinitionParameters struct {

	// A valid container properties
	// provided as a single valid JSON document. This parameter is required if the type parameter is container.
	// +kubebuilder:validation:Optional
	ContainerProperties *string `json:"containerProperties,omitempty" tf:"container_properties,omitempty"`

	// Specifies the name of the job definition.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A valid node properties
	// provided as a single valid JSON document. This parameter is required if the type parameter is multinode.
	// +kubebuilder:validation:Optional
	NodeProperties *string `json:"nodeProperties,omitempty" tf:"node_properties,omitempty"`

	// Specifies the parameter substitution placeholders to set in the job definition.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The platform capabilities required by the job definition. If no value is specified, it defaults to EC2. To run the job on Fargate resources, specify FARGATE.
	// +kubebuilder:validation:Optional
	// +listType=set
	PlatformCapabilities []*string `json:"platformCapabilities,omitempty" tf:"platform_capabilities,omitempty"`

	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is false.
	// +kubebuilder:validation:Optional
	PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of retry_strategy is 1.  Defined below.
	// +kubebuilder:validation:Optional
	RetryStrategy []RetryStrategyParameters `json:"retryStrategy,omitempty" tf:"retry_strategy,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of timeout is 1. Defined below.
	// +kubebuilder:validation:Optional
	Timeout []TimeoutParameters `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of job definition. Must be container or multinode.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RetryStrategyInitParameters struct {

	// The number of times to move a job to the RUNNABLE status. You may specify between 1 and 10 attempts.
	Attempts *float64 `json:"attempts,omitempty" tf:"attempts,omitempty"`

	// The evaluate on exit conditions under which the job should be retried or failed. If this parameter is specified, then the attempts parameter must also be specified. You may specify up to 5 configuration blocks.
	EvaluateOnExit []EvaluateOnExitInitParameters `json:"evaluateOnExit,omitempty" tf:"evaluate_on_exit,omitempty"`
}

type RetryStrategyObservation struct {

	// The number of times to move a job to the RUNNABLE status. You may specify between 1 and 10 attempts.
	Attempts *float64 `json:"attempts,omitempty" tf:"attempts,omitempty"`

	// The evaluate on exit conditions under which the job should be retried or failed. If this parameter is specified, then the attempts parameter must also be specified. You may specify up to 5 configuration blocks.
	EvaluateOnExit []EvaluateOnExitObservation `json:"evaluateOnExit,omitempty" tf:"evaluate_on_exit,omitempty"`
}

type RetryStrategyParameters struct {

	// The number of times to move a job to the RUNNABLE status. You may specify between 1 and 10 attempts.
	// +kubebuilder:validation:Optional
	Attempts *float64 `json:"attempts,omitempty" tf:"attempts,omitempty"`

	// The evaluate on exit conditions under which the job should be retried or failed. If this parameter is specified, then the attempts parameter must also be specified. You may specify up to 5 configuration blocks.
	// +kubebuilder:validation:Optional
	EvaluateOnExit []EvaluateOnExitParameters `json:"evaluateOnExit,omitempty" tf:"evaluate_on_exit,omitempty"`
}

type TimeoutInitParameters struct {

	// The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is 60 seconds.
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds,omitempty" tf:"attempt_duration_seconds,omitempty"`
}

type TimeoutObservation struct {

	// The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is 60 seconds.
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds,omitempty" tf:"attempt_duration_seconds,omitempty"`
}

type TimeoutParameters struct {

	// The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is 60 seconds.
	// +kubebuilder:validation:Optional
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds,omitempty" tf:"attempt_duration_seconds,omitempty"`
}

// JobDefinitionSpec defines the desired state of JobDefinition
type JobDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobDefinitionParameters `json:"forProvider"`
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
	InitProvider JobDefinitionInitParameters `json:"initProvider,omitempty"`
}

// JobDefinitionStatus defines the observed state of JobDefinition.
type JobDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JobDefinition is the Schema for the JobDefinitions API. Provides a Batch Job Definition resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type JobDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   JobDefinitionSpec   `json:"spec"`
	Status JobDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobDefinitionList contains a list of JobDefinitions
type JobDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobDefinition `json:"items"`
}

// Repository type metadata.
var (
	JobDefinition_Kind             = "JobDefinition"
	JobDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobDefinition_Kind}.String()
	JobDefinition_KindAPIVersion   = JobDefinition_Kind + "." + CRDGroupVersion.String()
	JobDefinition_GroupVersionKind = CRDGroupVersion.WithKind(JobDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&JobDefinition{}, &JobDefinitionList{})
}
