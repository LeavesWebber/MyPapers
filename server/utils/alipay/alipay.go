package wxpay

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"go.uber.org/zap"
	"server/global"

	"github.com/go-pay/gopay/alipay"
)

// Config 支付宝支付配置
type Config struct {
	AppID              string // 公众号ID
	PrivateKey         string // 应用私钥
	NotifyURL          string // 回调通知地址
	PublicCert         string // 应用公钥证书
	PayPublicCert      string //支付宝公钥证书
	PayRootCert        string //支付宝根证书
	TradeType          string // 交易类型
	SignType           string // 签名类型
	Charset            string // 请求使用的编码格式
	Format             string //请求格式
	QrPayMode          string // PC扫码方式
	QrcodeWidth        string // 自定义二维码宽度,qr_pay_mode=4时该参数有效
	ProductCode        string // 销售产品码,目前仅支持FAST_INSTANT_TRADE_PAY
	isSandbox          bool   //是否是正式环境
	EthNodeURL         string // 以太坊节点URL
	MPSContractAddress string // MPS合约地址
	AdminPrivateKey    string // 管理员私钥
}

var AlipayConfig = &Config{
	AppID:         global.MPS_CONFIG.AliPay.AppID,
	PrivateKey:    global.MPS_CONFIG.AliPay.PrivateKey,
	NotifyURL:     global.MPS_CONFIG.AliPay.NotifyURL,
	PublicCert:    global.MPS_CONFIG.AliPay.PublicCert,
	PayPublicCert: global.MPS_CONFIG.AliPay.PayPublicCert,
	PayRootCert:   global.MPS_CONFIG.AliPay.PayRootCert,
	//TradeType:          global.MPS_CONFIG.WxPay.TradeType,
	SignType:           global.MPS_CONFIG.AliPay.SignType,
	Charset:            global.MPS_CONFIG.AliPay.Charset,
	Format:             global.MPS_CONFIG.AliPay.Format,
	QrPayMode:          global.MPS_CONFIG.AliPay.QrPayMode,
	QrcodeWidth:        global.MPS_CONFIG.AliPay.QrcodeWidth,
	ProductCode:        global.MPS_CONFIG.AliPay.ProductCode,
	isSandbox:          global.MPS_CONFIG.AliPay.IsSandbox,
	EthNodeURL:         global.MPS_CONFIG.Blockchain.EthNodeURL,
	MPSContractAddress: global.MPS_CONFIG.Blockchain.MPSContractAddress,
	AdminPrivateKey:    global.MPS_CONFIG.Blockchain.AdminPrivateKey,
}

// GeneratePayParams 生成支付参数
func GeneratePayParams(orderNo string, amount float64, openID string) gopay.BodyMap {
	params := make(gopay.BodyMap)
	params["appid"] = AlipayConfig.AppID
	params["subject"] = "MPS充值"
	params["out_trade_no"] = orderNo
	params["total_amount"] = fmt.Sprintf("%.0f", amount*100) // 转换为分
	//todo
	//switch payType {
	// case constants.PayTypeAlipayWap:
	// 	m["product_code"] = AlipayWapProductCode
	// case constants.PayTypeAlipayApp:
	// 	m["product_code"] = AlipayAppProductCode
	// case constants.PayTypeAlipayMini:
	// 	m["buyer_id"] = buyerId
	// }
	params["product_code"] = AlipayConfig.ProductCode
	return params
}

// initAliPayCliny 初始化支付宝客户端。
// 该函数使用支付宝的AppID、私钥以及是否是沙箱环境来创建一个新的支付宝客户端。
// 如果初始化过程中遇到错误，会记录错误日志并返回nil。
func initAliPayCliny() *alipay.Client {
	// 创建支付宝客户端
	client, err := alipay.NewClient(AlipayConfig.AppID, AlipayConfig.PrivateKey, AlipayConfig.isSandbox)
	if err != nil {
		global.MPS_LOG.Error("initAliPayCliny error", zap.Error(err))
		return nil
	}

	// 配置支付宝客户端
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).  // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2). // 设置签名类型，不设置默认 RSA2
							SetNotifyUrl(AlipayConfig.NotifyURL)

	// 加载证书
	CRTerr := client.SetCertSnByPath("appPublicCert.crt", "alipayRootCert.crt", "alipayPublicCert.crt")
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
	ok, err := alipay.VerifySignWithCert("alipayPublicCert.crt", notifyReq)
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

// QRCodePay 二维码支付功能
// 该函数初始化支付宝客户端，生成支付参数，并调用支付宝接口进行交易预创建
// 参数:
//
//	orderNo - 订单编号
//	amount - 支付金额
//	openID - 用户在支付宝的唯一标识
//
// 返回值:
//
//	*alipay.TradePrecreateResponse - 支付宝交易预创建响应对象，包含二维码信息等
//	error - 错误对象，如果执行过程中发生错误
func QRCodePay(orderNo string, amount float64, openID string) (*alipay.TradePrecreateResponse, error) {
	// 初始化支付宝客户端
	client := initAliPayCliny()
	// 生成支付参数
	params := GeneratePayParams(orderNo, amount, openID)
	// 调用支付宝交易预创建接口
	aliRsp, err := client.TradePrecreate(context.Background(), params)
	if err != nil {
		// 判断是否为业务逻辑错误
		if bizErr, ok := alipay.IsBizError(err); ok {
			global.MPS_LOG.Error("业务逻辑出错:", zap.Error(bizErr))
			// do something
			return nil, err
		}
		global.MPS_LOG.Error("交易预创建出错:", zap.Error(err))
	}
	return aliRsp, err
}
