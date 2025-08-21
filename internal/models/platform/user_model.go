package platform

import (
	"context"

	"github.com/mdboynton/cortex-cloud-go/platform"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// *********************************************************
// Structs
// *********************************************************
type UserModel struct {
	UserEmail     types.String `tfsdk:"user_email"`
	UserFirstName types.String `tfsdk:"user_first_name"`
	UserLastName  types.String `tfsdk:"user_last_name"`
	RoleName      types.String `tfsdk:"role_name"`
	LastLoggedIn  types.Int32  `tfsdk:"last_logged_in"`
	UserType      types.String `tfsdk:"user_type"`
	Groups        types.Set    `tfsdk:"groups"`
	Scope         types.Set    `tfsdk:"scope"`
}

func (m *UserModel) ToGetRequest(ctx context.Context, diagnostics *diag.Diagnostics) platform.GetUserRequest {
	return platform.GetUserRequest{
		Email: m.UserEmail.ValueString(),
	}
}

func (m *UserModel) RefreshPropertyValues(ctx context.Context, diagnostics *diag.Diagnostics, response platform.GetUserResponse) {
	data, err := response.Marshal()
	if err != nil {
		diagnostics.AddError(
			"Value Conversion Error",
			err.Error(),
		)
	}

	groups, diags := types.SetValueFrom(ctx, m.Groups.ElementType(ctx), data.Groups)
	diagnostics.Append(diags...)
	if diagnostics.HasError() {
		return
	}

	scope, diags := types.SetValueFrom(ctx, m.Scope.ElementType(ctx), data.Scope)
	diagnostics.Append(diags...)
	if diagnostics.HasError() {
		return
	}

	m.UserEmail = types.StringValue(data.UserEmail)
	m.UserFirstName = types.StringValue(data.UserFirstName)
	m.UserLastName = types.StringValue(data.UserLastName)
	m.RoleName = types.StringValue(data.RoleName)
	m.LastLoggedIn = types.Int32Value(data.LastLoggedIn)
	m.UserType = types.StringValue(data.UserType)
	m.Groups = groups
	m.Scope = scope
}
