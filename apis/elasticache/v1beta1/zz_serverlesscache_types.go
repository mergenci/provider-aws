// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CacheUsageLimitsInitParameters struct {

	// The maximum data storage limit in the cache, expressed in Gigabytes. See data_storage Block for details.
	DataStorage []DataStorageInitParameters `json:"dataStorage,omitempty" tf:"data_storage,omitempty"`

	// The configuration for the number of ElastiCache Processing Units (ECPU) the cache can consume per second. See ecpu_per_second Block for details.
	EcpuPerSecond []EcpuPerSecondInitParameters `json:"ecpuPerSecond,omitempty" tf:"ecpu_per_second,omitempty"`
}

type CacheUsageLimitsObservation struct {

	// The maximum data storage limit in the cache, expressed in Gigabytes. See data_storage Block for details.
	DataStorage []DataStorageObservation `json:"dataStorage,omitempty" tf:"data_storage,omitempty"`

	// The configuration for the number of ElastiCache Processing Units (ECPU) the cache can consume per second. See ecpu_per_second Block for details.
	EcpuPerSecond []EcpuPerSecondObservation `json:"ecpuPerSecond,omitempty" tf:"ecpu_per_second,omitempty"`
}

type CacheUsageLimitsParameters struct {

	// The maximum data storage limit in the cache, expressed in Gigabytes. See data_storage Block for details.
	// +kubebuilder:validation:Optional
	DataStorage []DataStorageParameters `json:"dataStorage,omitempty" tf:"data_storage,omitempty"`

	// The configuration for the number of ElastiCache Processing Units (ECPU) the cache can consume per second. See ecpu_per_second Block for details.
	// +kubebuilder:validation:Optional
	EcpuPerSecond []EcpuPerSecondParameters `json:"ecpuPerSecond,omitempty" tf:"ecpu_per_second,omitempty"`
}

type DataStorageInitParameters struct {

	// The upper limit for data storage the cache is set to use. Must be between 1 and 5,000.
	Maximum *float64 `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// The lower limit for data storage the cache is set to use. Must be between 1 and 5,000.
	Minimum *float64 `json:"minimum,omitempty" tf:"minimum,omitempty"`

	// The unit that the storage is measured in, in GB.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type DataStorageObservation struct {

	// The upper limit for data storage the cache is set to use. Must be between 1 and 5,000.
	Maximum *float64 `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// The lower limit for data storage the cache is set to use. Must be between 1 and 5,000.
	Minimum *float64 `json:"minimum,omitempty" tf:"minimum,omitempty"`

	// The unit that the storage is measured in, in GB.
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type DataStorageParameters struct {

	// The upper limit for data storage the cache is set to use. Must be between 1 and 5,000.
	// +kubebuilder:validation:Optional
	Maximum *float64 `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// The lower limit for data storage the cache is set to use. Must be between 1 and 5,000.
	// +kubebuilder:validation:Optional
	Minimum *float64 `json:"minimum,omitempty" tf:"minimum,omitempty"`

	// The unit that the storage is measured in, in GB.
	// +kubebuilder:validation:Optional
	Unit *string `json:"unit" tf:"unit,omitempty"`
}

type EcpuPerSecondInitParameters struct {

	// The maximum number of ECPUs the cache can consume per second. Must be between 1,000 and 15,000,000.
	Maximum *float64 `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// The minimum number of ECPUs the cache can consume per second. Must be between 1,000 and 15,000,000.
	Minimum *float64 `json:"minimum,omitempty" tf:"minimum,omitempty"`
}

type EcpuPerSecondObservation struct {

	// The maximum number of ECPUs the cache can consume per second. Must be between 1,000 and 15,000,000.
	Maximum *float64 `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// The minimum number of ECPUs the cache can consume per second. Must be between 1,000 and 15,000,000.
	Minimum *float64 `json:"minimum,omitempty" tf:"minimum,omitempty"`
}

type EcpuPerSecondParameters struct {

	// The maximum number of ECPUs the cache can consume per second. Must be between 1,000 and 15,000,000.
	// +kubebuilder:validation:Optional
	Maximum *float64 `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// The minimum number of ECPUs the cache can consume per second. Must be between 1,000 and 15,000,000.
	// +kubebuilder:validation:Optional
	Minimum *float64 `json:"minimum,omitempty" tf:"minimum,omitempty"`
}

type EndpointInitParameters struct {
}

type EndpointObservation struct {

	// The DNS hostname of the cache node.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The port number that the cache engine is listening on. Set as integer.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type EndpointParameters struct {
}

type ReaderEndpointInitParameters struct {
}

type ReaderEndpointObservation struct {

	// The DNS hostname of the cache node.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The port number that the cache engine is listening on. Set as integer.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type ReaderEndpointParameters struct {
}

type ServerlessCacheInitParameters struct {

	// Sets the cache usage limits for storage and ElastiCache Processing Units for the cache. See cache_usage_limits Block for details.
	CacheUsageLimits []CacheUsageLimitsInitParameters `json:"cacheUsageLimits,omitempty" tf:"cache_usage_limits,omitempty"`

	// The daily time that snapshots will be created from the new serverless cache. Only supported for engine types "redis" or "valkey". Defaults to 0.
	DailySnapshotTime *string `json:"dailySnapshotTime,omitempty" tf:"daily_snapshot_time,omitempty"`

	// User-provided description for the serverless cache. The default is NULL.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  Name of the cache engine to be used for this cache cluster. Valid values are memcached, redis or valkey.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// ARN of the customer managed key for encrypting the data at rest. If no KMS key is provided, a default service key is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// –  The version of the cache engine that will be used to create the serverless cache.
	// See Describe Cache Engine Versions in the AWS Documentation for supported versions.
	MajorEngineVersion *string `json:"majorEngineVersion,omitempty" tf:"major_engine_version,omitempty"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// A list of the one or more VPC security groups to be associated with the serverless cache. The security group will authorize traffic access for the VPC end-point (private-link). If no other information is given this will be the VPC’s Default Security Group that is associated with the cluster VPC end-point.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The list of ARN(s) of the snapshot that the new serverless cache will be created from. Available for Redis only.
	SnapshotArnsToRestore []*string `json:"snapshotArnsToRestore,omitempty" tf:"snapshot_arns_to_restore,omitempty"`

	// The number of snapshots that will be retained for the serverless cache that is being created. As new snapshots beyond this limit are added, the oldest snapshots will be deleted on a rolling basis. Available for Redis only.
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// –  A list of the identifiers of the subnets where the VPC endpoint for the serverless cache will be deployed. All the subnetIds must belong to the same VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The identifier of the UserGroup to be associated with the serverless cache. Available for Redis only. Default is NULL.
	UserGroupID *string `json:"userGroupId,omitempty" tf:"user_group_id,omitempty"`
}

type ServerlessCacheObservation struct {

	// The Amazon Resource Name (ARN) of the serverless cache.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Sets the cache usage limits for storage and ElastiCache Processing Units for the cache. See cache_usage_limits Block for details.
	CacheUsageLimits []CacheUsageLimitsObservation `json:"cacheUsageLimits,omitempty" tf:"cache_usage_limits,omitempty"`

	// Timestamp of when the serverless cache was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The daily time that snapshots will be created from the new serverless cache. Only supported for engine types "redis" or "valkey". Defaults to 0.
	DailySnapshotTime *string `json:"dailySnapshotTime,omitempty" tf:"daily_snapshot_time,omitempty"`

	// User-provided description for the serverless cache. The default is NULL.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Represents the information required for client programs to connect to a cache node. See endpoint Block for details.
	Endpoint []EndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// –  Name of the cache engine to be used for this cache cluster. Valid values are memcached, redis or valkey.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The name and version number of the engine the serverless cache is compatible with.
	FullEngineVersion *string `json:"fullEngineVersion,omitempty" tf:"full_engine_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the customer managed key for encrypting the data at rest. If no KMS key is provided, a default service key is used.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// –  The version of the cache engine that will be used to create the serverless cache.
	// See Describe Cache Engine Versions in the AWS Documentation for supported versions.
	MajorEngineVersion *string `json:"majorEngineVersion,omitempty" tf:"major_engine_version,omitempty"`

	// Represents the information required for client programs to connect to a cache node. See reader_endpoint Block for details.
	ReaderEndpoint []ReaderEndpointObservation `json:"readerEndpoint,omitempty" tf:"reader_endpoint,omitempty"`

	// A list of the one or more VPC security groups to be associated with the serverless cache. The security group will authorize traffic access for the VPC end-point (private-link). If no other information is given this will be the VPC’s Default Security Group that is associated with the cluster VPC end-point.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The list of ARN(s) of the snapshot that the new serverless cache will be created from. Available for Redis only.
	SnapshotArnsToRestore []*string `json:"snapshotArnsToRestore,omitempty" tf:"snapshot_arns_to_restore,omitempty"`

	// The number of snapshots that will be retained for the serverless cache that is being created. As new snapshots beyond this limit are added, the oldest snapshots will be deleted on a rolling basis. Available for Redis only.
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// The current status of the serverless cache. The allowed values are CREATING, AVAILABLE, DELETING, CREATE-FAILED and MODIFYING.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// –  A list of the identifiers of the subnets where the VPC endpoint for the serverless cache will be deployed. All the subnetIds must belong to the same VPC.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The identifier of the UserGroup to be associated with the serverless cache. Available for Redis only. Default is NULL.
	UserGroupID *string `json:"userGroupId,omitempty" tf:"user_group_id,omitempty"`
}

type ServerlessCacheParameters struct {

	// Sets the cache usage limits for storage and ElastiCache Processing Units for the cache. See cache_usage_limits Block for details.
	// +kubebuilder:validation:Optional
	CacheUsageLimits []CacheUsageLimitsParameters `json:"cacheUsageLimits,omitempty" tf:"cache_usage_limits,omitempty"`

	// The daily time that snapshots will be created from the new serverless cache. Only supported for engine types "redis" or "valkey". Defaults to 0.
	// +kubebuilder:validation:Optional
	DailySnapshotTime *string `json:"dailySnapshotTime,omitempty" tf:"daily_snapshot_time,omitempty"`

	// User-provided description for the serverless cache. The default is NULL.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// –  Name of the cache engine to be used for this cache cluster. Valid values are memcached, redis or valkey.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// ARN of the customer managed key for encrypting the data at rest. If no KMS key is provided, a default service key is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// –  The version of the cache engine that will be used to create the serverless cache.
	// See Describe Cache Engine Versions in the AWS Documentation for supported versions.
	// +kubebuilder:validation:Optional
	MajorEngineVersion *string `json:"majorEngineVersion,omitempty" tf:"major_engine_version,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// A list of the one or more VPC security groups to be associated with the serverless cache. The security group will authorize traffic access for the VPC end-point (private-link). If no other information is given this will be the VPC’s Default Security Group that is associated with the cluster VPC end-point.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The list of ARN(s) of the snapshot that the new serverless cache will be created from. Available for Redis only.
	// +kubebuilder:validation:Optional
	SnapshotArnsToRestore []*string `json:"snapshotArnsToRestore,omitempty" tf:"snapshot_arns_to_restore,omitempty"`

	// The number of snapshots that will be retained for the serverless cache that is being created. As new snapshots beyond this limit are added, the oldest snapshots will be deleted on a rolling basis. Available for Redis only.
	// +kubebuilder:validation:Optional
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// References to Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in ec2 to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// –  A list of the identifiers of the subnets where the VPC endpoint for the serverless cache will be deployed. All the subnetIds must belong to the same VPC.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIDRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIDSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The identifier of the UserGroup to be associated with the serverless cache. Available for Redis only. Default is NULL.
	// +kubebuilder:validation:Optional
	UserGroupID *string `json:"userGroupId,omitempty" tf:"user_group_id,omitempty"`
}

// ServerlessCacheSpec defines the desired state of ServerlessCache
type ServerlessCacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerlessCacheParameters `json:"forProvider"`
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
	InitProvider ServerlessCacheInitParameters `json:"initProvider,omitempty"`
}

// ServerlessCacheStatus defines the observed state of ServerlessCache.
type ServerlessCacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerlessCacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServerlessCache is the Schema for the ServerlessCaches API. Provides an ElastiCache Serverless Cache resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServerlessCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engine) || (has(self.initProvider) && has(self.initProvider.engine))",message="spec.forProvider.engine is a required parameter"
	Spec   ServerlessCacheSpec   `json:"spec"`
	Status ServerlessCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerlessCacheList contains a list of ServerlessCaches
type ServerlessCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerlessCache `json:"items"`
}

// Repository type metadata.
var (
	ServerlessCache_Kind             = "ServerlessCache"
	ServerlessCache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerlessCache_Kind}.String()
	ServerlessCache_KindAPIVersion   = ServerlessCache_Kind + "." + CRDGroupVersion.String()
	ServerlessCache_GroupVersionKind = CRDGroupVersion.WithKind(ServerlessCache_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerlessCache{}, &ServerlessCacheList{})
}
