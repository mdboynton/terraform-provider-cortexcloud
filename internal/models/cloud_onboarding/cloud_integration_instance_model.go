// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	//"github.com/hashicorp/terraform-plugin-log/tflog"

	api "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/cloud_onboarding/cloud_integration"
)

type CloudIntegrationInstanceModel struct {
	AccountDetails          types.Object `tfsdk:"account_details"`
	AdditionalCapabilities  types.Object `tfsdk:"additional_capabilities"`
	CloudProvider           types.String `tfsdk:"cloud_provider"`
	CollectionConfiguration types.Object `tfsdk:"collection_configuration"`
	CustomResourceTags      types.Set    `tfsdk:"custom_resource_tags"`
	InstanceName            types.String `tfsdk:"instance_name"`
	ScanMode                types.String `tfsdk:"scan_mode"`
	Scope                   types.String `tfsdk:"scope"`
	ScopeModifications      types.Object `tfsdk:"scope_modifications"`
	Status                  types.String `tfsdk:"status"`
	InstanceId              types.String `tfsdk:"instance_id"`
	AccountName             types.String `tfsdk:"account_name"`
	OutpostId               types.String `tfsdk:"outpost_id"`
	CreationTime            types.String `tfsdk:"creation_time"`
	CloudFormationLink      types.String `tfsdk:"cloud_formation_link"`
	TemplateInstanceId      types.String `tfsdk:"template_instance_id"`
}

func (m *CloudIntegrationInstanceModel) ToCreateRequest(ctx context.Context, diagnostics *diag.Diagnostics) api.CreateCloudOnboardingIntegrationTemplateRequest {
	nullAsEmptyOpts := basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty: true,
	}

	accountDetails := api.CloudIntegrationAccountDetails{}
	diagnostics.Append(m.AccountDetails.As(ctx, &accountDetails, nullAsEmptyOpts)...)

	additionalCapabilities := api.CloudIntegrationAdditionalCapabilities{}
	diagnostics.Append(m.AdditionalCapabilities.As(ctx, &additionalCapabilities, basetypes.ObjectAsOptions{})...)

	collectionConfiguration := api.CloudIntegrationCollectionConfiguration{}
	diagnostics.Append(m.CollectionConfiguration.As(ctx, &collectionConfiguration, basetypes.ObjectAsOptions{})...)

	customResourcesTags := []api.CloudIntegrationCustomResourcesTag{}
	diagnostics.Append(m.CustomResourceTags.ElementsAs(ctx, &customResourcesTags, false)...)

	scopeModifications := api.CloudIntegrationScopeModifications{}
	diagnostics.Append(m.ScopeModifications.As(ctx, &scopeModifications, basetypes.ObjectAsOptions{})...)

	if diagnostics.HasError() {
		return api.CreateCloudOnboardingIntegrationTemplateRequest{}
	}

	request := api.CreateCloudOnboardingIntegrationTemplateRequest{
		RequestData: api.CreateCloudOnboardingIntegrationTemplateRequestData{
			AdditionalCapabilities:  additionalCapabilities,
			CloudProvider:           m.CloudProvider.ValueString(),
			CollectionConfiguration: collectionConfiguration,
			CustomResourcesTags:     customResourcesTags,
			InstanceName:            m.InstanceName.ValueString(),
			ScanMode:                m.ScanMode.ValueString(),
			Scope:                   m.Scope.ValueString(),
			ScopeModifications:      scopeModifications,
		},
	}

	if !m.AccountDetails.IsNull() {
		request.RequestData.AccountDetails = &accountDetails
	} else {
		request.RequestData.AccountDetails = nil
	}

	return request
}

func (m *CloudIntegrationInstanceModel) ToGetRequest(ctx context.Context, diagnostics *diag.Diagnostics, instanceId *string) api.CloudIntegrationInstancesRequest {
	var _instanceId string
	if instanceId != nil {
		_instanceId = *instanceId
	} else {
		_instanceId = m.InstanceId.ValueString()
	}

	andFilters := []api.CloudIntegrationInstancesAndFilter{
		{
			SearchField: "ID",
			SearchType:  "EQ",
			SearchValue: _instanceId,
		},
	}

	// Check if the status attribute is null or unknown
	statusIsNullOrUnknown := (m.Status.IsNull() || m.Status.IsUnknown())
	if statusIsNullOrUnknown {
		andFilters = append(andFilters, api.CloudIntegrationInstancesAndFilter{
			SearchField: "STATUS",
			SearchType:  "EQ",
			SearchValue: "PENDING",
		})
	} else {
		if m.CloudProvider.ValueString() == CloudIntegrationCloudProviderEnumAws {
			andFilters = append(andFilters, api.CloudIntegrationInstancesAndFilter{
				SearchField: "PROVISIONING_METHOD",
				SearchType:  "EQ",
				SearchValue: "CF",
			})
		}
	}

	return api.CloudIntegrationInstancesRequest{
		RequestData: api.CloudIntegrationInstancesRequestData{
			FilterData: api.CloudIntegrationInstancesFilterData{
				Filter: api.CloudIntegrationInstancesFilter{
					And: andFilters,
				},
				Paging: api.CloudIntegrationInstancesPaging{
					From: 0,
					To:   1000,
				},
			},
		},
	}
}

func (m *CloudIntegrationInstanceModel) ToUpdateRequest(ctx context.Context, diagnostics *diag.Diagnostics) api.CloudIntegrationEditRequest {
	additionalCapabilities := api.CloudIntegrationAdditionalCapabilities{}
	diagnostics.Append(m.AdditionalCapabilities.As(ctx, &additionalCapabilities, basetypes.ObjectAsOptions{})...)

	collectionConfiguration := api.CloudIntegrationCollectionConfiguration{}
	diagnostics.Append(m.CollectionConfiguration.As(ctx, &collectionConfiguration, basetypes.ObjectAsOptions{})...)

	customResourcesTags := []api.CloudIntegrationCustomResourcesTag{}
	diagnostics.Append(m.CustomResourceTags.ElementsAs(ctx, &customResourcesTags, false)...)

	scopeModifications := api.CloudIntegrationScopeModifications{}
	diagnostics.Append(m.ScopeModifications.As(ctx, &scopeModifications, basetypes.ObjectAsOptions{})...)

	if diagnostics.HasError() {
		return api.CloudIntegrationEditRequest{}
	}

	return api.CloudIntegrationEditRequest{
		RequestData: api.CloudIntegrationEditRequestData{
			AdditionalCapabilities:  additionalCapabilities,
			CloudProvider:           m.CloudProvider.ValueString(),
			CollectionConfiguration: collectionConfiguration,
			CustomResourcesTags:     customResourcesTags,
			InstanceId:              m.InstanceId.ValueString(),
			InstanceName:            m.InstanceName.ValueString(),
			ScopeModifications:      scopeModifications,
			ScanEnvId:               m.OutpostId.ValueString(),
		},
	}
}

func (m *CloudIntegrationInstanceModel) ToDeleteRequest(ctx context.Context, diagnostics *diag.Diagnostics) api.CloudIntegrationDeleteRequest {
	return api.CloudIntegrationDeleteRequest{
		RequestData: api.CloudIntegrationDeleteRequestData{
			Ids: []string{m.InstanceId.ValueString()},
		},
	}
}

// func (m *CloudIntegrationInstanceModel) RefreshPropertyValues(diagnostics *diag.Diagnostics, response api.CloudIntegrationInstancesResponse, instanceId, cloudFormationLink *string) {
// TEMPORARY
func (m *CloudIntegrationInstanceModel) RefreshPropertyValues(diagnostics *diag.Diagnostics, response api.CloudIntegrationInstancesResponse, instanceId, cloudFormationLink *string, isCreate bool) {
	// END TEMPORARY
	if instanceId != nil {
		m.InstanceId = types.StringValue(*instanceId)
	}

	if cloudFormationLink != nil {
		m.CloudFormationLink = types.StringValue(*cloudFormationLink)
	}

	// TEMPORARY
	if isCreate {
		m.TemplateInstanceId = types.StringValue(*instanceId)
	}
	// END TEMPORARY

	if len(response.Reply.Data) == 0 || len(response.Reply.Data) > 1 {
		m.Status = types.StringNull()
		m.InstanceName = types.StringNull()
		m.AccountName = types.StringNull()
		m.OutpostId = types.StringNull()
		m.CreationTime = types.StringNull()

		if len(response.Reply.Data) == 0 {
			diagnostics.AddWarning(
				"Integration Status Unknown",
				"Unable to retrieve computed values for the following arguments "+
					"from the Cortex Cloud API: status, instance_name, account_name, "+
					"outpost_id, creation_time\n\n"+
					"The provider will attempt to populate these arguments during "+
					"the next terraform apply operation.",
			)
		} else if len(response.Reply.Data) > 1 {
			diagnostics.AddWarning(
				"Integration Status Unknown",
				"Multiple values returned for the following arguments: "+
					"status, instance_name, account_name, outpost_id, creation_time\n\n"+
					"The provider will attempt to populate these arguments during "+
					"the next terraform refresh or apply operation.",
			)
		}

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
