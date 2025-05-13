package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"server/contracts"
	"server/core"
	"server/global"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. 初始化配置
	fmt.Println("正在加载配置文件 (config.yaml)...")
	global.MPS_VP = core.Viper()

	if global.MPS_CONFIG.Blockchain.EthNodeURL == "" {
		log.Fatal("配置加载失败或 EthNodeURL 为空")
	}
	fmt.Println("配置文件加载成功。")

	// 2. 调用测试函数
	fmt.Println("开始执行MPS合约测试...")
	testMPSContract()
	fmt.Println("测试执行完毕。")
}

func testMPSContract() {
	// 1. 连接到以太坊节点
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		log.Fatalf("无法连接到以太坊客户端: %v", err)
	}
	fmt.Println("成功连接到以太坊节点")

	// 2. 加载管理员私钥
	privateKeyHex := global.MPS_CONFIG.Blockchain.AdminPrivateKey
	if privateKeyHex == "" {
		log.Fatal("管理员私钥未配置")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("加载私钥失败: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法将公钥转换为ECDSA类型")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("管理员地址: %s\n", fromAddress.Hex())

	// 3. 获取合约地址
	contractAddressHex := global.MPS_CONFIG.Blockchain.MPSContractAddress
	if contractAddressHex == "" {
		log.Fatal("MPS合约地址未配置")
	}
	contractAddress := common.HexToAddress(contractAddressHex)
	fmt.Printf("MPS合约地址: %s\n", contractAddress.Hex())

	// 4. 创建合约实例
	instance, err := contracts.NewMPS(contractAddress, client)
	if err != nil {
		log.Fatalf("创建合约实例失败: %v", err)
	}
	fmt.Println("合约实例创建成功")

	// 5. 准备交易授权
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("获取chain ID失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("创建授权交易者失败: %v", err)
	}

	// === 测试合约功能 ===

	// 测试1: 查询代币名称和符号
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Printf("查询代币名称失败: %v", err)
	} else {
		fmt.Printf("代币名称: %s\n", name)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Printf("查询代币符号失败: %v", err)
	} else {
		fmt.Printf("代币符号: %s\n", symbol)
	}

	// 测试2: 存储哈希
	testHash := "test_hash_124"
	fmt.Printf("尝试存储哈希 '%s'...\n", testHash)
	txStoreHash, err := instance.StoreHash(auth, testHash)
	if err != nil {
		log.Printf("存储哈希失败: %v", err)
	} else {
		fmt.Printf("存储哈希交易已发送, 交易哈希: %s\n", txStoreHash.Hash().Hex())
	}

	// 测试3: 查询哈希对应的地址
	recipient, err := instance.GetRecipientByHash(&bind.CallOpts{}, testHash)
	if err != nil {
		log.Printf("查询哈希对应地址失败: %v", err)
	} else {
		fmt.Printf("哈希 '%s' 对应的地址: %s\n", testHash, recipient.Hex())
	}

	// 测试4: Mint代币
	testAddress := common.HexToAddress("0xdC0FfF81Cf674d822177F9a50fD9680201B4a1A2")
	amount := big.NewInt(1000000000000000000) // 1 MPS (假设18位小数)

	fmt.Printf("尝试向地址 %s Mint %s 个代币...\n", testAddress.Hex(), amount.String())
	txMint, err := instance.Mint(auth, []common.Address{testAddress}, amount)
	if err != nil {
		log.Printf("Mint代币失败: %v", err)
	} else {
		fmt.Printf("Mint交易已发送, 交易哈希: %s\n", txMint.Hash().Hex())
	}

	// 测试5: 查询余额
	balance, err := instance.BalanceOf(&bind.CallOpts{}, testAddress)
	if err != nil {
		log.Printf("查询余额失败: %v", err)
	} else {
		fmt.Printf("地址 %s 的余额: %s\n", testAddress.Hex(), balance.String())
	}
}
