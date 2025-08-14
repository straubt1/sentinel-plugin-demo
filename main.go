package main

import (
	"sentinel-plugin-demo/plugin"

	sdk "github.com/hashicorp/sentinel-sdk"
	"github.com/hashicorp/sentinel-sdk/framework"
	"github.com/hashicorp/sentinel-sdk/rpc"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		PluginFunc: func() sdk.Plugin {
			return &framework.Plugin{Root: &plugin.Root{}}
		},
	})
}
