package test

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/pkg/ctx"
)

type Test struct {
	app *ctx.AppCtx
}

func (b *Test) SetApp(app *ctx.AppCtx) *Test {
	b.app = app
	return b
}

func (b *Test) Install() error {

	b.app.GsR.GET("asfasf", func(context *gin.Context) {
		context.JSON(200, gin.H{"a": "b"})
	})
	return nil
}

func (b *Test) UnInstall() error {

	return nil
}

func (b *Test) GetPluginName() string {
	return "test"
}
