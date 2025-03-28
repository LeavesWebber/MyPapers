package request

// CreateRechargeOrderReq 创建充值订单请求
type CreateRechargeOrderReq struct {
	Amount     float64 `json:"amount" binding:"required,gt=0"`    // 充值金额(CNY)
	WalletAddr string  `json:"wallet_address" binding:"required"` // 钱包地址
<<<<<<< HEAD
	PayType    string  `json:"pay_type" binding:"required"`       //支付方式
=======
>>>>>>> 2f2fc5568856c67644fbada0a6a60a301279d468
}

// GetOrderStatusReq 获取订单状态请求
type GetOrderStatusReq struct {
	OrderNo string `form:"order_no" binding:"required"` // 订单号
}

<<<<<<< HEAD
// GetMPSBalanceReq 获取订单状态请求
type GetMPSBalanceReq struct {
	WalletAddr string `json:"wallet_address" binding:"required"` // 钱包地址
}

// GetMPSTransactionsReq 获取订单状态请求
type GetMPSTransactionsReq struct {
	UserId string `json:"user_id" binding:"required"` // 用户ID
=======
// SellMPSToFiatReq 卖出 MPS 换取法币请求
type SellMPSToFiatReq struct {
	Amount     float64 `json:"amount" binding:"required,gt=0"`    // 卖出金额
	WalletAddr string  `json:"wallet_address" binding:"required"` // 钱包地址
}

// BuyMPSWithFiatReq 使用法币购买虚拟币请求
type BuyMPSWithFiatReq struct {
	Amount     float64 `json:"amount" binding:"required,gt=0"`    // 购买金额
	WalletAddr string  `json:"wallet_address" binding:"required"` // 钱包地址
>>>>>>> 2f2fc5568856c67644fbada0a6a60a301279d468
}
