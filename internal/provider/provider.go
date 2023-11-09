// Copyright (c) drzln, LOL.

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/drzln/terraform-provider-wpengine/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				"wpengine_data_source": dataSourceScaffolding(),
			},
			ResourcesMap: map[string]*schema.Resource{
				// "wpengine_account":      resourceWPEngineAccount(),
				"wpengine_account_user": resourceWPEngineAccountUser(),
				// "wpengine_site":         resourceWPEngineSite(),
				// "wpengine_install":      resourceWPEngineInstall(),
				// "wpengine_domain":       resourceWPEngineDomain(),
				// "wpengine_ssh_key":      resourceWPEngineSshKey(),
				// "wpengine_cdn":          resourceWPEngineCdn(),
				// potentially unCRUDable
				// "wpengine_cache": resourceWPEngineCache(),
				// "wpengine_backup": resourceWPEngineBackup(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)
		return p
	}
}

// #############################################################################
// resourceWPEngineAccount
// #############################################################################

// func resourceWPEngineAccount() *schema.Resource {
// 	return &schema.Resource{
// 		CreateContext: resourceWPEngineAccountCreate,
// 		// ReadContext:   resourceWPEngineAccountRead,
// 		// UpdateContext: resourceWPEngineAccountUpdate,
// 		// DeleteContext: resourceWPEngineAccountDelete,
// 	}
// }

// end resourceWPEngineAccount


// ############################################################################
// config
// ############################################################################

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-wpengine", version)
		// TODO: myClient.UserAgent = userAgent

		return &client.ApiClient{}, nil
	}
}

// end config
