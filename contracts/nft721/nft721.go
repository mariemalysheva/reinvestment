// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_nft721

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

// ContractNft721MetaData contains all meta data concerning the ContractNft721 contract.
var ContractNft721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"maxSupply_\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BeanAddressNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRedeemer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMoreTokenIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedByRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedeemBeanNotOpen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistryNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongFrom\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"beanId\",\"type\":\"uint256\"}],\"name\":\"BeanRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRegistryActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorFilteringEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"realOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"beanIds\",\"type\":\"uint256[]\"}],\"name\":\"redeemBeans\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"redeemBeanOpen\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"beanAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURIPermanent\",\"type\":\"string\"}],\"name\":\"setBaseURIPermanent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBeanAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isRegistryActive\",\"type\":\"bool\"}],\"name\":\"setIsRegistryActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"setIsUriPermanent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_newSymbol\",\"type\":\"string\"}],\"name\":\"setNameAndSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setOperatorFilteringEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_redeemBeanOpen\",\"type\":\"bool\"}],\"name\":\"setRedeemBeanState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setTokenRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferLowerOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRealOwner\",\"type\":\"address\"}],\"name\":\"transferRealOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractNft721ABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractNft721MetaData.ABI instead.
var ContractNft721ABI = ContractNft721MetaData.ABI

// ContractNft721 is an auto generated Go binding around an Ethereum contract.
type ContractNft721 struct {
	ContractNft721Caller     // Read-only binding to the contract
	ContractNft721Transactor // Write-only binding to the contract
	ContractNft721Filterer   // Log filterer for contract events
}

// ContractNft721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ContractNft721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractNft721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractNft721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractNft721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractNft721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractNft721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractNft721Session struct {
	Contract     *ContractNft721   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractNft721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractNft721CallerSession struct {
	Contract *ContractNft721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ContractNft721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractNft721TransactorSession struct {
	Contract     *ContractNft721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractNft721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ContractNft721Raw struct {
	Contract *ContractNft721 // Generic contract binding to access the raw methods on
}

// ContractNft721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractNft721CallerRaw struct {
	Contract *ContractNft721Caller // Generic read-only contract binding to access the raw methods on
}

// ContractNft721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractNft721TransactorRaw struct {
	Contract *ContractNft721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewContractNft721 creates a new instance of ContractNft721, bound to a specific deployed contract.
func NewContractNft721(address common.Address, backend bind.ContractBackend) (*ContractNft721, error) {
	contract, err := bindContractNft721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractNft721{ContractNft721Caller: ContractNft721Caller{contract: contract}, ContractNft721Transactor: ContractNft721Transactor{contract: contract}, ContractNft721Filterer: ContractNft721Filterer{contract: contract}}, nil
}

// NewContractNft721Caller creates a new read-only instance of ContractNft721, bound to a specific deployed contract.
func NewContractNft721Caller(address common.Address, caller bind.ContractCaller) (*ContractNft721Caller, error) {
	contract, err := bindContractNft721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractNft721Caller{contract: contract}, nil
}

// NewContractNft721Transactor creates a new write-only instance of ContractNft721, bound to a specific deployed contract.
func NewContractNft721Transactor(address common.Address, transactor bind.ContractTransactor) (*ContractNft721Transactor, error) {
	contract, err := bindContractNft721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractNft721Transactor{contract: contract}, nil
}

// NewContractNft721Filterer creates a new log filterer instance of ContractNft721, bound to a specific deployed contract.
func NewContractNft721Filterer(address common.Address, filterer bind.ContractFilterer) (*ContractNft721Filterer, error) {
	contract, err := bindContractNft721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractNft721Filterer{contract: contract}, nil
}

// bindContractNft721 binds a generic wrapper to an already deployed contract.
func bindContractNft721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractNft721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractNft721 *ContractNft721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractNft721.Contract.ContractNft721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractNft721 *ContractNft721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractNft721.Contract.ContractNft721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractNft721 *ContractNft721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractNft721.Contract.ContractNft721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractNft721 *ContractNft721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractNft721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractNft721 *ContractNft721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractNft721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractNft721 *ContractNft721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractNft721.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint16)
func (_ContractNft721 *ContractNft721Caller) MAXSUPPLY(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint16)
func (_ContractNft721 *ContractNft721Session) MAXSUPPLY() (uint16, error) {
	return _ContractNft721.Contract.MAXSUPPLY(&_ContractNft721.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint16)
func (_ContractNft721 *ContractNft721CallerSession) MAXSUPPLY() (uint16, error) {
	return _ContractNft721.Contract.MAXSUPPLY(&_ContractNft721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ContractNft721 *ContractNft721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ContractNft721 *ContractNft721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ContractNft721.Contract.BalanceOf(&_ContractNft721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ContractNft721 *ContractNft721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ContractNft721.Contract.BalanceOf(&_ContractNft721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_ContractNft721 *ContractNft721Caller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_ContractNft721 *ContractNft721Session) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _ContractNft721.Contract.GetApproved(&_ContractNft721.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_ContractNft721 *ContractNft721CallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _ContractNft721.Contract.GetApproved(&_ContractNft721.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_ContractNft721 *ContractNft721Caller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_ContractNft721 *ContractNft721Session) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _ContractNft721.Contract.IsApprovedForAll(&_ContractNft721.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_ContractNft721 *ContractNft721CallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _ContractNft721.Contract.IsApprovedForAll(&_ContractNft721.CallOpts, arg0, arg1)
}

// IsRegistryActive is a free data retrieval call binding the contract method 0xabd017ea.
//
// Solidity: function isRegistryActive() view returns(bool)
func (_ContractNft721 *ContractNft721Caller) IsRegistryActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "isRegistryActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistryActive is a free data retrieval call binding the contract method 0xabd017ea.
//
// Solidity: function isRegistryActive() view returns(bool)
func (_ContractNft721 *ContractNft721Session) IsRegistryActive() (bool, error) {
	return _ContractNft721.Contract.IsRegistryActive(&_ContractNft721.CallOpts)
}

// IsRegistryActive is a free data retrieval call binding the contract method 0xabd017ea.
//
// Solidity: function isRegistryActive() view returns(bool)
func (_ContractNft721 *ContractNft721CallerSession) IsRegistryActive() (bool, error) {
	return _ContractNft721.Contract.IsRegistryActive(&_ContractNft721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractNft721 *ContractNft721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractNft721 *ContractNft721Session) Name() (string, error) {
	return _ContractNft721.Contract.Name(&_ContractNft721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractNft721 *ContractNft721CallerSession) Name() (string, error) {
	return _ContractNft721.Contract.Name(&_ContractNft721.CallOpts)
}

// OperatorFilteringEnabled is a free data retrieval call binding the contract method 0xfb796e6c.
//
// Solidity: function operatorFilteringEnabled() view returns(bool)
func (_ContractNft721 *ContractNft721Caller) OperatorFilteringEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "operatorFilteringEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorFilteringEnabled is a free data retrieval call binding the contract method 0xfb796e6c.
//
// Solidity: function operatorFilteringEnabled() view returns(bool)
func (_ContractNft721 *ContractNft721Session) OperatorFilteringEnabled() (bool, error) {
	return _ContractNft721.Contract.OperatorFilteringEnabled(&_ContractNft721.CallOpts)
}

// OperatorFilteringEnabled is a free data retrieval call binding the contract method 0xfb796e6c.
//
// Solidity: function operatorFilteringEnabled() view returns(bool)
func (_ContractNft721 *ContractNft721CallerSession) OperatorFilteringEnabled() (bool, error) {
	return _ContractNft721.Contract.OperatorFilteringEnabled(&_ContractNft721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractNft721 *ContractNft721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractNft721 *ContractNft721Session) Owner() (common.Address, error) {
	return _ContractNft721.Contract.Owner(&_ContractNft721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractNft721 *ContractNft721CallerSession) Owner() (common.Address, error) {
	return _ContractNft721.Contract.Owner(&_ContractNft721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_ContractNft721 *ContractNft721Caller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_ContractNft721 *ContractNft721Session) OwnerOf(id *big.Int) (common.Address, error) {
	return _ContractNft721.Contract.OwnerOf(&_ContractNft721.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_ContractNft721 *ContractNft721CallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _ContractNft721.Contract.OwnerOf(&_ContractNft721.CallOpts, id)
}

// RealOwner is a free data retrieval call binding the contract method 0x1df270f3.
//
// Solidity: function realOwner() view returns(address)
func (_ContractNft721 *ContractNft721Caller) RealOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "realOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RealOwner is a free data retrieval call binding the contract method 0x1df270f3.
//
// Solidity: function realOwner() view returns(address)
func (_ContractNft721 *ContractNft721Session) RealOwner() (common.Address, error) {
	return _ContractNft721.Contract.RealOwner(&_ContractNft721.CallOpts)
}

// RealOwner is a free data retrieval call binding the contract method 0x1df270f3.
//
// Solidity: function realOwner() view returns(address)
func (_ContractNft721 *ContractNft721CallerSession) RealOwner() (common.Address, error) {
	return _ContractNft721.Contract.RealOwner(&_ContractNft721.CallOpts)
}

// RedeemInfo is a free data retrieval call binding the contract method 0xefe7aa49.
//
// Solidity: function redeemInfo() view returns(bool redeemBeanOpen, address beanAddress)
func (_ContractNft721 *ContractNft721Caller) RedeemInfo(opts *bind.CallOpts) (struct {
	RedeemBeanOpen bool
	BeanAddress    common.Address
}, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "redeemInfo")

	outstruct := new(struct {
		RedeemBeanOpen bool
		BeanAddress    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RedeemBeanOpen = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BeanAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RedeemInfo is a free data retrieval call binding the contract method 0xefe7aa49.
//
// Solidity: function redeemInfo() view returns(bool redeemBeanOpen, address beanAddress)
func (_ContractNft721 *ContractNft721Session) RedeemInfo() (struct {
	RedeemBeanOpen bool
	BeanAddress    common.Address
}, error) {
	return _ContractNft721.Contract.RedeemInfo(&_ContractNft721.CallOpts)
}

// RedeemInfo is a free data retrieval call binding the contract method 0xefe7aa49.
//
// Solidity: function redeemInfo() view returns(bool redeemBeanOpen, address beanAddress)
func (_ContractNft721 *ContractNft721CallerSession) RedeemInfo() (struct {
	RedeemBeanOpen bool
	BeanAddress    common.Address
}, error) {
	return _ContractNft721.Contract.RedeemInfo(&_ContractNft721.CallOpts)
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() view returns(address)
func (_ContractNft721 *ContractNft721Caller) RegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "registryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() view returns(address)
func (_ContractNft721 *ContractNft721Session) RegistryAddress() (common.Address, error) {
	return _ContractNft721.Contract.RegistryAddress(&_ContractNft721.CallOpts)
}

// RegistryAddress is a free data retrieval call binding the contract method 0xed9aab51.
//
// Solidity: function registryAddress() view returns(address)
func (_ContractNft721 *ContractNft721CallerSession) RegistryAddress() (common.Address, error) {
	return _ContractNft721.Contract.RegistryAddress(&_ContractNft721.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_ContractNft721 *ContractNft721Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_ContractNft721 *ContractNft721Session) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ContractNft721.Contract.RoyaltyInfo(&_ContractNft721.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_ContractNft721 *ContractNft721CallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ContractNft721.Contract.RoyaltyInfo(&_ContractNft721.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractNft721 *ContractNft721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractNft721 *ContractNft721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractNft721.Contract.SupportsInterface(&_ContractNft721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractNft721 *ContractNft721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractNft721.Contract.SupportsInterface(&_ContractNft721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractNft721 *ContractNft721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractNft721 *ContractNft721Session) Symbol() (string, error) {
	return _ContractNft721.Contract.Symbol(&_ContractNft721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractNft721 *ContractNft721CallerSession) Symbol() (string, error) {
	return _ContractNft721.Contract.Symbol(&_ContractNft721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ContractNft721 *ContractNft721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ContractNft721 *ContractNft721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _ContractNft721.Contract.TokenURI(&_ContractNft721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ContractNft721 *ContractNft721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ContractNft721.Contract.TokenURI(&_ContractNft721.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractNft721 *ContractNft721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractNft721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractNft721 *ContractNft721Session) TotalSupply() (*big.Int, error) {
	return _ContractNft721.Contract.TotalSupply(&_ContractNft721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractNft721 *ContractNft721CallerSession) TotalSupply() (*big.Int, error) {
	return _ContractNft721.Contract.TotalSupply(&_ContractNft721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_ContractNft721 *ContractNft721Transactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_ContractNft721 *ContractNft721Session) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.Approve(&_ContractNft721.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_ContractNft721 *ContractNft721TransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.Approve(&_ContractNft721.TransactOpts, operator, tokenId)
}

// RedeemBeans is a paid mutator transaction binding the contract method 0xd443af80.
//
// Solidity: function redeemBeans(address to, uint256[] beanIds) returns(uint256[])
func (_ContractNft721 *ContractNft721Transactor) RedeemBeans(opts *bind.TransactOpts, to common.Address, beanIds []*big.Int) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "redeemBeans", to, beanIds)
}

// RedeemBeans is a paid mutator transaction binding the contract method 0xd443af80.
//
// Solidity: function redeemBeans(address to, uint256[] beanIds) returns(uint256[])
func (_ContractNft721 *ContractNft721Session) RedeemBeans(to common.Address, beanIds []*big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.RedeemBeans(&_ContractNft721.TransactOpts, to, beanIds)
}

// RedeemBeans is a paid mutator transaction binding the contract method 0xd443af80.
//
// Solidity: function redeemBeans(address to, uint256[] beanIds) returns(uint256[])
func (_ContractNft721 *ContractNft721TransactorSession) RedeemBeans(to common.Address, beanIds []*big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.RedeemBeans(&_ContractNft721.TransactOpts, to, beanIds)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractNft721 *ContractNft721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractNft721 *ContractNft721Session) RenounceOwnership() (*types.Transaction, error) {
	return _ContractNft721.Contract.RenounceOwnership(&_ContractNft721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractNft721 *ContractNft721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractNft721.Contract.RenounceOwnership(&_ContractNft721.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_ContractNft721 *ContractNft721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_ContractNft721 *ContractNft721Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.SafeTransferFrom(&_ContractNft721.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.SafeTransferFrom(&_ContractNft721.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_ContractNft721 *ContractNft721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_ContractNft721 *ContractNft721Session) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _ContractNft721.Contract.SafeTransferFrom0(&_ContractNft721.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _ContractNft721.Contract.SafeTransferFrom0(&_ContractNft721.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ContractNft721 *ContractNft721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ContractNft721 *ContractNft721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetApprovalForAll(&_ContractNft721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetApprovalForAll(&_ContractNft721.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_ContractNft721 *ContractNft721Transactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_ContractNft721 *ContractNft721Session) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetBaseURI(&_ContractNft721.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetBaseURI(&_ContractNft721.TransactOpts, baseURI)
}

// SetBaseURIPermanent is a paid mutator transaction binding the contract method 0x3c115fa6.
//
// Solidity: function setBaseURIPermanent(string baseURIPermanent) returns()
func (_ContractNft721 *ContractNft721Transactor) SetBaseURIPermanent(opts *bind.TransactOpts, baseURIPermanent string) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setBaseURIPermanent", baseURIPermanent)
}

// SetBaseURIPermanent is a paid mutator transaction binding the contract method 0x3c115fa6.
//
// Solidity: function setBaseURIPermanent(string baseURIPermanent) returns()
func (_ContractNft721 *ContractNft721Session) SetBaseURIPermanent(baseURIPermanent string) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetBaseURIPermanent(&_ContractNft721.TransactOpts, baseURIPermanent)
}

// SetBaseURIPermanent is a paid mutator transaction binding the contract method 0x3c115fa6.
//
// Solidity: function setBaseURIPermanent(string baseURIPermanent) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetBaseURIPermanent(baseURIPermanent string) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetBaseURIPermanent(&_ContractNft721.TransactOpts, baseURIPermanent)
}

// SetBeanAddress is a paid mutator transaction binding the contract method 0x54ecf309.
//
// Solidity: function setBeanAddress(address contractAddress) returns()
func (_ContractNft721 *ContractNft721Transactor) SetBeanAddress(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setBeanAddress", contractAddress)
}

// SetBeanAddress is a paid mutator transaction binding the contract method 0x54ecf309.
//
// Solidity: function setBeanAddress(address contractAddress) returns()
func (_ContractNft721 *ContractNft721Session) SetBeanAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetBeanAddress(&_ContractNft721.TransactOpts, contractAddress)
}

// SetBeanAddress is a paid mutator transaction binding the contract method 0x54ecf309.
//
// Solidity: function setBeanAddress(address contractAddress) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetBeanAddress(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetBeanAddress(&_ContractNft721.TransactOpts, contractAddress)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_ContractNft721 *ContractNft721Transactor) SetDefaultRoyalty(opts *bind.TransactOpts, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setDefaultRoyalty", receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_ContractNft721 *ContractNft721Session) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetDefaultRoyalty(&_ContractNft721.TransactOpts, receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetDefaultRoyalty(&_ContractNft721.TransactOpts, receiver, feeNumerator)
}

// SetIsRegistryActive is a paid mutator transaction binding the contract method 0x46fff98d.
//
// Solidity: function setIsRegistryActive(bool _isRegistryActive) returns()
func (_ContractNft721 *ContractNft721Transactor) SetIsRegistryActive(opts *bind.TransactOpts, _isRegistryActive bool) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setIsRegistryActive", _isRegistryActive)
}

// SetIsRegistryActive is a paid mutator transaction binding the contract method 0x46fff98d.
//
// Solidity: function setIsRegistryActive(bool _isRegistryActive) returns()
func (_ContractNft721 *ContractNft721Session) SetIsRegistryActive(_isRegistryActive bool) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetIsRegistryActive(&_ContractNft721.TransactOpts, _isRegistryActive)
}

// SetIsRegistryActive is a paid mutator transaction binding the contract method 0x46fff98d.
//
// Solidity: function setIsRegistryActive(bool _isRegistryActive) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetIsRegistryActive(_isRegistryActive bool) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetIsRegistryActive(&_ContractNft721.TransactOpts, _isRegistryActive)
}

// SetIsUriPermanent is a paid mutator transaction binding the contract method 0x8e0d9fcc.
//
// Solidity: function setIsUriPermanent(uint256[] tokenIds) returns()
func (_ContractNft721 *ContractNft721Transactor) SetIsUriPermanent(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setIsUriPermanent", tokenIds)
}

// SetIsUriPermanent is a paid mutator transaction binding the contract method 0x8e0d9fcc.
//
// Solidity: function setIsUriPermanent(uint256[] tokenIds) returns()
func (_ContractNft721 *ContractNft721Session) SetIsUriPermanent(tokenIds []*big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetIsUriPermanent(&_ContractNft721.TransactOpts, tokenIds)
}

// SetIsUriPermanent is a paid mutator transaction binding the contract method 0x8e0d9fcc.
//
// Solidity: function setIsUriPermanent(uint256[] tokenIds) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetIsUriPermanent(tokenIds []*big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetIsUriPermanent(&_ContractNft721.TransactOpts, tokenIds)
}

// SetNameAndSymbol is a paid mutator transaction binding the contract method 0x5a446215.
//
// Solidity: function setNameAndSymbol(string _newName, string _newSymbol) returns()
func (_ContractNft721 *ContractNft721Transactor) SetNameAndSymbol(opts *bind.TransactOpts, _newName string, _newSymbol string) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setNameAndSymbol", _newName, _newSymbol)
}

// SetNameAndSymbol is a paid mutator transaction binding the contract method 0x5a446215.
//
// Solidity: function setNameAndSymbol(string _newName, string _newSymbol) returns()
func (_ContractNft721 *ContractNft721Session) SetNameAndSymbol(_newName string, _newSymbol string) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetNameAndSymbol(&_ContractNft721.TransactOpts, _newName, _newSymbol)
}

// SetNameAndSymbol is a paid mutator transaction binding the contract method 0x5a446215.
//
// Solidity: function setNameAndSymbol(string _newName, string _newSymbol) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetNameAndSymbol(_newName string, _newSymbol string) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetNameAndSymbol(&_ContractNft721.TransactOpts, _newName, _newSymbol)
}

// SetOperatorFilteringEnabled is a paid mutator transaction binding the contract method 0xb7c0b8e8.
//
// Solidity: function setOperatorFilteringEnabled(bool value) returns()
func (_ContractNft721 *ContractNft721Transactor) SetOperatorFilteringEnabled(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setOperatorFilteringEnabled", value)
}

// SetOperatorFilteringEnabled is a paid mutator transaction binding the contract method 0xb7c0b8e8.
//
// Solidity: function setOperatorFilteringEnabled(bool value) returns()
func (_ContractNft721 *ContractNft721Session) SetOperatorFilteringEnabled(value bool) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetOperatorFilteringEnabled(&_ContractNft721.TransactOpts, value)
}

// SetOperatorFilteringEnabled is a paid mutator transaction binding the contract method 0xb7c0b8e8.
//
// Solidity: function setOperatorFilteringEnabled(bool value) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetOperatorFilteringEnabled(value bool) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetOperatorFilteringEnabled(&_ContractNft721.TransactOpts, value)
}

// SetRedeemBeanState is a paid mutator transaction binding the contract method 0x90d9c86a.
//
// Solidity: function setRedeemBeanState(bool _redeemBeanOpen) returns()
func (_ContractNft721 *ContractNft721Transactor) SetRedeemBeanState(opts *bind.TransactOpts, _redeemBeanOpen bool) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setRedeemBeanState", _redeemBeanOpen)
}

// SetRedeemBeanState is a paid mutator transaction binding the contract method 0x90d9c86a.
//
// Solidity: function setRedeemBeanState(bool _redeemBeanOpen) returns()
func (_ContractNft721 *ContractNft721Session) SetRedeemBeanState(_redeemBeanOpen bool) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetRedeemBeanState(&_ContractNft721.TransactOpts, _redeemBeanOpen)
}

// SetRedeemBeanState is a paid mutator transaction binding the contract method 0x90d9c86a.
//
// Solidity: function setRedeemBeanState(bool _redeemBeanOpen) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetRedeemBeanState(_redeemBeanOpen bool) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetRedeemBeanState(&_ContractNft721.TransactOpts, _redeemBeanOpen)
}

// SetRegistryAddress is a paid mutator transaction binding the contract method 0xab7b4993.
//
// Solidity: function setRegistryAddress(address _registryAddress) returns()
func (_ContractNft721 *ContractNft721Transactor) SetRegistryAddress(opts *bind.TransactOpts, _registryAddress common.Address) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setRegistryAddress", _registryAddress)
}

// SetRegistryAddress is a paid mutator transaction binding the contract method 0xab7b4993.
//
// Solidity: function setRegistryAddress(address _registryAddress) returns()
func (_ContractNft721 *ContractNft721Session) SetRegistryAddress(_registryAddress common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetRegistryAddress(&_ContractNft721.TransactOpts, _registryAddress)
}

// SetRegistryAddress is a paid mutator transaction binding the contract method 0xab7b4993.
//
// Solidity: function setRegistryAddress(address _registryAddress) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetRegistryAddress(_registryAddress common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetRegistryAddress(&_ContractNft721.TransactOpts, _registryAddress)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_ContractNft721 *ContractNft721Transactor) SetTokenRoyalty(opts *bind.TransactOpts, tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "setTokenRoyalty", tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_ContractNft721 *ContractNft721Session) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetTokenRoyalty(&_ContractNft721.TransactOpts, tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_ContractNft721 *ContractNft721TransactorSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.SetTokenRoyalty(&_ContractNft721.TransactOpts, tokenId, receiver, feeNumerator)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_ContractNft721 *ContractNft721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_ContractNft721 *ContractNft721Session) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.TransferFrom(&_ContractNft721.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_ContractNft721 *ContractNft721TransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _ContractNft721.Contract.TransferFrom(&_ContractNft721.TransactOpts, from, to, id)
}

// TransferLowerOwnership is a paid mutator transaction binding the contract method 0x09af3f9a.
//
// Solidity: function transferLowerOwnership(address newOwner) returns()
func (_ContractNft721 *ContractNft721Transactor) TransferLowerOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "transferLowerOwnership", newOwner)
}

// TransferLowerOwnership is a paid mutator transaction binding the contract method 0x09af3f9a.
//
// Solidity: function transferLowerOwnership(address newOwner) returns()
func (_ContractNft721 *ContractNft721Session) TransferLowerOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.TransferLowerOwnership(&_ContractNft721.TransactOpts, newOwner)
}

// TransferLowerOwnership is a paid mutator transaction binding the contract method 0x09af3f9a.
//
// Solidity: function transferLowerOwnership(address newOwner) returns()
func (_ContractNft721 *ContractNft721TransactorSession) TransferLowerOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.TransferLowerOwnership(&_ContractNft721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractNft721 *ContractNft721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractNft721 *ContractNft721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.TransferOwnership(&_ContractNft721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractNft721 *ContractNft721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.TransferOwnership(&_ContractNft721.TransactOpts, newOwner)
}

// TransferRealOwnership is a paid mutator transaction binding the contract method 0x2cff6770.
//
// Solidity: function transferRealOwnership(address newRealOwner) returns()
func (_ContractNft721 *ContractNft721Transactor) TransferRealOwnership(opts *bind.TransactOpts, newRealOwner common.Address) (*types.Transaction, error) {
	return _ContractNft721.contract.Transact(opts, "transferRealOwnership", newRealOwner)
}

// TransferRealOwnership is a paid mutator transaction binding the contract method 0x2cff6770.
//
// Solidity: function transferRealOwnership(address newRealOwner) returns()
func (_ContractNft721 *ContractNft721Session) TransferRealOwnership(newRealOwner common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.TransferRealOwnership(&_ContractNft721.TransactOpts, newRealOwner)
}

// TransferRealOwnership is a paid mutator transaction binding the contract method 0x2cff6770.
//
// Solidity: function transferRealOwnership(address newRealOwner) returns()
func (_ContractNft721 *ContractNft721TransactorSession) TransferRealOwnership(newRealOwner common.Address) (*types.Transaction, error) {
	return _ContractNft721.Contract.TransferRealOwnership(&_ContractNft721.TransactOpts, newRealOwner)
}

// ContractNft721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ContractNft721 contract.
type ContractNft721ApprovalIterator struct {
	Event *ContractNft721Approval // Event containing the contract specifics and raw log

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
func (it *ContractNft721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNft721Approval)
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
		it.Event = new(ContractNft721Approval)
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
func (it *ContractNft721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNft721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNft721Approval represents a Approval event raised by the ContractNft721 contract.
type ContractNft721Approval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_ContractNft721 *ContractNft721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*ContractNft721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractNft721.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ContractNft721ApprovalIterator{contract: _ContractNft721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_ContractNft721 *ContractNft721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractNft721Approval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractNft721.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNft721Approval)
				if err := _ContractNft721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_ContractNft721 *ContractNft721Filterer) ParseApproval(log types.Log) (*ContractNft721Approval, error) {
	event := new(ContractNft721Approval)
	if err := _ContractNft721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNft721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ContractNft721 contract.
type ContractNft721ApprovalForAllIterator struct {
	Event *ContractNft721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ContractNft721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNft721ApprovalForAll)
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
		it.Event = new(ContractNft721ApprovalForAll)
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
func (it *ContractNft721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNft721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNft721ApprovalForAll represents a ApprovalForAll event raised by the ContractNft721 contract.
type ContractNft721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ContractNft721 *ContractNft721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractNft721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractNft721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractNft721ApprovalForAllIterator{contract: _ContractNft721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ContractNft721 *ContractNft721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractNft721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractNft721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNft721ApprovalForAll)
				if err := _ContractNft721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ContractNft721 *ContractNft721Filterer) ParseApprovalForAll(log types.Log) (*ContractNft721ApprovalForAll, error) {
	event := new(ContractNft721ApprovalForAll)
	if err := _ContractNft721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNft721BeanRedeemedIterator is returned from FilterBeanRedeemed and is used to iterate over the raw logs and unpacked data for BeanRedeemed events raised by the ContractNft721 contract.
type ContractNft721BeanRedeemedIterator struct {
	Event *ContractNft721BeanRedeemed // Event containing the contract specifics and raw log

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
func (it *ContractNft721BeanRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNft721BeanRedeemed)
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
		it.Event = new(ContractNft721BeanRedeemed)
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
func (it *ContractNft721BeanRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNft721BeanRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNft721BeanRedeemed represents a BeanRedeemed event raised by the ContractNft721 contract.
type ContractNft721BeanRedeemed struct {
	To      common.Address
	TokenId *big.Int
	BeanId  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBeanRedeemed is a free log retrieval operation binding the contract event 0x71e92ec5a8e5a2b5c4717c4520cea4ce2cfe75c53bba10efe28c4328b31047cb.
//
// Solidity: event BeanRedeemed(address indexed to, uint256 indexed tokenId, uint256 indexed beanId)
func (_ContractNft721 *ContractNft721Filterer) FilterBeanRedeemed(opts *bind.FilterOpts, to []common.Address, tokenId []*big.Int, beanId []*big.Int) (*ContractNft721BeanRedeemedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beanIdRule []interface{}
	for _, beanIdItem := range beanId {
		beanIdRule = append(beanIdRule, beanIdItem)
	}

	logs, sub, err := _ContractNft721.contract.FilterLogs(opts, "BeanRedeemed", toRule, tokenIdRule, beanIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractNft721BeanRedeemedIterator{contract: _ContractNft721.contract, event: "BeanRedeemed", logs: logs, sub: sub}, nil
}

// WatchBeanRedeemed is a free log subscription operation binding the contract event 0x71e92ec5a8e5a2b5c4717c4520cea4ce2cfe75c53bba10efe28c4328b31047cb.
//
// Solidity: event BeanRedeemed(address indexed to, uint256 indexed tokenId, uint256 indexed beanId)
func (_ContractNft721 *ContractNft721Filterer) WatchBeanRedeemed(opts *bind.WatchOpts, sink chan<- *ContractNft721BeanRedeemed, to []common.Address, tokenId []*big.Int, beanId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var beanIdRule []interface{}
	for _, beanIdItem := range beanId {
		beanIdRule = append(beanIdRule, beanIdItem)
	}

	logs, sub, err := _ContractNft721.contract.WatchLogs(opts, "BeanRedeemed", toRule, tokenIdRule, beanIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNft721BeanRedeemed)
				if err := _ContractNft721.contract.UnpackLog(event, "BeanRedeemed", log); err != nil {
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

// ParseBeanRedeemed is a log parse operation binding the contract event 0x71e92ec5a8e5a2b5c4717c4520cea4ce2cfe75c53bba10efe28c4328b31047cb.
//
// Solidity: event BeanRedeemed(address indexed to, uint256 indexed tokenId, uint256 indexed beanId)
func (_ContractNft721 *ContractNft721Filterer) ParseBeanRedeemed(log types.Log) (*ContractNft721BeanRedeemed, error) {
	event := new(ContractNft721BeanRedeemed)
	if err := _ContractNft721.contract.UnpackLog(event, "BeanRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNft721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractNft721 contract.
type ContractNft721OwnershipTransferredIterator struct {
	Event *ContractNft721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractNft721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNft721OwnershipTransferred)
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
		it.Event = new(ContractNft721OwnershipTransferred)
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
func (it *ContractNft721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNft721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNft721OwnershipTransferred represents a OwnershipTransferred event raised by the ContractNft721 contract.
type ContractNft721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractNft721 *ContractNft721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractNft721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractNft721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractNft721OwnershipTransferredIterator{contract: _ContractNft721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractNft721 *ContractNft721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractNft721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractNft721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNft721OwnershipTransferred)
				if err := _ContractNft721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractNft721 *ContractNft721Filterer) ParseOwnershipTransferred(log types.Log) (*ContractNft721OwnershipTransferred, error) {
	event := new(ContractNft721OwnershipTransferred)
	if err := _ContractNft721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNft721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ContractNft721 contract.
type ContractNft721TransferIterator struct {
	Event *ContractNft721Transfer // Event containing the contract specifics and raw log

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
func (it *ContractNft721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNft721Transfer)
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
		it.Event = new(ContractNft721Transfer)
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
func (it *ContractNft721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNft721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNft721Transfer represents a Transfer event raised by the ContractNft721 contract.
type ContractNft721Transfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_ContractNft721 *ContractNft721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*ContractNft721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractNft721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ContractNft721TransferIterator{contract: _ContractNft721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_ContractNft721 *ContractNft721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractNft721Transfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ContractNft721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNft721Transfer)
				if err := _ContractNft721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_ContractNft721 *ContractNft721Filterer) ParseTransfer(log types.Log) (*ContractNft721Transfer, error) {
	event := new(ContractNft721Transfer)
	if err := _ContractNft721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
