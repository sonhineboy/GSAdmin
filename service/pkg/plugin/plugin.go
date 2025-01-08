package plugin

import (
	"github.com/pkg/errors"
	"sync"
)

type Installer struct {
	Plugins []Plugin
	once    sync.Once
}

func NewInstallerPlugins() *Installer {
	return new(Installer)
}

func (i *Installer) Register(plugins ...Plugin) {
	i.once.Do(func() {
		i.Plugins = append(i.Plugins, plugins...)
	})
}

func (i *Installer) Install() {
	var err error
	for _, plugin := range i.Plugins {
		err = plugin.Install()
		if err != nil {
			panic(errors.WithStack(err))
		}
	}
}

var ErrPluginNotName = errors.New("plugin name not set")

type Plugin interface {
	Install() error
	UnInstall() error
	GetPluginName() string
}
