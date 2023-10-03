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

type CatalogTableInitParameters struct {

	// Description of the table.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block for open table formats. See open_table_format_input below.
	OpenTableFormatInput []OpenTableFormatInputInitParameters `json:"openTableFormatInput,omitempty" tf:"open_table_format_input,omitempty"`

	// Owner of the table.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Properties associated with this table, as a list of key-value pairs.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Configuration block for a maximum of 3 partition indexes. See partition_index below.
	PartitionIndex []PartitionIndexInitParameters `json:"partitionIndex,omitempty" tf:"partition_index,omitempty"`

	// Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See partition_keys below.
	PartitionKeys []PartitionKeysInitParameters `json:"partitionKeys,omitempty" tf:"partition_keys,omitempty"`

	// Retention time for this table.
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// Configuration block for information about the physical storage of this table. For more information, refer to the Glue Developer Guide. See storage_descriptor below.
	StorageDescriptor []StorageDescriptorInitParameters `json:"storageDescriptor,omitempty" tf:"storage_descriptor,omitempty"`

	// Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as ALTER TABLE and SHOW CREATE TABLE will fail if this argument is empty.
	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	// Configuration block of a target table for resource linking. See target_table below.
	TargetTable []TargetTableInitParameters `json:"targetTable,omitempty" tf:"target_table,omitempty"`

	// If the table is a view, the expanded text of the view; otherwise null.
	ViewExpandedText *string `json:"viewExpandedText,omitempty" tf:"view_expanded_text,omitempty"`

	// If the table is a view, the original text of the view; otherwise null.
	ViewOriginalText *string `json:"viewOriginalText,omitempty" tf:"view_original_text,omitempty"`
}

type CatalogTableObservation struct {

	// The ARN of the Glue Table.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Description of the table.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Catalog ID, Database name and of the name table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block for open table formats. See open_table_format_input below.
	OpenTableFormatInput []OpenTableFormatInputObservation `json:"openTableFormatInput,omitempty" tf:"open_table_format_input,omitempty"`

	// Owner of the table.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Properties associated with this table, as a list of key-value pairs.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Configuration block for a maximum of 3 partition indexes. See partition_index below.
	PartitionIndex []PartitionIndexObservation `json:"partitionIndex,omitempty" tf:"partition_index,omitempty"`

	// Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See partition_keys below.
	PartitionKeys []PartitionKeysObservation `json:"partitionKeys,omitempty" tf:"partition_keys,omitempty"`

	// Retention time for this table.
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// Configuration block for information about the physical storage of this table. For more information, refer to the Glue Developer Guide. See storage_descriptor below.
	StorageDescriptor []StorageDescriptorObservation `json:"storageDescriptor,omitempty" tf:"storage_descriptor,omitempty"`

	// Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as ALTER TABLE and SHOW CREATE TABLE will fail if this argument is empty.
	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	// Configuration block of a target table for resource linking. See target_table below.
	TargetTable []TargetTableObservation `json:"targetTable,omitempty" tf:"target_table,omitempty"`

	// If the table is a view, the expanded text of the view; otherwise null.
	ViewExpandedText *string `json:"viewExpandedText,omitempty" tf:"view_expanded_text,omitempty"`

	// If the table is a view, the original text of the view; otherwise null.
	ViewOriginalText *string `json:"viewOriginalText,omitempty" tf:"view_original_text,omitempty"`
}

type CatalogTableParameters struct {

	// ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
	// +kubebuilder:validation:Required
	CatalogID *string `json:"catalogId" tf:"catalog_id,omitempty"`

	// Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glue/v1beta1.CatalogDatabase
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a CatalogDatabase in glue to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a CatalogDatabase in glue to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// Description of the table.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Configuration block for open table formats. See open_table_format_input below.
	// +kubebuilder:validation:Optional
	OpenTableFormatInput []OpenTableFormatInputParameters `json:"openTableFormatInput,omitempty" tf:"open_table_format_input,omitempty"`

	// Owner of the table.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Properties associated with this table, as a list of key-value pairs.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Configuration block for a maximum of 3 partition indexes. See partition_index below.
	// +kubebuilder:validation:Optional
	PartitionIndex []PartitionIndexParameters `json:"partitionIndex,omitempty" tf:"partition_index,omitempty"`

	// Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See partition_keys below.
	// +kubebuilder:validation:Optional
	PartitionKeys []PartitionKeysParameters `json:"partitionKeys,omitempty" tf:"partition_keys,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Retention time for this table.
	// +kubebuilder:validation:Optional
	Retention *float64 `json:"retention,omitempty" tf:"retention,omitempty"`

	// Configuration block for information about the physical storage of this table. For more information, refer to the Glue Developer Guide. See storage_descriptor below.
	// +kubebuilder:validation:Optional
	StorageDescriptor []StorageDescriptorParameters `json:"storageDescriptor,omitempty" tf:"storage_descriptor,omitempty"`

	// Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as ALTER TABLE and SHOW CREATE TABLE will fail if this argument is empty.
	// +kubebuilder:validation:Optional
	TableType *string `json:"tableType,omitempty" tf:"table_type,omitempty"`

	// Configuration block of a target table for resource linking. See target_table below.
	// +kubebuilder:validation:Optional
	TargetTable []TargetTableParameters `json:"targetTable,omitempty" tf:"target_table,omitempty"`

	// If the table is a view, the expanded text of the view; otherwise null.
	// +kubebuilder:validation:Optional
	ViewExpandedText *string `json:"viewExpandedText,omitempty" tf:"view_expanded_text,omitempty"`

	// If the table is a view, the original text of the view; otherwise null.
	// +kubebuilder:validation:Optional
	ViewOriginalText *string `json:"viewOriginalText,omitempty" tf:"view_original_text,omitempty"`
}

type ColumnsInitParameters struct {

	// Free-form text comment.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the Column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value pairs defining properties associated with the column.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Datatype of data in the Column.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ColumnsObservation struct {

	// Free-form text comment.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the Column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value pairs defining properties associated with the column.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Datatype of data in the Column.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ColumnsParameters struct {

	// Free-form text comment.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the Column.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Key-value pairs defining properties associated with the column.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Datatype of data in the Column.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IcebergInputInitParameters struct {

	// A required metadata operation. Can only be set to CREATE.
	MetadataOperation *string `json:"metadataOperation,omitempty" tf:"metadata_operation,omitempty"`

	// The table version for the Iceberg table. Defaults to 2.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type IcebergInputObservation struct {

	// A required metadata operation. Can only be set to CREATE.
	MetadataOperation *string `json:"metadataOperation,omitempty" tf:"metadata_operation,omitempty"`

	// The table version for the Iceberg table. Defaults to 2.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type IcebergInputParameters struct {

	// A required metadata operation. Can only be set to CREATE.
	// +kubebuilder:validation:Optional
	MetadataOperation *string `json:"metadataOperation" tf:"metadata_operation,omitempty"`

	// The table version for the Iceberg table. Defaults to 2.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type OpenTableFormatInputInitParameters struct {

	// Configuration block for iceberg table config. See iceberg_input below.
	IcebergInput []IcebergInputInitParameters `json:"icebergInput,omitempty" tf:"iceberg_input,omitempty"`
}

type OpenTableFormatInputObservation struct {

	// Configuration block for iceberg table config. See iceberg_input below.
	IcebergInput []IcebergInputObservation `json:"icebergInput,omitempty" tf:"iceberg_input,omitempty"`
}

type OpenTableFormatInputParameters struct {

	// Configuration block for iceberg table config. See iceberg_input below.
	// +kubebuilder:validation:Optional
	IcebergInput []IcebergInputParameters `json:"icebergInput" tf:"iceberg_input,omitempty"`
}

type PartitionIndexInitParameters struct {

	// Name of the partition index.
	IndexName *string `json:"indexName,omitempty" tf:"index_name,omitempty"`

	// Keys for the partition index.
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`
}

type PartitionIndexObservation struct {

	// Name of the partition index.
	IndexName *string `json:"indexName,omitempty" tf:"index_name,omitempty"`

	IndexStatus *string `json:"indexStatus,omitempty" tf:"index_status,omitempty"`

	// Keys for the partition index.
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`
}

type PartitionIndexParameters struct {

	// Name of the partition index.
	// +kubebuilder:validation:Optional
	IndexName *string `json:"indexName" tf:"index_name,omitempty"`

	// Keys for the partition index.
	// +kubebuilder:validation:Optional
	Keys []*string `json:"keys" tf:"keys,omitempty"`
}

type PartitionKeysInitParameters struct {

	// Free-form text comment.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the Partition Key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Datatype of data in the Partition Key.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PartitionKeysObservation struct {

	// Free-form text comment.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the Partition Key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Datatype of data in the Partition Key.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PartitionKeysParameters struct {

	// Free-form text comment.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// Name of the Partition Key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Datatype of data in the Partition Key.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SchemaIDInitParameters struct {

	// Name of the schema registry that contains the schema. Must be provided when schema_name is specified and conflicts with schema_arn.
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// ARN of the schema. One of schema_arn or schema_name has to be provided.
	SchemaArn *string `json:"schemaArn,omitempty" tf:"schema_arn,omitempty"`

	// Name of the schema. One of schema_arn or schema_name has to be provided.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`
}

type SchemaIDObservation struct {

	// Name of the schema registry that contains the schema. Must be provided when schema_name is specified and conflicts with schema_arn.
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// ARN of the schema. One of schema_arn or schema_name has to be provided.
	SchemaArn *string `json:"schemaArn,omitempty" tf:"schema_arn,omitempty"`

	// Name of the schema. One of schema_arn or schema_name has to be provided.
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`
}

type SchemaIDParameters struct {

	// Name of the schema registry that contains the schema. Must be provided when schema_name is specified and conflicts with schema_arn.
	// +kubebuilder:validation:Optional
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// ARN of the schema. One of schema_arn or schema_name has to be provided.
	// +kubebuilder:validation:Optional
	SchemaArn *string `json:"schemaArn,omitempty" tf:"schema_arn,omitempty"`

	// Name of the schema. One of schema_arn or schema_name has to be provided.
	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`
}

type SchemaReferenceInitParameters struct {

	// Configuration block that contains schema identity fields. Either this or the schema_version_id has to be provided. See schema_id below.
	SchemaID []SchemaIDInitParameters `json:"schemaId,omitempty" tf:"schema_id,omitempty"`

	// Unique ID assigned to a version of the schema. Either this or the schema_id has to be provided.
	SchemaVersionID *string `json:"schemaVersionId,omitempty" tf:"schema_version_id,omitempty"`

	// Version number of the schema.
	SchemaVersionNumber *float64 `json:"schemaVersionNumber,omitempty" tf:"schema_version_number,omitempty"`
}

type SchemaReferenceObservation struct {

	// Configuration block that contains schema identity fields. Either this or the schema_version_id has to be provided. See schema_id below.
	SchemaID []SchemaIDObservation `json:"schemaId,omitempty" tf:"schema_id,omitempty"`

	// Unique ID assigned to a version of the schema. Either this or the schema_id has to be provided.
	SchemaVersionID *string `json:"schemaVersionId,omitempty" tf:"schema_version_id,omitempty"`

	// Version number of the schema.
	SchemaVersionNumber *float64 `json:"schemaVersionNumber,omitempty" tf:"schema_version_number,omitempty"`
}

type SchemaReferenceParameters struct {

	// Configuration block that contains schema identity fields. Either this or the schema_version_id has to be provided. See schema_id below.
	// +kubebuilder:validation:Optional
	SchemaID []SchemaIDParameters `json:"schemaId,omitempty" tf:"schema_id,omitempty"`

	// Unique ID assigned to a version of the schema. Either this or the schema_id has to be provided.
	// +kubebuilder:validation:Optional
	SchemaVersionID *string `json:"schemaVersionId,omitempty" tf:"schema_version_id,omitempty"`

	// Version number of the schema.
	// +kubebuilder:validation:Optional
	SchemaVersionNumber *float64 `json:"schemaVersionNumber" tf:"schema_version_number,omitempty"`
}

type SerDeInfoInitParameters struct {

	// Name of the SerDe.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Usually the class that implements the SerDe. An example is org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe.
	SerializationLibrary *string `json:"serializationLibrary,omitempty" tf:"serialization_library,omitempty"`
}

type SerDeInfoObservation struct {

	// Name of the SerDe.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of initialization parameters for the SerDe, in key-value form.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Usually the class that implements the SerDe. An example is org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe.
	SerializationLibrary *string `json:"serializationLibrary,omitempty" tf:"serialization_library,omitempty"`
}

type SerDeInfoParameters struct {

	// Name of the SerDe.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Map of initialization parameters for the SerDe, in key-value form.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Usually the class that implements the SerDe. An example is org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe.
	// +kubebuilder:validation:Optional
	SerializationLibrary *string `json:"serializationLibrary,omitempty" tf:"serialization_library,omitempty"`
}

type SkewedInfoInitParameters struct {

	// List of names of columns that contain skewed values.
	SkewedColumnNames []*string `json:"skewedColumnNames,omitempty" tf:"skewed_column_names,omitempty"`

	// List of values that appear so frequently as to be considered skewed.
	SkewedColumnValueLocationMaps map[string]*string `json:"skewedColumnValueLocationMaps,omitempty" tf:"skewed_column_value_location_maps,omitempty"`

	// Map of skewed values to the columns that contain them.
	SkewedColumnValues []*string `json:"skewedColumnValues,omitempty" tf:"skewed_column_values,omitempty"`
}

type SkewedInfoObservation struct {

	// List of names of columns that contain skewed values.
	SkewedColumnNames []*string `json:"skewedColumnNames,omitempty" tf:"skewed_column_names,omitempty"`

	// List of values that appear so frequently as to be considered skewed.
	SkewedColumnValueLocationMaps map[string]*string `json:"skewedColumnValueLocationMaps,omitempty" tf:"skewed_column_value_location_maps,omitempty"`

	// Map of skewed values to the columns that contain them.
	SkewedColumnValues []*string `json:"skewedColumnValues,omitempty" tf:"skewed_column_values,omitempty"`
}

type SkewedInfoParameters struct {

	// List of names of columns that contain skewed values.
	// +kubebuilder:validation:Optional
	SkewedColumnNames []*string `json:"skewedColumnNames,omitempty" tf:"skewed_column_names,omitempty"`

	// List of values that appear so frequently as to be considered skewed.
	// +kubebuilder:validation:Optional
	SkewedColumnValueLocationMaps map[string]*string `json:"skewedColumnValueLocationMaps,omitempty" tf:"skewed_column_value_location_maps,omitempty"`

	// Map of skewed values to the columns that contain them.
	// +kubebuilder:validation:Optional
	SkewedColumnValues []*string `json:"skewedColumnValues,omitempty" tf:"skewed_column_values,omitempty"`
}

type SortColumnsInitParameters struct {

	// Name of the column.
	Column *string `json:"column,omitempty" tf:"column,omitempty"`

	// Whether the column is sorted in ascending (1) or descending order (0).
	SortOrder *float64 `json:"sortOrder,omitempty" tf:"sort_order,omitempty"`
}

type SortColumnsObservation struct {

	// Name of the column.
	Column *string `json:"column,omitempty" tf:"column,omitempty"`

	// Whether the column is sorted in ascending (1) or descending order (0).
	SortOrder *float64 `json:"sortOrder,omitempty" tf:"sort_order,omitempty"`
}

type SortColumnsParameters struct {

	// Name of the column.
	// +kubebuilder:validation:Optional
	Column *string `json:"column" tf:"column,omitempty"`

	// Whether the column is sorted in ascending (1) or descending order (0).
	// +kubebuilder:validation:Optional
	SortOrder *float64 `json:"sortOrder" tf:"sort_order,omitempty"`
}

type StorageDescriptorInitParameters struct {

	// List of reducer grouping columns, clustering columns, and bucketing columns in the table.
	BucketColumns []*string `json:"bucketColumns,omitempty" tf:"bucket_columns,omitempty"`

	// Configuration block for columns in the table. See columns below.
	Columns []ColumnsInitParameters `json:"columns,omitempty" tf:"columns,omitempty"`

	// Whether the data in the table is compressed.
	Compressed *bool `json:"compressed,omitempty" tf:"compressed,omitempty"`

	// Input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format.
	InputFormat *string `json:"inputFormat,omitempty" tf:"input_format,omitempty"`

	// Physical location of the table. By default this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Must be specified if the table contains any dimension columns.
	NumberOfBuckets *float64 `json:"numberOfBuckets,omitempty" tf:"number_of_buckets,omitempty"`

	// Output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format.
	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format,omitempty"`

	// User-supplied properties in key-value form.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Object that references a schema stored in the AWS Glue Schema Registry. When creating a table, you can pass an empty list of columns for the schema, and instead use a schema reference. See Schema Reference below.
	SchemaReference []SchemaReferenceInitParameters `json:"schemaReference,omitempty" tf:"schema_reference,omitempty"`

	// Configuration block for serialization and deserialization ("SerDe") information. See ser_de_info below.
	SerDeInfo []SerDeInfoInitParameters `json:"serDeInfo,omitempty" tf:"ser_de_info,omitempty"`

	// Configuration block with information about values that appear very frequently in a column (skewed values). See skewed_info below.
	SkewedInfo []SkewedInfoInitParameters `json:"skewedInfo,omitempty" tf:"skewed_info,omitempty"`

	// Configuration block for the sort order of each bucket in the table. See sort_columns below.
	SortColumns []SortColumnsInitParameters `json:"sortColumns,omitempty" tf:"sort_columns,omitempty"`

	// Whether the table data is stored in subdirectories.
	StoredAsSubDirectories *bool `json:"storedAsSubDirectories,omitempty" tf:"stored_as_sub_directories,omitempty"`
}

type StorageDescriptorObservation struct {

	// List of reducer grouping columns, clustering columns, and bucketing columns in the table.
	BucketColumns []*string `json:"bucketColumns,omitempty" tf:"bucket_columns,omitempty"`

	// Configuration block for columns in the table. See columns below.
	Columns []ColumnsObservation `json:"columns,omitempty" tf:"columns,omitempty"`

	// Whether the data in the table is compressed.
	Compressed *bool `json:"compressed,omitempty" tf:"compressed,omitempty"`

	// Input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format.
	InputFormat *string `json:"inputFormat,omitempty" tf:"input_format,omitempty"`

	// Physical location of the table. By default this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Must be specified if the table contains any dimension columns.
	NumberOfBuckets *float64 `json:"numberOfBuckets,omitempty" tf:"number_of_buckets,omitempty"`

	// Output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format.
	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format,omitempty"`

	// User-supplied properties in key-value form.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Object that references a schema stored in the AWS Glue Schema Registry. When creating a table, you can pass an empty list of columns for the schema, and instead use a schema reference. See Schema Reference below.
	SchemaReference []SchemaReferenceObservation `json:"schemaReference,omitempty" tf:"schema_reference,omitempty"`

	// Configuration block for serialization and deserialization ("SerDe") information. See ser_de_info below.
	SerDeInfo []SerDeInfoObservation `json:"serDeInfo,omitempty" tf:"ser_de_info,omitempty"`

	// Configuration block with information about values that appear very frequently in a column (skewed values). See skewed_info below.
	SkewedInfo []SkewedInfoObservation `json:"skewedInfo,omitempty" tf:"skewed_info,omitempty"`

	// Configuration block for the sort order of each bucket in the table. See sort_columns below.
	SortColumns []SortColumnsObservation `json:"sortColumns,omitempty" tf:"sort_columns,omitempty"`

	// Whether the table data is stored in subdirectories.
	StoredAsSubDirectories *bool `json:"storedAsSubDirectories,omitempty" tf:"stored_as_sub_directories,omitempty"`
}

type StorageDescriptorParameters struct {

	// List of reducer grouping columns, clustering columns, and bucketing columns in the table.
	// +kubebuilder:validation:Optional
	BucketColumns []*string `json:"bucketColumns,omitempty" tf:"bucket_columns,omitempty"`

	// Configuration block for columns in the table. See columns below.
	// +kubebuilder:validation:Optional
	Columns []ColumnsParameters `json:"columns,omitempty" tf:"columns,omitempty"`

	// Whether the data in the table is compressed.
	// +kubebuilder:validation:Optional
	Compressed *bool `json:"compressed,omitempty" tf:"compressed,omitempty"`

	// Input format: SequenceFileInputFormat (binary), or TextInputFormat, or a custom format.
	// +kubebuilder:validation:Optional
	InputFormat *string `json:"inputFormat,omitempty" tf:"input_format,omitempty"`

	// Physical location of the table. By default this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Must be specified if the table contains any dimension columns.
	// +kubebuilder:validation:Optional
	NumberOfBuckets *float64 `json:"numberOfBuckets,omitempty" tf:"number_of_buckets,omitempty"`

	// Output format: SequenceFileOutputFormat (binary), or IgnoreKeyTextOutputFormat, or a custom format.
	// +kubebuilder:validation:Optional
	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format,omitempty"`

	// User-supplied properties in key-value form.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Object that references a schema stored in the AWS Glue Schema Registry. When creating a table, you can pass an empty list of columns for the schema, and instead use a schema reference. See Schema Reference below.
	// +kubebuilder:validation:Optional
	SchemaReference []SchemaReferenceParameters `json:"schemaReference,omitempty" tf:"schema_reference,omitempty"`

	// Configuration block for serialization and deserialization ("SerDe") information. See ser_de_info below.
	// +kubebuilder:validation:Optional
	SerDeInfo []SerDeInfoParameters `json:"serDeInfo,omitempty" tf:"ser_de_info,omitempty"`

	// Configuration block with information about values that appear very frequently in a column (skewed values). See skewed_info below.
	// +kubebuilder:validation:Optional
	SkewedInfo []SkewedInfoParameters `json:"skewedInfo,omitempty" tf:"skewed_info,omitempty"`

	// Configuration block for the sort order of each bucket in the table. See sort_columns below.
	// +kubebuilder:validation:Optional
	SortColumns []SortColumnsParameters `json:"sortColumns,omitempty" tf:"sort_columns,omitempty"`

	// Whether the table data is stored in subdirectories.
	// +kubebuilder:validation:Optional
	StoredAsSubDirectories *bool `json:"storedAsSubDirectories,omitempty" tf:"stored_as_sub_directories,omitempty"`
}

type TargetTableInitParameters struct {

	// Name of the target table.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TargetTableObservation struct {

	// ID of the Data Catalog in which the table resides.
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Name of the catalog database that contains the target table.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Name of the target table.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TargetTableParameters struct {

	// ID of the Data Catalog in which the table resides.
	// +kubebuilder:validation:Required
	CatalogID *string `json:"catalogId" tf:"catalog_id,omitempty"`

	// Name of the catalog database that contains the target table.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// Name of the target table.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// CatalogTableSpec defines the desired state of CatalogTable
type CatalogTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogTableParameters `json:"forProvider"`
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
	InitProvider CatalogTableInitParameters `json:"initProvider,omitempty"`
}

// CatalogTableStatus defines the observed state of CatalogTable.
type CatalogTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogTable is the Schema for the CatalogTables API. Provides a Glue Catalog Table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CatalogTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogTableSpec   `json:"spec"`
	Status            CatalogTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogTableList contains a list of CatalogTables
type CatalogTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogTable `json:"items"`
}

// Repository type metadata.
var (
	CatalogTable_Kind             = "CatalogTable"
	CatalogTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogTable_Kind}.String()
	CatalogTable_KindAPIVersion   = CatalogTable_Kind + "." + CRDGroupVersion.String()
	CatalogTable_GroupVersionKind = CRDGroupVersion.WithKind(CatalogTable_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogTable{}, &CatalogTableList{})
}
