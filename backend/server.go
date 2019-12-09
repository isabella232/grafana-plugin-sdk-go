package backend

import (
	"github.com/grafana/grafana-plugin-sdk-go/common"
	plugin "github.com/hashicorp/go-plugin"
)

// Serve starts serving the datasource plugin over gRPC.
//
// The plugin ID should be in the format <org>-<name>-datasource.
func Serve(pluginID string, checkHandler CheckHandler, dataHandler DataQueryHandler, resourceHandler ResourceHandler) error {
	versionedPlugins := map[int]plugin.PluginSet{
		common.ProtocolVersion: {
			pluginID: &DatasourcePluginImpl{
				Impl: backendPluginWrapper{
					dataHandler:     dataHandler,
					checkHandler:    checkHandler,
					resourceHandler: resourceHandler,
				},
			},
		},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig:  common.Handshake,
		VersionedPlugins: versionedPlugins,
		GRPCServer:       plugin.DefaultGRPCServer,
	})

	return nil
}