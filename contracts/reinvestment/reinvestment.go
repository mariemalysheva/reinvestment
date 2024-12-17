// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_reinvestment

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

// AssetNFTAssetMetadata is an auto generated low-level Go binding around an user-defined struct.
type AssetNFTAssetMetadata struct {
	UserID         *big.Int
	Salt           string
	Amount         *big.Int
	SavingsBalance *big.Int
}

// ReinvestmentManagerReinvestmentPeriod is an auto generated low-level Go binding around an user-defined struct.
type ReinvestmentManagerReinvestmentPeriod struct {
	ID           *big.Int
	Start        *big.Int
	End          *big.Int
	Rate         *big.Int
	AssetPrice   *big.Int
	CurrentAsset common.Address
}

// ContractReinvestmentMetaData contains all meta data concerning the ContractReinvestment contract.
var ContractReinvestmentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentAsset\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReinvestmentManager.ReinvestmentPeriod\",\"name\":\"reinvestmentPeriod\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"userIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"newBalances\",\"type\":\"uint256[]\"}],\"name\":\"savingsReinvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reinvestmentPeriodID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"userID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"salt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"savingsBalance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAssetNFT.AssetMetadata[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"}],\"name\":\"userBatchAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"userIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"userBatchTransfered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_user_ids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"userIDs_\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"salt_\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balance_\",\"type\":\"uint256[]\"}],\"name\":\"addUserBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userID_\",\"type\":\"uint256\"}],\"name\":\"isUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"passedReinvestmentPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentAsset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinvestSavings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinvestmentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentAsset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"setAssetPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate_\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"userIDs_\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"transferUserBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userPortfolio\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentAsset\",\"type\":\"address\"}],\"internalType\":\"structReinvestmentManager.ReinvestmentPeriod\",\"name\":\"reinvestmentPeriod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"userID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"salt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"savingsBalance\",\"type\":\"uint256\"}],\"internalType\":\"structAssetNFT.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractReinvestmentABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractReinvestmentMetaData.ABI instead.
var ContractReinvestmentABI = ContractReinvestmentMetaData.ABI

// ContractReinvestment is an auto generated Go binding around an Ethereum contract.
type ContractReinvestment struct {
	ContractReinvestmentCaller     // Read-only binding to the contract
	ContractReinvestmentTransactor // Write-only binding to the contract
	ContractReinvestmentFilterer   // Log filterer for contract events
}

// ContractReinvestmentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractReinvestmentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReinvestmentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractReinvestmentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReinvestmentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractReinvestmentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReinvestmentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractReinvestmentSession struct {
	Contract     *ContractReinvestment // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContractReinvestmentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractReinvestmentCallerSession struct {
	Contract *ContractReinvestmentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ContractReinvestmentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractReinvestmentTransactorSession struct {
	Contract     *ContractReinvestmentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractReinvestmentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractReinvestmentRaw struct {
	Contract *ContractReinvestment // Generic contract binding to access the raw methods on
}

// ContractReinvestmentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractReinvestmentCallerRaw struct {
	Contract *ContractReinvestmentCaller // Generic read-only contract binding to access the raw methods on
}

// ContractReinvestmentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractReinvestmentTransactorRaw struct {
	Contract *ContractReinvestmentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractReinvestment creates a new instance of ContractReinvestment, bound to a specific deployed contract.
func NewContractReinvestment(address common.Address, backend bind.ContractBackend) (*ContractReinvestment, error) {
	contract, err := bindContractReinvestment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractReinvestment{ContractReinvestmentCaller: ContractReinvestmentCaller{contract: contract}, ContractReinvestmentTransactor: ContractReinvestmentTransactor{contract: contract}, ContractReinvestmentFilterer: ContractReinvestmentFilterer{contract: contract}}, nil
}

// NewContractReinvestmentCaller creates a new read-only instance of ContractReinvestment, bound to a specific deployed contract.
func NewContractReinvestmentCaller(address common.Address, caller bind.ContractCaller) (*ContractReinvestmentCaller, error) {
	contract, err := bindContractReinvestment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractReinvestmentCaller{contract: contract}, nil
}

// NewContractReinvestmentTransactor creates a new write-only instance of ContractReinvestment, bound to a specific deployed contract.
func NewContractReinvestmentTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractReinvestmentTransactor, error) {
	contract, err := bindContractReinvestment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractReinvestmentTransactor{contract: contract}, nil
}

// NewContractReinvestmentFilterer creates a new log filterer instance of ContractReinvestment, bound to a specific deployed contract.
func NewContractReinvestmentFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractReinvestmentFilterer, error) {
	contract, err := bindContractReinvestment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractReinvestmentFilterer{contract: contract}, nil
}

// bindContractReinvestment binds a generic wrapper to an already deployed contract.
func bindContractReinvestment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractReinvestmentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractReinvestment *ContractReinvestmentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractReinvestment.Contract.ContractReinvestmentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractReinvestment *ContractReinvestmentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.ContractReinvestmentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractReinvestment *ContractReinvestmentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.ContractReinvestmentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractReinvestment *ContractReinvestmentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractReinvestment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractReinvestment *ContractReinvestmentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractReinvestment *ContractReinvestmentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.contract.Transact(opts, method, params...)
}

// UserIds is a free data retrieval call binding the contract method 0x7671fe77.
//
// Solidity: function _user_ids(uint256 ) view returns(uint256)
func (_ContractReinvestment *ContractReinvestmentCaller) UserIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "_user_ids", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserIds is a free data retrieval call binding the contract method 0x7671fe77.
//
// Solidity: function _user_ids(uint256 ) view returns(uint256)
func (_ContractReinvestment *ContractReinvestmentSession) UserIds(arg0 *big.Int) (*big.Int, error) {
	return _ContractReinvestment.Contract.UserIds(&_ContractReinvestment.CallOpts, arg0)
}

// UserIds is a free data retrieval call binding the contract method 0x7671fe77.
//
// Solidity: function _user_ids(uint256 ) view returns(uint256)
func (_ContractReinvestment *ContractReinvestmentCallerSession) UserIds(arg0 *big.Int) (*big.Int, error) {
	return _ContractReinvestment.Contract.UserIds(&_ContractReinvestment.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ContractReinvestment *ContractReinvestmentCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ContractReinvestment *ContractReinvestmentSession) Admin() (common.Address, error) {
	return _ContractReinvestment.Contract.Admin(&_ContractReinvestment.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ContractReinvestment *ContractReinvestmentCallerSession) Admin() (common.Address, error) {
	return _ContractReinvestment.Contract.Admin(&_ContractReinvestment.CallOpts)
}

// AssetVault is a free data retrieval call binding the contract method 0xf61ebc36.
//
// Solidity: function assetVault() view returns(address)
func (_ContractReinvestment *ContractReinvestmentCaller) AssetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "assetVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetVault is a free data retrieval call binding the contract method 0xf61ebc36.
//
// Solidity: function assetVault() view returns(address)
func (_ContractReinvestment *ContractReinvestmentSession) AssetVault() (common.Address, error) {
	return _ContractReinvestment.Contract.AssetVault(&_ContractReinvestment.CallOpts)
}

// AssetVault is a free data retrieval call binding the contract method 0xf61ebc36.
//
// Solidity: function assetVault() view returns(address)
func (_ContractReinvestment *ContractReinvestmentCallerSession) AssetVault() (common.Address, error) {
	return _ContractReinvestment.Contract.AssetVault(&_ContractReinvestment.CallOpts)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_ContractReinvestment *ContractReinvestmentCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "assets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_ContractReinvestment *ContractReinvestmentSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _ContractReinvestment.Contract.Assets(&_ContractReinvestment.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_ContractReinvestment *ContractReinvestmentCallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _ContractReinvestment.Contract.Assets(&_ContractReinvestment.CallOpts, arg0)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_ContractReinvestment *ContractReinvestmentCaller) GetUserLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "getUserLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_ContractReinvestment *ContractReinvestmentSession) GetUserLength() (*big.Int, error) {
	return _ContractReinvestment.Contract.GetUserLength(&_ContractReinvestment.CallOpts)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_ContractReinvestment *ContractReinvestmentCallerSession) GetUserLength() (*big.Int, error) {
	return _ContractReinvestment.Contract.GetUserLength(&_ContractReinvestment.CallOpts)
}

// IsUser is a free data retrieval call binding the contract method 0x18a9cc1b.
//
// Solidity: function isUser(uint256 userID_) view returns(bool)
func (_ContractReinvestment *ContractReinvestmentCaller) IsUser(opts *bind.CallOpts, userID_ *big.Int) (bool, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "isUser", userID_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUser is a free data retrieval call binding the contract method 0x18a9cc1b.
//
// Solidity: function isUser(uint256 userID_) view returns(bool)
func (_ContractReinvestment *ContractReinvestmentSession) IsUser(userID_ *big.Int) (bool, error) {
	return _ContractReinvestment.Contract.IsUser(&_ContractReinvestment.CallOpts, userID_)
}

// IsUser is a free data retrieval call binding the contract method 0x18a9cc1b.
//
// Solidity: function isUser(uint256 userID_) view returns(bool)
func (_ContractReinvestment *ContractReinvestmentCallerSession) IsUser(userID_ *big.Int) (bool, error) {
	return _ContractReinvestment.Contract.IsUser(&_ContractReinvestment.CallOpts, userID_)
}

// PassedReinvestmentPeriods is a free data retrieval call binding the contract method 0x20e5e0eb.
//
// Solidity: function passedReinvestmentPeriods(uint256 ) view returns(uint256 ID, uint256 start, uint256 end, uint256 rate, uint256 assetPrice, address currentAsset)
func (_ContractReinvestment *ContractReinvestmentCaller) PassedReinvestmentPeriods(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ID           *big.Int
	Start        *big.Int
	End          *big.Int
	Rate         *big.Int
	AssetPrice   *big.Int
	CurrentAsset common.Address
}, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "passedReinvestmentPeriods", arg0)

	outstruct := new(struct {
		ID           *big.Int
		Start        *big.Int
		End          *big.Int
		Rate         *big.Int
		AssetPrice   *big.Int
		CurrentAsset common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Start = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AssetPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CurrentAsset = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PassedReinvestmentPeriods is a free data retrieval call binding the contract method 0x20e5e0eb.
//
// Solidity: function passedReinvestmentPeriods(uint256 ) view returns(uint256 ID, uint256 start, uint256 end, uint256 rate, uint256 assetPrice, address currentAsset)
func (_ContractReinvestment *ContractReinvestmentSession) PassedReinvestmentPeriods(arg0 *big.Int) (struct {
	ID           *big.Int
	Start        *big.Int
	End          *big.Int
	Rate         *big.Int
	AssetPrice   *big.Int
	CurrentAsset common.Address
}, error) {
	return _ContractReinvestment.Contract.PassedReinvestmentPeriods(&_ContractReinvestment.CallOpts, arg0)
}

// PassedReinvestmentPeriods is a free data retrieval call binding the contract method 0x20e5e0eb.
//
// Solidity: function passedReinvestmentPeriods(uint256 ) view returns(uint256 ID, uint256 start, uint256 end, uint256 rate, uint256 assetPrice, address currentAsset)
func (_ContractReinvestment *ContractReinvestmentCallerSession) PassedReinvestmentPeriods(arg0 *big.Int) (struct {
	ID           *big.Int
	Start        *big.Int
	End          *big.Int
	Rate         *big.Int
	AssetPrice   *big.Int
	CurrentAsset common.Address
}, error) {
	return _ContractReinvestment.Contract.PassedReinvestmentPeriods(&_ContractReinvestment.CallOpts, arg0)
}

// ReinvestmentPeriod is a free data retrieval call binding the contract method 0x74669cdc.
//
// Solidity: function reinvestmentPeriod() view returns(uint256 ID, uint256 start, uint256 end, uint256 rate, uint256 assetPrice, address currentAsset)
func (_ContractReinvestment *ContractReinvestmentCaller) ReinvestmentPeriod(opts *bind.CallOpts) (struct {
	ID           *big.Int
	Start        *big.Int
	End          *big.Int
	Rate         *big.Int
	AssetPrice   *big.Int
	CurrentAsset common.Address
}, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "reinvestmentPeriod")

	outstruct := new(struct {
		ID           *big.Int
		Start        *big.Int
		End          *big.Int
		Rate         *big.Int
		AssetPrice   *big.Int
		CurrentAsset common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Start = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AssetPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CurrentAsset = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ReinvestmentPeriod is a free data retrieval call binding the contract method 0x74669cdc.
//
// Solidity: function reinvestmentPeriod() view returns(uint256 ID, uint256 start, uint256 end, uint256 rate, uint256 assetPrice, address currentAsset)
func (_ContractReinvestment *ContractReinvestmentSession) ReinvestmentPeriod() (struct {
	ID           *big.Int
	Start        *big.Int
	End          *big.Int
	Rate         *big.Int
	AssetPrice   *big.Int
	CurrentAsset common.Address
}, error) {
	return _ContractReinvestment.Contract.ReinvestmentPeriod(&_ContractReinvestment.CallOpts)
}

// ReinvestmentPeriod is a free data retrieval call binding the contract method 0x74669cdc.
//
// Solidity: function reinvestmentPeriod() view returns(uint256 ID, uint256 start, uint256 end, uint256 rate, uint256 assetPrice, address currentAsset)
func (_ContractReinvestment *ContractReinvestmentCallerSession) ReinvestmentPeriod() (struct {
	ID           *big.Int
	Start        *big.Int
	End          *big.Int
	Rate         *big.Int
	AssetPrice   *big.Int
	CurrentAsset common.Address
}, error) {
	return _ContractReinvestment.Contract.ReinvestmentPeriod(&_ContractReinvestment.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0xd17106e9.
//
// Solidity: function userBalance(uint256 ) view returns(uint256)
func (_ContractReinvestment *ContractReinvestmentCaller) UserBalance(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "userBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalance is a free data retrieval call binding the contract method 0xd17106e9.
//
// Solidity: function userBalance(uint256 ) view returns(uint256)
func (_ContractReinvestment *ContractReinvestmentSession) UserBalance(arg0 *big.Int) (*big.Int, error) {
	return _ContractReinvestment.Contract.UserBalance(&_ContractReinvestment.CallOpts, arg0)
}

// UserBalance is a free data retrieval call binding the contract method 0xd17106e9.
//
// Solidity: function userBalance(uint256 ) view returns(uint256)
func (_ContractReinvestment *ContractReinvestmentCallerSession) UserBalance(arg0 *big.Int) (*big.Int, error) {
	return _ContractReinvestment.Contract.UserBalance(&_ContractReinvestment.CallOpts, arg0)
}

// UserPortfolio is a free data retrieval call binding the contract method 0x212355cc.
//
// Solidity: function userPortfolio(uint256 , uint256 ) view returns((uint256,uint256,uint256,uint256,uint256,address) reinvestmentPeriod, (uint256,string,uint256,uint256) assetMetadata)
func (_ContractReinvestment *ContractReinvestmentCaller) UserPortfolio(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	ReinvestmentPeriod ReinvestmentManagerReinvestmentPeriod
	AssetMetadata      AssetNFTAssetMetadata
}, error) {
	var out []interface{}
	err := _ContractReinvestment.contract.Call(opts, &out, "userPortfolio", arg0, arg1)

	outstruct := new(struct {
		ReinvestmentPeriod ReinvestmentManagerReinvestmentPeriod
		AssetMetadata      AssetNFTAssetMetadata
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReinvestmentPeriod = *abi.ConvertType(out[0], new(ReinvestmentManagerReinvestmentPeriod)).(*ReinvestmentManagerReinvestmentPeriod)
	outstruct.AssetMetadata = *abi.ConvertType(out[1], new(AssetNFTAssetMetadata)).(*AssetNFTAssetMetadata)

	return *outstruct, err

}

// UserPortfolio is a free data retrieval call binding the contract method 0x212355cc.
//
// Solidity: function userPortfolio(uint256 , uint256 ) view returns((uint256,uint256,uint256,uint256,uint256,address) reinvestmentPeriod, (uint256,string,uint256,uint256) assetMetadata)
func (_ContractReinvestment *ContractReinvestmentSession) UserPortfolio(arg0 *big.Int, arg1 *big.Int) (struct {
	ReinvestmentPeriod ReinvestmentManagerReinvestmentPeriod
	AssetMetadata      AssetNFTAssetMetadata
}, error) {
	return _ContractReinvestment.Contract.UserPortfolio(&_ContractReinvestment.CallOpts, arg0, arg1)
}

// UserPortfolio is a free data retrieval call binding the contract method 0x212355cc.
//
// Solidity: function userPortfolio(uint256 , uint256 ) view returns((uint256,uint256,uint256,uint256,uint256,address) reinvestmentPeriod, (uint256,string,uint256,uint256) assetMetadata)
func (_ContractReinvestment *ContractReinvestmentCallerSession) UserPortfolio(arg0 *big.Int, arg1 *big.Int) (struct {
	ReinvestmentPeriod ReinvestmentManagerReinvestmentPeriod
	AssetMetadata      AssetNFTAssetMetadata
}, error) {
	return _ContractReinvestment.Contract.UserPortfolio(&_ContractReinvestment.CallOpts, arg0, arg1)
}

// AddUserBatch is a paid mutator transaction binding the contract method 0x0c3da375.
//
// Solidity: function addUserBatch(uint256[] userIDs_, string[] salt_, uint256[] balance_) returns()
func (_ContractReinvestment *ContractReinvestmentTransactor) AddUserBatch(opts *bind.TransactOpts, userIDs_ []*big.Int, salt_ []string, balance_ []*big.Int) (*types.Transaction, error) {
	return _ContractReinvestment.contract.Transact(opts, "addUserBatch", userIDs_, salt_, balance_)
}

// AddUserBatch is a paid mutator transaction binding the contract method 0x0c3da375.
//
// Solidity: function addUserBatch(uint256[] userIDs_, string[] salt_, uint256[] balance_) returns()
func (_ContractReinvestment *ContractReinvestmentSession) AddUserBatch(userIDs_ []*big.Int, salt_ []string, balance_ []*big.Int) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.AddUserBatch(&_ContractReinvestment.TransactOpts, userIDs_, salt_, balance_)
}

// AddUserBatch is a paid mutator transaction binding the contract method 0x0c3da375.
//
// Solidity: function addUserBatch(uint256[] userIDs_, string[] salt_, uint256[] balance_) returns()
func (_ContractReinvestment *ContractReinvestmentTransactorSession) AddUserBatch(userIDs_ []*big.Int, salt_ []string, balance_ []*big.Int) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.AddUserBatch(&_ContractReinvestment.TransactOpts, userIDs_, salt_, balance_)
}

// ReinvestSavings is a paid mutator transaction binding the contract method 0x197460e0.
//
// Solidity: function reinvestSavings() returns()
func (_ContractReinvestment *ContractReinvestmentTransactor) ReinvestSavings(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractReinvestment.contract.Transact(opts, "reinvestSavings")
}

// ReinvestSavings is a paid mutator transaction binding the contract method 0x197460e0.
//
// Solidity: function reinvestSavings() returns()
func (_ContractReinvestment *ContractReinvestmentSession) ReinvestSavings() (*types.Transaction, error) {
	return _ContractReinvestment.Contract.ReinvestSavings(&_ContractReinvestment.TransactOpts)
}

// ReinvestSavings is a paid mutator transaction binding the contract method 0x197460e0.
//
// Solidity: function reinvestSavings() returns()
func (_ContractReinvestment *ContractReinvestmentTransactorSession) ReinvestSavings() (*types.Transaction, error) {
	return _ContractReinvestment.Contract.ReinvestSavings(&_ContractReinvestment.TransactOpts)
}

// SetAssetPrice is a paid mutator transaction binding the contract method 0x2eed8b95.
//
// Solidity: function setAssetPrice(uint256 price_) returns()
func (_ContractReinvestment *ContractReinvestmentTransactor) SetAssetPrice(opts *bind.TransactOpts, price_ *big.Int) (*types.Transaction, error) {
	return _ContractReinvestment.contract.Transact(opts, "setAssetPrice", price_)
}

// SetAssetPrice is a paid mutator transaction binding the contract method 0x2eed8b95.
//
// Solidity: function setAssetPrice(uint256 price_) returns()
func (_ContractReinvestment *ContractReinvestmentSession) SetAssetPrice(price_ *big.Int) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.SetAssetPrice(&_ContractReinvestment.TransactOpts, price_)
}

// SetAssetPrice is a paid mutator transaction binding the contract method 0x2eed8b95.
//
// Solidity: function setAssetPrice(uint256 price_) returns()
func (_ContractReinvestment *ContractReinvestmentTransactorSession) SetAssetPrice(price_ *big.Int) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.SetAssetPrice(&_ContractReinvestment.TransactOpts, price_)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 rate_) returns()
func (_ContractReinvestment *ContractReinvestmentTransactor) SetRate(opts *bind.TransactOpts, rate_ *big.Int) (*types.Transaction, error) {
	return _ContractReinvestment.contract.Transact(opts, "setRate", rate_)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 rate_) returns()
func (_ContractReinvestment *ContractReinvestmentSession) SetRate(rate_ *big.Int) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.SetRate(&_ContractReinvestment.TransactOpts, rate_)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 rate_) returns()
func (_ContractReinvestment *ContractReinvestmentTransactorSession) SetRate(rate_ *big.Int) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.SetRate(&_ContractReinvestment.TransactOpts, rate_)
}

// TransferUserBatch is a paid mutator transaction binding the contract method 0xa9cc8c27.
//
// Solidity: function transferUserBatch(uint256[] userIDs_, address to_) returns()
func (_ContractReinvestment *ContractReinvestmentTransactor) TransferUserBatch(opts *bind.TransactOpts, userIDs_ []*big.Int, to_ common.Address) (*types.Transaction, error) {
	return _ContractReinvestment.contract.Transact(opts, "transferUserBatch", userIDs_, to_)
}

// TransferUserBatch is a paid mutator transaction binding the contract method 0xa9cc8c27.
//
// Solidity: function transferUserBatch(uint256[] userIDs_, address to_) returns()
func (_ContractReinvestment *ContractReinvestmentSession) TransferUserBatch(userIDs_ []*big.Int, to_ common.Address) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.TransferUserBatch(&_ContractReinvestment.TransactOpts, userIDs_, to_)
}

// TransferUserBatch is a paid mutator transaction binding the contract method 0xa9cc8c27.
//
// Solidity: function transferUserBatch(uint256[] userIDs_, address to_) returns()
func (_ContractReinvestment *ContractReinvestmentTransactorSession) TransferUserBatch(userIDs_ []*big.Int, to_ common.Address) (*types.Transaction, error) {
	return _ContractReinvestment.Contract.TransferUserBatch(&_ContractReinvestment.TransactOpts, userIDs_, to_)
}

// ContractReinvestmentSavingsReinvestedIterator is returned from FilterSavingsReinvested and is used to iterate over the raw logs and unpacked data for SavingsReinvested events raised by the ContractReinvestment contract.
type ContractReinvestmentSavingsReinvestedIterator struct {
	Event *ContractReinvestmentSavingsReinvested // Event containing the contract specifics and raw log

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
func (it *ContractReinvestmentSavingsReinvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReinvestmentSavingsReinvested)
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
		it.Event = new(ContractReinvestmentSavingsReinvested)
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
func (it *ContractReinvestmentSavingsReinvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractReinvestmentSavingsReinvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractReinvestmentSavingsReinvested represents a SavingsReinvested event raised by the ContractReinvestment contract.
type ContractReinvestmentSavingsReinvested struct {
	ReinvestmentPeriod ReinvestmentManagerReinvestmentPeriod
	UserIDs            []*big.Int
	NewBalances        []*big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSavingsReinvested is a free log retrieval operation binding the contract event 0x25611233c14737befe0be4b005a57d52efdeb9b588bbae6a1bc6d98d83ad1d9f.
//
// Solidity: event savingsReinvested((uint256,uint256,uint256,uint256,uint256,address) reinvestmentPeriod, uint256[] userIDs, uint256[] newBalances)
func (_ContractReinvestment *ContractReinvestmentFilterer) FilterSavingsReinvested(opts *bind.FilterOpts) (*ContractReinvestmentSavingsReinvestedIterator, error) {

	logs, sub, err := _ContractReinvestment.contract.FilterLogs(opts, "savingsReinvested")
	if err != nil {
		return nil, err
	}
	return &ContractReinvestmentSavingsReinvestedIterator{contract: _ContractReinvestment.contract, event: "savingsReinvested", logs: logs, sub: sub}, nil
}

// WatchSavingsReinvested is a free log subscription operation binding the contract event 0x25611233c14737befe0be4b005a57d52efdeb9b588bbae6a1bc6d98d83ad1d9f.
//
// Solidity: event savingsReinvested((uint256,uint256,uint256,uint256,uint256,address) reinvestmentPeriod, uint256[] userIDs, uint256[] newBalances)
func (_ContractReinvestment *ContractReinvestmentFilterer) WatchSavingsReinvested(opts *bind.WatchOpts, sink chan<- *ContractReinvestmentSavingsReinvested) (event.Subscription, error) {

	logs, sub, err := _ContractReinvestment.contract.WatchLogs(opts, "savingsReinvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractReinvestmentSavingsReinvested)
				if err := _ContractReinvestment.contract.UnpackLog(event, "savingsReinvested", log); err != nil {
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

// ParseSavingsReinvested is a log parse operation binding the contract event 0x25611233c14737befe0be4b005a57d52efdeb9b588bbae6a1bc6d98d83ad1d9f.
//
// Solidity: event savingsReinvested((uint256,uint256,uint256,uint256,uint256,address) reinvestmentPeriod, uint256[] userIDs, uint256[] newBalances)
func (_ContractReinvestment *ContractReinvestmentFilterer) ParseSavingsReinvested(log types.Log) (*ContractReinvestmentSavingsReinvested, error) {
	event := new(ContractReinvestmentSavingsReinvested)
	if err := _ContractReinvestment.contract.UnpackLog(event, "savingsReinvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractReinvestmentUserBatchAddedIterator is returned from FilterUserBatchAdded and is used to iterate over the raw logs and unpacked data for UserBatchAdded events raised by the ContractReinvestment contract.
type ContractReinvestmentUserBatchAddedIterator struct {
	Event *ContractReinvestmentUserBatchAdded // Event containing the contract specifics and raw log

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
func (it *ContractReinvestmentUserBatchAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReinvestmentUserBatchAdded)
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
		it.Event = new(ContractReinvestmentUserBatchAdded)
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
func (it *ContractReinvestmentUserBatchAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractReinvestmentUserBatchAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractReinvestmentUserBatchAdded represents a UserBatchAdded event raised by the ContractReinvestment contract.
type ContractReinvestmentUserBatchAdded struct {
	ReinvestmentPeriodID *big.Int
	Metadata             []AssetNFTAssetMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUserBatchAdded is a free log retrieval operation binding the contract event 0xd1d207a485dc3ce78f408c94eccbd1459359297b930a35e94bb0c19ef31022bc.
//
// Solidity: event userBatchAdded(uint256 reinvestmentPeriodID, (uint256,string,uint256,uint256)[] metadata)
func (_ContractReinvestment *ContractReinvestmentFilterer) FilterUserBatchAdded(opts *bind.FilterOpts) (*ContractReinvestmentUserBatchAddedIterator, error) {

	logs, sub, err := _ContractReinvestment.contract.FilterLogs(opts, "userBatchAdded")
	if err != nil {
		return nil, err
	}
	return &ContractReinvestmentUserBatchAddedIterator{contract: _ContractReinvestment.contract, event: "userBatchAdded", logs: logs, sub: sub}, nil
}

// WatchUserBatchAdded is a free log subscription operation binding the contract event 0xd1d207a485dc3ce78f408c94eccbd1459359297b930a35e94bb0c19ef31022bc.
//
// Solidity: event userBatchAdded(uint256 reinvestmentPeriodID, (uint256,string,uint256,uint256)[] metadata)
func (_ContractReinvestment *ContractReinvestmentFilterer) WatchUserBatchAdded(opts *bind.WatchOpts, sink chan<- *ContractReinvestmentUserBatchAdded) (event.Subscription, error) {

	logs, sub, err := _ContractReinvestment.contract.WatchLogs(opts, "userBatchAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractReinvestmentUserBatchAdded)
				if err := _ContractReinvestment.contract.UnpackLog(event, "userBatchAdded", log); err != nil {
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

// ParseUserBatchAdded is a log parse operation binding the contract event 0xd1d207a485dc3ce78f408c94eccbd1459359297b930a35e94bb0c19ef31022bc.
//
// Solidity: event userBatchAdded(uint256 reinvestmentPeriodID, (uint256,string,uint256,uint256)[] metadata)
func (_ContractReinvestment *ContractReinvestmentFilterer) ParseUserBatchAdded(log types.Log) (*ContractReinvestmentUserBatchAdded, error) {
	event := new(ContractReinvestmentUserBatchAdded)
	if err := _ContractReinvestment.contract.UnpackLog(event, "userBatchAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractReinvestmentUserBatchTransferedIterator is returned from FilterUserBatchTransfered and is used to iterate over the raw logs and unpacked data for UserBatchTransfered events raised by the ContractReinvestment contract.
type ContractReinvestmentUserBatchTransferedIterator struct {
	Event *ContractReinvestmentUserBatchTransfered // Event containing the contract specifics and raw log

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
func (it *ContractReinvestmentUserBatchTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReinvestmentUserBatchTransfered)
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
		it.Event = new(ContractReinvestmentUserBatchTransfered)
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
func (it *ContractReinvestmentUserBatchTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractReinvestmentUserBatchTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractReinvestmentUserBatchTransfered represents a UserBatchTransfered event raised by the ContractReinvestment contract.
type ContractReinvestmentUserBatchTransfered struct {
	UserIDs []*big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserBatchTransfered is a free log retrieval operation binding the contract event 0xcc838d9b3148e86bf2fbbff1e2e44c48a0fa2df6a550544ac2a77994221a8e21.
//
// Solidity: event userBatchTransfered(uint256[] userIDs, address to)
func (_ContractReinvestment *ContractReinvestmentFilterer) FilterUserBatchTransfered(opts *bind.FilterOpts) (*ContractReinvestmentUserBatchTransferedIterator, error) {

	logs, sub, err := _ContractReinvestment.contract.FilterLogs(opts, "userBatchTransfered")
	if err != nil {
		return nil, err
	}
	return &ContractReinvestmentUserBatchTransferedIterator{contract: _ContractReinvestment.contract, event: "userBatchTransfered", logs: logs, sub: sub}, nil
}

// WatchUserBatchTransfered is a free log subscription operation binding the contract event 0xcc838d9b3148e86bf2fbbff1e2e44c48a0fa2df6a550544ac2a77994221a8e21.
//
// Solidity: event userBatchTransfered(uint256[] userIDs, address to)
func (_ContractReinvestment *ContractReinvestmentFilterer) WatchUserBatchTransfered(opts *bind.WatchOpts, sink chan<- *ContractReinvestmentUserBatchTransfered) (event.Subscription, error) {

	logs, sub, err := _ContractReinvestment.contract.WatchLogs(opts, "userBatchTransfered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractReinvestmentUserBatchTransfered)
				if err := _ContractReinvestment.contract.UnpackLog(event, "userBatchTransfered", log); err != nil {
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

// ParseUserBatchTransfered is a log parse operation binding the contract event 0xcc838d9b3148e86bf2fbbff1e2e44c48a0fa2df6a550544ac2a77994221a8e21.
//
// Solidity: event userBatchTransfered(uint256[] userIDs, address to)
func (_ContractReinvestment *ContractReinvestmentFilterer) ParseUserBatchTransfered(log types.Log) (*ContractReinvestmentUserBatchTransfered, error) {
	event := new(ContractReinvestmentUserBatchTransfered)
	if err := _ContractReinvestment.contract.UnpackLog(event, "userBatchTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
