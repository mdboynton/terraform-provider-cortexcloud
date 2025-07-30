package models

import (
	sdk "github.com/mdboynton/cortex-cloud-go/api"
	"github.com/mdboynton/cortex-cloud-go/appsec"
	"github.com/mdboynton/cortex-cloud-go/cloudonboarding"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CortexCloudProviderModel struct {
	ApiUrl               types.String `tfsdk:"api_url"`
	ApiPort              types.Int32  `tfsdk:"api_port"`
	ApiKey               types.String `tfsdk:"api_key"`
	ApiKeyId             types.Int32  `tfsdk:"api_key_id"`
	Insecure             types.Bool   `tfsdk:"insecure"`
	RequestTimeout       types.Int32  `tfsdk:"request_timeout"`
	RequestRetryInterval types.Int32  `tfsdk:"request_retry_interval"`
	CrashStackDir        types.String `tfsdk:"crash_stack_dir"`
	ConfigFile           types.String `tfsdk:"config_file"`
	CheckEnvironment     types.Bool   `tfsdk:"check_environment"`
}

type CortexCloudSDKClients struct {
	Config          sdk.Config
	AppSec          *appsec.Client
	CloudOnboarding *cloudonboarding.Client
}
