package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"net"
)

var _ function.Function = (*ReverseDnsFunction)(nil)

func NewReverseDnsFunction() function.Function {
	return &ReverseDnsFunction{}
}

type ReverseDnsFunction struct{}

func (f *ReverseDnsFunction) Metadata(_ context.Context, _ function.MetadataRequest, response *function.MetadataResponse) {
	response.Name = "reverse_dns"
}

func (f *ReverseDnsFunction) Definition(_ context.Context, _ function.DefinitionRequest, response *function.DefinitionResponse) {
	response.Definition = function.Definition{
		Summary:     "Calculate the reverse DNS name of an IP address",
		Description: "Accepts both IPv4 and IPv6 addresses and outputs their reverse DNS entry.",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "ip_address",
				Description: "The IPv4 or IPv6 address itself",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *ReverseDnsFunction) Run(ctx context.Context, request function.RunRequest, response *function.RunResponse) {
	var input string

	response.Error = function.ConcatFuncErrors(response.Error, request.Arguments.Get(ctx, &input))

	ipAddress := net.ParseIP(input)
	if ipAddress == nil {
		response.Error = function.ConcatFuncErrors(response.Error,
			function.NewArgumentFuncError(0, fmt.Sprintf("Cannot parse IP address '%s'", input)))
		return
	}

	var result string
	if ipv4 := ipAddress.To4(); ipv4 != nil {
		result = ReverseDNSIPv4(ipv4.String())
	} else {
		result = ReverseDNSIPv6(ipAddress)
	}

	response.Error = function.ConcatFuncErrors(response.Error, response.Result.Set(ctx, result))
}
