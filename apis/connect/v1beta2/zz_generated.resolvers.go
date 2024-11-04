// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *BotAssociation) ResolveReferences( // ResolveReferences of this BotAssociation.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.LexBot != nil {
		{
			m, l, err = apisresolver.GetManagedResource("lexmodels.aws.upbound.io", "v1beta2", "Bot", "BotList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LexBot.Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.LexBot.NameRef,
				Selector:     mg.Spec.ForProvider.LexBot.NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LexBot.Name")
		}
		mg.Spec.ForProvider.LexBot.Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LexBot.NameRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.LexBot != nil {
		{
			m, l, err = apisresolver.GetManagedResource("lexmodels.aws.upbound.io", "v1beta2", "Bot", "BotList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LexBot.Name),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.LexBot.NameRef,
				Selector:     mg.Spec.InitProvider.LexBot.NameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.LexBot.Name")
		}
		mg.Spec.InitProvider.LexBot.Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.LexBot.NameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this HoursOfOperation.
func (mg *HoursOfOperation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InstanceStorageConfig.
func (mg *InstanceStorageConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference

	if mg.Spec.ForProvider.StorageConfig != nil {
		if mg.Spec.ForProvider.StorageConfig.KinesisFirehoseConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta2", "DeliveryStream", "DeliveryStreamList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArn),
					Extract:      resource.ExtractParamPath("arn", false),
					Reference:    mg.Spec.ForProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArnRef,
					Selector:     mg.Spec.ForProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArn")
			}
			mg.Spec.ForProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.StorageConfig != nil {
		if mg.Spec.ForProvider.StorageConfig.KinesisStreamConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kinesis.aws.upbound.io", "v1beta2", "Stream", "StreamList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig.KinesisStreamConfig.StreamArn),
					Extract:      resource.ExtractParamPath("arn", false),
					Reference:    mg.Spec.ForProvider.StorageConfig.KinesisStreamConfig.StreamArnRef,
					Selector:     mg.Spec.ForProvider.StorageConfig.KinesisStreamConfig.StreamArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig.KinesisStreamConfig.StreamArn")
			}
			mg.Spec.ForProvider.StorageConfig.KinesisStreamConfig.StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.StorageConfig.KinesisStreamConfig.StreamArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.StorageConfig != nil {
		if mg.Spec.ForProvider.StorageConfig.KinesisVideoStreamConfig != nil {
			if mg.Spec.ForProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig != nil {
				{
					m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyID),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyIDRef,
						Selector:     mg.Spec.ForProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyID")
				}
				mg.Spec.ForProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyIDRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.ForProvider.StorageConfig != nil {
		if mg.Spec.ForProvider.StorageConfig.S3Config != nil {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig.S3Config.BucketName),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.StorageConfig.S3Config.BucketNameRef,
					Selector:     mg.Spec.ForProvider.StorageConfig.S3Config.BucketNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig.S3Config.BucketName")
			}
			mg.Spec.ForProvider.StorageConfig.S3Config.BucketName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.StorageConfig.S3Config.BucketNameRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.ForProvider.StorageConfig != nil {
		if mg.Spec.ForProvider.StorageConfig.S3Config != nil {
			if mg.Spec.ForProvider.StorageConfig.S3Config.EncryptionConfig != nil {
				{
					m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageConfig.S3Config.EncryptionConfig.KeyID),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.StorageConfig.S3Config.EncryptionConfig.KeyIDRef,
						Selector:     mg.Spec.ForProvider.StorageConfig.S3Config.EncryptionConfig.KeyIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.StorageConfig.S3Config.EncryptionConfig.KeyID")
				}
				mg.Spec.ForProvider.StorageConfig.S3Config.EncryptionConfig.KeyID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.StorageConfig.S3Config.EncryptionConfig.KeyIDRef = rsp.ResolvedReference

			}
		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.StorageConfig != nil {
		if mg.Spec.InitProvider.StorageConfig.KinesisFirehoseConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("firehose.aws.upbound.io", "v1beta2", "DeliveryStream", "DeliveryStreamList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArn),
					Extract:      resource.ExtractParamPath("arn", false),
					Reference:    mg.Spec.InitProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArnRef,
					Selector:     mg.Spec.InitProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArn")
			}
			mg.Spec.InitProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.StorageConfig.KinesisFirehoseConfig.FirehoseArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.StorageConfig != nil {
		if mg.Spec.InitProvider.StorageConfig.KinesisStreamConfig != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kinesis.aws.upbound.io", "v1beta2", "Stream", "StreamList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageConfig.KinesisStreamConfig.StreamArn),
					Extract:      resource.ExtractParamPath("arn", false),
					Reference:    mg.Spec.InitProvider.StorageConfig.KinesisStreamConfig.StreamArnRef,
					Selector:     mg.Spec.InitProvider.StorageConfig.KinesisStreamConfig.StreamArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.StorageConfig.KinesisStreamConfig.StreamArn")
			}
			mg.Spec.InitProvider.StorageConfig.KinesisStreamConfig.StreamArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.StorageConfig.KinesisStreamConfig.StreamArnRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.StorageConfig != nil {
		if mg.Spec.InitProvider.StorageConfig.KinesisVideoStreamConfig != nil {
			if mg.Spec.InitProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig != nil {
				{
					m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyID),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.InitProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyIDRef,
						Selector:     mg.Spec.InitProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyID")
				}
				mg.Spec.InitProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.StorageConfig.KinesisVideoStreamConfig.EncryptionConfig.KeyIDRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.InitProvider.StorageConfig != nil {
		if mg.Spec.InitProvider.StorageConfig.S3Config != nil {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageConfig.S3Config.BucketName),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.StorageConfig.S3Config.BucketNameRef,
					Selector:     mg.Spec.InitProvider.StorageConfig.S3Config.BucketNameSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.StorageConfig.S3Config.BucketName")
			}
			mg.Spec.InitProvider.StorageConfig.S3Config.BucketName = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.StorageConfig.S3Config.BucketNameRef = rsp.ResolvedReference

		}
	}
	if mg.Spec.InitProvider.StorageConfig != nil {
		if mg.Spec.InitProvider.StorageConfig.S3Config != nil {
			if mg.Spec.InitProvider.StorageConfig.S3Config.EncryptionConfig != nil {
				{
					m, l, err = apisresolver.GetManagedResource("kms.aws.upbound.io", "v1beta1", "Key", "KeyList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageConfig.S3Config.EncryptionConfig.KeyID),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.InitProvider.StorageConfig.S3Config.EncryptionConfig.KeyIDRef,
						Selector:     mg.Spec.InitProvider.StorageConfig.S3Config.EncryptionConfig.KeyIDSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.StorageConfig.S3Config.EncryptionConfig.KeyID")
				}
				mg.Spec.InitProvider.StorageConfig.S3Config.EncryptionConfig.KeyID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.StorageConfig.S3Config.EncryptionConfig.KeyIDRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}

// ResolveReferences of this Queue.
func (mg *Queue) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta2", "HoursOfOperation", "HoursOfOperationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HoursOfOperationID),
			Extract:      resource.ExtractParamPath("hours_of_operation_id", true),
			Reference:    mg.Spec.ForProvider.HoursOfOperationIDRef,
			Selector:     mg.Spec.ForProvider.HoursOfOperationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HoursOfOperationID")
	}
	mg.Spec.ForProvider.HoursOfOperationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HoursOfOperationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta2", "HoursOfOperation", "HoursOfOperationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HoursOfOperationID),
			Extract:      resource.ExtractParamPath("hours_of_operation_id", true),
			Reference:    mg.Spec.InitProvider.HoursOfOperationIDRef,
			Selector:     mg.Spec.InitProvider.HoursOfOperationIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HoursOfOperationID")
	}
	mg.Spec.InitProvider.HoursOfOperationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HoursOfOperationIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this QuickConnect.
func (mg *QuickConnect) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoutingProfile.
func (mg *RoutingProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta3", "Queue", "QueueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultOutboundQueueID),
			Extract:      resource.ExtractParamPath("queue_id", true),
			Reference:    mg.Spec.ForProvider.DefaultOutboundQueueIDRef,
			Selector:     mg.Spec.ForProvider.DefaultOutboundQueueIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DefaultOutboundQueueID")
	}
	mg.Spec.ForProvider.DefaultOutboundQueueID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DefaultOutboundQueueIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta3", "Queue", "QueueList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DefaultOutboundQueueID),
			Extract:      resource.ExtractParamPath("queue_id", true),
			Reference:    mg.Spec.InitProvider.DefaultOutboundQueueIDRef,
			Selector:     mg.Spec.InitProvider.DefaultOutboundQueueIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DefaultOutboundQueueID")
	}
	mg.Spec.InitProvider.DefaultOutboundQueueID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DefaultOutboundQueueIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta2", "RoutingProfile", "RoutingProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoutingProfileID),
			Extract:      resource.ExtractParamPath("routing_profile_id", true),
			Reference:    mg.Spec.ForProvider.RoutingProfileIDRef,
			Selector:     mg.Spec.ForProvider.RoutingProfileIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoutingProfileID")
	}
	mg.Spec.ForProvider.RoutingProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoutingProfileIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "SecurityProfile", "SecurityProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityProfileIds),
			Extract:       resource.ExtractParamPath("security_profile_id", true),
			References:    mg.Spec.ForProvider.SecurityProfileIdsRefs,
			Selector:      mg.Spec.ForProvider.SecurityProfileIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityProfileIds")
	}
	mg.Spec.ForProvider.SecurityProfileIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityProfileIdsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta2", "RoutingProfile", "RoutingProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoutingProfileID),
			Extract:      resource.ExtractParamPath("routing_profile_id", true),
			Reference:    mg.Spec.InitProvider.RoutingProfileIDRef,
			Selector:     mg.Spec.InitProvider.RoutingProfileIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoutingProfileID")
	}
	mg.Spec.InitProvider.RoutingProfileID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoutingProfileIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "SecurityProfile", "SecurityProfileList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SecurityProfileIds),
			Extract:       resource.ExtractParamPath("security_profile_id", true),
			References:    mg.Spec.InitProvider.SecurityProfileIdsRefs,
			Selector:      mg.Spec.InitProvider.SecurityProfileIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityProfileIds")
	}
	mg.Spec.InitProvider.SecurityProfileIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SecurityProfileIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this UserHierarchyStructure.
func (mg *UserHierarchyStructure) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceIDRef,
			Selector:     mg.Spec.ForProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceID")
	}
	mg.Spec.ForProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("connect.aws.upbound.io", "v1beta1", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceIDRef,
			Selector:     mg.Spec.InitProvider.InstanceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceID")
	}
	mg.Spec.InitProvider.InstanceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceIDRef = rsp.ResolvedReference

	return nil
}
