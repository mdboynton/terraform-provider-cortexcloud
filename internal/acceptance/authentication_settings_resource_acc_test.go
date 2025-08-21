// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	name              string = "test-tf-provider-auth-settings"
	domain            string = "test.com"
	defaultRole       string = "Instance Administrator"
	isAccountRole     bool   = false
	mappingsEmail            = "email"
	mappingsFirstName        = "firstName"
	mappingsLastName         = "lastName"
	mappingsGroupName        = "group"
	idpSsoUrl                = "https://paloaltonetworks.com/app/signin"
	idpCertificate           = "MIIDuzCCAqOgAwIBAgIUQs1LRebZYRe1emleU6a8mBHxRJwwDQYJKoZIhvcNAQELBQAwbTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkFSMRQwEgYDVQQHDAtMaXR0bGUgUm9jazEbMBkGA1UECgwSUGFsbyBBbHRvIE5ldHdvcmtzMR4wHAYDVQQLDBVQcm9mZXNzaW9uYWwgU2VydmljZXMwHhcNMjUwODIxMTYxNzI4WhcNMjYwODIxMTYxNzI4WjBtMQswCQYDVQQGEwJVUzELMAkGA1UECAwCQVIxFDASBgNVBAcMC0xpdHRsZSBSb2NrMRswGQYDVQQKDBJQYWxvIEFsdG8gTmV0d29ya3MxHjAcBgNVBAsMFVByb2Zlc3Npb25hbCBTZXJ2aWNlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKvLWHO1gwgMv5thWY3XR5+NCIuYZFxFHi1mva8w1e0b0A1nwwqO+eowzzxsEBcSIKf7rBkezDYrRJqnEuCqDKOI/jV/HAKU1h/ZJW3qgRFogO7eEmhPvvSOWXkExJmJ8ic7jS48pAbG9+dg9fZtAN6waMeB93mHAS0aY4sPuuyCbl8uyc0ovXJ2nqHTdB1Ff4W4wLtKJJsoK9N8E+Pz0YAI5dp2Ir2fgoERKDU9JjN5dwMGraQNa3LJCMpQj+1vrkBrL0bLfKI8daRk6MicYDTVnuBo/YDBRu+aLXftBpw7hyvmMivgwktDJDziBhmPoRv29A1bTOsxVEUt6w59QP8CAwEAAaNTMFEwHQYDVR0OBBYEFK3rXPMkuBe+hDNW3B4eYbnawCMAMB8GA1UdIwQYMBaAFK3rXPMkuBe+hDNW3B4eYbnawCMAMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAAJaB3Q3DvEu3LwPHFNCRnHwN8fAcGP1ItPB1G7f/EqBhzIOU0ZaxvqLnECPesy8LQqoHX8vdcAv6n1U+zPZp2zCSVVPd+sbZECqBhs9VsI2WF6LlkLvKDti0PESgkj1+K54xFDMiK56YmIUBP0rpOxL+MT+fzJCiG4H86uvc/jtMGVS96mWbl32Cf9MFhONhwDaxXDVhycWv197unvfIjrpMCOOxQCWGsU2MHGgidrtApHU9sufROCbjrm6wnuV0ndbHWUkK+oWmExIY/5h4qv7kQgme9tRkoKD3zyV8oqPt9NRzHTBR/5QDoaC/tvwb52Ohp+zKH/iD3CNrLPY5LQ="
	idpIssuer                = "https://www.test.com/a1b2c3d4e5f6g7h8i9j0"
	metadataUrl              = ""

	nameUpdated   string = "test-tf-provider-auth-settings-updated"
	domainUpdated string = "test1.com"
)

func TestAccAuthenticationSettingsResource(t *testing.T) {
	resourceName := "cortexcloud_authentication_settings.test"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccAuthenticationSettingsResourceConfig(name, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "domain", domain),
					resource.TestCheckResourceAttr(resourceName, "default_role", defaultRole),
					resource.TestCheckResourceAttr(resourceName, "mappings.email", mappingsEmail),
					resource.TestCheckResourceAttr(resourceName, "mappings.first_name", mappingsFirstName),
					resource.TestCheckResourceAttr(resourceName, "mappings.last_name", mappingsLastName),
					resource.TestCheckResourceAttr(resourceName, "mappings.group_name", mappingsGroupName),
				),
			},
			// Update and Read testing
			{
				Config: testAccAuthenticationSettingsResourceConfig(nameUpdated, domainUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", nameUpdated),
					resource.TestCheckResourceAttr(resourceName, "domain", domainUpdated),
				),
			},
		},
	})
}

func testAccAuthenticationSettingsResourceConfig(name, domain string) string {
	return fmt.Sprintf(`
resource "cortexcloud_authentication_settings" "test" {
  name   = "%s"
  domain = "%s"
  default_role = "%s"
  is_account_role = %t
  mappings = {
    email      = "%s"
    first_name = "%s"
    last_name  = "%s"
    group_name = "%s"
  }
  idp_sso_url = "%s"
  idp_certificate = "%s"
  idp_issuer = "%s"
  metadata_url = "%s"
}
`, name, domain, defaultRole, isAccountRole, mappingsEmail, mappingsFirstName, mappingsLastName, mappingsGroupName, idpSsoUrl, idpCertificate, idpIssuer, metadataUrl)
}
