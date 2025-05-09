package application_security

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
	appSecAPI "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/application_security"
	models "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models/application_security"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &ApplicationSecurityRuleResource{}
)

// NewApplicationSecurityRuleResource is a helper function to simplify the provider implementation.
func NewApplicationSecurityRuleResource() resource.Resource {
	return &ApplicationSecurityRuleResource{}
}

// ApplicationSecurityRuleResource is the resource implementation.
type ApplicationSecurityRuleResource struct {
	client *api.CortexCloudAPIClient
}

// Metadata returns the resource type name.
func (r *ApplicationSecurityRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_security_rule"
}

// Schema defines the schema for the resource.
func (r *ApplicationSecurityRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "TODO",
		Attributes: map[string]schema.Attribute{
			"category": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Required:    true,
                //Computed:    true,
			},
			"cloud_provider": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				//Required:    true,
				Optional:    true,
                Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "TODO",
                Computed:    true,
			},
			"description": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			"detection_method": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			"doc_link": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			"domain": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			"finding_category": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			"finding_docs": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			"finding_type_id": schema.Int32Attribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			"finding_type_name": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
            "frameworks": schema.SetNestedAttribute{
				Description: "TODO",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "TODO",
							Required:    true,
                            //Computed:    true,
						},
						"definition": schema.StringAttribute{
                            // TODO: validate yaml
							Description: "TODO",
							Required:    true,
                            //Computed:    true,
						},
						"definition_link": schema.StringAttribute{
							Description: "TODO",
							Optional:    true,
                            Computed:    true,
                            Default:     stringdefault.StaticString(""),
						},
						"remediation_description": schema.StringAttribute{
							Description: "TODO",
							Optional:    true,
                            Computed:    true,
                            Default:     stringdefault.StaticString(""),
						},
					},
				},
            },
			"id": schema.StringAttribute{
				// TODO: validation
                // TODO: should this be modifiable? can you change it via API?
				Description: "TODO",
				//Optional:    true,
                Computed:    true,
			},
			"is_custom": schema.BoolAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			"is_enabled": schema.BoolAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
            "labels": schema.SetAttribute{
				Description: "TODO",
				Optional:    true,
                ElementType: types.StringType,
            },
            "mitre_tactics": schema.SetAttribute{
				Description: "TODO", 
                Optional:    true, 
                Computed:    true,
                ElementType: types.StringType,
            },
            "mitre_techniques": schema.SetAttribute{
				Description: "TODO",
				Optional:    true,
                Computed:    true,
                ElementType: types.StringType,
            },
			"name": schema.StringAttribute{
				// TODO: validation
                // TODO: should this be modifiable? does it require replace?
				Description: "TODO",
				Required:    true,
			},
			"owner": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				//Optional:    true,
                Computed:    true,
			},
			"scanner": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Required:    true,
                //Computed:    true,
			},
			//"scanner_rule_id": schema.StringAttribute{
			//	// TODO: validation
			//	Description: "TODO",
			//	Optional:    true,
            //    Computed:    true,
			//},
			"severity": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Required:    true,
                //Computed:    true,
			},
			"source": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			//"source_version": schema.StringAttribute{
			//	// TODO: validation
			//	Description: "TODO",
			//	Optional:    true,
            //    Computed:    true,
			//},
			"sub_category": schema.StringAttribute{
				// TODO: validation
                // The valid inputs for this attribute are determined by the "category" value
				Description: "TODO",
				Optional:    true,
                Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "TODO",
                Computed:    true,
			},
        },
	}
}

// Configure adds the provider-configured client to the resource.
func (r *ApplicationSecurityRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.CortexCloudAPIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create creates the resource and sets the initial Terraform state.
func (r *ApplicationSecurityRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	defer util.PanicHandler(&resp.Diagnostics)
    

	// Read Terraform plan data into model
	var plan models.ApplicationSecurityRuleModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
    request := plan.ToCreateRequest(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

    // Create new application security rule
	response := appSecAPI.Create(ctx, &resp.Diagnostics, r.client, request)
	if resp.Diagnostics.HasError() {
		return
	}

	// Populate API response values in model
	plan.RefreshPropertyValues(ctx, &resp.Diagnostics, response)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest data.
func (r *ApplicationSecurityRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	defer util.PanicHandler(&resp.Diagnostics)

    // Get current state
	var state models.ApplicationSecurityRuleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

    // Retrieve rule details from API
    rule := appSecAPI.Get(ctx, &resp.Diagnostics, r.client, state.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}

	// Refresh state values
	state.RefreshPropertyValues(ctx, &resp.Diagnostics, rule)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set refreshed state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *ApplicationSecurityRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	defer util.PanicHandler(&resp.Diagnostics)
}

// Delete deletes the resource and removes it from the Terraform state on success.
func (r *ApplicationSecurityRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	defer util.PanicHandler(&resp.Diagnostics)
}
