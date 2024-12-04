package request

import (
	"server/global"
	"server/model/tables"
)

// CreateAuthority 创建用户请求参数
type CreateAuthority struct {
	AuthorityId   uint   `json:"authority_id" binding:"required"`   // 角色ID
	AuthorityName string `json:"authority_name" binding:"required"` // 角色名
	ParentId      uint   `json:"parent_id"`                         // 父角色ID
}

// GetAuthorityList 获取角色列表
type GetAuthorityList struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// DeleteAuthority 删除角色
type DeleteAuthority struct {
	AuthorityId uint `json:"authority_id"`
}

// ChangeAuthority 切换角色
type ChangeAuthority struct {
	AuthorityId uint `json:"authority_id"`
}

// DefaultMenu 默认角色配置的菜单
func DefaultMenu() []tables.BaseMenu {
	return []tables.BaseMenu{{
		MPS_MODEL: global.MPS_MODEL{ID: 1},
		ParentId:  0,
		Path:      "dashboard",
		Name:      "dashboard",
		Url:       "Home.vue",
		Title:     "仪表盘",
		Icon:      "setting",
	}}
}
