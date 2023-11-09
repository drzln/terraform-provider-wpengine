package account_user

import (
	"context"

	"github.com/drzln/terraform-provider-wpengine/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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

	client := m.(*client.ApiClient)

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

func resourceWPEngineAccountUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*client.ApiClient)

	// Get the user ID from the resource data
	userID := d.Id()

	// Call the client method to get the user details
	user, err := client.GetAccountUser(userID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set the resource data from the user details
	d.Set("first_name", user["first_name"])
	d.Set("last_name", user["last_name"])
	d.Set("email", user["email"])

	return diags
}

func resourceWPEngineAccountUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// var diags diag.Diagnostics

	client := m.(*client.ApiClient)

	userID := d.Id()

	// Check which fields have changed
	if d.HasChanges("first_name", "last_name", "email") {
		userData := map[string]interface{}{
			"first_name": d.Get("first_name").(string),
			"last_name":  d.Get("last_name").(string),
			"email":      d.Get("email").(string),
		}

		// Call the client method to update the user details
		_, err := client.UpdateAccountUser(userID, userData)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Call read to update the state
	return resourceWPEngineAccountUserRead(ctx, d, m)
}

func resourceWPEngineAccountUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*client.ApiClient)

	userID := d.Id()

	err := client.DeleteAccountUser(userID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Remove the user ID from the state as it no longer exists
	d.SetId("")

	return diags
}

// end resourceWPEngineAccountUser
