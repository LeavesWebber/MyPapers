package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/logic"
	"server/model/request"
	"server/utils"
	"strconv"
)

type AuthorityMenuApi struct {
}

// AddBaseMenu 新增菜单
func (a *AuthorityMenuApi) AddBaseMenu(c *gin.Context) {
	// 请求参数校验
	p := new(request.Menu)
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
	err := logic.AddBaseMenu(p)
	if err != nil {
		global.MPS_LOG.Error("logic.AddBaseMenu() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// SetMenuAuthority 设置menu和角色关联关系
func (a *AuthorityMenuApi) SetMenuAuthority(c *gin.Context) {
	// 请求参数校验
	p := new(request.SetMenuAuthorityInfo)
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
	err := logic.SetMenuAuthority(p)
	if err != nil {
		global.MPS_LOG.Error("logic.SetMenuAuthority() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// DeleteBaseMenu 删除菜单
func (a *AuthorityMenuApi) DeleteBaseMenu(c *gin.Context) {
	// 请求参数校验
	p := new(request.DeleteBaseMenu)
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
	err := logic.DeleteBaseMenu(p.ID)
	if err != nil {
		global.MPS_LOG.Error("logic.DeleteBaseMenu() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// GetMenu 获取菜单树
func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	// 获取用户角色id
	userInfo, err := utils.GetCurrentUserInfo(c)
	if err != nil {
		global.MPS_LOG.Error("logic.GetCurrentUserInfo() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	authorityId := c.Query("authorityId")
	if authorityId != "" {
		// 字符串转uint
		i, err := strconv.Atoi(authorityId)
		if err != nil {
			ResponseError(c, CodeServerBusy)
			return
		}
		userInfo.AuthorityId = uint(i)
	}
	// 逻辑处理
	menus, err := logic.GetMenuTree(userInfo.AuthorityId)
	if err != nil {
		global.MPS_LOG.Error("logic.GetMenuTree() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, menus)
}
