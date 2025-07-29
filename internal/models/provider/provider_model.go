package models

import (
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
