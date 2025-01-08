package initialize

import (
	"github.com/sonhineboy/gsadmin/service/pkg/ctx"
	"github.com/sonhineboy/gsadmin/service/pkg/plugin"
	"github.com/sonhineboy/gsadmin/service/plugins"
)

func PluginInit(ctx *ctx.AppCtx) {
	installer := plugin.NewInstallerPlugins()
	installer.Register(plugins.GetPlugins(ctx)...)
	installer.Install()
}
