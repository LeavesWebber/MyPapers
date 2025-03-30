package Alipay

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"go.uber.org/zap"
	"server/global"

	"github.com/go-pay/gopay/alipay"
)

// GeneratePayParams 生成支付参数
func GeneratePayParams(orderNo string, amount float64) gopay.BodyMap {
	params := make(gopay.BodyMap)
	params.Set("subject", "MPS充值").
		Set("out_trade_no", orderNo).
		Set("total_amount", amount).
		Set("product_code", global.MPS_CONFIG.AliPay.ProductCode).
		Set("qr_pay_mode", global.MPS_CONFIG.AliPay.QrPayMode).
		Set("qrcode_width", global.MPS_CONFIG.AliPay.QrcodeWidth)
	//todo
	//switch payType {
	// case constants.PayTypeAlipayWap:
	// 	m["product_code"] = AlipayWapProductCode
	// case constants.PayTypeAlipayApp:
	// 	m["product_code"] = AlipayAppProductCode
	// case constants.PayTypeAlipayMini:
	// 	m["buyer_id"] = buyerId
	// }
	return params
}

// initAliPayCliny 初始化支付宝客户端。
// 该函数使用支付宝的AppID、私钥以及是否是沙箱环境来创建一个新的支付宝客户端。
// 如果初始化过程中遇到错误，会记录错误日志并返回nil。
func InitAliPayClient() *alipay.Client {
	// 创建支付宝客户端
	client, err := alipay.NewClient(global.MPS_CONFIG.AliPay.AppID, global.MPS_CONFIG.AliPay.PrivateKey, global.MPS_CONFIG.AliPay.IsProd)
	if err != nil {
		global.MPS_LOG.Error("initAliPayClint error", zap.Error(err))
		return nil
	}

	// 配置支付宝客户端
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).  // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2). // 设置签名类型，不设置默认 RSA2
							SetNotifyUrl(global.MPS_CONFIG.AliPay.NotifyURL)
	// 加载证书
	CRTerr := client.SetCertSnByPath(global.MPS_CONFIG.AliPay.PublicCert, global.MPS_CONFIG.AliPay.PayRootCert, global.MPS_CONFIG.AliPay.PayPublicCert)
	if CRTerr != nil {
		global.MPS_LOG.Error("SetCertSnByPath error", zap.Error(CRTerr))
		return nil
	}

	// 返回配置好的支付宝客户端
	return client
}
func VerifySign(c *gin.Context) (gopay.BodyMap, bool) {
	var notifyReq gopay.BodyMap
	var err error
	// 解析通知请求体为 BodyMap
	notifyReq, err = alipay.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		global.MPS_LOG.Error("支付宝异步通知解析失败", zap.Error(err))
		return nil, false
	}
	// 验证签名
	ok, err := alipay.VerifySignWithCert(global.MPS_CONFIG.AliPay.PayPublicCert, notifyReq)
	if err != nil {
		global.MPS_LOG.Error("支付宝异步通知签名出错", zap.Error(err))
		return nil, false
	}

	if !ok {
		global.MPS_LOG.Error("支付宝异步通知签名失败", zap.Error(err))
		return nil, false
	}
	// 签名验证成功后的逻辑处理
	global.MPS_LOG.Info("签名验证成功")
	return notifyReq, true
}

// FastInstantTradePay 快捷即时交易支付函数
// 该函数初始化支付宝客户端，生成支付参数，并调用支付宝交易预创建接口完成支付过程
// 参数:
//
//	orderNo: 订单号
//	amount: 支付金额
//
// 返回值:
//
//	成功时返回支付宝响应参数，错误时返回错误信息
func FastInstantTradePay(orderNo string, amount float64) (string, error) {
	// 初始化支付宝客户端
	client := InitAliPayClient()
	// 生成支付参数
	params := GeneratePayParams(orderNo, amount)
	// 调用支付宝交易预创建接口
	aliRsp, err := client.TradePagePay(context.Background(), params)
	if err != nil {
		// 判断是否为业务逻辑错误
		if bizErr, ok := alipay.IsBizError(err); ok {
			global.MPS_LOG.Error("业务逻辑出错:", zap.Error(bizErr))
			// do something
			return "", err
		}
		global.MPS_LOG.Error("交易预创建出错:", zap.Error(err))
	}
	return aliRsp, err
}
