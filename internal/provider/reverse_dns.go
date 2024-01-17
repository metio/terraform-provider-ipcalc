/*
 * SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"net"
	"strings"
)

func ReverseDNSInvalidIPAddressError(ipAddress string) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"Invalid IP address",
		fmt.Sprintf("%v is not a valid IP address", ipAddress),
	)
}

func ReverseDNSIPv4(ipAddress string) string {
	splitted := strings.Split(ipAddress, ".")
	parts := make([]string, 0)
	for i := len(splitted) - 1; i >= 0; i-- {
		parts = append(parts, splitted[i])
	}
	joined := strings.Join(parts, ".")

	return fmt.Sprintf("%v.in-addr.arpa.", joined)
}

func ReverseDNSIPv6(ipAddress net.IP) string {
	expandedAddress := expandIPv6Address(ipAddress)
	splitted := strings.Split(strings.ReplaceAll(expandedAddress, ":", ""), "")
	parts := make([]string, 0)
	for i := len(splitted) - 1; i >= 0; i-- {
		parts = append(parts, splitted[i])
	}
	joined := strings.Join(parts, ".")

	return fmt.Sprintf("%v.ip6.arpa.", joined)
}

func expandIPv6Address(ip net.IP) string {
	b := make([]byte, 0, len(ip))

	for i := 0; i < len(ip); i += 2 {
		if i > 0 {
			b = append(b, ':')
		}
		s := (uint32(ip[i]) << 8) | uint32(ip[i+1])
		bHex := fmt.Sprintf("%04x", s)
		b = append(b, bHex...)
	}
	return string(b)
}
