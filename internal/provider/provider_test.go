/*
 * SPDX-FileCopyrightText: The terraform-provider-ipcalc Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider_test

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	internal "github.com/metio/terraform-provider-ipcalc/internal/provider"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPCalcProvider_Metadata(t *testing.T) {
	t.Parallel()
	p := &internal.IPCalcProvider{}
	request := provider.MetadataRequest{}
	response := &provider.MetadataResponse{}
	p.Metadata(context.TODO(), request, response)

	assert.Equal(t, "ipcalc", response.TypeName, "TypeName")
}

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"ipcalc": providerserver.NewProtocol6WithError(internal.New()),
	}
)
