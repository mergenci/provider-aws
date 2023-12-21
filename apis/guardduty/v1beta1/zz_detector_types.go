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

type AuditLogsInitParameters struct {

	// If true, enables Malware Protection as data source for the detector.
	// Defaults to true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type AuditLogsObservation struct {

	// If true, enables Malware Protection as data source for the detector.
	// Defaults to true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type AuditLogsParameters struct {

	// If true, enables Malware Protection as data source for the detector.
	// Defaults to true.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable" tf:"enable,omitempty"`
}

type DatasourcesInitParameters struct {

	// Configures Kubernetes protection.
	// See Kubernetes and Kubernetes Audit Logs below for more details.
	Kubernetes []KubernetesInitParameters `json:"kubernetes,omitempty" tf:"kubernetes,omitempty"`

	// Configures Malware Protection.
	// See Malware Protection, Scan EC2 instance with findings and EBS volumes below for more details.
	MalwareProtection []MalwareProtectionInitParameters `json:"malwareProtection,omitempty" tf:"malware_protection,omitempty"`

	// Configures S3 protection.
	// See S3 Logs below for more details.
	S3Logs []S3LogsInitParameters `json:"s3Logs,omitempty" tf:"s3_logs,omitempty"`
}

type DatasourcesObservation struct {

	// Configures Kubernetes protection.
	// See Kubernetes and Kubernetes Audit Logs below for more details.
	Kubernetes []KubernetesObservation `json:"kubernetes,omitempty" tf:"kubernetes,omitempty"`

	// Configures Malware Protection.
	// See Malware Protection, Scan EC2 instance with findings and EBS volumes below for more details.
	MalwareProtection []MalwareProtectionObservation `json:"malwareProtection,omitempty" tf:"malware_protection,omitempty"`

	// Configures S3 protection.
	// See S3 Logs below for more details.
	S3Logs []S3LogsObservation `json:"s3Logs,omitempty" tf:"s3_logs,omitempty"`
}

type DatasourcesParameters struct {

	// Configures Kubernetes protection.
	// See Kubernetes and Kubernetes Audit Logs below for more details.
	// +kubebuilder:validation:Optional
	Kubernetes []KubernetesParameters `json:"kubernetes,omitempty" tf:"kubernetes,omitempty"`

	// Configures Malware Protection.
	// See Malware Protection, Scan EC2 instance with findings and EBS volumes below for more details.
	// +kubebuilder:validation:Optional
	MalwareProtection []MalwareProtectionParameters `json:"malwareProtection,omitempty" tf:"malware_protection,omitempty"`

	// Configures S3 protection.
	// See S3 Logs below for more details.
	// +kubebuilder:validation:Optional
	S3Logs []S3LogsParameters `json:"s3Logs,omitempty" tf:"s3_logs,omitempty"`
}

type DetectorInitParameters struct {

	// Describes which data sources will be enabled for the detector. See Data Sources below for more details. Deprecated in favor of aws_guardduty_detector_feature resources.
	Datasources []DatasourcesInitParameters `json:"datasources,omitempty" tf:"datasources,omitempty"`

	// Enable monitoring and feedback reporting. Setting to false is equivalent to "suspending" GuardDuty. Defaults to true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to SIX_HOURS. Valid values for standalone and primary accounts: FIFTEEN_MINUTES, ONE_HOUR, SIX_HOURS. See AWS Documentation for more information.
	FindingPublishingFrequency *string `json:"findingPublishingFrequency,omitempty" tf:"finding_publishing_frequency,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DetectorObservation struct {

	// The AWS account ID of the GuardDuty detector
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Amazon Resource Name (ARN) of the GuardDuty detector
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Describes which data sources will be enabled for the detector. See Data Sources below for more details. Deprecated in favor of aws_guardduty_detector_feature resources.
	Datasources []DatasourcesObservation `json:"datasources,omitempty" tf:"datasources,omitempty"`

	// Enable monitoring and feedback reporting. Setting to false is equivalent to "suspending" GuardDuty. Defaults to true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to SIX_HOURS. Valid values for standalone and primary accounts: FIFTEEN_MINUTES, ONE_HOUR, SIX_HOURS. See AWS Documentation for more information.
	FindingPublishingFrequency *string `json:"findingPublishingFrequency,omitempty" tf:"finding_publishing_frequency,omitempty"`

	// The ID of the GuardDuty detector
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DetectorParameters struct {

	// Describes which data sources will be enabled for the detector. See Data Sources below for more details. Deprecated in favor of aws_guardduty_detector_feature resources.
	// +kubebuilder:validation:Optional
	Datasources []DatasourcesParameters `json:"datasources,omitempty" tf:"datasources,omitempty"`

	// Enable monitoring and feedback reporting. Setting to false is equivalent to "suspending" GuardDuty. Defaults to true.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`

	// Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to SIX_HOURS. Valid values for standalone and primary accounts: FIFTEEN_MINUTES, ONE_HOUR, SIX_HOURS. See AWS Documentation for more information.
	// +kubebuilder:validation:Optional
	FindingPublishingFrequency *string `json:"findingPublishingFrequency,omitempty" tf:"finding_publishing_frequency,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EBSVolumesInitParameters struct {

	// If true, enables Malware Protection as data source for the detector.
	// Defaults to true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type EBSVolumesObservation struct {

	// If true, enables Malware Protection as data source for the detector.
	// Defaults to true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type EBSVolumesParameters struct {

	// If true, enables Malware Protection as data source for the detector.
	// Defaults to true.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable" tf:"enable,omitempty"`
}

type KubernetesInitParameters struct {

	// Configures Kubernetes audit logs as a data source for Kubernetes protection.
	// See Kubernetes Audit Logs below for more details.
	AuditLogs []AuditLogsInitParameters `json:"auditLogs,omitempty" tf:"audit_logs,omitempty"`
}

type KubernetesObservation struct {

	// Configures Kubernetes audit logs as a data source for Kubernetes protection.
	// See Kubernetes Audit Logs below for more details.
	AuditLogs []AuditLogsObservation `json:"auditLogs,omitempty" tf:"audit_logs,omitempty"`
}

type KubernetesParameters struct {

	// Configures Kubernetes audit logs as a data source for Kubernetes protection.
	// See Kubernetes Audit Logs below for more details.
	// +kubebuilder:validation:Optional
	AuditLogs []AuditLogsParameters `json:"auditLogs" tf:"audit_logs,omitempty"`
}

type MalwareProtectionInitParameters struct {

	// Configure whether Malware Protection is enabled as data source for EC2 instances with findings for the detector.
	// See Scan EC2 instance with findings below for more details.
	ScanEC2InstanceWithFindings []ScanEC2InstanceWithFindingsInitParameters `json:"scanEc2InstanceWithFindings,omitempty" tf:"scan_ec2_instance_with_findings,omitempty"`
}

type MalwareProtectionObservation struct {

	// Configure whether Malware Protection is enabled as data source for EC2 instances with findings for the detector.
	// See Scan EC2 instance with findings below for more details.
	ScanEC2InstanceWithFindings []ScanEC2InstanceWithFindingsObservation `json:"scanEc2InstanceWithFindings,omitempty" tf:"scan_ec2_instance_with_findings,omitempty"`
}

type MalwareProtectionParameters struct {

	// Configure whether Malware Protection is enabled as data source for EC2 instances with findings for the detector.
	// See Scan EC2 instance with findings below for more details.
	// +kubebuilder:validation:Optional
	ScanEC2InstanceWithFindings []ScanEC2InstanceWithFindingsParameters `json:"scanEc2InstanceWithFindings" tf:"scan_ec2_instance_with_findings,omitempty"`
}

type S3LogsInitParameters struct {

	// If true, enables S3 protection.
	// Defaults to true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type S3LogsObservation struct {

	// If true, enables S3 protection.
	// Defaults to true.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type S3LogsParameters struct {

	// If true, enables S3 protection.
	// Defaults to true.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable" tf:"enable,omitempty"`
}

type ScanEC2InstanceWithFindingsInitParameters struct {

	// Configure whether scanning EBS volumes is enabled as data source for the detector for instances with findings.
	// See EBS volumes below for more details.
	EBSVolumes []EBSVolumesInitParameters `json:"ebsVolumes,omitempty" tf:"ebs_volumes,omitempty"`
}

type ScanEC2InstanceWithFindingsObservation struct {

	// Configure whether scanning EBS volumes is enabled as data source for the detector for instances with findings.
	// See EBS volumes below for more details.
	EBSVolumes []EBSVolumesObservation `json:"ebsVolumes,omitempty" tf:"ebs_volumes,omitempty"`
}

type ScanEC2InstanceWithFindingsParameters struct {

	// Configure whether scanning EBS volumes is enabled as data source for the detector for instances with findings.
	// See EBS volumes below for more details.
	// +kubebuilder:validation:Optional
	EBSVolumes []EBSVolumesParameters `json:"ebsVolumes" tf:"ebs_volumes,omitempty"`
}

// DetectorSpec defines the desired state of Detector
type DetectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DetectorParameters `json:"forProvider"`
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
	InitProvider DetectorInitParameters `json:"initProvider,omitempty"`
}

// DetectorStatus defines the observed state of Detector.
type DetectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DetectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Detector is the Schema for the Detectors API. Provides a resource to manage an Amazon GuardDuty detector
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Detector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DetectorSpec   `json:"spec"`
	Status            DetectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DetectorList contains a list of Detectors
type DetectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Detector `json:"items"`
}

// Repository type metadata.
var (
	Detector_Kind             = "Detector"
	Detector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Detector_Kind}.String()
	Detector_KindAPIVersion   = Detector_Kind + "." + CRDGroupVersion.String()
	Detector_GroupVersionKind = CRDGroupVersion.WithKind(Detector_Kind)
)

func init() {
	SchemeBuilder.Register(&Detector{}, &DetectorList{})
}
