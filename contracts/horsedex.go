// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// HorseDexABI is the input ABI used to generate the binding from.
const HorseDexABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"orders\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tradeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OrderPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"OrderCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"OrderRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"OrderProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"placeOrder\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"rejectOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"purchaseHORSE\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"processOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HorseDex is an auto generated Go binding around an Ethereum contract.
type HorseDex struct {
	HorseDexCaller     // Read-only binding to the contract
	HorseDexTransactor // Write-only binding to the contract
	HorseDexFilterer   // Log filterer for contract events
}

// HorseDexCaller is an auto generated read-only Go binding around an Ethereum contract.
type HorseDexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HorseDexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HorseDexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HorseDexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HorseDexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HorseDexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HorseDexSession struct {
	Contract     *HorseDex         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HorseDexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HorseDexCallerSession struct {
	Contract *HorseDexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HorseDexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HorseDexTransactorSession struct {
	Contract     *HorseDexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HorseDexRaw is an auto generated low-level Go binding around an Ethereum contract.
type HorseDexRaw struct {
	Contract *HorseDex // Generic contract binding to access the raw methods on
}

// HorseDexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HorseDexCallerRaw struct {
	Contract *HorseDexCaller // Generic read-only contract binding to access the raw methods on
}

// HorseDexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HorseDexTransactorRaw struct {
	Contract *HorseDexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHorseDex creates a new instance of HorseDex, bound to a specific deployed contract.
func NewHorseDex(address common.Address, backend bind.ContractBackend) (*HorseDex, error) {
	contract, err := bindHorseDex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HorseDex{HorseDexCaller: HorseDexCaller{contract: contract}, HorseDexTransactor: HorseDexTransactor{contract: contract}, HorseDexFilterer: HorseDexFilterer{contract: contract}}, nil
}

// NewHorseDexCaller creates a new read-only instance of HorseDex, bound to a specific deployed contract.
func NewHorseDexCaller(address common.Address, caller bind.ContractCaller) (*HorseDexCaller, error) {
	contract, err := bindHorseDex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HorseDexCaller{contract: contract}, nil
}

// NewHorseDexTransactor creates a new write-only instance of HorseDex, bound to a specific deployed contract.
func NewHorseDexTransactor(address common.Address, transactor bind.ContractTransactor) (*HorseDexTransactor, error) {
	contract, err := bindHorseDex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HorseDexTransactor{contract: contract}, nil
}

// NewHorseDexFilterer creates a new log filterer instance of HorseDex, bound to a specific deployed contract.
func NewHorseDexFilterer(address common.Address, filterer bind.ContractFilterer) (*HorseDexFilterer, error) {
	contract, err := bindHorseDex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HorseDexFilterer{contract: contract}, nil
}

// bindHorseDex binds a generic wrapper to an already deployed contract.
func bindHorseDex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HorseDexABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HorseDex *HorseDexRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HorseDex.Contract.HorseDexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HorseDex *HorseDexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HorseDex.Contract.HorseDexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HorseDex *HorseDexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HorseDex.Contract.HorseDexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HorseDex *HorseDexCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HorseDex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HorseDex *HorseDexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HorseDex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HorseDex *HorseDexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HorseDex.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_HorseDex *HorseDexCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HorseDex.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_HorseDex *HorseDexSession) IsOwner() (bool, error) {
	return _HorseDex.Contract.IsOwner(&_HorseDex.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_HorseDex *HorseDexCallerSession) IsOwner() (bool, error) {
	return _HorseDex.Contract.IsOwner(&_HorseDex.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_HorseDex *HorseDexCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HorseDex.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_HorseDex *HorseDexSession) IsPauser(account common.Address) (bool, error) {
	return _HorseDex.Contract.IsPauser(&_HorseDex.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(account address) constant returns(bool)
func (_HorseDex *HorseDexCallerSession) IsPauser(account common.Address) (bool, error) {
	return _HorseDex.Contract.IsPauser(&_HorseDex.CallOpts, account)
}

// Orders is a free data retrieval call binding the contract method 0xf40e8471.
//
// Solidity: function orders( address) constant returns(amount uint256, value uint256, buyer address)
func (_HorseDex *HorseDexCaller) Orders(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
	Buyer  common.Address
}, error) {
	ret := new(struct {
		Amount *big.Int
		Value  *big.Int
		Buyer  common.Address
	})
	out := ret
	err := _HorseDex.contract.Call(opts, out, "orders", arg0)
	return *ret, err
}

// Orders is a free data retrieval call binding the contract method 0xf40e8471.
//
// Solidity: function orders( address) constant returns(amount uint256, value uint256, buyer address)
func (_HorseDex *HorseDexSession) Orders(arg0 common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
	Buyer  common.Address
}, error) {
	return _HorseDex.Contract.Orders(&_HorseDex.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xf40e8471.
//
// Solidity: function orders( address) constant returns(amount uint256, value uint256, buyer address)
func (_HorseDex *HorseDexCallerSession) Orders(arg0 common.Address) (struct {
	Amount *big.Int
	Value  *big.Int
	Buyer  common.Address
}, error) {
	return _HorseDex.Contract.Orders(&_HorseDex.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HorseDex *HorseDexCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HorseDex.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HorseDex *HorseDexSession) Owner() (common.Address, error) {
	return _HorseDex.Contract.Owner(&_HorseDex.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HorseDex *HorseDexCallerSession) Owner() (common.Address, error) {
	return _HorseDex.Contract.Owner(&_HorseDex.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_HorseDex *HorseDexCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HorseDex.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_HorseDex *HorseDexSession) Paused() (bool, error) {
	return _HorseDex.Contract.Paused(&_HorseDex.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_HorseDex *HorseDexCallerSession) Paused() (bool, error) {
	return _HorseDex.Contract.Paused(&_HorseDex.CallOpts)
}

// TradeContract is a free data retrieval call binding the contract method 0xffa640d8.
//
// Solidity: function tradeContract() constant returns(address)
func (_HorseDex *HorseDexCaller) TradeContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HorseDex.contract.Call(opts, out, "tradeContract")
	return *ret0, err
}

// TradeContract is a free data retrieval call binding the contract method 0xffa640d8.
//
// Solidity: function tradeContract() constant returns(address)
func (_HorseDex *HorseDexSession) TradeContract() (common.Address, error) {
	return _HorseDex.Contract.TradeContract(&_HorseDex.CallOpts)
}

// TradeContract is a free data retrieval call binding the contract method 0xffa640d8.
//
// Solidity: function tradeContract() constant returns(address)
func (_HorseDex *HorseDexCallerSession) TradeContract() (common.Address, error) {
	return _HorseDex.Contract.TradeContract(&_HorseDex.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_HorseDex *HorseDexTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_HorseDex *HorseDexSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _HorseDex.Contract.AddPauser(&_HorseDex.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(account address) returns()
func (_HorseDex *HorseDexTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _HorseDex.Contract.AddPauser(&_HorseDex.TransactOpts, account)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x6a816548.
//
// Solidity: function cancelOrder() returns()
func (_HorseDex *HorseDexTransactor) CancelOrder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "cancelOrder")
}

// CancelOrder is a paid mutator transaction binding the contract method 0x6a816548.
//
// Solidity: function cancelOrder() returns()
func (_HorseDex *HorseDexSession) CancelOrder() (*types.Transaction, error) {
	return _HorseDex.Contract.CancelOrder(&_HorseDex.TransactOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x6a816548.
//
// Solidity: function cancelOrder() returns()
func (_HorseDex *HorseDexTransactorSession) CancelOrder() (*types.Transaction, error) {
	return _HorseDex.Contract.CancelOrder(&_HorseDex.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HorseDex *HorseDexTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HorseDex *HorseDexSession) Pause() (*types.Transaction, error) {
	return _HorseDex.Contract.Pause(&_HorseDex.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_HorseDex *HorseDexTransactorSession) Pause() (*types.Transaction, error) {
	return _HorseDex.Contract.Pause(&_HorseDex.TransactOpts)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0xba2255f4.
//
// Solidity: function placeOrder(amount uint256) returns()
func (_HorseDex *HorseDexTransactor) PlaceOrder(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "placeOrder", amount)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0xba2255f4.
//
// Solidity: function placeOrder(amount uint256) returns()
func (_HorseDex *HorseDexSession) PlaceOrder(amount *big.Int) (*types.Transaction, error) {
	return _HorseDex.Contract.PlaceOrder(&_HorseDex.TransactOpts, amount)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0xba2255f4.
//
// Solidity: function placeOrder(amount uint256) returns()
func (_HorseDex *HorseDexTransactorSession) PlaceOrder(amount *big.Int) (*types.Transaction, error) {
	return _HorseDex.Contract.PlaceOrder(&_HorseDex.TransactOpts, amount)
}

// ProcessOrder is a paid mutator transaction binding the contract method 0xea9f6b1f.
//
// Solidity: function processOrder(buyer address) returns()
func (_HorseDex *HorseDexTransactor) ProcessOrder(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "processOrder", buyer)
}

// ProcessOrder is a paid mutator transaction binding the contract method 0xea9f6b1f.
//
// Solidity: function processOrder(buyer address) returns()
func (_HorseDex *HorseDexSession) ProcessOrder(buyer common.Address) (*types.Transaction, error) {
	return _HorseDex.Contract.ProcessOrder(&_HorseDex.TransactOpts, buyer)
}

// ProcessOrder is a paid mutator transaction binding the contract method 0xea9f6b1f.
//
// Solidity: function processOrder(buyer address) returns()
func (_HorseDex *HorseDexTransactorSession) ProcessOrder(buyer common.Address) (*types.Transaction, error) {
	return _HorseDex.Contract.ProcessOrder(&_HorseDex.TransactOpts, buyer)
}

// PurchaseHORSE is a paid mutator transaction binding the contract method 0x70b16eaa.
//
// Solidity: function purchaseHORSE(amountGet uint256, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_HorseDex *HorseDexTransactor) PurchaseHORSE(opts *bind.TransactOpts, amountGet *big.Int, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "purchaseHORSE", amountGet, amountGive, expires, nonce, user, v, r, s, amount)
}

// PurchaseHORSE is a paid mutator transaction binding the contract method 0x70b16eaa.
//
// Solidity: function purchaseHORSE(amountGet uint256, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_HorseDex *HorseDexSession) PurchaseHORSE(amountGet *big.Int, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _HorseDex.Contract.PurchaseHORSE(&_HorseDex.TransactOpts, amountGet, amountGive, expires, nonce, user, v, r, s, amount)
}

// PurchaseHORSE is a paid mutator transaction binding the contract method 0x70b16eaa.
//
// Solidity: function purchaseHORSE(amountGet uint256, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_HorseDex *HorseDexTransactorSession) PurchaseHORSE(amountGet *big.Int, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _HorseDex.Contract.PurchaseHORSE(&_HorseDex.TransactOpts, amountGet, amountGive, expires, nonce, user, v, r, s, amount)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x6f00fb0d.
//
// Solidity: function rejectOrder(buyer address) returns()
func (_HorseDex *HorseDexTransactor) RejectOrder(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "rejectOrder", buyer)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x6f00fb0d.
//
// Solidity: function rejectOrder(buyer address) returns()
func (_HorseDex *HorseDexSession) RejectOrder(buyer common.Address) (*types.Transaction, error) {
	return _HorseDex.Contract.RejectOrder(&_HorseDex.TransactOpts, buyer)
}

// RejectOrder is a paid mutator transaction binding the contract method 0x6f00fb0d.
//
// Solidity: function rejectOrder(buyer address) returns()
func (_HorseDex *HorseDexTransactorSession) RejectOrder(buyer common.Address) (*types.Transaction, error) {
	return _HorseDex.Contract.RejectOrder(&_HorseDex.TransactOpts, buyer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HorseDex *HorseDexTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HorseDex *HorseDexSession) RenounceOwnership() (*types.Transaction, error) {
	return _HorseDex.Contract.RenounceOwnership(&_HorseDex.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HorseDex *HorseDexTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HorseDex.Contract.RenounceOwnership(&_HorseDex.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_HorseDex *HorseDexTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_HorseDex *HorseDexSession) RenouncePauser() (*types.Transaction, error) {
	return _HorseDex.Contract.RenouncePauser(&_HorseDex.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_HorseDex *HorseDexTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _HorseDex.Contract.RenouncePauser(&_HorseDex.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HorseDex *HorseDexTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HorseDex *HorseDexSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HorseDex.Contract.TransferOwnership(&_HorseDex.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HorseDex *HorseDexTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HorseDex.Contract.TransferOwnership(&_HorseDex.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HorseDex *HorseDexTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HorseDex.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HorseDex *HorseDexSession) Unpause() (*types.Transaction, error) {
	return _HorseDex.Contract.Unpause(&_HorseDex.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_HorseDex *HorseDexTransactorSession) Unpause() (*types.Transaction, error) {
	return _HorseDex.Contract.Unpause(&_HorseDex.TransactOpts)
}

// HorseDexOrderCanceledIterator is returned from FilterOrderCanceled and is used to iterate over the raw logs and unpacked data for OrderCanceled events raised by the HorseDex contract.
type HorseDexOrderCanceledIterator struct {
	Event *HorseDexOrderCanceled // Event containing the contract specifics and raw log

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
func (it *HorseDexOrderCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HorseDexOrderCanceled)
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
		it.Event = new(HorseDexOrderCanceled)
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
func (it *HorseDexOrderCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HorseDexOrderCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HorseDexOrderCanceled represents a OrderCanceled event raised by the HorseDex contract.
type HorseDexOrderCanceled struct {
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrderCanceled is a free log retrieval operation binding the contract event 0x14142eeeab120e46bf322bcea8744b9b566e11f71db000c6acdcb92384cb8fe8.
//
// Solidity: e OrderCanceled(buyer indexed address)
func (_HorseDex *HorseDexFilterer) FilterOrderCanceled(opts *bind.FilterOpts, buyer []common.Address) (*HorseDexOrderCanceledIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _HorseDex.contract.FilterLogs(opts, "OrderCanceled", buyerRule)
	if err != nil {
		return nil, err
	}
	return &HorseDexOrderCanceledIterator{contract: _HorseDex.contract, event: "OrderCanceled", logs: logs, sub: sub}, nil
}

// WatchOrderCanceled is a free log subscription operation binding the contract event 0x14142eeeab120e46bf322bcea8744b9b566e11f71db000c6acdcb92384cb8fe8.
//
// Solidity: e OrderCanceled(buyer indexed address)
func (_HorseDex *HorseDexFilterer) WatchOrderCanceled(opts *bind.WatchOpts, sink chan<- *HorseDexOrderCanceled, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _HorseDex.contract.WatchLogs(opts, "OrderCanceled", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HorseDexOrderCanceled)
				if err := _HorseDex.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
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

// HorseDexOrderPlacedIterator is returned from FilterOrderPlaced and is used to iterate over the raw logs and unpacked data for OrderPlaced events raised by the HorseDex contract.
type HorseDexOrderPlacedIterator struct {
	Event *HorseDexOrderPlaced // Event containing the contract specifics and raw log

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
func (it *HorseDexOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HorseDexOrderPlaced)
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
		it.Event = new(HorseDexOrderPlaced)
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
func (it *HorseDexOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HorseDexOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HorseDexOrderPlaced represents a OrderPlaced event raised by the HorseDex contract.
type HorseDexOrderPlaced struct {
	Buyer  common.Address
	Amount *big.Int
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOrderPlaced is a free log retrieval operation binding the contract event 0x03f23a9b67af957c8605376a4f515b53f102d4ad7f4da099bd1aff8e3ee13fec.
//
// Solidity: e OrderPlaced(buyer indexed address, amount uint256, value uint256)
func (_HorseDex *HorseDexFilterer) FilterOrderPlaced(opts *bind.FilterOpts, buyer []common.Address) (*HorseDexOrderPlacedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _HorseDex.contract.FilterLogs(opts, "OrderPlaced", buyerRule)
	if err != nil {
		return nil, err
	}
	return &HorseDexOrderPlacedIterator{contract: _HorseDex.contract, event: "OrderPlaced", logs: logs, sub: sub}, nil
}

// WatchOrderPlaced is a free log subscription operation binding the contract event 0x03f23a9b67af957c8605376a4f515b53f102d4ad7f4da099bd1aff8e3ee13fec.
//
// Solidity: e OrderPlaced(buyer indexed address, amount uint256, value uint256)
func (_HorseDex *HorseDexFilterer) WatchOrderPlaced(opts *bind.WatchOpts, sink chan<- *HorseDexOrderPlaced, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _HorseDex.contract.WatchLogs(opts, "OrderPlaced", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HorseDexOrderPlaced)
				if err := _HorseDex.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
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

// HorseDexOrderProcessedIterator is returned from FilterOrderProcessed and is used to iterate over the raw logs and unpacked data for OrderProcessed events raised by the HorseDex contract.
type HorseDexOrderProcessedIterator struct {
	Event *HorseDexOrderProcessed // Event containing the contract specifics and raw log

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
func (it *HorseDexOrderProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HorseDexOrderProcessed)
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
		it.Event = new(HorseDexOrderProcessed)
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
func (it *HorseDexOrderProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HorseDexOrderProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HorseDexOrderProcessed represents a OrderProcessed event raised by the HorseDex contract.
type HorseDexOrderProcessed struct {
	Buyer  common.Address
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOrderProcessed is a free log retrieval operation binding the contract event 0x933fa41f9d6c587605bfcea9a1c4f7c6b02f95e24fede95b754267aa5284c93b.
//
// Solidity: e OrderProcessed(buyer indexed address, amount uint256, price uint256)
func (_HorseDex *HorseDexFilterer) FilterOrderProcessed(opts *bind.FilterOpts, buyer []common.Address) (*HorseDexOrderProcessedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _HorseDex.contract.FilterLogs(opts, "OrderProcessed", buyerRule)
	if err != nil {
		return nil, err
	}
	return &HorseDexOrderProcessedIterator{contract: _HorseDex.contract, event: "OrderProcessed", logs: logs, sub: sub}, nil
}

// WatchOrderProcessed is a free log subscription operation binding the contract event 0x933fa41f9d6c587605bfcea9a1c4f7c6b02f95e24fede95b754267aa5284c93b.
//
// Solidity: e OrderProcessed(buyer indexed address, amount uint256, price uint256)
func (_HorseDex *HorseDexFilterer) WatchOrderProcessed(opts *bind.WatchOpts, sink chan<- *HorseDexOrderProcessed, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _HorseDex.contract.WatchLogs(opts, "OrderProcessed", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HorseDexOrderProcessed)
				if err := _HorseDex.contract.UnpackLog(event, "OrderProcessed", log); err != nil {
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

// HorseDexOrderRejectedIterator is returned from FilterOrderRejected and is used to iterate over the raw logs and unpacked data for OrderRejected events raised by the HorseDex contract.
type HorseDexOrderRejectedIterator struct {
	Event *HorseDexOrderRejected // Event containing the contract specifics and raw log

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
func (it *HorseDexOrderRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HorseDexOrderRejected)
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
		it.Event = new(HorseDexOrderRejected)
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
func (it *HorseDexOrderRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HorseDexOrderRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HorseDexOrderRejected represents a OrderRejected event raised by the HorseDex contract.
type HorseDexOrderRejected struct {
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrderRejected is a free log retrieval operation binding the contract event 0x460e3f6737e1f438756ca3583f21a237ba79517150a13b5cb280ec61c38c6010.
//
// Solidity: e OrderRejected(buyer indexed address)
func (_HorseDex *HorseDexFilterer) FilterOrderRejected(opts *bind.FilterOpts, buyer []common.Address) (*HorseDexOrderRejectedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _HorseDex.contract.FilterLogs(opts, "OrderRejected", buyerRule)
	if err != nil {
		return nil, err
	}
	return &HorseDexOrderRejectedIterator{contract: _HorseDex.contract, event: "OrderRejected", logs: logs, sub: sub}, nil
}

// WatchOrderRejected is a free log subscription operation binding the contract event 0x460e3f6737e1f438756ca3583f21a237ba79517150a13b5cb280ec61c38c6010.
//
// Solidity: e OrderRejected(buyer indexed address)
func (_HorseDex *HorseDexFilterer) WatchOrderRejected(opts *bind.WatchOpts, sink chan<- *HorseDexOrderRejected, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _HorseDex.contract.WatchLogs(opts, "OrderRejected", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HorseDexOrderRejected)
				if err := _HorseDex.contract.UnpackLog(event, "OrderRejected", log); err != nil {
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

// HorseDexOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HorseDex contract.
type HorseDexOwnershipTransferredIterator struct {
	Event *HorseDexOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HorseDexOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HorseDexOwnershipTransferred)
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
		it.Event = new(HorseDexOwnershipTransferred)
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
func (it *HorseDexOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HorseDexOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HorseDexOwnershipTransferred represents a OwnershipTransferred event raised by the HorseDex contract.
type HorseDexOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HorseDex *HorseDexFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HorseDexOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HorseDex.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HorseDexOwnershipTransferredIterator{contract: _HorseDex.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HorseDex *HorseDexFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HorseDexOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HorseDex.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HorseDexOwnershipTransferred)
				if err := _HorseDex.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// HorseDexPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the HorseDex contract.
type HorseDexPausedIterator struct {
	Event *HorseDexPaused // Event containing the contract specifics and raw log

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
func (it *HorseDexPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HorseDexPaused)
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
		it.Event = new(HorseDexPaused)
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
func (it *HorseDexPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HorseDexPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HorseDexPaused represents a Paused event raised by the HorseDex contract.
type HorseDexPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: e Paused()
func (_HorseDex *HorseDexFilterer) FilterPaused(opts *bind.FilterOpts) (*HorseDexPausedIterator, error) {

	logs, sub, err := _HorseDex.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HorseDexPausedIterator{contract: _HorseDex.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: e Paused()
func (_HorseDex *HorseDexFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HorseDexPaused) (event.Subscription, error) {

	logs, sub, err := _HorseDex.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HorseDexPaused)
				if err := _HorseDex.contract.UnpackLog(event, "Paused", log); err != nil {
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

// HorseDexPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the HorseDex contract.
type HorseDexPauserAddedIterator struct {
	Event *HorseDexPauserAdded // Event containing the contract specifics and raw log

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
func (it *HorseDexPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HorseDexPauserAdded)
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
		it.Event = new(HorseDexPauserAdded)
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
func (it *HorseDexPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HorseDexPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HorseDexPauserAdded represents a PauserAdded event raised by the HorseDex contract.
type HorseDexPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_HorseDex *HorseDexFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*HorseDexPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HorseDex.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &HorseDexPauserAddedIterator{contract: _HorseDex.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: e PauserAdded(account indexed address)
func (_HorseDex *HorseDexFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *HorseDexPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HorseDex.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HorseDexPauserAdded)
				if err := _HorseDex.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// HorseDexPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the HorseDex contract.
type HorseDexPauserRemovedIterator struct {
	Event *HorseDexPauserRemoved // Event containing the contract specifics and raw log

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
func (it *HorseDexPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HorseDexPauserRemoved)
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
		it.Event = new(HorseDexPauserRemoved)
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
func (it *HorseDexPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HorseDexPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HorseDexPauserRemoved represents a PauserRemoved event raised by the HorseDex contract.
type HorseDexPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_HorseDex *HorseDexFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*HorseDexPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HorseDex.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &HorseDexPauserRemovedIterator{contract: _HorseDex.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: e PauserRemoved(account indexed address)
func (_HorseDex *HorseDexFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *HorseDexPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _HorseDex.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HorseDexPauserRemoved)
				if err := _HorseDex.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// HorseDexUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the HorseDex contract.
type HorseDexUnpausedIterator struct {
	Event *HorseDexUnpaused // Event containing the contract specifics and raw log

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
func (it *HorseDexUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HorseDexUnpaused)
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
		it.Event = new(HorseDexUnpaused)
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
func (it *HorseDexUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HorseDexUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HorseDexUnpaused represents a Unpaused event raised by the HorseDex contract.
type HorseDexUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: e Unpaused()
func (_HorseDex *HorseDexFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HorseDexUnpausedIterator, error) {

	logs, sub, err := _HorseDex.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HorseDexUnpausedIterator{contract: _HorseDex.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: e Unpaused()
func (_HorseDex *HorseDexFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HorseDexUnpaused) (event.Subscription, error) {

	logs, sub, err := _HorseDex.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HorseDexUnpaused)
				if err := _HorseDex.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
