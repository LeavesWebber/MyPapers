package response

import (
	"server/model/request"
	"time"
)

// ConferenceDetail 委员会详情
type ConferenceDetail struct {
	CreatorId      uint            `json:"creator_id" comment:"创建人Id"`
	Name           string          `json:"name" comment:"会议名称"`
	Description    string          `json:"description" comment:"会议简介"`
	Category       string          `json:"category" comment:"会议类别"`
	Location       string          `json:"location" gorm:"comment:会议地点"`
	StartTime      time.Time       `json:"start_time" gorm:"comment:会议开始时间"`
	EndTime        time.Time       `json:"end_time" gorm:"comment:会议结束时间"`
	Presidents     []*request.Info `json:"presidents" comment:"主席信息"`
	VicePresidents []*request.Info `json:"vice_presidents" comment:"副主席信息"`
	Members        []*request.Info `json:"members" comment:"成员信息"`
}
