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
	Bin: "0x60e060409080825234610113578181610d528038038091610020828561012b565b83398101031261011357602061003582610164565b910151825163051ed6a360e41b81529091906020816004816001600160a01b0386165afa908115610120576000916100e1575b5060805281156100d05760a05260c05251610bd9908161017982396080518181816101d9015281816104fe015281816105d40152610712015260a05181818160bc0152818161019e0152818161086f0152610924015260c05181818161038c01526106510152f35b8251630b12999960e11b8152600490fd5b90506020813d602011610118575b816100fc6020938361012b565b810103126101135761010d90610164565b38610068565b600080fd5b3d91506100ef565b84513d6000823e3d90fd5b601f909101601f19168101906001600160401b0382119082101761014e57604052565b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101135756fe608060408181526004918236101561001657600080fd5b600092833560e01c9182632113ed211461063a5750816326c0e5c51461061557816330fc43ed146105f857816351ed6a30146105b45781636b74d515146105935781637476083b1461046e578163839159711461015d578163930412af146101445781639451944d14610125578163956501bb146100e4575063cb23bcb51461009e57600080fd5b346100e057816003193601126100e057602090516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5080fd5b90503461012157602060031936011261012157356001600160a01b03811680910361011c578282916020945280845220549051908152f35b600080fd5b8280fd5b833461014157806003193601126101415761013e6108ec565b80f35b80fd5b833461014157806003193601126101415761013e610865565b91905034610121576102606003193601126101215780517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230838201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038181166024840181905292602092606435917f00000000000000000000000000000000000000000000000000000000000000008416908581604481855afa9081156104645790849392918b9161042b575b5061022761027794610272926107e0565b89517f095ea7b300000000000000000000000000000000000000000000000000000000898201526001600160a01b03949094166024850152604480850191909152835260648361082d565b61097a565b833b15610406578451957f50f32f6800000000000000000000000000000000000000000000000000000000875281818801523560248701526024356044870152604435606487015260848601526084359081168091036104275760a485015260a4359067ffffffffffffffff918281168091036104065760c486015260c4358281168091036104065760e48601528360e46101048701376101248661014487015b836002831061040a57505050506101643560038110156104065790848794939261018490818901526101a49035818901526101c4880137836101e461020488015b600283106103dc57505050505061022435600381101561012157848381936102c4936102449081850152356102648401527f0000000000000000000000000000000000000000000000000000000000000000610284840152306102a48401525af19081156103d357506103ca575080f35b61013e90610803565b513d84823e3d90fd5b84959650928082939481966103f2600195610850565b168152019201920190918895949392610359565b8680fd5b806001928761041887610850565b16815201930191019091610318565b8580fd5b809450878092503d831161045d575b610444818361082d565b81010312610459579151839290610227610216565b8980fd5b503d61043a565b88513d8c823e3d90fd5b8383346100e05760206003193601126100e057823590811561056b573383528260205280832061049f8382546107e0565b905580517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201528260648201526064815260a0810181811067ffffffffffffffff82111761055857825261052b907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661097a565b519081527f0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc260203392a280f35b602485604188634e487b7160e01b835252fd5b8390517f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b83346101415780600319360112610141576105ac610865565b61013e6108ec565b5050346100e057816003193601126100e057602090516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b8390346100e05760206003193601126100e05761013e9035610674565b5050346100e057816003193601126100e05761013e9033835282602052822054610674565b8490346100e057816003193601126100e0576020907f00000000000000000000000000000000000000000000000000000000000000008152f35b80156107b65760003381528060205260408120549081831161077957828203918211610765576040903381528060205220556107376040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820152610708816106fa853360248401602090939291936001600160a01b0360408201951681520152565b03601f19810183528261082d565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001661097a565b6040519081527f8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc60203392a2565b80634e487b7160e01b602492526011600452fd5b60648383604051917fa47b7c6500000000000000000000000000000000000000000000000000000000835233600484015260248301526044820152fd5b60046040517f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b919082018092116107ed57565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff811161081757604052565b634e487b7160e01b600052604160045260246000fd5b90601f601f19910116810190811067ffffffffffffffff82111761081757604052565b359067ffffffffffffffff8216820361011c57565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1561011c57600080916004604051809481937f57ef4ab90000000000000000000000000000000000000000000000000000000083525af180156108e0576108d55750565b6108de90610803565b565b6040513d6000823e3d90fd5b6040517f6137391900000000000000000000000000000000000000000000000000000000815260208160048160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af180156108e0576109535750565b602090813d8311610973575b610969818361082d565b8101031261011c57565b503d61095f565b6001600160a01b0316906040516040810167ffffffffffffffff9082811082821117610817576040526020938483527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858401526000808587829751910182855af1903d15610ac8573d928311610ab45790610a1693929160405192610a0988601f19601f840116018561082d565b83523d868885013e610ad3565b805180610a24575b50505050565b818491810103126100e057820151908115918215036101415750610a4a57808080610a1e565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b602485634e487b7160e01b81526041600452fd5b90610a169392506060915b91929015610b345750815115610ae7575090565b3b15610af05790565b606460405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015610b475750805190602001fd5b6040519062461bcd60e51b825281602080600483015282519283602484015260005b848110610b8c57505050601f19601f836000604480968601015201168101030190fd5b818101830151868201604401528593508201610b6956fea2646970667358221220c7ba63ab38a99229f25ec7cb215c7f5787907b18abbafe1feac37cdfad96df3464736f6c63430008190033",
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
	Bin: "0x608080604052346015576110ab908161001b8239f35b600080fdfe60806040818152600436101561001457600080fd5b600091823560e01c9081639b505aa1146100d4575063dc082ad31461003857600080fd5b346100d05760209073ffffffffffffffffffffffffffffffffffffffff6100c8610061366101cc565b906100c3610d529186519261007889820185610222565b808452610324898501396100b587519485928a84016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03601f198101845283610222565b61029f565b169051908152f35b5080fd5b919050346101c8576100e5366101cc565b9092610d528082019082821067ffffffffffffffff83111761019b5791610138848783948a9661032486396020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b039082f592831561019157602094507fd628317c6ebae87acc5dbfadeb835cb97692cc6935ea72bf37461e14a0bbee1e8573ffffffffffffffffffffffffffffffffffffffff809616958551938785521692a351908152f35b82513d86823e3d90fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8280fd5b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc604091011261021d5760043573ffffffffffffffffffffffffffffffffffffffff8116810361021d579060243590565b600080fd5b90601f601f19910116810190811067ffffffffffffffff82111761024557604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b9081519160005b83811061028c575050016000815290565b806020809284010151818501520161027b565b600b906102c56102d36055946040519283916102bf602084018097610274565b90610274565b03601f198101835282610222565b519020604051906040820152600060208201523081520160ff815320803b156102f95790565b60046040517f215db331000000000000000000000000000000000000000000000000000000008152fdfe60e060409080825234610113578181610d528038038091610020828561012b565b83398101031261011357602061003582610164565b910151825163051ed6a360e41b81529091906020816004816001600160a01b0386165afa908115610120576000916100e1575b5060805281156100d05760a05260c05251610bd9908161017982396080518181816101d9015281816104fe015281816105d40152610712015260a05181818160bc0152818161019e0152818161086f0152610924015260c05181818161038c01526106510152f35b8251630b12999960e11b8152600490fd5b90506020813d602011610118575b816100fc6020938361012b565b810103126101135761010d90610164565b38610068565b600080fd5b3d91506100ef565b84513d6000823e3d90fd5b601f909101601f19168101906001600160401b0382119082101761014e57604052565b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101135756fe608060408181526004918236101561001657600080fd5b600092833560e01c9182632113ed211461063a5750816326c0e5c51461061557816330fc43ed146105f857816351ed6a30146105b45781636b74d515146105935781637476083b1461046e578163839159711461015d578163930412af146101445781639451944d14610125578163956501bb146100e4575063cb23bcb51461009e57600080fd5b346100e057816003193601126100e057602090516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5080fd5b90503461012157602060031936011261012157356001600160a01b03811680910361011c578282916020945280845220549051908152f35b600080fd5b8280fd5b833461014157806003193601126101415761013e6108ec565b80f35b80fd5b833461014157806003193601126101415761013e610865565b91905034610121576102606003193601126101215780517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230838201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038181166024840181905292602092606435917f00000000000000000000000000000000000000000000000000000000000000008416908581604481855afa9081156104645790849392918b9161042b575b5061022761027794610272926107e0565b89517f095ea7b300000000000000000000000000000000000000000000000000000000898201526001600160a01b03949094166024850152604480850191909152835260648361082d565b61097a565b833b15610406578451957f50f32f6800000000000000000000000000000000000000000000000000000000875281818801523560248701526024356044870152604435606487015260848601526084359081168091036104275760a485015260a4359067ffffffffffffffff918281168091036104065760c486015260c4358281168091036104065760e48601528360e46101048701376101248661014487015b836002831061040a57505050506101643560038110156104065790848794939261018490818901526101a49035818901526101c4880137836101e461020488015b600283106103dc57505050505061022435600381101561012157848381936102c4936102449081850152356102648401527f0000000000000000000000000000000000000000000000000000000000000000610284840152306102a48401525af19081156103d357506103ca575080f35b61013e90610803565b513d84823e3d90fd5b84959650928082939481966103f2600195610850565b168152019201920190918895949392610359565b8680fd5b806001928761041887610850565b16815201930191019091610318565b8580fd5b809450878092503d831161045d575b610444818361082d565b81010312610459579151839290610227610216565b8980fd5b503d61043a565b88513d8c823e3d90fd5b8383346100e05760206003193601126100e057823590811561056b573383528260205280832061049f8382546107e0565b905580517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201528260648201526064815260a0810181811067ffffffffffffffff82111761055857825261052b907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661097a565b519081527f0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc260203392a280f35b602485604188634e487b7160e01b835252fd5b8390517f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b83346101415780600319360112610141576105ac610865565b61013e6108ec565b5050346100e057816003193601126100e057602090516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b8390346100e05760206003193601126100e05761013e9035610674565b5050346100e057816003193601126100e05761013e9033835282602052822054610674565b8490346100e057816003193601126100e0576020907f00000000000000000000000000000000000000000000000000000000000000008152f35b80156107b65760003381528060205260408120549081831161077957828203918211610765576040903381528060205220556107376040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820152610708816106fa853360248401602090939291936001600160a01b0360408201951681520152565b03601f19810183528261082d565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001661097a565b6040519081527f8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc60203392a2565b80634e487b7160e01b602492526011600452fd5b60648383604051917fa47b7c6500000000000000000000000000000000000000000000000000000000835233600484015260248301526044820152fd5b60046040517f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b919082018092116107ed57565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff811161081757604052565b634e487b7160e01b600052604160045260246000fd5b90601f601f19910116810190811067ffffffffffffffff82111761081757604052565b359067ffffffffffffffff8216820361011c57565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1561011c57600080916004604051809481937f57ef4ab90000000000000000000000000000000000000000000000000000000083525af180156108e0576108d55750565b6108de90610803565b565b6040513d6000823e3d90fd5b6040517f6137391900000000000000000000000000000000000000000000000000000000815260208160048160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af180156108e0576109535750565b602090813d8311610973575b610969818361082d565b8101031261011c57565b503d61095f565b6001600160a01b0316906040516040810167ffffffffffffffff9082811082821117610817576040526020938483527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858401526000808587829751910182855af1903d15610ac8573d928311610ab45790610a1693929160405192610a0988601f19601f840116018561082d565b83523d868885013e610ad3565b805180610a24575b50505050565b818491810103126100e057820151908115918215036101415750610a4a57808080610a1e565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b602485634e487b7160e01b81526041600452fd5b90610a169392506060915b91929015610b345750815115610ae7575090565b3b15610af05790565b606460405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015610b475750805190602001fd5b6040519062461bcd60e51b825281602080600483015282519283602484015260005b848110610b8c57505050601f19601f836000604480968601015201168101030190fd5b818101830151868201604401528593508201610b6956fea2646970667358221220c7ba63ab38a99229f25ec7cb215c7f5787907b18abbafe1feac37cdfad96df3464736f6c63430008190033a2646970667358221220d8c30d5b4790a14f2689389e52fb9cecab39acdab1f81728bd1824caa5aa3ff764736f6c63430008190033",
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
	Bin: "0x60e060409080825234610121578181610ceb80380380916100208285610126565b8339810103126101215780516001600160a01b03808216928383036101215760208060049201519486519283809263051ed6a360e41b82525afa908115610116576000916100d4575b501660805281156100c45760a05260c05251610b8b908161016082396080518181816101410152818161053c0152818161060b015261077c015260a05181818160b301526106a0015260c0518181816102fd015261043f0152f35b8251620d29f560e71b8152600490fd5b6020813d60201161010e575b816100ed60209383610126565b8101031261010a5751908282168203610107575038610069565b80fd5b5080fd5b3d91506100e0565b85513d6000823e3d90fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761014957604052565b634e487b7160e01b600052604160045260246000fdfe608060408181526004908136101561001657600080fd5b600092833560e01c908163023a96fe146106745750806326c0e5c51461064f57806330fc43ed1461062f57806351ed6a30146105de5780637476083b146104ab578063956501bb146104665780639cfa2a2a146104275763bd3eec7d1461007c57600080fd5b3461041f5760206003199080823601126103805783359167ffffffffffffffff83116104235760c0838601918436030112610423577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff93878583169184359360ff851680950361041f578851907f1c1b4f3a000000000000000000000000000000000000000000000000000000008252858b83015260249888838b81895afa9283156104155785936103e6575b507f000000000000000000000000000000000000000000000000000000000000000016918a517fdd62ed3e000000000000000000000000000000000000000000000000000000008152308d820152868b8201528981604481875afa9081156103dc57928b8e95938e938998968e9c9b9a91610395575b508b610227819a99989661023a967f095ea7b300000000000000000000000000000000000000000000000000000000966101f56102e79b976102359761084a565b925197889586015284016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03601f19810184528361086d565b61091c565b6102b78c51998a98899788967f05fae1410000000000000000000000000000000000000000000000000000000088528701528d8601528c8301356044860152604483013560648601526064830135608486015260a46102af61029f60848601846108a6565b60c0848a015260e48901916108fb565b9301906108a6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc8584030160c48601526108fb565b03925af191821561038b578692610359575b50507f000000000000000000000000000000000000000000000000000000000000000092838203610328578580f35b6044955051937f75c0811b000000000000000000000000000000000000000000000000000000008552840152820152fd5b90809250813d8311610384575b610370818361086d565b81010312610380575138806102f9565b8480fd5b503d610366565b84513d88823e3d90fd5b9b9398975050939492505088813d83116103d5575b6103b4818361086d565b810103126103d1579651899793948d9493928d928d91908b6101b4565b8380fd5b503d6103aa565b8c513d88823e3d90fd5b9092508881813d831161040e575b6103fe818361086d565b810103126103805751913861013e565b503d6103f4565b8b513d87823e3d90fd5b8280fd5b8580fd5b838234610462578160031936011261046257602090517f00000000000000000000000000000000000000000000000000000000000000008152f35b5080fd5b50903461041f57602060031936011261041f573573ffffffffffffffffffffffffffffffffffffffff811680910361041f578282916020945280845220549051908152f35b508290346104625760206003193601126104625782359081156105b657338352826020528083206104dd83825461084a565b905580517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201528260648201526064815260a0810181811067ffffffffffffffff8211176105a3578252610576907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661091c565b519081527f0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc260203392a280f35b602485604188634e487b7160e01b835252fd5b8390517f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b8382346104625781600319360112610462576020905173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5050346104625760206003193601126104625761064c90356106c4565b80f35b83823461046257816003193601126104625761064c90338352826020528220546106c4565b84903461046257816003193601126104625760209073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b8015610820576000338152806020526040812054908183116107e3578282039182116107cf576040903381528060205220556107a16040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820152610765816107578533602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03601f19810183528261086d565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001661091c565b6040519081527f8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc60203392a2565b80634e487b7160e01b602492526011600452fd5b60648383604051917fa47b7c6500000000000000000000000000000000000000000000000000000000835233600484015260248301526044820152fd5b60046040517f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b9190820180921161085757565b634e487b7160e01b600052601160045260246000fd5b90601f601f19910116810190811067ffffffffffffffff82111761089057604052565b634e487b7160e01b600052604160045260246000fd5b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156108f657016020813591019167ffffffffffffffff82116108f65781360383136108f657565b600080fd5b601f8260209493601f19938186528686013760008582860101520116010190565b73ffffffffffffffffffffffffffffffffffffffff16906040516040810167ffffffffffffffff9082811082821117610890576040526020938483527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858401526000808587829751910182855af1903d15610a7a573d928311610a6657906109c5939291604051926109b888601f19601f840116018561086d565b83523d868885013e610a85565b8051806109d3575b50505050565b818491810103126104625782015190811591821503610a6357506109f9578080806109cd565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80fd5b602485634e487b7160e01b81526041600452fd5b906109c59392506060915b91929015610ae65750815115610a99575090565b3b15610aa25790565b606460405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015610af95750805190602001fd5b6040519062461bcd60e51b825281602080600483015282519283602484015260005b848110610b3e57505050601f19601f836000604480968601015201168101030190fd5b818101830151868201604401528593508201610b1b56fea2646970667358221220ac4310514ca2e32862c90f2c9c183da61993484b9eaa18bccd7851bfae55906464736f6c63430008190033",
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
	Bin: "0x6080806040523460155761103c908161001b8239f35b600080fdfe60806040818152600436101561001457600080fd5b600091823560e01c9081639b505aa1146100d4575063dc082ad31461003857600080fd5b346100d05760209073ffffffffffffffffffffffffffffffffffffffff6100c8610061366101c4565b906100c3610ceb918651926100788982018561021a565b80845261031c898501396100b587519485928a84016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03601f19810184528361021a565b610297565b169051908152f35b5080fd5b919050346101c0576100e5366101c4565b9092610ceb8082019082821067ffffffffffffffff8311176101935791610138848783948a9661031c86396020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b039082f5908115610189576020947f15e71db3d71eb3b7985105d763101e1d6c1c491ab3e6a0d682558c12cc0bb8d673ffffffffffffffffffffffffffffffffffffffff80955196169180a3168152f35b82513d86823e3d90fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8280fd5b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60409101126102155760043573ffffffffffffffffffffffffffffffffffffffff81168103610215579060243590565b600080fd5b90601f601f19910116810190811067ffffffffffffffff82111761023d57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b9081519160005b838110610284575050016000815290565b8060208092840101518185015201610273565b600b906102bd6102cb6055946040519283916102b760208401809761026c565b9061026c565b03601f19810183528261021a565b519020604051906040820152600060208201523081520160ff815320803b156102f15790565b60046040517f215db331000000000000000000000000000000000000000000000000000000008152fdfe60e060409080825234610121578181610ceb80380380916100208285610126565b8339810103126101215780516001600160a01b03808216928383036101215760208060049201519486519283809263051ed6a360e41b82525afa908115610116576000916100d4575b501660805281156100c45760a05260c05251610b8b908161016082396080518181816101410152818161053c0152818161060b015261077c015260a05181818160b301526106a0015260c0518181816102fd015261043f0152f35b8251620d29f560e71b8152600490fd5b6020813d60201161010e575b816100ed60209383610126565b8101031261010a5751908282168203610107575038610069565b80fd5b5080fd5b3d91506100e0565b85513d6000823e3d90fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761014957604052565b634e487b7160e01b600052604160045260246000fdfe608060408181526004908136101561001657600080fd5b600092833560e01c908163023a96fe146106745750806326c0e5c51461064f57806330fc43ed1461062f57806351ed6a30146105de5780637476083b146104ab578063956501bb146104665780639cfa2a2a146104275763bd3eec7d1461007c57600080fd5b3461041f5760206003199080823601126103805783359167ffffffffffffffff83116104235760c0838601918436030112610423577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff93878583169184359360ff851680950361041f578851907f1c1b4f3a000000000000000000000000000000000000000000000000000000008252858b83015260249888838b81895afa9283156104155785936103e6575b507f000000000000000000000000000000000000000000000000000000000000000016918a517fdd62ed3e000000000000000000000000000000000000000000000000000000008152308d820152868b8201528981604481875afa9081156103dc57928b8e95938e938998968e9c9b9a91610395575b508b610227819a99989661023a967f095ea7b300000000000000000000000000000000000000000000000000000000966101f56102e79b976102359761084a565b925197889586015284016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03601f19810184528361086d565b61091c565b6102b78c51998a98899788967f05fae1410000000000000000000000000000000000000000000000000000000088528701528d8601528c8301356044860152604483013560648601526064830135608486015260a46102af61029f60848601846108a6565b60c0848a015260e48901916108fb565b9301906108a6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc8584030160c48601526108fb565b03925af191821561038b578692610359575b50507f000000000000000000000000000000000000000000000000000000000000000092838203610328578580f35b6044955051937f75c0811b000000000000000000000000000000000000000000000000000000008552840152820152fd5b90809250813d8311610384575b610370818361086d565b81010312610380575138806102f9565b8480fd5b503d610366565b84513d88823e3d90fd5b9b9398975050939492505088813d83116103d5575b6103b4818361086d565b810103126103d1579651899793948d9493928d928d91908b6101b4565b8380fd5b503d6103aa565b8c513d88823e3d90fd5b9092508881813d831161040e575b6103fe818361086d565b810103126103805751913861013e565b503d6103f4565b8b513d87823e3d90fd5b8280fd5b8580fd5b838234610462578160031936011261046257602090517f00000000000000000000000000000000000000000000000000000000000000008152f35b5080fd5b50903461041f57602060031936011261041f573573ffffffffffffffffffffffffffffffffffffffff811680910361041f578282916020945280845220549051908152f35b508290346104625760206003193601126104625782359081156105b657338352826020528083206104dd83825461084a565b905580517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201528260648201526064815260a0810181811067ffffffffffffffff8211176105a3578252610576907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661091c565b519081527f0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc260203392a280f35b602485604188634e487b7160e01b835252fd5b8390517f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b8382346104625781600319360112610462576020905173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5050346104625760206003193601126104625761064c90356106c4565b80f35b83823461046257816003193601126104625761064c90338352826020528220546106c4565b84903461046257816003193601126104625760209073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b8015610820576000338152806020526040812054908183116107e3578282039182116107cf576040903381528060205220556107a16040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820152610765816107578533602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03601f19810183528261086d565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001661091c565b6040519081527f8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc60203392a2565b80634e487b7160e01b602492526011600452fd5b60648383604051917fa47b7c6500000000000000000000000000000000000000000000000000000000835233600484015260248301526044820152fd5b60046040517f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b9190820180921161085757565b634e487b7160e01b600052601160045260246000fd5b90601f601f19910116810190811067ffffffffffffffff82111761089057604052565b634e487b7160e01b600052604160045260246000fd5b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156108f657016020813591019167ffffffffffffffff82116108f65781360383136108f657565b600080fd5b601f8260209493601f19938186528686013760008582860101520116010190565b73ffffffffffffffffffffffffffffffffffffffff16906040516040810167ffffffffffffffff9082811082821117610890576040526020938483527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858401526000808587829751910182855af1903d15610a7a573d928311610a6657906109c5939291604051926109b888601f19601f840116018561086d565b83523d868885013e610a85565b8051806109d3575b50505050565b818491810103126104625782015190811591821503610a6357506109f9578080806109cd565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80fd5b602485634e487b7160e01b81526041600452fd5b906109c59392506060915b91929015610ae65750815115610a99575090565b3b15610aa25790565b606460405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015610af95750805190602001fd5b6040519062461bcd60e51b825281602080600483015282519283602484015260005b848110610b3e57505050601f19601f836000604480968601015201168101030190fd5b818101830151868201604401528593508201610b1b56fea2646970667358221220ac4310514ca2e32862c90f2c9c183da61993484b9eaa18bccd7851bfae55906464736f6c63430008190033a2646970667358221220dde426bfae5e63a5533a733570b5dd22065bdacb16058a625bf99369f569f92764736f6c63430008190033",
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
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220556184542fb9064583e2d89db785e7da0e61a5e40fc35c501465876b8dc8185a64736f6c63430008190033",
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
