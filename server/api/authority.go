package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/logic"
	"server/model/request"
	"server/model/response"
	"server/utils"
)

type AuthorityApi struct{}

// CreateAuthority 创建角色
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	// 请求参数校验
	p := new(request.CreateAuthority)
	if err := c.ShouldBindJSON(p); err != nil {
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
	if authority, err := logic.CreateAuthority(p); err != nil {
		global.MPS_LOG.Error("logic.CreateAuthority() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 返回响应
		ResponseSuccess(c, authority)
	}
}

// DeleteAuthority 删除角色
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	// 请求参数校验
	p := new(request.DeleteAuthority)
	if err := c.ShouldBindJSON(p); err != nil {
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
	err := logic.DeleteAuthority(p.AuthorityId)
	if err != nil {
		global.MPS_LOG.Error("logic.DeleteAuthority() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// UpdateAuthority 更新角色信息
func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	// 请求参数校验
	p := new(request.CreateAuthority)
	if err := c.ShouldBindJSON(p); err != nil {
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
	if authority, err := logic.UpdateAuthority(p); err != nil {
		global.MPS_LOG.Error("logic.UpdateAuthority() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 返回响应
		ResponseSuccess(c, authority)
	}
}

// GetAuthorityList 获取角色列表
func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	// 请求参数校验
	p := new(request.GetAuthorityList)
	if err := c.ShouldBindJSON(p); err != nil {
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
	list, total, err := logic.GetAuthorityList(p)
	if err != nil {
		global.MPS_LOG.Error("logic.GetAuthorityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, response.GetAuthorityList{
		List:     list,
		Total:    total,
		Page:     p.Page,
		PageSize: p.PageSize,
	})
}

// ChangeAuthority 切换角色
func (a *AuthorityApi) ChangeAuthority(c *gin.Context) {
	// 请求参数校验
	p := new(request.ChangeAuthority)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 从上下文中获取用户id
	userInfo, _ := utils.GetCurrentUserInfo(c)
	// 逻辑处理
	if err := logic.ChangeAuthority(userInfo.ID, p); err != nil {
		global.MPS_LOG.Error("logic.ChangeAuthority() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}
