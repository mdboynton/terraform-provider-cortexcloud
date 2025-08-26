// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package acceptance

import (
	"fmt"
	"strings"

	"github.com/mdboynton/cortex-cloud-go/enums"
)

var (
	AppSecRule1Name     = "test-rule"
	AppSecRule1Description     = "test description"
	AppSecRule1Category = enums.IacCategoryCompute.String()
	AppSecRule1SubCategory = enums.IacSubCategoryComputeOverprovisioned.String()
	AppSecRule1Scanner  = enums.ScannerIAC.String()
	AppSecRule1Severity = enums.SeverityInfo.String()
	AppSecRule1Labels = []string{
		"label1", 
		"label2",
	}
	AppSecRule1Framework1Name   = enums.FrameworkNameTerraform.String()
	AppSecRule1Framework1Definition = "scope:\n  provider: aws\ndefinition:\n  or:\n    - cond_type: attribute\n      resource_types:\n        - aws_instance\n      attribute: instance_type\n      operator: equals\n      value: t2.micro\n"
	AppSecRule1Framework1DefinitionLink = "http://docs.com/framework"
	AppSecRule1Framework1RemediationDescription = "fix it"

	AppSecUnitTestConfigTmpl = `provider "cortexcloud" {
	api_url = %s
	api_port = 443
	api_key = "test"
	api_key_id = 123
}
resource "cortexcloud_application_security_rule" "test" {
	name     = %s
	description = %s
	severity = %s
	scanner  = %s
	frameworks = [
		{
			name = %s
			definition = %s
			definition_link = %s
			remediation_description = %s
		}
	]
	category = %s
	sub_category = %s
	labels   = %s
}`

	AppSecUnitTestCreateOrCloneResponseTmpl = `{
	"id": "test-rule-id",
	"name": %s,
	"category": %s,
	"cloudProvider": "aws",
	"createdAt": {
		"value": "2025-08-26T00:00:00.000Z"
	},
	"description": %s,
	"detectionMethod": "some-method",
	"docLink": "http://docs.com",
	"domain": "test-domain",
	"findingCategory": "test-finding-category",
	"findingDocs": "http://docs.com/finding",
	"findingTypeId": 123,
	"findingTypeName": "test-finding",
	"frameworks": [{
		"name": %s,
		"definition": %s,
		"definitionLink": %s,
		"remediationDescription": %s
	}],
	"isCustom": true,
	"isEnabled": true,
	"labels": %s,
	"mitreTactics": ["tactic1"],
	"mitreTechniques": ["technique1"],
	"owner": "test-owner",
	"scanner": %s,
	"severity": %s,
	"source": "custom",
	"subCategory": %s,
	"updatedAt": {
		"value": "2025-08-26T00:00:00.000Z"
	}
}`

	AppSecUnitTestGetResponseTmpl = `{
	"id": "test-rule-id",
	"name": %s,
	"category": %s,
	"cloudProvider": "aws",
	"createdAt": {
		"value": "2025-08-26T00:00:00.000Z"
	},
	"description": %s,
	"detectionMethod": "some-method",
	"docLink": "http://docs.com",
	"domain": "test-domain",
	"findingCategory": "test-finding-category",
	"findingDocs": "http://docs.com/finding",
	"findingTypeId": 123,
	"findingTypeName": "test-finding",
	"frameworks": [{
		"name": %s,
		"definition": %s,
		"definitionLink": %s,
		"remediationDescription": %s
	}],
	"isCustom": true,
	"isEnabled": true,
	"labels": %s,
	"mitreTactics": ["tactic1"],
	"mitreTechniques": ["technique1"],
	"owner": "test-owner",
	"scanner": %s,
	"severity": %s,
	"source": "custom",
	"subCategory": %s,
	"updatedAt": {
		"value": "2025-08-26T00:00:00.000Z"
	}
}`

)

// AppSecRuleLabelsHCL returns the provided labels as a HCL string.
func AppSecRuleLabelsHCL(labels []string) string {
	return fmt.Sprintf(`["%s"]`, strings.Join(labels, "\", \""))
}
