// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hedgey

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

// HedgeyMetaData contains all meta data concerning the Hedgey contract.
var HedgeyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_holder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockDate\",\"type\":\"uint256\"}],\"name\":\"NFTCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_holder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unlockDate\",\"type\":\"uint256\"}],\"name\":\"NFTRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"URISet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDate\",\"type\":\"uint256\"}],\"name\":\"createNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"futures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"redeemNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"updateBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// HedgeyABI is the input ABI used to generate the binding from.
// Deprecated: Use HedgeyMetaData.ABI instead.
var HedgeyABI = HedgeyMetaData.ABI

// Hedgey is an auto generated Go binding around an Ethereum contract.
type Hedgey struct {
	HedgeyCaller     // Read-only binding to the contract
	HedgeyTransactor // Write-only binding to the contract
	HedgeyFilterer   // Log filterer for contract events
}

// HedgeyCaller is an auto generated read-only Go binding around an Ethereum contract.
type HedgeyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HedgeyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HedgeyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HedgeyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HedgeyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HedgeySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HedgeySession struct {
	Contract     *Hedgey           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HedgeyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HedgeyCallerSession struct {
	Contract *HedgeyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HedgeyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HedgeyTransactorSession struct {
	Contract     *HedgeyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HedgeyRaw is an auto generated low-level Go binding around an Ethereum contract.
type HedgeyRaw struct {
	Contract *Hedgey // Generic contract binding to access the raw methods on
}

// HedgeyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HedgeyCallerRaw struct {
	Contract *HedgeyCaller // Generic read-only contract binding to access the raw methods on
}

// HedgeyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HedgeyTransactorRaw struct {
	Contract *HedgeyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHedgey creates a new instance of Hedgey, bound to a specific deployed contract.
func NewHedgey(address common.Address, backend bind.ContractBackend) (*Hedgey, error) {
	contract, err := bindHedgey(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hedgey{HedgeyCaller: HedgeyCaller{contract: contract}, HedgeyTransactor: HedgeyTransactor{contract: contract}, HedgeyFilterer: HedgeyFilterer{contract: contract}}, nil
}

// NewHedgeyCaller creates a new read-only instance of Hedgey, bound to a specific deployed contract.
func NewHedgeyCaller(address common.Address, caller bind.ContractCaller) (*HedgeyCaller, error) {
	contract, err := bindHedgey(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HedgeyCaller{contract: contract}, nil
}

// NewHedgeyTransactor creates a new write-only instance of Hedgey, bound to a specific deployed contract.
func NewHedgeyTransactor(address common.Address, transactor bind.ContractTransactor) (*HedgeyTransactor, error) {
	contract, err := bindHedgey(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HedgeyTransactor{contract: contract}, nil
}

// NewHedgeyFilterer creates a new log filterer instance of Hedgey, bound to a specific deployed contract.
func NewHedgeyFilterer(address common.Address, filterer bind.ContractFilterer) (*HedgeyFilterer, error) {
	contract, err := bindHedgey(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HedgeyFilterer{contract: contract}, nil
}

// bindHedgey binds a generic wrapper to an already deployed contract.
func bindHedgey(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HedgeyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hedgey *HedgeyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hedgey.Contract.HedgeyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hedgey *HedgeyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hedgey.Contract.HedgeyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hedgey *HedgeyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hedgey.Contract.HedgeyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hedgey *HedgeyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hedgey.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hedgey *HedgeyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hedgey.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hedgey *HedgeyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hedgey.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hedgey *HedgeyCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hedgey *HedgeySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Hedgey.Contract.BalanceOf(&_Hedgey.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Hedgey *HedgeyCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Hedgey.Contract.BalanceOf(&_Hedgey.CallOpts, owner)
}

// Futures is a free data retrieval call binding the contract method 0x8a1a110d.
//
// Solidity: function futures(uint256 ) view returns(uint256 amount, address token, uint256 unlockDate)
func (_Hedgey *HedgeyCaller) Futures(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount     *big.Int
	Token      common.Address
	UnlockDate *big.Int
}, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "futures", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		Token      common.Address
		UnlockDate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.UnlockDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Futures is a free data retrieval call binding the contract method 0x8a1a110d.
//
// Solidity: function futures(uint256 ) view returns(uint256 amount, address token, uint256 unlockDate)
func (_Hedgey *HedgeySession) Futures(arg0 *big.Int) (struct {
	Amount     *big.Int
	Token      common.Address
	UnlockDate *big.Int
}, error) {
	return _Hedgey.Contract.Futures(&_Hedgey.CallOpts, arg0)
}

// Futures is a free data retrieval call binding the contract method 0x8a1a110d.
//
// Solidity: function futures(uint256 ) view returns(uint256 amount, address token, uint256 unlockDate)
func (_Hedgey *HedgeyCallerSession) Futures(arg0 *big.Int) (struct {
	Amount     *big.Int
	Token      common.Address
	UnlockDate *big.Int
}, error) {
	return _Hedgey.Contract.Futures(&_Hedgey.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hedgey *HedgeyCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hedgey *HedgeySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Hedgey.Contract.GetApproved(&_Hedgey.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Hedgey *HedgeyCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Hedgey.Contract.GetApproved(&_Hedgey.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hedgey *HedgeyCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hedgey *HedgeySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Hedgey.Contract.IsApprovedForAll(&_Hedgey.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Hedgey *HedgeyCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Hedgey.Contract.IsApprovedForAll(&_Hedgey.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hedgey *HedgeyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hedgey *HedgeySession) Name() (string, error) {
	return _Hedgey.Contract.Name(&_Hedgey.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hedgey *HedgeyCallerSession) Name() (string, error) {
	return _Hedgey.Contract.Name(&_Hedgey.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hedgey *HedgeyCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hedgey *HedgeySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Hedgey.Contract.OwnerOf(&_Hedgey.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Hedgey *HedgeyCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Hedgey.Contract.OwnerOf(&_Hedgey.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hedgey *HedgeyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hedgey *HedgeySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hedgey.Contract.SupportsInterface(&_Hedgey.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Hedgey *HedgeyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Hedgey.Contract.SupportsInterface(&_Hedgey.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hedgey *HedgeyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hedgey *HedgeySession) Symbol() (string, error) {
	return _Hedgey.Contract.Symbol(&_Hedgey.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Hedgey *HedgeyCallerSession) Symbol() (string, error) {
	return _Hedgey.Contract.Symbol(&_Hedgey.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Hedgey *HedgeyCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Hedgey *HedgeySession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Hedgey.Contract.TokenByIndex(&_Hedgey.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Hedgey *HedgeyCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Hedgey.Contract.TokenByIndex(&_Hedgey.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Hedgey *HedgeyCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Hedgey *HedgeySession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Hedgey.Contract.TokenOfOwnerByIndex(&_Hedgey.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Hedgey *HedgeyCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Hedgey.Contract.TokenOfOwnerByIndex(&_Hedgey.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hedgey *HedgeyCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hedgey *HedgeySession) TokenURI(tokenId *big.Int) (string, error) {
	return _Hedgey.Contract.TokenURI(&_Hedgey.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Hedgey *HedgeyCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Hedgey.Contract.TokenURI(&_Hedgey.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hedgey *HedgeyCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hedgey *HedgeySession) TotalSupply() (*big.Int, error) {
	return _Hedgey.Contract.TotalSupply(&_Hedgey.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Hedgey *HedgeyCallerSession) TotalSupply() (*big.Int, error) {
	return _Hedgey.Contract.TotalSupply(&_Hedgey.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Hedgey *HedgeyCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hedgey.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Hedgey *HedgeySession) Weth() (common.Address, error) {
	return _Hedgey.Contract.Weth(&_Hedgey.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Hedgey *HedgeyCallerSession) Weth() (common.Address, error) {
	return _Hedgey.Contract.Weth(&_Hedgey.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hedgey *HedgeyTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hedgey.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hedgey *HedgeySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.Approve(&_Hedgey.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Hedgey *HedgeyTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.Approve(&_Hedgey.TransactOpts, to, tokenId)
}

// CreateNFT is a paid mutator transaction binding the contract method 0xb273e653.
//
// Solidity: function createNFT(address _holder, uint256 _amount, address _token, uint256 _unlockDate) returns(uint256)
func (_Hedgey *HedgeyTransactor) CreateNFT(opts *bind.TransactOpts, _holder common.Address, _amount *big.Int, _token common.Address, _unlockDate *big.Int) (*types.Transaction, error) {
	return _Hedgey.contract.Transact(opts, "createNFT", _holder, _amount, _token, _unlockDate)
}

// CreateNFT is a paid mutator transaction binding the contract method 0xb273e653.
//
// Solidity: function createNFT(address _holder, uint256 _amount, address _token, uint256 _unlockDate) returns(uint256)
func (_Hedgey *HedgeySession) CreateNFT(_holder common.Address, _amount *big.Int, _token common.Address, _unlockDate *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.CreateNFT(&_Hedgey.TransactOpts, _holder, _amount, _token, _unlockDate)
}

// CreateNFT is a paid mutator transaction binding the contract method 0xb273e653.
//
// Solidity: function createNFT(address _holder, uint256 _amount, address _token, uint256 _unlockDate) returns(uint256)
func (_Hedgey *HedgeyTransactorSession) CreateNFT(_holder common.Address, _amount *big.Int, _token common.Address, _unlockDate *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.CreateNFT(&_Hedgey.TransactOpts, _holder, _amount, _token, _unlockDate)
}

// RedeemNFT is a paid mutator transaction binding the contract method 0x58dc2cdb.
//
// Solidity: function redeemNFT(uint256 _id) returns(bool)
func (_Hedgey *HedgeyTransactor) RedeemNFT(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Hedgey.contract.Transact(opts, "redeemNFT", _id)
}

// RedeemNFT is a paid mutator transaction binding the contract method 0x58dc2cdb.
//
// Solidity: function redeemNFT(uint256 _id) returns(bool)
func (_Hedgey *HedgeySession) RedeemNFT(_id *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.RedeemNFT(&_Hedgey.TransactOpts, _id)
}

// RedeemNFT is a paid mutator transaction binding the contract method 0x58dc2cdb.
//
// Solidity: function redeemNFT(uint256 _id) returns(bool)
func (_Hedgey *HedgeyTransactorSession) RedeemNFT(_id *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.RedeemNFT(&_Hedgey.TransactOpts, _id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hedgey *HedgeyTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hedgey.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hedgey *HedgeySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.SafeTransferFrom(&_Hedgey.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Hedgey *HedgeyTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.SafeTransferFrom(&_Hedgey.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hedgey *HedgeyTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hedgey.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hedgey *HedgeySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hedgey.Contract.SafeTransferFrom0(&_Hedgey.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Hedgey *HedgeyTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Hedgey.Contract.SafeTransferFrom0(&_Hedgey.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hedgey *HedgeyTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hedgey.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hedgey *HedgeySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hedgey.Contract.SetApprovalForAll(&_Hedgey.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Hedgey *HedgeyTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Hedgey.Contract.SetApprovalForAll(&_Hedgey.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hedgey *HedgeyTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hedgey.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hedgey *HedgeySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.TransferFrom(&_Hedgey.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Hedgey *HedgeyTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Hedgey.Contract.TransferFrom(&_Hedgey.TransactOpts, from, to, tokenId)
}

// UpdateBaseURI is a paid mutator transaction binding the contract method 0x931688cb.
//
// Solidity: function updateBaseURI(string _uri) returns()
func (_Hedgey *HedgeyTransactor) UpdateBaseURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _Hedgey.contract.Transact(opts, "updateBaseURI", _uri)
}

// UpdateBaseURI is a paid mutator transaction binding the contract method 0x931688cb.
//
// Solidity: function updateBaseURI(string _uri) returns()
func (_Hedgey *HedgeySession) UpdateBaseURI(_uri string) (*types.Transaction, error) {
	return _Hedgey.Contract.UpdateBaseURI(&_Hedgey.TransactOpts, _uri)
}

// UpdateBaseURI is a paid mutator transaction binding the contract method 0x931688cb.
//
// Solidity: function updateBaseURI(string _uri) returns()
func (_Hedgey *HedgeyTransactorSession) UpdateBaseURI(_uri string) (*types.Transaction, error) {
	return _Hedgey.Contract.UpdateBaseURI(&_Hedgey.TransactOpts, _uri)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Hedgey *HedgeyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hedgey.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Hedgey *HedgeySession) Receive() (*types.Transaction, error) {
	return _Hedgey.Contract.Receive(&_Hedgey.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Hedgey *HedgeyTransactorSession) Receive() (*types.Transaction, error) {
	return _Hedgey.Contract.Receive(&_Hedgey.TransactOpts)
}

// HedgeyApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Hedgey contract.
type HedgeyApprovalIterator struct {
	Event *HedgeyApproval // Event containing the contract specifics and raw log

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
func (it *HedgeyApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgeyApproval)
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
		it.Event = new(HedgeyApproval)
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
func (it *HedgeyApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgeyApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgeyApproval represents a Approval event raised by the Hedgey contract.
type HedgeyApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hedgey *HedgeyFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*HedgeyApprovalIterator, error) {

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

	logs, sub, err := _Hedgey.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HedgeyApprovalIterator{contract: _Hedgey.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Hedgey *HedgeyFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HedgeyApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Hedgey.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgeyApproval)
				if err := _Hedgey.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Hedgey *HedgeyFilterer) ParseApproval(log types.Log) (*HedgeyApproval, error) {
	event := new(HedgeyApproval)
	if err := _Hedgey.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgeyApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Hedgey contract.
type HedgeyApprovalForAllIterator struct {
	Event *HedgeyApprovalForAll // Event containing the contract specifics and raw log

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
func (it *HedgeyApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgeyApprovalForAll)
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
		it.Event = new(HedgeyApprovalForAll)
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
func (it *HedgeyApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgeyApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgeyApprovalForAll represents a ApprovalForAll event raised by the Hedgey contract.
type HedgeyApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hedgey *HedgeyFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*HedgeyApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Hedgey.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &HedgeyApprovalForAllIterator{contract: _Hedgey.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Hedgey *HedgeyFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *HedgeyApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Hedgey.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgeyApprovalForAll)
				if err := _Hedgey.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Hedgey *HedgeyFilterer) ParseApprovalForAll(log types.Log) (*HedgeyApprovalForAll, error) {
	event := new(HedgeyApprovalForAll)
	if err := _Hedgey.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgeyNFTCreatedIterator is returned from FilterNFTCreated and is used to iterate over the raw logs and unpacked data for NFTCreated events raised by the Hedgey contract.
type HedgeyNFTCreatedIterator struct {
	Event *HedgeyNFTCreated // Event containing the contract specifics and raw log

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
func (it *HedgeyNFTCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgeyNFTCreated)
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
		it.Event = new(HedgeyNFTCreated)
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
func (it *HedgeyNFTCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgeyNFTCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgeyNFTCreated represents a NFTCreated event raised by the Hedgey contract.
type HedgeyNFTCreated struct {
	I          *big.Int
	Holder     common.Address
	Amount     *big.Int
	Token      common.Address
	UnlockDate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTCreated is a free log retrieval operation binding the contract event 0xcd994fa262ab94dc214c94042fdc0563334c749730f5b49663732eda398fe129.
//
// Solidity: event NFTCreated(uint256 _i, address _holder, uint256 _amount, address _token, uint256 _unlockDate)
func (_Hedgey *HedgeyFilterer) FilterNFTCreated(opts *bind.FilterOpts) (*HedgeyNFTCreatedIterator, error) {

	logs, sub, err := _Hedgey.contract.FilterLogs(opts, "NFTCreated")
	if err != nil {
		return nil, err
	}
	return &HedgeyNFTCreatedIterator{contract: _Hedgey.contract, event: "NFTCreated", logs: logs, sub: sub}, nil
}

// WatchNFTCreated is a free log subscription operation binding the contract event 0xcd994fa262ab94dc214c94042fdc0563334c749730f5b49663732eda398fe129.
//
// Solidity: event NFTCreated(uint256 _i, address _holder, uint256 _amount, address _token, uint256 _unlockDate)
func (_Hedgey *HedgeyFilterer) WatchNFTCreated(opts *bind.WatchOpts, sink chan<- *HedgeyNFTCreated) (event.Subscription, error) {

	logs, sub, err := _Hedgey.contract.WatchLogs(opts, "NFTCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgeyNFTCreated)
				if err := _Hedgey.contract.UnpackLog(event, "NFTCreated", log); err != nil {
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

// ParseNFTCreated is a log parse operation binding the contract event 0xcd994fa262ab94dc214c94042fdc0563334c749730f5b49663732eda398fe129.
//
// Solidity: event NFTCreated(uint256 _i, address _holder, uint256 _amount, address _token, uint256 _unlockDate)
func (_Hedgey *HedgeyFilterer) ParseNFTCreated(log types.Log) (*HedgeyNFTCreated, error) {
	event := new(HedgeyNFTCreated)
	if err := _Hedgey.contract.UnpackLog(event, "NFTCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgeyNFTRedeemedIterator is returned from FilterNFTRedeemed and is used to iterate over the raw logs and unpacked data for NFTRedeemed events raised by the Hedgey contract.
type HedgeyNFTRedeemedIterator struct {
	Event *HedgeyNFTRedeemed // Event containing the contract specifics and raw log

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
func (it *HedgeyNFTRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgeyNFTRedeemed)
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
		it.Event = new(HedgeyNFTRedeemed)
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
func (it *HedgeyNFTRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgeyNFTRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgeyNFTRedeemed represents a NFTRedeemed event raised by the Hedgey contract.
type HedgeyNFTRedeemed struct {
	I          *big.Int
	Holder     common.Address
	Amount     *big.Int
	Token      common.Address
	UnlockDate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTRedeemed is a free log retrieval operation binding the contract event 0x59e3d61947ca8bf2d075b2e99dd29f02f9be1be5206aa07df0fff70a8b1ca5b9.
//
// Solidity: event NFTRedeemed(uint256 _i, address _holder, uint256 _amount, address _token, uint256 _unlockDate)
func (_Hedgey *HedgeyFilterer) FilterNFTRedeemed(opts *bind.FilterOpts) (*HedgeyNFTRedeemedIterator, error) {

	logs, sub, err := _Hedgey.contract.FilterLogs(opts, "NFTRedeemed")
	if err != nil {
		return nil, err
	}
	return &HedgeyNFTRedeemedIterator{contract: _Hedgey.contract, event: "NFTRedeemed", logs: logs, sub: sub}, nil
}

// WatchNFTRedeemed is a free log subscription operation binding the contract event 0x59e3d61947ca8bf2d075b2e99dd29f02f9be1be5206aa07df0fff70a8b1ca5b9.
//
// Solidity: event NFTRedeemed(uint256 _i, address _holder, uint256 _amount, address _token, uint256 _unlockDate)
func (_Hedgey *HedgeyFilterer) WatchNFTRedeemed(opts *bind.WatchOpts, sink chan<- *HedgeyNFTRedeemed) (event.Subscription, error) {

	logs, sub, err := _Hedgey.contract.WatchLogs(opts, "NFTRedeemed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgeyNFTRedeemed)
				if err := _Hedgey.contract.UnpackLog(event, "NFTRedeemed", log); err != nil {
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

// ParseNFTRedeemed is a log parse operation binding the contract event 0x59e3d61947ca8bf2d075b2e99dd29f02f9be1be5206aa07df0fff70a8b1ca5b9.
//
// Solidity: event NFTRedeemed(uint256 _i, address _holder, uint256 _amount, address _token, uint256 _unlockDate)
func (_Hedgey *HedgeyFilterer) ParseNFTRedeemed(log types.Log) (*HedgeyNFTRedeemed, error) {
	event := new(HedgeyNFTRedeemed)
	if err := _Hedgey.contract.UnpackLog(event, "NFTRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgeyTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Hedgey contract.
type HedgeyTransferIterator struct {
	Event *HedgeyTransfer // Event containing the contract specifics and raw log

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
func (it *HedgeyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgeyTransfer)
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
		it.Event = new(HedgeyTransfer)
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
func (it *HedgeyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgeyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgeyTransfer represents a Transfer event raised by the Hedgey contract.
type HedgeyTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hedgey *HedgeyFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*HedgeyTransferIterator, error) {

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

	logs, sub, err := _Hedgey.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &HedgeyTransferIterator{contract: _Hedgey.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Hedgey *HedgeyFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HedgeyTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Hedgey.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgeyTransfer)
				if err := _Hedgey.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Hedgey *HedgeyFilterer) ParseTransfer(log types.Log) (*HedgeyTransfer, error) {
	event := new(HedgeyTransfer)
	if err := _Hedgey.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HedgeyURISetIterator is returned from FilterURISet and is used to iterate over the raw logs and unpacked data for URISet events raised by the Hedgey contract.
type HedgeyURISetIterator struct {
	Event *HedgeyURISet // Event containing the contract specifics and raw log

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
func (it *HedgeyURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HedgeyURISet)
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
		it.Event = new(HedgeyURISet)
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
func (it *HedgeyURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HedgeyURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HedgeyURISet represents a URISet event raised by the Hedgey contract.
type HedgeyURISet struct {
	NewURI string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterURISet is a free log retrieval operation binding the contract event 0xde63cc2d19581e57e158d078c2df83f9ab70addd6257f7f12bfecb21c06c9128.
//
// Solidity: event URISet(string newURI)
func (_Hedgey *HedgeyFilterer) FilterURISet(opts *bind.FilterOpts) (*HedgeyURISetIterator, error) {

	logs, sub, err := _Hedgey.contract.FilterLogs(opts, "URISet")
	if err != nil {
		return nil, err
	}
	return &HedgeyURISetIterator{contract: _Hedgey.contract, event: "URISet", logs: logs, sub: sub}, nil
}

// WatchURISet is a free log subscription operation binding the contract event 0xde63cc2d19581e57e158d078c2df83f9ab70addd6257f7f12bfecb21c06c9128.
//
// Solidity: event URISet(string newURI)
func (_Hedgey *HedgeyFilterer) WatchURISet(opts *bind.WatchOpts, sink chan<- *HedgeyURISet) (event.Subscription, error) {

	logs, sub, err := _Hedgey.contract.WatchLogs(opts, "URISet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HedgeyURISet)
				if err := _Hedgey.contract.UnpackLog(event, "URISet", log); err != nil {
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

// ParseURISet is a log parse operation binding the contract event 0xde63cc2d19581e57e158d078c2df83f9ab70addd6257f7f12bfecb21c06c9128.
//
// Solidity: event URISet(string newURI)
func (_Hedgey *HedgeyFilterer) ParseURISet(log types.Log) (*HedgeyURISet, error) {
	event := new(HedgeyURISet)
	if err := _Hedgey.contract.UnpackLog(event, "URISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
