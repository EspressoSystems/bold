// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package assertionStakingPoolgen

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

// AssertionInputs is an auto generated low-level Go binding around an user-defined struct.
type AssertionInputs struct {
	BeforeStateData BeforeStateData
	BeforeState     AssertionState
	AfterState      AssertionState
}

// AssertionState is an auto generated low-level Go binding around an user-defined struct.
type AssertionState struct {
	GlobalState    GlobalState
	MachineStatus  uint8
	EndHistoryRoot [32]byte
}

// BeforeStateData is an auto generated low-level Go binding around an user-defined struct.
type BeforeStateData struct {
	PrevPrevAssertionHash [32]byte
	SequencerBatchAcc     [32]byte
	ConfigData            ConfigData
}

// ConfigData is an auto generated low-level Go binding around an user-defined struct.
type ConfigData struct {
	WasmModuleRoot      [32]byte
	RequiredStake       *big.Int
	ChallengeManager    common.Address
	ConfirmPeriodBlocks uint64
	NextInboxPosition   uint64
}

// CreateEdgeArgs is an auto generated low-level Go binding around an user-defined struct.
type CreateEdgeArgs struct {
	Level          uint8
	EndHistoryRoot [32]byte
	EndHeight      *big.Int
	ClaimId        [32]byte
	PrefixProof    []byte
	Proof          []byte
}

// GlobalState is an auto generated low-level Go binding around an user-defined struct.
type GlobalState struct {
	Bytes32Vals [2][32]byte
	U64Vals     [2]uint64
}

// AbsBoldStakingPoolMetaData contains all meta data concerning the AbsBoldStakingPool contract.
var AbsBoldStakingPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"AmountExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositIntoPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AbsBoldStakingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AbsBoldStakingPoolMetaData.ABI instead.
var AbsBoldStakingPoolABI = AbsBoldStakingPoolMetaData.ABI

// AbsBoldStakingPool is an auto generated Go binding around an Ethereum contract.
type AbsBoldStakingPool struct {
	AbsBoldStakingPoolCaller     // Read-only binding to the contract
	AbsBoldStakingPoolTransactor // Write-only binding to the contract
	AbsBoldStakingPoolFilterer   // Log filterer for contract events
}

// AbsBoldStakingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbsBoldStakingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsBoldStakingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbsBoldStakingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsBoldStakingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbsBoldStakingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbsBoldStakingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbsBoldStakingPoolSession struct {
	Contract     *AbsBoldStakingPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AbsBoldStakingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbsBoldStakingPoolCallerSession struct {
	Contract *AbsBoldStakingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AbsBoldStakingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbsBoldStakingPoolTransactorSession struct {
	Contract     *AbsBoldStakingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AbsBoldStakingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbsBoldStakingPoolRaw struct {
	Contract *AbsBoldStakingPool // Generic contract binding to access the raw methods on
}

// AbsBoldStakingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbsBoldStakingPoolCallerRaw struct {
	Contract *AbsBoldStakingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// AbsBoldStakingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbsBoldStakingPoolTransactorRaw struct {
	Contract *AbsBoldStakingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbsBoldStakingPool creates a new instance of AbsBoldStakingPool, bound to a specific deployed contract.
func NewAbsBoldStakingPool(address common.Address, backend bind.ContractBackend) (*AbsBoldStakingPool, error) {
	contract, err := bindAbsBoldStakingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbsBoldStakingPool{AbsBoldStakingPoolCaller: AbsBoldStakingPoolCaller{contract: contract}, AbsBoldStakingPoolTransactor: AbsBoldStakingPoolTransactor{contract: contract}, AbsBoldStakingPoolFilterer: AbsBoldStakingPoolFilterer{contract: contract}}, nil
}

// NewAbsBoldStakingPoolCaller creates a new read-only instance of AbsBoldStakingPool, bound to a specific deployed contract.
func NewAbsBoldStakingPoolCaller(address common.Address, caller bind.ContractCaller) (*AbsBoldStakingPoolCaller, error) {
	contract, err := bindAbsBoldStakingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbsBoldStakingPoolCaller{contract: contract}, nil
}

// NewAbsBoldStakingPoolTransactor creates a new write-only instance of AbsBoldStakingPool, bound to a specific deployed contract.
func NewAbsBoldStakingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AbsBoldStakingPoolTransactor, error) {
	contract, err := bindAbsBoldStakingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbsBoldStakingPoolTransactor{contract: contract}, nil
}

// NewAbsBoldStakingPoolFilterer creates a new log filterer instance of AbsBoldStakingPool, bound to a specific deployed contract.
func NewAbsBoldStakingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AbsBoldStakingPoolFilterer, error) {
	contract, err := bindAbsBoldStakingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbsBoldStakingPoolFilterer{contract: contract}, nil
}

// bindAbsBoldStakingPool binds a generic wrapper to an already deployed contract.
func bindAbsBoldStakingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbsBoldStakingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbsBoldStakingPool *AbsBoldStakingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbsBoldStakingPool.Contract.AbsBoldStakingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbsBoldStakingPool *AbsBoldStakingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.AbsBoldStakingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbsBoldStakingPool *AbsBoldStakingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.AbsBoldStakingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbsBoldStakingPool *AbsBoldStakingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbsBoldStakingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbsBoldStakingPool *AbsBoldStakingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbsBoldStakingPool *AbsBoldStakingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.contract.Transact(opts, method, params...)
}

// DepositBalance is a free data retrieval call binding the contract method 0x956501bb.
//
// Solidity: function depositBalance(address ) view returns(uint256)
func (_AbsBoldStakingPool *AbsBoldStakingPoolCaller) DepositBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AbsBoldStakingPool.contract.Call(opts, &out, "depositBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositBalance is a free data retrieval call binding the contract method 0x956501bb.
//
// Solidity: function depositBalance(address ) view returns(uint256)
func (_AbsBoldStakingPool *AbsBoldStakingPoolSession) DepositBalance(arg0 common.Address) (*big.Int, error) {
	return _AbsBoldStakingPool.Contract.DepositBalance(&_AbsBoldStakingPool.CallOpts, arg0)
}

// DepositBalance is a free data retrieval call binding the contract method 0x956501bb.
//
// Solidity: function depositBalance(address ) view returns(uint256)
func (_AbsBoldStakingPool *AbsBoldStakingPoolCallerSession) DepositBalance(arg0 common.Address) (*big.Int, error) {
	return _AbsBoldStakingPool.Contract.DepositBalance(&_AbsBoldStakingPool.CallOpts, arg0)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AbsBoldStakingPool *AbsBoldStakingPoolCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbsBoldStakingPool.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AbsBoldStakingPool *AbsBoldStakingPoolSession) StakeToken() (common.Address, error) {
	return _AbsBoldStakingPool.Contract.StakeToken(&_AbsBoldStakingPool.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AbsBoldStakingPool *AbsBoldStakingPoolCallerSession) StakeToken() (common.Address, error) {
	return _AbsBoldStakingPool.Contract.StakeToken(&_AbsBoldStakingPool.CallOpts)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x7476083b.
//
// Solidity: function depositIntoPool(uint256 amount) returns()
func (_AbsBoldStakingPool *AbsBoldStakingPoolTransactor) DepositIntoPool(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AbsBoldStakingPool.contract.Transact(opts, "depositIntoPool", amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x7476083b.
//
// Solidity: function depositIntoPool(uint256 amount) returns()
func (_AbsBoldStakingPool *AbsBoldStakingPoolSession) DepositIntoPool(amount *big.Int) (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.DepositIntoPool(&_AbsBoldStakingPool.TransactOpts, amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x7476083b.
//
// Solidity: function depositIntoPool(uint256 amount) returns()
func (_AbsBoldStakingPool *AbsBoldStakingPoolTransactorSession) DepositIntoPool(amount *big.Int) (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.DepositIntoPool(&_AbsBoldStakingPool.TransactOpts, amount)
}

// WithdrawFromPool26c0e5c5 is a paid mutator transaction binding the contract method 0x26c0e5c5.
//
// Solidity: function withdrawFromPool() returns()
func (_AbsBoldStakingPool *AbsBoldStakingPoolTransactor) WithdrawFromPool26c0e5c5(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbsBoldStakingPool.contract.Transact(opts, "withdrawFromPool")
}

// WithdrawFromPool26c0e5c5 is a paid mutator transaction binding the contract method 0x26c0e5c5.
//
// Solidity: function withdrawFromPool() returns()
func (_AbsBoldStakingPool *AbsBoldStakingPoolSession) WithdrawFromPool26c0e5c5() (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.WithdrawFromPool26c0e5c5(&_AbsBoldStakingPool.TransactOpts)
}

// WithdrawFromPool26c0e5c5 is a paid mutator transaction binding the contract method 0x26c0e5c5.
//
// Solidity: function withdrawFromPool() returns()
func (_AbsBoldStakingPool *AbsBoldStakingPoolTransactorSession) WithdrawFromPool26c0e5c5() (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.WithdrawFromPool26c0e5c5(&_AbsBoldStakingPool.TransactOpts)
}

// WithdrawFromPool30fc43ed is a paid mutator transaction binding the contract method 0x30fc43ed.
//
// Solidity: function withdrawFromPool(uint256 amount) returns()
func (_AbsBoldStakingPool *AbsBoldStakingPoolTransactor) WithdrawFromPool30fc43ed(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AbsBoldStakingPool.contract.Transact(opts, "withdrawFromPool0", amount)
}

// WithdrawFromPool30fc43ed is a paid mutator transaction binding the contract method 0x30fc43ed.
//
// Solidity: function withdrawFromPool(uint256 amount) returns()
func (_AbsBoldStakingPool *AbsBoldStakingPoolSession) WithdrawFromPool30fc43ed(amount *big.Int) (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.WithdrawFromPool30fc43ed(&_AbsBoldStakingPool.TransactOpts, amount)
}

// WithdrawFromPool30fc43ed is a paid mutator transaction binding the contract method 0x30fc43ed.
//
// Solidity: function withdrawFromPool(uint256 amount) returns()
func (_AbsBoldStakingPool *AbsBoldStakingPoolTransactorSession) WithdrawFromPool30fc43ed(amount *big.Int) (*types.Transaction, error) {
	return _AbsBoldStakingPool.Contract.WithdrawFromPool30fc43ed(&_AbsBoldStakingPool.TransactOpts, amount)
}

// AbsBoldStakingPoolStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the AbsBoldStakingPool contract.
type AbsBoldStakingPoolStakeDepositedIterator struct {
	Event *AbsBoldStakingPoolStakeDeposited // Event containing the contract specifics and raw log

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
func (it *AbsBoldStakingPoolStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsBoldStakingPoolStakeDeposited)
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
		it.Event = new(AbsBoldStakingPoolStakeDeposited)
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
func (it *AbsBoldStakingPoolStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsBoldStakingPoolStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsBoldStakingPoolStakeDeposited represents a StakeDeposited event raised by the AbsBoldStakingPool contract.
type AbsBoldStakingPoolStakeDeposited struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed sender, uint256 amount)
func (_AbsBoldStakingPool *AbsBoldStakingPoolFilterer) FilterStakeDeposited(opts *bind.FilterOpts, sender []common.Address) (*AbsBoldStakingPoolStakeDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AbsBoldStakingPool.contract.FilterLogs(opts, "StakeDeposited", senderRule)
	if err != nil {
		return nil, err
	}
	return &AbsBoldStakingPoolStakeDepositedIterator{contract: _AbsBoldStakingPool.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed sender, uint256 amount)
func (_AbsBoldStakingPool *AbsBoldStakingPoolFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *AbsBoldStakingPoolStakeDeposited, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AbsBoldStakingPool.contract.WatchLogs(opts, "StakeDeposited", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsBoldStakingPoolStakeDeposited)
				if err := _AbsBoldStakingPool.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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

// ParseStakeDeposited is a log parse operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed sender, uint256 amount)
func (_AbsBoldStakingPool *AbsBoldStakingPoolFilterer) ParseStakeDeposited(log types.Log) (*AbsBoldStakingPoolStakeDeposited, error) {
	event := new(AbsBoldStakingPoolStakeDeposited)
	if err := _AbsBoldStakingPool.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbsBoldStakingPoolStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the AbsBoldStakingPool contract.
type AbsBoldStakingPoolStakeWithdrawnIterator struct {
	Event *AbsBoldStakingPoolStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *AbsBoldStakingPoolStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbsBoldStakingPoolStakeWithdrawn)
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
		it.Event = new(AbsBoldStakingPoolStakeWithdrawn)
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
func (it *AbsBoldStakingPoolStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbsBoldStakingPoolStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbsBoldStakingPoolStakeWithdrawn represents a StakeWithdrawn event raised by the AbsBoldStakingPool contract.
type AbsBoldStakingPoolStakeWithdrawn struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed sender, uint256 amount)
func (_AbsBoldStakingPool *AbsBoldStakingPoolFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, sender []common.Address) (*AbsBoldStakingPoolStakeWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AbsBoldStakingPool.contract.FilterLogs(opts, "StakeWithdrawn", senderRule)
	if err != nil {
		return nil, err
	}
	return &AbsBoldStakingPoolStakeWithdrawnIterator{contract: _AbsBoldStakingPool.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed sender, uint256 amount)
func (_AbsBoldStakingPool *AbsBoldStakingPoolFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *AbsBoldStakingPoolStakeWithdrawn, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AbsBoldStakingPool.contract.WatchLogs(opts, "StakeWithdrawn", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbsBoldStakingPoolStakeWithdrawn)
				if err := _AbsBoldStakingPool.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed sender, uint256 amount)
func (_AbsBoldStakingPool *AbsBoldStakingPoolFilterer) ParseStakeWithdrawn(log types.Log) (*AbsBoldStakingPoolStakeWithdrawn, error) {
	event := new(AbsBoldStakingPoolStakeWithdrawn)
	if err := _AbsBoldStakingPool.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssertionStakingPoolMetaData contains all meta data concerning the AssertionStakingPool contract.
var AssertionStakingPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"AmountExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAssertionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"assertionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"prevPrevAssertionHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sequencerBatchAcc\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"requiredStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"challengeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nextInboxPosition\",\"type\":\"uint64\"}],\"internalType\":\"structConfigData\",\"name\":\"configData\",\"type\":\"tuple\"}],\"internalType\":\"structBeforeStateData\",\"name\":\"beforeStateData\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"endHistoryRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structAssertionState\",\"name\":\"beforeState\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"endHistoryRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structAssertionState\",\"name\":\"afterState\",\"type\":\"tuple\"}],\"internalType\":\"structAssertionInputs\",\"name\":\"assertionInputs\",\"type\":\"tuple\"}],\"name\":\"createAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositIntoPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"makeStakeWithdrawable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"makeStakeWithdrawableAndWithdrawBackIntoPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStakeBackIntoPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060409080825234610128578181610c4a8038038091610020828561012d565b83398101031261012857602061003582610166565b910151825163051ed6a360e41b81529091906020816004816001600160a01b0386165afa90811561011d576000916100e1575b5060805281156100d05760a05260c05251610acf908161017b82396080518181816101c1015281816104920152818161054401526108e2015260a05181818160b101528181610186015281816109b80152610a3b015260c05181818161033a01526105cc0152f35b8251630b12999960e11b8152600490fd5b906020823d8211610115575b816100fa6020938361012d565b81010312610112575061010c90610166565b38610068565b80fd5b3d91506100ed565b84513d6000823e3d90fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761015057604052565b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101285756fe608060408181526004918236101561001657600080fd5b600092833560e01c9182632113ed21146105b55750816326c0e5c51461059057816330fc43ed1461057357816351ed6a301461052f5781636b74d5151461050e5781637476083b1461041c578163839159711461015e578163930412af146101455781639451944d14610126578163956501bb146100e4575063cb23bcb51461009e57600080fd5b346100e057816003193601126100e057517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b90503461012257602036600319011261012257356001600160a01b0381169081900361011d578282916020945280845220549051908152f35b600080fd5b8280fd5b833461014257806003193601126101425761013f610a24565b80f35b80fd5b833461014257806003193601126101425761013f6109b6565b9190503461012257610260366003190112610122578051636eb1769f60e11b815230838201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038181166024840181905292602092606435917f00000000000000000000000000000000000000000000000000000000000000008416908581604481855afa9081156104125790849392918b916103d9575b5061020f6102409461023b926105ef565b895163095ea7b360e01b8982015293849161022d9160248401610987565b03601f19810184528361063b565b61065e565b833b156103b457845195630a1e65ed60e31b875281818801523560248701526024356044870152604435606487015260848601526084359081168091036103d55760a485015260a4359060018060401b03918281168091036103b45760c486015260c4358281168091036103b45760e48601528360e46101048701376101248661014487015b83600283106103b857505050506101643560038110156103b45790848794939261018490818901526101a49035818901526101c4880137836101e461020488015b6002831061038a57505050505061022435600381101561012257848381936102c4936102449081850152356102648401527f0000000000000000000000000000000000000000000000000000000000000000610284840152306102a48401525af19081156103815750610378575080f35b61013f90610612565b513d84823e3d90fd5b84959650928082939481966103a06001956109a2565b168152019201920190918895949392610307565b8680fd5b80600192876103c6876109a2565b168152019301910190916102c6565b8580fd5b809450878092503d831161040b575b6103f2818361063b565b8101031261040757915183929061020f6101fe565b8980fd5b503d6103e8565b88513d8c823e3d90fd5b8383346100e05760203660031901126100e05782359081156104ff573383528260205280832061044d8382546105ef565b905580516323b872dd60e01b60208201523360248201523060448201526064808201849052815260a081016001600160401b038111828210176104ec5782526104bf907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661065e565b519081527f0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc260203392a280f35b634e487b7160e01b855260418652602485fd5b51631f2a200560e01b81528390fd5b83346101425780600319360112610142576105276109b6565b61013f610a24565b5050346100e057816003193601126100e057517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b8390346100e05760203660031901126100e05761013f903561087b565b5050346100e057816003193601126100e05761013f903383528260205282205461087b565b8490346100e057816003193601126100e0576020907f00000000000000000000000000000000000000000000000000000000000000008152f35b919082018092116105fc57565b634e487b7160e01b600052601160045260246000fd5b6001600160401b03811161062557604052565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761062557604052565b604080516001600160a01b03929092169291906001600160401b0390820181811183821017610625576040526020938483527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858401526000808587829751910182855af1903d1561079c573d92831161078857906106fd939291604051926106f088601f19601f840116018561063b565b83523d868885013e6107a7565b80518061070b575b50505050565b818491810103126100e05782015190811591821503610142575061073157808080610705565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b634e487b7160e01b85526041600452602485fd5b906106fd9392506060915b9192901561080957508151156107bb575090565b3b156107c45790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561081c5750805190602001fd5b6040519062461bcd60e51b82528160208060048301528251908160248401526000935b828510610862575050604492506000838284010152601f80199101168101030190fd5b848101820151868601604401529381019385935061083f565b8015610975576000338152806020526040812054908183116109515782820391821161093d5760409033815280602052205561090f60405163a9059cbb60e01b60208201526108e0816108d2853360248401610987565b03601f19810183528261063b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661065e565b6040519081527f8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc60203392a2565b634e487b7160e01b81526011600452602490fd5b606483836040519163a47b7c6560e01b835233600484015260248301526044820152fd5b604051631f2a200560e01b8152600490fd5b6001600160a01b039091168152602081019190915260400190565b35906001600160401b038216820361011d57565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561011d57600080916004604051809481936357ef4ab960e01b83525af18015610a1857610a0d5750565b610a1690610612565b565b6040513d6000823e3d90fd5b604051636137391960e01b815260208160048160007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af18015610a1857610a725750565b602090813d8111610a92575b610a88818361063b565b8101031261011d57565b503d610a7e56fea264697066735822122047947bd537bb23979399f1a88df55bfb9e9cb3991ec93a4fb5bf4e4eff0bdad364736f6c63430008130033",
}

// AssertionStakingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AssertionStakingPoolMetaData.ABI instead.
var AssertionStakingPoolABI = AssertionStakingPoolMetaData.ABI

// AssertionStakingPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssertionStakingPoolMetaData.Bin instead.
var AssertionStakingPoolBin = AssertionStakingPoolMetaData.Bin

// DeployAssertionStakingPool deploys a new Ethereum contract, binding an instance of AssertionStakingPool to it.
func DeployAssertionStakingPool(auth *bind.TransactOpts, backend bind.ContractBackend, _rollup common.Address, _assertionHash [32]byte) (common.Address, *types.Transaction, *AssertionStakingPool, error) {
	parsed, err := AssertionStakingPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssertionStakingPoolBin), backend, _rollup, _assertionHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssertionStakingPool{AssertionStakingPoolCaller: AssertionStakingPoolCaller{contract: contract}, AssertionStakingPoolTransactor: AssertionStakingPoolTransactor{contract: contract}, AssertionStakingPoolFilterer: AssertionStakingPoolFilterer{contract: contract}}, nil
}

// AssertionStakingPool is an auto generated Go binding around an Ethereum contract.
type AssertionStakingPool struct {
	AssertionStakingPoolCaller     // Read-only binding to the contract
	AssertionStakingPoolTransactor // Write-only binding to the contract
	AssertionStakingPoolFilterer   // Log filterer for contract events
}

// AssertionStakingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssertionStakingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionStakingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssertionStakingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionStakingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssertionStakingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionStakingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssertionStakingPoolSession struct {
	Contract     *AssertionStakingPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AssertionStakingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssertionStakingPoolCallerSession struct {
	Contract *AssertionStakingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AssertionStakingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssertionStakingPoolTransactorSession struct {
	Contract     *AssertionStakingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AssertionStakingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssertionStakingPoolRaw struct {
	Contract *AssertionStakingPool // Generic contract binding to access the raw methods on
}

// AssertionStakingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssertionStakingPoolCallerRaw struct {
	Contract *AssertionStakingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// AssertionStakingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssertionStakingPoolTransactorRaw struct {
	Contract *AssertionStakingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssertionStakingPool creates a new instance of AssertionStakingPool, bound to a specific deployed contract.
func NewAssertionStakingPool(address common.Address, backend bind.ContractBackend) (*AssertionStakingPool, error) {
	contract, err := bindAssertionStakingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPool{AssertionStakingPoolCaller: AssertionStakingPoolCaller{contract: contract}, AssertionStakingPoolTransactor: AssertionStakingPoolTransactor{contract: contract}, AssertionStakingPoolFilterer: AssertionStakingPoolFilterer{contract: contract}}, nil
}

// NewAssertionStakingPoolCaller creates a new read-only instance of AssertionStakingPool, bound to a specific deployed contract.
func NewAssertionStakingPoolCaller(address common.Address, caller bind.ContractCaller) (*AssertionStakingPoolCaller, error) {
	contract, err := bindAssertionStakingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolCaller{contract: contract}, nil
}

// NewAssertionStakingPoolTransactor creates a new write-only instance of AssertionStakingPool, bound to a specific deployed contract.
func NewAssertionStakingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AssertionStakingPoolTransactor, error) {
	contract, err := bindAssertionStakingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolTransactor{contract: contract}, nil
}

// NewAssertionStakingPoolFilterer creates a new log filterer instance of AssertionStakingPool, bound to a specific deployed contract.
func NewAssertionStakingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AssertionStakingPoolFilterer, error) {
	contract, err := bindAssertionStakingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolFilterer{contract: contract}, nil
}

// bindAssertionStakingPool binds a generic wrapper to an already deployed contract.
func bindAssertionStakingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssertionStakingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionStakingPool *AssertionStakingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionStakingPool.Contract.AssertionStakingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionStakingPool *AssertionStakingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.AssertionStakingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionStakingPool *AssertionStakingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.AssertionStakingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionStakingPool *AssertionStakingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionStakingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionStakingPool *AssertionStakingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionStakingPool *AssertionStakingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.contract.Transact(opts, method, params...)
}

// AssertionHash is a free data retrieval call binding the contract method 0x2113ed21.
//
// Solidity: function assertionHash() view returns(bytes32)
func (_AssertionStakingPool *AssertionStakingPoolCaller) AssertionHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssertionStakingPool.contract.Call(opts, &out, "assertionHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AssertionHash is a free data retrieval call binding the contract method 0x2113ed21.
//
// Solidity: function assertionHash() view returns(bytes32)
func (_AssertionStakingPool *AssertionStakingPoolSession) AssertionHash() ([32]byte, error) {
	return _AssertionStakingPool.Contract.AssertionHash(&_AssertionStakingPool.CallOpts)
}

// AssertionHash is a free data retrieval call binding the contract method 0x2113ed21.
//
// Solidity: function assertionHash() view returns(bytes32)
func (_AssertionStakingPool *AssertionStakingPoolCallerSession) AssertionHash() ([32]byte, error) {
	return _AssertionStakingPool.Contract.AssertionHash(&_AssertionStakingPool.CallOpts)
}

// DepositBalance is a free data retrieval call binding the contract method 0x956501bb.
//
// Solidity: function depositBalance(address ) view returns(uint256)
func (_AssertionStakingPool *AssertionStakingPoolCaller) DepositBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AssertionStakingPool.contract.Call(opts, &out, "depositBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositBalance is a free data retrieval call binding the contract method 0x956501bb.
//
// Solidity: function depositBalance(address ) view returns(uint256)
func (_AssertionStakingPool *AssertionStakingPoolSession) DepositBalance(arg0 common.Address) (*big.Int, error) {
	return _AssertionStakingPool.Contract.DepositBalance(&_AssertionStakingPool.CallOpts, arg0)
}

// DepositBalance is a free data retrieval call binding the contract method 0x956501bb.
//
// Solidity: function depositBalance(address ) view returns(uint256)
func (_AssertionStakingPool *AssertionStakingPoolCallerSession) DepositBalance(arg0 common.Address) (*big.Int, error) {
	return _AssertionStakingPool.Contract.DepositBalance(&_AssertionStakingPool.CallOpts, arg0)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_AssertionStakingPool *AssertionStakingPoolCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssertionStakingPool.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_AssertionStakingPool *AssertionStakingPoolSession) Rollup() (common.Address, error) {
	return _AssertionStakingPool.Contract.Rollup(&_AssertionStakingPool.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_AssertionStakingPool *AssertionStakingPoolCallerSession) Rollup() (common.Address, error) {
	return _AssertionStakingPool.Contract.Rollup(&_AssertionStakingPool.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AssertionStakingPool *AssertionStakingPoolCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssertionStakingPool.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AssertionStakingPool *AssertionStakingPoolSession) StakeToken() (common.Address, error) {
	return _AssertionStakingPool.Contract.StakeToken(&_AssertionStakingPool.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_AssertionStakingPool *AssertionStakingPoolCallerSession) StakeToken() (common.Address, error) {
	return _AssertionStakingPool.Contract.StakeToken(&_AssertionStakingPool.CallOpts)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0x83915971.
//
// Solidity: function createAssertion(((bytes32,bytes32,(bytes32,uint256,address,uint64,uint64)),((bytes32[2],uint64[2]),uint8,bytes32),((bytes32[2],uint64[2]),uint8,bytes32)) assertionInputs) returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactor) CreateAssertion(opts *bind.TransactOpts, assertionInputs AssertionInputs) (*types.Transaction, error) {
	return _AssertionStakingPool.contract.Transact(opts, "createAssertion", assertionInputs)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0x83915971.
//
// Solidity: function createAssertion(((bytes32,bytes32,(bytes32,uint256,address,uint64,uint64)),((bytes32[2],uint64[2]),uint8,bytes32),((bytes32[2],uint64[2]),uint8,bytes32)) assertionInputs) returns()
func (_AssertionStakingPool *AssertionStakingPoolSession) CreateAssertion(assertionInputs AssertionInputs) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.CreateAssertion(&_AssertionStakingPool.TransactOpts, assertionInputs)
}

// CreateAssertion is a paid mutator transaction binding the contract method 0x83915971.
//
// Solidity: function createAssertion(((bytes32,bytes32,(bytes32,uint256,address,uint64,uint64)),((bytes32[2],uint64[2]),uint8,bytes32),((bytes32[2],uint64[2]),uint8,bytes32)) assertionInputs) returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactorSession) CreateAssertion(assertionInputs AssertionInputs) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.CreateAssertion(&_AssertionStakingPool.TransactOpts, assertionInputs)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x7476083b.
//
// Solidity: function depositIntoPool(uint256 amount) returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactor) DepositIntoPool(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AssertionStakingPool.contract.Transact(opts, "depositIntoPool", amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x7476083b.
//
// Solidity: function depositIntoPool(uint256 amount) returns()
func (_AssertionStakingPool *AssertionStakingPoolSession) DepositIntoPool(amount *big.Int) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.DepositIntoPool(&_AssertionStakingPool.TransactOpts, amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x7476083b.
//
// Solidity: function depositIntoPool(uint256 amount) returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactorSession) DepositIntoPool(amount *big.Int) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.DepositIntoPool(&_AssertionStakingPool.TransactOpts, amount)
}

// MakeStakeWithdrawable is a paid mutator transaction binding the contract method 0x930412af.
//
// Solidity: function makeStakeWithdrawable() returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactor) MakeStakeWithdrawable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionStakingPool.contract.Transact(opts, "makeStakeWithdrawable")
}

// MakeStakeWithdrawable is a paid mutator transaction binding the contract method 0x930412af.
//
// Solidity: function makeStakeWithdrawable() returns()
func (_AssertionStakingPool *AssertionStakingPoolSession) MakeStakeWithdrawable() (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.MakeStakeWithdrawable(&_AssertionStakingPool.TransactOpts)
}

// MakeStakeWithdrawable is a paid mutator transaction binding the contract method 0x930412af.
//
// Solidity: function makeStakeWithdrawable() returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactorSession) MakeStakeWithdrawable() (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.MakeStakeWithdrawable(&_AssertionStakingPool.TransactOpts)
}

// MakeStakeWithdrawableAndWithdrawBackIntoPool is a paid mutator transaction binding the contract method 0x6b74d515.
//
// Solidity: function makeStakeWithdrawableAndWithdrawBackIntoPool() returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactor) MakeStakeWithdrawableAndWithdrawBackIntoPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionStakingPool.contract.Transact(opts, "makeStakeWithdrawableAndWithdrawBackIntoPool")
}

// MakeStakeWithdrawableAndWithdrawBackIntoPool is a paid mutator transaction binding the contract method 0x6b74d515.
//
// Solidity: function makeStakeWithdrawableAndWithdrawBackIntoPool() returns()
func (_AssertionStakingPool *AssertionStakingPoolSession) MakeStakeWithdrawableAndWithdrawBackIntoPool() (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.MakeStakeWithdrawableAndWithdrawBackIntoPool(&_AssertionStakingPool.TransactOpts)
}

// MakeStakeWithdrawableAndWithdrawBackIntoPool is a paid mutator transaction binding the contract method 0x6b74d515.
//
// Solidity: function makeStakeWithdrawableAndWithdrawBackIntoPool() returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactorSession) MakeStakeWithdrawableAndWithdrawBackIntoPool() (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.MakeStakeWithdrawableAndWithdrawBackIntoPool(&_AssertionStakingPool.TransactOpts)
}

// WithdrawFromPool26c0e5c5 is a paid mutator transaction binding the contract method 0x26c0e5c5.
//
// Solidity: function withdrawFromPool() returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactor) WithdrawFromPool26c0e5c5(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionStakingPool.contract.Transact(opts, "withdrawFromPool")
}

// WithdrawFromPool26c0e5c5 is a paid mutator transaction binding the contract method 0x26c0e5c5.
//
// Solidity: function withdrawFromPool() returns()
func (_AssertionStakingPool *AssertionStakingPoolSession) WithdrawFromPool26c0e5c5() (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.WithdrawFromPool26c0e5c5(&_AssertionStakingPool.TransactOpts)
}

// WithdrawFromPool26c0e5c5 is a paid mutator transaction binding the contract method 0x26c0e5c5.
//
// Solidity: function withdrawFromPool() returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactorSession) WithdrawFromPool26c0e5c5() (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.WithdrawFromPool26c0e5c5(&_AssertionStakingPool.TransactOpts)
}

// WithdrawFromPool30fc43ed is a paid mutator transaction binding the contract method 0x30fc43ed.
//
// Solidity: function withdrawFromPool(uint256 amount) returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactor) WithdrawFromPool30fc43ed(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AssertionStakingPool.contract.Transact(opts, "withdrawFromPool0", amount)
}

// WithdrawFromPool30fc43ed is a paid mutator transaction binding the contract method 0x30fc43ed.
//
// Solidity: function withdrawFromPool(uint256 amount) returns()
func (_AssertionStakingPool *AssertionStakingPoolSession) WithdrawFromPool30fc43ed(amount *big.Int) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.WithdrawFromPool30fc43ed(&_AssertionStakingPool.TransactOpts, amount)
}

// WithdrawFromPool30fc43ed is a paid mutator transaction binding the contract method 0x30fc43ed.
//
// Solidity: function withdrawFromPool(uint256 amount) returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactorSession) WithdrawFromPool30fc43ed(amount *big.Int) (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.WithdrawFromPool30fc43ed(&_AssertionStakingPool.TransactOpts, amount)
}

// WithdrawStakeBackIntoPool is a paid mutator transaction binding the contract method 0x9451944d.
//
// Solidity: function withdrawStakeBackIntoPool() returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactor) WithdrawStakeBackIntoPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionStakingPool.contract.Transact(opts, "withdrawStakeBackIntoPool")
}

// WithdrawStakeBackIntoPool is a paid mutator transaction binding the contract method 0x9451944d.
//
// Solidity: function withdrawStakeBackIntoPool() returns()
func (_AssertionStakingPool *AssertionStakingPoolSession) WithdrawStakeBackIntoPool() (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.WithdrawStakeBackIntoPool(&_AssertionStakingPool.TransactOpts)
}

// WithdrawStakeBackIntoPool is a paid mutator transaction binding the contract method 0x9451944d.
//
// Solidity: function withdrawStakeBackIntoPool() returns()
func (_AssertionStakingPool *AssertionStakingPoolTransactorSession) WithdrawStakeBackIntoPool() (*types.Transaction, error) {
	return _AssertionStakingPool.Contract.WithdrawStakeBackIntoPool(&_AssertionStakingPool.TransactOpts)
}

// AssertionStakingPoolStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the AssertionStakingPool contract.
type AssertionStakingPoolStakeDepositedIterator struct {
	Event *AssertionStakingPoolStakeDeposited // Event containing the contract specifics and raw log

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
func (it *AssertionStakingPoolStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssertionStakingPoolStakeDeposited)
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
		it.Event = new(AssertionStakingPoolStakeDeposited)
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
func (it *AssertionStakingPoolStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssertionStakingPoolStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssertionStakingPoolStakeDeposited represents a StakeDeposited event raised by the AssertionStakingPool contract.
type AssertionStakingPoolStakeDeposited struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed sender, uint256 amount)
func (_AssertionStakingPool *AssertionStakingPoolFilterer) FilterStakeDeposited(opts *bind.FilterOpts, sender []common.Address) (*AssertionStakingPoolStakeDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssertionStakingPool.contract.FilterLogs(opts, "StakeDeposited", senderRule)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolStakeDepositedIterator{contract: _AssertionStakingPool.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed sender, uint256 amount)
func (_AssertionStakingPool *AssertionStakingPoolFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *AssertionStakingPoolStakeDeposited, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssertionStakingPool.contract.WatchLogs(opts, "StakeDeposited", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssertionStakingPoolStakeDeposited)
				if err := _AssertionStakingPool.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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

// ParseStakeDeposited is a log parse operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed sender, uint256 amount)
func (_AssertionStakingPool *AssertionStakingPoolFilterer) ParseStakeDeposited(log types.Log) (*AssertionStakingPoolStakeDeposited, error) {
	event := new(AssertionStakingPoolStakeDeposited)
	if err := _AssertionStakingPool.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssertionStakingPoolStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the AssertionStakingPool contract.
type AssertionStakingPoolStakeWithdrawnIterator struct {
	Event *AssertionStakingPoolStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *AssertionStakingPoolStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssertionStakingPoolStakeWithdrawn)
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
		it.Event = new(AssertionStakingPoolStakeWithdrawn)
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
func (it *AssertionStakingPoolStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssertionStakingPoolStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssertionStakingPoolStakeWithdrawn represents a StakeWithdrawn event raised by the AssertionStakingPool contract.
type AssertionStakingPoolStakeWithdrawn struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed sender, uint256 amount)
func (_AssertionStakingPool *AssertionStakingPoolFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, sender []common.Address) (*AssertionStakingPoolStakeWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssertionStakingPool.contract.FilterLogs(opts, "StakeWithdrawn", senderRule)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolStakeWithdrawnIterator{contract: _AssertionStakingPool.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed sender, uint256 amount)
func (_AssertionStakingPool *AssertionStakingPoolFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *AssertionStakingPoolStakeWithdrawn, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssertionStakingPool.contract.WatchLogs(opts, "StakeWithdrawn", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssertionStakingPoolStakeWithdrawn)
				if err := _AssertionStakingPool.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed sender, uint256 amount)
func (_AssertionStakingPool *AssertionStakingPoolFilterer) ParseStakeWithdrawn(log types.Log) (*AssertionStakingPoolStakeWithdrawn, error) {
	event := new(AssertionStakingPoolStakeWithdrawn)
	if err := _AssertionStakingPool.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssertionStakingPoolCreatorMetaData contains all meta data concerning the AssertionStakingPoolCreator contract.
var AssertionStakingPoolCreatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PoolDoesntExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assertionPool\",\"type\":\"address\"}],\"name\":\"NewAssertionPoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"contractIAssertionStakingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_assertionHash\",\"type\":\"bytes32\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractIAssertionStakingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657610eda908161001c8239f35b600080fdfe60806040818152600436101561001457600080fd5b600091823560e01c9081639b505aa1146100a5575063dc082ad31461003857600080fd5b346100a15760209061009961004c36610152565b8351929161008b90610c4a61006388820187610198565b80865261025b8887013961007d865193849289840161017d565b03601f198101835282610198565b6001600160a01b03926101fc565b169051908152f35b5080fd5b9190503461014e576100b636610152565b9092610c4a808201906001600160401b0382118383101761013a57916100e5848783948a9661025b863961017d565b039082f592831561013057602094507fd628317c6ebae87acc5dbfadeb835cb97692cc6935ea72bf37461e14a0bbee1e8560018060a01b03809616958551938785521692a351908152f35b82513d86823e3d90fd5b634e487b7160e01b87526041600452602487fd5b8280fd5b6040906003190112610178576004356001600160a01b0381168103610178579060243590565b600080fd5b6001600160a01b039091168152602081019190915260400190565b601f909101601f19168101906001600160401b038211908210176101bb57604052565b634e487b7160e01b600052604160045260246000fd5b9081519160005b8381106101e9575050016000815290565b80602080928401015181850152016101d8565b600b9061007d61022260559460405192839161021c6020840180976101d1565b906101d1565b519020604051906040820152600060208201523081520160ff815320803b156102485790565b60405163215db33160e01b8152600490fdfe60e060409080825234610128578181610c4a8038038091610020828561012d565b83398101031261012857602061003582610166565b910151825163051ed6a360e41b81529091906020816004816001600160a01b0386165afa90811561011d576000916100e1575b5060805281156100d05760a05260c05251610acf908161017b82396080518181816101c1015281816104920152818161054401526108e2015260a05181818160b101528181610186015281816109b80152610a3b015260c05181818161033a01526105cc0152f35b8251630b12999960e11b8152600490fd5b906020823d8211610115575b816100fa6020938361012d565b81010312610112575061010c90610166565b38610068565b80fd5b3d91506100ed565b84513d6000823e3d90fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761015057604052565b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101285756fe608060408181526004918236101561001657600080fd5b600092833560e01c9182632113ed21146105b55750816326c0e5c51461059057816330fc43ed1461057357816351ed6a301461052f5781636b74d5151461050e5781637476083b1461041c578163839159711461015e578163930412af146101455781639451944d14610126578163956501bb146100e4575063cb23bcb51461009e57600080fd5b346100e057816003193601126100e057517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b5080fd5b90503461012257602036600319011261012257356001600160a01b0381169081900361011d578282916020945280845220549051908152f35b600080fd5b8280fd5b833461014257806003193601126101425761013f610a24565b80f35b80fd5b833461014257806003193601126101425761013f6109b6565b9190503461012257610260366003190112610122578051636eb1769f60e11b815230838201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038181166024840181905292602092606435917f00000000000000000000000000000000000000000000000000000000000000008416908581604481855afa9081156104125790849392918b916103d9575b5061020f6102409461023b926105ef565b895163095ea7b360e01b8982015293849161022d9160248401610987565b03601f19810184528361063b565b61065e565b833b156103b457845195630a1e65ed60e31b875281818801523560248701526024356044870152604435606487015260848601526084359081168091036103d55760a485015260a4359060018060401b03918281168091036103b45760c486015260c4358281168091036103b45760e48601528360e46101048701376101248661014487015b83600283106103b857505050506101643560038110156103b45790848794939261018490818901526101a49035818901526101c4880137836101e461020488015b6002831061038a57505050505061022435600381101561012257848381936102c4936102449081850152356102648401527f0000000000000000000000000000000000000000000000000000000000000000610284840152306102a48401525af19081156103815750610378575080f35b61013f90610612565b513d84823e3d90fd5b84959650928082939481966103a06001956109a2565b168152019201920190918895949392610307565b8680fd5b80600192876103c6876109a2565b168152019301910190916102c6565b8580fd5b809450878092503d831161040b575b6103f2818361063b565b8101031261040757915183929061020f6101fe565b8980fd5b503d6103e8565b88513d8c823e3d90fd5b8383346100e05760203660031901126100e05782359081156104ff573383528260205280832061044d8382546105ef565b905580516323b872dd60e01b60208201523360248201523060448201526064808201849052815260a081016001600160401b038111828210176104ec5782526104bf907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661065e565b519081527f0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc260203392a280f35b634e487b7160e01b855260418652602485fd5b51631f2a200560e01b81528390fd5b83346101425780600319360112610142576105276109b6565b61013f610a24565b5050346100e057816003193601126100e057517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b8390346100e05760203660031901126100e05761013f903561087b565b5050346100e057816003193601126100e05761013f903383528260205282205461087b565b8490346100e057816003193601126100e0576020907f00000000000000000000000000000000000000000000000000000000000000008152f35b919082018092116105fc57565b634e487b7160e01b600052601160045260246000fd5b6001600160401b03811161062557604052565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761062557604052565b604080516001600160a01b03929092169291906001600160401b0390820181811183821017610625576040526020938483527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858401526000808587829751910182855af1903d1561079c573d92831161078857906106fd939291604051926106f088601f19601f840116018561063b565b83523d868885013e6107a7565b80518061070b575b50505050565b818491810103126100e05782015190811591821503610142575061073157808080610705565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b634e487b7160e01b85526041600452602485fd5b906106fd9392506060915b9192901561080957508151156107bb575090565b3b156107c45790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561081c5750805190602001fd5b6040519062461bcd60e51b82528160208060048301528251908160248401526000935b828510610862575050604492506000838284010152601f80199101168101030190fd5b848101820151868601604401529381019385935061083f565b8015610975576000338152806020526040812054908183116109515782820391821161093d5760409033815280602052205561090f60405163a9059cbb60e01b60208201526108e0816108d2853360248401610987565b03601f19810183528261063b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661065e565b6040519081527f8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc60203392a2565b634e487b7160e01b81526011600452602490fd5b606483836040519163a47b7c6560e01b835233600484015260248301526044820152fd5b604051631f2a200560e01b8152600490fd5b6001600160a01b039091168152602081019190915260400190565b35906001600160401b038216820361011d57565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316803b1561011d57600080916004604051809481936357ef4ab960e01b83525af18015610a1857610a0d5750565b610a1690610612565b565b6040513d6000823e3d90fd5b604051636137391960e01b815260208160048160007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165af18015610a1857610a725750565b602090813d8111610a92575b610a88818361063b565b8101031261011d57565b503d610a7e56fea264697066735822122047947bd537bb23979399f1a88df55bfb9e9cb3991ec93a4fb5bf4e4eff0bdad364736f6c63430008130033a26469706673582212204fdc0fa0bbe3d90afb079eae05660298a0eb0b66bb34db5997a9f43516bb5d9964736f6c63430008130033",
}

// AssertionStakingPoolCreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AssertionStakingPoolCreatorMetaData.ABI instead.
var AssertionStakingPoolCreatorABI = AssertionStakingPoolCreatorMetaData.ABI

// AssertionStakingPoolCreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssertionStakingPoolCreatorMetaData.Bin instead.
var AssertionStakingPoolCreatorBin = AssertionStakingPoolCreatorMetaData.Bin

// DeployAssertionStakingPoolCreator deploys a new Ethereum contract, binding an instance of AssertionStakingPoolCreator to it.
func DeployAssertionStakingPoolCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AssertionStakingPoolCreator, error) {
	parsed, err := AssertionStakingPoolCreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssertionStakingPoolCreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssertionStakingPoolCreator{AssertionStakingPoolCreatorCaller: AssertionStakingPoolCreatorCaller{contract: contract}, AssertionStakingPoolCreatorTransactor: AssertionStakingPoolCreatorTransactor{contract: contract}, AssertionStakingPoolCreatorFilterer: AssertionStakingPoolCreatorFilterer{contract: contract}}, nil
}

// AssertionStakingPoolCreator is an auto generated Go binding around an Ethereum contract.
type AssertionStakingPoolCreator struct {
	AssertionStakingPoolCreatorCaller     // Read-only binding to the contract
	AssertionStakingPoolCreatorTransactor // Write-only binding to the contract
	AssertionStakingPoolCreatorFilterer   // Log filterer for contract events
}

// AssertionStakingPoolCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssertionStakingPoolCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionStakingPoolCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssertionStakingPoolCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionStakingPoolCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssertionStakingPoolCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionStakingPoolCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssertionStakingPoolCreatorSession struct {
	Contract     *AssertionStakingPoolCreator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AssertionStakingPoolCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssertionStakingPoolCreatorCallerSession struct {
	Contract *AssertionStakingPoolCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// AssertionStakingPoolCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssertionStakingPoolCreatorTransactorSession struct {
	Contract     *AssertionStakingPoolCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// AssertionStakingPoolCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssertionStakingPoolCreatorRaw struct {
	Contract *AssertionStakingPoolCreator // Generic contract binding to access the raw methods on
}

// AssertionStakingPoolCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssertionStakingPoolCreatorCallerRaw struct {
	Contract *AssertionStakingPoolCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// AssertionStakingPoolCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssertionStakingPoolCreatorTransactorRaw struct {
	Contract *AssertionStakingPoolCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssertionStakingPoolCreator creates a new instance of AssertionStakingPoolCreator, bound to a specific deployed contract.
func NewAssertionStakingPoolCreator(address common.Address, backend bind.ContractBackend) (*AssertionStakingPoolCreator, error) {
	contract, err := bindAssertionStakingPoolCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolCreator{AssertionStakingPoolCreatorCaller: AssertionStakingPoolCreatorCaller{contract: contract}, AssertionStakingPoolCreatorTransactor: AssertionStakingPoolCreatorTransactor{contract: contract}, AssertionStakingPoolCreatorFilterer: AssertionStakingPoolCreatorFilterer{contract: contract}}, nil
}

// NewAssertionStakingPoolCreatorCaller creates a new read-only instance of AssertionStakingPoolCreator, bound to a specific deployed contract.
func NewAssertionStakingPoolCreatorCaller(address common.Address, caller bind.ContractCaller) (*AssertionStakingPoolCreatorCaller, error) {
	contract, err := bindAssertionStakingPoolCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolCreatorCaller{contract: contract}, nil
}

// NewAssertionStakingPoolCreatorTransactor creates a new write-only instance of AssertionStakingPoolCreator, bound to a specific deployed contract.
func NewAssertionStakingPoolCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AssertionStakingPoolCreatorTransactor, error) {
	contract, err := bindAssertionStakingPoolCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolCreatorTransactor{contract: contract}, nil
}

// NewAssertionStakingPoolCreatorFilterer creates a new log filterer instance of AssertionStakingPoolCreator, bound to a specific deployed contract.
func NewAssertionStakingPoolCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AssertionStakingPoolCreatorFilterer, error) {
	contract, err := bindAssertionStakingPoolCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolCreatorFilterer{contract: contract}, nil
}

// bindAssertionStakingPoolCreator binds a generic wrapper to an already deployed contract.
func bindAssertionStakingPoolCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssertionStakingPoolCreatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionStakingPoolCreator.Contract.AssertionStakingPoolCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionStakingPoolCreator.Contract.AssertionStakingPoolCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionStakingPoolCreator.Contract.AssertionStakingPoolCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionStakingPoolCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionStakingPoolCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionStakingPoolCreator.Contract.contract.Transact(opts, method, params...)
}

// GetPool is a free data retrieval call binding the contract method 0xdc082ad3.
//
// Solidity: function getPool(address _rollup, bytes32 _assertionHash) view returns(address)
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorCaller) GetPool(opts *bind.CallOpts, _rollup common.Address, _assertionHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AssertionStakingPoolCreator.contract.Call(opts, &out, "getPool", _rollup, _assertionHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0xdc082ad3.
//
// Solidity: function getPool(address _rollup, bytes32 _assertionHash) view returns(address)
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorSession) GetPool(_rollup common.Address, _assertionHash [32]byte) (common.Address, error) {
	return _AssertionStakingPoolCreator.Contract.GetPool(&_AssertionStakingPoolCreator.CallOpts, _rollup, _assertionHash)
}

// GetPool is a free data retrieval call binding the contract method 0xdc082ad3.
//
// Solidity: function getPool(address _rollup, bytes32 _assertionHash) view returns(address)
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorCallerSession) GetPool(_rollup common.Address, _assertionHash [32]byte) (common.Address, error) {
	return _AssertionStakingPoolCreator.Contract.GetPool(&_AssertionStakingPoolCreator.CallOpts, _rollup, _assertionHash)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9b505aa1.
//
// Solidity: function createPool(address _rollup, bytes32 _assertionHash) returns(address)
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorTransactor) CreatePool(opts *bind.TransactOpts, _rollup common.Address, _assertionHash [32]byte) (*types.Transaction, error) {
	return _AssertionStakingPoolCreator.contract.Transact(opts, "createPool", _rollup, _assertionHash)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9b505aa1.
//
// Solidity: function createPool(address _rollup, bytes32 _assertionHash) returns(address)
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorSession) CreatePool(_rollup common.Address, _assertionHash [32]byte) (*types.Transaction, error) {
	return _AssertionStakingPoolCreator.Contract.CreatePool(&_AssertionStakingPoolCreator.TransactOpts, _rollup, _assertionHash)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9b505aa1.
//
// Solidity: function createPool(address _rollup, bytes32 _assertionHash) returns(address)
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorTransactorSession) CreatePool(_rollup common.Address, _assertionHash [32]byte) (*types.Transaction, error) {
	return _AssertionStakingPoolCreator.Contract.CreatePool(&_AssertionStakingPoolCreator.TransactOpts, _rollup, _assertionHash)
}

// AssertionStakingPoolCreatorNewAssertionPoolCreatedIterator is returned from FilterNewAssertionPoolCreated and is used to iterate over the raw logs and unpacked data for NewAssertionPoolCreated events raised by the AssertionStakingPoolCreator contract.
type AssertionStakingPoolCreatorNewAssertionPoolCreatedIterator struct {
	Event *AssertionStakingPoolCreatorNewAssertionPoolCreated // Event containing the contract specifics and raw log

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
func (it *AssertionStakingPoolCreatorNewAssertionPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssertionStakingPoolCreatorNewAssertionPoolCreated)
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
		it.Event = new(AssertionStakingPoolCreatorNewAssertionPoolCreated)
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
func (it *AssertionStakingPoolCreatorNewAssertionPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssertionStakingPoolCreatorNewAssertionPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssertionStakingPoolCreatorNewAssertionPoolCreated represents a NewAssertionPoolCreated event raised by the AssertionStakingPoolCreator contract.
type AssertionStakingPoolCreatorNewAssertionPoolCreated struct {
	Rollup        common.Address
	AssertionHash [32]byte
	AssertionPool common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewAssertionPoolCreated is a free log retrieval operation binding the contract event 0xd628317c6ebae87acc5dbfadeb835cb97692cc6935ea72bf37461e14a0bbee1e.
//
// Solidity: event NewAssertionPoolCreated(address indexed rollup, bytes32 indexed _assertionHash, address assertionPool)
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorFilterer) FilterNewAssertionPoolCreated(opts *bind.FilterOpts, rollup []common.Address, _assertionHash [][32]byte) (*AssertionStakingPoolCreatorNewAssertionPoolCreatedIterator, error) {

	var rollupRule []interface{}
	for _, rollupItem := range rollup {
		rollupRule = append(rollupRule, rollupItem)
	}
	var _assertionHashRule []interface{}
	for _, _assertionHashItem := range _assertionHash {
		_assertionHashRule = append(_assertionHashRule, _assertionHashItem)
	}

	logs, sub, err := _AssertionStakingPoolCreator.contract.FilterLogs(opts, "NewAssertionPoolCreated", rollupRule, _assertionHashRule)
	if err != nil {
		return nil, err
	}
	return &AssertionStakingPoolCreatorNewAssertionPoolCreatedIterator{contract: _AssertionStakingPoolCreator.contract, event: "NewAssertionPoolCreated", logs: logs, sub: sub}, nil
}

// WatchNewAssertionPoolCreated is a free log subscription operation binding the contract event 0xd628317c6ebae87acc5dbfadeb835cb97692cc6935ea72bf37461e14a0bbee1e.
//
// Solidity: event NewAssertionPoolCreated(address indexed rollup, bytes32 indexed _assertionHash, address assertionPool)
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorFilterer) WatchNewAssertionPoolCreated(opts *bind.WatchOpts, sink chan<- *AssertionStakingPoolCreatorNewAssertionPoolCreated, rollup []common.Address, _assertionHash [][32]byte) (event.Subscription, error) {

	var rollupRule []interface{}
	for _, rollupItem := range rollup {
		rollupRule = append(rollupRule, rollupItem)
	}
	var _assertionHashRule []interface{}
	for _, _assertionHashItem := range _assertionHash {
		_assertionHashRule = append(_assertionHashRule, _assertionHashItem)
	}

	logs, sub, err := _AssertionStakingPoolCreator.contract.WatchLogs(opts, "NewAssertionPoolCreated", rollupRule, _assertionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssertionStakingPoolCreatorNewAssertionPoolCreated)
				if err := _AssertionStakingPoolCreator.contract.UnpackLog(event, "NewAssertionPoolCreated", log); err != nil {
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

// ParseNewAssertionPoolCreated is a log parse operation binding the contract event 0xd628317c6ebae87acc5dbfadeb835cb97692cc6935ea72bf37461e14a0bbee1e.
//
// Solidity: event NewAssertionPoolCreated(address indexed rollup, bytes32 indexed _assertionHash, address assertionPool)
func (_AssertionStakingPoolCreator *AssertionStakingPoolCreatorFilterer) ParseNewAssertionPoolCreated(log types.Log) (*AssertionStakingPoolCreatorNewAssertionPoolCreated, error) {
	event := new(AssertionStakingPoolCreatorNewAssertionPoolCreated)
	if err := _AssertionStakingPoolCreator.contract.UnpackLog(event, "NewAssertionPoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeStakingPoolMetaData contains all meta data concerning the EdgeStakingPool contract.
var EdgeStakingPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_challengeManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_edgeId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"AmountExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyEdgeId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"}],\"name\":\"IncorrectEdgeId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"challengeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"endHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"endHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structCreateEdgeArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"createEdge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositIntoPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"edgeId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060409080825234610120578181610b2780380380916100208285610125565b8339810103126101205780516001600160a01b03808216928383036101205760208060049201519486519283809263051ed6a360e41b82525afa908115610115576000916100d4575b501660805281156100c45760a05260c052516109c8908161015f82396080518181816101180152818161044f015281816104e0015261087a015260a05181818160b10152610568015260c05181818161024801526103780152f35b8251620d29f560e71b8152600490fd5b6020813d821161010d575b816100ec60209383610125565b810103126101095751908282168203610106575038610069565b80fd5b5080fd5b3d91506100df565b85513d6000823e3d90fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761014857604052565b634e487b7160e01b600052604160045260246000fdfe608060408181526004908136101561001657600080fd5b600092833560e01c908163023a96fe146105545750806326c0e5c51461052f57806330fc43ed1461050f57806351ed6a30146104cb5780637476083b146103d8578063956501bb1461039f5780639cfa2a2a146103605763bd3eec7d1461007c57600080fd5b3461035857600319602036820181136102b2578335916001600160401b03831161035c5760c083860191843603011261035c577f000000000000000000000000000000000000000000000000000000000000000060018060a01b0393878583169184359360ff8516809503610358578851630e0da79d60e11b81528a810186905260249888828b81895afa91821561034e57859261031f575b507f000000000000000000000000000000000000000000000000000000000000000016908a51636eb1769f60e11b8152308d820152868b8201528981604481865afa908115610315578c8e9593889795938e938e9c9b9a916102c7575b506101bc936101a961023297946101918f9c9b9a98956101b795610597565b905163095ea7b360e01b8d820152958693840161091f565b03601f1981018452836105ba565b6105f3565b6102208c51998a98899788966305fae14160e01b88528701528d8601528c8301356044860152604483013560648601526064830135608486015260a4610218610208608486018461093a565b60c0848a015260e4890191610971565b93019061093a565b8483036023190160c486015290610971565b03925af19182156102bd57869261028b575b50507f000000000000000000000000000000000000000000000000000000000000000092838203610273578580f35b516375c0811b60e01b81529384015282015260449150fd5b90809250813d83116102b6575b6102a281836105ba565b810103126102b257513880610244565b8480fd5b503d610298565b84513d88823e3d90fd5b9598975050998092508491509492943d831161030e575b6102e881836105ba565b8101031261030a579051899793948d949293909290918c918e916101bc610172565b8380fd5b503d6102de565b8c513d88823e3d90fd5b9091508881813d8311610347575b61033781836105ba565b810103126102b257519038610115565b503d61032d565b8b513d87823e3d90fd5b8280fd5b8580fd5b83823461039b578160031936011261039b57602090517f00000000000000000000000000000000000000000000000000000000000000008152f35b5080fd5b50903461035857602036600319011261035857356001600160a01b03811690819003610358578282916020945280845220549051908152f35b5082903461039b57602036600319011261039b5782359081156104bc573383528260205280832061040a838254610597565b905580516323b872dd60e01b60208201523360248201523060448201526064808201849052815260a081016001600160401b038111828210176104a957825261047c907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105f3565b519081527f0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc260203392a280f35b634e487b7160e01b855260418652602485fd5b51631f2a200560e01b81528390fd5b83823461039b578160031936011261039b57517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50503461039b57602036600319011261039b5761052c9035610813565b80f35b83823461039b578160031936011261039b5761052c9033835282602052822054610813565b84903461039b578160031936011261039b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b919082018092116105a457565b634e487b7160e01b600052601160045260246000fd5b601f909101601f19168101906001600160401b038211908210176105dd57604052565b634e487b7160e01b600052604160045260246000fd5b604080516001600160a01b03929092169291906001600160401b03908201818111838210176105dd576040526020938483527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858401526000808587829751910182855af1903d15610734573d92831161072057906106929392916040519261068588601f19601f84011601856105ba565b83523d868885013e61073f565b8051806106a0575b50505050565b8184918101031261039b578201519081159182150361071d57506106c65780808061069a565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b80fd5b634e487b7160e01b85526041600452602485fd5b906106929392506060915b919290156107a15750815115610753575090565b3b1561075c5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156107b45750805190602001fd5b6040519062461bcd60e51b82528160208060048301528251908160248401526000935b8285106107fa575050604492506000838284010152601f80199101168101030190fd5b84810182015186860160440152938101938593506107d7565b801561090d576000338152806020526040812054908183116108e9578282039182116108d5576040903381528060205220556108a760405163a9059cbb60e01b60208201526108788161086a85336024840161091f565b03601f1981018352826105ba565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105f3565b6040519081527f8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc60203392a2565b634e487b7160e01b81526011600452602490fd5b606483836040519163a47b7c6560e01b835233600484015260248301526044820152fd5b604051631f2a200560e01b8152600490fd5b6001600160a01b039091168152602081019190915260400190565b9035601e198236030181121561096c570160208101919035906001600160401b03821161096c57813603831361096c57565b600080fd5b908060209392818452848401376000828201840152601f01601f191601019056fea26469706673582212209632da7ad69b2f1191cdaa449622d3d3c4563257ac2720527080276d4f004b9364736f6c63430008130033",
}

// EdgeStakingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use EdgeStakingPoolMetaData.ABI instead.
var EdgeStakingPoolABI = EdgeStakingPoolMetaData.ABI

// EdgeStakingPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EdgeStakingPoolMetaData.Bin instead.
var EdgeStakingPoolBin = EdgeStakingPoolMetaData.Bin

// DeployEdgeStakingPool deploys a new Ethereum contract, binding an instance of EdgeStakingPool to it.
func DeployEdgeStakingPool(auth *bind.TransactOpts, backend bind.ContractBackend, _challengeManager common.Address, _edgeId [32]byte) (common.Address, *types.Transaction, *EdgeStakingPool, error) {
	parsed, err := EdgeStakingPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EdgeStakingPoolBin), backend, _challengeManager, _edgeId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EdgeStakingPool{EdgeStakingPoolCaller: EdgeStakingPoolCaller{contract: contract}, EdgeStakingPoolTransactor: EdgeStakingPoolTransactor{contract: contract}, EdgeStakingPoolFilterer: EdgeStakingPoolFilterer{contract: contract}}, nil
}

// EdgeStakingPool is an auto generated Go binding around an Ethereum contract.
type EdgeStakingPool struct {
	EdgeStakingPoolCaller     // Read-only binding to the contract
	EdgeStakingPoolTransactor // Write-only binding to the contract
	EdgeStakingPoolFilterer   // Log filterer for contract events
}

// EdgeStakingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type EdgeStakingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeStakingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EdgeStakingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeStakingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EdgeStakingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeStakingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EdgeStakingPoolSession struct {
	Contract     *EdgeStakingPool  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EdgeStakingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EdgeStakingPoolCallerSession struct {
	Contract *EdgeStakingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EdgeStakingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EdgeStakingPoolTransactorSession struct {
	Contract     *EdgeStakingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EdgeStakingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type EdgeStakingPoolRaw struct {
	Contract *EdgeStakingPool // Generic contract binding to access the raw methods on
}

// EdgeStakingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EdgeStakingPoolCallerRaw struct {
	Contract *EdgeStakingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// EdgeStakingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EdgeStakingPoolTransactorRaw struct {
	Contract *EdgeStakingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEdgeStakingPool creates a new instance of EdgeStakingPool, bound to a specific deployed contract.
func NewEdgeStakingPool(address common.Address, backend bind.ContractBackend) (*EdgeStakingPool, error) {
	contract, err := bindEdgeStakingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPool{EdgeStakingPoolCaller: EdgeStakingPoolCaller{contract: contract}, EdgeStakingPoolTransactor: EdgeStakingPoolTransactor{contract: contract}, EdgeStakingPoolFilterer: EdgeStakingPoolFilterer{contract: contract}}, nil
}

// NewEdgeStakingPoolCaller creates a new read-only instance of EdgeStakingPool, bound to a specific deployed contract.
func NewEdgeStakingPoolCaller(address common.Address, caller bind.ContractCaller) (*EdgeStakingPoolCaller, error) {
	contract, err := bindEdgeStakingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolCaller{contract: contract}, nil
}

// NewEdgeStakingPoolTransactor creates a new write-only instance of EdgeStakingPool, bound to a specific deployed contract.
func NewEdgeStakingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*EdgeStakingPoolTransactor, error) {
	contract, err := bindEdgeStakingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolTransactor{contract: contract}, nil
}

// NewEdgeStakingPoolFilterer creates a new log filterer instance of EdgeStakingPool, bound to a specific deployed contract.
func NewEdgeStakingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*EdgeStakingPoolFilterer, error) {
	contract, err := bindEdgeStakingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolFilterer{contract: contract}, nil
}

// bindEdgeStakingPool binds a generic wrapper to an already deployed contract.
func bindEdgeStakingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EdgeStakingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EdgeStakingPool *EdgeStakingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EdgeStakingPool.Contract.EdgeStakingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EdgeStakingPool *EdgeStakingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.EdgeStakingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EdgeStakingPool *EdgeStakingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.EdgeStakingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EdgeStakingPool *EdgeStakingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EdgeStakingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EdgeStakingPool *EdgeStakingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EdgeStakingPool *EdgeStakingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.contract.Transact(opts, method, params...)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() view returns(address)
func (_EdgeStakingPool *EdgeStakingPoolCaller) ChallengeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EdgeStakingPool.contract.Call(opts, &out, "challengeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() view returns(address)
func (_EdgeStakingPool *EdgeStakingPoolSession) ChallengeManager() (common.Address, error) {
	return _EdgeStakingPool.Contract.ChallengeManager(&_EdgeStakingPool.CallOpts)
}

// ChallengeManager is a free data retrieval call binding the contract method 0x023a96fe.
//
// Solidity: function challengeManager() view returns(address)
func (_EdgeStakingPool *EdgeStakingPoolCallerSession) ChallengeManager() (common.Address, error) {
	return _EdgeStakingPool.Contract.ChallengeManager(&_EdgeStakingPool.CallOpts)
}

// DepositBalance is a free data retrieval call binding the contract method 0x956501bb.
//
// Solidity: function depositBalance(address ) view returns(uint256)
func (_EdgeStakingPool *EdgeStakingPoolCaller) DepositBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EdgeStakingPool.contract.Call(opts, &out, "depositBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositBalance is a free data retrieval call binding the contract method 0x956501bb.
//
// Solidity: function depositBalance(address ) view returns(uint256)
func (_EdgeStakingPool *EdgeStakingPoolSession) DepositBalance(arg0 common.Address) (*big.Int, error) {
	return _EdgeStakingPool.Contract.DepositBalance(&_EdgeStakingPool.CallOpts, arg0)
}

// DepositBalance is a free data retrieval call binding the contract method 0x956501bb.
//
// Solidity: function depositBalance(address ) view returns(uint256)
func (_EdgeStakingPool *EdgeStakingPoolCallerSession) DepositBalance(arg0 common.Address) (*big.Int, error) {
	return _EdgeStakingPool.Contract.DepositBalance(&_EdgeStakingPool.CallOpts, arg0)
}

// EdgeId is a free data retrieval call binding the contract method 0x9cfa2a2a.
//
// Solidity: function edgeId() view returns(bytes32)
func (_EdgeStakingPool *EdgeStakingPoolCaller) EdgeId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EdgeStakingPool.contract.Call(opts, &out, "edgeId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EdgeId is a free data retrieval call binding the contract method 0x9cfa2a2a.
//
// Solidity: function edgeId() view returns(bytes32)
func (_EdgeStakingPool *EdgeStakingPoolSession) EdgeId() ([32]byte, error) {
	return _EdgeStakingPool.Contract.EdgeId(&_EdgeStakingPool.CallOpts)
}

// EdgeId is a free data retrieval call binding the contract method 0x9cfa2a2a.
//
// Solidity: function edgeId() view returns(bytes32)
func (_EdgeStakingPool *EdgeStakingPoolCallerSession) EdgeId() ([32]byte, error) {
	return _EdgeStakingPool.Contract.EdgeId(&_EdgeStakingPool.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_EdgeStakingPool *EdgeStakingPoolCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EdgeStakingPool.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_EdgeStakingPool *EdgeStakingPoolSession) StakeToken() (common.Address, error) {
	return _EdgeStakingPool.Contract.StakeToken(&_EdgeStakingPool.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_EdgeStakingPool *EdgeStakingPoolCallerSession) StakeToken() (common.Address, error) {
	return _EdgeStakingPool.Contract.StakeToken(&_EdgeStakingPool.CallOpts)
}

// CreateEdge is a paid mutator transaction binding the contract method 0xbd3eec7d.
//
// Solidity: function createEdge((uint8,bytes32,uint256,bytes32,bytes,bytes) args) returns()
func (_EdgeStakingPool *EdgeStakingPoolTransactor) CreateEdge(opts *bind.TransactOpts, args CreateEdgeArgs) (*types.Transaction, error) {
	return _EdgeStakingPool.contract.Transact(opts, "createEdge", args)
}

// CreateEdge is a paid mutator transaction binding the contract method 0xbd3eec7d.
//
// Solidity: function createEdge((uint8,bytes32,uint256,bytes32,bytes,bytes) args) returns()
func (_EdgeStakingPool *EdgeStakingPoolSession) CreateEdge(args CreateEdgeArgs) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.CreateEdge(&_EdgeStakingPool.TransactOpts, args)
}

// CreateEdge is a paid mutator transaction binding the contract method 0xbd3eec7d.
//
// Solidity: function createEdge((uint8,bytes32,uint256,bytes32,bytes,bytes) args) returns()
func (_EdgeStakingPool *EdgeStakingPoolTransactorSession) CreateEdge(args CreateEdgeArgs) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.CreateEdge(&_EdgeStakingPool.TransactOpts, args)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x7476083b.
//
// Solidity: function depositIntoPool(uint256 amount) returns()
func (_EdgeStakingPool *EdgeStakingPoolTransactor) DepositIntoPool(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EdgeStakingPool.contract.Transact(opts, "depositIntoPool", amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x7476083b.
//
// Solidity: function depositIntoPool(uint256 amount) returns()
func (_EdgeStakingPool *EdgeStakingPoolSession) DepositIntoPool(amount *big.Int) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.DepositIntoPool(&_EdgeStakingPool.TransactOpts, amount)
}

// DepositIntoPool is a paid mutator transaction binding the contract method 0x7476083b.
//
// Solidity: function depositIntoPool(uint256 amount) returns()
func (_EdgeStakingPool *EdgeStakingPoolTransactorSession) DepositIntoPool(amount *big.Int) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.DepositIntoPool(&_EdgeStakingPool.TransactOpts, amount)
}

// WithdrawFromPool26c0e5c5 is a paid mutator transaction binding the contract method 0x26c0e5c5.
//
// Solidity: function withdrawFromPool() returns()
func (_EdgeStakingPool *EdgeStakingPoolTransactor) WithdrawFromPool26c0e5c5(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EdgeStakingPool.contract.Transact(opts, "withdrawFromPool")
}

// WithdrawFromPool26c0e5c5 is a paid mutator transaction binding the contract method 0x26c0e5c5.
//
// Solidity: function withdrawFromPool() returns()
func (_EdgeStakingPool *EdgeStakingPoolSession) WithdrawFromPool26c0e5c5() (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.WithdrawFromPool26c0e5c5(&_EdgeStakingPool.TransactOpts)
}

// WithdrawFromPool26c0e5c5 is a paid mutator transaction binding the contract method 0x26c0e5c5.
//
// Solidity: function withdrawFromPool() returns()
func (_EdgeStakingPool *EdgeStakingPoolTransactorSession) WithdrawFromPool26c0e5c5() (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.WithdrawFromPool26c0e5c5(&_EdgeStakingPool.TransactOpts)
}

// WithdrawFromPool30fc43ed is a paid mutator transaction binding the contract method 0x30fc43ed.
//
// Solidity: function withdrawFromPool(uint256 amount) returns()
func (_EdgeStakingPool *EdgeStakingPoolTransactor) WithdrawFromPool30fc43ed(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EdgeStakingPool.contract.Transact(opts, "withdrawFromPool0", amount)
}

// WithdrawFromPool30fc43ed is a paid mutator transaction binding the contract method 0x30fc43ed.
//
// Solidity: function withdrawFromPool(uint256 amount) returns()
func (_EdgeStakingPool *EdgeStakingPoolSession) WithdrawFromPool30fc43ed(amount *big.Int) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.WithdrawFromPool30fc43ed(&_EdgeStakingPool.TransactOpts, amount)
}

// WithdrawFromPool30fc43ed is a paid mutator transaction binding the contract method 0x30fc43ed.
//
// Solidity: function withdrawFromPool(uint256 amount) returns()
func (_EdgeStakingPool *EdgeStakingPoolTransactorSession) WithdrawFromPool30fc43ed(amount *big.Int) (*types.Transaction, error) {
	return _EdgeStakingPool.Contract.WithdrawFromPool30fc43ed(&_EdgeStakingPool.TransactOpts, amount)
}

// EdgeStakingPoolStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the EdgeStakingPool contract.
type EdgeStakingPoolStakeDepositedIterator struct {
	Event *EdgeStakingPoolStakeDeposited // Event containing the contract specifics and raw log

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
func (it *EdgeStakingPoolStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeStakingPoolStakeDeposited)
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
		it.Event = new(EdgeStakingPoolStakeDeposited)
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
func (it *EdgeStakingPoolStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeStakingPoolStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeStakingPoolStakeDeposited represents a StakeDeposited event raised by the EdgeStakingPool contract.
type EdgeStakingPoolStakeDeposited struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed sender, uint256 amount)
func (_EdgeStakingPool *EdgeStakingPoolFilterer) FilterStakeDeposited(opts *bind.FilterOpts, sender []common.Address) (*EdgeStakingPoolStakeDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EdgeStakingPool.contract.FilterLogs(opts, "StakeDeposited", senderRule)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolStakeDepositedIterator{contract: _EdgeStakingPool.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed sender, uint256 amount)
func (_EdgeStakingPool *EdgeStakingPoolFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *EdgeStakingPoolStakeDeposited, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EdgeStakingPool.contract.WatchLogs(opts, "StakeDeposited", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeStakingPoolStakeDeposited)
				if err := _EdgeStakingPool.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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

// ParseStakeDeposited is a log parse operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed sender, uint256 amount)
func (_EdgeStakingPool *EdgeStakingPoolFilterer) ParseStakeDeposited(log types.Log) (*EdgeStakingPoolStakeDeposited, error) {
	event := new(EdgeStakingPoolStakeDeposited)
	if err := _EdgeStakingPool.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeStakingPoolStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the EdgeStakingPool contract.
type EdgeStakingPoolStakeWithdrawnIterator struct {
	Event *EdgeStakingPoolStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *EdgeStakingPoolStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeStakingPoolStakeWithdrawn)
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
		it.Event = new(EdgeStakingPoolStakeWithdrawn)
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
func (it *EdgeStakingPoolStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeStakingPoolStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeStakingPoolStakeWithdrawn represents a StakeWithdrawn event raised by the EdgeStakingPool contract.
type EdgeStakingPoolStakeWithdrawn struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed sender, uint256 amount)
func (_EdgeStakingPool *EdgeStakingPoolFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, sender []common.Address) (*EdgeStakingPoolStakeWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EdgeStakingPool.contract.FilterLogs(opts, "StakeWithdrawn", senderRule)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolStakeWithdrawnIterator{contract: _EdgeStakingPool.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed sender, uint256 amount)
func (_EdgeStakingPool *EdgeStakingPoolFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *EdgeStakingPoolStakeWithdrawn, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EdgeStakingPool.contract.WatchLogs(opts, "StakeWithdrawn", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeStakingPoolStakeWithdrawn)
				if err := _EdgeStakingPool.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed sender, uint256 amount)
func (_EdgeStakingPool *EdgeStakingPoolFilterer) ParseStakeWithdrawn(log types.Log) (*EdgeStakingPoolStakeWithdrawn, error) {
	event := new(EdgeStakingPoolStakeWithdrawn)
	if err := _EdgeStakingPool.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeStakingPoolCreatorMetaData contains all meta data concerning the EdgeStakingPoolCreator contract.
var EdgeStakingPoolCreatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PoolDoesntExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"NewEdgeStakingPoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challengeManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"contractIEdgeStakingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challengeManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"edgeId\",\"type\":\"bytes32\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractIEdgeStakingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657610daf908161001c8239f35b600080fdfe60806040818152600436101561001457600080fd5b600091823560e01c9081639b505aa1146100a5575063dc082ad31461003857600080fd5b346100a15760209061009961004c3661014a565b8351929161008b90610b2761006388820187610190565b8086526102538887013961007d8651938492898401610175565b03601f198101835282610190565b6001600160a01b03926101f4565b169051908152f35b5080fd5b839150346100a1576100b63661014a565b9091610b27808201906001600160401b0382118383101761013657916100e58486839489966102538639610175565b039082f590811561012c57935160209490936001600160a01b03938416907f15e71db3d71eb3b7985105d763101e1d6c1c491ab3e6a0d682558c12cc0bb8d69080a3168152f35b84513d85823e3d90fd5b634e487b7160e01b86526041600452602486fd5b6040906003190112610170576004356001600160a01b0381168103610170579060243590565b600080fd5b6001600160a01b039091168152602081019190915260400190565b601f909101601f19168101906001600160401b038211908210176101b357604052565b634e487b7160e01b600052604160045260246000fd5b9081519160005b8381106101e1575050016000815290565b80602080928401015181850152016101d0565b600b9061007d61021a6055946040519283916102146020840180976101c9565b906101c9565b519020604051906040820152600060208201523081520160ff815320803b156102405790565b60405163215db33160e01b8152600490fdfe60e060409080825234610120578181610b2780380380916100208285610125565b8339810103126101205780516001600160a01b03808216928383036101205760208060049201519486519283809263051ed6a360e41b82525afa908115610115576000916100d4575b501660805281156100c45760a05260c052516109c8908161015f82396080518181816101180152818161044f015281816104e0015261087a015260a05181818160b10152610568015260c05181818161024801526103780152f35b8251620d29f560e71b8152600490fd5b6020813d821161010d575b816100ec60209383610125565b810103126101095751908282168203610106575038610069565b80fd5b5080fd5b3d91506100df565b85513d6000823e3d90fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761014857604052565b634e487b7160e01b600052604160045260246000fdfe608060408181526004908136101561001657600080fd5b600092833560e01c908163023a96fe146105545750806326c0e5c51461052f57806330fc43ed1461050f57806351ed6a30146104cb5780637476083b146103d8578063956501bb1461039f5780639cfa2a2a146103605763bd3eec7d1461007c57600080fd5b3461035857600319602036820181136102b2578335916001600160401b03831161035c5760c083860191843603011261035c577f000000000000000000000000000000000000000000000000000000000000000060018060a01b0393878583169184359360ff8516809503610358578851630e0da79d60e11b81528a810186905260249888828b81895afa91821561034e57859261031f575b507f000000000000000000000000000000000000000000000000000000000000000016908a51636eb1769f60e11b8152308d820152868b8201528981604481865afa908115610315578c8e9593889795938e938e9c9b9a916102c7575b506101bc936101a961023297946101918f9c9b9a98956101b795610597565b905163095ea7b360e01b8d820152958693840161091f565b03601f1981018452836105ba565b6105f3565b6102208c51998a98899788966305fae14160e01b88528701528d8601528c8301356044860152604483013560648601526064830135608486015260a4610218610208608486018461093a565b60c0848a015260e4890191610971565b93019061093a565b8483036023190160c486015290610971565b03925af19182156102bd57869261028b575b50507f000000000000000000000000000000000000000000000000000000000000000092838203610273578580f35b516375c0811b60e01b81529384015282015260449150fd5b90809250813d83116102b6575b6102a281836105ba565b810103126102b257513880610244565b8480fd5b503d610298565b84513d88823e3d90fd5b9598975050998092508491509492943d831161030e575b6102e881836105ba565b8101031261030a579051899793948d949293909290918c918e916101bc610172565b8380fd5b503d6102de565b8c513d88823e3d90fd5b9091508881813d8311610347575b61033781836105ba565b810103126102b257519038610115565b503d61032d565b8b513d87823e3d90fd5b8280fd5b8580fd5b83823461039b578160031936011261039b57602090517f00000000000000000000000000000000000000000000000000000000000000008152f35b5080fd5b50903461035857602036600319011261035857356001600160a01b03811690819003610358578282916020945280845220549051908152f35b5082903461039b57602036600319011261039b5782359081156104bc573383528260205280832061040a838254610597565b905580516323b872dd60e01b60208201523360248201523060448201526064808201849052815260a081016001600160401b038111828210176104a957825261047c907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105f3565b519081527f0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc260203392a280f35b634e487b7160e01b855260418652602485fd5b51631f2a200560e01b81528390fd5b83823461039b578160031936011261039b57517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50503461039b57602036600319011261039b5761052c9035610813565b80f35b83823461039b578160031936011261039b5761052c9033835282602052822054610813565b84903461039b578160031936011261039b577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b919082018092116105a457565b634e487b7160e01b600052601160045260246000fd5b601f909101601f19168101906001600160401b038211908210176105dd57604052565b634e487b7160e01b600052604160045260246000fd5b604080516001600160a01b03929092169291906001600160401b03908201818111838210176105dd576040526020938483527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858401526000808587829751910182855af1903d15610734573d92831161072057906106929392916040519261068588601f19601f84011601856105ba565b83523d868885013e61073f565b8051806106a0575b50505050565b8184918101031261039b578201519081159182150361071d57506106c65780808061069a565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b80fd5b634e487b7160e01b85526041600452602485fd5b906106929392506060915b919290156107a15750815115610753575090565b3b1561075c5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156107b45750805190602001fd5b6040519062461bcd60e51b82528160208060048301528251908160248401526000935b8285106107fa575050604492506000838284010152601f80199101168101030190fd5b84810182015186860160440152938101938593506107d7565b801561090d576000338152806020526040812054908183116108e9578282039182116108d5576040903381528060205220556108a760405163a9059cbb60e01b60208201526108788161086a85336024840161091f565b03601f1981018352826105ba565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105f3565b6040519081527f8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc60203392a2565b634e487b7160e01b81526011600452602490fd5b606483836040519163a47b7c6560e01b835233600484015260248301526044820152fd5b604051631f2a200560e01b8152600490fd5b6001600160a01b039091168152602081019190915260400190565b9035601e198236030181121561096c570160208101919035906001600160401b03821161096c57813603831361096c57565b600080fd5b908060209392818452848401376000828201840152601f01601f191601019056fea26469706673582212209632da7ad69b2f1191cdaa449622d3d3c4563257ac2720527080276d4f004b9364736f6c63430008130033a2646970667358221220da86ce6007f27572938d290dad32a85aa0f1d10fddcdabca17445a58768ef59664736f6c63430008130033",
}

// EdgeStakingPoolCreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use EdgeStakingPoolCreatorMetaData.ABI instead.
var EdgeStakingPoolCreatorABI = EdgeStakingPoolCreatorMetaData.ABI

// EdgeStakingPoolCreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EdgeStakingPoolCreatorMetaData.Bin instead.
var EdgeStakingPoolCreatorBin = EdgeStakingPoolCreatorMetaData.Bin

// DeployEdgeStakingPoolCreator deploys a new Ethereum contract, binding an instance of EdgeStakingPoolCreator to it.
func DeployEdgeStakingPoolCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EdgeStakingPoolCreator, error) {
	parsed, err := EdgeStakingPoolCreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EdgeStakingPoolCreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EdgeStakingPoolCreator{EdgeStakingPoolCreatorCaller: EdgeStakingPoolCreatorCaller{contract: contract}, EdgeStakingPoolCreatorTransactor: EdgeStakingPoolCreatorTransactor{contract: contract}, EdgeStakingPoolCreatorFilterer: EdgeStakingPoolCreatorFilterer{contract: contract}}, nil
}

// EdgeStakingPoolCreator is an auto generated Go binding around an Ethereum contract.
type EdgeStakingPoolCreator struct {
	EdgeStakingPoolCreatorCaller     // Read-only binding to the contract
	EdgeStakingPoolCreatorTransactor // Write-only binding to the contract
	EdgeStakingPoolCreatorFilterer   // Log filterer for contract events
}

// EdgeStakingPoolCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type EdgeStakingPoolCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeStakingPoolCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EdgeStakingPoolCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeStakingPoolCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EdgeStakingPoolCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeStakingPoolCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EdgeStakingPoolCreatorSession struct {
	Contract     *EdgeStakingPoolCreator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EdgeStakingPoolCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EdgeStakingPoolCreatorCallerSession struct {
	Contract *EdgeStakingPoolCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// EdgeStakingPoolCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EdgeStakingPoolCreatorTransactorSession struct {
	Contract     *EdgeStakingPoolCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// EdgeStakingPoolCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type EdgeStakingPoolCreatorRaw struct {
	Contract *EdgeStakingPoolCreator // Generic contract binding to access the raw methods on
}

// EdgeStakingPoolCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EdgeStakingPoolCreatorCallerRaw struct {
	Contract *EdgeStakingPoolCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// EdgeStakingPoolCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EdgeStakingPoolCreatorTransactorRaw struct {
	Contract *EdgeStakingPoolCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEdgeStakingPoolCreator creates a new instance of EdgeStakingPoolCreator, bound to a specific deployed contract.
func NewEdgeStakingPoolCreator(address common.Address, backend bind.ContractBackend) (*EdgeStakingPoolCreator, error) {
	contract, err := bindEdgeStakingPoolCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolCreator{EdgeStakingPoolCreatorCaller: EdgeStakingPoolCreatorCaller{contract: contract}, EdgeStakingPoolCreatorTransactor: EdgeStakingPoolCreatorTransactor{contract: contract}, EdgeStakingPoolCreatorFilterer: EdgeStakingPoolCreatorFilterer{contract: contract}}, nil
}

// NewEdgeStakingPoolCreatorCaller creates a new read-only instance of EdgeStakingPoolCreator, bound to a specific deployed contract.
func NewEdgeStakingPoolCreatorCaller(address common.Address, caller bind.ContractCaller) (*EdgeStakingPoolCreatorCaller, error) {
	contract, err := bindEdgeStakingPoolCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolCreatorCaller{contract: contract}, nil
}

// NewEdgeStakingPoolCreatorTransactor creates a new write-only instance of EdgeStakingPoolCreator, bound to a specific deployed contract.
func NewEdgeStakingPoolCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*EdgeStakingPoolCreatorTransactor, error) {
	contract, err := bindEdgeStakingPoolCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolCreatorTransactor{contract: contract}, nil
}

// NewEdgeStakingPoolCreatorFilterer creates a new log filterer instance of EdgeStakingPoolCreator, bound to a specific deployed contract.
func NewEdgeStakingPoolCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*EdgeStakingPoolCreatorFilterer, error) {
	contract, err := bindEdgeStakingPoolCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolCreatorFilterer{contract: contract}, nil
}

// bindEdgeStakingPoolCreator binds a generic wrapper to an already deployed contract.
func bindEdgeStakingPoolCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EdgeStakingPoolCreatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EdgeStakingPoolCreator.Contract.EdgeStakingPoolCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EdgeStakingPoolCreator.Contract.EdgeStakingPoolCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EdgeStakingPoolCreator.Contract.EdgeStakingPoolCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EdgeStakingPoolCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EdgeStakingPoolCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EdgeStakingPoolCreator.Contract.contract.Transact(opts, method, params...)
}

// GetPool is a free data retrieval call binding the contract method 0xdc082ad3.
//
// Solidity: function getPool(address challengeManager, bytes32 edgeId) view returns(address)
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorCaller) GetPool(opts *bind.CallOpts, challengeManager common.Address, edgeId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EdgeStakingPoolCreator.contract.Call(opts, &out, "getPool", challengeManager, edgeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0xdc082ad3.
//
// Solidity: function getPool(address challengeManager, bytes32 edgeId) view returns(address)
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorSession) GetPool(challengeManager common.Address, edgeId [32]byte) (common.Address, error) {
	return _EdgeStakingPoolCreator.Contract.GetPool(&_EdgeStakingPoolCreator.CallOpts, challengeManager, edgeId)
}

// GetPool is a free data retrieval call binding the contract method 0xdc082ad3.
//
// Solidity: function getPool(address challengeManager, bytes32 edgeId) view returns(address)
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorCallerSession) GetPool(challengeManager common.Address, edgeId [32]byte) (common.Address, error) {
	return _EdgeStakingPoolCreator.Contract.GetPool(&_EdgeStakingPoolCreator.CallOpts, challengeManager, edgeId)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9b505aa1.
//
// Solidity: function createPool(address challengeManager, bytes32 edgeId) returns(address)
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorTransactor) CreatePool(opts *bind.TransactOpts, challengeManager common.Address, edgeId [32]byte) (*types.Transaction, error) {
	return _EdgeStakingPoolCreator.contract.Transact(opts, "createPool", challengeManager, edgeId)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9b505aa1.
//
// Solidity: function createPool(address challengeManager, bytes32 edgeId) returns(address)
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorSession) CreatePool(challengeManager common.Address, edgeId [32]byte) (*types.Transaction, error) {
	return _EdgeStakingPoolCreator.Contract.CreatePool(&_EdgeStakingPoolCreator.TransactOpts, challengeManager, edgeId)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9b505aa1.
//
// Solidity: function createPool(address challengeManager, bytes32 edgeId) returns(address)
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorTransactorSession) CreatePool(challengeManager common.Address, edgeId [32]byte) (*types.Transaction, error) {
	return _EdgeStakingPoolCreator.Contract.CreatePool(&_EdgeStakingPoolCreator.TransactOpts, challengeManager, edgeId)
}

// EdgeStakingPoolCreatorNewEdgeStakingPoolCreatedIterator is returned from FilterNewEdgeStakingPoolCreated and is used to iterate over the raw logs and unpacked data for NewEdgeStakingPoolCreated events raised by the EdgeStakingPoolCreator contract.
type EdgeStakingPoolCreatorNewEdgeStakingPoolCreatedIterator struct {
	Event *EdgeStakingPoolCreatorNewEdgeStakingPoolCreated // Event containing the contract specifics and raw log

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
func (it *EdgeStakingPoolCreatorNewEdgeStakingPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeStakingPoolCreatorNewEdgeStakingPoolCreated)
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
		it.Event = new(EdgeStakingPoolCreatorNewEdgeStakingPoolCreated)
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
func (it *EdgeStakingPoolCreatorNewEdgeStakingPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeStakingPoolCreatorNewEdgeStakingPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeStakingPoolCreatorNewEdgeStakingPoolCreated represents a NewEdgeStakingPoolCreated event raised by the EdgeStakingPoolCreator contract.
type EdgeStakingPoolCreatorNewEdgeStakingPoolCreated struct {
	ChallengeManager common.Address
	EdgeId           [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewEdgeStakingPoolCreated is a free log retrieval operation binding the contract event 0x15e71db3d71eb3b7985105d763101e1d6c1c491ab3e6a0d682558c12cc0bb8d6.
//
// Solidity: event NewEdgeStakingPoolCreated(address indexed challengeManager, bytes32 indexed edgeId)
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorFilterer) FilterNewEdgeStakingPoolCreated(opts *bind.FilterOpts, challengeManager []common.Address, edgeId [][32]byte) (*EdgeStakingPoolCreatorNewEdgeStakingPoolCreatedIterator, error) {

	var challengeManagerRule []interface{}
	for _, challengeManagerItem := range challengeManager {
		challengeManagerRule = append(challengeManagerRule, challengeManagerItem)
	}
	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}

	logs, sub, err := _EdgeStakingPoolCreator.contract.FilterLogs(opts, "NewEdgeStakingPoolCreated", challengeManagerRule, edgeIdRule)
	if err != nil {
		return nil, err
	}
	return &EdgeStakingPoolCreatorNewEdgeStakingPoolCreatedIterator{contract: _EdgeStakingPoolCreator.contract, event: "NewEdgeStakingPoolCreated", logs: logs, sub: sub}, nil
}

// WatchNewEdgeStakingPoolCreated is a free log subscription operation binding the contract event 0x15e71db3d71eb3b7985105d763101e1d6c1c491ab3e6a0d682558c12cc0bb8d6.
//
// Solidity: event NewEdgeStakingPoolCreated(address indexed challengeManager, bytes32 indexed edgeId)
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorFilterer) WatchNewEdgeStakingPoolCreated(opts *bind.WatchOpts, sink chan<- *EdgeStakingPoolCreatorNewEdgeStakingPoolCreated, challengeManager []common.Address, edgeId [][32]byte) (event.Subscription, error) {

	var challengeManagerRule []interface{}
	for _, challengeManagerItem := range challengeManager {
		challengeManagerRule = append(challengeManagerRule, challengeManagerItem)
	}
	var edgeIdRule []interface{}
	for _, edgeIdItem := range edgeId {
		edgeIdRule = append(edgeIdRule, edgeIdItem)
	}

	logs, sub, err := _EdgeStakingPoolCreator.contract.WatchLogs(opts, "NewEdgeStakingPoolCreated", challengeManagerRule, edgeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeStakingPoolCreatorNewEdgeStakingPoolCreated)
				if err := _EdgeStakingPoolCreator.contract.UnpackLog(event, "NewEdgeStakingPoolCreated", log); err != nil {
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

// ParseNewEdgeStakingPoolCreated is a log parse operation binding the contract event 0x15e71db3d71eb3b7985105d763101e1d6c1c491ab3e6a0d682558c12cc0bb8d6.
//
// Solidity: event NewEdgeStakingPoolCreated(address indexed challengeManager, bytes32 indexed edgeId)
func (_EdgeStakingPoolCreator *EdgeStakingPoolCreatorFilterer) ParseNewEdgeStakingPoolCreated(log types.Log) (*EdgeStakingPoolCreatorNewEdgeStakingPoolCreated, error) {
	event := new(EdgeStakingPoolCreatorNewEdgeStakingPoolCreated)
	if err := _EdgeStakingPoolCreator.contract.UnpackLog(event, "NewEdgeStakingPoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolCreatorUtilsMetaData contains all meta data concerning the StakingPoolCreatorUtils contract.
var StakingPoolCreatorUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PoolDoesntExist\",\"type\":\"error\"}]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220ef1326ee40c3be504d952c3a3fd3bbabf31873ad8881457c37d58b3b88763c3964736f6c63430008130033",
}

// StakingPoolCreatorUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingPoolCreatorUtilsMetaData.ABI instead.
var StakingPoolCreatorUtilsABI = StakingPoolCreatorUtilsMetaData.ABI

// StakingPoolCreatorUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingPoolCreatorUtilsMetaData.Bin instead.
var StakingPoolCreatorUtilsBin = StakingPoolCreatorUtilsMetaData.Bin

// DeployStakingPoolCreatorUtils deploys a new Ethereum contract, binding an instance of StakingPoolCreatorUtils to it.
func DeployStakingPoolCreatorUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StakingPoolCreatorUtils, error) {
	parsed, err := StakingPoolCreatorUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingPoolCreatorUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingPoolCreatorUtils{StakingPoolCreatorUtilsCaller: StakingPoolCreatorUtilsCaller{contract: contract}, StakingPoolCreatorUtilsTransactor: StakingPoolCreatorUtilsTransactor{contract: contract}, StakingPoolCreatorUtilsFilterer: StakingPoolCreatorUtilsFilterer{contract: contract}}, nil
}

// StakingPoolCreatorUtils is an auto generated Go binding around an Ethereum contract.
type StakingPoolCreatorUtils struct {
	StakingPoolCreatorUtilsCaller     // Read-only binding to the contract
	StakingPoolCreatorUtilsTransactor // Write-only binding to the contract
	StakingPoolCreatorUtilsFilterer   // Log filterer for contract events
}

// StakingPoolCreatorUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingPoolCreatorUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolCreatorUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingPoolCreatorUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolCreatorUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingPoolCreatorUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolCreatorUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingPoolCreatorUtilsSession struct {
	Contract     *StakingPoolCreatorUtils // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingPoolCreatorUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingPoolCreatorUtilsCallerSession struct {
	Contract *StakingPoolCreatorUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// StakingPoolCreatorUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingPoolCreatorUtilsTransactorSession struct {
	Contract     *StakingPoolCreatorUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// StakingPoolCreatorUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingPoolCreatorUtilsRaw struct {
	Contract *StakingPoolCreatorUtils // Generic contract binding to access the raw methods on
}

// StakingPoolCreatorUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingPoolCreatorUtilsCallerRaw struct {
	Contract *StakingPoolCreatorUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// StakingPoolCreatorUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingPoolCreatorUtilsTransactorRaw struct {
	Contract *StakingPoolCreatorUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingPoolCreatorUtils creates a new instance of StakingPoolCreatorUtils, bound to a specific deployed contract.
func NewStakingPoolCreatorUtils(address common.Address, backend bind.ContractBackend) (*StakingPoolCreatorUtils, error) {
	contract, err := bindStakingPoolCreatorUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingPoolCreatorUtils{StakingPoolCreatorUtilsCaller: StakingPoolCreatorUtilsCaller{contract: contract}, StakingPoolCreatorUtilsTransactor: StakingPoolCreatorUtilsTransactor{contract: contract}, StakingPoolCreatorUtilsFilterer: StakingPoolCreatorUtilsFilterer{contract: contract}}, nil
}

// NewStakingPoolCreatorUtilsCaller creates a new read-only instance of StakingPoolCreatorUtils, bound to a specific deployed contract.
func NewStakingPoolCreatorUtilsCaller(address common.Address, caller bind.ContractCaller) (*StakingPoolCreatorUtilsCaller, error) {
	contract, err := bindStakingPoolCreatorUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolCreatorUtilsCaller{contract: contract}, nil
}

// NewStakingPoolCreatorUtilsTransactor creates a new write-only instance of StakingPoolCreatorUtils, bound to a specific deployed contract.
func NewStakingPoolCreatorUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingPoolCreatorUtilsTransactor, error) {
	contract, err := bindStakingPoolCreatorUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolCreatorUtilsTransactor{contract: contract}, nil
}

// NewStakingPoolCreatorUtilsFilterer creates a new log filterer instance of StakingPoolCreatorUtils, bound to a specific deployed contract.
func NewStakingPoolCreatorUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingPoolCreatorUtilsFilterer, error) {
	contract, err := bindStakingPoolCreatorUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingPoolCreatorUtilsFilterer{contract: contract}, nil
}

// bindStakingPoolCreatorUtils binds a generic wrapper to an already deployed contract.
func bindStakingPoolCreatorUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingPoolCreatorUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolCreatorUtils *StakingPoolCreatorUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolCreatorUtils.Contract.StakingPoolCreatorUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolCreatorUtils *StakingPoolCreatorUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolCreatorUtils.Contract.StakingPoolCreatorUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolCreatorUtils *StakingPoolCreatorUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolCreatorUtils.Contract.StakingPoolCreatorUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolCreatorUtils *StakingPoolCreatorUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolCreatorUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolCreatorUtils *StakingPoolCreatorUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolCreatorUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolCreatorUtils *StakingPoolCreatorUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolCreatorUtils.Contract.contract.Transact(opts, method, params...)
}
