# SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
# SPDX-License-Identifier: 0BSD

output "ip_address" {
  value = data.ipcalc_reverse_dns.reverse_dns.ip_address
}

output "reverse_dns" {
  value = data.ipcalc_reverse_dns.reverse_dns.reverse_dns
}
