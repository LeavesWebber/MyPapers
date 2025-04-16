package tables

import (
	"server/global"
	"time"
)

type ConferenceIssue struct {
	global.MPS_MODEL
	ConferenceId        uint      `json:"conference_id" gorm:"comment:会议id"`
	Name                string    `json:"name" gorm:"comment:会议名称"`
	SubmissionStartTime time.Time `json:"submission_start_time" gorm:"comment:投稿开始时间"`
	SubmissionEndTime   time.Time `json:"submission_end_time" gorm:"comment:投稿截止时间"`
	Description         string    `json:"description" gorm:"comment:描述"`
	Year                int       `json:"year" gorm:"comment:年份"`
	Volume              uint      `json:"volume" gorm:"comment:卷号"`
}
