/*
This resource configure generic AWX settings.
Please note that resource deletion only delete object from terraform state and do not reset setting to his initial value.

See available settings list here: https://docs.ansible.com/ansible-tower/latest/html/towerapi/api_ref.html#/Settings/Settings_settings_update

Example Usage

```hcl
resource "awx_setting" "social_auth_saml_technical_contact" {
  name       = "SOCIAL_AUTH_SAML_TECHNICAL_CONTACT"
  value_json = {
    givenName = "Myorg"
    emailAddress = "test@foo.com"
  }
}

resource "awx_setting" "social_auth_saml_sp_entity_id" {
  name  = "SOCIAL_AUTH_SAML_SP_ENTITY_ID"
  value = "test"
}

resource "awx_setting" "schedule_max_jobs" {
  name  = "SCHEDULE_MAX_JOBS"
  value = 15
}
```

*/
package awx

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/mrcrilly/goawx/client"
)

func resourceSetting() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSettingUpdate,
		ReadContext:   resourceSettingRead,
		DeleteContext: resourceSettingDelete,
		UpdateContext: resourceSettingUpdate,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of setting to modify",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Value to be modified for given setting. This field is mutually exclusive with the field `value_json`",
				ConflictsWith: []string{
					"value_json",
				},
			},
			"value_json": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Value to be modified for given setting json formated. This field is mutually exclusive with the field `value`",
				ConflictsWith: []string{
					"value",
				},
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
	}
}

type setting map[string]string

func resourceSettingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.SettingService

	_, err := awxService.GetSettingsBySlug("all", make(map[string]string))
	if err != nil {
		return buildDiagnosticsMessage(
			"Create: failed to fetch settings",
			"Failed to fetch setting, got: %s", err.Error(),
		)
	}

	var value interface{}
	name := d.Get("name").(string)

	if v, ok := d.GetOk("value"); ok {
		value = v
	} else if v, ok := d.GetOk("value_json"); ok {
		value = v
	} else {
		return buildDiagnosticsMessage(
			"Wrong input value",
			"`value` or `value_json` need to be define, got: value: %s, value_json: %s", d.Get("value"), d.Get("value_json"),
		)
	}

	payload := map[string]interface{}{
		name: value,
	}

	_, err = awxService.UpdateSettings("all", payload, make(map[string]string))
	if err != nil {
		return buildDiagnosticsMessage(
			"Create: setting not created",
			"failed to save setting data, got: %s, %s", err.Error(), value,
		)
	}

	d.SetId(name)
	return resourceSettingRead(ctx, d, m)
}

func resourceSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	awxService := client.SettingService

	_, err := awxService.GetSettingsBySlug("all", make(map[string]string))
	if err != nil {
		return buildDiagnosticsMessage(
			"Unable to fetch settings",
			"Unable to load settings with slug all: got %s", err.Error(),
		)
	}

	d.Set("name", d.Id())
	d.Set("value", d.Get("value").(string))
	return diags
}

func resourceSettingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	d.SetId("")
	return diags
}
