package request

import "time"

// CreateJournal 创建期刊
type CreateJournal struct {
	CreatorId      uint    `json:"creator_id" comment:"创建者id"`
	Name           string  `json:"name" comment:"期刊名称" binding:"required"`
	CommitteeId    uint    `json:"committee_id" comment:"委员会id" binding:"required"`
	Description    string  `json:"description" comment:"期刊简介" binding:"required"`
	Category       string  `json:"category" gorm:"comment:期刊类别"`
	Presidents     []*Info `json:"presidents" comment:"主席信息" binding:"required"`
	VicePresidents []*Info `json:"vice_presidents" comment:"副主席信息" binding:"required"`
	Members        []*Info `json:"members" comment:"成员信息" binding:"required"`
}

// UpdateJournal 更新期刊
type UpdateJournal struct {
	ID             uint    `json:"id" comment:"期刊ID" binding:"required"`
	Name           string  `json:"name" comment:"期刊名称"`
	Description    string  `json:"description" comment:"期刊简介"`
	Category       string  `json:"category" gorm:"comment:期刊类别"`
	Presidents     []*Info `json:"presidents" comment:"主席信息"`
	VicePresidents []*Info `json:"vice_presidents" comment:"副主席信息"`
	Members        []*Info `json:"members" comment:"成员信息"`
}

// CreateJournalIssue 创建期刊Issue
type CreateJournalIssue struct {
	JournalId           uint      `json:"journal_id" comment:"期刊id" binding:"required"`
	Name                string    `json:"name" comment:"期刊名称" binding:"required"`
	SubmissionStartTime time.Time `json:"submission_start_time" comment:"投稿开始时间" binding:"required"`
	Description         string    `json:"description" comment:"描述"`
	Volume              uint      `json:"volume" comment:"卷号" binding:"required"`
	SubmissionEndTime   time.Time `json:"submission_end_time" comment:"投稿截止时间" binding:"required"`
}

// UpdateJournalIssue 更新期刊Issue
type UpdateJournalIssue struct {
	ID                  uint      `json:"id" comment:"期刊Issue ID" binding:"required"`
	Name                string    `json:"name" comment:"期刊名称"`
	Description         string    `json:"description" comment:"描述"`
	Volume              uint      `json:"volume" comment:"卷号"`
	SubmissionStartTime time.Time `json:"submission_start_time" comment:"投稿开始时间"`
	SubmissionEndTime   time.Time `json:"submission_end_time" comment:"投稿截止时间"`
}
