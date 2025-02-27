package middleware

import (
	"server/model/response"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.FailWithMessage("请求头中auth为空", c)
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithMessage("请求头中auth格式有误", c)
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := utils.ParseToken(parts[1])
		if err != nil {
			response.FailWithMessage("无效的Token", c)
			c.Abort()
			return
		}

		// 将当前请求的claims信息保存到请求的上下文c上
		c.Set("claims", mc)
		c.Next() // 后续的处理函数可以用过c.Get("claims")来获取当前请求的用户信息
	}
}
