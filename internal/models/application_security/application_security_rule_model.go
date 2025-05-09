package models

import (
	"context"
    "fmt"

	api "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/application_security"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	//"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type ApplicationSecurityRuleModel struct {
    Category types.String `tfsdk:"category"`
    CloudProvider types.String `tfsdk:"cloud_provider"`
    CreatedAt types.String `tfsdk:"created_at"`
    Description types.String `tfsdk:"description"`
    DetectionMethod types.String `tfsdk:"detection_method"`
    DocLink types.String `tfsdk:"doc_link"`
    Domain types.String `tfsdk:"domain"`
    FindingCategory types.String `tfsdk:"finding_category"`
    FindingDocs types.String `tfsdk:"finding_docs"`
    FindingTypeId types.Int32 `tfsdk:"finding_type_id"`
    FindingTypeName types.String `tfsdk:"finding_type_name"`
    Frameworks types.Set `tfsdk:"frameworks"`
    Id types.String `tfsdk:"id"`
    IsCustom types.Bool `tfsdk:"is_custom"`
    IsEnabled types.Bool `tfsdk:"is_enabled"`
    Labels types.Set `tfsdk:"labels"`
    MitreTactics types.Set `tfsdk:"mitre_tactics"`
    MitreTechniques types.Set `tfsdk:"mitre_techniques"`
    Name types.String `tfsdk:"name"`
    Owner types.String `tfsdk:"owner"`
    Scanner types.String `tfsdk:"scanner"`
    //ScannerRuleId types.String `tfsdk:"scannerRuleId"`
    Severity types.String `tfsdk:"severity"`
    Source types.String `tfsdk:"source"`
    //SourceVersion types.String `tfsdk:"sourceVersion"`
    SubCategory types.String `tfsdk:"sub_category"`
    UpdatedAt types.String`tfsdk:"updated_at"`
}

func (m *ApplicationSecurityRuleModel) ToCreateRequest(ctx context.Context, diagnostics *diag.Diagnostics) api.CreateApplicationSecurityRuleRequest {
    frameworks := []api.ApplicationSecurityRuleFramework{}
    diagnostics.Append(m.Frameworks.ElementsAs(ctx, &frameworks, false)...)
    tflog.Debug(ctx, fmt.Sprintf("\n\n\ne type: %+v\n\n\n", m.Frameworks.ElementType(ctx)))

    labels := []string{}
    diagnostics.Append(m.Frameworks.ElementsAs(ctx, &labels, false)...)

	if diagnostics.HasError() {
		return api.CreateApplicationSecurityRuleRequest{}
	}

    return api.CreateApplicationSecurityRuleRequest{
        Name: m.Name.ValueString(),
        Severity: m.Severity.ValueString(),
        Scanner: m.Scanner.ValueString(),
        Frameworks: frameworks,
        Category: m.Category.ValueString(),
        SubCategory: m.SubCategory.ValueString(),
        Description: m.Description.ValueString(),
        Labels: labels,
    }
}

func (m *ApplicationSecurityRuleModel) RefreshPropertyValues(ctx context.Context, diagnostics *diag.Diagnostics, response api.ApplicationSecurityRule) {
    labels, diags := types.SetValueFrom(ctx, types.StringType, response.Labels)
    diagnostics.Append(diags...)

    // TODO: fix this
    frameworks, diags := types.SetValueFrom(ctx, types.StringType, response.Frameworks)
    //frameworks, diags := types.SetValueFrom(ctx, m.Frameworks.ElementType(ctx), response.Frameworks)
    //frameworks, diags := types.SetValueFrom(ctx, x, response.Frameworks)
    //frameworkValues := []attr.Value{}
    //for _, framework := range response.Frameworks {
    //    types.ObjectValueFrom(ctx, m.Frameworks.Type(ctx).ValueType(ctx), framework)
    //    frameworkValues = append(frameworkValues, types.ObjectValueMust(m.Frameworks.ElementType(), ))
    //}

    mitreTactics, diags := types.SetValueFrom(ctx, types.StringType, response.MitreTactics)
    diagnostics.Append(diags...)

    mitreTechniques, diags := types.SetValueFrom(ctx, types.StringType, response.MitreTechniques)
    diagnostics.Append(diags...)

    if diagnostics.HasError() {
        return
    }

    m.Category = types.StringValue(response.Category)
    m.CloudProvider = types.StringValue(response.CloudProvider)
    m.CreatedAt = types.StringValue(response.CreatedAt.Value)
    m.Description = types.StringValue(response.Description)
    m.DetectionMethod = types.StringValue(response.DetectionMethod)
    m.DocLink = types.StringValue(response.DocLink)
    m.Domain = types.StringValue(response.Domain)
    m.FindingCategory = types.StringValue(response.FindingCategory)
    m.FindingDocs = types.StringValue(response.FindingDocs)
    m.FindingTypeId = types.Int32Value(int32(response.FindingTypeId))
    m.FindingTypeName = types.StringValue(response.FindingTypeName)
    m.Frameworks = frameworks
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
