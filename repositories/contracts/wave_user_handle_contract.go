// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// WaveUserHandleContractMetaData contains all meta data concerning the WaveUserHandleContract contract.
var WaveUserHandleContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userHandle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"claimUserHandle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userHandle\",\"type\":\"string\"}],\"name\":\"getUUIDByUserHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getUserHandleByTokenId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WaveUserHandleContractABI is the input ABI used to generate the binding from.
// Deprecated: Use WaveUserHandleContractMetaData.ABI instead.
var WaveUserHandleContractABI = WaveUserHandleContractMetaData.ABI

// WaveUserHandleContract is an auto generated Go binding around an Ethereum contract.
type WaveUserHandleContract struct {
	WaveUserHandleContractCaller     // Read-only binding to the contract
	WaveUserHandleContractTransactor // Write-only binding to the contract
	WaveUserHandleContractFilterer   // Log filterer for contract events
}

// WaveUserHandleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type WaveUserHandleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WaveUserHandleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WaveUserHandleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WaveUserHandleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WaveUserHandleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WaveUserHandleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WaveUserHandleContractSession struct {
	Contract     *WaveUserHandleContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WaveUserHandleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WaveUserHandleContractCallerSession struct {
	Contract *WaveUserHandleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// WaveUserHandleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WaveUserHandleContractTransactorSession struct {
	Contract     *WaveUserHandleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// WaveUserHandleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type WaveUserHandleContractRaw struct {
	Contract *WaveUserHandleContract // Generic contract binding to access the raw methods on
}

// WaveUserHandleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WaveUserHandleContractCallerRaw struct {
	Contract *WaveUserHandleContractCaller // Generic read-only contract binding to access the raw methods on
}

// WaveUserHandleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WaveUserHandleContractTransactorRaw struct {
	Contract *WaveUserHandleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWaveUserHandleContract creates a new instance of WaveUserHandleContract, bound to a specific deployed contract.
func NewWaveUserHandleContract(address common.Address, backend bind.ContractBackend) (*WaveUserHandleContract, error) {
	contract, err := bindWaveUserHandleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WaveUserHandleContract{WaveUserHandleContractCaller: WaveUserHandleContractCaller{contract: contract}, WaveUserHandleContractTransactor: WaveUserHandleContractTransactor{contract: contract}, WaveUserHandleContractFilterer: WaveUserHandleContractFilterer{contract: contract}}, nil
}

// NewWaveUserHandleContractCaller creates a new read-only instance of WaveUserHandleContract, bound to a specific deployed contract.
func NewWaveUserHandleContractCaller(address common.Address, caller bind.ContractCaller) (*WaveUserHandleContractCaller, error) {
	contract, err := bindWaveUserHandleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WaveUserHandleContractCaller{contract: contract}, nil
}

// NewWaveUserHandleContractTransactor creates a new write-only instance of WaveUserHandleContract, bound to a specific deployed contract.
func NewWaveUserHandleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*WaveUserHandleContractTransactor, error) {
	contract, err := bindWaveUserHandleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WaveUserHandleContractTransactor{contract: contract}, nil
}

// NewWaveUserHandleContractFilterer creates a new log filterer instance of WaveUserHandleContract, bound to a specific deployed contract.
func NewWaveUserHandleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*WaveUserHandleContractFilterer, error) {
	contract, err := bindWaveUserHandleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WaveUserHandleContractFilterer{contract: contract}, nil
}

// bindWaveUserHandleContract binds a generic wrapper to an already deployed contract.
func bindWaveUserHandleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WaveUserHandleContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WaveUserHandleContract *WaveUserHandleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WaveUserHandleContract.Contract.WaveUserHandleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WaveUserHandleContract *WaveUserHandleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.WaveUserHandleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WaveUserHandleContract *WaveUserHandleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.WaveUserHandleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WaveUserHandleContract *WaveUserHandleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WaveUserHandleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WaveUserHandleContract *WaveUserHandleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WaveUserHandleContract *WaveUserHandleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_WaveUserHandleContract *WaveUserHandleContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WaveUserHandleContract.Contract.BalanceOf(&_WaveUserHandleContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _WaveUserHandleContract.Contract.BalanceOf(&_WaveUserHandleContract.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_WaveUserHandleContract *WaveUserHandleContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _WaveUserHandleContract.Contract.GetApproved(&_WaveUserHandleContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _WaveUserHandleContract.Contract.GetApproved(&_WaveUserHandleContract.CallOpts, tokenId)
}

// GetUUIDByUserHandle is a free data retrieval call binding the contract method 0x98574de8.
//
// Solidity: function getUUIDByUserHandle(string userHandle) view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) GetUUIDByUserHandle(opts *bind.CallOpts, userHandle string) (string, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "getUUIDByUserHandle", userHandle)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUUIDByUserHandle is a free data retrieval call binding the contract method 0x98574de8.
//
// Solidity: function getUUIDByUserHandle(string userHandle) view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractSession) GetUUIDByUserHandle(userHandle string) (string, error) {
	return _WaveUserHandleContract.Contract.GetUUIDByUserHandle(&_WaveUserHandleContract.CallOpts, userHandle)
}

// GetUUIDByUserHandle is a free data retrieval call binding the contract method 0x98574de8.
//
// Solidity: function getUUIDByUserHandle(string userHandle) view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) GetUUIDByUserHandle(userHandle string) (string, error) {
	return _WaveUserHandleContract.Contract.GetUUIDByUserHandle(&_WaveUserHandleContract.CallOpts, userHandle)
}

// GetUserHandleByTokenId is a free data retrieval call binding the contract method 0x8a4cc184.
//
// Solidity: function getUserHandleByTokenId(uint256 tokenId) view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) GetUserHandleByTokenId(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "getUserHandleByTokenId", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserHandleByTokenId is a free data retrieval call binding the contract method 0x8a4cc184.
//
// Solidity: function getUserHandleByTokenId(uint256 tokenId) view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractSession) GetUserHandleByTokenId(tokenId *big.Int) (string, error) {
	return _WaveUserHandleContract.Contract.GetUserHandleByTokenId(&_WaveUserHandleContract.CallOpts, tokenId)
}

// GetUserHandleByTokenId is a free data retrieval call binding the contract method 0x8a4cc184.
//
// Solidity: function getUserHandleByTokenId(uint256 tokenId) view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) GetUserHandleByTokenId(tokenId *big.Int) (string, error) {
	return _WaveUserHandleContract.Contract.GetUserHandleByTokenId(&_WaveUserHandleContract.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_WaveUserHandleContract *WaveUserHandleContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WaveUserHandleContract.Contract.IsApprovedForAll(&_WaveUserHandleContract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _WaveUserHandleContract.Contract.IsApprovedForAll(&_WaveUserHandleContract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractSession) Name() (string, error) {
	return _WaveUserHandleContract.Contract.Name(&_WaveUserHandleContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) Name() (string, error) {
	return _WaveUserHandleContract.Contract.Name(&_WaveUserHandleContract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_WaveUserHandleContract *WaveUserHandleContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _WaveUserHandleContract.Contract.OwnerOf(&_WaveUserHandleContract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _WaveUserHandleContract.Contract.OwnerOf(&_WaveUserHandleContract.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WaveUserHandleContract *WaveUserHandleContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WaveUserHandleContract.Contract.SupportsInterface(&_WaveUserHandleContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WaveUserHandleContract.Contract.SupportsInterface(&_WaveUserHandleContract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractSession) Symbol() (string, error) {
	return _WaveUserHandleContract.Contract.Symbol(&_WaveUserHandleContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) Symbol() (string, error) {
	return _WaveUserHandleContract.Contract.Symbol(&_WaveUserHandleContract.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _WaveUserHandleContract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _WaveUserHandleContract.Contract.TokenURI(&_WaveUserHandleContract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_WaveUserHandleContract *WaveUserHandleContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _WaveUserHandleContract.Contract.TokenURI(&_WaveUserHandleContract.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WaveUserHandleContract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_WaveUserHandleContract *WaveUserHandleContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.Approve(&_WaveUserHandleContract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.Approve(&_WaveUserHandleContract.TransactOpts, to, tokenId)
}

// ClaimUserHandle is a paid mutator transaction binding the contract method 0xb372c02f.
//
// Solidity: function claimUserHandle(string userHandle, string uuid, string tokenURI) returns(uint256)
func (_WaveUserHandleContract *WaveUserHandleContractTransactor) ClaimUserHandle(opts *bind.TransactOpts, userHandle string, uuid string, tokenURI string) (*types.Transaction, error) {
	return _WaveUserHandleContract.contract.Transact(opts, "claimUserHandle", userHandle, uuid, tokenURI)
}

// ClaimUserHandle is a paid mutator transaction binding the contract method 0xb372c02f.
//
// Solidity: function claimUserHandle(string userHandle, string uuid, string tokenURI) returns(uint256)
func (_WaveUserHandleContract *WaveUserHandleContractSession) ClaimUserHandle(userHandle string, uuid string, tokenURI string) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.ClaimUserHandle(&_WaveUserHandleContract.TransactOpts, userHandle, uuid, tokenURI)
}

// ClaimUserHandle is a paid mutator transaction binding the contract method 0xb372c02f.
//
// Solidity: function claimUserHandle(string userHandle, string uuid, string tokenURI) returns(uint256)
func (_WaveUserHandleContract *WaveUserHandleContractTransactorSession) ClaimUserHandle(userHandle string, uuid string, tokenURI string) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.ClaimUserHandle(&_WaveUserHandleContract.TransactOpts, userHandle, uuid, tokenURI)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WaveUserHandleContract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WaveUserHandleContract *WaveUserHandleContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.SafeTransferFrom(&_WaveUserHandleContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.SafeTransferFrom(&_WaveUserHandleContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _WaveUserHandleContract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_WaveUserHandleContract *WaveUserHandleContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.SafeTransferFrom0(&_WaveUserHandleContract.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.SafeTransferFrom0(&_WaveUserHandleContract.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _WaveUserHandleContract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_WaveUserHandleContract *WaveUserHandleContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.SetApprovalForAll(&_WaveUserHandleContract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.SetApprovalForAll(&_WaveUserHandleContract.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WaveUserHandleContract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WaveUserHandleContract *WaveUserHandleContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.TransferFrom(&_WaveUserHandleContract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_WaveUserHandleContract *WaveUserHandleContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _WaveUserHandleContract.Contract.TransferFrom(&_WaveUserHandleContract.TransactOpts, from, to, tokenId)
}

// WaveUserHandleContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WaveUserHandleContract contract.
type WaveUserHandleContractApprovalIterator struct {
	Event *WaveUserHandleContractApproval // Event containing the contract specifics and raw log

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
func (it *WaveUserHandleContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WaveUserHandleContractApproval)
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
		it.Event = new(WaveUserHandleContractApproval)
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
func (it *WaveUserHandleContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WaveUserHandleContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WaveUserHandleContractApproval represents a Approval event raised by the WaveUserHandleContract contract.
type WaveUserHandleContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WaveUserHandleContract *WaveUserHandleContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*WaveUserHandleContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WaveUserHandleContract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WaveUserHandleContractApprovalIterator{contract: _WaveUserHandleContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WaveUserHandleContract *WaveUserHandleContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WaveUserHandleContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WaveUserHandleContract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WaveUserHandleContractApproval)
				if err := _WaveUserHandleContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_WaveUserHandleContract *WaveUserHandleContractFilterer) ParseApproval(log types.Log) (*WaveUserHandleContractApproval, error) {
	event := new(WaveUserHandleContractApproval)
	if err := _WaveUserHandleContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WaveUserHandleContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the WaveUserHandleContract contract.
type WaveUserHandleContractApprovalForAllIterator struct {
	Event *WaveUserHandleContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *WaveUserHandleContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WaveUserHandleContractApprovalForAll)
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
		it.Event = new(WaveUserHandleContractApprovalForAll)
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
func (it *WaveUserHandleContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WaveUserHandleContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WaveUserHandleContractApprovalForAll represents a ApprovalForAll event raised by the WaveUserHandleContract contract.
type WaveUserHandleContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WaveUserHandleContract *WaveUserHandleContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*WaveUserHandleContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WaveUserHandleContract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &WaveUserHandleContractApprovalForAllIterator{contract: _WaveUserHandleContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_WaveUserHandleContract *WaveUserHandleContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *WaveUserHandleContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _WaveUserHandleContract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WaveUserHandleContractApprovalForAll)
				if err := _WaveUserHandleContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_WaveUserHandleContract *WaveUserHandleContractFilterer) ParseApprovalForAll(log types.Log) (*WaveUserHandleContractApprovalForAll, error) {
	event := new(WaveUserHandleContractApprovalForAll)
	if err := _WaveUserHandleContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WaveUserHandleContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WaveUserHandleContract contract.
type WaveUserHandleContractTransferIterator struct {
	Event *WaveUserHandleContractTransfer // Event containing the contract specifics and raw log

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
func (it *WaveUserHandleContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WaveUserHandleContractTransfer)
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
		it.Event = new(WaveUserHandleContractTransfer)
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
func (it *WaveUserHandleContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WaveUserHandleContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WaveUserHandleContractTransfer represents a Transfer event raised by the WaveUserHandleContract contract.
type WaveUserHandleContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WaveUserHandleContract *WaveUserHandleContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*WaveUserHandleContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WaveUserHandleContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &WaveUserHandleContractTransferIterator{contract: _WaveUserHandleContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WaveUserHandleContract *WaveUserHandleContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WaveUserHandleContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _WaveUserHandleContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WaveUserHandleContractTransfer)
				if err := _WaveUserHandleContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_WaveUserHandleContract *WaveUserHandleContractFilterer) ParseTransfer(log types.Log) (*WaveUserHandleContractTransfer, error) {
	event := new(WaveUserHandleContractTransfer)
	if err := _WaveUserHandleContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
