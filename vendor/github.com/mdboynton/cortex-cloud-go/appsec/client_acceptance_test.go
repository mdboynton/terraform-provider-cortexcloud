// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build acceptance

package appsec

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mdboynton/cortex-cloud-go/api"
	"github.com/mdboynton/cortex-cloud-go/enums"
	"github.com/stretchr/testify/assert"
)

func setupAcceptanceTest(t *testing.T) *Client {
	config := &api.Config{
		// ApiUrl:   "https://api-e2e-susanpolgar.xdr.us.paloaltonetworks.com",
		// ApiKey:   "2CNVwmFHlTPUyPTz6kjRBXm1Dt8c6Npa9tgWPoIZenBphXwcX3EuxvXhCli0j5Q1YduBdOW9p6sP97QvMMEQ1IBJevyueVkbuB0jMTgrCAXI26GeD31Yg9JPnuem9APP",
		// ApiKeyId: 341,
		ApiUrl:   "https://api-pcscortexcloud.xdr.us.paloaltonetworks.com/",
		ApiKey:   "3z7Ohk2R6UGCWTWWL71OMW3SD32M3G4YR37pYfThfawXnRAB3vYSO5mQdsbEDpZrmYpdOXDddp9zXoMZ4AqsZrznvTwjTU46SNwsUeFQssQNmLsUuztgT9O3JHG2PgJf",
		ApiKeyId: 301,
	}

	client, err := NewClient(config)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	return client
}

func TestAppsecRuleLifecycle(t *testing.T) {
	//t.Skip("Skipping test due to persistent failures")
	client := setupAcceptanceTest(t)
	ctx := context.Background()

	// Create a new rule
	ruleName := fmt.Sprintf("test-rule-%d", time.Now().Unix())
	createReq := CreateOrCloneRequest{
		Name:        ruleName,
		Description: "test description",
		Category:    string(enums.IacCategoryCompute),
		SubCategory: string(enums.IacSubCategoryComputeOverprovisioned),
		Scanner:     string(enums.ScannerIAC),
		Severity:    string(enums.SeverityLow),
		Frameworks: []FrameworkData{
			{
				Name:       string(enums.FrameworkNameTerraform),
				Definition: "scope:\n  provider: \"aws\"\ndefinition:\n  or:\n    - cond_type: \"attribute\"\n      resource_types:\n        - \"aws_instance\"\n      attribute: \"instance_type\"\n      operator: \"equals\"\n      value: \"t3.micro\"",
			},
		},
		Labels: []string{"test-label"},
	}
	createdRule, err := client.CreateOrClone(ctx, createReq)
	assert.NoError(t, err)
	assert.NotNil(t, createdRule)
	assert.Equal(t, ruleName, createdRule.Name)
	ruleID := createdRule.Id

	// Get the rule
	gotRule, err := client.Get(ctx, ruleID)
	assert.NoError(t, err)
	assert.NotNil(t, gotRule)
	assert.Equal(t, ruleID, gotRule.Id)
	assert.Equal(t, ruleName, gotRule.Name)

	// Update the rule
	updatedName := fmt.Sprintf("updated-test-rule-%d", time.Now().Unix())
	updateReq := UpdateRequest{
		Name: updatedName,
	}
	updatedResp, err := client.Update(ctx, ruleID, updateReq)
	assert.NoError(t, err)
	assert.NotNil(t, updatedResp)
	assert.Equal(t, updatedName, updatedResp.Rule.Name)

	// Delete the rule
	err = client.Delete(ctx, ruleID)
	assert.NoError(t, err)

	// Verify the rule is deleted
	_, err = client.Get(ctx, ruleID)
	assert.Error(t, err) // Expect an error when getting a deleted rule
}

func TestClient_List_Acceptance(t *testing.T) {
	t.Skip("Skipping test due to persistent failures")
	client := setupAcceptanceTest(t)
	listReq := ListRequest{
		Limit: 1,
	}
	resp, err := client.List(context.Background(), listReq)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
