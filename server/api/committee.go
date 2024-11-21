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

type CommitteeApi struct{}

// CreateCommittee 创建委员会
func (com *CommitteeApi) CreateCommittee(c *gin.Context) {
	in := new(request.CreateCommittee)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("CreateCommittee with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 获取当前userId
	userInfo, _ := utils.GetCurrentUserInfo(c)
	in.CreatorId = userInfo.ID
	// 2. 业务处理
	if committee, err := logic.CreateCommittee(in); err != nil {
		global.MPS_LOG.Error("logic.CreateCommittee() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, committee)
	}
}

// UpdateCommittee 更新委员会
func (com *CommitteeApi) UpdateCommittee(c *gin.Context) {
	in := new(request.UpdateCommittee)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("UpdateCommittee with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}

	// 2. 业务处理
	if out, err := logic.UpdateCommittee(in); err != nil {
		global.MPS_LOG.Error("logic.UpdateCommittee failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, out)
	}
}

// GetCommitteesByUser 查询自己所在的委员会列表
func (com *CommitteeApi) GetCommitteesByUser(c *gin.Context) {
	// 1. 获取当前userId
	userInfo, _ := utils.GetCurrentUserInfo(c)
	// 2. 业务处理
	if journalList, err := logic.GetCommitteesByUser(userInfo.ID); err != nil {
		global.MPS_LOG.Error("logic.GetCommitteesByUser failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, journalList)
	}
}

// GetCommittee 获取委员会详情
func (com *CommitteeApi) GetCommittee(c *gin.Context) {
	// 1. 获取参数和校验参数
	committeeID := c.Query("committee_id")
	inID, err := strconv.Atoi(committeeID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if committeeDetail, err := logic.GetCommittee(id); err != nil {
		global.MPS_LOG.Error("logic.GetCommittee failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, committeeDetail)
	}
}

// GetAllCommittees 获取委员会列表
func (com *CommitteeApi) GetAllCommittees(c *gin.Context) {
	//  业务处理
	if committeeList, err := logic.GetAllCommittees(); err != nil {
		global.MPS_LOG.Error("logic.CommitteeList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, committeeList)
	}
}

// DeleteCommittee 删除委员会
func (com *CommitteeApi) DeleteCommittee(c *gin.Context) {
	// 1. 获取参数和校验参数
	committeeID := c.Query("committee_id")
	inID, err := strconv.Atoi(committeeID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if err := logic.DeleteCommittee(id); err != nil {
		global.MPS_LOG.Error("logic.DeleteCommittee failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, nil)
	}
}
