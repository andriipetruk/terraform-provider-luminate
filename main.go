package main

import (
	"github.com/andriipetruk/terraform-provider-luminate/luminate"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: luminate.Provider})
}
