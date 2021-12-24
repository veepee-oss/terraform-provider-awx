---
layout: "awx"
page_title: "AWX: awx_notification_template"
sidebar_current: "docs-awx-resource-notification-template"
description: |-
  *TBD*
---

# awx_notification_template

*TBD*

## Example Usage

```hcl
resource "awx_notification_template" "default" {
    name                      = "schedule-test"
    notification_type         = "webhook"
    organization_id           = data.awx_organization.default.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) 
* `notification_type` - (Required) 
* `organization_id` - (Required) 
* `description` - (Optional) 
* `notification_configuration` - (Optional) 
