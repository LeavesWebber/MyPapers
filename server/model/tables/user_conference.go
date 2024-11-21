package tables

import "time"

// UserConference 是 User 和 UserConference 的连接表
type UserConference struct {
	UserId       uint      `gorm:"comment:成员id;column:user_id"`
	ConferenceId uint      `gorm:"column:conference_id"`
	StartTime    time.Time `gorm:"comment:成员加入会议时间;column:start_time"`
	EndTime      time.Time `gorm:"comment:成员离开会议时间;column:end_time"`
	Position     string    `gorm:"comment:成员职位;column:position"`
	Level        string    `gorm:"comment:成员级别;column:level"`
}

func (s *UserConference) TableName() string {
	return "user_conference"
}
