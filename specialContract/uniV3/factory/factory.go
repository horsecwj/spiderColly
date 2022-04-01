// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniV3

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

// UniV3MetaData contains all meta data concerning the UniV3 contract.
var UniV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"FeeAmountEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniV3MetaData.ABI instead.
var UniV3ABI = UniV3MetaData.ABI

// UniV3 is an auto generated Go binding around an Ethereum contract.
type UniV3 struct {
	UniV3Caller     // Read-only binding to the contract
	UniV3Transactor // Write-only binding to the contract
	UniV3Filterer   // Log filterer for contract events
}

// UniV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniV3Session struct {
	Contract     *UniV3            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniV3CallerSession struct {
	Contract *UniV3Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UniV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniV3TransactorSession struct {
	Contract     *UniV3Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniV3Raw struct {
	Contract *UniV3 // Generic contract binding to access the raw methods on
}

// UniV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniV3CallerRaw struct {
	Contract *UniV3Caller // Generic read-only contract binding to access the raw methods on
}

// UniV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniV3TransactorRaw struct {
	Contract *UniV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniV3 creates a new instance of UniV3, bound to a specific deployed contract.
func NewUniV3(address common.Address, backend bind.ContractBackend) (*UniV3, error) {
	contract, err := bindUniV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniV3{UniV3Caller: UniV3Caller{contract: contract}, UniV3Transactor: UniV3Transactor{contract: contract}, UniV3Filterer: UniV3Filterer{contract: contract}}, nil
}

// NewUniV3Caller creates a new read-only instance of UniV3, bound to a specific deployed contract.
func NewUniV3Caller(address common.Address, caller bind.ContractCaller) (*UniV3Caller, error) {
	contract, err := bindUniV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniV3Caller{contract: contract}, nil
}

// NewUniV3Transactor creates a new write-only instance of UniV3, bound to a specific deployed contract.
func NewUniV3Transactor(address common.Address, transactor bind.ContractTransactor) (*UniV3Transactor, error) {
	contract, err := bindUniV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniV3Transactor{contract: contract}, nil
}

// NewUniV3Filterer creates a new log filterer instance of UniV3, bound to a specific deployed contract.
func NewUniV3Filterer(address common.Address, filterer bind.ContractFilterer) (*UniV3Filterer, error) {
	contract, err := bindUniV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniV3Filterer{contract: contract}, nil
}

// bindUniV3 binds a generic wrapper to an already deployed contract.
func bindUniV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniV3 *UniV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniV3.Contract.UniV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniV3 *UniV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniV3.Contract.UniV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniV3 *UniV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniV3.Contract.UniV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniV3 *UniV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniV3 *UniV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniV3 *UniV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniV3.Contract.contract.Transact(opts, method, params...)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_UniV3 *UniV3Caller) FeeAmountTickSpacing(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniV3.contract.Call(opts, &out, "feeAmountTickSpacing", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_UniV3 *UniV3Session) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _UniV3.Contract.FeeAmountTickSpacing(&_UniV3.CallOpts, arg0)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_UniV3 *UniV3CallerSession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _UniV3.Contract.FeeAmountTickSpacing(&_UniV3.CallOpts, arg0)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniV3 *UniV3Caller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniV3.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniV3 *UniV3Session) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _UniV3.Contract.GetPool(&_UniV3.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_UniV3 *UniV3CallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _UniV3.Contract.GetPool(&_UniV3.CallOpts, arg0, arg1, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniV3 *UniV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniV3 *UniV3Session) Owner() (common.Address, error) {
	return _UniV3.Contract.Owner(&_UniV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniV3 *UniV3CallerSession) Owner() (common.Address, error) {
	return _UniV3.Contract.Owner(&_UniV3.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniV3 *UniV3Caller) Parameters(opts *bind.CallOpts) (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	var out []interface{}
	err := _UniV3.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		Factory     common.Address
		Token0      common.Address
		Token1      common.Address
		Fee         *big.Int
		TickSpacing *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickSpacing = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniV3 *UniV3Session) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _UniV3.Contract.Parameters(&_UniV3.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_UniV3 *UniV3CallerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _UniV3.Contract.Parameters(&_UniV3.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_UniV3 *UniV3Transactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _UniV3.contract.Transact(opts, "createPool", tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_UniV3 *UniV3Session) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _UniV3.Contract.CreatePool(&_UniV3.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_UniV3 *UniV3TransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _UniV3.Contract.CreatePool(&_UniV3.TransactOpts, tokenA, tokenB, fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_UniV3 *UniV3Transactor) EnableFeeAmount(opts *bind.TransactOpts, fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _UniV3.contract.Transact(opts, "enableFeeAmount", fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_UniV3 *UniV3Session) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _UniV3.Contract.EnableFeeAmount(&_UniV3.TransactOpts, fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_UniV3 *UniV3TransactorSession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _UniV3.Contract.EnableFeeAmount(&_UniV3.TransactOpts, fee, tickSpacing)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_UniV3 *UniV3Transactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _UniV3.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_UniV3 *UniV3Session) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _UniV3.Contract.SetOwner(&_UniV3.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_UniV3 *UniV3TransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _UniV3.Contract.SetOwner(&_UniV3.TransactOpts, _owner)
}

// UniV3FeeAmountEnabledIterator is returned from FilterFeeAmountEnabled and is used to iterate over the raw logs and unpacked data for FeeAmountEnabled events raised by the UniV3 contract.
type UniV3FeeAmountEnabledIterator struct {
	Event *UniV3FeeAmountEnabled // Event containing the contract specifics and raw log

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
func (it *UniV3FeeAmountEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniV3FeeAmountEnabled)
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
		it.Event = new(UniV3FeeAmountEnabled)
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
func (it *UniV3FeeAmountEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniV3FeeAmountEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniV3FeeAmountEnabled represents a FeeAmountEnabled event raised by the UniV3 contract.
type UniV3FeeAmountEnabled struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountEnabled is a free log retrieval operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_UniV3 *UniV3Filterer) FilterFeeAmountEnabled(opts *bind.FilterOpts, fee []*big.Int, tickSpacing []*big.Int) (*UniV3FeeAmountEnabledIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _UniV3.contract.FilterLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return &UniV3FeeAmountEnabledIterator{contract: _UniV3.contract, event: "FeeAmountEnabled", logs: logs, sub: sub}, nil
}

// WatchFeeAmountEnabled is a free log subscription operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_UniV3 *UniV3Filterer) WatchFeeAmountEnabled(opts *bind.WatchOpts, sink chan<- *UniV3FeeAmountEnabled, fee []*big.Int, tickSpacing []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _UniV3.contract.WatchLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniV3FeeAmountEnabled)
				if err := _UniV3.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
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

// ParseFeeAmountEnabled is a log parse operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_UniV3 *UniV3Filterer) ParseFeeAmountEnabled(log types.Log) (*UniV3FeeAmountEnabled, error) {
	event := new(UniV3FeeAmountEnabled)
	if err := _UniV3.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniV3OwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the UniV3 contract.
type UniV3OwnerChangedIterator struct {
	Event *UniV3OwnerChanged // Event containing the contract specifics and raw log

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
func (it *UniV3OwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniV3OwnerChanged)
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
		it.Event = new(UniV3OwnerChanged)
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
func (it *UniV3OwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniV3OwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniV3OwnerChanged represents a OwnerChanged event raised by the UniV3 contract.
type UniV3OwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_UniV3 *UniV3Filterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*UniV3OwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniV3.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniV3OwnerChangedIterator{contract: _UniV3.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_UniV3 *UniV3Filterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *UniV3OwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniV3.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniV3OwnerChanged)
				if err := _UniV3.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_UniV3 *UniV3Filterer) ParseOwnerChanged(log types.Log) (*UniV3OwnerChanged, error) {
	event := new(UniV3OwnerChanged)
	if err := _UniV3.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniV3PoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the UniV3 contract.
type UniV3PoolCreatedIterator struct {
	Event *UniV3PoolCreated // Event containing the contract specifics and raw log

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
func (it *UniV3PoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniV3PoolCreated)
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
		it.Event = new(UniV3PoolCreated)
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
func (it *UniV3PoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniV3PoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniV3PoolCreated represents a PoolCreated event raised by the UniV3 contract.
type UniV3PoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_UniV3 *UniV3Filterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, fee []*big.Int) (*UniV3PoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _UniV3.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return &UniV3PoolCreatedIterator{contract: _UniV3.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_UniV3 *UniV3Filterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *UniV3PoolCreated, token0 []common.Address, token1 []common.Address, fee []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _UniV3.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniV3PoolCreated)
				if err := _UniV3.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_UniV3 *UniV3Filterer) ParsePoolCreated(log types.Log) (*UniV3PoolCreated, error) {
	event := new(UniV3PoolCreated)
	if err := _UniV3.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
