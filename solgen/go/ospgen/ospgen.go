// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ospgen

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

// ExecutionContext is an auto generated low-level Go binding around an user-defined struct.
type ExecutionContext struct {
	MaxInboxMessagesRead  *big.Int
	Bridge                common.Address
	InitialWasmModuleRoot [32]byte
}

// ExecutionState is an auto generated low-level Go binding around an user-defined struct.
type ExecutionState struct {
	GlobalState   GlobalState
	MachineStatus uint8
}

// GlobalState is an auto generated low-level Go binding around an user-defined struct.
type GlobalState struct {
	Bytes32Vals [2][32]byte
	U64Vals     [2]uint64
}

// Instruction is an auto generated low-level Go binding around an user-defined struct.
type Instruction struct {
	Opcode       uint16
	ArgumentData *big.Int
}

// Machine is an auto generated low-level Go binding around an user-defined struct.
type Machine struct {
	Status          uint8
	ValueStack      ValueStack
	ValueMultiStack MultiStack
	InternalStack   ValueStack
	FrameStack      StackFrameWindow
	FrameMultiStack MultiStack
	GlobalStateHash [32]byte
	ModuleIdx       uint32
	FunctionIdx     uint32
	FunctionPc      uint32
	RecoveryPc      [32]byte
	ModulesRoot     [32]byte
}

// Module is an auto generated low-level Go binding around an user-defined struct.
type Module struct {
	GlobalsMerkleRoot   [32]byte
	ModuleMemory        ModuleMemory
	TablesMerkleRoot    [32]byte
	FunctionsMerkleRoot [32]byte
	ExtraHash           [32]byte
	InternalsOffset     uint32
}

// ModuleMemory is an auto generated low-level Go binding around an user-defined struct.
type ModuleMemory struct {
	Size       uint64
	MaxSize    uint64
	MerkleRoot [32]byte
}

// MultiStack is an auto generated low-level Go binding around an user-defined struct.
type MultiStack struct {
	InactiveStackHash [32]byte
	RemainingHash     [32]byte
}

// StackFrame is an auto generated low-level Go binding around an user-defined struct.
type StackFrame struct {
	ReturnPc              Value
	LocalsMerkleRoot      [32]byte
	CallerModule          uint32
	CallerModuleInternals uint32
}

// StackFrameWindow is an auto generated low-level Go binding around an user-defined struct.
type StackFrameWindow struct {
	Proved        []StackFrame
	RemainingHash [32]byte
}

// Value is an auto generated low-level Go binding around an user-defined struct.
type Value struct {
	ValueType uint8
	Contents  *big.Int
}

// ValueArray is an auto generated low-level Go binding around an user-defined struct.
type ValueArray struct {
	Inner []Value
}

// ValueStack is an auto generated low-level Go binding around an user-defined struct.
type ValueStack struct {
	Proved        ValueArray
	RemainingHash [32]byte
}

// HashProofHelperMetaData contains all meta data concerning the HashProofHelper contract.
var HashProofHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"NotProven\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"}],\"name\":\"PreimagePartProven\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clearSplitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"getPreimagePart\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keccakStates\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"proveWithFullPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"proveWithSplitPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610016576118ef908161001c8239f35b600080fdfe608060408181526004918236101561001657600080fd5b600090813560e01c908163740085d71461038d5750806379754cba1461033e578063ae364ac214610312578063b7465799146102a15763d4e5dd2b1461005b57600080fd5b3461029e578160031936011261029e576001600160401b0391833583811161028f5761008a903690860161046c565b92610093610411565b9061009f368685610604565b95865160208098012095816060941694858211610248575b50508451916100c5836104d3565b6001908184528189850194868652898352828b528883208884528b528883209051151560ff8019835416911617815501935180519384116102355790849392916101118a979654610499565b8b601f82116101fc575b50508a91601f8511600114610182579361016e95938193829360008051602061189a833981519152999794610177575b50501b916000199060031b1c19161790555b85519182918983528983019061042c565b0390a351908152f35b015192508d8061014b565b9190601f9493941984168684528c8420935b8d8282106101e55750509160008051602061189a833981519152979593918561016e989694106101cc575b505050811b01905561015d565b015160001960f88460031b161c191690558b80806101bf565b8484015186558d9a50948701949384019301610194565b86845280842061022492601f880160051c820192881061022b575b601f0160051c019061067a565b8b8b61011b565b9091508190610217565b634e487b7160e01b825260418b52602482fd5b90919350610256858361064a565b888111610297575b610268908661066d565b9182861161029357821161028f57906102879185803693039101610604565b9138806100b7565b8280fd5b8380fd5b508761025e565b80fd5b50913461028f57602036600319011261028f57356001600160a01b0381169081900361028f5791819281526001602052209060018060401b038254169060096102ec6001850161055e565b93015461030a8251948594855260606020860152606085019061042c565b918301520390f35b50903461033a578160031936011261033a576103379033835260016020528220610691565b80f35b5080fd5b503461029e57606036600319011261029e578235906001600160401b03821161029e57506103746020936103869236910161046c565b61037c610411565b906044359261074e565b9051908152f35b9190503461029e578260031936011261029e578335836103ab610411565b92828152806020528181209360018060401b031693848252602052209460ff865416156103f9576103f5856103e26001890161055e565b905191829160208352602083019061042c565b0390f35b6309cb23c960e11b8452830152602482015260449150fd5b602435906001600160401b038216820361042757565b600080fd5b919082519283825260005b848110610458575050826000602080949584010152601f8019910116010190565b602081830181015184830182015201610437565b9181601f84011215610427578235916001600160401b038311610427576020838186019501011161042757565b90600182811c921680156104c9575b60208310146104b357565b634e487b7160e01b600052602260045260246000fd5b91607f16916104a8565b604081019081106001600160401b038211176104ee57604052565b634e487b7160e01b600052604160045260246000fd5b61032081019081106001600160401b038211176104ee57604052565b60a081019081106001600160401b038211176104ee57604052565b601f909101601f19168101906001600160401b038211908210176104ee57604052565b906040519182600082549261057284610499565b9081845260019485811690816000146105e1575060011461059e575b505061059c9250038361053b565b565b9093915060005260209081600020936000915b8183106105c957505061059c9350820101388061058e565b855488840185015294850194879450918301916105b1565b91505061059c94506020925060ff191682840152151560051b820101388061058e565b9192916001600160401b0382116104ee576040519161062d601f8201601f19166020018461053b565b829481845281830111610427578281602093846000960137010152565b9190820391821161065757565b634e487b7160e01b600052601160045260246000fd5b9190820180921161065757565b818110610685575050565b6000815560010161067a565b60026106c160009283815583600182016106ab8154610499565b806106c4575b505050600981019283910161067a565b55565b82601f82116001146106dc575050555b8338806106b1565b90918082526106fa601f60208420940160051c84016001850161067a565b55556106d4565b60001981146106575760010190565b9082101561071c570190565b634e487b7160e01b600052603260045260246000fd5b919091601983101561071c576018908360021c019260031b1690565b919092936000916002861661184f575b6001861615801590611844575b1561180b57336000526001602052604060002095600987015480156000146117c45787546001600160401b0319166001600160401b0384161788555b8587966107b88960098c015461066d565b60098b01555b8715806117b9575b6117ac5760005b608881106116e857506040516107e281610504565b61032036823760005b8b601982106116a857505061032060405161080581610504565b36903760405161081481610520565b60a036823760405161082581610520565b60a03682376040519061083782610504565b6103203683376040516001600160401b036103008201908111908211176104ee576103008101604090815260018252618082602083015267800000000000808a90820152678000000080008000606082015261808b6080820152638000000160a0820181905267800000008000808160c0830181905267800000000000800960e0840152608a61010084015260886101208401526380008009610140840152638000000a610160840152638000808b610180840152608b6001603f1b016101a08401526780000000000080896101c08401526780000000000080036101e084015267800000000000800261020084015260806001603f1b0161022084015261800a61024084015267800000008000000a6102608401526102808301526780000000000080806102a08301526102c08201526780000000800080086102e082015260005b60188110610e4857505050505060005b60198110610e08575050608888106109b25787608811610427576088019660871901966107be565b50919395509193955b6001600160401b0394851692602084018085116106575781811180610dfb575b610cb7575b505050505060011615610cae5760005b60208110610c2a57506040516001850191610a0a826104d3565b60018252610a178361055e565b956020830196875285600052600060205260406000208282541660005260205260016040600020935115159360ff199460ff8683541691161781550196519687518381116104ee57610a698254610499565b601f8111610bf8575b506020601f8211600114610b8c578190899a60009a98999a92610b81575b50508160011b916000199060031b1c19161790555b54169260405191602083526000918054610abe81610499565b92836020870152600182169182600014610b53575050600114610b09575b5050908060008051602061189a833981519152920390a333600052600160205261059c6040600020610691565b6000908152602081209092505b818310610b3957505081016040018160008051602061189a833981519152610adc565b805460408585010152879450602090920191600101610b16565b8695506040935060008051602061189a833981519152969492501682840152151560051b8201019192610adc565b015190503880610a90565b601f198216998360005260206000209a60005b818110610be05750918a9b9184600195949c9a9b9c10610bc7575b505050811b019055610aa5565b015160001960f88460031b161c19169055388080610bba565b838301518d556001909c019b60209384019301610b9f565b610c2490836000526020600020601f840160051c8101916020851061022b57601f0160051c019061067a565b38610a72565b9260039084821c6005808206918282029282840403610657578592610c5d92610c5492049061066d565b60028901610732565b905490841b1c169185901b603881166008600788168015908304821417156106575787830414871517156106575760f891820391821161065757610ca99360ff911c16901b1793610701565b6109f0565b50915050600090565b600094828111610de2575b5090610ccd9161064a565b92818411610dd9575b929160018901935b838110156109e057610cf1818484610710565b3590855491610cff83610499565b90600160401b8210156104ee57601f9384831115610d7c576002610d269101808a55610499565b938483101561071c576020610d5f9510600014610d64578892811690035b60ff83549160031b9260f81c831b921b19161790555b610701565b610cde565b886000528060206000208460051c0193169003610d44565b610d5f946001600160f81b031990921660001a929091601f8214610db857906002910360031b9260ff908116841b931b19910116178655610701565b5050908760005260206000209160ff19169060ff1617905560418655610701565b92508092610cd6565b82610ccd93929650610df39161064a565b949091610cc2565b5060098a015485106109db565b989a91999198610e3d90610d5a8b610e378360026001600160401b03610e2e838a611888565b51169301610732565b9061186a565b9a989a99919961098a565b9c9e9c600190969e968651602088015118604088015118606088015118608088015118865260a087015160c08801511860e088015118610100880151186101208801511880602088015261014088015161016089015118610180890151186101a0890151186101c08901511860408801526101e08801516102008901511861022089015118610240890151186102608901511860608801526102808801516102a0890151186102c0890151186102e089015118610300890151189081608089015280603f1c90848060401b0390851b1617188085528651604088015180603f1c90858060401b0390861b16171860208601526020870151606088015180603f1c90858060401b0390861b16171860408601526040870151608088015180603f1c90858060401b0390861b16171860608601526060870151875180603f1c90858060401b0390861b16171860808601528751188088526020880151855118602089015260408801518551186040890152606088015185511860608901526080880151855118608089015260a088015160208601511860a089015260c088015160208601511860c089015260e088015160208601511860e08901526101008801516020860151186101008901526101208801516020860151186101208901526101408801516040860151186101408901526101608801516040860151186101608901526101808801516040860151186101808901526101a08801516040860151186101a08901526101c08801516040860151186101c08901526101e08801516060860151186101e08901526102008801516060860151186102008901526102208801516060860151186102208901526102408801516060860151186102408901526102608801516060860151186102608901526102808801516080860151186102808901526102a08801516080860151186102a08901526102c08801516080860151186102c08901526102e08801516080860151186102e0890152610300880151608086015118610300890152808652602088015180601c1c90848060401b039060241b1617610100870152604088015180603d1c90848060401b039060031b161761016087015260608801518060171c90848060401b039060291b1617610260870152608088015180602e1c90848060401b039060121b16176102c087015260a088015180603f1c90848060401b0390851b1617604087015260c0880151908160141c848060401b0383602c1b161760a088015260e08901518060361c90858060401b0390600a1b16176101a08801526101008901518060131c90858060401b0390602d1b1617610200880152610120890151603e9080821c90868060401b039060021b16176103008901526101408a0151908160021c91868060401b03911b1617608088015261016089015180603a1c90858060401b039060061b161760e0880152610180890151918260151c858060401b0384602b1b16176101408901526101a08a01518060311c90868060401b0390600f1b16176102408901526101c08a01518060031c90868060401b0390603d1b16176102a08901526101e08a01518060241c90868060401b0390601c1b161760208901526102008a01518060091c90868060401b039060371b16176101208901526102208a01518060271c90868060401b039060191b16176101808901526102408a015180602b1c90868060401b039060151b16176101e08901526102608a01518060081c90868060401b039060381b16176102e08901526102808a01518060251c90868060401b0390601b1b161760608901526102a08a015180602c1c90868060401b039060141b161760c08901526102c08a01518060191c90868060401b039060271b16176101c08901526102e08a01518060381c90868060401b039060081b16176102208901526103008a01518060321c90868060401b0390600e1b16176102808901528260151c858060401b0384602b1b16178160141c868060401b0383602c1b1617191682188a52602088015160c0890151196101608a0151161860208b0152604088015160e0890151196101808a0151161860408b01526060880151610100890151196101a08a0151161860608b01526080880151610120890151196101c08a0151161860808b015260a0880151610140890151196101e08a0151161860a08b015260c0880151610160890151196102008a0151161860c08b015260e0880151610180890151196102208a0151161860e08b01526101008801516101a0890151196102408a015116186101008b01526101208801516101c0890151196102608a015116186101208b01526101408801516101e0890151196102808a015116186101408b0152610160880151610200890151196102a08a015116186101608b0152610180880151610220890151196102c08a015116186101808b01526101a0880151610240890151196102e08a015116186101a08b01526101c0880151610260890151196103008a015116186101c08b01526101e088015161028089015119895116186101e08b01526102008801516102a08901511960208a015116186102008b01526102208801516102c08901511960408a015116186102208b01526102408801516102e08901511960608a015116186102408b01526102608801516103008901511960808a015116186102608b015261028088015188511960a08a015116186102808b01526102a088015160208901511960c08a015116186102a08b01526102c088015160408901511960e08a015116186102c08b01526102e08801516060890151196101008a015116186102e08b01526103008801516080890151196101208a015116186103008b01528360051b860151928060151c90868060401b0390602b1b1617908060141c90868060401b0390602c1b1617191618188752019e9c9e9d959d61097a565b906116bd8160026116df949e969d9e01610732565b905460039190911b1c6001600160401b03166116d98285611888565b52610701565b999891996107eb565b9798909760008a8210156117885750611702818b8b610710565b3560f81c905b8060031c60058082069081810291818304036106575761172992049061066d565b600782168060388460031b1604600814901517156106575761177f9260ff6117588f936002610d5a9501610732565b91909260018060401b039182911660388760031b161b169083548360031b1c16189161186a565b989790986107cd565b908a81146117a3575b60878103611708579060801790611708565b60019150611791565b50919395509193956109bb565b5060018416156107c6565b87546001600160401b038481169116146107a75760405162461bcd60e51b815260206004820152600b60248201526a1112519197d3d19194d15560aa1b6044820152606490fd5b60405162461bcd60e51b81526020600482015260116024820152701393d517d09313d0d2d7d0531251d39151607a1b6044820152606490fd5b50608885061561076b565b3360005260016020526118656040600020610691565b61075e565b919082549060031b9160018060401b03809116831b921b1916179055565b90601981101561071c5760051b019056fef88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802ca2646970667358221220d0ef62077a0d8a5f1db61457f73c6e110865ba63f5b6bce1774c1b5a7a0da09664736f6c63430008130033",
}

// HashProofHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use HashProofHelperMetaData.ABI instead.
var HashProofHelperABI = HashProofHelperMetaData.ABI

// HashProofHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashProofHelperMetaData.Bin instead.
var HashProofHelperBin = HashProofHelperMetaData.Bin

// DeployHashProofHelper deploys a new Ethereum contract, binding an instance of HashProofHelper to it.
func DeployHashProofHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashProofHelper, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashProofHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// HashProofHelper is an auto generated Go binding around an Ethereum contract.
type HashProofHelper struct {
	HashProofHelperCaller     // Read-only binding to the contract
	HashProofHelperTransactor // Write-only binding to the contract
	HashProofHelperFilterer   // Log filterer for contract events
}

// HashProofHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashProofHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashProofHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashProofHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashProofHelperSession struct {
	Contract     *HashProofHelper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashProofHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashProofHelperCallerSession struct {
	Contract *HashProofHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// HashProofHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashProofHelperTransactorSession struct {
	Contract     *HashProofHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// HashProofHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashProofHelperRaw struct {
	Contract *HashProofHelper // Generic contract binding to access the raw methods on
}

// HashProofHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashProofHelperCallerRaw struct {
	Contract *HashProofHelperCaller // Generic read-only contract binding to access the raw methods on
}

// HashProofHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashProofHelperTransactorRaw struct {
	Contract *HashProofHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashProofHelper creates a new instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelper(address common.Address, backend bind.ContractBackend) (*HashProofHelper, error) {
	contract, err := bindHashProofHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// NewHashProofHelperCaller creates a new read-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperCaller(address common.Address, caller bind.ContractCaller) (*HashProofHelperCaller, error) {
	contract, err := bindHashProofHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperCaller{contract: contract}, nil
}

// NewHashProofHelperTransactor creates a new write-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*HashProofHelperTransactor, error) {
	contract, err := bindHashProofHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperTransactor{contract: contract}, nil
}

// NewHashProofHelperFilterer creates a new log filterer instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*HashProofHelperFilterer, error) {
	contract, err := bindHashProofHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperFilterer{contract: contract}, nil
}

// bindHashProofHelper binds a generic wrapper to an already deployed contract.
func bindHashProofHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.HashProofHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transact(opts, method, params...)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCaller) GetPreimagePart(opts *bind.CallOpts, fullHash [32]byte, offset uint64) ([]byte, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "getPreimagePart", fullHash, offset)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCallerSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCaller) KeccakStates(opts *bind.CallOpts, arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "keccakStates", arg0)

	outstruct := new(struct {
		Offset uint64
		Part   []byte
		Length *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Offset = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Part = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Length = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCallerSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactor) ClearSplitProof(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "clearSplitProof")
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactorSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithFullPreimage(opts *bind.TransactOpts, data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithFullPreimage", data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithSplitPreimage(opts *bind.TransactOpts, data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithSplitPreimage", data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// HashProofHelperPreimagePartProvenIterator is returned from FilterPreimagePartProven and is used to iterate over the raw logs and unpacked data for PreimagePartProven events raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProvenIterator struct {
	Event *HashProofHelperPreimagePartProven // Event containing the contract specifics and raw log

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
func (it *HashProofHelperPreimagePartProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashProofHelperPreimagePartProven)
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
		it.Event = new(HashProofHelperPreimagePartProven)
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
func (it *HashProofHelperPreimagePartProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashProofHelperPreimagePartProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashProofHelperPreimagePartProven represents a PreimagePartProven event raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProven struct {
	FullHash [32]byte
	Offset   uint64
	Part     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreimagePartProven is a free log retrieval operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) FilterPreimagePartProven(opts *bind.FilterOpts, fullHash [][32]byte, offset []uint64) (*HashProofHelperPreimagePartProvenIterator, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.FilterLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperPreimagePartProvenIterator{contract: _HashProofHelper.contract, event: "PreimagePartProven", logs: logs, sub: sub}, nil
}

// WatchPreimagePartProven is a free log subscription operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) WatchPreimagePartProven(opts *bind.WatchOpts, sink chan<- *HashProofHelperPreimagePartProven, fullHash [][32]byte, offset []uint64) (event.Subscription, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.WatchLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashProofHelperPreimagePartProven)
				if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
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

// ParsePreimagePartProven is a log parse operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) ParsePreimagePartProven(log types.Log) (*HashProofHelperPreimagePartProven, error) {
	event := new(HashProofHelperPreimagePartProven)
	if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOneStepProofEntryMetaData contains all meta data concerning the IOneStepProofEntry contract.
var IOneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"}],\"internalType\":\"structExecutionState\",\"name\":\"execState\",\"type\":\"tuple\"}],\"name\":\"getMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getStartMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProofEntryMetaData.ABI instead.
var IOneStepProofEntryABI = IOneStepProofEntryMetaData.ABI

// IOneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type IOneStepProofEntry struct {
	IOneStepProofEntryCaller     // Read-only binding to the contract
	IOneStepProofEntryTransactor // Write-only binding to the contract
	IOneStepProofEntryFilterer   // Log filterer for contract events
}

// IOneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProofEntrySession struct {
	Contract     *IOneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IOneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProofEntryCallerSession struct {
	Contract *IOneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IOneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProofEntryTransactorSession struct {
	Contract     *IOneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IOneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProofEntryRaw struct {
	Contract *IOneStepProofEntry // Generic contract binding to access the raw methods on
}

// IOneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCallerRaw struct {
	Contract *IOneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactorRaw struct {
	Contract *IOneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProofEntry creates a new instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*IOneStepProofEntry, error) {
	contract, err := bindIOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntry{IOneStepProofEntryCaller: IOneStepProofEntryCaller{contract: contract}, IOneStepProofEntryTransactor: IOneStepProofEntryTransactor{contract: contract}, IOneStepProofEntryFilterer: IOneStepProofEntryFilterer{contract: contract}}, nil
}

// NewIOneStepProofEntryCaller creates a new read-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProofEntryCaller, error) {
	contract, err := bindIOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryCaller{contract: contract}, nil
}

// NewIOneStepProofEntryTransactor creates a new write-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProofEntryTransactor, error) {
	contract, err := bindIOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryTransactor{contract: contract}, nil
}

// NewIOneStepProofEntryFilterer creates a new log filterer instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProofEntryFilterer, error) {
	contract, err := bindIOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryFilterer{contract: contract}, nil
}

// bindIOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindIOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) GetMachineHash(opts *bind.CallOpts, execState ExecutionState) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "getMachineHash", execState)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntrySession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetMachineHash(&_IOneStepProofEntry.CallOpts, execState)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetMachineHash(&_IOneStepProofEntry.CallOpts, execState)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) GetStartMachineHash(opts *bind.CallOpts, globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "getStartMachineHash", globalStateHash, wasmModuleRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntrySession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetStartMachineHash(&_IOneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetStartMachineHash(&_IOneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// IOneStepProverMetaData contains all meta data concerning the IOneStepProver contract.
var IOneStepProverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"instruction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"result\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"resultMod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProverABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProverMetaData.ABI instead.
var IOneStepProverABI = IOneStepProverMetaData.ABI

// IOneStepProver is an auto generated Go binding around an Ethereum contract.
type IOneStepProver struct {
	IOneStepProverCaller     // Read-only binding to the contract
	IOneStepProverTransactor // Write-only binding to the contract
	IOneStepProverFilterer   // Log filterer for contract events
}

// IOneStepProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProverSession struct {
	Contract     *IOneStepProver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOneStepProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProverCallerSession struct {
	Contract *IOneStepProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IOneStepProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProverTransactorSession struct {
	Contract     *IOneStepProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IOneStepProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProverRaw struct {
	Contract *IOneStepProver // Generic contract binding to access the raw methods on
}

// IOneStepProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProverCallerRaw struct {
	Contract *IOneStepProverCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProverTransactorRaw struct {
	Contract *IOneStepProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProver creates a new instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProver(address common.Address, backend bind.ContractBackend) (*IOneStepProver, error) {
	contract, err := bindIOneStepProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProver{IOneStepProverCaller: IOneStepProverCaller{contract: contract}, IOneStepProverTransactor: IOneStepProverTransactor{contract: contract}, IOneStepProverFilterer: IOneStepProverFilterer{contract: contract}}, nil
}

// NewIOneStepProverCaller creates a new read-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProverCaller, error) {
	contract, err := bindIOneStepProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverCaller{contract: contract}, nil
}

// NewIOneStepProverTransactor creates a new write-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProverTransactor, error) {
	contract, err := bindIOneStepProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverTransactor{contract: contract}, nil
}

// NewIOneStepProverFilterer creates a new log filterer instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProverFilterer, error) {
	contract, err := bindIOneStepProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverFilterer{contract: contract}, nil
}

// bindIOneStepProver binds a generic wrapper to an already deployed contract.
func bindIOneStepProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.IOneStepProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	var out []interface{}
	err := _IOneStepProver.contract.Call(opts, &out, "executeOneStep", execCtx, mach, mod, instruction, proof)

	outstruct := new(struct {
		Result    Machine
		ResultMod Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.ResultMod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCallerSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// OneStepProofEntryMetaData contains all meta data concerning the OneStepProofEntry contract.
var OneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"prover0_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMem_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMath_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverHostIo_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState\",\"name\":\"globalState\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus\",\"name\":\"machineStatus\",\"type\":\"uint8\"}],\"internalType\":\"structExecutionState\",\"name\":\"execState\",\"type\":\"tuple\"}],\"name\":\"getMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getStartMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prover0\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverHostIo\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMath\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMem\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608034620000c657601f62002af738819003918201601f19168301916001600160401b03831184841017620000cb57808492608094604052833981010312620000c6576200004d81620000e1565b906200005c60208201620000e1565b6200007860606200007060408501620000e1565b9301620000e1565b9060018060a01b03928380928160018060a01b03199716876000541617600055168560015416176001551683600254161760025516906003541617600355604051612a009081620000f78239f35b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b0382168203620000c65756fe6080604052600436101561001257600080fd5b60003560e01c806304997be4146100875780631f128bc01461008257806330a5509f1461007d5780635f52fd7c1461007857806366e5d9c314610073578063b5112fd21461006e5763c39619c41461006957600080fd5b6101d1565b610159565b610130565b610107565b6100de565b6100b5565b346100b05760403660031901126100b05760206100a8602435600435610437565b604051908152f35b600080fd5b346100b05760003660031901126100b0576001546040516001600160a01b039091168152602090f35b346100b05760003660031901126100b0576000546040516001600160a01b039091168152602090f35b346100b05760003660031901126100b0576003546040516001600160a01b039091168152602090f35b346100b05760003660031901126100b0576002546040516001600160a01b039091168152602090f35b346100b057366003190160c081126100b0576060136100b0576001600160401b0360a4358181116100b057366023820112156100b05780600401359182116100b05736602483830101116100b0576101cd9160246101bd9201608435606435611093565b6040519081529081906020820190565b0390f35b346100b05760a03660031901126100b05760206100a8610648565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b0382111761021d57604052565b6101ec565b602081019081106001600160401b0382111761021d57604052565b606081019081106001600160401b0382111761021d57604052565b60c081019081106001600160401b0382111761021d57604052565b608081019081106001600160401b0382111761021d57604052565b601f909101601f19168101906001600160401b0382119082101761021d57604052565b604051906102be82610222565b565b604051906102be82610202565b6040519061018082016001600160401b0381118382101761021d57604052565b6001600160401b03811161021d5760051b60200190565b6040519061031182610202565b60006020838281520152565b6040519061032a82610273565b600382528160005b6060811061033e575050565b602090610349610304565b82828501015201610332565b634e487b7160e01b600052603260045260246000fd5b8051156103785760200190565b610355565b8051600110156103785760400190565b8051600210156103785760600190565b80518210156103785760209160051b010190565b604051906103be82610222565b60608252565b604051906103d182610202565b60006020836040516103e281610222565b6060815281520152565b604051906103f982610202565b6000602083606081520152565b634e487b7160e01b600052602160045260246000fd5b6003111561042657565b610406565b60038210156104265752565b6105319161044361031d565b9161044c6128fc565b6104558461036b565b5261045f8361036b565b5061046861291f565b6104718461037d565b5261047b8361037d565b5061048461291f565b61048d8461038d565b526104978361038d565b506104a06102b1565b9283526104ab6102c0565b928352600060208401526104bd6103c4565b6104c56103ec565b6104cd610304565b600019815260006020820152916104e26102cd565b956000875260208701528260408701526060860152608085015260a084015260c0830152600060e083015260006101008301526000610120830152600019610140830152610160820152611fab565b90565b600311156100b057565b60843561053181610534565b6001600160401b038116036100b057565b9060806003198301126100b0576040519161057583610202565b8281602312156100b0576040519161058c83610202565b826044938285116100b0576004905b8582106105f3575050825280606312156100b057604051926105bc84610202565b839060849283116100b057905b8282106105d95750505060200152565b6020809183356105e88161054a565b8152019101906105c9565b813581526020918201910161059b565b6030916f26b0b1b434b7329032b93937b932b21d60811b825260108201520190565b6031917026b0b1b434b732903334b734b9b432b21d60791b825260118201520190565b600161065261053e565b61065b8161041c565b0361069a5761067161066c3661055b565b611f47565b60405161069481610686602082019485610625565b03601f19810183528261028e565b51902090565b60026106a461053e565b6106ad8161041c565b036106d3576106be61066c3661055b565b60405161069481610686602082019485610603565b60405162461bcd60e51b81526020600482015260126024820152714241445f4d414348494e455f53544154555360701b6044820152606490fd5b6040519061018082016001600160401b0381118382101761021d5760405260008083526101608361073c6103c4565b6020820152610749610304565b60408201526107566103c4565b60608201526107636103ec565b6080820152610770610304565b60a08201528260c08201528260e08201528261010082015282610120820152826101408201520152565b604051906107a78261023d565b60006040838281528260208201520152565b604051906107c682610258565b600060a0838281526107d661079a565b60208201528260408201528260608201528260808201520152565b156107f857565b60405162461bcd60e51b815260206004820152601360248201527209a828690929c8abe848a8c9ea48abe9082a69606b1b6044820152606490fd5b6040519061084082610202565b6040368337565b6040519061085482610202565b8160405161086181610202565b6040368237815260206040519161087783610202565b60403684370152565b1561088757565b60405162461bcd60e51b815260206004820152601060248201526f4241445f474c4f42414c5f535441544560801b6044820152606490fd5b634e487b7160e01b600052601160045260246000fd5b90600182018092116108e357565b6108bf565b600e019081600e116108e357565b90600282018092116108e357565b90602082018092116108e357565b1561091957565b60405162461bcd60e51b815260206004820152600c60248201526b1353d115531154d7d493d3d560a21b6044820152606490fd5b1561095457565b60405162461bcd60e51b815260206004820152601260248201527110905117d1955390d51253d394d7d493d3d560721b6044820152606490fd5b909392938483116100b05784116100b0578101920390565b90600163ffffffff809316019182116108e357565b51906102be82610534565b91908260409103126100b0576040516109de81610202565b8092805160078110156100b0578252602090810151910152565b6040929181810384136100b057835191610a1183610202565b8051929485936001600160401b0393908481116100b05783019160209485848403126100b057815193610a4385610222565b80519182116100b0570182601f820112156100b0578051610a63816102ed565b93610a708451958661028e565b818552878086019260061b840101928184116100b0579088809897969594939201915b838310610aaa575050505050815284520151910152565b978495969798610abe8385969794956109c6565b81520192019088979695949392610a93565b91908260409103126100b057604051610ae881610202565b6020808294805184520151910152565b519063ffffffff821682036100b057565b6040929181810384136100b057835191610b2283610202565b8051929485936001600160401b0381116100b057820183601f820112156100b057805190610b4f826102ed565b92610b5c8151948561028e565b828452602095868501918760a0809602850101938285116100b0579088809897969594939201925b848410610b9a5750505050505084520151910152565b90919293948096979850848403126100b0578886918351610bba81610273565b610bc486886109c6565b815284870151838201526060610bdb818901610af8565b86830152610beb60808901610af8565b9082015281520193019190889796959493610b84565b809291039161010083126100b05760405190610c1c82610258565b6060829482518452601f1901126100b057610c9660e060a092604051610c418161023d565b6020820151610c4f8161054a565b81526040820151610c5f8161054a565b60208201526060820151604082015260208601526080810151604086015283810151606086015260c0810151608086015201610af8565b910152565b91906101209081848203126100b05783516001600160401b0392908381116100b0578501906101c0828403126100b057610cd36102cd565b90610cdd836109bb565b825260208301518581116100b05784610cf79185016109f8565b6020830152610d098460408501610ad0565b604083015260808301518581116100b05784610d269185016109f8565b606083015260a08301519485116100b0576101a083610d4c866105319860209701610b09565b6080850152610d5e8660c08301610ad0565b60a08501526101008082015160c0860152610d7a848301610af8565b60e086015261014090610d8e828401610af8565b9086015261016093610da1858401610af8565b90860152610180820151908501520151908201529401610c01565b9060038210156104265752565b80516007811015610426578252602090810151910152565b908151604092838352606083019151936020928382860152855180915283608086019601916000905b828210610e1e575050505081015191015290565b90919296858282610e326001948c51610dc9565b01980193920190610e0a565b906040808201928051918084528251809552606094858501956020809501926000915b838310610e7657505050505081015191015290565b90919293978660a06001928b51610e8e828251610dc9565b80840151828801528681015163ffffffff90811687840152908601511660808201520199019493019190610e61565b908060209392818452848401376000828201840152601f01601f1916010190565b909391949294600435825260243560018060a01b0381168091036100b0576105319661107c9160208501526044356040850152610f266101e080606087015285018851610dbc565b611064602088015193610160610f99610f4e6101c097886102008b01526103a08a0190610de1565b60408c015180516102208b0152602001516102408a0152610f8360608d0151916101df1992838c8303016102608d0152610de1565b9060808d0151908a8303016102808b0152610e3e565b60a08b015180516102a08a0152602001516102c08901529960c08101516102e089015260e081015163ffffffff1661030089015261010081015163ffffffff1661032089015261012081015163ffffffff166103408901526101408101516103608901520151610380870152608086019063ffffffff60a060e0928051855260406020820151600180831b03808251166020890152602082015116828801520151606086015260408101516080860152606081015182860152608081015160c0860152015116910152565b805161ffff16610180850152602001516101a0840152565b818503910152610ebd565b6040513d6000823e3d90fd5b929061109d61070d565b506110a66107b9565b506110af6103b1565b506110b8610304565b506110c38284611a10565b6110da81969296936110d488611fab565b146107f1565b85516110e58161041c565b6110ee8161041c565b61159957506001602b1b90611102906108d5565b14611589576111129082846117cb565b61111d908385611ea4565b9290918160e08701958287516111369063ffffffff1690565b63ffffffff169061114791876124ca565b9561016096878a01511461115a90610912565b6111626103b1565b5061116b6103b1565b50611177908383611640565b611185908484979397611ea4565b611193908385989398611ea4565b91906101208c01978289516111ab9063ffffffff1690565b60061c6303ffffff1663ffffffff166111c3926122ab565b6101008d015163ffffffff1663ffffffff166111de92612435565b6060870151146111ed9061094d565b86516111fe91603f9091169061039d565b51936112099361098e565b9096516112199063ffffffff1690565b63ffffffff1693805161122f9063ffffffff1690565b611238906109a6565b63ffffffff169052815161ffff1697600093908460288b10158061157e575b8015611567575b801561155d575b8015611553575b861461137c57506001546112a2906001600160a01b0316935b60405163a92cb50160e01b81529a8b968795869560048701610ede565b03916001600160a01b03165afa9081156113775761053195600095869361134e575b506180238114908115611342575b501561132c575b50505050600281516112ea8161041c565b6112f38161041c565b148061131b575b15611fab576113088161217a565b611311816121fc565b5060008152611fab565b5061014081015160001914156112fa565b611335926124ca565b90820152388080806112d9565b618024915014386112d2565b90925061136d91953d8091833e611365818361028e565b810190610c9b565b94909491386112c4565b611087565b60458b148015611549575b8015611532575b801561151b575b8015611504575b80156114ed575b80156114d6575b80156114bf575b80156114b5575b80156114a1575b801561148a575b8015611473575b86146113eb57506002546112a2906001600160a01b03165b93611285565b6180108b101580611467575b801561144e575b8015611435575b861461142157506003546112a2906001600160a01b03166113e5565b546112a2906001600160a01b031693611285565b506180308b1015801561140557506180328b1115611405565b506180208b101580156113fe57506180248b11156113fe565b506180138b11156113f7565b5060bc8b101580156113cd575060bf8b11156113cd565b5060c08b101580156113c6575060c48b11156113c6565b5060ac8b14806113bf575060ad8b146113bf565b5060a78b146113b8565b50607c8b101580156113b15750608a8b11156113b1565b5060798b101580156113aa5750607b8b11156113aa565b5060518b101580156113a35750605a8b11156113a3565b50606a8b1015801561139c575060788b111561139c565b5060678b10158015611395575060698b1115611395565b5060468b1015801561138e5750604f8b111561138e565b5060508b14611387565b5060408b1461126c565b50603f8b14611265565b5060368b1015801561125e5750603e8b111561125e565b5060358b1115611257565b5050600282525061053190611fab565b915092916115af926115a9610847565b5061194f565b50906115ba82611f47565b916115cb60c0850193845114610880565b600184516115d88161041c565b6115e18161041c565b149182611628575b508161160e575b506115ff575061053190611fab565b61053191505160443590610437565b60200151516001600160401b0316600435119050386115f0565b159150386115e9565b60001981146108e35760010190565b61165060ff939492948583611724565b93169061165c826102ed565b9061166a604051928361028e565b828252601f19611679846102ed565b0160005b8181106116f1575050819560005b848110611699575050505050565b6116ec906116ab6116b598848661173f565b989098848661178a565b98906116cc6116c26102c0565b61ffff9093168352565b60208201526116db828761039d565b526116e6818661039d565b50611631565b61168b565b6020906116fc610304565b8282870101520161167d565b90821015610378570190565b9015610378573560f81c90600190565b826105319261173592959495611708565b3560f81c92611631565b600093918491905b6002831061175457505050565b9091939461177d6117839161ff0061176d898688611708565b3560f81c9160081b161796611631565b94611631565b9190611747565b600093918491905b6020831061179f57505050565b9091939461177d6117c4916117b5888587611708565b3560f81c9060081b1796611631565b9190611792565b61186263ffffffff9161186c6118746118179561185861182161180461182b9a6117f36107b9565b506117fc61079a565b50848861178a565b61180f9a919a61079a565b5084886118db565b84889392936118db565b84889c929c61178a565b916040519b6118398d61023d565b6001600160401b039182168d521660208c015260408b0152828661178a565b828697929761178a565b828695929561178a565b9190946118a4565b939093966040519661188588610258565b875260208701526040860152606085015260808401521660a082015291565b600093918491905b600483106118b957505050565b9091939461177d6118d49163ffffff0061176d898688611708565b91906118ac565b9192600091825b600890818510156119245761191d9167ffffffffffffff00611917926119098a878b611708565b3560f81c921b161796611631565b93611631565b92946118e2565b95945050509050565b60ff1660ff81146108e35760010190565b9060028110156103785760051b0190565b90929161195a610847565b506040519361196885610202565b6040368637611975610833565b9260005b60ff81169360028510156119ae57906119a36119996119a993868661178a565b919091968a61193e565b5261192d565b611979565b93505060005b60ff81169360028510156119f557906119de6119d46119f09386866118db565b919091968861193e565b6001600160401b03909116905261192d565b6119b4565b9596949350505050611a056102c0565b918252602082015291565b9160ff9192611a1d61070d565b50611a288482611714565b931680611b7257506000925b611a3c6103c4565b50611a456103c4565b50611a4e610304565b50611a576103ec565b50611a60610304565b50611a6c908583611bcd565b611a7a908684979397611d1e565b611a88908385989398611bcd565b611a93908486611dc9565b611aa1908587949394611d1e565b989093611aac6102cd565b98611ab7908a61042b565b602089015260408801526060870152608086015260a085015260c08401906000825260e0850160008152610100860160008152610120870191600083526101408801946000865261016089019860008a5299611b1490868961178a565b9152611b219085886118a4565b63ffffffff909116909152611b379084876118a4565b63ffffffff909116909152611b4d9083866118a4565b63ffffffff909116909152611b6390828561178a565b9252611b6e9261178a565b9252565b60018103611b835750600192611a34565b600203611b9257600292611a34565b60405162461bcd60e51b8152602060048201526013602482015272554e4b4e4f574e5f4d4143485f53544154555360681b6044820152606490fd5b90611be5611bef93611bdd6103c4565b50828461178a565b828495929561178a565b92909294611bfc846102ed565b93611c0a604051958661028e565b808552611c19601f19916102ed565b0160005b818110611c855750506000955b8451871015611c5957611c41611c53918585611c9c565b9790611c4d828861039d565b52611631565b95611c2a565b94925094505060405191611c6c83610222565b825260405191611c7b83610202565b8252602082015291565b602090611c90610304565b82828901015201611c1d565b90611ca5610304565b50611cbe611cb4848385611708565b3560f81c93611631565b9060068411611ce857611cd09261178a565b919060078210156104265760405191611c7b83610202565b60405162461bcd60e51b815260206004820152600e60248201526d4241445f56414c55455f5459504560901b6044820152606490fd5b9190611d37611d3f92611d2f610304565b50828561178a565b91909361178a565b919060405191611c7b83610202565b60405190611d5b82610273565b6000606083611d68610304565b81528260208201528260408201520152565b604051611d8681610222565b6000815290565b60405190611d9a82610202565b600182528160005b60209081811015611dc457602091611db8611d4e565b90828501015201611da2565b505050565b91611dd690611bdd6103ec565b909290806001600160f81b0319611dee828686611708565b351615611e885750611e3a9391611e4c611e0a611e4493611631565b94611e30611e16611d8d565b96611e1f611d4e565b50611e28610304565b508285611c9c565b828599929961178a565b82859692966118a4565b9190936118a4565b9260405196611e5a88610273565b8752602087015263ffffffff809216604087015216606085015292611e7e8361036b565b525b611a056102c0565b915050611e959150611631565b90611e9e611d7a565b90611e80565b909291926060604051611eb681610222565b52611ecf611ec5858385611708565b3560f81c94611631565b90611ed9856102ed565b92611ee7604051948561028e565b858452601f19611ef6876102ed565b0136602086013760005b60ff81169387851015611f2e57906119a3611f1f611f2993868661178a565b919091968861039d565b611f00565b9596505050505060405190611f4282610222565b815291565b8051906020808351930151910151602081519101516040519260208401946c23b637b130b61039ba30ba329d60991b8652602d850152604d84015260018060c01b0319809260c01b16606d84015260c01b166075820152605d815261069481610273565b8051611fb68161041c565b611fbf8161041c565b6120db576040810151610694611fd86020840151612942565b92610686612016611ff761014084019560001997888851141591612779565b9560a08401519061200b60808601516127dc565b908751141591612779565b936120246060840151612942565b9260c08101519161203c60e083015163ffffffff1690565b61010083015163ffffffff169061016061205e61012086015163ffffffff1690565b935194015194604051998a9860208a019c8d97959492909160dc999794926f26b0b1b434b73290393ab73734b7339d60811b8a5260108a015260308901526050880152607087015263ffffffff60e01b9283809260e01b16609088015260e01b16609486015260e01b166098840152609c83015260bc8201520190565b600181516120e88161041c565b6120f18161041c565b0361210f5760c0015160405161069481610686602082019485610625565b6002815161211c8161041c565b6121258161041c565b036121435760c0015160405161069481610686602082019485610603565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b6044820152606490fd5b604081018051519060a08301805151906000198083149081156121f2575b506121e8576020856080606095969701926121b384516127dc565b90515201936121c28551612942565b905152519060208201525251906020820152604051906121e182610222565b6060825252565b5050505060029052565b9050841438612198565b61220d610140820191825190612222565b1561221c576000199052600190565b50600090565b908060601c6122555760e09063ffffffff90818116610120850152818160201c1661010085015260401c16910152600190565b5050600090565b6001600160401b03811161021d57601f01601f191660200190565b6040519061228482610202565b601882527724b739ba393ab1ba34b7b71036b2b935b632903a3932b29d60411b6020830152565b9291926122c86122c36122be86516123d1565b6108e8565b6123e2565b91602094858401946c24b739ba393ab1ba34b7b7399d60991b86526123016122f1835160ff1690565b60f81b6001600160f81b03191690565b93600094851a61231087612414565b538493600e5b84518610156123b1576123a5878b612397848c6123686123ab9761233a8e8e61039d565b51966123606122f161235a6123518b5161ffff1690565b60081c60ff1690565b60ff1690565b901a92612424565b538c6123918d6123806122f161235a895161ffff1690565b612389856108d5565b911a92612424565b536108f6565b9101518c828c010152610904565b95611631565b94612316565b5094509491509550610531949150519020906123cb612277565b92612623565b90816022029160228304036108e357565b906123ec8261225c565b6123f9604051918261028e565b828152809261240a601f199161225c565b0190602036910137565b8051600d101561037857602d0190565b908151811015610378570160200190565b9061053192604051602081019168233ab731ba34b7b71d60b91b83526029820152602981526124638161023d565b519020906040519261247484610202565b6015845274233ab731ba34b7b71036b2b935b632903a3932b29d60591b6020850152612623565b604051906124a882610202565b601382527226b7b23ab6329036b2b935b632903a3932b29d60691b6020830152565b90610531928051906125a460208201518051604060208301519201516040519160208301936626b2b6b7b93c9d60c91b855260018060c01b0319809260c01b16602785015260c01b16602f83015260378201526037815261252a8161023d565b51902061068660408401519360608101519061255260a0608083015192015163ffffffff1690565b9160405196879560208701998a9492909160ab9694926626b7b23ab6329d60c91b87526007870152602786015260478501526067840152608783015263ffffffff60e01b9060e01b1660a78201520190565b519020906123cb61249b565b9392909384519060005b8281106125d257500191825260208201526040019150565b80602080928901015181840152016125ba565b156125ec57565b60405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606490fd5b92939091936000925b84519586518510156126b1579061267e91600197868984161560001461268457610686915061265f612672918a5161039d565b51604051928391602083019589876125b0565b519020965b1c93611631565b9261262c565b6126946126a8916106869361039d565b5192604051928391602083019589876125b0565b51902096612677565b9550919350506102be9150156125e5565b156126c957565b60405162461bcd60e51b81526020600482015260196024820152784d554c5449535441434b5f4e4f535441434b5f41435449564560381b6044820152606490fd5b91606b93916a36bab63a34b9ba30b1b59d60a91b8452600b840152602b830152604b8201520190565b1561273a57565b60405162461bcd60e51b815260206004820152601760248201527626aaa62a24a9aa20a1a5afa727a9aa20a1a5afa6a0a4a760491b6044820152606490fd5b916000199061278a838314156126c2565b156127bb5761279c9083511415612733565b610694602083519301519161068660405193849260208401968761270a565b5090610694602082519201519161068660405193849260208401968761270a565b602080820151926000935b835180518610156128b557906106866128a6612806886128af9561039d565b519261281284516128bd565b878501516040956060878201519101518751928b8401946b29ba30b1b590333930b6b29d60a11b8652602c850152604c84015263ffffffff60e01b918260e091821b16606c8501521b1660708201526054815261286e81610273565b5190209351928391888301958690916052927129ba30b1b590333930b6b29039ba30b1b59d60711b8352601283015260328201520190565b51902094611631565b936127e7565b509350915050565b80519060078210156104265760200151604051906020820192652b30b63ab29d60d11b845260f81b60268301526027820152602781526106948161023d565b612904610304565b5060405161291181610202565b600481526000602082015290565b612927610304565b5060405161293481610202565b600081526000602082015290565b90602091828101519281515151906000925b8284106129615750505050565b909192946129c29061298761298188855161297a610304565b505161039d565b516128bd565b9060405190858201926b2b30b63ab29039ba30b1b59d60a11b8452602c830152604c908183015281526129b981610273565b51902095611631565b92919061295456fea2646970667358221220d71eb2636b83616113449598a7ed42c06f55531812cafb4cddd536f48ae2d41364736f6c63430008130033",
}

// OneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryMetaData.ABI instead.
var OneStepProofEntryABI = OneStepProofEntryMetaData.ABI

// OneStepProofEntryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryMetaData.Bin instead.
var OneStepProofEntryBin = OneStepProofEntryMetaData.Bin

// DeployOneStepProofEntry deploys a new Ethereum contract, binding an instance of OneStepProofEntry to it.
func DeployOneStepProofEntry(auth *bind.TransactOpts, backend bind.ContractBackend, prover0_ common.Address, proverMem_ common.Address, proverMath_ common.Address, proverHostIo_ common.Address) (common.Address, *types.Transaction, *OneStepProofEntry, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryBin), backend, prover0_, proverMem_, proverMath_, proverHostIo_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// OneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntry struct {
	OneStepProofEntryCaller     // Read-only binding to the contract
	OneStepProofEntryTransactor // Write-only binding to the contract
	OneStepProofEntryFilterer   // Log filterer for contract events
}

// OneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntrySession struct {
	Contract     *OneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryCallerSession struct {
	Contract *OneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryTransactorSession struct {
	Contract     *OneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryRaw struct {
	Contract *OneStepProofEntry // Generic contract binding to access the raw methods on
}

// OneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryCallerRaw struct {
	Contract *OneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactorRaw struct {
	Contract *OneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntry creates a new instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*OneStepProofEntry, error) {
	contract, err := bindOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryCaller creates a new read-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryCaller, error) {
	contract, err := bindOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryCaller{contract: contract}, nil
}

// NewOneStepProofEntryTransactor creates a new write-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryTransactor, error) {
	contract, err := bindOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryTransactor{contract: contract}, nil
}

// NewOneStepProofEntryFilterer creates a new log filterer instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryFilterer, error) {
	contract, err := bindOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryFilterer{contract: contract}, nil
}

// bindOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.OneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCaller) GetMachineHash(opts *bind.CallOpts, execState ExecutionState) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "getMachineHash", execState)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntrySession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetMachineHash(&_OneStepProofEntry.CallOpts, execState)
}

// GetMachineHash is a free data retrieval call binding the contract method 0xc39619c4.
//
// Solidity: function getMachineHash(((bytes32[2],uint64[2]),uint8) execState) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) GetMachineHash(execState ExecutionState) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetMachineHash(&_OneStepProofEntry.CallOpts, execState)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCaller) GetStartMachineHash(opts *bind.CallOpts, globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "getStartMachineHash", globalStateHash, wasmModuleRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntrySession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetStartMachineHash(&_OneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetStartMachineHash(&_OneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0xb5112fd2.
//
// Solidity: function proveOneStep((uint256,address,bytes32) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) Prover0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "prover0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverHostIo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverHostIo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMath(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMath")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// OneStepProofEntryLibMetaData contains all meta data concerning the OneStepProofEntryLib contract.
var OneStepProofEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212209bb9b119673aaad5ddef738902253657ee9db5ea420fbe0e9835fc7845f202f264736f6c63430008130033",
}

// OneStepProofEntryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryLibMetaData.ABI instead.
var OneStepProofEntryLibABI = OneStepProofEntryLibMetaData.ABI

// OneStepProofEntryLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryLibMetaData.Bin instead.
var OneStepProofEntryLibBin = OneStepProofEntryLibMetaData.Bin

// DeployOneStepProofEntryLib deploys a new Ethereum contract, binding an instance of OneStepProofEntryLib to it.
func DeployOneStepProofEntryLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProofEntryLib, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// OneStepProofEntryLib is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntryLib struct {
	OneStepProofEntryLibCaller     // Read-only binding to the contract
	OneStepProofEntryLibTransactor // Write-only binding to the contract
	OneStepProofEntryLibFilterer   // Log filterer for contract events
}

// OneStepProofEntryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntryLibSession struct {
	Contract     *OneStepProofEntryLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryLibCallerSession struct {
	Contract *OneStepProofEntryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OneStepProofEntryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryLibTransactorSession struct {
	Contract     *OneStepProofEntryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryLibRaw struct {
	Contract *OneStepProofEntryLib // Generic contract binding to access the raw methods on
}

// OneStepProofEntryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCallerRaw struct {
	Contract *OneStepProofEntryLibCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactorRaw struct {
	Contract *OneStepProofEntryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntryLib creates a new instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLib(address common.Address, backend bind.ContractBackend) (*OneStepProofEntryLib, error) {
	contract, err := bindOneStepProofEntryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryLibCaller creates a new read-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryLibCaller, error) {
	contract, err := bindOneStepProofEntryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibCaller{contract: contract}, nil
}

// NewOneStepProofEntryLibTransactor creates a new write-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryLibTransactor, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibTransactor{contract: contract}, nil
}

// NewOneStepProofEntryLibFilterer creates a new log filterer instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryLibFilterer, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibFilterer{contract: contract}, nil
}

// bindOneStepProofEntryLib binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transact(opts, method, params...)
}

// OneStepProver0MetaData contains all meta data concerning the OneStepProver0 contract.
var OneStepProver0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60808060405234610016576125d8908161001c8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c63a92cb5011461002757600080fd5b346112775736600319016101e0811261127757606013611277576001600160401b0360643511611277576101c0606435360360031901126112775761010036608319011261127757604036610183190112611277576001600160401b036101c43511611277573660236101c435011215611277576001600160401b0360046101c435013511611277573660246101c435600401356101c435010111611277576100d06080611354565b60006080526100dd61141a565b60a0526040516100ec81611370565b6000808252602082015260c05261010161141a565b60e05260405161011081611370565b60608152600060208201526101005260405161012b81611370565b600080825260208201819052610120919091526101408190526101608190526101808190526101a08190526101c08190526101e052610168611461565b5060405161017581611354565b600360643560040135101561127757606435600481013582526001600160401b0360249091013511611277576101b6366064356024810135016004016114e2565b60208201526101ca366044606435016115b7565b60408201526001600160401b036064356084013511611277576101f8366064356084810135016004016114e2565b60608201526001600160401b0360643560a401351161127757604060643560a481013501360360031901126112775760405161023381611370565b60643560a481013501600401356001600160401b038111611277573660643560a481013501820160230112156112775760643560a4810135018101600401359061027c82611499565b9161028a60405193846113f7565b80835260208301913660643560a481013501820160a0840201602401116112775760643560a481013501810160240192905b60643560a481013501810160a0840201602401841061127c57505050508152602460a46064350135606435010135602082015260808201526103033660c4606435016115b7565b60a0820152610104606435013560c0820152610324610124606435016115df565b60e0820152610338610144606435016115df565b61010082015261034d610164606435016115df565b6101208201526064356101848101356101408301526101a401356101608201526040519061037a826113a6565b608435825260603660a319011261127757604051610397816113c1565b60a4356001600160401b03811690036112775760a435815260c4356001600160401b03811690036112775760c435602082015260e4356040820152602083015261010435604083015261012435606083015261014435608083015263ffffffff61016435166101643503611277576101643560a083015261ffff6104196115f0565b16806110d4575060155b806015146110ca5780601414610ffe57806003146110045780600514610ff05780600414610f9b5780601314610f565780601214610f0b5780601114610ec35780601014610e395780600f14610a5d5780600e14610a525780600d14610a385780600c14610a215780600b146109df5780600a146109ac57806009146109665780600814610942578060071461090d57806006146107de5780600214610723576001146104e057634e487b7160e01b600052605160045260246000fd5b6020810151600060206040516104f581611370565b8281520152518051518060001981011161070d5761053c91610531916000602060405161052181611370565b8281520152600019019051611fb9565b51602083015161252b565b604051906101208252805160038110156106f75761012083015261057260208201516101c06101408501526102e08401906112f7565b9060206040820151805161016086015201516101808401526105a860608201519261011f199384868303016101a08701526112f7565b608082015192848203016101c08501528251926040825260408201845180915260206060840195019060005b8181106106b0575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c08601528051602086015260406020820151600180831b0381511682880152600180831b03602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b516106ce8482516112df565b8581015160408501528260408201511682850152015116608082015201970191019190916105d4565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b506000602060405161073481611370565b828152015261800561ffff6107476115f0565b160361076c5761076761075d6020830151612402565b606083015161252b565b61053c565b61800661ffff61077a6115f0565b160361079a576107676107906060830151612402565b602083015161252b565b60405162461bcd60e51b815260206004820152601c60248201527b4d4f56455f494e5445524e414c5f494e56414c49445f4f50434f444560201b6044820152606490fd5b50906107ed6020830151612402565b9163ffffffff6108006020830151612402565b938161082161081b6108156020870151612402565b97612347565b92612347565b916040519661082f8861138b565b87526101a43560208801521660408601521660608401526080810151918251519160018301831161070d5761086660018401611499565b9261087460405194856113f7565b600101808452601f199061088790611499565b0160005b8181106108f657505060005b845180518210156108d157906108b0816108cc93611fb9565b516108bb8287611fb9565b526108c68186611fb9565b50611e21565b610897565b50509290936108ef90825151906108e88286611fb9565b5283611fb9565b505261053c565b602090610901612264565b8282880101520161088b565b5061093b61091e6020830151612402565b83516101c435600401359160246101c43501916101a43590611da9565b825261053c565b5061076761079083516101c435600401359060246101c43501906101a43590611d42565b506109746020820151612402565b6109a5602061098660808501516122dd565b019182516101c435600401359160246101c43501916101a43590611da9565b905261053c565b5061076761079060206109c260808501516122dd565b01516101c435600401359060246101c43501906101a43590611d42565b5063ffffffff6109fa6109f56020840151612402565b612347565b161561053c576101a435610a1663ffffffff821691821461160e565b61012082015261053c565b506101a435610a1663ffffffff821691821461160e565b506107676101c4356004013560246101c435018484611808565b506107678282611724565b50610a6e6109f56020830151612402565b610a76611461565b506060604051610a85816113dc565b52610a8e611461565b50610a97611442565b506000906000805b60208110610dfe5750610aef610b3e610aef610b06610ad8610b5595610ac3611442565b506101c4356004013560246101c43501611e30565b6101c4959195356004013560246101c43501611e30565b6101c4929192356004013560246101c43501611e82565b9160405195610b14876113c1565b6001600160401b03908116875216602086015260408501526101c435600481013590602401611e82565b6101c4949194356004013560246101c43501611e82565b9390916000949360005b60048110610dbb57509163ffffffff9391610bb2969360405199610b828b6113a6565b8a5260208a01526040890152606088015260808701521660a08501526101c4356004013560246101c43501611fcd565b5082519060208401518051604060208301519201516040519160208301936626b2b6b7b93c9d60c91b855260018060c01b0319809260c01b16602785015260c01b16602f830152603782015260378152610c0b816113c1565b5190206040850151606086015160808701519160a088015193604051966626b7b23ab6329d60c91b6020890152602788015260478701526067860152608785015260a784015263ffffffff60e01b9060e01b1660c783015260ab82528160e081011060018060401b0360e084011117610da5578160e0610cd0930160405260e08151602083012091610c9e828201611370565b6013828201527226b7b23ab6329036b2b935b632903a3932b29d60691b610100820152019163ffffffff851690612184565b61016084015103610d555763ffffffff60a0819382610d4194610d006020890151610cfa8a61164a565b9061252b565b610d176020890151610cfa8460e08c0151166123ce565b610d2d6020890151610cfa84878d0151166123ce565b1660e0870152015116826101a4351661170c565b16610100820152600061012082015261053c565b60405162461bcd60e51b815260206004820152602260248201527f43524f53535f4d4f44554c455f494e5445524e414c5f4d4f44554c45535f524f60448201526113d560f21b6064820152608490fd5b634e487b7160e01b600052604160045260246000fd5b95610df0610df69163ffffff00610de0896101c4356004013560246101c43501611dff565b3560f81c9160081b161796611e21565b96611e21565b959495610b5f565b9092610e2e610e3491610e1f866101c4356004013560246101c43501611dff565b3560f81c9060081b1794611e21565b91611e21565b610a9f565b50610e4b6020820151610cfa8361164a565b610e8e610e5b60808301516122dd565b610e766020840151610cfa63ffffffff6040850151166123ce565b610cfa63ffffffff60606020860151930151166123ce565b63ffffffff6101a435610ea48160401c156116c3565b818160201c1660e084015216610100820152600061012082015261053c565b50610ed56020820151610cfa8361164a565b610ef06020820151610cfa63ffffffff60e0850151166123ce565b610e8e6020820151610cfa63ffffffff60a0860151166123ce565b50610f1d6020820151610cfa8361164a565b610f2d610e5b60808301516122dd565b6101a435610f4363ffffffff821691821461160e565b610100820152600061012082015261053c565b506107676080820151610f67612264565b50610f7660018251511461229d565b610f808151611fac565b519060405190610f8f826113dc565b6000825252518261208e565b50610fac6109f56020830151612402565b610fb96020830151612402565b63ffffffff610fcb6020850151612402565b921615610fe1575061076790602083015161252b565b6107679150602083015161252b565b50610ffe6020820151612402565b5061053c565b5061ffff6110106115f0565b1660418103611052575061076760005b602083015161103b6040519261103584611370565b83611602565b6101a4356001600160401b0316602083015261252b565b6042810361106557506107676001611020565b6043810361107857506107676002611020565b604403611089576107676003611020565b60405162461bcd60e51b8152602060048201526019602482015278434f4e53545f505553485f494e56414c49445f4f50434f444560381b6044820152606490fd5b506002815261053c565b600181036110e457506014610423565b600f81036110f457506013610423565b6010810361110457506012610423565b618009810361111557506011610423565b61800b810361112657506010610423565b61800c81036111375750600f610423565b61800a81036111485750600e610423565b601181036111585750600d610423565b61800381036111695750600c610423565b618004810361117a5750600b610423565b6020810361118a5750600a610423565b6021810361119a57506009610423565b602381036111aa57506008610423565b602481036111ba57506007610423565b61800281036111cb57506006610423565b601a81036111db57506005610423565b601b81036111eb57506004610423565b60418110158061126c575b1561120357506003610423565b61800581148015611261575b1561121c57506002610423565b6180080361122b576001610423565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b6044820152606490fd5b50618006811461120f565b5060448111156111f6565b600080fd5b60a0843603126112775760a080602060249460405161129a8161138b565b6112a4368a6114b0565b81526040890135838201526112bb60608a016115df565b60408201526112cc60808a016115df565b60608201528152019501949250506102bc565b805160078110156106f7578252602090810151910152565b908151604092838352606083019151936020928382860152855180915283608086019601916000905b828210611334575050505081015191015290565b909192968582826113486001948c516112df565b01980193920190611320565b61018081019081106001600160401b03821117610da557604052565b604081019081106001600160401b03821117610da557604052565b608081019081106001600160401b03821117610da557604052565b60c081019081106001600160401b03821117610da557604052565b606081019081106001600160401b03821117610da557604052565b602081019081106001600160401b03821117610da557604052565b601f909101601f19168101906001600160401b03821190821017610da557604052565b6040519061142782611370565b6000602083604051611438816113dc565b6060815281520152565b6040519061144f826113c1565b60006040838281528260208201520152565b6040519061146e826113a6565b600060a08382815261147e611442565b60208201528260408201528260608201528260808201520152565b6001600160401b038111610da55760051b60200190565b9190826040910312611277576040516114c881611370565b809280356007811015611277578252602090810135910152565b604092918181038413611277578351916114fb83611370565b829481359260018060401b03938481116112775783019160209485848403126112775781519361152a856113dc565b8035918211611277570182601f8201121561127757803561154a81611499565b93611557845195866113f7565b818552878086019260061b84010192818411611277579088809897969594939201915b838310611591575050505050815284520135910152565b9784959697986115a58385969794956114b0565b8152019201908897969594939261157a565b9190826040910312611277576040516115cf81611370565b6020808294803584520135910152565b359063ffffffff8216820361127757565b6101843561ffff811681036112775790565b60078210156106f75752565b1561161557565b60405162461bcd60e51b815260206004820152600d60248201526c4241445f43414c4c5f4441544160981b6044820152606490fd5b6000602060405161165a81611370565b828152015263ffffffff610120820151169060e0610100820151910151916000602060405161168881611370565b82815201526040519261169a84611370565b6006845263ffffffff60401b9060401b169163ffffffff60201b9060201b161717602082015290565b156116ca57565b60405162461bcd60e51b815260206004820152601a6024820152794241445f43524f53535f4d4f44554c455f43414c4c5f4441544160301b6044820152606490fd5b91909163ffffffff8080941691160191821161070d57565b9061175f91611769602082019261173f8451610cfa8561164a565b610cfa845160a060e086019663ffffffff988993610cfa858b51166123ce565b51930151166123ce565b61177660808201516122dd565b92606084019281845116156117fe576101a4358281169081036117b95782610120956117ae9382604060009a0151169052511661170c565b166101008201520152565b60405162461bcd60e51b815260206004820152601d60248201527f4241445f43414c4c45525f494e5445524e414c5f43414c4c5f444154410000006044820152606490fd5b5050600290525050565b6000936020938483019261181f6109f58551612402565b6040966060885161182f816113dc565b5288978896875b600890818c1015611878576118719167ffffffffffffff0061186b9261185d8d8d8d611dff565b3560f81c921b161799611e21565b9a611e21565b9997611836565b90506118c3949a506118976118cd929893999c969c9a949a8489611e82565b9190946118b96118b26118ab85888d611dff565b3594611e21565b868b611e30565b868b999299611e82565b868b969296611fcd565b938d8d519081016d21b0b6361034b73234b932b1ba1d60911b815260018060c01b031991828760c01b16602e8201528a60368201526036815261190f816113c1565b5190206101a43503611cc6578d928f919261199d9385519084820192652a30b136329d60d11b845260ff60f81b1660268301528d60c01b16602782015289602f820152602f815261195f816113c1565b51902090712a30b136329036b2b935b632903a3932b29d60711b85519361198585611370565b601285528401526001600160401b0396871690612184565b91015103611c905763ffffffff809c169516851015611c7f576119e190868b8b516119c781611370565b828152015260608a516119d9816113dc565b528388611e82565b949092868b8b516119f181611370565b8281520152611a0e611a0487838b611dff565b3560f81c96611e21565b60068711611c4a57611a2190828a611e82565b9190986007881015611c36578c92611acf9492611a5892611a4d8f519b611a478d611370565b8c611602565b858b019c8d52611fcd565b5090611a6388612302565b8c51848101916d2a30b136329032b632b6b2b73a1d60911b835288602e830152604e820152604e8152611a958161138b565b51902091792a30b136329032b632b6b2b73a1036b2b935b632903a3932b29d60311b8d5194611ac386611370565b601a8652850152612184565b03611bfe5703611bf15780516007811015611bdd57600403611af8575050505050505060029052565b51906007821015611bc95750600503611b95575192848416938403611b5a5750916117ae8492610cfa61012096606060009997611b398151610cfa8b61164a565b61175f611b4960808b01516122dd565b95610cfa85845192890151166123ce565b60649083519062461bcd60e51b8252600482015260156024820152744241445f46554e435f5245465f434f4e54454e545360581b6044820152fd5b825162461bcd60e51b815260048101859052600d60248201526c4241445f454c454d5f5459504560981b6044820152606490fd5b634e487b7160e01b81526021600452602490fd5b634e487b7160e01b83526021600452602483fd5b5050505050505060029052565b865162461bcd60e51b815260048101899052601160248201527010905117d153115351539514d7d493d3d5607a1b6044820152606490fd5b634e487b7160e01b89526021600452602489fd5b8a5162461bcd60e51b8152600481018d9052600e60248201526d4241445f56414c55455f5459504560901b6044820152606490fd5b505050505050505050505060029052565b895162461bcd60e51b8152600481018c9052600f60248201526e10905117d51050931154d7d493d3d5608a1b6044820152606490fd5b60648f8f519062461bcd60e51b8252600482015260166024820152754241445f43414c4c5f494e4449524543545f4441544160501b6044820152fd5b15611d0957565b60405162461bcd60e51b815260206004820152601160248201527015d493d391d7d3515492d31157d493d3d5607a1b6044820152606490fd5b90611da692611da09160006020604051611d5b81611370565b828152015260006020604051611d7081611370565b82815201526060604051611d83816113dc565b52611d9a611d918784611ec9565b90978894611fcd565b50612110565b14611d02565b90565b611dfa90611da0611df1611da6979694959660006020604051611dcb81611370565b82815201526060604051611dde816113dc565b52611de98187611ec9565b919096611fcd565b50938585612110565b612110565b90821015611e0b570190565b634e487b7160e01b600052603260045260246000fd5b600019811461070d5760010190565b9192600091825b60089081851015611e7957611e729167ffffffffffffff00611e6c92611e5e8a878b611dff565b3560f81c921b161796611e21565b93611e21565b9294611e37565b95945050509050565b600093918491905b60208310611e9757505050565b90919394611ebc611ec291611ead888587611dff565b3560f81c9060081b1796611e21565b94611e21565b9190611e8a565b60405191611ed683611370565b600092838152836020809201528115611f985760f8918335831c9360068511611f62579290859360019287925b858410611f355750505050936007841015611bc95750611f2f60405193611f2985611370565b84611602565b82015291565b9091929395611ebc611f5a91611f4c898588611dff565b35851c9060081b1797611e21565b929190611f03565b60405162461bcd60e51b815260048101849052600e60248201526d4241445f56414c55455f5459504560901b6044820152606490fd5b634e487b7160e01b84526032600452602484fd5b805115611e0b5760200190565b8051821015611e0b5760209160051b010190565b916060604051611fdc816113dc565b52611ff5611feb828486611dff565b3560f81c91611e21565b611ffe82611499565b9261200c60405194856113f7565b828452601f1961201b84611499565b01366020860137600094855b60ff811693858510156120755761204261204c918585611e82565b9190919588611fb9565b5260ff80911690811461206157600101612027565b634e487b7160e01b87526011600452602487fd5b9650505050505060405190612089826113dc565b815291565b90805160078110156106f7576004146120cf57805160078110156106f7576006036120cf5760206120c1910151826120d6565b156120c95750565b60029052565b5060029052565b908060601c6121095760e09063ffffffff90818116610120850152818160201c1661010085015260401c16910152600190565b5050600090565b9061211d611da693612302565b906040519261212b84611370565b60128452712b30b63ab29036b2b935b632903a3932b29d60711b6020850152612184565b9392909384519060005b82811061217157500191825260208201526040019150565b8060208092890101518184015201612159565b92939091936000925b845195865185101561222057906121ed9160019786898416156000146121f3576121d391506121c06121e1918a51611fb9565b516040519283916020830195898761214f565b03601f1981018352826113f7565b519020965b1c93611e21565b9261218d565b612203612217916121d393611fb9565b51926040519283916020830195898761214f565b519020966121e6565b9550925050915061222d57565b60405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606490fd5b604051906122718261138b565b600060608360405161228281611370565b83815283602082015281528260208201528260408201520152565b156122a457565b60405162461bcd60e51b81526020600482015260116024820152700848288beae929c889eaebe988a9c8ea89607b1b6044820152606490fd5b6122fe906122e9612264565b506122f860018251511461229d565b51611fac565b5190565b80519060078210156106f75760200151604051906020820192652b30b63ab29d60d11b845260f81b6026830152602782015260278152612341816113c1565b51902090565b6020810151905160078110156106f75761239f57600160201b8110156123705763ffffffff1690565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606490fd5b60405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606490fd5b600060206040516123de81611370565b828152015263ffffffff604051916123f583611370565b6000835216602082015290565b9060209160405161241281611370565b60009381858093520152519180602060405161242d81611370565b82815201528251805160001993918482019182116124b6579061244f91611fb9565b51928451519081019081116124a257612467906124ca565b915b825181101561249c5780612481612497928751611fb9565b5161248c8286611fb9565b526108c68185611fb9565b612469565b50925290565b634e487b7160e01b83526011600452602483fd5b634e487b7160e01b84526011600452602484fd5b906124d482611499565b6040906124e3825191826113f7565b83815280936124f4601f1991611499565b0191600090815b848110612509575050505050565b602090825161251781611370565b8481528285818301528287010152016124fb565b518051519160019283810180911161070d57612546906124ca565b926000815b612567575b5050815151612563916108e88286611fb9565b5052565b8351805182101561259c57906125808161259693611fb9565b5161258b8288611fb9565b526108c68187611fb9565b8161254b565b5061255056fea26469706673582212202d8da9f96fbaff104f40d0c54927f58a062a988a75867e882530b405b94b7a1e64736f6c63430008130033",
}

// OneStepProver0ABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProver0MetaData.ABI instead.
var OneStepProver0ABI = OneStepProver0MetaData.ABI

// OneStepProver0Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProver0MetaData.Bin instead.
var OneStepProver0Bin = OneStepProver0MetaData.Bin

// DeployOneStepProver0 deploys a new Ethereum contract, binding an instance of OneStepProver0 to it.
func DeployOneStepProver0(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProver0, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProver0Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// OneStepProver0 is an auto generated Go binding around an Ethereum contract.
type OneStepProver0 struct {
	OneStepProver0Caller     // Read-only binding to the contract
	OneStepProver0Transactor // Write-only binding to the contract
	OneStepProver0Filterer   // Log filterer for contract events
}

// OneStepProver0Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProver0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProver0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProver0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProver0Session struct {
	Contract     *OneStepProver0   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProver0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProver0CallerSession struct {
	Contract *OneStepProver0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OneStepProver0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProver0TransactorSession struct {
	Contract     *OneStepProver0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OneStepProver0Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProver0Raw struct {
	Contract *OneStepProver0 // Generic contract binding to access the raw methods on
}

// OneStepProver0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProver0CallerRaw struct {
	Contract *OneStepProver0Caller // Generic read-only contract binding to access the raw methods on
}

// OneStepProver0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProver0TransactorRaw struct {
	Contract *OneStepProver0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProver0 creates a new instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0(address common.Address, backend bind.ContractBackend) (*OneStepProver0, error) {
	contract, err := bindOneStepProver0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// NewOneStepProver0Caller creates a new read-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Caller(address common.Address, caller bind.ContractCaller) (*OneStepProver0Caller, error) {
	contract, err := bindOneStepProver0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Caller{contract: contract}, nil
}

// NewOneStepProver0Transactor creates a new write-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Transactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProver0Transactor, error) {
	contract, err := bindOneStepProver0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Transactor{contract: contract}, nil
}

// NewOneStepProver0Filterer creates a new log filterer instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Filterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProver0Filterer, error) {
	contract, err := bindOneStepProver0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Filterer{contract: contract}, nil
}

// bindOneStepProver0 binds a generic wrapper to an already deployed contract.
func bindOneStepProver0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.OneStepProver0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Caller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProver0.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Session) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0CallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverHostIoMetaData contains all meta data concerning the OneStepProverHostIo contract.
var OneStepProverHostIoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BLOBSTREAM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CELESTIA_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657614165908161001c8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c9081635fd9e56d14610ebf57508063a92cb501146100745763dc3465471461004057600080fd5b3461006f57600036600319011261006f57602060405173c3e209eb245fd59c8586777b499d6a665df3abd28152f35b600080fd5b3461006f5736600319016101e0811261006f5760601361006f576001600160401b036064351161006f576101c06064353603600319011261006f5761010036608319011261006f5760403661018319011261006f576001600160401b036101c4351161006f573660236101c43501121561006f576001600160401b0360046101c43501351161006f573660246101c435600401356101c43501011161006f57600061016060405161012481610fa7565b82815261012f611052565b602082015260405161014081610f71565b8381528360208201526040820152610156611052565b606082015260405161016781610f71565b60608152836020820152608082015260405161018281610f71565b83815283602082015260a08201528260c08201528260e082015282610100820152826101208201528261014082015201526101bb611099565b506040516101c881610fa7565b600360643560040135101561006f57606435600481013582526001600160401b036024909101351161006f576102093660643560248101350160040161111a565b602082015261021d366044606435016111ef565b60408201526001600160401b03606435608401351161006f5761024b3660643560848101350160040161111a565b60608201526001600160401b0360643560a401351161006f57604060643560a4810135013603600319011261006f5760405161028681610f71565b60643560a481013501600401356001600160401b03811161006f573660643560a4810135018201602301121561006f5760643560a481013501810160040135906102cf826110d1565b916102dd604051938461102f565b80835260208301913660643560a481013501820160a08402016024011161006f5760643560a481013501810160240192905b60643560a481013501810160a08402016024018410610e3057505050508152602460a46064350135606435010135602082015260808201526103563660c4606435016111ef565b60a0820152610104606435013560c082015261037761012460643501611217565b60e082015261038b61014460643501611217565b6101008201526103a061016460643501611217565b6101208201526064356101848101356101408301526101a40135610160820152604051906103cd82610fde565b608435825260603660a319011261006f576040516103ea81610fc3565b60a4356001600160401b038116900361006f5760a435815260c4356001600160401b038116900361006f5760c435602082015260e4356040820152602083015261010435604083015261012435606083015261014435608083015263ffffffff6101643516610164350361006f576101643560a08301526101843561ffff8116810361006f5761801061ffff8216101580610e20575b15610d435760095b9081600814610d285781600714610d0d5781600614610d025781600514610ce757816004146108ee57816009146106d35750806003146106c957806002146106ab576001146104e757634e487b7160e01b600052605160045260246000fd5b6104f0816136d6565b604051906101208252805161050481610edc565b61012083015261052660208201516101c06101408501526102e0840190610f14565b90602060408201518051610160860152015161018084015261055c60608201519261011f199384868303016101a0870152610f14565b608082015192848203016101c08501528251926040825260408201845180915260206060840195019060005b818110610664575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c08601528051602086015260406020820151600180831b0381511682880152600180831b03602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b51610682848251610efc565b858101516040850152826040820151168285015201511660808201520197019101919091610588565b506106c46101c4356004013560246101c4350183613566565b6104f0565b506106c481613539565b90506106dd6132ff565b506106e66132ff565b506000604051916106f683610f71565b604036843760405161070781610f71565b60403682376000925b600260ff851610156107525761073761074c916101c4356004013560246101c4350161380b565b949061074660ff831688613338565b5261383d565b92610710565b9250929390936000925b600260ff851610156107aa576107836107a4916101c4356004013560246101c43501613789565b949061079260ff831688613338565b6001600160401b03909116905261383d565b9261075c565b925092939093604051936107bd85610f71565b845260208401526107cd836138ce565b60c0850151036108b65761801061ffff82161480156108a7575b1561082d57509061081e610815610823936101c43560040135906101c4356004013560246101c435016112c9565b90838787613349565b6138ce565b60c08201526104f0565b905061ffff81166180120361084b57508061081e6108239284613473565b61ffff1661801303610865578061081e61082392846134ec565b60405162461bcd60e51b815260206004820152601a602482015279494e56414c49445f474c4f42414c53544154455f4f50434f444560301b6044820152606490fd5b5061801161ffff8216146107e7565b60405162461bcd60e51b815260206004820152601060248201526f4241445f474c4f42414c5f535441544560801b6044820152606490fd5b50506108f8612ebe565b60405161090481610f8c565b6060905260405161091481610f8c565b60609052610920612ebe565b9061016083015191610930611099565b50610939611099565b5061094261107a565b50826109596101c4356004810135906024016137d5565b61096495919561107a565b5061097b906101c435600481013590602401613789565b610991906101c435600481013590602401613789565b6109a7906101c43560048101359060240161380b565b91604051936109b585610fc3565b6001600160401b03908116855216602084015260408301526109e3906101c43560048101359060240161380b565b9091906109fc906101c43560048101359060240161380b565b610a12906101c43560048101359060240161380b565b909190610a2b906101c43560048101359060240161373c565b9490936040519a610a3b8c610fde565b8b5260208b015260408a01526060890152608088015263ffffffff1660a0870152610a72906101c43560048101359060240161373c565b909390610a8b906101c43560048101359060240161384e565b93909680519060208101518051602082015191604001516040519160208301936626b2b6b7b93c9d60c91b8552600160c01b6001900319809260c01b16602785015260c01b16602f830152603782015260378152610ae881610fc3565b51902090604081015190606081015160808201519160a00151926040519460208601966626b7b23ab6329d60c91b8852602787015260478601526067850152608784015260a783015263ffffffff60e01b9060e01b1660c782015260ab8152610b5081610ff9565b519020610b5b612ebe565b610b6c9163ffffffff88168a613cde565b03610cac57600163ffffffff851601928363ffffffff861611610be557610b92846132df565b15610c495750505060018451511b03610c17575b610bb563ffffffff82166132df565b15610bfb5750505180519081600019810111610be557610bd9916000190190612c37565b516101608201526104f0565b634e487b7160e01b600052601160045260246000fd5b9163ffffffff610c0c931690613c0c565b6101608201526104f0565b60405162461bcd60e51b815260206004820152600a6024820152692ba927a723afa622a0a360b11b6044820152606490fd5b92610c66610c6c93946101c4356004013560246101c4350161384e565b50613c0c565b14610ba65760405162461bcd60e51b815260206004820152601360248201527257524f4e475f524f4f545f464f525f5a45524f60681b6044820152606490fd5b60405162461bcd60e51b81526020600482015260136024820152722ba927a723afa927a7aa2fa327a92fa622a0a360691b6044820152606490fd5b50506106c46101c4356004013560246101c435018484612eed565b5050600181526104f0565b50506106c46101c4356004013560246101c435018484611a77565b50506106c46101c4356004013560246101c4350184846113b1565b61ffff811661802003610d57576008610488565b61ffff811661802103610d6b576007610488565b61ffff811661802203610d7f576006610488565b61ffff811661802303610d93576005610488565b61ffff811661802403610da7576004610488565b61ffff811661803003610dbb576003610488565b61ffff811661803103610dcf576002610488565b61ffff811661803203610de3576001610488565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b6044820152606490fd5b5061801361ffff82161115610480565b60a08436031261006f576040516001600160401b036080820190811190821117610ea95760a06020602494836080849501604052610e6e368a6110e8565b8152604089013583820152610e8560608a01611217565b6040820152610e9660808a01611217565b606082015281520195019492505061030f565b634e487b7160e01b600052604160045260246000fd5b3461006f57600036600319011261006f57606360f81b8152602090f35b60031115610ee657565b634e487b7160e01b600052602160045260246000fd5b80516007811015610ee6578252602090810151910152565b908151604092838352606083019151936020928382860152855180915283608086019601916000905b828210610f51575050505081015191015290565b90919296858282610f656001948c51610efc565b01980193920190610f3d565b604081019081106001600160401b03821117610ea957604052565b602081019081106001600160401b03821117610ea957604052565b61018081019081106001600160401b03821117610ea957604052565b606081019081106001600160401b03821117610ea957604052565b60c081019081106001600160401b03821117610ea957604052565b60e081019081106001600160401b03821117610ea957604052565b608081019081106001600160401b03821117610ea957604052565b601f909101601f19168101906001600160401b03821190821017610ea957604052565b6040519061105f82610f71565b600060208360405161107081610f8c565b6060815281520152565b6040519061108782610fc3565b60006040838281528260208201520152565b604051906110a682610fde565b600060a0838281526110b661107a565b60208201528260408201528260608201528260808201520152565b6001600160401b038111610ea95760051b60200190565b919082604091031261006f5760405161110081610f71565b80928035600781101561006f578252602090810135910152565b60409291818103841361006f5783519161113383610f71565b829481359260018060401b039384811161006f57830191602094858484031261006f5781519361116285610f8c565b803591821161006f570182601f8201121561006f578035611182816110d1565b9361118f8451958661102f565b818552878086019260061b8401019281841161006f579088809897969594939201915b8383106111c9575050505050815284520135910152565b9784959697986111dd8385969794956110e8565b815201920190889796959493926111b2565b919082604091031261006f5760405161120781610f71565b6020808294803584520135910152565b359063ffffffff8216820361006f57565b6001019081600111610be557565b91908201809211610be557565b9082101561124f570190565b634e487b7160e01b600052603260045260246000fd5b6000198114610be55760010190565b6060906020815260166020820152752aa725a727aba72fa82922a4a6a0a3a2afa82927a7a360511b60408201520190565b156112ac57565b60405162461bcd60e51b8152806112c560048201611274565b0390fd5b9093929384831161006f57841161006f578101920390565b6001600160401b038111610ea957601f01601f191660200190565b3d15611327573d9061130d826112e1565b9161131b604051938461102f565b82523d6000602084013e565b606090565b81810292918115918404141715610be557565b92919261134b826112e1565b91611359604051938461102f565b82948184528183011161006f578281602093846000960137010152565b1561137d57565b60405162461bcd60e51b815260206004820152600c60248201526b4241445f505245494d41474560a01b6044820152606490fd5b91936020928381016113cb6113c68251613f82565b613e3a565b6113d86113c68351613f82565b63ffffffff93848216601f841615908115916119b2575b5080156119a6575b61199657506040906060825161140c81610f8c565b52876307ffffff809260051c169701926114298a8c8a8751613d20565b9b9193909c849d6060611447611440838786611243565b3592611265565b936101a435806115305750505060f81c61151857826114879594928a9461146d936112c9565b94909361147b36878761133f565b8d815191012014611376565b1689810190818111610be557836114a7936114ae95841161150f576112c9565b369161133f565b955b6000995b87518b10156114e0576114d46114da918c8b818c01015160f81c916119ff565b9a611265565b996114b4565b906114fa92939495969a5061150797985061150d99613b82565b9151015251925116613f4e565b906140b1565b565b925080926112c9565b855162461bcd60e51b8152806112c560048201611274565b9697966001908082036115b857505050509280929161155561155a9560f81c156112a5565b6112c9565b9290918a600086518686823780878101838152039060025afa156115ad5790611587889260005114611376565b1689810190818111610be557836114a7936115a795841161150f576112c9565b956114b0565b84513d6000823e3d90fd5b929e989491969390959260021460001461195a57918f979695949380926115556115e59560f81c156112a5565b809591961161006f5785350361191f57600080885186888237808781018381520390600a5afa916116146112fc565b92156118e7578251156118aa57878380518101031261006f57878e840151930151937f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001850361187157908e918460051b85159386820414841715610be5578d851610611688575b50505050505050506114b0565b60009e989e9360051c1690805b85821061184d575050506118375760009283928e926116b991600160201b0461132c565b9088519183830193808552808b850152898401527f16a2a19edfe81f20d09b681922c813b4b63683508c2280b93829971f439f0d2b608084015260a083015260c082015260c0815261170a81610ff9565b519060055afa6117186112fc565b9015611803578a8151036117c9578a8151910151908b81106117b8575b5081851161006f578a8301350361178057821161006f5782519161175883610f71565b898352369082011161006f57828991018183013760008282015295388080808080808061167b565b835162461bcd60e51b8152600481018b9052601160248201527025ad23afa82927a7a32faba927a723afad60791b6044820152606490fd5b600019908c0360031b1b1638611735565b845162461bcd60e51b8152600481018c9052601360248201527209a9e888ab0a0beaea49e9c8ebe988a9c8ea89606b1b6044820152606490fd5b845162461bcd60e51b8152600481018c9052600d60248201526c1353d111561417d19052531151609a1b6044820152606490fd5b634e487b7160e01b600052601260045260246000fd5b909193811b938180821614611868575b811c91811b90611695565b9381179361185d565b60648f8a519062461bcd60e51b825260048201526013602482015272554e4b4e4f574e5f424c535f4d4f44554c555360681b6044820152fd5b875162461bcd60e51b8152600481018f905260166024820152754b5a475f505245434f4d50494c455f4d495353494e4760501b6044820152606490fd5b875162461bcd60e51b8152600481018f9052601160248201527024a72b20a624a22fa5ad23afa82927a7a360791b6044820152606490fd5b865162461bcd60e51b8152600481018e90526014602482015273096b48ebea0a49e9e8cbeaea49e9c8ebe9082a6960631b6044820152606490fd5b5060648f8a519062461bcd60e51b825260048201526015602482015274554e4b4e4f574e5f505245494d4147455f5459504560581b6044820152fd5b9650505050509350506002915052565b50601f821615156113f7565b9050888101809111610be55787890151516001600160401b031610386113ef565b90610100918203918211610be557565b600019810191908211610be557565b91908203918211610be557565b90916020831015611a3a5782601f0392601f8411610be5578360031b93840460081490601f141715610be55760ff809116831b921b19161790565b60405162461bcd60e51b81526020600482015260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b6044820152606490fd5b9390929363ffffffff611a906113c66020840151613f82565b16611aa16113c66020840151613f82565b63ffffffff81169460018060401b03611ac5611ac06020870151613f82565b613ec1565b16976101a435159687806123ce575b6123be5760208101809111610be5576020890151516001600160401b03161080156123b2575b6123a3576060604051611b0c81610f8c565b52611b2681836307ffffff8660051c1660208c0151613d20565b989195906001600160f81b0319611b3e828688611243565b351661236857611b4d90611265565b9183906000901561209657506028830190818411610be557606360f81b80611b7684888a611243565b3516611fa6575b50611b8b91508385876112c9565b98909760288a10611f6c5760009b60209b8d6008815b1015611be657611bb28e8e8e611243565b3560f81c9060081b67ffffffffffffff0016179c611bcf90611265565b9e611bdb60089f611265565b9e909f9d908f611ba1565b93989d50949990959a611c0192979c509d93989d369161133f565b60208151910120600091829084611ec2575b6001600160401b038116611e32575b5060405191602083019384526040830152606082015260608152611c4581611014565b5190206024356001600160a01b038116929083900361006f576020906024604051809581936316bf557960e01b835260048301525afa918215611e2657600092611def575b5003611db35760015b15611da457858710611d6b57611ca986886119f2565b9860009a5b602063ffffffff8d16108b8d82611d51575b505015611d1657611d008c63ffffffff9283611cf48e8e611cee8f8f611ce990878a1692611236565b611236565b91611243565b3560f81c9216906119ff565b9b1663ffffffff8114610be5576001019a611cae565b9061150d9a50602097508793959b995061150798506040949650916307ffffff611d449360051c1690613b82565b9201510152015191613f4e565b611d6391925063ffffffff1689611236565b108b8d611cc0565b60405162461bcd60e51b81526020600482015260116024820152702120a22fa6a2a9a9a0a3a2afa82927a7a360791b6044820152606490fd5b50506002905295505050505050565b60405162461bcd60e51b81526020600482015260146024820152734241445f534551494e424f585f4d45535341474560601b6044820152606490fd5b90916020823d602011611e1e575b81611e0a6020938361102f565b81010312611e1b5750519038611c8a565b80fd5b3d9150611dfd565b6040513d6000823e3d90fd5b6024356001600160a01b038116925082900361006f57611e536020916123da565b604051636ab8cee160e11b81526001600160401b03909116600482015291829060249082905afa908115611e2657600091611e90575b5038611c22565b906020823d602011611eba575b81611eaa6020938361102f565b81010312611e1b57505138611e89565b3d9150611e9d565b926024356001600160a01b0381168103611f68576020611ee1876123da565b6040516316bf557960e01b81526001600160401b03909116600482015291829060249082906001600160a01b03165afa918215611f5c578092611f27575b505092611c13565b9091506020823d602011611f54575b81611f436020938361102f565b81010312611e1b5750513880611f1f565b3d9150611f36565b604051903d90823e3d90fd5b5080fd5b60405162461bcd60e51b81526020600482015260126024820152712120a22fa9a2a8a4a72127ac2fa82927a7a360711b6044820152606490fd5b9050611fb4858381896112c9565b600092811561208257823516611fd2575b5050611b8b915038611b7d565b8060011161207e57906001611fec926000190191016124fe565b611ff581610edc565b600281146120425761200681610edc565b6001811461203a575b61201881610edc565b15612025575b8180611fc5565b5060818301809111610be557611b8b9061201e565b82915061200f565b60405162461bcd60e51b8152602060048201526014602482015273109313d094d51491505357d553911150d251115160621b6044820152606490fd5b8280fd5b634e487b7160e01b84526032600452602484fd5b91969b905060016101a49b9499969b9a9398959a3514600014612356576120bf8989818d6112c9565b6071811061231d5783918361227b575b81607111612277576120e93660701984016071840161133f565b602081519101208215612263578592600190875b60208110612231575050506050602160405193602085019560ff60f81b823516875260018060601b03199060601b168286015201603584013760858201526085815261214881610fde565b519020604051906020820192835260408201526040815261216881610fc3565b519020906024356001600160a01b038116919082900361222d57602090602460405180948193636ab8cee160e11b835260048301525afa928315611f5c5780936121f6575b5050036121bb576001611c93565b60405162461bcd60e51b81526020600482015260136024820152724241445f44454c415945445f4d45535341474560681b6044820152606490fd5b909192506020823d602011612225575b816122136020938361102f565b81010312611e1b5750519038806121ad565b3d9150612206565b8380fd5b90919461225561225b91612246888689611243565b3560f81c9060081b1796611265565b91611265565b9190916120fd565b634e487b7160e01b86526032600452602486fd5b8480fd5b91506024356001600160a01b038116810361227757602061229b856123da565b604051636ab8cee160e11b81526001600160401b03909116600482015291829060249082906001600160a01b03165afa9081156123125785916122e0575b50916120cf565b90506020813d60201161230a575b816122fb6020938361102f565b810103126122775751386122d9565b3d91506122ee565b6040513d87823e3d90fd5b60405162461bcd60e51b81526020600482015260116024820152702120a22fa222a620aca2a22fa82927a7a360791b6044820152606490fd5b50505050945095505050506002915052565b60405162461bcd60e51b81526020600482015260136024820152722aa725a727aba72fa4a72127ac2fa82927a7a360691b6044820152606490fd5b50505050915092506002915052565b50601f83161515611afa565b5050505050915092506002915052565b506004358a1015611ad4565b6001600160401b039081166000190191908211610be557565b9081602091031261006f57516001600160401b038116810361006f5790565b919082604091031261006f5760405161242a81610f71565b9182908035906001600160f81b03198216820361006f57908252602001359063ffffffff198216820361006f5760200152565b91909160608184031261006f576040519061247782610fc3565b909283919081356001600160401b03811161006f5782019080601f8301121561006f578135906124a6826110d1565b916124b4604051938461102f565b808352602093848085019260051b82010192831161006f578401905b8282106124ef5750505083528082013590830152604090810135910152565b813581529084019084016124d0565b8160081161006f57803560c01c906040908151926303f16d4b60e11b92838552602060049473c3e209eb245fd59c8586777b499d6a665df3abd29082888881855afa9788156115ad57600098612832575b506001600160401b039788166103e80188811161281d578816851161280f5782908785518094819382525afa908115612804579087916000916127d7575b501683116127cb578660581161006f578387018481039360571993916101008587011261006f5760588701356001600160a01b0381160361006f5760a06077198097011261006f578251946125e186610fc3565b6125ee8260788a01612412565b86526125fd8260b88a01612412565b8587015260f8880135848701526101188801358a811161006f57826058612626928b010161245d565b966101388901358b811161006f5789019060808285039384011261006f57859081519361265285610fc3565b60588401358552011261006f5784519661266b88610f71565b6078820135885260988201358789015286830197885260b88201358c811161006f576126c29460586126a0928c95010161245d565b8684015260008680516126b281610fc3565b60608152828a8201520152612851565b5093519083825192015194156127985703612761575050603883013503612757576126ec906129cd565b50938415612757578060101161006f5760181161006f576008601082013560c01c91013560c01c01908282116127425750612729612730916123da565b928061132c565b9116101561273d57600090565b600190565b601190634e487b7160e01b6000525260246000fd5b5050505050600190565b5162461bcd60e51b81528086019190915260116024820152704d69736d6174636865644865696768747360781b6044820152606490fd5b825162461bcd60e51b8152808901859052600d60248201526c24a72b20a624a22fa82927a7a360991b6044820152606490fd5b50505050505050600290565b6127f79150833d85116127fd575b6127ef818361102f565b8101906123f3565b3861258d565b503d6127e5565b83513d6000823e3d90fd5b505050505050505050600190565b601188634e487b7160e01b6000525260246000fd5b61284a919850833d85116127fd576127ef818361102f565b963861254f565b9082519260209081808201956040875193015182604051948593631f3302a960e01b8552600485015280516024850152015160448301526080606483015260e482019080519160606084850152825180915284610104850193019060005b81811061294e575050508381015160a48401526040015160c48301528190038173c3e209eb245fd59c8586777b499d6a665df3abd25afa908115611e2657600091612918575b501561290c576129089351015191612967565b9091565b50505050600090600490565b8281813d8311612947575b61292d818361102f565b81010312611f685751908115158203611e1b5750386128f5565b503d612923565b82518552889688965094850194909201916001016128af565b906129b79261297683516129eb565b604061298560208601516129eb565b940151906040519462ffffff19809216602087015216603d850152605a840152605a83526129b283611014565b612a5e565b50156129c557600190600090565b600090600390565b60400151600381166129e25760021c90600090565b50600090600890565b80516020918201516040516001600160f81b031990921692820192835263ffffffff19166021820152601d8152612a2181610f71565b51905162ffffff19918282169190601d8110612a3f575b5050905090565b83919250601d0360031b1b1616803880612a38565b60061115610ee657565b60408201805193949360018111612b5d5750825151612b50575b6020830190815181511115612b4257602060405196612ad482890189612ac3602160009c8d96878652612ab3815180928b8686019101612e59565b810103600181018452018261102f565b604051928392839251928391612e59565b8101039060025afa15612312578551935191825115612b2b5790612afd93949151905190612c4b565b6006811015612b175780612b115750149190565b92915050565b634e487b7160e01b85526021600452602485fd5b5051600114159050612b3c57149190565b50508190565b505050509050600090600290565b5050509050600090600190565b612b6e845151916020860151612b80565b14612a78575050509050600090600190565b9060019081811115612c2f5760005b8183821b10612c0e57610100908103908111610be557612bae906119d3565b9282612bb9856119e3565b1b90612bc4826119e3565b8111612bd1575050505090565b9293509091838203612be35750505090565b612c02935090612bf681612bfc936119f2565b926119f2565b90612b80565b612c0b90611228565b90565b82810180911115612b8f57634e487b7160e01b600052601160045260246000fd5b505050600090565b805182101561124f5760209160051b010190565b9392938115612d325760018214612d1d57845115612d125782612c6d83612d3d565b612c80612c7a88516119e3565b88612db3565b92818110612cdd5781612bf6612ca09693612c9a936119f2565b90612c4b565b9091612cab82612a54565b81612cd6575050612cca83612cc4612cd09495516119e3565b90612c37565b51612e7c565b90600090565b9350919050565b612ce79450612c4b565b9091612cf282612a54565b81612cd6575050612d0b83612cc4612cd09495516119e3565b5190612e7c565b505090915090600590565b5050909151612d2c5790600090565b90600490565b505090915090600390565b600180821061006f57600082805b612d6e57506000198101908111610be55781901b918214612d6a575090565b1c90565b90612d7890611265565b90821c80612d4b565b90612d8b826110d1565b612d98604051918261102f565b8281528092612da9601f19916110d1565b0190602036910137565b919082518111612dfc57612dc681612d81565b9060005b818110612dd8575090925050565b80612de6612df79287612c37565b51612df18286612c37565b52611265565b612dca565b60405162461bcd60e51b815260206004820152602f60248201527f496e76616c69642072616e67653a205f626567696e206f72205f656e6420617260448201526e65206f7574206f6620626f756e647360881b6064820152608490fd5b60005b838110612e6c5750506000910152565b8181015183820152602001612e5c565b612eab6000916020936040519085820192600160f81b84526021830152604182015260418152612ac381611014565b8101039060025afa15611e265760005190565b60405190612ecb82610f71565b601382527226b7b23ab6329036b2b935b632903a3932b29d60691b6020830152565b90612ef6612ebe565b90610160830151906020612f0f6113c682870151613f82565b9101612f2263ffffffff83168251613d99565b156132d257858791519260051c6307ffffff16612f3e93613d20565b509190604051612f4d81610f8c565b60609052604051612f5d81610f8c565b606081529587612f6b612ebe565b6101608801518096612f7b611099565b50612f84611099565b50612f8d61107a565b50612f9990848661380b565b612fa494919461107a565b50612fb0908287613789565b612fbb908388613789565b612fc990848994939461380b565b9160405193612fd785610fc3565b6001600160401b0391821685521660208401526040830152612ffa90838861380b565b61300890848995939561380b565b61301390838a61380b565b61301e91938a61373c565b9490936040519861302e8a610fde565b8952602089015260408801526060870152608086015263ffffffff1660a085015261305a908c8661373c565b613068908d879a939a61384e565b91909480519060208101518051602082015191604001516040519160208301936626b2b6b7b93c9d60c91b8552600160c01b6001900319809260c01b16602785015260c01b16602f8301526037820152603781526130c581610fc3565b51902090604081015190606081015160808201519160a00151926040519460208601966626b7b23ab6329d60c91b8852602787015260478601526067850152608784015260a783015263ffffffff60e01b9060e01b1660c782015260ab815261312d81610ff9565b519020613138612ebe565b6131499163ffffffff8c1688613cde565b03610cac5763ffffffff9b600198898e8216019d8e911611610be5578c95613170876132df565b15613274575050505050849051511b03610c17575b60009587613192816132df565b1561325857505085610be55791869060005b8183116131fb57505050906131d860209493926131ca6040519384928884019687613be9565b03601f19810183528261102f565b5190206101608201525b015190610be55761150763ffffffff61150d9316613f4e565b60405194856132108360208301938a85613be9565b0395613224601f199788810183528261102f565b5190209461324c60405191826132406020820195808c88613be9565b0390810183528261102f565b51902091811c916131a4565b602096956132699550935090613cde565b6101608201526131e2565b9091929550613289939b50613292945061384e565b50978989613c0c565b146131855760405162461bcd60e51b815260206004820152601360248201527257524f4e475f524f4f545f464f525f5a45524f60681b6044820152606490fd5b5050505090506002915052565b80151590816132ec575090565b60001981019150808211610be557161590565b6040519061330c82610f71565b8160405161331981610f71565b6040368237815260206040519161332f83610f71565b60403684370152565b90600281101561124f5760051b0190565b936113c69493929360208101906133636113c68351613f82565b9161337563ffffffff98899251613f82565b1693600285101561346657602061339191019783168851613d99565b1561345a5750906307ffffff6133bd9260051c169360606040516133b481610f8c565b52848751613d20565b949190506101843561ffff811680910361006f5761801081036133fb575050916133ee604094926133f59451613338565b5191613b82565b91510152565b91945094506180119192501460001461341b576134189151613338565b52565b60405162461bcd60e51b81526020600482015260176024820152764241445f474c4f42414c5f53544154455f4f50434f444560481b6044820152606490fd5b60029052505050505050565b5060029052505050505050565b602081019163ffffffff61348a6113c68551613f82565b169160028310156134e357506134b061150d935192602060018060401b03930151613338565b511690600060206040516134c381610f71565b8281520152604051916134d583610f71565b6001835260208301526140b1565b60029052505050565b602081019163ffffffff6135106113c6613509611ac08751613f82565b9551613f82565b169160028310156134e3575090602061352a920151613338565b6001600160401b039091169052565b6101408101516001016135605760408161355960a061150d940151613dc4565b0151613dc4565b60029052565b6000199081610140820151036134e35760a0810191825151146134e35783836040613592930151613612565b518260401161006f5761150d92603f19019160400190613612565b90916049926831b7ba343932b0b21d60b91b8352600983015260298201520190565b156135d657565b60405162461bcd60e51b815260206004820152601460248201527357524f4e475f434f5448524541445f454d50545960601b6044820152606490fd5b90600092600091825b602081106136b757509061362f929161380b565b5060001983036136595761364381156135cf565b6136516020830151156135cf565b602082015252565b6040516020810190613670816131ca8588866135ad565b5190206020830151146136515760405162461bcd60e51b8152602060048201526012602482015271057524f4e475f434f5448524541445f504f560741b6044820152606490fd5b92946136cb6136d191612246888686611243565b93611265565b61361b565b600019908160a08201515114613735576101a435806137105750610140810191808351146137085761150d9252613938565b506002915052565b91610140820151036137355761372f63ffffffff61150d931682613b16565b50613938565b6002915052565b600093918491905b6004831061375157505050565b9091939461377c6137829163ffffff0061376c898688611243565b3560f81c9160081b161796611265565b94611265565b9190613744565b9192600091825b600890818510156137cc576137c59167ffffffffffffff006136cb926137b78a878b611243565b3560f81c921b161796611265565b9294613790565b95945050509050565b90916000926000926000915b602083106137ee57505050565b9091939461377c61380491612246888587611243565b91906137e1565b600093918491905b6020831061382057505050565b9091939461377c61383691612246888587611243565b9190613813565b60ff1660ff8114610be55760010190565b90929192606060405161386081610f8c565b5261387961386f858385611243565b3560f81c94611265565b9061388385612d81565b9260005b60ff811693878510156138b557906107466138a66138b093868661380b565b9190919688612c37565b613887565b95965050505050604051906138c982610f8c565b815291565b8051906020808351930151910151602081519101516040519260208401946c23b637b130b61039ba30ba329d60991b8652602d850152604d84015260018060c01b0319809260c01b16606d84015260c01b166075820152605d815261393281611014565b51902090565b604080820180515160a0840180515194959091600019808714908115613b0c575b50613b00576080870193845196602095868901519960009a5b8a5180518d1015613a3a579061398b8d613a3493612c37565b518a6139978251613dfb565b918d828201516060828401519301519151938401946b29ba30b1b590333930b6b29d60a11b8652602c850152604c84015263ffffffff60e01b918260e091821b16606c8501521b166070820152605481526139f181611014565b519020908b51908b8201927129ba30b1b590333930b6b29039ba30b1b59d60711b84526032830152605290818301528152613a2b81611014565b5190209b611265565b9a613972565b50919497995091949950889295515201928351908782015191805151516000915b898b838510613a8f57505050505090606093929151525190868201525251928301525190613a8882610f8c565b6060825252565b613ab98584959698936000613abf94613af997519251613aae81610f71565b828152015251612c37565b51613dfb565b908b51908d8201926b2b30b63ab29039ba30b1b59d60a11b8452602c830152604c90818301528152613af081611014565b51902094611265565b9190613a5b565b50600290955250505050565b9050811438613959565b61014081019160001980845103613b79576101209163ffffffff60401b60e085015160401b1663ffffffff60201b61010086015160201b16179363ffffffff80948192015116911601828111610be557821601818111610be55716179052600190565b50505050600090565b90612c0b9260405160208101916b26b2b6b7b93c903632b0b31d60a11b8352602c820152602c8152613bb381610fc3565b5190209060405192613bc484610f71565b601384527226b2b6b7b93c9036b2b935b632903a3932b29d60691b6020850152613cde565b602090613bff6040959382815194859201612e59565b0191825260208201520190565b91926000936000925b8451958651851015613c9a5790613c67916001978689841615600014613c6d576131ca9150613c48613c5b918a51612c37565b5160405192839160208301958987613be9565b519020965b1c93611265565b92613c15565b613c7d613c91916131ca93612c37565b519260405192839160208301958987613be9565b51902096613c60565b95509250509150613ca757565b60405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606490fd5b92939091936000925b8451958651851015613c9a5790613d1a916001978689841615600014613c6d576131ca9150613c48613c5b918a51612c37565b92613ce7565b93909291936060604051613d3381610f8c565b526040613d59613d4f613d4686896137d5565b9096879961384e565b9590959686613b82565b91015103613d6357565b60405162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d3515357d493d3d560921b6044820152606490fd5b6020820190818311610be557516001600160401b031610159081613dbb575090565b601f9150161590565b600090805182198103613dd5575052565b60208201906131ca613df383516040519283916020830195866135ad565b519020905252565b8051906007821015610ee65760200151604051906020820192652b30b63ab29d60d11b845260f81b602683015260278201526027815261393281610fc3565b602081015190516007811015610ee657613e9257600160201b811015613e635763ffffffff1690565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606490fd5b60405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606490fd5b602081015190516007811015610ee657600103613f1f57600160401b811015613ef0576001600160401b031690565b60405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b6044820152606490fd5b60405162461bcd60e51b81526020600482015260076024820152661393d517d24d8d60ca1b6044820152606490fd5b60006020604051613f5e81610f71565b828152015263ffffffff60405191613f7583610f71565b6000835216602082015290565b90602091604051613f9281610f71565b600093818580935201525191806020604051613fad81610f71565b828152015282518051600019939184820191821161403c5790613fcf91612c37565b519284515190810190811161402857613fe790614050565b915b8251811015614022578061400161401d928751612c37565b5161400c8286612c37565b526140178185612c37565b50611265565b613fe9565b50925290565b634e487b7160e01b83526011600452602483fd5b634e487b7160e01b84526011600452602484fd5b9061405a826110d1565b6040906140698251918261102f565b838152809361407a601f19916110d1565b0191600090815b84811061408f575050505050565b602090825161409d81610f71565b848152828581830152828701015201614081565b5180515191600192838101809111610be5576140cc90614050565b926000815b6140f4575b50508151516140f0916140e98286612c37565b5283612c37565b5052565b83518051821015614129579061410d8161412393612c37565b516141188288612c37565b526140178187612c37565b816140d1565b506140d656fea26469706673582212207e6b4ebd82f27459282f76a25cfa39fb25968101272d30d056b40675568fbdbb64736f6c63430008130033",
}

// OneStepProverHostIoABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverHostIoMetaData.ABI instead.
var OneStepProverHostIoABI = OneStepProverHostIoMetaData.ABI

// OneStepProverHostIoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverHostIoMetaData.Bin instead.
var OneStepProverHostIoBin = OneStepProverHostIoMetaData.Bin

// DeployOneStepProverHostIo deploys a new Ethereum contract, binding an instance of OneStepProverHostIo to it.
func DeployOneStepProverHostIo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverHostIo, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverHostIoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// OneStepProverHostIo is an auto generated Go binding around an Ethereum contract.
type OneStepProverHostIo struct {
	OneStepProverHostIoCaller     // Read-only binding to the contract
	OneStepProverHostIoTransactor // Write-only binding to the contract
	OneStepProverHostIoFilterer   // Log filterer for contract events
}

// OneStepProverHostIoCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverHostIoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverHostIoSession struct {
	Contract     *OneStepProverHostIo // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverHostIoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverHostIoCallerSession struct {
	Contract *OneStepProverHostIoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverHostIoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverHostIoTransactorSession struct {
	Contract     *OneStepProverHostIoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverHostIoRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverHostIoRaw struct {
	Contract *OneStepProverHostIo // Generic contract binding to access the raw methods on
}

// OneStepProverHostIoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCallerRaw struct {
	Contract *OneStepProverHostIoCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverHostIoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactorRaw struct {
	Contract *OneStepProverHostIoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverHostIo creates a new instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIo(address common.Address, backend bind.ContractBackend) (*OneStepProverHostIo, error) {
	contract, err := bindOneStepProverHostIo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// NewOneStepProverHostIoCaller creates a new read-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverHostIoCaller, error) {
	contract, err := bindOneStepProverHostIo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoCaller{contract: contract}, nil
}

// NewOneStepProverHostIoTransactor creates a new write-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverHostIoTransactor, error) {
	contract, err := bindOneStepProverHostIo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoTransactor{contract: contract}, nil
}

// NewOneStepProverHostIoFilterer creates a new log filterer instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverHostIoFilterer, error) {
	contract, err := bindOneStepProverHostIo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoFilterer{contract: contract}, nil
}

// bindOneStepProverHostIo binds a generic wrapper to an already deployed contract.
func bindOneStepProverHostIo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transact(opts, method, params...)
}

// BLOBSTREAM is a free data retrieval call binding the contract method 0xdc346547.
//
// Solidity: function BLOBSTREAM() view returns(address)
func (_OneStepProverHostIo *OneStepProverHostIoCaller) BLOBSTREAM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProverHostIo.contract.Call(opts, &out, "BLOBSTREAM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BLOBSTREAM is a free data retrieval call binding the contract method 0xdc346547.
//
// Solidity: function BLOBSTREAM() view returns(address)
func (_OneStepProverHostIo *OneStepProverHostIoSession) BLOBSTREAM() (common.Address, error) {
	return _OneStepProverHostIo.Contract.BLOBSTREAM(&_OneStepProverHostIo.CallOpts)
}

// BLOBSTREAM is a free data retrieval call binding the contract method 0xdc346547.
//
// Solidity: function BLOBSTREAM() view returns(address)
func (_OneStepProverHostIo *OneStepProverHostIoCallerSession) BLOBSTREAM() (common.Address, error) {
	return _OneStepProverHostIo.Contract.BLOBSTREAM(&_OneStepProverHostIo.CallOpts)
}

// CELESTIAMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x5fd9e56d.
//
// Solidity: function CELESTIA_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_OneStepProverHostIo *OneStepProverHostIoCaller) CELESTIAMESSAGEHEADERFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _OneStepProverHostIo.contract.Call(opts, &out, "CELESTIA_MESSAGE_HEADER_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// CELESTIAMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x5fd9e56d.
//
// Solidity: function CELESTIA_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_OneStepProverHostIo *OneStepProverHostIoSession) CELESTIAMESSAGEHEADERFLAG() ([1]byte, error) {
	return _OneStepProverHostIo.Contract.CELESTIAMESSAGEHEADERFLAG(&_OneStepProverHostIo.CallOpts)
}

// CELESTIAMESSAGEHEADERFLAG is a free data retrieval call binding the contract method 0x5fd9e56d.
//
// Solidity: function CELESTIA_MESSAGE_HEADER_FLAG() view returns(bytes1)
func (_OneStepProverHostIo *OneStepProverHostIoCallerSession) CELESTIAMESSAGEHEADERFLAG() ([1]byte, error) {
	return _OneStepProverHostIo.Contract.CELESTIAMESSAGEHEADERFLAG(&_OneStepProverHostIo.CallOpts)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverHostIo.contract.Call(opts, &out, "executeOneStep", execCtx, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCallerSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// OneStepProverMathMetaData contains all meta data concerning the OneStepProverMath contract.
var OneStepProverMathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657611ee8908161001c8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c63a92cb5011461002757600080fd5b34610de45736600319016101e08112610de457606013610de4576001600160401b0360643511610de4576101c060643536036003190112610de457610100366083190112610de457604036610183190112610de4576101c4356001600160401b038111610de45736602382011215610de45760048101356001600160401b038111610de45736910160240111610de45760006101606100c4610eed565b8281526100cf610fae565b60208201526100dc610f0d565b83815283602082015260408201526100f2610fae565b60608201526100ff610f0d565b606081528360208201526080820152610116610f0d565b83815283602082015260a08201528260c08201528260e08201528261010082015282610120820152826101408201520152600060a0610153610f2c565b82815261015e610f4b565b8381528360208201528360408201526020820152826040820152826060820152826080820152015261018e610eed565b6003606435600401351015610de457606435600481013582526001600160401b0360249091013511610de4576101cf36606435602481013501600401611014565b60208201526101e3366044606435016110da565b60408201526001600160401b036064356084013511610de45761021136606435608481013501600401611014565b60608201526001600160401b0360643560a4013511610de457604060643560a48101350136036003190112610de457610248610f0d565b60643560a481013501600401356001600160401b038111610de4573660643560a48101350182016023011215610de45760643560a4810135018101600401359061029961029483610fce565b610f89565b8281529160208301913660643560a481013501820160a084020160240111610de45760643560a481013501810160240192905b60643560a481013501810160a08402016024018410610de957505050508152602460a46064350135606435010135602082015260808201526103133660c4606435016110da565b60a0820152610104606435013560c0820152610334610124606435016110fe565b60e0820152610348610144606435016110fe565b61010082015261035d610164606435016110fe565b6101208201526064356101848101356101408301526101a40135610160820152610385610f2c565b60843581529060603660a3190112610de45761039f610f4b565b60a4356001600160401b0381169003610de45760a435815260c4356001600160401b0381169003610de45760c435602082015260e4356040820152602083015261010435604083015261012435606083015261014435608083015263ffffffff61016435166101643503610de4576101643560a083015261ffff61042161110f565b16604581148015610dda575b15610c585750600b5b80600b14610b9b5780600a14610aec5780600714610a975780600914610a4b5780600614610a0057806008146109f657806005146109ec57806004146109bf578060031461096b578060021461079b576001146104a357634e487b7160e01b600052605160045260246000fd5b60bc61ffff6104b061110f565b160361070c57600060025b6104c86020840151611cf9565b90815160078110156106b65760078210156106b657036106cc5760078210156106b6576104fb9181526020830151611e34565b604051906101208252805160038110156106b65761012083015261053160208201516101c06101408501526102e0840190610e90565b90602060408201518051610160860152015161018084015261056760608201519261011f199384868303016101a0870152610e90565b608082015192848203016101c08501528251926040825260408201845180915260206060840195019060005b81811061066f575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c08601528051602086015260406020820151600180831b0381511682880152600180831b03602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b5161068d848251610e78565b858101516040850152826040820151168285015201511660808201520197019101919091610593565b634e487b7160e01b600052602160045260246000fd5b60405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f5245494e544552505245545f5459504560401b6044820152606490fd5b60bd61ffff61071961110f565b160361072857600160036104bb565b60be61ffff61073561110f565b160361074457600260006104bb565b60bf61ffff61075161110f565b160361076057600360016104bb565b60405162461bcd60e51b81526020600482015260136024820152721253959053125117d491525395115494149155606a1b6044820152606490fd5b5060c061ffff6107a961110f565b16036108bb57600860005b60078110156106b657806108ae5763ffffffff5b6107d56020850151611cf9565b91825160078110156106b6570361086d57600160ff84161b8060001981011161085757600019810160208401511680602085015260ff600019818716011161085757600160ff61083a9681600019911601161b1661083f575b50506020830151611e34565b6104fb565b6000190119166020820151176020820152388061082e565b634e487b7160e01b600052601160045260246000fd5b60405162461bcd60e51b81526020600482015260196024820152784241445f455854454e445f53414d455f545950455f5459504560381b6044820152606490fd5b6001600160401b036107c8565b60c161ffff6108c861110f565b16036108d757601060006107b4565b60c261ffff6108e461110f565b16036108f357600860016107b4565b60c361ffff61090061110f565b160361090f57601060016107b4565b60c461ffff61091c61110f565b160361092b57602060016107b4565b60405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f455854454e445f53414d455f5459504560401b6044820152606490fd5b5061083a61098461097f6020840151611cf9565b611b79565b60ac61ffff61099161110f565b16036109b4576109a09061118d565b6109ae602084015191611c66565b90611e34565b63ffffffff166109a0565b5061083a63ffffffff6109dd6109d86020850151611cf9565b611bd7565b166109ae602084015191611c3a565b5061083a81611949565b5061083a81611626565b50610a116109d86020830151611cf9565b60781961ffff610a1f61110f565b160161ffff811161085757610a3c63ffffffff9161083a936114ce565b166109ae602084015191611c66565b50610a5c61097f6020830151611cf9565b60661961ffff610a6a61110f565b16019061ffff82116108575761083a9163ffffffff610a899216611342565b6109ae602084015191611c3a565b50610aa86109d86020830151611cf9565b610ab86109d86020840151611cf9565b60501961ffff610ac661110f565b160161ffff81116108575761083a92610ade926111b3565b6109ae602084015191611c96565b50610afd61097f6020830151611cf9565b610b0d61097f6020840151611cf9565b60451961ffff610b1b61110f565b160161ffff81116108575761083a92610ade9261ffff831660028114908115610b90575b8115610b85575b8115610b7a575b5015610b6b57610b5f610b659161118d565b9161118d565b906111b3565b63ffffffff91821691166111b3565b600891501438610b4d565b600681149150610b46565b600481149150610b3f565b50610ba96020820151611cf9565b604561ffff610bb661110f565b1603610bf75780519060078210156106b657610bd460209215611157565b0151610bed5761083a60016109ae602084015191611c3a565b61083a6000610a89565b605061ffff610c0461110f565b1603610c295780519060078210156106b657610c24600160209314611121565b610bd4565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa2a8ad60c91b6044820152606490fd5b604681101580610dcf575b15610c705750600a610436565b606781101580610dc4575b15610c8857506009610436565b606a81101580610db9575b15610ca057506008610436565b605181101580610dae575b15610cb857506007610436565b607981101580610da3575b15610cd057506006610436565b607c81101580610d98575b15610ce857506005610436565b60a78103610cf857506004610436565b60ac81148015610d8e575b15610d1057506003610436565b60c081101580610d83575b15610d2857506002610436565b60bc8110159081610d77575b5015610d41576001610436565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b6044820152606490fd5b60bf9150111538610d34565b5060c4811115610d1b565b5060ad8114610d03565b50608a811115610cdb565b50607b811115610cc3565b50605a811115610cab565b506078811115610c93565b506069811115610c7b565b50604f811115610c63565b506050811461042d565b600080fd5b60a084360312610de4576040516001600160401b036080820190811190821117610e625760a06020602494836080849501604052610e27368a610fe5565b8152604089013583820152610e3e60608a016110fe565b6040820152610e4f60808a016110fe565b60608201528152019501949250506102cc565b634e487b7160e01b600052604160045260246000fd5b805160078110156106b6578252602090810151910152565b908151604092838352606083019151936020928382860152855180915283608086019601916000905b828210610ecd575050505081015191015290565b90919296858282610ee16001948c51610e78565b01980193920190610eb9565b6040519061018082016001600160401b03811183821017610e6257604052565b60408051919082016001600160401b03811183821017610e6257604052565b6040519060c082016001600160401b03811183821017610e6257604052565b60405190606082016001600160401b03811183821017610e6257604052565b60405190602082016001600160401b03811183821017610e6257604052565b6040519190601f01601f191682016001600160401b03811183821017610e6257604052565b610fb6610f0d565b90610fbf610f6a565b60608152825260006020830152565b6001600160401b038111610e625760051b60200190565b9190826040910312610de457610ff9610f0d565b918035906007821015610de457602091845201356020830152565b91906040928381830312610de45761102a610f0d565b936001600160401b03928235848111610de4578301916020948584840312610de457611054610f6a565b938035918211610de4570182601f82011215610de457803561107861029482610fce565b93878086848152019260061b84010192818411610de4579088809897969594939201915b8383106110b457505050505081528552013590830152565b9784959697986110c8838596979495610fe5565b8152019201908897969594939261109c565b9190826040910312610de45760206110f0610f0d565b928035845201356020830152565b359063ffffffff82168203610de457565b6101843561ffff81168103610de45790565b1561112857565b60405162461bcd60e51b81526020600482015260076024820152661393d517d24d8d60ca1b6044820152606490fd5b1561115e57565b60405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606490fd5b638000000081166111a15763ffffffff1690565b63ffffffff1663ffffffff60201b1790565b9161ffff16806111cf57506001600160401b0391821691161490565b600181036111ea57506001600160401b039182169116141590565b60028103611209575060018060401b0380911660070b911660070b1290565b6003810361122357506001600160401b0390811691161090565b60048103611242575060018060401b0380911660070b911660070b1390565b6005810361125c57506001600160401b0390811691161190565b6006810361127c575060018060401b0380911660070b911660070b131590565b6007810361129757506001600160401b039081169116111590565b600881036112b7575060018060401b0380911660070b911660070b121590565b6009036112d0576001600160401b039081169116101590565b60405162461bcd60e51b815260206004820152600a6024820152690424144204952454c4f560b41b6044820152606490fd5b90600163ffffffff8093160191821161085757565b63ffffffff908116600019019190821161085757565b9063ffffffff80921660200391821161085757565b9060009061ffff16806113c5575060205b63ffffffff808216151590816113a3575b50156113785761137390611317565b611353565b63ffffffff9081166020039250821161138f575090565b634e487b7160e01b81526011600452602490fd5b6001600160401b0391506001906113b984611317565b161b8416161538611364565b600192839290918381036114235750925b6113df57505090565b63ffffffff831660208110908161140c575b5015611407576114018293611302565b926113d6565b505090565b83901b82166001600160401b0316159050386113f1565b6002919293501461145f5760405162461bcd60e51b815260206004820152600960248201526804241442049556e4f760bc1b6044820152606490fd5b82918182945b611471575b5050505090565b63ffffffff8091169060208210156114c85784821b83166001600160401b03166114b8575b81146114a457830183611465565b634e487b7160e01b83526011600452602483fd5b946114c290611302565b94611496565b5061146a565b9060009061ffff168061153d575060405b63ffffffff8082161515908161151b575b5015611504576114ff90611317565b6114df565b63ffffffff9081166040039250821161138f575090565b6001600160401b03915060019061153184611317565b161b84161615386114f0565b600192839290918381036115965750925b61155757505090565b63ffffffff831660408110908161157f575b5015611407576115798293611302565b9261154e565b83901b82166001600160401b031615905038611569565b600291929350146115d25760405162461bcd60e51b815260206004820152600960248201526804241442049556e4f760bc1b6044820152606490fd5b82918182945b6115e3575050505090565b63ffffffff8091169060408210156114c85784821b83166001600160401b0316611616575b81146114a4578301836115d8565b9461162090611302565b94611608565b6020810161163761097f8251611cf9565b61164461097f8351611cf9565b906000936101843561ffff9081811680910361184957606919019581871161138f57506003959081168681036116ef57505063ffffffff80921690811580156116cf575b6116c45750840b80156116ae576116ac94826109ae9416900b0516915b5191611c3a565b565b634e487b7160e01b600052601260045260246000fd5b600290525050505050565b50838316860b637fffffff19148015611688575081860b60001914611688565b6005810361172557505063ffffffff8092169081156116c45750840b80156116ae576116ac94826109ae9416900b0716916116a5565b929592600a810361175057505050506109ae906116ac93601f63ffffffff80931691161b16916116a5565b600c810361177857505050506109ae906116ac93601f63ffffffff80931691161c16916116a5565b600b81036117a257505050906116ac93601f6109ae9363ffffffff809416900b91161d16916116a5565b91925090600d81036117de575050506109ae90601f6116ac94169063ffffffff8091168181816117d18661132d565b161c16921b1617916116a5565b600e036118145750506109ae90601f6116ac94169063ffffffff8091168181816118078661132d565b161b16921c1617916116a5565b611829919263ffffffff80809716911661184d565b91909161184057506116ac926109ae9116916116a5565b60029052505050565b8680fd5b6000939261ffff16806118695750016001600160401b03169190565b600181036118805750036001600160401b03169190565b600281036118975750026001600160401b03169190565b600481036118c157506001600160401b03918216919082156118b95716049190565b505050600190565b600681036118e357506001600160401b03918216919082156118b95716069190565b600781036118f15750169190565b600881036118ff5750179190565b60090361190b57189190565b60405162461bcd60e51b81526020600482015260166024820152750494e56414c49445f47454e455249435f42494e5f4f560541b6044820152606490fd5b6020810161195a6109d88251611cf9565b916119686109d88351611cf9565b9060006101843561ffff90818116809103611b5d57607b19019181831161138f5750811660038103611a0257506001600160401b03948516919050811580156119dd575b6119d3575060070b9283156116ae576116ac93816109ae931660070b0516915b5191611c66565b6002905250505050565b5082851660070b60016001603f1b03191480156119ac57506000198260070b146119ac565b60058103611a3f57506001600160401b0394851691905081156119d3575060070b9283156116ae576116ac93816109ae931660070b0716916119cc565b600a8103611a68575050506109ae906116ac93603f60018060401b0380931691161b16916119cc565b600c8103611a91575050506109ae906116ac93603f60018060401b0380931691161c16916119cc565b600b8103611abd575050506109ae906116ac93603f60018060401b0380931660070b91161d16916119cc565b600d8103611afa57506116ac946109ae93603f90911692506001600160401b0391508116818181611aed86611b61565b161c16921b1617916119cc565b600e03611b3557506116ac936109ae92603f90911691506001600160401b03908116818181611b2886611b61565b161b16921c1617916119cc565b909293611b419261184d565b919091611b5557506109ae6116ac926119cc565b600290525050565b8280fd5b6001600160401b039081166040039190821161085757565b6020810151905160078110156106b657611b939015611157565b600160201b811015611ba85763ffffffff1690565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606490fd5b6020810151905160078110156106b6576001611bf39114611121565b600160401b811015611c0b576001600160401b031690565b60405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b6044820152606490fd5b60006020611c46610f0d565b828152015263ffffffff611c58610f0d565b916000835216602082015290565b60006020611c72610f0d565b8281520152611c7f610f0d565b600181526001600160401b03909116602082015290565b602090611ca1610f0d565b6000808252920182905215611cd557806020611cbb610f0d565b8281520152611cc8610f0d565b9081526001602082015290565b806020611ce0610f0d565b8281520152611ced610f0d565b90808252602082015290565b90602091611d05610f0d565b600093818580935201525191806020611d1c610f0d565b8281520152825180516000199391848201918211611d975790611d3e91611dab565b51928451519081019081116114a457611d5690611dd5565b915b8251811015611d915780611d70611d8c928751611dab565b51611d7b8286611dab565b52611d868185611dab565b50611e25565b611d58565b50925290565b634e487b7160e01b84526011600452602484fd5b8051821015611dbf5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b90611de261029483610fce565b8281528092611df3601f1991610fce565b01906000805b838110611e065750505050565b602090611e11610f0d565b838152828481830152828601015201611df9565b60001981146108575760010190565b518051519160019283810180911161085757611e4f90611dd5565b926000815b611e77575b5050815151611e7391611e6c8286611dab565b5283611dab565b5052565b83518051821015611eac5790611e9081611ea693611dab565b51611e9b8288611dab565b52611d868187611dab565b81611e54565b50611e5956fea264697066735822122059fd774e8b47f2284a9db82d2c1c5ca9aaa287de1296d00c17eeeb5b6f3d3ce364736f6c63430008130033",
}

// OneStepProverMathABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMathMetaData.ABI instead.
var OneStepProverMathABI = OneStepProverMathMetaData.ABI

// OneStepProverMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMathMetaData.Bin instead.
var OneStepProverMathBin = OneStepProverMathMetaData.Bin

// DeployOneStepProverMath deploys a new Ethereum contract, binding an instance of OneStepProverMath to it.
func DeployOneStepProverMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMath, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// OneStepProverMath is an auto generated Go binding around an Ethereum contract.
type OneStepProverMath struct {
	OneStepProverMathCaller     // Read-only binding to the contract
	OneStepProverMathTransactor // Write-only binding to the contract
	OneStepProverMathFilterer   // Log filterer for contract events
}

// OneStepProverMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMathSession struct {
	Contract     *OneStepProverMath // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProverMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMathCallerSession struct {
	Contract *OneStepProverMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProverMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMathTransactorSession struct {
	Contract     *OneStepProverMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProverMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMathRaw struct {
	Contract *OneStepProverMath // Generic contract binding to access the raw methods on
}

// OneStepProverMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMathCallerRaw struct {
	Contract *OneStepProverMathCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactorRaw struct {
	Contract *OneStepProverMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMath creates a new instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMath(address common.Address, backend bind.ContractBackend) (*OneStepProverMath, error) {
	contract, err := bindOneStepProverMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// NewOneStepProverMathCaller creates a new read-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMathCaller, error) {
	contract, err := bindOneStepProverMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathCaller{contract: contract}, nil
}

// NewOneStepProverMathTransactor creates a new write-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMathTransactor, error) {
	contract, err := bindOneStepProverMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathTransactor{contract: contract}, nil
}

// NewOneStepProverMathFilterer creates a new log filterer instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMathFilterer, error) {
	contract, err := bindOneStepProverMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathFilterer{contract: contract}, nil
}

// bindOneStepProverMath binds a generic wrapper to an already deployed contract.
func bindOneStepProverMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.OneStepProverMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMath.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverMemoryMetaData contains all meta data concerning the OneStepProverMemory contract.
var OneStepProverMemoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608080604052346100165761199b908161001c8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c63a92cb5011461002757600080fd5b346108505736600319016101e0811261085057606013610850576001600160401b0360643511610850576101c0606435360360031901126108505761010036608319011261085057604036610183190112610850576001600160401b036101c43511610850573660236101c435011215610850576001600160401b0360046101c435013511610850573660246101c435600401356101c435010111610850576100d06080610959565b60006080526100dd610a04565b60a0526040516100ec81610975565b6000808252602082015260c052610101610a04565b60e05260405161011081610975565b60608152600060208201526101005260405161012b81610975565b600080825260208201819052610120919091526101408190526101608190526101808190526101a08190526101c08190526101e081905260405160a09061017181610990565b828152604051610180816109ab565b838152836020820152836040820152602082015282604082015282606082015282608082015201526040516101b481610959565b600360643560040135101561085057606435600481013582526001600160401b0360249091013511610850576101f536606435602481013501600401610a75565b602082015261020936604460643501610b4a565b60408201526001600160401b0360643560840135116108505761023736606435608481013501600401610a75565b60608201526001600160401b0360643560a401351161085057604060643560a481013501360360031901126108505760405161027281610975565b60643560a481013501600401356001600160401b038111610850573660643560a481013501820160230112156108505760643560a481013501810160040135906102bb82610a2c565b916102c960405193846109e1565b80835260208301913660643560a481013501820160a0840201602401116108505760643560a481013501810160240192905b60643560a481013501810160a0840201602401841061085557505050508152602460a46064350135606435010135602082015260808201526103423660c460643501610b4a565b60a0820152610104606435013560c082015261036361012460643501610b72565b60e082015261037761014460643501610b72565b61010082015261038c61016460643501610b72565b6101208201526064356101848101356101408301526101a40135610160820152604051906103b982610990565b608435825260603660a3190112610850576040516103d6816109ab565b60a4356001600160401b03811690036108505760a435815260c4356001600160401b03811690036108505760c435602082015260e4356040820152602083015261010435604083015261012435606083015261014435608083015263ffffffff61016435166101643503610850576101643560a08301526101843561ffff8116810361085057602861ffff8216101580610841575b156107b3575060045b80600414610799578060031461077f578060021461075e576001146104a957634e487b7160e01b600052605160045260246000fd5b63ffffffff60208301515160101c166104dc63ffffffff6104d56104d060208601516117b8565b6116fd565b1682610b83565b60208481015101519091906001600160401b0316821161071a57818060101b04620100001482151715610704576105339160018060401b039060101b1660208501515261052d602084015191611784565b906118e7565b604051906101208252805160038110156106ee5761012083015261056960208201516101c06101408501526102e08401906108fc565b90602060408201518051610160860152015161018084015261059f60608201519261011f199384868303016101a08701526108fc565b608082015192848203016101c08501528251926040825260408201845180915260206060840195019060005b8181106106a7575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c08601528051602086015260406020820151600180831b0381511682880152600180831b03602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b516106c58482516108e4565b8581015160408501528260408201511682850152015116608082015201970191019190916105cb565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b505061075960208201516000602060405161073481610975565b82815201526040519061074682610975565b6000825263ffffffff60208301526118e7565b610533565b5061075963ffffffff60208401515160101c1661052d602084015191611784565b506107596101c4356004013560246101c435018484610f3d565b506107596101c4356004013560246101c435018484610b90565b603661ffff8216101580610832575b156107cf57506003610474565b61ffff8116603f036107e357506002610474565b61ffff166040036107f5576001610474565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b6044820152606490fd5b50603e61ffff821611156107c2565b50603561ffff8216111561046b565b600080fd5b60a084360312610850576040516001600160401b0360808201908111908211176108ce5760a06020602494836080849501604052610893368a610a43565b81526040890135838201526108aa60608a01610b72565b60408201526108bb60808a01610b72565b60608201528152019501949250506102fb565b634e487b7160e01b600052604160045260246000fd5b805160078110156106ee578252602090810151910152565b908151604092838352606083019151936020928382860152855180915283608086019601916000905b828210610939575050505081015191015290565b9091929685828261094d6001948c516108e4565b01980193920190610925565b61018081019081106001600160401b038211176108ce57604052565b604081019081106001600160401b038211176108ce57604052565b60c081019081106001600160401b038211176108ce57604052565b606081019081106001600160401b038211176108ce57604052565b602081019081106001600160401b038211176108ce57604052565b601f909101601f19168101906001600160401b038211908210176108ce57604052565b60405190610a1182610975565b6000602083604051610a22816109c6565b6060815281520152565b6001600160401b0381116108ce5760051b60200190565b919082604091031261085057604051610a5b81610975565b809280356007811015610850578252602090810135910152565b60409291818103841361085057835191610a8e83610975565b829481359260018060401b039384811161085057830191602094858484031261085057815193610abd856109c6565b8035918211610850570182601f82011215610850578035610add81610a2c565b93610aea845195866109e1565b818552878086019260061b84010192818411610850579088809897969594939201915b838310610b24575050505050815284520135910152565b978495969798610b38838596979495610a43565b81520192019088979695949392610b0d565b919082604091031261085057604051610b6281610975565b6020808294803584520135910152565b359063ffffffff8216820361085057565b9190820180921161070457565b9092916000936101843561ffff95868216809203610f2b575060288103610da15750600094600490610bef87965b60208701958463ffffffff966020610be789610bdd6104d08d516117b8565b166101a435610b83565b9101516115d2565b509590610d9457506001600160401b0385811696909590610c3a575b50505050519060405192610c1e84610975565b60078510156106ee57610c389484521660208301526118e7565b565b90859394969291600192838114808091610d7f575b15610c6d57505050505060ff91501660000b16915b38808080610c0b565b80610d69575b15610c8b575050505060ff1660000b16929050610c64565b9294919392600281148080610d54575b15610cae575050505016900b1691610c64565b9296919280610d3e575b15610cca5750505016900b1691610c64565b91955093925060049150149081610d2b575b5015610cee571660030b811691610c64565b60405162461bcd60e51b815260206004820152601560248201527410905117d491505117d096551154d7d4d251d39151605a1b6044820152606490fd5b905060078610156106ee57851438610cdc565b50955060078910156106ee578695858a14610cb8565b50925060078a10156106ee5787928a15610c9b565b50945060078910156106ee578694838a14610c73565b50955060078a10156106ee5787958a15610c4f565b6002905250505050505050565b60298103610dbb5750600194600890610bef600096610bbe565b602a8103610dd55750600294600490610bef600096610bbe565b602b8103610def5750600394600890610bef600096610bbe565b602c8103610e085750600094600190610bef8296610bbe565b602d8103610e215750600094600190610bef8796610bbe565b602e8103610e3b5750600094600290610bef600196610bbe565b602f8103610e545750600094600290610bef8796610bbe565b60308103610e6c57506001948590610bef8796610bbe565b60318103610e8557506001948590610bef600096610bbe565b60328103610e9e5750600194600290610bef8796610bbe565b60338103610eb85750600194600290610bef600096610bbe565b60348103610ed15750600194600490610bef8796610bbe565b603503610ee957600194600490610bef600096610bbe565b60405162461bcd60e51b815260206004820152601a602482015279494e56414c49445f4d454d4f52595f4c4f41445f4f50434f444560301b6044820152606490fd5b80fd5b60001981146107045760010190565b9290916101843561ffff811680910361085057603681036111a85750600460005b60209081870190610f6f82516117b8565b90815160078110156106ee5760078210156106ee5703611172578201516001600160401b0397600894808a16948a8416949093909187871061112f575b505063ffffffff610bdd6104d0610fc393516117b8565b9782610fcf868b610b83565b9101998a5151161061112157509493956000939291939560001996604051610ff6816109c6565b6060815296819982945b868610611023575050505050505050509161101d91604093611286565b91510152565b90919293949596979a6110368783610b83565b8d8160051c908d82036110dc575b5050601f168681101561109f5761105a90611278565b908160031b918083048b1490151715610704576110919166ffffffffffffff918e60ff809116831b921b1916179c8a1c1696610f2e565b949392919097969597611000565b60405162461bcd60e51b81526004810188905260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b6044820152606490fd5b818691601f99949f9e89949f9661110097600019810361110a575b50505051611421565b9a9b90958e611044565b61111392611286565b6040825101523880806110f7565b600290525050505050505050565b9194509060031b6008600160401b038116906008600160431b03168103610704576001901b8a16600019018a8111610704571689169263ffffffff610bdd610fac565b60405162461bcd60e51b815260048101849052600e60248201526d4241445f53544f52455f5459504560901b6044820152606490fd5b603781036111ba575060086001610f5e565b603881036111cc575060046002610f5e565b603981036111de575060086003610f5e565b603a81036111f0575060016000610f5e565b603b8103611202575060026000610f5e565b603c81036112135750600180610f5e565b603d8103611225575060026001610f5e565b603e036112355760046001610f5e565b60405162461bcd60e51b815260206004820152601b60248201527a494e56414c49445f4d454d4f52595f53544f52455f4f50434f444560281b6044820152606490fd5b601f0390601f821161070457565b92604091825193602094858101916b26b2b6b7b93c903632b0b31d60a11b8352602c820152602c81526112b8816109ab565b519020928051916112c883610975565b601383527226b2b6b7b93c9036b2b935b632903a3932b29d60691b86840152936000945b875191825187101561137b579061134a91600193888584161560001461135057611330915061131f61133e918d516113c2565b5187519283918d8301958b876113ec565b03601f1981018352826109e1565b519020925b1c95610f2e565b946112ec565b61136061137291611330936113c2565b519287519283918d8301958b876113ec565b51902092611343565b96955092505091945061138d57505090565b60649250519062461bcd60e51b82526004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152fd5b80518210156113d65760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b9392909384519060005b82811061140e57500191825260208201526040019150565b80602080928901015181840152016113f6565b91929490939460409160608351611437816109c6565b526000969560005b6020811061158e57509161147f9695939188959360ff999160608551611464816109c6565b526114708985896115c6565b359460f895861c9a8b9a610f2e565b966114898b610a2c565b9961149683519b8c6109e1565b8b8b52601f196114a58d610a2c565b60209a9101368d8c013760009e8f5b169d8e10156115245760009060005b898d8d83106114f7575050506114da909e8d6113c2565b5260ff809e169d8e1461070457600160ff9e019d9c8c9d8f6114b4565b9361150b846115199361151f9596976115c6565b358d1c9060081b1793610f2e565b91610f2e565b6114c3565b939a9b50939b50955095965061154e929a508593508351986115458a6109c6565b89529888611286565b9101510361155a575050565b60649250519062461bcd60e51b82526004820152600e60248201526d15d493d391d7d3515357d493d3d560921b6044820152fd5b95939692976115b46115ba916115a5868c896115c6565b3560f81c9060081b1794610f2e565b96610f2e565b9792979693959661143f565b908210156113d6570190565b94929190946000946115e48388610b83565b82516001600160401b0316106116f05791906000969396198693879888965b8588106116195750505050505050506000929190565b909192939495969961162b8b83610b83565b8060051c8681036116d2575b50601f16602080821015611695575061164f90611278565b9060039180831b92600891808504831490151715610704578d901b908d8204148d1517156107045760ff8a611689941c16901b179a610f2e565b96959493929190611603565b6064906040519062461bcd60e51b8252600482015260166024820152750848288bea0aa9898be988a828cbe84b2a88abe9288b60531b6044820152fd5b985094506116e5601f9a85858b8a611421565b509a90989590611637565b5050505091506001918190565b6020810151905160078110156106ee5761175557600160201b8110156117265763ffffffff1690565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606490fd5b60405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606490fd5b6000602060405161179481610975565b828152015263ffffffff604051916117ab83610975565b6000835216602082015290565b906020916040516117c881610975565b6000938185809352015251918060206040516117e381610975565b82815201528251805160001993918482019182116118725790611805916113c2565b519284515190810190811161185e5761181d90611886565b915b825181101561185857806118376118539287516113c2565b5161184282866113c2565b5261184d81856113c2565b50610f2e565b61181f565b50925290565b634e487b7160e01b83526011600452602483fd5b634e487b7160e01b84526011600452602484fd5b9061189082610a2c565b60409061189f825191826109e1565b83815280936118b0601f1991610a2c565b0191600090815b8481106118c5575050505050565b60209082516118d381610975565b8481528285818301528287010152016118b7565b51805151916001928381018091116107045761190290611886565b926000815b61192a575b50508151516119269161191f82866113c2565b52836113c2565b5052565b8351805182101561195f579061194381611959936113c2565b5161194e82886113c2565b5261184d81876113c2565b81611907565b5061190c56fea2646970667358221220b0017d4c040eabd6f7e980c5516459a648d1bec14ed2b38d05b1fde6b1ef516764736f6c63430008130033",
}

// OneStepProverMemoryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMemoryMetaData.ABI instead.
var OneStepProverMemoryABI = OneStepProverMemoryMetaData.ABI

// OneStepProverMemoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMemoryMetaData.Bin instead.
var OneStepProverMemoryBin = OneStepProverMemoryMetaData.Bin

// DeployOneStepProverMemory deploys a new Ethereum contract, binding an instance of OneStepProverMemory to it.
func DeployOneStepProverMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMemory, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// OneStepProverMemory is an auto generated Go binding around an Ethereum contract.
type OneStepProverMemory struct {
	OneStepProverMemoryCaller     // Read-only binding to the contract
	OneStepProverMemoryTransactor // Write-only binding to the contract
	OneStepProverMemoryFilterer   // Log filterer for contract events
}

// OneStepProverMemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMemorySession struct {
	Contract     *OneStepProverMemory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverMemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMemoryCallerSession struct {
	Contract *OneStepProverMemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverMemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMemoryTransactorSession struct {
	Contract     *OneStepProverMemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverMemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMemoryRaw struct {
	Contract *OneStepProverMemory // Generic contract binding to access the raw methods on
}

// OneStepProverMemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCallerRaw struct {
	Contract *OneStepProverMemoryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactorRaw struct {
	Contract *OneStepProverMemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMemory creates a new instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemory(address common.Address, backend bind.ContractBackend) (*OneStepProverMemory, error) {
	contract, err := bindOneStepProverMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// NewOneStepProverMemoryCaller creates a new read-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMemoryCaller, error) {
	contract, err := bindOneStepProverMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryCaller{contract: contract}, nil
}

// NewOneStepProverMemoryTransactor creates a new write-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMemoryTransactor, error) {
	contract, err := bindOneStepProverMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryTransactor{contract: contract}, nil
}

// NewOneStepProverMemoryFilterer creates a new log filterer instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMemoryFilterer, error) {
	contract, err := bindOneStepProverMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryFilterer{contract: contract}, nil
}

// bindOneStepProverMemory binds a generic wrapper to an already deployed contract.
func bindOneStepProverMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.OneStepProverMemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMemory.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemorySession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0xa92cb501.
//
// Solidity: function executeOneStep((uint256,address,bytes32) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}
