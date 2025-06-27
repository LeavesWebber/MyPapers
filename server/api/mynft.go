package api

import (
	"math/big"
	"server/contracts"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"strconv"
	"strings"

	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MyNFTApi struct{}

// MintNFT 铸造NFT
func (n *MyNFTApi) MintNFT(c *gin.Context) {
	var req request.MyNFTMintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.MPS_LOG.Error("MyNFT MintNFT参数绑定失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证地址格式
	if !common.IsHexAddress(req.To) {
		global.MPS_LOG.Error("无效的接收地址", zap.String("address", req.To))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证权限 - 只有管理员可以铸造NFT
	userInfo, _ := utils.GetCurrentUserInfo(c)
	if !utils.IsAdmin(userInfo.ID) {
		global.MPS_LOG.Error("非管理员用户尝试铸造NFT", zap.Uint("userID", userInfo.ID))
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
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
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

	// 执行铸造NFT交易
	txn, err := nftContract.SafeMint(auth, common.HexToAddress(req.To), req.URI)
	if err != nil {
		global.MPS_LOG.Error("铸造NFT交易失败", zap.Error(err))
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
		global.MPS_LOG.Error("铸造NFT交易失败", zap.String("txHash", txn.Hash().Hex()))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 获取新铸造的tokenId（通过查询总供应量）
	totalSupply, err := nftContract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		global.MPS_LOG.Error("查询总供应量失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 新铸造的tokenId是总供应量减1
	tokenId := new(big.Int).Sub(totalSupply, big.NewInt(1))

	ResponseSuccess(c, response.MyNFTMintResponse{
		TokenId:         tokenId.Uint64(),
		To:              req.To,
		URI:             req.URI,
		TransactionHash: txn.Hash().Hex(),
		BlockNumber:     receipt.BlockNumber.Uint64(),
		GasUsed:         receipt.GasUsed,
	})
}

// BulkMintNFT 批量铸造NFT
func (n *MyNFTApi) BulkMintNFT(c *gin.Context) {
	var req request.MyNFTBulkMintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.MPS_LOG.Error("MyNFT BulkMintNFT参数绑定失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证数组长度匹配
	if len(req.Recipients) != len(req.URIs) {
		global.MPS_LOG.Error("接收者数组和URI数组长度不匹配",
			zap.Int("recipients", len(req.Recipients)),
			zap.Int("uris", len(req.URIs)))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证地址格式
	for _, addr := range req.Recipients {
		if !common.IsHexAddress(addr) {
			global.MPS_LOG.Error("无效的接收地址", zap.String("address", addr))
			ResponseError(c, CodeInvalidParam)
			return
		}
	}

	// 验证权限 - 只有管理员可以批量铸造NFT
	userInfo, _ := utils.GetCurrentUserInfo(c)
	if !utils.IsAdmin(userInfo.ID) {
		global.MPS_LOG.Error("非管理员用户尝试批量铸造NFT", zap.Uint("userID", userInfo.ID))
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
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
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

	// 批量铸造NFT
	var results []response.MyNFTMintResult
	for i, recipient := range req.Recipients {
		// 执行铸造NFT交易
		txn, err := nftContract.SafeMint(auth, common.HexToAddress(recipient), req.URIs[i])
		if err != nil {
			global.MPS_LOG.Error("批量铸造NFT交易失败", zap.Error(err), zap.String("recipient", recipient))
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
			global.MPS_LOG.Error("批量铸造NFT交易失败", zap.String("txHash", txn.Hash().Hex()))
			ResponseError(c, CodeServerBusy)
			return
		}

		// 获取新铸造的tokenId
		totalSupply, err := nftContract.TotalSupply(&bind.CallOpts{})
		if err != nil {
			global.MPS_LOG.Error("查询总供应量失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}

		tokenId := new(big.Int).Sub(totalSupply, big.NewInt(1))

		results = append(results, response.MyNFTMintResult{
			TokenId:         tokenId.Uint64(),
			To:              recipient,
			URI:             req.URIs[i],
			TransactionHash: txn.Hash().Hex(),
			BlockNumber:     receipt.BlockNumber.Uint64(),
			GasUsed:         receipt.GasUsed,
		})
	}

	ResponseSuccess(c, response.MyNFTBulkMintResponse{
		Results: results,
		Count:   len(results),
	})
}

// SetContractMetadata 设置合约元数据
func (n *MyNFTApi) SetContractMetadata(c *gin.Context) {
	var req request.MyNFTSetMetadataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.MPS_LOG.Error("MyNFT SetContractMetadata参数绑定失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证权限 - 只有管理员可以设置合约元数据
	userInfo, _ := utils.GetCurrentUserInfo(c)
	if !utils.IsAdmin(userInfo.ID) {
		global.MPS_LOG.Error("非管理员用户尝试设置合约元数据", zap.Uint("userID", userInfo.ID))
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
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
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

	// 执行设置合约元数据交易
	txn, err := nftContract.SetContractMetadata(auth, req.MetadataURI)
	if err != nil {
		global.MPS_LOG.Error("设置合约元数据交易失败", zap.Error(err))
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
		global.MPS_LOG.Error("设置合约元数据交易失败", zap.String("txHash", txn.Hash().Hex()))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MyNFTSetMetadataResponse{
		MetadataURI:     req.MetadataURI,
		TransactionHash: txn.Hash().Hex(),
		BlockNumber:     receipt.BlockNumber.Uint64(),
		GasUsed:         receipt.GasUsed,
	})
}

// UpdateRoyaltyPercentage 更新版税百分比
func (n *MyNFTApi) UpdateRoyaltyPercentage(c *gin.Context) {
	var req request.MyNFTUpdateRoyaltyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.MPS_LOG.Error("MyNFT UpdateRoyaltyPercentage参数绑定失败", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证版税百分比范围
	if req.RoyaltyPercentage > 20 {
		global.MPS_LOG.Error("版税百分比过高", zap.Uint64("percentage", req.RoyaltyPercentage))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证权限 - 只有管理员可以更新版税百分比
	userInfo, _ := utils.GetCurrentUserInfo(c)
	if !utils.IsAdmin(userInfo.ID) {
		global.MPS_LOG.Error("非管理员用户尝试更新版税百分比", zap.Uint("userID", userInfo.ID))
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
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
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

	// 执行更新版税百分比交易
	txn, err := nftContract.UpdateRoyaltyPercentage(auth, big.NewInt(int64(req.RoyaltyPercentage)))
	if err != nil {
		global.MPS_LOG.Error("更新版税百分比交易失败", zap.Error(err))
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
		global.MPS_LOG.Error("更新版税百分比交易失败", zap.String("txHash", txn.Hash().Hex()))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MyNFTUpdateRoyaltyResponse{
		RoyaltyPercentage: req.RoyaltyPercentage,
		TransactionHash:   txn.Hash().Hex(),
		BlockNumber:       receipt.BlockNumber.Uint64(),
		GasUsed:           receipt.GasUsed,
	})
}

// GetTokenURI 查询tokenURI
func (n *MyNFTApi) GetTokenURI(c *gin.Context) {
	tokenIdStr := c.Query("tokenId")
	if tokenIdStr == "" {
		global.MPS_LOG.Error("缺少tokenId参数")
		ResponseError(c, CodeInvalidParam)
		return
	}

	tokenId, err := strconv.ParseUint(tokenIdStr, 10, 64)
	if err != nil {
		global.MPS_LOG.Error("无效的tokenId", zap.String("tokenId", tokenIdStr))
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
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 查询tokenURI
	uri, err := nftContract.TokenURI(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	if err != nil {
		global.MPS_LOG.Error("查询tokenURI失败", zap.Error(err), zap.Uint64("tokenId", tokenId))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MyNFTTokenURIResponse{
		TokenId: tokenId,
		URI:     uri,
	})
}

// SupportsInterface 查询接口支持
func (n *MyNFTApi) SupportsInterface(c *gin.Context) {
	interfaceId := c.Query("interfaceId")
	if interfaceId == "" {
		global.MPS_LOG.Error("缺少interfaceId参数")
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 解析interfaceId为[4]byte
	var id [4]byte
	bytes, err := hex.DecodeString(strings.TrimPrefix(interfaceId, "0x"))
	if err != nil || len(bytes) != 4 {
		global.MPS_LOG.Error("interfaceId格式错误", zap.String("interfaceId", interfaceId))
		ResponseError(c, CodeInvalidParam)
		return
	}
	copy(id[:], bytes)

	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 查询接口支持
	supported, err := nftContract.SupportsInterface(&bind.CallOpts{}, id)
	if err != nil {
		global.MPS_LOG.Error("查询接口支持失败", zap.Error(err), zap.String("interfaceId", interfaceId))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MyNFTSupportsInterfaceResponse{
		InterfaceId: interfaceId,
		Supported:   supported,
	})
}

// GetOwner 查询合约所有者
func (n *MyNFTApi) GetOwner(c *gin.Context) {
	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 查询合约所有者
	owner, err := nftContract.Owner(&bind.CallOpts{})
	if err != nil {
		global.MPS_LOG.Error("查询合约所有者失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MyNFTOwnerResponse{
		Owner: owner.Hex(),
	})
}

// GetBalanceOf 查询用户NFT余额
func (n *MyNFTApi) GetBalanceOf(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		global.MPS_LOG.Error("缺少address参数")
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 验证地址格式
	if !common.IsHexAddress(address) {
		global.MPS_LOG.Error("无效的地址格式", zap.String("address", address))
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
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 查询用户NFT余额
	balance, err := nftContract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		global.MPS_LOG.Error("查询用户NFT余额失败", zap.Error(err), zap.String("address", address))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MyNFTBalanceResponse{
		Address: address,
		Balance: balance.Uint64(),
	})
}

// GetOwnerOf 查询NFT所有者
func (n *MyNFTApi) GetOwnerOf(c *gin.Context) {
	tokenIdStr := c.Query("tokenId")
	if tokenIdStr == "" {
		global.MPS_LOG.Error("缺少tokenId参数")
		ResponseError(c, CodeInvalidParam)
		return
	}

	tokenId, err := strconv.ParseUint(tokenIdStr, 10, 64)
	if err != nil {
		global.MPS_LOG.Error("无效的tokenId", zap.String("tokenId", tokenIdStr))
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
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 查询NFT所有者
	owner, err := nftContract.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	if err != nil {
		global.MPS_LOG.Error("查询NFT所有者失败", zap.Error(err), zap.Uint64("tokenId", tokenId))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MyNFTOwnerOfResponse{
		TokenId: tokenId,
		Owner:   owner.Hex(),
	})
}

// GetTotalSupply 查询总供应量
func (n *MyNFTApi) GetTotalSupply(c *gin.Context) {
	// 连接到区块链
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		global.MPS_LOG.Error("连接区块链失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer client.Close()

	// 创建合约实例
	nftContract, err := contracts.NewMyNFT(common.HexToAddress(global.MPS_CONFIG.Blockchain.MyNFTContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建MyNFT合约实例失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 查询总供应量
	totalSupply, err := nftContract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		global.MPS_LOG.Error("查询总供应量失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, response.MyNFTTotalSupplyResponse{
		TotalSupply: totalSupply.Uint64(),
	})
}
