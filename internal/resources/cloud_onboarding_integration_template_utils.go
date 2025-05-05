package resources

import (
    "context"
    //"fmt"
    "net/url"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
	cloudAccountsAPI "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/cloud_accounts"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func createCloudOnboardingIntegrationTemplate(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, plan models.CloudOnboardingIntegrationTemplateModel) string {
    additionalCapabilities := cloudAccountsAPI.AdditionalCapabilities{}
    diagnostics.Append(plan.AdditionalCapabilities.As(ctx, &additionalCapabilities, basetypes.ObjectAsOptions{})...)
    if diagnostics.HasError() {
        return ""
    }

    collectionConfiguration := cloudAccountsAPI.CollectionConfiguration{}
    diagnostics.Append(plan.CollectionConfiguration.As(ctx, &collectionConfiguration, basetypes.ObjectAsOptions{})...)
    if diagnostics.HasError() {
        return ""
    }

    customResourcesTags := []cloudAccountsAPI.CustomResourcesTag{}
    diagnostics.Append(plan.CustomResourceTags.ElementsAs(ctx, &customResourcesTags, false)...)
    if diagnostics.HasError() {
        return ""
    }

    scopeModifications := cloudAccountsAPI.ScopeModifications{}
    diagnostics.Append(plan.ScopeModifications.As(ctx, &scopeModifications, basetypes.ObjectAsOptions{})...)
    if diagnostics.HasError() {
        return ""
    }

    request := cloudAccountsAPI.CreateCloudOnboardingIntegrationTemplateRequest{
        RequestData: cloudAccountsAPI.RequestData{
            AdditionalCapabilities: additionalCapabilities,
            CloudProvider: plan.CloudProvider.ValueString(),
            CollectionConfiguration: collectionConfiguration,
            CustomResourcesTags: customResourcesTags,
            InstanceName: plan.InstanceName.ValueString(),
            ScanMode: plan.ScanMode.ValueString(),
            Scope: plan.Scope.ValueString(),
            ScopeModifications: scopeModifications,
        },
    }

    response, err := cloudAccountsAPI.Create(ctx, client, request)
    if err != nil {
        diagnostics.AddError(
            "Error creating Cloud Onboarding Integration Template",
            err.Error(),
        )
        return ""
    }

    templateAutomatedUrl, err := url.Parse(response.Reply.Automated.Link)
    if err != nil {
        diagnostics.AddError(
            "Error creating Cloud Onboarding Integration Template",
            err.Error(),
        )
        return ""
    }

    templateAutomatedUrlParameters, err := url.ParseQuery(templateAutomatedUrl.RawFragment)
    if err != nil {
        diagnostics.AddError(
            "Error creating Cloud Onboarding Integration Template",
            err.Error(),
        )
        return ""
    }
    
    templateBucketUrl := templateAutomatedUrlParameters.Get("/stacks/quickcreate?templateURL")

    return templateBucketUrl
}
