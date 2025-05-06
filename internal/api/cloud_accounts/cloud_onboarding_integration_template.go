package cloud_accounts

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
)

// Create Cloud Onboarding Integration Template request structs
type CreateCloudOnboardingIntegrationTemplateRequest struct {
    RequestData CreateCloudOnboardingIntegrationTemplateRequestData `json:"request_data"`
}

type CreateCloudOnboardingIntegrationTemplateRequestData struct {
    AdditionalCapabilities CloudIntegrationAdditionalCapabilities `json:"additional_capabilities"`
    CloudProvider string `json:"cloud_provider"`
    CollectionConfiguration CloudIntegrationCollectionConfiguration `json:"collection_configuration"`
    //ConnectorName string `json:"connector_name"`
    CustomResourcesTags []CloudIntegrationCustomResourcesTag `json:"custom_resources_tags"`
    InstanceName string `json:"instance_name"`
    ScanMode string `json:"scan_mode"`
    Scope string `json:"scope"`
    ScopeModifications CloudIntegrationScopeModifications `json:"scope_modifications"`
}

type CloudIntegrationAdditionalCapabilities struct {
    DataSecurityPostureManagement bool `json:"data_security_posture_management" tfsdk:"data_security_posture_management"`
    RegistryScanning bool `json:"registry_scanning" tfsdk:"registry_scanning"`
    RegistryScanningOptions CloudIntegrationRegistryScanningOptions `json:"registry_scanning_options" tfsdk:"registry_scanning_options"`
    ServerlessScanning bool `json:"serverless_scanning" tfsdk:"serverless_scanning"`
    XSIAMAnalytics bool `json:"xsiam_analytics" tfsdk:"xsiam_analytics"`
}

type CloudIntegrationRegistryScanningOptions struct {
    Type string `json:"type" tfsdk:"type"`
}

type CloudIntegrationCollectionConfiguration struct {
    AuditLogs CloudIntegrationAuditLogs `json:"audit_logs" tfsdk:"audit_logs"`
}

type CloudIntegrationAuditLogs struct {
    Enabled bool `json:"enabled" tfsdk:"enabled"`
}

type CloudIntegrationCustomResourcesTag struct {
    Key string `json:"key" tfsdk:"key"`
    Value string `json:"value" tfsdk:"value"`
}

type CloudIntegrationScopeModifications struct {
    Regions CloudIntegrationScopeModificationsRegions `json:"regions" tfsdk:"regions"`
}

type CloudIntegrationScopeModificationsRegions struct {
    Enabled bool `json:"enabled" tfsdk:"enabled"`
}
    
// Create Cloud Onboarding Integration Template response structs
type CreateCloudOnboardingIntegrationTemplateResponse struct {
    Reply CreateCloudOnboardingIntegrationTemplateReply `json:"reply" tfsdk:"reply"`
}

type CreateCloudOnboardingIntegrationTemplateReply struct {
    Automated CreateCloudOnboardingIntegrationTemplateAutomated `json:"automated" tfsdk:"automated"`
    Manual CreateCloudOnboardingIntegrationTemplateManual `json:"manual" tfsdk:"manual"`
}

type CreateCloudOnboardingIntegrationTemplateAutomated struct {
    Link string `json:"link" tfsdk:"link"`
    TrackingGuid string `json:"tracking_guid" tfsdk:"tracking_guid"`
}

type CreateCloudOnboardingIntegrationTemplateManual struct {
    //TF_ARM string `json:"TF/ARM" tfsdk:"tf_arm"`
    CF string `json:"CF" tfsdk:"tf_arm"`
}

// Get Instance Details request structs
type CloudIntegrationInstanceDetailsRequest struct {
    RequestData CloudIntegrationInstanceDetailsRequestData `json:"request_data" tfsdk:"request_data"`
}

type CloudIntegrationInstanceDetailsRequestData struct {
    InstanceId string `json:"instance_id" tfsdk:"instance_id"`
}

type CloudIntegrationInstanceDetailsResponse struct {
    Reply CloudIntegrationInstanceDetailsResponseData `json:"reply" tfsdk:"reply"`
}

type CloudIntegrationInstanceDetailsResponseData struct {
    Id string `json:"id" tfsdk:"id"`
    Collector string `json:"collector" tfsdk:"collector"`
    InstanceName string `json:"instance_name" tfsdk:"instance_name"`
    Scope string `json:"scope" tfsdk:"scope"`
    Tags []CloudIntegrationCustomResourcesTag `json:"tags" tfsdk:"tags"`
    Scan CloudIntegrationInstanceDetailsScan `json:"scan" tfsdk:"scan"`
    SecurityCapabilities []CloudIntegrationInstanceDetailsSecurityCapability `json:"security_capabilities" tfsdk:"security_capabilities"`
    Status string `json:"status" tfsdk:"status"`
    CloudProvider string `json:"cloud_provider" tfsdk:"cloud_provider"`
    CollectionConfiguration string `json:"collection_configuration" tfsdk:"collection_configuration"`
    AdditionalCapabilities string `json:"additional_capabilities" tfsdk:"additional_capabilities"`
}

type CloudIntegrationInstanceDetailsScan struct {
    ScanMethod string `json:"scan_method" tfsdk:"scan_method"`
}

type CloudIntegrationInstanceDetailsSecurityCapability struct {
    Name string `json:"name" tfsdk:"name"`
    Description string `json:"description" tfsdk:"description"`
    Status string `json:"status" tfsdk:"status"`
}

// Edit Integration Instance Template request structs
type CloudIntegrationEditRequest struct {
    RequestData CloudIntegrationEditRequestData `json:"request_data" tfsdk:"request_data"`
}

type CloudIntegrationEditRequestData struct {
    AdditionalCapabilities CloudIntegrationAdditionalCapabilities `json:"additional_capabilities"`
    CloudProvider string `json:"cloud_provider"`
    CollectionConfiguration CloudIntegrationCollectionConfiguration `json:"collection_configuration"`
    CustomResourcesTags []CloudIntegrationCustomResourcesTag `json:"custom_resources_tags"`
    InstanceId string `json:"instance_id" tfsdk:"instance_id"`
    //ScanEnvId string `json:"instance_id" tfsdk:"instance_id"`
    InstanceName string `json:"instance_name" tfsdk:"instance_name"`
    ScopeModifications CloudIntegrationScopeModifications `json:"scope_modifications"`
}

// Functions
func Create(ctx context.Context, client *api.CortexCloudAPIClient, req CreateCloudOnboardingIntegrationTemplateRequest) (CreateCloudOnboardingIntegrationTemplateResponse, error) {
    var response CreateCloudOnboardingIntegrationTemplateResponse
    if err := client.Request(ctx, "POST", api.CreateCloudOnboardingIntegrationTemplateEndpoint, nil, req, &response); err != nil {
        return response, fmt.Errorf("creating cloud onboarding integration template: %s", err.Error())
    }

    return response, nil
}

func GetInstanceDetails(ctx context.Context, client *api.CortexCloudAPIClient, req CloudIntegrationInstanceDetailsRequest) (CloudIntegrationInstanceDetailsResponse, error) {
    var response CloudIntegrationInstanceDetailsResponse
    if err := client.Request(ctx, "POST", api.GetCloudIntegrationInstanceDetailsEndpoint, nil, req, &response); err != nil {
        return response, fmt.Errorf("getting cloud integration instance details: %s", err.Error())
    }

    return response, nil
}

func UpdateInstanceTemplate(ctx context.Context, client *api.CortexCloudAPIClient, req CloudIntegrationEditRequest) (CreateCloudOnboardingIntegrationTemplateResponse, error) {
    var response CreateCloudOnboardingIntegrationTemplateResponse
    if err := client.Request(ctx, "POST", api.EditCloudIntegrationInstanceTemplateEndpoint, nil, req, &response); err != nil {
        return response, fmt.Errorf("updating cloud onboarding integration template: %s", err.Error())
    }

    return response, nil
}
