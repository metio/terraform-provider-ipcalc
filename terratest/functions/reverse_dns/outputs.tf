# SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
# SPDX-License-Identifier: 0BSD

output "reverse_dns" {
  value = provider::ipcalc::reverse_dns(var.ip_address)
}
