---
layout: "cfsecurity"
page_title: "Cloud Foundry security entitlement: cfsecurity_entitle_asg"
sidebar_current: "docs-cfsecurity-resource-asg"
description: Entitle a security group to an org. Resource only manage entitlement previously set in resource. If entitlements has been added by an other way the provider will not override it.
---

# cfsecurity\_entitle\_asg

Entitle a security group to an org. Resource only manage entitlement previously set in resource. If entitlements has been added by an other way the provider will not override it.

## Example Usage

Basic usage

```hcl
resource "cfsecurity_entitle_asg" "my-entitlements" {
  entitle {
    asg_id = "dcee7d89-149b-4bab-9eb9-1e5e73c22aae"
    org_id = "7e0477b9-fff8-41b1-8fd8-969095ba62e5"
  }
  entitle {
    asg_id = "ce9ee907-74a2-4226-a5b2-5b6336973a9e"
    org_id = "11ce76d1-3e17-4479-b090-ff971da597ca"
  }
}
```

## Argument Reference

The following arguments are supported:

* `entitle` - (Required) A list of entitlements.
    - `asg_id` - (Required, String) a security group to be entitle on the org
    - `org_id` - (Required, String) an organisation guid

## Attributes Reference

The following attributes are exported:

* `id` - A generated GUID
