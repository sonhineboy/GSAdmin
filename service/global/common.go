package global

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sonhineboy/gsadmin/service/config"
	"github.com/sonhineboy/gsadmin/service/pkg/event"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"time"
)

var (
	// GsR 全局web 引擎
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
	EventDispatcher *event.Dispatcher
	// Limiter 限流器
	Limiter *rate.Limiter
	// Logger 日志工具
	Logger   *zap.SugaredLogger
	ormTrans = map[string]string{
		"record not found": "数据不存在",
	}
)

// GetError 获取错误信息
func GetError(errs error, r interface{}) string {
	var (
		v validator.ValidationErrors
	)
	if errors.As(errs, &v) {
		return getValidateMsg(v, r)
	} else {
		return errs.Error()
	}
}

func getValidateMsg(errs validator.ValidationErrors, r interface{}) string {
	s := reflect.TypeOf(r)
	for _, fieldError := range errs {
		filed, _ := s.FieldByName(fieldError.Field())
		errTag := fieldError.Tag() + "_msg"
		// 获取对应binding得错误消息
		errTagText := filed.Tag.Get(errTag)
		// 获取统一错误消息
		errText := filed.Tag.Get("msg")
		if errTagText != "" {
			return filed.Tag.Get("json") + ":" + errTagText
		}
		if errText != "" {
			return errText
		}
		return filed.Tag.Get("json") + ":" + fieldError.Tag()
	}
	return ""
}

// Pages 通用分页
func Pages(page int, pageSize int, total int, rows interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	data["page"] = page
	data["pageSize"] = pageSize
	data["rows"] = rows
	data["total"] = total
	return data
}

// IsSuperAdmin 即将废弃，请勿使用
func IsSuperAdmin(roles []string, role string) bool {
	for _, v := range roles {
		if v == role {
			return true
		}

	}
	return false
}

// CaptchaServe 验证码
func CaptchaServe(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		err := captcha.WriteImage(&content, id, width, height)
		if err != nil {
			println(err.Error())
		}
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func GetEventDispatcher(c *gin.Context) *event.Dispatcher {

	v, ok := c.Get("e")

	if ok == false {
		fmt.Print("无法获取对象")
		return nil
	}

	e, ok := v.(event.Dispatcher)

	if ok == false {
		fmt.Print("类型不正确")
		return nil
	}

	return &e
}

func GormTans(err error) error {
	if err != nil {
		if v, ok := ormTrans[err.Error()]; ok {
			return errors.New(v)
		}
	}
	return err
}

func SlicesHasStr(s interface{}, str string) bool {

	if v, ok := s.([]string); ok {
		for _, ss := range v {
			if ss == str {
				return true
			}
		}
	}

	return false

}

func IsSlice(v interface{}) bool {
	_, ok := v.([]interface{})
	return ok
}
