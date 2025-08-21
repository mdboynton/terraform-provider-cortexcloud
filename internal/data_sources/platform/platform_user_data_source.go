package platform

import (
	"context"

	models "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models/platform"
	providerModels "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models/provider"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/util"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mdboynton/cortex-cloud-go/platform"
)

var (
	_ datasource.DataSource = &UserDataSource{}
)

func NewPlatformUserDataSource() datasource.DataSource {
	return &UserDataSource{}
}

type UserDataSource struct {
	client *platform.Client
}

func (r *UserDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "TODO",
		Attributes: map[string]schema.Attribute{
			"user_email": schema.StringAttribute{
				Description: "Email address of the user.",
				Required:    true,
			},
			"user_first_name": schema.StringAttribute{
				Description: "First name of the user.",
				Required:    true,
			},
			"user_last_name": schema.StringAttribute{
				Description: "Last name of the user.",
				Required:    true,
			},
			"role_name": schema.StringAttribute{
				Description: "Role name associated with the user.",
				Required:    true,
			},
			"last_logged_in": schema.Int32Attribute{
				Description: "Timestamp of when the user last logged in.",
				Required:    true,
			},
			"user_type": schema.StringAttribute{
				Description: "Type of user.",
				Required:    true,
			},
			"groups": schema.SetAttribute{
				Description: "Name of user groups associated with the user, if applicable.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"scope": schema.SetAttribute{
				Description: "Name of scope associated with the user, if applicable.",
				ElementType: types.StringType,
				Optional:    true,
			},
		},
	}
}

// Configure adds the provider-configured client to the data store.
func (r *UserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*providerModels.CortexCloudSDKClients)

	if !ok {
		util.AddUnexpectedResourceConfigureTypeError(&resp.Diagnostics, "*http.Client", req.ProviderData)
		return
	}

	r.client = client.Platform
}

// Read refreshes the Terraform state with the latest data.
func (r *UserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	defer util.PanicHandler(&resp.Diagnostics)

	// Populate data source configuration into model
	var config models.UserModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve user details from API
	request := config.ToGetRequest(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := r.client.GetUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Platform Data Source Read Error",
			err.Error(),
		)
		return
	}

	// Refresh state values
	config.RefreshPropertyValues(ctx, &resp.Diagnostics, response)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set refreshed state
	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
