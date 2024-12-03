// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFreeipaHostDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccFreeipaHostDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.freeipa_host.test", "id", "duba-nfws-sgwa01.corp.example.com"),
					resource.TestCheckResourceAttr("data.freeipa_host.test", "fqdn", "duba-nfws-sgwa01.corp.example.com"),
					resource.TestCheckResourceAttr("data.freeipa_host.test", "hostname", "duba-nfws-sgwa01"),
				),
			},
		},
	})
}

const testAccFreeipaHostDataSourceConfig = `
data "freeipa_host" "test" {
  fqdn = "duba-nfws-sgwa01.corp.example.com"
}
`
