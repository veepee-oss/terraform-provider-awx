/*
*TBD*

Example Usage

```hcl
*TBD*
```

*/
package awx

import (
	"context"
	"fmt"
	"strconv"

	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		DeleteContext: resourceUserDelete,
		UpdateContext: resourceUserUpdate,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_superuser": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_system_auditor": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*awx.AWX)
	awxService := client.UserService

	result, err := awxService.CreateUser(map[string]interface{}{
		"username":          d.Get("username").(string),
		"password":          d.Get("password").(string),
		"first_name":        d.Get("first_name").(string),
		"last_name":         d.Get("last_name").(string),
		"email":             d.Get("email").(string),
		"is_superuser":      d.Get("is_superuser").(bool),
		"is_system_auditor": d.Get("is_system_auditor").(bool),
	}, map[string]string{})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create new user",
			Detail:   fmt.Sprintf("Unable to create new user: %s", err.Error()),
		})
		return diags
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceUserRead(ctx, d, m)
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.UserService
	var diags diag.Diagnostics
	if diags.HasError() {
		return diags
	}
	id, _ := strconv.Atoi(d.Id())
	_, err := awxService.UpdateUser(id, map[string]interface{}{
		"username":          d.Get("username").(string),
		"password":          d.Get("password").(string),
		"first_name":        d.Get("first_name").(string),
		"last_name":         d.Get("last_name").(string),
		"email":             d.Get("email").(string),
		"is_superuser":      d.Get("is_superuser").(bool),
		"is_system_auditor": d.Get("is_system_auditor").(bool),
	}, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update user",
			Detail:   fmt.Sprintf("Unable to update new user: %s", err.Error()),
		})
		return diags
	}

	return resourceUserRead(ctx, d, m)

}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	var diags diag.Diagnostics
	awxService := client.UserService
	id, _ := strconv.Atoi(d.Id())
	res, err := awxService.GetUserByID(id, make(map[string]string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to fetch user",
			Detail:   fmt.Sprintf("Unable to fetch user: %s", err.Error()),
		})
		return diags
	}
	d.Set("username", res.Username)
	d.Set("password", res.Password)
	d.Set("first_name", res.FirstName)
	d.Set("last_name", res.LastName)
	d.Set("email", res.Email)
	d.Set("is_superuser", res.IsSuperUser)
	d.Set("is_system_auditor", res.IsSystemAuditor)

	return nil
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.UserService
	id, diags := convertStateIDToNummeric("Delete User", d)

	if diags.HasError() {
		return diags
	}

	if _, err := awxService.DeleteUser(id); err != nil {
		return buildDiagDeleteFail(
			"User",
			fmt.Sprintf("id %v, got %s ",
				id, err.Error()))
	}
	d.SetId("")
	return diags
}
