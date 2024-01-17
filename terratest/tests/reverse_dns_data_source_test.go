/*
 * SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
 * SPDX-License-Identifier: 0BSD
 */

package acceptance_test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

			terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
				TerraformDir: "../data-sources/ipcalc_reverse_dns",
				Vars: map[string]interface{}{
					"ip_address": testCase.ipAddress,
				},
			})

			defer terraform.Destroy(t, terraformOptions)
			terraform.InitAndApplyAndIdempotent(t, terraformOptions)

			assert.Equal(t, testCase.ipAddress, terraform.Output(t, terraformOptions, "ip_address"), "ip_address")
			assert.Equal(t, testCase.reverseDNS, terraform.Output(t, terraformOptions, "reverse_dns"), "reverse_dns")
		})
	}
}
