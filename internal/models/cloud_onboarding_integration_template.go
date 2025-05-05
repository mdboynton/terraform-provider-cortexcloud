package models

import (
    "github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudOnboardingIntegrationTemplateModel struct {
    AdditionalCapabilities types.Object `tfsdk:"additional_capabilities"`
    CloudProvider types.String `tfsdk:"cloud_provider"`
    CollectionConfiguration types.Object `tfsdk:"collection_configuration"`
    CustomResourceTags types.Set `tfsdk:"custom_resource_tags"`
    InstanceName types.String `tfsdk:"instance_name"`
    ScanMode types.String `tfsdk:"scan_mode"`
    Scope types.String `tfsdk:"scope"`
    ScopeModifications types.Object `tfsdk:"scope_modifications"`
    CloudFormationLink types.String `tfsdk:"cloud_formation_link"`
}
