package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/orange-cloudfoundry/terraform-provider-cfsecurity/cfsecurity"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cfsecurity.Provider})

}
