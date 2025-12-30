// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// SybilMetaData contains all meta data concerning the Sybil contract.
var SybilMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_s\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_t\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_batchSize\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"accountIdx\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelLink\",\"inputs\":[{\"name\":\"counterparty\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelVouch\",\"inputs\":[{\"name\":\"counterparty\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeWithoutSteal\",\"inputs\":[{\"name\":\"counterparty\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"finalize\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"b\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasLink\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"b\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFinalizeReady\",\"inputs\":[{\"name\":\"x\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"y\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLinked\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastForgedId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestBatchId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestGraphRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextEdgeId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextIdx\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"windowStart\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"loFunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"hiFunded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"stakeAmt\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingEdges\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requiredStake\",\"inputs\":[{\"name\":\"x\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"y\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scoreRootAt\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scoreSnapshot\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"score\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scoreSnapshotOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"score\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batch\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBatchSize\",\"inputs\":[{\"name\":\"_n\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setParams\",\"inputs\":[{\"name\":\"_s\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_t\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"steal\",\"inputs\":[{\"name\":\"counterparty\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitBatch\",\"inputs\":[{\"name\":\"newGraph\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newScore\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"n\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalAccounts\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unforged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint72\",\"internalType\":\"uint72\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vouch\",\"inputs\":[{\"name\":\"subject\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"windowT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AccountCreated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"idx\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchSizeUpdated\",\"inputs\":[{\"name\":\"batchSize\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchSubmitted\",\"inputs\":[{\"name\":\"batchId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"count\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"storageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"newGraphRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"newScoreRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"edgesPacked\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClosedNoLink\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"counterparty\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LinkCancelled\",\"inputs\":[{\"name\":\"lo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hi\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Linked\",\"inputs\":[{\"name\":\"lo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hi\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"windowStart\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"windowEnd\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParamsUpdated\",\"inputs\":[{\"name\":\"stakeS\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"windowT\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ScoreSynced\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"batchId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"score\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Stolen\",\"inputs\":[{\"name\":\"thief\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"victim\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payout\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Vouched\",\"inputs\":[{\"name\":\"attester\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"subject\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WindowOpened\",\"inputs\":[{\"name\":\"lo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hi\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"start\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"end\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyHi\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyLinked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyLo\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Early\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EthXferFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MissingIdx\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoWindow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotBothFunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotHiOnly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotLinked\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotLoOnly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PastWindow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Self\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerifyFail\",\"inputs\":[]}]",
}

// SybilABI is the input ABI used to generate the binding from.
// Deprecated: Use SybilMetaData.ABI instead.
var SybilABI = SybilMetaData.ABI

// Sybil is an auto generated Go binding around an Ethereum contract.
type Sybil struct {
	SybilCaller     // Read-only binding to the contract
	SybilTransactor // Write-only binding to the contract
	SybilFilterer   // Log filterer for contract events
}

// SybilCaller is an auto generated read-only Go binding around an Ethereum contract.
type SybilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SybilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SybilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SybilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SybilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SybilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SybilSession struct {
	Contract     *Sybil            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SybilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SybilCallerSession struct {
	Contract *SybilCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SybilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SybilTransactorSession struct {
	Contract     *SybilTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SybilRaw is an auto generated low-level Go binding around an Ethereum contract.
type SybilRaw struct {
	Contract *Sybil // Generic contract binding to access the raw methods on
}

// SybilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SybilCallerRaw struct {
	Contract *SybilCaller // Generic read-only contract binding to access the raw methods on
}

// SybilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SybilTransactorRaw struct {
	Contract *SybilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSybil creates a new instance of Sybil, bound to a specific deployed contract.
func NewSybil(address common.Address, backend bind.ContractBackend) (*Sybil, error) {
	contract, err := bindSybil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sybil{SybilCaller: SybilCaller{contract: contract}, SybilTransactor: SybilTransactor{contract: contract}, SybilFilterer: SybilFilterer{contract: contract}}, nil
}

// NewSybilCaller creates a new read-only instance of Sybil, bound to a specific deployed contract.
func NewSybilCaller(address common.Address, caller bind.ContractCaller) (*SybilCaller, error) {
	contract, err := bindSybil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SybilCaller{contract: contract}, nil
}

// NewSybilTransactor creates a new write-only instance of Sybil, bound to a specific deployed contract.
func NewSybilTransactor(address common.Address, transactor bind.ContractTransactor) (*SybilTransactor, error) {
	contract, err := bindSybil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SybilTransactor{contract: contract}, nil
}

// NewSybilFilterer creates a new log filterer instance of Sybil, bound to a specific deployed contract.
func NewSybilFilterer(address common.Address, filterer bind.ContractFilterer) (*SybilFilterer, error) {
	contract, err := bindSybil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SybilFilterer{contract: contract}, nil
}

// bindSybil binds a generic wrapper to an already deployed contract.
func bindSybil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SybilMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sybil *SybilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sybil.Contract.SybilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sybil *SybilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.Contract.SybilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sybil *SybilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sybil.Contract.SybilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sybil *SybilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sybil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sybil *SybilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sybil *SybilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sybil.Contract.contract.Transact(opts, method, params...)
}

// AccountIdx is a free data retrieval call binding the contract method 0x9ca71f18.
//
// Solidity: function accountIdx(address ) view returns(uint32)
func (_Sybil *SybilCaller) AccountIdx(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "accountIdx", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AccountIdx is a free data retrieval call binding the contract method 0x9ca71f18.
//
// Solidity: function accountIdx(address ) view returns(uint32)
func (_Sybil *SybilSession) AccountIdx(arg0 common.Address) (uint32, error) {
	return _Sybil.Contract.AccountIdx(&_Sybil.CallOpts, arg0)
}

// AccountIdx is a free data retrieval call binding the contract method 0x9ca71f18.
//
// Solidity: function accountIdx(address ) view returns(uint32)
func (_Sybil *SybilCallerSession) AccountIdx(arg0 common.Address) (uint32, error) {
	return _Sybil.Contract.AccountIdx(&_Sybil.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Sybil *SybilCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Sybil *SybilSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Sybil.Contract.BalanceOf(&_Sybil.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Sybil *SybilCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Sybil.Contract.BalanceOf(&_Sybil.CallOpts, arg0)
}

// BatchId is a free data retrieval call binding the contract method 0x4972134a.
//
// Solidity: function batchId() view returns(uint64)
func (_Sybil *SybilCaller) BatchId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "batchId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BatchId is a free data retrieval call binding the contract method 0x4972134a.
//
// Solidity: function batchId() view returns(uint64)
func (_Sybil *SybilSession) BatchId() (uint64, error) {
	return _Sybil.Contract.BatchId(&_Sybil.CallOpts)
}

// BatchId is a free data retrieval call binding the contract method 0x4972134a.
//
// Solidity: function batchId() view returns(uint64)
func (_Sybil *SybilCallerSession) BatchId() (uint64, error) {
	return _Sybil.Contract.BatchId(&_Sybil.CallOpts)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint32)
func (_Sybil *SybilCaller) BatchSize(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "batchSize")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint32)
func (_Sybil *SybilSession) BatchSize() (uint32, error) {
	return _Sybil.Contract.BatchSize(&_Sybil.CallOpts)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint32)
func (_Sybil *SybilCallerSession) BatchSize() (uint32, error) {
	return _Sybil.Contract.BatchSize(&_Sybil.CallOpts)
}

// HasLink is a free data retrieval call binding the contract method 0x636453ed.
//
// Solidity: function hasLink(address a, address b) view returns(bool)
func (_Sybil *SybilCaller) HasLink(opts *bind.CallOpts, a common.Address, b common.Address) (bool, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "hasLink", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasLink is a free data retrieval call binding the contract method 0x636453ed.
//
// Solidity: function hasLink(address a, address b) view returns(bool)
func (_Sybil *SybilSession) HasLink(a common.Address, b common.Address) (bool, error) {
	return _Sybil.Contract.HasLink(&_Sybil.CallOpts, a, b)
}

// HasLink is a free data retrieval call binding the contract method 0x636453ed.
//
// Solidity: function hasLink(address a, address b) view returns(bool)
func (_Sybil *SybilCallerSession) HasLink(a common.Address, b common.Address) (bool, error) {
	return _Sybil.Contract.HasLink(&_Sybil.CallOpts, a, b)
}

// IsFinalizeReady is a free data retrieval call binding the contract method 0xa37d6ac8.
//
// Solidity: function isFinalizeReady(address x, address y) view returns(bool)
func (_Sybil *SybilCaller) IsFinalizeReady(opts *bind.CallOpts, x common.Address, y common.Address) (bool, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "isFinalizeReady", x, y)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalizeReady is a free data retrieval call binding the contract method 0xa37d6ac8.
//
// Solidity: function isFinalizeReady(address x, address y) view returns(bool)
func (_Sybil *SybilSession) IsFinalizeReady(x common.Address, y common.Address) (bool, error) {
	return _Sybil.Contract.IsFinalizeReady(&_Sybil.CallOpts, x, y)
}

// IsFinalizeReady is a free data retrieval call binding the contract method 0xa37d6ac8.
//
// Solidity: function isFinalizeReady(address x, address y) view returns(bool)
func (_Sybil *SybilCallerSession) IsFinalizeReady(x common.Address, y common.Address) (bool, error) {
	return _Sybil.Contract.IsFinalizeReady(&_Sybil.CallOpts, x, y)
}

// IsLinked is a free data retrieval call binding the contract method 0x047f8183.
//
// Solidity: function isLinked(address , address ) view returns(bool)
func (_Sybil *SybilCaller) IsLinked(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "isLinked", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLinked is a free data retrieval call binding the contract method 0x047f8183.
//
// Solidity: function isLinked(address , address ) view returns(bool)
func (_Sybil *SybilSession) IsLinked(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Sybil.Contract.IsLinked(&_Sybil.CallOpts, arg0, arg1)
}

// IsLinked is a free data retrieval call binding the contract method 0x047f8183.
//
// Solidity: function isLinked(address , address ) view returns(bool)
func (_Sybil *SybilCallerSession) IsLinked(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Sybil.Contract.IsLinked(&_Sybil.CallOpts, arg0, arg1)
}

// LastForgedId is a free data retrieval call binding the contract method 0x23b250b6.
//
// Solidity: function lastForgedId() view returns(uint32)
func (_Sybil *SybilCaller) LastForgedId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "lastForgedId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastForgedId is a free data retrieval call binding the contract method 0x23b250b6.
//
// Solidity: function lastForgedId() view returns(uint32)
func (_Sybil *SybilSession) LastForgedId() (uint32, error) {
	return _Sybil.Contract.LastForgedId(&_Sybil.CallOpts)
}

// LastForgedId is a free data retrieval call binding the contract method 0x23b250b6.
//
// Solidity: function lastForgedId() view returns(uint32)
func (_Sybil *SybilCallerSession) LastForgedId() (uint32, error) {
	return _Sybil.Contract.LastForgedId(&_Sybil.CallOpts)
}

// LatestBatchId is a free data retrieval call binding the contract method 0x4d4e7019.
//
// Solidity: function latestBatchId() view returns(uint64)
func (_Sybil *SybilCaller) LatestBatchId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "latestBatchId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestBatchId is a free data retrieval call binding the contract method 0x4d4e7019.
//
// Solidity: function latestBatchId() view returns(uint64)
func (_Sybil *SybilSession) LatestBatchId() (uint64, error) {
	return _Sybil.Contract.LatestBatchId(&_Sybil.CallOpts)
}

// LatestBatchId is a free data retrieval call binding the contract method 0x4d4e7019.
//
// Solidity: function latestBatchId() view returns(uint64)
func (_Sybil *SybilCallerSession) LatestBatchId() (uint64, error) {
	return _Sybil.Contract.LatestBatchId(&_Sybil.CallOpts)
}

// LatestGraphRoot is a free data retrieval call binding the contract method 0xca660d92.
//
// Solidity: function latestGraphRoot() view returns(bytes32)
func (_Sybil *SybilCaller) LatestGraphRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "latestGraphRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestGraphRoot is a free data retrieval call binding the contract method 0xca660d92.
//
// Solidity: function latestGraphRoot() view returns(bytes32)
func (_Sybil *SybilSession) LatestGraphRoot() ([32]byte, error) {
	return _Sybil.Contract.LatestGraphRoot(&_Sybil.CallOpts)
}

// LatestGraphRoot is a free data retrieval call binding the contract method 0xca660d92.
//
// Solidity: function latestGraphRoot() view returns(bytes32)
func (_Sybil *SybilCallerSession) LatestGraphRoot() ([32]byte, error) {
	return _Sybil.Contract.LatestGraphRoot(&_Sybil.CallOpts)
}

// NextEdgeId is a free data retrieval call binding the contract method 0xad0449eb.
//
// Solidity: function nextEdgeId() view returns(uint32)
func (_Sybil *SybilCaller) NextEdgeId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "nextEdgeId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextEdgeId is a free data retrieval call binding the contract method 0xad0449eb.
//
// Solidity: function nextEdgeId() view returns(uint32)
func (_Sybil *SybilSession) NextEdgeId() (uint32, error) {
	return _Sybil.Contract.NextEdgeId(&_Sybil.CallOpts)
}

// NextEdgeId is a free data retrieval call binding the contract method 0xad0449eb.
//
// Solidity: function nextEdgeId() view returns(uint32)
func (_Sybil *SybilCallerSession) NextEdgeId() (uint32, error) {
	return _Sybil.Contract.NextEdgeId(&_Sybil.CallOpts)
}

// NextIdx is a free data retrieval call binding the contract method 0x6f403b5c.
//
// Solidity: function nextIdx() view returns(uint32)
func (_Sybil *SybilCaller) NextIdx(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "nextIdx")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextIdx is a free data retrieval call binding the contract method 0x6f403b5c.
//
// Solidity: function nextIdx() view returns(uint32)
func (_Sybil *SybilSession) NextIdx() (uint32, error) {
	return _Sybil.Contract.NextIdx(&_Sybil.CallOpts)
}

// NextIdx is a free data retrieval call binding the contract method 0x6f403b5c.
//
// Solidity: function nextIdx() view returns(uint32)
func (_Sybil *SybilCallerSession) NextIdx() (uint32, error) {
	return _Sybil.Contract.NextIdx(&_Sybil.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sybil *SybilCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sybil *SybilSession) Owner() (common.Address, error) {
	return _Sybil.Contract.Owner(&_Sybil.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sybil *SybilCallerSession) Owner() (common.Address, error) {
	return _Sybil.Contract.Owner(&_Sybil.CallOpts)
}

// Pairs is a free data retrieval call binding the contract method 0x69454b86.
//
// Solidity: function pairs(address , address ) view returns(uint64 windowStart, bool loFunded, bool hiFunded, uint128 stakeAmt)
func (_Sybil *SybilCaller) Pairs(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	WindowStart uint64
	LoFunded    bool
	HiFunded    bool
	StakeAmt    *big.Int
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "pairs", arg0, arg1)

	outstruct := new(struct {
		WindowStart uint64
		LoFunded    bool
		HiFunded    bool
		StakeAmt    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WindowStart = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LoFunded = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.HiFunded = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.StakeAmt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pairs is a free data retrieval call binding the contract method 0x69454b86.
//
// Solidity: function pairs(address , address ) view returns(uint64 windowStart, bool loFunded, bool hiFunded, uint128 stakeAmt)
func (_Sybil *SybilSession) Pairs(arg0 common.Address, arg1 common.Address) (struct {
	WindowStart uint64
	LoFunded    bool
	HiFunded    bool
	StakeAmt    *big.Int
}, error) {
	return _Sybil.Contract.Pairs(&_Sybil.CallOpts, arg0, arg1)
}

// Pairs is a free data retrieval call binding the contract method 0x69454b86.
//
// Solidity: function pairs(address , address ) view returns(uint64 windowStart, bool loFunded, bool hiFunded, uint128 stakeAmt)
func (_Sybil *SybilCallerSession) Pairs(arg0 common.Address, arg1 common.Address) (struct {
	WindowStart uint64
	LoFunded    bool
	HiFunded    bool
	StakeAmt    *big.Int
}, error) {
	return _Sybil.Contract.Pairs(&_Sybil.CallOpts, arg0, arg1)
}

// PendingEdges is a free data retrieval call binding the contract method 0x9cf05d3f.
//
// Solidity: function pendingEdges() view returns(uint32)
func (_Sybil *SybilCaller) PendingEdges(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "pendingEdges")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PendingEdges is a free data retrieval call binding the contract method 0x9cf05d3f.
//
// Solidity: function pendingEdges() view returns(uint32)
func (_Sybil *SybilSession) PendingEdges() (uint32, error) {
	return _Sybil.Contract.PendingEdges(&_Sybil.CallOpts)
}

// PendingEdges is a free data retrieval call binding the contract method 0x9cf05d3f.
//
// Solidity: function pendingEdges() view returns(uint32)
func (_Sybil *SybilCallerSession) PendingEdges() (uint32, error) {
	return _Sybil.Contract.PendingEdges(&_Sybil.CallOpts)
}

// RequiredStake is a free data retrieval call binding the contract method 0xb4b432e0.
//
// Solidity: function requiredStake(address x, address y) view returns(uint256)
func (_Sybil *SybilCaller) RequiredStake(opts *bind.CallOpts, x common.Address, y common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "requiredStake", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredStake is a free data retrieval call binding the contract method 0xb4b432e0.
//
// Solidity: function requiredStake(address x, address y) view returns(uint256)
func (_Sybil *SybilSession) RequiredStake(x common.Address, y common.Address) (*big.Int, error) {
	return _Sybil.Contract.RequiredStake(&_Sybil.CallOpts, x, y)
}

// RequiredStake is a free data retrieval call binding the contract method 0xb4b432e0.
//
// Solidity: function requiredStake(address x, address y) view returns(uint256)
func (_Sybil *SybilCallerSession) RequiredStake(x common.Address, y common.Address) (*big.Int, error) {
	return _Sybil.Contract.RequiredStake(&_Sybil.CallOpts, x, y)
}

// ScoreRootAt is a free data retrieval call binding the contract method 0x53d38a0e.
//
// Solidity: function scoreRootAt(uint64 ) view returns(bytes32)
func (_Sybil *SybilCaller) ScoreRootAt(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "scoreRootAt", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ScoreRootAt is a free data retrieval call binding the contract method 0x53d38a0e.
//
// Solidity: function scoreRootAt(uint64 ) view returns(bytes32)
func (_Sybil *SybilSession) ScoreRootAt(arg0 uint64) ([32]byte, error) {
	return _Sybil.Contract.ScoreRootAt(&_Sybil.CallOpts, arg0)
}

// ScoreRootAt is a free data retrieval call binding the contract method 0x53d38a0e.
//
// Solidity: function scoreRootAt(uint64 ) view returns(bytes32)
func (_Sybil *SybilCallerSession) ScoreRootAt(arg0 uint64) ([32]byte, error) {
	return _Sybil.Contract.ScoreRootAt(&_Sybil.CallOpts, arg0)
}

// ScoreSnapshot is a free data retrieval call binding the contract method 0x3222c8fe.
//
// Solidity: function scoreSnapshot(address ) view returns(uint256 score, uint64 batchId)
func (_Sybil *SybilCaller) ScoreSnapshot(opts *bind.CallOpts, arg0 common.Address) (struct {
	Score   *big.Int
	BatchId uint64
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "scoreSnapshot", arg0)

	outstruct := new(struct {
		Score   *big.Int
		BatchId uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Score = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BatchId = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// ScoreSnapshot is a free data retrieval call binding the contract method 0x3222c8fe.
//
// Solidity: function scoreSnapshot(address ) view returns(uint256 score, uint64 batchId)
func (_Sybil *SybilSession) ScoreSnapshot(arg0 common.Address) (struct {
	Score   *big.Int
	BatchId uint64
}, error) {
	return _Sybil.Contract.ScoreSnapshot(&_Sybil.CallOpts, arg0)
}

// ScoreSnapshot is a free data retrieval call binding the contract method 0x3222c8fe.
//
// Solidity: function scoreSnapshot(address ) view returns(uint256 score, uint64 batchId)
func (_Sybil *SybilCallerSession) ScoreSnapshot(arg0 common.Address) (struct {
	Score   *big.Int
	BatchId uint64
}, error) {
	return _Sybil.Contract.ScoreSnapshot(&_Sybil.CallOpts, arg0)
}

// ScoreSnapshotOf is a free data retrieval call binding the contract method 0x620f2b36.
//
// Solidity: function scoreSnapshotOf(address user) view returns(uint256 score, uint64 batch)
func (_Sybil *SybilCaller) ScoreSnapshotOf(opts *bind.CallOpts, user common.Address) (struct {
	Score *big.Int
	Batch uint64
}, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "scoreSnapshotOf", user)

	outstruct := new(struct {
		Score *big.Int
		Batch uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Score = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Batch = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// ScoreSnapshotOf is a free data retrieval call binding the contract method 0x620f2b36.
//
// Solidity: function scoreSnapshotOf(address user) view returns(uint256 score, uint64 batch)
func (_Sybil *SybilSession) ScoreSnapshotOf(user common.Address) (struct {
	Score *big.Int
	Batch uint64
}, error) {
	return _Sybil.Contract.ScoreSnapshotOf(&_Sybil.CallOpts, user)
}

// ScoreSnapshotOf is a free data retrieval call binding the contract method 0x620f2b36.
//
// Solidity: function scoreSnapshotOf(address user) view returns(uint256 score, uint64 batch)
func (_Sybil *SybilCallerSession) ScoreSnapshotOf(user common.Address) (struct {
	Score *big.Int
	Batch uint64
}, error) {
	return _Sybil.Contract.ScoreSnapshotOf(&_Sybil.CallOpts, user)
}

// StakeS is a free data retrieval call binding the contract method 0x6483879b.
//
// Solidity: function stakeS() view returns(uint256)
func (_Sybil *SybilCaller) StakeS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "stakeS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeS is a free data retrieval call binding the contract method 0x6483879b.
//
// Solidity: function stakeS() view returns(uint256)
func (_Sybil *SybilSession) StakeS() (*big.Int, error) {
	return _Sybil.Contract.StakeS(&_Sybil.CallOpts)
}

// StakeS is a free data retrieval call binding the contract method 0x6483879b.
//
// Solidity: function stakeS() view returns(uint256)
func (_Sybil *SybilCallerSession) StakeS() (*big.Int, error) {
	return _Sybil.Contract.StakeS(&_Sybil.CallOpts)
}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint32)
func (_Sybil *SybilCaller) TotalAccounts(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "totalAccounts")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint32)
func (_Sybil *SybilSession) TotalAccounts() (uint32, error) {
	return _Sybil.Contract.TotalAccounts(&_Sybil.CallOpts)
}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint32)
func (_Sybil *SybilCallerSession) TotalAccounts() (uint32, error) {
	return _Sybil.Contract.TotalAccounts(&_Sybil.CallOpts)
}

// Unforged is a free data retrieval call binding the contract method 0xaa8d6c4f.
//
// Solidity: function unforged(uint32 ) view returns(uint72)
func (_Sybil *SybilCaller) Unforged(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "unforged", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Unforged is a free data retrieval call binding the contract method 0xaa8d6c4f.
//
// Solidity: function unforged(uint32 ) view returns(uint72)
func (_Sybil *SybilSession) Unforged(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.Unforged(&_Sybil.CallOpts, arg0)
}

// Unforged is a free data retrieval call binding the contract method 0xaa8d6c4f.
//
// Solidity: function unforged(uint32 ) view returns(uint72)
func (_Sybil *SybilCallerSession) Unforged(arg0 uint32) (*big.Int, error) {
	return _Sybil.Contract.Unforged(&_Sybil.CallOpts, arg0)
}

// WindowT is a free data retrieval call binding the contract method 0x19411573.
//
// Solidity: function windowT() view returns(uint64)
func (_Sybil *SybilCaller) WindowT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Sybil.contract.Call(opts, &out, "windowT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WindowT is a free data retrieval call binding the contract method 0x19411573.
//
// Solidity: function windowT() view returns(uint64)
func (_Sybil *SybilSession) WindowT() (uint64, error) {
	return _Sybil.Contract.WindowT(&_Sybil.CallOpts)
}

// WindowT is a free data retrieval call binding the contract method 0x19411573.
//
// Solidity: function windowT() view returns(uint64)
func (_Sybil *SybilCallerSession) WindowT() (uint64, error) {
	return _Sybil.Contract.WindowT(&_Sybil.CallOpts)
}

// CancelLink is a paid mutator transaction binding the contract method 0xe21f2e29.
//
// Solidity: function cancelLink(address counterparty) returns()
func (_Sybil *SybilTransactor) CancelLink(opts *bind.TransactOpts, counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "cancelLink", counterparty)
}

// CancelLink is a paid mutator transaction binding the contract method 0xe21f2e29.
//
// Solidity: function cancelLink(address counterparty) returns()
func (_Sybil *SybilSession) CancelLink(counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.CancelLink(&_Sybil.TransactOpts, counterparty)
}

// CancelLink is a paid mutator transaction binding the contract method 0xe21f2e29.
//
// Solidity: function cancelLink(address counterparty) returns()
func (_Sybil *SybilTransactorSession) CancelLink(counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.CancelLink(&_Sybil.TransactOpts, counterparty)
}

// CancelVouch is a paid mutator transaction binding the contract method 0x63ee8ad1.
//
// Solidity: function cancelVouch(address counterparty) returns()
func (_Sybil *SybilTransactor) CancelVouch(opts *bind.TransactOpts, counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "cancelVouch", counterparty)
}

// CancelVouch is a paid mutator transaction binding the contract method 0x63ee8ad1.
//
// Solidity: function cancelVouch(address counterparty) returns()
func (_Sybil *SybilSession) CancelVouch(counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.CancelVouch(&_Sybil.TransactOpts, counterparty)
}

// CancelVouch is a paid mutator transaction binding the contract method 0x63ee8ad1.
//
// Solidity: function cancelVouch(address counterparty) returns()
func (_Sybil *SybilTransactorSession) CancelVouch(counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.CancelVouch(&_Sybil.TransactOpts, counterparty)
}

// CloseWithoutSteal is a paid mutator transaction binding the contract method 0x75a5c10f.
//
// Solidity: function closeWithoutSteal(address counterparty) returns()
func (_Sybil *SybilTransactor) CloseWithoutSteal(opts *bind.TransactOpts, counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "closeWithoutSteal", counterparty)
}

// CloseWithoutSteal is a paid mutator transaction binding the contract method 0x75a5c10f.
//
// Solidity: function closeWithoutSteal(address counterparty) returns()
func (_Sybil *SybilSession) CloseWithoutSteal(counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.CloseWithoutSteal(&_Sybil.TransactOpts, counterparty)
}

// CloseWithoutSteal is a paid mutator transaction binding the contract method 0x75a5c10f.
//
// Solidity: function closeWithoutSteal(address counterparty) returns()
func (_Sybil *SybilTransactorSession) CloseWithoutSteal(counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.CloseWithoutSteal(&_Sybil.TransactOpts, counterparty)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Sybil *SybilTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Sybil *SybilSession) Deposit() (*types.Transaction, error) {
	return _Sybil.Contract.Deposit(&_Sybil.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Sybil *SybilTransactorSession) Deposit() (*types.Transaction, error) {
	return _Sybil.Contract.Deposit(&_Sybil.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0xbcdf569b.
//
// Solidity: function finalize(address a, address b) returns()
func (_Sybil *SybilTransactor) Finalize(opts *bind.TransactOpts, a common.Address, b common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "finalize", a, b)
}

// Finalize is a paid mutator transaction binding the contract method 0xbcdf569b.
//
// Solidity: function finalize(address a, address b) returns()
func (_Sybil *SybilSession) Finalize(a common.Address, b common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Finalize(&_Sybil.TransactOpts, a, b)
}

// Finalize is a paid mutator transaction binding the contract method 0xbcdf569b.
//
// Solidity: function finalize(address a, address b) returns()
func (_Sybil *SybilTransactorSession) Finalize(a common.Address, b common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Finalize(&_Sybil.TransactOpts, a, b)
}

// SetBatchSize is a paid mutator transaction binding the contract method 0x8ff0ab7f.
//
// Solidity: function setBatchSize(uint32 _n) returns()
func (_Sybil *SybilTransactor) SetBatchSize(opts *bind.TransactOpts, _n uint32) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "setBatchSize", _n)
}

// SetBatchSize is a paid mutator transaction binding the contract method 0x8ff0ab7f.
//
// Solidity: function setBatchSize(uint32 _n) returns()
func (_Sybil *SybilSession) SetBatchSize(_n uint32) (*types.Transaction, error) {
	return _Sybil.Contract.SetBatchSize(&_Sybil.TransactOpts, _n)
}

// SetBatchSize is a paid mutator transaction binding the contract method 0x8ff0ab7f.
//
// Solidity: function setBatchSize(uint32 _n) returns()
func (_Sybil *SybilTransactorSession) SetBatchSize(_n uint32) (*types.Transaction, error) {
	return _Sybil.Contract.SetBatchSize(&_Sybil.TransactOpts, _n)
}

// SetParams is a paid mutator transaction binding the contract method 0x0088c6be.
//
// Solidity: function setParams(uint256 _s, uint64 _t) returns()
func (_Sybil *SybilTransactor) SetParams(opts *bind.TransactOpts, _s *big.Int, _t uint64) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "setParams", _s, _t)
}

// SetParams is a paid mutator transaction binding the contract method 0x0088c6be.
//
// Solidity: function setParams(uint256 _s, uint64 _t) returns()
func (_Sybil *SybilSession) SetParams(_s *big.Int, _t uint64) (*types.Transaction, error) {
	return _Sybil.Contract.SetParams(&_Sybil.TransactOpts, _s, _t)
}

// SetParams is a paid mutator transaction binding the contract method 0x0088c6be.
//
// Solidity: function setParams(uint256 _s, uint64 _t) returns()
func (_Sybil *SybilTransactorSession) SetParams(_s *big.Int, _t uint64) (*types.Transaction, error) {
	return _Sybil.Contract.SetParams(&_Sybil.TransactOpts, _s, _t)
}

// Steal is a paid mutator transaction binding the contract method 0x5dc8abb2.
//
// Solidity: function steal(address counterparty) returns()
func (_Sybil *SybilTransactor) Steal(opts *bind.TransactOpts, counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "steal", counterparty)
}

// Steal is a paid mutator transaction binding the contract method 0x5dc8abb2.
//
// Solidity: function steal(address counterparty) returns()
func (_Sybil *SybilSession) Steal(counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Steal(&_Sybil.TransactOpts, counterparty)
}

// Steal is a paid mutator transaction binding the contract method 0x5dc8abb2.
//
// Solidity: function steal(address counterparty) returns()
func (_Sybil *SybilTransactorSession) Steal(counterparty common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Steal(&_Sybil.TransactOpts, counterparty)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x29bc9852.
//
// Solidity: function submitBatch(bytes32 newGraph, bytes32 newScore, uint32 n, bytes proof) returns()
func (_Sybil *SybilTransactor) SubmitBatch(opts *bind.TransactOpts, newGraph [32]byte, newScore [32]byte, n uint32, proof []byte) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "submitBatch", newGraph, newScore, n, proof)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x29bc9852.
//
// Solidity: function submitBatch(bytes32 newGraph, bytes32 newScore, uint32 n, bytes proof) returns()
func (_Sybil *SybilSession) SubmitBatch(newGraph [32]byte, newScore [32]byte, n uint32, proof []byte) (*types.Transaction, error) {
	return _Sybil.Contract.SubmitBatch(&_Sybil.TransactOpts, newGraph, newScore, n, proof)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x29bc9852.
//
// Solidity: function submitBatch(bytes32 newGraph, bytes32 newScore, uint32 n, bytes proof) returns()
func (_Sybil *SybilTransactorSession) SubmitBatch(newGraph [32]byte, newScore [32]byte, n uint32, proof []byte) (*types.Transaction, error) {
	return _Sybil.Contract.SubmitBatch(&_Sybil.TransactOpts, newGraph, newScore, n, proof)
}

// Vouch is a paid mutator transaction binding the contract method 0xdd66e16b.
//
// Solidity: function vouch(address subject) payable returns()
func (_Sybil *SybilTransactor) Vouch(opts *bind.TransactOpts, subject common.Address) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "vouch", subject)
}

// Vouch is a paid mutator transaction binding the contract method 0xdd66e16b.
//
// Solidity: function vouch(address subject) payable returns()
func (_Sybil *SybilSession) Vouch(subject common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Vouch(&_Sybil.TransactOpts, subject)
}

// Vouch is a paid mutator transaction binding the contract method 0xdd66e16b.
//
// Solidity: function vouch(address subject) payable returns()
func (_Sybil *SybilTransactorSession) Vouch(subject common.Address) (*types.Transaction, error) {
	return _Sybil.Contract.Vouch(&_Sybil.TransactOpts, subject)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Sybil *SybilTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Sybil.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Sybil *SybilSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Withdraw(&_Sybil.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Sybil *SybilTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Sybil.Contract.Withdraw(&_Sybil.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sybil *SybilTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sybil.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sybil *SybilSession) Receive() (*types.Transaction, error) {
	return _Sybil.Contract.Receive(&_Sybil.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sybil *SybilTransactorSession) Receive() (*types.Transaction, error) {
	return _Sybil.Contract.Receive(&_Sybil.TransactOpts)
}

// SybilAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Sybil contract.
type SybilAccountCreatedIterator struct {
	Event *SybilAccountCreated // Event containing the contract specifics and raw log

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
func (it *SybilAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilAccountCreated)
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
		it.Event = new(SybilAccountCreated)
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
func (it *SybilAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilAccountCreated represents a AccountCreated event raised by the Sybil contract.
type SybilAccountCreated struct {
	Owner common.Address
	Idx   uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x1c6dbe19b5b78f3f71631ae073544d89823a3a2b9c28f7bd4a01d219779b4e58.
//
// Solidity: event AccountCreated(address indexed owner, uint32 indexed idx)
func (_Sybil *SybilFilterer) FilterAccountCreated(opts *bind.FilterOpts, owner []common.Address, idx []uint32) (*SybilAccountCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "AccountCreated", ownerRule, idxRule)
	if err != nil {
		return nil, err
	}
	return &SybilAccountCreatedIterator{contract: _Sybil.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x1c6dbe19b5b78f3f71631ae073544d89823a3a2b9c28f7bd4a01d219779b4e58.
//
// Solidity: event AccountCreated(address indexed owner, uint32 indexed idx)
func (_Sybil *SybilFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *SybilAccountCreated, owner []common.Address, idx []uint32) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "AccountCreated", ownerRule, idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilAccountCreated)
				if err := _Sybil.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0x1c6dbe19b5b78f3f71631ae073544d89823a3a2b9c28f7bd4a01d219779b4e58.
//
// Solidity: event AccountCreated(address indexed owner, uint32 indexed idx)
func (_Sybil *SybilFilterer) ParseAccountCreated(log types.Log) (*SybilAccountCreated, error) {
	event := new(SybilAccountCreated)
	if err := _Sybil.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilBatchSizeUpdatedIterator is returned from FilterBatchSizeUpdated and is used to iterate over the raw logs and unpacked data for BatchSizeUpdated events raised by the Sybil contract.
type SybilBatchSizeUpdatedIterator struct {
	Event *SybilBatchSizeUpdated // Event containing the contract specifics and raw log

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
func (it *SybilBatchSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilBatchSizeUpdated)
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
		it.Event = new(SybilBatchSizeUpdated)
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
func (it *SybilBatchSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilBatchSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilBatchSizeUpdated represents a BatchSizeUpdated event raised by the Sybil contract.
type SybilBatchSizeUpdated struct {
	BatchSize uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchSizeUpdated is a free log retrieval operation binding the contract event 0x12c716d53525eb23412a351798dd6de5dd019613ae5ea581fa3fa808bd59b3b8.
//
// Solidity: event BatchSizeUpdated(uint32 batchSize)
func (_Sybil *SybilFilterer) FilterBatchSizeUpdated(opts *bind.FilterOpts) (*SybilBatchSizeUpdatedIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "BatchSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &SybilBatchSizeUpdatedIterator{contract: _Sybil.contract, event: "BatchSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchBatchSizeUpdated is a free log subscription operation binding the contract event 0x12c716d53525eb23412a351798dd6de5dd019613ae5ea581fa3fa808bd59b3b8.
//
// Solidity: event BatchSizeUpdated(uint32 batchSize)
func (_Sybil *SybilFilterer) WatchBatchSizeUpdated(opts *bind.WatchOpts, sink chan<- *SybilBatchSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "BatchSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilBatchSizeUpdated)
				if err := _Sybil.contract.UnpackLog(event, "BatchSizeUpdated", log); err != nil {
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

// ParseBatchSizeUpdated is a log parse operation binding the contract event 0x12c716d53525eb23412a351798dd6de5dd019613ae5ea581fa3fa808bd59b3b8.
//
// Solidity: event BatchSizeUpdated(uint32 batchSize)
func (_Sybil *SybilFilterer) ParseBatchSizeUpdated(log types.Log) (*SybilBatchSizeUpdated, error) {
	event := new(SybilBatchSizeUpdated)
	if err := _Sybil.contract.UnpackLog(event, "BatchSizeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilBatchSubmittedIterator is returned from FilterBatchSubmitted and is used to iterate over the raw logs and unpacked data for BatchSubmitted events raised by the Sybil contract.
type SybilBatchSubmittedIterator struct {
	Event *SybilBatchSubmitted // Event containing the contract specifics and raw log

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
func (it *SybilBatchSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilBatchSubmitted)
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
		it.Event = new(SybilBatchSubmitted)
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
func (it *SybilBatchSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilBatchSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilBatchSubmitted represents a BatchSubmitted event raised by the Sybil contract.
type SybilBatchSubmitted struct {
	BatchId      uint64
	Count        uint32
	StorageHash  [32]byte
	NewGraphRoot [32]byte
	NewScoreRoot [32]byte
	EdgesPacked  []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBatchSubmitted is a free log retrieval operation binding the contract event 0xe94308f806199d660ecc4b105960dd1d833e84ff46305358498521934d629049.
//
// Solidity: event BatchSubmitted(uint64 indexed batchId, uint32 count, bytes32 storageHash, bytes32 newGraphRoot, bytes32 newScoreRoot, bytes edgesPacked)
func (_Sybil *SybilFilterer) FilterBatchSubmitted(opts *bind.FilterOpts, batchId []uint64) (*SybilBatchSubmittedIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "BatchSubmitted", batchIdRule)
	if err != nil {
		return nil, err
	}
	return &SybilBatchSubmittedIterator{contract: _Sybil.contract, event: "BatchSubmitted", logs: logs, sub: sub}, nil
}

// WatchBatchSubmitted is a free log subscription operation binding the contract event 0xe94308f806199d660ecc4b105960dd1d833e84ff46305358498521934d629049.
//
// Solidity: event BatchSubmitted(uint64 indexed batchId, uint32 count, bytes32 storageHash, bytes32 newGraphRoot, bytes32 newScoreRoot, bytes edgesPacked)
func (_Sybil *SybilFilterer) WatchBatchSubmitted(opts *bind.WatchOpts, sink chan<- *SybilBatchSubmitted, batchId []uint64) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "BatchSubmitted", batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilBatchSubmitted)
				if err := _Sybil.contract.UnpackLog(event, "BatchSubmitted", log); err != nil {
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

// ParseBatchSubmitted is a log parse operation binding the contract event 0xe94308f806199d660ecc4b105960dd1d833e84ff46305358498521934d629049.
//
// Solidity: event BatchSubmitted(uint64 indexed batchId, uint32 count, bytes32 storageHash, bytes32 newGraphRoot, bytes32 newScoreRoot, bytes edgesPacked)
func (_Sybil *SybilFilterer) ParseBatchSubmitted(log types.Log) (*SybilBatchSubmitted, error) {
	event := new(SybilBatchSubmitted)
	if err := _Sybil.contract.UnpackLog(event, "BatchSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilClosedNoLinkIterator is returned from FilterClosedNoLink and is used to iterate over the raw logs and unpacked data for ClosedNoLink events raised by the Sybil contract.
type SybilClosedNoLinkIterator struct {
	Event *SybilClosedNoLink // Event containing the contract specifics and raw log

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
func (it *SybilClosedNoLinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilClosedNoLink)
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
		it.Event = new(SybilClosedNoLink)
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
func (it *SybilClosedNoLinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilClosedNoLinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilClosedNoLink represents a ClosedNoLink event raised by the Sybil contract.
type SybilClosedNoLink struct {
	Caller       common.Address
	Counterparty common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClosedNoLink is a free log retrieval operation binding the contract event 0x725fa7b1236efef1a62b7ed5eabe8a7a9ab6f621e5149ca4d55ad57d7f6db943.
//
// Solidity: event ClosedNoLink(address indexed caller, address indexed counterparty)
func (_Sybil *SybilFilterer) FilterClosedNoLink(opts *bind.FilterOpts, caller []common.Address, counterparty []common.Address) (*SybilClosedNoLinkIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var counterpartyRule []interface{}
	for _, counterpartyItem := range counterparty {
		counterpartyRule = append(counterpartyRule, counterpartyItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "ClosedNoLink", callerRule, counterpartyRule)
	if err != nil {
		return nil, err
	}
	return &SybilClosedNoLinkIterator{contract: _Sybil.contract, event: "ClosedNoLink", logs: logs, sub: sub}, nil
}

// WatchClosedNoLink is a free log subscription operation binding the contract event 0x725fa7b1236efef1a62b7ed5eabe8a7a9ab6f621e5149ca4d55ad57d7f6db943.
//
// Solidity: event ClosedNoLink(address indexed caller, address indexed counterparty)
func (_Sybil *SybilFilterer) WatchClosedNoLink(opts *bind.WatchOpts, sink chan<- *SybilClosedNoLink, caller []common.Address, counterparty []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var counterpartyRule []interface{}
	for _, counterpartyItem := range counterparty {
		counterpartyRule = append(counterpartyRule, counterpartyItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "ClosedNoLink", callerRule, counterpartyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilClosedNoLink)
				if err := _Sybil.contract.UnpackLog(event, "ClosedNoLink", log); err != nil {
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

// ParseClosedNoLink is a log parse operation binding the contract event 0x725fa7b1236efef1a62b7ed5eabe8a7a9ab6f621e5149ca4d55ad57d7f6db943.
//
// Solidity: event ClosedNoLink(address indexed caller, address indexed counterparty)
func (_Sybil *SybilFilterer) ParseClosedNoLink(log types.Log) (*SybilClosedNoLink, error) {
	event := new(SybilClosedNoLink)
	if err := _Sybil.contract.UnpackLog(event, "ClosedNoLink", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Sybil contract.
type SybilDepositedIterator struct {
	Event *SybilDeposited // Event containing the contract specifics and raw log

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
func (it *SybilDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilDeposited)
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
		it.Event = new(SybilDeposited)
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
func (it *SybilDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilDeposited represents a Deposited event raised by the Sybil contract.
type SybilDeposited struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_Sybil *SybilFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*SybilDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &SybilDepositedIterator{contract: _Sybil.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_Sybil *SybilFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *SybilDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilDeposited)
				if err := _Sybil.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed user, uint256 amount)
func (_Sybil *SybilFilterer) ParseDeposited(log types.Log) (*SybilDeposited, error) {
	event := new(SybilDeposited)
	if err := _Sybil.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilLinkCancelledIterator is returned from FilterLinkCancelled and is used to iterate over the raw logs and unpacked data for LinkCancelled events raised by the Sybil contract.
type SybilLinkCancelledIterator struct {
	Event *SybilLinkCancelled // Event containing the contract specifics and raw log

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
func (it *SybilLinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilLinkCancelled)
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
		it.Event = new(SybilLinkCancelled)
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
func (it *SybilLinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilLinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilLinkCancelled represents a LinkCancelled event raised by the Sybil contract.
type SybilLinkCancelled struct {
	Lo     common.Address
	Hi     common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLinkCancelled is a free log retrieval operation binding the contract event 0x0c796753764d7c3d1579290417d4cd325202cee1ad6a7d5b0fc40a31e312f1d7.
//
// Solidity: event LinkCancelled(address indexed lo, address indexed hi, address indexed caller)
func (_Sybil *SybilFilterer) FilterLinkCancelled(opts *bind.FilterOpts, lo []common.Address, hi []common.Address, caller []common.Address) (*SybilLinkCancelledIterator, error) {

	var loRule []interface{}
	for _, loItem := range lo {
		loRule = append(loRule, loItem)
	}
	var hiRule []interface{}
	for _, hiItem := range hi {
		hiRule = append(hiRule, hiItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "LinkCancelled", loRule, hiRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &SybilLinkCancelledIterator{contract: _Sybil.contract, event: "LinkCancelled", logs: logs, sub: sub}, nil
}

// WatchLinkCancelled is a free log subscription operation binding the contract event 0x0c796753764d7c3d1579290417d4cd325202cee1ad6a7d5b0fc40a31e312f1d7.
//
// Solidity: event LinkCancelled(address indexed lo, address indexed hi, address indexed caller)
func (_Sybil *SybilFilterer) WatchLinkCancelled(opts *bind.WatchOpts, sink chan<- *SybilLinkCancelled, lo []common.Address, hi []common.Address, caller []common.Address) (event.Subscription, error) {

	var loRule []interface{}
	for _, loItem := range lo {
		loRule = append(loRule, loItem)
	}
	var hiRule []interface{}
	for _, hiItem := range hi {
		hiRule = append(hiRule, hiItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "LinkCancelled", loRule, hiRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilLinkCancelled)
				if err := _Sybil.contract.UnpackLog(event, "LinkCancelled", log); err != nil {
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

// ParseLinkCancelled is a log parse operation binding the contract event 0x0c796753764d7c3d1579290417d4cd325202cee1ad6a7d5b0fc40a31e312f1d7.
//
// Solidity: event LinkCancelled(address indexed lo, address indexed hi, address indexed caller)
func (_Sybil *SybilFilterer) ParseLinkCancelled(log types.Log) (*SybilLinkCancelled, error) {
	event := new(SybilLinkCancelled)
	if err := _Sybil.contract.UnpackLog(event, "LinkCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilLinkedIterator is returned from FilterLinked and is used to iterate over the raw logs and unpacked data for Linked events raised by the Sybil contract.
type SybilLinkedIterator struct {
	Event *SybilLinked // Event containing the contract specifics and raw log

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
func (it *SybilLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilLinked)
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
		it.Event = new(SybilLinked)
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
func (it *SybilLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilLinked represents a Linked event raised by the Sybil contract.
type SybilLinked struct {
	Lo          common.Address
	Hi          common.Address
	WindowStart uint64
	WindowEnd   uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLinked is a free log retrieval operation binding the contract event 0xd6242ceaf1369834c676d8f04c80c084b4876c59e1df27b08bca55cde345028d.
//
// Solidity: event Linked(address indexed lo, address indexed hi, uint64 windowStart, uint64 windowEnd)
func (_Sybil *SybilFilterer) FilterLinked(opts *bind.FilterOpts, lo []common.Address, hi []common.Address) (*SybilLinkedIterator, error) {

	var loRule []interface{}
	for _, loItem := range lo {
		loRule = append(loRule, loItem)
	}
	var hiRule []interface{}
	for _, hiItem := range hi {
		hiRule = append(hiRule, hiItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "Linked", loRule, hiRule)
	if err != nil {
		return nil, err
	}
	return &SybilLinkedIterator{contract: _Sybil.contract, event: "Linked", logs: logs, sub: sub}, nil
}

// WatchLinked is a free log subscription operation binding the contract event 0xd6242ceaf1369834c676d8f04c80c084b4876c59e1df27b08bca55cde345028d.
//
// Solidity: event Linked(address indexed lo, address indexed hi, uint64 windowStart, uint64 windowEnd)
func (_Sybil *SybilFilterer) WatchLinked(opts *bind.WatchOpts, sink chan<- *SybilLinked, lo []common.Address, hi []common.Address) (event.Subscription, error) {

	var loRule []interface{}
	for _, loItem := range lo {
		loRule = append(loRule, loItem)
	}
	var hiRule []interface{}
	for _, hiItem := range hi {
		hiRule = append(hiRule, hiItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "Linked", loRule, hiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilLinked)
				if err := _Sybil.contract.UnpackLog(event, "Linked", log); err != nil {
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

// ParseLinked is a log parse operation binding the contract event 0xd6242ceaf1369834c676d8f04c80c084b4876c59e1df27b08bca55cde345028d.
//
// Solidity: event Linked(address indexed lo, address indexed hi, uint64 windowStart, uint64 windowEnd)
func (_Sybil *SybilFilterer) ParseLinked(log types.Log) (*SybilLinked, error) {
	event := new(SybilLinked)
	if err := _Sybil.contract.UnpackLog(event, "Linked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilParamsUpdatedIterator is returned from FilterParamsUpdated and is used to iterate over the raw logs and unpacked data for ParamsUpdated events raised by the Sybil contract.
type SybilParamsUpdatedIterator struct {
	Event *SybilParamsUpdated // Event containing the contract specifics and raw log

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
func (it *SybilParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilParamsUpdated)
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
		it.Event = new(SybilParamsUpdated)
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
func (it *SybilParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilParamsUpdated represents a ParamsUpdated event raised by the Sybil contract.
type SybilParamsUpdated struct {
	StakeS  *big.Int
	WindowT uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterParamsUpdated is a free log retrieval operation binding the contract event 0x49cff52c74120d8c8fa989c5aa8f254add8c0ba6abe75eeb06f420d9abf5d52b.
//
// Solidity: event ParamsUpdated(uint256 stakeS, uint64 windowT)
func (_Sybil *SybilFilterer) FilterParamsUpdated(opts *bind.FilterOpts) (*SybilParamsUpdatedIterator, error) {

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "ParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &SybilParamsUpdatedIterator{contract: _Sybil.contract, event: "ParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchParamsUpdated is a free log subscription operation binding the contract event 0x49cff52c74120d8c8fa989c5aa8f254add8c0ba6abe75eeb06f420d9abf5d52b.
//
// Solidity: event ParamsUpdated(uint256 stakeS, uint64 windowT)
func (_Sybil *SybilFilterer) WatchParamsUpdated(opts *bind.WatchOpts, sink chan<- *SybilParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "ParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilParamsUpdated)
				if err := _Sybil.contract.UnpackLog(event, "ParamsUpdated", log); err != nil {
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

// ParseParamsUpdated is a log parse operation binding the contract event 0x49cff52c74120d8c8fa989c5aa8f254add8c0ba6abe75eeb06f420d9abf5d52b.
//
// Solidity: event ParamsUpdated(uint256 stakeS, uint64 windowT)
func (_Sybil *SybilFilterer) ParseParamsUpdated(log types.Log) (*SybilParamsUpdated, error) {
	event := new(SybilParamsUpdated)
	if err := _Sybil.contract.UnpackLog(event, "ParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilScoreSyncedIterator is returned from FilterScoreSynced and is used to iterate over the raw logs and unpacked data for ScoreSynced events raised by the Sybil contract.
type SybilScoreSyncedIterator struct {
	Event *SybilScoreSynced // Event containing the contract specifics and raw log

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
func (it *SybilScoreSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilScoreSynced)
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
		it.Event = new(SybilScoreSynced)
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
func (it *SybilScoreSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilScoreSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilScoreSynced represents a ScoreSynced event raised by the Sybil contract.
type SybilScoreSynced struct {
	User    common.Address
	BatchId uint64
	Score   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterScoreSynced is a free log retrieval operation binding the contract event 0x1b9a85024c07566c217b6364cf1724c6d910a1389dc84ec361f3296242f1b746.
//
// Solidity: event ScoreSynced(address indexed user, uint64 indexed batchId, uint256 score)
func (_Sybil *SybilFilterer) FilterScoreSynced(opts *bind.FilterOpts, user []common.Address, batchId []uint64) (*SybilScoreSyncedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "ScoreSynced", userRule, batchIdRule)
	if err != nil {
		return nil, err
	}
	return &SybilScoreSyncedIterator{contract: _Sybil.contract, event: "ScoreSynced", logs: logs, sub: sub}, nil
}

// WatchScoreSynced is a free log subscription operation binding the contract event 0x1b9a85024c07566c217b6364cf1724c6d910a1389dc84ec361f3296242f1b746.
//
// Solidity: event ScoreSynced(address indexed user, uint64 indexed batchId, uint256 score)
func (_Sybil *SybilFilterer) WatchScoreSynced(opts *bind.WatchOpts, sink chan<- *SybilScoreSynced, user []common.Address, batchId []uint64) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "ScoreSynced", userRule, batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilScoreSynced)
				if err := _Sybil.contract.UnpackLog(event, "ScoreSynced", log); err != nil {
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

// ParseScoreSynced is a log parse operation binding the contract event 0x1b9a85024c07566c217b6364cf1724c6d910a1389dc84ec361f3296242f1b746.
//
// Solidity: event ScoreSynced(address indexed user, uint64 indexed batchId, uint256 score)
func (_Sybil *SybilFilterer) ParseScoreSynced(log types.Log) (*SybilScoreSynced, error) {
	event := new(SybilScoreSynced)
	if err := _Sybil.contract.UnpackLog(event, "ScoreSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilStolenIterator is returned from FilterStolen and is used to iterate over the raw logs and unpacked data for Stolen events raised by the Sybil contract.
type SybilStolenIterator struct {
	Event *SybilStolen // Event containing the contract specifics and raw log

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
func (it *SybilStolenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilStolen)
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
		it.Event = new(SybilStolen)
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
func (it *SybilStolenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilStolenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilStolen represents a Stolen event raised by the Sybil contract.
type SybilStolen struct {
	Thief  common.Address
	Victim common.Address
	Payout *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStolen is a free log retrieval operation binding the contract event 0xc818f8b73119babbe2632ee0c7e73786aa3a64aa5a3bcc8c2963df4967fbe223.
//
// Solidity: event Stolen(address indexed thief, address indexed victim, uint256 payout)
func (_Sybil *SybilFilterer) FilterStolen(opts *bind.FilterOpts, thief []common.Address, victim []common.Address) (*SybilStolenIterator, error) {

	var thiefRule []interface{}
	for _, thiefItem := range thief {
		thiefRule = append(thiefRule, thiefItem)
	}
	var victimRule []interface{}
	for _, victimItem := range victim {
		victimRule = append(victimRule, victimItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "Stolen", thiefRule, victimRule)
	if err != nil {
		return nil, err
	}
	return &SybilStolenIterator{contract: _Sybil.contract, event: "Stolen", logs: logs, sub: sub}, nil
}

// WatchStolen is a free log subscription operation binding the contract event 0xc818f8b73119babbe2632ee0c7e73786aa3a64aa5a3bcc8c2963df4967fbe223.
//
// Solidity: event Stolen(address indexed thief, address indexed victim, uint256 payout)
func (_Sybil *SybilFilterer) WatchStolen(opts *bind.WatchOpts, sink chan<- *SybilStolen, thief []common.Address, victim []common.Address) (event.Subscription, error) {

	var thiefRule []interface{}
	for _, thiefItem := range thief {
		thiefRule = append(thiefRule, thiefItem)
	}
	var victimRule []interface{}
	for _, victimItem := range victim {
		victimRule = append(victimRule, victimItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "Stolen", thiefRule, victimRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilStolen)
				if err := _Sybil.contract.UnpackLog(event, "Stolen", log); err != nil {
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

// ParseStolen is a log parse operation binding the contract event 0xc818f8b73119babbe2632ee0c7e73786aa3a64aa5a3bcc8c2963df4967fbe223.
//
// Solidity: event Stolen(address indexed thief, address indexed victim, uint256 payout)
func (_Sybil *SybilFilterer) ParseStolen(log types.Log) (*SybilStolen, error) {
	event := new(SybilStolen)
	if err := _Sybil.contract.UnpackLog(event, "Stolen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilVouchedIterator is returned from FilterVouched and is used to iterate over the raw logs and unpacked data for Vouched events raised by the Sybil contract.
type SybilVouchedIterator struct {
	Event *SybilVouched // Event containing the contract specifics and raw log

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
func (it *SybilVouchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilVouched)
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
		it.Event = new(SybilVouched)
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
func (it *SybilVouchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilVouchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilVouched represents a Vouched event raised by the Sybil contract.
type SybilVouched struct {
	Attester common.Address
	Subject  common.Address
	Stake    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVouched is a free log retrieval operation binding the contract event 0x49517064a7e568cc243d6d42d1e1388d887dea9ef158bc6f15e34da4055bde83.
//
// Solidity: event Vouched(address indexed attester, address indexed subject, uint256 stake)
func (_Sybil *SybilFilterer) FilterVouched(opts *bind.FilterOpts, attester []common.Address, subject []common.Address) (*SybilVouchedIterator, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "Vouched", attesterRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return &SybilVouchedIterator{contract: _Sybil.contract, event: "Vouched", logs: logs, sub: sub}, nil
}

// WatchVouched is a free log subscription operation binding the contract event 0x49517064a7e568cc243d6d42d1e1388d887dea9ef158bc6f15e34da4055bde83.
//
// Solidity: event Vouched(address indexed attester, address indexed subject, uint256 stake)
func (_Sybil *SybilFilterer) WatchVouched(opts *bind.WatchOpts, sink chan<- *SybilVouched, attester []common.Address, subject []common.Address) (event.Subscription, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "Vouched", attesterRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilVouched)
				if err := _Sybil.contract.UnpackLog(event, "Vouched", log); err != nil {
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

// ParseVouched is a log parse operation binding the contract event 0x49517064a7e568cc243d6d42d1e1388d887dea9ef158bc6f15e34da4055bde83.
//
// Solidity: event Vouched(address indexed attester, address indexed subject, uint256 stake)
func (_Sybil *SybilFilterer) ParseVouched(log types.Log) (*SybilVouched, error) {
	event := new(SybilVouched)
	if err := _Sybil.contract.UnpackLog(event, "Vouched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilWindowOpenedIterator is returned from FilterWindowOpened and is used to iterate over the raw logs and unpacked data for WindowOpened events raised by the Sybil contract.
type SybilWindowOpenedIterator struct {
	Event *SybilWindowOpened // Event containing the contract specifics and raw log

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
func (it *SybilWindowOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilWindowOpened)
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
		it.Event = new(SybilWindowOpened)
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
func (it *SybilWindowOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilWindowOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilWindowOpened represents a WindowOpened event raised by the Sybil contract.
type SybilWindowOpened struct {
	Lo    common.Address
	Hi    common.Address
	Start uint64
	End   uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWindowOpened is a free log retrieval operation binding the contract event 0x548dbda810ccd7a77b123b86cb90a4c4f9bf87b9b1ed9d00f57c86bcb7842084.
//
// Solidity: event WindowOpened(address indexed lo, address indexed hi, uint64 start, uint64 end)
func (_Sybil *SybilFilterer) FilterWindowOpened(opts *bind.FilterOpts, lo []common.Address, hi []common.Address) (*SybilWindowOpenedIterator, error) {

	var loRule []interface{}
	for _, loItem := range lo {
		loRule = append(loRule, loItem)
	}
	var hiRule []interface{}
	for _, hiItem := range hi {
		hiRule = append(hiRule, hiItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "WindowOpened", loRule, hiRule)
	if err != nil {
		return nil, err
	}
	return &SybilWindowOpenedIterator{contract: _Sybil.contract, event: "WindowOpened", logs: logs, sub: sub}, nil
}

// WatchWindowOpened is a free log subscription operation binding the contract event 0x548dbda810ccd7a77b123b86cb90a4c4f9bf87b9b1ed9d00f57c86bcb7842084.
//
// Solidity: event WindowOpened(address indexed lo, address indexed hi, uint64 start, uint64 end)
func (_Sybil *SybilFilterer) WatchWindowOpened(opts *bind.WatchOpts, sink chan<- *SybilWindowOpened, lo []common.Address, hi []common.Address) (event.Subscription, error) {

	var loRule []interface{}
	for _, loItem := range lo {
		loRule = append(loRule, loItem)
	}
	var hiRule []interface{}
	for _, hiItem := range hi {
		hiRule = append(hiRule, hiItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "WindowOpened", loRule, hiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilWindowOpened)
				if err := _Sybil.contract.UnpackLog(event, "WindowOpened", log); err != nil {
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

// ParseWindowOpened is a log parse operation binding the contract event 0x548dbda810ccd7a77b123b86cb90a4c4f9bf87b9b1ed9d00f57c86bcb7842084.
//
// Solidity: event WindowOpened(address indexed lo, address indexed hi, uint64 start, uint64 end)
func (_Sybil *SybilFilterer) ParseWindowOpened(log types.Log) (*SybilWindowOpened, error) {
	event := new(SybilWindowOpened)
	if err := _Sybil.contract.UnpackLog(event, "WindowOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SybilWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Sybil contract.
type SybilWithdrawnIterator struct {
	Event *SybilWithdrawn // Event containing the contract specifics and raw log

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
func (it *SybilWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SybilWithdrawn)
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
		it.Event = new(SybilWithdrawn)
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
func (it *SybilWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SybilWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SybilWithdrawn represents a Withdrawn event raised by the Sybil contract.
type SybilWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Sybil *SybilFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*SybilWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Sybil.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &SybilWithdrawnIterator{contract: _Sybil.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Sybil *SybilFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *SybilWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Sybil.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SybilWithdrawn)
				if err := _Sybil.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_Sybil *SybilFilterer) ParseWithdrawn(log types.Log) (*SybilWithdrawn, error) {
	event := new(SybilWithdrawn)
	if err := _Sybil.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
