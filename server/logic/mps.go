package logic

import (
	"errors"
	"server/dao/mysql"
	"server/global"
	"server/model/request"
	"server/model/tables"
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
