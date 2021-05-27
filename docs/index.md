---
layout: "cfsecurity"
page_title: "Provider: Cloud Foundry security entitlement"
sidebar_current: "docs-cfsecurity-index"
description: |- Provider for using cloud foundry security entitlement api.
---

# Cloud Foundry security entitlement Provider

Provider for using cloud foundry security entitlement api, this will let you interact with api to handle entitlement and binding of security groups.

You will be able to use this provider alongside cloud foundry provider.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "cfsecurity" {
  cf_api_url          = "https://api.[your domain]"
  user                = "admin user cloud foundry"
  password            = "admin password"
  skip_ssl_validation = false
}
```

## Argument Reference

The following arguments are supported:

* `cf_api_url` - (Required) API endpoint (e.g. https://api.local.pcfdev.io). This can also be specified with the `CF_API_URL` shell environment variable.

* `cf_security_url` - (Optional) This is by default set to `https://cfsecurity.[your domain]` (e.g.: https://cfsecurity.local.pcfdev.io). This is the URL to cfsecurity server. Can be defined with the `CF_SECURITY_URL` shell environment variable.

* `user` - (Optional) Cloud Foundry user. Defaults to "admin". This can also be specified with the `CF_USER` shell environment variable. Unless mentionned explicitly in a resource, CF admin permissions are not required.

* `password` - (Optional) Cloud Foundry user's password. This can also be specified with the `CF_PASSWORD` shell environment variable.

* `cf_client_id` - (Optional) The cf client ID to make request with a client instead of user. This can also be specified with the `CF_CLIENT_ID` shell environment variable.

* `cf_client_secret` - (Optional) The cf client secret to make request with a client instead of user. This can also be specified with the `CF_CLIENT_SECRET` shell environment variable.

* `skip_ssl_validation` - (Optional) Skip verification of the API endpoint - Not recommended!. Defaults to "false". This can also be specified with the `CF_SKIP_SSL_VALIDATION` shell environment variable.