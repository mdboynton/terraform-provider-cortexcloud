package resources

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	//"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	//"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Ensure the implementation satisfies the expected interfaces.
var (
    _ resource.Resource = &CloudOnboardingIntegrationTemplateResource{}
)

// NewCloudOnboardingIntegrationTemplateResource is a helper function to simplify the provider implementation.
func NewCloudOnboardingIntegrationTemplateResource() resource.Resource {
    return &CloudOnboardingIntegrationTemplateResource{}
}

// CloudOnboardingIntegrationTemplateResource is the resource implementation.
type CloudOnboardingIntegrationTemplateResource struct {
    client *api.CortexCloudAPIClient
}

// Metadata returns the resource type name.
func (r *CloudOnboardingIntegrationTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_cloud_onboarding_integration_template"
}

// Schema defines the schema for the resource.
func (r *CloudOnboardingIntegrationTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Description: "TODO",
        Attributes: map[string]schema.Attribute{
            "additional_capabilities": schema.SingleNestedAttribute{
                Description: "TODO",
                //Required: true,
                Optional: true,
                Computed: true,
                Attributes: map[string]schema.Attribute{
                    "data_security_posture_management": schema.BoolAttribute{
                        Description: "TODO",
                        Optional: true,
                        Computed: true,
                    },
                    "registry_scanning": schema.BoolAttribute{
                        Description: "TODO",
                        Optional: true,
                        Computed: true,
                    },
                    "registry_scanning_options": schema.SingleNestedAttribute{
                        Description: "TODO",
                        Optional: true,
                        Computed: true,
                        Attributes: map[string]schema.Attribute{
                            "type": schema.StringAttribute{
                                // TODO: validation ("ALL", etc)
                                Description: "TODO",
                                Optional: true,
                                Computed: true,
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
                        Optional: true,
                        Computed: true,
                    },
                    "xsiam_analytics": schema.BoolAttribute{
                        Description: "TODO",
                        Optional: true,
                        Computed: true,
                    },
                },
            },
            "cloud_provider": schema.StringAttribute{
                // TODO: validation
                Description: "TODO",
                Required: true,
            },
            "collection_configuration": schema.SingleNestedAttribute{
                Description: "TODO",
                Optional: true,
                Computed: true,
                Attributes: map[string]schema.Attribute{
                    "audit_logs": schema.SingleNestedAttribute{
                        Description: "TODO",
                        Optional: true,
                        Computed: true,
                        Attributes: map[string]schema.Attribute{
                            "enabled": schema.BoolAttribute{
                                Description: "TODO",
                                Optional: true,
                                Computed: true,
                            },
                        },
                    },
                },
                //Attributes: map[string]schema.Attribute{
                //    "audit_logs": schema.BoolAttribute{
                //        Description: "TODO",
                //        Optional: true,
                //        Computed: true,
                //    },
                //},
            },
            "custom_resource_tags": schema.SetNestedAttribute{
                Description: "TODO",
                Optional: true,
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "key": schema.StringAttribute{
                            Description: "TODO",
                            Optional: true,
                            Computed: true,
                        },
                        "value": schema.StringAttribute{
                            Description: "TODO",
                            Optional: true,
                            Computed: true,
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
                Optional: true,
                Computed: true,
                //PlanModifiers: []planmodifier.String{
                //    stringplanmodifier.RequiresReplace(),
                //},
            },
            "scan_mode": schema.StringAttribute{
                // TODO: validation ("MANAGED", "OUTPOST")
                // TODO: include warning about additional costs when using outpost
                Description: "TODO",
                Required: true,
            },
            "scope": schema.StringAttribute{
                // TODO: validation ("ACCOUNT", "ORGANIZATION", "ACCOUNT_GROUP")
                Description: "TODO",
                Required: true,
            },
            "scope_modifications": schema.SingleNestedAttribute{
                Description: "TODO",
                Optional: true,
                Computed: true,
                Attributes: map[string]schema.Attribute{
                    // TODO: projects, subscriptions (not currently in UI)
                    //"accounts": schema.SingleNestedAttribute{
                    //    Description: "TODO",
                    //    Optional: true,
                    //    Computed: true,
                    //    Attributes: map[string]schema.Attribute{
                    //        // TODO: do we need an enabled attribute or is it
                    //        // not needed since it's optional?
                    //        // TODO: project_ids, subscription_ids (not currently in UI)
                    //        "enabled": schema.BoolAttribute{
                    //            Description: "TODO",
                    //            Optional: true,
                    //            Computed: true,
                    //        },
                    //        "type": schema.StringAttribute{
                    //            // TODO: validation ("INCLUDE", "EXCLUDE")
                    //            Description: "TODO",
                    //            Optional: true,
                    //            Computed: true,
                    //        },
                    //        "account_ids": schema.SetAttribute{
                    //            Description: "TODO",
                    //            Optional: true,
                    //            Computed: true,
                    //            ElementType: types.StringType,
                    //        },
                    //    },
                    //},
                    "regions": schema.SingleNestedAttribute{
                        Description: "TODO",
                        Optional: true,
                        Computed: true,
                        Attributes: map[string]schema.Attribute{
                            // TODO: do we need an enabled attribute or is it
                            // not needed since it's optional?
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
                Computed: true,
            },
            "instance_id": schema.StringAttribute{
                Description: "TODO",
                Computed: true,
            },
            "account_name": schema.StringAttribute{
                Description: "TODO",
                Computed: true,
            },
            "outpost_id": schema.StringAttribute{
                Description: "TODO",
                Computed: true,
            },
            "creation_time": schema.StringAttribute{
                Description: "TODO",
                Computed: true,
            },
            "cloud_formation_link": schema.StringAttribute{
                Description: "TODO",
                Computed: true,
                Sensitive: true,
            },
        },
    }
}

// Configure adds the provider-configured client to the resource.
func (r *CloudOnboardingIntegrationTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *CloudOnboardingIntegrationTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    defer util.PanicHandler(&resp.Diagnostics)

    // Read Terraform plan data into model
    var data models.CloudOnboardingIntegrationTemplateModel
    resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    // Create cloud onboarding integration template
    response := createCloudOnboardingIntegrationTemplate(ctx, &resp.Diagnostics, r.client, data)
	if resp.Diagnostics.HasError() {
        return
	}

    // Retrieve instance ID and template URL from API response
    instanceId := response.Reply.Automated.TrackingGuid
    templateUrl := parseTemplateURL(&resp.Diagnostics, response.Reply.Automated.Link)
	if resp.Diagnostics.HasError() {
        return
	}

    // Retrieve cloud integration details from API
    integrationStatus := getCloudIntegrationsByInstanceId(ctx, &resp.Diagnostics, r.client, instanceId)
	if resp.Diagnostics.HasError() {
        return
	}

    // Populate API response values in model
    data.RefreshPropertyValues(&resp.Diagnostics, integrationStatus, &instanceId, &templateUrl)
	if resp.Diagnostics.HasError() {
        return
	}

    // Set state to fully populated data
    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read refreshes the Terraform state with the latest data.
func (r *CloudOnboardingIntegrationTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
    defer util.PanicHandler(&resp.Diagnostics)

    // Get current state
    var state models.CloudOnboardingIntegrationTemplateModel
    resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
    if resp.Diagnostics.HasError() {
        return
    }

    // Retrieve integration status from remote
    integrationStatus := getCloudIntegrationsByInstanceId(ctx, &resp.Diagnostics, r.client, state.InstanceId.ValueString())
    if resp.Diagnostics.HasError() {
        return
    }

    // Refresh state values
    state.RefreshPropertyValues(&resp.Diagnostics, integrationStatus, nil, nil)
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
func (r *CloudOnboardingIntegrationTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
    defer util.PanicHandler(&resp.Diagnostics)

    // Read Terraform plan data into model
    var plan models.CloudOnboardingIntegrationTemplateModel
    resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
    if resp.Diagnostics.HasError() {
        return
    }

    // Get current state
    // TODO: do we need this?
    var state models.CloudOnboardingIntegrationTemplateModel
    resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
    if resp.Diagnostics.HasError() {
        return
    }

    // Update integration
    updatedTemplate := updateCloudIntegration(ctx, &resp.Diagnostics, r.client, plan)
    if resp.Diagnostics.HasError() {
        return
    }
    
    // Set state to updated values 
    resp.Diagnostics.Append(resp.State.Set(ctx, &updatedTemplate)...)
    if resp.Diagnostics.HasError() {
        return
    }
}

// Delete deletes the resource and removes it from the Terraform state on success.
func (r *CloudOnboardingIntegrationTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    defer util.PanicHandler(&resp.Diagnostics)
}
