package Alipay

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"go.uber.org/zap"
	"math/big"
	"server/global"
	"server/model/request"
	"server/model/response"
)

// 支付宝支付状态
const (
	ALIPAY_TRADE_SUCCESS  = "TRADE_SUCCESS"  //交易支付成功
	ALIPAY_TRADE_CLOSED   = "TRADE_CLOSED"   //交易结束，不可退款
	ALIPAY_TRADE_FINISHED = "TRADE_FINISHED" //未付款交易超时关闭，或支付完成后全额退款。
	ALIPAY_WAIT_BUYER_PAY = "WAIT_BUYER_PAY" //交易创建，等待买家付款
)

// 支付宝商家转账状态
const (
	SUCCESS = "SUCCESS" //转账成功
	FAIL    = "FAIL"    //转账失败
)

// 商家转账参数   product_code
const (
	FAST_INSTANT_TRADE_PAY = "FAST_INSTANT_TRADE_PAY"
	TRANS_ACCOUNT_NO_PWD   = "TRANS_ACCOUNT_NO_PWD"  //单笔无密转账到支付宝账户
	TRANS_BANKCARD_NO_PWD  = "TRANS_BANKCARD_NO_PWD" //单笔无密转账到银行卡
)

// 商家转账参数  biz_scene
const (
	DIRECT_TRANSFER     = "DIRECT_TRANSFER"     //单笔无密转账到支付宝,B2C现金红包
	PERSONAL_COLLECTION = "PERSONAL_COLLECTION" //C2C现金红包-领红包
)

// GeneratePayParams 生成支付参数
func GeneratePayParams(orderNo string, amount float64) gopay.BodyMap {
	params := make(gopay.BodyMap)
	params.Set("subject", "MPS充值").
		Set("out_trade_no", orderNo).
		Set("total_amount", amount).
		Set("product_code", FAST_INSTANT_TRADE_PAY).
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

// InitAliPayClient 初始化支付宝客户端。
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

// FastInstantTradePay 执行快速即时交易支付。
// 该函数接收订单号和金额作为参数，初始化支付宝客户端，并调用交易预创建接口生成支付URL。
// 返回值是一个包含支付信息的结构体指针和一个错误类型，如果执行过程中遇到错误，则返回nil和错误信息。
func FastInstantTradePay(orderNo string, amount float64) (*response.BuyMPSWithFiatResp, error) {
	// 生成支付参数
	params := GeneratePayParams(orderNo, amount)
	// 初始化支付宝客户端
	aliPayClient := InitAliPayClient()
	// 调用支付宝交易预创建接口
	aliRsp, err := aliPayClient.TradePagePay(context.Background(), params)
	if err != nil {
		// 判断是否为业务逻辑错误
		if bizErr, ok := alipay.IsBizError(err); ok {
			global.MPS_LOG.Error("业务逻辑出错:", zap.Error(bizErr))
			// do something
			return nil, err
		}
		global.MPS_LOG.Error("交易预创建出错:", zap.Error(err))
	}
	// 构造返回的支付响应结构体
	buyMPSWithFiatResp := &response.BuyMPSWithFiatResp{
		OrderNo: orderNo,
		PayUrl:  aliRsp,
		AliPayParams: response.AliPayParams{
			OutTradeNo:  orderNo,
			ProductCode: FAST_INSTANT_TRADE_PAY,
			Subject:     "MPS充值",
			TotalAmount: fmt.Sprintf("%f", amount),
		},
	}
	// 返回支付响应结构体指针和nil错误，表示执行成功
	return buyMPSWithFiatResp, err
}

// FundTransUniTransfer 执行资金转账操作，将MPS货币转换为法定货币并转移到指定账户。
// 该函数接收一个SellMPSToFiatReq请求对象和一个订单号作为参数，生成转账参数并调用支付宝客户端完成转账。
// 如果转账成功，返回SellMPSToFiatResp响应对象；如果失败，返回错误。
// 参数:
//
//	req *request.SellMPSToFiatReq - 转账请求对象，包含转账所需的信息。
//	orderNo string - 转账订单号，用于唯一标识此次转账操作。
//
// 返回值:
//
//	*response.SellMPSToFiatResp - 转账响应对象，包含支付宝返回的转账结果信息。
//	error - 如果转账过程中发生错误，返回该错误。
func FundTransUniTransfer(req *request.SellMPSToFiatReq, orderNo string) (*response.SellMPSToFiatResp, error) {
	// 生成转账所需的参数，包括订单号等信息。
	bm := GenerateTransferParams(req, orderNo)
	// 初始化支付宝客户端，用于发送转账请求。
	aliPayClient := InitAliPayClient()
	// 调用支付宝客户端的资金转账接口，执行转账操作。
	resp, err := aliPayClient.FundTransUniTransfer(context.Background(), bm)
	if err != nil {
		// 判断是否为业务逻辑错误。
		if bizErr, ok := alipay.IsBizError(err); ok {
			// 记录业务逻辑错误日志。
			global.MPS_LOG.Error("业务逻辑出错:", zap.Error(bizErr))
			return nil, err
		}
		// 记录资金转账错误日志。
		global.MPS_LOG.Error("资金转账出错:", zap.Error(err))
		return nil, err
	}
	// 构建并返回转账响应对象。
	sellMPSToFiatResp := &response.SellMPSToFiatResp{
		AlipayResp: *resp.Response,
	}
	return sellMPSToFiatResp, nil
}

// GenerateTransferParams 生成转账参数
// GenerateTransferParams 根据请求和订单号生成转账参数。
// 该函数主要用于处理从MPS货币到法定货币的转账，通过计算转账金额并设置必要的参数来完成转账。
// 参数:
//
//	req - 包含转账请求信息的结构体。
//	orderNo - 转账业务的订单号。
//
// 返回值:
//
//	返回一个填充了转账参数的BodyMap对象。
func GenerateTransferParams(req *request.SellMPSToFiatReq, orderNo string) gopay.BodyMap {
	// 将请求中的金额转换为大浮点数，以确保精度。
	amount := big.NewFloat(req.MpsAmount)
	// 获取MPS到法定货币的汇率，并将其转换为大浮点数。
	mpsTOFiatRate := big.NewFloat(global.MPS_CONFIG.Business.MPSExchangeRate)
	// 计算转账金额，通过请求金额除以汇率得到。
	transferAmount, _ := amount.Quo(amount, mpsTOFiatRate).Float64()

	// 初始化转账参数的BodyMap对象。
	params := make(gopay.BodyMap)
	// 设置转账业务的订单号。
	params.Set("out_biz_no", orderNo)
	// 设置转账金额。
	params.Set("trans_amount", transferAmount)
	// 设置转账的业务场景为直接转账。
	params.Set("biz_scene", DIRECT_TRANSFER)
	// 设置转账的产品代码。
	params.Set("product_code", TRANS_ACCOUNT_NO_PWD)
	// 设置转账的标题。
	params.Set("order_title", "MPS出售")

	// 设置收款人信息。
	params.SetBodyMap("payee_info", func(bm gopay.BodyMap) {
		// 设置收款人的身份信息。
		bm.Set("identity", req.PayeeInfo.Identity)
		// 设置收款人的身份类型。
		bm.Set("identity_type", req.PayeeInfo.IdentityType)
		// 设置收款人的姓名。
		bm.Set("name", req.PayeeInfo.Name)

		// 设置收款人的银行卡扩展信息。
		bm.SetBodyMap("bankcard_ext_info", func(bm gopay.BodyMap) {
			// 设置账户类型。
			bm.Set("account_type", global.MPS_CONFIG.Business.AccountType)
			// 设置金融机构名称。
			bm.Set("inst_name", global.MPS_CONFIG.Business.InstName)
			// TODO: 根据需求添加其他必要的信息。
		})
	})
	// 返回填充了转账参数的BodyMap对象。
	return params
}
