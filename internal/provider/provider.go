// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"strings"

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
				"wpengine_account":      resourceWPEngineAccount(),
				"wpengine_account_user": resourceWPEngineAccountUser(),
				"wpengine_site":         resourceWPEngineSite(),
				"wpengine_install":      resourceWPEngineInstall(),
				"wpengine_domain":       resourceWPEngineDomain(),
				"wpengine_ssh_key":      resourceWPEngineSshKey(),
				"wpengine_cdn":          resourceWPEngineCdn(),
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
// resourceWPEngineAccountUser
// #############################################################################

func resourceWPEngineAccountUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWPEngineAccountUserCreate,
		// ReadContext:   resourceWPEngineAccountUserRead,
		// UpdateContext: resourceWPEngineAccountUserUpdate,
		// DeleteContext: resourceWPEngineAccountUserDelete,
	}
}

func resourceWPEngineAccountUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*apiClient)

	accountID := d.Get("account_id").(string)
	userData := map[string]interface{}{
		"first_name": d.Get("first_name").(string),
		"last_name":  d.Get("last_name").(string),
		"email":      d.Get("email").(string),
	}

	user, err := client.CreateAccountUser(accountID, userData)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(user["user_id"].(string))
	return diags
}

// func resourceWPEngineAccountUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	// Implement user read logic using the API client
// }

// func resourceWPEngineAccountUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	// Implement user update logic using the API client
// }

//	func resourceWPEngineAccountUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
//		// Implement user deletion logic using the API client
//	}
func resourceWPEngineAccount() {}

// end resourceWPEngineAccountUser

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-wpengine", version)
		// TODO: myClient.UserAgent = userAgent

		return &apiClient{}, nil
	}
}
