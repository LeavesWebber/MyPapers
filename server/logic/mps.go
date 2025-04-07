package logic

import (
	"errors"
	"math/big"
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	Alipay "server/utils/alipay"
	"time"
)

func GetTXHashList(userId string) (out []string, err error) {
	// 查询数据库
	if out, err = mysql.GetTXHashList(userId); err != nil {
		return nil, err
	}
	// 返回信息
	return out, err
}

func CreateMPSRechargeOrder(req *request.BuyMPSWithFiatReq, userID uint, orderNo string) error {
	// 获取配置
	mpsTOFiatRate := global.MPS_CONFIG.Business.MPSExchangeRate
	if mpsTOFiatRate <= 0 {
		return errors.New("invalid MPS to fiat exchange rate")
	}
	// 计算 MPS 金额
	mpsAmount := int64(req.Amount * mpsTOFiatRate)
	// 生成订单号
	// 创建订单记录
	order := &tables.MPSRechargeOrder{
		UserID:     userID,
		OrderNo:    orderNo,
		Amount:     req.Amount,
		MPSAmount:  mpsAmount,                      // 动态计算兑换金额
		Status:     tables.MPSRechargeOrderPending, // 待支付状态
		WalletAddr: req.WalletAddr,
	}
	return mysql.CreateMPSRechargeOrder(order)
}

// CreateMPSBusinessTransferOrder 创建MPS商家转账订单
// 该函数根据请求信息生成转账订单，并记录订单信息
// 参数:
//
//	req: 转账请求对象，包含转账详情
//	resp: 转账响应对象，包含转账结果
//	id: 用户ID
//	no: 订单编号
//
// 返回值:
//
//	error: 在执行过程中遇到的错误（如果有）
func CreateMPSBusinessTransferOrder(req *request.SellMPSToFiatReq, resp *response.SellMPSToFiatResp, id uint, no string) error {
	// 金额

	amount := big.NewFloat(0).SetInt64(req.MpsAmount)
	mpsTOFiatRate := big.NewFloat(global.MPS_CONFIG.Business.MPSExchangeRate)
	transferAmount, _ := amount.Quo(amount, mpsTOFiatRate).Float64()

	// 初始化状态和支付宝资金订单ID
	var status string
	var aliPayFundOrderId string
	var transDate string

	// 根据支付类型获取相应的状态和订单ID
	switch req.PayType {
	case global.MPS_CONFIG.AliPay.AliPayType:
		status = resp.AlipayResp.Status
		aliPayFundOrderId = resp.AlipayResp.PayFundOrderId
		transDate = resp.AlipayResp.TransDate
	default:
		// 默认情况下，设置状态为失败，并生成当前时间作为转账日期
		status = Alipay.FAIL
		aliPayFundOrderId = ""
		transDate = time.Now().Format("2025-03-30 21:06:33")
	}

	// 创建订单记录
	order := &tables.MPSBusinessTransferOrder{
		UserID:            id,
		OrderNo:           no,
		Identity:          req.PayeeInfo.Identity,
		IdentityType:      req.PayeeInfo.IdentityType,
		FaitAmount:        transferAmount,
		MpsAmount:         req.MpsAmount,
		Status:            status,
		AliPayFundOrderId: aliPayFundOrderId,
		TransDate:         transDate,
	}

	// 调用数据库操作，创建MPS商家转账订单记录
	return mysql.CreateMPSBusinessTransferOrder(order)
}
