// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package util

import (
	"regexp"
	"github.com/go-playground/validator/v10"
)

// ---------------------------
// Regex Patterns
// ---------------------------

const (
	// TODO: update this to also work with whatever endpoint ID is
	RegexpPatternSystemManagementUserOrEndpointID = `^[^/]+/[^/]+$`
)

var (
	RegexpSystemManagementUserOrEndpointID *regexp.Regexp
)

func CompileRegex() error {
	var err error

	RegexpSystemManagementUserOrEndpointID, err = regexp.Compile(RegexpPatternSystemManagementUserOrEndpointID)

	return err
}

func ValidateRiskScoreID(fl validator.FieldLevel) bool {
	return RegexpSystemManagementUserOrEndpointID.MatchString(fl.Field().String())
}
