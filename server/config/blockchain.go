package config

import (
	"os"
)

// BlockchainConfig 区块链配置
type BlockchainConfig struct {
	RPCURL        string // RPC节点URL
	PrivateKey    string // 私钥
	ProxyContract string // 代理合约地址
	Network       string // 网络类型
}

// GetBlockchainConfig 获取区块链配置
func GetBlockchainConfig() *BlockchainConfig {
	return &BlockchainConfig{
		RPCURL:        getEnv("RPC_URL", "http://localhost:8545"),
		PrivateKey:    getEnv("PRIVATE_KEY", ""),
		ProxyContract: getEnv("PROXY_CONTRACT", ""),
		Network:       getEnv("NETWORK", "localhost"),
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
