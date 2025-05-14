// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloud_integration

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api"
	cloudIntegrationAPI "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/api/cloud_onboarding/cloud_integration"
	models "github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/models/cloud_onboarding"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/planmodifiers"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/util"
	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/validators"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"

	//"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"organization_id": schema.StringAttribute{
						// TODO: validation
						Description: "TODO",
						Optional:    true,
						Computed:    true,
					},
				},
				Default: objectdefault.StaticValue(types.ObjectNull(map[string]attr.Type{"organization_id": types.StringType})),
			},
			"additional_capabilities": schema.SingleNestedAttribute{
				Description: "Define which additional security capabilities " +
					"to enable.",
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"data_security_posture_management": schema.BoolAttribute{
						Description: "Whether to enable data security " +
							"posture management, an agentless data security " +
							"scanner that discovers, classifies, protects, " +
							"and governs sensitive data.",
						Optional: true,
						Computed: true,
					},
					"registry_scanning": schema.BoolAttribute{
						Description: "Whether to enable registry scanning, " +
							"a container registry scanner that scans " +
							"registry images for vulnerabilities, malware, " +
							"and secrets.",
						Optional: true,
						Computed: true,
					},
					"registry_scanning_options": schema.SingleNestedAttribute{
						Description: "TODO",
						Optional:    true,
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description: "Type of registry scanning. " +
									"Must be one of `ALL`, `LATEST_TAG` or " +
									"`TAGS_MODIFIED_DAYS`. If set to " +
									"`TAGS_MODIFIED_DAYS`, `last_days` must " +
									"be configured.",
								Optional: true,
								Computed: true,
								Validators: []validator.String{
									stringvalidator.OneOf(
										models.CloudIntegrationRegistryScanningTypeEnums...,
									),
									validators.AlsoRequiresOnStringValues(
										[]string{
											models.CloudIntegrationRegistryScanningTypeEnumTagsModifiedDays,
										},
										path.MatchRelative().AtParent().AtName("last_days"),
									),
								},
							},
							"last_days": schema.Int32Attribute{
								Description: "Number of days within which " +
									"the tags on a registry image must have " +
									"been created or updated for the image " +
									"to be scanned. Minimum value is 0 and " +
									"maximum value is 90. Cannot be " +
									"configured if `type` is not set to " +
									"`TAGS_MODIFIED_DAYS`.",
								Optional: true,
								Computed: true,
								Validators: []validator.Int32{
									int32validator.Between(0, 90),
									int32validator.AlsoRequires(path.MatchRelative().AtParent().AtName("type")),
								},
								PlanModifiers: []planmodifier.Int32{
									planmodifiers.NullIfAlsoSetInt32(
										[]string{
											models.CloudIntegrationRegistryScanningTypeEnumAll,
											models.CloudIntegrationRegistryScanningTypeEnumLatestTag,
										},
									),
								},
								//Default: int32default.StaticInt32(90),
							},
						},
					},
					"serverless_scanning": schema.BoolAttribute{
						Description: "TODO",
						Optional:    true,
						Computed:    true,
					},
					"xsiam_analytics": schema.BoolAttribute{
						Description: "Whether to enable XSIAM analytics to " +
							"analyze your endpoint data to develop a " +
							"baseline and raise Analytics and Analytics " +
							"BIOC alerts when anomalies and malicious " +
							"behaviors are detected.",
						Optional: true,
						Computed: true,
					},
				},
			},
			"cloud_provider": schema.StringAttribute{
				Description: "The cloud service provider that is being " +
					"integrated. Must be one of `AWS`, `AZURE` or `GCP`.",
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						models.CloudIntegrationCloudProviderEnums...,
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"collection_configuration": schema.SingleNestedAttribute{
				Description: "Configure the data that will be collected.",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"audit_logs": schema.SingleNestedAttribute{
						Description: "Configuration for audit logs " +
							"collection.",
						Optional: true,
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Whether to enable audit log " +
									"collection.",
								Optional: true,
								Computed: true,
							},
						},
					},
				},
			},
			"custom_resource_tags": schema.SetNestedAttribute{
				// TODO: prevent duplicate tag keys
				Description: "Custom tags that will be applied to any new " +
					"resource created by Cortex in the cloud environment. " +
					"By default, the `managed_by` tag will always be " +
					"applied with the value `paloaltonetworks`.",
				Optional: true,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description: "The key of the custom resource tag.",
							Optional:    true,
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "The value of the custom resource tag.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"instance_name": schema.StringAttribute{
				// TODO: validation
				Description: "Name of the integration instance. If left " +
					"empty, the name will be auto-populated.",
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					validators.ValidateCloudIntegrationInstanceName(),
				},
				Default: stringdefault.StaticString(""),
			},
			"scan_mode": schema.StringAttribute{
				// TODO: add description of outpost
				// TODO: verify that an outpost id is specified if set to MANAGED
				// TODO: find out how to get the default outpost id
				Description: "Define where the scanning for the cloud " +
					"environment will occur. Must be either `MANAGED` or " +
					"`OUTPOST`. If set to `MANAGED`, scanning will be done " +
					"in the Cortex Cloud tenant's environment. If set to " +
					"`OUTPOST`, scanning will be done on the cloud " +
					"infrastructure owned and managed by you." +
					"\n\nNOTE: Scanning with an outpost may require " +
					"additional CSP permissions and may incur additional " +
					"costs.",
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						models.CloudIntegrationScanModeEnums...,
					),
				},
			},
			"scope": schema.StringAttribute{
				Description: "Define the scope for this integration " +
					"instance. Must be one of `ACCOUNT`, `ORGANIZATION` or " +
					"`ACCOUNT_GROUP`.",
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						models.CloudIntegrationScopeEnums...,
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"scope_modifications": schema.SingleNestedAttribute{
				Description: "Define the scope of scans by including/excluding " +
					"accounts or regions.",
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					// TODO: projects, subscriptions (not currently in UI)
					"accounts": schema.SingleNestedAttribute{
						Description: "Configuration for account-level scope " +
							"modifications for AWS integrations. Cannot be " +
							"configured if `cloud_type` is not set to `AWS`.",
						Optional: true,
						Computed: true,
						Attributes: map[string]schema.Attribute{
							// TODO: project_ids, subscription_ids (not currently in UI)
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
							//"account_ids": schema.SetAttribute{
							//    Description: "TODO",
							//    Optional: true,
							//    Computed: true,
							//    ElementType: types.StringType,
							//},
						},
						Default: objectdefault.StaticValue(
							types.ObjectNull(
								map[string]attr.Type{
									"enabled": types.BoolType,
								},
							),
						),
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
						Default: objectdefault.StaticValue(
							types.ObjectNull(
								map[string]attr.Type{
									"enabled": types.BoolType,
								},
							),
						),
					},
				},
			},
			"status": schema.StringAttribute{
				Description: "Status of the integration.",
				Computed:    true,
			},
			"instance_id": schema.StringAttribute{
				Description: "A unique identifier of the integration.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"account_name": schema.StringAttribute{
				Description: "TODO",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"outpost_id": schema.StringAttribute{
				Description: "A unique identifier of the Outpost instance " +
					"assigned to the integration.",
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"creation_time": schema.StringAttribute{
				Description: "Timestamp representing the creation date and " +
					"time of the cloud integration template.",
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cloud_formation_link": schema.StringAttribute{
				Description: "Link to the CloudFormation stack generated by " +
					"Cortex to deploy the necessary infrastructure for " +
					"integrating the AWS account, OU or organization into " +
					"Cortex. Only populated for AWS integegrations. Assign " +
					"this value to the `template_url` argument in a " +
					"`aws_cloudformation_stack` resource from the official " +
					"AWS Terraform Provider to automatically set up the " +
					"integration.",
				Computed:  true,
				Sensitive: true,
				// TODO: this might actually change if the edit endpoint returns a new CF template
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// TEMPORARY
			"template_instance_id": schema.StringAttribute{
				Description: "TEMPORARY: A unique identifier of the integration template.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// END TEMPORARY
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
	getRequest := plan.ToGetRequest(ctx, &resp.Diagnostics, &instanceId)
	integrationDetails := cloudIntegrationAPI.Get(ctx, &resp.Diagnostics, r.client, getRequest)
	if resp.Diagnostics.HasError() {
		return
	}

	// Populate API response values into model
	plan.RefreshPropertyValues(&resp.Diagnostics, integrationDetails, &instanceId, &templateUrl, true)
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
	instanceId := state.InstanceId.ValueString()
	getRequest := state.ToGetRequest(ctx, &resp.Diagnostics, &instanceId)
	//integrationDetails := cloudIntegrationAPI.Get(ctx, &resp.Diagnostics, r.client, getRequest)
	integrationDetails := cloudIntegrationAPI.Get(ctx, &resp.Diagnostics, r.client, getRequest)
	if resp.Diagnostics.HasError() {
		return
	}

	//integrationDetails := cloudIntegrationAPI.GetByInstanceId(ctx, &resp.Diagnostics, r.client, state.InstanceId.ValueString())
	//if resp.Diagnostics.HasError() {
	//	return
	//}

	// Refresh state values
	state.RefreshPropertyValues(&resp.Diagnostics, integrationDetails, nil, nil, false)
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

	// Get current state
	var state models.CloudIntegrationInstanceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from state
	request := state.ToDeleteRequest(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete integration
	cloudIntegrationAPI.Delete(ctx, &resp.Diagnostics, r.client, request)
	if resp.Diagnostics.HasError() {
		return
	}

	// Retrieve integration details from API and print warning if the
	// integration still appears in the results
	integrationDetails := cloudIntegrationAPI.GetByInstanceId(ctx, &resp.Diagnostics, r.client, state.InstanceName.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}

	if len(integrationDetails.Reply.Data) > 0 {
		resp.Diagnostics.AddWarning(
			"Resource Not Deleted",
			"Terraform resource has been destroyed, but integration still "+
				"exists in Cortex. This may occur when deleting integrations "+
				"that are in the PENDING state. Navigate to the Data Sources "+
				"menu in the Cortex UI and search for your integration to "+
				"confirm that it has been deleted, or manually delete it if "+
				"not.",
		)
	}
}
