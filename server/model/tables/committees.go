package tables

import "server/global"

type Committee struct {
	global.MPS_MODEL
	CreatorId   uint   `json:"creator_id" gorm:"comment:创建者ID"`
	Name        string `json:"name" gorm:"comment:委员会名称"`
	Description string `json:"description" gorm:"comment:描述;type:text"`
	Users       []User `json:"-" gorm:"many2many:user_committee;"`
}
