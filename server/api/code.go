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

	CodeInsufficientPermissions
	CodeDeleteSelf
	ErrorArticleNoExist
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",

	CodeCaptchaFailed:  "获取验证码失败",
	CodeInvalidCaptcha: "验证码错误",

	CodeInsufficientPermissions: "权限不足",
	CodeDeleteSelf:              "不可删除自己",
	ErrorArticleNoExist:         "文章不存在",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
