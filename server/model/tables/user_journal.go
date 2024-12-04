package tables

import "time"

// UserJournal 是 User 和 Journal 的连接表
type UserJournal struct {
	UserId    uint      `gorm:"comment:成员id;column:user_id"`
	JournalId uint      `gorm:"column:journal_id"`
	StartTime time.Time `gorm:"comment:成员加入期刊时间;column:start_time"`
	EndTime   time.Time `gorm:"comment:成员离开期刊时间;column:end_time"`
	Position  string    `gorm:"comment:成员职位;column:position"`
	Level     string    `gorm:"comment:成员级别;column:level"`
}

func (s *UserJournal) TableName() string {
	return "user_journal"
}
