package main

import (
	"github.com/RedVentures/shield-provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: shield.Provider,
	})
}
