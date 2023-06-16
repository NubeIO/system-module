package main

import (
	"github.com/NubeIO/module-core-system/logger"
	"github.com/NubeIO/module-core-system/pkg"
	"github.com/NubeIO/rubix-os/module/shared"
	"github.com/hashicorp/go-plugin"
)

func ServePlugin(module *pkg.Module) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         plugin.PluginSet{"module-core-system": &shared.NubeModule{Impl: module}},
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}

func main() {
	module := &pkg.Module{}
	logger.SetLogger(module.Config)
	ServePlugin(module)
}
