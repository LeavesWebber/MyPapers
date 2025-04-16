package core

import (
	"fmt"
	"os"
	"server/core/internal"
	"server/global"
	"server/utils"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

// Zap 初始化
func Zap() (logger *zap.Logger) {
	// 判断有没有日志文件夹
	if ok, _ := utils.PathExists(global.MPS_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.MPS_CONFIG.Zap.Director)
		_ = os.Mkdir(global.MPS_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...)) // 不同级别的log记录到不同的日志文件中，cores中包含好几种级别
	if global.MPS_CONFIG.Zap.ShowLine {        // 显示行
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
