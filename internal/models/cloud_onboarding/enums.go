package models

import (
    //"fmt"
)

const (
    //BaseEndpoint = "public_api"
)

var (
    CloudIntegrationCloudProviderEnums = []string{
        "AWS",
        "AZURE",
        "GCP",
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
)
