// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengeV2gen

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

// AddLeafArgs is an auto generated low-level Go binding around an user-defined struct.
type AddLeafArgs struct {
	ChallengeId            [32]byte
	ClaimId                [32]byte
	Height                 *big.Int
	HistoryRoot            [32]byte
	FirstState             [32]byte
	FirstStatehistoryProof []byte
	LastState              [32]byte
	LastStatehistoryProof  []byte
}

// Assertion is an auto generated low-level Go binding around an user-defined struct.
type Assertion struct {
	PredecessorId           [32]byte
	SuccessionChallenge     [32]byte
	IsFirstChild            bool
	SecondChildCreationTime *big.Int
	FirstChildCreationTime  *big.Int
	StateHash               [32]byte
	Height                  *big.Int
	Status                  uint8
	InboxMsgCountSeen       *big.Int
}

// Challenge is an auto generated low-level Go binding around an user-defined struct.
type Challenge struct {
	RootId        [32]byte
	WinningClaim  [32]byte
	ChallengeType uint8
}

// ChallengeVertex is an auto generated low-level Go binding around an user-defined struct.
type ChallengeVertex struct {
	ChallengeId             [32]byte
	HistoryRoot             [32]byte
	Height                  *big.Int
	SuccessionChallenge     [32]byte
	PredecessorId           [32]byte
	ClaimId                 [32]byte
	Staker                  common.Address
	Status                  uint8
	PsId                    [32]byte
	PsLastUpdatedTimestamp  *big.Int
	FlushedPsTimeSec        *big.Int
	LowestHeightSuccessorId [32]byte
}

// ExecutionContext is an auto generated low-level Go binding around an user-defined struct.
type ExecutionContext struct {
	MaxInboxMessagesRead *big.Int
	Bridge               common.Address
}

// OneStepData is an auto generated low-level Go binding around an user-defined struct.
type OneStepData struct {
	ExecCtx     ExecutionContext
	MachineStep *big.Int
	BeforeHash  [32]byte
	Proof       []byte
}

// AssertionChainMetaData contains all meta data concerning the AssertionChain contract.
var AssertionChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"NotConfirmable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"NotRejectable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"assertionExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"assertions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"successionChallenge\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"enumStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"inboxMsgCountSeen\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeManagerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriodSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"confirmAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"}],\"name\":\"createNewAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"createSuccessionChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAssertion\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"successionChallenge\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isFirstChild\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"secondChildCreationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstChildCreationTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"enumStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"inboxMsgCountSeen\",\"type\":\"uint256\"}],\"internalType\":\"structAssertion\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getFirstChildCreationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getInboxMsgCountSeen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getPredecessorId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getStateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getSuccessionChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"isFirstChild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"rejectAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"_challengeManager\",\"type\":\"address\"}],\"name\":\"updateChallengeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405268056bc75e2d6310000060805234801561001d57600080fd5b5060405161162c38038061162c83398101604081905261003c916101f9565b6002818155604080516101208101825260008082526020808301828152938301828152606084018381526080850184815260a086018a815260c08701868152600160e089018181526101008a018990528880529681905288517fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4990815599517fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4a5594517fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4b805491151560ff1992831617905593517fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4c5591517fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4d55517fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4e55517fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4f5591517fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb50805494979596959194909316919084908111156101de576101de61021d565b02179055506101008201518160080155905050505050610233565b6000806040838503121561020c57600080fd5b505080516020909101519092909150565b634e487b7160e01b600052602160045260246000fd5b6080516113d76102556000396000818161025f01526108b801526113d76000f3fe6080604052600436106101045760003560e01c80636894bdd5116100a05780639ca565d4116100645780639ca565d41461032e578063d60715b51461034e578063f9bce634146103cf578063fb601294146103ef578063ff8aef871461040557600080fd5b80636894bdd51461028157806375dc6098146102a15780637cfd5ab9146102c157806388302884146102e1578063896efbf21461030e57600080fd5b8063295dfd32146101095780632fefa18a14610148578063308362281461017b57806343ed6ad9146101ab57806349635f9a146101cb5780635625c360146101eb5780635a4038f5146102135780635a627dbc1461024557806360c7dc471461024d575b600080fd5b34801561011557600080fd5b506101466101243660046111a2565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b005b34801561015457600080fd5b506101686101633660046111d2565b610425565b6040519081526020015b60405180910390f35b34801561018757600080fd5b5061019b6101963660046111d2565b610473565b6040519015158152602001610172565b3480156101b757600080fd5b506101686101c63660046111d2565b6104ba565b3480156101d757600080fd5b506101466101e63660046111eb565b6104fe565b3480156101f757600080fd5b506000546040516001600160a01b039091168152602001610172565b34801561021f57600080fd5b5061019b61022e3660046111d2565b600090815260016020526040902060050154151590565b6101466108b6565b34801561025957600080fd5b506101687f000000000000000000000000000000000000000000000000000000000000000081565b34801561028d57600080fd5b5061014661029c3660046111d2565b610927565b3480156102ad57600080fd5b506101466102bc3660046111d2565b610b38565b3480156102cd57600080fd5b506101686102dc3660046111d2565b610d43565b3480156102ed57600080fd5b506103016102fc3660046111d2565b610d87565b604051610172919061124f565b34801561031a57600080fd5b506101686103293660046111d2565b610e7b565b34801561033a57600080fd5b506101686103493660046111d2565b610ebf565b34801561035a57600080fd5b506103ba6103693660046111d2565b6001602081905260009182526040909120805491810154600282015460038301546004840154600585015460068601546007870154600890970154959660ff95861696949593949293919291169089565b604051610172999897969594939291906112bc565b3480156103db57600080fd5b506101686103ea3660046111d2565b610f00565b3480156103fb57600080fd5b5061016860025481565b34801561041157600080fd5b506101466104203660046111d2565b610f44565b60008181526001602052604081206005015461045c5760405162461bcd60e51b815260040161045390611310565b60405180910390fd5b506000908152600160208190526040909120015490565b6000818152600160205260408120600501546104a15760405162461bcd60e51b815260040161045390611310565b5060009081526001602052604090206002015460ff1690565b6000818152600160205260408120600501546104e85760405162461bcd60e51b815260040161045390611310565b5060009081526001602052604090206004015490565b60408051602081018590529081018390526060810182905260009060800160405160208183030381529060405280519060200120905061054f81600090815260016020526040902060050154151590565b156105975760405162461bcd60e51b8152602060048201526018602482015277417373657274696f6e20616c72656164792065786973747360401b6044820152606401610453565b6000828152600160205260409020600501546105ff5760405162461bcd60e51b815260206004820152602160248201527f50726576696f757320617373657274696f6e20646f6573206e6f7420657869736044820152601d60fa1b6064820152608401610453565b600260008281526001602052604080822054825290206007015460ff16600281111561062d5761062d611217565b0361067a5760405162461bcd60e51b815260206004820152601b60248201527f50726576696f757320617373657274696f6e2072656a656374656400000000006044820152606401610453565b6000818152600160205260408082205482529020839060060154106106ed5760405162461bcd60e51b815260206004820152602360248201527f486569676874206e6f742067726561746572207468616e20707265646563657360448201526239b7b960e91b6064820152608401610453565b600082815260016020526040902060040154151580610720576000838152600160205260409020426004909101556107c1565b6002546000838152600160205260408082205482529020600401546107459190611358565b42106107935760405162461bcd60e51b815260206004820152601a60248201527f546f6f206c61746520746f20637265617465207369626c696e670000000000006044820152606401610453565b60008381526001602052604081206003015490036107c1576000838152600160205260409020426003909101555b6040518061012001604052808481526020016000801b815260200182151515815260200160008152602001600081526020018681526020018581526020016000600281111561081257610812611217565b8152600060209182018190528481526001808352604091829020845181559284015183820155908301516002808401805492151560ff19938416179055606085015160038501556080850151600485015560a0850151600585015560c0850151600685015560e0850151600785018054919490939190911691849081111561089c5761089c611217565b021790555061010082015181600801559050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000034146109255760405162461bcd60e51b815260206004820152601a60248201527f436f7272656374207374616b65206e6f742070726f76696465640000000000006044820152606401610453565b565b6000818152600160205260409020600501546109555760405162461bcd60e51b815260040161045390611310565b600160008281526001602052604080822054825290206007015460ff16600281111561098357610983611217565b146109d05760405162461bcd60e51b815260206004820181905260248201527f50726576696f757320617373657274696f6e206e6f7420636f6e6669726d65646044820152606401610453565b600081815260016020526040808220548252902060030154158015610a185750600254600082815260016020526040808220548252902060040154610a159190611358565b42115b15610a42576000818152600160208190526040909120600701805460ff191682805b021790555050565b60008181526001602081905260408083205483528220015490819003610a7e57604051631895e8f560e21b815260048101839052602401610453565b60008054604051630e7a2a9d60e31b8152600481018490526001600160a01b03909116906373d154e890602401602060405180830381865afa158015610ac8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aec9190611371565b9050828114610b1157604051632158b7ff60e11b815260048101849052602401610453565b6000838152600160208190526040909120600701805460ff191682805b0217905550505050565b600081815260016020526040902060050154610b665760405162461bcd60e51b815260040161045390611310565b60008181526001602052604081206007015460ff166002811115610b8c57610b8c611217565b14610bd45760405162461bcd60e51b8152602060048201526018602482015277417373657274696f6e206973206e6f742070656e64696e6760401b6044820152606401610453565b600260008281526001602052604080822054825290206007015460ff166002811115610c0257610c02611217565b03610c2d576000818152600160208190526040909120600701805460029260ff199091169083610a3a565b60008181526001602081905260408083205483528220015490819003610c6957604051632158b7ff60e11b815260048101839052602401610453565b60008054604051630e7a2a9d60e31b8152600481018490526001600160a01b03909116906373d154e890602401602060405180830381865afa158015610cb3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd79190611371565b905080610cfa57604051632158b7ff60e11b815260048101849052602401610453565b828103610d1d57604051632158b7ff60e11b815260048101849052602401610453565b6000838152600160208190526040909120600701805460029260ff199091169083610b2e565b600081815260016020526040812060050154610d715760405162461bcd60e51b815260040161045390611310565b5060009081526001602052604090206008015490565b6040805161012081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081019190915260008281526001602081815260409283902083516101208101855281548152928101549183019190915260028082015460ff9081161515948401949094526003820154606084015260048201546080840152600582015460a0840152600682015460c084015260078201549293919260e08501921690811115610e5557610e55611217565b6002811115610e6657610e66611217565b81526020016008820154815250509050919050565b600081815260016020526040812060050154610ea95760405162461bcd60e51b815260040161045390611310565b5060009081526001602052604090206006015490565b600081815260016020526040812060050154610eed5760405162461bcd60e51b815260040161045390611310565b5060009081526001602052604090205490565b600081815260016020526040812060050154610f2e5760405162461bcd60e51b815260040161045390611310565b5060009081526001602052604090206005015490565b600081815260016020526040902060050154610f725760405162461bcd60e51b815260040161045390611310565b600260008281526001602052604090206007015460ff166002811115610f9a57610f9a611217565b03610fe75760405162461bcd60e51b815260206004820152601a60248201527f417373657274696f6e20616c72656164792072656a65637465640000000000006044820152606401610453565b60008181526001602081905260409091200154156110435760405162461bcd60e51b815260206004820152601960248201527810da185b1b195b99d948185b1c9958591e4818dc99585d1959603a1b6044820152606401610453565b60008181526001602052604081206003015490036110ad5760405162461bcd60e51b815260206004820152602160248201527f4174206c656173742074776f206368696c6472656e206e6f74206372656174656044820152601960fa1b6064820152608401610453565b600280546110ba9161138a565b6000828152600160205260409020600401546110d69190611358565b421061111c5760405162461bcd60e51b8152602060048201526015602482015274546f6f206c61746520746f206368616c6c656e676560581b6044820152606401610453565b60005460405163f696dc5560e01b8152600481018390526001600160a01b039091169063f696dc55906024016020604051808303816000875af1158015611167573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118b9190611371565b600091825260016020819052604090922090910155565b6000602082840312156111b457600080fd5b81356001600160a01b03811681146111cb57600080fd5b9392505050565b6000602082840312156111e457600080fd5b5035919050565b60008060006060848603121561120057600080fd5b505081359360208301359350604090920135919050565b634e487b7160e01b600052602160045260246000fd5b6003811061124b57634e487b7160e01b600052602160045260246000fd5b9052565b6000610120820190508251825260208301516020830152604083015115156040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015260e08301516112ac60e084018261122d565b5061010092830151919092015290565b6000610120820190508a825289602083015288151560408301528760608301528660808301528560a08301528460c08301526112fb60e083018561122d565b826101008301529a9950505050505050505050565b602080825260189082015277105cdcd95c9d1a5bdb88191bd95cc81b9bdd08195e1a5cdd60421b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b8082018082111561136b5761136b611342565b92915050565b60006020828403121561138357600080fd5b5051919050565b808202811582820484141761136b5761136b61134256fea2646970667358221220cbfe3d57fcd9ffe61762e23ecfa588c83e0a2312fabb708c1cbc775f1e0070c364736f6c63430008110033",
}

// AssertionChainABI is the input ABI used to generate the binding from.
// Deprecated: Use AssertionChainMetaData.ABI instead.
var AssertionChainABI = AssertionChainMetaData.ABI

// AssertionChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssertionChainMetaData.Bin instead.
var AssertionChainBin = AssertionChainMetaData.Bin

// DeployAssertionChain deploys a new Ethereum contract, binding an instance of AssertionChain to it.
func DeployAssertionChain(auth *bind.TransactOpts, backend bind.ContractBackend, stateHash [32]byte, _challengePeriodSeconds *big.Int) (common.Address, *types.Transaction, *AssertionChain, error) {
	parsed, err := AssertionChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssertionChainBin), backend, stateHash, _challengePeriodSeconds)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssertionChain{AssertionChainCaller: AssertionChainCaller{contract: contract}, AssertionChainTransactor: AssertionChainTransactor{contract: contract}, AssertionChainFilterer: AssertionChainFilterer{contract: contract}}, nil
}

// AssertionChain is an auto generated Go binding around an Ethereum contract.
type AssertionChain struct {
	AssertionChainCaller     // Read-only binding to the contract
	AssertionChainTransactor // Write-only binding to the contract
	AssertionChainFilterer   // Log filterer for contract events
}

// AssertionChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssertionChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssertionChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssertionChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssertionChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssertionChainSession struct {
	Contract     *AssertionChain   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssertionChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssertionChainCallerSession struct {
	Contract *AssertionChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AssertionChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssertionChainTransactorSession struct {
	Contract     *AssertionChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AssertionChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssertionChainRaw struct {
	Contract *AssertionChain // Generic contract binding to access the raw methods on
}

// AssertionChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssertionChainCallerRaw struct {
	Contract *AssertionChainCaller // Generic read-only contract binding to access the raw methods on
}

// AssertionChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssertionChainTransactorRaw struct {
	Contract *AssertionChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssertionChain creates a new instance of AssertionChain, bound to a specific deployed contract.
func NewAssertionChain(address common.Address, backend bind.ContractBackend) (*AssertionChain, error) {
	contract, err := bindAssertionChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssertionChain{AssertionChainCaller: AssertionChainCaller{contract: contract}, AssertionChainTransactor: AssertionChainTransactor{contract: contract}, AssertionChainFilterer: AssertionChainFilterer{contract: contract}}, nil
}

// NewAssertionChainCaller creates a new read-only instance of AssertionChain, bound to a specific deployed contract.
func NewAssertionChainCaller(address common.Address, caller bind.ContractCaller) (*AssertionChainCaller, error) {
	contract, err := bindAssertionChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionChainCaller{contract: contract}, nil
}

// NewAssertionChainTransactor creates a new write-only instance of AssertionChain, bound to a specific deployed contract.
func NewAssertionChainTransactor(address common.Address, transactor bind.ContractTransactor) (*AssertionChainTransactor, error) {
	contract, err := bindAssertionChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssertionChainTransactor{contract: contract}, nil
}

// NewAssertionChainFilterer creates a new log filterer instance of AssertionChain, bound to a specific deployed contract.
func NewAssertionChainFilterer(address common.Address, filterer bind.ContractFilterer) (*AssertionChainFilterer, error) {
	contract, err := bindAssertionChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssertionChainFilterer{contract: contract}, nil
}

// bindAssertionChain binds a generic wrapper to an already deployed contract.
func bindAssertionChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssertionChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionChain *AssertionChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionChain.Contract.AssertionChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionChain *AssertionChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionChain.Contract.AssertionChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionChain *AssertionChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionChain.Contract.AssertionChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssertionChain *AssertionChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssertionChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssertionChain *AssertionChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssertionChain *AssertionChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssertionChain.Contract.contract.Transact(opts, method, params...)
}

// AssertionExists is a free data retrieval call binding the contract method 0x5a4038f5.
//
// Solidity: function assertionExists(bytes32 assertionId) view returns(bool)
func (_AssertionChain *AssertionChainCaller) AssertionExists(opts *bind.CallOpts, assertionId [32]byte) (bool, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "assertionExists", assertionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssertionExists is a free data retrieval call binding the contract method 0x5a4038f5.
//
// Solidity: function assertionExists(bytes32 assertionId) view returns(bool)
func (_AssertionChain *AssertionChainSession) AssertionExists(assertionId [32]byte) (bool, error) {
	return _AssertionChain.Contract.AssertionExists(&_AssertionChain.CallOpts, assertionId)
}

// AssertionExists is a free data retrieval call binding the contract method 0x5a4038f5.
//
// Solidity: function assertionExists(bytes32 assertionId) view returns(bool)
func (_AssertionChain *AssertionChainCallerSession) AssertionExists(assertionId [32]byte) (bool, error) {
	return _AssertionChain.Contract.AssertionExists(&_AssertionChain.CallOpts, assertionId)
}

// Assertions is a free data retrieval call binding the contract method 0xd60715b5.
//
// Solidity: function assertions(bytes32 ) view returns(bytes32 predecessorId, bytes32 successionChallenge, bool isFirstChild, uint256 secondChildCreationTime, uint256 firstChildCreationTime, bytes32 stateHash, uint256 height, uint8 status, uint256 inboxMsgCountSeen)
func (_AssertionChain *AssertionChainCaller) Assertions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	PredecessorId           [32]byte
	SuccessionChallenge     [32]byte
	IsFirstChild            bool
	SecondChildCreationTime *big.Int
	FirstChildCreationTime  *big.Int
	StateHash               [32]byte
	Height                  *big.Int
	Status                  uint8
	InboxMsgCountSeen       *big.Int
}, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "assertions", arg0)

	outstruct := new(struct {
		PredecessorId           [32]byte
		SuccessionChallenge     [32]byte
		IsFirstChild            bool
		SecondChildCreationTime *big.Int
		FirstChildCreationTime  *big.Int
		StateHash               [32]byte
		Height                  *big.Int
		Status                  uint8
		InboxMsgCountSeen       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PredecessorId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.SuccessionChallenge = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.IsFirstChild = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.SecondChildCreationTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FirstChildCreationTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StateHash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Height = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.InboxMsgCountSeen = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assertions is a free data retrieval call binding the contract method 0xd60715b5.
//
// Solidity: function assertions(bytes32 ) view returns(bytes32 predecessorId, bytes32 successionChallenge, bool isFirstChild, uint256 secondChildCreationTime, uint256 firstChildCreationTime, bytes32 stateHash, uint256 height, uint8 status, uint256 inboxMsgCountSeen)
func (_AssertionChain *AssertionChainSession) Assertions(arg0 [32]byte) (struct {
	PredecessorId           [32]byte
	SuccessionChallenge     [32]byte
	IsFirstChild            bool
	SecondChildCreationTime *big.Int
	FirstChildCreationTime  *big.Int
	StateHash               [32]byte
	Height                  *big.Int
	Status                  uint8
	InboxMsgCountSeen       *big.Int
}, error) {
	return _AssertionChain.Contract.Assertions(&_AssertionChain.CallOpts, arg0)
}

// Assertions is a free data retrieval call binding the contract method 0xd60715b5.
//
// Solidity: function assertions(bytes32 ) view returns(bytes32 predecessorId, bytes32 successionChallenge, bool isFirstChild, uint256 secondChildCreationTime, uint256 firstChildCreationTime, bytes32 stateHash, uint256 height, uint8 status, uint256 inboxMsgCountSeen)
func (_AssertionChain *AssertionChainCallerSession) Assertions(arg0 [32]byte) (struct {
	PredecessorId           [32]byte
	SuccessionChallenge     [32]byte
	IsFirstChild            bool
	SecondChildCreationTime *big.Int
	FirstChildCreationTime  *big.Int
	StateHash               [32]byte
	Height                  *big.Int
	Status                  uint8
	InboxMsgCountSeen       *big.Int
}, error) {
	return _AssertionChain.Contract.Assertions(&_AssertionChain.CallOpts, arg0)
}

// ChallengeManagerAddr is a free data retrieval call binding the contract method 0x5625c360.
//
// Solidity: function challengeManagerAddr() view returns(address)
func (_AssertionChain *AssertionChainCaller) ChallengeManagerAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "challengeManagerAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeManagerAddr is a free data retrieval call binding the contract method 0x5625c360.
//
// Solidity: function challengeManagerAddr() view returns(address)
func (_AssertionChain *AssertionChainSession) ChallengeManagerAddr() (common.Address, error) {
	return _AssertionChain.Contract.ChallengeManagerAddr(&_AssertionChain.CallOpts)
}

// ChallengeManagerAddr is a free data retrieval call binding the contract method 0x5625c360.
//
// Solidity: function challengeManagerAddr() view returns(address)
func (_AssertionChain *AssertionChainCallerSession) ChallengeManagerAddr() (common.Address, error) {
	return _AssertionChain.Contract.ChallengeManagerAddr(&_AssertionChain.CallOpts)
}

// ChallengePeriodSeconds is a free data retrieval call binding the contract method 0xfb601294.
//
// Solidity: function challengePeriodSeconds() view returns(uint256)
func (_AssertionChain *AssertionChainCaller) ChallengePeriodSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "challengePeriodSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengePeriodSeconds is a free data retrieval call binding the contract method 0xfb601294.
//
// Solidity: function challengePeriodSeconds() view returns(uint256)
func (_AssertionChain *AssertionChainSession) ChallengePeriodSeconds() (*big.Int, error) {
	return _AssertionChain.Contract.ChallengePeriodSeconds(&_AssertionChain.CallOpts)
}

// ChallengePeriodSeconds is a free data retrieval call binding the contract method 0xfb601294.
//
// Solidity: function challengePeriodSeconds() view returns(uint256)
func (_AssertionChain *AssertionChainCallerSession) ChallengePeriodSeconds() (*big.Int, error) {
	return _AssertionChain.Contract.ChallengePeriodSeconds(&_AssertionChain.CallOpts)
}

// GetAssertion is a free data retrieval call binding the contract method 0x88302884.
//
// Solidity: function getAssertion(bytes32 id) view returns((bytes32,bytes32,bool,uint256,uint256,bytes32,uint256,uint8,uint256))
func (_AssertionChain *AssertionChainCaller) GetAssertion(opts *bind.CallOpts, id [32]byte) (Assertion, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getAssertion", id)

	if err != nil {
		return *new(Assertion), err
	}

	out0 := *abi.ConvertType(out[0], new(Assertion)).(*Assertion)

	return out0, err

}

// GetAssertion is a free data retrieval call binding the contract method 0x88302884.
//
// Solidity: function getAssertion(bytes32 id) view returns((bytes32,bytes32,bool,uint256,uint256,bytes32,uint256,uint8,uint256))
func (_AssertionChain *AssertionChainSession) GetAssertion(id [32]byte) (Assertion, error) {
	return _AssertionChain.Contract.GetAssertion(&_AssertionChain.CallOpts, id)
}

// GetAssertion is a free data retrieval call binding the contract method 0x88302884.
//
// Solidity: function getAssertion(bytes32 id) view returns((bytes32,bytes32,bool,uint256,uint256,bytes32,uint256,uint8,uint256))
func (_AssertionChain *AssertionChainCallerSession) GetAssertion(id [32]byte) (Assertion, error) {
	return _AssertionChain.Contract.GetAssertion(&_AssertionChain.CallOpts, id)
}

// GetFirstChildCreationTime is a free data retrieval call binding the contract method 0x43ed6ad9.
//
// Solidity: function getFirstChildCreationTime(bytes32 assertionId) view returns(uint256)
func (_AssertionChain *AssertionChainCaller) GetFirstChildCreationTime(opts *bind.CallOpts, assertionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getFirstChildCreationTime", assertionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFirstChildCreationTime is a free data retrieval call binding the contract method 0x43ed6ad9.
//
// Solidity: function getFirstChildCreationTime(bytes32 assertionId) view returns(uint256)
func (_AssertionChain *AssertionChainSession) GetFirstChildCreationTime(assertionId [32]byte) (*big.Int, error) {
	return _AssertionChain.Contract.GetFirstChildCreationTime(&_AssertionChain.CallOpts, assertionId)
}

// GetFirstChildCreationTime is a free data retrieval call binding the contract method 0x43ed6ad9.
//
// Solidity: function getFirstChildCreationTime(bytes32 assertionId) view returns(uint256)
func (_AssertionChain *AssertionChainCallerSession) GetFirstChildCreationTime(assertionId [32]byte) (*big.Int, error) {
	return _AssertionChain.Contract.GetFirstChildCreationTime(&_AssertionChain.CallOpts, assertionId)
}

// GetHeight is a free data retrieval call binding the contract method 0x896efbf2.
//
// Solidity: function getHeight(bytes32 assertionId) view returns(uint256)
func (_AssertionChain *AssertionChainCaller) GetHeight(opts *bind.CallOpts, assertionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getHeight", assertionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHeight is a free data retrieval call binding the contract method 0x896efbf2.
//
// Solidity: function getHeight(bytes32 assertionId) view returns(uint256)
func (_AssertionChain *AssertionChainSession) GetHeight(assertionId [32]byte) (*big.Int, error) {
	return _AssertionChain.Contract.GetHeight(&_AssertionChain.CallOpts, assertionId)
}

// GetHeight is a free data retrieval call binding the contract method 0x896efbf2.
//
// Solidity: function getHeight(bytes32 assertionId) view returns(uint256)
func (_AssertionChain *AssertionChainCallerSession) GetHeight(assertionId [32]byte) (*big.Int, error) {
	return _AssertionChain.Contract.GetHeight(&_AssertionChain.CallOpts, assertionId)
}

// GetInboxMsgCountSeen is a free data retrieval call binding the contract method 0x7cfd5ab9.
//
// Solidity: function getInboxMsgCountSeen(bytes32 assertionId) view returns(uint256)
func (_AssertionChain *AssertionChainCaller) GetInboxMsgCountSeen(opts *bind.CallOpts, assertionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getInboxMsgCountSeen", assertionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInboxMsgCountSeen is a free data retrieval call binding the contract method 0x7cfd5ab9.
//
// Solidity: function getInboxMsgCountSeen(bytes32 assertionId) view returns(uint256)
func (_AssertionChain *AssertionChainSession) GetInboxMsgCountSeen(assertionId [32]byte) (*big.Int, error) {
	return _AssertionChain.Contract.GetInboxMsgCountSeen(&_AssertionChain.CallOpts, assertionId)
}

// GetInboxMsgCountSeen is a free data retrieval call binding the contract method 0x7cfd5ab9.
//
// Solidity: function getInboxMsgCountSeen(bytes32 assertionId) view returns(uint256)
func (_AssertionChain *AssertionChainCallerSession) GetInboxMsgCountSeen(assertionId [32]byte) (*big.Int, error) {
	return _AssertionChain.Contract.GetInboxMsgCountSeen(&_AssertionChain.CallOpts, assertionId)
}

// GetPredecessorId is a free data retrieval call binding the contract method 0x9ca565d4.
//
// Solidity: function getPredecessorId(bytes32 assertionId) view returns(bytes32)
func (_AssertionChain *AssertionChainCaller) GetPredecessorId(opts *bind.CallOpts, assertionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getPredecessorId", assertionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPredecessorId is a free data retrieval call binding the contract method 0x9ca565d4.
//
// Solidity: function getPredecessorId(bytes32 assertionId) view returns(bytes32)
func (_AssertionChain *AssertionChainSession) GetPredecessorId(assertionId [32]byte) ([32]byte, error) {
	return _AssertionChain.Contract.GetPredecessorId(&_AssertionChain.CallOpts, assertionId)
}

// GetPredecessorId is a free data retrieval call binding the contract method 0x9ca565d4.
//
// Solidity: function getPredecessorId(bytes32 assertionId) view returns(bytes32)
func (_AssertionChain *AssertionChainCallerSession) GetPredecessorId(assertionId [32]byte) ([32]byte, error) {
	return _AssertionChain.Contract.GetPredecessorId(&_AssertionChain.CallOpts, assertionId)
}

// GetStateHash is a free data retrieval call binding the contract method 0xf9bce634.
//
// Solidity: function getStateHash(bytes32 assertionId) view returns(bytes32)
func (_AssertionChain *AssertionChainCaller) GetStateHash(opts *bind.CallOpts, assertionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getStateHash", assertionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStateHash is a free data retrieval call binding the contract method 0xf9bce634.
//
// Solidity: function getStateHash(bytes32 assertionId) view returns(bytes32)
func (_AssertionChain *AssertionChainSession) GetStateHash(assertionId [32]byte) ([32]byte, error) {
	return _AssertionChain.Contract.GetStateHash(&_AssertionChain.CallOpts, assertionId)
}

// GetStateHash is a free data retrieval call binding the contract method 0xf9bce634.
//
// Solidity: function getStateHash(bytes32 assertionId) view returns(bytes32)
func (_AssertionChain *AssertionChainCallerSession) GetStateHash(assertionId [32]byte) ([32]byte, error) {
	return _AssertionChain.Contract.GetStateHash(&_AssertionChain.CallOpts, assertionId)
}

// GetSuccessionChallenge is a free data retrieval call binding the contract method 0x2fefa18a.
//
// Solidity: function getSuccessionChallenge(bytes32 assertionId) view returns(bytes32)
func (_AssertionChain *AssertionChainCaller) GetSuccessionChallenge(opts *bind.CallOpts, assertionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "getSuccessionChallenge", assertionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSuccessionChallenge is a free data retrieval call binding the contract method 0x2fefa18a.
//
// Solidity: function getSuccessionChallenge(bytes32 assertionId) view returns(bytes32)
func (_AssertionChain *AssertionChainSession) GetSuccessionChallenge(assertionId [32]byte) ([32]byte, error) {
	return _AssertionChain.Contract.GetSuccessionChallenge(&_AssertionChain.CallOpts, assertionId)
}

// GetSuccessionChallenge is a free data retrieval call binding the contract method 0x2fefa18a.
//
// Solidity: function getSuccessionChallenge(bytes32 assertionId) view returns(bytes32)
func (_AssertionChain *AssertionChainCallerSession) GetSuccessionChallenge(assertionId [32]byte) ([32]byte, error) {
	return _AssertionChain.Contract.GetSuccessionChallenge(&_AssertionChain.CallOpts, assertionId)
}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionId) view returns(bool)
func (_AssertionChain *AssertionChainCaller) IsFirstChild(opts *bind.CallOpts, assertionId [32]byte) (bool, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "isFirstChild", assertionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionId) view returns(bool)
func (_AssertionChain *AssertionChainSession) IsFirstChild(assertionId [32]byte) (bool, error) {
	return _AssertionChain.Contract.IsFirstChild(&_AssertionChain.CallOpts, assertionId)
}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionId) view returns(bool)
func (_AssertionChain *AssertionChainCallerSession) IsFirstChild(assertionId [32]byte) (bool, error) {
	return _AssertionChain.Contract.IsFirstChild(&_AssertionChain.CallOpts, assertionId)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_AssertionChain *AssertionChainCaller) StakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssertionChain.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_AssertionChain *AssertionChainSession) StakeAmount() (*big.Int, error) {
	return _AssertionChain.Contract.StakeAmount(&_AssertionChain.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_AssertionChain *AssertionChainCallerSession) StakeAmount() (*big.Int, error) {
	return _AssertionChain.Contract.StakeAmount(&_AssertionChain.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x5a627dbc.
//
// Solidity: function addStake() payable returns()
func (_AssertionChain *AssertionChainTransactor) AddStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "addStake")
}

// AddStake is a paid mutator transaction binding the contract method 0x5a627dbc.
//
// Solidity: function addStake() payable returns()
func (_AssertionChain *AssertionChainSession) AddStake() (*types.Transaction, error) {
	return _AssertionChain.Contract.AddStake(&_AssertionChain.TransactOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x5a627dbc.
//
// Solidity: function addStake() payable returns()
func (_AssertionChain *AssertionChainTransactorSession) AddStake() (*types.Transaction, error) {
	return _AssertionChain.Contract.AddStake(&_AssertionChain.TransactOpts)
}

// ConfirmAssertion is a paid mutator transaction binding the contract method 0x6894bdd5.
//
// Solidity: function confirmAssertion(bytes32 assertionId) returns()
func (_AssertionChain *AssertionChainTransactor) ConfirmAssertion(opts *bind.TransactOpts, assertionId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "confirmAssertion", assertionId)
}

// ConfirmAssertion is a paid mutator transaction binding the contract method 0x6894bdd5.
//
// Solidity: function confirmAssertion(bytes32 assertionId) returns()
func (_AssertionChain *AssertionChainSession) ConfirmAssertion(assertionId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmAssertion(&_AssertionChain.TransactOpts, assertionId)
}

// ConfirmAssertion is a paid mutator transaction binding the contract method 0x6894bdd5.
//
// Solidity: function confirmAssertion(bytes32 assertionId) returns()
func (_AssertionChain *AssertionChainTransactorSession) ConfirmAssertion(assertionId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.ConfirmAssertion(&_AssertionChain.TransactOpts, assertionId)
}

// CreateNewAssertion is a paid mutator transaction binding the contract method 0x49635f9a.
//
// Solidity: function createNewAssertion(bytes32 stateHash, uint256 height, bytes32 predecessorId) returns()
func (_AssertionChain *AssertionChainTransactor) CreateNewAssertion(opts *bind.TransactOpts, stateHash [32]byte, height *big.Int, predecessorId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "createNewAssertion", stateHash, height, predecessorId)
}

// CreateNewAssertion is a paid mutator transaction binding the contract method 0x49635f9a.
//
// Solidity: function createNewAssertion(bytes32 stateHash, uint256 height, bytes32 predecessorId) returns()
func (_AssertionChain *AssertionChainSession) CreateNewAssertion(stateHash [32]byte, height *big.Int, predecessorId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.CreateNewAssertion(&_AssertionChain.TransactOpts, stateHash, height, predecessorId)
}

// CreateNewAssertion is a paid mutator transaction binding the contract method 0x49635f9a.
//
// Solidity: function createNewAssertion(bytes32 stateHash, uint256 height, bytes32 predecessorId) returns()
func (_AssertionChain *AssertionChainTransactorSession) CreateNewAssertion(stateHash [32]byte, height *big.Int, predecessorId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.CreateNewAssertion(&_AssertionChain.TransactOpts, stateHash, height, predecessorId)
}

// CreateSuccessionChallenge is a paid mutator transaction binding the contract method 0xff8aef87.
//
// Solidity: function createSuccessionChallenge(bytes32 assertionId) returns()
func (_AssertionChain *AssertionChainTransactor) CreateSuccessionChallenge(opts *bind.TransactOpts, assertionId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "createSuccessionChallenge", assertionId)
}

// CreateSuccessionChallenge is a paid mutator transaction binding the contract method 0xff8aef87.
//
// Solidity: function createSuccessionChallenge(bytes32 assertionId) returns()
func (_AssertionChain *AssertionChainSession) CreateSuccessionChallenge(assertionId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.CreateSuccessionChallenge(&_AssertionChain.TransactOpts, assertionId)
}

// CreateSuccessionChallenge is a paid mutator transaction binding the contract method 0xff8aef87.
//
// Solidity: function createSuccessionChallenge(bytes32 assertionId) returns()
func (_AssertionChain *AssertionChainTransactorSession) CreateSuccessionChallenge(assertionId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.CreateSuccessionChallenge(&_AssertionChain.TransactOpts, assertionId)
}

// RejectAssertion is a paid mutator transaction binding the contract method 0x75dc6098.
//
// Solidity: function rejectAssertion(bytes32 assertionId) returns()
func (_AssertionChain *AssertionChainTransactor) RejectAssertion(opts *bind.TransactOpts, assertionId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "rejectAssertion", assertionId)
}

// RejectAssertion is a paid mutator transaction binding the contract method 0x75dc6098.
//
// Solidity: function rejectAssertion(bytes32 assertionId) returns()
func (_AssertionChain *AssertionChainSession) RejectAssertion(assertionId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.RejectAssertion(&_AssertionChain.TransactOpts, assertionId)
}

// RejectAssertion is a paid mutator transaction binding the contract method 0x75dc6098.
//
// Solidity: function rejectAssertion(bytes32 assertionId) returns()
func (_AssertionChain *AssertionChainTransactorSession) RejectAssertion(assertionId [32]byte) (*types.Transaction, error) {
	return _AssertionChain.Contract.RejectAssertion(&_AssertionChain.TransactOpts, assertionId)
}

// UpdateChallengeManager is a paid mutator transaction binding the contract method 0x295dfd32.
//
// Solidity: function updateChallengeManager(address _challengeManager) returns()
func (_AssertionChain *AssertionChainTransactor) UpdateChallengeManager(opts *bind.TransactOpts, _challengeManager common.Address) (*types.Transaction, error) {
	return _AssertionChain.contract.Transact(opts, "updateChallengeManager", _challengeManager)
}

// UpdateChallengeManager is a paid mutator transaction binding the contract method 0x295dfd32.
//
// Solidity: function updateChallengeManager(address _challengeManager) returns()
func (_AssertionChain *AssertionChainSession) UpdateChallengeManager(_challengeManager common.Address) (*types.Transaction, error) {
	return _AssertionChain.Contract.UpdateChallengeManager(&_AssertionChain.TransactOpts, _challengeManager)
}

// UpdateChallengeManager is a paid mutator transaction binding the contract method 0x295dfd32.
//
// Solidity: function updateChallengeManager(address _challengeManager) returns()
func (_AssertionChain *AssertionChainTransactorSession) UpdateChallengeManager(_challengeManager common.Address) (*types.Transaction, error) {
	return _AssertionChain.Contract.UpdateChallengeManager(&_AssertionChain.TransactOpts, _challengeManager)
}

// ChallengeManagerImplMetaData contains all meta data concerning the ChallengeManagerImpl contract.
var ChallengeManagerImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAssertionChain\",\"name\":\"_assertionChain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_miniStakeValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodSec\",\"type\":\"uint256\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"_oneStepProofEntry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fromId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"toId\",\"type\":\"bytes32\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fromId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"toId\",\"type\":\"bytes32\"}],\"name\":\"Merged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"VertexAdded\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"historyRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"firstStatehistoryProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"lastState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"lastStatehistoryProof\",\"type\":\"bytes\"}],\"internalType\":\"structAddLeafArgs\",\"name\":\"leafData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof2\",\"type\":\"bytes\"}],\"name\":\"addLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assertionChain\",\"outputs\":[{\"internalType\":\"contractIAssertionChain\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prefixHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"bisect\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"},{\"internalType\":\"enumChallengeType\",\"name\":\"typ\",\"type\":\"uint8\"}],\"name\":\"calculateChallengeId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"commitmentMerkle\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commitmentHeight\",\"type\":\"uint256\"}],\"name\":\"calculateChallengeVertexId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"challengeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengePeriodSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"rootId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"winningClaim\",\"type\":\"bytes32\"},{\"internalType\":\"enumChallengeType\",\"name\":\"challengeType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"childrenAreAtOneStepFork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"confirmForPsTimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"confirmForSucessionChallengeWin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"createSubChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"winnerVId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structOneStepData\",\"name\":\"oneStepData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"beforeHistoryInclusionProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"afterHistoryInclusionProof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"getChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"winningClaim\",\"type\":\"bytes32\"},{\"internalType\":\"enumChallengeType\",\"name\":\"challengeType\",\"type\":\"uint8\"}],\"internalType\":\"structChallenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentPsTimer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"getVertex\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"historyRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"successionChallenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"enumVertexStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"psId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"psLastUpdatedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flushedPsTimeSec\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lowestHeightSuccessorId\",\"type\":\"bytes32\"}],\"internalType\":\"structChallengeVertex\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"hasConfirmedSibling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAssertionChain\",\"name\":\"_assertionChain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_miniStakeValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriodSec\",\"type\":\"uint256\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"_oneStepProofEntry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"isPresumptiveSuccessor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prefixHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"merge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miniStakeValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"vertexExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"vertices\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"historyRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"successionChallenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"enumVertexStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"psId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"psLastUpdatedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flushedPsTimeSec\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lowestHeightSuccessorId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"winningClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162004b2138038062004b218339810160408190526200003491620000ec565b62000042848484846200004c565b505050506200013d565b6002546001600160a01b031615620000995760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015260640160405180910390fd5b600280546001600160a01b039586166001600160a01b03199182161790915560049390935560059190915560038054919093169116179055565b6001600160a01b0381168114620000e957600080fd5b50565b600080600080608085870312156200010357600080fd5b84516200011081620000d3565b80945050602085015192506040850151915060608501516200013281620000d3565b939692955090935050565b6149d4806200014d6000396000f3fe6080604052600436106101565760003560e01c806373d154e8116100c1578063b241493b1161007a578063b241493b146103fe578063bd62325114610411578063c1e69b6614610431578063d1bac9a41461047e578063e41b50581461049e578063f4f81db2146104be578063f696dc551461055b57600080fd5b806373d154e8146103205780637a4d47dc1461035157806386f048ed146103715780638ac043491461039e57806398b67d59146103be5780639e3d87cd146103de57600080fd5b80634a658788116101135780634a65878814610274578063597e1e0b1461029457806359c69996146102b4578063654f0dc2146102ca5780636b0b2592146102e05780636e49e3f21461030057600080fd5b806316ef55341461015b5780631b7bbecb1461018e5780631d5618ac146101cd578063359076cf146101ef578063458d2bf11461020f57806348dd29241461023c575b600080fd5b34801561016757600080fd5b5061017b610176366004613f62565b61057b565b6040519081526020015b60405180910390f35b34801561019a57600080fd5b506101bd6101a9366004613f96565b600090815260016020526040902054151590565b6040519015158152602001610185565b3480156101d957600080fd5b506101ed6101e8366004613f96565b610590565b005b3480156101fb57600080fd5b5061017b61020a36600461407a565b6105aa565b34801561021b57600080fd5b5061022f61022a366004613f96565b61067a565b60405161018591906140f3565b34801561024857600080fd5b5060025461025c906001600160a01b031681565b6040516001600160a01b039091168152602001610185565b34801561028057600080fd5b5061017b61028f36600461411b565b61075a565b3480156102a057600080fd5b5061017b6102af36600461407a565b61076f565b3480156102c057600080fd5b5061017b60045481565b3480156102d657600080fd5b5061017b60055481565b3480156102ec57600080fd5b506101bd6102fb366004613f96565b61080a565b34801561030c57600080fd5b5061017b61031b36600461418f565b610823565b34801561032c57600080fd5b5061017b61033b366004613f96565b6000908152600160208190526040909120015490565b34801561035d57600080fd5b506101bd61036c366004613f96565b610868565b34801561037d57600080fd5b5061039161038c366004613f96565b61087c565b6040516101859190614244565b3480156103aa57600080fd5b5061017b6103b9366004613f96565b61097e565b3480156103ca57600080fd5b506101bd6103d9366004613f96565b61098a565b3480156103ea57600080fd5b506101ed6103f93660046142f9565b610b59565b61017b61040c366004614343565b610bdb565b34801561041d57600080fd5b5061017b61042c366004613f96565b610f05565b34801561043d57600080fd5b5061046f61044c366004613f96565b600160208190526000918252604090912080549181015460029091015460ff1683565b604051610185939291906143e0565b34801561048a57600080fd5b506101ed610499366004613f96565b6110ad565b3480156104aa57600080fd5b506101bd6104b9366004613f96565b6110ba565b3480156104ca57600080fd5b506105436104d9366004613f96565b600060208190529081526040902080546001820154600283015460038401546004850154600586015460068701546007880154600889015460098a0154600a909a01549899979896979596949593946001600160a01b03841694600160a01b90940460ff1693908c565b6040516101859c9b9a999897969594939291906143fb565b34801561056757600080fd5b5061017b610576366004613f96565b61113a565b60006105878383611323565b90505b92915050565b61059e600082600554611356565b6105a781611443565b50565b60008060006105c16000600188888860055461149e565b6000888152602081905260408120600401549294509092506105e38189611525565b6000898152602081905260408120549192509061060290898685611637565b905061061e81846005546000611703909392919063ffffffff16565b506005546106329060009087908c9061183f565b604080518a8152602081018790527f69d5465c81edf7aaaf2e5c6c8829500df87d84c87f8d5b1221b59eaeaca70d27910160405180910390a1509293505050505b9392505050565b61069b60408051606081018252600080825260208201819052909182015290565b6000828152600160205260409020546106f65760405162461bcd60e51b815260206004820152601860248201527710da185b1b195b99d948191bd95cc81b9bdd08195e1a5cdd60421b60448201526064015b60405180910390fd5b60008281526001602081815260409283902083516060810185528154815292810154918301919091526002810154919290919083019060ff166003811115610740576107406140c9565b6003811115610751576107516140c9565b90525092915050565b6000610767848484611ce3565b949350505050565b60008061078460006001878787600554611d1a565b5090506107a18186600554600061183f909392919063ffffffff16565b60008181526020819052604080822060040154878352908220600901546107c9929190611dfe565b60408051868152602081018390527f72b50597145599e4288d411331c925b40b33b0fa3cccadc1f57d2a1ab973553a910160405180910390a1949350505050565b600081815260208190526040812060010154151561058a565b60035460009081906108489082906001906001600160a01b03168b8b8b8b8b8b611f4b565b600090815260016020819052604090912001979097559695505050505050565b60006108748183612267565b506001919050565b610884613f01565b6000828152602081905260409020600101546108b25760405162461bcd60e51b81526004016106ed9061446f565b6000828152602081815260409182902082516101808101845281548152600180830154938201939093526002820154938101939093526003810154606084015260048101546080840152600581015460a084015260068101546001600160a01b03811660c0850152909160e0840191600160a01b900460ff169081111561093b5761093b6140c9565b600181111561094c5761094c6140c9565b8152600782015460208201526008820154604082015260098201546060820152600a9091015460809091015292915050565b600061058a8183611525565b6000818152602081905260408120600101546109b85760405162461bcd60e51b81526004016106ed9061446f565b600082815260208190526040808220600401548083529120600101546109f05760405162461bcd60e51b81526004016106ed9061449e565b6000818152602081905260409020600301548015610acc57600081815260016020819052604090912001548015610aca57600081815260208190526040902060010154610a7f5760405162461bcd60e51b815260206004820152601c60248201527f57696e6e696e6720636c61696d20646f6573206e6f742065786973740000000060448201526064016106ed565b848103610a9157506000949350505050565b6001600082815260208190526040902060060154600160a01b900460ff166001811115610ac057610ac06140c9565b1495945050505050565b505b6000828152602081905260409020600701548015610b4e57600081815260208190526040902060010154610a7f5760405162461bcd60e51b8152602060048201526024808201527f50726573756d707469766520737563636573736f7220646f6573206e6f7420656044820152631e1a5cdd60e21b60648201526084016106ed565b506000949350505050565b6002546001600160a01b031615610ba15760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b60448201526064016106ed565b600280546001600160a01b039586166001600160a01b03199182161790915560049390935560059190915560038054919093169116179055565b600080863560009081526001602052604090206002015460ff166003811115610c0657610c066140c9565b03610d04576000610cc2600060016040518060a00160405280600454815260200160055481526020018b610c39906144d5565b81526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8b0181900481028201810190925289815291810191908a908a908190840183828082843760009201919091525050509152506002546001600160a01b0316612487565b90507f4383ba11a7cd16be5880c5f674b93be38b3b1fcafd7a7b06151998fa2a67534981604051610cf591815260200190565b60405180910390a19050610efc565b6001863560009081526001602052604090206002015460ff166003811115610d2e57610d2e6140c9565b03610dde576000610cc2600060016040518060a00160405280600454815260200160055481526020018b610d61906144d5565b81526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8b0181900481028201810190925289815291810191908a908a90819084018382808284376000920191909152505050915250612948565b6002863560009081526001602052604090206002015460ff166003811115610e0857610e086140c9565b03610eb8576000610cc2600060016040518060a00160405280600454815260200160055481526020018b610e3b906144d5565b81526020018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8b0181900481028201810190925289815291810191908a908a90819084018382808284376000920191909152505050915250612ab4565b60405162461bcd60e51b8152602060048201526019602482015278556e6578706563746564206368616c6c656e6765207479706560381b60448201526064016106ed565b95945050505050565b6000806000610f1a6000600186600554612c63565b600086815260208190526040812060010154929450909250610f3d848383611ce3565b9050610f4a848388612db7565b600082815260208181526040918290208351815590830151600180830191909155918301516002820155606083015160038201556080830151600482015560a0830151600582015560c08301516006820180546001600160a01b039092166001600160a01b031983168117825560e0860151939491926001600160a81b0319161790600160a01b908490811115610fe357610fe36140c9565b021790555061010082015160078201556101208201516008820155610140820151600982015561016090910151600a90910155604080516060810182528281526000602082015290810184600381111561103f5761103f6140c9565b9052600085815260016020818152604092839020845181559084015181830155918301516002830180549192909160ff191690836003811115611084576110846140c9565b0217905550505060008681526020819052604090206110a39085612e62565b5091949350505050565b61059e6000600183612eb1565b6000818152602081905260408120600101546110e85760405162461bcd60e51b81526004016106ed9061446f565b600082815260208190526040808220600401548083529120600101546111205760405162461bcd60e51b81526004016106ed9061457e565b600090815260208190526040902060070154909114919050565b60025460009081906111599060019085906001600160a01b0316612f7c565b600254604051633e6f398d60e21b8152600481018690529192506000916001600160a01b039091169063f9bce63490602401602060405180830381865afa1580156111a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111cc91906145bf565b905060006111dc83836000611ce3565b90506111e9838387612db7565b600082815260208181526040918290208351815590830151600180830191909155918301516002820155606083015160038201556080830151600482015560a0830151600582015560c08301516006820180546001600160a01b039092166001600160a01b031983168117825560e0860151939491926001600160a81b0319161790600160a01b908490811115611282576112826140c9565b021790555061010082015160078201556101208201516008820155610140820151600982015561016090910151600a90910155604080516060810182528281526000602082018190529091820152600084815260016020818152604092839020845181559084015181830155918301516002830180549192909160ff191690836003811115611313576113136140c9565b0217905550939695505050505050565b600082826040516020016113389291906145d8565b60405160208183030381529060405280519060200120905092915050565b6113608383613025565b600082815260208490526040808220600401548252902060030154156113d45760405162461bcd60e51b815260206004820152602360248201527f53756363657373696f6e206368616c6c656e676520616c7265616479206f70656044820152621b995960ea1b60648201526084016106ed565b806113df8484611525565b1161143e5760405162461bcd60e51b815260206004820152602960248201527f507354696d6572206e6f742067726561746572207468616e206368616c6c656e60448201526819d9481c195c9a5bd960ba1b60648201526084016106ed565b505050565b600081815260208190526040902061145a90613176565b600081815260208190526040902080549061147490613236565b1561149a5760008281526020818152604080832060050154848452600192839052922001555b5050565b6000806000806114b28a8a8a8a8a8a6132be565b600082815260208d905260409020600101549193509150156115165760405162461bcd60e51b815260206004820152601f60248201527f426973656374696f6e2076657274657820616c7265616479206578697374730060448201526064016106ed565b90999098509650505050505050565b60008181526020839052604081206001015461158e5760405162461bcd60e51b815260206004820152602260248201527f56657274657820646f6573206e6f7420657869737420666f722070732074696d60448201526132b960f11b60648201526084016106ed565b600082815260208490526040808220600401548083529120600101546115c65760405162461bcd60e51b81526004016106ed9061457e565b60008181526020859052604090206007015483900361161757600083815260208590526040808220600901548383529120600801546116059042614617565b61160f919061462a565b91505061058a565b505060008181526020839052604090206009015461058a565b5092915050565b61163f613f01565b60008590036116605760405162461bcd60e51b81526004016106ed9061463d565b60008490036116815760405162461bcd60e51b81526004016106ed90614668565b826000036116a15760405162461bcd60e51b81526004016106ed90614693565b5060408051610180810182529485526020850193909352918301526000606083018190526080830181905260a0830181905260c0830181905260e083018190526101008301819052610120830181905261014083019190915261016082015290565b60008061170f85613401565b600081815260208890526040902060010154909150156117695760405162461bcd60e51b815260206004820152601560248201527456657274657820616c72656164792065786973747360581b60448201526064016106ed565b600081815260208781526040918290208751815590870151600180830191909155918701516002820155606087015160038201556080870151600482015560a0870151600582015560c08701516006820180546001600160a01b039092166001600160a01b031983168117825560e08a01518a9590936001600160a81b03191690911790600160a01b908490811115611804576118046140c9565b021790555061010082015160078201556101208201516008820155610140820151600982015561016090910151600a90910155610efc868583865b60008381526020859052604090206001015461189d5760405162461bcd60e51b815260206004820152601b60248201527f53746172742076657274657820646f6573206e6f74206578697374000000000060448201526064016106ed565b60008381526020859052604090206118b490613236565b1561190d5760405162461bcd60e51b8152602060048201526024808201527f43616e6e6f7420636f6e6e656374206120737563636573736f7220746f2061206044820152633632b0b360e11b60648201526084016106ed565b6000828152602085905260409020600101546119675760405162461bcd60e51b8152602060048201526019602482015278115b99081d995c9d195e08191bd95cc81b9bdd08195e1a5cdd603a1b60448201526064016106ed565b6000828152602085905260409020600401548390036119c85760405162461bcd60e51b815260206004820152601a60248201527f566572746963657320616c726561647920636f6e6e656374656400000000000060448201526064016106ed565b600082815260208590526040808220600290810154868452919092209091015410611a445760405162461bcd60e51b815260206004820152602660248201527f537461727420686569676874206e6f74206c6f776572207468616e20656e64206044820152651a195a59da1d60d21b60648201526084016106ed565b6000828152602085905260408082205485835291205414611ac55760405162461bcd60e51b815260206004820152603560248201527f5072656465636573736f7220616e6420737563636573736f722061726520696e60448201527420646966666572656e74206368616c6c656e67657360581b60648201526084016106ed565b6000828152602085905260409020611add908461341a565b6000838152602085905260408120600a01549003611b1e57611b0184846000611dfe565b6000838152602085905260409020611b1990836134e6565b611cdd565b600082815260208590526040808220600290810154868452828420600a01548452919092209091015480821015611c1257611b5a8686856135c1565b15611be75760405162461bcd60e51b815260206004820152605160248201527f5374617274207665727465782068617320707320776974682074696d6572206760448201527f726561746572207468616e206368616c6c656e676520706572696f642c2063616064820152706e6e6f7420736574206c6f77657220707360781b608482015260a4016106ed565b611bf386866000611dfe565b6000858152602087905260409020611c0b90856134e6565b5050611cdd565b808203611cda57611c248686856135c1565b15611cb75760405162461bcd60e51b815260206004820152605760248201527f5374617274207665727465782068617320707320776974682074696d6572206760448201527f726561746572207468616e206368616c6c656e676520706572696f642c2063616064820152766e6e6f74207365742073616d652068656967687420707360481b608482015260a4016106ed565b611cc386866000611dfe565b6000858152602087905260408120611c0b916134e6565b50505b50505050565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b600080600080611d2e8a8a8a8a8a8a6132be565b600082815260208d905260409020600101549193509150611da15760405162461bcd60e51b815260206004820152602760248201527f426973656374696f6e2076657274657820646f6573206e6f7420616c726561646044820152661e48195e1a5cdd60ca1b60648201526084016106ed565b600082815260208b905260409020611db890613236565b156115165760405162461bcd60e51b815260206004820152601660248201527521b0b73737ba1036b2b933b2903a379030903632b0b360511b60448201526064016106ed565b600082815260208490526040902060010154611e2c5760405162461bcd60e51b81526004016106ed9061446f565b6000828152602084905260409020611e4390613236565b15611ea55760405162461bcd60e51b815260206004820152602c60248201527f43616e6e6f7420666c757368206c6561662061732069742077696c6c206e657660448201526b65722068617665206120505360a01b60648201526084016106ed565b60008281526020849052604090206007015415611f3357600082815260208490526040812060080154611ed89042614617565b60008481526020869052604080822060070154825281206009015491925090611f0290839061462a565b905082811015611f0f5750815b6000848152602086905260408082206007015482529020611f309082613638565b50505b600082815260208490526040902061143e90426136c3565b600086815260208a90526040812060010154611f795760405162461bcd60e51b81526004016106ed9061446f565b600087815260208b9052604080822060040154808352912060010154611fb15760405162461bcd60e51b81526004016106ed9061449e565b600081815260208c9052604081206003015490819003611fe35760405162461bcd60e51b81526004016106ed906146b8565b6003600082815260208d9052604090206002015460ff16600381111561200b5761200b6140c9565b1461206d5760405162461bcd60e51b815260206004820152602c60248201527f4368616c6c656e6765206973206e6f74206174206f6e6520737465702065786560448201526b18dd5d1a5bdb881c1bda5b9d60a21b60648201526084016106ed565b6120cc8c60008481526020019081526020016000206001015489606001358a604001358a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061371292505050565b6121185760405162461bcd60e51b815260206004820152601b60248201527f4265666f7265207374617465206e6f7420696e20686973746f7279000000000060448201526064016106ed565b60006001600160a01b038b16635d3adcfb8a6040810135606082013561214160808401846146fb565b6040518663ffffffff1660e01b8152600401612161959493929190614741565b602060405180830381865afa15801561217e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121a291906145bf565b905061220b8d60008c815260200190815260200160002060010154828b6040013560016121cf919061462a565b89898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061371292505050565b6122575760405162461bcd60e51b815260206004820152601a60248201527f4166746572207374617465206e6f7420696e20686973746f727900000000000060448201526064016106ed565b509b9a5050505050505050505050565b6000818152602083905260409020600101546122d15760405162461bcd60e51b8152602060048201526024808201527f466f726b2063616e6469646174652076657274657820646f6573206e6f7420656044820152631e1a5cdd60e21b60648201526084016106ed565b60008181526020839052604090206122e890613236565b156123405760405162461bcd60e51b815260206004820152602260248201527f4c6561662063616e206e65766572206265206120666f726b2063616e64696461604482015261746560f01b60648201526084016106ed565b600081815260208390526040808220600a015482529020600101546123975760405162461bcd60e51b815260206004820152600d60248201526c4e6f20737563636573736f727360981b60448201526064016106ed565b600081815260208390526040808220600a81015483529082206002908101549284905201546123c69082614617565b60011461242c5760405162461bcd60e51b815260206004820152602e60248201527f4c6f7765737420686569676874206e6f74206f6e652061626f7665207468652060448201526d18dd5c9c995b9d081a195a59da1d60921b60648201526084016106ed565b6000828152602084905260409020600701541561143e5760405162461bcd60e51b81526020600482015260196024820152782430b990383932b9bab6b83a34bb329039bab1b1b2b9b9b7b960391b60448201526064016106ed565b600080826001600160a01b0316639ca565d48560400151602001516040518263ffffffff1660e01b81526004016124c091815260200190565b602060405180830381865afa1580156124dd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061250191906145bf565b6040858101515190516317f7d0c560e11b815260048101839052919250906001600160a01b03851690632fefa18a90602401602060405180830381865afa158015612550573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061257491906145bf565b146125d85760405162461bcd60e51b815260206004820152602e60248201527f436c61696d207072656465636573736f72206e6f74206c696e6b656420746f2060448201526d74686973206368616c6c656e676560901b60648201526084016106ed565b6040808501516020015190516344b77df960e11b81526000916001600160a01b0386169163896efbf2916126129160040190815260200190565b602060405180830381865afa15801561262f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061265391906145bf565b6040516344b77df960e11b8152600481018490529091506000906001600160a01b0386169063896efbf290602401602060405180830381865afa15801561269e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126c291906145bf565b905060006126d08284614617565b9050866040015160400151811461271a5760405162461bcd60e51b815260206004820152600e60248201526d125b9d985b1a59081a195a59da1d60921b60448201526064016106ed565b604080880151602001519051633e6f398d60e21b81526000916001600160a01b0389169163f9bce634916127549160040190815260200190565b602060405180830381865afa158015612771573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061279591906145bf565b604051637cfd5ab960e01b8152600481018790529091506001600160a01b03881690637cfd5ab990602401602060405180830381865afa1580156127dd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061280191906145bf565b61280f828a6080015161371c565b1461285c5760405162461bcd60e51b815260206004820181905260248201527f496e76616c696420696e626f78206d657373616765732070726f63657373656460448201526064016106ed565b876040015160c0015181146128cc5760405162461bcd60e51b815260206004820152603060248201527f4c617374207374617465206973206e6f742074686520617373657274696f6e2060448201526f0c6d8c2d2da40c4d8dec6d640d0c2e6d60831b60648201526084016106ed565b6128df8989604001518a60000151613727565b505050506040848101518051606082015192820151602090920151600094506129159391929033612910828a6138b2565b613a23565b6040808601515160009081526020888152919020549086015191925061293e9188918491611703565b9695505050505050565b604080820151602090810151600090815290859052908120600101546129805760405162461bcd60e51b81526004016106ed906147a8565b60408281015160209081015160009081529086905281812060040154808252919020600101546129c25760405162461bcd60e51b81526004016106ed906147d6565b600081815260208681526040808320600290810154878301519093015184529220909101546129f19190614617565b600114612a105760405162461bcd60e51b81526004016106ed9061480b565b6040808401515160008381526020889052919091206003015414612a465760405162461bcd60e51b81526004016106ed90614851565b6000805b612a5d8686604001518760000151613727565b5050506040808301518051606082015192820151602090920151600093612a8b939091336129108b83611525565b60408085015151600090815260208781529190205490850151919250610efc9187918491611703565b60408082015160209081015160009081529085905290812060010154612aec5760405162461bcd60e51b81526004016106ed906147a8565b6040828101516020908101516000908152908690528181206004015480825291902060010154612b2e5760405162461bcd60e51b81526004016106ed906147d6565b60008181526020868152604080832060029081015487830151909301518452922090910154612b5d9190614617565b600114612b7c5760405162461bcd60e51b81526004016106ed9061480b565b6040808401515160008381526020889052919091206003015414612bb25760405162461bcd60e51b81526004016106ed90614851565b60408301516020908101516000528590526000612bdb846040015160c00151856080015161371c565b60008381526020889052604081206002015491925090612bff906210000090614897565b90508185604001516040015182612c16919061462a565b14612a4a5760405162461bcd60e51b815260206004820152601c60248201527f496e636f6e73697374656e742070726f6772616d20636f756e7465720000000060448201526064016106ed565b600080612c708685612267565b60008481526020869052604090206001015415612c9f5760405162461bcd60e51b81526004016106ed906148ae565b612caa8685856135c1565b15612d015760405162461bcd60e51b815260206004820152602160248201527f50726573756d707469766520737563636573736f7220636f6e6669726d61626c6044820152606560f81b60648201526084016106ed565b60008481526020879052604090206003015415612d305760405162461bcd60e51b81526004016106ed906148df565b6000848152602087815260408083205480845291889052822060020154909190612d6d9060ff166003811115612d6857612d686140c9565b613b89565b90506000612d7b8783611323565b600081815260208a9052604090205490915015612daa5760405162461bcd60e51b81526004016106ed906148df565b9890975095505050505050565b612dbf613f01565b6000849003612de05760405162461bcd60e51b81526004016106ed9061463d565b6000839003612e015760405162461bcd60e51b81526004016106ed90614668565b50604080516101808101825293845260208401929092526000918301829052606083018290526080830182905260a083015260c08201819052600160e083015261010082018190526101208201819052610140820181905261016082015290565b6001820154612e835760405162461bcd60e51b81526004016106ed9061446f565b612e8c82613236565b15612ea95760405162461bcd60e51b81526004016106ed90614911565b600390910155565b612ebb8382613025565b60008181526020849052604080822060040154825281206003015490819003612ef65760405162461bcd60e51b81526004016106ed906146b8565b6000818152602084905260409020600101548214611cdd5760405162461bcd60e51b815260206004820152603b60248201527f53756363657373696f6e206368616c6c656e676520646964206e6f742064656360448201527f6c617265207468697320766572746578207468652077696e6e6572000000000060648201526084016106ed565b6000336001600160a01b03831614612fe95760405162461bcd60e51b815260206004820152602a60248201527f4f6e6c7920617373657274696f6e20636861696e2063616e20637265617465206044820152696368616c6c656e67657360b01b60648201526084016106ed565b6000612ff6846000611323565b600081815260208790526040902054909150156107675760405162461bcd60e51b81526004016106ed906148df565b6000818152602083905260409020600101546130535760405162461bcd60e51b81526004016106ed9061446f565b60008082815260208490526040902060060154600160a01b900460ff166001811115613081576130816140c9565b146130c65760405162461bcd60e51b8152602060048201526015602482015274566572746578206973206e6f742070656e64696e6760581b60448201526064016106ed565b600081815260208390526040808220600401548083529120600101546130fe5760405162461bcd60e51b81526004016106ed9061457e565b6001600082815260208590526040902060060154600160a01b900460ff16600181111561312d5761312d6140c9565b1461143e5760405162461bcd60e51b8152602060048201526019602482015278141c99591958d95cdcdbdc881b9bdd0818dbdb999a5c9b5959603a1b60448201526064016106ed565b60018101546131975760405162461bcd60e51b81526004016106ed9061446f565b60006006820154600160a01b900460ff1660018111156131b9576131b96140c9565b146132205760405162461bcd60e51b815260206004820152603160248201527f566572746578206d7573742062652050656e64696e67206265666f72652062656044820152701a5b99c81cd95d0810dbdb999a5c9b5959607a1b60648201526084016106ed565b600601805460ff60a01b1916600160a01b179055565b60006132458260010154151590565b61329d5760405162461bcd60e51b8152602060048201526024808201527f506f74656e7469616c206c6561662076657274657820646f6573206e6f7420656044820152631e1a5cdd60e21b60648201526084016106ed565b60018201541515801561058a575050600601546001600160a01b0316151590565b60008481526020879052604081206001015481906132ee5760405162461bcd60e51b81526004016106ed9061446f565b60008681526020898152604080832054808452918a905290912060010154156133295760405162461bcd60e51b81526004016106ed906148ae565b600087815260208a90526040808220600401548083529120600101546133615760405162461bcd60e51b81526004016106ed9061457e565b600081815260208b905260409020600701548890036133ce5760405162461bcd60e51b815260206004820152602360248201527f43616e6e6f74206269736563742070726573756d70746976652073756363657360448201526239b7b960e91b60648201526084016106ed565b60006133da8b8a613c58565b60008a905260208c905290506133f1838983611ce3565b9b909a5098505050505050505050565b600061058a826000015183602001518460400151611ce3565b600182015461343b5760405162461bcd60e51b81526004016106ed9061446f565b808260040154036134885760405162461bcd60e51b8152602060048201526017602482015276141c99591958d95cdcdbdc88185b1c9958591e481cd95d604a1b60448201526064016106ed565b61349182613ce8565b156134de5760405162461bcd60e51b815260206004820152601e60248201527f43616e6e6f7420736574207072656465636573736f72206f6e20726f6f74000060448201526064016106ed565b600490910155565b60018201546135075760405162461bcd60e51b81526004016106ed9061446f565b801580613518575080826007015414155b6135555760405162461bcd60e51b815260206004820152600e60248201526d141cc8185b1c9958591e481cd95d60921b60448201526064016106ed565b61355e82613236565b156135ab5760405162461bcd60e51b815260206004820152601a60248201527f43616e6e6f7420736574207073206964206f6e2061206c65616600000000000060448201526064016106ed565b60078201819055801561149a57600a9190910155565b6000828152602084905260408120600101546135ef5760405162461bcd60e51b81526004016106ed9061457e565b600083815260208590526040812060070154900361360f57506000610673565b8161362f8586600087815260200190815260200160002060070154611525565b11949350505050565b60018201546136595760405162461bcd60e51b81526004016106ed9061446f565b61366282613ce8565b156136bb5760405162461bcd60e51b8152602060048201526024808201527f43616e6e6f742073657420707320666c75736865642074696d65206f6e2061206044820152631c9bdbdd60e21b60648201526084016106ed565b600990910155565b60018201546136e45760405162461bcd60e51b81526004016106ed9061446f565b6136ed82613236565b1561370a5760405162461bcd60e51b81526004016106ed90614911565b600890910155565b6001949350505050565b600061058782614955565b606082015160000361376f5760405162461bcd60e51b8152602060048201526011602482015270115b5c1d1e481a1a5cdd1bdc9e549bdbdd607a1b60448201526064016106ed565b81604001516000036137b25760405162461bcd60e51b815260206004820152600c60248201526b115b5c1d1e481a195a59da1d60a21b60448201526064016106ed565b8034146138015760405162461bcd60e51b815260206004820152601b60248201527f496e636f7272656374206d696e692d7374616b6520616d6f756e74000000000060448201526064016106ed565b8151600090815260208490526040902060010154156138325760405162461bcd60e51b81526004016106ed906148ae565b613846826000015183608001516000611ce3565b82516000908152602085905260409020541461143e5760405162461bcd60e51b815260206004820152602560248201527f4669727374207374617465206973206e6f7420746865206368616c6c656e6765604482015264081c9bdbdd60da1b60648201526084016106ed565b6040516306106c4560e31b81526004810183905260009081906001600160a01b03841690633083622890602401602060405180830381865afa1580156138fc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613920919061497c565b90508015613a1957604051632729597560e21b8152600481018590526000906001600160a01b03851690639ca565d490602401602060405180830381865afa158015613970573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061399491906145bf565b6040516343ed6ad960e01b8152600481018290529091506000906001600160a01b038616906343ed6ad990602401602060405180830381865afa1580156139df573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a0391906145bf565b9050613a0f8142614617565b935050505061058a565b600091505061058a565b613a2b613f01565b6000879003613a4c5760405162461bcd60e51b81526004016106ed9061463d565b6000869003613a6d5760405162461bcd60e51b81526004016106ed90614668565b84600003613a8d5760405162461bcd60e51b81526004016106ed90614693565b6000849003613ace5760405162461bcd60e51b815260206004820152600d60248201526c16995c9bc818db185a5b481a59609a1b60448201526064016106ed565b6001600160a01b038316613b1a5760405162461bcd60e51b81526020600482015260136024820152725a65726f207374616b6572206164647265737360681b60448201526064016106ed565b5060408051610180810182529687526020870195909552938501929092526000606085018190526080850181905260a08501919091526001600160a01b0390911660c084015260e083018190526101008301819052610120830181905261014083019190915261016082015290565b600080826003811115613b9e57613b9e6140c9565b03613bab57506001919050565b6001826003811115613bbf57613bbf6140c9565b03613bcc57506002919050565b6002826003811115613be057613be06140c9565b03613bed57506003919050565b60405162461bcd60e51b815260206004820152603560248201527f43616e6e6f7420676574206e657874206368616c6c656e6765207479706520666044820152746f72206f6e652073746570206368616c6c656e676560581b60648201526084016106ed565b919050565b600081815260208390526040812060010154613c865760405162461bcd60e51b81526004016106ed9061446f565b60008281526020849052604080822060040154808352912060010154613cbe5760405162461bcd60e51b81526004016106ed9061457e565b60008181526020859052604080822060029081015486845291909220909101546107679190613d6f565b6000613cf78260010154151590565b613d4f5760405162461bcd60e51b8152602060048201526024808201527f506f74656e7469616c20726f6f742076657274657820646f6573206e6f7420656044820152631e1a5cdd60e21b60648201526084016106ed565b60068201546001600160a01b031615801561058a57505060050154151590565b60006002613d7d8484614617565b1015613dcb5760405162461bcd60e51b815260206004820181905260248201527f48656967687420646966666572656e74206e6f742074776f206f72206d6f726560448201526064016106ed565b613dd58383614617565b600203613dee57613de783600161462a565b905061058a565b6000613e0584613dff600186614617565b18613e22565b9050600019811b80613e18600186614617565b1695945050505050565b6000600160801b8210613e4257608091821c91613e3f908261462a565b90505b600160401b8210613e6057604091821c91613e5d908261462a565b90505b6401000000008210613e7f57602091821c91613e7c908261462a565b90505b620100008210613e9c57601091821c91613e99908261462a565b90505b6101008210613eb857600891821c91613eb5908261462a565b90505b60108210613ed357600491821c91613ed0908261462a565b90505b60048210613eee57600291821c91613eeb908261462a565b90505b60028210613c535761058a60018261462a565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081018290529060e08201908152600060208201819052604082018190526060820181905260809091015290565b60008060408385031215613f7557600080fd5b82359150602083013560048110613f8b57600080fd5b809150509250929050565b600060208284031215613fa857600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60405161010081016001600160401b0381118282101715613fe857613fe8613faf565b60405290565b600082601f830112613fff57600080fd5b81356001600160401b038082111561401957614019613faf565b604051601f8301601f19908116603f0116810190828211818310171561404157614041613faf565b8160405283815286602085880101111561405a57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561408f57600080fd5b833592506020840135915060408401356001600160401b038111156140b357600080fd5b6140bf86828701613fee565b9150509250925092565b634e487b7160e01b600052602160045260246000fd5b600481106140ef576140ef6140c9565b9052565b60006060820190508251825260208301516020830152604083015161163060408401826140df565b60008060006060848603121561413057600080fd5b505081359360208301359350604090920135919050565b60008083601f84011261415957600080fd5b5081356001600160401b0381111561417057600080fd5b60208301915083602082850101111561418857600080fd5b9250929050565b600080600080600080608087890312156141a857600080fd5b8635955060208701356001600160401b03808211156141c657600080fd5b9088019060a0828b0312156141da57600080fd5b909550604088013590808211156141f057600080fd5b6141fc8a838b01614147565b9096509450606089013591508082111561421557600080fd5b5061422289828a01614147565b979a9699509497509295939492505050565b600281106140ef576140ef6140c9565b600061018082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015161429d60c08401826001600160a01b03169052565b5060e08301516142b060e0840182614234565b5061010083810151908301526101208084015190830152610140808401519083015261016092830151929091019190915290565b6001600160a01b03811681146105a757600080fd5b6000806000806080858703121561430f57600080fd5b843561431a816142e4565b935060208501359250604085013591506060850135614338816142e4565b939692955090935050565b60008060008060006060868803121561435b57600080fd5b85356001600160401b038082111561437257600080fd5b90870190610100828a03121561438757600080fd5b9095506020870135908082111561439d57600080fd5b6143a989838a01614147565b909650945060408801359150808211156143c257600080fd5b506143cf88828901614147565b969995985093965092949392505050565b838152602081018390526060810161076760408301846140df565b8c8152602081018c9052604081018b9052606081018a90526080810189905260a081018890526001600160a01b03871660c0820152610180810161444260e0830188614234565b856101008301528461012083015283610140830152826101608301529d9c50505050505050505050505050565b60208082526015908201527415995c9d195e08191bd95cc81b9bdd08195e1a5cdd605a1b604082015260600190565b6020808252601a908201527f5072656465636573736f7220646f6573206e6f74206578697374000000000000604082015260600190565b600061010082360312156144e857600080fd5b6144f0613fc5565b823581526020830135602082015260408301356040820152606083013560608201526080830135608082015260a08301356001600160401b038082111561453657600080fd5b61454236838701613fee565b60a084015260c085013560c084015260e085013591508082111561456557600080fd5b5061457236828601613fee565b60e08301525092915050565b60208082526021908201527f5072656465636573736f722076657274657820646f6573206e6f7420657869736040820152601d60fa1b606082015260800190565b6000602082840312156145d157600080fd5b5051919050565b8281526000600483106145ed576145ed6140c9565b5060f89190911b6020820152602101919050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561058a5761058a614601565b8082018082111561058a5761058a614601565b60208082526011908201527016995c9bc818da185b1b195b99d9481a59607a1b604082015260600190565b60208082526011908201527016995c9bc81a1a5cdd1bdc9e481c9bdbdd607a1b604082015260600190565b6020808252600b908201526a16995c9bc81a195a59da1d60aa1b604082015260600190565b60208082526023908201527f53756363657373696f6e206368616c6c656e676520646f6573206e6f742065786040820152621a5cdd60ea1b606082015260800190565b6000808335601e1984360301811261471257600080fd5b8301803591506001600160401b0382111561472c57600080fd5b60200191503681900382131561418857600080fd5b8535815260006020870135614755816142e4565b6001600160a01b03166020830152604082018690526060820185905260a0608083018190528201839052828460c0840137600060c0848401015260c0601f19601f85011683010190509695505050505050565b60208082526014908201527310db185a5b48191bd95cc81b9bdd08195e1a5cdd60621b604082015260600190565b6020808252818101527f436c61696d207072656465636573736f7220646f6573206e6f74206578697374604082015260600190565b60208082526026908201527f436c61696d206e6f7420686569676874206f6e652061626f766520707265646560408201526531b2b9b9b7b960d11b606082015260800190565b60208082526026908201527f436c61696d2068617320696e76616c69642073756363657373696f6e206368616040820152656c6c656e676560d01b606082015260800190565b808202811582820484141761058a5761058a614601565b60208082526017908201527615da5b9b995c88185b1c9958591e48191958db185c9959604a1b604082015260600190565b6020808252601890820152774368616c6c656e676520616c72656164792065786973747360401b604082015260600190565b60208082526024908201527f43616e6e6f7420736574207073206c6173742075706461746564206f6e2061206040820152633632b0b360e11b606082015260800190565b80516020808301519190811015614976576000198160200360031b1b821691505b50919050565b60006020828403121561498e57600080fd5b8151801515811461067357600080fdfea26469706673582212201c4f5a9419045789d770a72b863358e2ba4edc698af5ea6138ceb6f12005ad9064736f6c63430008110033",
}

// ChallengeManagerImplABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeManagerImplMetaData.ABI instead.
var ChallengeManagerImplABI = ChallengeManagerImplMetaData.ABI

// ChallengeManagerImplBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChallengeManagerImplMetaData.Bin instead.
var ChallengeManagerImplBin = ChallengeManagerImplMetaData.Bin

// DeployChallengeManagerImpl deploys a new Ethereum contract, binding an instance of ChallengeManagerImpl to it.
func DeployChallengeManagerImpl(auth *bind.TransactOpts, backend bind.ContractBackend, _assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriodSec *big.Int, _oneStepProofEntry common.Address) (common.Address, *types.Transaction, *ChallengeManagerImpl, error) {
	parsed, err := ChallengeManagerImplMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChallengeManagerImplBin), backend, _assertionChain, _miniStakeValue, _challengePeriodSec, _oneStepProofEntry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeManagerImpl{ChallengeManagerImplCaller: ChallengeManagerImplCaller{contract: contract}, ChallengeManagerImplTransactor: ChallengeManagerImplTransactor{contract: contract}, ChallengeManagerImplFilterer: ChallengeManagerImplFilterer{contract: contract}}, nil
}

// ChallengeManagerImpl is an auto generated Go binding around an Ethereum contract.
type ChallengeManagerImpl struct {
	ChallengeManagerImplCaller     // Read-only binding to the contract
	ChallengeManagerImplTransactor // Write-only binding to the contract
	ChallengeManagerImplFilterer   // Log filterer for contract events
}

// ChallengeManagerImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeManagerImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeManagerImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeManagerImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeManagerImplSession struct {
	Contract     *ChallengeManagerImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ChallengeManagerImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeManagerImplCallerSession struct {
	Contract *ChallengeManagerImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ChallengeManagerImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeManagerImplTransactorSession struct {
	Contract     *ChallengeManagerImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ChallengeManagerImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeManagerImplRaw struct {
	Contract *ChallengeManagerImpl // Generic contract binding to access the raw methods on
}

// ChallengeManagerImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeManagerImplCallerRaw struct {
	Contract *ChallengeManagerImplCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeManagerImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeManagerImplTransactorRaw struct {
	Contract *ChallengeManagerImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeManagerImpl creates a new instance of ChallengeManagerImpl, bound to a specific deployed contract.
func NewChallengeManagerImpl(address common.Address, backend bind.ContractBackend) (*ChallengeManagerImpl, error) {
	contract, err := bindChallengeManagerImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerImpl{ChallengeManagerImplCaller: ChallengeManagerImplCaller{contract: contract}, ChallengeManagerImplTransactor: ChallengeManagerImplTransactor{contract: contract}, ChallengeManagerImplFilterer: ChallengeManagerImplFilterer{contract: contract}}, nil
}

// NewChallengeManagerImplCaller creates a new read-only instance of ChallengeManagerImpl, bound to a specific deployed contract.
func NewChallengeManagerImplCaller(address common.Address, caller bind.ContractCaller) (*ChallengeManagerImplCaller, error) {
	contract, err := bindChallengeManagerImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerImplCaller{contract: contract}, nil
}

// NewChallengeManagerImplTransactor creates a new write-only instance of ChallengeManagerImpl, bound to a specific deployed contract.
func NewChallengeManagerImplTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeManagerImplTransactor, error) {
	contract, err := bindChallengeManagerImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerImplTransactor{contract: contract}, nil
}

// NewChallengeManagerImplFilterer creates a new log filterer instance of ChallengeManagerImpl, bound to a specific deployed contract.
func NewChallengeManagerImplFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeManagerImplFilterer, error) {
	contract, err := bindChallengeManagerImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerImplFilterer{contract: contract}, nil
}

// bindChallengeManagerImpl binds a generic wrapper to an already deployed contract.
func bindChallengeManagerImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManagerImpl *ChallengeManagerImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeManagerImpl.Contract.ChallengeManagerImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManagerImpl *ChallengeManagerImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.ChallengeManagerImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManagerImpl *ChallengeManagerImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.ChallengeManagerImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManagerImpl *ChallengeManagerImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeManagerImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.contract.Transact(opts, method, params...)
}

// AssertionChain is a free data retrieval call binding the contract method 0x48dd2924.
//
// Solidity: function assertionChain() view returns(address)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) AssertionChain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "assertionChain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssertionChain is a free data retrieval call binding the contract method 0x48dd2924.
//
// Solidity: function assertionChain() view returns(address)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) AssertionChain() (common.Address, error) {
	return _ChallengeManagerImpl.Contract.AssertionChain(&_ChallengeManagerImpl.CallOpts)
}

// AssertionChain is a free data retrieval call binding the contract method 0x48dd2924.
//
// Solidity: function assertionChain() view returns(address)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) AssertionChain() (common.Address, error) {
	return _ChallengeManagerImpl.Contract.AssertionChain(&_ChallengeManagerImpl.CallOpts)
}

// CalculateChallengeId is a free data retrieval call binding the contract method 0x16ef5534.
//
// Solidity: function calculateChallengeId(bytes32 assertionId, uint8 typ) pure returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) CalculateChallengeId(opts *bind.CallOpts, assertionId [32]byte, typ uint8) ([32]byte, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "calculateChallengeId", assertionId, typ)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateChallengeId is a free data retrieval call binding the contract method 0x16ef5534.
//
// Solidity: function calculateChallengeId(bytes32 assertionId, uint8 typ) pure returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) CalculateChallengeId(assertionId [32]byte, typ uint8) ([32]byte, error) {
	return _ChallengeManagerImpl.Contract.CalculateChallengeId(&_ChallengeManagerImpl.CallOpts, assertionId, typ)
}

// CalculateChallengeId is a free data retrieval call binding the contract method 0x16ef5534.
//
// Solidity: function calculateChallengeId(bytes32 assertionId, uint8 typ) pure returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) CalculateChallengeId(assertionId [32]byte, typ uint8) ([32]byte, error) {
	return _ChallengeManagerImpl.Contract.CalculateChallengeId(&_ChallengeManagerImpl.CallOpts, assertionId, typ)
}

// CalculateChallengeVertexId is a free data retrieval call binding the contract method 0x4a658788.
//
// Solidity: function calculateChallengeVertexId(bytes32 challengeId, bytes32 commitmentMerkle, uint256 commitmentHeight) pure returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) CalculateChallengeVertexId(opts *bind.CallOpts, challengeId [32]byte, commitmentMerkle [32]byte, commitmentHeight *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "calculateChallengeVertexId", challengeId, commitmentMerkle, commitmentHeight)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateChallengeVertexId is a free data retrieval call binding the contract method 0x4a658788.
//
// Solidity: function calculateChallengeVertexId(bytes32 challengeId, bytes32 commitmentMerkle, uint256 commitmentHeight) pure returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) CalculateChallengeVertexId(challengeId [32]byte, commitmentMerkle [32]byte, commitmentHeight *big.Int) ([32]byte, error) {
	return _ChallengeManagerImpl.Contract.CalculateChallengeVertexId(&_ChallengeManagerImpl.CallOpts, challengeId, commitmentMerkle, commitmentHeight)
}

// CalculateChallengeVertexId is a free data retrieval call binding the contract method 0x4a658788.
//
// Solidity: function calculateChallengeVertexId(bytes32 challengeId, bytes32 commitmentMerkle, uint256 commitmentHeight) pure returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) CalculateChallengeVertexId(challengeId [32]byte, commitmentMerkle [32]byte, commitmentHeight *big.Int) ([32]byte, error) {
	return _ChallengeManagerImpl.Contract.CalculateChallengeVertexId(&_ChallengeManagerImpl.CallOpts, challengeId, commitmentMerkle, commitmentHeight)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(bytes32 challengeId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) ChallengeExists(opts *bind.CallOpts, challengeId [32]byte) (bool, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "challengeExists", challengeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(bytes32 challengeId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) ChallengeExists(challengeId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.ChallengeExists(&_ChallengeManagerImpl.CallOpts, challengeId)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(bytes32 challengeId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) ChallengeExists(challengeId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.ChallengeExists(&_ChallengeManagerImpl.CallOpts, challengeId)
}

// ChallengePeriodSec is a free data retrieval call binding the contract method 0x654f0dc2.
//
// Solidity: function challengePeriodSec() view returns(uint256)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) ChallengePeriodSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "challengePeriodSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengePeriodSec is a free data retrieval call binding the contract method 0x654f0dc2.
//
// Solidity: function challengePeriodSec() view returns(uint256)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) ChallengePeriodSec() (*big.Int, error) {
	return _ChallengeManagerImpl.Contract.ChallengePeriodSec(&_ChallengeManagerImpl.CallOpts)
}

// ChallengePeriodSec is a free data retrieval call binding the contract method 0x654f0dc2.
//
// Solidity: function challengePeriodSec() view returns(uint256)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) ChallengePeriodSec() (*big.Int, error) {
	return _ChallengeManagerImpl.Contract.ChallengePeriodSec(&_ChallengeManagerImpl.CallOpts)
}

// Challenges is a free data retrieval call binding the contract method 0xc1e69b66.
//
// Solidity: function challenges(bytes32 ) view returns(bytes32 rootId, bytes32 winningClaim, uint8 challengeType)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) Challenges(opts *bind.CallOpts, arg0 [32]byte) (struct {
	RootId        [32]byte
	WinningClaim  [32]byte
	ChallengeType uint8
}, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "challenges", arg0)

	outstruct := new(struct {
		RootId        [32]byte
		WinningClaim  [32]byte
		ChallengeType uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RootId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.WinningClaim = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ChallengeType = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0xc1e69b66.
//
// Solidity: function challenges(bytes32 ) view returns(bytes32 rootId, bytes32 winningClaim, uint8 challengeType)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) Challenges(arg0 [32]byte) (struct {
	RootId        [32]byte
	WinningClaim  [32]byte
	ChallengeType uint8
}, error) {
	return _ChallengeManagerImpl.Contract.Challenges(&_ChallengeManagerImpl.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0xc1e69b66.
//
// Solidity: function challenges(bytes32 ) view returns(bytes32 rootId, bytes32 winningClaim, uint8 challengeType)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) Challenges(arg0 [32]byte) (struct {
	RootId        [32]byte
	WinningClaim  [32]byte
	ChallengeType uint8
}, error) {
	return _ChallengeManagerImpl.Contract.Challenges(&_ChallengeManagerImpl.CallOpts, arg0)
}

// ChildrenAreAtOneStepFork is a free data retrieval call binding the contract method 0x7a4d47dc.
//
// Solidity: function childrenAreAtOneStepFork(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) ChildrenAreAtOneStepFork(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "childrenAreAtOneStepFork", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChildrenAreAtOneStepFork is a free data retrieval call binding the contract method 0x7a4d47dc.
//
// Solidity: function childrenAreAtOneStepFork(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) ChildrenAreAtOneStepFork(vId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.ChildrenAreAtOneStepFork(&_ChallengeManagerImpl.CallOpts, vId)
}

// ChildrenAreAtOneStepFork is a free data retrieval call binding the contract method 0x7a4d47dc.
//
// Solidity: function childrenAreAtOneStepFork(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) ChildrenAreAtOneStepFork(vId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.ChildrenAreAtOneStepFork(&_ChallengeManagerImpl.CallOpts, vId)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 challengeId) view returns((bytes32,bytes32,uint8))
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) GetChallenge(opts *bind.CallOpts, challengeId [32]byte) (Challenge, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "getChallenge", challengeId)

	if err != nil {
		return *new(Challenge), err
	}

	out0 := *abi.ConvertType(out[0], new(Challenge)).(*Challenge)

	return out0, err

}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 challengeId) view returns((bytes32,bytes32,uint8))
func (_ChallengeManagerImpl *ChallengeManagerImplSession) GetChallenge(challengeId [32]byte) (Challenge, error) {
	return _ChallengeManagerImpl.Contract.GetChallenge(&_ChallengeManagerImpl.CallOpts, challengeId)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 challengeId) view returns((bytes32,bytes32,uint8))
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) GetChallenge(challengeId [32]byte) (Challenge, error) {
	return _ChallengeManagerImpl.Contract.GetChallenge(&_ChallengeManagerImpl.CallOpts, challengeId)
}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x8ac04349.
//
// Solidity: function getCurrentPsTimer(bytes32 vId) view returns(uint256)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) GetCurrentPsTimer(opts *bind.CallOpts, vId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "getCurrentPsTimer", vId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x8ac04349.
//
// Solidity: function getCurrentPsTimer(bytes32 vId) view returns(uint256)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) GetCurrentPsTimer(vId [32]byte) (*big.Int, error) {
	return _ChallengeManagerImpl.Contract.GetCurrentPsTimer(&_ChallengeManagerImpl.CallOpts, vId)
}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x8ac04349.
//
// Solidity: function getCurrentPsTimer(bytes32 vId) view returns(uint256)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) GetCurrentPsTimer(vId [32]byte) (*big.Int, error) {
	return _ChallengeManagerImpl.Contract.GetCurrentPsTimer(&_ChallengeManagerImpl.CallOpts, vId)
}

// GetVertex is a free data retrieval call binding the contract method 0x86f048ed.
//
// Solidity: function getVertex(bytes32 vId) view returns((bytes32,bytes32,uint256,bytes32,bytes32,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) GetVertex(opts *bind.CallOpts, vId [32]byte) (ChallengeVertex, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "getVertex", vId)

	if err != nil {
		return *new(ChallengeVertex), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeVertex)).(*ChallengeVertex)

	return out0, err

}

// GetVertex is a free data retrieval call binding the contract method 0x86f048ed.
//
// Solidity: function getVertex(bytes32 vId) view returns((bytes32,bytes32,uint256,bytes32,bytes32,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_ChallengeManagerImpl *ChallengeManagerImplSession) GetVertex(vId [32]byte) (ChallengeVertex, error) {
	return _ChallengeManagerImpl.Contract.GetVertex(&_ChallengeManagerImpl.CallOpts, vId)
}

// GetVertex is a free data retrieval call binding the contract method 0x86f048ed.
//
// Solidity: function getVertex(bytes32 vId) view returns((bytes32,bytes32,uint256,bytes32,bytes32,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) GetVertex(vId [32]byte) (ChallengeVertex, error) {
	return _ChallengeManagerImpl.Contract.GetVertex(&_ChallengeManagerImpl.CallOpts, vId)
}

// HasConfirmedSibling is a free data retrieval call binding the contract method 0x98b67d59.
//
// Solidity: function hasConfirmedSibling(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) HasConfirmedSibling(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "hasConfirmedSibling", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasConfirmedSibling is a free data retrieval call binding the contract method 0x98b67d59.
//
// Solidity: function hasConfirmedSibling(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) HasConfirmedSibling(vId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.HasConfirmedSibling(&_ChallengeManagerImpl.CallOpts, vId)
}

// HasConfirmedSibling is a free data retrieval call binding the contract method 0x98b67d59.
//
// Solidity: function hasConfirmedSibling(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) HasConfirmedSibling(vId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.HasConfirmedSibling(&_ChallengeManagerImpl.CallOpts, vId)
}

// IsPresumptiveSuccessor is a free data retrieval call binding the contract method 0xe41b5058.
//
// Solidity: function isPresumptiveSuccessor(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) IsPresumptiveSuccessor(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "isPresumptiveSuccessor", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPresumptiveSuccessor is a free data retrieval call binding the contract method 0xe41b5058.
//
// Solidity: function isPresumptiveSuccessor(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) IsPresumptiveSuccessor(vId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.IsPresumptiveSuccessor(&_ChallengeManagerImpl.CallOpts, vId)
}

// IsPresumptiveSuccessor is a free data retrieval call binding the contract method 0xe41b5058.
//
// Solidity: function isPresumptiveSuccessor(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) IsPresumptiveSuccessor(vId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.IsPresumptiveSuccessor(&_ChallengeManagerImpl.CallOpts, vId)
}

// MiniStakeValue is a free data retrieval call binding the contract method 0x59c69996.
//
// Solidity: function miniStakeValue() view returns(uint256)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) MiniStakeValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "miniStakeValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiniStakeValue is a free data retrieval call binding the contract method 0x59c69996.
//
// Solidity: function miniStakeValue() view returns(uint256)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) MiniStakeValue() (*big.Int, error) {
	return _ChallengeManagerImpl.Contract.MiniStakeValue(&_ChallengeManagerImpl.CallOpts)
}

// MiniStakeValue is a free data retrieval call binding the contract method 0x59c69996.
//
// Solidity: function miniStakeValue() view returns(uint256)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) MiniStakeValue() (*big.Int, error) {
	return _ChallengeManagerImpl.Contract.MiniStakeValue(&_ChallengeManagerImpl.CallOpts)
}

// VertexExists is a free data retrieval call binding the contract method 0x6b0b2592.
//
// Solidity: function vertexExists(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) VertexExists(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "vertexExists", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VertexExists is a free data retrieval call binding the contract method 0x6b0b2592.
//
// Solidity: function vertexExists(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) VertexExists(vId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.VertexExists(&_ChallengeManagerImpl.CallOpts, vId)
}

// VertexExists is a free data retrieval call binding the contract method 0x6b0b2592.
//
// Solidity: function vertexExists(bytes32 vId) view returns(bool)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) VertexExists(vId [32]byte) (bool, error) {
	return _ChallengeManagerImpl.Contract.VertexExists(&_ChallengeManagerImpl.CallOpts, vId)
}

// Vertices is a free data retrieval call binding the contract method 0xf4f81db2.
//
// Solidity: function vertices(bytes32 ) view returns(bytes32 challengeId, bytes32 historyRoot, uint256 height, bytes32 successionChallenge, bytes32 predecessorId, bytes32 claimId, address staker, uint8 status, bytes32 psId, uint256 psLastUpdatedTimestamp, uint256 flushedPsTimeSec, bytes32 lowestHeightSuccessorId)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) Vertices(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ChallengeId             [32]byte
	HistoryRoot             [32]byte
	Height                  *big.Int
	SuccessionChallenge     [32]byte
	PredecessorId           [32]byte
	ClaimId                 [32]byte
	Staker                  common.Address
	Status                  uint8
	PsId                    [32]byte
	PsLastUpdatedTimestamp  *big.Int
	FlushedPsTimeSec        *big.Int
	LowestHeightSuccessorId [32]byte
}, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "vertices", arg0)

	outstruct := new(struct {
		ChallengeId             [32]byte
		HistoryRoot             [32]byte
		Height                  *big.Int
		SuccessionChallenge     [32]byte
		PredecessorId           [32]byte
		ClaimId                 [32]byte
		Staker                  common.Address
		Status                  uint8
		PsId                    [32]byte
		PsLastUpdatedTimestamp  *big.Int
		FlushedPsTimeSec        *big.Int
		LowestHeightSuccessorId [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChallengeId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.HistoryRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Height = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SuccessionChallenge = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.PredecessorId = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.ClaimId = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Staker = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.PsId = *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	outstruct.PsLastUpdatedTimestamp = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.FlushedPsTimeSec = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.LowestHeightSuccessorId = *abi.ConvertType(out[11], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Vertices is a free data retrieval call binding the contract method 0xf4f81db2.
//
// Solidity: function vertices(bytes32 ) view returns(bytes32 challengeId, bytes32 historyRoot, uint256 height, bytes32 successionChallenge, bytes32 predecessorId, bytes32 claimId, address staker, uint8 status, bytes32 psId, uint256 psLastUpdatedTimestamp, uint256 flushedPsTimeSec, bytes32 lowestHeightSuccessorId)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) Vertices(arg0 [32]byte) (struct {
	ChallengeId             [32]byte
	HistoryRoot             [32]byte
	Height                  *big.Int
	SuccessionChallenge     [32]byte
	PredecessorId           [32]byte
	ClaimId                 [32]byte
	Staker                  common.Address
	Status                  uint8
	PsId                    [32]byte
	PsLastUpdatedTimestamp  *big.Int
	FlushedPsTimeSec        *big.Int
	LowestHeightSuccessorId [32]byte
}, error) {
	return _ChallengeManagerImpl.Contract.Vertices(&_ChallengeManagerImpl.CallOpts, arg0)
}

// Vertices is a free data retrieval call binding the contract method 0xf4f81db2.
//
// Solidity: function vertices(bytes32 ) view returns(bytes32 challengeId, bytes32 historyRoot, uint256 height, bytes32 successionChallenge, bytes32 predecessorId, bytes32 claimId, address staker, uint8 status, bytes32 psId, uint256 psLastUpdatedTimestamp, uint256 flushedPsTimeSec, bytes32 lowestHeightSuccessorId)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) Vertices(arg0 [32]byte) (struct {
	ChallengeId             [32]byte
	HistoryRoot             [32]byte
	Height                  *big.Int
	SuccessionChallenge     [32]byte
	PredecessorId           [32]byte
	ClaimId                 [32]byte
	Staker                  common.Address
	Status                  uint8
	PsId                    [32]byte
	PsLastUpdatedTimestamp  *big.Int
	FlushedPsTimeSec        *big.Int
	LowestHeightSuccessorId [32]byte
}, error) {
	return _ChallengeManagerImpl.Contract.Vertices(&_ChallengeManagerImpl.CallOpts, arg0)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplCaller) WinningClaim(opts *bind.CallOpts, challengeId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ChallengeManagerImpl.contract.Call(opts, &out, "winningClaim", challengeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _ChallengeManagerImpl.Contract.WinningClaim(&_ChallengeManagerImpl.CallOpts, challengeId)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplCallerSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _ChallengeManagerImpl.Contract.WinningClaim(&_ChallengeManagerImpl.CallOpts, challengeId)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) payable returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactor) AddLeaf(opts *bind.TransactOpts, leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.contract.Transact(opts, "addLeaf", leafData, proof1, proof2)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) payable returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) AddLeaf(leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.AddLeaf(&_ChallengeManagerImpl.TransactOpts, leafData, proof1, proof2)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) payable returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorSession) AddLeaf(leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.AddLeaf(&_ChallengeManagerImpl.TransactOpts, leafData, proof1, proof2)
}

// Bisect is a paid mutator transaction binding the contract method 0x359076cf.
//
// Solidity: function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactor) Bisect(opts *bind.TransactOpts, vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.contract.Transact(opts, "bisect", vId, prefixHistoryRoot, prefixProof)
}

// Bisect is a paid mutator transaction binding the contract method 0x359076cf.
//
// Solidity: function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) Bisect(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.Bisect(&_ChallengeManagerImpl.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// Bisect is a paid mutator transaction binding the contract method 0x359076cf.
//
// Solidity: function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorSession) Bisect(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.Bisect(&_ChallengeManagerImpl.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x1d5618ac.
//
// Solidity: function confirmForPsTimer(bytes32 vId) returns()
func (_ChallengeManagerImpl *ChallengeManagerImplTransactor) ConfirmForPsTimer(opts *bind.TransactOpts, vId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.contract.Transact(opts, "confirmForPsTimer", vId)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x1d5618ac.
//
// Solidity: function confirmForPsTimer(bytes32 vId) returns()
func (_ChallengeManagerImpl *ChallengeManagerImplSession) ConfirmForPsTimer(vId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.ConfirmForPsTimer(&_ChallengeManagerImpl.TransactOpts, vId)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x1d5618ac.
//
// Solidity: function confirmForPsTimer(bytes32 vId) returns()
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorSession) ConfirmForPsTimer(vId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.ConfirmForPsTimer(&_ChallengeManagerImpl.TransactOpts, vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0xd1bac9a4.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 vId) returns()
func (_ChallengeManagerImpl *ChallengeManagerImplTransactor) ConfirmForSucessionChallengeWin(opts *bind.TransactOpts, vId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.contract.Transact(opts, "confirmForSucessionChallengeWin", vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0xd1bac9a4.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 vId) returns()
func (_ChallengeManagerImpl *ChallengeManagerImplSession) ConfirmForSucessionChallengeWin(vId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.ConfirmForSucessionChallengeWin(&_ChallengeManagerImpl.TransactOpts, vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0xd1bac9a4.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 vId) returns()
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorSession) ConfirmForSucessionChallengeWin(vId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.ConfirmForSucessionChallengeWin(&_ChallengeManagerImpl.TransactOpts, vId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 assertionId) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactor) CreateChallenge(opts *bind.TransactOpts, assertionId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.contract.Transact(opts, "createChallenge", assertionId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 assertionId) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) CreateChallenge(assertionId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.CreateChallenge(&_ChallengeManagerImpl.TransactOpts, assertionId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 assertionId) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorSession) CreateChallenge(assertionId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.CreateChallenge(&_ChallengeManagerImpl.TransactOpts, assertionId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xbd623251.
//
// Solidity: function createSubChallenge(bytes32 vId) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactor) CreateSubChallenge(opts *bind.TransactOpts, vId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.contract.Transact(opts, "createSubChallenge", vId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xbd623251.
//
// Solidity: function createSubChallenge(bytes32 vId) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) CreateSubChallenge(vId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.CreateSubChallenge(&_ChallengeManagerImpl.TransactOpts, vId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xbd623251.
//
// Solidity: function createSubChallenge(bytes32 vId) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorSession) CreateSubChallenge(vId [32]byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.CreateSubChallenge(&_ChallengeManagerImpl.TransactOpts, vId)
}

// ExecuteOneStep is a paid mutator transaction binding the contract method 0x6e49e3f2.
//
// Solidity: function executeOneStep(bytes32 winnerVId, ((uint256,address),uint256,bytes32,bytes) oneStepData, bytes beforeHistoryInclusionProof, bytes afterHistoryInclusionProof) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactor) ExecuteOneStep(opts *bind.TransactOpts, winnerVId [32]byte, oneStepData OneStepData, beforeHistoryInclusionProof []byte, afterHistoryInclusionProof []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.contract.Transact(opts, "executeOneStep", winnerVId, oneStepData, beforeHistoryInclusionProof, afterHistoryInclusionProof)
}

// ExecuteOneStep is a paid mutator transaction binding the contract method 0x6e49e3f2.
//
// Solidity: function executeOneStep(bytes32 winnerVId, ((uint256,address),uint256,bytes32,bytes) oneStepData, bytes beforeHistoryInclusionProof, bytes afterHistoryInclusionProof) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) ExecuteOneStep(winnerVId [32]byte, oneStepData OneStepData, beforeHistoryInclusionProof []byte, afterHistoryInclusionProof []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.ExecuteOneStep(&_ChallengeManagerImpl.TransactOpts, winnerVId, oneStepData, beforeHistoryInclusionProof, afterHistoryInclusionProof)
}

// ExecuteOneStep is a paid mutator transaction binding the contract method 0x6e49e3f2.
//
// Solidity: function executeOneStep(bytes32 winnerVId, ((uint256,address),uint256,bytes32,bytes) oneStepData, bytes beforeHistoryInclusionProof, bytes afterHistoryInclusionProof) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorSession) ExecuteOneStep(winnerVId [32]byte, oneStepData OneStepData, beforeHistoryInclusionProof []byte, afterHistoryInclusionProof []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.ExecuteOneStep(&_ChallengeManagerImpl.TransactOpts, winnerVId, oneStepData, beforeHistoryInclusionProof, afterHistoryInclusionProof)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e3d87cd.
//
// Solidity: function initialize(address _assertionChain, uint256 _miniStakeValue, uint256 _challengePeriodSec, address _oneStepProofEntry) returns()
func (_ChallengeManagerImpl *ChallengeManagerImplTransactor) Initialize(opts *bind.TransactOpts, _assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriodSec *big.Int, _oneStepProofEntry common.Address) (*types.Transaction, error) {
	return _ChallengeManagerImpl.contract.Transact(opts, "initialize", _assertionChain, _miniStakeValue, _challengePeriodSec, _oneStepProofEntry)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e3d87cd.
//
// Solidity: function initialize(address _assertionChain, uint256 _miniStakeValue, uint256 _challengePeriodSec, address _oneStepProofEntry) returns()
func (_ChallengeManagerImpl *ChallengeManagerImplSession) Initialize(_assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriodSec *big.Int, _oneStepProofEntry common.Address) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.Initialize(&_ChallengeManagerImpl.TransactOpts, _assertionChain, _miniStakeValue, _challengePeriodSec, _oneStepProofEntry)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e3d87cd.
//
// Solidity: function initialize(address _assertionChain, uint256 _miniStakeValue, uint256 _challengePeriodSec, address _oneStepProofEntry) returns()
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorSession) Initialize(_assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriodSec *big.Int, _oneStepProofEntry common.Address) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.Initialize(&_ChallengeManagerImpl.TransactOpts, _assertionChain, _miniStakeValue, _challengePeriodSec, _oneStepProofEntry)
}

// Merge is a paid mutator transaction binding the contract method 0x597e1e0b.
//
// Solidity: function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactor) Merge(opts *bind.TransactOpts, vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.contract.Transact(opts, "merge", vId, prefixHistoryRoot, prefixProof)
}

// Merge is a paid mutator transaction binding the contract method 0x597e1e0b.
//
// Solidity: function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplSession) Merge(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.Merge(&_ChallengeManagerImpl.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// Merge is a paid mutator transaction binding the contract method 0x597e1e0b.
//
// Solidity: function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_ChallengeManagerImpl *ChallengeManagerImplTransactorSession) Merge(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _ChallengeManagerImpl.Contract.Merge(&_ChallengeManagerImpl.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// ChallengeManagerImplBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the ChallengeManagerImpl contract.
type ChallengeManagerImplBisectedIterator struct {
	Event *ChallengeManagerImplBisected // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerImplBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerImplBisected)
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
		it.Event = new(ChallengeManagerImplBisected)
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
func (it *ChallengeManagerImplBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerImplBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerImplBisected represents a Bisected event raised by the ChallengeManagerImpl contract.
type ChallengeManagerImplBisected struct {
	FromId [32]byte
	ToId   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x69d5465c81edf7aaaf2e5c6c8829500df87d84c87f8d5b1221b59eaeaca70d27.
//
// Solidity: event Bisected(bytes32 fromId, bytes32 toId)
func (_ChallengeManagerImpl *ChallengeManagerImplFilterer) FilterBisected(opts *bind.FilterOpts) (*ChallengeManagerImplBisectedIterator, error) {

	logs, sub, err := _ChallengeManagerImpl.contract.FilterLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerImplBisectedIterator{contract: _ChallengeManagerImpl.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x69d5465c81edf7aaaf2e5c6c8829500df87d84c87f8d5b1221b59eaeaca70d27.
//
// Solidity: event Bisected(bytes32 fromId, bytes32 toId)
func (_ChallengeManagerImpl *ChallengeManagerImplFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *ChallengeManagerImplBisected) (event.Subscription, error) {

	logs, sub, err := _ChallengeManagerImpl.contract.WatchLogs(opts, "Bisected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerImplBisected)
				if err := _ChallengeManagerImpl.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x69d5465c81edf7aaaf2e5c6c8829500df87d84c87f8d5b1221b59eaeaca70d27.
//
// Solidity: event Bisected(bytes32 fromId, bytes32 toId)
func (_ChallengeManagerImpl *ChallengeManagerImplFilterer) ParseBisected(log types.Log) (*ChallengeManagerImplBisected, error) {
	event := new(ChallengeManagerImplBisected)
	if err := _ChallengeManagerImpl.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerImplMergedIterator is returned from FilterMerged and is used to iterate over the raw logs and unpacked data for Merged events raised by the ChallengeManagerImpl contract.
type ChallengeManagerImplMergedIterator struct {
	Event *ChallengeManagerImplMerged // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerImplMergedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerImplMerged)
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
		it.Event = new(ChallengeManagerImplMerged)
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
func (it *ChallengeManagerImplMergedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerImplMergedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerImplMerged represents a Merged event raised by the ChallengeManagerImpl contract.
type ChallengeManagerImplMerged struct {
	FromId [32]byte
	ToId   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMerged is a free log retrieval operation binding the contract event 0x72b50597145599e4288d411331c925b40b33b0fa3cccadc1f57d2a1ab973553a.
//
// Solidity: event Merged(bytes32 fromId, bytes32 toId)
func (_ChallengeManagerImpl *ChallengeManagerImplFilterer) FilterMerged(opts *bind.FilterOpts) (*ChallengeManagerImplMergedIterator, error) {

	logs, sub, err := _ChallengeManagerImpl.contract.FilterLogs(opts, "Merged")
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerImplMergedIterator{contract: _ChallengeManagerImpl.contract, event: "Merged", logs: logs, sub: sub}, nil
}

// WatchMerged is a free log subscription operation binding the contract event 0x72b50597145599e4288d411331c925b40b33b0fa3cccadc1f57d2a1ab973553a.
//
// Solidity: event Merged(bytes32 fromId, bytes32 toId)
func (_ChallengeManagerImpl *ChallengeManagerImplFilterer) WatchMerged(opts *bind.WatchOpts, sink chan<- *ChallengeManagerImplMerged) (event.Subscription, error) {

	logs, sub, err := _ChallengeManagerImpl.contract.WatchLogs(opts, "Merged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerImplMerged)
				if err := _ChallengeManagerImpl.contract.UnpackLog(event, "Merged", log); err != nil {
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

// ParseMerged is a log parse operation binding the contract event 0x72b50597145599e4288d411331c925b40b33b0fa3cccadc1f57d2a1ab973553a.
//
// Solidity: event Merged(bytes32 fromId, bytes32 toId)
func (_ChallengeManagerImpl *ChallengeManagerImplFilterer) ParseMerged(log types.Log) (*ChallengeManagerImplMerged, error) {
	event := new(ChallengeManagerImplMerged)
	if err := _ChallengeManagerImpl.contract.UnpackLog(event, "Merged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerImplVertexAddedIterator is returned from FilterVertexAdded and is used to iterate over the raw logs and unpacked data for VertexAdded events raised by the ChallengeManagerImpl contract.
type ChallengeManagerImplVertexAddedIterator struct {
	Event *ChallengeManagerImplVertexAdded // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerImplVertexAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerImplVertexAdded)
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
		it.Event = new(ChallengeManagerImplVertexAdded)
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
func (it *ChallengeManagerImplVertexAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerImplVertexAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerImplVertexAdded represents a VertexAdded event raised by the ChallengeManagerImpl contract.
type ChallengeManagerImplVertexAdded struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVertexAdded is a free log retrieval operation binding the contract event 0x4383ba11a7cd16be5880c5f674b93be38b3b1fcafd7a7b06151998fa2a675349.
//
// Solidity: event VertexAdded(bytes32 id)
func (_ChallengeManagerImpl *ChallengeManagerImplFilterer) FilterVertexAdded(opts *bind.FilterOpts) (*ChallengeManagerImplVertexAddedIterator, error) {

	logs, sub, err := _ChallengeManagerImpl.contract.FilterLogs(opts, "VertexAdded")
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerImplVertexAddedIterator{contract: _ChallengeManagerImpl.contract, event: "VertexAdded", logs: logs, sub: sub}, nil
}

// WatchVertexAdded is a free log subscription operation binding the contract event 0x4383ba11a7cd16be5880c5f674b93be38b3b1fcafd7a7b06151998fa2a675349.
//
// Solidity: event VertexAdded(bytes32 id)
func (_ChallengeManagerImpl *ChallengeManagerImplFilterer) WatchVertexAdded(opts *bind.WatchOpts, sink chan<- *ChallengeManagerImplVertexAdded) (event.Subscription, error) {

	logs, sub, err := _ChallengeManagerImpl.contract.WatchLogs(opts, "VertexAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerImplVertexAdded)
				if err := _ChallengeManagerImpl.contract.UnpackLog(event, "VertexAdded", log); err != nil {
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

// ParseVertexAdded is a log parse operation binding the contract event 0x4383ba11a7cd16be5880c5f674b93be38b3b1fcafd7a7b06151998fa2a675349.
//
// Solidity: event VertexAdded(bytes32 id)
func (_ChallengeManagerImpl *ChallengeManagerImplFilterer) ParseVertexAdded(log types.Log) (*ChallengeManagerImplVertexAdded, error) {
	event := new(ChallengeManagerImplVertexAdded)
	if err := _ChallengeManagerImpl.contract.UnpackLog(event, "VertexAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerLibMetaData contains all meta data concerning the ChallengeManagerLib contract.
var ChallengeManagerLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c68cef5474305697a6c5465466a1833c55b1bb2cd641dfe1f3eee909aa1980c464736f6c63430008110033",
}

// ChallengeManagerLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeManagerLibMetaData.ABI instead.
var ChallengeManagerLibABI = ChallengeManagerLibMetaData.ABI

// ChallengeManagerLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChallengeManagerLibMetaData.Bin instead.
var ChallengeManagerLibBin = ChallengeManagerLibMetaData.Bin

// DeployChallengeManagerLib deploys a new Ethereum contract, binding an instance of ChallengeManagerLib to it.
func DeployChallengeManagerLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeManagerLib, error) {
	parsed, err := ChallengeManagerLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChallengeManagerLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeManagerLib{ChallengeManagerLibCaller: ChallengeManagerLibCaller{contract: contract}, ChallengeManagerLibTransactor: ChallengeManagerLibTransactor{contract: contract}, ChallengeManagerLibFilterer: ChallengeManagerLibFilterer{contract: contract}}, nil
}

// ChallengeManagerLib is an auto generated Go binding around an Ethereum contract.
type ChallengeManagerLib struct {
	ChallengeManagerLibCaller     // Read-only binding to the contract
	ChallengeManagerLibTransactor // Write-only binding to the contract
	ChallengeManagerLibFilterer   // Log filterer for contract events
}

// ChallengeManagerLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeManagerLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeManagerLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeManagerLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeManagerLibSession struct {
	Contract     *ChallengeManagerLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChallengeManagerLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeManagerLibCallerSession struct {
	Contract *ChallengeManagerLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ChallengeManagerLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeManagerLibTransactorSession struct {
	Contract     *ChallengeManagerLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ChallengeManagerLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeManagerLibRaw struct {
	Contract *ChallengeManagerLib // Generic contract binding to access the raw methods on
}

// ChallengeManagerLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeManagerLibCallerRaw struct {
	Contract *ChallengeManagerLibCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeManagerLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeManagerLibTransactorRaw struct {
	Contract *ChallengeManagerLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeManagerLib creates a new instance of ChallengeManagerLib, bound to a specific deployed contract.
func NewChallengeManagerLib(address common.Address, backend bind.ContractBackend) (*ChallengeManagerLib, error) {
	contract, err := bindChallengeManagerLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerLib{ChallengeManagerLibCaller: ChallengeManagerLibCaller{contract: contract}, ChallengeManagerLibTransactor: ChallengeManagerLibTransactor{contract: contract}, ChallengeManagerLibFilterer: ChallengeManagerLibFilterer{contract: contract}}, nil
}

// NewChallengeManagerLibCaller creates a new read-only instance of ChallengeManagerLib, bound to a specific deployed contract.
func NewChallengeManagerLibCaller(address common.Address, caller bind.ContractCaller) (*ChallengeManagerLibCaller, error) {
	contract, err := bindChallengeManagerLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerLibCaller{contract: contract}, nil
}

// NewChallengeManagerLibTransactor creates a new write-only instance of ChallengeManagerLib, bound to a specific deployed contract.
func NewChallengeManagerLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeManagerLibTransactor, error) {
	contract, err := bindChallengeManagerLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerLibTransactor{contract: contract}, nil
}

// NewChallengeManagerLibFilterer creates a new log filterer instance of ChallengeManagerLib, bound to a specific deployed contract.
func NewChallengeManagerLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeManagerLibFilterer, error) {
	contract, err := bindChallengeManagerLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerLibFilterer{contract: contract}, nil
}

// bindChallengeManagerLib binds a generic wrapper to an already deployed contract.
func bindChallengeManagerLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManagerLib *ChallengeManagerLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeManagerLib.Contract.ChallengeManagerLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManagerLib *ChallengeManagerLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManagerLib.Contract.ChallengeManagerLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManagerLib *ChallengeManagerLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManagerLib.Contract.ChallengeManagerLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManagerLib *ChallengeManagerLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeManagerLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManagerLib *ChallengeManagerLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManagerLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManagerLib *ChallengeManagerLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManagerLib.Contract.contract.Transact(opts, method, params...)
}

// IAssertionChainMetaData contains all meta data concerning the IAssertionChain contract.
var IAssertionChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getFirstChildCreationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getInboxMsgCountSeen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getPredecessorId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getStateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"getSuccessionChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"isFirstChild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAssertionChainABI is the input ABI used to generate the binding from.
// Deprecated: Use IAssertionChainMetaData.ABI instead.
var IAssertionChainABI = IAssertionChainMetaData.ABI

// IAssertionChain is an auto generated Go binding around an Ethereum contract.
type IAssertionChain struct {
	IAssertionChainCaller     // Read-only binding to the contract
	IAssertionChainTransactor // Write-only binding to the contract
	IAssertionChainFilterer   // Log filterer for contract events
}

// IAssertionChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAssertionChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssertionChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAssertionChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssertionChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAssertionChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssertionChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAssertionChainSession struct {
	Contract     *IAssertionChain  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAssertionChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAssertionChainCallerSession struct {
	Contract *IAssertionChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAssertionChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAssertionChainTransactorSession struct {
	Contract     *IAssertionChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAssertionChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAssertionChainRaw struct {
	Contract *IAssertionChain // Generic contract binding to access the raw methods on
}

// IAssertionChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAssertionChainCallerRaw struct {
	Contract *IAssertionChainCaller // Generic read-only contract binding to access the raw methods on
}

// IAssertionChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAssertionChainTransactorRaw struct {
	Contract *IAssertionChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAssertionChain creates a new instance of IAssertionChain, bound to a specific deployed contract.
func NewIAssertionChain(address common.Address, backend bind.ContractBackend) (*IAssertionChain, error) {
	contract, err := bindIAssertionChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAssertionChain{IAssertionChainCaller: IAssertionChainCaller{contract: contract}, IAssertionChainTransactor: IAssertionChainTransactor{contract: contract}, IAssertionChainFilterer: IAssertionChainFilterer{contract: contract}}, nil
}

// NewIAssertionChainCaller creates a new read-only instance of IAssertionChain, bound to a specific deployed contract.
func NewIAssertionChainCaller(address common.Address, caller bind.ContractCaller) (*IAssertionChainCaller, error) {
	contract, err := bindIAssertionChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAssertionChainCaller{contract: contract}, nil
}

// NewIAssertionChainTransactor creates a new write-only instance of IAssertionChain, bound to a specific deployed contract.
func NewIAssertionChainTransactor(address common.Address, transactor bind.ContractTransactor) (*IAssertionChainTransactor, error) {
	contract, err := bindIAssertionChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAssertionChainTransactor{contract: contract}, nil
}

// NewIAssertionChainFilterer creates a new log filterer instance of IAssertionChain, bound to a specific deployed contract.
func NewIAssertionChainFilterer(address common.Address, filterer bind.ContractFilterer) (*IAssertionChainFilterer, error) {
	contract, err := bindIAssertionChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAssertionChainFilterer{contract: contract}, nil
}

// bindIAssertionChain binds a generic wrapper to an already deployed contract.
func bindIAssertionChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAssertionChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssertionChain *IAssertionChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssertionChain.Contract.IAssertionChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssertionChain *IAssertionChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssertionChain.Contract.IAssertionChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssertionChain *IAssertionChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssertionChain.Contract.IAssertionChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssertionChain *IAssertionChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssertionChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssertionChain *IAssertionChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssertionChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssertionChain *IAssertionChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssertionChain.Contract.contract.Transact(opts, method, params...)
}

// GetFirstChildCreationTime is a free data retrieval call binding the contract method 0x43ed6ad9.
//
// Solidity: function getFirstChildCreationTime(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCaller) GetFirstChildCreationTime(opts *bind.CallOpts, assertionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getFirstChildCreationTime", assertionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFirstChildCreationTime is a free data retrieval call binding the contract method 0x43ed6ad9.
//
// Solidity: function getFirstChildCreationTime(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainSession) GetFirstChildCreationTime(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetFirstChildCreationTime(&_IAssertionChain.CallOpts, assertionId)
}

// GetFirstChildCreationTime is a free data retrieval call binding the contract method 0x43ed6ad9.
//
// Solidity: function getFirstChildCreationTime(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCallerSession) GetFirstChildCreationTime(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetFirstChildCreationTime(&_IAssertionChain.CallOpts, assertionId)
}

// GetHeight is a free data retrieval call binding the contract method 0x896efbf2.
//
// Solidity: function getHeight(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCaller) GetHeight(opts *bind.CallOpts, assertionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getHeight", assertionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHeight is a free data retrieval call binding the contract method 0x896efbf2.
//
// Solidity: function getHeight(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainSession) GetHeight(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetHeight(&_IAssertionChain.CallOpts, assertionId)
}

// GetHeight is a free data retrieval call binding the contract method 0x896efbf2.
//
// Solidity: function getHeight(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCallerSession) GetHeight(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetHeight(&_IAssertionChain.CallOpts, assertionId)
}

// GetInboxMsgCountSeen is a free data retrieval call binding the contract method 0x7cfd5ab9.
//
// Solidity: function getInboxMsgCountSeen(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCaller) GetInboxMsgCountSeen(opts *bind.CallOpts, assertionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getInboxMsgCountSeen", assertionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInboxMsgCountSeen is a free data retrieval call binding the contract method 0x7cfd5ab9.
//
// Solidity: function getInboxMsgCountSeen(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainSession) GetInboxMsgCountSeen(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetInboxMsgCountSeen(&_IAssertionChain.CallOpts, assertionId)
}

// GetInboxMsgCountSeen is a free data retrieval call binding the contract method 0x7cfd5ab9.
//
// Solidity: function getInboxMsgCountSeen(bytes32 assertionId) view returns(uint256)
func (_IAssertionChain *IAssertionChainCallerSession) GetInboxMsgCountSeen(assertionId [32]byte) (*big.Int, error) {
	return _IAssertionChain.Contract.GetInboxMsgCountSeen(&_IAssertionChain.CallOpts, assertionId)
}

// GetPredecessorId is a free data retrieval call binding the contract method 0x9ca565d4.
//
// Solidity: function getPredecessorId(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCaller) GetPredecessorId(opts *bind.CallOpts, assertionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getPredecessorId", assertionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPredecessorId is a free data retrieval call binding the contract method 0x9ca565d4.
//
// Solidity: function getPredecessorId(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainSession) GetPredecessorId(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetPredecessorId(&_IAssertionChain.CallOpts, assertionId)
}

// GetPredecessorId is a free data retrieval call binding the contract method 0x9ca565d4.
//
// Solidity: function getPredecessorId(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCallerSession) GetPredecessorId(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetPredecessorId(&_IAssertionChain.CallOpts, assertionId)
}

// GetStateHash is a free data retrieval call binding the contract method 0xf9bce634.
//
// Solidity: function getStateHash(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCaller) GetStateHash(opts *bind.CallOpts, assertionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getStateHash", assertionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStateHash is a free data retrieval call binding the contract method 0xf9bce634.
//
// Solidity: function getStateHash(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainSession) GetStateHash(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetStateHash(&_IAssertionChain.CallOpts, assertionId)
}

// GetStateHash is a free data retrieval call binding the contract method 0xf9bce634.
//
// Solidity: function getStateHash(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCallerSession) GetStateHash(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetStateHash(&_IAssertionChain.CallOpts, assertionId)
}

// GetSuccessionChallenge is a free data retrieval call binding the contract method 0x2fefa18a.
//
// Solidity: function getSuccessionChallenge(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCaller) GetSuccessionChallenge(opts *bind.CallOpts, assertionId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "getSuccessionChallenge", assertionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSuccessionChallenge is a free data retrieval call binding the contract method 0x2fefa18a.
//
// Solidity: function getSuccessionChallenge(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainSession) GetSuccessionChallenge(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetSuccessionChallenge(&_IAssertionChain.CallOpts, assertionId)
}

// GetSuccessionChallenge is a free data retrieval call binding the contract method 0x2fefa18a.
//
// Solidity: function getSuccessionChallenge(bytes32 assertionId) view returns(bytes32)
func (_IAssertionChain *IAssertionChainCallerSession) GetSuccessionChallenge(assertionId [32]byte) ([32]byte, error) {
	return _IAssertionChain.Contract.GetSuccessionChallenge(&_IAssertionChain.CallOpts, assertionId)
}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionId) view returns(bool)
func (_IAssertionChain *IAssertionChainCaller) IsFirstChild(opts *bind.CallOpts, assertionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IAssertionChain.contract.Call(opts, &out, "isFirstChild", assertionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionId) view returns(bool)
func (_IAssertionChain *IAssertionChainSession) IsFirstChild(assertionId [32]byte) (bool, error) {
	return _IAssertionChain.Contract.IsFirstChild(&_IAssertionChain.CallOpts, assertionId)
}

// IsFirstChild is a free data retrieval call binding the contract method 0x30836228.
//
// Solidity: function isFirstChild(bytes32 assertionId) view returns(bool)
func (_IAssertionChain *IAssertionChainCallerSession) IsFirstChild(assertionId [32]byte) (bool, error) {
	return _IAssertionChain.Contract.IsFirstChild(&_IAssertionChain.CallOpts, assertionId)
}

// IChallengeManagerMetaData contains all meta data concerning the IChallengeManager contract.
var IChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"historyRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"firstStatehistoryProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"lastState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"lastStatehistoryProof\",\"type\":\"bytes\"}],\"internalType\":\"structAddLeafArgs\",\"name\":\"leafData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof2\",\"type\":\"bytes\"}],\"name\":\"addLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prefixHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"bisect\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"challengeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"childrenAreAtOneStepFork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"confirmForPsTimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"confirmForSucessionChallengeWin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"createSubChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"getChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"winningClaim\",\"type\":\"bytes32\"},{\"internalType\":\"enumChallengeType\",\"name\":\"challengeType\",\"type\":\"uint8\"}],\"internalType\":\"structChallenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentPsTimer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"getVertex\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"historyRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"successionChallenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"enumVertexStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"psId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"psLastUpdatedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flushedPsTimeSec\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lowestHeightSuccessorId\",\"type\":\"bytes32\"}],\"internalType\":\"structChallengeVertex\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"hasConfirmedSibling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAssertionChain\",\"name\":\"_assertionChain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_miniStakeValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriod\",\"type\":\"uint256\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"_oneStepProofEntry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prefixHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"merge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"vertexExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"winningClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IChallengeManagerMetaData.ABI instead.
var IChallengeManagerABI = IChallengeManagerMetaData.ABI

// IChallengeManager is an auto generated Go binding around an Ethereum contract.
type IChallengeManager struct {
	IChallengeManagerCaller     // Read-only binding to the contract
	IChallengeManagerTransactor // Write-only binding to the contract
	IChallengeManagerFilterer   // Log filterer for contract events
}

// IChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerSession struct {
	Contract     *IChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCallerSession struct {
	Contract *IChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerTransactorSession struct {
	Contract     *IChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerRaw struct {
	Contract *IChallengeManager // Generic contract binding to access the raw methods on
}

// IChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCallerRaw struct {
	Contract *IChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactorRaw struct {
	Contract *IChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManager creates a new instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManager(address common.Address, backend bind.ContractBackend) (*IChallengeManager, error) {
	contract, err := bindIChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCaller creates a new read-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCaller, error) {
	contract, err := bindIChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCaller{contract: contract}, nil
}

// NewIChallengeManagerTransactor creates a new write-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerTransactor, error) {
	contract, err := bindIChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerTransactor{contract: contract}, nil
}

// NewIChallengeManagerFilterer creates a new log filterer instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerFilterer, error) {
	contract, err := bindIChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerFilterer{contract: contract}, nil
}

// bindIChallengeManager binds a generic wrapper to an already deployed contract.
func bindIChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.IChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(bytes32 challengeId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCaller) ChallengeExists(opts *bind.CallOpts, challengeId [32]byte) (bool, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "challengeExists", challengeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(bytes32 challengeId) view returns(bool)
func (_IChallengeManager *IChallengeManagerSession) ChallengeExists(challengeId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.ChallengeExists(&_IChallengeManager.CallOpts, challengeId)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(bytes32 challengeId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCallerSession) ChallengeExists(challengeId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.ChallengeExists(&_IChallengeManager.CallOpts, challengeId)
}

// ChildrenAreAtOneStepFork is a free data retrieval call binding the contract method 0x7a4d47dc.
//
// Solidity: function childrenAreAtOneStepFork(bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCaller) ChildrenAreAtOneStepFork(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "childrenAreAtOneStepFork", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChildrenAreAtOneStepFork is a free data retrieval call binding the contract method 0x7a4d47dc.
//
// Solidity: function childrenAreAtOneStepFork(bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerSession) ChildrenAreAtOneStepFork(vId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.ChildrenAreAtOneStepFork(&_IChallengeManager.CallOpts, vId)
}

// ChildrenAreAtOneStepFork is a free data retrieval call binding the contract method 0x7a4d47dc.
//
// Solidity: function childrenAreAtOneStepFork(bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCallerSession) ChildrenAreAtOneStepFork(vId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.ChildrenAreAtOneStepFork(&_IChallengeManager.CallOpts, vId)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 challengeId) view returns((bytes32,bytes32,uint8))
func (_IChallengeManager *IChallengeManagerCaller) GetChallenge(opts *bind.CallOpts, challengeId [32]byte) (Challenge, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "getChallenge", challengeId)

	if err != nil {
		return *new(Challenge), err
	}

	out0 := *abi.ConvertType(out[0], new(Challenge)).(*Challenge)

	return out0, err

}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 challengeId) view returns((bytes32,bytes32,uint8))
func (_IChallengeManager *IChallengeManagerSession) GetChallenge(challengeId [32]byte) (Challenge, error) {
	return _IChallengeManager.Contract.GetChallenge(&_IChallengeManager.CallOpts, challengeId)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 challengeId) view returns((bytes32,bytes32,uint8))
func (_IChallengeManager *IChallengeManagerCallerSession) GetChallenge(challengeId [32]byte) (Challenge, error) {
	return _IChallengeManager.Contract.GetChallenge(&_IChallengeManager.CallOpts, challengeId)
}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x8ac04349.
//
// Solidity: function getCurrentPsTimer(bytes32 vId) view returns(uint256)
func (_IChallengeManager *IChallengeManagerCaller) GetCurrentPsTimer(opts *bind.CallOpts, vId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "getCurrentPsTimer", vId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x8ac04349.
//
// Solidity: function getCurrentPsTimer(bytes32 vId) view returns(uint256)
func (_IChallengeManager *IChallengeManagerSession) GetCurrentPsTimer(vId [32]byte) (*big.Int, error) {
	return _IChallengeManager.Contract.GetCurrentPsTimer(&_IChallengeManager.CallOpts, vId)
}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x8ac04349.
//
// Solidity: function getCurrentPsTimer(bytes32 vId) view returns(uint256)
func (_IChallengeManager *IChallengeManagerCallerSession) GetCurrentPsTimer(vId [32]byte) (*big.Int, error) {
	return _IChallengeManager.Contract.GetCurrentPsTimer(&_IChallengeManager.CallOpts, vId)
}

// GetVertex is a free data retrieval call binding the contract method 0x86f048ed.
//
// Solidity: function getVertex(bytes32 vId) view returns((bytes32,bytes32,uint256,bytes32,bytes32,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_IChallengeManager *IChallengeManagerCaller) GetVertex(opts *bind.CallOpts, vId [32]byte) (ChallengeVertex, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "getVertex", vId)

	if err != nil {
		return *new(ChallengeVertex), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeVertex)).(*ChallengeVertex)

	return out0, err

}

// GetVertex is a free data retrieval call binding the contract method 0x86f048ed.
//
// Solidity: function getVertex(bytes32 vId) view returns((bytes32,bytes32,uint256,bytes32,bytes32,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_IChallengeManager *IChallengeManagerSession) GetVertex(vId [32]byte) (ChallengeVertex, error) {
	return _IChallengeManager.Contract.GetVertex(&_IChallengeManager.CallOpts, vId)
}

// GetVertex is a free data retrieval call binding the contract method 0x86f048ed.
//
// Solidity: function getVertex(bytes32 vId) view returns((bytes32,bytes32,uint256,bytes32,bytes32,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_IChallengeManager *IChallengeManagerCallerSession) GetVertex(vId [32]byte) (ChallengeVertex, error) {
	return _IChallengeManager.Contract.GetVertex(&_IChallengeManager.CallOpts, vId)
}

// HasConfirmedSibling is a free data retrieval call binding the contract method 0x98b67d59.
//
// Solidity: function hasConfirmedSibling(bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCaller) HasConfirmedSibling(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "hasConfirmedSibling", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasConfirmedSibling is a free data retrieval call binding the contract method 0x98b67d59.
//
// Solidity: function hasConfirmedSibling(bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerSession) HasConfirmedSibling(vId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.HasConfirmedSibling(&_IChallengeManager.CallOpts, vId)
}

// HasConfirmedSibling is a free data retrieval call binding the contract method 0x98b67d59.
//
// Solidity: function hasConfirmedSibling(bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCallerSession) HasConfirmedSibling(vId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.HasConfirmedSibling(&_IChallengeManager.CallOpts, vId)
}

// VertexExists is a free data retrieval call binding the contract method 0x6b0b2592.
//
// Solidity: function vertexExists(bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCaller) VertexExists(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "vertexExists", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VertexExists is a free data retrieval call binding the contract method 0x6b0b2592.
//
// Solidity: function vertexExists(bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerSession) VertexExists(vId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.VertexExists(&_IChallengeManager.CallOpts, vId)
}

// VertexExists is a free data retrieval call binding the contract method 0x6b0b2592.
//
// Solidity: function vertexExists(bytes32 vId) view returns(bool)
func (_IChallengeManager *IChallengeManagerCallerSession) VertexExists(vId [32]byte) (bool, error) {
	return _IChallengeManager.Contract.VertexExists(&_IChallengeManager.CallOpts, vId)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_IChallengeManager *IChallengeManagerCaller) WinningClaim(opts *bind.CallOpts, challengeId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "winningClaim", challengeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_IChallengeManager *IChallengeManagerSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _IChallengeManager.Contract.WinningClaim(&_IChallengeManager.CallOpts, challengeId)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_IChallengeManager *IChallengeManagerCallerSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _IChallengeManager.Contract.WinningClaim(&_IChallengeManager.CallOpts, challengeId)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) payable returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactor) AddLeaf(opts *bind.TransactOpts, leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "addLeaf", leafData, proof1, proof2)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) payable returns(bytes32)
func (_IChallengeManager *IChallengeManagerSession) AddLeaf(leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.AddLeaf(&_IChallengeManager.TransactOpts, leafData, proof1, proof2)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) payable returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactorSession) AddLeaf(leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.AddLeaf(&_IChallengeManager.TransactOpts, leafData, proof1, proof2)
}

// Bisect is a paid mutator transaction binding the contract method 0x359076cf.
//
// Solidity: function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactor) Bisect(opts *bind.TransactOpts, vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "bisect", vId, prefixHistoryRoot, prefixProof)
}

// Bisect is a paid mutator transaction binding the contract method 0x359076cf.
//
// Solidity: function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManager *IChallengeManagerSession) Bisect(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Bisect(&_IChallengeManager.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// Bisect is a paid mutator transaction binding the contract method 0x359076cf.
//
// Solidity: function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactorSession) Bisect(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Bisect(&_IChallengeManager.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x1d5618ac.
//
// Solidity: function confirmForPsTimer(bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerTransactor) ConfirmForPsTimer(opts *bind.TransactOpts, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "confirmForPsTimer", vId)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x1d5618ac.
//
// Solidity: function confirmForPsTimer(bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerSession) ConfirmForPsTimer(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ConfirmForPsTimer(&_IChallengeManager.TransactOpts, vId)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x1d5618ac.
//
// Solidity: function confirmForPsTimer(bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) ConfirmForPsTimer(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ConfirmForPsTimer(&_IChallengeManager.TransactOpts, vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0xd1bac9a4.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerTransactor) ConfirmForSucessionChallengeWin(opts *bind.TransactOpts, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "confirmForSucessionChallengeWin", vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0xd1bac9a4.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerSession) ConfirmForSucessionChallengeWin(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ConfirmForSucessionChallengeWin(&_IChallengeManager.TransactOpts, vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0xd1bac9a4.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 vId) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) ConfirmForSucessionChallengeWin(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ConfirmForSucessionChallengeWin(&_IChallengeManager.TransactOpts, vId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 assertionId) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactor) CreateChallenge(opts *bind.TransactOpts, assertionId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "createChallenge", assertionId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 assertionId) returns(bytes32)
func (_IChallengeManager *IChallengeManagerSession) CreateChallenge(assertionId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateChallenge(&_IChallengeManager.TransactOpts, assertionId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 assertionId) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactorSession) CreateChallenge(assertionId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateChallenge(&_IChallengeManager.TransactOpts, assertionId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xbd623251.
//
// Solidity: function createSubChallenge(bytes32 vId) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactor) CreateSubChallenge(opts *bind.TransactOpts, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "createSubChallenge", vId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xbd623251.
//
// Solidity: function createSubChallenge(bytes32 vId) returns(bytes32)
func (_IChallengeManager *IChallengeManagerSession) CreateSubChallenge(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateSubChallenge(&_IChallengeManager.TransactOpts, vId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xbd623251.
//
// Solidity: function createSubChallenge(bytes32 vId) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactorSession) CreateSubChallenge(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateSubChallenge(&_IChallengeManager.TransactOpts, vId)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e3d87cd.
//
// Solidity: function initialize(address _assertionChain, uint256 _miniStakeValue, uint256 _challengePeriod, address _oneStepProofEntry) returns()
func (_IChallengeManager *IChallengeManagerTransactor) Initialize(opts *bind.TransactOpts, _assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriod *big.Int, _oneStepProofEntry common.Address) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initialize", _assertionChain, _miniStakeValue, _challengePeriod, _oneStepProofEntry)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e3d87cd.
//
// Solidity: function initialize(address _assertionChain, uint256 _miniStakeValue, uint256 _challengePeriod, address _oneStepProofEntry) returns()
func (_IChallengeManager *IChallengeManagerSession) Initialize(_assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriod *big.Int, _oneStepProofEntry common.Address) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Initialize(&_IChallengeManager.TransactOpts, _assertionChain, _miniStakeValue, _challengePeriod, _oneStepProofEntry)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e3d87cd.
//
// Solidity: function initialize(address _assertionChain, uint256 _miniStakeValue, uint256 _challengePeriod, address _oneStepProofEntry) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) Initialize(_assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriod *big.Int, _oneStepProofEntry common.Address) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Initialize(&_IChallengeManager.TransactOpts, _assertionChain, _miniStakeValue, _challengePeriod, _oneStepProofEntry)
}

// Merge is a paid mutator transaction binding the contract method 0x597e1e0b.
//
// Solidity: function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactor) Merge(opts *bind.TransactOpts, vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "merge", vId, prefixHistoryRoot, prefixProof)
}

// Merge is a paid mutator transaction binding the contract method 0x597e1e0b.
//
// Solidity: function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManager *IChallengeManagerSession) Merge(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Merge(&_IChallengeManager.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// Merge is a paid mutator transaction binding the contract method 0x597e1e0b.
//
// Solidity: function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManager *IChallengeManagerTransactorSession) Merge(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Merge(&_IChallengeManager.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// IChallengeManagerCoreMetaData contains all meta data concerning the IChallengeManagerCore contract.
var IChallengeManagerCoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"historyRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"firstState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"firstStatehistoryProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"lastState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"lastStatehistoryProof\",\"type\":\"bytes\"}],\"internalType\":\"structAddLeafArgs\",\"name\":\"leafData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof2\",\"type\":\"bytes\"}],\"name\":\"addLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prefixHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"bisect\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"confirmForPsTimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"confirmForSucessionChallengeWin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assertionId\",\"type\":\"bytes32\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"createSubChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAssertionChain\",\"name\":\"_assertionChain\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_miniStakeValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengePeriod\",\"type\":\"uint256\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"_oneStepProofEntry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prefixHistoryRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefixProof\",\"type\":\"bytes\"}],\"name\":\"merge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IChallengeManagerCoreABI is the input ABI used to generate the binding from.
// Deprecated: Use IChallengeManagerCoreMetaData.ABI instead.
var IChallengeManagerCoreABI = IChallengeManagerCoreMetaData.ABI

// IChallengeManagerCore is an auto generated Go binding around an Ethereum contract.
type IChallengeManagerCore struct {
	IChallengeManagerCoreCaller     // Read-only binding to the contract
	IChallengeManagerCoreTransactor // Write-only binding to the contract
	IChallengeManagerCoreFilterer   // Log filterer for contract events
}

// IChallengeManagerCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerCoreSession struct {
	Contract     *IChallengeManagerCore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IChallengeManagerCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCoreCallerSession struct {
	Contract *IChallengeManagerCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IChallengeManagerCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerCoreTransactorSession struct {
	Contract     *IChallengeManagerCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IChallengeManagerCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerCoreRaw struct {
	Contract *IChallengeManagerCore // Generic contract binding to access the raw methods on
}

// IChallengeManagerCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCoreCallerRaw struct {
	Contract *IChallengeManagerCoreCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerCoreTransactorRaw struct {
	Contract *IChallengeManagerCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManagerCore creates a new instance of IChallengeManagerCore, bound to a specific deployed contract.
func NewIChallengeManagerCore(address common.Address, backend bind.ContractBackend) (*IChallengeManagerCore, error) {
	contract, err := bindIChallengeManagerCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCore{IChallengeManagerCoreCaller: IChallengeManagerCoreCaller{contract: contract}, IChallengeManagerCoreTransactor: IChallengeManagerCoreTransactor{contract: contract}, IChallengeManagerCoreFilterer: IChallengeManagerCoreFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCoreCaller creates a new read-only instance of IChallengeManagerCore, bound to a specific deployed contract.
func NewIChallengeManagerCoreCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCoreCaller, error) {
	contract, err := bindIChallengeManagerCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCoreCaller{contract: contract}, nil
}

// NewIChallengeManagerCoreTransactor creates a new write-only instance of IChallengeManagerCore, bound to a specific deployed contract.
func NewIChallengeManagerCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerCoreTransactor, error) {
	contract, err := bindIChallengeManagerCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCoreTransactor{contract: contract}, nil
}

// NewIChallengeManagerCoreFilterer creates a new log filterer instance of IChallengeManagerCore, bound to a specific deployed contract.
func NewIChallengeManagerCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerCoreFilterer, error) {
	contract, err := bindIChallengeManagerCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCoreFilterer{contract: contract}, nil
}

// bindIChallengeManagerCore binds a generic wrapper to an already deployed contract.
func bindIChallengeManagerCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManagerCore *IChallengeManagerCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManagerCore.Contract.IChallengeManagerCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManagerCore *IChallengeManagerCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.IChallengeManagerCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManagerCore *IChallengeManagerCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.IChallengeManagerCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManagerCore *IChallengeManagerCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManagerCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.contract.Transact(opts, method, params...)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) payable returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactor) AddLeaf(opts *bind.TransactOpts, leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.contract.Transact(opts, "addLeaf", leafData, proof1, proof2)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) payable returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreSession) AddLeaf(leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.AddLeaf(&_IChallengeManagerCore.TransactOpts, leafData, proof1, proof2)
}

// AddLeaf is a paid mutator transaction binding the contract method 0xb241493b.
//
// Solidity: function addLeaf((bytes32,bytes32,uint256,bytes32,bytes32,bytes,bytes32,bytes) leafData, bytes proof1, bytes proof2) payable returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorSession) AddLeaf(leafData AddLeafArgs, proof1 []byte, proof2 []byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.AddLeaf(&_IChallengeManagerCore.TransactOpts, leafData, proof1, proof2)
}

// Bisect is a paid mutator transaction binding the contract method 0x359076cf.
//
// Solidity: function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactor) Bisect(opts *bind.TransactOpts, vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.contract.Transact(opts, "bisect", vId, prefixHistoryRoot, prefixProof)
}

// Bisect is a paid mutator transaction binding the contract method 0x359076cf.
//
// Solidity: function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreSession) Bisect(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.Bisect(&_IChallengeManagerCore.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// Bisect is a paid mutator transaction binding the contract method 0x359076cf.
//
// Solidity: function bisect(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorSession) Bisect(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.Bisect(&_IChallengeManagerCore.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x1d5618ac.
//
// Solidity: function confirmForPsTimer(bytes32 vId) returns()
func (_IChallengeManagerCore *IChallengeManagerCoreTransactor) ConfirmForPsTimer(opts *bind.TransactOpts, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.contract.Transact(opts, "confirmForPsTimer", vId)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x1d5618ac.
//
// Solidity: function confirmForPsTimer(bytes32 vId) returns()
func (_IChallengeManagerCore *IChallengeManagerCoreSession) ConfirmForPsTimer(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.ConfirmForPsTimer(&_IChallengeManagerCore.TransactOpts, vId)
}

// ConfirmForPsTimer is a paid mutator transaction binding the contract method 0x1d5618ac.
//
// Solidity: function confirmForPsTimer(bytes32 vId) returns()
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorSession) ConfirmForPsTimer(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.ConfirmForPsTimer(&_IChallengeManagerCore.TransactOpts, vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0xd1bac9a4.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 vId) returns()
func (_IChallengeManagerCore *IChallengeManagerCoreTransactor) ConfirmForSucessionChallengeWin(opts *bind.TransactOpts, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.contract.Transact(opts, "confirmForSucessionChallengeWin", vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0xd1bac9a4.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 vId) returns()
func (_IChallengeManagerCore *IChallengeManagerCoreSession) ConfirmForSucessionChallengeWin(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.ConfirmForSucessionChallengeWin(&_IChallengeManagerCore.TransactOpts, vId)
}

// ConfirmForSucessionChallengeWin is a paid mutator transaction binding the contract method 0xd1bac9a4.
//
// Solidity: function confirmForSucessionChallengeWin(bytes32 vId) returns()
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorSession) ConfirmForSucessionChallengeWin(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.ConfirmForSucessionChallengeWin(&_IChallengeManagerCore.TransactOpts, vId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 assertionId) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactor) CreateChallenge(opts *bind.TransactOpts, assertionId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.contract.Transact(opts, "createChallenge", assertionId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 assertionId) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreSession) CreateChallenge(assertionId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.CreateChallenge(&_IChallengeManagerCore.TransactOpts, assertionId)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xf696dc55.
//
// Solidity: function createChallenge(bytes32 assertionId) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorSession) CreateChallenge(assertionId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.CreateChallenge(&_IChallengeManagerCore.TransactOpts, assertionId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xbd623251.
//
// Solidity: function createSubChallenge(bytes32 vId) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactor) CreateSubChallenge(opts *bind.TransactOpts, vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.contract.Transact(opts, "createSubChallenge", vId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xbd623251.
//
// Solidity: function createSubChallenge(bytes32 vId) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreSession) CreateSubChallenge(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.CreateSubChallenge(&_IChallengeManagerCore.TransactOpts, vId)
}

// CreateSubChallenge is a paid mutator transaction binding the contract method 0xbd623251.
//
// Solidity: function createSubChallenge(bytes32 vId) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorSession) CreateSubChallenge(vId [32]byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.CreateSubChallenge(&_IChallengeManagerCore.TransactOpts, vId)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e3d87cd.
//
// Solidity: function initialize(address _assertionChain, uint256 _miniStakeValue, uint256 _challengePeriod, address _oneStepProofEntry) returns()
func (_IChallengeManagerCore *IChallengeManagerCoreTransactor) Initialize(opts *bind.TransactOpts, _assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriod *big.Int, _oneStepProofEntry common.Address) (*types.Transaction, error) {
	return _IChallengeManagerCore.contract.Transact(opts, "initialize", _assertionChain, _miniStakeValue, _challengePeriod, _oneStepProofEntry)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e3d87cd.
//
// Solidity: function initialize(address _assertionChain, uint256 _miniStakeValue, uint256 _challengePeriod, address _oneStepProofEntry) returns()
func (_IChallengeManagerCore *IChallengeManagerCoreSession) Initialize(_assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriod *big.Int, _oneStepProofEntry common.Address) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.Initialize(&_IChallengeManagerCore.TransactOpts, _assertionChain, _miniStakeValue, _challengePeriod, _oneStepProofEntry)
}

// Initialize is a paid mutator transaction binding the contract method 0x9e3d87cd.
//
// Solidity: function initialize(address _assertionChain, uint256 _miniStakeValue, uint256 _challengePeriod, address _oneStepProofEntry) returns()
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorSession) Initialize(_assertionChain common.Address, _miniStakeValue *big.Int, _challengePeriod *big.Int, _oneStepProofEntry common.Address) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.Initialize(&_IChallengeManagerCore.TransactOpts, _assertionChain, _miniStakeValue, _challengePeriod, _oneStepProofEntry)
}

// Merge is a paid mutator transaction binding the contract method 0x597e1e0b.
//
// Solidity: function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactor) Merge(opts *bind.TransactOpts, vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.contract.Transact(opts, "merge", vId, prefixHistoryRoot, prefixProof)
}

// Merge is a paid mutator transaction binding the contract method 0x597e1e0b.
//
// Solidity: function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreSession) Merge(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.Merge(&_IChallengeManagerCore.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// Merge is a paid mutator transaction binding the contract method 0x597e1e0b.
//
// Solidity: function merge(bytes32 vId, bytes32 prefixHistoryRoot, bytes prefixProof) returns(bytes32)
func (_IChallengeManagerCore *IChallengeManagerCoreTransactorSession) Merge(vId [32]byte, prefixHistoryRoot [32]byte, prefixProof []byte) (*types.Transaction, error) {
	return _IChallengeManagerCore.Contract.Merge(&_IChallengeManagerCore.TransactOpts, vId, prefixHistoryRoot, prefixProof)
}

// IChallengeManagerExternalViewMetaData contains all meta data concerning the IChallengeManagerExternalView contract.
var IChallengeManagerExternalViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"challengeExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"childrenAreAtOneStepFork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"getChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"winningClaim\",\"type\":\"bytes32\"},{\"internalType\":\"enumChallengeType\",\"name\":\"challengeType\",\"type\":\"uint8\"}],\"internalType\":\"structChallenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentPsTimer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"getVertex\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"historyRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"successionChallenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"enumVertexStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"psId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"psLastUpdatedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flushedPsTimeSec\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lowestHeightSuccessorId\",\"type\":\"bytes32\"}],\"internalType\":\"structChallengeVertex\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"hasConfirmedSibling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vId\",\"type\":\"bytes32\"}],\"name\":\"vertexExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"winningClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IChallengeManagerExternalViewABI is the input ABI used to generate the binding from.
// Deprecated: Use IChallengeManagerExternalViewMetaData.ABI instead.
var IChallengeManagerExternalViewABI = IChallengeManagerExternalViewMetaData.ABI

// IChallengeManagerExternalView is an auto generated Go binding around an Ethereum contract.
type IChallengeManagerExternalView struct {
	IChallengeManagerExternalViewCaller     // Read-only binding to the contract
	IChallengeManagerExternalViewTransactor // Write-only binding to the contract
	IChallengeManagerExternalViewFilterer   // Log filterer for contract events
}

// IChallengeManagerExternalViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerExternalViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerExternalViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerExternalViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerExternalViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerExternalViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerExternalViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerExternalViewSession struct {
	Contract     *IChallengeManagerExternalView // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IChallengeManagerExternalViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerExternalViewCallerSession struct {
	Contract *IChallengeManagerExternalViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// IChallengeManagerExternalViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerExternalViewTransactorSession struct {
	Contract     *IChallengeManagerExternalViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// IChallengeManagerExternalViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerExternalViewRaw struct {
	Contract *IChallengeManagerExternalView // Generic contract binding to access the raw methods on
}

// IChallengeManagerExternalViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerExternalViewCallerRaw struct {
	Contract *IChallengeManagerExternalViewCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerExternalViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerExternalViewTransactorRaw struct {
	Contract *IChallengeManagerExternalViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManagerExternalView creates a new instance of IChallengeManagerExternalView, bound to a specific deployed contract.
func NewIChallengeManagerExternalView(address common.Address, backend bind.ContractBackend) (*IChallengeManagerExternalView, error) {
	contract, err := bindIChallengeManagerExternalView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerExternalView{IChallengeManagerExternalViewCaller: IChallengeManagerExternalViewCaller{contract: contract}, IChallengeManagerExternalViewTransactor: IChallengeManagerExternalViewTransactor{contract: contract}, IChallengeManagerExternalViewFilterer: IChallengeManagerExternalViewFilterer{contract: contract}}, nil
}

// NewIChallengeManagerExternalViewCaller creates a new read-only instance of IChallengeManagerExternalView, bound to a specific deployed contract.
func NewIChallengeManagerExternalViewCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerExternalViewCaller, error) {
	contract, err := bindIChallengeManagerExternalView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerExternalViewCaller{contract: contract}, nil
}

// NewIChallengeManagerExternalViewTransactor creates a new write-only instance of IChallengeManagerExternalView, bound to a specific deployed contract.
func NewIChallengeManagerExternalViewTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerExternalViewTransactor, error) {
	contract, err := bindIChallengeManagerExternalView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerExternalViewTransactor{contract: contract}, nil
}

// NewIChallengeManagerExternalViewFilterer creates a new log filterer instance of IChallengeManagerExternalView, bound to a specific deployed contract.
func NewIChallengeManagerExternalViewFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerExternalViewFilterer, error) {
	contract, err := bindIChallengeManagerExternalView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerExternalViewFilterer{contract: contract}, nil
}

// bindIChallengeManagerExternalView binds a generic wrapper to an already deployed contract.
func bindIChallengeManagerExternalView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerExternalViewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManagerExternalView.Contract.IChallengeManagerExternalViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManagerExternalView.Contract.IChallengeManagerExternalViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManagerExternalView.Contract.IChallengeManagerExternalViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManagerExternalView.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManagerExternalView.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManagerExternalView.Contract.contract.Transact(opts, method, params...)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(bytes32 challengeId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCaller) ChallengeExists(opts *bind.CallOpts, challengeId [32]byte) (bool, error) {
	var out []interface{}
	err := _IChallengeManagerExternalView.contract.Call(opts, &out, "challengeExists", challengeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(bytes32 challengeId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewSession) ChallengeExists(challengeId [32]byte) (bool, error) {
	return _IChallengeManagerExternalView.Contract.ChallengeExists(&_IChallengeManagerExternalView.CallOpts, challengeId)
}

// ChallengeExists is a free data retrieval call binding the contract method 0x1b7bbecb.
//
// Solidity: function challengeExists(bytes32 challengeId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCallerSession) ChallengeExists(challengeId [32]byte) (bool, error) {
	return _IChallengeManagerExternalView.Contract.ChallengeExists(&_IChallengeManagerExternalView.CallOpts, challengeId)
}

// ChildrenAreAtOneStepFork is a free data retrieval call binding the contract method 0x7a4d47dc.
//
// Solidity: function childrenAreAtOneStepFork(bytes32 vId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCaller) ChildrenAreAtOneStepFork(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _IChallengeManagerExternalView.contract.Call(opts, &out, "childrenAreAtOneStepFork", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChildrenAreAtOneStepFork is a free data retrieval call binding the contract method 0x7a4d47dc.
//
// Solidity: function childrenAreAtOneStepFork(bytes32 vId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewSession) ChildrenAreAtOneStepFork(vId [32]byte) (bool, error) {
	return _IChallengeManagerExternalView.Contract.ChildrenAreAtOneStepFork(&_IChallengeManagerExternalView.CallOpts, vId)
}

// ChildrenAreAtOneStepFork is a free data retrieval call binding the contract method 0x7a4d47dc.
//
// Solidity: function childrenAreAtOneStepFork(bytes32 vId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCallerSession) ChildrenAreAtOneStepFork(vId [32]byte) (bool, error) {
	return _IChallengeManagerExternalView.Contract.ChildrenAreAtOneStepFork(&_IChallengeManagerExternalView.CallOpts, vId)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 challengeId) view returns((bytes32,bytes32,uint8))
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCaller) GetChallenge(opts *bind.CallOpts, challengeId [32]byte) (Challenge, error) {
	var out []interface{}
	err := _IChallengeManagerExternalView.contract.Call(opts, &out, "getChallenge", challengeId)

	if err != nil {
		return *new(Challenge), err
	}

	out0 := *abi.ConvertType(out[0], new(Challenge)).(*Challenge)

	return out0, err

}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 challengeId) view returns((bytes32,bytes32,uint8))
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewSession) GetChallenge(challengeId [32]byte) (Challenge, error) {
	return _IChallengeManagerExternalView.Contract.GetChallenge(&_IChallengeManagerExternalView.CallOpts, challengeId)
}

// GetChallenge is a free data retrieval call binding the contract method 0x458d2bf1.
//
// Solidity: function getChallenge(bytes32 challengeId) view returns((bytes32,bytes32,uint8))
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCallerSession) GetChallenge(challengeId [32]byte) (Challenge, error) {
	return _IChallengeManagerExternalView.Contract.GetChallenge(&_IChallengeManagerExternalView.CallOpts, challengeId)
}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x8ac04349.
//
// Solidity: function getCurrentPsTimer(bytes32 vId) view returns(uint256)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCaller) GetCurrentPsTimer(opts *bind.CallOpts, vId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IChallengeManagerExternalView.contract.Call(opts, &out, "getCurrentPsTimer", vId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x8ac04349.
//
// Solidity: function getCurrentPsTimer(bytes32 vId) view returns(uint256)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewSession) GetCurrentPsTimer(vId [32]byte) (*big.Int, error) {
	return _IChallengeManagerExternalView.Contract.GetCurrentPsTimer(&_IChallengeManagerExternalView.CallOpts, vId)
}

// GetCurrentPsTimer is a free data retrieval call binding the contract method 0x8ac04349.
//
// Solidity: function getCurrentPsTimer(bytes32 vId) view returns(uint256)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCallerSession) GetCurrentPsTimer(vId [32]byte) (*big.Int, error) {
	return _IChallengeManagerExternalView.Contract.GetCurrentPsTimer(&_IChallengeManagerExternalView.CallOpts, vId)
}

// GetVertex is a free data retrieval call binding the contract method 0x86f048ed.
//
// Solidity: function getVertex(bytes32 vId) view returns((bytes32,bytes32,uint256,bytes32,bytes32,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCaller) GetVertex(opts *bind.CallOpts, vId [32]byte) (ChallengeVertex, error) {
	var out []interface{}
	err := _IChallengeManagerExternalView.contract.Call(opts, &out, "getVertex", vId)

	if err != nil {
		return *new(ChallengeVertex), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeVertex)).(*ChallengeVertex)

	return out0, err

}

// GetVertex is a free data retrieval call binding the contract method 0x86f048ed.
//
// Solidity: function getVertex(bytes32 vId) view returns((bytes32,bytes32,uint256,bytes32,bytes32,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewSession) GetVertex(vId [32]byte) (ChallengeVertex, error) {
	return _IChallengeManagerExternalView.Contract.GetVertex(&_IChallengeManagerExternalView.CallOpts, vId)
}

// GetVertex is a free data retrieval call binding the contract method 0x86f048ed.
//
// Solidity: function getVertex(bytes32 vId) view returns((bytes32,bytes32,uint256,bytes32,bytes32,bytes32,address,uint8,bytes32,uint256,uint256,bytes32))
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCallerSession) GetVertex(vId [32]byte) (ChallengeVertex, error) {
	return _IChallengeManagerExternalView.Contract.GetVertex(&_IChallengeManagerExternalView.CallOpts, vId)
}

// HasConfirmedSibling is a free data retrieval call binding the contract method 0x98b67d59.
//
// Solidity: function hasConfirmedSibling(bytes32 vId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCaller) HasConfirmedSibling(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _IChallengeManagerExternalView.contract.Call(opts, &out, "hasConfirmedSibling", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasConfirmedSibling is a free data retrieval call binding the contract method 0x98b67d59.
//
// Solidity: function hasConfirmedSibling(bytes32 vId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewSession) HasConfirmedSibling(vId [32]byte) (bool, error) {
	return _IChallengeManagerExternalView.Contract.HasConfirmedSibling(&_IChallengeManagerExternalView.CallOpts, vId)
}

// HasConfirmedSibling is a free data retrieval call binding the contract method 0x98b67d59.
//
// Solidity: function hasConfirmedSibling(bytes32 vId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCallerSession) HasConfirmedSibling(vId [32]byte) (bool, error) {
	return _IChallengeManagerExternalView.Contract.HasConfirmedSibling(&_IChallengeManagerExternalView.CallOpts, vId)
}

// VertexExists is a free data retrieval call binding the contract method 0x6b0b2592.
//
// Solidity: function vertexExists(bytes32 vId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCaller) VertexExists(opts *bind.CallOpts, vId [32]byte) (bool, error) {
	var out []interface{}
	err := _IChallengeManagerExternalView.contract.Call(opts, &out, "vertexExists", vId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VertexExists is a free data retrieval call binding the contract method 0x6b0b2592.
//
// Solidity: function vertexExists(bytes32 vId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewSession) VertexExists(vId [32]byte) (bool, error) {
	return _IChallengeManagerExternalView.Contract.VertexExists(&_IChallengeManagerExternalView.CallOpts, vId)
}

// VertexExists is a free data retrieval call binding the contract method 0x6b0b2592.
//
// Solidity: function vertexExists(bytes32 vId) view returns(bool)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCallerSession) VertexExists(vId [32]byte) (bool, error) {
	return _IChallengeManagerExternalView.Contract.VertexExists(&_IChallengeManagerExternalView.CallOpts, vId)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCaller) WinningClaim(opts *bind.CallOpts, challengeId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IChallengeManagerExternalView.contract.Call(opts, &out, "winningClaim", challengeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _IChallengeManagerExternalView.Contract.WinningClaim(&_IChallengeManagerExternalView.CallOpts, challengeId)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_IChallengeManagerExternalView *IChallengeManagerExternalViewCallerSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _IChallengeManagerExternalView.Contract.WinningClaim(&_IChallengeManagerExternalView.CallOpts, challengeId)
}

// IInboxMetaData contains all meta data concerning the IInbox contract.
var IInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"msgCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use IInboxMetaData.ABI instead.
var IInboxABI = IInboxMetaData.ABI

// IInbox is an auto generated Go binding around an Ethereum contract.
type IInbox struct {
	IInboxCaller     // Read-only binding to the contract
	IInboxTransactor // Write-only binding to the contract
	IInboxFilterer   // Log filterer for contract events
}

// IInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInboxSession struct {
	Contract     *IInbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInboxCallerSession struct {
	Contract *IInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInboxTransactorSession struct {
	Contract     *IInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInboxRaw struct {
	Contract *IInbox // Generic contract binding to access the raw methods on
}

// IInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInboxCallerRaw struct {
	Contract *IInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInboxTransactorRaw struct {
	Contract *IInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInbox creates a new instance of IInbox, bound to a specific deployed contract.
func NewIInbox(address common.Address, backend bind.ContractBackend) (*IInbox, error) {
	contract, err := bindIInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInbox{IInboxCaller: IInboxCaller{contract: contract}, IInboxTransactor: IInboxTransactor{contract: contract}, IInboxFilterer: IInboxFilterer{contract: contract}}, nil
}

// NewIInboxCaller creates a new read-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxCaller(address common.Address, caller bind.ContractCaller) (*IInboxCaller, error) {
	contract, err := bindIInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxCaller{contract: contract}, nil
}

// NewIInboxTransactor creates a new write-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IInboxTransactor, error) {
	contract, err := bindIInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxTransactor{contract: contract}, nil
}

// NewIInboxFilterer creates a new log filterer instance of IInbox, bound to a specific deployed contract.
func NewIInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IInboxFilterer, error) {
	contract, err := bindIInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInboxFilterer{contract: contract}, nil
}

// bindIInbox binds a generic wrapper to an already deployed contract.
func bindIInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.IInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transact(opts, method, params...)
}

// MsgCount is a paid mutator transaction binding the contract method 0x8f1a2810.
//
// Solidity: function msgCount() returns(uint256)
func (_IInbox *IInboxTransactor) MsgCount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "msgCount")
}

// MsgCount is a paid mutator transaction binding the contract method 0x8f1a2810.
//
// Solidity: function msgCount() returns(uint256)
func (_IInbox *IInboxSession) MsgCount() (*types.Transaction, error) {
	return _IInbox.Contract.MsgCount(&_IInbox.TransactOpts)
}

// MsgCount is a paid mutator transaction binding the contract method 0x8f1a2810.
//
// Solidity: function msgCount() returns(uint256)
func (_IInbox *IInboxTransactorSession) MsgCount() (*types.Transaction, error) {
	return _IInbox.Contract.MsgCount(&_IInbox.TransactOpts)
}

// OneStepProofManagerMetaData contains all meta data concerning the OneStepProofManager contract.
var OneStepProofManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"startState\",\"type\":\"bytes32\"}],\"name\":\"createOneStepProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"startState\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_winner\",\"type\":\"bytes32\"}],\"name\":\"setWinningClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeId\",\"type\":\"bytes32\"}],\"name\":\"winningClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"winningClaims\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061018a806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632f3069611461005157806335025bde1461007357806373d154e814610098578063a4714dbb146100b8575b600080fd5b61007161005f366004610119565b60009182526020829052604090912055565b005b61008661008136600461013b565b6100d8565b60405190815260200160405180910390f35b6100866100a636600461013b565b60009081526020819052604090205490565b6100866100c636600461013b565b60006020819052908152604090205481565b60405162461bcd60e51b815260206004820152600f60248201526e1393d517d253541311535153951151608a1b604482015260009060640160405180910390fd5b6000806040838503121561012c57600080fd5b50508035926020909101359150565b60006020828403121561014d57600080fd5b503591905056fea26469706673582212209e5324e842829dc8e8eb1d68cb646baaee1e5f228d460fbcfee0ee0c5143370264736f6c63430008110033",
}

// OneStepProofManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofManagerMetaData.ABI instead.
var OneStepProofManagerABI = OneStepProofManagerMetaData.ABI

// OneStepProofManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofManagerMetaData.Bin instead.
var OneStepProofManagerBin = OneStepProofManagerMetaData.Bin

// DeployOneStepProofManager deploys a new Ethereum contract, binding an instance of OneStepProofManager to it.
func DeployOneStepProofManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProofManager, error) {
	parsed, err := OneStepProofManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofManager{OneStepProofManagerCaller: OneStepProofManagerCaller{contract: contract}, OneStepProofManagerTransactor: OneStepProofManagerTransactor{contract: contract}, OneStepProofManagerFilterer: OneStepProofManagerFilterer{contract: contract}}, nil
}

// OneStepProofManager is an auto generated Go binding around an Ethereum contract.
type OneStepProofManager struct {
	OneStepProofManagerCaller     // Read-only binding to the contract
	OneStepProofManagerTransactor // Write-only binding to the contract
	OneStepProofManagerFilterer   // Log filterer for contract events
}

// OneStepProofManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofManagerSession struct {
	Contract     *OneStepProofManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProofManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofManagerCallerSession struct {
	Contract *OneStepProofManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProofManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofManagerTransactorSession struct {
	Contract     *OneStepProofManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProofManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofManagerRaw struct {
	Contract *OneStepProofManager // Generic contract binding to access the raw methods on
}

// OneStepProofManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofManagerCallerRaw struct {
	Contract *OneStepProofManagerCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofManagerTransactorRaw struct {
	Contract *OneStepProofManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofManager creates a new instance of OneStepProofManager, bound to a specific deployed contract.
func NewOneStepProofManager(address common.Address, backend bind.ContractBackend) (*OneStepProofManager, error) {
	contract, err := bindOneStepProofManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofManager{OneStepProofManagerCaller: OneStepProofManagerCaller{contract: contract}, OneStepProofManagerTransactor: OneStepProofManagerTransactor{contract: contract}, OneStepProofManagerFilterer: OneStepProofManagerFilterer{contract: contract}}, nil
}

// NewOneStepProofManagerCaller creates a new read-only instance of OneStepProofManager, bound to a specific deployed contract.
func NewOneStepProofManagerCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofManagerCaller, error) {
	contract, err := bindOneStepProofManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofManagerCaller{contract: contract}, nil
}

// NewOneStepProofManagerTransactor creates a new write-only instance of OneStepProofManager, bound to a specific deployed contract.
func NewOneStepProofManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofManagerTransactor, error) {
	contract, err := bindOneStepProofManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofManagerTransactor{contract: contract}, nil
}

// NewOneStepProofManagerFilterer creates a new log filterer instance of OneStepProofManager, bound to a specific deployed contract.
func NewOneStepProofManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofManagerFilterer, error) {
	contract, err := bindOneStepProofManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofManagerFilterer{contract: contract}, nil
}

// bindOneStepProofManager binds a generic wrapper to an already deployed contract.
func bindOneStepProofManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofManager *OneStepProofManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofManager.Contract.OneStepProofManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofManager *OneStepProofManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofManager.Contract.OneStepProofManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofManager *OneStepProofManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofManager.Contract.OneStepProofManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofManager *OneStepProofManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofManager *OneStepProofManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofManager *OneStepProofManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofManager.Contract.contract.Transact(opts, method, params...)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_OneStepProofManager *OneStepProofManagerCaller) WinningClaim(opts *bind.CallOpts, challengeId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofManager.contract.Call(opts, &out, "winningClaim", challengeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_OneStepProofManager *OneStepProofManagerSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _OneStepProofManager.Contract.WinningClaim(&_OneStepProofManager.CallOpts, challengeId)
}

// WinningClaim is a free data retrieval call binding the contract method 0x73d154e8.
//
// Solidity: function winningClaim(bytes32 challengeId) view returns(bytes32)
func (_OneStepProofManager *OneStepProofManagerCallerSession) WinningClaim(challengeId [32]byte) ([32]byte, error) {
	return _OneStepProofManager.Contract.WinningClaim(&_OneStepProofManager.CallOpts, challengeId)
}

// WinningClaims is a free data retrieval call binding the contract method 0xa4714dbb.
//
// Solidity: function winningClaims(bytes32 ) view returns(bytes32)
func (_OneStepProofManager *OneStepProofManagerCaller) WinningClaims(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofManager.contract.Call(opts, &out, "winningClaims", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinningClaims is a free data retrieval call binding the contract method 0xa4714dbb.
//
// Solidity: function winningClaims(bytes32 ) view returns(bytes32)
func (_OneStepProofManager *OneStepProofManagerSession) WinningClaims(arg0 [32]byte) ([32]byte, error) {
	return _OneStepProofManager.Contract.WinningClaims(&_OneStepProofManager.CallOpts, arg0)
}

// WinningClaims is a free data retrieval call binding the contract method 0xa4714dbb.
//
// Solidity: function winningClaims(bytes32 ) view returns(bytes32)
func (_OneStepProofManager *OneStepProofManagerCallerSession) WinningClaims(arg0 [32]byte) ([32]byte, error) {
	return _OneStepProofManager.Contract.WinningClaims(&_OneStepProofManager.CallOpts, arg0)
}

// CreateOneStepProof is a paid mutator transaction binding the contract method 0x35025bde.
//
// Solidity: function createOneStepProof(bytes32 startState) returns(bytes32)
func (_OneStepProofManager *OneStepProofManagerTransactor) CreateOneStepProof(opts *bind.TransactOpts, startState [32]byte) (*types.Transaction, error) {
	return _OneStepProofManager.contract.Transact(opts, "createOneStepProof", startState)
}

// CreateOneStepProof is a paid mutator transaction binding the contract method 0x35025bde.
//
// Solidity: function createOneStepProof(bytes32 startState) returns(bytes32)
func (_OneStepProofManager *OneStepProofManagerSession) CreateOneStepProof(startState [32]byte) (*types.Transaction, error) {
	return _OneStepProofManager.Contract.CreateOneStepProof(&_OneStepProofManager.TransactOpts, startState)
}

// CreateOneStepProof is a paid mutator transaction binding the contract method 0x35025bde.
//
// Solidity: function createOneStepProof(bytes32 startState) returns(bytes32)
func (_OneStepProofManager *OneStepProofManagerTransactorSession) CreateOneStepProof(startState [32]byte) (*types.Transaction, error) {
	return _OneStepProofManager.Contract.CreateOneStepProof(&_OneStepProofManager.TransactOpts, startState)
}

// SetWinningClaim is a paid mutator transaction binding the contract method 0x2f306961.
//
// Solidity: function setWinningClaim(bytes32 startState, bytes32 _winner) returns()
func (_OneStepProofManager *OneStepProofManagerTransactor) SetWinningClaim(opts *bind.TransactOpts, startState [32]byte, _winner [32]byte) (*types.Transaction, error) {
	return _OneStepProofManager.contract.Transact(opts, "setWinningClaim", startState, _winner)
}

// SetWinningClaim is a paid mutator transaction binding the contract method 0x2f306961.
//
// Solidity: function setWinningClaim(bytes32 startState, bytes32 _winner) returns()
func (_OneStepProofManager *OneStepProofManagerSession) SetWinningClaim(startState [32]byte, _winner [32]byte) (*types.Transaction, error) {
	return _OneStepProofManager.Contract.SetWinningClaim(&_OneStepProofManager.TransactOpts, startState, _winner)
}

// SetWinningClaim is a paid mutator transaction binding the contract method 0x2f306961.
//
// Solidity: function setWinningClaim(bytes32 startState, bytes32 _winner) returns()
func (_OneStepProofManager *OneStepProofManagerTransactorSession) SetWinningClaim(startState [32]byte, _winner [32]byte) (*types.Transaction, error) {
	return _OneStepProofManager.Contract.SetWinningClaim(&_OneStepProofManager.TransactOpts, startState, _winner)
}
