package mysql

import (
	"errors"
	"server/global"
	"server/model/request"
	"server/model/tables"

	"gorm.io/gorm"
)

// AddBaseMenu 新增菜单
func AddBaseMenu(menu *request.Menu) error {
	m := tables.BaseMenu{
		ParentId: menu.ParentId,
		Path:     menu.Path,
		Name:     menu.Name,
		Url:      menu.Url,
		Title:    menu.Title,
		Icon:     menu.Icon,
	}
	if !errors.Is(global.MPS_DB.Where("name = ?", menu.Name).First(&m).Error, gorm.ErrRecordNotFound) {
		return global.ErrMenuExistence
	}
	return global.MPS_DB.Create(&m).Error
}

// SetMenuAuthority 设置menu和角色的对应关系
func SetMenuAuthority(info *request.SetMenuAuthorityInfo) error {
	var s tables.Authority
	global.MPS_DB.Preload("BaseMenus").First(&s, "authority_id = ?", info.AuthorityId)
	err := global.MPS_DB.Model(&s).Association("BaseMenus").Replace(&info.BaseMenus) // 用一个新的关联替换原来的关联
	return err
}

// DeleteBaseMenu 删除菜单
func DeleteBaseMenu(menuId uint) (err error) {
	err = global.MPS_DB.Where("parent_id = ?", menuId).First(&tables.BaseMenu{}).Error
	if err != nil {
		var menu tables.BaseMenu
		db := global.MPS_DB.Preload("Authorities").Where("id = ?", menuId).First(&menu).Delete(&menu)
		if len(menu.Authorities) > 0 {
			err = global.MPS_DB.Model(&menu).Association("Authorities").Delete(&menu.Authorities)
		} else {
			err = db.Error
			if err != nil {
				return
			}
		}
	} else {
		return global.ErrMenuHasChild
	}
	return err
}

// GetMenuTreeMap 获取菜单的父子对应关系（此时还是乱序的）
func GetMenuTreeMap(authorityId uint) (treeMap map[uint][]tables.BaseMenu, err error) {
	var baseMenu []tables.BaseMenu
	treeMap = make(map[uint][]tables.BaseMenu)

	var AuthorityMenus []tables.AuthorityMenu
	err = global.MPS_DB.Where("authority_authority_id = ?", authorityId).Find(&AuthorityMenus).Error // 先找到角色与菜单id的对应关系
	if err != nil {
		return
	}

	var MenuIds []string

	for i := range AuthorityMenus {
		MenuIds = append(MenuIds, AuthorityMenus[i].MenuId)
	}

	err = global.MPS_DB.Where("id in (?)", MenuIds).Order("sort").Find(&baseMenu).Error // 再从所有菜单中找出刚刚的菜单id对应的数据
	if err != nil {
		return
	}
	for _, v := range baseMenu {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v) // 建立父子菜单关系
	}
	return treeMap, err
}
