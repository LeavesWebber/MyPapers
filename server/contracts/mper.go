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

// MPERMetaData contains all meta data concerning the MPER contract.
var MPERMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"HashStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"ReviewStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferWithHash\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"getRecipientByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"getReviewByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"toAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"storeHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"storeReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MPERABI is the input ABI used to generate the binding from.
// Deprecated: Use MPERMetaData.ABI instead.
var MPERABI = MPERMetaData.ABI

// MPER is an auto generated Go binding around an Ethereum contract.
type MPER struct {
	MPERCaller     // Read-only binding to the contract
	MPERTransactor // Write-only binding to the contract
	MPERFilterer   // Log filterer for contract events
}

// MPERCaller is an auto generated read-only Go binding around an Ethereum contract.
type MPERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MPERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MPERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MPERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MPERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MPERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MPERSession struct {
	Contract     *MPER             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MPERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MPERCallerSession struct {
	Contract *MPERCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MPERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MPERTransactorSession struct {
	Contract     *MPERTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MPERRaw is an auto generated low-level Go binding around an Ethereum contract.
type MPERRaw struct {
	Contract *MPER // Generic contract binding to access the raw methods on
}

// MPERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MPERCallerRaw struct {
	Contract *MPERCaller // Generic read-only contract binding to access the raw methods on
}

// MPERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MPERTransactorRaw struct {
	Contract *MPERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMPER creates a new instance of MPER, bound to a specific deployed contract.
func NewMPER(address common.Address, backend bind.ContractBackend) (*MPER, error) {
	contract, err := bindMPER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MPER{MPERCaller: MPERCaller{contract: contract}, MPERTransactor: MPERTransactor{contract: contract}, MPERFilterer: MPERFilterer{contract: contract}}, nil
}

// NewMPERCaller creates a new read-only instance of MPER, bound to a specific deployed contract.
func NewMPERCaller(address common.Address, caller bind.ContractCaller) (*MPERCaller, error) {
	contract, err := bindMPER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MPERCaller{contract: contract}, nil
}

// NewMPERTransactor creates a new write-only instance of MPER, bound to a specific deployed contract.
func NewMPERTransactor(address common.Address, transactor bind.ContractTransactor) (*MPERTransactor, error) {
	contract, err := bindMPER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MPERTransactor{contract: contract}, nil
}

// NewMPERFilterer creates a new log filterer instance of MPER, bound to a specific deployed contract.
func NewMPERFilterer(address common.Address, filterer bind.ContractFilterer) (*MPERFilterer, error) {
	contract, err := bindMPER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MPERFilterer{contract: contract}, nil
}

// bindMPER binds a generic wrapper to an already deployed contract.
func bindMPER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MPERMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MPER *MPERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MPER.Contract.MPERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MPER *MPERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MPER.Contract.MPERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MPER *MPERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MPER.Contract.MPERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MPER *MPERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MPER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MPER *MPERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MPER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MPER *MPERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MPER.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MPER *MPERCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MPER.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MPER *MPERSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MPER.Contract.Allowance(&_MPER.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MPER *MPERCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MPER.Contract.Allowance(&_MPER.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MPER *MPERCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MPER.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MPER *MPERSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MPER.Contract.BalanceOf(&_MPER.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MPER *MPERCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MPER.Contract.BalanceOf(&_MPER.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MPER *MPERCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MPER.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MPER *MPERSession) Decimals() (uint8, error) {
	return _MPER.Contract.Decimals(&_MPER.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MPER *MPERCallerSession) Decimals() (uint8, error) {
	return _MPER.Contract.Decimals(&_MPER.CallOpts)
}

// GetRecipientByHash is a free data retrieval call binding the contract method 0x882514f2.
//
// Solidity: function getRecipientByHash(string hash) view returns(address)
func (_MPER *MPERCaller) GetRecipientByHash(opts *bind.CallOpts, hash string) (common.Address, error) {
	var out []interface{}
	err := _MPER.contract.Call(opts, &out, "getRecipientByHash", hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecipientByHash is a free data retrieval call binding the contract method 0x882514f2.
//
// Solidity: function getRecipientByHash(string hash) view returns(address)
func (_MPER *MPERSession) GetRecipientByHash(hash string) (common.Address, error) {
	return _MPER.Contract.GetRecipientByHash(&_MPER.CallOpts, hash)
}

// GetRecipientByHash is a free data retrieval call binding the contract method 0x882514f2.
//
// Solidity: function getRecipientByHash(string hash) view returns(address)
func (_MPER *MPERCallerSession) GetRecipientByHash(hash string) (common.Address, error) {
	return _MPER.Contract.GetRecipientByHash(&_MPER.CallOpts, hash)
}

// GetReviewByHash is a free data retrieval call binding the contract method 0xdb836881.
//
// Solidity: function getReviewByHash(string content) view returns(address)
func (_MPER *MPERCaller) GetReviewByHash(opts *bind.CallOpts, content string) (common.Address, error) {
	var out []interface{}
	err := _MPER.contract.Call(opts, &out, "getReviewByHash", content)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReviewByHash is a free data retrieval call binding the contract method 0xdb836881.
//
// Solidity: function getReviewByHash(string content) view returns(address)
func (_MPER *MPERSession) GetReviewByHash(content string) (common.Address, error) {
	return _MPER.Contract.GetReviewByHash(&_MPER.CallOpts, content)
}

// GetReviewByHash is a free data retrieval call binding the contract method 0xdb836881.
//
// Solidity: function getReviewByHash(string content) view returns(address)
func (_MPER *MPERCallerSession) GetReviewByHash(content string) (common.Address, error) {
	return _MPER.Contract.GetReviewByHash(&_MPER.CallOpts, content)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MPER *MPERCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MPER.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MPER *MPERSession) Name() (string, error) {
	return _MPER.Contract.Name(&_MPER.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MPER *MPERCallerSession) Name() (string, error) {
	return _MPER.Contract.Name(&_MPER.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MPER *MPERCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MPER.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MPER *MPERSession) Symbol() (string, error) {
	return _MPER.Contract.Symbol(&_MPER.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MPER *MPERCallerSession) Symbol() (string, error) {
	return _MPER.Contract.Symbol(&_MPER.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MPER *MPERCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MPER.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MPER *MPERSession) TotalSupply() (*big.Int, error) {
	return _MPER.Contract.TotalSupply(&_MPER.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MPER *MPERCallerSession) TotalSupply() (*big.Int, error) {
	return _MPER.Contract.TotalSupply(&_MPER.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MPER *MPERTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPER.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MPER *MPERSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.Approve(&_MPER.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MPER *MPERTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.Approve(&_MPER.TransactOpts, spender, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_MPER *MPERTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPER.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_MPER *MPERSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.BurnFrom(&_MPER.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_MPER *MPERTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.BurnFrom(&_MPER.TransactOpts, account, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialSupply) returns()
func (_MPER *MPERTransactor) Initialize(opts *bind.TransactOpts, initialSupply *big.Int) (*types.Transaction, error) {
	return _MPER.contract.Transact(opts, "initialize", initialSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialSupply) returns()
func (_MPER *MPERSession) Initialize(initialSupply *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.Initialize(&_MPER.TransactOpts, initialSupply)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialSupply) returns()
func (_MPER *MPERTransactorSession) Initialize(initialSupply *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.Initialize(&_MPER.TransactOpts, initialSupply)
}

// Mint is a paid mutator transaction binding the contract method 0xea66696c.
//
// Solidity: function mint(address[] toAddresses, uint256 amount) returns(bool)
func (_MPER *MPERTransactor) Mint(opts *bind.TransactOpts, toAddresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPER.contract.Transact(opts, "mint", toAddresses, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xea66696c.
//
// Solidity: function mint(address[] toAddresses, uint256 amount) returns(bool)
func (_MPER *MPERSession) Mint(toAddresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.Mint(&_MPER.TransactOpts, toAddresses, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xea66696c.
//
// Solidity: function mint(address[] toAddresses, uint256 amount) returns(bool)
func (_MPER *MPERTransactorSession) Mint(toAddresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.Mint(&_MPER.TransactOpts, toAddresses, amount)
}

// StoreHash is a paid mutator transaction binding the contract method 0x71dc61cb.
//
// Solidity: function storeHash(string hash) returns()
func (_MPER *MPERTransactor) StoreHash(opts *bind.TransactOpts, hash string) (*types.Transaction, error) {
	return _MPER.contract.Transact(opts, "storeHash", hash)
}

// StoreHash is a paid mutator transaction binding the contract method 0x71dc61cb.
//
// Solidity: function storeHash(string hash) returns()
func (_MPER *MPERSession) StoreHash(hash string) (*types.Transaction, error) {
	return _MPER.Contract.StoreHash(&_MPER.TransactOpts, hash)
}

// StoreHash is a paid mutator transaction binding the contract method 0x71dc61cb.
//
// Solidity: function storeHash(string hash) returns()
func (_MPER *MPERTransactorSession) StoreHash(hash string) (*types.Transaction, error) {
	return _MPER.Contract.StoreHash(&_MPER.TransactOpts, hash)
}

// StoreReview is a paid mutator transaction binding the contract method 0x8c23b081.
//
// Solidity: function storeReview(string content) returns()
func (_MPER *MPERTransactor) StoreReview(opts *bind.TransactOpts, content string) (*types.Transaction, error) {
	return _MPER.contract.Transact(opts, "storeReview", content)
}

// StoreReview is a paid mutator transaction binding the contract method 0x8c23b081.
//
// Solidity: function storeReview(string content) returns()
func (_MPER *MPERSession) StoreReview(content string) (*types.Transaction, error) {
	return _MPER.Contract.StoreReview(&_MPER.TransactOpts, content)
}

// StoreReview is a paid mutator transaction binding the contract method 0x8c23b081.
//
// Solidity: function storeReview(string content) returns()
func (_MPER *MPERTransactorSession) StoreReview(content string) (*types.Transaction, error) {
	return _MPER.Contract.StoreReview(&_MPER.TransactOpts, content)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MPER *MPERTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPER.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MPER *MPERSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.Transfer(&_MPER.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MPER *MPERTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.Transfer(&_MPER.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MPER *MPERTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPER.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MPER *MPERSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.TransferFrom(&_MPER.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MPER *MPERTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MPER.Contract.TransferFrom(&_MPER.TransactOpts, from, to, value)
}

// MPERApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MPER contract.
type MPERApprovalIterator struct {
	Event *MPERApproval // Event containing the contract specifics and raw log

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
func (it *MPERApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPERApproval)
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
		it.Event = new(MPERApproval)
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
func (it *MPERApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPERApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPERApproval represents a Approval event raised by the MPER contract.
type MPERApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MPER *MPERFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MPERApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MPER.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MPERApprovalIterator{contract: _MPER.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MPER *MPERFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MPERApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MPER.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPERApproval)
				if err := _MPER.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MPER *MPERFilterer) ParseApproval(log types.Log) (*MPERApproval, error) {
	event := new(MPERApproval)
	if err := _MPER.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPERBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the MPER contract.
type MPERBurnIterator struct {
	Event *MPERBurn // Event containing the contract specifics and raw log

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
func (it *MPERBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPERBurn)
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
		it.Event = new(MPERBurn)
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
func (it *MPERBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPERBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPERBurn represents a Burn event raised by the MPER contract.
type MPERBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_MPER *MPERFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*MPERBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MPER.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &MPERBurnIterator{contract: _MPER.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_MPER *MPERFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *MPERBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MPER.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPERBurn)
				if err := _MPER.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_MPER *MPERFilterer) ParseBurn(log types.Log) (*MPERBurn, error) {
	event := new(MPERBurn)
	if err := _MPER.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPERHashStoredIterator is returned from FilterHashStored and is used to iterate over the raw logs and unpacked data for HashStored events raised by the MPER contract.
type MPERHashStoredIterator struct {
	Event *MPERHashStored // Event containing the contract specifics and raw log

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
func (it *MPERHashStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPERHashStored)
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
		it.Event = new(MPERHashStored)
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
func (it *MPERHashStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPERHashStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPERHashStored represents a HashStored event raised by the MPER contract.
type MPERHashStored struct {
	Hash   string
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHashStored is a free log retrieval operation binding the contract event 0x86fe26a4c2c4da61049e3786668239929ce2f0465c18b8c655ad4cd126df9ea9.
//
// Solidity: event HashStored(string hash, address indexed sender)
func (_MPER *MPERFilterer) FilterHashStored(opts *bind.FilterOpts, sender []common.Address) (*MPERHashStoredIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MPER.contract.FilterLogs(opts, "HashStored", senderRule)
	if err != nil {
		return nil, err
	}
	return &MPERHashStoredIterator{contract: _MPER.contract, event: "HashStored", logs: logs, sub: sub}, nil
}

// WatchHashStored is a free log subscription operation binding the contract event 0x86fe26a4c2c4da61049e3786668239929ce2f0465c18b8c655ad4cd126df9ea9.
//
// Solidity: event HashStored(string hash, address indexed sender)
func (_MPER *MPERFilterer) WatchHashStored(opts *bind.WatchOpts, sink chan<- *MPERHashStored, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MPER.contract.WatchLogs(opts, "HashStored", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPERHashStored)
				if err := _MPER.contract.UnpackLog(event, "HashStored", log); err != nil {
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
func (_MPER *MPERFilterer) ParseHashStored(log types.Log) (*MPERHashStored, error) {
	event := new(MPERHashStored)
	if err := _MPER.contract.UnpackLog(event, "HashStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPERInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MPER contract.
type MPERInitializedIterator struct {
	Event *MPERInitialized // Event containing the contract specifics and raw log

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
func (it *MPERInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPERInitialized)
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
		it.Event = new(MPERInitialized)
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
func (it *MPERInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPERInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPERInitialized represents a Initialized event raised by the MPER contract.
type MPERInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MPER *MPERFilterer) FilterInitialized(opts *bind.FilterOpts) (*MPERInitializedIterator, error) {

	logs, sub, err := _MPER.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MPERInitializedIterator{contract: _MPER.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MPER *MPERFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MPERInitialized) (event.Subscription, error) {

	logs, sub, err := _MPER.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPERInitialized)
				if err := _MPER.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MPER *MPERFilterer) ParseInitialized(log types.Log) (*MPERInitialized, error) {
	event := new(MPERInitialized)
	if err := _MPER.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPERMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the MPER contract.
type MPERMintIterator struct {
	Event *MPERMint // Event containing the contract specifics and raw log

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
func (it *MPERMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPERMint)
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
		it.Event = new(MPERMint)
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
func (it *MPERMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPERMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPERMint represents a Mint event raised by the MPER contract.
type MPERMint struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_MPER *MPERFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*MPERMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPER.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &MPERMintIterator{contract: _MPER.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 value)
func (_MPER *MPERFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *MPERMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPER.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPERMint)
				if err := _MPER.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_MPER *MPERFilterer) ParseMint(log types.Log) (*MPERMint, error) {
	event := new(MPERMint)
	if err := _MPER.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPERReviewStoredIterator is returned from FilterReviewStored and is used to iterate over the raw logs and unpacked data for ReviewStored events raised by the MPER contract.
type MPERReviewStoredIterator struct {
	Event *MPERReviewStored // Event containing the contract specifics and raw log

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
func (it *MPERReviewStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPERReviewStored)
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
		it.Event = new(MPERReviewStored)
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
func (it *MPERReviewStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPERReviewStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPERReviewStored represents a ReviewStored event raised by the MPER contract.
type MPERReviewStored struct {
	Content string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReviewStored is a free log retrieval operation binding the contract event 0xcb08fdcb10363722336d254fcc0aa3a2f24fe63245fdbb1fdd1a1064d8c75ea2.
//
// Solidity: event ReviewStored(string content)
func (_MPER *MPERFilterer) FilterReviewStored(opts *bind.FilterOpts) (*MPERReviewStoredIterator, error) {

	logs, sub, err := _MPER.contract.FilterLogs(opts, "ReviewStored")
	if err != nil {
		return nil, err
	}
	return &MPERReviewStoredIterator{contract: _MPER.contract, event: "ReviewStored", logs: logs, sub: sub}, nil
}

// WatchReviewStored is a free log subscription operation binding the contract event 0xcb08fdcb10363722336d254fcc0aa3a2f24fe63245fdbb1fdd1a1064d8c75ea2.
//
// Solidity: event ReviewStored(string content)
func (_MPER *MPERFilterer) WatchReviewStored(opts *bind.WatchOpts, sink chan<- *MPERReviewStored) (event.Subscription, error) {

	logs, sub, err := _MPER.contract.WatchLogs(opts, "ReviewStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPERReviewStored)
				if err := _MPER.contract.UnpackLog(event, "ReviewStored", log); err != nil {
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
func (_MPER *MPERFilterer) ParseReviewStored(log types.Log) (*MPERReviewStored, error) {
	event := new(MPERReviewStored)
	if err := _MPER.contract.UnpackLog(event, "ReviewStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPERTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MPER contract.
type MPERTransferIterator struct {
	Event *MPERTransfer // Event containing the contract specifics and raw log

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
func (it *MPERTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPERTransfer)
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
		it.Event = new(MPERTransfer)
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
func (it *MPERTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPERTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPERTransfer represents a Transfer event raised by the MPER contract.
type MPERTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MPER *MPERFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MPERTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPER.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MPERTransferIterator{contract: _MPER.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MPER *MPERFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MPERTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPER.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPERTransfer)
				if err := _MPER.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MPER *MPERFilterer) ParseTransfer(log types.Log) (*MPERTransfer, error) {
	event := new(MPERTransfer)
	if err := _MPER.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MPERTransferWithHashIterator is returned from FilterTransferWithHash and is used to iterate over the raw logs and unpacked data for TransferWithHash events raised by the MPER contract.
type MPERTransferWithHashIterator struct {
	Event *MPERTransferWithHash // Event containing the contract specifics and raw log

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
func (it *MPERTransferWithHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MPERTransferWithHash)
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
		it.Event = new(MPERTransferWithHash)
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
func (it *MPERTransferWithHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MPERTransferWithHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MPERTransferWithHash represents a TransferWithHash event raised by the MPER contract.
type MPERTransferWithHash struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferWithHash is a free log retrieval operation binding the contract event 0x6b525005bc93437a219538e3a08ce19ff9b23abacb849261956e3eedd93b16db.
//
// Solidity: event TransferWithHash(address indexed from, address indexed to, uint256 value)
func (_MPER *MPERFilterer) FilterTransferWithHash(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MPERTransferWithHashIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPER.contract.FilterLogs(opts, "TransferWithHash", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MPERTransferWithHashIterator{contract: _MPER.contract, event: "TransferWithHash", logs: logs, sub: sub}, nil
}

// WatchTransferWithHash is a free log subscription operation binding the contract event 0x6b525005bc93437a219538e3a08ce19ff9b23abacb849261956e3eedd93b16db.
//
// Solidity: event TransferWithHash(address indexed from, address indexed to, uint256 value)
func (_MPER *MPERFilterer) WatchTransferWithHash(opts *bind.WatchOpts, sink chan<- *MPERTransferWithHash, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MPER.contract.WatchLogs(opts, "TransferWithHash", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MPERTransferWithHash)
				if err := _MPER.contract.UnpackLog(event, "TransferWithHash", log); err != nil {
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
func (_MPER *MPERFilterer) ParseTransferWithHash(log types.Log) (*MPERTransferWithHash, error) {
	event := new(MPERTransferWithHash)
	if err := _MPER.contract.UnpackLog(event, "TransferWithHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
