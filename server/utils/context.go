package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	CtxTokenKey = "mc"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUserInfo 获取当前登录的用户基本信息
func GetCurrentUserInfo(c *gin.Context) (mc BaseClaims, err error) {
	myClime, ok := c.Get(CtxTokenKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	mc, ok = myClime.(BaseClaims)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
