// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_airdrop

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

// ContractAirdropMetaData contains all meta data concerning the ContractAirdrop contract.
var ContractAirdropMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_id\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amount\",\"type\":\"uint256[]\"}],\"name\":\"bulkAirdropERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_value\",\"type\":\"uint256[]\"}],\"name\":\"bulkAirdropERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_id\",\"type\":\"uint256[]\"}],\"name\":\"bulkAirdropERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractAirdropABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAirdropMetaData.ABI instead.
var ContractAirdropABI = ContractAirdropMetaData.ABI

// ContractAirdrop is an auto generated Go binding around an Ethereum contract.
type ContractAirdrop struct {
	ContractAirdropCaller     // Read-only binding to the contract
	ContractAirdropTransactor // Write-only binding to the contract
	ContractAirdropFilterer   // Log filterer for contract events
}

// ContractAirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAirdropSession struct {
	Contract     *ContractAirdrop  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractAirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAirdropCallerSession struct {
	Contract *ContractAirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContractAirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAirdropTransactorSession struct {
	Contract     *ContractAirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractAirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAirdropRaw struct {
	Contract *ContractAirdrop // Generic contract binding to access the raw methods on
}

// ContractAirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAirdropCallerRaw struct {
	Contract *ContractAirdropCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAirdropTransactorRaw struct {
	Contract *ContractAirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAirdrop creates a new instance of ContractAirdrop, bound to a specific deployed contract.
func NewContractAirdrop(address common.Address, backend bind.ContractBackend) (*ContractAirdrop, error) {
	contract, err := bindContractAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAirdrop{ContractAirdropCaller: ContractAirdropCaller{contract: contract}, ContractAirdropTransactor: ContractAirdropTransactor{contract: contract}, ContractAirdropFilterer: ContractAirdropFilterer{contract: contract}}, nil
}

// NewContractAirdropCaller creates a new read-only instance of ContractAirdrop, bound to a specific deployed contract.
func NewContractAirdropCaller(address common.Address, caller bind.ContractCaller) (*ContractAirdropCaller, error) {
	contract, err := bindContractAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAirdropCaller{contract: contract}, nil
}

// NewContractAirdropTransactor creates a new write-only instance of ContractAirdrop, bound to a specific deployed contract.
func NewContractAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAirdropTransactor, error) {
	contract, err := bindContractAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAirdropTransactor{contract: contract}, nil
}

// NewContractAirdropFilterer creates a new log filterer instance of ContractAirdrop, bound to a specific deployed contract.
func NewContractAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAirdropFilterer, error) {
	contract, err := bindContractAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAirdropFilterer{contract: contract}, nil
}

// bindContractAirdrop binds a generic wrapper to an already deployed contract.
func bindContractAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAirdropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAirdrop *ContractAirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAirdrop.Contract.ContractAirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAirdrop *ContractAirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.ContractAirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAirdrop *ContractAirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.ContractAirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAirdrop *ContractAirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAirdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAirdrop *ContractAirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAirdrop *ContractAirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.contract.Transact(opts, method, params...)
}

// BulkAirdropERC1155 is a paid mutator transaction binding the contract method 0x7561a9dc.
//
// Solidity: function bulkAirdropERC1155(address _token, address[] _to, uint256[] _id, uint256[] _amount) returns()
func (_ContractAirdrop *ContractAirdropTransactor) BulkAirdropERC1155(opts *bind.TransactOpts, _token common.Address, _to []common.Address, _id []*big.Int, _amount []*big.Int) (*types.Transaction, error) {
	return _ContractAirdrop.contract.Transact(opts, "bulkAirdropERC1155", _token, _to, _id, _amount)
}

// BulkAirdropERC1155 is a paid mutator transaction binding the contract method 0x7561a9dc.
//
// Solidity: function bulkAirdropERC1155(address _token, address[] _to, uint256[] _id, uint256[] _amount) returns()
func (_ContractAirdrop *ContractAirdropSession) BulkAirdropERC1155(_token common.Address, _to []common.Address, _id []*big.Int, _amount []*big.Int) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.BulkAirdropERC1155(&_ContractAirdrop.TransactOpts, _token, _to, _id, _amount)
}

// BulkAirdropERC1155 is a paid mutator transaction binding the contract method 0x7561a9dc.
//
// Solidity: function bulkAirdropERC1155(address _token, address[] _to, uint256[] _id, uint256[] _amount) returns()
func (_ContractAirdrop *ContractAirdropTransactorSession) BulkAirdropERC1155(_token common.Address, _to []common.Address, _id []*big.Int, _amount []*big.Int) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.BulkAirdropERC1155(&_ContractAirdrop.TransactOpts, _token, _to, _id, _amount)
}

// BulkAirdropERC20 is a paid mutator transaction binding the contract method 0x242ab770.
//
// Solidity: function bulkAirdropERC20(address _token, address[] _to, uint256[] _value) returns()
func (_ContractAirdrop *ContractAirdropTransactor) BulkAirdropERC20(opts *bind.TransactOpts, _token common.Address, _to []common.Address, _value []*big.Int) (*types.Transaction, error) {
	return _ContractAirdrop.contract.Transact(opts, "bulkAirdropERC20", _token, _to, _value)
}

// BulkAirdropERC20 is a paid mutator transaction binding the contract method 0x242ab770.
//
// Solidity: function bulkAirdropERC20(address _token, address[] _to, uint256[] _value) returns()
func (_ContractAirdrop *ContractAirdropSession) BulkAirdropERC20(_token common.Address, _to []common.Address, _value []*big.Int) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.BulkAirdropERC20(&_ContractAirdrop.TransactOpts, _token, _to, _value)
}

// BulkAirdropERC20 is a paid mutator transaction binding the contract method 0x242ab770.
//
// Solidity: function bulkAirdropERC20(address _token, address[] _to, uint256[] _value) returns()
func (_ContractAirdrop *ContractAirdropTransactorSession) BulkAirdropERC20(_token common.Address, _to []common.Address, _value []*big.Int) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.BulkAirdropERC20(&_ContractAirdrop.TransactOpts, _token, _to, _value)
}

// BulkAirdropERC721 is a paid mutator transaction binding the contract method 0xcdb970d3.
//
// Solidity: function bulkAirdropERC721(address _token, address[] _to, uint256[] _id) returns()
func (_ContractAirdrop *ContractAirdropTransactor) BulkAirdropERC721(opts *bind.TransactOpts, _token common.Address, _to []common.Address, _id []*big.Int) (*types.Transaction, error) {
	return _ContractAirdrop.contract.Transact(opts, "bulkAirdropERC721", _token, _to, _id)
}

// BulkAirdropERC721 is a paid mutator transaction binding the contract method 0xcdb970d3.
//
// Solidity: function bulkAirdropERC721(address _token, address[] _to, uint256[] _id) returns()
func (_ContractAirdrop *ContractAirdropSession) BulkAirdropERC721(_token common.Address, _to []common.Address, _id []*big.Int) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.BulkAirdropERC721(&_ContractAirdrop.TransactOpts, _token, _to, _id)
}

// BulkAirdropERC721 is a paid mutator transaction binding the contract method 0xcdb970d3.
//
// Solidity: function bulkAirdropERC721(address _token, address[] _to, uint256[] _id) returns()
func (_ContractAirdrop *ContractAirdropTransactorSession) BulkAirdropERC721(_token common.Address, _to []common.Address, _id []*big.Int) (*types.Transaction, error) {
	return _ContractAirdrop.Contract.BulkAirdropERC721(&_ContractAirdrop.TransactOpts, _token, _to, _id)
}
