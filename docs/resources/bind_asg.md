---
layout: "cfsecurity"
page_title: "Cloud Foundry security entitlement: cfsecurity_bind_asg"
sidebar_current: "docs-cfsecurity-resource-asg"
description: Bind a security group to an org through cfsecurity server (useful only for org manager who wants to use terraform).
---

# cfsecurity\_bind\_asg

Bind a security group to an org through cfsecurity server (useful only for org manager who wants to use terraform). Resource only manage entitlement previously set in resource when `force` is to `false`. If entitlements has been added by an other way the provider will not override it.

## Example Usage

Basic usage

```hcl
resource "cfsecurity_bind_asg" "my-bindings" {
  bind {
    asg_id   = "dcee7d89-149b-4bab-9eb9-1e5e73c22aae"
    space_id = "7e0477b9-fff8-41b1-8fd8-969095ba62e5"
  }
  bind {
    asg_id   = "ce9ee907-74a2-4226-a5b2-5b6336973a9e"
    space_id = "11ce76d1-3e17-4479-b090-ff971da597ca"
  }
  force = false
}
```

## Argument Reference

The following arguments are supported:

* `bind` - (Required) A list of entitlements.
    - `asg_id` - (Required, String) a security group to be entitle on the org
    - `space_id` - (Required, String) an organisation guid
* `force` - (Optionnal, boolean) if set to true, resource will overrides security groups assigments for org manager.

## Attributes Reference

The following attributes are exported:

* `id` - A generated GUID
