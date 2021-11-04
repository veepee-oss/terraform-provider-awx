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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/mrcrilly/goawx/client"
)

func resourceInventorySource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInventorySourceCreate,
		ReadContext:   resourceInventorySourceRead,
		UpdateContext: resourceInventorySourceUpdate,
		DeleteContext: resourceInventorySourceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled_var": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"overwrite": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"overwrite_vars": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"update_on_launch": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"inventory_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"credential_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "scm",
				Optional: true,
			},
			"source_vars": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_cache_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"verbosity": &schema.Schema{
				Type:     schema.TypeInt,
				Default:  1,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceInventorySourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventorySourcesService

	result, err := awxService.CreateInventorySource(map[string]interface{}{
		"name":                 d.Get("name").(string),
		"description":          d.Get("description").(string),
		"enabled_var":          d.Get("enabled_var").(string),
		"enabled_value":        d.Get("enabled_value").(string),
		"overwrite":            d.Get("overwrite").(bool),
		"overwrite_vars":       d.Get("overwrite_vars").(bool),
		"update_on_launch":     d.Get("update_on_launch").(bool),
		"inventory":            d.Get("inventory_id").(int),
		"credential":           d.Get("credential_id").(int),
		"source":               d.Get("source").(string),
		"source_vars":          d.Get("source_vars").(string),
		"host_filter":          d.Get("host_filter").(string),
		"update_cache_timeout": d.Get("update_cache_timeout").(int),
		"verbosity":            d.Get("verbosity").(int),
		// "source_project":   d.Get("source_project_id").(int),
		// "source_path":      d.Get("source_path").(string),
	}, map[string]string{})
	if err != nil {
		return buildDiagCreateFail(diagElementInventorySourceTitle, err)
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceInventorySourceRead(ctx, d, m)

}

func resourceInventorySourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventorySourcesService
	id, diags := convertStateIDToNummeric(diagElementInventorySourceTitle, d)
	if diags.HasError() {
		return diags
	}

	_, err := awxService.UpdateInventorySource(id, map[string]interface{}{
		"name":                 d.Get("name").(string),
		"description":          d.Get("description").(string),
		"enabled_var":          d.Get("enabled_var").(string),
		"enabled_value":        d.Get("enabled_value").(string),
		"overwrite":            d.Get("overwrite").(bool),
		"overwrite_vars":       d.Get("overwrite_vars").(bool),
		"update_on_launch":     d.Get("update_on_launch").(bool),
		"inventory":            d.Get("inventory_id").(int),
		"credential":           d.Get("credential_id").(int),
		"source":               d.Get("source").(string),
		"source_vars":          d.Get("source_vars").(string),
		"host_filter":          d.Get("host_filter").(string),
		"update_cache_timeout": d.Get("update_cache_timeout").(int),
		"verbosity":            d.Get("verbosity").(int),
		// "source_project":   d.Get("source_project_id").(int),
		// "source_path":      d.Get("source_path").(string),
	}, nil)
	if err != nil {
		return buildDiagUpdateFail(diagElementInventorySourceTitle, id, err)
	}

	return resourceInventorySourceRead(ctx, d, m)
}

func resourceInventorySourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventorySourcesService
	id, diags := convertStateIDToNummeric(diagElementInventorySourceTitle, d)
	if diags.HasError() {
		return diags
	}
	if _, err := awxService.DeleteInventorySource(id); err != nil {
		return buildDiagDeleteFail(
			"inventroy source",
			fmt.Sprintf("inventroy source %v, got %s ",
				id, err.Error()))
	}
	d.SetId("")
	return nil
}

func resourceInventorySourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventorySourcesService
	id, diags := convertStateIDToNummeric(diagElementInventorySourceTitle, d)
	if diags.HasError() {
		return diags
	}
	res, err := awxService.GetInventorySourceByID(id, make(map[string]string))
	if err != nil {
		return buildDiagNotFoundFail(diagElementInventorySourceTitle, id, err)
	}
	d = setInventorySourceResourceData(d, res)
	return nil
}

func setInventorySourceResourceData(d *schema.ResourceData, r *awx.InventorySource) *schema.ResourceData {
	d.Set("name", r.Name)
	d.Set("description", r.Description)
	d.Set("enabled_var", r.EnabledVar)
	d.Set("enabled_value", r.EnabledValue)
	d.Set("overwrite", r.Overwrite)
	d.Set("overwrite_vars", r.OverwriteVars)
	d.Set("update_on_launch", r.UpdateOnLaunch)
	d.Set("inventory_id", r.Inventory)
	d.Set("credential_id", r.Credential)
	d.Set("source", r.Source)
	d.Set("source_vars", normalizeJsonYaml(r.SourceVars))
	d.Set("host_filter", r.HostFilter)
	d.Set("update_cache_timeout", r.UpdateCacheTimeout)
	d.Set("verbosity", r.Verbosity)
	// d.Set("source_project_id", r.SourceProject)
	// d.Set("source_path", r.SourcePath)

	return d
}
