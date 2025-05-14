// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"fmt"
)

const (
	BaseEndpoint = "public_api"
)

var (
	// Health Check
	HealthCheckEndpoint = fmt.Sprintf("%s/v1/healthcheck", BaseEndpoint)

	// Cloud Integration Instances
	CreateCloudOnboardingIntegrationTemplateEndpoint = fmt.Sprintf("%s/v1/cloud_onboarding/create_instance_template", BaseEndpoint)
	GetCloudIntegrationInstancesEndpoint             = fmt.Sprintf("%s/v1/cloud_onboarding/get_instances", BaseEndpoint)
	GetCloudIntegrationInstanceDetailsEndpoint       = fmt.Sprintf("%s/v1/cloud_onboarding/get_instance_details", BaseEndpoint)
	EditCloudIntegrationInstanceTemplateEndpoint     = fmt.Sprintf("%s/v1/cloud_onboarding/edit_instance", BaseEndpoint)
	DeleteCloudIntegrationInstanceEndpoint           = fmt.Sprintf("%s/v1/cloud_onboarding/delete_instance", BaseEndpoint)

	// Application Security Rules
	ApplicationSecurityRulesEndpoint        = fmt.Sprintf("%s/appsec/v1/rules", BaseEndpoint)
	ValidateApplicationSecurityRuleEndpoint = fmt.Sprintf("%s/appsec/v1/rules/rule_actions", BaseEndpoint)
)
