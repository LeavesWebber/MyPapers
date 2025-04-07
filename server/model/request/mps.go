package request

import "github.com/go-pay/gopay/alipay/v3"

// CreateRechargeOrderReq 创建充值订单请求
type CreateRechargeOrderReq struct {
	Amount     float64 `json:"amount" binding:"required,gt=0"`    // 充值金额(CNY)
	WalletAddr string  `json:"wallet_address" binding:"required"` // 钱包地址
	PayType    string  `json:"pay_type" binding:"required"`       //支付方式

}

// GetOrderStatusReq 获取订单状态请求
type GetOrderStatusReq struct {
	OrderNo string `form:"order_no" binding:"required"` // 订单号
}

// GetMPSBalanceReq 获取订单状态请求
type GetMPSBalanceReq struct {
	WalletAddr string `json:"wallet_address" binding:"required"` // 钱包地址
}

// GetMPSTransactionsReq 获取订单状态请求
type GetMPSTransactionsReq struct {
	UserId string `json:"user_id" binding:"required"` // 用户ID
}

// SellMPSToFiatReq 卖出 MPS 换取法币请求
type SellMPSToFiatReq struct {
	MpsAmount  int64            `json:"amount" binding:"required,gt=0"`    // 卖出mps金额
	WalletAddr string           `json:"wallet_address" binding:"required"` // 钱包地址
	PayeeInfo  alipay.PayeeInfo `json:"payee_info" binding:"required"`
	PayType    string           `json:"pay_type" binding:"required"`
}

// BuyMPSWithFiatReq 使用法币购买虚拟币请求
type BuyMPSWithFiatReq struct {
	Amount     float64 `json:"amount" binding:"required,gt=0"`    // 购买金额
	WalletAddr string  `json:"wallet_address" binding:"required"` // 钱包地址
	PayType    string  `json:"pay_type" binding:"required"`
}
