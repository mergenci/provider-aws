/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta12 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/elb/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/elbv2/v1beta1"
	v1beta13 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta14 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Attachment.
func (mg *Attachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.ForProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AutoscalingGroupName")
	}
	mg.Spec.ForProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ELB),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ELBRef,
		Selector:     mg.Spec.ForProvider.ELBSelector,
		To: reference.To{
			List:    &v1beta1.ELBList{},
			Managed: &v1beta1.ELB{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ELB")
	}
	mg.Spec.ForProvider.ELB = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ELBRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LBTargetGroupArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.LBTargetGroupArnRef,
		Selector:     mg.Spec.ForProvider.LBTargetGroupArnSelector,
		To: reference.To{
			List:    &v1beta11.LBTargetGroupList{},
			Managed: &v1beta11.LBTargetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LBTargetGroupArn")
	}
	mg.Spec.ForProvider.LBTargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LBTargetGroupArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ALBTargetGroupArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.ALBTargetGroupArnRef,
		Selector:     mg.Spec.InitProvider.ALBTargetGroupArnSelector,
		To: reference.To{
			List:    &v1beta1.LBTargetGroupList{},
			Managed: &v1beta1.LBTargetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ALBTargetGroupArn")
	}
	mg.Spec.InitProvider.ALBTargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ALBTargetGroupArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.InitProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AutoscalingGroupName")
	}
	mg.Spec.InitProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ELB),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.ELBRef,
		Selector:     mg.Spec.InitProvider.ELBSelector,
		To: reference.To{
			List:    &v1beta11.ELBList{},
			Managed: &v1beta11.ELB{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ELB")
	}
	mg.Spec.InitProvider.ELB = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ELBRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LBTargetGroupArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.LBTargetGroupArnRef,
		Selector:     mg.Spec.InitProvider.LBTargetGroupArnSelector,
		To: reference.To{
			List:    &v1beta1.LBTargetGroupList{},
			Managed: &v1beta1.LBTargetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LBTargetGroupArn")
	}
	mg.Spec.InitProvider.LBTargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LBTargetGroupArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AutoscalingGroup.
func (mg *AutoscalingGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LaunchConfiguration),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LaunchConfigurationRef,
		Selector:     mg.Spec.ForProvider.LaunchConfigurationSelector,
		To: reference.To{
			List:    &LaunchConfigurationList{},
			Managed: &LaunchConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LaunchConfiguration")
	}
	mg.Spec.ForProvider.LaunchConfiguration = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LaunchConfigurationRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LaunchTemplate); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LaunchTemplate[i3].ID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.LaunchTemplate[i3].IDRef,
			Selector:     mg.Spec.ForProvider.LaunchTemplate[i3].IDSelector,
			To: reference.To{
				List:    &v1beta12.LaunchTemplateList{},
				Managed: &v1beta12.LaunchTemplate{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LaunchTemplate[i3].ID")
		}
		mg.Spec.ForProvider.LaunchTemplate[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LaunchTemplate[i3].IDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.MixedInstancesPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDRef,
					Selector:     mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDSelector,
					To: reference.To{
						List:    &v1beta12.LaunchTemplateList{},
						Managed: &v1beta12.LaunchTemplate{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID")
				}
				mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.MixedInstancesPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID),
						Extract:      resource.ExtractResourceID(),
						Reference:    mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDRef,
						Selector:     mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDSelector,
						To: reference.To{
							List:    &v1beta12.LaunchTemplateList{},
							Managed: &v1beta12.LaunchTemplate{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID")
					}
					mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDRef = rsp.ResolvedReference

				}
			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PlacementGroup),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PlacementGroupRef,
		Selector:     mg.Spec.ForProvider.PlacementGroupSelector,
		To: reference.To{
			List:    &v1beta12.PlacementGroupList{},
			Managed: &v1beta12.PlacementGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PlacementGroup")
	}
	mg.Spec.ForProvider.PlacementGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PlacementGroupRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceLinkedRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ServiceLinkedRoleArnRef,
		Selector:     mg.Spec.ForProvider.ServiceLinkedRoleArnSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceLinkedRoleArn")
	}
	mg.Spec.ForProvider.ServiceLinkedRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceLinkedRoleArnRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCZoneIdentifier),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCZoneIdentifierRefs,
		Selector:      mg.Spec.ForProvider.VPCZoneIdentifierSelector,
		To: reference.To{
			List:    &v1beta12.SubnetList{},
			Managed: &v1beta12.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCZoneIdentifier")
	}
	mg.Spec.ForProvider.VPCZoneIdentifier = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCZoneIdentifierRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LaunchConfiguration),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.LaunchConfigurationRef,
		Selector:     mg.Spec.InitProvider.LaunchConfigurationSelector,
		To: reference.To{
			List:    &LaunchConfigurationList{},
			Managed: &LaunchConfiguration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LaunchConfiguration")
	}
	mg.Spec.InitProvider.LaunchConfiguration = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LaunchConfigurationRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.LaunchTemplate); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LaunchTemplate[i3].ID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.LaunchTemplate[i3].IDRef,
			Selector:     mg.Spec.InitProvider.LaunchTemplate[i3].IDSelector,
			To: reference.To{
				List:    &v1beta12.LaunchTemplateList{},
				Managed: &v1beta12.LaunchTemplate{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.LaunchTemplate[i3].ID")
		}
		mg.Spec.InitProvider.LaunchTemplate[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.LaunchTemplate[i3].IDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.MixedInstancesPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDRef,
					Selector:     mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDSelector,
					To: reference.To{
						List:    &v1beta12.LaunchTemplateList{},
						Managed: &v1beta12.LaunchTemplate{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID")
				}
				mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].LaunchTemplateSpecification[i5].LaunchTemplateIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.MixedInstancesPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID),
						Extract:      resource.ExtractResourceID(),
						Reference:    mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDRef,
						Selector:     mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDSelector,
						To: reference.To{
							List:    &v1beta12.LaunchTemplateList{},
							Managed: &v1beta12.LaunchTemplate{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID")
					}
					mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateID = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.MixedInstancesPolicy[i3].LaunchTemplate[i4].Override[i5].LaunchTemplateSpecification[i6].LaunchTemplateIDRef = rsp.ResolvedReference

				}
			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PlacementGroup),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.PlacementGroupRef,
		Selector:     mg.Spec.InitProvider.PlacementGroupSelector,
		To: reference.To{
			List:    &v1beta12.PlacementGroupList{},
			Managed: &v1beta12.PlacementGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PlacementGroup")
	}
	mg.Spec.InitProvider.PlacementGroup = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PlacementGroupRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceLinkedRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.ServiceLinkedRoleArnRef,
		Selector:     mg.Spec.InitProvider.ServiceLinkedRoleArnSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceLinkedRoleArn")
	}
	mg.Spec.InitProvider.ServiceLinkedRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceLinkedRoleArnRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.VPCZoneIdentifier),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.VPCZoneIdentifierRefs,
		Selector:      mg.Spec.InitProvider.VPCZoneIdentifierSelector,
		To: reference.To{
			List:    &v1beta12.SubnetList{},
			Managed: &v1beta12.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCZoneIdentifier")
	}
	mg.Spec.InitProvider.VPCZoneIdentifier = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.VPCZoneIdentifierRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this GroupTag.
func (mg *GroupTag) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.ForProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AutoscalingGroupName")
	}
	mg.Spec.ForProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.InitProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AutoscalingGroupName")
	}
	mg.Spec.InitProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LifecycleHook.
func (mg *LifecycleHook) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.ForProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AutoscalingGroupName")
	}
	mg.Spec.ForProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.InitProvider.RoleArnRef,
		Selector:     mg.Spec.InitProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta13.RoleList{},
			Managed: &v1beta13.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RoleArn")
	}
	mg.Spec.InitProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Notification.
func (mg *Notification) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TopicArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.TopicArnRef,
		Selector:     mg.Spec.ForProvider.TopicArnSelector,
		To: reference.To{
			List:    &v1beta14.TopicList{},
			Managed: &v1beta14.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TopicArn")
	}
	mg.Spec.ForProvider.TopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TopicArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TopicArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.InitProvider.TopicArnRef,
		Selector:     mg.Spec.InitProvider.TopicArnSelector,
		To: reference.To{
			List:    &v1beta14.TopicList{},
			Managed: &v1beta14.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TopicArn")
	}
	mg.Spec.InitProvider.TopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TopicArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Policy.
func (mg *Policy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.ForProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AutoscalingGroupName")
	}
	mg.Spec.ForProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Schedule.
func (mg *Schedule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.ForProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AutoscalingGroupName")
	}
	mg.Spec.ForProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	return nil
}
