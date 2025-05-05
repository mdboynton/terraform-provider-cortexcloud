package cloud_accounts

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
)

//"fmt"
//"net/http"

//"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"

type CreateCloudOnboardingIntegrationTemplateRequest struct {
    RequestData RequestData `json:"request_data"`
}

type RequestData struct {
    AdditionalCapabilities AdditionalCapabilities `json:"additional_capabilities"`
    CloudProvider string `json:"cloud_provider"`
    CollectionConfiguration CollectionConfiguration `json:"collection_configuration"`
    //ConnectorName string `json:"connector_name"`
    CustomResourcesTags []CustomResourcesTag `json:"custom_resources_tags"`
    InstanceName string `json:"instance_name"`
    ScanMode string `json:"scan_mode"`
    Scope string `json:"scope"`
    ScopeModifications ScopeModifications `json:"scope_modifications"`
}

type AdditionalCapabilities struct {
    DataSecurityPostureManagement bool `json:"data_security_posture_management" tfsdk:"data_security_posture_management"`
    RegistryScanning bool `json:"registry_scanning" tfsdk:"registry_scanning"`
    RegistryScanningOptions RegistryScanningOptions `json:"registry_scanning_options" tfsdk:"registry_scanning_options"`
    ServerlessScanning bool `json:"serverless_scanning" tfsdk:"serverless_scanning"`
    XSIAMAnalytics bool `json:"xsiam_analytics" tfsdk:"xsiam_analytics"`
}

type RegistryScanningOptions struct {
    Type string `json:"type" tfsdk:"type"`
}

type CollectionConfiguration struct {
    AuditLogs AuditLogs `json:"audit_logs" tfsdk:"audit_logs"`
}

type AuditLogs struct {
    Enabled bool `json:"enabled" tfsdk:"enabled"`
}

type CustomResourcesTag struct {
    Key string `json:"key" tfsdk:"key"`
    Value string `json:"value" tfsdk:"value"`
}

type ScopeModifications struct {
    Regions ScopeModificationsRegions `json:"regions" tfsdk:"regions"`
}

type ScopeModificationsRegions struct {
    Enabled bool `json:"enabled" tfsdk:"enabled"`
}
    



type CreateCloudOnboardingIntegrationTemplateResponse struct {
    Reply Reply `json:"reply" tfsdk:"reply"`
}

type Reply struct {
    Automated Automated `json:"automated" tfsdk:"automated"`
    Manual Manual `json:"manual" tfsdk:"manual"`
}

type Automated struct {
    Link string `json:"link" tfsdk:"link"`
    TrackingGuid string `json:"tracking_guid" tfsdk:"tracking_guid"`
}

type Manual struct {
    //TF_ARM string `json:"TF/ARM" tfsdk:"tf_arm"`
    CF string `json:"CF" tfsdk:"tf_arm"`
}



func Create(ctx context.Context, client *api.CortexCloudAPIClient, req CreateCloudOnboardingIntegrationTemplateRequest) (CreateCloudOnboardingIntegrationTemplateResponse, error) {
    var response CreateCloudOnboardingIntegrationTemplateResponse
    if err := client.Request(ctx, "POST", api.CreateCloudOnboardingIntegrationTemplateEndpoint, nil, req, &response); err != nil {
        return response, fmt.Errorf("creating cloud onboarding integration template: %s", err.Error())
    }

    return response, nil
}
