package tables

import (
	"server/global"
)

type Journal struct {
	global.MPS_MODEL
	CommitteeId uint   `json:"committee_id" gorm:"comment:委员会id"`
	CreatorId   uint   `json:"creator_id" gorm:"comment:创建者id"`
	Name        string `json:"name" gorm:"comment:期刊名称"`
	Description string `json:"description" gorm:"comment:描述;type:text"`
	Category    string `json:"category" gorm:"comment:期刊类别"`
	Users       []User `json:"-" gorm:"many2many:user_journal;"`
}
