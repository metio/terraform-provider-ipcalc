/*
 * SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider_test

import (
	"context"
	"fmt"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/metio/terraform-provider-ipcalc/internal/provider"
	"regexp"
	"testing"
)

func TestReverseDNSDataSource_Schema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	provider.NewReverseDNSDataSource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)
	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

func TestReverseDNSDataSource(t *testing.T) {
	testCases := map[string]struct {
		ipAddress  string
		reverseDNS string
	}{
		"ipv4-short": {
			ipAddress:  "1.2.3.4",
			reverseDNS: "4.3.2.1.in-addr.arpa.",
		},
		"ipv4-long": {
			ipAddress:  "192.168.128.64",
			reverseDNS: "64.128.168.192.in-addr.arpa.",
		},
		"ipv4-wikipedia": {
			ipAddress:  "203.0.113.1",
			reverseDNS: "1.113.0.203.in-addr.arpa.",
		},
		"ipv6": {
			ipAddress:  "2a02:8108:8ac0:295:d60c:76e5:1daf:6b11",
			reverseDNS: "1.1.b.6.f.a.d.1.5.e.6.7.c.0.6.d.5.9.2.0.0.c.a.8.8.0.1.8.2.0.a.2.ip6.arpa.",
		},
		"ipv6-wikipedia": {
			ipAddress:  "2001:db8::567:89ab",
			reverseDNS: "b.a.9.8.7.6.5.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa.",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			resource.UnitTest(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: providerConfig() + fmt.Sprintf(`
							data "ipcalc_reverse_dns" "test" {
								ip_address  = "%s"
							}
						`, testCase.ipAddress),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.ipcalc_reverse_dns.test", "ip_address", testCase.ipAddress),
							resource.TestCheckResourceAttr("data.ipcalc_reverse_dns.test", "reverse_dns", testCase.reverseDNS),
						),
					},
				},
			})
		})
	}
}

func TestReverseDNSDataSource_Configuration_Errors(t *testing.T) {
	testCases := map[string]ConfigurationErrorTestCase{
		"empty-ip-address": {
			Configuration: `
				ip_address = ""
			`,
			ErrorRegex: "Attribute ip_address string length must be at least 1",
		},
		"missing-ip-address": {
			Configuration: `
			`,
			ErrorRegex: `The argument "ip_address" is required, but no definition was found`,
		},
		"invalid-ip-address": {
			Configuration: `
				local_part  = "a.b.c.d"
			`,
			ErrorRegex: "Domain names must be convertible to ASCII",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			resource.UnitTest(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: providerConfig() + fmt.Sprintf(`
							data "ipcalc_reverse_dns" "test" {
								%s
							}
						`, testCase.Configuration),
						ExpectError: regexp.MustCompile(testCase.ErrorRegex),
					},
				},
			})
		})
	}
}
