/*
 * SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"net"
)

var (
	_ datasource.DataSource = (*ReverseDNSDataSource)(nil)
)

func NewReverseDNSDataSource() datasource.DataSource {
	return &ReverseDNSDataSource{}
}

type ReverseDNSDataSource struct{}

type ReverseDNSDataSourceModel struct {
	IPAddress  types.String `tfsdk:"ip_address"`
	ReverseDNS types.String `tfsdk:"reverse_dns"`
}

func (d *ReverseDNSDataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_reverse_dns"
}

func (d *ReverseDNSDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Calculate the reverse DNS name of an IP address.",
		MarkdownDescription: "Calculate the reverse DNS name of an IP address.",
		Attributes: map[string]schema.Attribute{
			"ip_address": schema.StringAttribute{
				Description:         "The IP address itself.",
				MarkdownDescription: "The IP address itself.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"reverse_dns": schema.StringAttribute{
				Description:         "The reverse DNS for the given IP address.",
				MarkdownDescription: "The reverse DNS for the given IP address.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},
		},
	}
}

func (d *ReverseDNSDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data ReverseDNSDataSourceModel
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	ipAddress := net.ParseIP(data.IPAddress.ValueString())
	if ipAddress == nil {
		response.Diagnostics.Append(ReverseDNSInvalidIPAddressError(data.IPAddress.ValueString()))
		return
	}

	if ipv4 := ipAddress.To4(); ipv4 != nil {
		data.ReverseDNS = types.StringValue(ReverseDNSIPv4(ipv4.String()))
	} else {
		data.ReverseDNS = types.StringValue(ReverseDNSIPv6(ipAddress))
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
