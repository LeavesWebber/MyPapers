package api

import (
	"math/big"
	"server/contracts"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/utils"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MPSApi struct{}

// Mint 批量铸币
func (m *MPSApi) Mint(c *gin.Context) {
	var req request.MPSMintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.MPS_LOG.Error("MPS Mint参数绑定失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证权限 - 只有管理员可以铸币
	userInfo, _ := utils.GetCurrentUserInfo(c)
	if !utils.IsAdmin(userInfo.ID) {
		global.MPS_LOG.Error("非管理员用户尝试铸币", zap.Uint("userID", userInfo.ID))
		ResponseError(c, CodeInsufficientPermissions)
		return
	}

	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MPS合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 准备管理员私钥
	privateKey, err := crypto.HexToECDSA(global.MPS_CONFIG.Blockchain.AdminPrivateKey)
	if err != nil {
		global.MPS_LOG.Error("加载管理员私钥失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 创建交易认证
	chainID := big.NewInt(global.MPS_CONFIG.Blockchain.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		global.MPS_LOG.Error("创建交易认证失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 设置交易参数
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit
	auth.GasPrice, err = client.SuggestGasPrice(c.Request.Context())
	if err != nil {
		global.MPS_LOG.Error("获取Gas价格失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 转换地址格式
	var addresses []common.Address
	for _, addr := range req.ToAddresses {
		if !common.IsHexAddress(addr) {
			global.MPS_LOG.Error("无效的以太坊地址", zap.String("address", addr))
			ResponseError(c, CodeInvalidParam)
			return
		}
		addresses = append(addresses, common.HexToAddress(addr))
	}

	// 转换金额为Wei
	amountWei := new(big.Int)
	amountWei.SetString(req.Amount, 10)

	// 执行铸币交易
	txn, err := mpsContract.Mint(auth, addresses, amountWei)
	if err != nil {
		global.MPS_LOG.Error("铸币交易失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(c.Request.Context(), client, txn)
	if err != nil {
		global.MPS_LOG.Error("等待交易确认失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	if receipt.Status == 0 {
		global.MPS_LOG.Error("铸币交易失败", zap.String("txHash", txn.Hash().Hex()))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MPSMintResponse{
		TransactionHash: txn.Hash().Hex(),
		BlockNumber:     receipt.BlockNumber.Uint64(),
		GasUsed:         receipt.GasUsed,
	})
}

// Transfer 转账
func (m *MPSApi) Transfer(c *gin.Context) {
	var req request.MPSTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.MPS_LOG.Error("MPS Transfer参数绑定失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证地址格式
	if !common.IsHexAddress(req.To) {
		global.MPS_LOG.Error("无效的接收地址", zap.String("address", req.To))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取当前用户信息
	userInfo, _ := utils.GetCurrentUserInfo(c)

	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MPS合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 获取用户私钥（这里需要从用户钱包或安全存储获取）
	// 注意：实际生产环境中应该使用更安全的方式管理用户私钥
	userPrivateKeyStr, err := utils.GetUserPrivateKey(userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("获取用户私钥失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	privateKey, err := crypto.HexToECDSA(userPrivateKeyStr)
	if err != nil {
		global.MPS_LOG.Error("用户私钥格式错误", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 创建交易认证
	chainID := big.NewInt(global.MPS_CONFIG.Blockchain.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		global.MPS_LOG.Error("创建交易认证失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 设置交易参数
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit
	auth.GasPrice, err = client.SuggestGasPrice(c.Request.Context())
	if err != nil {
		global.MPS_LOG.Error("获取Gas价格失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 转换金额为Wei
	amountWei := new(big.Int)
	amountWei.SetString(req.Amount, 10)

	// 执行转账交易
	txn, err := mpsContract.Transfer(auth, common.HexToAddress(req.To), amountWei)
	if err != nil {
		global.MPS_LOG.Error("转账交易失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(c.Request.Context(), client, txn)
	if err != nil {
		global.MPS_LOG.Error("等待交易确认失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	if receipt.Status == 0 {
		global.MPS_LOG.Error("转账交易失败", zap.String("txHash", txn.Hash().Hex()))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MPSTransferResponse{
		TransactionHash: txn.Hash().Hex(),
		BlockNumber:     receipt.BlockNumber.Uint64(),
		GasUsed:         receipt.GasUsed,
	})
}

// GetBalanceOf 查询余额
func (m *MPSApi) GetBalanceOf(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		global.MPS_LOG.Error("无效的以太坊地址", zap.String("address", address))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MPS合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 查询余额
	balance, err := mpsContract.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		global.MPS_LOG.Error("查询余额失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MPSBalanceResponse{
		Address: address,
		Balance: balance.String(),
	})
}

// StoreHash 存储哈希
func (m *MPSApi) StoreHash(c *gin.Context) {
	var req request.MPSStoreHashRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.MPS_LOG.Error("MPS StoreHash参数绑定失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取当前用户信息
	userInfo, _ := utils.GetCurrentUserInfo(c)

	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MPS合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 获取用户私钥
	userPrivateKeyStr, err := utils.GetUserPrivateKey(userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("获取用户私钥失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	privateKey, err := crypto.HexToECDSA(userPrivateKeyStr)
	if err != nil {
		global.MPS_LOG.Error("用户私钥格式错误", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 创建交易认证
	chainID := big.NewInt(global.MPS_CONFIG.Blockchain.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		global.MPS_LOG.Error("创建交易认证失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 设置交易参数
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit
	auth.GasPrice, err = client.SuggestGasPrice(c.Request.Context())
	if err != nil {
		global.MPS_LOG.Error("获取Gas价格失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 执行存储哈希交易
	txn, err := mpsContract.StoreHash(auth, req.Hash)
	if err != nil {
		global.MPS_LOG.Error("存储哈希交易失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(c.Request.Context(), client, txn)
	if err != nil {
		global.MPS_LOG.Error("等待交易确认失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	if receipt.Status == 0 {
		global.MPS_LOG.Error("存储哈希交易失败", zap.String("txHash", txn.Hash().Hex()))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MPSStoreHashResponse{
		TransactionHash: txn.Hash().Hex(),
		BlockNumber:     receipt.BlockNumber.Uint64(),
		GasUsed:         receipt.GasUsed,
	})
}

// GetRecipientByHash 查询哈希归属
func (m *MPSApi) GetRecipientByHash(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		global.MPS_LOG.Error("哈希参数为空")
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MPS合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 查询哈希归属
	recipient, err := mpsContract.GetRecipientByHash(nil, hash)
	if err != nil {
		global.MPS_LOG.Error("查询哈希归属失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MPSBalanceResponse{
		Address: hash,
		Balance: recipient.Hex(),
	})
}

// StoreReview 存储审稿内容
func (m *MPSApi) StoreReview(c *gin.Context) {
	var req request.MPSStoreReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.MPS_LOG.Error("MPS StoreReview参数绑定失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取当前用户信息
	userInfo, _ := utils.GetCurrentUserInfo(c)

	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MPS合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 获取用户私钥
	userPrivateKeyStr, err := utils.GetUserPrivateKey(userInfo.ID)
	if err != nil {
		global.MPS_LOG.Error("获取用户私钥失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	privateKey, err := crypto.HexToECDSA(userPrivateKeyStr)
	if err != nil {
		global.MPS_LOG.Error("用户私钥格式错误", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 创建交易认证
	chainID := big.NewInt(global.MPS_CONFIG.Blockchain.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		global.MPS_LOG.Error("创建交易认证失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 设置交易参数
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit
	auth.GasPrice, err = client.SuggestGasPrice(c.Request.Context())
	if err != nil {
		global.MPS_LOG.Error("获取Gas价格失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 执行存储审稿内容交易
	txn, err := mpsContract.StoreReview(auth, req.Content)
	if err != nil {
		global.MPS_LOG.Error("存储审稿内容交易失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(c.Request.Context(), client, txn)
	if err != nil {
		global.MPS_LOG.Error("等待交易确认失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	if receipt.Status == 0 {
		global.MPS_LOG.Error("存储审稿内容交易失败", zap.String("txHash", txn.Hash().Hex()))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MPSStoreReviewResponse{
		TransactionHash: txn.Hash().Hex(),
		BlockNumber:     receipt.BlockNumber.Uint64(),
		GasUsed:         receipt.GasUsed,
	})
}

// GetReviewByHash 查询审稿归属
func (m *MPSApi) GetReviewByHash(c *gin.Context) {
	content := c.Param("content")
	if content == "" {
		global.MPS_LOG.Error("审稿内容参数为空")
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MPS合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 查询审稿归属
	recipient, err := mpsContract.GetReviewByHash(nil, content)
	if err != nil {
		global.MPS_LOG.Error("查询审稿归属失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MPSBalanceResponse{
		Address: content,
		Balance: recipient.Hex(),
	})
}

// RegisterUser 注册用户
func (m *MPSApi) RegisterUser(c *gin.Context) {
	var req request.MPSRegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.MPS_LOG.Error("MPS RegisterUser参数绑定失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证权限 - 只有管理员可以注册用户
	userInfo, _ := utils.GetCurrentUserInfo(c)
	if !utils.IsAdmin(userInfo.ID) {
		global.MPS_LOG.Error("非管理员用户尝试注册用户", zap.Uint("userID", userInfo.ID))
		ResponseError(c, CodeInsufficientPermissions)
		return
	}

	// 验证地址格式
	if !common.IsHexAddress(req.UserAddress) {
		global.MPS_LOG.Error("无效的用户地址", zap.String("address", req.UserAddress))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MPS合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 准备管理员私钥
	privateKey, err := crypto.HexToECDSA(global.MPS_CONFIG.Blockchain.AdminPrivateKey)
	if err != nil {
		global.MPS_LOG.Error("加载管理员私钥失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 创建交易认证
	chainID := big.NewInt(global.MPS_CONFIG.Blockchain.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		global.MPS_LOG.Error("创建交易认证失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 设置交易参数
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit
	auth.GasPrice, err = client.SuggestGasPrice(c.Request.Context())
	if err != nil {
		global.MPS_LOG.Error("获取Gas价格失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 执行注册用户交易
	txn, err := mpsContract.RegisterUser(auth, common.HexToAddress(req.UserAddress))
	if err != nil {
		// 检查错误信息中是否包含MPS余额不足的revert
		if err.Error() != "" && strings.Contains(err.Error(), "MPS: Insufficient tokens in the contract for reward") {
			global.MPS_LOG.Error("注册用户交易失败-合约MPS余额不足", zap.Error(err))
			ResponseErrorWithMsg(c, CodeServerBusy, "合约MPS余额不足，无法完成注册奖励")
			return
		}
		global.MPS_LOG.Error("注册用户交易失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(c.Request.Context(), client, txn)
	if err != nil {
		global.MPS_LOG.Error("等待交易确认失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	if receipt.Status == 0 {
		global.MPS_LOG.Error("注册用户交易失败", zap.String("txHash", txn.Hash().Hex()))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MPSRegisterUserResponse{
		TransactionHash: txn.Hash().Hex(),
		BlockNumber:     receipt.BlockNumber.Uint64(),
		GasUsed:         receipt.GasUsed,
	})
}
