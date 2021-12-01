---
layout: "awx"
page_title: "AWX: awx_settings_ldap_team_map"
sidebar_current: "docs-awx-resource-settings_ldap_team_map"
description: |-
  *TBD*
---

# awx_settings_ldap_team_map

*TBD*

## Example Usage

```hcl
data "awx_organization" "default" {
  name = "Default"
}

resource "awx_team" "admin_team" {
        name = "Admins"
        organization_id = data.awx_organization.default.id
}

resource "awx_settings_ldap_team_map" "admin_team_map" {
  name         = resource.awx_team.admin_team.name
  users        = ["CN=MyTeam,OU=Groups,DC=example,DC=com"]
  organization = data.awx_organization.default.name
  remove       = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this team
* `organization` - (Required) Name of the team organization
* `users` - (Optional) Optional list of Group DNs to map access to this team
* `remove` - (Optional) When True, a user who is not a member of the given groups will be removed from the team

