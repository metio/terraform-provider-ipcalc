/*
 * SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"regexp"
	"testing"
)

func TestReverseDnsFunction_Valid(t *testing.T) {
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
				TerraformVersionChecks: []tfversion.TerraformVersionCheck{
					tfversion.SkipBelow(tfversion.Version1_8_0),
				},
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`
							output "result" {
								value = provider::ipcalc::reverse_dns("%s")
							}
						`, testCase.ipAddress),
						Check: resource.TestCheckOutput("result", testCase.reverseDNS),
					},
				},
			})
		})
	}
}

func TestReverseDnsFunction_Invalid(t *testing.T) {
	testCases := map[string]struct {
		ipAddress string
		error     string
	}{
		"empty-ip-address": {
			ipAddress: "",
			error:     "Cannot parse IP address ''",
		},
		"invalid-ip-address": {
			ipAddress: "a.b.c.d",
			error:     "Cannot parse IP address 'a.b.c.d'",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			resource.UnitTest(t, resource.TestCase{
				TerraformVersionChecks: []tfversion.TerraformVersionCheck{
					tfversion.SkipBelow(tfversion.Version1_8_0),
				},
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`
							output "result" {
								value = provider::ipcalc::reverse_dns("%s")
							}
						`, testCase.ipAddress),
						ExpectError: regexp.MustCompile(testCase.error),
					},
				},
			})
		})
	}
}
