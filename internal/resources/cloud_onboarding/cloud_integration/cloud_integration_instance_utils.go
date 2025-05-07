package resources

import (
	"context"
	//"fmt"
	"net/url"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
	cloudIntegrationInstancesAPI "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/cloud_onboarding/cloud_integration"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func createCloudOnboardingIntegrationTemplate(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, plan models.CloudIntegrationInstanceModel) cloudIntegrationInstancesAPI.CreateCloudOnboardingIntegrationTemplateResponse {
	var response cloudIntegrationInstancesAPI.CreateCloudOnboardingIntegrationTemplateResponse

	additionalCapabilities := cloudIntegrationInstancesAPI.CloudIntegrationAdditionalCapabilities{}
	diagnostics.Append(plan.AdditionalCapabilities.As(ctx, &additionalCapabilities, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	collectionConfiguration := cloudIntegrationInstancesAPI.CloudIntegrationCollectionConfiguration{}
	diagnostics.Append(plan.CollectionConfiguration.As(ctx, &collectionConfiguration, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	customResourcesTags := []cloudIntegrationInstancesAPI.CloudIntegrationCustomResourcesTag{}
	diagnostics.Append(plan.CustomResourceTags.ElementsAs(ctx, &customResourcesTags, false)...)
	if diagnostics.HasError() {
		return response
	}

	scopeModifications := cloudIntegrationInstancesAPI.CloudIntegrationScopeModifications{}
	diagnostics.Append(plan.ScopeModifications.As(ctx, &scopeModifications, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	request := cloudIntegrationInstancesAPI.CreateCloudOnboardingIntegrationTemplateRequest{
		RequestData: cloudIntegrationInstancesAPI.CreateCloudOnboardingIntegrationTemplateRequestData{
			AdditionalCapabilities:  additionalCapabilities,
			CloudProvider:           plan.CloudProvider.ValueString(),
			CollectionConfiguration: collectionConfiguration,
			CustomResourcesTags:     customResourcesTags,
			InstanceName:            plan.InstanceName.ValueString(),
			ScanMode:                plan.ScanMode.ValueString(),
			Scope:                   plan.Scope.ValueString(),
			ScopeModifications:      scopeModifications,
		},
	}

	response, err := cloudIntegrationInstancesAPI.Create(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error creating Cloud Onboarding Integration Template",
			err.Error(),
		)
		return response
	}

	return response
}

func getCloudIntegrations(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, request cloudIntegrationInstancesAPI.CloudIntegrationInstancesRequest) cloudIntegrationInstancesAPI.CloudIntegrationInstancesResponse {
	response, err := cloudIntegrationInstancesAPI.GetInstances(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error retrieving Cloud Integrations",
			err.Error(),
		)
		return response
	}

	return response
}

func getCloudIntegrationsByInstanceId(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, instanceId string) cloudIntegrationInstancesAPI.CloudIntegrationInstancesResponse {
	request := cloudIntegrationInstancesAPI.CloudIntegrationInstancesRequest{
		RequestData: cloudIntegrationInstancesAPI.CloudIntegrationInstancesRequestData{
			FilterData: cloudIntegrationInstancesAPI.CloudIntegrationInstancesFilterData{
				Filter: cloudIntegrationInstancesAPI.CloudIntegrationInstancesFilter{
					And: []cloudIntegrationInstancesAPI.CloudIntegrationInstancesAndFilter{
						{
							SearchField: "ID",
							SearchType:  "EQ",
							SearchValue: instanceId,
						},
					},
				},
				Paging: cloudIntegrationInstancesAPI.CloudIntegrationInstancesPaging{
					From: 0,
					To:   1000,
				},
			},
		},
	}

	response, err := cloudIntegrationInstancesAPI.GetInstances(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error retrieving Cloud Integrations by instance ID",
			err.Error(),
		)
		return response
	}

	return response
}

func getCloudIntegrationStatus(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, instanceId string) cloudIntegrationInstancesAPI.CloudIntegrationInstanceDetailsResponse {
	request := cloudIntegrationInstancesAPI.CloudIntegrationInstanceDetailsRequest{
		RequestData: cloudIntegrationInstancesAPI.CloudIntegrationInstanceDetailsRequestData{
			InstanceId: instanceId,
		},
	}

	response, err := cloudIntegrationInstancesAPI.GetInstanceDetails(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error reading Cloud Integration status",
			err.Error(),
		)
		return response
	}

	return response
}

func updateCloudIntegration(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, plan models.CloudIntegrationInstanceModel) cloudIntegrationInstancesAPI.CreateCloudOnboardingIntegrationTemplateResponse {
	var response cloudIntegrationInstancesAPI.CreateCloudOnboardingIntegrationTemplateResponse

	additionalCapabilities := cloudIntegrationInstancesAPI.CloudIntegrationAdditionalCapabilities{}
	diagnostics.Append(plan.AdditionalCapabilities.As(ctx, &additionalCapabilities, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	collectionConfiguration := cloudIntegrationInstancesAPI.CloudIntegrationCollectionConfiguration{}
	diagnostics.Append(plan.CollectionConfiguration.As(ctx, &collectionConfiguration, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	customResourcesTags := []cloudIntegrationInstancesAPI.CloudIntegrationCustomResourcesTag{}
	diagnostics.Append(plan.CustomResourceTags.ElementsAs(ctx, &customResourcesTags, false)...)
	if diagnostics.HasError() {
		return response
	}

	scopeModifications := cloudIntegrationInstancesAPI.CloudIntegrationScopeModifications{}
	diagnostics.Append(plan.ScopeModifications.As(ctx, &scopeModifications, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	request := cloudIntegrationInstancesAPI.CloudIntegrationEditRequest{
		RequestData: cloudIntegrationInstancesAPI.CloudIntegrationEditRequestData{
			AdditionalCapabilities:  additionalCapabilities,
			CloudProvider:           plan.CloudProvider.ValueString(),
			CollectionConfiguration: collectionConfiguration,
			CustomResourcesTags:     customResourcesTags,
			InstanceId:              plan.InstanceId.ValueString(),
			InstanceName:            plan.InstanceName.ValueString(),
			ScopeModifications:      scopeModifications,
		},
	}

	response, err := cloudIntegrationInstancesAPI.UpdateInstanceTemplate(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error updating Cloud Onboarding Integration Template",
			err.Error(),
		)
		return response
	}

	return response
}

func parseTemplateURL(diagnostics *diag.Diagnostics, templateAutomatedLink string) string {
	var templateBucketUrl string

	templateAutomatedUrl, err := url.Parse(templateAutomatedLink)
	if err != nil {
		diagnostics.AddError(
			"Error creating Cloud Onboarding Integration Template",
			err.Error(),
		)
		return templateBucketUrl
	}

	templateAutomatedUrlParameters, err := url.ParseQuery(templateAutomatedUrl.RawFragment)
	if err != nil {
		diagnostics.AddError(
			"Error creating Cloud Onboarding Integration Template",
			err.Error(),
		)
		return templateBucketUrl
	}

	// TODO: verify with regex
	templateBucketUrl = templateAutomatedUrlParameters.Get("/stacks/quickcreate?templateURL")

	return templateBucketUrl
}
