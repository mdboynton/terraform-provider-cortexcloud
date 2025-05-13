package provider

import (
	"context"
	//"fmt"
	"os"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
	appSecResources "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/resources/application_security"
	cloudIntegrationResources "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/resources/cloud_onboarding/cloud_integration"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	//"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	ApiUrlEnvVar               = "CORTEX_API_BASE_URL"
	ApiKeyEnvVar               = "CORTEX_API_KEY"
	ApiKeyIdEnvVar             = "CORTEX_API_KEY_ID"
	InsecureEnvVar             = "CORTEX_TF_INSECURE"
	RequestTimeoutEnvVar       = "CORTEX_TF_REQUEST_TIMEOUT"
	RequestRetryIntervalEnvVar = "CORTEX_TF_REQUEST_RETRY_INTERVAL"
	CrashStackDirEnvVar        = "CORTEX_TF_CRASH_STACK_DIR"

	InsecureDefault             = false
	RequestTimeoutDefault       = 60
	RequestRetryIntervalDefault = 3
	CrashStackDirDefault        = ""
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

type CortexCloudProviderModel struct {
	ApiUrl               types.String `tfsdk:"api_url"`
	ApiKey               types.String `tfsdk:"api_key"`
	ApiKeyId             types.Int32  `tfsdk:"api_key_id"`
	Insecure             types.Bool   `tfsdk:"insecure"`
	RequestTimeout       types.Int32  `tfsdk:"request_timeout"`
	RequestRetryInterval types.Int32  `tfsdk:"request_retry_interval"`
	CrashStackDir        types.String `tfsdk:"crash_stack_dir"`
}

func (p *CortexCloudProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_url": schema.StringAttribute{
				Optional: true,
				Description: "The FQDN of your Cortex Cloud tenant. You can retrieve " +
					"this from the Cortex Cloud console by navigating to Settings > " +
					"Configurations > Integrations > API Keys and clicking the " +
					"\"Copy API URL\" button. Can also be configured using the " +
					"`CORTEX_API_BASE_URL` environment variable.",
			},
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				Description: "The API key for the user in Cortex Cloud that the " +
					"provider will use. You can create this from the Cortex Cloud " +
					"console by navigating to Settings > Configurations > Integrations " +
					"> API Keys. Can also be configured using the `CORTEX_API_KEY` " +
					"environment variable. \n\nNote: If you are running the provider " +
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
		},
	}
}

func (p *CortexCloudProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cortexcloud"
	resp.Version = p.version
}

//func (p *CortexCloudProvider) ValidateConfig(ctx context.Context, req provider.ValidateConfigRequest, resp *provider.ValidateConfigResponse) {
//    util.HCLogDebug(ctx, "ValidateConfig")
//
//    var config CortexCloudProviderModel
//    resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
//    if resp.Diagnostics.HasError() {
//        return
//    }
//
//    // Validate required provider configuration arguments
//    if (config.ConsoleUrl.IsNull() || config.ConsoleUrl.IsUnknown()) {
//        resp.Diagnostics.AddError(
//            "Missing Provider Configuration Attribute",
//            fmt.Sprintf("Missing value for provider configuration attribute 'console_url'. Please provide a value in the provider configuration block, in the file specified by the 'config_file' attribute, or set the %s environment variable.", ConsoleUrlEnvVar),
//        )
//    }
//
//    if (config.Username.IsNull() || config.Username.IsUnknown()) {
//        resp.Diagnostics.AddError(
//            "Missing Provider Configuration Attribute",
//            fmt.Sprintf("Missing value for provider configuration attribute 'username'. Please provide a value in the provider configuration block, in the file specified by the 'config_file' attribute, or set the %s environment variable.", UsernameEnvVar),
//        )
//    }
//
//    if (config.Password.IsNull() || config.Password.IsUnknown()) {
//        resp.Diagnostics.AddError(
//            "Missing Provider Configuration Attribute",
//            fmt.Sprintf("Missing value for provider configuration attribute 'username'. Please provide a value in the provider configuration block, in the file specified by the 'config_file' attribute, or set the %s environment variable.", UsernameEnvVar),
//        )
//    }
//
//    if resp.Diagnostics.HasError() {
//        return
//    }
//
//    // Check optional parameters and assign default value if null or unknown
//    if (config.Insecure.IsNull() || config.Insecure.IsUnknown()) {
//        resp.Diagnostics.AddWarning(
//            "Missing Optional Provider Configuration Attribute",
//            fmt.Sprintf("Missing value for provider configuration attribute 'insecure'. Using default value of %t.", InsecureDefault),
//        )
//        config.Insecure = types.BoolValue(InsecureDefault)
//    }
//
//    if (config.RequestTimeout.IsNull() || config.RequestTimeout.IsUnknown()) {
//        resp.Diagnostics.AddWarning(
//            "Missing Optional Provider Configuration Attribute",
//            fmt.Sprintf("Missing value for provider configuration attribute 'request_timeout'. Using default value of %d.", RequestTimeoutDefault),
//        )
//        config.RequestTimeout = types.Int32Value(RequestTimeoutDefault)
//    }
//}

//func (p *CortexCloudProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
//    return []func() ephemeral.EphemeralResource{
//        //resources.NewCloudOnboardingIntegrationTemplateResource,
//    }
//}

func (p *CortexCloudProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		cloudIntegrationResources.NewCloudIntegrationInstanceResource,
		appSecResources.NewApplicationSecurityRuleResource,
	}
}

func (p *CortexCloudProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *CortexCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Debug(ctx, "Starting provider configuration")

	var (
		apiConfig api.CortexCloudAPIClientConfig
		diags     diag.Diagnostics
	)

	//if p.version == "test" {
	//	consoleUrl := os.Getenv(ConsoleUrlEnvVar)
	//	username := os.Getenv(UsernameEnvVar)
	//	password := os.Getenv(PasswordEnvVar)
	//	insecure := true
	//	requestTimeout := 60

	//	config = api.PrismaCloudComputeAPIClientConfig{
	//		ConsoleURL:     &consoleUrl,
	//		Username:       &username,
	//		Password:       &password,
	//		Insecure:       &insecure,
	//		RequestTimeout: &requestTimeout,
	//	}

	//	client, err := api.NewPrismaCloudComputeAPIClient(ctx, config)
	//	if err != nil {
	//		resp.Diagnostics.AddError("API Client Initliaization Error", err.Error())
	//	}

	//	resp.DataSourceData = client
	//	resp.ResourceData = client

	//	return
	//}

	apiConfig, diags = GetAPIClientConfiguration(ctx, req)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	//// TODO: fix this
	////// Overwrite console URL value in config
	////err := validateConsoleUrl(config.ConsoleURL)
	////if err != nil {
	////    resp.Diagnostics.AddError("API Client Configuration Error", *err)
	////}

	// Set request timeout to 60 if not specified in provider configuration
	if apiConfig.RequestTimeout == nil {
		tflog.Warn(ctx, "No request timeout apiConfigured. Using default value of 60 seconds.")
		defaultTimeout := 60
		apiConfig.RequestTimeout = &defaultTimeout
	}

	// Initialize API client
	tflog.Debug(ctx, "Provider apiConfig created, initializing API client")
	client, err := api.NewCortexCloudAPIClient(ctx, apiConfig)
	if err != nil {
		resp.Diagnostics.AddError("API Client Initalization Error", err.Error())
		return
	}
	tflog.Debug(ctx, "API client initialized successfully")

	// Print warning if debug logs are enabled
	if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" || os.Getenv("TF_LOG_PROVIDER") == "DEBUG" || os.Getenv("TF_LOG_PROVIDER") == "TRACE" {
		tflog.Warn(ctx, "Debug logging enabled. Be aware that your API key and key ID will be visible in the provider log output!")
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func GetAPIClientConfiguration(ctx context.Context, req provider.ConfigureRequest) (api.CortexCloudAPIClientConfig, diag.Diagnostics) {
	// Retrieve configuration values from provider block
	var config api.CortexCloudAPIClientConfig
	diags := req.Config.Get(ctx, &config)
	if diags.HasError() {
		return config, diags
	}
	//// Attempt to retrieve configuration values from environment variables and
	//// overwrite the provider block values if they are succesfully retrieved
	//// and validated
	//diags.Append(overwriteApiClientConfigWithEnvVars(ctx, &config)...)
	//if diags.HasError() {
	//	return config, diags
	//}

	//// Raise diagnostic errors for any unconfigured required values
	//unconfiguredValues := getUnconfiguredValues(config)
	//if len(unconfiguredValues) > 0 {
	//	diags.AddError(
	//		"Provider Configuration Error",
	//		fmt.Sprintf("Required provider configuration values not found in provider block, config file or environment variables: %s", strings.Join(unconfiguredValues, ", ")),
	//	)
	//}

	return config, diags
}

//func overwriteApiClientConfigWithFile(ctx context.Context, config *api.PrismaCloudComputeAPIClientConfig) diag.Diagnostics {
//	var diags diag.Diagnostics
//
//	if config == nil {
//		diags.AddError(
//			"Provider Configuration Error",
//			"Error configuring provider: Expected *api.PrismaCloudComputeAPIClientConfig, got nil pointer. Please report this issue to the provider developers.",
//		)
//		return diags
//	}
//
//	if config.ConfigFile != nil {
//		tflog.Debug(ctx, fmt.Sprintf("Configuring provider from file %s", *config.ConfigFile))
//
//		// Open config file
//		configFile, err := os.Open(*config.ConfigFile)
//		if err != nil {
//			diags.AddWarning(
//				"Provider Configuration File Error",
//				fmt.Sprintf("Error configuring provider: Configuration file specified but could not be opened. Provider will default to using configuration values in provider block or environment variables.\nError: %s", err),
//			)
//			return diags
//		}
//
//		defer configFile.Close()
//
//		// Read contents of config file
//		configFileContent, err := io.ReadAll(configFile)
//		if err != nil {
//			diags.AddWarning(
//				"Provider Configuration File Error",
//				fmt.Sprintf("Error configuring provider: Failed to read configuration file. Provider will default to using configuration values in provider block or environment variables.\nError: %s", err),
//			)
//			return diags
//		}
//
//		// Unmarshal config file contents
//		err = json.Unmarshal(configFileContent, &config)
//		if err != nil {
//			diags.AddWarning(
//				"Provider Configuration File Error",
//				fmt.Sprintf("Error configuring provider: Failed to unmarshal configuration file. Provider will default to using configuration values in provider block or environment variables.\nError: %s", err),
//			)
//		}
//	}
//
//	return diags
//}
//
//// Overwrite API client configuration values with values from environment variables if they're set, non-empty and valid
//func overwriteApiClientConfigWithEnvVars(ctx context.Context, config *api.PrismaCloudComputeAPIClientConfig) diag.Diagnostics {
//	var (
//		diags          diag.Diagnostics
//		consoleUrl     string
//		username       string
//		password       string
//		insecure       bool
//		requestTimeout int
//		err            error
//	)
//
//	// For each nil/empty provider configuration parameter, check its
//	// respective environment variable for a non-empty/valid value. If a valid
//	// value is found, overwrite the relevent configuration value. Otherwise,
//	// raise an error, as this is the final place where this value can be
//	// retrieved from.
//	if util.IsNilOrEmpty(config.ConsoleURL) {
//		err = util.GetEnvironmentVariable(ConsoleUrlEnvVar, &consoleUrl)
//		if err != nil {
//			diags.AddError(
//				"Provider Configuration Error",
//				fmt.Sprintf("Error occured while attempting to parse provider configuration value \"console_url\" from environment variable \"%s\": %s", ConsoleUrlEnvVar, err.Error()),
//			)
//		} else {
//			util.HCLogInfo(ctx, "Using environment variable for provider configuration value \"console_url\"")
//		}
//	}
//
//	if util.IsNilOrEmpty(config.Username) {
//		err = util.GetEnvironmentVariable(UsernameEnvVar, &username)
//		if err != nil {
//			diags.AddError(
//				"Provider Configuration Error",
//				fmt.Sprintf("Error occured while attempting to parse provider configuration value \"username\" from environment variable \"%s\": %s", UsernameEnvVar, err.Error()),
//			)
//		} else {
//			util.HCLogInfo(ctx, "Using environment variable for provider configuration value \"password\"")
//		}
//	}
//
//	if util.IsNilOrEmpty(config.Password) {
//		err = util.GetEnvironmentVariable(PasswordEnvVar, &password)
//		if err != nil {
//			diags.AddError(
//				"Provider Configuration Error",
//				fmt.Sprintf("Error occured while attempting to parse provider configuration value \"password\" from environment variable \"%s\": %s", PasswordEnvVar, err.Error()),
//			)
//		} else {
//			util.HCLogInfo(ctx, "Using environment variable for provider configuration value \"password\"")
//		}
//	}
//
//	if config.Insecure == nil {
//		err = util.GetEnvironmentVariable(InsecureEnvVar, &insecure)
//		if err != nil {
//			diags.AddWarning(
//				"Provider Configuration Error",
//				fmt.Sprintf("Error occured while attempting to parse provider configuration value \"insecure\" from environment variable \"%s\": %s", InsecureEnvVar, err.Error()),
//			)
//		} else {
//			util.HCLogInfo(ctx, "Using environment variable for provider configuration value \"insecure\"")
//		}
//	}
//
//	if config.RequestTimeout == nil {
//		err = util.GetEnvironmentVariable(RequestTimeoutEnvVar, &requestTimeout)
//		if err != nil {
//			diags.AddWarning(
//				"Provider Configuration Error",
//				fmt.Sprintf("Error occured while attempting to parse provider configuration value \"request_timeout\" from environment variable \"%s\": %s", RequestTimeoutEnvVar, err.Error()),
//			)
//		} else {
//			util.HCLogInfo(ctx, "Using environment variable for provider configuration value \"request_timeout\"")
//		}
//	}
//
//	return diags
//}
//
//func getUnconfiguredValues(config api.PrismaCloudComputeAPIClientConfig) []string {
//	unconfiguredValues := []string{}
//
//	if util.IsNilOrEmpty(config.ConsoleURL) {
//		unconfiguredValues = append(unconfiguredValues, "console_url")
//	}
//
//	if util.IsNilOrEmpty(config.Username) {
//		unconfiguredValues = append(unconfiguredValues, "username")
//	}
//
//	if util.IsNilOrEmpty(config.Password) {
//		unconfiguredValues = append(unconfiguredValues, "password")
//	}
//
//	return unconfiguredValues
//}

//func validateConsoleUrl(consoleUrl *string) *string {
//    if consoleUrl == nil {
//        errorMessage := "Error occured while attempting to parse console URL: nil pointer reference"
//        return &errorMessage
//    }
//
//    parsedConsoleUrl, err := url.Parse(*consoleUrl)
//    if err != nil {
//        errorMessage := fmt.Sprintf("Error occured while attempting to parse console URL: %s", err.Error())
//        return &errorMessage
//    }
//
//    // Set URL scheme to https if not specified
//    if parsedConsoleUrl.Scheme == "" {
//        parsedConsoleUrl.Scheme = "https"
//    }
//
//    resp := parsedConsoleUrl.String()
//    consoleUrl = &resp
//
//    return nil
//}
