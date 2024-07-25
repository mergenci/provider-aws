// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package eks

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the eks group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_eks_cluster", func(r *config.Resource) {
		r.References = config.References{
			"role_arn": {
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
			"vpc_config.subnet_ids": {
				TerraformName:     "aws_subnet",
				RefFieldName:      "SubnetIDRefs",
				SelectorFieldName: "SubnetIDSelector",
			},
			"vpc_config.security_group_ids": {
				TerraformName:     "aws_security_group",
				RefFieldName:      "SecurityGroupIDRefs",
				SelectorFieldName: "SecurityGroupIDSelector",
			},
		}
		r.UseAsync = true
		// r.Conversions = append(r.Conversions,
		// 	conversion.NewCustomConverter("v1beta1", "v1beta2", func(src, target resource.Managed) error {
		// 		annotations := src.GetAnnotations()
		// 		if annotations == nil {
		// 			return nil
		// 		}

		// 		fieldName := "bootstrapSelfManagedAddons"
		// 		annotationKey := "conversion.upjet.crossplane.io/" + fieldName
		// 		jsonString, ok := annotations[annotationKey]
		// 		if !ok {
		// 			return nil
		// 		}

		// 		var annotationValue struct {
		// 			ForProvider  *bool `json:"spec.forProvider,omitempty"`
		// 			InitProvider *bool `json:"spec.initProvider,omitempty"`
		// 			AtProvider   *bool `json:"status.atProvider,omitempty"`
		// 		}
		// 		if err := json.Unmarshal([]byte(jsonString), &annotationValue); err != nil {
		// 			return errors.Wrap(err, "cannot unmarshal annotation as JSON")
		// 		}

		// 		targetTyped := target.(*v1beta2.Cluster)
		// 		if annotationValue.ForProvider != nil {
		// 			targetTyped.Spec.ForProvider.BootstrapSelfManagedAddons = annotationValue.ForProvider
		// 		}
		// 		if annotationValue.InitProvider != nil {
		// 			targetTyped.Spec.InitProvider.BootstrapSelfManagedAddons = annotationValue.InitProvider
		// 		}
		// 		if annotationValue.AtProvider != nil {
		// 			targetTyped.Status.AtProvider.BootstrapSelfManagedAddons = annotationValue.AtProvider
		// 		}

		// 		// Removing the annotation is necessary. Consider that user
		// 		// deletes the field, but leaves the annotation intact.
		// 		// Converting back to storage version will preserve the
		// 		// annotation. When a non-storage version is requested,
		// 		// conversion function will find the annotation and will
		// 		// convert it to the field, which is supposed to be deleted.
		// 		meta.RemoveAnnotations(target, annotationKey)

		// 		return nil
		// 	}),
		// 	conversion.NewCustomConverter("v1beta2", "v1beta1", func(src, target resource.Managed) error {
		// 		fieldName := "bootstrapSelfManagedAddons"
		// 		srcTyped := src.(*v1beta2.Cluster)

		// 		annotation := make(map[string]any)
		// 		if srcTyped.Spec.ForProvider.BootstrapSelfManagedAddons != nil {
		// 			annotation["spec.forProvider"] = srcTyped.Spec.ForProvider.BootstrapSelfManagedAddons
		// 		}
		// 		if srcTyped.Spec.InitProvider.BootstrapSelfManagedAddons != nil {
		// 			annotation["spec.initProvider"] = srcTyped.Spec.InitProvider.BootstrapSelfManagedAddons
		// 		}
		// 		if srcTyped.Status.AtProvider.BootstrapSelfManagedAddons != nil {
		// 			annotation["status.atProvider"] = srcTyped.Status.AtProvider.BootstrapSelfManagedAddons
		// 		}

		// 		if len(annotation) == 0 {
		// 			return nil
		// 		}

		// 		jsonBytes, err := json.Marshal(annotation)
		// 		if err != nil {
		// 			return errors.Wrap(err, "cannot marshal subnetConfiguration to JSON")
		// 		}

		// 		annotations := map[string]string{
		// 			"conversion.upjet.crossplane.io/" + fieldName: string(jsonBytes),
		// 		}
		// 		meta.AddAnnotations(target, annotations)

		// 		return nil
		// 	}),
		// )
	})
	p.AddResourceConfigurator("aws_eks_node_group", func(r *config.Resource) {
		r.References["cluster_name"] = config.Reference{
			TerraformName: "aws_eks_cluster",
			Extractor:     "ExternalNameIfClusterActive()",
		}
		r.References["node_role_arn"] = config.Reference{
			TerraformName: "aws_iam_role",
			Extractor:     common.PathARNExtractor,
		}
		r.References["remote_access.source_security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SourceSecurityGroupIDRefs",
			SelectorFieldName: "SourceSecurityGroupIDSelector",
		}
		r.References["subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"release_version",
				"version",
			},
		}
		r.UseAsync = true
		r.MetaResource.ArgumentDocs["launch_template.version"] = `- (Required) EC2 Launch Template version number.`
		r.MetaResource.ArgumentDocs["subnet_ids"] = `- Identifiers of EC2 Subnets to associate with the EKS Node Group. Amazon EKS managed node groups can be launched in both public and private subnets. If you plan to deploy load balancers to a subnet, the private subnet must have tag kubernetes.io/role/internal-elb, the public subnet must have tag kubernetes.io/role/elb.`
	})
	p.AddResourceConfigurator("aws_eks_identity_provider_config", func(r *config.Resource) {
		r.Version = common.VersionV1Beta1
		// OmittedFields in config.ExternalName works only for the top-level fields.
		r.References = config.References{
			"cluster_name": {
				TerraformName: "aws_eks_cluster",
			},
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_eks_fargate_profile", func(r *config.Resource) {
		r.References = config.References{
			"cluster_name": {
				TerraformName: "aws_eks_cluster",
			},
			"pod_execution_role_arn": {
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
			"subnet_ids": {
				TerraformName:     "aws_subnet",
				RefFieldName:      "SubnetIDRefs",
				SelectorFieldName: "SubnetIDSelector",
			},
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_eks_addon", func(r *config.Resource) {
		r.References = config.References{
			"cluster_name": {
				TerraformName: "aws_eks_cluster",
			},
			"service_account_role_arn": {
				TerraformName: "aws_iam_role",
				Extractor:     common.PathARNExtractor,
			},
		}
		r.UseAsync = true
	})
}
