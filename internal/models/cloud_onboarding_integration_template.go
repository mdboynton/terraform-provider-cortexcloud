package models

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/diag"

	cloudAccountsAPI "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/cloud_accounts"
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
    Status types.String `tfsdk:"status"`
    InstanceId types.String `tfsdk:"instance_id"`
    CloudFormationLink types.String `tfsdk:"cloud_formation_link"`
}

func (m *CloudOnboardingIntegrationTemplateModel) RefreshPropertyValues(ctx context.Context, diagnostics *diag.Diagnostics, integrationStatus cloudAccountsAPI.CloudIntegrationInstanceDetailsResponse) {
    integrationData := integrationStatus.Reply

    tags, diags := types.SetValueFrom(ctx, m.CustomResourceTags.ElementType(ctx), integrationData.Tags)
    diagnostics.Append(diags...)
    if diagnostics.HasError() {
        return
    }

    m.InstanceName = types.StringValue(integrationData.InstanceName)
    m.Scope = types.StringValue(integrationData.Scope)
    m.Status = types.StringValue(integrationData.Status)
    m.CustomResourceTags = tags
    m.CloudProvider = types.StringValue(integrationData.CloudProvider)
    m.ScanMode = types.StringValue(integrationData.Scan.ScanMethod)
    // TODO: refresh AdditionalCapabilities with integrationData.AdditionalCapabilities
    // TODO: refresh CollectionConfiguration with integrationData.CollectionConfiguration
    // (need to convert API response values from string to API struct for both)
}
