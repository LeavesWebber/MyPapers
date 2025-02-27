package request

// CreateRechargeOrderReq 创建充值订单请求
type CreateRechargeOrderReq struct {
	Amount      float64 `json:"amount" binding:"required,gt=0"`      // 充值金额(CNY)
	WalletAddr  string  `json:"wallet_address" binding:"required"`   // 钱包地址
}

// GetOrderStatusReq 获取订单状态请求
type GetOrderStatusReq struct {
	OrderNo string `form:"order_no" binding:"required"` // 订单号
} 