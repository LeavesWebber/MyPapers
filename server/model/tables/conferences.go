package tables

import (
	"server/global"
	"time"
)

type Conference struct {
	global.MPS_MODEL
	CommitteeId uint      `json:"committee_id" gorm:"comment:委员会id"`
	CreatorId   uint      `json:"creator_id" gorm:"comment:创建者id"`
	Name        string    `json:"name" gorm:"comment:会议名称"`
	Description string    `json:"description" gorm:"comment:描述;type:text"`
	Location    string    `json:"location" gorm:"comment:会议地点"`
	Category    string    `json:"category" gorm:"comment:期刊类别"`
	StartTime   time.Time `json:"start_time" gorm:"comment:会议开始时间"`
	EndTime     time.Time `json:"end_time" gorm:"comment:会议结束时间"`
	Users       []User    `json:"-" gorm:"many2many:user_conference;"`
}
