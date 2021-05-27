---
layout: "cfsecurity"
page_title: "Cloud Foundry security entitlement: cfsecurity_asg"
sidebar_current: "docs-cfsecurity-datasource-asg"
description: Get information on a Cloud Foundry Application Security Group from security entitlement api.
---

# cfsecurity\_asg

Retrieve a security group id by its name (useful only for org managers who wants to use terraform).

## Example Usage

```hcl
data "cfsecurity_asg" "entitled" {
  name = "entitled"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The name of the application security group to lookup

## Attributes Reference

The following attributes are exported:

- `id` - The GUID of the application security group
