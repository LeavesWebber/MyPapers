package main

import (
	"go.uber.org/zap"
	"server/core"
	"server/global"
	"server/initialize"
	"server/utils"
)

// @title Swagger for MPS
// @version 0.0.1
// @description This is a simple family tree system
// @contact.name blingder
// @BasePath /
func main() {
	global.MPS_VP = core.Viper()                  // 读取配置初始化
	global.MPS_LOG = core.Zap()                   // zap日志库初始化
	global.MPS_TRAN = core.InitTrans("zh")        // 翻译器初始化
	global.MPS_DB = initialize.Gorm()             // gorm连接数据库
	global.MPS_REDIS = initialize.Redis()         //redis 初始化
	if err := utils.InitSnowFlake(); err != nil { // SnowFlake初始化
		global.MPS_LOG.Error("initialize.Init() failed err:", zap.Error(err))
		return
	}

	if global.MPS_DB != nil {
		initialize.RegisterTables(global.MPS_DB) // 初始化表
		db, _ := global.MPS_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
