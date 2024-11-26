package tables

import "server/global"

type Review struct {
	global.MPS_MODEL
	ReviewerId uint   `json:"reviewer_id" gorm:"comment:审稿人id"`
	PaperId    uint   `json:"paper_id" gorm:"comment:论文id"`
	Comment    string `json:"comment" gorm:"comment:审稿意见;type:text"`
	Status     string `json:"status" gorm:"comment:审稿状态"`
	OldVersion bool   `gorm:"column:old_version"`
}
