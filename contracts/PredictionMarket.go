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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumPredictionMarket.Position\",\"name\":\"position\",\"type\":\"uint8\"}],\"name\":\"BetPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"CloseRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"LockRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"StartRound\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"closeRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intervalSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumPredictionMarket.Position\",\"name\":\"position\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"lockRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBetAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"enumPredictionMarket.Position\",\"name\":\"position\",\"type\":\"uint8\"}],\"name\":\"placeBet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closeTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"lockPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"closePrice\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bullAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bearAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"oracleCalled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contracts *ContractsCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contracts *ContractsSession) CurrentEpoch() (*big.Int, error) {
	return _Contracts.Contract.CurrentEpoch(&_Contracts.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contracts *ContractsCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Contracts.Contract.CurrentEpoch(&_Contracts.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(int256)
func (_Contracts *ContractsCaller) GetLatestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLatestPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(int256)
func (_Contracts *ContractsSession) GetLatestPrice() (*big.Int, error) {
	return _Contracts.Contract.GetLatestPrice(&_Contracts.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(int256)
func (_Contracts *ContractsCallerSession) GetLatestPrice() (*big.Int, error) {
	return _Contracts.Contract.GetLatestPrice(&_Contracts.CallOpts)
}

// IntervalSeconds is a free data retrieval call binding the contract method 0x7d1cd04f.
//
// Solidity: function intervalSeconds() view returns(uint256)
func (_Contracts *ContractsCaller) IntervalSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "intervalSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntervalSeconds is a free data retrieval call binding the contract method 0x7d1cd04f.
//
// Solidity: function intervalSeconds() view returns(uint256)
func (_Contracts *ContractsSession) IntervalSeconds() (*big.Int, error) {
	return _Contracts.Contract.IntervalSeconds(&_Contracts.CallOpts)
}

// IntervalSeconds is a free data retrieval call binding the contract method 0x7d1cd04f.
//
// Solidity: function intervalSeconds() view returns(uint256)
func (_Contracts *ContractsCallerSession) IntervalSeconds() (*big.Int, error) {
	return _Contracts.Contract.IntervalSeconds(&_Contracts.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint256 amount, uint8 position, bool claimed)
func (_Contracts *ContractsCaller) Ledger(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount   *big.Int
	Position uint8
	Claimed  bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ledger", arg0, arg1)

	outstruct := new(struct {
		Amount   *big.Int
		Position uint8
		Claimed  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Position = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Claimed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint256 amount, uint8 position, bool claimed)
func (_Contracts *ContractsSession) Ledger(arg0 *big.Int, arg1 common.Address) (struct {
	Amount   *big.Int
	Position uint8
	Claimed  bool
}, error) {
	return _Contracts.Contract.Ledger(&_Contracts.CallOpts, arg0, arg1)
}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint256 amount, uint8 position, bool claimed)
func (_Contracts *ContractsCallerSession) Ledger(arg0 *big.Int, arg1 common.Address) (struct {
	Amount   *big.Int
	Position uint8
	Claimed  bool
}, error) {
	return _Contracts.Contract.Ledger(&_Contracts.CallOpts, arg0, arg1)
}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_Contracts *ContractsCaller) MinBetAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "minBetAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_Contracts *ContractsSession) MinBetAmount() (*big.Int, error) {
	return _Contracts.Contract.MinBetAmount(&_Contracts.CallOpts)
}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_Contracts *ContractsCallerSession) MinBetAmount() (*big.Int, error) {
	return _Contracts.Contract.MinBetAmount(&_Contracts.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Contracts *ContractsCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Contracts *ContractsSession) Operator() (common.Address, error) {
	return _Contracts.Contract.Operator(&_Contracts.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Contracts *ContractsCallerSession) Operator() (common.Address, error) {
	return _Contracts.Contract.Operator(&_Contracts.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contracts *ContractsCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contracts *ContractsSession) Oracle() (common.Address, error) {
	return _Contracts.Contract.Oracle(&_Contracts.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contracts *ContractsCallerSession) Oracle() (common.Address, error) {
	return _Contracts.Contract.Oracle(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startTimestamp, uint256 lockTimestamp, uint256 closeTimestamp, int256 lockPrice, int256 closePrice, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, bool oracleCalled)
func (_Contracts *ContractsCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Epoch          *big.Int
	StartTimestamp *big.Int
	LockTimestamp  *big.Int
	CloseTimestamp *big.Int
	LockPrice      *big.Int
	ClosePrice     *big.Int
	TotalAmount    *big.Int
	BullAmount     *big.Int
	BearAmount     *big.Int
	OracleCalled   bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "rounds", arg0)

	outstruct := new(struct {
		Epoch          *big.Int
		StartTimestamp *big.Int
		LockTimestamp  *big.Int
		CloseTimestamp *big.Int
		LockPrice      *big.Int
		ClosePrice     *big.Int
		TotalAmount    *big.Int
		BullAmount     *big.Int
		BearAmount     *big.Int
		OracleCalled   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LockTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CloseTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ClosePrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.BullAmount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.BearAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.OracleCalled = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startTimestamp, uint256 lockTimestamp, uint256 closeTimestamp, int256 lockPrice, int256 closePrice, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, bool oracleCalled)
func (_Contracts *ContractsSession) Rounds(arg0 *big.Int) (struct {
	Epoch          *big.Int
	StartTimestamp *big.Int
	LockTimestamp  *big.Int
	CloseTimestamp *big.Int
	LockPrice      *big.Int
	ClosePrice     *big.Int
	TotalAmount    *big.Int
	BullAmount     *big.Int
	BearAmount     *big.Int
	OracleCalled   bool
}, error) {
	return _Contracts.Contract.Rounds(&_Contracts.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startTimestamp, uint256 lockTimestamp, uint256 closeTimestamp, int256 lockPrice, int256 closePrice, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, bool oracleCalled)
func (_Contracts *ContractsCallerSession) Rounds(arg0 *big.Int) (struct {
	Epoch          *big.Int
	StartTimestamp *big.Int
	LockTimestamp  *big.Int
	CloseTimestamp *big.Int
	LockPrice      *big.Int
	ClosePrice     *big.Int
	TotalAmount    *big.Int
	BullAmount     *big.Int
	BearAmount     *big.Int
	OracleCalled   bool
}, error) {
	return _Contracts.Contract.Rounds(&_Contracts.CallOpts, arg0)
}

// TreasuryFee is a free data retrieval call binding the contract method 0xcc32d176.
//
// Solidity: function treasuryFee() view returns(uint256)
func (_Contracts *ContractsCaller) TreasuryFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "treasuryFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryFee is a free data retrieval call binding the contract method 0xcc32d176.
//
// Solidity: function treasuryFee() view returns(uint256)
func (_Contracts *ContractsSession) TreasuryFee() (*big.Int, error) {
	return _Contracts.Contract.TreasuryFee(&_Contracts.CallOpts)
}

// TreasuryFee is a free data retrieval call binding the contract method 0xcc32d176.
//
// Solidity: function treasuryFee() view returns(uint256)
func (_Contracts *ContractsCallerSession) TreasuryFee() (*big.Int, error) {
	return _Contracts.Contract.TreasuryFee(&_Contracts.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 epoch) returns()
func (_Contracts *ContractsTransactor) Claim(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "claim", epoch)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 epoch) returns()
func (_Contracts *ContractsSession) Claim(epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Claim(&_Contracts.TransactOpts, epoch)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 epoch) returns()
func (_Contracts *ContractsTransactorSession) Claim(epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Claim(&_Contracts.TransactOpts, epoch)
}

// CloseRound is a paid mutator transaction binding the contract method 0x88e01a98.
//
// Solidity: function closeRound(uint256 epoch) returns()
func (_Contracts *ContractsTransactor) CloseRound(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "closeRound", epoch)
}

// CloseRound is a paid mutator transaction binding the contract method 0x88e01a98.
//
// Solidity: function closeRound(uint256 epoch) returns()
func (_Contracts *ContractsSession) CloseRound(epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CloseRound(&_Contracts.TransactOpts, epoch)
}

// CloseRound is a paid mutator transaction binding the contract method 0x88e01a98.
//
// Solidity: function closeRound(uint256 epoch) returns()
func (_Contracts *ContractsTransactorSession) CloseRound(epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CloseRound(&_Contracts.TransactOpts, epoch)
}

// LockRound is a paid mutator transaction binding the contract method 0x8ab4ca8a.
//
// Solidity: function lockRound(uint256 epoch) returns()
func (_Contracts *ContractsTransactor) LockRound(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "lockRound", epoch)
}

// LockRound is a paid mutator transaction binding the contract method 0x8ab4ca8a.
//
// Solidity: function lockRound(uint256 epoch) returns()
func (_Contracts *ContractsSession) LockRound(epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.LockRound(&_Contracts.TransactOpts, epoch)
}

// LockRound is a paid mutator transaction binding the contract method 0x8ab4ca8a.
//
// Solidity: function lockRound(uint256 epoch) returns()
func (_Contracts *ContractsTransactorSession) LockRound(epoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.LockRound(&_Contracts.TransactOpts, epoch)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x03edf914.
//
// Solidity: function placeBet(uint256 epoch, uint8 position) payable returns()
func (_Contracts *ContractsTransactor) PlaceBet(opts *bind.TransactOpts, epoch *big.Int, position uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "placeBet", epoch, position)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x03edf914.
//
// Solidity: function placeBet(uint256 epoch, uint8 position) payable returns()
func (_Contracts *ContractsSession) PlaceBet(epoch *big.Int, position uint8) (*types.Transaction, error) {
	return _Contracts.Contract.PlaceBet(&_Contracts.TransactOpts, epoch, position)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x03edf914.
//
// Solidity: function placeBet(uint256 epoch, uint8 position) payable returns()
func (_Contracts *ContractsTransactorSession) PlaceBet(epoch *big.Int, position uint8) (*types.Transaction, error) {
	return _Contracts.Contract.PlaceBet(&_Contracts.TransactOpts, epoch, position)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Contracts *ContractsTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Contracts *ContractsSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetOperator(&_Contracts.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Contracts *ContractsTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetOperator(&_Contracts.TransactOpts, _operator)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Contracts *ContractsTransactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Contracts *ContractsSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetOracle(&_Contracts.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Contracts *ContractsTransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetOracle(&_Contracts.TransactOpts, _oracle)
}

// StartRound is a paid mutator transaction binding the contract method 0x55e3f086.
//
// Solidity: function startRound() returns()
func (_Contracts *ContractsTransactor) StartRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "startRound")
}

// StartRound is a paid mutator transaction binding the contract method 0x55e3f086.
//
// Solidity: function startRound() returns()
func (_Contracts *ContractsSession) StartRound() (*types.Transaction, error) {
	return _Contracts.Contract.StartRound(&_Contracts.TransactOpts)
}

// StartRound is a paid mutator transaction binding the contract method 0x55e3f086.
//
// Solidity: function startRound() returns()
func (_Contracts *ContractsTransactorSession) StartRound() (*types.Transaction, error) {
	return _Contracts.Contract.StartRound(&_Contracts.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactorSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// ContractsBetPlacedIterator is returned from FilterBetPlaced and is used to iterate over the raw logs and unpacked data for BetPlaced events raised by the Contracts contract.
type ContractsBetPlacedIterator struct {
	Event *ContractsBetPlaced // Event containing the contract specifics and raw log

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
func (it *ContractsBetPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBetPlaced)
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
		it.Event = new(ContractsBetPlaced)
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
func (it *ContractsBetPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBetPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBetPlaced represents a BetPlaced event raised by the Contracts contract.
type ContractsBetPlaced struct {
	User     common.Address
	Epoch    *big.Int
	Amount   *big.Int
	Position uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBetPlaced is a free log retrieval operation binding the contract event 0xe8bcec081119392599b60e1206c80b5ce85381e4a7371cc28ec1050e93681189.
//
// Solidity: event BetPlaced(address indexed user, uint256 indexed epoch, uint256 amount, uint8 position)
func (_Contracts *ContractsFilterer) FilterBetPlaced(opts *bind.FilterOpts, user []common.Address, epoch []*big.Int) (*ContractsBetPlacedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "BetPlaced", userRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ContractsBetPlacedIterator{contract: _Contracts.contract, event: "BetPlaced", logs: logs, sub: sub}, nil
}

// WatchBetPlaced is a free log subscription operation binding the contract event 0xe8bcec081119392599b60e1206c80b5ce85381e4a7371cc28ec1050e93681189.
//
// Solidity: event BetPlaced(address indexed user, uint256 indexed epoch, uint256 amount, uint8 position)
func (_Contracts *ContractsFilterer) WatchBetPlaced(opts *bind.WatchOpts, sink chan<- *ContractsBetPlaced, user []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "BetPlaced", userRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBetPlaced)
				if err := _Contracts.contract.UnpackLog(event, "BetPlaced", log); err != nil {
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

// ParseBetPlaced is a log parse operation binding the contract event 0xe8bcec081119392599b60e1206c80b5ce85381e4a7371cc28ec1050e93681189.
//
// Solidity: event BetPlaced(address indexed user, uint256 indexed epoch, uint256 amount, uint8 position)
func (_Contracts *ContractsFilterer) ParseBetPlaced(log types.Log) (*ContractsBetPlaced, error) {
	event := new(ContractsBetPlaced)
	if err := _Contracts.contract.UnpackLog(event, "BetPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Contracts contract.
type ContractsClaimedIterator struct {
	Event *ContractsClaimed // Event containing the contract specifics and raw log

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
func (it *ContractsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsClaimed)
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
		it.Event = new(ContractsClaimed)
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
func (it *ContractsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsClaimed represents a Claimed event raised by the Contracts contract.
type ContractsClaimed struct {
	User   common.Address
	Epoch  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address indexed user, uint256 indexed epoch, uint256 amount)
func (_Contracts *ContractsFilterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address, epoch []*big.Int) (*ContractsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Claimed", userRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ContractsClaimedIterator{contract: _Contracts.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address indexed user, uint256 indexed epoch, uint256 amount)
func (_Contracts *ContractsFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ContractsClaimed, user []common.Address, epoch []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Claimed", userRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsClaimed)
				if err := _Contracts.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address indexed user, uint256 indexed epoch, uint256 amount)
func (_Contracts *ContractsFilterer) ParseClaimed(log types.Log) (*ContractsClaimed, error) {
	event := new(ContractsClaimed)
	if err := _Contracts.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCloseRoundIterator is returned from FilterCloseRound and is used to iterate over the raw logs and unpacked data for CloseRound events raised by the Contracts contract.
type ContractsCloseRoundIterator struct {
	Event *ContractsCloseRound // Event containing the contract specifics and raw log

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
func (it *ContractsCloseRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCloseRound)
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
		it.Event = new(ContractsCloseRound)
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
func (it *ContractsCloseRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCloseRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCloseRound represents a CloseRound event raised by the Contracts contract.
type ContractsCloseRound struct {
	Epoch *big.Int
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCloseRound is a free log retrieval operation binding the contract event 0xce329a017dea018178f0609b4213bdde80d33425d51d24f7090d644ade2e003b.
//
// Solidity: event CloseRound(uint256 indexed epoch, int256 price)
func (_Contracts *ContractsFilterer) FilterCloseRound(opts *bind.FilterOpts, epoch []*big.Int) (*ContractsCloseRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CloseRound", epochRule)
	if err != nil {
		return nil, err
	}
	return &ContractsCloseRoundIterator{contract: _Contracts.contract, event: "CloseRound", logs: logs, sub: sub}, nil
}

// WatchCloseRound is a free log subscription operation binding the contract event 0xce329a017dea018178f0609b4213bdde80d33425d51d24f7090d644ade2e003b.
//
// Solidity: event CloseRound(uint256 indexed epoch, int256 price)
func (_Contracts *ContractsFilterer) WatchCloseRound(opts *bind.WatchOpts, sink chan<- *ContractsCloseRound, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CloseRound", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCloseRound)
				if err := _Contracts.contract.UnpackLog(event, "CloseRound", log); err != nil {
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

// ParseCloseRound is a log parse operation binding the contract event 0xce329a017dea018178f0609b4213bdde80d33425d51d24f7090d644ade2e003b.
//
// Solidity: event CloseRound(uint256 indexed epoch, int256 price)
func (_Contracts *ContractsFilterer) ParseCloseRound(log types.Log) (*ContractsCloseRound, error) {
	event := new(ContractsCloseRound)
	if err := _Contracts.contract.UnpackLog(event, "CloseRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLockRoundIterator is returned from FilterLockRound and is used to iterate over the raw logs and unpacked data for LockRound events raised by the Contracts contract.
type ContractsLockRoundIterator struct {
	Event *ContractsLockRound // Event containing the contract specifics and raw log

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
func (it *ContractsLockRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLockRound)
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
		it.Event = new(ContractsLockRound)
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
func (it *ContractsLockRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLockRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLockRound represents a LockRound event raised by the Contracts contract.
type ContractsLockRound struct {
	Epoch *big.Int
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLockRound is a free log retrieval operation binding the contract event 0xeef94cbda40aa01b9b02b7dbd6075caa3662c0fabf77a84ed674b09b0ff9cb1e.
//
// Solidity: event LockRound(uint256 indexed epoch, int256 price)
func (_Contracts *ContractsFilterer) FilterLockRound(opts *bind.FilterOpts, epoch []*big.Int) (*ContractsLockRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LockRound", epochRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLockRoundIterator{contract: _Contracts.contract, event: "LockRound", logs: logs, sub: sub}, nil
}

// WatchLockRound is a free log subscription operation binding the contract event 0xeef94cbda40aa01b9b02b7dbd6075caa3662c0fabf77a84ed674b09b0ff9cb1e.
//
// Solidity: event LockRound(uint256 indexed epoch, int256 price)
func (_Contracts *ContractsFilterer) WatchLockRound(opts *bind.WatchOpts, sink chan<- *ContractsLockRound, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LockRound", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLockRound)
				if err := _Contracts.contract.UnpackLog(event, "LockRound", log); err != nil {
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

// ParseLockRound is a log parse operation binding the contract event 0xeef94cbda40aa01b9b02b7dbd6075caa3662c0fabf77a84ed674b09b0ff9cb1e.
//
// Solidity: event LockRound(uint256 indexed epoch, int256 price)
func (_Contracts *ContractsFilterer) ParseLockRound(log types.Log) (*ContractsLockRound, error) {
	event := new(ContractsLockRound)
	if err := _Contracts.contract.UnpackLog(event, "LockRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsStartRoundIterator is returned from FilterStartRound and is used to iterate over the raw logs and unpacked data for StartRound events raised by the Contracts contract.
type ContractsStartRoundIterator struct {
	Event *ContractsStartRound // Event containing the contract specifics and raw log

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
func (it *ContractsStartRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsStartRound)
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
		it.Event = new(ContractsStartRound)
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
func (it *ContractsStartRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsStartRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsStartRound represents a StartRound event raised by the Contracts contract.
type ContractsStartRound struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStartRound is a free log retrieval operation binding the contract event 0x939f42374aa9bf1d8d8cd56d8a9110cb040cd8dfeae44080c6fcf2645e51b452.
//
// Solidity: event StartRound(uint256 indexed epoch)
func (_Contracts *ContractsFilterer) FilterStartRound(opts *bind.FilterOpts, epoch []*big.Int) (*ContractsStartRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "StartRound", epochRule)
	if err != nil {
		return nil, err
	}
	return &ContractsStartRoundIterator{contract: _Contracts.contract, event: "StartRound", logs: logs, sub: sub}, nil
}

// WatchStartRound is a free log subscription operation binding the contract event 0x939f42374aa9bf1d8d8cd56d8a9110cb040cd8dfeae44080c6fcf2645e51b452.
//
// Solidity: event StartRound(uint256 indexed epoch)
func (_Contracts *ContractsFilterer) WatchStartRound(opts *bind.WatchOpts, sink chan<- *ContractsStartRound, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "StartRound", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsStartRound)
				if err := _Contracts.contract.UnpackLog(event, "StartRound", log); err != nil {
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

// ParseStartRound is a log parse operation binding the contract event 0x939f42374aa9bf1d8d8cd56d8a9110cb040cd8dfeae44080c6fcf2645e51b452.
//
// Solidity: event StartRound(uint256 indexed epoch)
func (_Contracts *ContractsFilterer) ParseStartRound(log types.Log) (*ContractsStartRound, error) {
	event := new(ContractsStartRound)
	if err := _Contracts.contract.UnpackLog(event, "StartRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
