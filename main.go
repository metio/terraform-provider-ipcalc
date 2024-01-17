/*
 * SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/metio/terraform-provider-ipcalc/internal/provider"
	"log"
)

// Run the documentation generation tool
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name=terraform-provider-ipcalc

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address:         "registry.terraform.io/metio/ipcalc",
		Debug:           debug,
		ProtocolVersion: 6,
	}

	err := providerserver.Serve(context.Background(), provider.New, opts)
	if err != nil {
		log.Fatal(err)
	}
}
