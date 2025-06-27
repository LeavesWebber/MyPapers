package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`    // 响应码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

// 响应码常量
const (
	CodeSuccess          = 200 // 成功
	CodeInvalidParam     = 400 // 参数错误
	CodeUnauthorized     = 401 // 未授权
	CodePermissionDenied = 403 // 权限不足
	CodeNotFound         = 404 // 未找到
	CodeServerBusy       = 500 // 服务器繁忙
	CodeNotImplemented   = 501 // 未实现
)

// 响应消息常量
var (
	MessageSuccess          = "操作成功"
	MessageInvalidParam     = "参数错误"
	MessageUnauthorized     = "未授权"
	MessagePermissionDenied = "权限不足"
	MessageNotFound         = "资源未找到"
	MessageServerBusy       = "服务器繁忙"
	MessageNotImplemented   = "功能未实现"
)

// ResponseSuccess 成功响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: MessageSuccess,
		Data:    data,
	})
}

// ResponseError 错误响应
func ResponseError(c *gin.Context, code int) {
	var message string
	switch code {
	case CodeInvalidParam:
		message = MessageInvalidParam
	case CodeUnauthorized:
		message = MessageUnauthorized
	case CodePermissionDenied:
		message = MessagePermissionDenied
	case CodeNotFound:
		message = MessageNotFound
	case CodeServerBusy:
		message = MessageServerBusy
	case CodeNotImplemented:
		message = MessageNotImplemented
	default:
		message = MessageServerBusy
	}

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ResponseErrorWithMessage 自定义错误消息响应
func ResponseErrorWithMessage(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
