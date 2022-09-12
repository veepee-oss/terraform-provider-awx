/*
*TBD*

Example Usage

```hcl
resource "awx_organization" "my_orga" {
  name = "My AWX Organization"
  ...
}

data "awx_organization_role" "orga_admins" {
  name       = "Admin"
  organization_id = resource.awx_organization.my_orga.id
}
```

*/
package awx

import (
	"context"
	"strconv"

	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOrganizationRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrganizationRolesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organization_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func dataSourceOrganizationRolesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	params := make(map[string]string)

	organization_id := d.Get("organization_id").(int)

	Organization, err := client.OrganizationsService.GetOrganizationsByID(organization_id, params)
	if err != nil {
		return buildDiagnosticsMessage(
			"Get: Fail to fetch Organization",
			"Fail to find the organization got: %s",
			err.Error(),
		)
	}

	roleslist := []*awx.ApplyRole{
		Organization.SummaryFields.ObjectRoles.ExecuteRole,
		Organization.SummaryFields.ObjectRoles.ProjectAdminRole,
		Organization.SummaryFields.ObjectRoles.InventoryAdminRole,
		Organization.SummaryFields.ObjectRoles.CredentialAdminRole,
		Organization.SummaryFields.ObjectRoles.WorkflowAdminRole,
		Organization.SummaryFields.ObjectRoles.NotificationAdminRole,
		Organization.SummaryFields.ObjectRoles.JobTemplateAdminRole,
		Organization.SummaryFields.ObjectRoles.ExecuteEnvironmentsAdminRole,
		Organization.SummaryFields.ObjectRoles.AuditorRole,
		Organization.SummaryFields.ObjectRoles.ReadRole,
		Organization.SummaryFields.ObjectRoles.ApprovalRole,

		// Only for User object
		Organization.SummaryFields.ObjectRoles.AdminRole,
		Organization.SummaryFields.ObjectRoles.MemberRole,
	}

	if roleID, okID := d.GetOk("id"); okID {
		id := roleID.(int)
		for _, v := range roleslist {
			if v != nil && id == v.ID {
				d = setOrganizationRoleData(d, v)
				return diags
			}
		}
	}

	if roleName, okName := d.GetOk("name"); okName {
		name := roleName.(string)

		for _, v := range roleslist {
			if v != nil && name == v.Name {
				d = setOrganizationRoleData(d, v)
				return diags
			}
		}
	}

	return buildDiagnosticsMessage(
		"Failed to fetch organization role - Not Found",
		"The organization role was not found",
	)
}

func setOrganizationRoleData(d *schema.ResourceData, r *awx.ApplyRole) *schema.ResourceData {
	d.Set("name", r.Name)
	d.SetId(strconv.Itoa(r.ID))
	return d
}
