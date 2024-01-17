# SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
# SPDX-License-Identifier: 0BSD

terraform {
  required_providers {
    ipcalc = {
      source  = "localhost/metio/ipcalc"
      version = "9999.99.99"
    }
  }
}

provider "ipcalc" {}
