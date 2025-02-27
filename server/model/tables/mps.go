package tables

import (
	"time"
)

// MPSRechargeOrder MPS充值订单表
type MPSRechargeOrder struct {
	ID          uint      `gorm:"primarykey"`
	UserID      uint      `gorm:"not null;comment:用户ID"`
	OrderNo     string    `gorm:"type:varchar(64);not null;uniqueIndex;comment:订单号"`
	Amount      float64   `gorm:"type:decimal(10,2);not null;comment:充值金额(CNY)"`
	MPSAmount   float64   `gorm:"type:decimal(10,2);not null;comment:MPS数量"`
	Status      int       `gorm:"type:tinyint;not null;default:0;comment:订单状态：0-待支付 1-支付成功 2-支付失败"`
	WxTradeNo   string    `gorm:"type:varchar(64);comment:微信支付交易号"`
	WalletAddr  string    `gorm:"type:varchar(42);not null;comment:钱包地址"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// MPSTransaction MPS交易记录表
type MPSTransaction struct {
	ID          uint      `gorm:"primarykey"`
	UserID      uint      `gorm:"not null;comment:用户ID"`
	Type        int       `gorm:"type:tinyint;not null;comment:交易类型：1-充值 2-消费 3-奖励"`
	Amount      float64   `gorm:"type:decimal(10,2);not null;comment:MPS数量"`
	OrderNo     string    `gorm:"type:varchar(64);comment:关联订单号"`
	Description string    `gorm:"type:varchar(255);comment:交易描述"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

// TableName 设置表名
func (MPSRechargeOrder) TableName() string {
	return "mps_recharge_orders"
}

func (MPSTransaction) TableName() string {
	return "mps_transactions"
} 