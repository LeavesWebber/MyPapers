package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/config"
	"server/global"
	"server/model/response"
)

type NFTApi struct{}

// 获取区块链配置
func getBlockchainConfig() *config.BlockchainConfig {
	return config.GetBlockchainConfig()
}

// MintRequest 铸造NFT请求结构
type MintRequest struct {
	To      string `json:"to" binding:"required"`
	TokenID int64  `json:"tokenId"`
	URI     string `json:"uri"`
}

// SetPriceRequest 设置价格请求结构
type SetPriceRequest struct {
	TokenID int64   `json:"tokenId" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
}

// BuyNFTRequest 购买NFT请求结构
type BuyNFTRequest struct {
	TokenID int64 `json:"tokenId" binding:"required"`
}

// SellNFTRequest 出售NFT请求结构
type SellNFTRequest struct {
	TokenID int64   `json:"tokenId" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
}

// TransferNFTRequest 转移NFT请求结构
type TransferNFTRequest struct {
	TokenID int64  `json:"tokenId" binding:"required"`
	From    string `json:"from" binding:"required"`
	To      string `json:"to" binding:"required"`
}

// 加载ABI
func loadABI(filePath string) (abi.ABI, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return abi.ABI{}, err
	}
	defer file.Close()

	var contractABI abi.ABI
	if err := json.NewDecoder(file).Decode(&contractABI); err != nil {
		return abi.ABI{}, err
	}
	return contractABI, nil
}

// 铸造NFT
func mintNFT(to common.Address, tokenID *big.Int, uri string) (string, error) {
	// 1. 获取配置
	cfg := getBlockchainConfig()

	// 2. 连接区块链
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// 3. 加载逻辑合约ABI
	contractABI, err := loadABI("hardhat/artifacts/contracts/MyNFT.sol/MyNFT.json")
	if err != nil {
		return "", err
	}

	// 4. 构造调用数据
	data, err := contractABI.Pack("safeMint", to, uri)
	if err != nil {
		return "", err
	}

	// 5. 解析私钥
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return "", err
	}

	// 6. 获取账户地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 7. 获取链信息
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	// 8. 构建交易
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// 重要：调用代理合约，但使用逻辑合约的ABI编码
	targetContract := common.HexToAddress(cfg.ProxyContract)
	tx := types.NewTransaction(nonce, targetContract, big.NewInt(0), 500000, gasPrice, data)

	// 9. 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	// 10. 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

// MintNFT 铸造NFT API
func (n *NFTApi) MintNFT(c *gin.Context) {
	var req MintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 转换地址
	toAddress := common.HexToAddress(req.To)
	txHash, err := mintNFT(toAddress, big.NewInt(req.TokenID), req.URI)

	if err != nil {
		global.MPS_LOG.Error("铸造NFT失败", zap.Error(err))
		response.FailWithMessage("铸造NFT失败: "+err.Error(), c)
		return
	}

	response.OkWithData(gin.H{
		"txHash":  txHash,
		"tokenId": req.TokenID,
		"to":      req.To,
		"uri":     req.URI,
	}, c)
}

// GetMyNFTs 获取我的NFT列表
func (n *NFTApi) GetMyNFTs(c *gin.Context) {
	// 从JWT中获取用户信息
	userInfo, exists := c.Get("userInfo")
	if !exists {
		response.FailWithMessage("用户未登录", c)
		return
	}

	// 这里应该调用数据库查询用户的NFT
	// 暂时返回模拟数据
	nfts := []gin.H{
		{
			"tokenId":   1,
			"name":      "MyNFT #1",
			"uri":       "ipfs://QmXYZ...",
			"owner":     userInfo,
			"price":     "0.1",
			"isForSale": false,
		},
		{
			"tokenId":   2,
			"name":      "MyNFT #2",
			"uri":       "ipfs://QmABC...",
			"owner":     userInfo,
			"price":     "0.2",
			"isForSale": true,
		},
	}

	response.OkWithData(nfts, c)
}

// GetNFTByTokenId 根据Token ID获取NFT信息
func (n *NFTApi) GetNFTByTokenId(c *gin.Context) {
	tokenId := c.Param("tokenId")

	// 这里应该调用智能合约查询NFT信息
	// 暂时返回模拟数据
	nft := gin.H{
		"tokenId":   tokenId,
		"name":      "MyNFT #" + tokenId,
		"uri":       "ipfs://QmXYZ...",
		"owner":     "0x123...",
		"price":     "0.1",
		"isForSale": false,
		"metadata": gin.H{
			"name":        "MyNFT #" + tokenId,
			"description": "This is a unique NFT",
			"image":       "ipfs://QmXYZ...",
			"attributes": []gin.H{
				{"trait_type": "Rarity", "value": "Common"},
				{"trait_type": "Level", "value": "1"},
			},
		},
	}

	response.OkWithData(nft, c)
}

// GetNFTMetadata 获取NFT元数据
func (n *NFTApi) GetNFTMetadata(c *gin.Context) {
	tokenId := c.Param("tokenId")

	// 这里应该从IPFS或其他存储获取元数据
	metadata := gin.H{
		"name":         "MyNFT #" + tokenId,
		"description":  "This is a unique NFT created on MyPapers platform",
		"image":        "ipfs://QmXYZ...",
		"external_url": "https://mypapers.com/nft/" + tokenId,
		"attributes": []gin.H{
			{"trait_type": "Rarity", "value": "Common"},
			{"trait_type": "Level", "value": "1"},
			{"trait_type": "Type", "value": "Paper Certificate"},
		},
	}

	response.OkWithData(metadata, c)
}

// SetNFTPrice 设置NFT价格
func (n *NFTApi) SetNFTPrice(c *gin.Context) {
	var req SetPriceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该调用智能合约设置价格
	// 暂时返回成功
	response.OkWithData(gin.H{
		"tokenId": req.TokenID,
		"price":   req.Price,
		"message": "价格设置成功",
	}, c)
}

// GetNFTMarketplace 获取NFT市场列表
func (n *NFTApi) GetNFTMarketplace(c *gin.Context) {
	// 获取查询参数
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	category := c.Query("category")
	sortBy := c.DefaultQuery("sortBy", "createdAt")
	sortOrder := c.DefaultQuery("sortOrder", "desc")

	// 这里应该查询数据库获取市场列表
	// 暂时返回模拟数据
	marketplace := gin.H{
		"page":      page,
		"limit":     limit,
		"category":  category,
		"sortBy":    sortBy,
		"sortOrder": sortOrder,
		"total":     100,
		"items": []gin.H{
			{
				"tokenId":   1,
				"name":      "MyNFT #1",
				"uri":       "ipfs://QmXYZ...",
				"owner":     "0x123...",
				"price":     "0.1",
				"isForSale": true,
				"createdAt": "2024-01-01T00:00:00Z",
			},
			{
				"tokenId":   2,
				"name":      "MyNFT #2",
				"uri":       "ipfs://QmABC...",
				"owner":     "0x456...",
				"price":     "0.2",
				"isForSale": true,
				"createdAt": "2024-01-02T00:00:00Z",
			},
		},
	}

	response.OkWithData(marketplace, c)
}

// BuyNFT 购买NFT
func (n *NFTApi) BuyNFT(c *gin.Context) {
	var req BuyNFTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该调用智能合约购买NFT
	// 暂时返回成功
	response.OkWithData(gin.H{
		"tokenId": req.TokenID,
		"message": "购买成功",
		"txHash":  "0x123...",
	}, c)
}

// SellNFT 出售NFT
func (n *NFTApi) SellNFT(c *gin.Context) {
	var req SellNFTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该调用智能合约出售NFT
	// 暂时返回成功
	response.OkWithData(gin.H{
		"tokenId": req.TokenID,
		"price":   req.Price,
		"message": "出售成功",
		"txHash":  "0x123...",
	}, c)
}

// GetNFTStats 获取NFT统计信息
func (n *NFTApi) GetNFTStats(c *gin.Context) {
	// 这里应该查询数据库和区块链获取统计信息
	// 暂时返回模拟数据
	stats := gin.H{
		"totalSupply": 1000,
		"totalOwners": 500,
		"floorPrice":  "0.1",
		"totalVolume": "1000",
		"todayVolume": "50",
		"todaySales":  25,
	}

	response.OkWithData(stats, c)
}
