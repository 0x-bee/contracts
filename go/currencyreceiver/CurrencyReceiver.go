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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"currency\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"orderIds\",\"type\":\"string[]\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"currency\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"currency\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"orderId\",\"type\":\"string[]\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"currency\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"billId\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Currencyreceiver *CurrencyreceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Currencyreceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Currencyreceiver *CurrencyreceiverSession) Owner() (common.Address, error) {
	return _Currencyreceiver.Contract.Owner(&_Currencyreceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Currencyreceiver *CurrencyreceiverCallerSession) Owner() (common.Address, error) {
	return _Currencyreceiver.Contract.Owner(&_Currencyreceiver.CallOpts)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Currencyreceiver *CurrencyreceiverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currencyreceiver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Currencyreceiver *CurrencyreceiverSession) RenounceOwnership() (*types.Transaction, error) {
	return _Currencyreceiver.Contract.RenounceOwnership(&_Currencyreceiver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Currencyreceiver *CurrencyreceiverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Currencyreceiver.Contract.RenounceOwnership(&_Currencyreceiver.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Currencyreceiver *CurrencyreceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Currencyreceiver *CurrencyreceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.TransferOwnership(&_Currencyreceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Currencyreceiver *CurrencyreceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.TransferOwnership(&_Currencyreceiver.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x17e78b6c.
//
// Solidity: function withdraw(address[] currency, uint256[] amount, string billId) returns()
func (_Currencyreceiver *CurrencyreceiverTransactor) Withdraw(opts *bind.TransactOpts, currency []common.Address, amount []*big.Int, billId string) (*types.Transaction, error) {
	return _Currencyreceiver.contract.Transact(opts, "withdraw", currency, amount, billId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x17e78b6c.
//
// Solidity: function withdraw(address[] currency, uint256[] amount, string billId) returns()
func (_Currencyreceiver *CurrencyreceiverSession) Withdraw(currency []common.Address, amount []*big.Int, billId string) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.Withdraw(&_Currencyreceiver.TransactOpts, currency, amount, billId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x17e78b6c.
//
// Solidity: function withdraw(address[] currency, uint256[] amount, string billId) returns()
func (_Currencyreceiver *CurrencyreceiverTransactorSession) Withdraw(currency []common.Address, amount []*big.Int, billId string) (*types.Transaction, error) {
	return _Currencyreceiver.Contract.Withdraw(&_Currencyreceiver.TransactOpts, currency, amount, billId)
}

// CurrencyreceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Currencyreceiver contract.
type CurrencyreceiverOwnershipTransferredIterator struct {
	Event *CurrencyreceiverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CurrencyreceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrencyreceiverOwnershipTransferred)
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
		it.Event = new(CurrencyreceiverOwnershipTransferred)
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
func (it *CurrencyreceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrencyreceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrencyreceiverOwnershipTransferred represents a OwnershipTransferred event raised by the Currencyreceiver contract.
type CurrencyreceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Currencyreceiver *CurrencyreceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CurrencyreceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Currencyreceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CurrencyreceiverOwnershipTransferredIterator{contract: _Currencyreceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Currencyreceiver *CurrencyreceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CurrencyreceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Currencyreceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrencyreceiverOwnershipTransferred)
				if err := _Currencyreceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Currencyreceiver *CurrencyreceiverFilterer) ParseOwnershipTransferred(log types.Log) (*CurrencyreceiverOwnershipTransferred, error) {
	event := new(CurrencyreceiverOwnershipTransferred)
	if err := _Currencyreceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Currency []common.Address
	Amount   []*big.Int
	To       []common.Address
	OrderIds []string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xa00a8aecfc1929afaf83d428e7b217b6626396c9dd39ad9f5d7327a0bb5e4d13.
//
// Solidity: event Refund(address[] indexed currency, uint256[] amount, address[] indexed to, string[] orderIds)
func (_Currencyreceiver *CurrencyreceiverFilterer) FilterRefund(opts *bind.FilterOpts, currency [][]common.Address, to [][]common.Address) (*CurrencyreceiverRefundIterator, error) {

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

// WatchRefund is a free log subscription operation binding the contract event 0xa00a8aecfc1929afaf83d428e7b217b6626396c9dd39ad9f5d7327a0bb5e4d13.
//
// Solidity: event Refund(address[] indexed currency, uint256[] amount, address[] indexed to, string[] orderIds)
func (_Currencyreceiver *CurrencyreceiverFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *CurrencyreceiverRefund, currency [][]common.Address, to [][]common.Address) (event.Subscription, error) {

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

// ParseRefund is a log parse operation binding the contract event 0xa00a8aecfc1929afaf83d428e7b217b6626396c9dd39ad9f5d7327a0bb5e4d13.
//
// Solidity: event Refund(address[] indexed currency, uint256[] amount, address[] indexed to, string[] orderIds)
func (_Currencyreceiver *CurrencyreceiverFilterer) ParseRefund(log types.Log) (*CurrencyreceiverRefund, error) {
	event := new(CurrencyreceiverRefund)
	if err := _Currencyreceiver.contract.UnpackLog(event, "Refund", log); err != nil {
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
	Currency []common.Address
	Amount   []*big.Int
	To       common.Address
	BillId   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x246131af8b217ee0901b4b5a97b3dc7ebf85c4565198d26948c19055597fc002.
//
// Solidity: event Withdraw(address[] indexed currency, uint256[] amount, address indexed to, string billId)
func (_Currencyreceiver *CurrencyreceiverFilterer) FilterWithdraw(opts *bind.FilterOpts, currency [][]common.Address, to []common.Address) (*CurrencyreceiverWithdrawIterator, error) {

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

// WatchWithdraw is a free log subscription operation binding the contract event 0x246131af8b217ee0901b4b5a97b3dc7ebf85c4565198d26948c19055597fc002.
//
// Solidity: event Withdraw(address[] indexed currency, uint256[] amount, address indexed to, string billId)
func (_Currencyreceiver *CurrencyreceiverFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CurrencyreceiverWithdraw, currency [][]common.Address, to []common.Address) (event.Subscription, error) {

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

// ParseWithdraw is a log parse operation binding the contract event 0x246131af8b217ee0901b4b5a97b3dc7ebf85c4565198d26948c19055597fc002.
//
// Solidity: event Withdraw(address[] indexed currency, uint256[] amount, address indexed to, string billId)
func (_Currencyreceiver *CurrencyreceiverFilterer) ParseWithdraw(log types.Log) (*CurrencyreceiverWithdraw, error) {
	event := new(CurrencyreceiverWithdraw)
	if err := _Currencyreceiver.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
