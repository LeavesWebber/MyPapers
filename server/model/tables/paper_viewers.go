package tables

import (
	"server/global"
)

type PaperViewers struct {
	global.MPS_MODEL
	PaperId  uint `json:"paper_id" gorm:"comment:论文id"`
	ViewerId uint `json:"viewer_id" gorm:"comment:可查看者id"`
}
