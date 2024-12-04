package logic

import (
	"server/dao/mysql"
	"server/model/request"
	"server/model/tables"
)

// AddBaseMenu 新增菜单
func AddBaseMenu(m *request.Menu) error {
	return mysql.AddBaseMenu(m)
}

// SetMenuAuthority 设置menu和角色的对应关系
func SetMenuAuthority(s *request.SetMenuAuthorityInfo) error {
	return mysql.SetMenuAuthority(s)
}

// AddMenuAuthority 为角色增加menu对应关系
func AddMenuAuthority(menus []tables.BaseMenu, authorityId uint) (err error) {
	var auth request.SetMenuAuthorityInfo
	auth.AuthorityId = authorityId
	auth.BaseMenus = menus
	err = SetMenuAuthority(&auth)
	return err
}

// DeleteBaseMenu 删除菜单
func DeleteBaseMenu(menuId uint) (err error) {
	err = mysql.DeleteBaseMenu(menuId)
	return err
}

// GetMenuTree 获取动态菜单树
func GetMenuTree(authorityId uint) (menus []tables.BaseMenu, err error) {
	menuTree, err := mysql.GetMenuTreeMap(authorityId) // 获取菜单的父子对应关系（此时还是乱序的）
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = getMenuChildrenList(&menus[i], menuTree) // 从根节点开始遍历整理菜单树（从根节点开始有序）
	}
	return menus, err
}

// getChildrenList 获取子菜单
func getMenuChildrenList(menu *tables.BaseMenu, treeMap map[uint][]tables.BaseMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getMenuChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
