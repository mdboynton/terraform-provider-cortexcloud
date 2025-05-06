package resources

import (
	"context"
	//"fmt"
	"net/url"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
	cloudAccountsAPI "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/cloud_accounts"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func createCloudOnboardingIntegrationTemplate(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, plan models.CloudOnboardingIntegrationTemplateModel) cloudAccountsAPI.CreateCloudOnboardingIntegrationTemplateResponse {
    var response cloudAccountsAPI.CreateCloudOnboardingIntegrationTemplateResponse

	additionalCapabilities := cloudAccountsAPI.CloudIntegrationAdditionalCapabilities{}
	diagnostics.Append(plan.AdditionalCapabilities.As(ctx, &additionalCapabilities, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	collectionConfiguration := cloudAccountsAPI.CloudIntegrationCollectionConfiguration{}
	diagnostics.Append(plan.CollectionConfiguration.As(ctx, &collectionConfiguration, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	customResourcesTags := []cloudAccountsAPI.CloudIntegrationCustomResourcesTag{}
	diagnostics.Append(plan.CustomResourceTags.ElementsAs(ctx, &customResourcesTags, false)...)
	if diagnostics.HasError() {
		return response
	}

	scopeModifications := cloudAccountsAPI.CloudIntegrationScopeModifications{}
	diagnostics.Append(plan.ScopeModifications.As(ctx, &scopeModifications, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

    request := cloudAccountsAPI.CreateCloudOnboardingIntegrationTemplateRequest{
		RequestData: cloudAccountsAPI.CreateCloudOnboardingIntegrationTemplateRequestData{
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

    response, err := cloudAccountsAPI.Create(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error creating Cloud Onboarding Integration Template",
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

func getCloudIntegrationStatus(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, instanceId string) (cloudAccountsAPI.CloudIntegrationInstanceDetailsResponse) {
    request := cloudAccountsAPI.CloudIntegrationInstanceDetailsRequest{
        RequestData: cloudAccountsAPI.CloudIntegrationInstanceDetailsRequestData{
            InstanceId: instanceId,
        },
    }

    response, err := cloudAccountsAPI.GetInstanceDetails(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error reading Cloud Integration status",
			err.Error(),
		)
		return response
	}

    return response
}

func updateCloudIntegration(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, plan models.CloudOnboardingIntegrationTemplateModel) cloudAccountsAPI.CreateCloudOnboardingIntegrationTemplateResponse {
    var response cloudAccountsAPI.CreateCloudOnboardingIntegrationTemplateResponse

	additionalCapabilities := cloudAccountsAPI.CloudIntegrationAdditionalCapabilities{}
	diagnostics.Append(plan.AdditionalCapabilities.As(ctx, &additionalCapabilities, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	collectionConfiguration := cloudAccountsAPI.CloudIntegrationCollectionConfiguration{}
	diagnostics.Append(plan.CollectionConfiguration.As(ctx, &collectionConfiguration, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

	customResourcesTags := []cloudAccountsAPI.CloudIntegrationCustomResourcesTag{}
	diagnostics.Append(plan.CustomResourceTags.ElementsAs(ctx, &customResourcesTags, false)...)
	if diagnostics.HasError() {
		return response
	}

	scopeModifications := cloudAccountsAPI.CloudIntegrationScopeModifications{}
	diagnostics.Append(plan.ScopeModifications.As(ctx, &scopeModifications, basetypes.ObjectAsOptions{})...)
	if diagnostics.HasError() {
		return response
	}

    request := cloudAccountsAPI.CloudIntegrationEditRequest{
        RequestData: cloudAccountsAPI.CloudIntegrationEditRequestData{
            AdditionalCapabilities: additionalCapabilities,
			CloudProvider:           plan.CloudProvider.ValueString(),
			CollectionConfiguration: collectionConfiguration,
			CustomResourcesTags:     customResourcesTags,
            InstanceId:              plan.InstanceId.ValueString(),
			InstanceName:            plan.InstanceName.ValueString(),
			ScopeModifications:      scopeModifications,
        },
    }

    response, err := cloudAccountsAPI.UpdateInstanceTemplate(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error updating Cloud Onboarding Integration Template",
			err.Error(),
		)
		return response
	}

    return response
}
