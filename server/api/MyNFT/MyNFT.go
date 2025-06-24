// main.go
package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// 从环境变量加载配置
var (
	rpcURL        = os.Getenv("RPC_URL")        // http://localhost:8545
	privateKeyStr = os.Getenv("PRIVATE_KEY")    // Hardhat账号私钥
	proxyAddress  = os.Getenv("PROXY_CONTRACT") // 代理合约地址
)

// MintRequest 前端请求结构
type MintRequest struct {
	To      string `json:"to"`
	TokenID int64  `json:"tokenId"`
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
func mintNFT(to common.Address, tokenID *big.Int) (string, error) {
	// 1. 连接区块链
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// 2. 加载逻辑合约ABI
	contractABI, err := loadABI("new git/hardhat/artifacts/contracts/MyNFT.sol/MyNFT.json")
	if err != nil {
		return "", err
	}

	// 3. 构造调用数据
	data, err := contractABI.Pack("mint", to, tokenID)
	if err != nil {
		return "", err
	}

	// 4. 解析私钥
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return "", err
	}

	// 5. 获取账户地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 6. 获取链信息
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	// 7. 构建交易
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	// 重要：调用代理合约，但使用逻辑合约的ABI编码
	targetContract := common.HexToAddress(proxyAddress)
	tx := types.NewTransaction(nonce, targetContract, big.NewInt(0), 500000, gasPrice, data)

	// 8. 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	// 9. 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

// API路由
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/mint", func(c *gin.Context) {
		var req MintRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 转换地址
		toAddress := common.HexToAddress(req.To)
		txHash, err := mintNFT(toAddress, big.NewInt(req.TokenID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"txHash": txHash})
	})

	return r
}

func main() {
	// 初始化环境变量
	if rpcURL == "" || privateKeyStr == "" || proxyAddress == "" {
		log.Fatal("Missing required environment variables")
	}

	r := setupRouter()
	r.Run(":8080") // 启动后端服务
}
