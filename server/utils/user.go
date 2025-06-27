package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"server/global"
	"server/model/tables"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserInfo 用户信息结构
type UserInfo struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Address    string `json:"address"`
	Role       string `json:"role"`
	PrivateKey string `json:"private_key"`
}

// GetCurrentUserInfoFromDB 从数据库获取当前用户信息
func GetCurrentUserInfoFromDB(c *gin.Context) (*UserInfo, error) {
	// 从JWT token中获取用户信息
	// 这里简化处理，实际应该从JWT中解析
	userID, exists := c.Get("user_id")
	if !exists {
		return nil, nil
	}

	// 从数据库获取用户信息
	var user tables.User
	if err := global.MPS_DB.Where("id = ?", userID).First(&user).Error; err != nil {
		global.MPS_LOG.Error("获取用户信息失败", zap.Error(err))
		return nil, err
	}

	return &UserInfo{
		ID:         user.ID,
		Username:   user.Username,
		Address:    user.Address,
		Role:       user.Role,
		PrivateKey: user.PrivateKey,
	}, nil
}

// IsAdmin 判断是否为管理员
func IsAdmin(userID uint) bool {
	// 简化处理，实际应该查询数据库
	return userID == 1 || userID == 2 // 假设ID为1和2的是管理员
}

// GetUserPrivateKey 获取用户私钥
func GetUserPrivateKey(userID uint) (string, error) {
	// 从数据库获取用户私钥
	var user tables.User
	if err := global.MPS_DB.Where("id = ?", userID).First(&user).Error; err != nil {
		global.MPS_LOG.Error("获取用户私钥失败", zap.Error(err))
		return "", err
	}

	return user.PrivateKey, nil
}

// GeneratePrivateKey 生成新的私钥
func GeneratePrivateKey() (string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	return hex.EncodeToString(privateKeyBytes), nil
}

// GetAddressFromPrivateKey 从私钥获取地址
func GetAddressFromPrivateKey(privateKeyHex string) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", nil
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address.Hex(), nil
}

// CreateTransactOpts 创建交易选项
func CreateTransactOpts(privateKeyHex string, chainID int64) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	chainIDBig := big.NewInt(chainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainIDBig)
	if err != nil {
		return nil, err
	}

	// 设置Gas限制
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit

	return auth, nil
}

// ValidateAddress 验证以太坊地址格式
func ValidateAddress(address string) bool {
	return common.IsHexAddress(address)
}

// GenerateRandomHash 生成随机哈希
func GenerateRandomHash() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
