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

type ConferenceApi struct{}

// CreateConference 创建委员会
func (con *ConferenceApi) CreateConference(c *gin.Context) {
	in := new(request.CreateConference)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("CreateConference with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam) // 非validator.ValidationErrors类型错误直接返回
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 2. 业务处理
	// 获取当前userId
	userInfo, _ := utils.GetCurrentUserInfo(c)
	in.CreatorId = userInfo.ID
	if conference, err := logic.CreateConference(in); err != nil {
		global.MPS_LOG.Error("logic.CreateConference() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, conference)
	}
}

// UpdateConference 更新委员会
func (con *ConferenceApi) UpdateConference(c *gin.Context) {
	in := new(request.UpdateConference)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("UpdateConference with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}

	// 2. 业务处理
	if out, err := logic.UpdateConference(in); err != nil {
		global.MPS_LOG.Error("logic.UpdateConference failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, out)
	}
}

// GetConference 获取委员会详情
func (con *ConferenceApi) GetConference(c *gin.Context) {
	// 1. 获取参数和校验参数
	conferenceID := c.Query("conference_id")
	inID, err := strconv.Atoi(conferenceID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if conferenceDetail, err := logic.GetConference(id); err != nil {
		global.MPS_LOG.Error("logic.GetConference failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, conferenceDetail)
	}
}

// GetAllConferences 获取会议列表
func (con *ConferenceApi) GetAllConferences(c *gin.Context) {
	//  业务处理
	if conferenceList, err := logic.GetAllConferences(); err != nil {
		global.MPS_LOG.Error("logic.ConferenceList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, conferenceList)
	}
}

// GetAllConferencesByCommittee 根据委员会查询会议列表
func (con *ConferenceApi) GetAllConferencesByCommittee(c *gin.Context) {
	// 1. 获取参数和校验参数
	committeeID := c.Query("committee_id")
	inID, err := strconv.Atoi(committeeID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if conferenceList, err := logic.GetAllConferencesByCommittee(id); err != nil {
		global.MPS_LOG.Error("logic.GetAllConferencesByCommittee failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, conferenceList)
	}

}

// DeleteConference 删除会议
func (con *ConferenceApi) DeleteConference(c *gin.Context) {
	// 1. 获取参数和校验参数
	conferenceID := c.Query("conference_id")
	inID, err := strconv.Atoi(conferenceID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if err := logic.DeleteConference(id); err != nil {
		global.MPS_LOG.Error("logic.DeleteConference failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, nil)
	}
}

// GetConferencesByUser 查询自己所在的会议列表
func (con *ConferenceApi) GetConferencesByUser(c *gin.Context) {
	// 1. 获取当前userId
	userInfo, _ := utils.GetCurrentUserInfo(c)
	// 2. 业务处理
	if conferenceList, err := logic.GetConferencesByUser(userInfo.ID); err != nil {
		global.MPS_LOG.Error("logic.GetConferencesByUser failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, conferenceList)
	}
}

// CreateConferenceIssue 创建Issue
func (con *ConferenceApi) CreateConferenceIssue(c *gin.Context) {
	in := new(request.CreateConferenceIssue)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("CreateConferenceIssue with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 2. 业务处理
	if err := logic.CreateConferenceIssue(in); err != nil {
		global.MPS_LOG.Error("logic.CreateConferenceIssue() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, nil)
	}
}

// GetAllConferenceIssues  查询会议Issue列表
func (con *ConferenceApi) GetAllConferenceIssues(c *gin.Context) {
	// 1. 获取参数和校验参数
	conferenceID := c.Query("conference_id")
	inID, err := strconv.Atoi(conferenceID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if conferenceIssues, err := logic.GetAllConferenceIssues(id); err != nil {
		global.MPS_LOG.Error("logic.GetConferenceIssues failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, conferenceIssues)
	}
}

// UpdateConferenceIssue 更新会议Issue
func (con *ConferenceApi) UpdateConferenceIssue(c *gin.Context) {
	in := new(request.UpdateConferenceIssue)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("UpdateConferenceIssue with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 2. 业务处理
	if err := logic.UpdateConferenceIssue(in); err != nil {
		global.MPS_LOG.Error("logic.UpdateConferenceIssue failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, nil)
	}
}

// DeleteConferenceIssue 删除会议Issue
func (con *ConferenceApi) DeleteConferenceIssue(c *gin.Context) {
	// 1. 获取参数和校验参数
	issueID := c.Query("issue_id")
	inID, err := strconv.Atoi(issueID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if err := logic.DeleteConferenceIssue(id); err != nil {
		global.MPS_LOG.Error("logic.DeleteConferenceIssue failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, nil)
	}
}

// GetLevelInConference 获取用户在会议的level
func (con *ConferenceApi) GetLevelInConference(c *gin.Context) {
	// 1. 获取参数和校验参数
	conferenceID := c.Query("conference_id")
	inID, err := strconv.Atoi(conferenceID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 获取当前userId
	userInfo, _ := utils.GetCurrentUserInfo(c)
	// 2. 业务处理
	if conferenceIssues, err := logic.GetLevelInConference(id, userInfo.ID); err != nil {
		global.MPS_LOG.Error("logic.GetConferenceIssues failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, conferenceIssues)
	}
}
