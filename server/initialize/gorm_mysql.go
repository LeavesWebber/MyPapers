package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/global"
	"server/initialize/internal"
)

// Gorm 初始化
func Gorm() *gorm.DB {
	m := global.MPS_CONFIG.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else { // new是自定义驱动
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
