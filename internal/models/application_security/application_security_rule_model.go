// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	//"slices"
	//"strings"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/util"
	"github.com/mdboynton/cortex-cloud-go/appsec"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// *********************************************************
// Structs
// *********************************************************
type ApplicationSecurityRuleModel struct {
	Category        types.String     `tfsdk:"category"`
	CloudProvider   types.String     `tfsdk:"cloud_provider"`
	CreatedAt       types.String     `tfsdk:"created_at"`
	Description     types.String     `tfsdk:"description"`
	DetectionMethod types.String     `tfsdk:"detection_method"`
	DocLink         types.String     `tfsdk:"doc_link"`
	Domain          types.String     `tfsdk:"domain"`
	FindingCategory types.String     `tfsdk:"finding_category"`
	FindingDocs     types.String     `tfsdk:"finding_docs"`
	FindingTypeId   types.Int32      `tfsdk:"finding_type_id"`
	FindingTypeName types.String     `tfsdk:"finding_type_name"`
	Frameworks      []FrameworkModel `tfsdk:"frameworks"`
	Id              types.String     `tfsdk:"id"`
	IsCustom        types.Bool       `tfsdk:"is_custom"`
	IsEnabled       types.Bool       `tfsdk:"is_enabled"`
	Labels          types.Set        `tfsdk:"labels"`
	MitreTactics    types.Set        `tfsdk:"mitre_tactics"`
	MitreTechniques types.Set        `tfsdk:"mitre_techniques"`
	Name            types.String     `tfsdk:"name"`
	Owner           types.String     `tfsdk:"owner"`
	Scanner         types.String     `tfsdk:"scanner"`
	Severity        types.String     `tfsdk:"severity"`
	Source          types.String     `tfsdk:"source"`
	SubCategory     types.String     `tfsdk:"sub_category"`
	UpdatedAt       types.String     `tfsdk:"updated_at"`
}

type FrameworkModel struct {
	Name types.String `tfsdk:"name"`
	//Definition             FrameworkDefinitionModel
	Definition             types.String `tfsdk:"definition"`
	DefinitionLink         types.String `tfsdk:"definition_link"`
	RemediationDescription types.String `tfsdk:"remediation_description"`
}

type FrameworkDefinitionModel struct {
	Metadata *FrameworkDefinitionMetadataModel `tfsdk:"metadata" yaml:"metadata,omitempty"`
	Scope    FrameworkDefinitionScopeModel     `tfsdk:"scope" yaml:"scope,omitempty"`
	//Definition FrameworkDefinitionLogicModel     `tfsdk:"definition" yaml:"definition,omitempty"`
	Definition types.Dynamic `tfsdk:"definition" yaml:"definition,omitempty"`
}

type FrameworkDefinitionMetadataModel struct {
	Name       types.String `tfsdk:"name" yaml:"name"`
	Guidelines types.String `tfsdk:"guidelines" yaml:"guidelines"`
	Category   types.String `tfsdk:"category" yaml:"category"`
	Severity   types.String `tfsdk:"severity" yaml:"severity"`
}

type FrameworkDefinitionScopeModel struct {
	Provider types.String `tfsdk:"provider" yaml:"provider,omitempty"`
}

type FrameworkDefinitionLogicModel struct {
	And []FrameworkDefinitionLogicConditionModel `tfsdk:"and" yaml:"and,omitempty"`
	Or  []FrameworkDefinitionLogicConditionModel `tfsdk:"or" yaml:"or,omitempty"`
}

type FrameworkDefinitionLogicConditionModel struct {
	ConditionType types.String                             `tfsdk:"condition_type" yaml:"cond_type,omitempty"`
	ResourceTypes types.List                               `tfsdk:"resource_types" yaml:"resource_types,omitempty"`
	Attribute     types.String                             `tfsdk:"attribute" yaml:"attribute,omitempty"`
	Operator      types.String                             `tfsdk:"operator" yaml:"operator,omitempty"`
	Value         types.String                             `tfsdk:"value" yaml:"value,omitempty"`
	And           []FrameworkDefinitionLogicConditionModel `tfsdk:"and" yaml:"and,omitempty"`
	Or            []FrameworkDefinitionLogicConditionModel `tfsdk:"or" yaml:"or,omitempty"`
}

// *********************************************************
// Request conversion functions
// *********************************************************

func (m *ApplicationSecurityRuleModel) ToCreateOrCloneRequest(ctx context.Context, diagnostics *diag.Diagnostics) appsec.CreateOrCloneRequest {
	labels := util.StringSetToStringArray(ctx, diagnostics, m.Labels)
	if diagnostics.HasError() {
		return appsec.CreateOrCloneRequest{}
	}

	var frameworks []appsec.FrameworkData
	for _, f := range m.Frameworks {
		frameworks = append(frameworks, appsec.FrameworkData{
			Name:                   f.Name.ValueString(),
			Definition:             f.Definition.ValueString(),
			RemediationDescription: f.RemediationDescription.ValueString(),
			DefinitionLink:         f.DefinitionLink.ValueString(),
		})
	}

	return appsec.CreateOrCloneRequest{
		Name:        m.Name.ValueString(),
		Severity:    m.Severity.ValueString(),
		Scanner:     m.Scanner.ValueString(),
		Frameworks:  frameworks,
		Category:    m.Category.ValueString(),
		SubCategory: m.SubCategory.ValueString(),
		Description: m.Description.ValueString(),
		Labels:      labels,
	}
}

func (m *ApplicationSecurityRuleModel) ToUpdateRequest(ctx context.Context, diagnostics *diag.Diagnostics) appsec.UpdateRequest {
	labels := util.StringSetToStringArray(ctx, diagnostics, m.Labels)
	if diagnostics.HasError() {
		return appsec.UpdateRequest{}
	}

	// If the target rule is a default rule, only the labels field may be updated
	if !m.IsCustom.ValueBool() {
		return appsec.UpdateRequest{
			Labels: labels,
		}
	}

	var frameworks []appsec.FrameworkData
	for _, f := range m.Frameworks {
		frameworks = append(frameworks, appsec.FrameworkData{
			Name:                   f.Name.ValueString(),
			Definition:             f.Definition.ValueString(),
			RemediationDescription: f.RemediationDescription.ValueString(),
			DefinitionLink:         f.DefinitionLink.ValueString(),
		})
	}

	name := m.Name.ValueString()
	severity := m.Severity.ValueString()
	scanner := m.Scanner.ValueString()
	category := m.Category.ValueString()
	subCategory := m.SubCategory.ValueString()
	description := m.Description.ValueString()

	return appsec.UpdateRequest{
		Name:        name,
		Severity:    severity,
		Scanner:     scanner,
		Frameworks:  frameworks,
		Category:    category,
		SubCategory: subCategory,
		Description: description,
		Labels:      labels,
	}
}

// *********************************************************
// Helper functions
// *********************************************************
func (m *ApplicationSecurityRuleModel) RefreshPropertyValues(ctx context.Context, diagnostics *diag.Diagnostics, response appsec.Rule) {
	// TODO: create member functions for conversion to schema

	//var frameworkValues []FrameworkModel
	////for _, framework := range response.Frameworks {
	//for idx, framework := range response.Frameworks {
	//	// If the TERRAFORM framework exists in the resource configuration and
	//	// the TERRAFORMPLAN framework does not exist in the current resource
	//	// state, do not include the TERRAFORMPLAN framework in the updated
	//	// Frameworks value, otherwise Terraform will error on recieving an
	//	// unexpected new value
	//	if framework.Name == "TERRAFORMPLAN" && slices.ContainsFunc(m.Frameworks, func(f FrameworkModel) bool { return strings.ToUpper(f.Name.ValueString()) == "TERRAFORM" }) {
	//		continue
	//	}

	//	//var remediationDescription string
	//	//if framework.RemediationDescription == nil {
	//	//	remediationDescription = ""
	//	//} else {
	//	//	remediationDescription = *framework.RemediationDescription
	//	//}

	//	frameworkValues = append(frameworkValues, FrameworkModel{
	//		Name:                   types.StringValue(framework.Name),
	//		//Definition:             types.StringValue(framework.Definition),
	//		Definition:             m.Frameworks[idx].Definition,
	//		RemediationDescription: types.StringValue(framework.RemediationDescription),
	//		DefinitionLink:         types.StringValue(framework.DefinitionLink),
	//	})
	//}

	var conversionDiags diag.Diagnostics
	labels, diags := types.SetValueFrom(ctx, types.StringType, response.Labels)
	conversionDiags.Append(diags...)

	mitreTactics, diags := types.SetValueFrom(ctx, types.StringType, response.MitreTactics)
	conversionDiags.Append(diags...)

	mitreTechniques, diags := types.SetValueFrom(ctx, types.StringType, response.MitreTechniques)
	conversionDiags.Append(diags...)

	diagnostics.Append(conversionDiags...)

	if diagnostics.HasError() {
		return
	}

	m.Category = types.StringValue(response.Category)
	m.CloudProvider = types.StringValue(response.CloudProvider)
	m.CreatedAt = types.StringValue(response.CreatedAt.Value)
	m.Description = types.StringValue(response.Description)
	m.DetectionMethod = types.StringPointerValue(response.DetectionMethod)
	m.DocLink = types.StringValue(response.DocLink)
	m.Domain = types.StringValue(response.Domain)
	m.FindingCategory = types.StringValue(response.FindingCategory)
	m.FindingDocs = types.StringValue(response.FindingDocs)
	m.FindingTypeId = types.Int32Value(int32(response.FindingTypeId))
	m.FindingTypeName = types.StringValue(response.FindingTypeName)
	//m.Frameworks = frameworkValues
	m.Id = types.StringValue(response.Id)
	m.IsCustom = types.BoolValue(response.IsCustom)
	m.IsEnabled = types.BoolValue(response.IsEnabled)
	m.Labels = labels
	m.MitreTactics = mitreTactics
	m.MitreTechniques = mitreTechniques
	m.Name = types.StringValue(response.Name)
	m.Owner = types.StringValue(response.Owner)
	m.Scanner = types.StringValue(response.Scanner)
	m.Severity = types.StringValue(response.Severity)
	m.Source = types.StringValue(response.Source)
	m.SubCategory = types.StringValue(response.SubCategory)
	m.UpdatedAt = types.StringValue(response.UpdatedAt.Value)
}
