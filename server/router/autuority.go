package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	//authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityApi := api.ApiGroupApp.AuthorityApi
	{
		authorityRouterWithoutRecord.POST("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouterWithoutRecord.DELETE("deleteAuthority", authorityApi.DeleteAuthority) // 删除角色
		authorityRouterWithoutRecord.PUT("updateAuthority", authorityApi.UpdateAuthority)    // 更新角色信息
		authorityRouterWithoutRecord.POST("changeAuthority", authorityApi.ChangeAuthority)   // 切换角色
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
