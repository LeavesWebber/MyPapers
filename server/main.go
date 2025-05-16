// 使用情去掉注释并注释mint.go文件
package main

import (
	"server/core"
	"server/global"
	"server/initialize"
	"server/task"
	"server/utils"

	"go.uber.org/zap"
)

// @title Swagger for mps
// @version 0.0.1
// @description This is a simple family tree system
// @contact.name blingder
// @BasePath /
func main() {
	global.MPS_VP = core.Viper("config.yaml")      // 读取配置初始化
	global.MPS_LOG = core.Zap()                    // zap日志库初始化
	global.MPS_TRAN = core.InitTrans("zh")         // 翻译器初始化
	global.MPS_REDIS = core.Redis()                //redis 初始化
	global.MPS_DB = initialize.Gorm()              // gorm连接数据库
	global.MPS_RABBITMQ = core.NewRabbitMQSimple() //rabittMQ 初始化
	if err := utils.InitSnowFlake(); err != nil {  // SnowFlake初始化
		global.MPS_LOG.Error("initialize.Init() failed err:", zap.Error(err))
		return
	}
	if global.MPS_DB != nil {
		initialize.RegisterTables(global.MPS_DB) // 初始化表
		db, _ := global.MPS_DB.DB()
		defer db.Close()
	}
	task.Task()
	core.RunServer()
}
