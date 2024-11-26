package global

import (
	"server/config"
	"time"

	ut "github.com/go-playground/universal-translator"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	MPS_VP     *viper.Viper
	MPS_CONFIG config.Server
	MPS_LOG    *zap.Logger
	MPS_DB     *gorm.DB
	MPS_TRAN   ut.Translator
)

type MPS_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
