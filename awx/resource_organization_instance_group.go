/*
*TBD*

Example Usage

```hcl
resource "awx_organization_instance_group" "baseconfig" {
  organization_id = awx_organization.baseconfig.id
  instance_group_id   = awx_instance_group.baseconfig.id
}
```

*/
package awx

import (
	"context"
	"fmt"
	"strconv"

	awx "github.com/veepee-oss/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrganizationsGalaxyCredentials() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrganizationsInstanceGroupsCreate,
		DeleteContext: resourceOrganizationsInstanceGroupsDelete,
		ReadContext:   resourceOrganizationsInstanceGroupsRead,

		Schema: map[string]*schema.Schema{

			"organization_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"instance_group_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceOrganizationsInstanceGroupsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	awxService := client.OrganizationsService
	OrganizationID := d.Get("organization_id").(int)
	_, err := awxService.GetOrganizationsByID(OrganizationID, make(map[string]string))
	if err != nil {
		return buildDiagNotFoundFail("organization", OrganizationID, err)
	}

	result, err := awxService.AssociateInstanceGroups(OrganizationID, map[string]interface{}{
		"id": d.Get("credential_id").(int),
	}, map[string]string{})

	if err != nil {
		return buildDiagnosticsMessage("Create: Organization not AssociateInstanceGroups", "Fail to add Instance Group with Id %v, for Organization ID %v, got error: %s", d.Get("credential_id").(int), OrganizationID, err.Error())
	}

	d.SetId(strconv.Itoa(result.ID))
	return diags
}

func resourceOrganizationsInstanceGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceOrganizationsInstanceGroupsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	awxService := client.OrganizationsService
	OrganizationID := d.Get("organization_id").(int)
	res, err := awxService.GetOrganizationsByID(OrganizationID, make(map[string]string))
	if err != nil {
		return buildDiagNotFoundFail("organization", OrganizationID, err)
	}

	_, err = awxService.DisAssociateInstanceGroups(res.ID, map[string]interface{}{
		"id": d.Get("credential_id").(int),
	}, map[string]string{})
	if err != nil {
		return buildDiagDeleteFail("Organization DisAssociateInstanceGroups", fmt.Sprintf("DisAssociateInstanceGroups %v, from OrganizationID %v got %s ", d.Get("credential_id").(int), d.Get("organization_id").(int), err.Error()))
	}

	d.SetId("")
	return diags
}
