// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingrewards

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

// StakingrewardsMetaData contains all meta data concerning the Stakingrewards contract.
var StakingrewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsDistribution\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"stakeWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingrewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingrewardsMetaData.ABI instead.
var StakingrewardsABI = StakingrewardsMetaData.ABI

// Stakingrewards is an auto generated Go binding around an Ethereum contract.
type Stakingrewards struct {
	StakingrewardsCaller     // Read-only binding to the contract
	StakingrewardsTransactor // Write-only binding to the contract
	StakingrewardsFilterer   // Log filterer for contract events
}

// StakingrewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingrewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingrewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingrewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingrewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingrewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingrewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingrewardsSession struct {
	Contract     *Stakingrewards   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingrewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingrewardsCallerSession struct {
	Contract *StakingrewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingrewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingrewardsTransactorSession struct {
	Contract     *StakingrewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingrewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingrewardsRaw struct {
	Contract *Stakingrewards // Generic contract binding to access the raw methods on
}

// StakingrewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingrewardsCallerRaw struct {
	Contract *StakingrewardsCaller // Generic read-only contract binding to access the raw methods on
}

// StakingrewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingrewardsTransactorRaw struct {
	Contract *StakingrewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingrewards creates a new instance of Stakingrewards, bound to a specific deployed contract.
func NewStakingrewards(address common.Address, backend bind.ContractBackend) (*Stakingrewards, error) {
	contract, err := bindStakingrewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakingrewards{StakingrewardsCaller: StakingrewardsCaller{contract: contract}, StakingrewardsTransactor: StakingrewardsTransactor{contract: contract}, StakingrewardsFilterer: StakingrewardsFilterer{contract: contract}}, nil
}

// NewStakingrewardsCaller creates a new read-only instance of Stakingrewards, bound to a specific deployed contract.
func NewStakingrewardsCaller(address common.Address, caller bind.ContractCaller) (*StakingrewardsCaller, error) {
	contract, err := bindStakingrewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingrewardsCaller{contract: contract}, nil
}

// NewStakingrewardsTransactor creates a new write-only instance of Stakingrewards, bound to a specific deployed contract.
func NewStakingrewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingrewardsTransactor, error) {
	contract, err := bindStakingrewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingrewardsTransactor{contract: contract}, nil
}

// NewStakingrewardsFilterer creates a new log filterer instance of Stakingrewards, bound to a specific deployed contract.
func NewStakingrewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingrewardsFilterer, error) {
	contract, err := bindStakingrewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingrewardsFilterer{contract: contract}, nil
}

// bindStakingrewards binds a generic wrapper to an already deployed contract.
func bindStakingrewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingrewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakingrewards *StakingrewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakingrewards.Contract.StakingrewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakingrewards *StakingrewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrewards.Contract.StakingrewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakingrewards *StakingrewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakingrewards.Contract.StakingrewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakingrewards *StakingrewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakingrewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakingrewards *StakingrewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakingrewards *StakingrewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakingrewards.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.BalanceOf(&_Stakingrewards.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.BalanceOf(&_Stakingrewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) Earned(account common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.Earned(&_Stakingrewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.Earned(&_Stakingrewards.CallOpts, account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) GetRewardForDuration() (*big.Int, error) {
	return _Stakingrewards.Contract.GetRewardForDuration(&_Stakingrewards.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _Stakingrewards.Contract.GetRewardForDuration(&_Stakingrewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Stakingrewards.Contract.LastTimeRewardApplicable(&_Stakingrewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Stakingrewards.Contract.LastTimeRewardApplicable(&_Stakingrewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) LastUpdateTime() (*big.Int, error) {
	return _Stakingrewards.Contract.LastUpdateTime(&_Stakingrewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) LastUpdateTime() (*big.Int, error) {
	return _Stakingrewards.Contract.LastUpdateTime(&_Stakingrewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) PeriodFinish() (*big.Int, error) {
	return _Stakingrewards.Contract.PeriodFinish(&_Stakingrewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) PeriodFinish() (*big.Int, error) {
	return _Stakingrewards.Contract.PeriodFinish(&_Stakingrewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) RewardPerToken() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardPerToken(&_Stakingrewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) RewardPerToken() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardPerToken(&_Stakingrewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) RewardPerTokenStored() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardPerTokenStored(&_Stakingrewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardPerTokenStored(&_Stakingrewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) RewardRate() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardRate(&_Stakingrewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) RewardRate() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardRate(&_Stakingrewards.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.Rewards(&_Stakingrewards.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.Rewards(&_Stakingrewards.CallOpts, arg0)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_Stakingrewards *StakingrewardsCaller) RewardsDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "rewardsDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_Stakingrewards *StakingrewardsSession) RewardsDistribution() (common.Address, error) {
	return _Stakingrewards.Contract.RewardsDistribution(&_Stakingrewards.CallOpts)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_Stakingrewards *StakingrewardsCallerSession) RewardsDistribution() (common.Address, error) {
	return _Stakingrewards.Contract.RewardsDistribution(&_Stakingrewards.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) RewardsDuration() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardsDuration(&_Stakingrewards.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) RewardsDuration() (*big.Int, error) {
	return _Stakingrewards.Contract.RewardsDuration(&_Stakingrewards.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_Stakingrewards *StakingrewardsCaller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_Stakingrewards *StakingrewardsSession) RewardsToken() (common.Address, error) {
	return _Stakingrewards.Contract.RewardsToken(&_Stakingrewards.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_Stakingrewards *StakingrewardsCallerSession) RewardsToken() (common.Address, error) {
	return _Stakingrewards.Contract.RewardsToken(&_Stakingrewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Stakingrewards *StakingrewardsCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Stakingrewards *StakingrewardsSession) StakingToken() (common.Address, error) {
	return _Stakingrewards.Contract.StakingToken(&_Stakingrewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Stakingrewards *StakingrewardsCallerSession) StakingToken() (common.Address, error) {
	return _Stakingrewards.Contract.StakingToken(&_Stakingrewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) TotalSupply() (*big.Int, error) {
	return _Stakingrewards.Contract.TotalSupply(&_Stakingrewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _Stakingrewards.Contract.TotalSupply(&_Stakingrewards.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Stakingrewards *StakingrewardsCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrewards.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Stakingrewards *StakingrewardsSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.UserRewardPerTokenPaid(&_Stakingrewards.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Stakingrewards *StakingrewardsCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Stakingrewards.Contract.UserRewardPerTokenPaid(&_Stakingrewards.CallOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Stakingrewards *StakingrewardsTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Stakingrewards *StakingrewardsSession) Exit() (*types.Transaction, error) {
	return _Stakingrewards.Contract.Exit(&_Stakingrewards.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Stakingrewards *StakingrewardsTransactorSession) Exit() (*types.Transaction, error) {
	return _Stakingrewards.Contract.Exit(&_Stakingrewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Stakingrewards *StakingrewardsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Stakingrewards *StakingrewardsSession) GetReward() (*types.Transaction, error) {
	return _Stakingrewards.Contract.GetReward(&_Stakingrewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Stakingrewards *StakingrewardsTransactorSession) GetReward() (*types.Transaction, error) {
	return _Stakingrewards.Contract.GetReward(&_Stakingrewards.TransactOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Stakingrewards *StakingrewardsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Stakingrewards *StakingrewardsSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.NotifyRewardAmount(&_Stakingrewards.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Stakingrewards *StakingrewardsTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.NotifyRewardAmount(&_Stakingrewards.TransactOpts, reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.Stake(&_Stakingrewards.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.Stake(&_Stakingrewards.TransactOpts, amount)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0xecd9ba82.
//
// Solidity: function stakeWithPermit(uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Stakingrewards *StakingrewardsTransactor) StakeWithPermit(opts *bind.TransactOpts, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "stakeWithPermit", amount, deadline, v, r, s)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0xecd9ba82.
//
// Solidity: function stakeWithPermit(uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Stakingrewards *StakingrewardsSession) StakeWithPermit(amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Stakingrewards.Contract.StakeWithPermit(&_Stakingrewards.TransactOpts, amount, deadline, v, r, s)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0xecd9ba82.
//
// Solidity: function stakeWithPermit(uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Stakingrewards *StakingrewardsTransactorSession) StakeWithPermit(amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Stakingrewards.Contract.StakeWithPermit(&_Stakingrewards.TransactOpts, amount, deadline, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.Withdraw(&_Stakingrewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Stakingrewards *StakingrewardsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Stakingrewards.Contract.Withdraw(&_Stakingrewards.TransactOpts, amount)
}

// StakingrewardsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Stakingrewards contract.
type StakingrewardsRewardAddedIterator struct {
	Event *StakingrewardsRewardAdded // Event containing the contract specifics and raw log

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
func (it *StakingrewardsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrewardsRewardAdded)
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
		it.Event = new(StakingrewardsRewardAdded)
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
func (it *StakingrewardsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrewardsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrewardsRewardAdded represents a RewardAdded event raised by the Stakingrewards contract.
type StakingrewardsRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Stakingrewards *StakingrewardsFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*StakingrewardsRewardAddedIterator, error) {

	logs, sub, err := _Stakingrewards.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &StakingrewardsRewardAddedIterator{contract: _Stakingrewards.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Stakingrewards *StakingrewardsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *StakingrewardsRewardAdded) (event.Subscription, error) {

	logs, sub, err := _Stakingrewards.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrewardsRewardAdded)
				if err := _Stakingrewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Stakingrewards *StakingrewardsFilterer) ParseRewardAdded(log types.Log) (*StakingrewardsRewardAdded, error) {
	event := new(StakingrewardsRewardAdded)
	if err := _Stakingrewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrewardsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the Stakingrewards contract.
type StakingrewardsRewardPaidIterator struct {
	Event *StakingrewardsRewardPaid // Event containing the contract specifics and raw log

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
func (it *StakingrewardsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrewardsRewardPaid)
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
		it.Event = new(StakingrewardsRewardPaid)
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
func (it *StakingrewardsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrewardsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrewardsRewardPaid represents a RewardPaid event raised by the Stakingrewards contract.
type StakingrewardsRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Stakingrewards *StakingrewardsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*StakingrewardsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingrewards.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingrewardsRewardPaidIterator{contract: _Stakingrewards.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Stakingrewards *StakingrewardsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *StakingrewardsRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingrewards.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrewardsRewardPaid)
				if err := _Stakingrewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Stakingrewards *StakingrewardsFilterer) ParseRewardPaid(log types.Log) (*StakingrewardsRewardPaid, error) {
	event := new(StakingrewardsRewardPaid)
	if err := _Stakingrewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrewardsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Stakingrewards contract.
type StakingrewardsStakedIterator struct {
	Event *StakingrewardsStaked // Event containing the contract specifics and raw log

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
func (it *StakingrewardsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrewardsStaked)
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
		it.Event = new(StakingrewardsStaked)
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
func (it *StakingrewardsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrewardsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrewardsStaked represents a Staked event raised by the Stakingrewards contract.
type StakingrewardsStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Stakingrewards *StakingrewardsFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*StakingrewardsStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingrewards.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingrewardsStakedIterator{contract: _Stakingrewards.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Stakingrewards *StakingrewardsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingrewardsStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingrewards.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrewardsStaked)
				if err := _Stakingrewards.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Stakingrewards *StakingrewardsFilterer) ParseStaked(log types.Log) (*StakingrewardsStaked, error) {
	event := new(StakingrewardsStaked)
	if err := _Stakingrewards.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrewardsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Stakingrewards contract.
type StakingrewardsWithdrawnIterator struct {
	Event *StakingrewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingrewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrewardsWithdrawn)
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
		it.Event = new(StakingrewardsWithdrawn)
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
func (it *StakingrewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrewardsWithdrawn represents a Withdrawn event raised by the Stakingrewards contract.
type StakingrewardsWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Stakingrewards *StakingrewardsFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*StakingrewardsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingrewards.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingrewardsWithdrawnIterator{contract: _Stakingrewards.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Stakingrewards *StakingrewardsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingrewardsWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stakingrewards.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrewardsWithdrawn)
				if err := _Stakingrewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Stakingrewards *StakingrewardsFilterer) ParseWithdrawn(log types.Log) (*StakingrewardsWithdrawn, error) {
	event := new(StakingrewardsWithdrawn)
	if err := _Stakingrewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
