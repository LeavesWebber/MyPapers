package request

import "server/model/tables"

// SetMenuAuthorityInfo 设置menu和角色关联关系
type SetMenuAuthorityInfo struct {
	BaseMenus   []tables.BaseMenu `json:"base_menus" binding:"required"`
	AuthorityId uint              `json:"authority_id" binding:"required"`
}

// Menu 新增菜单
type Menu struct {
	ParentId uint   `json:"parentId"`                 // 父菜单ID
	Path     string `json:"path" binding:"required"`  // 路由path
	Name     string `json:"name" binding:"required"`  // 路由name
	Url      string `json:"url" binding:"required"`   // 对应前端文件路径
	Title    string `json:"title" binding:"required"` // 菜单名
	Icon     string `json:"icon"`                     // 菜单图标
}

// DeleteBaseMenu 删除菜单
type DeleteBaseMenu struct {
	ID uint `json:"id"`
}

// 到后面可能有用
//type Menu struct {
//	BaseMenu
//	MenuId      string `json:"menuId" gorm:"comment:菜单ID"`
//	AuthorityId uint   `json:"-" gorm:"comment:角色ID"`
//	Children    []Menu `json:"children" gorm:"-"`
//}

//func SetAuthorityMenu() {
//	Menus
//}
