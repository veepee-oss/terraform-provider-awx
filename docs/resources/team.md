---
layout: "awx"
page_title: "AWX: awx_team"
sidebar_current: "docs-awx-resource-team"
description: |-
  *TBD*
---

# awx_team

*TBD*

## Example Usage

```hcl
data "awx_organization" "default" {
  name = "Default"
}

data "awx_inventory" "myinv" {
  name = "My Inventory"
}

data "awx_inventory_role" "myinv_admins" {
  name         = "Admin"
  inventory_id = data.awx_inventory.myinv.id
}

data "awx_project" "myproj" {
  name = "My Project"
}

data "awx_project_role" "myproj_admins" {
  name = "Admin"
  project_id = data.awx_project.myproj.id
}

resource "awx_team" "admins_team" {
  name                 = "admins-team"
  organization_id      = data.awx_organization.default.id

  role_entitlement {
    role_id = data.awx_inventory_role.myinv_admins.id
  }
  role_entitlement {
    role_id = data.awx_project_role.myproj_admins.id
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this team
* `organization_id` - (Required) Numeric ID of the team organization
* `description` - (Optional) Optional description of this team
* `role_entitlement` - (Optional) Set of role IDs for access by this team

