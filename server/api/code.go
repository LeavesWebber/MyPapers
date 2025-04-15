package api

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserEmailExist
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
	CodeCaptchaFailed
	CodeInvalidCaptcha
	CodeInvalidEmailTime
	CodeInsufficientPermissions
	CodeDeleteSelf
	ErrorArticleNoExist
	CodeInnerError
	CodeInvalidEmailCode
	//增加新的错误码，用来判断用户修改权限
	CodeNoPermission
	CodeCannotModifySelf
	CodeCannotSetSuperAdmin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:          "success",
	CodeInvalidParam:     "请求参数错误",
	CodeUserExist:        "用户已存在",
	CodeUserNotExist:     "用户不存在",
	CodeInvalidPassword:  "用户名或密码错误",
	CodeServerBusy:       "服务繁忙",
	CodeInvalidEmailTime: "验证码请求频繁,请稍后再试",
	CodeInvalidEmailCode: "验证码错误",
	CodeNeedLogin:        "需要登录",
	CodeInvalidToken:     "无效的token",
	CodeInnerError:       "内部错误",
	CodeCaptchaFailed:    "获取验证码失败",
	CodeInvalidCaptcha:   "验证码错误",

	CodeInsufficientPermissions: "权限不足",
	CodeDeleteSelf:              "不可删除自己",
	ErrorArticleNoExist:         "文章不存在",

	CodeNoPermission:        "没有操作权限",
	CodeCannotModifySelf:    "不能修改自己的权限",
	CodeCannotSetSuperAdmin: "不能设置其他用户为超级管理员",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
