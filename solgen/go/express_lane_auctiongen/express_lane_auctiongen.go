// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package express_lane_auctiongen

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

// Bid is an auto generated low-level Go binding around an user-defined struct.
type Bid struct {
	ExpressLaneController common.Address
	Amount                *big.Int
	Signature             []byte
}

// ELCRound is an auto generated low-level Go binding around an user-defined struct.
type ELCRound struct {
	ExpressLaneController common.Address
	Round                 uint64
}

// InitArgs is an auto generated low-level Go binding around an user-defined struct.
type InitArgs struct {
	Auctioneer              common.Address
	BiddingToken            common.Address
	Beneficiary             common.Address
	RoundTimingInfo         RoundTimingInfo
	MinReservePrice         *big.Int
	AuctioneerAdmin         common.Address
	MinReservePriceSetter   common.Address
	ReservePriceSetter      common.Address
	ReservePriceSetterAdmin common.Address
	BeneficiarySetter       common.Address
	RoundTimingSetter       common.Address
	MasterAdmin             common.Address
}

// RoundTimingInfo is an auto generated low-level Go binding around an user-defined struct.
type RoundTimingInfo struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}

// Transferor is an auto generated low-level Go binding around an user-defined struct.
type Transferor struct {
	Addr            common.Address
	FixedUntilRound uint64
}

// BalanceLibMetaData contains all meta data concerning the BalanceLib contract.
var BalanceLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212203398e9be405f49fb037dff58fafb86014420b29b99c037c4d7d5ee4473ed1bbb64736f6c63430008190033",
}

// BalanceLibABI is the input ABI used to generate the binding from.
// Deprecated: Use BalanceLibMetaData.ABI instead.
var BalanceLibABI = BalanceLibMetaData.ABI

// BalanceLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BalanceLibMetaData.Bin instead.
var BalanceLibBin = BalanceLibMetaData.Bin

// DeployBalanceLib deploys a new Ethereum contract, binding an instance of BalanceLib to it.
func DeployBalanceLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BalanceLib, error) {
	parsed, err := BalanceLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BalanceLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BalanceLib{BalanceLibCaller: BalanceLibCaller{contract: contract}, BalanceLibTransactor: BalanceLibTransactor{contract: contract}, BalanceLibFilterer: BalanceLibFilterer{contract: contract}}, nil
}

// BalanceLib is an auto generated Go binding around an Ethereum contract.
type BalanceLib struct {
	BalanceLibCaller     // Read-only binding to the contract
	BalanceLibTransactor // Write-only binding to the contract
	BalanceLibFilterer   // Log filterer for contract events
}

// BalanceLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceLibSession struct {
	Contract     *BalanceLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceLibCallerSession struct {
	Contract *BalanceLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BalanceLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceLibTransactorSession struct {
	Contract     *BalanceLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BalanceLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceLibRaw struct {
	Contract *BalanceLib // Generic contract binding to access the raw methods on
}

// BalanceLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceLibCallerRaw struct {
	Contract *BalanceLibCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceLibTransactorRaw struct {
	Contract *BalanceLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceLib creates a new instance of BalanceLib, bound to a specific deployed contract.
func NewBalanceLib(address common.Address, backend bind.ContractBackend) (*BalanceLib, error) {
	contract, err := bindBalanceLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceLib{BalanceLibCaller: BalanceLibCaller{contract: contract}, BalanceLibTransactor: BalanceLibTransactor{contract: contract}, BalanceLibFilterer: BalanceLibFilterer{contract: contract}}, nil
}

// NewBalanceLibCaller creates a new read-only instance of BalanceLib, bound to a specific deployed contract.
func NewBalanceLibCaller(address common.Address, caller bind.ContractCaller) (*BalanceLibCaller, error) {
	contract, err := bindBalanceLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceLibCaller{contract: contract}, nil
}

// NewBalanceLibTransactor creates a new write-only instance of BalanceLib, bound to a specific deployed contract.
func NewBalanceLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceLibTransactor, error) {
	contract, err := bindBalanceLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceLibTransactor{contract: contract}, nil
}

// NewBalanceLibFilterer creates a new log filterer instance of BalanceLib, bound to a specific deployed contract.
func NewBalanceLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceLibFilterer, error) {
	contract, err := bindBalanceLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceLibFilterer{contract: contract}, nil
}

// bindBalanceLib binds a generic wrapper to an already deployed contract.
func bindBalanceLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalanceLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceLib *BalanceLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceLib.Contract.BalanceLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceLib *BalanceLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceLib.Contract.BalanceLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceLib *BalanceLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceLib.Contract.BalanceLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceLib *BalanceLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceLib *BalanceLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceLib *BalanceLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceLib.Contract.contract.Transact(opts, method, params...)
}

// BurnerMetaData contains all meta data concerning the Burner contract.
var BurnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20BurnableUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a034608457601f61032a38819003918201601f19168301916001600160401b03831184841017608957808492602094604052833981010312608457516001600160a01b03811690819003608457801560725760805260405161028a90816100a082396080518181816063015260b40152f35b60405163d92e233d60e01b8152600490fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe60806040818152600436101561001457600080fd5b600091823560e01c90816344df8e701461008b575063fc0c546a1461003857600080fd5b346100875781600319360112610087576020905173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5080fd5b83833461008757816003193601126100875773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167f70a08231000000000000000000000000000000000000000000000000000000008452306004850152602084602481845afa9081156102485783916101b3575b839450803b156101ae5760248492845195869384927f42966c6800000000000000000000000000000000000000000000000000000000845260048401525af180156101a457610161578280f35b67ffffffffffffffff8211610177575281808280f35b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b81513d85823e3d90fd5b505050fd5b90506020933d602011610240575b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116810181811067ffffffffffffffff821117610213578185966020928652810103126101ae575190610114565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b3d94506101c1565b505051903d90823e3d90fdfea2646970667358221220ecc1b56457d6535a6c38aa4a04222c9c62ad6189b77788bdb18347ebeae112f764736f6c63430008190033",
}

// BurnerABI is the input ABI used to generate the binding from.
// Deprecated: Use BurnerMetaData.ABI instead.
var BurnerABI = BurnerMetaData.ABI

// BurnerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BurnerMetaData.Bin instead.
var BurnerBin = BurnerMetaData.Bin

// DeployBurner deploys a new Ethereum contract, binding an instance of Burner to it.
func DeployBurner(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Burner, error) {
	parsed, err := BurnerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnerBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Burner{BurnerCaller: BurnerCaller{contract: contract}, BurnerTransactor: BurnerTransactor{contract: contract}, BurnerFilterer: BurnerFilterer{contract: contract}}, nil
}

// Burner is an auto generated Go binding around an Ethereum contract.
type Burner struct {
	BurnerCaller     // Read-only binding to the contract
	BurnerTransactor // Write-only binding to the contract
	BurnerFilterer   // Log filterer for contract events
}

// BurnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnerSession struct {
	Contract     *Burner           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnerCallerSession struct {
	Contract *BurnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BurnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnerTransactorSession struct {
	Contract     *BurnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnerRaw struct {
	Contract *Burner // Generic contract binding to access the raw methods on
}

// BurnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnerCallerRaw struct {
	Contract *BurnerCaller // Generic read-only contract binding to access the raw methods on
}

// BurnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnerTransactorRaw struct {
	Contract *BurnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurner creates a new instance of Burner, bound to a specific deployed contract.
func NewBurner(address common.Address, backend bind.ContractBackend) (*Burner, error) {
	contract, err := bindBurner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Burner{BurnerCaller: BurnerCaller{contract: contract}, BurnerTransactor: BurnerTransactor{contract: contract}, BurnerFilterer: BurnerFilterer{contract: contract}}, nil
}

// NewBurnerCaller creates a new read-only instance of Burner, bound to a specific deployed contract.
func NewBurnerCaller(address common.Address, caller bind.ContractCaller) (*BurnerCaller, error) {
	contract, err := bindBurner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerCaller{contract: contract}, nil
}

// NewBurnerTransactor creates a new write-only instance of Burner, bound to a specific deployed contract.
func NewBurnerTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnerTransactor, error) {
	contract, err := bindBurner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerTransactor{contract: contract}, nil
}

// NewBurnerFilterer creates a new log filterer instance of Burner, bound to a specific deployed contract.
func NewBurnerFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnerFilterer, error) {
	contract, err := bindBurner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnerFilterer{contract: contract}, nil
}

// bindBurner binds a generic wrapper to an already deployed contract.
func bindBurner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Burner *BurnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Burner.Contract.BurnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Burner *BurnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burner.Contract.BurnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Burner *BurnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Burner.Contract.BurnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Burner *BurnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Burner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Burner *BurnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Burner *BurnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Burner.Contract.contract.Transact(opts, method, params...)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Burner *BurnerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Burner.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Burner *BurnerSession) Token() (common.Address, error) {
	return _Burner.Contract.Token(&_Burner.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Burner *BurnerCallerSession) Token() (common.Address, error) {
	return _Burner.Contract.Token(&_Burner.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_Burner *BurnerTransactor) Burn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Burner.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_Burner *BurnerSession) Burn() (*types.Transaction, error) {
	return _Burner.Contract.Burn(&_Burner.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_Burner *BurnerTransactorSession) Burn() (*types.Transaction, error) {
	return _Burner.Contract.Burn(&_Burner.TransactOpts)
}

// ExpressLaneAuctionMetaData contains all meta data concerning the ExpressLaneAuction contract.
var ExpressLaneAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AuctionNotClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidsWrongOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"name\":\"FixedTransferor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountRequested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountRequested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalanceAcc\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"currentRound\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newRound\",\"type\":\"uint64\"}],\"name\":\"InvalidNewRound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"currentStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newStart\",\"type\":\"uint64\"}],\"name\":\"InvalidNewStart\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NegativeOffset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"roundStart\",\"type\":\"int64\"}],\"name\":\"NegativeRoundStart\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotExpressLaneController\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"expectedTransferor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"NotTransferor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReserveBlackout\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"ReservePriceNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReservePrice\",\"type\":\"uint256\"}],\"name\":\"ReservePriceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"RoundAlreadyResolved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RoundDurationTooShort\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"RoundNotResolved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"RoundTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"currentRound\",\"type\":\"uint64\"}],\"name\":\"RoundTooOld\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameBidder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TieBidsWrongOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalMaxRound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAuctionClosingSeconds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroBiddingToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isMultiBidAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"firstPriceBidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"firstPriceExpressLaneController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstPriceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundStartTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundEndTimestamp\",\"type\":\"uint64\"}],\"name\":\"AuctionResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBeneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"SetBeneficiary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousExpressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newExpressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"}],\"name\":\"SetExpressLaneController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"SetMinReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReservePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReservePrice\",\"type\":\"uint256\"}],\"name\":\"SetReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"currentRound\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"name\":\"SetRoundTimingInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"name\":\"SetTransferor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundWithdrawable\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUCTIONEER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTIONEER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BENEFICIARY_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_RESERVE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_SETTER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUND_TIMING_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"balanceOfAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiaryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"biddingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flushBeneficiaryBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBidHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_auctioneer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_biddingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structRoundTimingInfo\",\"name\":\"_roundTimingInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_minReservePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_auctioneerAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minReservePriceSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePriceSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePriceSetterAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiarySetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_roundTimingSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_masterAdmin\",\"type\":\"address\"}],\"internalType\":\"structInitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiateWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAuctionRoundClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReserveBlackout\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minReservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"firstPriceBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"secondPriceBid\",\"type\":\"tuple\"}],\"name\":\"resolveMultiBidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"firstPriceBid\",\"type\":\"tuple\"}],\"name\":\"resolveSingleBidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolvedRounds\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"internalType\":\"structELCRound\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"internalType\":\"structELCRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"roundTimestamps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundTimingInfo\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinReservePrice\",\"type\":\"uint256\"}],\"name\":\"setMinReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReservePrice\",\"type\":\"uint256\"}],\"name\":\"setReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structRoundTimingInfo\",\"name\":\"newRoundTimingInfo\",\"type\":\"tuple\"}],\"name\":\"setRoundTimingInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"internalType\":\"structTransferor\",\"name\":\"transferor\",\"type\":\"tuple\"}],\"name\":\"setTransferor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newExpressLaneController\",\"type\":\"address\"}],\"name\":\"transferExpressLaneController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferorOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"withdrawableBalanceAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0806040523460225730608052614533908161002882396080518161106c0152f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80627be2fe14612ec45780630152682d14612e7b57806301ffc9a714612da757806302b6293814612d6257806304c584ad14612d335780630d253fbe14612c5757806314d9631614612c1c5780631682e50b14612be15780631c31f71014612a42578063248a9ca314612a135780632d668ce7146129f05780632f2ff15d14612927578063336a5b5e146128ec57806336568abe1461284d57806338af3eed14612825578063447a709e146123215780635633c337146122b9578063639d7566146122915780636a514beb146122435780636ad72517146122065780636dc4fc4e14611ea75780636e8cace514611df757806370a0823114611da75780637b617f9414611d6157806383af0a1f14611d425780638948cc4e14611d075780638a19c8bc14611cd75780639010d07c14611c9157806391d1485414611c435780639a1fadd31461101e578063a217fddf14611002578063b3ee252f14610fc7578063b51d1d4f14610ef7578063b6b55f2514610dd9578063bef0ec7414610c67578063c5b6aa2f14610b7e578063ca15c87314610b52578063ce9c7c0d146109d6578063cfe9232b1461099b578063d547741f1461095c578063db2e1eed1461093d578063e2fc6f681461091e578063e3f7bb55146108e3578063e460d2c51461089d578063e4d20c1d14610711578063f698da25146106ee5763fed87be81461022057600080fd5b3461042a57608060031936011261042a573360009081527fb75f9f346f983b8df66279e29440225cdc713a9aca5d61beaf930b9fd293970660209081526040909120547f6d8dad7188c7ed005c55bf77fbf589583d8668b0dad30a9b9dd016321a5c256f9060ff161561053f575061029661322a565b61029f81613522565b906102b16102ac366134bb565b613522565b9067ffffffffffffffff92838316848216036104fb576102f1916102d76102dd9261335c565b90613612565b50916102d76102eb366134bb565b9161335c565b5090828216838216036104b4575050604435908082169283830361042a57831561048a576024359382851680860361042a5762015180811161045957606435958487169384880361042a578286610348898b613375565b161161042f57600435918260070b9283810361042a578960a09961040d937fffffffffffffffff0000000000000000000000000000000000000000000000007f982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd479d60c01b168b85166fffffffffffffffff00000000000000008360401b161777ffffffffffffffff000000000000000000000000000000008460801b16171761010455604051936103f88561319d565b84528584015260408301526060820152613522565b9560405196168652850152604084015260608301526080820152a1005b600080fd5b60046040517f326de360000000000000000000000000000000000000000000000000000000008152fd5b602490604051907fc34a76cf0000000000000000000000000000000000000000000000000000000082526004820152fd5b60046040517f047bad52000000000000000000000000000000000000000000000000000000008152fd5b6040517fa0e269d800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff918216600482015291166024820152604490fd5b0390fd5b6040517f68c18ca900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff91821660048201529216602483015250604490fd5b61054833613fea565b604051906105558261319d565b604282528382019260603685378251156106d857603084538251600190600110156106d857607860218501536041905b80821161067e57505061063b576104f7936105df936106239361061460489460405197857f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008a978801528251928391603789019101613e47565b8401917f206973206d697373696e6720726f6c6520000000000000000000000000000000603784015251809386840190613e47565b01036028810184520182613207565b60405191829162461bcd60e51b835260048301613e6a565b6064846040519062461bcd60e51b825280600483015260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b9091600f811660108110156106d8576f181899199a1a9b1b9c1cb0b131b232b360811b901a6106ad8487613fd9565b5360041c9180156106c2576000190190610585565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b3461042a57600060031936011261042a576020610709613da6565b604051908152f35b3461042a5760208060031936011261042a576004357fb07567e7223e21f7dce4c0a89131ce9c32d0d3484085f3f331dea8caef56d1418060005260658352604060002033600052835260ff60406000205416156107b7575080917f5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0604061010392848454918351928352820152a1556101025481116107ac57005b6107b590613d2c565b005b90506107c233613fea565b604051906107cf8261319d565b604282528382019260603685378251156106d857603084538251600190600110156106d857607860218501536041905b80821161085957505061063b576104f7936105df936106239361061460489460405197857f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008a978801528251928391603789019101613e47565b9091600f811660108110156106d8576f181899199a1a9b1b9c1cb0b131b232b360811b901a6108888487613fd9565b5360041c9180156106c25760001901906107ff565b3461042a57600060031936011261042a5760206108d96108bb613c75565b5067ffffffffffffffff6108cd61322a565b915460a01c1690613ca8565b6040519015158152f35b3461042a57600060031936011261042a5760206040517fa8131bb4589277d6866d942849029b416b39e61eb3969a32787130bbdd292a968152f35b3461042a57600060031936011261042a57602061010554604051908152f35b3461042a57600060031936011261042a57602061010254604051908152f35b3461042a57604060031936011261042a576107b560043561097b613171565b908060005260656020526109966001604060002001546138c8565b613a46565b3461042a57600060031936011261042a5760206040517f1d693f62a755e2b3c6494da41af454605b9006057cb3c79b6adda1378f2a50a78152f35b3461042a5760208060031936011261042a577f19e6f23df7275b48d1c33822c6ad041a743378552246ac819f578ae1d6709cf98060005260658252604060002033600052825260ff6040600020541615610a6e57610a356108bb613c75565b610a44576107b5600435613d2c565b60046040517f4f006978000000000000000000000000000000000000000000000000000000008152fd5b610a7733613fea565b60405190610a848261319d565b604282528382019260603685378251156106d857603084538251600190600110156106d857607860218501536041905b808211610b0e57505061063b576104f7936105df936106239361061460489460405197857f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008a978801528251928391603789019101613e47565b9091600f811660108110156106d8576f181899199a1a9b1b9c1cb0b131b232b360811b901a610b3d8487613fd9565b5360041c9180156106c2576000190190610ab4565b3461042a57602060031936011261042a5760043560005260976020526020604060002054604051908152f35b3461042a57600060031936011261042a573360005260fd6020526040600020610ba86102ac61322a565b67ffffffffffffffff80821614610c3d57610bc390826136b4565b908115610c135760009055610be581336001600160a01b036101015416613bb4565b6040519081527f9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda860203392a2005b60046040517fd0d04f60000000000000000000000000000000000000000000000000000000008152fd5b60046040517f3d89ddde000000000000000000000000000000000000000000000000000000008152fd5b3461042a57604060031936011261042a573360005261010680602052604060002054906001600160a01b0391828116151580610db6575b610d78575033600052602052610d2c604060002082610cbb61347a565b1673ffffffffffffffffffffffffffffffffffffffff19825416178155610ce06134a4565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff000000000000000000000000000000000000000083549260a01b169116179055565b610d3461347a565b90610d3d6134a4565b9167ffffffffffffffff6040519316835216907ff6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c060203392a3005b60249067ffffffffffffffff604051917f75d899f200000000000000000000000000000000000000000000000000000000835260a01c166004820152fd5b5067ffffffffffffffff80610dcc6102ac61322a565b16908260a01c1611610c9e565b3461042a57602060031936011261042a576004353360005260fd60205260406000208115610ecd5760018101805467ffffffffffffffff908180821603610eb7575b505050610e29828254613b9a565b9055610e896001600160a01b036101015416604051907f23b872dd00000000000000000000000000000000000000000000000000000000602083015233602483015230604483015283606483015260648252610e84826131eb565b613e96565b6040519081527fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c60203392a2005b67ffffffffffffffff1916179055828080610e1b565b60046040517f1f2a2005000000000000000000000000000000000000000000000000000000008152fd5b3461042a57600060031936011261042a5767ffffffffffffffff600281610f1f6102ac61322a565b1601908082116106c2573360005260fd6020526040600020908154928315610ecd57811691818314610c3d576001019081549080821603610f9d5767ffffffffffffffff1983911617905560405191825260208201527f31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f540160403392a2005b60046040517f04eb6b3f000000000000000000000000000000000000000000000000000000008152fd5b3461042a57600060031936011261042a5760206040517f19e6f23df7275b48d1c33822c6ad041a743378552246ac819f578ae1d6709cf98152f35b3461042a57600060031936011261042a57602060405160008152f35b3461042a576101e060031936011261042a5760005460ff8160081c161590818092611c36575b8015611c1f575b15611bb55781600160ff19831617600055611b86575b506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163014611b1c5760ff60005460081c166110a481613c04565b604051906110b1826131cf565b60128252602082017f457870726573734c616e6541756374696f6e00000000000000000000000000008152604051926110e9846131cf565b6001845261112660208501937f3100000000000000000000000000000000000000000000000000000000000000855261112181613c04565b613c04565b519020915190209060c95560ca556001600160a01b0361114461344e565b1615611af2576001600160a01b0361115a61344e565b1673ffffffffffffffffffffffffffffffffffffffff199061010190828254161790556001600160a01b0361118d613464565b16610100918254161790557f8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e660406111c3613464565b6001600160a01b0382519160008352166020820152a17f9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794604060e43580610103557f5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f082805160008152836020820152a18061010255815190600082526020820152a160643560070b6064350361042a57600060643560070b12611ac85767ffffffffffffffff60a4351660a4350361042a5767ffffffffffffffff60a435161561048a5767ffffffffffffffff608435166084350361042a576201518067ffffffffffffffff6084351611611a8c5767ffffffffffffffff60c4351660c4350361042a5767ffffffffffffffff6084351667ffffffffffffffff6112eb60a43560c435613375565b161161042f577fffffffffffffffff00000000000000000000000000000000000000000000000060c43560c01b1667ffffffffffffffff606435166fffffffffffffffff000000000000000060843560401b161777ffffffffffffffff0000000000000000000000000000000060a43560801b1617176101045560807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c36011261042a577f982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd4760a06113e36040516113c18161319d565b6064358152608435602082015260a435604082015260c4356060820152613522565b67ffffffffffffffff6040519116815260643560070b602082015267ffffffffffffffff60843516604082015267ffffffffffffffff60a43516606082015267ffffffffffffffff60c435166080820152a16101c4356001600160a01b03811680910361042a5760008181527fffdfc1249c027f9191656349feb0761381bb32c9f557e01f419fd08754bf5a1b60205260409020546065919060ff1615611a3b575b6000805261149d6097918260205260406000206140fa565b50610124356001600160a01b03811680910361042a5761150c907fb07567e7223e21f7dce4c0a89131ce9c32d0d3484085f3f331dea8caef56d141806000528460205260406000208260005260205260ff60406000205416156119ea575b6000528260205260406000206140fa565b50610184356001600160a01b03811680910361042a5761157a907fc1b97c934675624ef2089089ac12ae8922988c11dc8a578dfbac10d9eecf4761806000528460205260406000208260005260205260ff60406000205416156119ea576000528260205260406000206140fa565b5061158361347a565b610104356001600160a01b03811680910361042a576115fd7f1d693f62a755e2b3c6494da41af454605b9006057cb3c79b6adda1378f2a50a79283600052856020526001600160a01b0360406000209116908160005260205260ff6040600020541615611999575b836000528460205260406000206140fa565b506116587f3fb9f0655b78e8eabe9e0f51d65db56c7690d4329012c3faf1fbd6d43f65826191826000528560205260406000208160005260205260ff6040600020541615611948575b826000528460205260406000206140fa565b508160005283602052600160406000200181815491557fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9283600080a461014435906001600160a01b03821680920361042a5761016435916001600160a01b03831680930361042a576117177f19e6f23df7275b48d1c33822c6ad041a743378552246ac819f578ae1d6709cf9918260005260406000208160005260205260ff60406000205416156118f7575b826000528560205260406000206140fa565b506117727fa8131bb4589277d6866d942849029b416b39e61eb3969a32787130bbdd292a9693846000528660205260406000208160005260205260ff60406000205416156118a6575b846000528560205260406000206140fa565b5080600052846020526001604060002001918383549355600080a46101a435906001600160a01b03821680920361042a576117f9927f6d8dad7188c7ed005c55bf77fbf589583d8668b0dad30a9b9dd016321a5c256f908160005260406000208460005260205260ff6040600020541615611855575b5060005260205260406000206140fa565b5061180057005b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff600054166000557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a1005b816000526020526040600020836000526020526040600020600160ff198254161790553383827f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4856117e8565b84600052866020526040600020816000526020526040600020600160ff198254161790553381867f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4611760565b82600052866020526040600020816000526020526040600020600160ff198254161790553381847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4611705565b82600052856020526040600020816000526020526040600020600160ff198254161790553381847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4611646565b83600052856020526040600020816000526020526040600020600160ff198254161790553381857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a46115eb565b80600052846020526040600020826000526020526040600020600160ff198254161790553382827f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a46114fb565b60008052816020526040600020816000526020526040600020600160ff19825416179055338160007f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4611485565b60246040517fc34a76cf00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff608435166004820152fd5b60046040517f16f46dfe000000000000000000000000000000000000000000000000000000008152fd5b60046040517f3fb3c7af000000000000000000000000000000000000000000000000000000008152fd5b608460405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201527f64656c656761746563616c6c00000000000000000000000000000000000000006064820152fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000166101011760005581611061565b608460405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152fd5b50303b15801561104b5750600160ff82161461104b565b50600160ff821610611044565b3461042a57604060031936011261042a57611c5c613171565b60043560005260656020526001600160a01b0360406000209116600052602052602060ff604060002054166040519015158152f35b3461042a57604060031936011261042a57600435600052609760205260206001600160a01b03611cc76024356040600020613fc1565b9190546040519260031b1c168152f35b3461042a57600060031936011261042a576020611cf56102ac61322a565b67ffffffffffffffff60405191168152f35b3461042a57600060031936011261042a5760206040517fb07567e7223e21f7dce4c0a89131ce9c32d0d3484085f3f331dea8caef56d1418152f35b3461042a57600060031936011261042a57602061010354604051908152f35b3461042a57602060031936011261042a57611d8a611d7d613143565b611d8561322a565b613612565b6040805167ffffffffffffffff9384168152919092166020820152f35b3461042a57602060031936011261042a576001600160a01b03611dc8613187565b1660005260fd60205260206107096040600020611df1611de96102ac61322a565b8254926136b4565b90613ba7565b3461042a57604060031936011261042a57611e10613187565b611e1861315a565b9067ffffffffffffffff80611e2e6102ac61322a565b1690831610611e58576020916001600160a01b03610709921660005260fd835260406000206136b4565b50611e646102ac61322a565b6040517f395f4fd600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff928316600482015291166024820152604490fd5b3461042a5760031960208136011261042a576004359067ffffffffffffffff80831161042a5760608360040192843603011261042a57611ee56136d8565b611eed61322a565b92611ef7846139f2565b156121dc57602401359261010254928385106121a557908291611f1982613522565b94611f40611f268761335c565b611f3981611f3436876133ad565b613ad0565b5094613612565b939092611f4c8861335c565b611f5582613490565b92611f5e613c75565b9190548984169a8b9160a01c16101561217457611f9c600160405193611f83856131cf565b6001600160a01b0380981685528c60208601521861331b565b92909261215e5785928a6020838661201b95511673ffffffffffffffffffffffffffffffffffffffff198554161784550151167fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff000000000000000000000000000000000000000083549260a01b169116179055565b16988960005260fd60205260406000209061204161203b835492846136b4565b82613ba7565b80158015612155575b61211e5750936120fd899692946000809c99957f7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d479c9761208d8760a09e9b613ba7565b905561010561209d878254613b9a565b9055817fb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b88876120cc87613490565b6040805167ffffffffffffffff97881681529387166020850152958e169583019590955290931692606090a4613490565b169a60405196875260208701526040860152166060840152166080820152a4005b86604491604051917fcf47918100000000000000000000000000000000000000000000000000000000835260048301526024820152fd5b5086811061204a565b634e487b7160e01b600052600060045260246000fd5b60248a604051907f451f87340000000000000000000000000000000000000000000000000000000082526004820152fd5b60448585604051917f56f9b75a00000000000000000000000000000000000000000000000000000000835260048301526024820152fd5b60046040517fb9adeefd000000000000000000000000000000000000000000000000000000008152fd5b3461042a57600060031936011261042a5761010580548015610ecd5760006107b592556001600160a01b0380610101541690610100541690613bb4565b3461042a57602060031936011261042a5760406001600160a01b0380612267613187565b1660005261010660205267ffffffffffffffff82600020548351928116835260a01c166020820152f35b3461042a57600060031936011261042a5760206001600160a01b036101015416604051908152f35b3461042a57604060031936011261042a576122d2613187565b6122da61315a565b9067ffffffffffffffff806122f06102ac61322a565b1690831610611e58576020916001600160a01b03610709921660005260fd8352611df16040600020918254926136b4565b3461042a5760031960408136011261042a5760043567ffffffffffffffff811161042a57606082823603011261042a576024359167ffffffffffffffff831161042a57606090833603011261042a576123786136d8565b61238061322a565b612389816139f2565b156121dc57602482013592602481013584106127fb5761010254806024830135106127c057506123b882613522565b926123c28461335c565b926123d484611f3436856004016133ad565b6123e886611f3497939736886004016133ad565b906001600160a01b0381166001600160a01b038916146127965760248701358a1492836126fc575b5050506126d25761242091613612565b92909161242c8661335c565b9561243982600401613490565b612441613c75565b905467ffffffffffffffff808b169160a01c16101561269757600161248e916001600160a01b0360405194612475866131cf565b16845267ffffffffffffffff8b1660208501521861331b565b61215e5767ffffffffffffffff6020836001600160a01b0361251795511673ffffffffffffffffffffffffffffffffffffffff198554161784550151167fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff000000000000000000000000000000000000000083549260a01b169116179055565b6001600160a01b03861660005260fd60205260406000209061253e61203b835492846136b4565b8015801561268a575b61264f57509467ffffffffffffffff856024866001600160a01b0361262860019c9860008f817f7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d479f60a09f879f938c9e6125c88f96948d7fb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b96013590613ba7565b90556101056125db8c8c01358254613b9a565b905561261d896125ed89600401613490565b169560405193849384916040919493606084019567ffffffffffffffff9283809216865216602085015216910152565b0390a4600401613490565b169c846040519d168d5260208d0152013560408b01521660608901521660808701521693a4005b846044916024604051927fcf479181000000000000000000000000000000000000000000000000000000008452013560048301526024820152fd5b5060248501358110612547565b60248967ffffffffffffffff604051917f451f8734000000000000000000000000000000000000000000000000000000008352166004820152fd5b60046040517f9185a0ae000000000000000000000000000000000000000000000000000000008152fd5b61277e9192935061278a90604051908161273860208201928d8490916bffffffffffffffffffffffff1960349360601b16825260148201520190565b039161274c601f1993848101835282613207565b5190209460405193849160208301968790916bffffffffffffffffffffffff1960349360601b16825260148201520190565b03908101835282613207565b51902011888080612410565b60046040517ff4a3e485000000000000000000000000000000000000000000000000000000008152fd5b604492506024604051927f56f9b75a000000000000000000000000000000000000000000000000000000008452013560048301526024820152fd5b60046040517fa234cb19000000000000000000000000000000000000000000000000000000008152fd5b3461042a57600060031936011261042a5760206001600160a01b036101005416604051908152f35b3461042a57604060031936011261042a57612866613171565b336001600160a01b03821603612882576107b590600435613a46565b608460405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c6600000000000000000000000000000000006064820152fd5b3461042a57600060031936011261042a5760206040517fc1b97c934675624ef2089089ac12ae8922988c11dc8a578dfbac10d9eecf47618152f35b3461042a57604060031936011261042a576107b56004356097612948613171565b9180600052602090606582526129656001604060002001546138c8565b80600052606582526001600160a01b03604060002094169384600052825260ff60406000205416156129a0575b6000525260406000206140fa565b806000526065825260406000208460005282526040600020600160ff198254161790553384827f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4612992565b3461042a57600060031936011261042a5760206108d9612a0e61322a565b6139f2565b3461042a57602060031936011261042a5760043560005260656020526020600160406000200154604051908152f35b3461042a5760208060031936011261042a57612a5c613187565b907fc1b97c934675624ef2089089ac12ae8922988c11dc8a578dfbac10d9eecf47618060005260658252604060002033600052825260ff6040600020541615612afd575073ffffffffffffffffffffffffffffffffffffffff19610100927f8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e660408554926001600160a01b038251918186168352168096820152a116179055005b612b0633613fea565b60405190612b138261319d565b604282528382019260603685378251156106d857603084538251600190600110156106d857607860218501536041905b808211612b9d57505061063b576104f7936105df936106239361061460489460405197857f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008a978801528251928391603789019101613e47565b9091600f811660108110156106d8576f181899199a1a9b1b9c1cb0b131b232b360811b901a612bcc8487613fd9565b5360041c9180156106c2576000190190612b43565b3461042a57600060031936011261042a5760206040517f6d8dad7188c7ed005c55bf77fbf589583d8668b0dad30a9b9dd016321a5c256f8152f35b3461042a57600060031936011261042a5760206040517f3fb9f0655b78e8eabe9e0f51d65db56c7690d4329012c3faf1fbd6d43f6582618152f35b3461042a57600060031936011261042a5760006020604051612c78816131cf565b828152015260006020604051612c8d816131cf565b828152015267ffffffffffffffff60fe54908060ff5460a01c169160a01c1611600014612d1f57608060ff612d1d612cce612cc860fe61332d565b9261332d565b612cf8604051809467ffffffffffffffff602080926001600160a01b038151168552015116910152565b80516001600160a01b031660408401526020015167ffffffffffffffff166060830152565bf35b608060fe612d1d612cce612cc860ff61332d565b3461042a57606060031936011261042a576020610709612d51613143565b612d59613171565b6044359161326a565b3461042a57602060031936011261042a576001600160a01b03612d83613187565b1660005260fd60205260206107096040600020612da16102ac61322a565b906136b4565b3461042a57602060031936011261042a576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361042a57807f5a05180f0000000000000000000000000000000000000000000000000000000060209214908115612e1e575b506040519015158152f35b7f7965db0b00000000000000000000000000000000000000000000000000000000811491508115612e51575b5082612e13565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501482612e4a565b3461042a57600060031936011261042a5760806101045467ffffffffffffffff90604051918160070b8352808260401c16602084015281841c16604083015260c01c6060820152f35b3461042a57604060031936011261042a57612edd613143565b612ee5613171565b612eed61322a565b90612ef782613522565b9167ffffffffffffffff92838116848616106130fe5750612f1784613577565b91825493856001600160a01b03938487169586600052610106602052856040600020541697881515958660001461305557338a03613005579373ffffffffffffffffffffffffffffffffffffffff19999388979693612fa4937fb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b9a612ff1985b169c8d9116179055613612565b929094600014612ffd57935b8142168092821610600014612ff657505b604051948594169884916040919493606084019567ffffffffffffffff9283809216865216602085015216910152565b0390a4005b9050612fc1565b508693612fb0565b6040517f7621d94a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff861660048201526001600160a01b038b166024820152336044820152606490fd5b9895969798338a036130ae57938896937fb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b989693612fa49373ffffffffffffffffffffffffffffffffffffffff199c9b612ff198612f97565b6040517f660af6d200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff861660048201526001600160a01b038b166024820152336044820152606490fd5b6040517f395f4fd600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8681166004830152919091166024820152604490fd5b6004359067ffffffffffffffff8216820361042a57565b6024359067ffffffffffffffff8216820361042a57565b602435906001600160a01b038216820361042a57565b600435906001600160a01b038216820361042a57565b6080810190811067ffffffffffffffff8211176131b957604052565b634e487b7160e01b600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176131b957604052565b60a0810190811067ffffffffffffffff8211176131b957604052565b90601f601f19910116810190811067ffffffffffffffff8211176131b957604052565b604051906132378261319d565b816060610104548060070b835267ffffffffffffffff808260401c1660208501528160801c16604084015260c01c910152565b916001600160a01b036040519267ffffffffffffffff60208501957f0358b2b705d5c5ef47651be44f418326852a390f3b4c933661a5f4f0d8fa1ee387521660408501521660608301526080820152608081526132c6816131eb565b5190206132d1613da6565b906040519060208201927f1901000000000000000000000000000000000000000000000000000000000000845260228301526042820152604281526133158161319d565b51902090565b60028110156106d85760fe0190600090565b9060405161333a816131cf565b91546001600160a01b038116835260a01c67ffffffffffffffff166020830152565b90600167ffffffffffffffff809316019182116106c257565b91909167ffffffffffffffff808094169116019182116106c257565b67ffffffffffffffff81116131b957601f01601f191660200190565b919060608382031261042a5760405167ffffffffffffffff939060608101858111828210176131b957604052809482356001600160a01b038116810361042a5782526020928381013584840152604081013591821161042a570183601f8201121561042a5780359261341e84613391565b9461342c6040519687613207565b84865281858401011161042a5783604094826000940183880137850101520152565b6024356001600160a01b038116810361042a5790565b6044356001600160a01b038116810361042a5790565b6004356001600160a01b038116810361042a5790565b356001600160a01b038116810361042a5790565b60243567ffffffffffffffff8116810361042a5790565b600319608091011261042a57604051906134d48261319d565b816004358060070b810361042a57815267ffffffffffffffff90602435828116810361042a576020820152604435828116810361042a576040820152606435918216820361042a5760600152565b805160070b67ffffffffffffffff908082421660070b1261356f5760206135498392613e12565b9301511691821561355957160490565b634e487b7160e01b600052601260045260246000fd5b505050600090565b60fe5467ffffffffffffffff918216919060a01c8116820361359a57505060fe90565b60ff5460a01c1681036135ad575060ff90565b602490604051907ffbb052d80000000000000000000000000000000000000000000000000000000082526004820152fd5b9060001967ffffffffffffffff809316019182116106c257565b67ffffffffffffffff91821690821603919082116106c257565b91906020835193019267ffffffffffffffff9182808651169116028281169081036106c25760070b9060070b01677fffffffffffffff198112677fffffffffffffff8213176106c2578060070b6000811261368357508161367b91613680931694511684613375565b6135de565b90565b602490604051907ff160ad790000000000000000000000000000000000000000000000000000000082526004820152fd5b6001810154909167ffffffffffffffff9182169116106136d2575490565b50600090565b3360009081527f7662a4858c3782df3275efafb25e2dcddc1f1962914c08167ab217a9d5a739ad60209081526040808320549092907f1d693f62a755e2b3c6494da41af454605b9006057cb3c79b6adda1378f2a50a79060ff161561373d5750505050565b61374633613fea565b8451916137528361319d565b604283528483019360603686378351156138b45760308553835190600191600110156138b45790607860218601536041915b81831161384657505050613804576105df9385936137ee936137df6048946104f7995198857f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008b978801528251928391603789019101613e47565b01036028810185520183613207565b5191829162461bcd60e51b835260048301613e6a565b60648486519062461bcd60e51b825280600483015260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b909192600f811660108110156138a0576f181899199a1a9b1b9c1cb0b131b232b360811b901a6138768588613fd9565b5360041c92801561388c57600019019190613784565b602482634e487b7160e01b81526011600452fd5b602483634e487b7160e01b81526032600452fd5b80634e487b7160e01b602492526032600452fd5b600081815260209060658252604092838220338352835260ff8483205416156138f15750505050565b6138fa33613fea565b8451916139068361319d565b604283528483019360603686378351156138b45760308553835190600191600110156138b45790607860218601536041915b81831161399357505050613804576105df9385936137ee936137df6048946104f7995198857f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008b978801528251928391603789019101613e47565b909192600f811660108110156138a0576f181899199a1a9b1b9c1cb0b131b232b360811b901a6139c38588613fd9565b5360041c92801561388c57600019019190613938565b9067ffffffffffffffff80911691821561355957160690565b805160070b67ffffffffffffffff908082421660070b1261356f57613a3e613a1a8392613e12565b9382604081613a316020850198828a5116906139d9565b97511692015116906135f8565b169116101590565b906040613a859260009080825260656020526001600160a01b0383832094169384835260205260ff8383205416613a88575b8152609760205220614184565b50565b808252606560205282822084835260205282822060ff1981541690553384827ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b8580a4613a78565b919061367b906001600160a01b0390613b0e8286511691613b066040613afd60208a01958651908561326a565b98015188614274565b9590956142ac565b918316918260005260fd602052613b2f6040600020611df1838254926136b4565b9151809210613b3f575050509190565b60649350613b61908360005260fd602052611df16040600020918254926136b4565b90604051927f36b24c14000000000000000000000000000000000000000000000000000000008452600484015260248301526044820152fd5b919082018092116106c257565b919082039182116106c257565b613c02926001600160a01b03604051937fa9059cbb000000000000000000000000000000000000000000000000000000006020860152166024840152604483015260448252610e848261319d565b565b15613c0b57565b608460405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152fd5b60fe9060009167ffffffffffffffff60fe54908060ff5460a01c169160a01c1610613c9e579190565b5060ff9150600190565b9081519067ffffffffffffffff9160070b82421660070b1261356f578180613cd7613cd286613522565b61335c565b1691161015613d255780613a3e613cf1845160070b613e12565b93826060613d1b82613d0b6020860199828b5116906139d9565b98511683604086015116906135f8565b92015116906135f8565b5050600090565b61010354808210613d6f57506101027f9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794604082548151908152846020820152a155565b60449250604051917fda4f272e00000000000000000000000000000000000000000000000000000000835260048301526024820152fd5b60c95460ca546040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815260c0810181811067ffffffffffffffff8211176131b95760405251902090565b67ffffffffffffffff9060070b81421660070b0390677fffffffffffffff8213677fffffffffffffff198312176106c2571690565b60005b838110613e5a5750506000910152565b8181015183820152602001613e4a565b601f19601f60409360208452613e8f8151809281602088015260208888019101613e47565b0116010190565b6001600160a01b031690613f14604051613eaf816131cf565b6020938482527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858301526000808587829751910182855af13d15613fb9573d91613ef983613391565b92613f076040519485613207565b83523d868885013e61446d565b805180613f22575b50505050565b81849181010312613fb55782015190811591821503613fb25750613f4857808080613f1c565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b80fd5b5080fd5b60609161446d565b80548210156106d85760005260206000200190600090565b9081518110156106d8570160200190565b604051906060820182811067ffffffffffffffff8211176131b957604052602a82526020820160403682378251156106d857603090538151600190600110156106d857607860218401536029905b80821161408c5750506140485790565b606460405162461bcd60e51b815260206004820152602060248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152fd5b9091600f811660108110156140e5576f181899199a1a9b1b9c1cb0b131b232b360811b901a6140bb8486613fd9565b5360041c9180156140d0576000190190614038565b60246000634e487b7160e01b81526011600452fd5b60246000634e487b7160e01b81526032600452fd5b9190600183016000908282528060205260408220541560001461417e578454946801000000000000000086101561416a578361415a614143886001604098999a01855584613fc1565b81939154906000199060031b92831b921b19161790565b9055549382526020522055600190565b602483634e487b7160e01b81526041600452fd5b50925050565b9060018201906000928184528260205260408420549081151560001461426d57600019918281018181116142595782549084820191821161424557808203614210575b505050805480156141fc578201916141df8383613fc1565b909182549160031b1b191690555582526020526040812055600190565b602486634e487b7160e01b81526031600452fd5b6142306142206141439386613fc1565b90549060031b1c92839286613fc1565b905586528460205260408620553880806141c7565b602488634e487b7160e01b81526011600452fd5b602487634e487b7160e01b81526011600452fd5b5050505090565b9060418151146000146142a25761429e916020820151906060604084015193015160001a906143de565b9091565b5050600090600290565b60058110156143c857806142bd5750565b6001810361430957606460405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152fd5b6002810361435557606460405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152fd5b60031461435e57565b608460405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152fd5b634e487b7160e01b600052602160045260246000fd5b9291907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083116144615791608094939160ff602094604051948552168484015260408301526060820152600093849182805260015afa156144545781516001600160a01b0381161561444e579190565b50600190565b50604051903d90823e3d90fd5b50505050600090600390565b919290156144ce5750815115614481575090565b3b1561448a5790565b606460405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156144e15750805190602001fd5b6104f79060405191829162461bcd60e51b835260048301613e6a56fea2646970667358221220fb51c5841dce1fd5dc96549e9fa5f46850b2275e640bbcef1c3a01d221d5a4c064736f6c63430008190033",
}

// ExpressLaneAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use ExpressLaneAuctionMetaData.ABI instead.
var ExpressLaneAuctionABI = ExpressLaneAuctionMetaData.ABI

// ExpressLaneAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExpressLaneAuctionMetaData.Bin instead.
var ExpressLaneAuctionBin = ExpressLaneAuctionMetaData.Bin

// DeployExpressLaneAuction deploys a new Ethereum contract, binding an instance of ExpressLaneAuction to it.
func DeployExpressLaneAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExpressLaneAuction, error) {
	parsed, err := ExpressLaneAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExpressLaneAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExpressLaneAuction{ExpressLaneAuctionCaller: ExpressLaneAuctionCaller{contract: contract}, ExpressLaneAuctionTransactor: ExpressLaneAuctionTransactor{contract: contract}, ExpressLaneAuctionFilterer: ExpressLaneAuctionFilterer{contract: contract}}, nil
}

// ExpressLaneAuction is an auto generated Go binding around an Ethereum contract.
type ExpressLaneAuction struct {
	ExpressLaneAuctionCaller     // Read-only binding to the contract
	ExpressLaneAuctionTransactor // Write-only binding to the contract
	ExpressLaneAuctionFilterer   // Log filterer for contract events
}

// ExpressLaneAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExpressLaneAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpressLaneAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExpressLaneAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpressLaneAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExpressLaneAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpressLaneAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExpressLaneAuctionSession struct {
	Contract     *ExpressLaneAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExpressLaneAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExpressLaneAuctionCallerSession struct {
	Contract *ExpressLaneAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ExpressLaneAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExpressLaneAuctionTransactorSession struct {
	Contract     *ExpressLaneAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ExpressLaneAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExpressLaneAuctionRaw struct {
	Contract *ExpressLaneAuction // Generic contract binding to access the raw methods on
}

// ExpressLaneAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExpressLaneAuctionCallerRaw struct {
	Contract *ExpressLaneAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// ExpressLaneAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExpressLaneAuctionTransactorRaw struct {
	Contract *ExpressLaneAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExpressLaneAuction creates a new instance of ExpressLaneAuction, bound to a specific deployed contract.
func NewExpressLaneAuction(address common.Address, backend bind.ContractBackend) (*ExpressLaneAuction, error) {
	contract, err := bindExpressLaneAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuction{ExpressLaneAuctionCaller: ExpressLaneAuctionCaller{contract: contract}, ExpressLaneAuctionTransactor: ExpressLaneAuctionTransactor{contract: contract}, ExpressLaneAuctionFilterer: ExpressLaneAuctionFilterer{contract: contract}}, nil
}

// NewExpressLaneAuctionCaller creates a new read-only instance of ExpressLaneAuction, bound to a specific deployed contract.
func NewExpressLaneAuctionCaller(address common.Address, caller bind.ContractCaller) (*ExpressLaneAuctionCaller, error) {
	contract, err := bindExpressLaneAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionCaller{contract: contract}, nil
}

// NewExpressLaneAuctionTransactor creates a new write-only instance of ExpressLaneAuction, bound to a specific deployed contract.
func NewExpressLaneAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*ExpressLaneAuctionTransactor, error) {
	contract, err := bindExpressLaneAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionTransactor{contract: contract}, nil
}

// NewExpressLaneAuctionFilterer creates a new log filterer instance of ExpressLaneAuction, bound to a specific deployed contract.
func NewExpressLaneAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*ExpressLaneAuctionFilterer, error) {
	contract, err := bindExpressLaneAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionFilterer{contract: contract}, nil
}

// bindExpressLaneAuction binds a generic wrapper to an already deployed contract.
func bindExpressLaneAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExpressLaneAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExpressLaneAuction *ExpressLaneAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExpressLaneAuction.Contract.ExpressLaneAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExpressLaneAuction *ExpressLaneAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ExpressLaneAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExpressLaneAuction *ExpressLaneAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ExpressLaneAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExpressLaneAuction *ExpressLaneAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExpressLaneAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.contract.Transact(opts, method, params...)
}

// AUCTIONEERADMINROLE is a free data retrieval call binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) AUCTIONEERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "AUCTIONEER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUCTIONEERADMINROLE is a free data retrieval call binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) AUCTIONEERADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.AUCTIONEERADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// AUCTIONEERADMINROLE is a free data retrieval call binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) AUCTIONEERADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.AUCTIONEERADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// AUCTIONEERROLE is a free data retrieval call binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) AUCTIONEERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "AUCTIONEER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUCTIONEERROLE is a free data retrieval call binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) AUCTIONEERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.AUCTIONEERROLE(&_ExpressLaneAuction.CallOpts)
}

// AUCTIONEERROLE is a free data retrieval call binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) AUCTIONEERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.AUCTIONEERROLE(&_ExpressLaneAuction.CallOpts)
}

// BENEFICIARYSETTERROLE is a free data retrieval call binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BENEFICIARYSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "BENEFICIARY_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BENEFICIARYSETTERROLE is a free data retrieval call binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BENEFICIARYSETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.BENEFICIARYSETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// BENEFICIARYSETTERROLE is a free data retrieval call binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BENEFICIARYSETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.BENEFICIARYSETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.DEFAULTADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.DEFAULTADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// MINRESERVESETTERROLE is a free data retrieval call binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) MINRESERVESETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "MIN_RESERVE_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINRESERVESETTERROLE is a free data retrieval call binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) MINRESERVESETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.MINRESERVESETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// MINRESERVESETTERROLE is a free data retrieval call binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) MINRESERVESETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.MINRESERVESETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// RESERVESETTERADMINROLE is a free data retrieval call binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) RESERVESETTERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "RESERVE_SETTER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESERVESETTERADMINROLE is a free data retrieval call binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RESERVESETTERADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.RESERVESETTERADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// RESERVESETTERADMINROLE is a free data retrieval call binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) RESERVESETTERADMINROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.RESERVESETTERADMINROLE(&_ExpressLaneAuction.CallOpts)
}

// RESERVESETTERROLE is a free data retrieval call binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) RESERVESETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "RESERVE_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESERVESETTERROLE is a free data retrieval call binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RESERVESETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.RESERVESETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// RESERVESETTERROLE is a free data retrieval call binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) RESERVESETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.RESERVESETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// ROUNDTIMINGSETTERROLE is a free data retrieval call binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) ROUNDTIMINGSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "ROUND_TIMING_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUNDTIMINGSETTERROLE is a free data retrieval call binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ROUNDTIMINGSETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.ROUNDTIMINGSETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// ROUNDTIMINGSETTERROLE is a free data retrieval call binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) ROUNDTIMINGSETTERROLE() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.ROUNDTIMINGSETTERROLE(&_ExpressLaneAuction.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BalanceOf(&_ExpressLaneAuction.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BalanceOf(&_ExpressLaneAuction.CallOpts, account)
}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BalanceOfAtRound(opts *bind.CallOpts, account common.Address, round uint64) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "balanceOfAtRound", account, round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BalanceOfAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BalanceOfAtRound(&_ExpressLaneAuction.CallOpts, account, round)
}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BalanceOfAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BalanceOfAtRound(&_ExpressLaneAuction.CallOpts, account, round)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) Beneficiary() (common.Address, error) {
	return _ExpressLaneAuction.Contract.Beneficiary(&_ExpressLaneAuction.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) Beneficiary() (common.Address, error) {
	return _ExpressLaneAuction.Contract.Beneficiary(&_ExpressLaneAuction.CallOpts)
}

// BeneficiaryBalance is a free data retrieval call binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BeneficiaryBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "beneficiaryBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BeneficiaryBalance is a free data retrieval call binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BeneficiaryBalance() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BeneficiaryBalance(&_ExpressLaneAuction.CallOpts)
}

// BeneficiaryBalance is a free data retrieval call binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BeneficiaryBalance() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.BeneficiaryBalance(&_ExpressLaneAuction.CallOpts)
}

// BiddingToken is a free data retrieval call binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) BiddingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "biddingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BiddingToken is a free data retrieval call binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) BiddingToken() (common.Address, error) {
	return _ExpressLaneAuction.Contract.BiddingToken(&_ExpressLaneAuction.CallOpts)
}

// BiddingToken is a free data retrieval call binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) BiddingToken() (common.Address, error) {
	return _ExpressLaneAuction.Contract.BiddingToken(&_ExpressLaneAuction.CallOpts)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) CurrentRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "currentRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) CurrentRound() (uint64, error) {
	return _ExpressLaneAuction.Contract.CurrentRound(&_ExpressLaneAuction.CallOpts)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) CurrentRound() (uint64, error) {
	return _ExpressLaneAuction.Contract.CurrentRound(&_ExpressLaneAuction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) DomainSeparator() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.DomainSeparator(&_ExpressLaneAuction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) DomainSeparator() ([32]byte, error) {
	return _ExpressLaneAuction.Contract.DomainSeparator(&_ExpressLaneAuction.CallOpts)
}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) GetBidHash(opts *bind.CallOpts, round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "getBidHash", round, expressLaneController, amount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GetBidHash(round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	return _ExpressLaneAuction.Contract.GetBidHash(&_ExpressLaneAuction.CallOpts, round, expressLaneController, amount)
}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) GetBidHash(round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	return _ExpressLaneAuction.Contract.GetBidHash(&_ExpressLaneAuction.CallOpts, round, expressLaneController, amount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ExpressLaneAuction.Contract.GetRoleAdmin(&_ExpressLaneAuction.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ExpressLaneAuction.Contract.GetRoleAdmin(&_ExpressLaneAuction.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ExpressLaneAuction.Contract.GetRoleMember(&_ExpressLaneAuction.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ExpressLaneAuction.Contract.GetRoleMember(&_ExpressLaneAuction.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.GetRoleMemberCount(&_ExpressLaneAuction.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.GetRoleMemberCount(&_ExpressLaneAuction.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ExpressLaneAuction.Contract.HasRole(&_ExpressLaneAuction.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ExpressLaneAuction.Contract.HasRole(&_ExpressLaneAuction.CallOpts, role, account)
}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) IsAuctionRoundClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "isAuctionRoundClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) IsAuctionRoundClosed() (bool, error) {
	return _ExpressLaneAuction.Contract.IsAuctionRoundClosed(&_ExpressLaneAuction.CallOpts)
}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) IsAuctionRoundClosed() (bool, error) {
	return _ExpressLaneAuction.Contract.IsAuctionRoundClosed(&_ExpressLaneAuction.CallOpts)
}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) IsReserveBlackout(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "isReserveBlackout")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) IsReserveBlackout() (bool, error) {
	return _ExpressLaneAuction.Contract.IsReserveBlackout(&_ExpressLaneAuction.CallOpts)
}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) IsReserveBlackout() (bool, error) {
	return _ExpressLaneAuction.Contract.IsReserveBlackout(&_ExpressLaneAuction.CallOpts)
}

// MinReservePrice is a free data retrieval call binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) MinReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "minReservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReservePrice is a free data retrieval call binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) MinReservePrice() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.MinReservePrice(&_ExpressLaneAuction.CallOpts)
}

// MinReservePrice is a free data retrieval call binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) MinReservePrice() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.MinReservePrice(&_ExpressLaneAuction.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) ReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "reservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ReservePrice() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.ReservePrice(&_ExpressLaneAuction.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) ReservePrice() (*big.Int, error) {
	return _ExpressLaneAuction.Contract.ReservePrice(&_ExpressLaneAuction.CallOpts)
}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) ResolvedRounds(opts *bind.CallOpts) (ELCRound, ELCRound, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "resolvedRounds")

	if err != nil {
		return *new(ELCRound), *new(ELCRound), err
	}

	out0 := *abi.ConvertType(out[0], new(ELCRound)).(*ELCRound)
	out1 := *abi.ConvertType(out[1], new(ELCRound)).(*ELCRound)

	return out0, out1, err

}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ResolvedRounds() (ELCRound, ELCRound, error) {
	return _ExpressLaneAuction.Contract.ResolvedRounds(&_ExpressLaneAuction.CallOpts)
}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) ResolvedRounds() (ELCRound, ELCRound, error) {
	return _ExpressLaneAuction.Contract.ResolvedRounds(&_ExpressLaneAuction.CallOpts)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64, uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) RoundTimestamps(opts *bind.CallOpts, round uint64) (uint64, uint64, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "roundTimestamps", round)

	if err != nil {
		return *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64, uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RoundTimestamps(round uint64) (uint64, uint64, error) {
	return _ExpressLaneAuction.Contract.RoundTimestamps(&_ExpressLaneAuction.CallOpts, round)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64, uint64)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) RoundTimestamps(round uint64) (uint64, uint64, error) {
	return _ExpressLaneAuction.Contract.RoundTimestamps(&_ExpressLaneAuction.CallOpts, round)
}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) RoundTimingInfo(opts *bind.CallOpts) (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "roundTimingInfo")

	outstruct := new(struct {
		OffsetTimestamp          int64
		RoundDurationSeconds     uint64
		AuctionClosingSeconds    uint64
		ReserveSubmissionSeconds uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OffsetTimestamp = *abi.ConvertType(out[0], new(int64)).(*int64)
	outstruct.RoundDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.AuctionClosingSeconds = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.ReserveSubmissionSeconds = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RoundTimingInfo() (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	return _ExpressLaneAuction.Contract.RoundTimingInfo(&_ExpressLaneAuction.CallOpts)
}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) RoundTimingInfo() (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	return _ExpressLaneAuction.Contract.RoundTimingInfo(&_ExpressLaneAuction.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExpressLaneAuction.Contract.SupportsInterface(&_ExpressLaneAuction.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExpressLaneAuction.Contract.SupportsInterface(&_ExpressLaneAuction.CallOpts, interfaceId)
}

// TransferorOf is a free data retrieval call binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address ) view returns(address addr, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) TransferorOf(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr            common.Address
	FixedUntilRound uint64
}, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "transferorOf", arg0)

	outstruct := new(struct {
		Addr            common.Address
		FixedUntilRound uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FixedUntilRound = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// TransferorOf is a free data retrieval call binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address ) view returns(address addr, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) TransferorOf(arg0 common.Address) (struct {
	Addr            common.Address
	FixedUntilRound uint64
}, error) {
	return _ExpressLaneAuction.Contract.TransferorOf(&_ExpressLaneAuction.CallOpts, arg0)
}

// TransferorOf is a free data retrieval call binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address ) view returns(address addr, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) TransferorOf(arg0 common.Address) (struct {
	Addr            common.Address
	FixedUntilRound uint64
}, error) {
	return _ExpressLaneAuction.Contract.TransferorOf(&_ExpressLaneAuction.CallOpts, arg0)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) WithdrawableBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "withdrawableBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) WithdrawableBalance(account common.Address) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.WithdrawableBalance(&_ExpressLaneAuction.CallOpts, account)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) WithdrawableBalance(account common.Address) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.WithdrawableBalance(&_ExpressLaneAuction.CallOpts, account)
}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCaller) WithdrawableBalanceAtRound(opts *bind.CallOpts, account common.Address, round uint64) (*big.Int, error) {
	var out []interface{}
	err := _ExpressLaneAuction.contract.Call(opts, &out, "withdrawableBalanceAtRound", account, round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionSession) WithdrawableBalanceAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.WithdrawableBalanceAtRound(&_ExpressLaneAuction.CallOpts, account, round)
}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_ExpressLaneAuction *ExpressLaneAuctionCallerSession) WithdrawableBalanceAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _ExpressLaneAuction.Contract.WithdrawableBalanceAtRound(&_ExpressLaneAuction.CallOpts, account, round)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.Deposit(&_ExpressLaneAuction.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.Deposit(&_ExpressLaneAuction.TransactOpts, amount)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) FinalizeWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "finalizeWithdrawal")
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) FinalizeWithdrawal() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.FinalizeWithdrawal(&_ExpressLaneAuction.TransactOpts)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) FinalizeWithdrawal() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.FinalizeWithdrawal(&_ExpressLaneAuction.TransactOpts)
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) FlushBeneficiaryBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "flushBeneficiaryBalance")
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) FlushBeneficiaryBalance() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.FlushBeneficiaryBalance(&_ExpressLaneAuction.TransactOpts)
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) FlushBeneficiaryBalance() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.FlushBeneficiaryBalance(&_ExpressLaneAuction.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.GrantRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.GrantRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) Initialize(opts *bind.TransactOpts, args InitArgs) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "initialize", args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) Initialize(args InitArgs) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.Initialize(&_ExpressLaneAuction.TransactOpts, args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) Initialize(args InitArgs) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.Initialize(&_ExpressLaneAuction.TransactOpts, args)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) InitiateWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "initiateWithdrawal")
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.InitiateWithdrawal(&_ExpressLaneAuction.TransactOpts)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.InitiateWithdrawal(&_ExpressLaneAuction.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.RenounceRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.RenounceRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) ResolveMultiBidAuction(opts *bind.TransactOpts, firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "resolveMultiBidAuction", firstPriceBid, secondPriceBid)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ResolveMultiBidAuction(firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ResolveMultiBidAuction(&_ExpressLaneAuction.TransactOpts, firstPriceBid, secondPriceBid)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) ResolveMultiBidAuction(firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ResolveMultiBidAuction(&_ExpressLaneAuction.TransactOpts, firstPriceBid, secondPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) ResolveSingleBidAuction(opts *bind.TransactOpts, firstPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "resolveSingleBidAuction", firstPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) ResolveSingleBidAuction(firstPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ResolveSingleBidAuction(&_ExpressLaneAuction.TransactOpts, firstPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) ResolveSingleBidAuction(firstPriceBid Bid) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.ResolveSingleBidAuction(&_ExpressLaneAuction.TransactOpts, firstPriceBid)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.RevokeRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.RevokeRole(&_ExpressLaneAuction.TransactOpts, role, account)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetBeneficiary(opts *bind.TransactOpts, newBeneficiary common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setBeneficiary", newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetBeneficiary(&_ExpressLaneAuction.TransactOpts, newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetBeneficiary(&_ExpressLaneAuction.TransactOpts, newBeneficiary)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetMinReservePrice(opts *bind.TransactOpts, newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setMinReservePrice", newMinReservePrice)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetMinReservePrice(newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetMinReservePrice(&_ExpressLaneAuction.TransactOpts, newMinReservePrice)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetMinReservePrice(newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetMinReservePrice(&_ExpressLaneAuction.TransactOpts, newMinReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetReservePrice(opts *bind.TransactOpts, newReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setReservePrice", newReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetReservePrice(newReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetReservePrice(&_ExpressLaneAuction.TransactOpts, newReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetReservePrice(newReservePrice *big.Int) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetReservePrice(&_ExpressLaneAuction.TransactOpts, newReservePrice)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetRoundTimingInfo(opts *bind.TransactOpts, newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setRoundTimingInfo", newRoundTimingInfo)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetRoundTimingInfo(newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetRoundTimingInfo(&_ExpressLaneAuction.TransactOpts, newRoundTimingInfo)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetRoundTimingInfo(newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetRoundTimingInfo(&_ExpressLaneAuction.TransactOpts, newRoundTimingInfo)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) SetTransferor(opts *bind.TransactOpts, transferor Transferor) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "setTransferor", transferor)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) SetTransferor(transferor Transferor) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetTransferor(&_ExpressLaneAuction.TransactOpts, transferor)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) SetTransferor(transferor Transferor) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.SetTransferor(&_ExpressLaneAuction.TransactOpts, transferor)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactor) TransferExpressLaneController(opts *bind.TransactOpts, round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.contract.Transact(opts, "transferExpressLaneController", round, newExpressLaneController)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionSession) TransferExpressLaneController(round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.TransferExpressLaneController(&_ExpressLaneAuction.TransactOpts, round, newExpressLaneController)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_ExpressLaneAuction *ExpressLaneAuctionTransactorSession) TransferExpressLaneController(round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _ExpressLaneAuction.Contract.TransferExpressLaneController(&_ExpressLaneAuction.TransactOpts, round, newExpressLaneController)
}

// ExpressLaneAuctionAuctionResolvedIterator is returned from FilterAuctionResolved and is used to iterate over the raw logs and unpacked data for AuctionResolved events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionAuctionResolvedIterator struct {
	Event *ExpressLaneAuctionAuctionResolved // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionAuctionResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionAuctionResolved)
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
		it.Event = new(ExpressLaneAuctionAuctionResolved)
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
func (it *ExpressLaneAuctionAuctionResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionAuctionResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionAuctionResolved represents a AuctionResolved event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionAuctionResolved struct {
	IsMultiBidAuction               bool
	Round                           uint64
	FirstPriceBidder                common.Address
	FirstPriceExpressLaneController common.Address
	FirstPriceAmount                *big.Int
	Price                           *big.Int
	RoundStartTimestamp             uint64
	RoundEndTimestamp               uint64
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterAuctionResolved is a free log retrieval operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterAuctionResolved(opts *bind.FilterOpts, isMultiBidAuction []bool, firstPriceBidder []common.Address, firstPriceExpressLaneController []common.Address) (*ExpressLaneAuctionAuctionResolvedIterator, error) {

	var isMultiBidAuctionRule []interface{}
	for _, isMultiBidAuctionItem := range isMultiBidAuction {
		isMultiBidAuctionRule = append(isMultiBidAuctionRule, isMultiBidAuctionItem)
	}

	var firstPriceBidderRule []interface{}
	for _, firstPriceBidderItem := range firstPriceBidder {
		firstPriceBidderRule = append(firstPriceBidderRule, firstPriceBidderItem)
	}
	var firstPriceExpressLaneControllerRule []interface{}
	for _, firstPriceExpressLaneControllerItem := range firstPriceExpressLaneController {
		firstPriceExpressLaneControllerRule = append(firstPriceExpressLaneControllerRule, firstPriceExpressLaneControllerItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "AuctionResolved", isMultiBidAuctionRule, firstPriceBidderRule, firstPriceExpressLaneControllerRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionAuctionResolvedIterator{contract: _ExpressLaneAuction.contract, event: "AuctionResolved", logs: logs, sub: sub}, nil
}

// WatchAuctionResolved is a free log subscription operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchAuctionResolved(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionAuctionResolved, isMultiBidAuction []bool, firstPriceBidder []common.Address, firstPriceExpressLaneController []common.Address) (event.Subscription, error) {

	var isMultiBidAuctionRule []interface{}
	for _, isMultiBidAuctionItem := range isMultiBidAuction {
		isMultiBidAuctionRule = append(isMultiBidAuctionRule, isMultiBidAuctionItem)
	}

	var firstPriceBidderRule []interface{}
	for _, firstPriceBidderItem := range firstPriceBidder {
		firstPriceBidderRule = append(firstPriceBidderRule, firstPriceBidderItem)
	}
	var firstPriceExpressLaneControllerRule []interface{}
	for _, firstPriceExpressLaneControllerItem := range firstPriceExpressLaneController {
		firstPriceExpressLaneControllerRule = append(firstPriceExpressLaneControllerRule, firstPriceExpressLaneControllerItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "AuctionResolved", isMultiBidAuctionRule, firstPriceBidderRule, firstPriceExpressLaneControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionAuctionResolved)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "AuctionResolved", log); err != nil {
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

// ParseAuctionResolved is a log parse operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseAuctionResolved(log types.Log) (*ExpressLaneAuctionAuctionResolved, error) {
	event := new(ExpressLaneAuctionAuctionResolved)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "AuctionResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionDepositIterator struct {
	Event *ExpressLaneAuctionDeposit // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionDeposit)
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
		it.Event = new(ExpressLaneAuctionDeposit)
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
func (it *ExpressLaneAuctionDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionDeposit represents a Deposit event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*ExpressLaneAuctionDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionDepositIterator{contract: _ExpressLaneAuction.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionDeposit)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseDeposit(log types.Log) (*ExpressLaneAuctionDeposit, error) {
	event := new(ExpressLaneAuctionDeposit)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionInitializedIterator struct {
	Event *ExpressLaneAuctionInitialized // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionInitialized)
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
		it.Event = new(ExpressLaneAuctionInitialized)
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
func (it *ExpressLaneAuctionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionInitialized represents a Initialized event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterInitialized(opts *bind.FilterOpts) (*ExpressLaneAuctionInitializedIterator, error) {

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionInitializedIterator{contract: _ExpressLaneAuction.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionInitialized) (event.Subscription, error) {

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionInitialized)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseInitialized(log types.Log) (*ExpressLaneAuctionInitialized, error) {
	event := new(ExpressLaneAuctionInitialized)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleAdminChangedIterator struct {
	Event *ExpressLaneAuctionRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionRoleAdminChanged)
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
		it.Event = new(ExpressLaneAuctionRoleAdminChanged)
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
func (it *ExpressLaneAuctionRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionRoleAdminChanged represents a RoleAdminChanged event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ExpressLaneAuctionRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionRoleAdminChangedIterator{contract: _ExpressLaneAuction.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionRoleAdminChanged)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseRoleAdminChanged(log types.Log) (*ExpressLaneAuctionRoleAdminChanged, error) {
	event := new(ExpressLaneAuctionRoleAdminChanged)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleGrantedIterator struct {
	Event *ExpressLaneAuctionRoleGranted // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionRoleGranted)
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
		it.Event = new(ExpressLaneAuctionRoleGranted)
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
func (it *ExpressLaneAuctionRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionRoleGranted represents a RoleGranted event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ExpressLaneAuctionRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionRoleGrantedIterator{contract: _ExpressLaneAuction.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionRoleGranted)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseRoleGranted(log types.Log) (*ExpressLaneAuctionRoleGranted, error) {
	event := new(ExpressLaneAuctionRoleGranted)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleRevokedIterator struct {
	Event *ExpressLaneAuctionRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionRoleRevoked)
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
		it.Event = new(ExpressLaneAuctionRoleRevoked)
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
func (it *ExpressLaneAuctionRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionRoleRevoked represents a RoleRevoked event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ExpressLaneAuctionRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionRoleRevokedIterator{contract: _ExpressLaneAuction.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionRoleRevoked)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseRoleRevoked(log types.Log) (*ExpressLaneAuctionRoleRevoked, error) {
	event := new(ExpressLaneAuctionRoleRevoked)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetBeneficiaryIterator is returned from FilterSetBeneficiary and is used to iterate over the raw logs and unpacked data for SetBeneficiary events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetBeneficiaryIterator struct {
	Event *ExpressLaneAuctionSetBeneficiary // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetBeneficiaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetBeneficiary)
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
		it.Event = new(ExpressLaneAuctionSetBeneficiary)
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
func (it *ExpressLaneAuctionSetBeneficiaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetBeneficiaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetBeneficiary represents a SetBeneficiary event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetBeneficiary struct {
	OldBeneficiary common.Address
	NewBeneficiary common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetBeneficiary is a free log retrieval operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetBeneficiary(opts *bind.FilterOpts) (*ExpressLaneAuctionSetBeneficiaryIterator, error) {

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetBeneficiary")
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetBeneficiaryIterator{contract: _ExpressLaneAuction.contract, event: "SetBeneficiary", logs: logs, sub: sub}, nil
}

// WatchSetBeneficiary is a free log subscription operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetBeneficiary(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetBeneficiary) (event.Subscription, error) {

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetBeneficiary")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetBeneficiary)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetBeneficiary", log); err != nil {
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

// ParseSetBeneficiary is a log parse operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetBeneficiary(log types.Log) (*ExpressLaneAuctionSetBeneficiary, error) {
	event := new(ExpressLaneAuctionSetBeneficiary)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetBeneficiary", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetExpressLaneControllerIterator is returned from FilterSetExpressLaneController and is used to iterate over the raw logs and unpacked data for SetExpressLaneController events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetExpressLaneControllerIterator struct {
	Event *ExpressLaneAuctionSetExpressLaneController // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetExpressLaneControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetExpressLaneController)
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
		it.Event = new(ExpressLaneAuctionSetExpressLaneController)
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
func (it *ExpressLaneAuctionSetExpressLaneControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetExpressLaneControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetExpressLaneController represents a SetExpressLaneController event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetExpressLaneController struct {
	Round                         uint64
	PreviousExpressLaneController common.Address
	NewExpressLaneController      common.Address
	Transferor                    common.Address
	StartTimestamp                uint64
	EndTimestamp                  uint64
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterSetExpressLaneController is a free log retrieval operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetExpressLaneController(opts *bind.FilterOpts, previousExpressLaneController []common.Address, newExpressLaneController []common.Address, transferor []common.Address) (*ExpressLaneAuctionSetExpressLaneControllerIterator, error) {

	var previousExpressLaneControllerRule []interface{}
	for _, previousExpressLaneControllerItem := range previousExpressLaneController {
		previousExpressLaneControllerRule = append(previousExpressLaneControllerRule, previousExpressLaneControllerItem)
	}
	var newExpressLaneControllerRule []interface{}
	for _, newExpressLaneControllerItem := range newExpressLaneController {
		newExpressLaneControllerRule = append(newExpressLaneControllerRule, newExpressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetExpressLaneController", previousExpressLaneControllerRule, newExpressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetExpressLaneControllerIterator{contract: _ExpressLaneAuction.contract, event: "SetExpressLaneController", logs: logs, sub: sub}, nil
}

// WatchSetExpressLaneController is a free log subscription operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetExpressLaneController(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetExpressLaneController, previousExpressLaneController []common.Address, newExpressLaneController []common.Address, transferor []common.Address) (event.Subscription, error) {

	var previousExpressLaneControllerRule []interface{}
	for _, previousExpressLaneControllerItem := range previousExpressLaneController {
		previousExpressLaneControllerRule = append(previousExpressLaneControllerRule, previousExpressLaneControllerItem)
	}
	var newExpressLaneControllerRule []interface{}
	for _, newExpressLaneControllerItem := range newExpressLaneController {
		newExpressLaneControllerRule = append(newExpressLaneControllerRule, newExpressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetExpressLaneController", previousExpressLaneControllerRule, newExpressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetExpressLaneController)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetExpressLaneController", log); err != nil {
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

// ParseSetExpressLaneController is a log parse operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetExpressLaneController(log types.Log) (*ExpressLaneAuctionSetExpressLaneController, error) {
	event := new(ExpressLaneAuctionSetExpressLaneController)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetExpressLaneController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetMinReservePriceIterator is returned from FilterSetMinReservePrice and is used to iterate over the raw logs and unpacked data for SetMinReservePrice events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetMinReservePriceIterator struct {
	Event *ExpressLaneAuctionSetMinReservePrice // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetMinReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetMinReservePrice)
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
		it.Event = new(ExpressLaneAuctionSetMinReservePrice)
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
func (it *ExpressLaneAuctionSetMinReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetMinReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetMinReservePrice represents a SetMinReservePrice event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetMinReservePrice struct {
	OldPrice *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMinReservePrice is a free log retrieval operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetMinReservePrice(opts *bind.FilterOpts) (*ExpressLaneAuctionSetMinReservePriceIterator, error) {

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetMinReservePrice")
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetMinReservePriceIterator{contract: _ExpressLaneAuction.contract, event: "SetMinReservePrice", logs: logs, sub: sub}, nil
}

// WatchSetMinReservePrice is a free log subscription operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetMinReservePrice(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetMinReservePrice) (event.Subscription, error) {

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetMinReservePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetMinReservePrice)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetMinReservePrice", log); err != nil {
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

// ParseSetMinReservePrice is a log parse operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetMinReservePrice(log types.Log) (*ExpressLaneAuctionSetMinReservePrice, error) {
	event := new(ExpressLaneAuctionSetMinReservePrice)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetMinReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetReservePriceIterator is returned from FilterSetReservePrice and is used to iterate over the raw logs and unpacked data for SetReservePrice events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetReservePriceIterator struct {
	Event *ExpressLaneAuctionSetReservePrice // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetReservePrice)
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
		it.Event = new(ExpressLaneAuctionSetReservePrice)
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
func (it *ExpressLaneAuctionSetReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetReservePrice represents a SetReservePrice event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetReservePrice struct {
	OldReservePrice *big.Int
	NewReservePrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetReservePrice is a free log retrieval operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetReservePrice(opts *bind.FilterOpts) (*ExpressLaneAuctionSetReservePriceIterator, error) {

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetReservePrice")
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetReservePriceIterator{contract: _ExpressLaneAuction.contract, event: "SetReservePrice", logs: logs, sub: sub}, nil
}

// WatchSetReservePrice is a free log subscription operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetReservePrice(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetReservePrice) (event.Subscription, error) {

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetReservePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetReservePrice)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetReservePrice", log); err != nil {
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

// ParseSetReservePrice is a log parse operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetReservePrice(log types.Log) (*ExpressLaneAuctionSetReservePrice, error) {
	event := new(ExpressLaneAuctionSetReservePrice)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetRoundTimingInfoIterator is returned from FilterSetRoundTimingInfo and is used to iterate over the raw logs and unpacked data for SetRoundTimingInfo events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetRoundTimingInfoIterator struct {
	Event *ExpressLaneAuctionSetRoundTimingInfo // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetRoundTimingInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetRoundTimingInfo)
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
		it.Event = new(ExpressLaneAuctionSetRoundTimingInfo)
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
func (it *ExpressLaneAuctionSetRoundTimingInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetRoundTimingInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetRoundTimingInfo represents a SetRoundTimingInfo event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetRoundTimingInfo struct {
	CurrentRound             uint64
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetRoundTimingInfo is a free log retrieval operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetRoundTimingInfo(opts *bind.FilterOpts) (*ExpressLaneAuctionSetRoundTimingInfoIterator, error) {

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetRoundTimingInfo")
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetRoundTimingInfoIterator{contract: _ExpressLaneAuction.contract, event: "SetRoundTimingInfo", logs: logs, sub: sub}, nil
}

// WatchSetRoundTimingInfo is a free log subscription operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetRoundTimingInfo(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetRoundTimingInfo) (event.Subscription, error) {

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetRoundTimingInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetRoundTimingInfo)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetRoundTimingInfo", log); err != nil {
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

// ParseSetRoundTimingInfo is a log parse operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetRoundTimingInfo(log types.Log) (*ExpressLaneAuctionSetRoundTimingInfo, error) {
	event := new(ExpressLaneAuctionSetRoundTimingInfo)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetRoundTimingInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionSetTransferorIterator is returned from FilterSetTransferor and is used to iterate over the raw logs and unpacked data for SetTransferor events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetTransferorIterator struct {
	Event *ExpressLaneAuctionSetTransferor // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionSetTransferorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionSetTransferor)
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
		it.Event = new(ExpressLaneAuctionSetTransferor)
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
func (it *ExpressLaneAuctionSetTransferorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionSetTransferorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionSetTransferor represents a SetTransferor event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionSetTransferor struct {
	ExpressLaneController common.Address
	Transferor            common.Address
	FixedUntilRound       uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetTransferor is a free log retrieval operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterSetTransferor(opts *bind.FilterOpts, expressLaneController []common.Address, transferor []common.Address) (*ExpressLaneAuctionSetTransferorIterator, error) {

	var expressLaneControllerRule []interface{}
	for _, expressLaneControllerItem := range expressLaneController {
		expressLaneControllerRule = append(expressLaneControllerRule, expressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "SetTransferor", expressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionSetTransferorIterator{contract: _ExpressLaneAuction.contract, event: "SetTransferor", logs: logs, sub: sub}, nil
}

// WatchSetTransferor is a free log subscription operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchSetTransferor(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionSetTransferor, expressLaneController []common.Address, transferor []common.Address) (event.Subscription, error) {

	var expressLaneControllerRule []interface{}
	for _, expressLaneControllerItem := range expressLaneController {
		expressLaneControllerRule = append(expressLaneControllerRule, expressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "SetTransferor", expressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionSetTransferor)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetTransferor", log); err != nil {
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

// ParseSetTransferor is a log parse operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseSetTransferor(log types.Log) (*ExpressLaneAuctionSetTransferor, error) {
	event := new(ExpressLaneAuctionSetTransferor)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "SetTransferor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionWithdrawalFinalizedIterator struct {
	Event *ExpressLaneAuctionWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionWithdrawalFinalized)
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
		it.Event = new(ExpressLaneAuctionWithdrawalFinalized)
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
func (it *ExpressLaneAuctionWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionWithdrawalFinalized represents a WithdrawalFinalized event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionWithdrawalFinalized struct {
	Account          common.Address
	WithdrawalAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, account []common.Address) (*ExpressLaneAuctionWithdrawalFinalizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "WithdrawalFinalized", accountRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionWithdrawalFinalizedIterator{contract: _ExpressLaneAuction.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionWithdrawalFinalized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "WithdrawalFinalized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionWithdrawalFinalized)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseWithdrawalFinalized(log types.Log) (*ExpressLaneAuctionWithdrawalFinalized, error) {
	event := new(ExpressLaneAuctionWithdrawalFinalized)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExpressLaneAuctionWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionWithdrawalInitiatedIterator struct {
	Event *ExpressLaneAuctionWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *ExpressLaneAuctionWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpressLaneAuctionWithdrawalInitiated)
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
		it.Event = new(ExpressLaneAuctionWithdrawalInitiated)
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
func (it *ExpressLaneAuctionWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpressLaneAuctionWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpressLaneAuctionWithdrawalInitiated represents a WithdrawalInitiated event raised by the ExpressLaneAuction contract.
type ExpressLaneAuctionWithdrawalInitiated struct {
	Account           common.Address
	WithdrawalAmount  *big.Int
	RoundWithdrawable *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, account []common.Address) (*ExpressLaneAuctionWithdrawalInitiatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.FilterLogs(opts, "WithdrawalInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return &ExpressLaneAuctionWithdrawalInitiatedIterator{contract: _ExpressLaneAuction.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *ExpressLaneAuctionWithdrawalInitiated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExpressLaneAuction.contract.WatchLogs(opts, "WithdrawalInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpressLaneAuctionWithdrawalInitiated)
				if err := _ExpressLaneAuction.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_ExpressLaneAuction *ExpressLaneAuctionFilterer) ParseWithdrawalInitiated(log types.Log) (*ExpressLaneAuctionWithdrawalInitiated, error) {
	event := new(ExpressLaneAuctionWithdrawalInitiated)
	if err := _ExpressLaneAuction.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionMetaData contains all meta data concerning the IExpressLaneAuction contract.
var IExpressLaneAuctionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isMultiBidAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"firstPriceBidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"firstPriceExpressLaneController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstPriceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundStartTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundEndTimestamp\",\"type\":\"uint64\"}],\"name\":\"AuctionResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBeneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"SetBeneficiary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousExpressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newExpressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"}],\"name\":\"SetExpressLaneController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"SetMinReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReservePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReservePrice\",\"type\":\"uint256\"}],\"name\":\"SetReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"currentRound\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"name\":\"SetRoundTimingInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"name\":\"SetTransferor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundWithdrawable\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUCTIONEER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTIONEER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BENEFICIARY_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_RESERVE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_SETTER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESERVE_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUND_TIMING_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"balanceOfAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiaryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"biddingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flushBeneficiaryBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getBidHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_auctioneer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_biddingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structRoundTimingInfo\",\"name\":\"_roundTimingInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_minReservePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_auctioneerAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minReservePriceSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePriceSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reservePriceSetterAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beneficiarySetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_roundTimingSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_masterAdmin\",\"type\":\"address\"}],\"internalType\":\"structInitArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiateWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAuctionRoundClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReserveBlackout\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minReservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"firstPriceBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"secondPriceBid\",\"type\":\"tuple\"}],\"name\":\"resolveMultiBidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBid\",\"name\":\"firstPriceBid\",\"type\":\"tuple\"}],\"name\":\"resolveSingleBidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolvedRounds\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"internalType\":\"structELCRound\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"internalType\":\"structELCRound\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"roundTimestamps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundTimingInfo\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinReservePrice\",\"type\":\"uint256\"}],\"name\":\"setMinReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReservePrice\",\"type\":\"uint256\"}],\"name\":\"setReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"offsetTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"roundDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"auctionClosingSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"reserveSubmissionSeconds\",\"type\":\"uint64\"}],\"internalType\":\"structRoundTimingInfo\",\"name\":\"newRoundTimingInfo\",\"type\":\"tuple\"}],\"name\":\"setRoundTimingInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fixedUntilRound\",\"type\":\"uint64\"}],\"internalType\":\"structTransferor\",\"name\":\"transferor\",\"type\":\"tuple\"}],\"name\":\"setTransferor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newExpressLaneController\",\"type\":\"address\"}],\"name\":\"transferExpressLaneController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expressLaneController\",\"type\":\"address\"}],\"name\":\"transferorOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"fixedUntil\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"round\",\"type\":\"uint64\"}],\"name\":\"withdrawableBalanceAtRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IExpressLaneAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use IExpressLaneAuctionMetaData.ABI instead.
var IExpressLaneAuctionABI = IExpressLaneAuctionMetaData.ABI

// IExpressLaneAuction is an auto generated Go binding around an Ethereum contract.
type IExpressLaneAuction struct {
	IExpressLaneAuctionCaller     // Read-only binding to the contract
	IExpressLaneAuctionTransactor // Write-only binding to the contract
	IExpressLaneAuctionFilterer   // Log filterer for contract events
}

// IExpressLaneAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExpressLaneAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExpressLaneAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExpressLaneAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExpressLaneAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExpressLaneAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExpressLaneAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExpressLaneAuctionSession struct {
	Contract     *IExpressLaneAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IExpressLaneAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExpressLaneAuctionCallerSession struct {
	Contract *IExpressLaneAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IExpressLaneAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExpressLaneAuctionTransactorSession struct {
	Contract     *IExpressLaneAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IExpressLaneAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExpressLaneAuctionRaw struct {
	Contract *IExpressLaneAuction // Generic contract binding to access the raw methods on
}

// IExpressLaneAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExpressLaneAuctionCallerRaw struct {
	Contract *IExpressLaneAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// IExpressLaneAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExpressLaneAuctionTransactorRaw struct {
	Contract *IExpressLaneAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExpressLaneAuction creates a new instance of IExpressLaneAuction, bound to a specific deployed contract.
func NewIExpressLaneAuction(address common.Address, backend bind.ContractBackend) (*IExpressLaneAuction, error) {
	contract, err := bindIExpressLaneAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuction{IExpressLaneAuctionCaller: IExpressLaneAuctionCaller{contract: contract}, IExpressLaneAuctionTransactor: IExpressLaneAuctionTransactor{contract: contract}, IExpressLaneAuctionFilterer: IExpressLaneAuctionFilterer{contract: contract}}, nil
}

// NewIExpressLaneAuctionCaller creates a new read-only instance of IExpressLaneAuction, bound to a specific deployed contract.
func NewIExpressLaneAuctionCaller(address common.Address, caller bind.ContractCaller) (*IExpressLaneAuctionCaller, error) {
	contract, err := bindIExpressLaneAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionCaller{contract: contract}, nil
}

// NewIExpressLaneAuctionTransactor creates a new write-only instance of IExpressLaneAuction, bound to a specific deployed contract.
func NewIExpressLaneAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*IExpressLaneAuctionTransactor, error) {
	contract, err := bindIExpressLaneAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionTransactor{contract: contract}, nil
}

// NewIExpressLaneAuctionFilterer creates a new log filterer instance of IExpressLaneAuction, bound to a specific deployed contract.
func NewIExpressLaneAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*IExpressLaneAuctionFilterer, error) {
	contract, err := bindIExpressLaneAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionFilterer{contract: contract}, nil
}

// bindIExpressLaneAuction binds a generic wrapper to an already deployed contract.
func bindIExpressLaneAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IExpressLaneAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExpressLaneAuction *IExpressLaneAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExpressLaneAuction.Contract.IExpressLaneAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExpressLaneAuction *IExpressLaneAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.IExpressLaneAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExpressLaneAuction *IExpressLaneAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.IExpressLaneAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExpressLaneAuction *IExpressLaneAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExpressLaneAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.BalanceOf(&_IExpressLaneAuction.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.BalanceOf(&_IExpressLaneAuction.CallOpts, account)
}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) BalanceOfAtRound(opts *bind.CallOpts, account common.Address, round uint64) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "balanceOfAtRound", account, round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BalanceOfAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.BalanceOfAtRound(&_IExpressLaneAuction.CallOpts, account, round)
}

// BalanceOfAtRound is a free data retrieval call binding the contract method 0x5633c337.
//
// Solidity: function balanceOfAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) BalanceOfAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.BalanceOfAtRound(&_IExpressLaneAuction.CallOpts, account, round)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) CurrentRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "currentRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) CurrentRound() (uint64, error) {
	return _IExpressLaneAuction.Contract.CurrentRound(&_IExpressLaneAuction.CallOpts)
}

// CurrentRound is a free data retrieval call binding the contract method 0x8a19c8bc.
//
// Solidity: function currentRound() view returns(uint64)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) CurrentRound() (uint64, error) {
	return _IExpressLaneAuction.Contract.CurrentRound(&_IExpressLaneAuction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) DomainSeparator() ([32]byte, error) {
	return _IExpressLaneAuction.Contract.DomainSeparator(&_IExpressLaneAuction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) DomainSeparator() ([32]byte, error) {
	return _IExpressLaneAuction.Contract.DomainSeparator(&_IExpressLaneAuction.CallOpts)
}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) GetBidHash(opts *bind.CallOpts, round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "getBidHash", round, expressLaneController, amount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GetBidHash(round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	return _IExpressLaneAuction.Contract.GetBidHash(&_IExpressLaneAuction.CallOpts, round, expressLaneController, amount)
}

// GetBidHash is a free data retrieval call binding the contract method 0x04c584ad.
//
// Solidity: function getBidHash(uint64 round, address expressLaneController, uint256 amount) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) GetBidHash(round uint64, expressLaneController common.Address, amount *big.Int) ([32]byte, error) {
	return _IExpressLaneAuction.Contract.GetBidHash(&_IExpressLaneAuction.CallOpts, round, expressLaneController, amount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IExpressLaneAuction.Contract.GetRoleAdmin(&_IExpressLaneAuction.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IExpressLaneAuction.Contract.GetRoleAdmin(&_IExpressLaneAuction.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IExpressLaneAuction.Contract.GetRoleMember(&_IExpressLaneAuction.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IExpressLaneAuction.Contract.GetRoleMember(&_IExpressLaneAuction.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.GetRoleMemberCount(&_IExpressLaneAuction.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.GetRoleMemberCount(&_IExpressLaneAuction.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IExpressLaneAuction.Contract.HasRole(&_IExpressLaneAuction.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IExpressLaneAuction.Contract.HasRole(&_IExpressLaneAuction.CallOpts, role, account)
}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) IsAuctionRoundClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "isAuctionRoundClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) IsAuctionRoundClosed() (bool, error) {
	return _IExpressLaneAuction.Contract.IsAuctionRoundClosed(&_IExpressLaneAuction.CallOpts)
}

// IsAuctionRoundClosed is a free data retrieval call binding the contract method 0x2d668ce7.
//
// Solidity: function isAuctionRoundClosed() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) IsAuctionRoundClosed() (bool, error) {
	return _IExpressLaneAuction.Contract.IsAuctionRoundClosed(&_IExpressLaneAuction.CallOpts)
}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) IsReserveBlackout(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "isReserveBlackout")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) IsReserveBlackout() (bool, error) {
	return _IExpressLaneAuction.Contract.IsReserveBlackout(&_IExpressLaneAuction.CallOpts)
}

// IsReserveBlackout is a free data retrieval call binding the contract method 0xe460d2c5.
//
// Solidity: function isReserveBlackout() view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) IsReserveBlackout() (bool, error) {
	return _IExpressLaneAuction.Contract.IsReserveBlackout(&_IExpressLaneAuction.CallOpts)
}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) ResolvedRounds(opts *bind.CallOpts) (ELCRound, ELCRound, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "resolvedRounds")

	if err != nil {
		return *new(ELCRound), *new(ELCRound), err
	}

	out0 := *abi.ConvertType(out[0], new(ELCRound)).(*ELCRound)
	out1 := *abi.ConvertType(out[1], new(ELCRound)).(*ELCRound)

	return out0, out1, err

}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ResolvedRounds() (ELCRound, ELCRound, error) {
	return _IExpressLaneAuction.Contract.ResolvedRounds(&_IExpressLaneAuction.CallOpts)
}

// ResolvedRounds is a free data retrieval call binding the contract method 0x0d253fbe.
//
// Solidity: function resolvedRounds() view returns((address,uint64), (address,uint64))
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) ResolvedRounds() (ELCRound, ELCRound, error) {
	return _IExpressLaneAuction.Contract.ResolvedRounds(&_IExpressLaneAuction.CallOpts)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64 start, uint64 end)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) RoundTimestamps(opts *bind.CallOpts, round uint64) (struct {
	Start uint64
	End   uint64
}, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "roundTimestamps", round)

	outstruct := new(struct {
		Start uint64
		End   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Start = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.End = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64 start, uint64 end)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RoundTimestamps(round uint64) (struct {
	Start uint64
	End   uint64
}, error) {
	return _IExpressLaneAuction.Contract.RoundTimestamps(&_IExpressLaneAuction.CallOpts, round)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0x7b617f94.
//
// Solidity: function roundTimestamps(uint64 round) view returns(uint64 start, uint64 end)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) RoundTimestamps(round uint64) (struct {
	Start uint64
	End   uint64
}, error) {
	return _IExpressLaneAuction.Contract.RoundTimestamps(&_IExpressLaneAuction.CallOpts, round)
}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) RoundTimingInfo(opts *bind.CallOpts) (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "roundTimingInfo")

	outstruct := new(struct {
		OffsetTimestamp          int64
		RoundDurationSeconds     uint64
		AuctionClosingSeconds    uint64
		ReserveSubmissionSeconds uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OffsetTimestamp = *abi.ConvertType(out[0], new(int64)).(*int64)
	outstruct.RoundDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.AuctionClosingSeconds = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.ReserveSubmissionSeconds = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RoundTimingInfo() (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	return _IExpressLaneAuction.Contract.RoundTimingInfo(&_IExpressLaneAuction.CallOpts)
}

// RoundTimingInfo is a free data retrieval call binding the contract method 0x0152682d.
//
// Solidity: function roundTimingInfo() view returns(int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) RoundTimingInfo() (struct {
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
}, error) {
	return _IExpressLaneAuction.Contract.RoundTimingInfo(&_IExpressLaneAuction.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IExpressLaneAuction.Contract.SupportsInterface(&_IExpressLaneAuction.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IExpressLaneAuction.Contract.SupportsInterface(&_IExpressLaneAuction.CallOpts, interfaceId)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) WithdrawableBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "withdrawableBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) WithdrawableBalance(account common.Address) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.WithdrawableBalance(&_IExpressLaneAuction.CallOpts, account)
}

// WithdrawableBalance is a free data retrieval call binding the contract method 0x02b62938.
//
// Solidity: function withdrawableBalance(address account) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) WithdrawableBalance(account common.Address) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.WithdrawableBalance(&_IExpressLaneAuction.CallOpts, account)
}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCaller) WithdrawableBalanceAtRound(opts *bind.CallOpts, account common.Address, round uint64) (*big.Int, error) {
	var out []interface{}
	err := _IExpressLaneAuction.contract.Call(opts, &out, "withdrawableBalanceAtRound", account, round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) WithdrawableBalanceAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.WithdrawableBalanceAtRound(&_IExpressLaneAuction.CallOpts, account, round)
}

// WithdrawableBalanceAtRound is a free data retrieval call binding the contract method 0x6e8cace5.
//
// Solidity: function withdrawableBalanceAtRound(address account, uint64 round) view returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionCallerSession) WithdrawableBalanceAtRound(account common.Address, round uint64) (*big.Int, error) {
	return _IExpressLaneAuction.Contract.WithdrawableBalanceAtRound(&_IExpressLaneAuction.CallOpts, account, round)
}

// AUCTIONEERADMINROLE is a paid mutator transaction binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) AUCTIONEERADMINROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "AUCTIONEER_ADMIN_ROLE")
}

// AUCTIONEERADMINROLE is a paid mutator transaction binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) AUCTIONEERADMINROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.AUCTIONEERADMINROLE(&_IExpressLaneAuction.TransactOpts)
}

// AUCTIONEERADMINROLE is a paid mutator transaction binding the contract method 0x14d96316.
//
// Solidity: function AUCTIONEER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) AUCTIONEERADMINROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.AUCTIONEERADMINROLE(&_IExpressLaneAuction.TransactOpts)
}

// AUCTIONEERROLE is a paid mutator transaction binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) AUCTIONEERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "AUCTIONEER_ROLE")
}

// AUCTIONEERROLE is a paid mutator transaction binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) AUCTIONEERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.AUCTIONEERROLE(&_IExpressLaneAuction.TransactOpts)
}

// AUCTIONEERROLE is a paid mutator transaction binding the contract method 0xcfe9232b.
//
// Solidity: function AUCTIONEER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) AUCTIONEERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.AUCTIONEERROLE(&_IExpressLaneAuction.TransactOpts)
}

// BENEFICIARYSETTERROLE is a paid mutator transaction binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) BENEFICIARYSETTERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "BENEFICIARY_SETTER_ROLE")
}

// BENEFICIARYSETTERROLE is a paid mutator transaction binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BENEFICIARYSETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BENEFICIARYSETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// BENEFICIARYSETTERROLE is a paid mutator transaction binding the contract method 0x336a5b5e.
//
// Solidity: function BENEFICIARY_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) BENEFICIARYSETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BENEFICIARYSETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// MINRESERVESETTERROLE is a paid mutator transaction binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) MINRESERVESETTERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "MIN_RESERVE_SETTER_ROLE")
}

// MINRESERVESETTERROLE is a paid mutator transaction binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) MINRESERVESETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.MINRESERVESETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// MINRESERVESETTERROLE is a paid mutator transaction binding the contract method 0x8948cc4e.
//
// Solidity: function MIN_RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) MINRESERVESETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.MINRESERVESETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// RESERVESETTERADMINROLE is a paid mutator transaction binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) RESERVESETTERADMINROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "RESERVE_SETTER_ADMIN_ROLE")
}

// RESERVESETTERADMINROLE is a paid mutator transaction binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RESERVESETTERADMINROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RESERVESETTERADMINROLE(&_IExpressLaneAuction.TransactOpts)
}

// RESERVESETTERADMINROLE is a paid mutator transaction binding the contract method 0xe3f7bb55.
//
// Solidity: function RESERVE_SETTER_ADMIN_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) RESERVESETTERADMINROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RESERVESETTERADMINROLE(&_IExpressLaneAuction.TransactOpts)
}

// RESERVESETTERROLE is a paid mutator transaction binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) RESERVESETTERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "RESERVE_SETTER_ROLE")
}

// RESERVESETTERROLE is a paid mutator transaction binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RESERVESETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RESERVESETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// RESERVESETTERROLE is a paid mutator transaction binding the contract method 0xb3ee252f.
//
// Solidity: function RESERVE_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) RESERVESETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RESERVESETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// ROUNDTIMINGSETTERROLE is a paid mutator transaction binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) ROUNDTIMINGSETTERROLE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "ROUND_TIMING_SETTER_ROLE")
}

// ROUNDTIMINGSETTERROLE is a paid mutator transaction binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ROUNDTIMINGSETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ROUNDTIMINGSETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// ROUNDTIMINGSETTERROLE is a paid mutator transaction binding the contract method 0x1682e50b.
//
// Solidity: function ROUND_TIMING_SETTER_ROLE() returns(bytes32)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) ROUNDTIMINGSETTERROLE() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ROUNDTIMINGSETTERROLE(&_IExpressLaneAuction.TransactOpts)
}

// Beneficiary is a paid mutator transaction binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) Beneficiary(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "beneficiary")
}

// Beneficiary is a paid mutator transaction binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) Beneficiary() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Beneficiary(&_IExpressLaneAuction.TransactOpts)
}

// Beneficiary is a paid mutator transaction binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) Beneficiary() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Beneficiary(&_IExpressLaneAuction.TransactOpts)
}

// BeneficiaryBalance is a paid mutator transaction binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) BeneficiaryBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "beneficiaryBalance")
}

// BeneficiaryBalance is a paid mutator transaction binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BeneficiaryBalance() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BeneficiaryBalance(&_IExpressLaneAuction.TransactOpts)
}

// BeneficiaryBalance is a paid mutator transaction binding the contract method 0xe2fc6f68.
//
// Solidity: function beneficiaryBalance() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) BeneficiaryBalance() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BeneficiaryBalance(&_IExpressLaneAuction.TransactOpts)
}

// BiddingToken is a paid mutator transaction binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) BiddingToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "biddingToken")
}

// BiddingToken is a paid mutator transaction binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) BiddingToken() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BiddingToken(&_IExpressLaneAuction.TransactOpts)
}

// BiddingToken is a paid mutator transaction binding the contract method 0x639d7566.
//
// Solidity: function biddingToken() returns(address)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) BiddingToken() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.BiddingToken(&_IExpressLaneAuction.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Deposit(&_IExpressLaneAuction.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Deposit(&_IExpressLaneAuction.TransactOpts, amount)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) FinalizeWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "finalizeWithdrawal")
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) FinalizeWithdrawal() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.FinalizeWithdrawal(&_IExpressLaneAuction.TransactOpts)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc5b6aa2f.
//
// Solidity: function finalizeWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) FinalizeWithdrawal() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.FinalizeWithdrawal(&_IExpressLaneAuction.TransactOpts)
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) FlushBeneficiaryBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "flushBeneficiaryBalance")
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) FlushBeneficiaryBalance() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.FlushBeneficiaryBalance(&_IExpressLaneAuction.TransactOpts)
}

// FlushBeneficiaryBalance is a paid mutator transaction binding the contract method 0x6ad72517.
//
// Solidity: function flushBeneficiaryBalance() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) FlushBeneficiaryBalance() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.FlushBeneficiaryBalance(&_IExpressLaneAuction.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.GrantRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.GrantRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) Initialize(opts *bind.TransactOpts, args InitArgs) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "initialize", args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) Initialize(args InitArgs) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Initialize(&_IExpressLaneAuction.TransactOpts, args)
}

// Initialize is a paid mutator transaction binding the contract method 0x9a1fadd3.
//
// Solidity: function initialize((address,address,address,(int64,uint64,uint64,uint64),uint256,address,address,address,address,address,address,address) args) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) Initialize(args InitArgs) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.Initialize(&_IExpressLaneAuction.TransactOpts, args)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) InitiateWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "initiateWithdrawal")
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.InitiateWithdrawal(&_IExpressLaneAuction.TransactOpts)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xb51d1d4f.
//
// Solidity: function initiateWithdrawal() returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) InitiateWithdrawal() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.InitiateWithdrawal(&_IExpressLaneAuction.TransactOpts)
}

// MinReservePrice is a paid mutator transaction binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) MinReservePrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "minReservePrice")
}

// MinReservePrice is a paid mutator transaction binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) MinReservePrice() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.MinReservePrice(&_IExpressLaneAuction.TransactOpts)
}

// MinReservePrice is a paid mutator transaction binding the contract method 0x83af0a1f.
//
// Solidity: function minReservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) MinReservePrice() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.MinReservePrice(&_IExpressLaneAuction.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RenounceRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RenounceRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// ReservePrice is a paid mutator transaction binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) ReservePrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "reservePrice")
}

// ReservePrice is a paid mutator transaction binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ReservePrice() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ReservePrice(&_IExpressLaneAuction.TransactOpts)
}

// ReservePrice is a paid mutator transaction binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() returns(uint256)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) ReservePrice() (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ReservePrice(&_IExpressLaneAuction.TransactOpts)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) ResolveMultiBidAuction(opts *bind.TransactOpts, firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "resolveMultiBidAuction", firstPriceBid, secondPriceBid)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ResolveMultiBidAuction(firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ResolveMultiBidAuction(&_IExpressLaneAuction.TransactOpts, firstPriceBid, secondPriceBid)
}

// ResolveMultiBidAuction is a paid mutator transaction binding the contract method 0x447a709e.
//
// Solidity: function resolveMultiBidAuction((address,uint256,bytes) firstPriceBid, (address,uint256,bytes) secondPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) ResolveMultiBidAuction(firstPriceBid Bid, secondPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ResolveMultiBidAuction(&_IExpressLaneAuction.TransactOpts, firstPriceBid, secondPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) ResolveSingleBidAuction(opts *bind.TransactOpts, firstPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "resolveSingleBidAuction", firstPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) ResolveSingleBidAuction(firstPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ResolveSingleBidAuction(&_IExpressLaneAuction.TransactOpts, firstPriceBid)
}

// ResolveSingleBidAuction is a paid mutator transaction binding the contract method 0x6dc4fc4e.
//
// Solidity: function resolveSingleBidAuction((address,uint256,bytes) firstPriceBid) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) ResolveSingleBidAuction(firstPriceBid Bid) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.ResolveSingleBidAuction(&_IExpressLaneAuction.TransactOpts, firstPriceBid)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RevokeRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.RevokeRole(&_IExpressLaneAuction.TransactOpts, role, account)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetBeneficiary(opts *bind.TransactOpts, newBeneficiary common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setBeneficiary", newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetBeneficiary(&_IExpressLaneAuction.TransactOpts, newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetBeneficiary(&_IExpressLaneAuction.TransactOpts, newBeneficiary)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetMinReservePrice(opts *bind.TransactOpts, newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setMinReservePrice", newMinReservePrice)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetMinReservePrice(newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetMinReservePrice(&_IExpressLaneAuction.TransactOpts, newMinReservePrice)
}

// SetMinReservePrice is a paid mutator transaction binding the contract method 0xe4d20c1d.
//
// Solidity: function setMinReservePrice(uint256 newMinReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetMinReservePrice(newMinReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetMinReservePrice(&_IExpressLaneAuction.TransactOpts, newMinReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetReservePrice(opts *bind.TransactOpts, newReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setReservePrice", newReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetReservePrice(newReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetReservePrice(&_IExpressLaneAuction.TransactOpts, newReservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 newReservePrice) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetReservePrice(newReservePrice *big.Int) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetReservePrice(&_IExpressLaneAuction.TransactOpts, newReservePrice)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetRoundTimingInfo(opts *bind.TransactOpts, newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setRoundTimingInfo", newRoundTimingInfo)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetRoundTimingInfo(newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetRoundTimingInfo(&_IExpressLaneAuction.TransactOpts, newRoundTimingInfo)
}

// SetRoundTimingInfo is a paid mutator transaction binding the contract method 0xfed87be8.
//
// Solidity: function setRoundTimingInfo((int64,uint64,uint64,uint64) newRoundTimingInfo) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetRoundTimingInfo(newRoundTimingInfo RoundTimingInfo) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetRoundTimingInfo(&_IExpressLaneAuction.TransactOpts, newRoundTimingInfo)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) SetTransferor(opts *bind.TransactOpts, transferor Transferor) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "setTransferor", transferor)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) SetTransferor(transferor Transferor) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetTransferor(&_IExpressLaneAuction.TransactOpts, transferor)
}

// SetTransferor is a paid mutator transaction binding the contract method 0xbef0ec74.
//
// Solidity: function setTransferor((address,uint64) transferor) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) SetTransferor(transferor Transferor) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.SetTransferor(&_IExpressLaneAuction.TransactOpts, transferor)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) TransferExpressLaneController(opts *bind.TransactOpts, round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "transferExpressLaneController", round, newExpressLaneController)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionSession) TransferExpressLaneController(round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.TransferExpressLaneController(&_IExpressLaneAuction.TransactOpts, round, newExpressLaneController)
}

// TransferExpressLaneController is a paid mutator transaction binding the contract method 0x007be2fe.
//
// Solidity: function transferExpressLaneController(uint64 round, address newExpressLaneController) returns()
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) TransferExpressLaneController(round uint64, newExpressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.TransferExpressLaneController(&_IExpressLaneAuction.TransactOpts, round, newExpressLaneController)
}

// TransferorOf is a paid mutator transaction binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address expressLaneController) returns(address addr, uint64 fixedUntil)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactor) TransferorOf(opts *bind.TransactOpts, expressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.contract.Transact(opts, "transferorOf", expressLaneController)
}

// TransferorOf is a paid mutator transaction binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address expressLaneController) returns(address addr, uint64 fixedUntil)
func (_IExpressLaneAuction *IExpressLaneAuctionSession) TransferorOf(expressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.TransferorOf(&_IExpressLaneAuction.TransactOpts, expressLaneController)
}

// TransferorOf is a paid mutator transaction binding the contract method 0x6a514beb.
//
// Solidity: function transferorOf(address expressLaneController) returns(address addr, uint64 fixedUntil)
func (_IExpressLaneAuction *IExpressLaneAuctionTransactorSession) TransferorOf(expressLaneController common.Address) (*types.Transaction, error) {
	return _IExpressLaneAuction.Contract.TransferorOf(&_IExpressLaneAuction.TransactOpts, expressLaneController)
}

// IExpressLaneAuctionAuctionResolvedIterator is returned from FilterAuctionResolved and is used to iterate over the raw logs and unpacked data for AuctionResolved events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionAuctionResolvedIterator struct {
	Event *IExpressLaneAuctionAuctionResolved // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionAuctionResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionAuctionResolved)
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
		it.Event = new(IExpressLaneAuctionAuctionResolved)
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
func (it *IExpressLaneAuctionAuctionResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionAuctionResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionAuctionResolved represents a AuctionResolved event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionAuctionResolved struct {
	IsMultiBidAuction               bool
	Round                           uint64
	FirstPriceBidder                common.Address
	FirstPriceExpressLaneController common.Address
	FirstPriceAmount                *big.Int
	Price                           *big.Int
	RoundStartTimestamp             uint64
	RoundEndTimestamp               uint64
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterAuctionResolved is a free log retrieval operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterAuctionResolved(opts *bind.FilterOpts, isMultiBidAuction []bool, firstPriceBidder []common.Address, firstPriceExpressLaneController []common.Address) (*IExpressLaneAuctionAuctionResolvedIterator, error) {

	var isMultiBidAuctionRule []interface{}
	for _, isMultiBidAuctionItem := range isMultiBidAuction {
		isMultiBidAuctionRule = append(isMultiBidAuctionRule, isMultiBidAuctionItem)
	}

	var firstPriceBidderRule []interface{}
	for _, firstPriceBidderItem := range firstPriceBidder {
		firstPriceBidderRule = append(firstPriceBidderRule, firstPriceBidderItem)
	}
	var firstPriceExpressLaneControllerRule []interface{}
	for _, firstPriceExpressLaneControllerItem := range firstPriceExpressLaneController {
		firstPriceExpressLaneControllerRule = append(firstPriceExpressLaneControllerRule, firstPriceExpressLaneControllerItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "AuctionResolved", isMultiBidAuctionRule, firstPriceBidderRule, firstPriceExpressLaneControllerRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionAuctionResolvedIterator{contract: _IExpressLaneAuction.contract, event: "AuctionResolved", logs: logs, sub: sub}, nil
}

// WatchAuctionResolved is a free log subscription operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchAuctionResolved(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionAuctionResolved, isMultiBidAuction []bool, firstPriceBidder []common.Address, firstPriceExpressLaneController []common.Address) (event.Subscription, error) {

	var isMultiBidAuctionRule []interface{}
	for _, isMultiBidAuctionItem := range isMultiBidAuction {
		isMultiBidAuctionRule = append(isMultiBidAuctionRule, isMultiBidAuctionItem)
	}

	var firstPriceBidderRule []interface{}
	for _, firstPriceBidderItem := range firstPriceBidder {
		firstPriceBidderRule = append(firstPriceBidderRule, firstPriceBidderItem)
	}
	var firstPriceExpressLaneControllerRule []interface{}
	for _, firstPriceExpressLaneControllerItem := range firstPriceExpressLaneController {
		firstPriceExpressLaneControllerRule = append(firstPriceExpressLaneControllerRule, firstPriceExpressLaneControllerItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "AuctionResolved", isMultiBidAuctionRule, firstPriceBidderRule, firstPriceExpressLaneControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionAuctionResolved)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "AuctionResolved", log); err != nil {
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

// ParseAuctionResolved is a log parse operation binding the contract event 0x7f5bdabbd27a8fc572781b177055488d7c6729a2bade4f57da9d200f31c15d47.
//
// Solidity: event AuctionResolved(bool indexed isMultiBidAuction, uint64 round, address indexed firstPriceBidder, address indexed firstPriceExpressLaneController, uint256 firstPriceAmount, uint256 price, uint64 roundStartTimestamp, uint64 roundEndTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseAuctionResolved(log types.Log) (*IExpressLaneAuctionAuctionResolved, error) {
	event := new(IExpressLaneAuctionAuctionResolved)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "AuctionResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionDepositIterator struct {
	Event *IExpressLaneAuctionDeposit // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionDeposit)
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
		it.Event = new(IExpressLaneAuctionDeposit)
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
func (it *IExpressLaneAuctionDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionDeposit represents a Deposit event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*IExpressLaneAuctionDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionDepositIterator{contract: _IExpressLaneAuction.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionDeposit)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseDeposit(log types.Log) (*IExpressLaneAuctionDeposit, error) {
	event := new(IExpressLaneAuctionDeposit)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleAdminChangedIterator struct {
	Event *IExpressLaneAuctionRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionRoleAdminChanged)
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
		it.Event = new(IExpressLaneAuctionRoleAdminChanged)
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
func (it *IExpressLaneAuctionRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionRoleAdminChanged represents a RoleAdminChanged event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IExpressLaneAuctionRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionRoleAdminChangedIterator{contract: _IExpressLaneAuction.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionRoleAdminChanged)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseRoleAdminChanged(log types.Log) (*IExpressLaneAuctionRoleAdminChanged, error) {
	event := new(IExpressLaneAuctionRoleAdminChanged)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleGrantedIterator struct {
	Event *IExpressLaneAuctionRoleGranted // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionRoleGranted)
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
		it.Event = new(IExpressLaneAuctionRoleGranted)
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
func (it *IExpressLaneAuctionRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionRoleGranted represents a RoleGranted event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IExpressLaneAuctionRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionRoleGrantedIterator{contract: _IExpressLaneAuction.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionRoleGranted)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseRoleGranted(log types.Log) (*IExpressLaneAuctionRoleGranted, error) {
	event := new(IExpressLaneAuctionRoleGranted)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleRevokedIterator struct {
	Event *IExpressLaneAuctionRoleRevoked // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionRoleRevoked)
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
		it.Event = new(IExpressLaneAuctionRoleRevoked)
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
func (it *IExpressLaneAuctionRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionRoleRevoked represents a RoleRevoked event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IExpressLaneAuctionRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionRoleRevokedIterator{contract: _IExpressLaneAuction.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionRoleRevoked)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseRoleRevoked(log types.Log) (*IExpressLaneAuctionRoleRevoked, error) {
	event := new(IExpressLaneAuctionRoleRevoked)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetBeneficiaryIterator is returned from FilterSetBeneficiary and is used to iterate over the raw logs and unpacked data for SetBeneficiary events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetBeneficiaryIterator struct {
	Event *IExpressLaneAuctionSetBeneficiary // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetBeneficiaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetBeneficiary)
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
		it.Event = new(IExpressLaneAuctionSetBeneficiary)
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
func (it *IExpressLaneAuctionSetBeneficiaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetBeneficiaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetBeneficiary represents a SetBeneficiary event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetBeneficiary struct {
	OldBeneficiary common.Address
	NewBeneficiary common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetBeneficiary is a free log retrieval operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetBeneficiary(opts *bind.FilterOpts) (*IExpressLaneAuctionSetBeneficiaryIterator, error) {

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetBeneficiary")
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetBeneficiaryIterator{contract: _IExpressLaneAuction.contract, event: "SetBeneficiary", logs: logs, sub: sub}, nil
}

// WatchSetBeneficiary is a free log subscription operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetBeneficiary(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetBeneficiary) (event.Subscription, error) {

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetBeneficiary")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetBeneficiary)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetBeneficiary", log); err != nil {
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

// ParseSetBeneficiary is a log parse operation binding the contract event 0x8a0149b2f3ddf2c9ee85738165131d82babbb938f749321d59f75750afa7f4e6.
//
// Solidity: event SetBeneficiary(address oldBeneficiary, address newBeneficiary)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetBeneficiary(log types.Log) (*IExpressLaneAuctionSetBeneficiary, error) {
	event := new(IExpressLaneAuctionSetBeneficiary)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetBeneficiary", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetExpressLaneControllerIterator is returned from FilterSetExpressLaneController and is used to iterate over the raw logs and unpacked data for SetExpressLaneController events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetExpressLaneControllerIterator struct {
	Event *IExpressLaneAuctionSetExpressLaneController // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetExpressLaneControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetExpressLaneController)
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
		it.Event = new(IExpressLaneAuctionSetExpressLaneController)
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
func (it *IExpressLaneAuctionSetExpressLaneControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetExpressLaneControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetExpressLaneController represents a SetExpressLaneController event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetExpressLaneController struct {
	Round                         uint64
	PreviousExpressLaneController common.Address
	NewExpressLaneController      common.Address
	Transferor                    common.Address
	StartTimestamp                uint64
	EndTimestamp                  uint64
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterSetExpressLaneController is a free log retrieval operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetExpressLaneController(opts *bind.FilterOpts, previousExpressLaneController []common.Address, newExpressLaneController []common.Address, transferor []common.Address) (*IExpressLaneAuctionSetExpressLaneControllerIterator, error) {

	var previousExpressLaneControllerRule []interface{}
	for _, previousExpressLaneControllerItem := range previousExpressLaneController {
		previousExpressLaneControllerRule = append(previousExpressLaneControllerRule, previousExpressLaneControllerItem)
	}
	var newExpressLaneControllerRule []interface{}
	for _, newExpressLaneControllerItem := range newExpressLaneController {
		newExpressLaneControllerRule = append(newExpressLaneControllerRule, newExpressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetExpressLaneController", previousExpressLaneControllerRule, newExpressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetExpressLaneControllerIterator{contract: _IExpressLaneAuction.contract, event: "SetExpressLaneController", logs: logs, sub: sub}, nil
}

// WatchSetExpressLaneController is a free log subscription operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetExpressLaneController(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetExpressLaneController, previousExpressLaneController []common.Address, newExpressLaneController []common.Address, transferor []common.Address) (event.Subscription, error) {

	var previousExpressLaneControllerRule []interface{}
	for _, previousExpressLaneControllerItem := range previousExpressLaneController {
		previousExpressLaneControllerRule = append(previousExpressLaneControllerRule, previousExpressLaneControllerItem)
	}
	var newExpressLaneControllerRule []interface{}
	for _, newExpressLaneControllerItem := range newExpressLaneController {
		newExpressLaneControllerRule = append(newExpressLaneControllerRule, newExpressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetExpressLaneController", previousExpressLaneControllerRule, newExpressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetExpressLaneController)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetExpressLaneController", log); err != nil {
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

// ParseSetExpressLaneController is a log parse operation binding the contract event 0xb59adc820ca642dad493a0a6e0bdf979dcae037dea114b70d5c66b1c0b791c4b.
//
// Solidity: event SetExpressLaneController(uint64 round, address indexed previousExpressLaneController, address indexed newExpressLaneController, address indexed transferor, uint64 startTimestamp, uint64 endTimestamp)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetExpressLaneController(log types.Log) (*IExpressLaneAuctionSetExpressLaneController, error) {
	event := new(IExpressLaneAuctionSetExpressLaneController)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetExpressLaneController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetMinReservePriceIterator is returned from FilterSetMinReservePrice and is used to iterate over the raw logs and unpacked data for SetMinReservePrice events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetMinReservePriceIterator struct {
	Event *IExpressLaneAuctionSetMinReservePrice // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetMinReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetMinReservePrice)
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
		it.Event = new(IExpressLaneAuctionSetMinReservePrice)
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
func (it *IExpressLaneAuctionSetMinReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetMinReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetMinReservePrice represents a SetMinReservePrice event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetMinReservePrice struct {
	OldPrice *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMinReservePrice is a free log retrieval operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetMinReservePrice(opts *bind.FilterOpts) (*IExpressLaneAuctionSetMinReservePriceIterator, error) {

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetMinReservePrice")
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetMinReservePriceIterator{contract: _IExpressLaneAuction.contract, event: "SetMinReservePrice", logs: logs, sub: sub}, nil
}

// WatchSetMinReservePrice is a free log subscription operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetMinReservePrice(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetMinReservePrice) (event.Subscription, error) {

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetMinReservePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetMinReservePrice)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetMinReservePrice", log); err != nil {
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

// ParseSetMinReservePrice is a log parse operation binding the contract event 0x5848068f11aa3ba9fe3fc33c5f9f2a3cd1aed67986b85b5e0cedc67dbe96f0f0.
//
// Solidity: event SetMinReservePrice(uint256 oldPrice, uint256 newPrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetMinReservePrice(log types.Log) (*IExpressLaneAuctionSetMinReservePrice, error) {
	event := new(IExpressLaneAuctionSetMinReservePrice)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetMinReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetReservePriceIterator is returned from FilterSetReservePrice and is used to iterate over the raw logs and unpacked data for SetReservePrice events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetReservePriceIterator struct {
	Event *IExpressLaneAuctionSetReservePrice // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetReservePrice)
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
		it.Event = new(IExpressLaneAuctionSetReservePrice)
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
func (it *IExpressLaneAuctionSetReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetReservePrice represents a SetReservePrice event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetReservePrice struct {
	OldReservePrice *big.Int
	NewReservePrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetReservePrice is a free log retrieval operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetReservePrice(opts *bind.FilterOpts) (*IExpressLaneAuctionSetReservePriceIterator, error) {

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetReservePrice")
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetReservePriceIterator{contract: _IExpressLaneAuction.contract, event: "SetReservePrice", logs: logs, sub: sub}, nil
}

// WatchSetReservePrice is a free log subscription operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetReservePrice(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetReservePrice) (event.Subscription, error) {

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetReservePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetReservePrice)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetReservePrice", log); err != nil {
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

// ParseSetReservePrice is a log parse operation binding the contract event 0x9725e37e079c5bda6009a8f54d86265849f30acf61c630f9e1ac91e67de98794.
//
// Solidity: event SetReservePrice(uint256 oldReservePrice, uint256 newReservePrice)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetReservePrice(log types.Log) (*IExpressLaneAuctionSetReservePrice, error) {
	event := new(IExpressLaneAuctionSetReservePrice)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetRoundTimingInfoIterator is returned from FilterSetRoundTimingInfo and is used to iterate over the raw logs and unpacked data for SetRoundTimingInfo events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetRoundTimingInfoIterator struct {
	Event *IExpressLaneAuctionSetRoundTimingInfo // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetRoundTimingInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetRoundTimingInfo)
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
		it.Event = new(IExpressLaneAuctionSetRoundTimingInfo)
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
func (it *IExpressLaneAuctionSetRoundTimingInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetRoundTimingInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetRoundTimingInfo represents a SetRoundTimingInfo event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetRoundTimingInfo struct {
	CurrentRound             uint64
	OffsetTimestamp          int64
	RoundDurationSeconds     uint64
	AuctionClosingSeconds    uint64
	ReserveSubmissionSeconds uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetRoundTimingInfo is a free log retrieval operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetRoundTimingInfo(opts *bind.FilterOpts) (*IExpressLaneAuctionSetRoundTimingInfoIterator, error) {

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetRoundTimingInfo")
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetRoundTimingInfoIterator{contract: _IExpressLaneAuction.contract, event: "SetRoundTimingInfo", logs: logs, sub: sub}, nil
}

// WatchSetRoundTimingInfo is a free log subscription operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetRoundTimingInfo(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetRoundTimingInfo) (event.Subscription, error) {

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetRoundTimingInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetRoundTimingInfo)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetRoundTimingInfo", log); err != nil {
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

// ParseSetRoundTimingInfo is a log parse operation binding the contract event 0x982cfb73783b8c64455c76cdeb1351467c4f1e6b3615fec07df232c1b46ffd47.
//
// Solidity: event SetRoundTimingInfo(uint64 currentRound, int64 offsetTimestamp, uint64 roundDurationSeconds, uint64 auctionClosingSeconds, uint64 reserveSubmissionSeconds)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetRoundTimingInfo(log types.Log) (*IExpressLaneAuctionSetRoundTimingInfo, error) {
	event := new(IExpressLaneAuctionSetRoundTimingInfo)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetRoundTimingInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionSetTransferorIterator is returned from FilterSetTransferor and is used to iterate over the raw logs and unpacked data for SetTransferor events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetTransferorIterator struct {
	Event *IExpressLaneAuctionSetTransferor // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionSetTransferorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionSetTransferor)
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
		it.Event = new(IExpressLaneAuctionSetTransferor)
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
func (it *IExpressLaneAuctionSetTransferorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionSetTransferorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionSetTransferor represents a SetTransferor event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionSetTransferor struct {
	ExpressLaneController common.Address
	Transferor            common.Address
	FixedUntilRound       uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetTransferor is a free log retrieval operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterSetTransferor(opts *bind.FilterOpts, expressLaneController []common.Address, transferor []common.Address) (*IExpressLaneAuctionSetTransferorIterator, error) {

	var expressLaneControllerRule []interface{}
	for _, expressLaneControllerItem := range expressLaneController {
		expressLaneControllerRule = append(expressLaneControllerRule, expressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "SetTransferor", expressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionSetTransferorIterator{contract: _IExpressLaneAuction.contract, event: "SetTransferor", logs: logs, sub: sub}, nil
}

// WatchSetTransferor is a free log subscription operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchSetTransferor(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionSetTransferor, expressLaneController []common.Address, transferor []common.Address) (event.Subscription, error) {

	var expressLaneControllerRule []interface{}
	for _, expressLaneControllerItem := range expressLaneController {
		expressLaneControllerRule = append(expressLaneControllerRule, expressLaneControllerItem)
	}
	var transferorRule []interface{}
	for _, transferorItem := range transferor {
		transferorRule = append(transferorRule, transferorItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "SetTransferor", expressLaneControllerRule, transferorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionSetTransferor)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetTransferor", log); err != nil {
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

// ParseSetTransferor is a log parse operation binding the contract event 0xf6d28df235d9fa45a42d45dbb7c4f4ac76edb51e528f09f25a0650d32b8b33c0.
//
// Solidity: event SetTransferor(address indexed expressLaneController, address indexed transferor, uint64 fixedUntilRound)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseSetTransferor(log types.Log) (*IExpressLaneAuctionSetTransferor, error) {
	event := new(IExpressLaneAuctionSetTransferor)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "SetTransferor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionWithdrawalFinalizedIterator struct {
	Event *IExpressLaneAuctionWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionWithdrawalFinalized)
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
		it.Event = new(IExpressLaneAuctionWithdrawalFinalized)
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
func (it *IExpressLaneAuctionWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionWithdrawalFinalized represents a WithdrawalFinalized event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionWithdrawalFinalized struct {
	Account          common.Address
	WithdrawalAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, account []common.Address) (*IExpressLaneAuctionWithdrawalFinalizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "WithdrawalFinalized", accountRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionWithdrawalFinalizedIterator{contract: _IExpressLaneAuction.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionWithdrawalFinalized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "WithdrawalFinalized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionWithdrawalFinalized)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0x9e5c4f9f4e46b8629d3dda85f43a69194f50254404a72dc62b9e932d9c94eda8.
//
// Solidity: event WithdrawalFinalized(address indexed account, uint256 withdrawalAmount)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseWithdrawalFinalized(log types.Log) (*IExpressLaneAuctionWithdrawalFinalized, error) {
	event := new(IExpressLaneAuctionWithdrawalFinalized)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExpressLaneAuctionWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionWithdrawalInitiatedIterator struct {
	Event *IExpressLaneAuctionWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *IExpressLaneAuctionWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IExpressLaneAuctionWithdrawalInitiated)
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
		it.Event = new(IExpressLaneAuctionWithdrawalInitiated)
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
func (it *IExpressLaneAuctionWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IExpressLaneAuctionWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IExpressLaneAuctionWithdrawalInitiated represents a WithdrawalInitiated event raised by the IExpressLaneAuction contract.
type IExpressLaneAuctionWithdrawalInitiated struct {
	Account           common.Address
	WithdrawalAmount  *big.Int
	RoundWithdrawable *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, account []common.Address) (*IExpressLaneAuctionWithdrawalInitiatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.FilterLogs(opts, "WithdrawalInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return &IExpressLaneAuctionWithdrawalInitiatedIterator{contract: _IExpressLaneAuction.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *IExpressLaneAuctionWithdrawalInitiated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IExpressLaneAuction.contract.WatchLogs(opts, "WithdrawalInitiated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IExpressLaneAuctionWithdrawalInitiated)
				if err := _IExpressLaneAuction.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x31f69201fab7912e3ec9850e3ab705964bf46d9d4276bdcbb6d05e965e5f5401.
//
// Solidity: event WithdrawalInitiated(address indexed account, uint256 withdrawalAmount, uint256 roundWithdrawable)
func (_IExpressLaneAuction *IExpressLaneAuctionFilterer) ParseWithdrawalInitiated(log types.Log) (*IExpressLaneAuctionWithdrawalInitiated, error) {
	event := new(IExpressLaneAuctionWithdrawalInitiated)
	if err := _IExpressLaneAuction.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LatestELCRoundsLibMetaData contains all meta data concerning the LatestELCRoundsLib contract.
var LatestELCRoundsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220eff5b8996e703c22fa5045c8ed2216b44f6564f8731686b8fe0363f974e00bbc64736f6c63430008190033",
}

// LatestELCRoundsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use LatestELCRoundsLibMetaData.ABI instead.
var LatestELCRoundsLibABI = LatestELCRoundsLibMetaData.ABI

// LatestELCRoundsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LatestELCRoundsLibMetaData.Bin instead.
var LatestELCRoundsLibBin = LatestELCRoundsLibMetaData.Bin

// DeployLatestELCRoundsLib deploys a new Ethereum contract, binding an instance of LatestELCRoundsLib to it.
func DeployLatestELCRoundsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LatestELCRoundsLib, error) {
	parsed, err := LatestELCRoundsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LatestELCRoundsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LatestELCRoundsLib{LatestELCRoundsLibCaller: LatestELCRoundsLibCaller{contract: contract}, LatestELCRoundsLibTransactor: LatestELCRoundsLibTransactor{contract: contract}, LatestELCRoundsLibFilterer: LatestELCRoundsLibFilterer{contract: contract}}, nil
}

// LatestELCRoundsLib is an auto generated Go binding around an Ethereum contract.
type LatestELCRoundsLib struct {
	LatestELCRoundsLibCaller     // Read-only binding to the contract
	LatestELCRoundsLibTransactor // Write-only binding to the contract
	LatestELCRoundsLibFilterer   // Log filterer for contract events
}

// LatestELCRoundsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type LatestELCRoundsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LatestELCRoundsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LatestELCRoundsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LatestELCRoundsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LatestELCRoundsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LatestELCRoundsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LatestELCRoundsLibSession struct {
	Contract     *LatestELCRoundsLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LatestELCRoundsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LatestELCRoundsLibCallerSession struct {
	Contract *LatestELCRoundsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LatestELCRoundsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LatestELCRoundsLibTransactorSession struct {
	Contract     *LatestELCRoundsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LatestELCRoundsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type LatestELCRoundsLibRaw struct {
	Contract *LatestELCRoundsLib // Generic contract binding to access the raw methods on
}

// LatestELCRoundsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LatestELCRoundsLibCallerRaw struct {
	Contract *LatestELCRoundsLibCaller // Generic read-only contract binding to access the raw methods on
}

// LatestELCRoundsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LatestELCRoundsLibTransactorRaw struct {
	Contract *LatestELCRoundsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLatestELCRoundsLib creates a new instance of LatestELCRoundsLib, bound to a specific deployed contract.
func NewLatestELCRoundsLib(address common.Address, backend bind.ContractBackend) (*LatestELCRoundsLib, error) {
	contract, err := bindLatestELCRoundsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LatestELCRoundsLib{LatestELCRoundsLibCaller: LatestELCRoundsLibCaller{contract: contract}, LatestELCRoundsLibTransactor: LatestELCRoundsLibTransactor{contract: contract}, LatestELCRoundsLibFilterer: LatestELCRoundsLibFilterer{contract: contract}}, nil
}

// NewLatestELCRoundsLibCaller creates a new read-only instance of LatestELCRoundsLib, bound to a specific deployed contract.
func NewLatestELCRoundsLibCaller(address common.Address, caller bind.ContractCaller) (*LatestELCRoundsLibCaller, error) {
	contract, err := bindLatestELCRoundsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LatestELCRoundsLibCaller{contract: contract}, nil
}

// NewLatestELCRoundsLibTransactor creates a new write-only instance of LatestELCRoundsLib, bound to a specific deployed contract.
func NewLatestELCRoundsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*LatestELCRoundsLibTransactor, error) {
	contract, err := bindLatestELCRoundsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LatestELCRoundsLibTransactor{contract: contract}, nil
}

// NewLatestELCRoundsLibFilterer creates a new log filterer instance of LatestELCRoundsLib, bound to a specific deployed contract.
func NewLatestELCRoundsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*LatestELCRoundsLibFilterer, error) {
	contract, err := bindLatestELCRoundsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LatestELCRoundsLibFilterer{contract: contract}, nil
}

// bindLatestELCRoundsLib binds a generic wrapper to an already deployed contract.
func bindLatestELCRoundsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LatestELCRoundsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LatestELCRoundsLib *LatestELCRoundsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LatestELCRoundsLib.Contract.LatestELCRoundsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LatestELCRoundsLib *LatestELCRoundsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LatestELCRoundsLib.Contract.LatestELCRoundsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LatestELCRoundsLib *LatestELCRoundsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LatestELCRoundsLib.Contract.LatestELCRoundsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LatestELCRoundsLib *LatestELCRoundsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LatestELCRoundsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LatestELCRoundsLib *LatestELCRoundsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LatestELCRoundsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LatestELCRoundsLib *LatestELCRoundsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LatestELCRoundsLib.Contract.contract.Transact(opts, method, params...)
}

// RoundTimingInfoLibMetaData contains all meta data concerning the RoundTimingInfoLib contract.
var RoundTimingInfoLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220a428d8ab673e80f0f4fcb90c73e455074ead10d196b3b1e78efcd53a843bc25564736f6c63430008190033",
}

// RoundTimingInfoLibABI is the input ABI used to generate the binding from.
// Deprecated: Use RoundTimingInfoLibMetaData.ABI instead.
var RoundTimingInfoLibABI = RoundTimingInfoLibMetaData.ABI

// RoundTimingInfoLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoundTimingInfoLibMetaData.Bin instead.
var RoundTimingInfoLibBin = RoundTimingInfoLibMetaData.Bin

// DeployRoundTimingInfoLib deploys a new Ethereum contract, binding an instance of RoundTimingInfoLib to it.
func DeployRoundTimingInfoLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoundTimingInfoLib, error) {
	parsed, err := RoundTimingInfoLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoundTimingInfoLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoundTimingInfoLib{RoundTimingInfoLibCaller: RoundTimingInfoLibCaller{contract: contract}, RoundTimingInfoLibTransactor: RoundTimingInfoLibTransactor{contract: contract}, RoundTimingInfoLibFilterer: RoundTimingInfoLibFilterer{contract: contract}}, nil
}

// RoundTimingInfoLib is an auto generated Go binding around an Ethereum contract.
type RoundTimingInfoLib struct {
	RoundTimingInfoLibCaller     // Read-only binding to the contract
	RoundTimingInfoLibTransactor // Write-only binding to the contract
	RoundTimingInfoLibFilterer   // Log filterer for contract events
}

// RoundTimingInfoLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoundTimingInfoLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundTimingInfoLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoundTimingInfoLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundTimingInfoLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoundTimingInfoLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundTimingInfoLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoundTimingInfoLibSession struct {
	Contract     *RoundTimingInfoLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RoundTimingInfoLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoundTimingInfoLibCallerSession struct {
	Contract *RoundTimingInfoLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RoundTimingInfoLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoundTimingInfoLibTransactorSession struct {
	Contract     *RoundTimingInfoLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RoundTimingInfoLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoundTimingInfoLibRaw struct {
	Contract *RoundTimingInfoLib // Generic contract binding to access the raw methods on
}

// RoundTimingInfoLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoundTimingInfoLibCallerRaw struct {
	Contract *RoundTimingInfoLibCaller // Generic read-only contract binding to access the raw methods on
}

// RoundTimingInfoLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoundTimingInfoLibTransactorRaw struct {
	Contract *RoundTimingInfoLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoundTimingInfoLib creates a new instance of RoundTimingInfoLib, bound to a specific deployed contract.
func NewRoundTimingInfoLib(address common.Address, backend bind.ContractBackend) (*RoundTimingInfoLib, error) {
	contract, err := bindRoundTimingInfoLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoundTimingInfoLib{RoundTimingInfoLibCaller: RoundTimingInfoLibCaller{contract: contract}, RoundTimingInfoLibTransactor: RoundTimingInfoLibTransactor{contract: contract}, RoundTimingInfoLibFilterer: RoundTimingInfoLibFilterer{contract: contract}}, nil
}

// NewRoundTimingInfoLibCaller creates a new read-only instance of RoundTimingInfoLib, bound to a specific deployed contract.
func NewRoundTimingInfoLibCaller(address common.Address, caller bind.ContractCaller) (*RoundTimingInfoLibCaller, error) {
	contract, err := bindRoundTimingInfoLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoundTimingInfoLibCaller{contract: contract}, nil
}

// NewRoundTimingInfoLibTransactor creates a new write-only instance of RoundTimingInfoLib, bound to a specific deployed contract.
func NewRoundTimingInfoLibTransactor(address common.Address, transactor bind.ContractTransactor) (*RoundTimingInfoLibTransactor, error) {
	contract, err := bindRoundTimingInfoLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoundTimingInfoLibTransactor{contract: contract}, nil
}

// NewRoundTimingInfoLibFilterer creates a new log filterer instance of RoundTimingInfoLib, bound to a specific deployed contract.
func NewRoundTimingInfoLibFilterer(address common.Address, filterer bind.ContractFilterer) (*RoundTimingInfoLibFilterer, error) {
	contract, err := bindRoundTimingInfoLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoundTimingInfoLibFilterer{contract: contract}, nil
}

// bindRoundTimingInfoLib binds a generic wrapper to an already deployed contract.
func bindRoundTimingInfoLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RoundTimingInfoLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoundTimingInfoLib *RoundTimingInfoLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoundTimingInfoLib.Contract.RoundTimingInfoLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoundTimingInfoLib *RoundTimingInfoLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoundTimingInfoLib.Contract.RoundTimingInfoLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoundTimingInfoLib *RoundTimingInfoLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoundTimingInfoLib.Contract.RoundTimingInfoLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoundTimingInfoLib *RoundTimingInfoLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoundTimingInfoLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoundTimingInfoLib *RoundTimingInfoLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoundTimingInfoLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoundTimingInfoLib *RoundTimingInfoLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoundTimingInfoLib.Contract.contract.Transact(opts, method, params...)
}
