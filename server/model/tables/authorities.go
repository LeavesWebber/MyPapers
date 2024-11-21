package tables

import "time"

type Authority struct {
	CreatedAt     time.Time   // 创建时间
	UpdatedAt     time.Time   // 更新时间
	DeletedAt     *time.Time  `sql:"index"`
	AuthorityId   uint        `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	ParentId      uint        `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	AuthorityName string      `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	Children      []Authority `json:"children" gorm:"-"`                                                   // 遍历时的节点
	BaseMenus     []BaseMenu  `json:"menus" gorm:"many2many:authority_menus;"`
	Users         []User      `json:"-" gorm:"many2many:user_authority;"`
}
