package core

import (
	"fmt"
	"server/global"
	"server/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

// RunServer 启动
func RunServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.MPS_CONFIG.System.Addr)
	s := initServer(address, Router)
	global.MPS_LOG.Info("server run success on ", zap.String("address", address))
	global.MPS_LOG.Error(s.ListenAndServe().Error())
}
