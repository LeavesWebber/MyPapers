package tables

// UserPaper 是 User 和 Paper 的连接表
type UserPaper struct {
	UserId     uint `gorm:"comment:作者id;column:user_id"`
	PaperId    uint `gorm:"column:paper_id"`
	OldVersion bool `gorm:"column:old_version"`
}

func (s *UserPaper) TableName() string {
	return "user_paper"
}
