package models

import (
    "time"

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
    AccountName types.String `tfsdk:"account_name"`
    OutpostId types.String `tfsdk:"outpost_id"`
    CreationTime types.String `tfsdk:"creation_time"`
    CloudFormationLink types.String `tfsdk:"cloud_formation_link"`
}

func (m *CloudOnboardingIntegrationTemplateModel) RefreshPropertyValues(diagnostics *diag.Diagnostics, response cloudAccountsAPI.CloudIntegrationInstancesResponse, instanceId, cloudFormationLink *string) {
    if instanceId != nil {
        m.InstanceId = types.StringValue(*instanceId)
    }

    if cloudFormationLink != nil {
        m.CloudFormationLink= types.StringValue(*cloudFormationLink)
    }

    if (len(response.Reply.Data) == 0 || len(response.Reply.Data) > 1) {
        m.Status = types.StringNull()
        m.InstanceName = types.StringNull()
        m.AccountName = types.StringNull()
        m.OutpostId = types.StringNull()
        m.CreationTime = types.StringNull()
    
        diagnostics.AddWarning(
            "Integration Status Unknown",
            "Unable to retrieve computed values for the following arguments " +
            "from the Cortex Cloud API: status, instance_name, account_name, " +
            "outpost_id, creation_time\n\n" +
            "The provider will attempt to populate these arguments during " +
            "the next terraform apply operation.",
        )

        return
    }

    integrationData := response.Reply.Data[0]

    if integrationData.Status != "" {
        m.Status = types.StringValue(integrationData.Status)
    } else {
        m.Status = types.StringNull()
    }

    if integrationData.InstanceName != "" {
        m.InstanceName = types.StringValue(integrationData.InstanceName)
    } else {
        m.InstanceName = types.StringNull()
    }

    if integrationData.AccountName != "" {
        m.AccountName = types.StringValue(integrationData.AccountName)
    } else {
        m.AccountName = types.StringNull()
    }

    if integrationData.OutpostId != "" {
        m.OutpostId = types.StringValue(integrationData.OutpostId)
    } else {
        m.OutpostId = types.StringNull()
    }

    if integrationData.CreationTime != 0 {
        m.CreationTime = types.StringValue(time.UnixMilli(int64(integrationData.CreationTime)).Format("2006-01-02T15:04:05.000Z"))
    } else {
        m.CreationTime = types.StringNull()
    }

    // TODO: refresh Tags
    
    // TODO: refresh AdditionalCapabilities with integrationData.AdditionalCapabilities
    // TODO: refresh CollectionConfiguration with integrationData.CollectionConfiguration
    // (need to convert API response values from string to API struct for both)
}
