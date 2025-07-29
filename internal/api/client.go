// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	models "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models/provider"

	"github.com/mdboynton/cortex-cloud-go/api"
	"github.com/mdboynton/cortex-cloud-go/appsec"
	"github.com/mdboynton/cortex-cloud-go/cloudonboarding"
)

// TODO: update docstring
// Client implements the HTTP client that will be used to execute
// requests to the Cortex Cloud API.
type Client struct {
	Config          api.Config
	AppSec          *appsec.Client
	CloudOnboarding *cloudonboarding.Client
}

func (c *Client) Setup(ctx context.Context, diags *diag.Diagnostics, providerConfig models.CortexCloudProviderModel) {
	tflog.Debug(ctx, "Setting up Cortex Cloud API client")

	var (
		clientConfig *api.Config
		err          error
	)

	if !providerConfig.ConfigFile.IsNull() && !providerConfig.ConfigFile.IsUnknown() {
		configFile := providerConfig.ConfigFile.ValueString()

		if configFile != "" {
			clientConfig, err = api.NewConfigFromFile(configFile, providerConfig.CheckEnvironment.ValueBool())
		}
	} else {
		clientConfig = api.NewConfig(
			providerConfig.ApiUrl.ValueString(),
			providerConfig.ApiKey.ValueString(),
			int(providerConfig.ApiKeyId.ValueInt32()),
			providerConfig.CheckEnvironment.ValueBool(),
			api.WithApiPort(int(providerConfig.ApiPort.ValueInt32())),
			api.WithSkipVerifyCertificate(providerConfig.Insecure.ValueBool()),
			api.WithTimeout(int(providerConfig.RequestTimeout.ValueInt32())),
			//api.WithRetryMaxDelay(providerConfig.RetryMaxDelay),
			api.WithCrashStackDir(providerConfig.CrashStackDir.ValueString()),
		)
	}

	if err = clientConfig.Validate(); err != nil {
		diags.AddError("Cortex Cloud SDK Configuration Error", err.Error())
		return
	}

	if appSecClient, err := appsec.NewClient(clientConfig); err != nil {
		diags.AddError("Cortex Cloud API Setup Error", err.Error())
		return
	} else {
		c.AppSec = appSecClient
	}

	if cloudOnboardingClient, err := cloudonboarding.NewClient(clientConfig); err != nil {
		diags.AddError("Cortex Cloud API Setup Error", err.Error())
		return
	} else {
		c.CloudOnboarding = cloudOnboardingClient
	}

	tflog.Debug(ctx, "Cortex Cloud API client setup complete")
}
