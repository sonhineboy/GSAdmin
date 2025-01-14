package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Model models.Role
	Where map[string]interface{}
}

func (r *RoleRepository) List(page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.Role
		offSet int
	)
	db := global.Db.Model(&r.Model)

	if r.Where != nil && len(r.Where) > 0 {
		db.Where(r.Where)
	}
	db.Count(&total)
	if page == 0 || pageSize == 0 {
		db.Preload("Menus").Order(sortField + " desc" + ",id desc")
	} else {
		offSet = (page - 1) * pageSize
		db.Preload("Menus").Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)
	}
	db.Find(&data)
	return global.Pages(page, pageSize, int(total), data)
}

/*
添加角色
*/
func (r *RoleRepository) Add(post requests.Role) error {
	db := global.Db.Create(&models.Role{
		Alias:  post.Alias,
		Label:  post.Label,
		Sort:   post.Sort,
		Remark: post.Remark,
		Status: &post.Status,
	})
	return db.Error
}

/*
更新角色
*/
func (r *RoleRepository) Update(post requests.Role) error {

	return global.Db.Transaction(func(sessionDb *gorm.DB) error {
		return sessionDb.Debug().Where("id = ?", post.Id).Updates(&models.Role{
			Alias:  post.Alias,
			Label:  post.Label,
			Sort:   post.Sort,
			Remark: post.Remark,
			Status: &post.Status,
		}).Error
	})

}

func (r *RoleRepository) UpMenus(post requests.RoleUpMenus) error {
	var role models.Role
	role.ID = post.Id

	if len(post.Menus) > 0 {

		var replace []models.AdminMenu

		for _, v := range post.Menus {
			var li models.AdminMenu
			li.ID = v
			replace = append(replace, li)

		}
		return global.Db.Model(&role).Omit("Menus.*").Association("Menus").Replace(replace)
	} else {
		return global.Db.Model(&role).Association("Menus").Clear()
	}

}
