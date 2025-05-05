package api

import (
    "fmt"
)

const (
    BaseEndpoint = "public_api/v1/"
)

var (
    CreateCloudOnboardingIntegrationTemplateEndpoint = fmt.Sprintf("%s/cloud_onboarding/create_instance_template", BaseEndpoint)
)
