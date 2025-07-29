// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"slices"
	"strings"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/util"
	"github.com/mdboynton/cortex-cloud-go/appsec"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
// Schema
// *********************************************************
func (m ApplicationSecurityRuleModel) GetSchema() schema.Schema {
	return schema.Schema{
		Description: "TODO",
		Attributes: map[string]schema.Attribute{
			"category": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Required:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cloud_provider": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				// TODO: rename to "Impact"
				// TODO: validation
				Description: "TODO",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Default: stringdefault.StaticString(""),
			},
			"detection_method": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"doc_link": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"finding_category": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"finding_docs": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"finding_type_id": schema.Int32Attribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.Int32{
					int32planmodifier.UseStateForUnknown(),
				},
			},
			"finding_type_name": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"frameworks": FrameworkModel{}.GetSchema(),
			"id": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_custom": schema.BoolAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"is_enabled": schema.BoolAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"labels": schema.SetAttribute{
				Description: "TODO",
				Required:    true,
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"mitre_tactics": schema.SetAttribute{
				Description: "TODO",
				Computed:    true,
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"mitre_techniques": schema.SetAttribute{
				Description: "TODO",
				Computed:    true,
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				// TODO: validation
				// TODO: should this be modifiable? does it require replace?
				Description: "TODO",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"owner": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"scanner": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"severity": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Required:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"source": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"sub_category": schema.StringAttribute{
				// TODO: validation
				// The valid inputs for this attribute are determined by the "category" value
				Description: "TODO",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				//PlanModifiers: []planmodifier.String{
				//	stringplanmodifier.UseStateForUnknown(),
				//},
			},
		},
	}
}

func (m FrameworkModel) GetSchema() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		// TODO: validator to make sure this is not null
		Description: "TODO",
		Required:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Description: "TODO",
					Required:    true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				//"definition": FrameworkDefinitionModel{}.GetSchema(),
				"definition": schema.StringAttribute{
					Description: "TODO",
					//Required:    true,
					Optional: true,
					Computed: true,
					//PlanModifiers: []planmodifier.String{
					//	stringplanmodifier.UseStateForUnknown(),
					//},
				},
				"definition_link": schema.StringAttribute{
					Description: "TODO",
					Optional:    true,
					Computed:    true,
					Default:     stringdefault.StaticString(""),
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				"remediation_description": schema.StringAttribute{
					Description: "TODO",
					Optional:    true,
					Computed:    true,
					Default:     stringdefault.StaticString(""),
				},
			},
		},
		//PlanModifiers: []planmodifier.Set{
		//	setplanmodifier.UseStateForUnknown(),
		//},
	}
}

func (m FrameworkDefinitionModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: "TODO",
		Required:    true,
		Attributes: map[string]schema.Attribute{
			"metadata":   FrameworkDefinitionMetadataModel{}.GetSchema(),
			"scope":      FrameworkDefinitionScopeModel{}.GetSchema(),
			"definition": FrameworkDefinitionLogicModel{}.GetSchema(),
		},
	}
}

func (m FrameworkDefinitionMetadataModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: "TODO",
		Optional:    true,
		Computed:    true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "TODO",
				Optional:    true,
				Computed:    true,
			},
			"guidelines": schema.StringAttribute{
				Description: "TODO",
				Optional:    true,
				Computed:    true,
			},
			"category": schema.StringAttribute{
				Description: "TODO",
				Optional:    true,
				Computed:    true,
			},
			"severity": schema.StringAttribute{
				Description: "TODO",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func (m FrameworkDefinitionScopeModel) GetSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: "TODO",
		Required:    true,
		Attributes: map[string]schema.Attribute{
			"provider": schema.StringAttribute{
				Description: "TODO",
				Required:    true,
			},
		},
	}
}

// func (m FrameworkDefinitionLogicModel) GetSchema() schema.SingleNestedAttribute {
func (m FrameworkDefinitionLogicModel) GetSchema() schema.DynamicAttribute {
	return schema.DynamicAttribute{
		Description: "TODO",
		Required:    true,
		//Attributes: map[string]schema.Attribute{
		//	"and": FrameworkDefinitionLogicConditionModel{}.GetSchema(),
		//	"or":  FrameworkDefinitionLogicConditionModel{}.GetSchema(),
		//},
	}
}

func (m FrameworkDefinitionLogicConditionModel) GetSchema() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Description: "TODO",
		Required:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"condition_type": schema.StringAttribute{
					Description: "TODO",
					Required:    true,
				},
				"resource_types": schema.ListAttribute{
					Description: "TODO",
					Required:    true,
					ElementType: types.StringType,
				},
				"attribute": schema.StringAttribute{
					Description: "TODO",
					Required:    true,
				},
				"operator": schema.StringAttribute{
					Description: "TODO",
					Required:    true,
				},
				"value": schema.StringAttribute{
					Description: "TODO",
					Required:    true,
				},
				"and": FrameworkDefinitionLogicConditionModel{}.GetSchema(),
				"or":  FrameworkDefinitionLogicConditionModel{}.GetSchema(),
			},
		},
	}
}

// TODO: put tfsdk tags on API models and replace explicit object creation with ElementsAs
//func (m FrameworkModel) ToTerraform(ctx context.Context, diagnostics *diag.Diagnostics, name, category, severity string) appsec.Framework {
//	var metadata *api.FrameworkDefinitionMetadata
//	if m.Definition.Metadata == nil {
//		metadata = nil
//	} else {
//		metadata = &api.FrameworkDefinitionMetadata{
//			Name:       m.Definition.Metadata.Name.ValueString(),
//			Guidelines: m.Definition.Metadata.Guidelines.ValueStringPointer(),
//			Category:   m.Definition.Metadata.Category.ValueString(),
//			Severity:   m.Definition.Metadata.Severity.ValueString(),
//		}
//	}
//
//	//switch underlyingType := m.Definition.Definition.UnderlyingValue().(type) {
//	//case types.Object:
//	//	break
//	//default:
//	//	diagnostics.AddError(
//	//		"Invalid Type",
//	//		fmt.Sprintf("Expected type %T, recieved %T", conditionsType, underlyingType),
//	//	)
//	//	return api.Framework{}
//	//}
//
//	andConditions, orConditions := []api.FrameworkDefinitionLogicCondition{}, []api.FrameworkDefinitionLogicCondition{}
//	//for _, andCondition := range m.Definition.Definition.And {
//	//	andConditions = append(andConditions, andCondition.ToTerraform(ctx, diagnostics))
//	//}
//
//	//for _, orCondition := range m.Definition.Definition.And {
//	//	orConditions = append(orConditions, orCondition.ToTerraform(ctx, diagnostics))
//	//}
//
//	//var definitionValue []string
//	//diagnostics.Append(m.Definition.Definition.Value.ElementsAs(ctx, &definitionValue, false)...)
//	//if diagnostics.HasError() {
//	//	return api.Framework{}
//	//}
//
//	return api.Framework{
//		Name: m.Name.ValueString(),
//		Definition: api.FrameworkDefinition{
//			Metadata: metadata,
//			Scope: api.FrameworkDefinitionScope{
//				Provider: m.Definition.Scope.Provider.ValueString(),
//			},
//			Definition: api.FrameworkDefinitionLogic{
//				And: andConditions,
//				//ConditionType: m.Definition.Definition.ConditionType.ValueString(),
//				Or: orConditions,
//				//Value:         definitionValue,
//			},
//		},
//		DefinitionLink:         m.DefinitionLink.ValueString(),
//		RemediationDescription: m.RemediationDescription.ValueStringPointer(),
//	}
//}

//func (m FrameworkDefinitionLogicConditionModel) ToTerraform(ctx context.Context, diagnostics *diag.Diagnostics) api.FrameworkDefinitionLogicCondition {
//	var (
//		resourceTypes       []string
//		nestedAndConditions []api.FrameworkDefinitionLogicCondition
//		nestedOrConditions  []api.FrameworkDefinitionLogicCondition
//	)
//
//	diagnostics.Append(m.ResourceTypes.ElementsAs(ctx, &resourceTypes, false)...)
//
//	for _, nestedAndCondition := range m.And {
//		nestedAndConditions = append(nestedAndConditions, nestedAndCondition.ToTerraform(ctx, diagnostics))
//	}
//
//	for _, nestedOrCondition := range m.Or {
//		nestedOrConditions = append(nestedOrConditions, nestedOrCondition.ToTerraform(ctx, diagnostics))
//	}
//
//	if diagnostics.HasError() {
//		return api.FrameworkDefinitionLogicCondition{}
//	}
//
//	return api.FrameworkDefinitionLogicCondition{
//		ConditionType: m.ConditionType.ValueString(),
//		ResourceTypes: resourceTypes,
//		Attribute:     m.Attribute.ValueString(),
//		Operator:      m.Operator.ValueString(),
//		Value:         m.Value.ValueString(),
//		And:           nestedAndConditions,
//		Or:            nestedOrConditions,
//	}
//}
//
//func (m *FrameworkModel) FromAPI(ctx context.Context, diagnostics *diag.Diagnostics, framework appsecAPI.Framework) {
//	var metadata *FrameworkDefinitionMetadataModel
//	if m.Definition.Metadata == nil {
//		metadata = nil
//	} else {
//		metadata = &FrameworkDefinitionMetadataModel{
//			Name:       types.StringValue(framework.Definition.Metadata.Name),
//			Guidelines: types.StringPointerValue(m.Definition.Metadata.Guidelines.ValueStringPointer()),
//			Category:   types.StringValue(m.Definition.Metadata.Category.ValueString()),
//			Severity:   types.StringValue(m.Definition.Metadata.Severity.ValueString()),
//		}
//	}
//
//	andConditions, orConditions := []FrameworkDefinitionLogicConditionModel{}, []FrameworkDefinitionLogicConditionModel{}
//	for _, andCondition := range framework.Definition.Definition.And {
//		nestedAndCondition := FrameworkDefinitionLogicConditionModel{}
//		nestedAndCondition.FromAPI(ctx, diagnostics, andCondition)
//		andConditions = append(andConditions, nestedAndCondition)
//	}
//
//	for _, orCondition := range framework.Definition.Definition.And {
//		nestedAndCondition := FrameworkDefinitionLogicConditionModel{}
//		nestedAndCondition.FromAPI(ctx, diagnostics, orCondition)
//		orConditions = append(orConditions, nestedAndCondition)
//	}
//
//	m.Name = types.StringValue(framework.Name)
//	m.Definition = FrameworkDefinitionModel{
//		Metadata: metadata,
//		Scope: FrameworkDefinitionScopeModel{
//			Provider: types.StringValue(m.Definition.Scope.Provider.ValueString()),
//		},
//		//Definition: FrameworkDefinitionLogicModel{
//		//	//And: andConditions,
//		//	//Or:  orConditions,
//		//},
//	}
//	m.DefinitionLink = types.StringValue(framework.DefinitionLink)
//	m.RemediationDescription = types.StringPointerValue(framework.RemediationDescription)
//}
//
//func (m *FrameworkDefinitionLogicConditionModel) FromAPI(ctx context.Context, diagnostics *diag.Diagnostics, frameworkDefinition api.FrameworkDefinitionLogicCondition) {
//	var (
//		nestedAndConditions []FrameworkDefinitionLogicConditionModel
//		nestedOrConditions  []FrameworkDefinitionLogicConditionModel
//	)
//
//	resourceTypes, diags := types.ListValueFrom(ctx, types.StringType, frameworkDefinition.ResourceTypes)
//	diagnostics.Append(diags...)
//
//	for _, nestedAndCondition := range frameworkDefinition.And {
//		nestedAndYamlConditionModel := FrameworkDefinitionLogicConditionModel{}
//		nestedAndYamlConditionModel.FromAPI(ctx, diagnostics, nestedAndCondition)
//		nestedAndConditions = append(nestedAndConditions, nestedAndYamlConditionModel)
//	}
//
//	for _, nestedOrCondition := range frameworkDefinition.Or {
//		nestedOrYamlConditionModel := FrameworkDefinitionLogicConditionModel{}
//		nestedOrYamlConditionModel.FromAPI(ctx, diagnostics, nestedOrCondition)
//		nestedOrConditions = append(nestedOrConditions, nestedOrYamlConditionModel)
//	}
//
//	if diagnostics.HasError() {
//		return
//	}
//
//	m.ConditionType = types.StringValue(m.ConditionType.ValueString())
//	m.ResourceTypes = resourceTypes
//	m.Attribute = types.StringValue(m.Attribute.ValueString())
//	m.Operator = types.StringValue(m.Operator.ValueString())
//	m.Value = types.StringValue(m.Value.ValueString())
//	m.And = nestedAndConditions
//	m.Or = nestedOrConditions
//}

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

	var frameworkValues []FrameworkModel
	for _, framework := range response.Frameworks {
		// If the TERRAFORM framework exists in the resource configuration and
		// the TERRAFORMPLAN framework does not exist in the current resource
		// state, do not include the TERRAFORMPLAN framework in the updated
		// Frameworks value, otherwise Terraform will error on recieving an
		// unexpected new value
		if framework.Name == "TERRAFORMPLAN" && slices.ContainsFunc(m.Frameworks, func(f FrameworkModel) bool { return strings.ToUpper(f.Name.ValueString()) == "TERRAFORM" }) {
			continue
		}

		var remediationDescription string
		if framework.RemediationDescription == nil {
			remediationDescription = ""
		} else {
			remediationDescription = *framework.RemediationDescription
		}

		frameworkValues = append(frameworkValues, FrameworkModel{
			Name:                   types.StringValue(framework.Name),
			Definition:             types.StringValue(framework.Definition),
			RemediationDescription: types.StringValue(remediationDescription),
			DefinitionLink:         types.StringValue(framework.DefinitionLink),
		})
	}

	labels, diags := types.SetValueFrom(ctx, types.StringType, response.Labels)
	mitreTactics, diags := types.SetValueFrom(ctx, types.StringType, response.MitreTactics)
	mitreTechniques, diags := types.SetValueFrom(ctx, types.StringType, response.MitreTechniques)
	diagnostics.Append(diags...)

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
	m.Frameworks = frameworkValues
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
