package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"server/config"
	"time"

	ut "github.com/go-playground/universal-translator"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

// rabbitMQ结构体
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}
type QueueMessage struct {
	Address     string
	MPSAmount   float64
	UUID        int64
	Description string
	OrderNo     string
}

var (
	MPS_VP       *viper.Viper
	MPS_CONFIG   config.Server
	MPS_LOG      *zap.Logger
	MPS_DB       *gorm.DB
	MPS_TRAN     ut.Translator
	MPS_REDIS    *redis.Client
	MPS_RABBITMQ *RabbitMQ
)

type MPS_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

const (
	SUPER_ADMIN   = 101 // 超级管理员
	CREATOR_ADMIN = 102 // 创建者管理员
	MEMBER        = 103 // 成员
	USER          = 104 // 普通用户
)
const (
	REDIS_PREFIX      = "mps:"
	REDIS_SMTP_PREFIX = REDIS_PREFIX + "smtp:"
)
const (
	SMTP_EXPIRED_TIME = time.Minute * 10
	SMTP_RETRY_TIME   = time.Minute
)
