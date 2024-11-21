package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type UserRouter struct{}

// InitUserRouter 用户相关api
func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	//userRouterWithoutRecord := Router.Group("user").Use(middleware.OperationRecord())
	userApi := api.ApiGroupApp.UserApi
	{
		userRouter.POST("changePassword", userApi.ChangePassword)   // 用户修改密码
		userRouter.POST("changeHeaderImg", userApi.ChangeHeaderImg) // 修改头像
		//userRouter.POST("setUserAuthority", userApi.SetUserAuthority) // 切换角色 重新生成的token这部分处理还没完善
		//userRouter.DELETE("deleteUser", userApi.DeleteUser)               // 删除用户√
		userRouter.PUT("setUserInfo", userApi.SetUserInfo) // 设置用户信息
		userRouter.GET("getSelfInfo", userApi.GetSelfInfo) // 获取自身信息
		userRouter.PUT("setSelfInfo", userApi.SetSelfInfo) // 设置自身信息
		//userRouter.POST("setUserAuthorities", userApi.SetUserAuthorities) // 设置用户权限组
		//userRouter.POST("resetPassword", userApi.ResetPassword)           // 重置用户密码
		//userRouter.GET("getUserTree", userApi.GetUserTree)                // 获取用户树
		userRouter.GET("list", userApi.GetAllUser) // 获取所有用户
	}
}
