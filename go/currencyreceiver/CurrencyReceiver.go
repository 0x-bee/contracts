// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package currencyreceiver

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

// CurrencyreceiverMetaData contains all meta data concerning the Currencyreceiver contract.
var CurrencyreceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUND_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"currency\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"orderId\",\"type\":\"string[]\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"currency\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"billId\",\"type\":\"string[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CurrencyreceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use CurrencyreceiverMetaData.ABI instead.
var CurrencyreceiverABI = CurrencyreceiverMetaData.ABI

// Currencyreceiver is an auto generated Go binding around an Ethereum contract.
type Currencyreceiver struct {
	CurrencyreceiverCaller     // Read-only binding to the contract
	CurrencyreceiverTransactor // Write-only binding to the contract
	CurrencyreceiverFilterer   // Log filterer for contract events
}

// CurrencyreceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurrencyreceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrencyreceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurrencyreceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrencyreceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurrencyreceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrencyreceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurrencyreceiverSession struct {
	Contract     *Currencyreceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurrencyreceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurrencyreceiverCallerSession struct {
	Contract *CurrencyreceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CurrencyreceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurrencyreceiverTransactorSession struct {
	Contract     *CurrencyreceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CurrencyreceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurrencyreceiverRaw struct {
	Contract *Currencyreceiver // Generic contract binding to access the raw methods on
}

// CurrencyreceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurrencyreceiverCallerRaw struct {
	Contract *CurrencyreceiverCaller // Generic read-only contract binding to access the raw methods on
}

// CurrencyreceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurrencyreceiverTransactorRaw struct {
	Contract *CurrencyreceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurrencyreceiver creates a new instance of Currencyreceiver, bound to a specific deployed contract.
func NewCurrencyreceiver(address common.Address, backend bind.ContractBackend) (*Currencyreceiver, error) {
	contract, err := bindCurrencyreceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Currencyreceiver{CurrencyreceiverCaller: CurrencyreceiverCaller{contract: contract}, CurrencyreceiverTransactor: CurrencyreceiverTransactor{contract: contract}, CurrencyreceiverFilterer: CurrencyreceiverFilterer{contract: contract}}, nil
}

// NewCurrencyreceiverCaller creates a new read-only instance of Currencyreceiver, bound to a specific deployed contract.
func NewCurrencyreceiverCaller(address common.Address, caller bind.ContractCaller) (*CurrencyreceiverCaller, error) {
	contract, err := bindCurrencyreceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverCaller{contract: contract}, nil
}

// NewCurrencyreceiverTransactor creates a new write-only instance of Currencyreceiver, bound to a specific deployed contract.
func NewCurrencyreceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*CurrencyreceiverTransactor, error) {
	contract, err := bindCurrencyreceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverTransactor{contract: contract}, nil
}

// NewCurrencyreceiverFilterer creates a new log filterer instance of Currencyreceiver, bound to a specific deployed contract.
func NewCurrencyreceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*CurrencyreceiverFilterer, error) {
	contract, err := bindCurrencyreceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverFilterer{contract: contract}, nil
}

// bindCurrencyreceiver binds a generic wrapper to an already deployed contract.
func bindCurrencyreceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurrencyreceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Currencyreceiver *CurrencyreceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Currencyreceiver.Contract.CurrencyreceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Currencyreceiver *CurrencyreceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.CurrencyreceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Currencyreceiver *CurrencyreceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.CurrencyreceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Currencyreceiver *CurrencyreceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Currencyreceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Currencyreceiver *CurrencyreceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Currencyreceiver *CurrencyreceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Currencyreceiver.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Currencyreceiver.Contract.DEFAULTADMINROLE(&_Currencyreceiver.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Currencyreceiver.Contract.DEFAULTADMINROLE(&_Currencyreceiver.CallOpts)
}

// REFUNDROLE is a free data retrieval call binding the contract method 0x627c0906.
//
// Solidity: function REFUND_ROLE() view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverCaller) REFUNDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Currencyreceiver.contract.Call(opts, &out, "REFUND_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REFUNDROLE is a free data retrieval call binding the contract method 0x627c0906.
//
// Solidity: function REFUND_ROLE() view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverSession) REFUNDROLE() ([32]byte, error) {
	return _Currencyreceiver.Contract.REFUNDROLE(&_Currencyreceiver.CallOpts)
}

// REFUNDROLE is a free data retrieval call binding the contract method 0x627c0906.
//
// Solidity: function REFUND_ROLE() view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverCallerSession) REFUNDROLE() ([32]byte, error) {
	return _Currencyreceiver.Contract.REFUNDROLE(&_Currencyreceiver.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverCaller) WITHDRAWROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Currencyreceiver.contract.Call(opts, &out, "WITHDRAW_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverSession) WITHDRAWROLE() ([32]byte, error) {
	return _Currencyreceiver.Contract.WITHDRAWROLE(&_Currencyreceiver.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverCallerSession) WITHDRAWROLE() ([32]byte, error) {
	return _Currencyreceiver.Contract.WITHDRAWROLE(&_Currencyreceiver.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address currency) view returns(uint256)
func (_Currencyreceiver *CurrencyreceiverCaller) BalanceOf(opts *bind.CallOpts, currency common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Currencyreceiver.contract.Call(opts, &out, "balanceOf", currency)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address currency) view returns(uint256)
func (_Currencyreceiver *CurrencyreceiverSession) BalanceOf(currency common.Address) (*big.Int, error) {
	return _Currencyreceiver.Contract.BalanceOf(&_Currencyreceiver.CallOpts, currency)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address currency) view returns(uint256)
func (_Currencyreceiver *CurrencyreceiverCallerSession) BalanceOf(currency common.Address) (*big.Int, error) {
	return _Currencyreceiver.Contract.BalanceOf(&_Currencyreceiver.CallOpts, currency)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Currencyreceiver.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Currencyreceiver.Contract.GetRoleAdmin(&_Currencyreceiver.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Currencyreceiver *CurrencyreceiverCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Currencyreceiver.Contract.GetRoleAdmin(&_Currencyreceiver.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Currencyreceiver *CurrencyreceiverCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Currencyreceiver.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Currencyreceiver *CurrencyreceiverSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Currencyreceiver.Contract.HasRole(&_Currencyreceiver.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Currencyreceiver *CurrencyreceiverCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Currencyreceiver.Contract.HasRole(&_Currencyreceiver.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Currencyreceiver *CurrencyreceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Currencyreceiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Currencyreceiver *CurrencyreceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Currencyreceiver.Contract.SupportsInterface(&_Currencyreceiver.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Currencyreceiver *CurrencyreceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Currencyreceiver.Contract.SupportsInterface(&_Currencyreceiver.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Currencyreceiver *CurrencyreceiverTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Currencyreceiver *CurrencyreceiverSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.GrantRole(&_Currencyreceiver.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Currencyreceiver *CurrencyreceiverTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.GrantRole(&_Currencyreceiver.TransactOpts, role, account)
}

// Pay is a paid mutator transaction binding the contract method 0x4a4bdb30.
//
// Solidity: function pay(address currency, uint256 amount, string orderId) returns()
func (_Currencyreceiver *CurrencyreceiverTransactor) Pay(opts *bind.TransactOpts, currency common.Address, amount *big.Int, orderId string) (*types.Transaction, error) {
	return _Currencyreceiver.contract.Transact(opts, "pay", currency, amount, orderId)
}

// Pay is a paid mutator transaction binding the contract method 0x4a4bdb30.
//
// Solidity: function pay(address currency, uint256 amount, string orderId) returns()
func (_Currencyreceiver *CurrencyreceiverSession) Pay(currency common.Address, amount *big.Int, orderId string) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.Pay(&_Currencyreceiver.TransactOpts, currency, amount, orderId)
}

// Pay is a paid mutator transaction binding the contract method 0x4a4bdb30.
//
// Solidity: function pay(address currency, uint256 amount, string orderId) returns()
func (_Currencyreceiver *CurrencyreceiverTransactorSession) Pay(currency common.Address, amount *big.Int, orderId string) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.Pay(&_Currencyreceiver.TransactOpts, currency, amount, orderId)
}

// Refund is a paid mutator transaction binding the contract method 0x56edb469.
//
// Solidity: function refund(address[] currency, uint256[] amount, address[] to, string[] orderId) returns()
func (_Currencyreceiver *CurrencyreceiverTransactor) Refund(opts *bind.TransactOpts, currency []common.Address, amount []*big.Int, to []common.Address, orderId []string) (*types.Transaction, error) {
	return _Currencyreceiver.contract.Transact(opts, "refund", currency, amount, to, orderId)
}

// Refund is a paid mutator transaction binding the contract method 0x56edb469.
//
// Solidity: function refund(address[] currency, uint256[] amount, address[] to, string[] orderId) returns()
func (_Currencyreceiver *CurrencyreceiverSession) Refund(currency []common.Address, amount []*big.Int, to []common.Address, orderId []string) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.Refund(&_Currencyreceiver.TransactOpts, currency, amount, to, orderId)
}

// Refund is a paid mutator transaction binding the contract method 0x56edb469.
//
// Solidity: function refund(address[] currency, uint256[] amount, address[] to, string[] orderId) returns()
func (_Currencyreceiver *CurrencyreceiverTransactorSession) Refund(currency []common.Address, amount []*big.Int, to []common.Address, orderId []string) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.Refund(&_Currencyreceiver.TransactOpts, currency, amount, to, orderId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Currencyreceiver *CurrencyreceiverTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Currencyreceiver *CurrencyreceiverSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.RenounceRole(&_Currencyreceiver.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Currencyreceiver *CurrencyreceiverTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.RenounceRole(&_Currencyreceiver.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Currencyreceiver *CurrencyreceiverTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Currencyreceiver *CurrencyreceiverSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.RevokeRole(&_Currencyreceiver.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Currencyreceiver *CurrencyreceiverTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.RevokeRole(&_Currencyreceiver.TransactOpts, role, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0xea4e6c5c.
//
// Solidity: function withdraw(address[] currency, uint256[] amount, address[] to, string[] billId) returns()
func (_Currencyreceiver *CurrencyreceiverTransactor) Withdraw(opts *bind.TransactOpts, currency []common.Address, amount []*big.Int, to []common.Address, billId []string) (*types.Transaction, error) {
	return _Currencyreceiver.contract.Transact(opts, "withdraw", currency, amount, to, billId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xea4e6c5c.
//
// Solidity: function withdraw(address[] currency, uint256[] amount, address[] to, string[] billId) returns()
func (_Currencyreceiver *CurrencyreceiverSession) Withdraw(currency []common.Address, amount []*big.Int, to []common.Address, billId []string) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.Withdraw(&_Currencyreceiver.TransactOpts, currency, amount, to, billId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xea4e6c5c.
//
// Solidity: function withdraw(address[] currency, uint256[] amount, address[] to, string[] billId) returns()
func (_Currencyreceiver *CurrencyreceiverTransactorSession) Withdraw(currency []common.Address, amount []*big.Int, to []common.Address, billId []string) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.Withdraw(&_Currencyreceiver.TransactOpts, currency, amount, to, billId)
}

// CurrencyreceiverPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the Currencyreceiver contract.
type CurrencyreceiverPayIterator struct {
	Event *CurrencyreceiverPay // Event containing the contract specifics and raw log

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
func (it *CurrencyreceiverPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrencyreceiverPay)
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
		it.Event = new(CurrencyreceiverPay)
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
func (it *CurrencyreceiverPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrencyreceiverPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrencyreceiverPay represents a Pay event raised by the Currencyreceiver contract.
type CurrencyreceiverPay struct {
	Currency common.Address
	From     common.Address
	Amount   *big.Int
	OrderId  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x5d2ae1cdca5222ec8fffbc4b088f5b5144ad5f4eaf2fc8afc49cd35857671082.
//
// Solidity: event Pay(address indexed currency, address indexed from, uint256 amount, string orderId)
func (_Currencyreceiver *CurrencyreceiverFilterer) FilterPay(opts *bind.FilterOpts, currency []common.Address, from []common.Address) (*CurrencyreceiverPayIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Currencyreceiver.contract.FilterLogs(opts, "Pay", currencyRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverPayIterator{contract: _Currencyreceiver.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x5d2ae1cdca5222ec8fffbc4b088f5b5144ad5f4eaf2fc8afc49cd35857671082.
//
// Solidity: event Pay(address indexed currency, address indexed from, uint256 amount, string orderId)
func (_Currencyreceiver *CurrencyreceiverFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *CurrencyreceiverPay, currency []common.Address, from []common.Address) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Currencyreceiver.contract.WatchLogs(opts, "Pay", currencyRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrencyreceiverPay)
				if err := _Currencyreceiver.contract.UnpackLog(event, "Pay", log); err != nil {
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

// ParsePay is a log parse operation binding the contract event 0x5d2ae1cdca5222ec8fffbc4b088f5b5144ad5f4eaf2fc8afc49cd35857671082.
//
// Solidity: event Pay(address indexed currency, address indexed from, uint256 amount, string orderId)
func (_Currencyreceiver *CurrencyreceiverFilterer) ParsePay(log types.Log) (*CurrencyreceiverPay, error) {
	event := new(CurrencyreceiverPay)
	if err := _Currencyreceiver.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrencyreceiverRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the Currencyreceiver contract.
type CurrencyreceiverRefundIterator struct {
	Event *CurrencyreceiverRefund // Event containing the contract specifics and raw log

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
func (it *CurrencyreceiverRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrencyreceiverRefund)
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
		it.Event = new(CurrencyreceiverRefund)
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
func (it *CurrencyreceiverRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrencyreceiverRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrencyreceiverRefund represents a Refund event raised by the Currencyreceiver contract.
type CurrencyreceiverRefund struct {
	Currency common.Address
	Amount   *big.Int
	To       common.Address
	OrderId  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x1a50ff908577c6023d20715c654df4d0ac65bbb4ee6422490751a0489f85dc67.
//
// Solidity: event Refund(address indexed currency, uint256 amount, address indexed to, string orderId)
func (_Currencyreceiver *CurrencyreceiverFilterer) FilterRefund(opts *bind.FilterOpts, currency []common.Address, to []common.Address) (*CurrencyreceiverRefundIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Currencyreceiver.contract.FilterLogs(opts, "Refund", currencyRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverRefundIterator{contract: _Currencyreceiver.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x1a50ff908577c6023d20715c654df4d0ac65bbb4ee6422490751a0489f85dc67.
//
// Solidity: event Refund(address indexed currency, uint256 amount, address indexed to, string orderId)
func (_Currencyreceiver *CurrencyreceiverFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *CurrencyreceiverRefund, currency []common.Address, to []common.Address) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Currencyreceiver.contract.WatchLogs(opts, "Refund", currencyRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrencyreceiverRefund)
				if err := _Currencyreceiver.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x1a50ff908577c6023d20715c654df4d0ac65bbb4ee6422490751a0489f85dc67.
//
// Solidity: event Refund(address indexed currency, uint256 amount, address indexed to, string orderId)
func (_Currencyreceiver *CurrencyreceiverFilterer) ParseRefund(log types.Log) (*CurrencyreceiverRefund, error) {
	event := new(CurrencyreceiverRefund)
	if err := _Currencyreceiver.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrencyreceiverRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Currencyreceiver contract.
type CurrencyreceiverRoleAdminChangedIterator struct {
	Event *CurrencyreceiverRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CurrencyreceiverRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrencyreceiverRoleAdminChanged)
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
		it.Event = new(CurrencyreceiverRoleAdminChanged)
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
func (it *CurrencyreceiverRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrencyreceiverRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrencyreceiverRoleAdminChanged represents a RoleAdminChanged event raised by the Currencyreceiver contract.
type CurrencyreceiverRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Currencyreceiver *CurrencyreceiverFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CurrencyreceiverRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Currencyreceiver.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverRoleAdminChangedIterator{contract: _Currencyreceiver.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Currencyreceiver *CurrencyreceiverFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CurrencyreceiverRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Currencyreceiver.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrencyreceiverRoleAdminChanged)
				if err := _Currencyreceiver.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Currencyreceiver *CurrencyreceiverFilterer) ParseRoleAdminChanged(log types.Log) (*CurrencyreceiverRoleAdminChanged, error) {
	event := new(CurrencyreceiverRoleAdminChanged)
	if err := _Currencyreceiver.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrencyreceiverRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Currencyreceiver contract.
type CurrencyreceiverRoleGrantedIterator struct {
	Event *CurrencyreceiverRoleGranted // Event containing the contract specifics and raw log

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
func (it *CurrencyreceiverRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrencyreceiverRoleGranted)
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
		it.Event = new(CurrencyreceiverRoleGranted)
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
func (it *CurrencyreceiverRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrencyreceiverRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrencyreceiverRoleGranted represents a RoleGranted event raised by the Currencyreceiver contract.
type CurrencyreceiverRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currencyreceiver *CurrencyreceiverFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CurrencyreceiverRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Currencyreceiver.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverRoleGrantedIterator{contract: _Currencyreceiver.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currencyreceiver *CurrencyreceiverFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CurrencyreceiverRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Currencyreceiver.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrencyreceiverRoleGranted)
				if err := _Currencyreceiver.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currencyreceiver *CurrencyreceiverFilterer) ParseRoleGranted(log types.Log) (*CurrencyreceiverRoleGranted, error) {
	event := new(CurrencyreceiverRoleGranted)
	if err := _Currencyreceiver.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrencyreceiverRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Currencyreceiver contract.
type CurrencyreceiverRoleRevokedIterator struct {
	Event *CurrencyreceiverRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CurrencyreceiverRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrencyreceiverRoleRevoked)
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
		it.Event = new(CurrencyreceiverRoleRevoked)
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
func (it *CurrencyreceiverRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrencyreceiverRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrencyreceiverRoleRevoked represents a RoleRevoked event raised by the Currencyreceiver contract.
type CurrencyreceiverRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currencyreceiver *CurrencyreceiverFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CurrencyreceiverRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Currencyreceiver.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverRoleRevokedIterator{contract: _Currencyreceiver.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currencyreceiver *CurrencyreceiverFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CurrencyreceiverRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Currencyreceiver.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrencyreceiverRoleRevoked)
				if err := _Currencyreceiver.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currencyreceiver *CurrencyreceiverFilterer) ParseRoleRevoked(log types.Log) (*CurrencyreceiverRoleRevoked, error) {
	event := new(CurrencyreceiverRoleRevoked)
	if err := _Currencyreceiver.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrencyreceiverWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Currencyreceiver contract.
type CurrencyreceiverWithdrawIterator struct {
	Event *CurrencyreceiverWithdraw // Event containing the contract specifics and raw log

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
func (it *CurrencyreceiverWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrencyreceiverWithdraw)
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
		it.Event = new(CurrencyreceiverWithdraw)
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
func (it *CurrencyreceiverWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrencyreceiverWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrencyreceiverWithdraw represents a Withdraw event raised by the Currencyreceiver contract.
type CurrencyreceiverWithdraw struct {
	Currency common.Address
	Amount   *big.Int
	To       common.Address
	BillId   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xa4c6cd4bfefcc09490a00cf1f79a859de6f34c1da3186bb65d5102b1b8445547.
//
// Solidity: event Withdraw(address indexed currency, uint256 amount, address indexed to, string billId)
func (_Currencyreceiver *CurrencyreceiverFilterer) FilterWithdraw(opts *bind.FilterOpts, currency []common.Address, to []common.Address) (*CurrencyreceiverWithdrawIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Currencyreceiver.contract.FilterLogs(opts, "Withdraw", currencyRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverWithdrawIterator{contract: _Currencyreceiver.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xa4c6cd4bfefcc09490a00cf1f79a859de6f34c1da3186bb65d5102b1b8445547.
//
// Solidity: event Withdraw(address indexed currency, uint256 amount, address indexed to, string billId)
func (_Currencyreceiver *CurrencyreceiverFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CurrencyreceiverWithdraw, currency []common.Address, to []common.Address) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Currencyreceiver.contract.WatchLogs(opts, "Withdraw", currencyRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrencyreceiverWithdraw)
				if err := _Currencyreceiver.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xa4c6cd4bfefcc09490a00cf1f79a859de6f34c1da3186bb65d5102b1b8445547.
//
// Solidity: event Withdraw(address indexed currency, uint256 amount, address indexed to, string billId)
func (_Currencyreceiver *CurrencyreceiverFilterer) ParseWithdraw(log types.Log) (*CurrencyreceiverWithdraw, error) {
	event := new(CurrencyreceiverWithdraw)
	if err := _Currencyreceiver.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
