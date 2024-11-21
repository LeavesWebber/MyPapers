package tables

import "server/global"

// BaseMenu 菜单表
type BaseMenu struct {
	global.MPS_MODEL
	ParentId    uint        `json:"parentId" gorm:"comment:父菜单ID"` // 父菜单ID
	Sort        int         `json:"sort" gorm:"comment:排序标记"`      // 排序标记
	Path        string      `json:"path" gorm:"comment:路由path"`    // 路由path
	Name        string      `json:"name" gorm:"comment:路由name"`    // 路由name
	Url         string      `json:"url" gorm:"comment:对应前端文件路径"`   // 对应前端文件路径
	Title       string      `json:"title" gorm:"comment:菜单名"`      // 菜单名
	Icon        string      `json:"icon" gorm:"comment:菜单图标"`      // 菜单图标
	Authorities []Authority `json:"authorities" gorm:"many2many:authority_menus;"`
	Children    []BaseMenu  `json:"children,omitempty" gorm:"-"`
}
