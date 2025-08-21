// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strconv"

	cloudOnboardingDataSources "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/data_sources/cloud_onboarding"
	models "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models/provider"
	appSecResources "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/resources/application_security"
	cloudOnboardingResources "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/resources/cloud_onboarding"
	platformResources "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/resources/platform"
	sdk "github.com/mdboynton/cortex-cloud-go/api"
	"github.com/mdboynton/cortex-cloud-go/appsec"
	"github.com/mdboynton/cortex-cloud-go/cloudonboarding"
	"github.com/mdboynton/cortex-cloud-go/log"
	"github.com/mdboynton/cortex-cloud-go/platform"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ provider.Provider = &CortexCloudProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CortexCloudProvider{
			version: version,
		}
	}
}

type CortexCloudProvider struct {
	version string
}

func (p *CortexCloudProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_url": schema.StringAttribute{
				Optional: true,
				Description: fmt.Sprintf("The API URL of your Cortex Cloud tenant. "+
					"You can retrieve this from the Cortex Cloud console by "+
					"navigating to Settings > Configurations > Integrations > "+
					"API Keys and clicking the \"Copy API URL\" button. Can "+
					"also be configured using the `%s` environment "+
					"variable.", sdk.CORTEXCLOUD_API_URL_ENV_VAR),
			},
			"api_port": schema.Int32Attribute{
				Optional:    true,
				Description: "TODO",
			},
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				Description: "The API key for the user in Cortex Cloud that the " +
					"provider will use. You can create this from the Cortex Cloud " +
					"console by navigating to Settings > Configurations > Integrations " +
					"> API Keys. Can also be configured using the `CORTEX_API_KEY` " +
					"environment variable. \n\nWARNING: If you are running the provider " +
					"with Terraform with the `TF_LOG` environment variable set to `DEBUG`, " +
					"the provider will output this value in the debug logs.",
			},
			"api_key_id": schema.Int32Attribute{
				Optional:  true,
				Sensitive: true,
				Description: "The ID of the API key provided in the \"api_key\" " +
					"argument. You can retrieve this from the Cortex Cloud console " +
					"by navigating to Settings > Configurations > Integrations > " +
					"API Keys. Can also be configured using the `CORTEX_API_KEY_ID` " +
					"environment variable.",
			},
			"insecure": schema.BoolAttribute{
				Optional: true,
				Description: "Explicity allow the provider to perform \"insecure\" " +
					"SSL requests. If omitted, the default value is `false`. Can also " +
					"be configured using the `CORTEX_TF_INSECURE` environment variable.",
			},
			"request_timeout": schema.Int32Attribute{
				Optional: true,
				Description: "Time (in seconds) to wait for requests to the Cortex " +
					"Cloud API to return before timing out. If omitted, the default value " +
					"is `60`. Can also be configured using the `CORTEX_TF_REQUEST_TIMEOUT` " +
					"environment variable.",
			},
			"request_retry_interval": schema.Int32Attribute{
				Optional: true,
				Description: "Time (in seconds) to wait between API requests in " +
					"the event of an HTTP 429 (Too Many Requests) response. If omitted, " +
					"the default value is `3`. Can also be configured using the " +
					"`CORTEX_TF_REQUEST_RETRY_INTERVAL` environment variable.",
			},
			"crash_stack_dir": schema.StringAttribute{
				Optional: true,
				Description: "The location on the filesystem where the crash stack " +
					"contents will be written in the event of the provider encountering " +
					"an unexpected error. If omitted, the default value is an empty " +
					"string, which will be interpreted as `$TMPDIR` on Unix systems (or " +
					"`/tmp` if `$TMPDIR` is empty). On Windows systems, an empty string " +
					"will be interpreted as the the first of the following values that is " +
					"non-empty, in order of evaluation: `%%TMP%%`, `%%TEMP%%`, " +
					"%%USERPROFILE%%`, or the Windows directory. Can also be configured " +
					"using the `CORTEX_TF_CRASH_STACK_DIR` environment variable.",
			},
			"config_file": schema.StringAttribute{
				Optional:    true,
				Description: "TODO",
			},
			"check_environment": schema.BoolAttribute{
				Optional:    true,
				Description: "TODO",
			},
		},
	}
}

func (p *CortexCloudProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cortexcloud"
	resp.Version = p.version
}

func (p *CortexCloudProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		cloudOnboardingResources.NewCloudIntegrationTemplateResource,
		appSecResources.NewApplicationSecurityRuleResource,
		platformResources.NewAuthenticationSettingsResource,
	}
}

func (p *CortexCloudProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		cloudOnboardingDataSources.NewCloudIntegrationInstanceDataSource,
	}
}

func (p *CortexCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Debug(ctx, "Starting provider configuration")

	// Set log level according to Terraform environment variables and print
	// warning message if debug logs are enabled
	var logLevel string
	if slices.ContainsFunc([]string{"DEBUG", "TRACE"}, func(s string) bool {
		return s == os.Getenv("TF_LOG") || s == os.Getenv("TF_LOG_PROVIDER")
	}) {
		logLevel = "debug"
		tflog.Warn(ctx, "Debug logging enabled. Be aware that your API key "+
			"and key ID will be visible in the provider log "+
			"output!")
	} else {
		logLevel = "quiet"
	}

	// Retrieve configuration values from provider block
	var providerConfig models.CortexCloudProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &providerConfig)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		clientConfig *sdk.Config
		err          error
	)

	// If the config_file argument is defined, initialize SDK client config
	// using values stored in the provided file
	if !providerConfig.ConfigFile.IsNull() && !providerConfig.ConfigFile.IsUnknown() {
		configFile := providerConfig.ConfigFile.ValueString()

		if configFile != "" {
			clientConfig, err = sdk.NewConfigFromFile(configFile, providerConfig.CheckEnvironment.ValueBool())
		}
		// Otherwise, configure SDK client using values from provider block and
		// environment variables
	} else {
		// Get values of required configuration arguments
		apiUrl := providerConfig.ApiUrl.ValueString()
		apiKey := providerConfig.ApiKey.ValueString()
		apiKeyId := int(providerConfig.ApiKeyId.ValueInt32())

		if apiUrl == "" {
			if v := os.Getenv("CORTEX_API_URL"); v == "" {
				tflog.Error(ctx, `No value provided for required configuration argument "api_url" in provider block or CORTEX_API_URL environment variable.`)
			} else {
				apiUrl = v
			}
		}
		tflog.Debug(ctx, fmt.Sprintf(`CORTEX_API_URL="%s"`, apiUrl))

		if apiKey == "" {
			if v := os.Getenv("CORTEX_API_KEY"); v == "" {
				tflog.Error(ctx, `No value provided for required configuration argument "api_key" in provider block or CORTEX_API_KEY environment variable.`)
			} else {
				apiKey = v
			}
		}
		tflog.Debug(ctx, fmt.Sprintf(`CORTEX_API_KEY="%s"`, apiKey))

		if apiKeyId == 0 {
			if v := os.Getenv("CORTEX_API_KEY_ID"); v == "" {
				tflog.Error(ctx, `No value provided for required configuration argument "api_key_id" in provider block or CORTEX_API_KEY_ID environment variable.`)
			} else {
				apiKeyId, err = strconv.Atoi(v)
				if err != nil {
					tflog.Error(ctx, fmt.Sprintf(`Error occured while converting CORTEX_API_KEY_ID value to int: %s`, err.Error()))
				}
			}
		}
		tflog.Debug(ctx, fmt.Sprintf(`CORTEX_API_KEY_ID=%d`, apiKeyId))

		clientConfig = sdk.NewConfig(
			apiUrl,
			apiKey,
			apiKeyId,
			providerConfig.CheckEnvironment.ValueBool(),
			sdk.WithApiPort(int(providerConfig.ApiPort.ValueInt32())),
			sdk.WithSkipVerifyCertificate(providerConfig.Insecure.ValueBool()),
			sdk.WithTimeout(int(providerConfig.RequestTimeout.ValueInt32())),
			//sdk.WithRetryMaxDelay(providerConfig.RetryMaxDelay),
			sdk.WithCrashStackDir(providerConfig.CrashStackDir.ValueString()),
			sdk.WithLogger(log.TflogAdapter{}),
			sdk.WithLogLevel(logLevel),
		)
	}

	// Validate SDK client configuration
	if err = clientConfig.Validate(); err != nil {
		resp.Diagnostics.AddError("Cortex Cloud SDK Configuration Error", err.Error())
		return
	}

	// Initialize SDK clients
	clients := models.CortexCloudSDKClients{}

	appSecClient, err := appsec.NewClient(clientConfig)
	if err != nil {
		resp.Diagnostics.AddError("Cortex Cloud API Setup Error", err.Error())
		return
	}

	cloudOnboardingClient, err := cloudonboarding.NewClient(clientConfig)
	if err != nil {
		resp.Diagnostics.AddError("Cortex Cloud API Setup Error", err.Error())
		return
	}

	platformClient, err := platform.NewClient(clientConfig)
	if err != nil {
		resp.Diagnostics.AddError("Cortex Cloud API Setup Error", err.Error())
		return
	}

	tflog.Debug(ctx, "Cortex Cloud API client setup complete")

	// Attach SDK clients to model
	clients.AppSec = appSecClient
	clients.CloudOnboarding = cloudOnboardingClient
	clients.Platform = platformClient

	// Assign clients model pointer to ProviderData to allow resources and
	// data sources to access SDK functions
	resp.DataSourceData = &clients
	resp.ResourceData = &clients
}
