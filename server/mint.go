package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	// "os" // 如果需要用 os.Exit 控制退出码

	"server/contracts"
	"server/global"
	"server/core" // 假设此包中的 Viper() 函数负责加载配置并填充 global.MPS_CONFIG

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. 初始化配置
	// 我们假设您的项目中有一个 core.Viper() 函数，
	// 它会读取 config.yaml 文件并填充 global.MPS_CONFIG。
	// 如果您的配置初始化方式不同，请相应调整此部分。
	// 例如，如果 core.Viper() 需要配置文件路径作为参数，可能是 core.Viper("config.yaml")
	fmt.Println("正在加载配置文件 (config.yaml)...")
	global.MPS_VP = core.Viper() // 这行代码应该会触发 Viper 加载配置并填充 global.MPS_CONFIG

	// 基本检查，确保配置似乎已加载
	if global.MPS_CONFIG.Blockchain.EthNodeURL == "" {
		log.Fatal("配置加载失败或 EthNodeURL 为空。请确保 config.yaml 文件存在且配置正确，并且 core.Viper() 按预期工作。")
	}
	fmt.Println("配置文件加载成功。")

	// 2. 调用 mintToken 函数
	fmt.Println("开始执行 mintToken 函数...")
	mintToken()
	fmt.Println("mintToken 函数执行完毕。")
}

// mintToken 函数 (您提供的代码)
func mintToken() {
	// 1. 连接到以太坊节点
	//    从 config.yaml 中读取节点 URL
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL) 
	if err != nil {
		log.Fatalf("无法连接到以太坊客户端: %v, URL: %s", err, global.MPS_CONFIG.Blockchain.EthNodeURL)
	}
	fmt.Println("成功连接到以太坊节点:", global.MPS_CONFIG.Blockchain.EthNodeURL)

	// 2. 加载您的私钥 (用于发送交易)
	//    从 config.yaml 中读取管理员私钥
	//    警告：确保 admin-private-key 在 config.yaml 中已配置且安全！
	privateKeyHex := global.MPS_CONFIG.Blockchain.AdminPrivateKey 
	if privateKeyHex == "" {
		log.Fatal("管理员私钥 (admin-private-key) 未在 config.yaml 中配置")
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("加载私钥失败: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法将公钥转换为 ECDSA 类型")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("使用地址: %s\n", fromAddress.Hex())

	// 3. 获取合约地址 (这是 MPERproxy 的地址)
	//    从 config.yaml 中读取 MPS 合约地址
	contractAddressHex := global.MPS_CONFIG.Blockchain.MPSContractAddress 
	if contractAddressHex == "" {
		log.Fatal("MPS 合约地址 (mps-contract-address) 未在 config.yaml 中配置")
	}
	contractAddress := common.HexToAddress(contractAddressHex)
	fmt.Printf("MPER 代理合约地址 (来自 config.yaml: mps-contract-address): %s\n", contractAddress.Hex())

	// 4. 创建合约实例
	//    使用 abigen 生成的 contracts.NewMPER 函数
	instance, err := contracts.NewMPER(contractAddress, client)
	if err != nil {
		log.Fatalf("创建合约实例失败: %v", err)
	}
	fmt.Println("合约实例创建成功")

	// 5. 准备交易选项 (用于写操作，即需要发送交易的方法)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("获取 nonce 失败: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("获取建议 gas price 失败: %v", err)
	}

	// 获取当前网络的 Chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("获取 chain ID 失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("创建授权交易者失败: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit 
	auth.GasPrice = gasPrice

	// === 调用合约方法示例 ===

	// 示例 1: 调用只读方法 (view function) - getRecipientByHash
	testHash := "my_unique_hash_123"
	fmt.Printf("查询哈希 '%s' 对应的接收者地址...\n", testHash)
	recipientAddress, err := instance.GetRecipientByHash(&bind.CallOpts{}, testHash)
	if err != nil {
		log.Printf("调用 GetRecipientByHash 失败: %v", err)
	} else {
		if recipientAddress == (common.Address{}) || recipientAddress.Hex() == "0x0000000000000000000000000000000000000000" {
			fmt.Printf("哈希 '%s' 没有关联的接收者地址。\n", testHash)
		} else {
			fmt.Printf("哈希 '%s' 对应的接收者地址: %s\n", testHash, recipientAddress.Hex())
		}
	}

	// 示例 2: 调用写方法 (发送交易) - storeHash
	fmt.Printf("尝试为哈希 '%s' 调用 storeHash...\n", testHash)
	txStoreHash, err := instance.StoreHash(auth, testHash)
	if err != nil {
		log.Fatalf("调用 StoreHash 失败: %v", err)
	}
	fmt.Printf("StoreHash 交易已发送, 交易哈希: %s\n", txStoreHash.Hash().Hex())
	fmt.Println("等待交易被矿工打包...")

	receiptStoreHash, err := bind.WaitMined(context.Background(), client, txStoreHash)
	if err != nil {
		log.Fatalf("StoreHash 交易打包失败: %v", err)
	}
	if receiptStoreHash.Status == 1 { 
		fmt.Printf("StoreHash 交易成功! 区块号: %d\n", receiptStoreHash.BlockNumber)
		recipientAddressAfterStore, err := instance.GetRecipientByHash(&bind.CallOpts{}, testHash)
		if err != nil {
			log.Printf("再次调用 GetRecipientByHash 失败: %v", err)
		} else {
			fmt.Printf("再次查询，哈希 '%s' 对应的接收者地址: %s\n", testHash, recipientAddressAfterStore.Hex())
		}
	} else {
		fmt.Printf("StoreHash 交易失败! 区块号: %d\n", receiptStoreHash.BlockNumber)
	}
	
	// 示例 3: 调用 mint 函数 (需要更新 nonce)
	newNonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("获取新的 nonce 失败: %v", err)
	}
	auth.Nonce = big.NewInt(int64(newNonce)) // 使用最新的 nonce
	mintToAddresses := []common.Address{common.HexToAddress("0xdC0FfF81Cf674d822177F9a50fD9680201B4a1A2")} // 示例接收地址
	amountToMint := big.NewInt(1000000000000000000) // 1 MPER (假设18位小数)
	
	fmt.Printf("尝试向地址 %s Mint %s 个代币...\n", mintToAddresses[0].Hex(), amountToMint.String())
	txMint, err := instance.Mint(auth, mintToAddresses, amountToMint)
	if err != nil {
	    log.Fatalf("调用 Mint 失败: %v", err)
	}
	fmt.Printf("Mint 交易已发送, 交易哈希: %s\n", txMint.Hash().Hex())
	fmt.Println("等待 Mint 交易被矿工打包...")
	receiptMint, err := bind.WaitMined(context.Background(), client, txMint)
	if err != nil {
	    log.Fatalf("Mint 交易打包失败: %v", err)
	}
	if receiptMint.Status == 1 {
	    fmt.Printf("Mint 交易成功! 区块号: %d\n", receiptMint.BlockNumber)
	} else {
	    fmt.Printf("Mint 交易失败! 区块号: %d\n", receiptMint.BlockNumber)
	}
}