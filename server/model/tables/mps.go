package tables

import (
	"server/global"
	"time"
)

// MPSRechargeOrderStatus 定义充值订单状态
const (
	MPSRechargeOrderPending = 0 // 待支付
	MPSRechargeOrderSuccess = 1 // 支付成功
	MPSRechargeOrderFailed  = 2 // 支付失败
)

// MPSTransactionType 定义交易类型
const (
	MPSTransactionRecharge = 1 // 充值
	MPSTransactionConsume  = 2 // 消费
	MPSTransactionReward   = 3 // 奖励
	MPSTransactionSell     = 4 // 出售
)

// MPSRechargeOrder MPS充值订单表
type MPSRechargeOrder struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	UserID     int64     `gorm:"not null;comment:用户ID" json:"user_id"`
	OrderNo    string    `gorm:"type:varchar(64);not null;uniqueIndex;comment:订单号" json:"order_no"`
	Amount     float64   `gorm:"type:float;not null;comment:充值金额(CNY)" json:"amount"`
	MPSAmount  float64   `gorm:"type:float;not null;comment:MPS数量" json:"mps_amount"`
	Status     int       `gorm:"type:tinyint;not null;default:0;comment:订单状态：0-待支付 1-支付成功 2-支付失败" json:"status"`
	WxTradeNo  string    `gorm:"type:varchar(64);comment:微信支付交易号" json:"wx_trade_no"`
	AliTradeNo string    `gorm:"type:varchar(64);comment:支付宝支付交易号" json:"ali_trade_no"`
	WalletAddr string    `gorm:"type:varchar(42);not null;comment:钱包地址" json:"wallet_addr"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

// MPSTransaction MPS交易记录表
type MPSTransaction struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UserID      int64     `gorm:"type:int;not null;comment:用户ID" json:"user_id"`
	Type        int       `gorm:"type:tinyint;not null;comment:交易类型：1-充值 2-消费 3-奖励 4-出售" json:"type"`
	MpsAmount   float64   `gorm:"type:float;not null;comment:MPS数量" json:"amount"`
	TxHash      string    `gorm:"type:varchar(255);comment:交易哈希" json:"tx_hash"`
	OrderNo     string    `gorm:"type:varchar(64);comment:关联订单号" json:"order_no"`
	Description string    `gorm:"type:varchar(255);comment:交易描述" json:"description"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// MPSBusinessTransferOrder MPS商家转账记录表
type MPSBusinessTransferOrder struct {
	global.MPS_MODEL
	UserID            int64   `gorm:"type:int;not null;comment:用户ID" json:"user_id"`
	OrderNo           string  `gorm:"type:varchar(64);not null;uniqueIndex;comment:商户订单号" json:"order_no"`
	Identity          string  `gorm:"type:varchar(64); not null;comment:转账的账户:当 identity_type=ALIPAY_USER_ID 时，填写支付宝用户 UID。示例值：2088123412341234。当 identity_type=ALIPAY_LOGON_ID 时，填写支付宝登录号。示例值：186xxxxxxxx" json:"identity"` //
	IdentityType      string  `gorm:"type:varchar(64); not null;comment:转账类型支付宝的会员ID: ALIPAY_USER_ID 支付宝登录号: ALIPAY_LOGON_ID 支付宝openid: ALIPAY_OPEN_ID" json:"identity_type"`
	MpsAmount         float64 `gorm:"type:float;not null;comment:出售的MPS数量" json:"mps_amount"`
	FaitAmount        float64 `gorm:"type:float;not null;comment:出售MPS收到的金额(CNY)" json:"transfer_amount"`
	WxTradeNo         string  `gorm:"type:varchar(64);comment:微信支付交易号" json:"wx_trade_no"`
	AliPayFundOrderId string  `gorm:"type:varchar(64);not null;comment:支付宝支付资金流水号" json:"ali_pay_fund_order_id"`
	TransDate         string  `gorm:"type:varchar(64);not null;comment:订单支付时间，格式为yyyy-MM-dd HH:mm:ss" json:"trans_date"`
	Status            string  `gorm:"type:varchar(64);not null;comment:转账单据状态。 SUCCESS（该笔转账交易成功）：成功； FAIL：失败" json:"status"`
}

const (
	TableMPSRechargeOrders = "mps_recharge_orders"
	TableMPSTransactions   = "mps_transactions"
)

// TableName 设置表名
func (MPSRechargeOrder) TableName() string {
	return TableMPSRechargeOrders
}

func (MPSTransaction) TableName() string {
	return TableMPSTransactions
}
