package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"server/global"
)

func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.MPS_CONFIG.Redis.Addr,
		Password: global.MPS_CONFIG.Redis.Password, // 没有密码，默认值
		DB:       global.MPS_CONFIG.Redis.DB,       // 默认DB 0
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		global.MPS_LOG.Error("redis连接失败", zap.Error(err))
	}
	return rdb
}
