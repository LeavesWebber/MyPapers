package api

import (
	"errors"
	"path/filepath"
	"server/dao/mysql"
	"server/global"
	"server/logic"
	"server/model/request"
	"server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type PaperApi struct{}

// SubmitPaper 投稿
func (p *PaperApi) SubmitPaper(c *gin.Context) {
	var err error
	in := new(request.SubmitPaper)
	// 1. 获取参数和校验参数
	if err = c.ShouldBind(in); err != nil {
		// 请求参数有误，直接返回响应
		global.MPS_LOG.Error("SubmitPaper with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	if in.JournalId == 0 && in.ConferenceId == 0 {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 3. 逻辑处理
	//out, err := logic.SubmitPaper(c, in)
	// 从token中获取用户信息
	userInfo, _ := utils.GetCurrentUserInfo(c)
	out, err := logic.SubmitPaper2(c, in, userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("logic.SubmitPaper failed", zap.Error(err))
		if errors.Is(err, global.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 4. 返回响应
	ResponseSuccess(c, out)
}

// GetPaper 查看投稿详情，包括审核状态、结果和文件路径
func (p *PaperApi) GetPaper(c *gin.Context) {
	// 1. 获取参数和校验参数
	paperID := c.Query("paper_id")
	inID, err := strconv.Atoi(paperID)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := uint(inID)
	// 2. 业务处理
	out, err := logic.GetPaper(id)
	if err != nil {
		global.MPS_LOG.Error("logic.GetPaper failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorArticleNoExist) {
			ResponseError(c, ErrorArticleNoExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, out)
}

// GetAllSelfPapers 查询自己的投稿列表
func (p *PaperApi) GetAllSelfPapers(c *gin.Context) {
	// 1. 获取当前userId
	userInfo, _ := utils.GetCurrentUserInfo(c)
	filter := c.Query("filter")
	// 2. 逻辑处理
	out, err := logic.GetAllSelfPapers(filter, userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("logic.GetAllPapers failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// GetAllPapers 查询投稿列表
func (p *PaperApi) GetAllPapers(c *gin.Context) {
	// 1. 获取参数和校验参数
	journalId := c.Query("journal_id")
	conferenceId := c.Query("conference_id")
	// 2. 逻辑处理
	out, err := logic.GetAllPapers(journalId, conferenceId)
	if err != nil {
		global.MPS_LOG.Error("logic.GetAllPapers failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// UpdatePaper 更新投稿
func (p *PaperApi) UpdatePaper(c *gin.Context) {
	// 1. 获取参数和校验参数
	in := new(request.UpdatePaper)
	if err := c.ShouldBind(in); err != nil {
		// 请求参数有误，直接返回响应
		global.MPS_LOG.Error("UpdatePaper with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 2. 生成文件名和保存路径
	filename := filepath.Base(in.Data.Filename)
	//finalName := fmt.Sprintf("%d_%s", currentUserID, filename)
	//saveFile := filepath.Join("./public/", finalName)
	saveFile := filepath.Join("./public/", filename)
	// 保存文件
	if err := c.SaveUploadedFile(in.Data, saveFile); err != nil {
		global.MPS_LOG.Error("SaveUploadedFile failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 业务处理
	// 从token中获取用户信息
	userInfo, _ := utils.GetCurrentUserInfo(c)
	//out, err := logic.UpdatePaper(saveFile, in)
	out, err := logic.UpdatePaper2(userInfo.ID, saveFile, in)
	if err != nil {
		global.MPS_LOG.Error("logic.UpdatePaper failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, out)
}

// DeletePaper 删除投稿
func (p *PaperApi) DeletePaper(c *gin.Context) {
	// 1. 获取参数和校验参数
	paperId := c.Query("paper_id")
	inId, err := strconv.Atoi(paperId)
	id := uint(inId)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 业务处理
	err = logic.DeletePaper(id)
	if err != nil {
		global.MPS_LOG.Error("logic.DeletePaper failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// GetAllAcceptPapers 查询已经审核通过的投稿列表
func (p *PaperApi) GetAllAcceptPapers(c *gin.Context) {
	// 1. 获取参数和校验参数
	//journalId := c.Query("journal_id")
	//conferenceId := c.Query("conference_id")
	// 2. 逻辑处理
	out, err := logic.GetAllAcceptPapers()
	if err != nil {
		global.MPS_LOG.Error("logic.GetAllAcceptPapers failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// GetAllAcceptPapersByJournalAndTime 按期刊和时间查询已经审核通过的投稿列表
func (p *PaperApi) GetAllAcceptPapersByJournalAndTime(c *gin.Context) {
	// 1. 获取参数和校验参数
	journalId := c.Query("journal_id")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	// 解析时间字符串为time.Time类型2023-08-02T15:04:05Z
	start, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	end, err := time.Parse(time.RFC3339, endTime)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 逻辑处理
	out, err := logic.GetAllAcceptPapersByJournalAndTime(journalId, start, end)
	if err != nil {
		global.MPS_LOG.Error("logic.GetAllAcceptPapersByJournalAndTime failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// GetAllAcceptPapersByConferenceAndTime 按会议和时间查询已经审核通过的投稿列表
func (p *PaperApi) GetAllAcceptPapersByConferenceAndTime(c *gin.Context) {
	// 1. 获取参数和校验参数
	conferenceId := c.Query("conference_id")
	startTime := c.Query("start_time")
	endTime := c.Query("end_time")
	// 解析时间字符串为time.Time类型2023-08-02T15:04:05Z
	start, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	end, err := time.Parse(time.RFC3339, endTime)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 逻辑处理
	out, err := logic.GetAllAcceptPapersByConferenceAndTime(conferenceId, start, end)
	if err != nil {
		global.MPS_LOG.Error("logic.GetAllAcceptPapersByConferenceAndTime failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// GetPaperVersions 查询某个投稿的所有版本
func (p *PaperApi) GetPaperVersions(c *gin.Context) {
	// 1. 获取参数和校验参数
	versionId := c.Query("version_id")
	inId, err := strconv.Atoi(versionId)
	id := uint(inId)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 逻辑处理
	out, err := logic.GetPaperVersions(id)
	if err != nil {
		global.MPS_LOG.Error("logic.GetPaperVersions failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// GetHonoraryCertificate 获取荣誉证书
func (p *PaperApi) GetHonoraryCertificate(c *gin.Context) {
	// 1. 获取参数和校验参数
	paperId := c.Query("paper_id")
	userinfo := c.Query("userInfo")
	inId, err := strconv.Atoi(paperId)
	id := uint(inId)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 逻辑处理
	out, err := logic.GetHonoraryCertificate(id, userinfo)
	if err != nil {
		global.MPS_LOG.Error("logic.GetHonoraryCertificate failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// PublishPaper 发布投稿
func (p *PaperApi) PublishPaper(c *gin.Context) {
	// 1. 获取参数和校验参数
	var err error
	in := new(request.PublishPaper)
	if err = c.ShouldBindJSON(in); err != nil {
		// 请求参数有误，直接返回响应
		global.MPS_LOG.Error("SubmitPaper with invalid param", zap.Error(err))
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
	err = logic.PublishPaper(in)
	if err != nil {
		global.MPS_LOG.Error("logic.PublishPaper failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// AddPaperViewer 设置投稿可查看者
func (p *PaperApi) AddPaperViewer(c *gin.Context) {
	// 1. 获取参数和校验参数
	var err error
	in := new(request.AddPaperViewer)
	if err = c.ShouldBind(in); err != nil {
		// 请求参数有误，直接返回响应
		global.MPS_LOG.Error("SetPaperViewers with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	// 从token中获取用户信息
	userInfo, _ := utils.GetCurrentUserInfo(c)
	in.ViewerId = userInfo.ID
	// 2. 逻辑处理
	err = logic.AddPaperViewer(in)
	if err != nil {
		global.MPS_LOG.Error("logic.AddPaperViewer failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// CheckPaperViewer 查看用户是否有权限查看投稿
func (p *PaperApi) CheckPaperViewer(c *gin.Context) {
	// 1. 获取参数和校验参数
	paperId := c.Query("paper_id")
	inId, err := strconv.Atoi(paperId)
	id := uint(inId)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 逻辑处理
	// 从token中获取用户信息
	userInfo, _ := utils.GetCurrentUserInfo(c)
	out, err := logic.CheckPaperViewer(id, userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("logic.CheckPaperViewer failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// GetMyNFTs 获取我的NFT
func (p *PaperApi) GetMyNFTs(c *gin.Context) {
	// 从token中获取用户信息
	userInfo, _ := utils.GetCurrentUserInfo(c)
	out, err := logic.GetMyNFTs(userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("logic.GetMyNFTs failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// UpdatePrice 更新价格
func (p *PaperApi) UpdatePrice(c *gin.Context) {
	// 1. 获取参数和校验参数
	var err error
	in := new(request.UpdatePrice)
	if err = c.ShouldBindJSON(in); err != nil {
		// 请求参数有误，直接返回响应
		global.MPS_LOG.Error("UpdatePrice with invalid param", zap.Error(err))
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
	err = logic.UpdatePrice(in)
	if err != nil {
		global.MPS_LOG.Error("logic.UpdatePrice failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// GetNFTInfoByTokenId 根据tokenId获取NFT信息
func (p *PaperApi) GetNFTInfoByTokenId(c *gin.Context) {
	// 1. 获取参数和校验参数
	tokenIds := c.Query("token_ids")
	if tokenIds == "" {
		ResponseSuccess(c, nil)
	}
	// 2. 逻辑处理
	out, err := logic.GetNFTInfoByTokenId(tokenIds)
	if err != nil {
		global.MPS_LOG.Error("logic.GetNFTInfoByTokenId failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, out)
}

// UpdatePaperUserId 修改投稿对应的user_id
func (p *PaperApi) UpdatePaperUserId(c *gin.Context) {
	// 1. 获取参数和校验参数
	var err error
	in := new(struct {
		PaperId uint `json:"paper_id"`
	})
	if err = c.ShouldBindJSON(in); err != nil {
		// 请求参数有误，直接返回响应
		global.MPS_LOG.Error("UpdatePaperUserId with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(global.MPS_TRAN)))
		return
	}
	userInfo, _ := utils.GetCurrentUserInfo(c)
	// 2. 逻辑处理
	err = logic.UpdatePaperUserId(in.PaperId, userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("logic.UpdatePaperUserId failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
