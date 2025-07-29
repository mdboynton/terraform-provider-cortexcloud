// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package planmodifiers

//import (
//	"context"
//	"fmt"
//	"strings"
//
//	"gopkg.in/yaml.v3"
//
//	appSecApi "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/application_security"
//	models "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models/application_security"
//
//	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
//	"github.com/hashicorp/terraform-plugin-framework/types"
//	//"github.com/hashicorp/terraform-plugin-framework/path"
//)
//
////var (
////	_ validator.String = NullIfAlsoSetPlanModifier{}
////	_ validator.Bool   = NullIfAlsoSetPlanModifier{}
////)
//
//func AddFrameworkDefinitionMetadata() planmodifier.String {
//	return &addFrameworkDefinitionMetadata{}
//}
//
//type addFrameworkDefinitionMetadata struct {
//}
//
//func (m *addFrameworkDefinitionMetadata) Description(ctx context.Context) string {
//	return m.MarkdownDescription(ctx)
//}
//
//func (m *addFrameworkDefinitionMetadata) MarkdownDescription(context.Context) string {
//	return ""
//}
//
//func (m *addFrameworkDefinitionMetadata) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
//	var definition customTypes.YamlString
//	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path, &definition)...)
//	if resp.Diagnostics.HasError() {
//		return
//	}
//
//	definitionYaml := appSecApi.ApplicationSecurityRuleYaml{}
//	err := yaml.Unmarshal([]byte(definition.ValueString()), &definitionYaml)
//	if err != nil {
//		resp.Diagnostics.AddAttributeError(
//			req.Path,
//			"Value Conversion Error",
//			fmt.Sprintf("Failed to convert rule definition \"%s\": %s", definition.ValueString(), err.Error()),
//		)
//		return
//	}
//
//	if definitionYaml.Metadata == nil {
//		var plan models.ApplicationSecurityRuleModel
//		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
//		if resp.Diagnostics.HasError() {
//			return
//		}
//
//		var guidelines *string
//		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.ParentPath().AtName("remediation_description"), &guidelines)...)
//
//		definitionYaml.Metadata = &appSecApi.ApplicationSecurityRuleYamlMetadata{
//			Name:       strings.ToLower(plan.Name.ValueString()),
//			Guidelines: guidelines,
//			Category:   strings.ToLower(plan.Category.ValueString()),
//			Severity:   strings.ToLower(plan.Severity.ValueString()),
//		}
//
//		var updatedDefinitionYaml strings.Builder
//		encoder := yaml.NewEncoder(&updatedDefinitionYaml)
//		encoder.SetIndent(2)
//		err := encoder.Encode(definitionYaml)
//		if err != nil {
//			resp.Diagnostics.AddAttributeError(
//				req.Path,
//				"Value Conversion Error",
//				fmt.Sprintf("Failed to convert modified rule definition back into YAML string: %s", err.Error()),
//			)
//		}
//		updateDefinitionString := updatedDefinitionYaml.String()
//		resp.PlanValue = types.StringValue(updateDefinitionString)
//	}
//}
