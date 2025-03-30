package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"math/big"
	"server/contracts"
	"server/global"
	"server/logic"
	"server/model/request"
	"server/model/response"
	"server/model/tables"
	Alipay "server/utils/alipay"
	"server/utils/wxpay"
	"strconv"
	"sync"
	"time"
)

type MPSService struct{}

const (
	AliPay_TRADE_SUCCESS  = "TRADE_SUCCESS"  //交易支付成功
	AliPay_TRADE_CLOSED   = "TRADE_CLOSED"   //交易结束，不可退款
	AliPay_TRADE_FINISHED = "TRADE_FINISHED" //未付款交易超时关闭，或支付完成后全额退款。
	AliPay_WAIT_BUYER_PAY = "WAIT_BUYER_PAY" //交易创建，等待买家付款
)

// CreateRechargeOrder  根据类型创建充值订单
func (s *MPSService) CreateRechargeOrder(userID uint, req *request.CreateRechargeOrderReq) (*response.CreateRechargeOrderResp, error) {
	// 生成订单号
	orderNo := GenerateOrderNo()
	mpstoFiatRate := global.MPS_CONFIG.Business.MPSExchangeRate
	mpsAmount := int64(mpstoFiatRate * req.Amount)
	// 创建订单记录
	order := &tables.MPSRechargeOrder{
		UserID:     userID,
		OrderNo:    orderNo,
		Amount:     req.Amount,
		MPSAmount:  mpsAmount, // 1:1 兑换
		Status:     0,         // 待支付
		WalletAddr: req.WalletAddr,
	}

	if err := global.MPS_DB.Create(order).Error; err != nil {
		return nil, err
	}
	// TODO: 获取用户openID
	openID := "test_open_id"

	var resp *response.CreateRechargeOrderResp
	switch req.PayType {
	case global.MPS_CONFIG.AliPay.AliPayType:
		aliParams := Alipay.GeneratePayParams(orderNo, req.Amount)
		resp = &response.CreateRechargeOrderResp{
			OrderNo: orderNo,
			AliPayParams: response.AliPayParams{
				Subject:     aliParams.Get("subject"),
				OutTradeNo:  aliParams.Get("out_trade_no"),
				TotalAmount: aliParams.Get("total_amount"),
				ProductCode: aliParams.Get("product_code"),
			},
		}
	case global.MPS_CONFIG.WxPay.WxPayType:
		wxParams := wxpay.GeneratePayParams(orderNo, req.Amount, openID)
		resp = &response.CreateRechargeOrderResp{
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
	amount.SetString(fmt.Sprintf("%.0f", order.MPSAmount*1e18), 10)

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
	// 记录交易
	transaction := &tables.MPSTransaction{
		UserID:      order.UserID,
		Type:        1, // 充值
		Amount:      order.MPSAmount,
		OrderNo:     orderNo,
		TxHash:      receipt.TxHash.Hex(),
		Description: "微信支付充值",
	}
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (s *MPSService) HandleAliPayNotify(c *gin.Context) error {
	// 验证签名

	global.MPS_LOG.Info("支付宝异步通知验证开始")
	notifyReq, ok := Alipay.VerifySign(c)
	if !ok {
		return errors.New("非法签名")
	}
	status := notifyReq.Get("trade_status")
	var statusID int
	switch status {
	case AliPay_TRADE_SUCCESS:
		statusID = 1
	default:
		return errors.New("支付宝交易状态非法")
	}
	orderNo := notifyReq.Get("out_trade_no")
	totalAmountStr := notifyReq.Get("total_amount")
	totalAmount, err := strconv.ParseFloat(totalAmountStr, 64)
	if err != nil {
		return err
	}
	appId := notifyReq.Get("app_id")
	var order tables.MPSRechargeOrder
	if err := global.MPS_DB.Where("order_no = ?", orderNo).First(&order).Error; err != nil {
		return err
	}
	if totalAmount != order.Amount || appId != global.MPS_CONFIG.AliPay.AppID {
		return errors.New("支付宝交易信息非法")
	}
	// 检查订单状态
	if order.Status != 0 {
		return nil // 订单已处理
	}

	// 更新订单状态
	if err := global.MPS_DB.Model(&order).Updates(map[string]interface{}{
		"status":       statusID, // 支付成功
		"ali_trade_no": notifyReq.Get("trade_no"),
	}).Error; err != nil {
		return err
	}
	// 开启事务
	tx := global.MPS_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			global.MPS_LOG.Error("事务回滚", zap.Any("reason", r))
		}
	}()
	// 调用智能合约发放代币
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer client.Close()
	// 从配置获取私钥
	privateKey, err := crypto.HexToECDSA(global.MPS_CONFIG.Blockchain.AdminPrivateKey)
	if err != nil {
		tx.Rollback()
		return err
	}

	chainID := big.NewInt(global.MPS_CONFIG.Blockchain.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 设置交易参数
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	auth.Value = big.NewInt(0)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建合约实例失败: ", zap.Error(err))
		tx.Rollback()
		return err
	}

	// 发放代币
	addresses := []common.Address{common.HexToAddress(order.WalletAddr)}
	// 将 float64 转换为 big.Int，考虑 18 位小数
	amount := new(big.Int)
	amount.SetInt64(order.MPSAmount)

	txn, err := mpsContract.Mint(auth, addresses, amount)
	if err != nil {
		global.MPS_LOG.Error("调用智能合约发放代币失败: ", zap.Error(err))
		tx.Rollback()
		return err
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, txn)
	if err != nil {
		global.MPS_LOG.Error("等待交易确认失败: ", zap.Error(err))
		tx.Rollback()
		return err
	}

	if receipt.Status == 0 {
		global.MPS_LOG.Error("mint transaction failed", zap.Error(err))
		tx.Rollback()
		return errors.New("mint transaction failed")
	}
	// 记录交易
	transaction := &tables.MPSTransaction{
		UserID:      order.UserID,
		Type:        1, // 充值
		Amount:      order.MPSAmount,
		OrderNo:     orderNo,
		TxHash:      receipt.TxHash.Hex(),
		Description: "支付宝充值",
	}
	if err := tx.Create(transaction).Error; err != nil {
		global.MPS_LOG.Error("记录交易失败: ", zap.Error(err))
		tx.Rollback()
		return err
	}
	global.MPS_LOG.Info("支付宝异步通知验证结束")
	return tx.Commit().Error
}

func (s *MPSService) GetMPSBalance(address common.Address) (*response.MPSBalanceResp, error) {
	var resp *response.MPSBalanceResp
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接以太坊节点失败: ", zap.Error(err))
	}
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建合约实例失败: ", zap.Error(err))
		return nil, err
	}
	balance, err := mpsContract.BalanceOf(nil, address)
	if err != nil {
		global.MPS_LOG.Error("查询余额失败: ", zap.Error(err))
	}
	resp = &response.MPSBalanceResp{
		Balance: balance.String(),
	}
	return resp, nil
}

func (s *MPSService) GetMPSTransactions(id string) (interface{}, error) {
	// 初始化响应对象
	resp := &response.TxList{TxList: make([]types.Transaction, 0)}

	// 创建以太坊客户端并确保资源释放
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接以太坊节点失败", zap.String("id", id), zap.Error(err))
		return nil, fmt.Errorf("连接以太坊节点失败: %w", err)
	}
	defer client.Close() // 确保客户端关闭

	// 获取交易哈希列表
	txHashList, err := logic.GetTXHashList(id)
	if err != nil {
		global.MPS_LOG.Error("查询交易哈希失败", zap.String("id", id), zap.Error(err))
		return nil, fmt.Errorf("查询交易哈希失败: %w", err)
	}

	// 如果交易哈希列表为空，直接返回
	if len(txHashList) == 0 {
		global.MPS_LOG.Info("交易哈希列表为空", zap.String("id", id))
		return resp, nil
	}

	// 并发查询交易
	var wg sync.WaitGroup
	errChan := make(chan error, len(txHashList))
	resultChan := make(chan *types.Transaction, len(txHashList))

	// 引入 Goroutine 限制机制
	sem := make(chan struct{}, 10) // 最大并发数为 10
	for _, hash := range txHashList {
		wg.Add(1)
		go func(hash string) {
			defer wg.Done()
			sem <- struct{}{}        // 占用一个 Goroutine 配额
			defer func() { <-sem }() // 释放配额

			txHash := common.HexToHash(hash)
			tx, _, err := client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				global.MPS_LOG.Error("查询交易失败", zap.String("id", id), zap.String("txHash", hash), zap.Error(err))
				errChan <- fmt.Errorf("查询交易失败: %w", err)
				return
			}
			if tx == nil {
				global.MPS_LOG.Warn("交易对象为空", zap.String("id", id), zap.String("txHash", hash))
				return
			}
			resultChan <- tx
		}(hash)
	}

	// 等待所有 Goroutine 完成
	go func() {
		wg.Wait()
		close(resultChan)
		close(errChan)
	}()

	// 收集结果和错误
	for tx := range resultChan {
		resp.TxList = append(resp.TxList, *tx)
	}

	// 检查是否有错误发生
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("发生错误: %v", errs)
	}
	return resp, nil
}

// SellMPSToFiat 卖出 MPS 换取法币
func (s *MPSService) SellMPSToFiat(userID uint, req *request.SellMPSToFiatReq) (*response.SellMPSToFiatResp, error) {
	// 生成订单号
	orderNo := GenerateOrderNo()
	mpstoFiatRate := global.MPS_CONFIG.Business.MPSExchangeRate
	mpsAmount := int64(mpstoFiatRate * req.Amount)
	// 创建订单记录
	order := &tables.MPSRechargeOrder{
		UserID:     userID,
		OrderNo:    orderNo,
		Amount:     req.Amount,
		MPSAmount:  mpsAmount, // 1:1 兑换
		Status:     0,         // 待处理
		WalletAddr: req.WalletAddr,
	}

	if err := global.MPS_DB.Create(order).Error; err != nil {
		global.MPS_LOG.Error("Failed to create order", zap.Error(err))
		return nil, err
	}

	// 调用智能合约销毁代币
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("Failed to create Ethereum client", zap.Error(err))
		return nil, err
	}

	// 从配置获取私钥
	privateKey, err := crypto.HexToECDSA(global.MPS_CONFIG.Blockchain.AdminPrivateKey)
	if err != nil {
		global.MPS_LOG.Error("Failed to parse private key", zap.Error(err))
		return nil, err
	}

	chainID := big.NewInt(int64(global.MPS_CONFIG.Blockchain.ChainID))
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		global.MPS_LOG.Error("Failed to create transactor", zap.Error(err))
		return nil, err
	}

	// 设置交易参数
	auth.GasLimit = uint64(global.MPS_CONFIG.Blockchain.GasLimit)
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		global.MPS_LOG.Error("Failed to suggest gas price", zap.Error(err))
		return nil, err
	}

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("Failed to create contract instance", zap.Error(err))
		return nil, err
	}

	// 销毁代币
	addresses := []common.Address{common.HexToAddress(req.WalletAddr)}
	// 将 float64 转换为 big.Int，考虑 18 位小数
	amount := new(big.Int)
	amount.SetString(fmt.Sprintf("%.0f", req.Amount*1e18), 10)

	txn, err := mpsContract.BurnFrom(auth, addresses[0], amount)
	if err != nil {
		global.MPS_LOG.Error("Failed to burn tokens", zap.Error(err))
		return nil, err
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, txn)
	if err != nil {
		global.MPS_LOG.Error("Failed to wait for transaction mining", zap.Error(err))
		return nil, err
	}

	if receipt.Status == 0 {
		global.MPS_LOG.Error("Burn transaction failed", zap.Any("receipt", receipt))
		return nil, errors.New("burn transaction failed")
	}

	// 更新订单状态
	if err := global.MPS_DB.Model(&order).Update("status", 1).Error; err != nil {
		global.MPS_LOG.Error("Failed to update order status", zap.Error(err))
		return nil, err
	}

	// 构造响应
	resp := &response.SellMPSToFiatResp{
		OrderNo: orderNo,
		Status:  1, // 处理成功
	}

	return resp, nil
}

// BuyMPSWithFiat 使用法币购买虚拟币
func (s *MPSService) BuyMPSWithFiat(userID uint, req *request.BuyMPSWithFiatReq) (*response.BuyMPSWithFiatResp, error) {
	// 生成订单号
	orderNo := GenerateOrderNo()
	mpstoFiatRate := global.MPS_CONFIG.Business.MPSExchangeRate
	mpsAmount := int64(mpstoFiatRate * req.Amount)
	// 创建订单记录
	order := &tables.MPSRechargeOrder{
		UserID:     userID,
		OrderNo:    orderNo,
		Amount:     req.Amount,
		MPSAmount:  mpsAmount, // 1:1 兑换
		Status:     0,         // 待支付
		WalletAddr: req.WalletAddr,
	}

	if err := global.MPS_DB.Create(order).Error; err != nil {
		global.MPS_LOG.Error("Failed to create order", zap.Error(err))
		return nil, err
	}

	// TODO: 获取用户openID
	openID := "test_open_id"

	// 生成微信支付参数
	wxParams := wxpay.GeneratePayParams(orderNo, req.Amount, openID)

	// 构造响应
	resp := &response.BuyMPSWithFiatResp{
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

// GenerateOrderNo 生成订单号
func GenerateOrderNo() string {
	return fmt.Sprintf("MPS%s%s",
		time.Now().Format("20060102150405"),
		uuid.New().String()[:8])
}
func (s *MPSService) Pay(userID uint, req *request.BuyMPSWithFiatReq) (interface{}, error) {
	orderNo := GenerateOrderNo()

	// 处理支付逻辑
	var resp interface{}
	var payErr error
	switch req.PayType {
	case global.MPS_CONFIG.AliPay.AliPayType:
		resp, payErr = Alipay.FastInstantTradePay(orderNo, req.Amount)
		if payErr != nil {
			global.MPS_LOG.Error("Alipay QRCodePay failed for order "+orderNo, zap.Error(payErr))
		}
	default:
		payErr = errors.New("unsupported payment type")
		global.MPS_LOG.Error("Alipay QRCodePay failed for order "+orderNo, zap.Error(payErr))
	}
	// 返回结果
	if payErr != nil {
		return nil, payErr
	}
	// 插入数据库
	err := logic.CreateMPSRechargeOrder(req, userID, orderNo)
	if err != nil {
		global.MPS_LOG.Error("Failed to CreateMPSRechargeOrder to DataBase", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
