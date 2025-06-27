# 区块链配置说明文档

## 概述

本文档详细说明了 MyPapers 项目中区块链相关的配置，包括智能合约地址、网络配置、Gas 设置等。

## 配置文件结构

### 1. 主要配置文件

- `config/blockchain.go` - 区块链配置结构体定义
- `config/config.go` - 主配置文件，包含 Blockchain 结构体
- `config-example.yaml` - 配置示例文件

### 2. 配置字段说明

#### 基础网络配置

```yaml
blockchain:
  # 以太坊节点配置
  eth-node-url: "https://rpc.paperschain.io"  # 以太坊节点URL
  rpc-url: "https://rpc.paperschain.io"       # RPC URL（与eth-node-url相同）
  network-name: "paperschain"                 # 网络名称
  chain-id: 408                               # 链ID
  gas-limit: 3000000                          # Gas限制
  decimals: 18                                # 代币小数位数
```

#### 智能合约地址配置

```yaml
blockchain:
  # 智能合约配置
  mps-contract-address: "0x146329c1C8e5bA10FA991B40476CD731822e1Bb0"    # MPS代币合约地址
  mynft-contract-address: "0xCf3150B891e5176545c9EC8BfD2321bf13168848"  # MyNFT合约地址
  erc20-contract-address: "0xe49D299E6Cc29fD264A77D43Ce76dB461C9F2434"  # ERC20合约地址
  erc721-contract-address: "0xCf3150B891e5176545c9EC8BfD2321bf13168848" # ERC721合约地址
  marketplace-contract-address: "0xe699ED3A02460E3AcE9bf157C5F1098CEb10C2B0" # 市场合约地址
```

#### 管理员配置

```yaml
blockchain:
  # 管理员配置
  admin-private-key: "your-admin-private-key-here"  # 管理员私钥（去掉0x前缀）
  admin-address: "your-admin-address-here"          # 管理员地址
```

#### Gas 配置

```yaml
blockchain:
  # Gas配置
  max-fee-per-gas: 20000000000        # 最大Gas费用（EIP-1559）
  max-priority-fee-per-gas: 1000000000 # 最大优先Gas费用（EIP-1559）
  gas-price: 20000000000              # Gas价格（传统交易）
```

#### 交易配置

```yaml
blockchain:
  # 交易配置
  confirmations: 1                    # 确认区块数
  timeout: 300                        # 交易超时时间（秒）
  debug: false                        # 是否开启调试模式
```

## 配置步骤

### 1. 复制配置文件

```bash
cp config-example.yaml config.yaml
```

### 2. 修改配置文件

根据你的实际环境修改 `config.yaml` 中的配置项：

#### 本地开发环境

```yaml
blockchain:
  eth-node-url: "http://localhost:8545"
  rpc-url: "http://localhost:8545"
  network-name: "localhost"
  chain-id: 1337
  gas-limit: 3000000
  decimals: 18
  
  # 使用本地部署的合约地址
  mps-contract-address: "0x5FbDB2315678afecb367f032d93F642f64180aa3"
  mynft-contract-address: "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
  erc20-contract-address: "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
  erc721-contract-address: "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"
  marketplace-contract-address: "0xDc64a140Aa3E981100a9becA4E685f962fC1914C"
  
  admin-private-key: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
  admin-address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
  
  max-fee-per-gas: 20000000000
  max-priority-fee-per-gas: 1000000000
  gas-price: 20000000000
  
  confirmations: 1
  timeout: 300
  debug: true
```

#### 测试网络环境

```yaml
blockchain:
  eth-node-url: "https://sepolia.infura.io/v3/YOUR-PROJECT-ID"
  rpc-url: "https://sepolia.infura.io/v3/YOUR-PROJECT-ID"
  network-name: "sepolia"
  chain-id: 11155111
  gas-limit: 3000000
  decimals: 18
  
  # 使用测试网络部署的合约地址
  mps-contract-address: "0x..."
  mynft-contract-address: "0x..."
  erc20-contract-address: "0x..."
  erc721-contract-address: "0x..."
  marketplace-contract-address: "0x..."
  
  admin-private-key: "your-test-private-key"
  admin-address: "your-test-address"
  
  max-fee-per-gas: 20000000000
  max-priority-fee-per-gas: 1000000000
  gas-price: 20000000000
  
  confirmations: 3
  timeout: 300
  debug: false
```

#### 生产环境

```yaml
blockchain:
  eth-node-url: "https://mainnet.infura.io/v3/YOUR-PROJECT-ID"
  rpc-url: "https://mainnet.infura.io/v3/YOUR-PROJECT-ID"
  network-name: "mainnet"
  chain-id: 1
  gas-limit: 3000000
  decimals: 18
  
  # 使用主网部署的合约地址
  mps-contract-address: "0x..."
  mynft-contract-address: "0x..."
  erc20-contract-address: "0x..."
  erc721-contract-address: "0x..."
  marketplace-contract-address: "0x..."
  
  admin-private-key: "your-production-private-key"
  admin-address: "your-production-address"
  
  max-fee-per-gas: 20000000000
  max-priority-fee-per-gas: 1000000000
  gas-price: 20000000000
  
  confirmations: 12
  timeout: 600
  debug: false
```

## 代码中使用配置

### 1. 获取配置实例

```go
import (
    "server/global"
)

// 在需要的地方使用配置
func someFunction() {
    blockchainConfig := global.MPS_CONFIG.Blockchain
    
    // 获取链ID
    chainID := blockchainConfig.GetChainIDBigInt()
    
    // 获取Gas限制
    gasLimit := blockchainConfig.GetGasLimit()
    
    // 获取节点URL
    nodeURL := blockchainConfig.GetNodeURL()
    
    // 判断网络类型
    if blockchainConfig.IsLocalhost() {
        // 本地网络逻辑
    } else if blockchainConfig.IsTestnet() {
        // 测试网络逻辑
    } else if blockchainConfig.IsMainnet() {
        // 主网逻辑
    }
}
```

### 2. 创建以太坊客户端

```go
import (
    "github.com/ethereum/go-ethereum/ethclient"
    "server/global"
)

func createClient() (*ethclient.Client, error) {
    blockchainConfig := global.MPS_CONFIG.Blockchain
    client, err := ethclient.Dial(blockchainConfig.GetNodeURL())
    if err != nil {
        return nil, err
    }
    return client, nil
}
```

### 3. 创建合约实例

```go
import (
    "server/contracts"
    "server/global"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func createMPSContract(client *ethclient.Client) (*contracts.MPS, error) {
    blockchainConfig := global.MPS_CONFIG.Blockchain
    contractAddress := common.HexToAddress(blockchainConfig.MPSContractAddress)
    
    contract, err := contracts.NewMPS(contractAddress, client)
    if err != nil {
        return nil, err
    }
    return contract, nil
}

func createMyNFTContract(client *ethclient.Client) (*contracts.MyNFT, error) {
    blockchainConfig := global.MPS_CONFIG.Blockchain
    contractAddress := common.HexToAddress(blockchainConfig.MyNFTContractAddress)
    
    contract, err := contracts.NewMyNFT(contractAddress, client)
    if err != nil {
        return nil, err
    }
    return contract, nil
}
```

## 安全注意事项

### 1. 私钥管理

- **永远不要**将私钥硬编码在代码中
- **永远不要**将私钥提交到版本控制系统
- 使用环境变量或安全的密钥管理服务
- 在生产环境中使用硬件钱包或多重签名

### 2. 网络配置

- 确保使用正确的网络配置
- 在部署前仔细检查合约地址
- 使用适当的 Gas 限制和价格
- 监控交易确认状态

### 3. 错误处理

- 始终检查合约调用的返回值
- 实现适当的错误处理和重试机制
- 记录所有区块链交易日志

## 常见问题

### 1. 配置不生效

- 检查配置文件路径是否正确
- 确认配置文件格式正确（YAML语法）
- 重启服务以加载新配置

### 2. 合约调用失败

- 检查合约地址是否正确
- 确认网络连接正常
- 检查Gas限制是否足够
- 验证调用者权限

### 3. 交易超时

- 增加 `timeout` 配置值
- 检查网络拥堵情况
- 调整Gas价格

## 相关文件

- `config/blockchain.go` - 区块链配置结构体
- `config/config.go` - 主配置文件
- `config-example.yaml` - 配置示例
- `contracts/` - 智能合约绑定文件
- `api/mps.go` - MPS API实现
- `api/mynft.go` - MyNFT API实现 