package validators

import (
    "context"
    "fmt"
    "slices"
    "strings"
	
    "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func StringEnumValidator(enums []string) stringEnumValidator {
    return stringEnumValidator{
        Enums: enums,
    }
}

type stringEnumValidator struct {
    Enums []string
}

func (v stringEnumValidator) Description(ctx context.Context) string {
    return ""
}

func (v stringEnumValidator) MarkdownDescription(ctx context.Context) string {
    return ""
}

func (v stringEnumValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
    if (req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown()) {
        return
    }

    if !slices.Contains(v.Enums, req.ConfigValue.String()) {
        resp.Diagnostics.AddAttributeError(
            req.Path,
            "Validation Error", 
            fmt.Sprintf(enumValidationErrorMessage, strings.Join(v.Enums[:], ", ")),
        )
    }

    return
}
