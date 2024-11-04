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

type AccountLevelDetailedStatusCodeMetricsInitParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AccountLevelDetailedStatusCodeMetricsObservation struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AccountLevelDetailedStatusCodeMetricsParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AccountLevelInitParameters struct {

	// S3 Storage Lens activity metrics. See Activity Metrics below for more details.
	ActivityMetrics *ActivityMetricsInitParameters `json:"activityMetrics,omitempty" tf:"activity_metrics,omitempty"`

	// optimization metrics for S3 Storage Lens. See Advanced Cost-Optimization Metrics below for more details.
	AdvancedCostOptimizationMetrics *AdvancedCostOptimizationMetricsInitParameters `json:"advancedCostOptimizationMetrics,omitempty" tf:"advanced_cost_optimization_metrics,omitempty"`

	// protection metrics for S3 Storage Lens. See Advanced Data-Protection Metrics below for more details.
	AdvancedDataProtectionMetrics *AdvancedDataProtectionMetricsInitParameters `json:"advancedDataProtectionMetrics,omitempty" tf:"advanced_data_protection_metrics,omitempty"`

	// level configuration. See Bucket Level below for more details.
	BucketLevel *BucketLevelInitParameters `json:"bucketLevel,omitempty" tf:"bucket_level,omitempty"`

	// Detailed status code metrics for S3 Storage Lens. See Detailed Status Code Metrics below for more details.
	DetailedStatusCodeMetrics *AccountLevelDetailedStatusCodeMetricsInitParameters `json:"detailedStatusCodeMetrics,omitempty" tf:"detailed_status_code_metrics,omitempty"`
}

type AccountLevelObservation struct {

	// S3 Storage Lens activity metrics. See Activity Metrics below for more details.
	ActivityMetrics *ActivityMetricsObservation `json:"activityMetrics,omitempty" tf:"activity_metrics,omitempty"`

	// optimization metrics for S3 Storage Lens. See Advanced Cost-Optimization Metrics below for more details.
	AdvancedCostOptimizationMetrics *AdvancedCostOptimizationMetricsObservation `json:"advancedCostOptimizationMetrics,omitempty" tf:"advanced_cost_optimization_metrics,omitempty"`

	// protection metrics for S3 Storage Lens. See Advanced Data-Protection Metrics below for more details.
	AdvancedDataProtectionMetrics *AdvancedDataProtectionMetricsObservation `json:"advancedDataProtectionMetrics,omitempty" tf:"advanced_data_protection_metrics,omitempty"`

	// level configuration. See Bucket Level below for more details.
	BucketLevel *BucketLevelObservation `json:"bucketLevel,omitempty" tf:"bucket_level,omitempty"`

	// Detailed status code metrics for S3 Storage Lens. See Detailed Status Code Metrics below for more details.
	DetailedStatusCodeMetrics *AccountLevelDetailedStatusCodeMetricsObservation `json:"detailedStatusCodeMetrics,omitempty" tf:"detailed_status_code_metrics,omitempty"`
}

type AccountLevelParameters struct {

	// S3 Storage Lens activity metrics. See Activity Metrics below for more details.
	// +kubebuilder:validation:Optional
	ActivityMetrics *ActivityMetricsParameters `json:"activityMetrics,omitempty" tf:"activity_metrics,omitempty"`

	// optimization metrics for S3 Storage Lens. See Advanced Cost-Optimization Metrics below for more details.
	// +kubebuilder:validation:Optional
	AdvancedCostOptimizationMetrics *AdvancedCostOptimizationMetricsParameters `json:"advancedCostOptimizationMetrics,omitempty" tf:"advanced_cost_optimization_metrics,omitempty"`

	// protection metrics for S3 Storage Lens. See Advanced Data-Protection Metrics below for more details.
	// +kubebuilder:validation:Optional
	AdvancedDataProtectionMetrics *AdvancedDataProtectionMetricsParameters `json:"advancedDataProtectionMetrics,omitempty" tf:"advanced_data_protection_metrics,omitempty"`

	// level configuration. See Bucket Level below for more details.
	// +kubebuilder:validation:Optional
	BucketLevel *BucketLevelParameters `json:"bucketLevel" tf:"bucket_level,omitempty"`

	// Detailed status code metrics for S3 Storage Lens. See Detailed Status Code Metrics below for more details.
	// +kubebuilder:validation:Optional
	DetailedStatusCodeMetrics *AccountLevelDetailedStatusCodeMetricsParameters `json:"detailedStatusCodeMetrics,omitempty" tf:"detailed_status_code_metrics,omitempty"`
}

type ActivityMetricsInitParameters struct {

	// Whether the activity metrics are enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ActivityMetricsObservation struct {

	// Whether the activity metrics are enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ActivityMetricsParameters struct {

	// Whether the activity metrics are enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdvancedCostOptimizationMetricsInitParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdvancedCostOptimizationMetricsObservation struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdvancedCostOptimizationMetricsParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdvancedDataProtectionMetricsInitParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdvancedDataProtectionMetricsObservation struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdvancedDataProtectionMetricsParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AwsOrgInitParameters struct {

	// The Amazon Resource Name (ARN) of the bucket.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type AwsOrgObservation struct {

	// The Amazon Resource Name (ARN) of the bucket.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type AwsOrgParameters struct {

	// The Amazon Resource Name (ARN) of the bucket.
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn" tf:"arn,omitempty"`
}

type BucketLevelActivityMetricsInitParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketLevelActivityMetricsObservation struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketLevelActivityMetricsParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketLevelAdvancedCostOptimizationMetricsInitParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketLevelAdvancedCostOptimizationMetricsObservation struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketLevelAdvancedCostOptimizationMetricsParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketLevelAdvancedDataProtectionMetricsInitParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketLevelAdvancedDataProtectionMetricsObservation struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketLevelAdvancedDataProtectionMetricsParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketLevelInitParameters struct {

	// S3 Storage Lens activity metrics. See Activity Metrics below for more details.
	ActivityMetrics *BucketLevelActivityMetricsInitParameters `json:"activityMetrics,omitempty" tf:"activity_metrics,omitempty"`

	// optimization metrics for S3 Storage Lens. See Advanced Cost-Optimization Metrics below for more details.
	AdvancedCostOptimizationMetrics *BucketLevelAdvancedCostOptimizationMetricsInitParameters `json:"advancedCostOptimizationMetrics,omitempty" tf:"advanced_cost_optimization_metrics,omitempty"`

	// protection metrics for S3 Storage Lens. See Advanced Data-Protection Metrics below for more details.
	AdvancedDataProtectionMetrics *BucketLevelAdvancedDataProtectionMetricsInitParameters `json:"advancedDataProtectionMetrics,omitempty" tf:"advanced_data_protection_metrics,omitempty"`

	// Detailed status code metrics for S3 Storage Lens. See Detailed Status Code Metrics below for more details.
	DetailedStatusCodeMetrics *DetailedStatusCodeMetricsInitParameters `json:"detailedStatusCodeMetrics,omitempty" tf:"detailed_status_code_metrics,omitempty"`

	// level metrics for S3 Storage Lens. See Prefix Level below for more details.
	PrefixLevel *PrefixLevelInitParameters `json:"prefixLevel,omitempty" tf:"prefix_level,omitempty"`
}

type BucketLevelObservation struct {

	// S3 Storage Lens activity metrics. See Activity Metrics below for more details.
	ActivityMetrics *BucketLevelActivityMetricsObservation `json:"activityMetrics,omitempty" tf:"activity_metrics,omitempty"`

	// optimization metrics for S3 Storage Lens. See Advanced Cost-Optimization Metrics below for more details.
	AdvancedCostOptimizationMetrics *BucketLevelAdvancedCostOptimizationMetricsObservation `json:"advancedCostOptimizationMetrics,omitempty" tf:"advanced_cost_optimization_metrics,omitempty"`

	// protection metrics for S3 Storage Lens. See Advanced Data-Protection Metrics below for more details.
	AdvancedDataProtectionMetrics *BucketLevelAdvancedDataProtectionMetricsObservation `json:"advancedDataProtectionMetrics,omitempty" tf:"advanced_data_protection_metrics,omitempty"`

	// Detailed status code metrics for S3 Storage Lens. See Detailed Status Code Metrics below for more details.
	DetailedStatusCodeMetrics *DetailedStatusCodeMetricsObservation `json:"detailedStatusCodeMetrics,omitempty" tf:"detailed_status_code_metrics,omitempty"`

	// level metrics for S3 Storage Lens. See Prefix Level below for more details.
	PrefixLevel *PrefixLevelObservation `json:"prefixLevel,omitempty" tf:"prefix_level,omitempty"`
}

type BucketLevelParameters struct {

	// S3 Storage Lens activity metrics. See Activity Metrics below for more details.
	// +kubebuilder:validation:Optional
	ActivityMetrics *BucketLevelActivityMetricsParameters `json:"activityMetrics,omitempty" tf:"activity_metrics,omitempty"`

	// optimization metrics for S3 Storage Lens. See Advanced Cost-Optimization Metrics below for more details.
	// +kubebuilder:validation:Optional
	AdvancedCostOptimizationMetrics *BucketLevelAdvancedCostOptimizationMetricsParameters `json:"advancedCostOptimizationMetrics,omitempty" tf:"advanced_cost_optimization_metrics,omitempty"`

	// protection metrics for S3 Storage Lens. See Advanced Data-Protection Metrics below for more details.
	// +kubebuilder:validation:Optional
	AdvancedDataProtectionMetrics *BucketLevelAdvancedDataProtectionMetricsParameters `json:"advancedDataProtectionMetrics,omitempty" tf:"advanced_data_protection_metrics,omitempty"`

	// Detailed status code metrics for S3 Storage Lens. See Detailed Status Code Metrics below for more details.
	// +kubebuilder:validation:Optional
	DetailedStatusCodeMetrics *DetailedStatusCodeMetricsParameters `json:"detailedStatusCodeMetrics,omitempty" tf:"detailed_status_code_metrics,omitempty"`

	// level metrics for S3 Storage Lens. See Prefix Level below for more details.
	// +kubebuilder:validation:Optional
	PrefixLevel *PrefixLevelParameters `json:"prefixLevel,omitempty" tf:"prefix_level,omitempty"`
}

type CloudWatchMetricsInitParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type CloudWatchMetricsObservation struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type CloudWatchMetricsParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type DataExportInitParameters struct {

	// Amazon CloudWatch publishing for S3 Storage Lens metrics. See Cloud Watch Metrics below for more details.
	CloudWatchMetrics *CloudWatchMetricsInitParameters `json:"cloudWatchMetrics,omitempty" tf:"cloud_watch_metrics,omitempty"`

	// The bucket where the S3 Storage Lens metrics export will be located. See S3 Bucket Destination below for more details.
	S3BucketDestination *S3BucketDestinationInitParameters `json:"s3BucketDestination,omitempty" tf:"s3_bucket_destination,omitempty"`
}

type DataExportObservation struct {

	// Amazon CloudWatch publishing for S3 Storage Lens metrics. See Cloud Watch Metrics below for more details.
	CloudWatchMetrics *CloudWatchMetricsObservation `json:"cloudWatchMetrics,omitempty" tf:"cloud_watch_metrics,omitempty"`

	// The bucket where the S3 Storage Lens metrics export will be located. See S3 Bucket Destination below for more details.
	S3BucketDestination *S3BucketDestinationObservation `json:"s3BucketDestination,omitempty" tf:"s3_bucket_destination,omitempty"`
}

type DataExportParameters struct {

	// Amazon CloudWatch publishing for S3 Storage Lens metrics. See Cloud Watch Metrics below for more details.
	// +kubebuilder:validation:Optional
	CloudWatchMetrics *CloudWatchMetricsParameters `json:"cloudWatchMetrics,omitempty" tf:"cloud_watch_metrics,omitempty"`

	// The bucket where the S3 Storage Lens metrics export will be located. See S3 Bucket Destination below for more details.
	// +kubebuilder:validation:Optional
	S3BucketDestination *S3BucketDestinationParameters `json:"s3BucketDestination,omitempty" tf:"s3_bucket_destination,omitempty"`
}

type DetailedStatusCodeMetricsInitParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DetailedStatusCodeMetricsObservation struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DetailedStatusCodeMetricsParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type EncryptionInitParameters struct {

	// KMS encryption. See SSE KMS below for more details.
	SseKMS *SseKMSInitParameters `json:"sseKms,omitempty" tf:"sse_kms,omitempty"`

	// S3 encryption. An empty configuration block {} should be used.
	SseS3 []SseS3InitParameters `json:"sseS3,omitempty" tf:"sse_s3,omitempty"`
}

type EncryptionObservation struct {

	// KMS encryption. See SSE KMS below for more details.
	SseKMS *SseKMSObservation `json:"sseKms,omitempty" tf:"sse_kms,omitempty"`

	// S3 encryption. An empty configuration block {} should be used.
	SseS3 []SseS3Parameters `json:"sseS3,omitempty" tf:"sse_s3,omitempty"`
}

type EncryptionParameters struct {

	// KMS encryption. See SSE KMS below for more details.
	// +kubebuilder:validation:Optional
	SseKMS *SseKMSParameters `json:"sseKms,omitempty" tf:"sse_kms,omitempty"`

	// S3 encryption. An empty configuration block {} should be used.
	// +kubebuilder:validation:Optional
	SseS3 []SseS3Parameters `json:"sseS3,omitempty" tf:"sse_s3,omitempty"`
}

type ExcludeInitParameters struct {

	// List of S3 bucket ARNs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +listType=set
	Buckets []*string `json:"buckets,omitempty" tf:"buckets,omitempty"`

	// References to Bucket in s3 to populate buckets.
	// +kubebuilder:validation:Optional
	BucketsRefs []v1.Reference `json:"bucketsRefs,omitempty" tf:"-"`

	// Selector for a list of Bucket in s3 to populate buckets.
	// +kubebuilder:validation:Optional
	BucketsSelector *v1.Selector `json:"bucketsSelector,omitempty" tf:"-"`

	// List of AWS Regions.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type ExcludeObservation struct {

	// List of S3 bucket ARNs.
	// +listType=set
	Buckets []*string `json:"buckets,omitempty" tf:"buckets,omitempty"`

	// List of AWS Regions.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type ExcludeParameters struct {

	// List of S3 bucket ARNs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	// +listType=set
	Buckets []*string `json:"buckets,omitempty" tf:"buckets,omitempty"`

	// References to Bucket in s3 to populate buckets.
	// +kubebuilder:validation:Optional
	BucketsRefs []v1.Reference `json:"bucketsRefs,omitempty" tf:"-"`

	// Selector for a list of Bucket in s3 to populate buckets.
	// +kubebuilder:validation:Optional
	BucketsSelector *v1.Selector `json:"bucketsSelector,omitempty" tf:"-"`

	// List of AWS Regions.
	// +kubebuilder:validation:Optional
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type IncludeInitParameters struct {

	// List of S3 bucket ARNs.
	// +listType=set
	Buckets []*string `json:"buckets,omitempty" tf:"buckets,omitempty"`

	// List of AWS Regions.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type IncludeObservation struct {

	// List of S3 bucket ARNs.
	// +listType=set
	Buckets []*string `json:"buckets,omitempty" tf:"buckets,omitempty"`

	// List of AWS Regions.
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type IncludeParameters struct {

	// List of S3 bucket ARNs.
	// +kubebuilder:validation:Optional
	// +listType=set
	Buckets []*string `json:"buckets,omitempty" tf:"buckets,omitempty"`

	// List of AWS Regions.
	// +kubebuilder:validation:Optional
	// +listType=set
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type PrefixLevelInitParameters struct {

	// level storage metrics for S3 Storage Lens. See Prefix Level Storage Metrics below for more details.
	StorageMetrics *StorageMetricsInitParameters `json:"storageMetrics,omitempty" tf:"storage_metrics,omitempty"`
}

type PrefixLevelObservation struct {

	// level storage metrics for S3 Storage Lens. See Prefix Level Storage Metrics below for more details.
	StorageMetrics *StorageMetricsObservation `json:"storageMetrics,omitempty" tf:"storage_metrics,omitempty"`
}

type PrefixLevelParameters struct {

	// level storage metrics for S3 Storage Lens. See Prefix Level Storage Metrics below for more details.
	// +kubebuilder:validation:Optional
	StorageMetrics *StorageMetricsParameters `json:"storageMetrics" tf:"storage_metrics,omitempty"`
}

type S3BucketDestinationInitParameters struct {

	// The account ID of the owner of the S3 Storage Lens metrics export bucket.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The Amazon Resource Name (ARN) of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Reference to a Bucket in s3 to populate arn.
	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate arn.
	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// Encryption of the metrics exports in this bucket. See Encryption below for more details.
	Encryption *EncryptionInitParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// The export format. Valid values: CSV, Parquet.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The schema version of the export file. Valid values: V_1.
	OutputSchemaVersion *string `json:"outputSchemaVersion,omitempty" tf:"output_schema_version,omitempty"`

	// The prefix of the destination bucket where the metrics export will be delivered.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3BucketDestinationObservation struct {

	// The account ID of the owner of the S3 Storage Lens metrics export bucket.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The Amazon Resource Name (ARN) of the bucket.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Encryption of the metrics exports in this bucket. See Encryption below for more details.
	Encryption *EncryptionObservation `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// The export format. Valid values: CSV, Parquet.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The schema version of the export file. Valid values: V_1.
	OutputSchemaVersion *string `json:"outputSchemaVersion,omitempty" tf:"output_schema_version,omitempty"`

	// The prefix of the destination bucket where the metrics export will be delivered.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3BucketDestinationParameters struct {

	// The account ID of the owner of the S3 Storage Lens metrics export bucket.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// The Amazon Resource Name (ARN) of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta2.Bucket
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Reference to a Bucket in s3 to populate arn.
	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate arn.
	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// Encryption of the metrics exports in this bucket. See Encryption below for more details.
	// +kubebuilder:validation:Optional
	Encryption *EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// The export format. Valid values: CSV, Parquet.
	// +kubebuilder:validation:Optional
	Format *string `json:"format" tf:"format,omitempty"`

	// The schema version of the export file. Valid values: V_1.
	// +kubebuilder:validation:Optional
	OutputSchemaVersion *string `json:"outputSchemaVersion" tf:"output_schema_version,omitempty"`

	// The prefix of the destination bucket where the metrics export will be delivered.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type SelectionCriteriaInitParameters struct {

	// The delimiter of the selection criteria being used.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// The max depth of the selection criteria.
	MaxDepth *float64 `json:"maxDepth,omitempty" tf:"max_depth,omitempty"`

	// The minimum number of storage bytes percentage whose metrics will be selected.
	MinStorageBytesPercentage *float64 `json:"minStorageBytesPercentage,omitempty" tf:"min_storage_bytes_percentage,omitempty"`
}

type SelectionCriteriaObservation struct {

	// The delimiter of the selection criteria being used.
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// The max depth of the selection criteria.
	MaxDepth *float64 `json:"maxDepth,omitempty" tf:"max_depth,omitempty"`

	// The minimum number of storage bytes percentage whose metrics will be selected.
	MinStorageBytesPercentage *float64 `json:"minStorageBytesPercentage,omitempty" tf:"min_storage_bytes_percentage,omitempty"`
}

type SelectionCriteriaParameters struct {

	// The delimiter of the selection criteria being used.
	// +kubebuilder:validation:Optional
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// The max depth of the selection criteria.
	// +kubebuilder:validation:Optional
	MaxDepth *float64 `json:"maxDepth,omitempty" tf:"max_depth,omitempty"`

	// The minimum number of storage bytes percentage whose metrics will be selected.
	// +kubebuilder:validation:Optional
	MinStorageBytesPercentage *float64 `json:"minStorageBytesPercentage,omitempty" tf:"min_storage_bytes_percentage,omitempty"`
}

type SseKMSInitParameters struct {

	// KMS key ARN.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type SseKMSObservation struct {

	// KMS key ARN.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type SseKMSParameters struct {

	// KMS key ARN.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId" tf:"key_id,omitempty"`
}

type SseS3InitParameters struct {
}

type SseS3Observation struct {
}

type SseS3Parameters struct {
}

type StorageLensConfigurationInitParameters struct {

	// The AWS account ID for the S3 Storage Lens configuration.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The ID of the S3 Storage Lens configuration.
	ConfigID *string `json:"configId,omitempty" tf:"config_id,omitempty"`

	// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
	StorageLensConfiguration *StorageLensConfigurationStorageLensConfigurationInitParameters `json:"storageLensConfiguration,omitempty" tf:"storage_lens_configuration,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StorageLensConfigurationObservation struct {

	// The AWS account ID for the S3 Storage Lens configuration.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Amazon Resource Name (ARN) of the S3 Storage Lens configuration.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the S3 Storage Lens configuration.
	ConfigID *string `json:"configId,omitempty" tf:"config_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
	StorageLensConfiguration *StorageLensConfigurationStorageLensConfigurationObservation `json:"storageLensConfiguration,omitempty" tf:"storage_lens_configuration,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type StorageLensConfigurationParameters struct {

	// The AWS account ID for the S3 Storage Lens configuration.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The ID of the S3 Storage Lens configuration.
	// +kubebuilder:validation:Optional
	ConfigID *string `json:"configId,omitempty" tf:"config_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
	// +kubebuilder:validation:Optional
	StorageLensConfiguration *StorageLensConfigurationStorageLensConfigurationParameters `json:"storageLensConfiguration,omitempty" tf:"storage_lens_configuration,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StorageLensConfigurationStorageLensConfigurationInitParameters struct {

	// level configurations of the S3 Storage Lens configuration. See Account Level below for more details.
	AccountLevel *AccountLevelInitParameters `json:"accountLevel,omitempty" tf:"account_level,omitempty"`

	// The Amazon Web Services organization for the S3 Storage Lens configuration. See AWS Org below for more details.
	AwsOrg *AwsOrgInitParameters `json:"awsOrg,omitempty" tf:"aws_org,omitempty"`

	// Properties of S3 Storage Lens metrics export including the destination, schema and format. See Data Export below for more details.
	DataExport *DataExportInitParameters `json:"dataExport,omitempty" tf:"data_export,omitempty"`

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// What is excluded in this configuration. Conflicts with include. See Exclude below for more details.
	Exclude *ExcludeInitParameters `json:"exclude,omitempty" tf:"exclude,omitempty"`

	// What is included in this configuration. Conflicts with exclude. See Include below for more details.
	Include *IncludeInitParameters `json:"include,omitempty" tf:"include,omitempty"`
}

type StorageLensConfigurationStorageLensConfigurationObservation struct {

	// level configurations of the S3 Storage Lens configuration. See Account Level below for more details.
	AccountLevel *AccountLevelObservation `json:"accountLevel,omitempty" tf:"account_level,omitempty"`

	// The Amazon Web Services organization for the S3 Storage Lens configuration. See AWS Org below for more details.
	AwsOrg *AwsOrgObservation `json:"awsOrg,omitempty" tf:"aws_org,omitempty"`

	// Properties of S3 Storage Lens metrics export including the destination, schema and format. See Data Export below for more details.
	DataExport *DataExportObservation `json:"dataExport,omitempty" tf:"data_export,omitempty"`

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// What is excluded in this configuration. Conflicts with include. See Exclude below for more details.
	Exclude *ExcludeObservation `json:"exclude,omitempty" tf:"exclude,omitempty"`

	// What is included in this configuration. Conflicts with exclude. See Include below for more details.
	Include *IncludeObservation `json:"include,omitempty" tf:"include,omitempty"`
}

type StorageLensConfigurationStorageLensConfigurationParameters struct {

	// level configurations of the S3 Storage Lens configuration. See Account Level below for more details.
	// +kubebuilder:validation:Optional
	AccountLevel *AccountLevelParameters `json:"accountLevel" tf:"account_level,omitempty"`

	// The Amazon Web Services organization for the S3 Storage Lens configuration. See AWS Org below for more details.
	// +kubebuilder:validation:Optional
	AwsOrg *AwsOrgParameters `json:"awsOrg,omitempty" tf:"aws_org,omitempty"`

	// Properties of S3 Storage Lens metrics export including the destination, schema and format. See Data Export below for more details.
	// +kubebuilder:validation:Optional
	DataExport *DataExportParameters `json:"dataExport,omitempty" tf:"data_export,omitempty"`

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// What is excluded in this configuration. Conflicts with include. See Exclude below for more details.
	// +kubebuilder:validation:Optional
	Exclude *ExcludeParameters `json:"exclude,omitempty" tf:"exclude,omitempty"`

	// What is included in this configuration. Conflicts with exclude. See Include below for more details.
	// +kubebuilder:validation:Optional
	Include *IncludeParameters `json:"include,omitempty" tf:"include,omitempty"`
}

type StorageMetricsInitParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Selection criteria. See Selection Criteria below for more details.
	SelectionCriteria *SelectionCriteriaInitParameters `json:"selectionCriteria,omitempty" tf:"selection_criteria,omitempty"`
}

type StorageMetricsObservation struct {

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Selection criteria. See Selection Criteria below for more details.
	SelectionCriteria *SelectionCriteriaObservation `json:"selectionCriteria,omitempty" tf:"selection_criteria,omitempty"`
}

type StorageMetricsParameters struct {

	// Whether the S3 Storage Lens configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Selection criteria. See Selection Criteria below for more details.
	// +kubebuilder:validation:Optional
	SelectionCriteria *SelectionCriteriaParameters `json:"selectionCriteria,omitempty" tf:"selection_criteria,omitempty"`
}

// StorageLensConfigurationSpec defines the desired state of StorageLensConfiguration
type StorageLensConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageLensConfigurationParameters `json:"forProvider"`
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
	InitProvider StorageLensConfigurationInitParameters `json:"initProvider,omitempty"`
}

// StorageLensConfigurationStatus defines the observed state of StorageLensConfiguration.
type StorageLensConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageLensConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StorageLensConfiguration is the Schema for the StorageLensConfigurations API. Provides a resource to manage an S3 Storage Lens configuration.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StorageLensConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configId) || (has(self.initProvider) && has(self.initProvider.configId))",message="spec.forProvider.configId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageLensConfiguration) || (has(self.initProvider) && has(self.initProvider.storageLensConfiguration))",message="spec.forProvider.storageLensConfiguration is a required parameter"
	Spec   StorageLensConfigurationSpec   `json:"spec"`
	Status StorageLensConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageLensConfigurationList contains a list of StorageLensConfigurations
type StorageLensConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageLensConfiguration `json:"items"`
}

// Repository type metadata.
var (
	StorageLensConfiguration_Kind             = "StorageLensConfiguration"
	StorageLensConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageLensConfiguration_Kind}.String()
	StorageLensConfiguration_KindAPIVersion   = StorageLensConfiguration_Kind + "." + CRDGroupVersion.String()
	StorageLensConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(StorageLensConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageLensConfiguration{}, &StorageLensConfigurationList{})
}
