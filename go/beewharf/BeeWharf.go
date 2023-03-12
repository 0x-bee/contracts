// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package beewharf

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
)

// BeewharfMetaData contains all meta data concerning the Beewharf contract.
var BeewharfMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"PayERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"}],\"name\":\"PayEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"RefundERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"}],\"name\":\"RefundEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"}],\"name\":\"WithdrawEth\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"addNewSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"balanceOfERC20\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"payERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"payERC20From\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"}],\"name\":\"payEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"refundERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"refundERC20From\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"}],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BeewharfABI is the input ABI used to generate the binding from.
// Deprecated: Use BeewharfMetaData.ABI instead.
var BeewharfABI = BeewharfMetaData.ABI

// Beewharf is an auto generated Go binding around an Ethereum contract.
type Beewharf struct {
	BeewharfCaller     // Read-only binding to the contract
	BeewharfTransactor // Write-only binding to the contract
	BeewharfFilterer   // Log filterer for contract events
}

// BeewharfCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeewharfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeewharfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeewharfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeewharfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeewharfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeewharfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeewharfSession struct {
	Contract     *Beewharf         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeewharfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeewharfCallerSession struct {
	Contract *BeewharfCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BeewharfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeewharfTransactorSession struct {
	Contract     *BeewharfTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BeewharfRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeewharfRaw struct {
	Contract *Beewharf // Generic contract binding to access the raw methods on
}

// BeewharfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeewharfCallerRaw struct {
	Contract *BeewharfCaller // Generic read-only contract binding to access the raw methods on
}

// BeewharfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeewharfTransactorRaw struct {
	Contract *BeewharfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeewharf creates a new instance of Beewharf, bound to a specific deployed contract.
func NewBeewharf(address common.Address, backend bind.ContractBackend) (*Beewharf, error) {
	contract, err := bindBeewharf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Beewharf{BeewharfCaller: BeewharfCaller{contract: contract}, BeewharfTransactor: BeewharfTransactor{contract: contract}, BeewharfFilterer: BeewharfFilterer{contract: contract}}, nil
}

// NewBeewharfCaller creates a new read-only instance of Beewharf, bound to a specific deployed contract.
func NewBeewharfCaller(address common.Address, caller bind.ContractCaller) (*BeewharfCaller, error) {
	contract, err := bindBeewharf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeewharfCaller{contract: contract}, nil
}

// NewBeewharfTransactor creates a new write-only instance of Beewharf, bound to a specific deployed contract.
func NewBeewharfTransactor(address common.Address, transactor bind.ContractTransactor) (*BeewharfTransactor, error) {
	contract, err := bindBeewharf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeewharfTransactor{contract: contract}, nil
}

// NewBeewharfFilterer creates a new log filterer instance of Beewharf, bound to a specific deployed contract.
func NewBeewharfFilterer(address common.Address, filterer bind.ContractFilterer) (*BeewharfFilterer, error) {
	contract, err := bindBeewharf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeewharfFilterer{contract: contract}, nil
}

// bindBeewharf binds a generic wrapper to an already deployed contract.
func bindBeewharf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BeewharfABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Beewharf *BeewharfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Beewharf.Contract.BeewharfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Beewharf *BeewharfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Beewharf.Contract.BeewharfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Beewharf *BeewharfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Beewharf.Contract.BeewharfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Beewharf *BeewharfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Beewharf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Beewharf *BeewharfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Beewharf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Beewharf *BeewharfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Beewharf.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfERC20 is a free data retrieval call binding the contract method 0xd93db24f.
//
// Solidity: function balanceOfERC20(address tokenAddress) view returns(uint256)
func (_Beewharf *BeewharfCaller) BalanceOfERC20(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Beewharf.contract.Call(opts, &out, "balanceOfERC20", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfERC20 is a free data retrieval call binding the contract method 0xd93db24f.
//
// Solidity: function balanceOfERC20(address tokenAddress) view returns(uint256)
func (_Beewharf *BeewharfSession) BalanceOfERC20(tokenAddress common.Address) (*big.Int, error) {
	return _Beewharf.Contract.BalanceOfERC20(&_Beewharf.CallOpts, tokenAddress)
}

// BalanceOfERC20 is a free data retrieval call binding the contract method 0xd93db24f.
//
// Solidity: function balanceOfERC20(address tokenAddress) view returns(uint256)
func (_Beewharf *BeewharfCallerSession) BalanceOfERC20(tokenAddress common.Address) (*big.Int, error) {
	return _Beewharf.Contract.BalanceOfERC20(&_Beewharf.CallOpts, tokenAddress)
}

// BalanceOfETH is a free data retrieval call binding the contract method 0x8ea5b802.
//
// Solidity: function balanceOfETH() view returns(uint256)
func (_Beewharf *BeewharfCaller) BalanceOfETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Beewharf.contract.Call(opts, &out, "balanceOfETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfETH is a free data retrieval call binding the contract method 0x8ea5b802.
//
// Solidity: function balanceOfETH() view returns(uint256)
func (_Beewharf *BeewharfSession) BalanceOfETH() (*big.Int, error) {
	return _Beewharf.Contract.BalanceOfETH(&_Beewharf.CallOpts)
}

// BalanceOfETH is a free data retrieval call binding the contract method 0x8ea5b802.
//
// Solidity: function balanceOfETH() view returns(uint256)
func (_Beewharf *BeewharfCallerSession) BalanceOfETH() (*big.Int, error) {
	return _Beewharf.Contract.BalanceOfETH(&_Beewharf.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Beewharf *BeewharfCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Beewharf.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Beewharf *BeewharfSession) Owner() (common.Address, error) {
	return _Beewharf.Contract.Owner(&_Beewharf.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Beewharf *BeewharfCallerSession) Owner() (common.Address, error) {
	return _Beewharf.Contract.Owner(&_Beewharf.CallOpts)
}

// AddNewSupportToken is a paid mutator transaction binding the contract method 0xb30917ad.
//
// Solidity: function addNewSupportToken(address tokenAddress) returns()
func (_Beewharf *BeewharfTransactor) AddNewSupportToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "addNewSupportToken", tokenAddress)
}

// AddNewSupportToken is a paid mutator transaction binding the contract method 0xb30917ad.
//
// Solidity: function addNewSupportToken(address tokenAddress) returns()
func (_Beewharf *BeewharfSession) AddNewSupportToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.AddNewSupportToken(&_Beewharf.TransactOpts, tokenAddress)
}

// AddNewSupportToken is a paid mutator transaction binding the contract method 0xb30917ad.
//
// Solidity: function addNewSupportToken(address tokenAddress) returns()
func (_Beewharf *BeewharfTransactorSession) AddNewSupportToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.AddNewSupportToken(&_Beewharf.TransactOpts, tokenAddress)
}

// PayERC20 is a paid mutator transaction binding the contract method 0x711e0ee7.
//
// Solidity: function payERC20(uint256 amount, string orderId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactor) PayERC20(opts *bind.TransactOpts, amount *big.Int, orderId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "payERC20", amount, orderId, tokenAddress)
}

// PayERC20 is a paid mutator transaction binding the contract method 0x711e0ee7.
//
// Solidity: function payERC20(uint256 amount, string orderId, address tokenAddress) returns()
func (_Beewharf *BeewharfSession) PayERC20(amount *big.Int, orderId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.PayERC20(&_Beewharf.TransactOpts, amount, orderId, tokenAddress)
}

// PayERC20 is a paid mutator transaction binding the contract method 0x711e0ee7.
//
// Solidity: function payERC20(uint256 amount, string orderId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactorSession) PayERC20(amount *big.Int, orderId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.PayERC20(&_Beewharf.TransactOpts, amount, orderId, tokenAddress)
}

// PayERC20From is a paid mutator transaction binding the contract method 0x793191f1.
//
// Solidity: function payERC20From(address from, uint256 amount, string orderId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactor) PayERC20From(opts *bind.TransactOpts, from common.Address, amount *big.Int, orderId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "payERC20From", from, amount, orderId, tokenAddress)
}

// PayERC20From is a paid mutator transaction binding the contract method 0x793191f1.
//
// Solidity: function payERC20From(address from, uint256 amount, string orderId, address tokenAddress) returns()
func (_Beewharf *BeewharfSession) PayERC20From(from common.Address, amount *big.Int, orderId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.PayERC20From(&_Beewharf.TransactOpts, from, amount, orderId, tokenAddress)
}

// PayERC20From is a paid mutator transaction binding the contract method 0x793191f1.
//
// Solidity: function payERC20From(address from, uint256 amount, string orderId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactorSession) PayERC20From(from common.Address, amount *big.Int, orderId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.PayERC20From(&_Beewharf.TransactOpts, from, amount, orderId, tokenAddress)
}

// PayEth is a paid mutator transaction binding the contract method 0xaf7d51f5.
//
// Solidity: function payEth(string orderId) payable returns()
func (_Beewharf *BeewharfTransactor) PayEth(opts *bind.TransactOpts, orderId string) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "payEth", orderId)
}

// PayEth is a paid mutator transaction binding the contract method 0xaf7d51f5.
//
// Solidity: function payEth(string orderId) payable returns()
func (_Beewharf *BeewharfSession) PayEth(orderId string) (*types.Transaction, error) {
	return _Beewharf.Contract.PayEth(&_Beewharf.TransactOpts, orderId)
}

// PayEth is a paid mutator transaction binding the contract method 0xaf7d51f5.
//
// Solidity: function payEth(string orderId) payable returns()
func (_Beewharf *BeewharfTransactorSession) PayEth(orderId string) (*types.Transaction, error) {
	return _Beewharf.Contract.PayEth(&_Beewharf.TransactOpts, orderId)
}

// RefundERC20 is a paid mutator transaction binding the contract method 0x70d47134.
//
// Solidity: function refundERC20(address to, uint256 amount, string billId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactor) RefundERC20(opts *bind.TransactOpts, to common.Address, amount *big.Int, billId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "refundERC20", to, amount, billId, tokenAddress)
}

// RefundERC20 is a paid mutator transaction binding the contract method 0x70d47134.
//
// Solidity: function refundERC20(address to, uint256 amount, string billId, address tokenAddress) returns()
func (_Beewharf *BeewharfSession) RefundERC20(to common.Address, amount *big.Int, billId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.RefundERC20(&_Beewharf.TransactOpts, to, amount, billId, tokenAddress)
}

// RefundERC20 is a paid mutator transaction binding the contract method 0x70d47134.
//
// Solidity: function refundERC20(address to, uint256 amount, string billId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactorSession) RefundERC20(to common.Address, amount *big.Int, billId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.RefundERC20(&_Beewharf.TransactOpts, to, amount, billId, tokenAddress)
}

// RefundERC20From is a paid mutator transaction binding the contract method 0xcc86a503.
//
// Solidity: function refundERC20From(address from, address to, uint256 amount, string billId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactor) RefundERC20From(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, billId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "refundERC20From", from, to, amount, billId, tokenAddress)
}

// RefundERC20From is a paid mutator transaction binding the contract method 0xcc86a503.
//
// Solidity: function refundERC20From(address from, address to, uint256 amount, string billId, address tokenAddress) returns()
func (_Beewharf *BeewharfSession) RefundERC20From(from common.Address, to common.Address, amount *big.Int, billId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.RefundERC20From(&_Beewharf.TransactOpts, from, to, amount, billId, tokenAddress)
}

// RefundERC20From is a paid mutator transaction binding the contract method 0xcc86a503.
//
// Solidity: function refundERC20From(address from, address to, uint256 amount, string billId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactorSession) RefundERC20From(from common.Address, to common.Address, amount *big.Int, billId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.RefundERC20From(&_Beewharf.TransactOpts, from, to, amount, billId, tokenAddress)
}

// RefundETH is a paid mutator transaction binding the contract method 0x0290579b.
//
// Solidity: function refundETH(address to, uint256 amount, string billId) returns()
func (_Beewharf *BeewharfTransactor) RefundETH(opts *bind.TransactOpts, to common.Address, amount *big.Int, billId string) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "refundETH", to, amount, billId)
}

// RefundETH is a paid mutator transaction binding the contract method 0x0290579b.
//
// Solidity: function refundETH(address to, uint256 amount, string billId) returns()
func (_Beewharf *BeewharfSession) RefundETH(to common.Address, amount *big.Int, billId string) (*types.Transaction, error) {
	return _Beewharf.Contract.RefundETH(&_Beewharf.TransactOpts, to, amount, billId)
}

// RefundETH is a paid mutator transaction binding the contract method 0x0290579b.
//
// Solidity: function refundETH(address to, uint256 amount, string billId) returns()
func (_Beewharf *BeewharfTransactorSession) RefundETH(to common.Address, amount *big.Int, billId string) (*types.Transaction, error) {
	return _Beewharf.Contract.RefundETH(&_Beewharf.TransactOpts, to, amount, billId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Beewharf *BeewharfTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Beewharf *BeewharfSession) RenounceOwnership() (*types.Transaction, error) {
	return _Beewharf.Contract.RenounceOwnership(&_Beewharf.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Beewharf *BeewharfTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Beewharf.Contract.RenounceOwnership(&_Beewharf.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Beewharf *BeewharfTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Beewharf *BeewharfSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.TransferOwnership(&_Beewharf.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Beewharf *BeewharfTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.TransferOwnership(&_Beewharf.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x58598dad.
//
// Solidity: function withdrawERC20(uint256 amount, string billId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactor) WithdrawERC20(opts *bind.TransactOpts, amount *big.Int, billId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "withdrawERC20", amount, billId, tokenAddress)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x58598dad.
//
// Solidity: function withdrawERC20(uint256 amount, string billId, address tokenAddress) returns()
func (_Beewharf *BeewharfSession) WithdrawERC20(amount *big.Int, billId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.WithdrawERC20(&_Beewharf.TransactOpts, amount, billId, tokenAddress)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x58598dad.
//
// Solidity: function withdrawERC20(uint256 amount, string billId, address tokenAddress) returns()
func (_Beewharf *BeewharfTransactorSession) WithdrawERC20(amount *big.Int, billId string, tokenAddress common.Address) (*types.Transaction, error) {
	return _Beewharf.Contract.WithdrawERC20(&_Beewharf.TransactOpts, amount, billId, tokenAddress)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xe6cd186f.
//
// Solidity: function withdrawEth(uint256 amount, string billId) returns()
func (_Beewharf *BeewharfTransactor) WithdrawEth(opts *bind.TransactOpts, amount *big.Int, billId string) (*types.Transaction, error) {
	return _Beewharf.contract.Transact(opts, "withdrawEth", amount, billId)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xe6cd186f.
//
// Solidity: function withdrawEth(uint256 amount, string billId) returns()
func (_Beewharf *BeewharfSession) WithdrawEth(amount *big.Int, billId string) (*types.Transaction, error) {
	return _Beewharf.Contract.WithdrawEth(&_Beewharf.TransactOpts, amount, billId)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xe6cd186f.
//
// Solidity: function withdrawEth(uint256 amount, string billId) returns()
func (_Beewharf *BeewharfTransactorSession) WithdrawEth(amount *big.Int, billId string) (*types.Transaction, error) {
	return _Beewharf.Contract.WithdrawEth(&_Beewharf.TransactOpts, amount, billId)
}

// BeewharfOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Beewharf contract.
type BeewharfOwnershipTransferredIterator struct {
	Event *BeewharfOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BeewharfOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeewharfOwnershipTransferred)
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
		it.Event = new(BeewharfOwnershipTransferred)
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
func (it *BeewharfOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeewharfOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeewharfOwnershipTransferred represents a OwnershipTransferred event raised by the Beewharf contract.
type BeewharfOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Beewharf *BeewharfFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BeewharfOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Beewharf.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BeewharfOwnershipTransferredIterator{contract: _Beewharf.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Beewharf *BeewharfFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BeewharfOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Beewharf.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeewharfOwnershipTransferred)
				if err := _Beewharf.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Beewharf *BeewharfFilterer) ParseOwnershipTransferred(log types.Log) (*BeewharfOwnershipTransferred, error) {
	event := new(BeewharfOwnershipTransferred)
	if err := _Beewharf.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeewharfPayERC20Iterator is returned from FilterPayERC20 and is used to iterate over the raw logs and unpacked data for PayERC20 events raised by the Beewharf contract.
type BeewharfPayERC20Iterator struct {
	Event *BeewharfPayERC20 // Event containing the contract specifics and raw log

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
func (it *BeewharfPayERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeewharfPayERC20)
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
		it.Event = new(BeewharfPayERC20)
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
func (it *BeewharfPayERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeewharfPayERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeewharfPayERC20 represents a PayERC20 event raised by the Beewharf contract.
type BeewharfPayERC20 struct {
	From         common.Address
	Amount       *big.Int
	OrderId      string
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPayERC20 is a free log retrieval operation binding the contract event 0x5813e31687b98924e93d090022c25c96ef51b871564c433d5823e936dc3b4397.
//
// Solidity: event PayERC20(address indexed from, uint256 amount, string orderId, address indexed tokenAddress)
func (_Beewharf *BeewharfFilterer) FilterPayERC20(opts *bind.FilterOpts, from []common.Address, tokenAddress []common.Address) (*BeewharfPayERC20Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Beewharf.contract.FilterLogs(opts, "PayERC20", fromRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BeewharfPayERC20Iterator{contract: _Beewharf.contract, event: "PayERC20", logs: logs, sub: sub}, nil
}

// WatchPayERC20 is a free log subscription operation binding the contract event 0x5813e31687b98924e93d090022c25c96ef51b871564c433d5823e936dc3b4397.
//
// Solidity: event PayERC20(address indexed from, uint256 amount, string orderId, address indexed tokenAddress)
func (_Beewharf *BeewharfFilterer) WatchPayERC20(opts *bind.WatchOpts, sink chan<- *BeewharfPayERC20, from []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Beewharf.contract.WatchLogs(opts, "PayERC20", fromRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeewharfPayERC20)
				if err := _Beewharf.contract.UnpackLog(event, "PayERC20", log); err != nil {
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

// ParsePayERC20 is a log parse operation binding the contract event 0x5813e31687b98924e93d090022c25c96ef51b871564c433d5823e936dc3b4397.
//
// Solidity: event PayERC20(address indexed from, uint256 amount, string orderId, address indexed tokenAddress)
func (_Beewharf *BeewharfFilterer) ParsePayERC20(log types.Log) (*BeewharfPayERC20, error) {
	event := new(BeewharfPayERC20)
	if err := _Beewharf.contract.UnpackLog(event, "PayERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeewharfPayEthIterator is returned from FilterPayEth and is used to iterate over the raw logs and unpacked data for PayEth events raised by the Beewharf contract.
type BeewharfPayEthIterator struct {
	Event *BeewharfPayEth // Event containing the contract specifics and raw log

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
func (it *BeewharfPayEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeewharfPayEth)
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
		it.Event = new(BeewharfPayEth)
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
func (it *BeewharfPayEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeewharfPayEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeewharfPayEth represents a PayEth event raised by the Beewharf contract.
type BeewharfPayEth struct {
	From    common.Address
	Amount  *big.Int
	OrderId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayEth is a free log retrieval operation binding the contract event 0x0fac5502d5aa0952c2aceb7677cfae0e136494c4b5ca45bb8f0c636ab017efa5.
//
// Solidity: event PayEth(address indexed from, uint256 amount, string orderId)
func (_Beewharf *BeewharfFilterer) FilterPayEth(opts *bind.FilterOpts, from []common.Address) (*BeewharfPayEthIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Beewharf.contract.FilterLogs(opts, "PayEth", fromRule)
	if err != nil {
		return nil, err
	}
	return &BeewharfPayEthIterator{contract: _Beewharf.contract, event: "PayEth", logs: logs, sub: sub}, nil
}

// WatchPayEth is a free log subscription operation binding the contract event 0x0fac5502d5aa0952c2aceb7677cfae0e136494c4b5ca45bb8f0c636ab017efa5.
//
// Solidity: event PayEth(address indexed from, uint256 amount, string orderId)
func (_Beewharf *BeewharfFilterer) WatchPayEth(opts *bind.WatchOpts, sink chan<- *BeewharfPayEth, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Beewharf.contract.WatchLogs(opts, "PayEth", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeewharfPayEth)
				if err := _Beewharf.contract.UnpackLog(event, "PayEth", log); err != nil {
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

// ParsePayEth is a log parse operation binding the contract event 0x0fac5502d5aa0952c2aceb7677cfae0e136494c4b5ca45bb8f0c636ab017efa5.
//
// Solidity: event PayEth(address indexed from, uint256 amount, string orderId)
func (_Beewharf *BeewharfFilterer) ParsePayEth(log types.Log) (*BeewharfPayEth, error) {
	event := new(BeewharfPayEth)
	if err := _Beewharf.contract.UnpackLog(event, "PayEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeewharfRefundERC20Iterator is returned from FilterRefundERC20 and is used to iterate over the raw logs and unpacked data for RefundERC20 events raised by the Beewharf contract.
type BeewharfRefundERC20Iterator struct {
	Event *BeewharfRefundERC20 // Event containing the contract specifics and raw log

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
func (it *BeewharfRefundERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeewharfRefundERC20)
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
		it.Event = new(BeewharfRefundERC20)
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
func (it *BeewharfRefundERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeewharfRefundERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeewharfRefundERC20 represents a RefundERC20 event raised by the Beewharf contract.
type BeewharfRefundERC20 struct {
	To           common.Address
	Amount       *big.Int
	BillId       string
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRefundERC20 is a free log retrieval operation binding the contract event 0x6b4b6b0c19db95f8115cbbc9222c0904280ced3214cae0a128151210a5b6ec30.
//
// Solidity: event RefundERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress)
func (_Beewharf *BeewharfFilterer) FilterRefundERC20(opts *bind.FilterOpts, to []common.Address, tokenAddress []common.Address) (*BeewharfRefundERC20Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Beewharf.contract.FilterLogs(opts, "RefundERC20", toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BeewharfRefundERC20Iterator{contract: _Beewharf.contract, event: "RefundERC20", logs: logs, sub: sub}, nil
}

// WatchRefundERC20 is a free log subscription operation binding the contract event 0x6b4b6b0c19db95f8115cbbc9222c0904280ced3214cae0a128151210a5b6ec30.
//
// Solidity: event RefundERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress)
func (_Beewharf *BeewharfFilterer) WatchRefundERC20(opts *bind.WatchOpts, sink chan<- *BeewharfRefundERC20, to []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Beewharf.contract.WatchLogs(opts, "RefundERC20", toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeewharfRefundERC20)
				if err := _Beewharf.contract.UnpackLog(event, "RefundERC20", log); err != nil {
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

// ParseRefundERC20 is a log parse operation binding the contract event 0x6b4b6b0c19db95f8115cbbc9222c0904280ced3214cae0a128151210a5b6ec30.
//
// Solidity: event RefundERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress)
func (_Beewharf *BeewharfFilterer) ParseRefundERC20(log types.Log) (*BeewharfRefundERC20, error) {
	event := new(BeewharfRefundERC20)
	if err := _Beewharf.contract.UnpackLog(event, "RefundERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeewharfRefundEthIterator is returned from FilterRefundEth and is used to iterate over the raw logs and unpacked data for RefundEth events raised by the Beewharf contract.
type BeewharfRefundEthIterator struct {
	Event *BeewharfRefundEth // Event containing the contract specifics and raw log

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
func (it *BeewharfRefundEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeewharfRefundEth)
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
		it.Event = new(BeewharfRefundEth)
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
func (it *BeewharfRefundEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeewharfRefundEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeewharfRefundEth represents a RefundEth event raised by the Beewharf contract.
type BeewharfRefundEth struct {
	To     common.Address
	Amount *big.Int
	BillId string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundEth is a free log retrieval operation binding the contract event 0xd4ffb8b83d27e6fca3ab24f0e700caca6e812178a99b11e0c04383fa37d4aae2.
//
// Solidity: event RefundEth(address indexed to, uint256 amount, string billId)
func (_Beewharf *BeewharfFilterer) FilterRefundEth(opts *bind.FilterOpts, to []common.Address) (*BeewharfRefundEthIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Beewharf.contract.FilterLogs(opts, "RefundEth", toRule)
	if err != nil {
		return nil, err
	}
	return &BeewharfRefundEthIterator{contract: _Beewharf.contract, event: "RefundEth", logs: logs, sub: sub}, nil
}

// WatchRefundEth is a free log subscription operation binding the contract event 0xd4ffb8b83d27e6fca3ab24f0e700caca6e812178a99b11e0c04383fa37d4aae2.
//
// Solidity: event RefundEth(address indexed to, uint256 amount, string billId)
func (_Beewharf *BeewharfFilterer) WatchRefundEth(opts *bind.WatchOpts, sink chan<- *BeewharfRefundEth, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Beewharf.contract.WatchLogs(opts, "RefundEth", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeewharfRefundEth)
				if err := _Beewharf.contract.UnpackLog(event, "RefundEth", log); err != nil {
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

// ParseRefundEth is a log parse operation binding the contract event 0xd4ffb8b83d27e6fca3ab24f0e700caca6e812178a99b11e0c04383fa37d4aae2.
//
// Solidity: event RefundEth(address indexed to, uint256 amount, string billId)
func (_Beewharf *BeewharfFilterer) ParseRefundEth(log types.Log) (*BeewharfRefundEth, error) {
	event := new(BeewharfRefundEth)
	if err := _Beewharf.contract.UnpackLog(event, "RefundEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeewharfWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the Beewharf contract.
type BeewharfWithdrawERC20Iterator struct {
	Event *BeewharfWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *BeewharfWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeewharfWithdrawERC20)
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
		it.Event = new(BeewharfWithdrawERC20)
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
func (it *BeewharfWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeewharfWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeewharfWithdrawERC20 represents a WithdrawERC20 event raised by the Beewharf contract.
type BeewharfWithdrawERC20 struct {
	To           common.Address
	Amount       *big.Int
	BillId       string
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0x3822da7c1846353da157717dd7e95db54b4ffe105a04f5042cdf2cb1adadb283.
//
// Solidity: event WithdrawERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress)
func (_Beewharf *BeewharfFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, to []common.Address, tokenAddress []common.Address) (*BeewharfWithdrawERC20Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Beewharf.contract.FilterLogs(opts, "WithdrawERC20", toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BeewharfWithdrawERC20Iterator{contract: _Beewharf.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0x3822da7c1846353da157717dd7e95db54b4ffe105a04f5042cdf2cb1adadb283.
//
// Solidity: event WithdrawERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress)
func (_Beewharf *BeewharfFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *BeewharfWithdrawERC20, to []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Beewharf.contract.WatchLogs(opts, "WithdrawERC20", toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeewharfWithdrawERC20)
				if err := _Beewharf.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0x3822da7c1846353da157717dd7e95db54b4ffe105a04f5042cdf2cb1adadb283.
//
// Solidity: event WithdrawERC20(address indexed to, uint256 amount, string billId, address indexed tokenAddress)
func (_Beewharf *BeewharfFilterer) ParseWithdrawERC20(log types.Log) (*BeewharfWithdrawERC20, error) {
	event := new(BeewharfWithdrawERC20)
	if err := _Beewharf.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeewharfWithdrawEthIterator is returned from FilterWithdrawEth and is used to iterate over the raw logs and unpacked data for WithdrawEth events raised by the Beewharf contract.
type BeewharfWithdrawEthIterator struct {
	Event *BeewharfWithdrawEth // Event containing the contract specifics and raw log

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
func (it *BeewharfWithdrawEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeewharfWithdrawEth)
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
		it.Event = new(BeewharfWithdrawEth)
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
func (it *BeewharfWithdrawEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeewharfWithdrawEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeewharfWithdrawEth represents a WithdrawEth event raised by the Beewharf contract.
type BeewharfWithdrawEth struct {
	To     common.Address
	Amount *big.Int
	BillId string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEth is a free log retrieval operation binding the contract event 0xd8f8af852ed19cf5c8faaabb0f17c48f257593ede85efc2f1fe7ba6830de23e9.
//
// Solidity: event WithdrawEth(address indexed to, uint256 amount, string billId)
func (_Beewharf *BeewharfFilterer) FilterWithdrawEth(opts *bind.FilterOpts, to []common.Address) (*BeewharfWithdrawEthIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Beewharf.contract.FilterLogs(opts, "WithdrawEth", toRule)
	if err != nil {
		return nil, err
	}
	return &BeewharfWithdrawEthIterator{contract: _Beewharf.contract, event: "WithdrawEth", logs: logs, sub: sub}, nil
}

// WatchWithdrawEth is a free log subscription operation binding the contract event 0xd8f8af852ed19cf5c8faaabb0f17c48f257593ede85efc2f1fe7ba6830de23e9.
//
// Solidity: event WithdrawEth(address indexed to, uint256 amount, string billId)
func (_Beewharf *BeewharfFilterer) WatchWithdrawEth(opts *bind.WatchOpts, sink chan<- *BeewharfWithdrawEth, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Beewharf.contract.WatchLogs(opts, "WithdrawEth", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeewharfWithdrawEth)
				if err := _Beewharf.contract.UnpackLog(event, "WithdrawEth", log); err != nil {
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

// ParseWithdrawEth is a log parse operation binding the contract event 0xd8f8af852ed19cf5c8faaabb0f17c48f257593ede85efc2f1fe7ba6830de23e9.
//
// Solidity: event WithdrawEth(address indexed to, uint256 amount, string billId)
func (_Beewharf *BeewharfFilterer) ParseWithdrawEth(log types.Log) (*BeewharfWithdrawEth, error) {
	event := new(BeewharfWithdrawEth)
	if err := _Beewharf.contract.UnpackLog(event, "WithdrawEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
