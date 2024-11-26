package tables

import (
	"server/global"
	"time"
)

type JournalIssue struct {
	global.MPS_MODEL
	JournalId           uint      `json:"journal_id" gorm:"comment:期刊id"`
	Name                string    `json:"name" gorm:"comment:期刊名称"`
	SubmissionStartTime time.Time `json:"submission_start_time" gorm:"comment:投稿开始时间"`
	SubmissionEndTime   time.Time `json:"submission_end_time" gorm:"comment:投稿截止时间"`
	Description         string    `json:"description" gorm:"comment:描述"`
	Year                int       `json:"year" gorm:"comment:年份"`
	Volume              uint      `json:"volume" gorm:"comment:卷号"`
}
