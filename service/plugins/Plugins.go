package plugins

import (
	"github.com/sonhineboy/gsadmin/service/pkg/ctx"
	"github.com/sonhineboy/gsadmin/service/pkg/plugin"
	"github.com/sonhineboy/gsadmin/service/plugins/test"
)

func GetPlugins(ctx *ctx.AppCtx) []plugin.Plugin {
	var (
		Plugins []plugin.Plugin
	)

	Plugins = append(Plugins,
		//plugin +
		new(test.Test).SetApp(ctx),
	)
	return Plugins
}
