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

func getCloudIntegrations(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, request cloudAccountsAPI.CloudIntegrationInstancesRequest) (cloudAccountsAPI.CloudIntegrationInstancesResponse) {
    response, err := cloudAccountsAPI.GetInstances(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error retrieving Cloud Integrations",
			err.Error(),
		)
		return response
	}

    return response
}

func getCloudIntegrationsByInstanceId(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, instanceId string) (cloudAccountsAPI.CloudIntegrationInstancesResponse) {
    request := cloudAccountsAPI.CloudIntegrationInstancesRequest{
        RequestData: cloudAccountsAPI.CloudIntegrationInstancesRequestData{
            FilterData: cloudAccountsAPI.CloudIntegrationInstancesFilterData{
                Filter: cloudAccountsAPI.CloudIntegrationInstancesFilter{
                    And: []cloudAccountsAPI.CloudIntegrationInstancesAndFilter{
                        {
                            SearchField: "ID",
                            SearchType: "EQ",
                            SearchValue: instanceId,
                        },
                    },
                },
                Paging: cloudAccountsAPI.CloudIntegrationInstancesPaging{
                    From: 0,
                    To: 1000,
                },
            },
        },
    }

    response, err := cloudAccountsAPI.GetInstances(ctx, client, request)
	if err != nil {
		diagnostics.AddError(
			"Error retrieving Cloud Integrations by instance ID",
			err.Error(),
		)
		return response
	}

    return response
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

//func populateComputedAttributeValues(diagnostics *diag.Diagnostics, model *models.CloudOnboardingIntegrationTemplateModel, response cloudAccountsAPI.CloudIntegrationInstancesResponse, instanceId, cloudFormationLink string) {
//    model.InstanceId = types.StringValue(instanceId)
//    model.CloudFormationLink= types.StringValue(cloudFormationLink)
//
//    if (len(response.Reply.Data) == 0 || len(response.Reply.Data) > 1) {
//        model.Status = types.StringNull()
//        model.InstanceName = types.StringNull()
//        model.AccountName = types.StringNull()
//        model.OutpostId = types.StringNull()
//        model.CreationTime = types.StringNull()
//    
//        diagnostics.AddWarning(
//            "Integration Status Unknown",
//            "Unable to retrieve computed values for the following arguments " +
//            "from the Cortex Cloud API: status, instance_name, account_name, " +
//            "outpost_id, creation_time\n\n" +
//            "The provider will attempt to populate these arguments during " +
//            "the next terraform apply operation.",
//        )
//
//        return
//    }
//
//    integrationData := response.Reply.Data[0]
//
//    if integrationData.Status != "" {
//        model.Status = types.StringValue(integrationData.Status)
//    } else {
//        model.Status = types.StringNull()
//    }
//
//    if integrationData.InstanceName != "" {
//        model.InstanceName = types.StringValue(integrationData.InstanceName)
//    } else {
//        model.InstanceName = types.StringNull()
//    }
//
//    if integrationData.AccountName != "" {
//        model.AccountName = types.StringValue(integrationData.AccountName)
//    } else {
//        model.AccountName = types.StringNull()
//    }
//
//    if integrationData.OutpostId != "" {
//        model.OutpostId = types.StringValue(integrationData.OutpostId)
//    } else {
//        model.OutpostId = types.StringNull()
//    }
//
//    if integrationData.CreationTime != 0 {
//        model.CreationTime = types.StringValue(time.Unix(int64(integrationData.CreationTime), 0).Format("2006-01-02T15:04:05.000Z"))
//    } else {
//        model.CreationTime = types.StringNull()
//    }
//}
