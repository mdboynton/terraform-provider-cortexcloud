package application_security

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type ApplicationSecurityRule struct {
	Category        string                             `json:"category"`
	CloudProvider   string                             `json:"cloudProvider"`
	CreatedAt       ApplicationSecurityRuleCreatedAt   `json:"createdAt"`
	Description     string                             `json:"description"`
	DetectionMethod string                             `json:"detectionMethod"`
	DocLink         string                             `json:"docLink"`
	Domain          string                             `json:"domain"`
	FindingCategory string                             `json:"findingCategory"`
	FindingDocs     string                             `json:"findingDocs"`
	FindingTypeId   int                                `json:"findingTypeId"`
	FindingTypeName string                             `json:"findingTypeName"`
	Frameworks      []ApplicationSecurityRuleFramework `json:"frameworks" tfsdk:"frameworks"`
	Id              string                             `json:"id"`
	IsCustom        bool                               `json:"isCustom"`
	IsEnabled       bool                               `json:"isEnabled"`
	Labels          []string                           `json:"labels"`
	MitreTactics    []string                           `json:"mitreTactics"`
	MitreTechniques []string                           `json:"mitreTechniques"`
	Name            string                             `json:"name"`
	Owner           string                             `json:"owner"`
	Scanner         string                             `json:"scanner"`
	//ScannerRuleId string `json:"scannerRuleId"`
	Severity string `json:"severity"`
	Source   string `json:"source"`
	//SourceVersion string `json:"sourceVersion"`
	SubCategory string                           `json:"subCategory"`
	UpdatedAt   ApplicationSecurityRuleUpdatedAt `json:"updatedAt"`
}

type ApplicationSecurityRuleFramework struct {
	Name                   string `json:"name" tfsdk:"name"`
	Definition             string `json:"definition" tfsdk:"definition"`
	DefinitionLink         string `json:"definitionLink" tfsdk:"definition_link"`
	RemediationDescription string `json:"remediationDescription" tfsdk:"remediation_description"`
	//RemediationIds []string `json:"remediationIds" tfsdk:"remediation_ids"`
	//ResourceTypes []string `json:"resourceTypes" tfsdk:"remediation_types"`
}

type ApplicationSecurityRuleCreatedAt struct {
	Value string `json:"value"`
}

type ApplicationSecurityRuleUpdatedAt struct {
	Value string `json:"value"`
}

type CreateApplicationSecurityRuleRequest struct {
	Name        string                             `json:"name"`
	Severity    string                             `json:"severity"`
	Scanner     string                             `json:"scanner"`
	Frameworks  []ApplicationSecurityRuleFramework `json:"frameworks"`
	Category    string                             `json:"category"`
	SubCategory string                             `json:"subCategory"`
	Description string                             `json:"description"`
	Labels      []string                           `json:"labels"`
}

// Functions

func Create(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, request CreateApplicationSecurityRuleRequest) ApplicationSecurityRule {
	var response ApplicationSecurityRule

	if err := client.Request(ctx, "POST", api.ApplicationSecurityRulesEndpoint, nil, request, &response); err != nil {
		diagnostics.AddError(
			"Error creating Application Security Rule",
			err.Error(),
		)
	}

	return response
}

func Get(ctx context.Context, diagnostics *diag.Diagnostics, client *api.CortexCloudAPIClient, ruleId string) ApplicationSecurityRule {
	var response ApplicationSecurityRule

	if err := client.Request(ctx, "GET", fmt.Sprintf("%s/%s", api.ApplicationSecurityRulesEndpoint, ruleId), nil, nil, &response); err != nil {
		diagnostics.AddError(
			"Error retrieving Application Security Rule",
			err.Error(),
		)
	}

	return response
}
