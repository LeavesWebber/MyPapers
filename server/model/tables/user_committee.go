package tables

import "time"

// UserCommittee 是 User 和 UserCommittee 的连接表
type UserCommittee struct {
	UserId      uint      `gorm:"comment:成员id;column:user_id"`
	CommitteeId uint      `gorm:"column:committee_id"`
	StartTime   time.Time `gorm:"comment:成员加入委员会时间;column:start_time"`
	EndTime     time.Time `gorm:"comment:成员离开委员会时间;column:end_time"`
	Position    string    `gorm:"comment:成员职位;column:position"`
	Level       string    `gorm:"comment:成员级别;column:level"`
}

func (s *UserCommittee) TableName() string {
	return "user_committee"
}
