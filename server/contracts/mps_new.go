// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MPSMetaData contains all meta data concerning the MPS contract.
var MPSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"HashStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"ReviewStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferWithHash\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"getRecipientByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"getReviewByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"toAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"storeHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"storeReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MPSABI is the input ABI used to generate the binding from.
// Deprecated: Use MPSMetaData.ABI instead.
var MPSABI = MPSMetaData.ABI

// MPS is an auto generated Go binding around an Ethereum contract.
type MPS struct {
	MPSCaller     // Read-only binding to the contract
	MPSTransactor // Write-only binding to the contract
	MPSFilterer   // Log filterer for contract events
}
func (m *MPS) BurnFrom(auth *bind.TransactOpts, addresses common.Address, mpsAmountToWei *big.Int) (any, error) {
	panic("unimplemented")
}

func (m *MPS) BurnFroms(auth *bind.TransactOpts, addresses common.Address, mpsAmountToWei *big.Int) (any, error) {
	panic("unimplemented")
}

// MPSCaller is an auto generated read-only Go binding around an Ethereum contract.
type MPSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MPSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MPSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MPSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MPSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MPSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MPSSession struct {
	Contract     *MPS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MPSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MPSCallerSession struct {
	Contract *MPSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MPSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MPSTransactorSession struct {
	Contract     *MPSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MPSRaw is an auto generated low-level Go binding around an Ethereum contract.
type MPSRaw struct {
	Contract *MPS // Generic contract binding to access the raw methods on
}

// MPSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MPSCallerRaw struct {
	Contract *MPSCaller // Generic read-only contract binding to access the raw methods on
}

// MPSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MPSTransactorRaw struct {
	Contract *MPSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMPS creates a new instance of MPS, bound to a specific deployed contract.
func NewMPS(address common.Address, backend bind.ContractBackend) (*MPS, error) {
	contract, err := bindMPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MPS{MPSCaller: MPSCaller{contract: contract}, MPSTransactor: MPSTransactor{contract: contract}, MPSFilterer: MPSFilterer{contract: contract}}, nil
}

// NewMPSCaller creates a new read-only instance of MPS, bound to a specific deployed contract.
func NewMPSCaller(address common.Address, caller bind.ContractCaller) (*MPSCaller, error) {
	contract, err := bindMPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MPSCaller{contract: contract}, nil
}

// NewMPSTransactor creates a new write-only instance of MPS, bound to a specific deployed contract.
func NewMPSTransactor(address common.Address, transactor bind.ContractTransactor) (*MPSTransactor, error) {
	contract, err := bindMPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MPSTransactor{contract: contract}, nil
}

// NewMPSFilterer creates a new log filterer instance of MPS, bound to a specific deployed contract.
func NewMPSFilterer(address common.Address, filterer bind.ContractFilterer) (*MPSFilterer, error) {
	contract, err := bindMPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MPSFilterer{contract: contract}, nil
}

// bindMPS binds a generic wrapper to an already deployed contract.
func bindMPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MPSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MPS *MPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MPS.Contract.MPSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MPS *MPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MPS.Contract.MPSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MPS *MPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MPS.Contract.MPSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MPS *MPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MPS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MPS *MPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MPS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MPS *MPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MPS.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MPS *MPSCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MPS.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MPS *MPSSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MPS.Contract.Allowance(&_MPS.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MPS *MPSCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MPS.Contract.Allowance(&_MPS.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MPS *MPSCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MPS.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MPS *MPSSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MPS.Contract.BalanceOf(&_MPS.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MPS *MPSCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MPS.Contract.BalanceOf(&_MPS.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MPS *MPSCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MPS.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MPS *MPSSession) Decimals() (uint8, error) {
	return _MPS.Contract.Decimals(&_MPS.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MPS *MPSCallerSession) Decimals() (uint8, error) {
	return _MPS.Contract.Decimals(&_MPS.CallOpts)
}

// GetRecipientByHash is a free data retrieval call binding the contract method 0x882514f2.
//
// Solidity: function getRecipientByHash(string hash) view returns(address)
func (_MPS *MPSCaller) GetRecipientByHash(opts *bind.CallOpts, hash string) (common.Address, error) {
	var out []interface{}
	err := _MPS.contract.Call(opts, &out, "getRecipientByHash", hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecipientByHash is a free data retrieval call binding the contract method 0x882514f2.
//
// Solidity: function getRecipientByHash(string hash) view returns(address)
func (_MPS *MPSSession) GetRecipientByHash(hash string) (common.Address, error) {
	return _MPS.Contract.GetRecipientByHash(&_MPS.CallOpts, hash)
}

// GetRecipientByHash is a free data retrieval call binding the contract method 0x882514f2.
//
// Solidity: function getRecipientByHash(string hash) view returns(address)
func (_MPS *MPSCallerSession) GetRecipientByHash(hash string) (common.Address, error) {
	return _MPS.Contract.GetRecipientByHash(&_MPS.CallOpts, hash)
}

// GetReviewByHash is a free data retrieval call binding the contract method 0xdb836881.
//
// Solidity: function getReviewByHash(string content) view returns(address)
func (_MPS *MPSCaller) GetReviewByHash(opts *bind.CallOpts, content string) (common.Address, error) {
	var out []interface{}
	err := _MPS.contract.Call(opts, &out, "getReviewByHash", content)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReviewByHash is a free data retrieval call binding the contract method 0xdb836881.
//
// Solidity: function getReviewByHash(string content) view returns(address)
func (_MPS *MPSSession) GetReviewByHash(content string) (common.Address, error) {
	return _MPS.Contract.GetReviewByHash(&_MPS.CallOpts, content)
}

// GetReviewByHash is a free data retrieval call binding the contract method 0xdb836881.
//
// Solidity: function getReviewByHash(string content) view returns(address)
func (_MPS *MPSCallerSession) GetReviewByHash(content string) (common.Address, error) {
	return _MPS.Contract.GetReviewByHash(&_MPS.CallOpts, content)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MPS *MPSCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MPS.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MPS *MPSSession) Name() (string, error) {
	return _MPS.Contract.Name(&_MPS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MPS *MPSCallerSession) Name() (string, error) {
	return _MPS.Contract.Name(&_MPS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MPS *MPSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MPS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MPS *MPSSession) Owner() (common.Address, error) {
	return _MPS.Contract.Owner(&_MPS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MPS *MPSCallerSession) Owner() (common.Address, error) {
	return _MPS.Contract.Owner(&_MPS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MPS *MPSCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MPS.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MPS *MPSSession) Symbol() (string, error) {
	return _MPS.Contract.Symbol(&_MPS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MPS *MPSCallerSession) Symbol() (string, error) {
	return _MPS.Contract.Symbol(&_MPS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MPS *MPSCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MPS.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MPS *MPSSession) TotalSupply() (*big.Int, error) {
	return _MPS.Contract.TotalSupply(&_MPS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MPS *MPSCallerSession) TotalSupply() (*big.Int, error) {
	return _MPS.Contract.TotalSupply(&_MPS.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MPS *MPSTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MPS *MPSSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.Approve(&_MPS.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MPS *MPSTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.Approve(&_MPS.TransactOpts, spender, value)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialSupply) returns()
func (_MPS *MPSTransactor) Initialize(opts *bind.TransactOpts, initialSupply *big.Int) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "initialize", initialSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialSupply) returns()
func (_MPS *MPSSession) Initialize(initialSupply *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.Initialize(&_MPS.TransactOpts, initialSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialSupply) returns()
func (_MPS *MPSTransactorSession) Initialize(initialSupply *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.Initialize(&_MPS.TransactOpts, initialSupply)
}

// Mint is a paid mutator transaction binding the contract method 0xea66696c.
//
// Solidity: function mint(address[] toAddresses, uint256 amount) returns()
func (_MPS *MPSTransactor) Mint(opts *bind.TransactOpts, toAddresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "mint", toAddresses, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xea66696c.
//
// Solidity: function mint(address[] toAddresses, uint256 amount) returns()
func (_MPS *MPSSession) Mint(toAddresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.Mint(&_MPS.TransactOpts, toAddresses, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xea66696c.
//
// Solidity: function mint(address[] toAddresses, uint256 amount) returns()
func (_MPS *MPSTransactorSession) Mint(toAddresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.Mint(&_MPS.TransactOpts, toAddresses, amount)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x2199d5cd.
//
// Solidity: function registerUser(address user) returns()
func (_MPS *MPSTransactor) RegisterUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "registerUser", user)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x2199d5cd.
//
// Solidity: function registerUser(address user) returns()
func (_MPS *MPSSession) RegisterUser(user common.Address) (*types.Transaction, error) {
	return _MPS.Contract.RegisterUser(&_MPS.TransactOpts, user)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x2199d5cd.
//
// Solidity: function registerUser(address user) returns()
func (_MPS *MPSTransactorSession) RegisterUser(user common.Address) (*types.Transaction, error) {
	return _MPS.Contract.RegisterUser(&_MPS.TransactOpts, user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MPS *MPSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MPS *MPSSession) RenounceOwnership() (*types.Transaction, error) {
	return _MPS.Contract.RenounceOwnership(&_MPS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MPS *MPSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MPS.Contract.RenounceOwnership(&_MPS.TransactOpts)
}

// StoreHash is a paid mutator transaction binding the contract method 0x71dc61cb.
//
// Solidity: function storeHash(string hash) returns()
func (_MPS *MPSTransactor) StoreHash(opts *bind.TransactOpts, hash string) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "storeHash", hash)
}

// StoreHash is a paid mutator transaction binding the contract method 0x71dc61cb.
//
// Solidity: function storeHash(string hash) returns()
func (_MPS *MPSSession) StoreHash(hash string) (*types.Transaction, error) {
	return _MPS.Contract.StoreHash(&_MPS.TransactOpts, hash)
}

// StoreHash is a paid mutator transaction binding the contract method 0x71dc61cb.
//
// Solidity: function storeHash(string hash) returns()
func (_MPS *MPSTransactorSession) StoreHash(hash string) (*types.Transaction, error) {
	return _MPS.Contract.StoreHash(&_MPS.TransactOpts, hash)
}

// StoreReview is a paid mutator transaction binding the contract method 0x8c23b081.
//
// Solidity: function storeReview(string content) returns()
func (_MPS *MPSTransactor) StoreReview(opts *bind.TransactOpts, content string) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "storeReview", content)
}

// StoreReview is a paid mutator transaction binding the contract method 0x8c23b081.
//
// Solidity: function storeReview(string content) returns()
func (_MPS *MPSSession) StoreReview(content string) (*types.Transaction, error) {
	return _MPS.Contract.StoreReview(&_MPS.TransactOpts, content)
}

// StoreReview is a paid mutator transaction binding the contract method 0x8c23b081.
//
// Solidity: function storeReview(string content) returns()
func (_MPS *MPSTransactorSession) StoreReview(content string) (*types.Transaction, error) {
	return _MPS.Contract.StoreReview(&_MPS.TransactOpts, content)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MPS *MPSTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MPS *MPSSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.Transfer(&_MPS.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MPS *MPSTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.Transfer(&_MPS.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MPS *MPSTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MPS *MPSSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.TransferFrom(&_MPS.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MPS *MPSTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPS.Contract.TransferFrom(&_MPS.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MPS *MPSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MPS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MPS *MPSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MPS.Contract.TransferOwnership(&_MPS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MPS *MPSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MPS.Contract.TransferOwnership(&_MPS.TransactOpts, newOwner)
}

// MPSApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MPS contract.
type MPSApprovalIterator struct {
	Event *MPSApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MPSApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPSApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MPSApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MPSApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPSApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPSApproval represents a Approval event raised by the MPS contract.
type MPSApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MPS *MPSFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MPSApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MPS.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MPSApprovalIterator{contract: _MPS.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MPS *MPSFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MPSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MPS.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPSApproval)
				if err := _MPS.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MPS *MPSFilterer) ParseApproval(log types.Log) (*MPSApproval, error) {
	event := new(MPSApproval)
	if err := _MPS.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPSBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the MPS contract.
type MPSBurnIterator struct {
	Event *MPSBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MPSBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPSBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MPSBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MPSBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPSBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPSBurn represents a Burn event raised by the MPS contract.
type MPSBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_MPS *MPSFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*MPSBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MPS.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &MPSBurnIterator{contract: _MPS.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_MPS *MPSFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *MPSBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MPS.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPSBurn)
				if err := _MPS.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_MPS *MPSFilterer) ParseBurn(log types.Log) (*MPSBurn, error) {
	event := new(MPSBurn)
	if err := _MPS.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPSHashStoredIterator is returned from FilterHashStored and is used to iterate over the raw logs and unpacked data for HashStored events raised by the MPS contract.
type MPSHashStoredIterator struct {
	Event *MPSHashStored // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MPSHashStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPSHashStored)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MPSHashStored)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MPSHashStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPSHashStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPSHashStored represents a HashStored event raised by the MPS contract.
type MPSHashStored struct {
	Hash   string
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHashStored is a free log retrieval operation binding the contract event 0x86fe26a4c2c4da61049e3786668239929ce2f0465c18b8c655ad4cd126df9ea9.
//
// Solidity: event HashStored(string hash, address indexed sender)
func (_MPS *MPSFilterer) FilterHashStored(opts *bind.FilterOpts, sender []common.Address) (*MPSHashStoredIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MPS.contract.FilterLogs(opts, "HashStored", senderRule)
	if err != nil {
		return nil, err
	}
	return &MPSHashStoredIterator{contract: _MPS.contract, event: "HashStored", logs: logs, sub: sub}, nil
}

// WatchHashStored is a free log subscription operation binding the contract event 0x86fe26a4c2c4da61049e3786668239929ce2f0465c18b8c655ad4cd126df9ea9.
//
// Solidity: event HashStored(string hash, address indexed sender)
func (_MPS *MPSFilterer) WatchHashStored(opts *bind.WatchOpts, sink chan<- *MPSHashStored, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MPS.contract.WatchLogs(opts, "HashStored", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPSHashStored)
				if err := _MPS.contract.UnpackLog(event, "HashStored", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHashStored is a log parse operation binding the contract event 0x86fe26a4c2c4da61049e3786668239929ce2f0465c18b8c655ad4cd126df9ea9.
//
// Solidity: event HashStored(string hash, address indexed sender)
func (_MPS *MPSFilterer) ParseHashStored(log types.Log) (*MPSHashStored, error) {
	event := new(MPSHashStored)
	if err := _MPS.contract.UnpackLog(event, "HashStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPSInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MPS contract.
type MPSInitializedIterator struct {
	Event *MPSInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MPSInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPSInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MPSInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MPSInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPSInitialized represents a Initialized event raised by the MPS contract.
type MPSInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MPS *MPSFilterer) FilterInitialized(opts *bind.FilterOpts) (*MPSInitializedIterator, error) {

	logs, sub, err := _MPS.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MPSInitializedIterator{contract: _MPS.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MPS *MPSFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MPSInitialized) (event.Subscription, error) {

	logs, sub, err := _MPS.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPSInitialized)
				if err := _MPS.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MPS *MPSFilterer) ParseInitialized(log types.Log) (*MPSInitialized, error) {
	event := new(MPSInitialized)
	if err := _MPS.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPSMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the MPS contract.
type MPSMintIterator struct {
	Event *MPSMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MPSMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPSMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MPSMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MPSMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPSMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPSMint represents a Mint event raised by the MPS contract.
type MPSMint struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_MPS *MPSFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*MPSMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPS.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &MPSMintIterator{contract: _MPS.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_MPS *MPSFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *MPSMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPS.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPSMint)
				if err := _MPS.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_MPS *MPSFilterer) ParseMint(log types.Log) (*MPSMint, error) {
	event := new(MPSMint)
	if err := _MPS.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MPS contract.
type MPSOwnershipTransferredIterator struct {
	Event *MPSOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MPSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPSOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MPSOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MPSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPSOwnershipTransferred represents a OwnershipTransferred event raised by the MPS contract.
type MPSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MPS *MPSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MPSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MPS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MPSOwnershipTransferredIterator{contract: _MPS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MPS *MPSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MPSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MPS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPSOwnershipTransferred)
				if err := _MPS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MPS *MPSFilterer) ParseOwnershipTransferred(log types.Log) (*MPSOwnershipTransferred, error) {
	event := new(MPSOwnershipTransferred)
	if err := _MPS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPSReviewStoredIterator is returned from FilterReviewStored and is used to iterate over the raw logs and unpacked data for ReviewStored events raised by the MPS contract.
type MPSReviewStoredIterator struct {
	Event *MPSReviewStored // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MPSReviewStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPSReviewStored)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MPSReviewStored)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MPSReviewStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPSReviewStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPSReviewStored represents a ReviewStored event raised by the MPS contract.
type MPSReviewStored struct {
	Content string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReviewStored is a free log retrieval operation binding the contract event 0xcb08fdcb10363722336d254fcc0aa3a2f24fe63245fdbb1fdd1a1064d8c75ea2.
//
// Solidity: event ReviewStored(string content)
func (_MPS *MPSFilterer) FilterReviewStored(opts *bind.FilterOpts) (*MPSReviewStoredIterator, error) {

	logs, sub, err := _MPS.contract.FilterLogs(opts, "ReviewStored")
	if err != nil {
		return nil, err
	}
	return &MPSReviewStoredIterator{contract: _MPS.contract, event: "ReviewStored", logs: logs, sub: sub}, nil
}

// WatchReviewStored is a free log subscription operation binding the contract event 0xcb08fdcb10363722336d254fcc0aa3a2f24fe63245fdbb1fdd1a1064d8c75ea2.
//
// Solidity: event ReviewStored(string content)
func (_MPS *MPSFilterer) WatchReviewStored(opts *bind.WatchOpts, sink chan<- *MPSReviewStored) (event.Subscription, error) {

	logs, sub, err := _MPS.contract.WatchLogs(opts, "ReviewStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPSReviewStored)
				if err := _MPS.contract.UnpackLog(event, "ReviewStored", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReviewStored is a log parse operation binding the contract event 0xcb08fdcb10363722336d254fcc0aa3a2f24fe63245fdbb1fdd1a1064d8c75ea2.
//
// Solidity: event ReviewStored(string content)
func (_MPS *MPSFilterer) ParseReviewStored(log types.Log) (*MPSReviewStored, error) {
	event := new(MPSReviewStored)
	if err := _MPS.contract.UnpackLog(event, "ReviewStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPSTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MPS contract.
type MPSTransferIterator struct {
	Event *MPSTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MPSTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPSTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MPSTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MPSTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPSTransfer represents a Transfer event raised by the MPS contract.
type MPSTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MPS *MPSFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MPSTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MPSTransferIterator{contract: _MPS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MPS *MPSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MPSTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPSTransfer)
				if err := _MPS.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MPS *MPSFilterer) ParseTransfer(log types.Log) (*MPSTransfer, error) {
	event := new(MPSTransfer)
	if err := _MPS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPSTransferWithHashIterator is returned from FilterTransferWithHash and is used to iterate over the raw logs and unpacked data for TransferWithHash events raised by the MPS contract.
type MPSTransferWithHashIterator struct {
	Event *MPSTransferWithHash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MPSTransferWithHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPSTransferWithHash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MPSTransferWithHash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MPSTransferWithHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPSTransferWithHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPSTransferWithHash represents a TransferWithHash event raised by the MPS contract.
type MPSTransferWithHash struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferWithHash is a free log retrieval operation binding the contract event 0x6b525005bc93437a219538e3a08ce19ff9b23abacb849261956e3eedd93b16db.
//
// Solidity: event TransferWithHash(address indexed from, address indexed to, uint256 value)
func (_MPS *MPSFilterer) FilterTransferWithHash(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MPSTransferWithHashIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPS.contract.FilterLogs(opts, "TransferWithHash", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MPSTransferWithHashIterator{contract: _MPS.contract, event: "TransferWithHash", logs: logs, sub: sub}, nil
}

// WatchTransferWithHash is a free log subscription operation binding the contract event 0x6b525005bc93437a219538e3a08ce19ff9b23abacb849261956e3eedd93b16db.
//
// Solidity: event TransferWithHash(address indexed from, address indexed to, uint256 value)
func (_MPS *MPSFilterer) WatchTransferWithHash(opts *bind.WatchOpts, sink chan<- *MPSTransferWithHash, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPS.contract.WatchLogs(opts, "TransferWithHash", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPSTransferWithHash)
				if err := _MPS.contract.UnpackLog(event, "TransferWithHash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferWithHash is a log parse operation binding the contract event 0x6b525005bc93437a219538e3a08ce19ff9b23abacb849261956e3eedd93b16db.
//
// Solidity: event TransferWithHash(address indexed from, address indexed to, uint256 value)
func (_MPS *MPSFilterer) ParseTransferWithHash(log types.Log) (*MPSTransferWithHash, error) {
	event := new(MPSTransferWithHash)
	if err := _MPS.contract.UnpackLog(event, "TransferWithHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
