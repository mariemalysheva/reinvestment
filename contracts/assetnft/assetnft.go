// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_assetnft

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// AssetNFTAssetMetadata is an auto generated low-level Go binding around an user-defined struct.
type AssetNFTAssetMetadata struct {
	UserID         *big.Int
	Salt           string
	Amount         *big.Int
	SavingsBalance *big.Int
}

// ContractAssetnftMetaData contains all meta data concerning the ContractAssetnft contract.
var ContractAssetnftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reinvestmentPeriod_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids_\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"userBatchAssetsMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids_\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"userBatchTransfered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_batch_size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assetMetadata\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"userID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"salt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"savingsBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetVault\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"userIds_\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"userID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"salt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"savingsBalance\",\"type\":\"uint256\"}],\"internalType\":\"structAssetNFT.AssetMetadata[]\",\"name\":\"metadata_\",\"type\":\"tuple[]\"}],\"name\":\"mintNFTBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinvestmentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinvestmentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size_\",\"type\":\"uint256\"}],\"name\":\"setBatchSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"userIDs_\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"transferUserAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractAssetnftABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAssetnftMetaData.ABI instead.
var ContractAssetnftABI = ContractAssetnftMetaData.ABI

// ContractAssetnft is an auto generated Go binding around an Ethereum contract.
type ContractAssetnft struct {
	ContractAssetnftCaller     // Read-only binding to the contract
	ContractAssetnftTransactor // Write-only binding to the contract
	ContractAssetnftFilterer   // Log filterer for contract events
}

// ContractAssetnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAssetnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAssetnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAssetnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAssetnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAssetnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAssetnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAssetnftSession struct {
	Contract     *ContractAssetnft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractAssetnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAssetnftCallerSession struct {
	Contract *ContractAssetnftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractAssetnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAssetnftTransactorSession struct {
	Contract     *ContractAssetnftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractAssetnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAssetnftRaw struct {
	Contract *ContractAssetnft // Generic contract binding to access the raw methods on
}

// ContractAssetnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAssetnftCallerRaw struct {
	Contract *ContractAssetnftCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAssetnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAssetnftTransactorRaw struct {
	Contract *ContractAssetnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAssetnft creates a new instance of ContractAssetnft, bound to a specific deployed contract.
func NewContractAssetnft(address common.Address, backend bind.ContractBackend) (*ContractAssetnft, error) {
	contract, err := bindContractAssetnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAssetnft{ContractAssetnftCaller: ContractAssetnftCaller{contract: contract}, ContractAssetnftTransactor: ContractAssetnftTransactor{contract: contract}, ContractAssetnftFilterer: ContractAssetnftFilterer{contract: contract}}, nil
}

// NewContractAssetnftCaller creates a new read-only instance of ContractAssetnft, bound to a specific deployed contract.
func NewContractAssetnftCaller(address common.Address, caller bind.ContractCaller) (*ContractAssetnftCaller, error) {
	contract, err := bindContractAssetnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAssetnftCaller{contract: contract}, nil
}

// NewContractAssetnftTransactor creates a new write-only instance of ContractAssetnft, bound to a specific deployed contract.
func NewContractAssetnftTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAssetnftTransactor, error) {
	contract, err := bindContractAssetnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAssetnftTransactor{contract: contract}, nil
}

// NewContractAssetnftFilterer creates a new log filterer instance of ContractAssetnft, bound to a specific deployed contract.
func NewContractAssetnftFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAssetnftFilterer, error) {
	contract, err := bindContractAssetnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAssetnftFilterer{contract: contract}, nil
}

// bindContractAssetnft binds a generic wrapper to an already deployed contract.
func bindContractAssetnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAssetnftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAssetnft *ContractAssetnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAssetnft.Contract.ContractAssetnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAssetnft *ContractAssetnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.ContractAssetnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAssetnft *ContractAssetnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.ContractAssetnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAssetnft *ContractAssetnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAssetnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAssetnft *ContractAssetnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAssetnft *ContractAssetnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.contract.Transact(opts, method, params...)
}

// BatchSize is a free data retrieval call binding the contract method 0x5eedd2a0.
//
// Solidity: function _batch_size() view returns(uint256)
func (_ContractAssetnft *ContractAssetnftCaller) BatchSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "_batch_size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchSize is a free data retrieval call binding the contract method 0x5eedd2a0.
//
// Solidity: function _batch_size() view returns(uint256)
func (_ContractAssetnft *ContractAssetnftSession) BatchSize() (*big.Int, error) {
	return _ContractAssetnft.Contract.BatchSize(&_ContractAssetnft.CallOpts)
}

// BatchSize is a free data retrieval call binding the contract method 0x5eedd2a0.
//
// Solidity: function _batch_size() view returns(uint256)
func (_ContractAssetnft *ContractAssetnftCallerSession) BatchSize() (*big.Int, error) {
	return _ContractAssetnft.Contract.BatchSize(&_ContractAssetnft.CallOpts)
}

// AssetMetadata is a free data retrieval call binding the contract method 0x504e35d3.
//
// Solidity: function assetMetadata(uint256 ) view returns(uint256 userID, string salt, uint256 amount, uint256 savingsBalance)
func (_ContractAssetnft *ContractAssetnftCaller) AssetMetadata(opts *bind.CallOpts, arg0 *big.Int) (struct {
	UserID         *big.Int
	Salt           string
	Amount         *big.Int
	SavingsBalance *big.Int
}, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "assetMetadata", arg0)

	outstruct := new(struct {
		UserID         *big.Int
		Salt           string
		Amount         *big.Int
		SavingsBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Salt = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SavingsBalance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AssetMetadata is a free data retrieval call binding the contract method 0x504e35d3.
//
// Solidity: function assetMetadata(uint256 ) view returns(uint256 userID, string salt, uint256 amount, uint256 savingsBalance)
func (_ContractAssetnft *ContractAssetnftSession) AssetMetadata(arg0 *big.Int) (struct {
	UserID         *big.Int
	Salt           string
	Amount         *big.Int
	SavingsBalance *big.Int
}, error) {
	return _ContractAssetnft.Contract.AssetMetadata(&_ContractAssetnft.CallOpts, arg0)
}

// AssetMetadata is a free data retrieval call binding the contract method 0x504e35d3.
//
// Solidity: function assetMetadata(uint256 ) view returns(uint256 userID, string salt, uint256 amount, uint256 savingsBalance)
func (_ContractAssetnft *ContractAssetnftCallerSession) AssetMetadata(arg0 *big.Int) (struct {
	UserID         *big.Int
	Salt           string
	Amount         *big.Int
	SavingsBalance *big.Int
}, error) {
	return _ContractAssetnft.Contract.AssetMetadata(&_ContractAssetnft.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ContractAssetnft *ContractAssetnftCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ContractAssetnft *ContractAssetnftSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ContractAssetnft.Contract.BalanceOf(&_ContractAssetnft.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ContractAssetnft *ContractAssetnftCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ContractAssetnft.Contract.BalanceOf(&_ContractAssetnft.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ContractAssetnft *ContractAssetnftCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ContractAssetnft *ContractAssetnftSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ContractAssetnft.Contract.BalanceOfBatch(&_ContractAssetnft.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ContractAssetnft *ContractAssetnftCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ContractAssetnft.Contract.BalanceOfBatch(&_ContractAssetnft.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ContractAssetnft *ContractAssetnftCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ContractAssetnft *ContractAssetnftSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ContractAssetnft.Contract.IsApprovedForAll(&_ContractAssetnft.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ContractAssetnft *ContractAssetnftCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ContractAssetnft.Contract.IsApprovedForAll(&_ContractAssetnft.CallOpts, account, operator)
}

// ReinvestmentManager is a free data retrieval call binding the contract method 0x345d9f5a.
//
// Solidity: function reinvestmentManager() view returns(address)
func (_ContractAssetnft *ContractAssetnftCaller) ReinvestmentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "reinvestmentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReinvestmentManager is a free data retrieval call binding the contract method 0x345d9f5a.
//
// Solidity: function reinvestmentManager() view returns(address)
func (_ContractAssetnft *ContractAssetnftSession) ReinvestmentManager() (common.Address, error) {
	return _ContractAssetnft.Contract.ReinvestmentManager(&_ContractAssetnft.CallOpts)
}

// ReinvestmentManager is a free data retrieval call binding the contract method 0x345d9f5a.
//
// Solidity: function reinvestmentManager() view returns(address)
func (_ContractAssetnft *ContractAssetnftCallerSession) ReinvestmentManager() (common.Address, error) {
	return _ContractAssetnft.Contract.ReinvestmentManager(&_ContractAssetnft.CallOpts)
}

// ReinvestmentPeriod is a free data retrieval call binding the contract method 0x74669cdc.
//
// Solidity: function reinvestmentPeriod() view returns(uint256)
func (_ContractAssetnft *ContractAssetnftCaller) ReinvestmentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "reinvestmentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReinvestmentPeriod is a free data retrieval call binding the contract method 0x74669cdc.
//
// Solidity: function reinvestmentPeriod() view returns(uint256)
func (_ContractAssetnft *ContractAssetnftSession) ReinvestmentPeriod() (*big.Int, error) {
	return _ContractAssetnft.Contract.ReinvestmentPeriod(&_ContractAssetnft.CallOpts)
}

// ReinvestmentPeriod is a free data retrieval call binding the contract method 0x74669cdc.
//
// Solidity: function reinvestmentPeriod() view returns(uint256)
func (_ContractAssetnft *ContractAssetnftCallerSession) ReinvestmentPeriod() (*big.Int, error) {
	return _ContractAssetnft.Contract.ReinvestmentPeriod(&_ContractAssetnft.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAssetnft *ContractAssetnftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAssetnft *ContractAssetnftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractAssetnft.Contract.SupportsInterface(&_ContractAssetnft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractAssetnft *ContractAssetnftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractAssetnft.Contract.SupportsInterface(&_ContractAssetnft.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ContractAssetnft *ContractAssetnftCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ContractAssetnft *ContractAssetnftSession) Uri(arg0 *big.Int) (string, error) {
	return _ContractAssetnft.Contract.Uri(&_ContractAssetnft.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ContractAssetnft *ContractAssetnftCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _ContractAssetnft.Contract.Uri(&_ContractAssetnft.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ContractAssetnft *ContractAssetnftCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAssetnft.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ContractAssetnft *ContractAssetnftSession) Vault() (common.Address, error) {
	return _ContractAssetnft.Contract.Vault(&_ContractAssetnft.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ContractAssetnft *ContractAssetnftCallerSession) Vault() (common.Address, error) {
	return _ContractAssetnft.Contract.Vault(&_ContractAssetnft.CallOpts)
}

// MintNFTBatch is a paid mutator transaction binding the contract method 0x2f67353a.
//
// Solidity: function mintNFTBatch(address assetVault, uint256[] userIds_, (uint256,string,uint256,uint256)[] metadata_) returns()
func (_ContractAssetnft *ContractAssetnftTransactor) MintNFTBatch(opts *bind.TransactOpts, assetVault common.Address, userIds_ []*big.Int, metadata_ []AssetNFTAssetMetadata) (*types.Transaction, error) {
	aAbi, err := ContractAssetnftMetaData.GetAbi()
	if err != nil {
	return nil, err
	}

	input, err := aAbi.Pack("mintNFTBatch", assetVault, userIds_, metadata_)
	if err != nil {
	return nil, err
	}

	fmt.Println(hexutil.Encode(input))

	return _ContractAssetnft.contract.Transact(opts, "mintNFTBatch", assetVault, userIds_, metadata_)
}

// MintNFTBatch is a paid mutator transaction binding the contract method 0x2f67353a.
//
// Solidity: function mintNFTBatch(address assetVault, uint256[] userIds_, (uint256,string,uint256,uint256)[] metadata_) returns()
func (_ContractAssetnft *ContractAssetnftSession) MintNFTBatch(assetVault common.Address, userIds_ []*big.Int, metadata_ []AssetNFTAssetMetadata) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.MintNFTBatch(&_ContractAssetnft.TransactOpts, assetVault, userIds_, metadata_)
}

// MintNFTBatch is a paid mutator transaction binding the contract method 0x2f67353a.
//
// Solidity: function mintNFTBatch(address assetVault, uint256[] userIds_, (uint256,string,uint256,uint256)[] metadata_) returns()
func (_ContractAssetnft *ContractAssetnftTransactorSession) MintNFTBatch(assetVault common.Address, userIds_ []*big.Int, metadata_ []AssetNFTAssetMetadata) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.MintNFTBatch(&_ContractAssetnft.TransactOpts, assetVault, userIds_, metadata_)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ContractAssetnft *ContractAssetnftTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ContractAssetnft.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ContractAssetnft *ContractAssetnftSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.SafeBatchTransferFrom(&_ContractAssetnft.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ContractAssetnft *ContractAssetnftTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.SafeBatchTransferFrom(&_ContractAssetnft.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ContractAssetnft *ContractAssetnftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ContractAssetnft.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ContractAssetnft *ContractAssetnftSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.SafeTransferFrom(&_ContractAssetnft.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ContractAssetnft *ContractAssetnftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.SafeTransferFrom(&_ContractAssetnft.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ContractAssetnft *ContractAssetnftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ContractAssetnft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ContractAssetnft *ContractAssetnftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.SetApprovalForAll(&_ContractAssetnft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ContractAssetnft *ContractAssetnftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.SetApprovalForAll(&_ContractAssetnft.TransactOpts, operator, approved)
}

// SetBatchSize is a paid mutator transaction binding the contract method 0x576f35e3.
//
// Solidity: function setBatchSize(uint256 size_) returns()
func (_ContractAssetnft *ContractAssetnftTransactor) SetBatchSize(opts *bind.TransactOpts, size_ *big.Int) (*types.Transaction, error) {
	return _ContractAssetnft.contract.Transact(opts, "setBatchSize", size_)
}

// SetBatchSize is a paid mutator transaction binding the contract method 0x576f35e3.
//
// Solidity: function setBatchSize(uint256 size_) returns()
func (_ContractAssetnft *ContractAssetnftSession) SetBatchSize(size_ *big.Int) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.SetBatchSize(&_ContractAssetnft.TransactOpts, size_)
}

// SetBatchSize is a paid mutator transaction binding the contract method 0x576f35e3.
//
// Solidity: function setBatchSize(uint256 size_) returns()
func (_ContractAssetnft *ContractAssetnftTransactorSession) SetBatchSize(size_ *big.Int) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.SetBatchSize(&_ContractAssetnft.TransactOpts, size_)
}

// TransferUserAssets is a paid mutator transaction binding the contract method 0x75431e5b.
//
// Solidity: function transferUserAssets(uint256[] userIDs_, address to_) returns()
func (_ContractAssetnft *ContractAssetnftTransactor) TransferUserAssets(opts *bind.TransactOpts, userIDs_ []*big.Int, to_ common.Address) (*types.Transaction, error) {
	return _ContractAssetnft.contract.Transact(opts, "transferUserAssets", userIDs_, to_)
}

// TransferUserAssets is a paid mutator transaction binding the contract method 0x75431e5b.
//
// Solidity: function transferUserAssets(uint256[] userIDs_, address to_) returns()
func (_ContractAssetnft *ContractAssetnftSession) TransferUserAssets(userIDs_ []*big.Int, to_ common.Address) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.TransferUserAssets(&_ContractAssetnft.TransactOpts, userIDs_, to_)
}

// TransferUserAssets is a paid mutator transaction binding the contract method 0x75431e5b.
//
// Solidity: function transferUserAssets(uint256[] userIDs_, address to_) returns()
func (_ContractAssetnft *ContractAssetnftTransactorSession) TransferUserAssets(userIDs_ []*big.Int, to_ common.Address) (*types.Transaction, error) {
	return _ContractAssetnft.Contract.TransferUserAssets(&_ContractAssetnft.TransactOpts, userIDs_, to_)
}

// ContractAssetnftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ContractAssetnft contract.
type ContractAssetnftApprovalForAllIterator struct {
	Event *ContractAssetnftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ContractAssetnftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAssetnftApprovalForAll)
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
		it.Event = new(ContractAssetnftApprovalForAll)
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
func (it *ContractAssetnftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAssetnftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAssetnftApprovalForAll represents a ApprovalForAll event raised by the ContractAssetnft contract.
type ContractAssetnftApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ContractAssetnft *ContractAssetnftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ContractAssetnftApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAssetnft.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAssetnftApprovalForAllIterator{contract: _ContractAssetnft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ContractAssetnft *ContractAssetnftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractAssetnftApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractAssetnft.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAssetnftApprovalForAll)
				if err := _ContractAssetnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ContractAssetnft *ContractAssetnftFilterer) ParseApprovalForAll(log types.Log) (*ContractAssetnftApprovalForAll, error) {
	event := new(ContractAssetnftApprovalForAll)
	if err := _ContractAssetnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAssetnftTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ContractAssetnft contract.
type ContractAssetnftTransferBatchIterator struct {
	Event *ContractAssetnftTransferBatch // Event containing the contract specifics and raw log

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
func (it *ContractAssetnftTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAssetnftTransferBatch)
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
		it.Event = new(ContractAssetnftTransferBatch)
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
func (it *ContractAssetnftTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAssetnftTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAssetnftTransferBatch represents a TransferBatch event raised by the ContractAssetnft contract.
type ContractAssetnftTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ContractAssetnft *ContractAssetnftFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ContractAssetnftTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractAssetnft.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractAssetnftTransferBatchIterator{contract: _ContractAssetnft.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ContractAssetnft *ContractAssetnftFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ContractAssetnftTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractAssetnft.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAssetnftTransferBatch)
				if err := _ContractAssetnft.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ContractAssetnft *ContractAssetnftFilterer) ParseTransferBatch(log types.Log) (*ContractAssetnftTransferBatch, error) {
	event := new(ContractAssetnftTransferBatch)
	if err := _ContractAssetnft.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAssetnftTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ContractAssetnft contract.
type ContractAssetnftTransferSingleIterator struct {
	Event *ContractAssetnftTransferSingle // Event containing the contract specifics and raw log

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
func (it *ContractAssetnftTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAssetnftTransferSingle)
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
		it.Event = new(ContractAssetnftTransferSingle)
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
func (it *ContractAssetnftTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAssetnftTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAssetnftTransferSingle represents a TransferSingle event raised by the ContractAssetnft contract.
type ContractAssetnftTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ContractAssetnft *ContractAssetnftFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ContractAssetnftTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractAssetnft.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractAssetnftTransferSingleIterator{contract: _ContractAssetnft.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ContractAssetnft *ContractAssetnftFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ContractAssetnftTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractAssetnft.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAssetnftTransferSingle)
				if err := _ContractAssetnft.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ContractAssetnft *ContractAssetnftFilterer) ParseTransferSingle(log types.Log) (*ContractAssetnftTransferSingle, error) {
	event := new(ContractAssetnftTransferSingle)
	if err := _ContractAssetnft.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAssetnftURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ContractAssetnft contract.
type ContractAssetnftURIIterator struct {
	Event *ContractAssetnftURI // Event containing the contract specifics and raw log

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
func (it *ContractAssetnftURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAssetnftURI)
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
		it.Event = new(ContractAssetnftURI)
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
func (it *ContractAssetnftURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAssetnftURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAssetnftURI represents a URI event raised by the ContractAssetnft contract.
type ContractAssetnftURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ContractAssetnft *ContractAssetnftFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ContractAssetnftURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractAssetnft.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractAssetnftURIIterator{contract: _ContractAssetnft.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ContractAssetnft *ContractAssetnftFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ContractAssetnftURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractAssetnft.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAssetnftURI)
				if err := _ContractAssetnft.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ContractAssetnft *ContractAssetnftFilterer) ParseURI(log types.Log) (*ContractAssetnftURI, error) {
	event := new(ContractAssetnftURI)
	if err := _ContractAssetnft.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAssetnftUserBatchAssetsMintedIterator is returned from FilterUserBatchAssetsMinted and is used to iterate over the raw logs and unpacked data for UserBatchAssetsMinted events raised by the ContractAssetnft contract.
type ContractAssetnftUserBatchAssetsMintedIterator struct {
	Event *ContractAssetnftUserBatchAssetsMinted // Event containing the contract specifics and raw log

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
func (it *ContractAssetnftUserBatchAssetsMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAssetnftUserBatchAssetsMinted)
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
		it.Event = new(ContractAssetnftUserBatchAssetsMinted)
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
func (it *ContractAssetnftUserBatchAssetsMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAssetnftUserBatchAssetsMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAssetnftUserBatchAssetsMinted represents a UserBatchAssetsMinted event raised by the ContractAssetnft contract.
type ContractAssetnftUserBatchAssetsMinted struct {
	Ids     []*big.Int
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserBatchAssetsMinted is a free log retrieval operation binding the contract event 0x2526a339db126eae8b23e7d397a78dbd0233906a4e8b6bd939f03fd242c61cd8.
//
// Solidity: event userBatchAssetsMinted(uint256[] ids_, uint256[] amounts_)
func (_ContractAssetnft *ContractAssetnftFilterer) FilterUserBatchAssetsMinted(opts *bind.FilterOpts) (*ContractAssetnftUserBatchAssetsMintedIterator, error) {

	logs, sub, err := _ContractAssetnft.contract.FilterLogs(opts, "userBatchAssetsMinted")
	if err != nil {
		return nil, err
	}
	return &ContractAssetnftUserBatchAssetsMintedIterator{contract: _ContractAssetnft.contract, event: "userBatchAssetsMinted", logs: logs, sub: sub}, nil
}

// WatchUserBatchAssetsMinted is a free log subscription operation binding the contract event 0x2526a339db126eae8b23e7d397a78dbd0233906a4e8b6bd939f03fd242c61cd8.
//
// Solidity: event userBatchAssetsMinted(uint256[] ids_, uint256[] amounts_)
func (_ContractAssetnft *ContractAssetnftFilterer) WatchUserBatchAssetsMinted(opts *bind.WatchOpts, sink chan<- *ContractAssetnftUserBatchAssetsMinted) (event.Subscription, error) {

	logs, sub, err := _ContractAssetnft.contract.WatchLogs(opts, "userBatchAssetsMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAssetnftUserBatchAssetsMinted)
				if err := _ContractAssetnft.contract.UnpackLog(event, "userBatchAssetsMinted", log); err != nil {
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

// ParseUserBatchAssetsMinted is a log parse operation binding the contract event 0x2526a339db126eae8b23e7d397a78dbd0233906a4e8b6bd939f03fd242c61cd8.
//
// Solidity: event userBatchAssetsMinted(uint256[] ids_, uint256[] amounts_)
func (_ContractAssetnft *ContractAssetnftFilterer) ParseUserBatchAssetsMinted(log types.Log) (*ContractAssetnftUserBatchAssetsMinted, error) {
	event := new(ContractAssetnftUserBatchAssetsMinted)
	if err := _ContractAssetnft.contract.UnpackLog(event, "userBatchAssetsMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAssetnftUserBatchTransferedIterator is returned from FilterUserBatchTransfered and is used to iterate over the raw logs and unpacked data for UserBatchTransfered events raised by the ContractAssetnft contract.
type ContractAssetnftUserBatchTransferedIterator struct {
	Event *ContractAssetnftUserBatchTransfered // Event containing the contract specifics and raw log

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
func (it *ContractAssetnftUserBatchTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAssetnftUserBatchTransfered)
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
		it.Event = new(ContractAssetnftUserBatchTransfered)
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
func (it *ContractAssetnftUserBatchTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAssetnftUserBatchTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAssetnftUserBatchTransfered represents a UserBatchTransfered event raised by the ContractAssetnft contract.
type ContractAssetnftUserBatchTransfered struct {
	To      common.Address
	Ids     []*big.Int
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserBatchTransfered is a free log retrieval operation binding the contract event 0x70bf7a161ef35c8cb6b9bfa3f5305b73da2122263278bd9dd323839978a4f4f9.
//
// Solidity: event userBatchTransfered(address to_, uint256[] ids_, uint256[] amounts_)
func (_ContractAssetnft *ContractAssetnftFilterer) FilterUserBatchTransfered(opts *bind.FilterOpts) (*ContractAssetnftUserBatchTransferedIterator, error) {

	logs, sub, err := _ContractAssetnft.contract.FilterLogs(opts, "userBatchTransfered")
	if err != nil {
		return nil, err
	}
	return &ContractAssetnftUserBatchTransferedIterator{contract: _ContractAssetnft.contract, event: "userBatchTransfered", logs: logs, sub: sub}, nil
}

// WatchUserBatchTransfered is a free log subscription operation binding the contract event 0x70bf7a161ef35c8cb6b9bfa3f5305b73da2122263278bd9dd323839978a4f4f9.
//
// Solidity: event userBatchTransfered(address to_, uint256[] ids_, uint256[] amounts_)
func (_ContractAssetnft *ContractAssetnftFilterer) WatchUserBatchTransfered(opts *bind.WatchOpts, sink chan<- *ContractAssetnftUserBatchTransfered) (event.Subscription, error) {

	logs, sub, err := _ContractAssetnft.contract.WatchLogs(opts, "userBatchTransfered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAssetnftUserBatchTransfered)
				if err := _ContractAssetnft.contract.UnpackLog(event, "userBatchTransfered", log); err != nil {
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

// ParseUserBatchTransfered is a log parse operation binding the contract event 0x70bf7a161ef35c8cb6b9bfa3f5305b73da2122263278bd9dd323839978a4f4f9.
//
// Solidity: event userBatchTransfered(address to_, uint256[] ids_, uint256[] amounts_)
func (_ContractAssetnft *ContractAssetnftFilterer) ParseUserBatchTransfered(log types.Log) (*ContractAssetnftUserBatchTransfered, error) {
	event := new(ContractAssetnftUserBatchTransfered)
	if err := _ContractAssetnft.contract.UnpackLog(event, "userBatchTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
