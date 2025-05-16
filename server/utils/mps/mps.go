package Mps

import (
	"context"
	"math/big"
	"server/contracts"
	"server/global"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func Burn(walletAddr string, MpsAmount float64) (*types.Receipt, error) {
	// 调用智能合约发放代币
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	// 从配置获取私钥
	privateKey, err := crypto.HexToECDSA(global.MPS_CONFIG.Blockchain.AdminPrivateKey)
	if err != nil {
		return nil, err
	}

	chainID := big.NewInt(global.MPS_CONFIG.Blockchain.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	// 设置交易参数
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	auth.Value = big.NewInt(0)
	if err != nil {
		global.MPS_LOG.Error("设置GasPrice失败: ", zap.Error(err))
		return nil, err
	}

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建合约实例失败: ", zap.Error(err))
		return nil, err
	}

	// 发放代币
	addresses := []common.Address{common.HexToAddress(walletAddr)}
	decimals := decimal.NewFromFloat(10).Pow(decimal.NewFromInt(global.MPS_CONFIG.Blockchain.Decimals))
	decimalAmount := decimal.NewFromFloat(MpsAmount).Mul(decimals)
	mpsAmountToWei := new(big.Int)
	mpsAmountToWei.SetString(decimalAmount.String(), 10)
	txn, err := mpsContract.Mint(auth, addresses, mpsAmountToWei)
	if err != nil {
		global.MPS_LOG.Error("调用智能合约销毁代币失败: ", zap.Error(err))
		return nil, err
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, txn)
	if err != nil {
		global.MPS_LOG.Error("等待交易确认失败: ", zap.Error(err))
		return nil, err
	}
	return receipt, nil
}

func Trans(walletAddr string, mpsAmount float64) (*types.Receipt, error) {
	// 调用智能合约发放代币
	client, err := ethclient.Dial(global.MPS_CONFIG.Blockchain.EthNodeURL)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// 从配置获取私钥
	privateKey, err := crypto.HexToECDSA(global.MPS_CONFIG.Blockchain.AdminPrivateKey)
	if err != nil {
		return nil, err
	}

	chainID := big.NewInt(global.MPS_CONFIG.Blockchain.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	// 设置交易参数
	auth.GasLimit = global.MPS_CONFIG.Blockchain.GasLimit
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	auth.Value = big.NewInt(0)
	if err != nil {
		return nil, err
	}

	// 创建合约实例
	mpsContract, err := contracts.NewMPS(common.HexToAddress(global.MPS_CONFIG.Blockchain.MPSContractAddress), client)
	if err != nil {
		global.MPS_LOG.Error("创建合约实例失败: ", zap.Error(err))
		return nil, err
	}

	// 发放代币
	addresses := []common.Address{common.HexToAddress(walletAddr)}
	// 将 float64 转换为 big.Int，考虑 18 位小数
	decimals := decimal.NewFromFloat(10).Pow(decimal.NewFromInt(global.MPS_CONFIG.Blockchain.Decimals))
	decimalAmount := decimal.NewFromFloat(mpsAmount).Mul(decimals)
	mpsAmountToWei := new(big.Int)
	mpsAmountToWei.SetString(decimalAmount.String(), 10)
	txn, err := mpsContract.Mint(auth, addresses, mpsAmountToWei)
	if err != nil {
		global.MPS_LOG.Error("调用智能合约发放代币失败: ", zap.Error(err))
		return nil, err
	}

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, txn)
	if err != nil {
		global.MPS_LOG.Error("等待交易确认失败: ", zap.Error(err))
		return nil, err
	}

	return receipt, nil
}
