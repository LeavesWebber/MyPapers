package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MPSABI is the input ABI used to generate the binding from.
const MPSABI = `[{"inputs":[{"internalType":"uint256","name":"initialSupply","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"allowance","type":"uint256"},{"internalType":"uint256","name":"needed","type":"uint256"}],"name":"ERC20InsufficientAllowance","type":"error"},{"inputs":[{"internalType":"address","name":"sender","type":"address"},{"internalType":"uint256","name":"balance","type":"uint256"},{"internalType":"uint256","name":"needed","type":"uint256"}],"name":"ERC20InsufficientBalance","type":"error"},{"inputs":[{"internalType":"address","name":"approver","type":"address"}],"name":"ERC20InvalidApprover","type":"error"},{"inputs":[{"internalType":"address","name":"receiver","type":"address"}],"name":"ERC20InvalidReceiver","type":"error"},{"inputs":[{"internalType":"address","name":"sender","type":"address"}],"name":"ERC20InvalidSender","type":"error"},{"inputs":[{"internalType":"address","name":"spender","type":"address"}],"name":"ERC20InvalidSpender","type":"error"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"user","type":"address"}],"name":"registerUser","type":"function","stateMutability":"nonpayable"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Burn","type":"event"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"burnFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"hash","type":"string"},{"indexed":true,"internalType":"address","name":"sender","type":"address"}],"name":"HashStored","type":"event"},{"inputs":[{"internalType":"address[]","name":"toAddresses","type":"address[]"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"mint","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Mint","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"content","type":"string"}],"name":"ReviewStored","type":"event"},{"inputs":[{"internalType":"string","name":"hash","type":"string"}],"name":"storeHash","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"content","type":"string"}],"name":"storeReview","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"recipient","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"TransferWithHash","type":"event"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"hash","type":"string"}],"name":"getRecipientByHash","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"content","type":"string"}],"name":"getReviewByHash","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"user","type":"address"}],"name":"registerUser","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"hash","type":"string"}],"name":"storeHash","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"content","type":"string"}],"name":"storeReview","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`

// MPS is an auto generated Go binding around an Ethereum contract.
type MPS struct {
	MPSCaller     // Read-only binding to the contract
	MPSTransactor // Write-only binding to the contract
	MPSFilterer   // Log filterer for contract events
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

// NewMPS creates a new instance of MPS, bound to a specific deployed contract.
func NewMPS(address common.Address, backend bind.ContractBackend) (*MPS, error) {
	contract, err := bindMPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MPS{MPSCaller: MPSCaller{contract: contract}, MPSTransactor: MPSTransactor{contract: contract}, MPSFilterer: MPSFilterer{contract: contract}}, nil
}

// bindMPS binds a generic wrapper to an already deployed contract.
func bindMPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MPSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
func (m *MPSCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := m.contract.Call(opts, &out, "balanceOf", account)
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

// GetRecipientByHash is a free data retrieval call binding the contract method.
func (m *MPSCaller) GetRecipientByHash(opts *bind.CallOpts, hash string) (common.Address, error) {
	var out []interface{}
	err := m.contract.Call(opts, &out, "getRecipientByHash", hash)
	if err != nil {
		return common.Address{}, err
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

// GetReviewByHash is a free data retrieval call binding the contract method.
func (m *MPSCaller) GetReviewByHash(opts *bind.CallOpts, content string) (common.Address, error) {
	var out []interface{}
	err := m.contract.Call(opts, &out, "getReviewByHash", content)
	if err != nil {
		return common.Address{}, err
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
func (m *MPSTransactor) Mint(opts *bind.TransactOpts, toAddresses []common.Address, amount *big.Int) (*types.Transaction, error) {
	return m.contract.Transact(opts, "mint", toAddresses, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
func (m *MPSTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return m.contract.Transact(opts, "transfer", recipient, amount)
}

// StoreHash is a paid mutator transaction binding the contract method.
func (m *MPSTransactor) StoreHash(opts *bind.TransactOpts, hash string) (*types.Transaction, error) {
	return m.contract.Transact(opts, "storeHash", hash)
}

// StoreReview is a paid mutator transaction binding the contract method.
func (m *MPSTransactor) StoreReview(opts *bind.TransactOpts, content string) (*types.Transaction, error) {
	return m.contract.Transact(opts, "storeReview", content)
}

// RegisterUser is a paid mutator transaction binding the contract method.
func (m *MPSTransactor) RegisterUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return m.contract.Transact(opts, "registerUser", user)
}
