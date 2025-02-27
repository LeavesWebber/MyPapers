package service

import (
	"errors"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	"server/utils/wxpay"
	"server/global"
	"server/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"context"
	"fmt"
)

type MPSService struct{}

// CreateRechargeOrder 创建充值订单
func (s *MPSService) CreateRechargeOrder(userID uint, req *request.CreateRechargeOrderReq) (*response.CreateRechargeOrderResp, error) {
	// 生成订单号
	orderNo := wxpay.GenerateOrderNo()

	// 创建订单记录
	order := &tables.MPSRechargeOrder{
		UserID:     userID,
		OrderNo:    orderNo,
		Amount:     req.Amount,
		MPSAmount:  req.Amount, // 1:1 兑换
		Status:     0,          // 待支付
		WalletAddr: req.WalletAddr,
	}

	if err := global.MPS_DB.Create(order).Error; err != nil {
		return nil, err
	}

	// TODO: 获取用户openID
	openID := "test_open_id"

	// 生成微信支付参数
	wxParams := wxpay.GeneratePayParams(orderNo, req.Amount, openID)

	// 构造响应
	resp := &response.CreateRechargeOrderResp{
		OrderNo: orderNo,
		PayParams: response.WxPayParams{
			AppID:     wxParams["appid"],
			TimeStamp: wxParams["timestamp"],
			NonceStr:  wxParams["nonce_str"],
			Package:   "prepay_id=" + wxParams["prepay_id"],
			SignType:  "MD5",
			PaySign:   wxParams["sign"],
		},
	}

	return resp, nil
}

// GetOrderStatus 获取订单状态
func (s *MPSService) GetOrderStatus(orderNo string) (*response.OrderStatusResp, error) {
	var order tables.MPSRechargeOrder
	if err := global.MPS_DB.Where("order_no = ?", orderNo).First(&order).Error; err != nil {
		return nil, err
	}

	return &response.OrderStatusResp{
		OrderNo:   order.OrderNo,
		Status:    order.Status,
		Amount:    order.Amount,
		MPSAmount: order.MPSAmount,
	}, nil
}

// HandleWxPayNotify 处理微信支付回调
func (s *MPSService) HandleWxPayNotify(params map[string]string) error {
	// 验证签名
	if !wxpay.VerifySign(params, params["sign"]) {
		return errors.New("invalid signature")
	}

	orderNo := params["out_trade_no"]
	var order tables.MPSRechargeOrder
	if err := global.MPS_DB.Where("order_no = ?", orderNo).First(&order).Error; err != nil {
		return err
	}

	// 检查订单状态
	if order.Status != 0 {
		return nil // 订单已处理
	}

	// 开启事务
	tx := global.MPS_DB.Begin()

	// 更新订单状态
	if err := tx.Model(&order).Updates(map[string]interface{}{
		"status":      1, // 支付成功
		"wx_trade_no": params["transaction_id"],
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 记录交易
	transaction := &tables.MPSTransaction{
		UserID:      order.UserID,
		Type:        1, // 充值
		Amount:      order.MPSAmount,
		OrderNo:     orderNo,
		Description: "微信支付充值",
	}

	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 调用智能合约发放代币
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 从配置获取私钥
	privateKey, err := crypto.HexToECDSA(global.MPS_CONFIG.Blockchain.AdminPrivateKey)
	if err != nil {
		tx.Rollback()
		return err
	}

	chainID := big.NewInt(int64(global.MPS_CONFIG.Blockchain.ChainID))
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 设置交易参数
	auth.GasLimit = uint64(global.MPS_CONFIG.Blockchain.GasLimit)
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		tx.Rollback()
		return err
	}

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 发放代币
	addresses := []common.Address{common.HexToAddress(order.WalletAddr)}
	// 将 float64 转换为 big.Int，考虑 18 位小数
	amount := new(big.Int)
	amount.SetString(fmt.Sprintf("%.0f", order.MPSAmount * 1e18), 10)
	
	txn, err := mpsContract.Mint(auth, addresses, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, txn)
	if err != nil {
		tx.Rollback()
		return err
	}

	if receipt.Status == 0 {
		tx.Rollback()
		return errors.New("mint transaction failed")
	}

	return tx.Commit().Error
} 