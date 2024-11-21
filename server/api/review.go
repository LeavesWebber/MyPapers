package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/logic"
	"server/model/request"
	"server/utils"
)

type ReviewApi struct{}

// GetAllReviews 查询审核列表
func (r *ReviewApi) GetAllReviews(c *gin.Context) {
	// 1. 获取当前userId
	userInfo, err := utils.GetCurrentUserInfo(c)
	filter := c.Query("filter")
	// 2. 逻辑处理
	out, err := logic.GetAllReviews(filter, userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("logic.GetAllReviews failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// AllotReviewers 分配审核人
func (r *ReviewApi) AllotReviewers(c *gin.Context) {
	// 1. 获取参数和校验参数
	in := new(request.AllotReviewers)
	if err := c.ShouldBindJSON(in); err != nil {
		// 请求参数有误，直接返回响应
		global.MPS_LOG.Error("AllotReviewers with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 2. 逻辑处理
	err := logic.AllotReviewers(in)
	if err != nil {
		global.MPS_LOG.Error("logic.AllotReview failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// SubmitReview 审核
func (r *ReviewApi) SubmitReview(c *gin.Context) {
	// 1. 获取当前userId
	userInfo, err := utils.GetCurrentUserInfo(c)
	// 2. 获取参数和校验参数
	in := new(request.SubmitReview)
	if err := c.ShouldBindJSON(in); err != nil {
		// 请求参数有误，直接返回响应
		global.MPS_LOG.Error("SubmitReview with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 3. 逻辑处理
	out, err := logic.SubmitReview(in, userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("logic.SubmitReview failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}
