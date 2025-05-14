// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

//"fmt"

const (
// BaseEndpoint = "public_api"
)

var (
	CloudIntegrationCloudProviderEnumAws   = "AWS"
	CloudIntegrationCloudProviderEnumAzure = "AZURE"
	CloudIntegrationCloudProviderEnumGcp   = "GCP"
	CloudIntegrationCloudProviderEnums     = []string{
		CloudIntegrationCloudProviderEnumAws,
		CloudIntegrationCloudProviderEnumAzure,
		CloudIntegrationCloudProviderEnumGcp,
	}

	CloudIntegrationScanModeEnums = []string{
		"MANAGED",
		"OUTPOST",
	}

	CloudIntegrationScopeEnums = []string{
		"ACCOUNT",
		"ORGANIZATION",
		"ACCOUNT_GROUP",
	}

	CloudIntegrationScopeModificationTypeEnums = []string{
		"INCLUDE",
		"EXCLUDE",
	}

	CloudIntegrationRegistryScanningTypeEnumAll              = "ALL"
	CloudIntegrationRegistryScanningTypeEnumLatestTag        = "LATEST_TAG"
	CloudIntegrationRegistryScanningTypeEnumTagsModifiedDays = "TAGS_MODIFIED_DAYS"
	CloudIntegrationRegistryScanningTypeEnums                = []string{
		CloudIntegrationRegistryScanningTypeEnumAll,
		CloudIntegrationRegistryScanningTypeEnumLatestTag,
		CloudIntegrationRegistryScanningTypeEnumTagsModifiedDays,
	}
)
