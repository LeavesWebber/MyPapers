package response

// WxPayParams 微信支付参数
type WxPayParams struct {
	AppID     string `json:"appId"`     // 公众号ID
	TimeStamp string `json:"timeStamp"` // 时间戳
	NonceStr  string `json:"nonceStr"`  // 随机字符串
	Package   string `json:"package"`   // 订单详情扩展字符串
	SignType  string `json:"signType"`  // 签名方式
	PaySign   string `json:"paySign"`   // 签名
}

// CreateRechargeOrderResp 创建充值订单响应
type CreateRechargeOrderResp struct {
	OrderNo   string      `json:"order_no"`   // 订单号
	PayParams WxPayParams `json:"pay_params"` // 支付参数
}

// OrderStatusResp 订单状态响应
type OrderStatusResp struct {
	OrderNo   string  `json:"order_no"`   // 订单号
	Status    int     `json:"status"`     // 订单状态：0-待支付 1-支付成功 2-支付失败
	Amount    float64 `json:"amount"`     // 充值金额
	MPSAmount float64 `json:"mps_amount"` // MPS数量
}

// SellMPSToFiatResp 卖出 MPS 换取法币响应
type SellMPSToFiatResp struct {
	OrderNo string `json:"order_no"` // 订单号
	Status  int    `json:"status"`   // 处理状态
}

// BuyMPSWithFiatResp 使用法币购买虚拟币响应
type BuyMPSWithFiatResp struct {
	OrderNo   string      `json:"order_no"`   // 订单号
	PayParams WxPayParams `json:"pay_params"` // 支付参数
}
