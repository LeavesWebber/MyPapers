package initialize

import (
	"os"
	"server/global"
	"server/model/tables"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		tables.User{},
		tables.Authority{},
		tables.UserAuthority{},
		tables.BaseMenu{},
		tables.Committee{},
		tables.UserCommittee{},
		tables.Journal{},
		tables.UserJournal{},
		tables.Conference{},
		tables.UserConference{},
		tables.Paper{},
		tables.UserPaper{},
		tables.Review{},
		tables.JournalIssue{},
		tables.ConferenceIssue{},
		tables.PaperViewers{},
		tables.MPSRechargeOrder{},
		tables.MPSTransaction{},
	)
	if err != nil {
		global.MPS_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.MPS_LOG.Info("register table success")
}
