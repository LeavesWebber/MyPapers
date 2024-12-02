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

type JournalApi struct{}

// CreateJournal 创建期刊
func (j *JournalApi) CreateJournal(c *gin.Context) {
	in := new(request.CreateJournal)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("CreateJournal with invalid param", zap.Error(err))
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
	if journal, err := logic.CreateJournal(in); err != nil {
		global.MPS_LOG.Error("logic.CreateJournal() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, journal)
	}
}

// UpdateJournal 更新期刊
func (j *JournalApi) UpdateJournal(c *gin.Context) {
	in := new(request.UpdateJournal)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("UpdateJournal with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}

	// 2. 业务处理
	if out, err := logic.UpdateJournal(in); err != nil {
		global.MPS_LOG.Error("logic.UpdateJournal failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, out)
	}
}

// GetJournalsByUser 查询自己所在的期刊列表
func (j *JournalApi) GetJournalsByUser(c *gin.Context) {
	// 1. 获取当前userId
	userInfo, _ := utils.GetCurrentUserInfo(c)
	// 2. 业务处理
	if journalList, err := logic.GetJournalsByUser(userInfo.ID); err != nil {
		global.MPS_LOG.Error("logic.GetJournalsByUser failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, journalList)
	}
}

// GetJournal 获取期刊详情
func (j *JournalApi) GetJournal(c *gin.Context) {
	// 1. 获取参数和校验参数
	journalID := c.Query("journal_id")
	inID, err := strconv.Atoi(journalID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if journalDetail, err := logic.GetJournal(id); err != nil {
		global.MPS_LOG.Error("logic.GetJournal failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, journalDetail)
	}
}

// GetAllJournals 获取期刊列表
func (j *JournalApi) GetAllJournals(c *gin.Context) {
	//  业务处理
	if journalList, err := logic.GetAllJournals(); err != nil {
		global.MPS_LOG.Error("logic.JournalList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, journalList)
	}
}

// GetAllJournalsByCommittee 根据委员会查询期刊列表
func (j *JournalApi) GetAllJournalsByCommittee(c *gin.Context) {
	// 1. 获取参数和校验参数
	committeeID := c.Query("committee_id")
	inID, err := strconv.Atoi(committeeID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if journalList, err := logic.GetAllJournalsByCommittee(id); err != nil {
		global.MPS_LOG.Error("logic.GetAllJournalsByCommittee failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, journalList)
	}

}

// DeleteJournal 删除期刊
func (j *JournalApi) DeleteJournal(c *gin.Context) {
	// 1. 获取参数和校验参数
	journalID := c.Query("journal_id")
	inID, err := strconv.Atoi(journalID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if err := logic.DeleteJournal(id); err != nil {
		global.MPS_LOG.Error("logic.DeleteJournal failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, nil)
	}
}

// CreateJournalIssue 创建Issue
func (j *JournalApi) CreateJournalIssue(c *gin.Context) {
	in := new(request.CreateJournalIssue)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("CreateJournalIssue with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 2. 业务处理
	if err := logic.CreateJournalIssue(in); err != nil {
		global.MPS_LOG.Error("logic.CreateJournalIssue() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, nil)
	}
}

// GetAllJournalIssues  查询期刊Issue列表
func (j *JournalApi) GetAllJournalIssues(c *gin.Context) {
	// 1. 获取参数和校验参数
	journalID := c.Query("journal_id")
	inID, err := strconv.Atoi(journalID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if journalIssues, err := logic.GetAllJournalIssues(id); err != nil {
		global.MPS_LOG.Error("logic.GetJournalIssues failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, journalIssues)
	}
}

// UpdateJournalIssue 更新期刊Issue
func (j *JournalApi) UpdateJournalIssue(c *gin.Context) {
	in := new(request.UpdateJournalIssue)
	// 1. 获取参数和校验参数
	if err := c.ShouldBind(in); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			global.MPS_LOG.Error("UpdateJournalIssue with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 2. 业务处理
	if err := logic.UpdateJournalIssue(in); err != nil {
		global.MPS_LOG.Error("logic.UpdateJournalIssue failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, nil)
	}
}

// DeleteJournalIssue 删除期刊Issue
func (j *JournalApi) DeleteJournalIssue(c *gin.Context) {
	// 1. 获取参数和校验参数
	issueID := c.Query("issue_id")
	inID, err := strconv.Atoi(issueID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	if err := logic.DeleteJournalIssue(id); err != nil {
		global.MPS_LOG.Error("logic.DeleteJournalIssue failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, nil)
	}
}

// GetLevelInJournal 获取用户在期刊的level
func (j *JournalApi) GetLevelInJournal(c *gin.Context) {
	// 1. 获取参数和校验参数
	journalID := c.Query("journal_id")
	inID, err := strconv.Atoi(journalID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 获取当前userId
	userInfo, _ := utils.GetCurrentUserInfo(c)
	// 2. 业务处理
	if journalIssues, err := logic.GetLevelInJournal(id, userInfo.ID); err != nil {
		global.MPS_LOG.Error("logic.GetJournalIssues failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} else {
		// 3. 返回响应
		ResponseSuccess(c, journalIssues)
	}
}
