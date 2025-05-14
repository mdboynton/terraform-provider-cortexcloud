package cloud_integration

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// *********************************************************
// CREATE request structs
// *********************************************************
type CreateCloudOnboardingIntegrationTemplateRequest struct {
	RequestData CreateCloudOnboardingIntegrationTemplateRequestData `json:"request_data"`
}

type CreateCloudOnboardingIntegrationTemplateRequestData struct {
	AccountDetails          *CloudIntegrationAccountDetails         `json:"account_details,omitempty"`
	AdditionalCapabilities  CloudIntegrationAdditionalCapabilities  `json:"additional_capabilities"`
	CloudProvider           string                                  `json:"cloud_provider"`
	CollectionConfiguration CloudIntegrationCollectionConfiguration `json:"collection_configuration"`
	//ConnectorName string `json:"connector_name"`
	CustomResourcesTags []CloudIntegrationCustomResourcesTag `json:"custom_resources_tags"`
	InstanceName        string                               `json:"instance_name"`
	ScanMode            string                               `json:"scan_mode"`
	Scope               string                               `json:"scope"`
	ScopeModifications  CloudIntegrationScopeModifications   `json:"scope_modifications"`
}

type CloudIntegrationAccountDetails struct {
	OrganizationId string `json:"organization_id" tfsdk:"organization_id"`
}

type CloudIntegrationAdditionalCapabilities struct {
	DataSecurityPostureManagement bool                                    `json:"data_security_posture_management" tfsdk:"data_security_posture_management"`
	RegistryScanning              bool                                    `json:"registry_scanning" tfsdk:"registry_scanning"`
	RegistryScanningOptions       CloudIntegrationRegistryScanningOptions `json:"registry_scanning_options" tfsdk:"registry_scanning_options"`
	ServerlessScanning            bool                                    `json:"serverless_scanning" tfsdk:"serverless_scanning"`
	XSIAMAnalytics                bool                                    `json:"xsiam_analytics" tfsdk:"xsiam_analytics"`
}

type CloudIntegrationRegistryScanningOptions struct {
	Type     string `json:"type" tfsdk:"type"`
	LastDays *int   `json:"last_days,omitempty" tfsdk:"last_days"`
}

type CloudIntegrationCollectionConfiguration struct {
	AuditLogs CloudIntegrationAuditLogs `json:"audit_logs" tfsdk:"audit_logs"`
}

type CloudIntegrationAuditLogs struct {
	Enabled bool `json:"enabled" tfsdk:"enabled"`
}

type CloudIntegrationCustomResourcesTag struct {
	Key   string `json:"key" tfsdk:"key"`
	Value string `json:"value" tfsdk:"value"`
}

type CloudIntegrationScopeModifications struct {
	Accounts *CloudIntegrationScopeModificationsAccounts `json:"accounts" tfsdk:"accounts"`
	Regions  *CloudIntegrationScopeModificationsRegions  `json:"regions" tfsdk:"regions"`
}

type CloudIntegrationScopeModificationsAccounts struct {
	Enabled bool `json:"enabled" tfsdk:"enabled"`
}

type CloudIntegrationScopeModificationsRegions struct {
	Enabled bool `json:"enabled" tfsdk:"enabled"`
}

// *********************************************************
// CREATE response structs
// *********************************************************

type CreateCloudOnboardingIntegrationTemplateResponse struct {
	Reply CreateCloudOnboardingIntegrationTemplateReply `json:"reply" tfsdk:"reply"`
}

type CreateCloudOnboardingIntegrationTemplateReply struct {
	Automated CreateCloudOnboardingIntegrationTemplateAutomated `json:"automated" tfsdk:"automated"`
	Manual    CreateCloudOnboardingIntegrationTemplateManual    `json:"manual" tfsdk:"manual"`
}

type CreateCloudOnboardingIntegrationTemplateAutomated struct {
	Link         string `json:"link" tfsdk:"link"`
	TrackingGuid string `json:"tracking_guid" tfsdk:"tracking_guid"`
}

type CreateCloudOnboardingIntegrationTemplateManual struct {
	//TF_ARM string `json:"TF/ARM" tfsdk:"tf_arm"`
	CF string `json:"CF" tfsdk:"tf_arm"`
}

// *********************************************************
// GET INSTANCES request structs
// *********************************************************

type CloudIntegrationInstancesRequest struct {
	RequestData CloudIntegrationInstancesRequestData `json:"request_data" tfsdk:"request_data"`
}

type CloudIntegrationInstancesRequestData struct {
	FilterData CloudIntegrationInstancesFilterData `json:"filter_data" tfsdk:"filter_data"`
}

type CloudIntegrationInstancesFilterData struct {
	Sort   []CloudIntegrationInstancesSort `json:"sort,omitempty" tfsdk:"sort"`
	Paging CloudIntegrationInstancesPaging `json:"paging,omitempty" tfsdk:"paging"`
	Filter CloudIntegrationInstancesFilter `json:"filter" tfsdk:"filter"`
}

type CloudIntegrationInstancesSort struct {
	Field string `json:"FIELD" tfsdk:"field"`
	Order string `json:"ORDER" tfsdk:"order"`
}

type CloudIntegrationInstancesPaging struct {
	From int `json:"from" tfsdk:"from"` // TODO: cant be less than 0
	To   int `json:"to" tfsdk:"to"`     // TODO: cant be more than 1000
}

type CloudIntegrationInstancesFilter struct {
	And []CloudIntegrationInstancesAndFilter `json:"AND" tfsdk:"and"`
}

type CloudIntegrationInstancesAndFilter struct {
	SearchField string `json:"SEARCH_FIELD" tfsdk:"search_field"`
	SearchType  string `json:"SEARCH_TYPE" tfsdk:"search_type"`
	SearchValue string `json:"SEARCH_VALUE" tfsdk:"search_value"`
}

// *********************************************************
// GET INSTANCES response structs
// *********************************************************

type CloudIntegrationInstancesResponse struct {
	Reply CloudIntegrationInstancesResponseReply `json:"reply" tfsdk:"reply"`
}

type CloudIntegrationInstancesResponseReply struct {
	Data        []CloudIntegrationInstancesResponseData `json:"DATA" tfsdk:"data"`
	FilterCount int                                     `json:"FILTER_COUNT" tfsdk:"FILTER_COUNT"`
	TotalCloud  int                                     `json:"TOTAL_COUNT" tfsdk:"TOTAL_COUNT"`
}

type CloudIntegrationInstancesResponseData struct {
	InstanceId    string `json:"instance_id" tfsdk:"instance_id"`
	CloudProvider string `json:"cloud_provider" tfsdk:"cloud_provider"`
	InstanceName  string `json:"instance_name" tfsdk:"instance_name"`
	AccountName   string `json:"account_name" tfsdk:"account_name"`
	//Accounts
	Scope              string `json:"scope" tfsdk:"scope"`
	ScanMode           string `json:"scan_mode" tfsdk:"scan_mode"`
	CreationTime       int    `json:"creation_time" tfsdk:"creation_time"`
	CustomResourceTags string `json:"custom_resource_tags" tfsdk:"custom_resource_tags"`
	//ProvisioningMethod
	AdditionalCapabilities string `json:"additional_capabilities" tfsdk:"additional_capabilities"`
	IsPendingChanges       string `json:"is_pending_changes" tfsdk:"is_pending_changes"` // TODO: might be a bool?
	Status                 string `json:"status" tfsdk:"status"`
	OutpostId              string `json:"outpost_id" tfsdk:"outpost_id"`
	//CollectionConfiguration string `json:"collection_configuration" tfsdk:"collection_configuration"`
}

// *********************************************************
// GET INSTANCE DETAILS request structs
// *********************************************************

type CloudIntegrationInstanceDetailsRequest struct {
	RequestData CloudIntegrationInstanceDetailsRequestData `json:"request_data" tfsdk:"request_data"`
}

type CloudIntegrationInstanceDetailsRequestData struct {
	InstanceId string `json:"id" tfsdk:"instance_id"`
}

// *********************************************************
// GET INSTANCE DETAILS response structs
// *********************************************************

type CloudIntegrationInstanceDetailsResponse struct {
	Reply CloudIntegrationInstanceDetailsResponseData `json:"reply" tfsdk:"reply"`
}

type CloudIntegrationInstanceDetailsResponseData struct {
	Id                      string                                              `json:"id" tfsdk:"id"`
	Collector               string                                              `json:"collector" tfsdk:"collector"`
	InstanceName            string                                              `json:"instance_name" tfsdk:"instance_name"`
	Scope                   string                                              `json:"scope" tfsdk:"scope"`
	Tags                    []CloudIntegrationCustomResourcesTag                `json:"tags" tfsdk:"tags"`
	Scan                    CloudIntegrationInstanceDetailsScan                 `json:"scan" tfsdk:"scan"`
	SecurityCapabilities    []CloudIntegrationInstanceDetailsSecurityCapability `json:"security_capabilities" tfsdk:"security_capabilities"`
	Status                  string                                              `json:"status" tfsdk:"status"`
	CloudProvider           string                                              `json:"cloud_provider" tfsdk:"cloud_provider"`
	CollectionConfiguration string                                              `json:"collection_configuration" tfsdk:"collection_configuration"`
	AdditionalCapabilities  string                                              `json:"additional_capabilities" tfsdk:"additional_capabilities"`
}

type CloudIntegrationInstanceDetailsScan struct {
	ScanMethod string `json:"scan_method" tfsdk:"scan_method"`
}

type CloudIntegrationInstanceDetailsSecurityCapability struct {
	Name        string `json:"name" tfsdk:"name"`
	Description string `json:"description" tfsdk:"description"`
	Status      string `json:"status" tfsdk:"status"`
}

// *********************************************************
// EDIT request structs
// *********************************************************

type CloudIntegrationEditRequest struct {
	RequestData CloudIntegrationEditRequestData `json:"request_data" tfsdk:"request_data"`
}

type CloudIntegrationEditRequestData struct {
	AdditionalCapabilities  CloudIntegrationAdditionalCapabilities  `json:"additional_capabilities"`
	CloudProvider           string                                  `json:"cloud_provider"`
	CollectionConfiguration CloudIntegrationCollectionConfiguration `json:"collection_configuration"`
	CustomResourcesTags     []CloudIntegrationCustomResourcesTag    `json:"custom_resources_tags"`
	//InstanceId              string                                  `json:"instance_id" tfsdk:"instance_id"`
	InstanceId         string                             `json:"id" tfsdk:"instance_id"`
	ScanEnvId          string                             `json:"scan_env_id" tfsdk:"scan_env_id"`
	InstanceName       string                             `json:"instance_name" tfsdk:"instance_name"`
	ScopeModifications CloudIntegrationScopeModifications `json:"scope_modifications"`
}

// *********************************************************
// DELETE request structs
// *********************************************************

type CloudIntegrationDeleteRequest struct {
	RequestData CloudIntegrationDeleteRequestData `json:"request_data" tfsdk:"request_data"`
}

type CloudIntegrationDeleteRequestData struct {
	Ids []string `json:"ids" tfsdk:"ids"`
}

// *********************************************************
// DELETE response structs
// *********************************************************

type CloudIntegrationDeleteResponse struct {
	Reply CloudIntegrationDeleteResponseReply `json:"reply" tfsdk:"reply"`
}

type CloudIntegrationDeleteResponseReply struct {
}

// *********************************************************
// Request functions
// *********************************************************

func CreateTemplate(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, request CreateCloudOnboardingIntegrationTemplateRequest) (CreateCloudOnboardingIntegrationTemplateResponse, string) {
	var response CreateCloudOnboardingIntegrationTemplateResponse

	if err := client.Request(ctx, http.MethodPost, api.CreateCloudOnboardingIntegrationTemplateEndpoint, nil, request, &response); err != nil {
		diagnostics.AddError(
			"Error creating Cloud Onboarding Integration Template",
			err.Error(),
		)
	}

	templateUrl, err := getCloudFormationTemplateUrl(response.Reply.Automated.Link)
	if err != nil {
		diagnostics.AddError(
			"Error creating Cloud Onboarding Integration Template",
			fmt.Sprintf("Failed to parse CloudFormation template URL from API response: %s", err.Error()),
		)
	}

	return response, templateUrl
}

func Get(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, request CloudIntegrationInstancesRequest) CloudIntegrationInstancesResponse {
	var response CloudIntegrationInstancesResponse

	if err := client.Request(ctx, http.MethodPost, api.GetCloudIntegrationInstancesEndpoint, nil, request, &response); err != nil {
		diagnostics.AddError(
			"Error retrieving Cloud Integrations",
			err.Error(),
		)
	}

	return response
}

func GetByInstanceId(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, instanceId string) CloudIntegrationInstancesResponse {
	var response CloudIntegrationInstancesResponse

	request := createGetByInstanceIdRequest(instanceId)

	if err := client.Request(ctx, http.MethodPost, api.GetCloudIntegrationInstancesEndpoint, nil, request, &response); err != nil {
		diagnostics.AddError(
			"Error retrieving Cloud Integrations",
			err.Error(),
		)
	}

	return response
}

func GetCloudFormationDeploymentByInstanceName(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, instanceName string) CloudIntegrationInstancesResponse {
	var response CloudIntegrationInstancesResponse

	request := createGetCloudFormationDeploymentByInstanceIdRequest(instanceName)

	if err := client.Request(ctx, http.MethodPost, api.GetCloudIntegrationInstancesEndpoint, nil, request, &response); err != nil {
		diagnostics.AddError(
			"Error retrieving Cloud Integrations",
			err.Error(),
		)
	}

	return response
}

func GetDetails(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, request CloudIntegrationInstanceDetailsRequest) CloudIntegrationInstanceDetailsResponse {
	var response CloudIntegrationInstanceDetailsResponse

	if err := client.Request(ctx, http.MethodPost, api.GetCloudIntegrationInstanceDetailsEndpoint, nil, request, &response); err != nil {
		diagnostics.AddError(
			"Error retrieving Cloud Integrations by instance ID",
			err.Error(),
		)
	}

	return response
}

func Update(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, request CloudIntegrationEditRequest) CreateCloudOnboardingIntegrationTemplateResponse {
	var response CreateCloudOnboardingIntegrationTemplateResponse

	if err := client.Request(ctx, http.MethodPost, api.EditCloudIntegrationInstanceTemplateEndpoint, nil, request, &response); err != nil {
		diagnostics.AddError(
			"Error updating Cloud Onboarding Integration Template",
			err.Error(),
		)
	}

	return response
}

func Delete(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, request CloudIntegrationDeleteRequest) CloudIntegrationDeleteResponse {
	var response CloudIntegrationDeleteResponse

	if err := client.Request(ctx, http.MethodPost, api.DeleteCloudIntegrationInstanceEndpoint, nil, request, &response); err != nil {
		diagnostics.AddError(
			"Error deleting Cloud Integration",
			err.Error(),
		)
	}

	return response
}

// *********************************************************
// Helper functions
// *********************************************************

func createGetByInstanceIdRequest(instanceId string) CloudIntegrationInstancesRequest {
	return CloudIntegrationInstancesRequest{
		RequestData: CloudIntegrationInstancesRequestData{
			FilterData: CloudIntegrationInstancesFilterData{
				Filter: CloudIntegrationInstancesFilter{
					And: []CloudIntegrationInstancesAndFilter{
						{
							SearchField: "ID",
							SearchType:  "EQ",
							SearchValue: instanceId,
						},
					},
				},
				Paging: CloudIntegrationInstancesPaging{
					From: 0,
					To:   1000,
				},
			},
		},
	}
}

func createGetCloudFormationDeploymentByInstanceIdRequest(instanceId string) CloudIntegrationInstancesRequest {
	return CloudIntegrationInstancesRequest{
		RequestData: CloudIntegrationInstancesRequestData{
			FilterData: CloudIntegrationInstancesFilterData{
				Filter: CloudIntegrationInstancesFilter{
					And: []CloudIntegrationInstancesAndFilter{
						{
							SearchField: "ID",
							SearchType:  "EQ",
							SearchValue: instanceId,
						},
						{
							SearchField: "PROVISIONING_METHOD",
							SearchType:  "EQ",
							SearchValue: "CF",
						},
					},
				},
				Paging: CloudIntegrationInstancesPaging{
					From: 0,
					To:   1000,
				},
			},
		},
	}
}

func getCloudFormationTemplateUrl(responseUrl string) (string, error) {
	var cloudFormationTemplateUrl string

	parsedResponseUrl, err := url.Parse(responseUrl)
	if err != nil {
		return cloudFormationTemplateUrl, err
	}

	queryValueMap, err := url.ParseQuery(parsedResponseUrl.RawFragment)
	if err != nil {
		return cloudFormationTemplateUrl, err
	}

	// TODO: verify with regex
	cloudFormationTemplateUrl = queryValueMap.Get("/stacks/quickcreate?templateURL")

	return cloudFormationTemplateUrl, nil
}
