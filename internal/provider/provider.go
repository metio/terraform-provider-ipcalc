/*
 * SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ provider.Provider = (*IPCalcProvider)(nil)
)

type IPCalcProvider struct{}

type IPCalcProviderModel struct{}

func New() provider.Provider {
	return &IPCalcProvider{}
}

func (p *IPCalcProvider) Metadata(_ context.Context, _ provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "ipcalc"
}

func (p *IPCalcProvider) Schema(_ context.Context, _ provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Provider for the ipcalc-like functionality. Requires Terraform 1.0 or later.",
		MarkdownDescription: "Provider for the [ipcalc](https://gitlab.com/ipcalc/ipcalc)-like functionality. Requires Terraform 1.0 or later.",
		Attributes:          map[string]schema.Attribute{},
	}
}

func (p *IPCalcProvider) Configure(ctx context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
	tflog.Info(ctx, "IPCalcProvider configured")
}

func (p *IPCalcProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewReverseDNSDataSource,
	}
}

func (p *IPCalcProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
