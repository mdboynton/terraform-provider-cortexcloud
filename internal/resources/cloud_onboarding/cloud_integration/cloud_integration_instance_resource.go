package cloud_integration

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
	cloudIntegrationAPI "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/cloud_onboarding/cloud_integration"
	models "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models/cloud_onboarding"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/validators"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &CloudIntegrationInstanceResource{}
)

// NewCloudIntegrationInstanceResource is a helper function to simplify the provider implementation.
func NewCloudIntegrationInstanceResource() resource.Resource {
	return &CloudIntegrationInstanceResource{}
}

// CloudIntegrationInstanceResource is the resource implementation.
type CloudIntegrationInstanceResource struct {
	client *api.CortexCloudAPIClient
}

// Metadata returns the resource type name.
func (r *CloudIntegrationInstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_integration_instance"
}

// Schema defines the schema for the resource.
func (r *CloudIntegrationInstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "TODO",
		Attributes: map[string]schema.Attribute{
            // TODO: currently can only be specified for Azure integrations
			"account_details": schema.SingleNestedAttribute{
				Description: "TODO",
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
                    "organization_id": schema.StringAttribute{
                        // TODO: validation
				        Description: "TODO",
				        Optional: true,
				        Computed: true,
                    },
                },
                Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{"organization_id": types.StringType})),
                //Default: objectdefault.StaticValue(types.ObjectValueMust(
                //        map[string]attr.Type{
                //            "organization_id": types.StringType,
                //        },
                //        map[string]attr.Value{
                //            "organization_id": types.StringNull(),
                //        },
                //    ),
                //),
            },
			"additional_capabilities": schema.SingleNestedAttribute{
				Description: "TODO",
				//Required: true,
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"data_security_posture_management": schema.BoolAttribute{
						Description: "TODO",
						Optional:    true,
						Computed:    true,
					},
					"registry_scanning": schema.BoolAttribute{
						Description: "TODO",
						Optional:    true,
						Computed:    true,
					},
					"registry_scanning_options": schema.SingleNestedAttribute{
						Description: "TODO",
						Optional:    true,
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								// TODO: validation ("ALL", etc)
								Description: "TODO",
								Optional:    true,
								Computed:    true,
							},
						},
					},
					//"registry_scanning": schema.SingleNestedAttribute{
					//    Description: "TODO",
					//    Optional: true,
					//    Computed: true,
					//    Attributes: map[string]schema.Attribute{
					//        "enabled": schema.BoolAttribute{
					//            Description: "TODO",
					//            Optional: true,
					//            Computed: true,
					//        },
					//        "initial_scanning_configuration": schema.StringAttribute{
					//            // TODO: validation ("ALL", etc)
					//            Description: "TODO",
					//            Optional: true,
					//            Computed: true,
					//        },
					//    },
					//},
					"serverless_scanning": schema.BoolAttribute{
						Description: "TODO",
						Optional:    true,
						Computed:    true,
					},
					"xsiam_analytics": schema.BoolAttribute{
						Description: "TODO",
						Optional:    true,
						Computed:    true,
					},
				},
			},
			"cloud_provider": schema.StringAttribute{
				// TODO: validation
				Description: "TODO",
				Required:    true,
                Validators: []validator.String{
                    validators.StringEnumValidator(models.CloudIntegrationCloudProviderEnums),
                },
			},
			"collection_configuration": schema.SingleNestedAttribute{
				Description: "TODO",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"audit_logs": schema.SingleNestedAttribute{
						Description: "TODO",
						Optional:    true,
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "TODO",
								Optional:    true,
								Computed:    true,
							},
						},
					},
				},
			},
			"custom_resource_tags": schema.SetNestedAttribute{
				Description: "TODO",
				Optional:    true,
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description: "TODO",
							Optional:    true,
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "TODO",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"instance_name": schema.StringAttribute{
				// TODO: validation
				// TODO: include message about auto-population if empty
				// TODO: this might not be able to be populated with auto-generated value
				// since its not returned in the response payload
				Description: "TODO",
				Optional:    true,
				Computed:    true,
				//PlanModifiers: []planmodifier.String{
				//    stringplanmodifier.RequiresReplace(),
				//},
			},
			"scan_mode": schema.StringAttribute{
				// TODO: include warning about additional costs when using outpost
				Description: "TODO",
				Required:    true,
                Validators: []validator.String{
                    validators.StringEnumValidator(models.CloudIntegrationScanModeEnums),
                },
			},
			"scope": schema.StringAttribute{
				Description: "TODO",
				Required:    true,
                Validators: []validator.String{
                    validators.StringEnumValidator(models.CloudIntegrationScopeEnums),
                },
			},
			"scope_modifications": schema.SingleNestedAttribute{
				Description: "TODO",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					// TODO: projects, subscriptions (not currently in UI)
					"accounts": schema.SingleNestedAttribute{
					    Description: "TODO",
					    Optional: true,
					    Computed: true,
					    Attributes: map[string]schema.Attribute{
					        // TODO: project_ids, subscription_ids (not currently in UI)
					        "enabled": schema.BoolAttribute{
					            Description: "TODO",
					            Optional: true,
					            Computed: true,
					        },
					        //"type": schema.StringAttribute{
					        //    // TODO: validation ("INCLUDE", "EXCLUDE")
					        //    Description: "TODO",
					        //    Optional: true,
					        //    Computed: true,
					        //},
					        //"account_ids": schema.SetAttribute{
					        //    Description: "TODO",
					        //    Optional: true,
					        //    Computed: true,
					        //    ElementType: types.StringType,
					        //},
					    },
					},
					"regions": schema.SingleNestedAttribute{
						Description: "TODO",
						Optional:    true,
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							// TODO: do we need an enabled attribute or is it
							// not needed since it's optional?
							"enabled": schema.BoolAttribute{
								Description: "TODO",
								Optional:    true,
								Computed:    true,
							},
							//"type": schema.StringAttribute{
							//    // TODO: validation ("INCLUDE", "EXCLUDE")
							//    Description: "TODO",
							//    Optional: true,
							//    Computed: true,
							//},
							//"regions": schema.SetAttribute{
							//    Description: "TODO",
							//    Optional: true,
							//    Computed: true,
							//    ElementType: types.StringType,
							//},
						},
					},
				},
			},
			"status": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
			},
			"instance_id": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
			},
			"account_name": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
			},
			"outpost_id": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
			},
			"creation_time": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
			},
			"cloud_formation_link": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

// Configure adds the provider-configured client to the resource.
func (r *CloudIntegrationInstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
    // Prevent panic if the provider has not been configured
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
func (r *CloudIntegrationInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	defer util.PanicHandler(&resp.Diagnostics)

	// Read Terraform plan data into model
	var plan models.CloudIntegrationInstanceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
    request := plan.ToCreateRequest(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new cloud onboarding integration template
    response, templateUrl := cloudIntegrationAPI.CreateTemplate(ctx, &resp.Diagnostics, r.client, request)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get instance ID from API response
	instanceId := response.Reply.Automated.TrackingGuid

	// Retrieve cloud integration details from API
    integrationDetails := cloudIntegrationAPI.GetByInstanceId(ctx, &resp.Diagnostics, r.client, instanceId)
	if resp.Diagnostics.HasError() {
		return
	}

	// Populate API response values into model
	plan.RefreshPropertyValues(&resp.Diagnostics, integrationDetails, &instanceId, &templateUrl)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

// Read refreshes the Terraform state with the latest data.
func (r *CloudIntegrationInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	defer util.PanicHandler(&resp.Diagnostics)

	// Get current state
	var state models.CloudIntegrationInstanceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve integration details from API
    integrationDetails := cloudIntegrationAPI.GetByInstanceId(ctx, &resp.Diagnostics, r.client, state.InstanceName.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}

	// Refresh state values
	state.RefreshPropertyValues(&resp.Diagnostics, integrationDetails, nil, nil)
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
func (r *CloudIntegrationInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	defer util.PanicHandler(&resp.Diagnostics)

	// Read Terraform plan data into model
	var plan models.CloudIntegrationInstanceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

    // Generate API request body from plan
    request := plan.ToUpdateRequest(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update integration
	updatedIntegration := cloudIntegrationAPI.Update(ctx, &resp.Diagnostics, r.client, request)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to updated values
	resp.Diagnostics.Append(resp.State.Set(ctx, &updatedIntegration)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes it from the Terraform state on success.
func (r *CloudIntegrationInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	defer util.PanicHandler(&resp.Diagnostics)
}
