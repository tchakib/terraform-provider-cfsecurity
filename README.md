# Terraform-provider-cfsecurity

This is a terraform provider for using [cf-security-entitlement](https://github.com/orange-cloudfoundry/cf-security-entitlement).

This provider has been made to be used with [cloud foundry provider](https://github.com/cloudfoundry-community/terraform-provider-cf).

You can find documentation at https://registry.terraform.io/providers/orange-cloudfoundry/cfsecurity/latest/docs

## Installations

**Requirements:** You need, of course, terraform (**>=0.13**) which is available here: https://www.terraform.io/downloads.html

Add to your terraform file:

```hcl
terraform {
  required_providers {
    cfsecurity = {
      source  = "orange-cloudfoundry/cfsecurity"
      version = "latest"
    }
  }
}
```

## Documentation

You can find documentation at https://registry.terraform.io/providers/orange-cloudfoundry/cfsecurity/latest/docs