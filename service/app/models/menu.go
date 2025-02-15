package models

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/datatypes"
)

type AdminMenu struct {
	global.GsModel
	Meta      datatypes.JSONMap `gorm:"type:varchar(500);column:meta;comment:元数据" json:"meta"`
	Component string            `gorm:"type:varchar(100);column:component;comment:组件" json:"component"`
	Name      string            `gorm:"type:varchar(80);column:name;comment:别名" json:"name"`
	ParentId  uint              `gorm:"type:int(11);column:parent_id;comment:上级id" json:"parent_id"`
	Sort      int               `gorm:"type:int(11);column:sort;comment:排序;default:0" json:"sort"`
	Path      string            `gorm:"type:varchar(100);column:path;comment:路径" json:"path"`
	Redirect  string            `gorm:"type:varchar(200);column:redirect;comment:重定向uri" json:"redirect"`
	ApiList   []MenuApiList     `gorm:"foreignKey:MenuId;references:ID" json:"apiList"`
}

type TreeMenu struct {
	AdminMenu
	Children []*TreeMenu `json:"children"`
}

func (a AdminMenu) Test() {

	fmt.Println("test-------")
}
