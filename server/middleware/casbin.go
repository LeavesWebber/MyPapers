package middleware

import (
	"server/api"
	"server/global"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := utils.GetCurrentUserInfo(c) // 从token中解析出用户信息
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(user.AuthorityId))
		e := utils.Casbin()                    // 实例化一个策略，把数据库中的数据读出来，下面的Enforce就可以通过这些数据去管理权限
		success, _ := e.Enforce(sub, obj, act) // 把这3个参数传递给e.Enforce， 就可以实现对Web页面和请求接口的权限控制管理
		if global.MPS_CONFIG.System.Env == "develop" || success {
			c.Next()
		} else {
			api.ResponseError(c, api.CodeInsufficientPermissions)
			c.Abort()
			return
		}
	}
}
