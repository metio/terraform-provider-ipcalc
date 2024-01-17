# SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
# SPDX-License-Identifier: 0BSD

data "ipcalc_reverse_dns" "reverse_dns" {
  ip_address = var.ip_address
}
