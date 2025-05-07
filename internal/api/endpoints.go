package api

import (
    "fmt"
)

const (
    BaseEndpoint = "public_api/v1"
)

var (
    HealthCheckEndpoint = fmt.Sprintf("%s/healthcheck", BaseEndpoint)

    CreateCloudOnboardingIntegrationTemplateEndpoint = fmt.Sprintf("%s/cloud_onboarding/create_instance_template", BaseEndpoint)
    GetCloudIntegrationInstancesEndpoint = fmt.Sprintf("%s/cloud_onboarding/get_instances", BaseEndpoint)
    GetCloudIntegrationInstanceDetailsEndpoint = fmt.Sprintf("%s/cloud_onboarding/get_instance_details", BaseEndpoint)
    EditCloudIntegrationInstanceTemplateEndpoint = fmt.Sprintf("%s/cloud_onboarding/edit_instance", BaseEndpoint)
)
