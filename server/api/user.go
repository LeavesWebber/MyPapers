package api

import (
	"errors"
	"server/global"
	"server/logic"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// Login 用户登录
func (u *UserApi) Login(c *gin.Context) {
	// 请求参数校验
	in := new(request.Login)
	if err := c.ShouldBindJSON(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 判断验证码
	//if store.Verify(in.CaptchaId, in.Captcha, true) == false {
	//	ResponseError(c, CodeInvalidCaptcha)
	//	return
	//}
	// 逻辑处理
	out, err := logic.Login(in)
	if err != nil {
		global.MPS_LOG.Error("logic.Login() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, response.Login{
		UserInfo: out.UserInfo,
		Token:    out.Token,
	})
}

// Register 注册
func (u *UserApi) Register(c *gin.Context) {
	// 请求参数校验
	in := new(request.Register)
	if err := c.ShouldBindJSON(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	//var authorities []tables.Authority
	//for _, v := range p.AuthorityIds {
	//	authorities = append(authorities, tables.Authority{
	//		AuthorityId: v,
	//	})
	//}
	// 逻辑处理
	if err := logic.Register(in); err != nil {
		if err == global.ErrorUserExist {
			global.MPS_LOG.Error("logic.Register() failed", zap.Error(err))
			ResponseError(c, CodeUserExist)
			return
		}
		if err == global.ErrorUserEmailExist {
			global.MPS_LOG.Error("logic.Register() failed", zap.Error(err))
			ResponseError(c, CodeUserEmailExist)
			return
		}
		global.MPS_LOG.Error("logic.Register() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// GetSelfInfo 获取自身信息
func (u *UserApi) GetSelfInfo(c *gin.Context) {
	baseInfo, _ := utils.GetCurrentUserInfo(c)
	out, err := logic.GetSelfInfo(baseInfo.UUID)
	if err != nil {
		global.MPS_LOG.Error("logic.GetSelfInfo() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//ResponseSuccess(c, response.GetSelfInfo{UserInfo: info})
	ResponseSuccess(c, out)
}

//发送邮件的函数

// GetUserTree 获取用户树
//func (u *UserApi) GetUserTree(c *gin.Context) {
//	userTree, err := logic.GetUserTree()
//	if err != nil {
//		global.MPS_LOG.Error("logic.GetAllUser() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, response.GetUserTree{UserTree: userTree})
//}

// GetAllUser 获取所有用户
func (u *UserApi) GetAllUser(c *gin.Context) {
	// 获取当前用户信息
	userInfo, err := utils.GetCurrentUserInfo(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 打印当前用户信息
	global.MPS_LOG.Info("current user info",
		zap.Int64("UUID", userInfo.UUID),
		zap.Uint("AuthorityId", userInfo.AuthorityId),
		zap.String("Username", userInfo.Username))

	// 检查是否有管理员权限
	isAdmin := false
	switch userInfo.AuthorityId {
	case global.SUPER_ADMIN:
		isAdmin = true
		global.MPS_LOG.Info("user is admin")
	}

	// 根据权限调用不同的逻辑
	var users []response.GetAllUser
	if isAdmin {
		users, err = logic.GetAllUser()
	} else {
		user, err := logic.GetSelfInfo(userInfo.UUID)
		if err == nil {
			// 将 tables.User 转换为 response.GetAllUser
			users = []response.GetAllUser{ConvertToGetAllUser(user)}
		}
	}

	if err != nil {
		global.MPS_LOG.Error("get user info failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 打印最终返回的数据
	global.MPS_LOG.Info("response data", zap.Any("users", users))
	ResponseSuccess(c, users)
}

// ChangePassword 用户修改密码
func (u *UserApi) ChangePassword(c *gin.Context) {
	// 请求参数校验
	in := new(request.ChangePassword)
	if err := c.ShouldBindJSON(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	user, err := utils.GetCurrentUserInfo(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 逻辑处理
	if err = logic.ChangePassword(in, user.UUID); err != nil {
		global.MPS_LOG.Error("logic.ChangePassword() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// ChangeHeaderImg 修改头像
func (u *UserApi) ChangeHeaderImg(c *gin.Context) {
	// 请求参数校验
	in := new(request.ChangeHeaderImg)
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 逻辑处理
	user, err := utils.GetCurrentUserInfo(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	url := ""
	if url, err = logic.ChangeHeaderImg(c, user.UUID, in); err != nil {
		global.MPS_LOG.Error("logic.ChangeHeaderImg() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, url)
}

// SendMail 发送邮箱验证
func (u *UserApi) SendMail(c *gin.Context) {

	in := new(request.SendMail)
	if err := c.ShouldBindJSON(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 逻辑处理
	if err := logic.SendMail(in); err != nil {
		if err == global.ErrorUserExist {
			global.MPS_LOG.Error("logic.SendMail() failed", zap.Error(err))
			ResponseError(c, CodeUserExist)
			return
		}
		global.MPS_LOG.Error("logic.SendMail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, nil)
}

// ResetPassword 重置用户密码
func (u *UserApi) ResetPassword(c *gin.Context) {
	//	// 请求参数校验
	//	p := new(request.ResetPassword)
	//	if err := c.ShouldBindJSON(p); err != nil {
	//		errs, ok := err.(validator.ValidationErrors)
	//		if !ok {
	//			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
	//			return
	//		}
	//		// validator.ValidationErrors类型错误则进行翻译
	//		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.GA_TRAN)))
	//		return
	//	}
	//	// 逻辑处理
	//	if err := logic.ResetPassword(p.UUID); err != nil {
	//		global.GA_LOG.Error("logic.ResetPassword() failed", zap.Error(err))
	//		ResponseError(c, CodeServerBusy)
	//		return
	//	}
	//	// 返回响应
	//	ResponseSuccess(c, nil)
}

// DeleteUser 删除用户
func (u *UserApi) DeleteUser(c *gin.Context) {
	//	// 请求参数校验
	//	p := new(request.DeleteUser)
	//	if err := c.ShouldBindJSON(p); err != nil {
	//		errs, ok := err.(validator.ValidationErrors)
	//		if !ok {
	//			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
	//			return
	//		}
	//		// validator.ValidationErrors类型错误则进行翻译
	//		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.GA_TRAN)))
	//		return
	//	}
	//	// 逻辑处理
	//	user, err := utils.GetCurrentUserInfo(c)
	//	if err != nil {
	//		ResponseError(c, CodeNeedLogin)
	//		return
	//	}
	//	if user.ID == p.ID {
	//		ResponseError(c, CodeDeleteSelf)
	//		return
	//	}
	//	if err = logic.DeleteUser(p.ID); err != nil {
	//		global.GA_LOG.Error("logic.DeleteUser() failed", zap.Error(err))
	//		ResponseError(c, CodeServerBusy)
	//		return
	//	}
	//	// 返回响应
	//	ResponseSuccess(c, nil)
}

// SetUserInfo 设置用户信息
func (u *UserApi) SetUserInfo(c *gin.Context) {
	// 1. 获取当前用户信息
	currentUser, err := utils.GetCurrentUserInfo(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 2. 验证是否为管理员
	if currentUser.AuthorityId != global.SUPER_ADMIN {
		ResponseError(c, CodeNoPermission)
		return
	}

	// 3. 请求参数校验
	in := new(request.SetUserInfo)
	if err := c.ShouldBindJSON(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}

	// 4. 不允许修改自己的权限
	if in.ID == currentUser.ID {
		ResponseError(c, CodeCannotModifySelf)
		return
	}

	// 5. 不允许将其他用户设置为超级管理员
	if in.AuthorityId == global.SUPER_ADMIN {
		ResponseError(c, CodeCannotSetSuperAdmin)
		return
	}

	// 6. 执行更新操作
	if err := logic.SetUserInfo(in); err != nil {
		switch err {
		case global.ErrorUserNotExist:
			ResponseError(c, CodeUserNotExist)
		case global.ErrorUserExist:
			ResponseError(c, CodeUserExist)
		default:
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// SetSelfInfo 设置自身信息
func (u *UserApi) SetSelfInfo(c *gin.Context) {
	// 请求参数校验
	in := new(request.SetUserInfo)
	if err := c.ShouldBindJSON(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	baseInfo, err := utils.GetCurrentUserInfo(c)
	if err != nil {
		global.MPS_LOG.Error("logic.GetCurrentUserInfo() failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	in.ID = baseInfo.ID
	// 逻辑处理
	if err = logic.SetUserInfo(in); err != nil {
		global.MPS_LOG.Error("logic.SetUserInfo() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// SetUserAuthorities 设置用户权限组
func (u *UserApi) SetUserAuthorities(c *gin.Context) {
	// 请求参数校验
	in := new(request.SetUserAuthorities)
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}

	// 逻辑处理
	if err := logic.SetUserAuthorities(in.ID, in.AuthorityIds); err != nil {
		global.MPS_LOG.Error("logic.SetUserAuthorities() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUserID 获取当前登录的用户ID
func getCurrentUserID(c *gin.Context) (userID uint, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(uint)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// SetUserAuthority 切换角色
func (u *UserApi) SetUserAuthority(c *gin.Context) {
	//	// 请求参数校验
	//	p := new(request.SetUserAuthority)
	//	if err := c.ShouldBindJSON(p); err != nil {
	//		errs, ok := err.(validator.ValidationErrors)
	//		if !ok {
	//			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
	//			return
	//		}
	//		// validator.ValidationErrors类型错误则进行翻译
	//		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.GA_TRAN)))
	//		return
	//	}
	//	// 逻辑处理
	//	user, err := utils.GetCurrentUserInfo(c)
	//	if err != nil {
	//		global.GA_LOG.Error("logic.GetCurrentUserInfo() failed", zap.Error(err))
	//		ResponseError(c, CodeServerBusy)
	//		return
	//	}
	//	token, err := logic.SetUserAuthority(user, p.AuthorityId)
	//	if err != nil {
	//		global.GA_LOG.Error("logic.SetUserAuthority() failed", zap.Error(err))
	//		ResponseError(c, CodeServerBusy)
	//		return
	//	}
	//	fmt.Println(token)
	//	c.Header("new-token", token)
	//	// 返回响应
	//	ResponseSuccess(c, nil)
}

// ConvertToGetAllUser 将 tables.User 转换为 response.GetAllUser
func ConvertToGetAllUser(user tables.User) response.GetAllUser {
	return response.GetAllUser{
		ID:          user.ID,
		AuthorityId: user.AuthorityId,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Department:  user.Department,
		Phone:       user.Phone,
		Address:     user.Address,
	}
}
