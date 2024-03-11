/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Cluster.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.BrokerNodeGroupInfo); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnets),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnetsRefs,
				Selector:      mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnetsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnets")
		}
		mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnets = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].ClientSubnetsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.BrokerNodeGroupInfo); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroups),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroupsRefs,
				Selector:      mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroupsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroups")
		}
		mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.BrokerNodeGroupInfo[i3].SecurityGroupsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ConfigurationInfo); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kafka.aws.upbound.io", "v1beta1", "Configuration", "ConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConfigurationInfo[i3].Arn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.ConfigurationInfo[i3].ArnRef,
				Selector:     mg.Spec.ForProvider.ConfigurationInfo[i3].ArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ConfigurationInfo[i3].Arn")
		}
		mg.Spec.ForProvider.ConfigurationInfo[i3].Arn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ConfigurationInfo[i3].ArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EncryptionInfo); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArnRef,
				Selector:     mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArn")
		}
		mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LoggingInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("cloudwatchlogs.aws.upbound.io", "v1beta1", "Group", "GroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroup),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroupRef,
						Selector:     mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroupSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroup")
				}
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroup = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroupRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LoggingInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta1", "DeliveryStream", "DeliveryStreamList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStream),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStreamRef,
						Selector:     mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStreamSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStream")
				}
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStream = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStreamRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.LoggingInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].Bucket),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].BucketRef,
						Selector:     mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].BucketSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].Bucket")
				}
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].BucketRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.BrokerNodeGroupInfo); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].ClientSubnets),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].ClientSubnetsRefs,
				Selector:      mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].ClientSubnetsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].ClientSubnets")
		}
		mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].ClientSubnets = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].ClientSubnetsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.BrokerNodeGroupInfo); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].SecurityGroups),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].SecurityGroupsRefs,
				Selector:      mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].SecurityGroupsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].SecurityGroups")
		}
		mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.BrokerNodeGroupInfo[i3].SecurityGroupsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.ConfigurationInfo); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kafka.aws.upbound.io", "v1beta1", "Configuration", "ConfigurationList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConfigurationInfo[i3].Arn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.InitProvider.ConfigurationInfo[i3].ArnRef,
				Selector:     mg.Spec.InitProvider.ConfigurationInfo[i3].ArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.ConfigurationInfo[i3].Arn")
		}
		mg.Spec.InitProvider.ConfigurationInfo[i3].Arn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.ConfigurationInfo[i3].ArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EncryptionInfo); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.InitProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArnRef,
				Selector:     mg.Spec.InitProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArn")
		}
		mg.Spec.InitProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EncryptionInfo[i3].EncryptionAtRestKMSKeyArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.LoggingInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("cloudwatchlogs.aws.upbound.io", "v1beta1", "Group", "GroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroup),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroupRef,
						Selector:     mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroupSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroup")
				}
				mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroup = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].CloudwatchLogs[i5].LogGroupRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.LoggingInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta1", "DeliveryStream", "DeliveryStreamList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStream),
						Extract:      resource.ExtractParamPath("name", false),
						Reference:    mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStreamRef,
						Selector:     mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStreamSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStream")
				}
				mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStream = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].Firehose[i5].DeliveryStreamRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.LoggingInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].S3); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].Bucket),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].BucketRef,
						Selector:     mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].BucketSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].Bucket")
				}
				mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.LoggingInfo[i3].BrokerLogs[i4].S3[i5].BucketRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}