package request

import "time"

// CreateConference 创建会议
type CreateConference struct {
	Name           string    `json:"name" comment:"会议名称" binding:"required"`
	CreatorId      uint      `json:"creator_id" comment:"创建者id"`
	CommitteeId    uint      `json:"committee_id" comment:"委员会id" binding:"required"`
	Description    string    `json:"description" comment:"会议简介" binding:"required"`
	Location       string    `json:"location" gorm:"comment:会议地点"`
	Category       string    `json:"category" gorm:"comment:期刊类别"`
	StartTime      time.Time `json:"start_time" gorm:"comment:会议开始时间"`
	EndTime        time.Time `json:"end_time" gorm:"comment:会议结束时间"`
	Presidents     []*Info   `json:"presidents" comment:"主席信息" binding:"required"`
	VicePresidents []*Info   `json:"vice_presidents" comment:"副主席信息" binding:"required"`
	Members        []*Info   `json:"members" comment:"成员信息" binding:"required"`
}

// UpdateConference 更新会议
type UpdateConference struct {
	ID             uint      `json:"id" comment:"会议ID" binding:"required"`
	Name           string    `json:"name" comment:"会议名称"`
	Description    string    `json:"description" comment:"会议简介"`
	Category       string    `json:"category" gorm:"comment:会议类别"`
	Location       string    `json:"location" gorm:"comment:会议地点"`
	StartTime      time.Time `json:"start_time" gorm:"comment:会议开始时间"`
	EndTime        time.Time `json:"end_time" gorm:"comment:会议结束时间"`
	Presidents     []*Info   `json:"presidents" comment:"主席信息"`
	VicePresidents []*Info   `json:"vice_presidents" comment:"副主席信息"`
	Members        []*Info   `json:"members" comment:"成员信息"`
}

// CreateConferenceIssue 创建期刊Issue
type CreateConferenceIssue struct {
	ConferenceId        uint      `json:"conference_id" comment:"期刊id" binding:"required"`
	Name                string    `json:"name" comment:"期刊名称" binding:"required"`
	SubmissionStartTime time.Time `json:"submission_start_time" comment:"投稿开始时间" binding:"required"`
	Description         string    `json:"description" comment:"描述"`
	Volume              uint      `json:"volume" comment:"卷号" binding:"required"`
	SubmissionEndTime   time.Time `json:"submission_end_time" comment:"投稿截止时间" binding:"required"`
}

// UpdateConferenceIssue 更新期刊Issue
type UpdateConferenceIssue struct {
	ID                  uint      `json:"id" comment:"期刊Issue ID" binding:"required"`
	Name                string    `json:"name" comment:"期刊名称"`
	Description         string    `json:"description" comment:"描述"`
	Volume              uint      `json:"volume" comment:"卷号"`
	SubmissionStartTime time.Time `json:"submission_start_time" comment:"投稿开始时间"`
	SubmissionEndTime   time.Time `json:"submission_end_time" comment:"投稿截止时间"`
}
