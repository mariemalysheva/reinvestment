// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_nft_factory

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

// ContractNftFactoryMetaData contains all meta data concerning the ContractNftFactory contract.
var ContractNftFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_album\",\"type\":\"address\"}],\"name\":\"AlbumCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"createAlbum\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractNftFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractNftFactoryMetaData.ABI instead.
var ContractNftFactoryABI = ContractNftFactoryMetaData.ABI

// ContractNftFactory is an auto generated Go binding around an Ethereum contract.
type ContractNftFactory struct {
	ContractNftFactoryCaller     // Read-only binding to the contract
	ContractNftFactoryTransactor // Write-only binding to the contract
	ContractNftFactoryFilterer   // Log filterer for contract events
}

// ContractNftFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractNftFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractNftFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractNftFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractNftFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractNftFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractNftFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractNftFactorySession struct {
	Contract     *ContractNftFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractNftFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractNftFactoryCallerSession struct {
	Contract *ContractNftFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContractNftFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractNftFactoryTransactorSession struct {
	Contract     *ContractNftFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractNftFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractNftFactoryRaw struct {
	Contract *ContractNftFactory // Generic contract binding to access the raw methods on
}

// ContractNftFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractNftFactoryCallerRaw struct {
	Contract *ContractNftFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractNftFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractNftFactoryTransactorRaw struct {
	Contract *ContractNftFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractNftFactory creates a new instance of ContractNftFactory, bound to a specific deployed contract.
func NewContractNftFactory(address common.Address, backend bind.ContractBackend) (*ContractNftFactory, error) {
	contract, err := bindContractNftFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractNftFactory{ContractNftFactoryCaller: ContractNftFactoryCaller{contract: contract}, ContractNftFactoryTransactor: ContractNftFactoryTransactor{contract: contract}, ContractNftFactoryFilterer: ContractNftFactoryFilterer{contract: contract}}, nil
}

// NewContractNftFactoryCaller creates a new read-only instance of ContractNftFactory, bound to a specific deployed contract.
func NewContractNftFactoryCaller(address common.Address, caller bind.ContractCaller) (*ContractNftFactoryCaller, error) {
	contract, err := bindContractNftFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractNftFactoryCaller{contract: contract}, nil
}

// NewContractNftFactoryTransactor creates a new write-only instance of ContractNftFactory, bound to a specific deployed contract.
func NewContractNftFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractNftFactoryTransactor, error) {
	contract, err := bindContractNftFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractNftFactoryTransactor{contract: contract}, nil
}

// NewContractNftFactoryFilterer creates a new log filterer instance of ContractNftFactory, bound to a specific deployed contract.
func NewContractNftFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractNftFactoryFilterer, error) {
	contract, err := bindContractNftFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractNftFactoryFilterer{contract: contract}, nil
}

// bindContractNftFactory binds a generic wrapper to an already deployed contract.
func bindContractNftFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractNftFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractNftFactory *ContractNftFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractNftFactory.Contract.ContractNftFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractNftFactory *ContractNftFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractNftFactory.Contract.ContractNftFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractNftFactory *ContractNftFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractNftFactory.Contract.ContractNftFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractNftFactory *ContractNftFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractNftFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractNftFactory *ContractNftFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractNftFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractNftFactory *ContractNftFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractNftFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateAlbum is a paid mutator transaction binding the contract method 0xe940ff43.
//
// Solidity: function createAlbum(string name_, string symbol_) returns(address)
func (_ContractNftFactory *ContractNftFactoryTransactor) CreateAlbum(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _ContractNftFactory.contract.Transact(opts, "createAlbum", name_, symbol_)
}

// CreateAlbum is a paid mutator transaction binding the contract method 0xe940ff43.
//
// Solidity: function createAlbum(string name_, string symbol_) returns(address)
func (_ContractNftFactory *ContractNftFactorySession) CreateAlbum(name_ string, symbol_ string) (*types.Transaction, error) {
	return _ContractNftFactory.Contract.CreateAlbum(&_ContractNftFactory.TransactOpts, name_, symbol_)
}

// CreateAlbum is a paid mutator transaction binding the contract method 0xe940ff43.
//
// Solidity: function createAlbum(string name_, string symbol_) returns(address)
func (_ContractNftFactory *ContractNftFactoryTransactorSession) CreateAlbum(name_ string, symbol_ string) (*types.Transaction, error) {
	return _ContractNftFactory.Contract.CreateAlbum(&_ContractNftFactory.TransactOpts, name_, symbol_)
}

// ContractNftFactoryAlbumCreatedIterator is returned from FilterAlbumCreated and is used to iterate over the raw logs and unpacked data for AlbumCreated events raised by the ContractNftFactory contract.
type ContractNftFactoryAlbumCreatedIterator struct {
	Event *ContractNftFactoryAlbumCreated // Event containing the contract specifics and raw log

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
func (it *ContractNftFactoryAlbumCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNftFactoryAlbumCreated)
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
		it.Event = new(ContractNftFactoryAlbumCreated)
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
func (it *ContractNftFactoryAlbumCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNftFactoryAlbumCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNftFactoryAlbumCreated represents a AlbumCreated event raised by the ContractNftFactory contract.
type ContractNftFactoryAlbumCreated struct {
	Creator common.Address
	Album   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAlbumCreated is a free log retrieval operation binding the contract event 0x1f3134bb590263d2dc9f908ea8c7d58802e656b4bc76dec136da5f9faacf1f7b.
//
// Solidity: event AlbumCreated(address _creator, address _album)
func (_ContractNftFactory *ContractNftFactoryFilterer) FilterAlbumCreated(opts *bind.FilterOpts) (*ContractNftFactoryAlbumCreatedIterator, error) {

	logs, sub, err := _ContractNftFactory.contract.FilterLogs(opts, "AlbumCreated")
	if err != nil {
		return nil, err
	}
	return &ContractNftFactoryAlbumCreatedIterator{contract: _ContractNftFactory.contract, event: "AlbumCreated", logs: logs, sub: sub}, nil
}

// WatchAlbumCreated is a free log subscription operation binding the contract event 0x1f3134bb590263d2dc9f908ea8c7d58802e656b4bc76dec136da5f9faacf1f7b.
//
// Solidity: event AlbumCreated(address _creator, address _album)
func (_ContractNftFactory *ContractNftFactoryFilterer) WatchAlbumCreated(opts *bind.WatchOpts, sink chan<- *ContractNftFactoryAlbumCreated) (event.Subscription, error) {

	logs, sub, err := _ContractNftFactory.contract.WatchLogs(opts, "AlbumCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNftFactoryAlbumCreated)
				if err := _ContractNftFactory.contract.UnpackLog(event, "AlbumCreated", log); err != nil {
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

// ParseAlbumCreated is a log parse operation binding the contract event 0x1f3134bb590263d2dc9f908ea8c7d58802e656b4bc76dec136da5f9faacf1f7b.
//
// Solidity: event AlbumCreated(address _creator, address _album)
func (_ContractNftFactory *ContractNftFactoryFilterer) ParseAlbumCreated(log types.Log) (*ContractNftFactoryAlbumCreated, error) {
	event := new(ContractNftFactoryAlbumCreated)
	if err := _ContractNftFactory.contract.UnpackLog(event, "AlbumCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
