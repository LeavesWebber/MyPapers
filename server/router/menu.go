package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type MenuRouter struct {
}

func (m *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	//menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := api.ApiGroupApp.AuthorityMenuApi
	{
		menuRouterWithoutRecord.POST("addBaseMenu", authorityMenuApi.AddBaseMenu)           // 新增菜单
		menuRouterWithoutRecord.POST("setMenuAuthority", authorityMenuApi.SetMenuAuthority) // 设置menu和角色关联关系
		menuRouterWithoutRecord.DELETE("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)   // 删除菜单
		//menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)     // 更新菜单(不太需要)
	}
	{
		menuRouterWithoutRecord.GET("getMenu", authorityMenuApi.GetMenu) // 获取菜单树√
		//menuRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList)           // 分页获取基础menu列表(不一定要)
		//menuRouterWithoutRecord.POST("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   // 获取用户动态路由(和获取菜单树差不多，一个是baseMenu，一个是Menu)
		//0menuRouterWithoutRecord.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // 获取指定角色menu
		//menuRouterWithoutRecord.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById)   // 根据id获取菜单
	}
}
