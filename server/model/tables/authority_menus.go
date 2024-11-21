package tables

type AuthorityMenu struct {
	MenuId      string `json:"menuId" gorm:"comment:菜单ID;column:base_menu_id"`
	AuthorityId string `json:"-" gorm:"comment:角色ID;column:authority_authority_id"`
}

func (s AuthorityMenu) TableName() string {
	return "authority_menus"
}
