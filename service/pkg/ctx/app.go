package ctx

import (
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadmin/service/config"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/pkg/event"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type WithOption func(ctx *AppCtx)

type AppCtx struct {
	// GsR web 引擎
	GsR *gin.Engine

	// GsAppPath 项目路径
	GsAppPath string

	// Config 全局配置
	Config *config.Config

	// Db 全局数据库
	Db *gorm.DB

	// SuperAdmin 超级管理员标识
	SuperAdmin string

	// EventDispatcher 事件分发器
	EventDispatcher *event.DispatcherEvent

	// Limiter 限流器
	Limiter *rate.Limiter

	// Logger 日志工具
	Logger *zap.SugaredLogger
}

func NewDefaultAppCtx(options ...WithOption) *AppCtx {
	appCtx := new(AppCtx)
	appCtx.GsR = global.GsE
	appCtx.Logger = global.Logger
	appCtx.Limiter = global.Limiter
	appCtx.EventDispatcher = &global.EventDispatcher
	appCtx.Db = global.Db
	appCtx.GsAppPath = global.GsAppPath
	appCtx.SuperAdmin = global.SuperAdmin
	appCtx.Config = global.Config

	for _, option := range options {
		option(appCtx)
	}
	return appCtx
}
