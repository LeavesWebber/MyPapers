package wxpay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"server/global"
	"sort"
	"strings"
)

// Config 微信支付配置
type Config struct {
	AppID              string // 公众号ID
	MchID              string // 商户号
	Key                string // API密钥
	NotifyURL          string // 回调通知地址
	TradeType          string // 交易类型
	SignType           string // 签名类型
	EthNodeURL         string // 以太坊节点URL
	MPSContractAddress string // MPS合约地址
	AdminPrivateKey    string // 管理员私钥
}

var WxConfig = &Config{
	AppID:              global.MPS_CONFIG.WxPay.AppID,
	MchID:              global.MPS_CONFIG.WxPay.MchID,
	Key:                global.MPS_CONFIG.WxPay.Key,
	NotifyURL:          global.MPS_CONFIG.WxPay.NotifyURL,
	TradeType:          global.MPS_CONFIG.WxPay.TradeType,
	SignType:           global.MPS_CONFIG.WxPay.SignType,
	EthNodeURL:         global.MPS_CONFIG.Blockchain.EthNodeURL,
	MPSContractAddress: global.MPS_CONFIG.Blockchain.MPSContractAddress,
	AdminPrivateKey:    global.MPS_CONFIG.Blockchain.AdminPrivateKey,
}

// GeneratePayParams 生成支付参数
func GeneratePayParams(orderNo string, amount float64, openID string) map[string]string {
	params := make(map[string]string)
	params["appid"] = WxConfig.AppID
	params["mch_id"] = WxConfig.MchID
	params["nonce_str"] = uuid.New().String()
	params["body"] = "MPS充值"
	params["out_trade_no"] = orderNo
	params["total_fee"] = fmt.Sprintf("%.0f", amount*100) // 转换为分
	params["spbill_create_ip"] = "127.0.0.1"
	params["notify_url"] = WxConfig.NotifyURL
	params["trade_type"] = WxConfig.TradeType
	params["openid"] = openID

	// 生成签名
	params["sign"] = generateSign(params)

	return params
}

// generateSign 生成签名
func generateSign(params map[string]string) string {
	// 按字典序排序参数
	var keys []string
	for k := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 拼接字符串
	var builder strings.Builder
	for _, k := range keys {
		if params[k] != "" {
			builder.WriteString(k)
			builder.WriteString("=")
			builder.WriteString(params[k])
			builder.WriteString("&")
		}
	}
	builder.WriteString("key=")
	builder.WriteString(WxConfig.Key)

	// MD5加密
	h := md5.New()
	h.Write([]byte(builder.String()))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// VerifySign 验证签名
func VerifySign(params map[string]string, sign string) bool {
	return generateSign(params) == sign
}
