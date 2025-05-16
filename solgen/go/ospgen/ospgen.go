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
	MaxInboxMessagesRead *big.Int
	Bridge               common.Address
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
	Bin: "0x608080604052346015576118e9908161001b8239f35b600080fdfe608060408181526004918236101561001657600080fd5b600090813560e01c908163740085d71461039c5750806379754cba1461034d578063ae364ac214610321578063b7465799146102b05763d4e5dd2b1461005b57600080fd5b346102ad57816003193601126102ad576001600160401b0391833583811161029e5761008a903690860161047b565b92610093610420565b9061009f36868561181f565b95865160208098012095816060941694858211610257575b50508451916100c5836104e2565b60019060018452600189850194868652898352828b528883208884528b528883209051151560ff801983541691161781550193519081519384116102445790849392916101148a9796546104a8565b8b601f821161020b575b50508a92601f851160011461018d57509280610179959361016593600080516020611894833981519152989692610182575b50508160011b916000199060031b1c19161790565b90555b85519182918983528983019061043b565b0390a351908152f35b015190508c80610150565b8582528b8220939291601f198616915b8d8383106101f05750505092600192859260008051602061189483398151915298966101799896106101d7575b505050811b019055610168565b015160001960f88460031b161c191690558b80806101ca565b8486015187558d9a509581019594850194919091019061019d565b86835280832061023392601f880160051c820192881061023a575b601f0160051c0190610612565b8b8b61011e565b9091508190610226565b634e487b7160e01b815260418b52602490fd5b9091935061026585836106bc565b8881116102a6575b6102779086610699565b918286116102a257821161029e5790610296918580369303910161181f565b9138806100b7565b8280fd5b8380fd5b508761026d565b80fd5b50913461029e57602036600319011261029e57356001600160a01b0381169081900361029e5791819281526001602052209060018060401b038254169060096102fb6001850161056d565b9301546103198251948594855260606020860152606085019061043b565b918301520390f35b5090346103495781600319360112610349576103469033835260016020528220610629565b80f35b5080fd5b50346102ad5760603660031901126102ad578235906001600160401b0382116102ad57506103836020936103959236910161047b565b61038b610420565b9060443592610707565b9051908152f35b919050346102ad57826003193601126102ad578335836103ba610420565b92828152806020528181209360018060401b031693848252602052209460ff8654161561040857610404856103f16001890161056d565b905191829160208352602083019061043b565b0390f35b6309cb23c960e11b8452830152602482015260449150fd5b602435906001600160401b038216820361043657565b600080fd5b919082519283825260005b848110610467575050826000602080949584010152601f8019910116010190565b602081830181015184830182015201610446565b9181601f84011215610436578235916001600160401b038311610436576020838186019501011161043657565b90600182811c921680156104d8575b60208310146104c257565b634e487b7160e01b600052602260045260246000fd5b91607f16916104b7565b604081019081106001600160401b038211176104fd57604052565b634e487b7160e01b600052604160045260246000fd5b61032081019081106001600160401b038211176104fd57604052565b60a081019081106001600160401b038211176104fd57604052565b601f909101601f19168101906001600160401b038211908210176104fd57604052565b90604051918260008254610580816104a8565b908184526020946001916001811690816000146105f057506001146105b1575b5050506105af9250038361054a565b565b600090815285812095935091905b8183106105d85750506105af93508201013880806105a0565b855488840185015294850194879450918301916105bf565b925050506105af94925060ff191682840152151560051b8201013880806105a0565b81811061061d575050565b60008155600101610612565b6002610659600092838155836001820161064381546104a8565b8061065c575b5050506009810192839101610612565b55565b82601f8211600114610674575050555b833880610649565b9091808252610692601f60208420940160051c840160018501610612565b555561066c565b919082018092116106a657565b634e487b7160e01b600052601160045260246000fd5b919082039182116106a657565b908210156106d5570190565b634e487b7160e01b600052603260045260246000fd5b91909160198310156106d5576018908360021c019260031b1690565b929391909160009160028616611804575b60018616158015906117f9575b156117c057336000526001602052604060002095600987015480156000146117705787546001600160401b0319166001600160401b0384161788559495945b8695806107758960098c0154610699565b60098b01555b871580611765575b6117585760005b60888110611697575060405161079f81610513565b61032036823760005b8b601982106116605750506103206040516107c281610513565b3690376040516107d18161052f565b60a03682376040516107e28161052f565b60a0368237604051906107f482610513565b6103203683376040516001600160401b036103008201908111908211176104fd576103008101604090815260018252618082602083015267800000000000808a90820152678000000080008000606082015261808b6080820152638000000160a0820181905267800000008000808160c0830181905267800000000000800960e0840152608a61010084015260886101208401526380008009610140840152638000000a610160840152638000808b610180840152608b6001603f1b016101a08401526780000000000080896101c08401526780000000000080036101e084015267800000000000800261020084015260806001603f1b0161022084015261800a61024084015267800000008000000a6102608401526102808301526780000000000080806102a08301526102c08201526780000000800080086102e082015260005b60188110610e0057505050505060005b60198110610dc05750506088881061096f57876088116104365760880196608719019661077b565b50919395509193955b6001600160401b0394851692602084018085116106a65781811180610db3575b610c75575b505050505060011615610c6c5760005b60208110610bed575060405160018501916109c7826104e2565b600182526109d48361056d565b956020830196875285600052600060205260406000208282541660005260205260016040600020935115159360ff199460ff8683541691161781550196519687518381116104fd57610a2682546104a8565b601f8111610bbb575b506020601f8211600114610b4f579080610a65928a9b60009b999a9b92610b445750508160011b916000199060031b1c19161790565b90555b54169260405191602083526000918054610a81816104a8565b92836020870152600182169182600014610b16575050600114610acc575b50509080600080516020611894833981519152920390a33360005260016020526105af6040600020610629565b6000908152602081209092505b818310610afc575050810160400181600080516020611894833981519152610a9f565b805460408585010152879450602090920191600101610ad9565b86955060409350600080516020611894833981519152969492501682840152151560051b8201019192610a9f565b015190503880610150565b601f198216998360005260206000209a60005b818110610ba35750918a9b9184600195949c9a9b9c10610b8a575b505050811b019055610a68565b015160001960f88460031b161c19169055388080610b7d565b838301518d556001909c019b60209384019301610b62565b610be790836000526020600020601f840160051c8101916020851061023a57601f0160051c0190610612565b38610a2f565b9260039084821c60058082069182820292828404036106a6578592610c2092610c17920490610699565b600289016106eb565b905490841b1c169185901b603881166008600788168015908304821417156106a65787830414871517156106a65760f89182039182116106a65760019360ff911c16901b1793016109ad565b50915050600090565b600094828111610d9a575b5090610c8b916106bc565b92818411610d91575b929160018901935b8381101561099d57610caf8184846106c9565b3590855491610cbd836104a8565b90600160401b8210156104fd57601f9384831115610d35576002610ce49101808a556104a8565b93848310156106d557602060019510600014610d1d578892811690035b60ff83549160031b9260f81c831b921b19161790555b01610c9c565b886000528060206000208460051c0193169003610d01565b6001946001600160f81b031990921660001a929091601f8214610d7057906002910360031b9260ff908116841b931b19910116178655610d17565b5050908760005260206000209160ff19169060ff1617905560418655610d17565b92508092610c94565b82610c8b93929650610dab916106bc565b949091610c80565b5060098a01548510610998565b989a91999198600190610df48b610dee8360026001600160401b03610de5838a611882565b511693016106eb565b90611865565b019a989a999199610947565b9c9e9c600190969e968651602088015118604088015118606088015118608088015118865260a087015160c08801511860e088015118610100880151186101208801511880602088015261014088015161016089015118610180890151186101a0890151186101c08901511860408801526101e08801516102008901511861022089015118610240890151186102608901511860608801526102808801516102a0890151186102c0890151186102e089015118610300890151189081608089015280603f1c90848060401b0390851b1617188085528651604088015180603f1c90858060401b0390861b16171860208601526020870151606088015180603f1c90858060401b0390861b16171860408601526040870151608088015180603f1c90858060401b0390861b16171860608601526060870151875180603f1c90858060401b0390861b16171860808601528751188088526020880151855118602089015260408801518551186040890152606088015185511860608901526080880151855118608089015260a088015160208601511860a089015260c088015160208601511860c089015260e088015160208601511860e08901526101008801516020860151186101008901526101208801516020860151186101208901526101408801516040860151186101408901526101608801516040860151186101608901526101808801516040860151186101808901526101a08801516040860151186101a08901526101c08801516040860151186101c08901526101e08801516060860151186101e08901526102008801516060860151186102008901526102208801516060860151186102208901526102408801516060860151186102408901526102608801516060860151186102608901526102808801516080860151186102808901526102a08801516080860151186102a08901526102c08801516080860151186102c08901526102e08801516080860151186102e0890152610300880151608086015118610300890152808652602088015180601c1c90848060401b039060241b1617610100870152604088015180603d1c90848060401b039060031b161761016087015260608801518060171c90848060401b039060291b1617610260870152608088015180602e1c90848060401b039060121b16176102c087015260a088015180603f1c90848060401b0390851b1617604087015260c0880151908160141c848060401b0383602c1b161760a088015260e08901518060361c90858060401b0390600a1b16176101a08801526101008901518060131c90858060401b0390602d1b1617610200880152610120890151603e9080821c90868060401b039060021b16176103008901526101408a0151908160021c91868060401b03911b1617608088015261016089015180603a1c90858060401b039060061b161760e0880152610180890151918260151c858060401b0384602b1b16176101408901526101a08a01518060311c90868060401b0390600f1b16176102408901526101c08a01518060031c90868060401b0390603d1b16176102a08901526101e08a01518060241c90868060401b0390601c1b161760208901526102008a01518060091c90868060401b039060371b16176101208901526102208a01518060271c90868060401b039060191b16176101808901526102408a015180602b1c90868060401b039060151b16176101e08901526102608a01518060081c90868060401b039060381b16176102e08901526102808a01518060251c90868060401b0390601b1b161760608901526102a08a015180602c1c90868060401b039060141b161760c08901526102c08a01518060191c90868060401b039060271b16176101c08901526102e08a01518060381c90868060401b039060081b16176102208901526103008a01518060321c90868060401b0390600e1b16176102808901528260151c858060401b0384602b1b16178160141c868060401b0383602c1b1617191682188a52602088015160c0890151196101608a0151161860208b0152604088015160e0890151196101808a0151161860408b01526060880151610100890151196101a08a0151161860608b01526080880151610120890151196101c08a0151161860808b015260a0880151610140890151196101e08a0151161860a08b015260c0880151610160890151196102008a0151161860c08b015260e0880151610180890151196102208a0151161860e08b01526101008801516101a0890151196102408a015116186101008b01526101208801516101c0890151196102608a015116186101208b01526101408801516101e0890151196102808a015116186101408b0152610160880151610200890151196102a08a015116186101608b0152610180880151610220890151196102c08a015116186101808b01526101a0880151610240890151196102e08a015116186101a08b01526101c0880151610260890151196103008a015116186101c08b01526101e088015161028089015119895116186101e08b01526102008801516102a08901511960208a015116186102008b01526102208801516102c08901511960408a015116186102208b01526102408801516102e08901511960608a015116186102408b01526102608801516103008901511960808a015116186102608b015261028088015188511960a08a015116186102808b01526102a088015160208901511960c08a015116186102a08b01526102c088015160408901511960e08a015116186102c08b01526102e08801516060890151196101008a015116186102e08b01526103008801516080890151196101208a015116186103008b01528360051b860151928060151c90868060401b0390602b1b1617908060141c90868060401b0390602c1b1617191618188752019e9c9e9d959d610937565b906116748160026001949e969d9e016106eb565b838060401b0391549060031b1c1661168c8285611882565b5201999891996107a8565b9798909760008a82101561173457506116b1818b8b6106c9565b3560f81c905b60039181831c60058082069081810291818304036106a6576116da920490610699565b9060078316603884861b1690808204600814901517156106a6576001948f60029060ff9361171461172a978a8060401b03948594016106eb565b95909616901b1691838554911b1c161891611865565b019897909861078a565b908a811461174f575b608781036116b75790608017906116b7565b6001915061173d565b5091939550919395610978565b506001841615610783565b87546001600160401b039081169084160361178d57959495610764565b60405162461bcd60e51b815260206004820152600b60248201526a1112519197d3d19194d15560aa1b6044820152606490fd5b60405162461bcd60e51b81526020600482015260116024820152701393d517d09313d0d2d7d0531251d39151607a1b6044820152606490fd5b506088840615610725565b33600052600160205261181a6040600020610629565b610718565b9192916001600160401b0382116104fd5760405191611848601f8201601f19166020018461054a565b829481845281830111610436578281602093846000960137010152565b919060018060401b038084549260031b9316831b921b1916179055565b9060198110156106d55760051b019056fef88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802ca2646970667358221220760987728ecbd99eb86e5b72cf13e6682c010174700b1d370a8d5ec0dd149abc64736f6c63430008190033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"}],\"name\":\"getEndMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getStartMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) GetEndMachineHash(opts *bind.CallOpts, status uint8, globalStateHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "getEndMachineHash", status, globalStateHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntrySession) GetEndMachineHash(status uint8, globalStateHash [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetEndMachineHash(&_IOneStepProofEntry.CallOpts, status, globalStateHash)
}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) GetEndMachineHash(status uint8, globalStateHash [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetEndMachineHash(&_IOneStepProofEntry.CallOpts, status, globalStateHash)
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

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// IOneStepProverMetaData contains all meta data concerning the IOneStepProver contract.
var IOneStepProverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"instruction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"result\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"resultMod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCallerSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// OneStepProofEntryMetaData contains all meta data concerning the OneStepProofEntry contract.
var OneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"prover0_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMem_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMath_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverHostIo_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"}],\"name\":\"getEndMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getStartMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prover0\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverHostIo\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMath\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMem\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60803460ae57601f61282138819003918201601f19168301916001600160401b0383118484101760b35780849260809460405283398101031260ae5760428160c9565b90604d6020820160c9565b60616060605b6040850160c9565b930160c9565b9060018060a01b03928380928160018060a01b0319971687600054161760005516856001541617600155168360025416176002551690600354161760035560405161274490816100dd8239f35b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820360ae5756fe6080604052600436101561001257600080fd5b6000803560e01c90816304997be41461008a575080631f128bc01461008557806330a5509f146100805780635d3adcfb1461007b5780635f52fd7c1461007657806366e5d9c3146100715763d8558b871461006c57600080fd5b6102c6565b610293565b61026a565b610206565b6101dd565b6101af565b346101ac5760403660031901126101ac576101986101a8916100aa61042a565b906100b3611492565b6100bc83610478565b526100c682610478565b506100cf6114b5565b6100d88361048a565b526100e28261048a565b506100eb6114b5565b6100f48361049a565b526100fe8261049a565b506101076103be565b9182526101126103cd565b9182528060208301526101236104d1565b61012b6104f9565b610133610411565b600019815260006020820152916101486103da565b9484865260208601528260408601526060850152608084015260a083015260043560c08301528060e0830152806101008301526101208201526000196101408201526024356101608201526114d8565b6040519081529081906020820190565b0390f35b80fd5b346101d85760003660031901126101d8576001546040516001600160a01b039091168152602090f35b600080fd5b346101d85760003660031901126101d8576000546040516001600160a01b039091168152602090f35b346101d857366003190160a081126101d8576040136101d8576001600160401b036084358181116101d857366023820112156101d85780600401359182116101d85736602483830101116101d8576101a89160246101989201606435604435610e63565b346101d85760003660031901126101d8576003546040516001600160a01b039091168152602090f35b346101d85760003660031901126101d8576002546040516001600160a01b039091168152602090f35b600411156101d857565b346101d85760403660031901126101d85760206102f16004356102e8816102bc565b602435906113d0565b604051908152f35b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b0382111761032a57604052565b6102f9565b602081019081106001600160401b0382111761032a57604052565b606081019081106001600160401b0382111761032a57604052565b60c081019081106001600160401b0382111761032a57604052565b608081019081106001600160401b0382111761032a57604052565b601f909101601f19168101906001600160401b0382119082101761032a57604052565b604051906103cb8261032f565b565b604051906103cb8261030f565b6040519061018082016001600160401b0381118382101761032a57604052565b6001600160401b03811161032a5760051b60200190565b6040519061041e8261030f565b60006020838281520152565b6040519061043782610380565b600382528160005b6060811061044b575050565b602090610456610411565b8282850101520161043f565b634e487b7160e01b600052603260045260246000fd5b8051156104855760200190565b610462565b8051600110156104855760400190565b8051600210156104855760600190565b80518210156104855760209160051b010190565b604051906104cb8261032f565b60608252565b604051906104de8261030f565b60006020836040516104ef8161032f565b6060815281520152565b604051906105068261030f565b6000602083606081520152565b634e487b7160e01b600052602160045260246000fd5b6004111561053357565b610513565b60048210156105335752565b6040519061018082016001600160401b0381118382101761032a576040526000808352610160836105736104d1565b6020820152610580610411565b604082015261058d6104d1565b606082015261059a6104f9565b60808201526105a7610411565b60a08201528260c08201528260e08201528261010082015282610120820152826101408201520152565b604051906105de8261034a565b60006040838281528260208201520152565b604051906105fd82610365565b600060a08382815261060d6105d1565b60208201528260408201528260608201528260808201520152565b1561062f57565b60405162461bcd60e51b815260206004820152601360248201527209a828690929c8abe848a8c9ea48abe9082a69606b1b6044820152606490fd5b634e487b7160e01b600052601160045260246000fd5b906001820180921161068e57565b61066a565b600e019081600e1161068e57565b906002820180921161068e57565b906020820180921161068e57565b156106c457565b60405162461bcd60e51b815260206004820152600c60248201526b1353d115531154d7d493d3d560a21b6044820152606490fd5b156106ff57565b60405162461bcd60e51b815260206004820152601260248201527110905117d1955390d51253d394d7d493d3d560721b6044820152606490fd5b909392938483116101d85784116101d8578101920390565b90600163ffffffff8093160191821161068e57565b51906103cb826102bc565b91908260409103126101d8576040516107898161030f565b8092805160078110156101d8578252602090810151910152565b6040929181810384136101d857604051916107bd8361030f565b8051929485936001600160401b0393908481116101d85783019160209485848403126101d857604051936107f08561032f565b80519182116101d8570182601f820112156101d8578051610810816103fa565b9361081e604051958661039b565b818552878086019260061b840101928184116101d8579088809897969594939201915b838310610858575050505050815284520151910152565b97849596979861086c838596979495610771565b81520192019088979695949392610841565b91908260409103126101d8576040516108968161030f565b6020808294805184520151910152565b519063ffffffff821682036101d857565b919091604080828503126101d8578051916108d18361030f565b8051929485936001600160401b0381116101d857820181601f820112156101d857805190602094610901836103fa565b9361090e8251958661039b565b838552868501918760a0809602850101938285116101d8579088809897969594939201925b8484106109495750505050505084520151910152565b90919293948096979850848403126101d857888691835161096981610380565b6109738688610771565b81528487015183820152606061098a8189016108a6565b8683015261099a608089016108a6565b9082015281520193019190889796959493610933565b51906001600160401b03821682036101d857565b809291039161010083126101d857604051906109df82610365565b6060829482518452601f1901126101d857610a5560e060a092604051610a048161034a565b610a10602083016109b0565b8152610a1e604083016109b0565b60208201526060820151604082015260208601526080810151604086015283810151606086015260c08101516080860152016108a6565b910152565b91906101209081848203126101d85783516001600160401b0392908381116101d8578501906101c0828403126101d857610a926103da565b90610a9c83610766565b825260208301518581116101d85784610ab69185016107a3565b6020830152610ac8846040850161087e565b604083015260808301518581116101d85784610ae59185016107a3565b606083015260a08301519485116101d8576101a083610b0b86610b7b98602097016108b7565b6080850152610b1d8660c0830161087e565b60a08501526101008082015160c0860152610b398483016108a6565b60e086015261014090610b4d8284016108a6565b9086015261016093610b608584016108a6565b908601526101808201519085015201519082015294016109c4565b90565b9060048210156105335752565b80516007811015610533578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b838210610be757505050505060208091015191015290565b9091929396838282610bfc6001948c51610b8b565b0198019493920190610bcf565b90604091604082019281519360408452845180915260609460608501956020809201936000915b848310610c495750505050505060208091015191015290565b9091929394978460a06001928b51610c62828251610b8b565b80840151828801528681015163ffffffff9081168784015290860151166080820152019901959493019190610c30565b908060209392818452848401376000828201840152601f01601f1916010190565b60043581526024356001600160a01b038116969395939491926101c092918890036101d857610e33610d63610d18610e4995610b7b9b6020890152806040890152610d018189018651610b7e565b6020850151906101e0890152610380880190610ba3565b6040840151805161020089015260200151610220880152610d4d6060850151916101bf1992838a8303016102408b0152610ba3565b9060808501519088830301610260890152610c09565b60a08301518051610280880152602001516102a08701529760c08301516102c087015260e083015163ffffffff166102e087015261010083015163ffffffff1661030087015261012083015163ffffffff1661032087015261014083015161034087015261016080930151610360870152606086019063ffffffff60a060e0928051855260406020820151600180831b03808251166020890152602082015116828801520151606086015260408101516080860152606081015182860152608081015160c0860152015116910152565b8301906020809161ffff81511684520151910152565b6101a0818503910152610c92565b6040513d6000823e3d90fd5b610eaa90939193610e72610544565b50610e7b6105f0565b50610e846104be565b50610e8d610411565b50610e9883856116d4565b929095610ea4876114d8565b14610628565b8451610eb581610529565b610ebe81610529565b611368576001602b1b90610ed190610680565b1461135857610ee19082846118a2565b610eec90838561197b565b9290918160e0870195828751610f059063ffffffff1690565b63ffffffff1690610f169187611a6c565b9561016096878a015114610f29906106bd565b610f316104be565b50610f3a6104be565b50610f46908383611b58565b610f5490848497939761197b565b610f6290838598939861197b565b91906101208c0197828951610f7a9063ffffffff1690565b60061c6303ffffff1663ffffffff16610f9292611c4f565b6101008d015163ffffffff1663ffffffff16610fad92611d6d565b606087015114610fbc906106f8565b8651610fcd91603f909116906104aa565b5193610fd893610739565b909651610fe89063ffffffff1690565b63ffffffff16938051610ffe9063ffffffff1690565b61100790610751565b63ffffffff169052815161ffff1697600093908460288b10158061134d575b8015611336575b801561132c575b8015611322575b861461114b5750600154611071906001600160a01b0316935b604051633604366f60e01b81529a8b968795869560048701610cb3565b03916001600160a01b03165afa90811561114657610b7b95600095869361111d575b506180238114908115611111575b50156110fb575b50505050600281516110b981610529565b6110c281610529565b14806110ea575b156114d8576110d781611dd3565b6110e081611e55565b50600081526114d8565b5061014081015160001914156110c9565b61110492611a6c565b90820152388080806110a8565b618024915014386110a1565b90925061113c91953d8091833e611134818361039b565b810190610a5a565b9490949138611093565b610e57565b60458b148015611318575b8015611301575b80156112ea575b80156112d3575b80156112bc575b80156112a5575b801561128e575b8015611284575b8015611270575b8015611259575b8015611242575b86146111ba5750600254611071906001600160a01b03165b93611054565b6180108b101580611236575b801561121d575b8015611204575b86146111f05750600354611071906001600160a01b03166111b4565b54611071906001600160a01b031693611054565b506180308b101580156111d457506180328b11156111d4565b506180208b101580156111cd57506180248b11156111cd565b506180138b11156111c6565b5060bc8b1015801561119c575060bf8b111561119c565b5060c08b10158015611195575060c48b1115611195565b5060ac8b148061118e575060ad8b1461118e565b5060a78b14611187565b50607c8b101580156111805750608a8b1115611180565b5060798b101580156111795750607b8b1115611179565b5060518b101580156111725750605a8b1115611172565b50606a8b1015801561116b575060788b111561116b565b5060678b10158015611164575060698b1115611164565b5060468b1015801561115d5750604f8b111561115d565b5060508b14611156565b5060408b1461103b565b50603f8b14611034565b5060368b1015801561102d5750603e8b111561102d565b5060358b1115611026565b50506002825250610b7b906114d8565b50505050610b7b906114d8565b6f26b0b1b434b732903a37b7903330b91d60811b815260100190565b6f26b0b1b434b7329032b93937b932b21d60811b815260100190565b6031917026b0b1b434b732903334b734b9b432b21d60791b825260118201520190565b6113d981610529565b6001810361140b5750604051611405816113f76020820194856113ad565b03601f19810183528261039b565b51902090565b905061141681610529565b6002810361143457506040516020810190611405816113f784611391565b80611440600392610529565b0361145a576040516020810190611405816113f784611375565b60405162461bcd60e51b815260206004820152601060248201526f4241445f424c4f434b5f53544154555360801b6044820152606490fd5b61149a610411565b506040516114a78161030f565b600481526000602082015290565b6114bd610411565b506040516114ca8161030f565b600081526000602082015290565b80516114e381610529565b6114ec81610529565b6116085760408101516114056115056020840151611e7b565b926113f761154361152461014084019560001997888851141591611fb6565b9560a0840151906115386080860151612019565b908751141591611fb6565b936115516060840151611e7b565b9260c08101519161156960e083015163ffffffff1690565b61010083015163ffffffff169061016061158b61012086015163ffffffff1690565b935194015194604051998a9860208a019c8d97959492909160dc999794926f26b0b1b434b73290393ab73734b7339d60811b8a5260108a015260308901526050880152607087015263ffffffff60e01b9283809260e01b16609088015260e01b16609486015260e01b166098840152609c83015260bc8201520190565b6001815161161581610529565b61161e81610529565b0361163c5760c00151604051611405816113f76020820194856113ad565b6002815161164981610529565b61165281610529565b0361166d57506040516020810190611405816113f784611391565b6003905161167a81610529565b61168381610529565b0361169d576040516020810190611405816113f784611375565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b6044820152606490fd5b9160ff91926116e1610544565b506116ec8482612110565b93168061183657506000925b6117006104d1565b506117096104d1565b50611712610411565b5061171b6104f9565b50611724610411565b5061173090858361213b565b61173e908684979397612205565b61174c90838598939861213b565b6117579084866122b0565b611765908587949394612205565b9890936117706103da565b9861177b908a610538565b602089015260408801526060870152608086015260a085015260c08401906000825260e0850160008152610100860160008152610120870191600083526101408801946000865261016089019860008a52996117d8908689612520565b91526117e5908588612396565b63ffffffff9091169091526117fb908487612396565b63ffffffff909116909152611811908386612396565b63ffffffff909116909152611827908285612520565b925261183292612520565b9252565b6001810361184757506001926116f8565b6002810361185857506002926116f8565b600303611867576003926116f8565b60405162461bcd60e51b8152602060048201526013602482015272554e4b4e4f574e5f4d4143485f53544154555360681b6044820152606490fd5b61193963ffffffff9161194361194b6118ee9561192f6118f86118db6119029a6118ca6105f0565b506118d36105d1565b508488612520565b6118e69a919a6105d1565b5084886126c1565b84889392936126c1565b84889c929c612520565b916040519b6119108d61034a565b6001600160401b039182168d521660208c015260408b01528286612520565b8286979297612520565b8286959295612520565b919094612396565b939093966040519661195c88610365565b875260208701526040860152606085015260808401521660a082015291565b90929192606060405161198d8161032f565b526119a661199c8583856120f5565b3560f81c94612101565b906119b0856103fa565b926119be604051948561039b565b858452601f196119cd876103fa565b0136602086013760005b60ff81169387851015611a095760ff91611a016119f76001938787612520565b91909197896104aa565b5201166119d7565b9596505050505060405190611a1d8261032f565b815291565b6001600160401b03811161032a57601f01601f191660200190565b60405190611a4a8261030f565b601382527226b7b23ab6329036b2b935b632903a3932b29d60691b6020830152565b90610b7b92805190611b4660208201518051604060208301519201516040519160208301936626b2b6b7b93c9d60c91b855260018060c01b0319809260c01b16602785015260c01b16602f830152603782015260378152611acc8161034a565b5190206113f7604084015193606081015190611af460a0608083015192015163ffffffff1690565b9160405196879560208701998a9492909160ab9694926626b7b23ab6329d60c91b87526007870152602786015260478501526067840152608783015263ffffffff60e01b9060e01b1660a78201520190565b51902090611b52611a3d565b92612451565b611b6860ff939492948583612120565b931690611b74826103fa565b90611b82604051928361039b565b828252601f19611b91846103fa565b0160005b818110611c04575050819560005b848110611bb1575050505050565b600190611bc2611bcc9884866124ea565b84869a929a612520565b9890611be3611bd96103cd565b61ffff9093168352565b6020820152611bf282876104aa565b52611bfd81866104aa565b5001611ba3565b602090611c0f610411565b82828701015201611b95565b60405190611c288261030f565b601882527724b739ba393ab1ba34b7b71036b2b935b632903a3932b29d60411b6020830152565b929192611c6c611c67611c628651612562565b610693565b612573565b9160209460208401946c24b739ba393ab1ba34b7b7399d60991b8652611ca6611c96835160ff1690565b60f81b6001600160f81b03191690565b9360009460001a611cb6876125a5565b53600093600e5b8451861015611d5357611d4b878b611d3d848c611d0e600197611ce08e8e6104aa565b5196611d06611c96611d00611cf78b5161ffff1690565b60081c60ff1690565b60ff1690565b901a926125b5565b538c611d378d611d26611c96611d00895161ffff1690565b611d2f85610680565b911a926125b5565b536106a1565b9101518c828c0101526106af565b950194611cbd565b5094509491509550610b7b94915051902090611b52611c1b565b90610b7b92604051602081019168233ab731ba34b7b71d60b91b8352602982015260298152611d9b8161034a565b5190209060405192611dac8461030f565b6015845274233ab731ba34b7b71036b2b935b632903a3932b29d60591b6020850152612451565b604081018051519060a0830180515190600019808314908115611e4b575b50611e4157602085608060609596970192611e0c8451612019565b9051520193611e1b8551611e7b565b90515251906020820152525190602082015260405190611e3a8261032f565b6060825252565b5050505060029052565b9050841438611df1565b611e666101408201918251906125c6565b15611e75576000199052600190565b50600090565b9060209160208101519281515151906000925b828410611e9b5750505050565b90919294600190611ec0611eba888551611eb3610411565b50516104aa565b51612600565b9060405190858201926b2b30b63ab29039ba30b1b59d60a11b8452602c830152604c90818301528152611ef281610380565b5190209501929190611e8e565b15611f0657565b60405162461bcd60e51b81526020600482015260196024820152784d554c5449535441434b5f4e4f535441434b5f41435449564560381b6044820152606490fd5b91606b93916a36bab63a34b9ba30b1b59d60a91b8452600b840152602b830152604b8201520190565b15611f7757565b60405162461bcd60e51b815260206004820152601760248201527626aaa62a24a9aa20a1a5afa727a9aa20a1a5afa6a0a4a760491b6044820152606490fd5b9160001990611fc783831415611eff565b15611ff857611fd99083511415611f70565b61140560208351930151916113f7604051938492602084019687611f47565b509061140560208251920151916113f7604051938492602084019687611f47565b602080820151926000935b835180518610156120ed57906113f76120e2612042886001956104aa565b519261204e8451612600565b878501516040956060878201519101518751928b8401946b29ba30b1b590333930b6b29d60a11b8652602c850152604c84015263ffffffff60e01b918260e091821b16606c8501521b166070820152605481526120aa81610380565b5190209351928391888301958690916052927129ba30b1b590333930b6b29039ba30b1b59d60711b8352601283015260328201520190565b519020940193612024565b509350915050565b90821015610485570190565b600019811461068e5760010190565b9015610485573560f81c90600190565b82610b7b92612131929594956120f5565b3560f81c92612101565b9061215361215d9361214b6104d1565b508284612520565b8284959295612520565b9290929461216a846103fa565b93612178604051958661039b565b808552612187601f19916103fa565b0160005b8181106121ee5750506000955b84518710156121c2576121ae600191858561263f565b97906121ba82886104aa565b520195612198565b949250945050604051916121d58361032f565b8252604051916121e48361030f565b8252602082015291565b6020906121f9610411565b8282890101520161218b565b919061221e61222692612216610411565b508285612520565b919093612520565b9190604051916121e48361030f565b6040519061224282610380565b600060608361224f610411565b81528260208201528260408201520152565b60405161226d8161032f565b6000815290565b604051906122818261030f565b600182528160005b602090818110156122ab5760209161229f612235565b90828501015201612289565b505050565b916122bd9061214b6104f9565b909290806001600160f81b03196122d58286866120f5565b35161561237a575061232193916123336122f161232b93612101565b946123176122fd612274565b96612306612235565b5061230f610411565b50828561263f565b8285999299612520565b8285969296612396565b919093612396565b926040519661234188610380565b8752602087015263ffffffff80921660408701521660608501529261236583610478565b525b61236f6103cd565b918252602082015291565b9150506123879150612101565b90612390612261565b90612367565b600093918491905b600483106123ab57505050565b909193946123d560019163ffffff006123c58986886120f5565b3560f81c9160081b161796612101565b9401919061239e565b9392909384519060005b82811061240057500191825260208201526040019150565b80602080928901015181840152016123e8565b1561241a57565b60405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606490fd5b92939091936000925b84519586518510156124d95790600191829786898416156000146124ac576113f7915061248b61249e918a516104aa565b51604051928391602083019589876123de565b519020965b1c93019261245a565b6124bc6124d0916113f7936104aa565b5192604051928391602083019589876123de565b519020966124a3565b9550919350506103cb915015612413565b600093918491905b600283106124ff57505050565b9091939461251760019161ff006123c58986886120f5565b940191906124f2565b600093918491905b6020831061253557505050565b9091939461255960019161254a8885876120f5565b3560f81c9060081b1796612101565b94019190612528565b908160220291602283040361068e57565b9061257d82611a22565b61258a604051918261039b565b828152809261259b601f1991611a22565b0190602036910137565b8051600d101561048557602d0190565b908151811015610485570160200190565b908060601c6125f95760e09063ffffffff90818116610120850152818160201c1661010085015260401c16910152600190565b5050600090565b80519060078210156105335760200151604051906020820192652b30b63ab29d60d11b845260f81b60268301526027820152602781526114058161034a565b90612648610411565b506126616126578483856120f5565b3560f81c93612101565b906006841161268b5761267392612520565b9190600782101561053357604051916121e48361030f565b60405162461bcd60e51b815260206004820152600e60248201526d4241445f56414c55455f5459504560901b6044820152606490fd5b9192600091825b600890818510156127055760019167ffffffffffffff006126fc926126ee8a878b6120f5565b3560f81c921b161796612101565b930192946126c8565b9594505050905056fea26469706673582212209e1e90eed93d01eb67566ec1ef2246277f1f816b166cc5e52e37dc2fc86a878c64736f6c63430008190033",
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

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCaller) GetEndMachineHash(opts *bind.CallOpts, status uint8, globalStateHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "getEndMachineHash", status, globalStateHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntrySession) GetEndMachineHash(status uint8, globalStateHash [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetEndMachineHash(&_OneStepProofEntry.CallOpts, status, globalStateHash)
}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) GetEndMachineHash(status uint8, globalStateHash [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetEndMachineHash(&_OneStepProofEntry.CallOpts, status, globalStateHash)
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

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
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
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220e173c0e1914966f9f0e6a9a456483a1923f81631a1137192ad034e88a4161bc864736f6c63430008190033",
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557612586908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c633604366f1461002757600080fd5b3461122d5736600319016101c0811261122d5760401361122d576001600160401b036044351161122d576101c06044353603600319011261122d5761010036606319011261122d5760403661016319011261122d576101a4356001600160401b03811161122d573660238201121561122d576001600160401b0360048201351161122d5736602482600401358301011161122d576100c56080611313565b60006080526100d26113d9565b60a0526040516100e18161132f565b6000808252602082015260c0526100f66113d9565b60e0526040516101058161132f565b6060815260006020820152610100526040516101208161132f565b600080825260208201819052610120919091526101408190526101608190526101808190526101a08190526101c08190526101e05261015d611420565b5060405161016a81611313565b6004604435810135101561122d57604435600481013582526001600160401b036024909101351161122d576101aa366044356024810135016004016114a1565b60208201526101bd366044803501611579565b60408201526001600160401b03604435608401351161122d576101eb366044356084810135016004016114a1565b60608201526001600160401b0360443560a401351161122d57604060443560a4810135013603600319011261122d576040516102268161132f565b60443560a481013501600401356001600160401b03811161122d573660443560a4810135018201602301121561122d5760443560a4810135018101600401359061026f82611458565b9161027d60405193846113b6565b80835260208301913660443560a481013501820160a08402016024011161122d5760443560a481013501810160240192905b60443560a481013501810160a0840201602401841061123257505050508152602460a46044350135604435010135602082015260808201526102f63660c460443501611579565b60a0820152610104604435013560c0820152610317610124604435016115a1565b60e082015261032b610144604435016115a1565b610100820152610340610164604435016115a1565b6101208201526044356101848101356101408301526101a401356101608201526040519161036d83611365565b6064358352606036608319011261122d5760405161038a81611380565b6084356001600160401b038116900361122d57608435815260a4356001600160401b038116900361122d5760a435602082015260c4356040820152602084015260e435604084015261010435606084015261012435608084015263ffffffff6101443516610144350361122d576101443560a084015261ffff61040b6115b2565b168061108a575060015b908160151461102c5781601414610f7f5781600314610eb85781600514610e625781600414610e535781601314610d125781601214610ce25781601114610cc35781601014610c815781600f14610c495781600e14610c165781600d14610bf35781600c14610bde5781600b14610bd25781600a1461081c5750806009146107fa578060081461077d57806007146106f957806006146106af57806002146106a9576001146104d457634e487b7160e01b600052605160045260246000fd5b600281525b604051906101208252805160048110156106935761012083015261050f60208201516101c06101408501526102e08401906112ad565b90602060408201518051610160860152015161018084015261054560608201519261011f199384868303016101a08701526112ad565b608082015192848203016101c085015282519260408252835180604084015260206060840195019060005b81811061064c575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c08601528051602086015260406020820151600180831b0381511682880152600180831b03602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b5161066a848251611295565b858101516040850152826040820151168285015201511660808201520197019101919091610570565b634e487b7160e01b600052602160045260246000fd5b506104d9565b506106f460808201516106c0611e20565b506106cf600182515114611f8d565b6106d98151611e59565b5190604051906106e88261139b565b60008252525182612296565b6104d9565b50610711602082015161070b836121e9565b90611c57565b6107546107216080830151611fcd565b61073c602084015161070b63ffffffff604085015116612262565b61070b63ffffffff6060602086015193015116612262565b6101843561076a63ffffffff82169182146115d0565b61010082015260006101208201526104d9565b5061078f602082015161070b836121e9565b6107aa602082015161070b63ffffffff60e085015116612262565b6107c5602082015161070b63ffffffff60a086015116612262565b63ffffffff610184356107db8160401c15611c0e565b818160201c1660e08401521661010082015260006101208201526104d9565b5061080c602082015161070b836121e9565b6107c56107216080830151611fcd565b905061083361082e6020840151611cd0565b611d99565b61083b611420565b50606060405161084a8161139b565b52610853611420565b5061085c611401565b506000916000805b60208110610ba25750906109106108ff6108ee6108ba6108986108a996610889611401565b5086600401356024880161200d565b86600498929801356024880161200d565b866004939293013560248801612380565b91604051976108c889611380565b6001600160401b0390811689521660208801526040870152600485013560248601612380565b846004939293013560248601612380565b846004949294013560248601612380565b9290916000939560005b60048110610b6c57509163ffffffff9391610969979695936040519a61093f8c611365565b8b5260208b015260408a0152606089015260808801521660a086015260248160040135910161205a565b5082519060208401518051604060208301519201516040519160208301936626b2b6b7b93c9d60c91b855260018060c01b0319809260c01b16602785015260c01b16602f8301526037820152603781526109c281611380565b5190206040850151606086015160808701519160a088015193604051966626b7b23ab6329d60c91b6020890152602788015260478701526067860152608785015260a784015263ffffffff60e01b9060e01b1660c783015260ab82528160e081011060018060401b0360e084011117610b56578160e0610a87930160405260e08151602083012091610a5582820161132f565b6013828201527226b7b23ab6329036b2b935b632903a3932b29d60691b610100820152019163ffffffff8516906123f7565b61016084015103610b065763ffffffff60a0819382610af294610ab1602089015161070b8a6121e9565b610ac8602089015161070b8460e08c015116612262565b610ade602089015161070b84878d015116612262565b1660e0870152015116826101843516611b27565b1661010082015260006101208201526104d9565b60405162461bcd60e51b815260206004820152602260248201527f43524f53535f4d4f44554c455f494e5445524e414c5f4d4f44554c45535f524f60448201526113d560f21b6064820152608490fd5b634e487b7160e01b600052604160045260246000fd5b9694610b9b60019163ffffff00610b8b898b6004013560248d01611ff2565b3560f81c9160081b161796611ffe565b970161091a565b9093610bcb600191610bbc87866004013560248801611ff2565b3560f81c9060081b1795611ffe565b9101610864565b50506106f48282611b3f565b6106f49150602481600401359101848461160c565b505061018435610c0b63ffffffff82169182146115d0565b6101208201526104d9565b505063ffffffff610c2d61082e6020840151611cd0565b16156104d95761018435610c0b63ffffffff82169182146115d0565b6106f49150610c77906020610c616080860151611fcd565b0151602482600401359201906101843590611f2f565b6020830151611c57565b9050610cbc610c936020840151611cd0565b916020610ca36080860151611fcd565b0192835190602483600401359301916101843590611ed0565b90526104d9565b6106f49150610c77908451602482600401359201906101843590611f2f565b610d0b9150610cf46020840151611cd0565b845190602483600401359301916101843590611ed0565b82526104d9565b505090610d226020830151611cd0565b9163ffffffff610d356020830151611cd0565b9381610d56610d50610d4a6020870151611cd0565b97611d99565b92611d99565b9160405196610d648861134a565b875261018435602088015216604086015216606084015260808101519182515191600183018311610e3d57610d9b60018401611458565b92610da960405194856113b6565b600101808452601f1990610dbc90611458565b0160005b818110610e2657505060005b84518051821015610e015790610de481600193611e7c565b51610def8287611e7c565b52610dfa8186611e7c565b5001610dcc565b5050929093610e1f9082515190610e188286611e7c565b5283611e7c565b50526104d9565b602090610e31611e20565b82828801015201610dc0565b634e487b7160e01b600052601160045260246000fd5b50506106a96020820151611cd0565b5050610e7461082e6020830151611cd0565b610e816020830151611cd0565b63ffffffff610e936020850151611cd0565b921615610ea957506106f4906020830151611c57565b6106f491506020830151611c57565b505061ffff610ec56115b2565b1660418103610f0757506106f460005b6020830151610ef060405192610eea8461132f565b836115c4565b610184356001600160401b03166020830152611c57565b60428103610f1a57506106f46001610ed5565b60438103610f2d57506106f46002610ed5565b604403610f3e576106f46003610ed5565b60405162461bcd60e51b8152602060048201526019602482015278434f4e53545f505553485f494e56414c49445f4f50434f444560381b6044820152606490fd5b505060006020604051610f918161132f565b828152015261800561ffff610fa46115b2565b1603610fc4576106f4610fba6020830151611cd0565b6060830151611c57565b61800661ffff610fd26115b2565b1603610fe8576106f4610c776060830151611cd0565b60405162461bcd60e51b815260206004820152601c60248201527b4d4f56455f494e5445524e414c5f494e56414c49445f4f50434f444560201b6044820152606490fd5b50506020810151600060206040516110438161132f565b82815201525180515180600019810111610e3d576106f49161107f916000602060405161106f8161132f565b8281520152600019019051611e7c565b516020830151611c57565b6001810361109a57506002610415565b600f81036110aa57506006610415565b601081036110ba57506007610415565b61800981036110cb57506008610415565b61800b81036110dc57506009610415565b61800c81036110ed5750600a610415565b61800a81036110fe5750600b610415565b6011810361110e5750600c610415565b618003810361111f5750600d610415565b61800481036111305750600e610415565b602081036111405750600f610415565b6021810361115057506010610415565b6023810361116057506011610415565b6024810361117057506012610415565b618002810361118157506013610415565b601a810361119157506004610415565b601b81036111a157506005610415565b604181101580611222575b156111b957506003610415565b61800581148015611217575b156111d257506014610415565b618008036111e1576015610415565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b6044820152606490fd5b5061800681146111c5565b5060448111156111ac565b600080fd5b60a08436031261122d5760a08060206024946040516112508161134a565b61125a368a61146f565b815260408901358382015261127160608a016115a1565b604082015261128260808a016115a1565b60608201528152019501949250506102af565b80516007811015610693578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b8382106112f157505050505060208091015191015290565b90919293968382826113066001948c51611295565b01980194939201906112d9565b61018081019081106001600160401b03821117610b5657604052565b604081019081106001600160401b03821117610b5657604052565b608081019081106001600160401b03821117610b5657604052565b60c081019081106001600160401b03821117610b5657604052565b606081019081106001600160401b03821117610b5657604052565b602081019081106001600160401b03821117610b5657604052565b601f909101601f19168101906001600160401b03821190821017610b5657604052565b604051906113e68261132f565b60006020836040516113f78161139b565b6060815281520152565b6040519061140e82611380565b60006040838281528260208201520152565b6040519061142d82611365565b600060a08382815261143d611401565b60208201528260408201528260608201528260808201520152565b6001600160401b038111610b565760051b60200190565b919082604091031261122d576040516114878161132f565b80928035600781101561122d578252602090810135910152565b60409291818103841361122d57604051916114bb8361132f565b829481359260018060401b039384811161122d57830191602094858484031261122d57604051936114eb8561139b565b803591821161122d570182601f8201121561122d57803561150b81611458565b9361151960405195866113b6565b818552878086019260061b8401019281841161122d579088809897969594939201915b838310611553575050505050815284520135910152565b97849596979861156783859697949561146f565b8152019201908897969594939261153c565b919082604091031261122d576040516115918161132f565b6020808294803584520135910152565b359063ffffffff8216820361122d57565b6101643561ffff8116810361122d5790565b60078210156106935752565b156115d757565b60405162461bcd60e51b815260206004820152600d60248201526c4241445f43414c4c5f4441544160981b6044820152606490fd5b9091926000936020958684019361162661082e8651611cd0565b60409860608a516116368161139b565b52889989978a809c5b60088092101561167c5760019167ffffffffffffff00611672926116648e8e8e611ff2565b3560f81c921b16179a611ffe565b9c01809c9961163f565b88979d508993999150936116cb959b6116d59399956116c1989e986116ba6116b3859f6116aa90888c612380565b9788919b611ff2565b3595611ffe565b878c61200d565b878c9a929a612380565b878c97929761205a565b8d516d21b0b6361034b73234b932b1ba1d60911b9581019586526001600160c01b031960c085901b8116602e83015260368083018c9052825291966001600160401b039690929161172581611380565b5190206101843503611aea57928f91928f946117ac948d899288519187830193652a30b136329d60d11b855260ff60f81b16602684015260c01b1660278201528b602f820152602f815261177881611380565b51902092712a30b136329036b2b935b632903a3932b29d60711b87519561179e8761132f565b6012875286015216906123f7565b91015103611ab45763ffffffff809c169516851015611aa3576117f090868b8b516117d68161132f565b828152015260608a516117e88161139b565b528388612380565b949092868b8b516118008161132f565b828152015261181d61181387838b611ff2565b3560f81c96611ffe565b60068711611a6e5761183090828a612380565b9190986007881015611a5a578c926118de94926118679261185c8f519b6118568d61132f565b8c6115c4565b858b019c8d5261205a565b5090611872886124d1565b8c51848101916d2a30b136329032b632b6b2b73a1d60911b835288602e830152604e820152604e81526118a48161134a565b51902091792a30b136329032b632b6b2b73a1036b2b935b632903a3932b29d60311b8d51946118d28661132f565b601a86528501526123f7565b03611a225703611a155780516007811015611a0157600403611907575050505050505060029052565b519060078210156119ed57506005036119b957519284841693840361197e575091611973849261070b61012096606060009997611948815161070b8b6121e9565b61196961195860808b0151611fcd565b9561070b8584519289015116612262565b5193015116612262565b166101008201520152565b60649083519062461bcd60e51b8252600482015260156024820152744241445f46554e435f5245465f434f4e54454e545360581b6044820152fd5b825162461bcd60e51b815260048101859052600d60248201526c4241445f454c454d5f5459504560981b6044820152606490fd5b634e487b7160e01b81526021600452602490fd5b634e487b7160e01b83526021600452602483fd5b5050505050505060029052565b865162461bcd60e51b815260048101899052601160248201527010905117d153115351539514d7d493d3d5607a1b6044820152606490fd5b634e487b7160e01b89526021600452602489fd5b8a5162461bcd60e51b8152600481018d9052600e60248201526d4241445f56414c55455f5459504560901b6044820152606490fd5b505050505050505050505060029052565b895162461bcd60e51b8152600481018c9052600f60248201526e10905117d51050931154d7d493d3d5608a1b6044820152606490fd5b5060648f8f519062461bcd60e51b8252600482015260166024820152754241445f43414c4c5f494e4449524543545f4441544160501b6044820152fd5b91909163ffffffff80809416911601918211610e3d57565b9061196991611b7a6020820192611b5a845161070b856121e9565b61070b845160a060e086019663ffffffff98899361070b858b5116612262565b611b876080820151611fcd565b9260608401928184511615611c045761018435828116908103611bbf5782610120956119739382604060009a01511690525116611b27565b60405162461bcd60e51b815260206004820152601d60248201527f4241445f43414c4c45525f494e5445524e414c5f43414c4c5f444154410000006044820152606490fd5b5050600290525050565b15611c1557565b60405162461bcd60e51b815260206004820152601a6024820152794241445f43524f53535f4d4f44554c455f43414c4c5f4441544160301b6044820152606490fd5b518051519160019260018101809111610e3d57611c73906122de565b926000815b611c94575b5050815151611c9091610e188286611e7c565b5052565b83518051821015611cca5781611cac84938493611e7c565b51611cb78289611e7c565b52611cc28188611e7c565b500190611c78565b50611c7d565b90602091604051611ce08161132f565b600093818580935201525191806020604051611cfb8161132f565b8281520152825180516000199391848201918211611d855790611d1d91611e7c565b5192845151908101908111611d7157611d35906122de565b915b8251811015611d6b5780611d4e6001928751611e7c565b51611d598286611e7c565b52611d648185611e7c565b5001611d37565b50925290565b634e487b7160e01b83526011600452602483fd5b634e487b7160e01b84526011600452602484fd5b60208101519051600781101561069357611df157600160201b811015611dc25763ffffffff1690565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606490fd5b60405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606490fd5b60405190611e2d8261134a565b6000606083604051611e3e8161132f565b83815283602082015281528260208201528260408201520152565b805115611e665760200190565b634e487b7160e01b600052603260045260246000fd5b8051821015611e665760209160051b010190565b15611e9757565b60405162461bcd60e51b815260206004820152601160248201527015d493d391d7d3515492d31157d493d3d5607a1b6044820152606490fd5b611f2790611f21611f18611f2c979694959660006020604051611ef28161132f565b82815201526060604051611f058161139b565b52611f108187612101565b91909661205a565b50938585612341565b14611e90565b612341565b90565b90611f2c92611f219160006020604051611f488161132f565b828152015260006020604051611f5d8161132f565b82815201526060604051611f708161139b565b52611f87611f7e8784612101565b9097889461205a565b50612341565b15611f9457565b60405162461bcd60e51b81526020600482015260116024820152700848288beae929c889eaebe988a9c8ea89607b1b6044820152606490fd5b611fee90611fd9611e20565b50611fe8600182515114611f8d565b51611e59565b5190565b90821015611e66570190565b6000198114610e3d5760010190565b9192600091825b600890818510156120515760019167ffffffffffffff006120489261203a8a878b611ff2565b3560f81c921b161796611ffe565b93019294612014565b95945050509050565b90929192606060405161206c8161139b565b5261208561207b858385611ff2565b3560f81c94611ffe565b9061208f85611458565b9261209d60405194856113b6565b858452601f196120ac87611458565b0136602086013760005b60ff811693878510156120e85760ff916120e06120d66001938787612380565b9190919789611e7c565b5201166120b6565b95965050505050604051906120fc8261139b565b815291565b906040519061210f8261132f565b6000928383528360208094015281156121d55760f891813560f81c936006851161219f5785936001938593859392915b83861061216f575050505050509260078310156119ed575061216760405192610eea8461132f565b602082015291565b9091929361219387986121858398998686611ff2565b35861c9060081b1798611ffe565b9601949392919061213f565b60405162461bcd60e51b815260206004820152600e60248201526d4241445f56414c55455f5459504560901b6044820152606490fd5b634e487b7160e01b84526032600452602484fd5b600060206040516121f98161132f565b828152015263ffffffff610120820151169060e061010082015191015191600060206040516122278161132f565b8281520152604051926122398461132f565b6006845263ffffffff60401b9060401b169163ffffffff60201b9060201b161717602082015290565b600060206040516122728161132f565b828152015263ffffffff604051916122898361132f565b6000835216602082015290565b9080516007811015610693576004146122d75780516007811015610693576006036122d75760206122c991015182612516565b156122d15750565b60029052565b5060029052565b906122e882611458565b6040906122f860405191826113b6565b8381528093612309601f1991611458565b019160009060005b84811061231f575050505050565b602090825161232d8161132f565b848152828581830152828701015201612311565b9061234e611f2c936124d1565b906040519261235c8461132f565b60128452712b30b63ab29036b2b935b632903a3932b29d60711b60208501526123f7565b600093918491905b6020831061239557505050565b909193946123b96001916123aa888587611ff2565b3560f81c9060081b1796611ffe565b94019190612388565b9392909384519060005b8281106123e457500191825260208201526040019150565b80602080928901015181840152016123cc565b92939091936000925b845195865185101561248d579060019182978689841615600014612460576124449150612431612452918a51611e7c565b51604051928391602083019589876123c2565b03601f1981018352826113b6565b519020965b1c930192612400565b6124706124849161244493611e7c565b5192604051928391602083019589876123c2565b51902096612457565b9550925050915061249a57565b60405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606490fd5b80519060078210156106935760200151604051906020820192652b30b63ab29d60d11b845260f81b602683015260278201526027815261251081611380565b51902090565b908060601c6125495760e09063ffffffff90818116610120850152818160201c1661010085015260401c16910152600190565b505060009056fea264697066735822122017252787de07d99616bf4eb90f26ec2611c338c66dc828c577e48d70e9bd750864736f6c63430008190033",
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Session) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0CallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverHostIoMetaData contains all meta data concerning the OneStepProverHostIo contract.
var OneStepProverHostIoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BLOBSTREAM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CELESTIA_MESSAGE_HEADER_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557614172908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c9081633604366f14610094575080635fd9e56d146100745763dc3465471461004057600080fd5b3461006f57600036600319011261006f57602060405173c3e209eb245fd59c8586777b499d6a665df3abd28152f35b600080fd5b3461006f57600036600319011261006f57604051606360f81b8152602090f35b3461006f5736600319016101c0811261006f5760401361006f576001600160401b036044351161006f576101c06044353603600319011261006f5761010036606319011261006f5760403661016319011261006f576101a435906001600160401b03821161006f573660238301121561006f576001600160401b0360048301351161006f5736602483600401358401011161006f5761016081610138600093610f64565b82815261014361100f565b602082015260405161015481610f2e565b838152836020820152604082015261016a61100f565b606082015260405161017b81610f2e565b60608152836020820152608082015260405161019681610f2e565b83815283602082015260a08201528260c08201528260e082015282610100820152826101208201528261014082015201526101cf611056565b506040516101dc81610f64565b6004604435810135101561006f57604435600481013582526001600160401b036024909101351161006f5761021c366044356024810135016004016110d7565b602082015261022f3660448035016111af565b60408201526001600160401b03604435608401351161006f5761025d366044356084810135016004016110d7565b60608201526001600160401b0360443560a401351161006f57604060443560a4810135013603600319011261006f5760405161029881610f2e565b60443560a481013501600401356001600160401b03811161006f573660443560a4810135018201602301121561006f5760443560a481013501810160040135906102e18261108e565b916102ef6040519384610fec565b80835260208301913660443560a481013501820160a08402016024011161006f5760443560a481013501810160240192905b60443560a481013501810160a08402016024018410610e2157505050508152602460a46044350135604435010135602082015260808201526103683660c4604435016111af565b60a0820152610104604435013560c0820152610389610124604435016111d7565b60e082015261039d610144604435016111d7565b6101008201526103b2610164604435016111d7565b6101208201526044356101848101356101408301526101a40135610160820152604051916103df83610f9b565b6064358352606036608319011261006f576040516103fc81610f80565b6084356001600160401b038116900361006f57608435815260a4356001600160401b038116900361006f5760a435602082015260c4356040820152602084015260e435604084015261010435606084015261012435608084015263ffffffff6101443516610144350361006f576101443560a084015261ffff6101643516610164350361006f5761801061ffff6101643516101580610e0e575b15610d195760085b80600b14610d0e5780600a14610cfa5780600914610cef5780600714610923578060061461090e578060051461090357806004146108ee57806001146108d4576008146104fb57634e487b7160e01b600052605160045260246000fd5b610503612867565b5061050c612867565b506000906040519361051d85610f2e565b60403686376040519061052f82610f2e565b60403683376000935b600260ff861610156105745761055961056e91856004013560248701613613565b959061056860ff83168a613219565b52613208565b93610538565b935090936000935b600260ff861610156105c45761059d6105be91856004013560248701613646565b95906105ac60ff831689613219565b6001600160401b039091169052613208565b9361057c565b9195935093604051936105d685610f2e565b845260208401526105e68361322a565b60c08501510361089c5761801061ffff610164351614801561088a575b1561080a57610634929161062f9161062691600481013591908290602401611250565b90838787613351565b61322a565b60c08201525b604051906101208252805160048110156107f45761012083015261067060208201516101c06101408501526102e0840190610ec8565b9060206040820151805161016086015201516101808401526106a660608201519261011f199384868303016101a0870152610ec8565b608082015192848203016101c085015282519260408252835180604084015260206060840195019060005b8181106107ad575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c08601528051602086015260406020820151600180831b0381511682880152600180831b03602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b516107cb848251610eb0565b8581015160408501528260408201511682850152015116608082015201970191019190916106d1565b634e487b7160e01b600052602160045260246000fd5b50506101643561ffff166180120361082a578061062f61063492846132e1565b6101643561ffff1661801303610848578061062f6106349284613294565b60405162461bcd60e51b815260206004820152601a602482015279494e56414c49445f474c4f42414c53544154455f4f50434f444560301b6044820152606490fd5b5061801161ffff610164351614610603565b60405162461bcd60e51b815260206004820152601060248201526f4241445f474c4f42414c5f535441544560801b6044820152606490fd5b508060246108e9926004013591018484612230565b61063a565b508060246108e99260040135910184846117b9565b50506001815261063a565b508060246108e99260040135910184846113a0565b5061092c611300565b60405161093881610f49565b6060905260405161094881610f49565b60609052610954611300565b9161016084015190610964611056565b5061096d611056565b50610976611037565b50816109896004830135602484016135dc565b610994969196611037565b506109a790600485013560248601613646565b6109b990600486013560248701613646565b6109cb90600487013560248801613613565b91604051936109d985610f80565b6001600160401b0390811685521660208401526040830152610a0390600486013560248701613613565b909190610a1890600487013560248801613613565b610a2a90600488013560248901613613565b909190610a3f90600489013560248a0161347b565b9490936040519b610a4f8d610f9b565b8c5260208c015260408b015260608a0152608089015263ffffffff1660a0880152610a829060048501356024860161347b565b909490610a97906004860135602487016134f5565b91909780519060208101518051602082015191604001516040519160208301936626b2b6b7b93c9d60c91b8552600160c01b6001900319809260c01b16602785015260c01b16602f830152603782015260378152610af481610f80565b51902090604081015190606081015160808201519160a00151926040519460208601966626b7b23ab6329d60c91b8852602787015260478601526067850152608784015260a783015263ffffffff60e01b9060e01b1660c782015260ab8152610b5c81610fb6565b519020610b67611300565b610b789163ffffffff89168b612d6d565b03610cb457600163ffffffff861601938463ffffffff871611610bf257610b9e85612c3b565b15610c56575050505060018451511b03610c24575b610bc263ffffffff8216612c3b565b15610c085750505180519081600019810111610bf257610be691600019019061135b565b5161016082015261063a565b634e487b7160e01b600052601160045260246000fd5b9163ffffffff610c19931690612ca1565b61016082015261063a565b60405162461bcd60e51b815260206004820152600a6024820152692ba927a723afa622a0a360b11b6044820152606490fd5b610c74939491816024610c6e936004013591016134f5565b50612ca1565b14610bb35760405162461bcd60e51b815260206004820152601360248201527257524f4e475f524f4f545f464f525f5a45524f60681b6044820152606490fd5b60405162461bcd60e51b81526020600482015260136024820152722ba927a723afa927a7aa2fa327a92fa622a0a360691b6044820152606490fd5b50506108e9816112b8565b508060246108e99260040135910183611268565b50506108e9816111e8565b6101643561ffff1661802003610d3057600161049e565b6101643561ffff1661802103610d4757600461049e565b6101643561ffff1661802203610d5e57600561049e565b6101643561ffff1661802303610d7557600661049e565b6101643561ffff1661802403610d8c57600761049e565b6101643561ffff1661803003610da357600961049e565b6101643561ffff1661803103610dba57600a61049e565b6101643561ffff1661803203610dd157600b61049e565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b6044820152606490fd5b5061801361ffff61016435161115610496565b60a08436031261006f576040516001600160401b036080820190811190821117610e9a5760a06020602494836080849501604052610e5f368a6110a5565b8152604089013583820152610e7660608a016111d7565b6040820152610e8760808a016111d7565b6060820152815201950194925050610321565b634e487b7160e01b600052604160045260246000fd5b805160078110156107f4578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b838210610f0c57505050505060208091015191015290565b9091929396838282610f216001948c51610eb0565b0198019493920190610ef4565b604081019081106001600160401b03821117610e9a57604052565b602081019081106001600160401b03821117610e9a57604052565b61018081019081106001600160401b03821117610e9a57604052565b606081019081106001600160401b03821117610e9a57604052565b60c081019081106001600160401b03821117610e9a57604052565b60e081019081106001600160401b03821117610e9a57604052565b608081019081106001600160401b03821117610e9a57604052565b601f909101601f19168101906001600160401b03821190821017610e9a57604052565b6040519061101c82610f2e565b600060208360405161102d81610f49565b6060815281520152565b6040519061104482610f80565b60006040838281528260208201520152565b6040519061106382610f9b565b600060a083828152611073611037565b60208201528260408201528260608201528260808201520152565b6001600160401b038111610e9a5760051b60200190565b919082604091031261006f576040516110bd81610f2e565b80928035600781101561006f578252602090810135910152565b60409291818103841361006f57604051916110f183610f2e565b829481359260018060401b039384811161006f57830191602094858484031261006f576040519361112185610f49565b803591821161006f570182601f8201121561006f5780356111418161108e565b9361114f6040519586610fec565b818552878086019260061b8401019281841161006f579088809897969594939201915b838310611189575050505050815284520135910152565b97849596979861119d8385969794956110a5565b81520192019088979695949392611172565b919082604091031261006f576040516111c781610f2e565b6020808294803584520135910152565b359063ffffffff8216820361006f57565b600019908160a0820151511461124957610184358061122457506101408101918083511461121c5761121a925261290c565b565b506002915052565b91610140820151036112495761124363ffffffff61121a9316826128a0565b5061290c565b6002915052565b9093929384831161006f57841161006f578101920390565b6000199081610140820151036112af5760a0810191825151146112af5783836040611294930151612b45565b518260401161006f5761121a92603f19019160400190612b45565b60029052505050565b6101408101516001016112df576040816112d860a061121a940151612c04565b0151612c04565b60029052565b6001600160401b038111610e9a57601f01601f191660200190565b6040519061130d82610f2e565b601382527226b7b23ab6329036b2b935b632903a3932b29d60691b6020830152565b90610100918203918211610bf257565b600019810191908211610bf257565b91908203918211610bf257565b805182101561136f5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6001019081600111610bf257565b91908201809211610bf257565b906113a9611300565b906101608301519060206113c76113c282870151612dce565b612e97565b91016113da63ffffffff83168251612f1e565b1561179157908686819493519260051c6307ffffff166113f993612f49565b509560405161140781610f49565b6060905260405161141781610f49565b6060815296611424611300565b6101608801518092611434611056565b5061143d611056565b50611446611037565b50611452908886613613565b61145a611037565b50611466908987613646565b611474908a88979397613646565b611482908b89989398613613565b916040519761149089610f80565b6001600160401b03918216895216602088015260408701526114b3908a88613613565b6114be908b89613613565b6114cc908c8a969396613613565b6114da908d8b95939561347b565b919098604051936114ea85610f9b565b845260208401526040830152606082019485526080820192835260a082019763ffffffff16885261151c908c8a61347b565b611527919c8a6134f5565b94909782519360208401518051602082015191604001516040519160208301936626b2b6b7b93c9d60c91b8552600160c01b6001900319809260c01b16602785015260c01b16602f83015260378201526037815261158481610f80565b5190209360400151925190519151926040519460208601966626b7b23ab6329d60c91b8852602787015260478601526067850152608784015260a783015263ffffffff60e01b9060e01b1660c782015260ab81526115e181610fb6565b5190206115ec611300565b6115fd9163ffffffff8c1688612d6d565b03610cb45763ffffffff9b60019860018e8216019d8e911611610bf2578c9561162587612c3b565b1561173357505050505060019051511b03610c24575b60009561164788612c3b565b15611712575085610bf25791869060005b8183116116b5575050509061168c602094939261167e6040519384928884019687612c7e565b03601f198101835282610fec565b5190206101608201525b015190610bf2576116af63ffffffff61121a9316612fc2565b90612ff6565b60405194856116ca8360208301938a85612c7e565b03956116de601f1997888101835282610fec565b5190209461170660405191826116fa6020820195808c88612c7e565b03908101835282610fec565b51902091811c91611658565b9550611728925060209493915086600096612d6d565b610160820152611696565b9091929550611748939b5061175194506134f5565b50978989612ca1565b1461163b5760405162461bcd60e51b815260206004820152601360248201527257524f4e475f524f4f545f464f525f5a45524f60681b6044820152606490fd5b5050505090506002915052565b9082101561136f570190565b6000198114610bf25760010190565b9092939163ffffffff6117d26113c26020850151612dce565b16926117e46113c26020850151612dce565b63ffffffff81169260018060401b036118086118036020880151612dce565b613076565b169761018435801595868061212a575b6121195760208101809111610bf25760208a0151516001600160401b031610801561210d575b6120fd57606060405161185081610f49565b5261186a82846307ffffff8760051c1660208d0151612f49565b999196906001600160f81b031961188282878961179e565b35166120c257611891906117aa565b92849860001461209e57506002976028840190818511610bf257606360f81b806118bc84898b61179e565b3516611fae575b506118d291505b848688611250565b999098600081600314611d0e57506002146118fd57634e487b7160e01b600052605160045260246000fd5b60288a10611cd45760009b60209b8d6008815b101561194e576119218e8e8e61179e565b3560f81c9060081b67ffffffffffffff0016179c61193e906117aa565b60019e909e019d9c60088f611910565b93989d50949990959a61196992979c509d93989d36916121be565b60208151910120600091829084611c27575b6001600160401b038116611b97575b50604051916020830193845260408301526060820152606081526119ad81610fd1565b5190206024356001600160a01b038116929083900361006f576020906024604051809581936316bf557960e01b835260048301525afa918215611b8b57600092611b57575b5003611b1b5760015b15611b0c57858710611ad357611a11868861134e565b9860009a5b602063ffffffff8d16108b8d82611ab9575b505015611a7e57611a688c63ffffffff9283611a5c8e8e611a568f8f611a5190878a1692611393565b611393565b9161179e565b3560f81c921690613126565b9b1663ffffffff8114610bf2576001019a611a16565b9061121a9a50602097508793959b99506116af98506040949650916307ffffff611aac9360051c169061319e565b9201510152015191612fc2565b611acb91925063ffffffff1689611393565b108b8d611a28565b60405162461bcd60e51b81526020600482015260116024820152702120a22fa6a2a9a9a0a3a2afa82927a7a360791b6044820152606490fd5b50506002905295505050505050565b60405162461bcd60e51b81526020600482015260146024820152734241445f534551494e424f585f4d45535341474560601b6044820152606490fd5b9091506020813d602011611b83575b81611b7360209383610fec565b8101031261006f575190386119f2565b3d9150611b66565b6040513d6000823e3d90fd5b6024356001600160a01b038116925082900361006f57611bb8602091613103565b604051636ab8cee160e11b81526001600160401b03909116600482015291829060249082905afa908115611b8b57600091611bf5575b503861198a565b90506020813d602011611c1f575b81611c1060209383610fec565b8101031261006f575138611bee565b3d9150611c03565b926024356001600160a01b0381168103611cd0576020611c4687613103565b6040516316bf557960e01b81526001600160401b03909116600482015291829060249082906001600160a01b03165afa918215611cc4578092611c8c575b50509261197b565b9091506020823d602011611cbc575b81611ca860209383610fec565b81010312611cb95750513880611c84565b80fd5b3d9150611c9b565b604051903d90823e3d90fd5b5080fd5b60405162461bcd60e51b81526020600482015260126024820152712120a22fa9a2a8a4a72127ac2fa82927a7a360711b6044820152606490fd5b93989d92979c959a90506071819c959a92979c10611f7557839183611ed3575b81607111611ecf57611d48366070198401607184016121be565b602081519101208215611ebb578592600190875b60208110611e90575050506050602160405193602085019560ff60f81b823516875260018060601b03199060601b1682860152016035840137608582015260858152611da781610f9b565b5190206040519060208201928352604082015260408152611dc781610f80565b519020906024356001600160a01b0381169190829003611e8c57602090602460405180948193636ab8cee160e11b835260048301525afa928315611cc4578093611e55575b505003611e1a5760016119fb565b60405162461bcd60e51b81526020600482015260136024820152724241445f44454c415945445f4d45535341474560681b6044820152606490fd5b909192506020823d602011611e84575b81611e7260209383610fec565b81010312611cb9575051903880611e0c565b3d9150611e65565b8380fd5b909194611eb3600191611ea488868961179e565b3560f81c9060081b17966117aa565b929101611d5c565b634e487b7160e01b86526032600452602486fd5b8480fd5b91506024356001600160a01b0381168103611ecf576020611ef385613103565b604051636ab8cee160e11b81526001600160401b03909116600482015291829060249082906001600160a01b03165afa908115611f6a578591611f38575b5091611d2e565b90506020813d602011611f62575b81611f5360209383610fec565b81010312611ecf575138611f31565b3d9150611f46565b6040513d87823e3d90fd5b60405162461bcd60e51b81526020600482015260116024820152702120a22fa222a620aca2a22fa82927a7a360791b6044820152606490fd5b9050611fbc8683818a611250565b600092811561208a57823516611fda575b50506118d29150386118c3565b8060011161208657906001611ff49260001901910161379c565b611ffd8161311c565b6002811461204a5761200e8161311c565b60018114612042575b6120208161311c565b1561202d575b8180611fcd565b5060818401809111610bf2576118d290612026565b829150612017565b60405162461bcd60e51b8152602060048201526014602482015273109313d094d51491505357d553911150d251115160621b6044820152606490fd5b8280fd5b634e487b7160e01b84526032600452602484fd5b6001036120b0576118d26003986118ca565b9a505050505050505090506002915052565b60405162461bcd60e51b81526020600482015260136024820152722aa725a727aba72fa4a72127ac2fa82927a7a360691b6044820152606490fd5b5050505094505090506002915052565b50601f8416151561183e565b505050505094505090506003915052565b506004358b1015611818565b1561213d57565b60405162461bcd60e51b81526020600482015260166024820152752aa725a727aba72fa82922a4a6a0a3a2afa82927a7a360511b6044820152606490fd5b3d156121a6573d9061218c826112e5565b9161219a6040519384610fec565b82523d6000602084013e565b606090565b81810292918115918404141715610bf257565b9291926121ca826112e5565b916121d86040519384610fec565b82948184528183011161006f578281602093846000960137010152565b156121fc57565b60405162461bcd60e51b815260206004820152600c60248201526b4241445f505245494d41474560a01b6044820152606490fd5b93909192936122456113c26020830151612dce565b926122566113c26020840151612dce565b9463ffffffff8616601f86161590811591612844575b508015612838575b61282c57606060405161228681610f49565b526122a087826307ffffff8960051c166020860151612f49565b9691909281996060946122be6122b782848661179e565b35916117aa565b9561018435806123a657505060f81c61213d576122fa94816122df93611250565b9390916122ed3686856121be565b60208151910120146121f5565b602063ffffffff831601908163ffffffff841611610bf2578361232e9361233595841161239c575b63ffffffff1691611250565b36916121be565b935b6000965b85518810156123635761235b600191896020818a01015160f81c91613126565b97019661233b565b6116af949750916040602061238d819563ffffffff97956307ffffff61121a9c60051c169061319e565b92015101520151925116612fc2565b9092508290612322565b600181989594969814600014612435575050928092916123cc6123d19560f81c15612136565b611250565b92602060006040518685823780878101838152039060025afa15611b8b576123fc90600051146121f5565b602063ffffffff831601908163ffffffff841611610bf2578361232e9361242f95841161239c5763ffffffff1691611250565b93612337565b909a969594929391906002036127ef57612457936123cc849260f81c15612136565b9190928260201161006f578335036127b35760008060405184868237808581018381520390600a5afa9061248961217b565b911561277a5781511561273c5760408280518101031261006f57604060208301519201517f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff000000018103612701578260051b831590848104602014821715610bf25763ffffffff841610612500575b505050505050612337565b6307ffffff60009b979b9360051c1660015b8581106126da5750506126c457600092839261253391600160201b046121ab565b9060405190602082019260208452602060408401526020898401527f16a2a19edfe81f20d09b681922c813b4b63683508c2280b93829971f439f0d2b608084015260a083015260c082015260c0815261258b81610fb6565b519060055afa61259961217b565b901561268f57602081510361265457602081519101519060208110612642575b508160401161006f5760208301350361260957821161006f57604051916125df83610f2e565b60208352369082011161006f576040602091018183013760006040820152933880808080806124f5565b60405162461bcd60e51b815260206004820152601160248201527025ad23afa82927a7a32faba927a723afad60791b6044820152606490fd5b6000199060200360031b1b16386125b9565b60405162461bcd60e51b815260206004820152601360248201527209a9e888ab0a0beaea49e9c8ebe988a9c8ea89606b1b6044820152606490fd5b60405162461bcd60e51b815260206004820152600d60248201526c1353d111561417d19052531151609a1b6044820152606490fd5b634e487b7160e01b600052601260045260246000fd5b909360011b936001808216146126f7575b60011c9060011b612512565b93600117936126eb565b60405162461bcd60e51b8152602060048201526013602482015272554e4b4e4f574e5f424c535f4d4f44554c555360681b6044820152606490fd5b60405162461bcd60e51b81526020600482015260166024820152754b5a475f505245434f4d50494c455f4d495353494e4760501b6044820152606490fd5b60405162461bcd60e51b815260206004820152601160248201527024a72b20a624a22fa5ad23afa82927a7a360791b6044820152606490fd5b60405162461bcd60e51b8152602060048201526014602482015273096b48ebea0a49e9e8cbeaea49e9c8ebe9082a6960631b6044820152606490fd5b60405162461bcd60e51b8152602060048201526015602482015274554e4b4e4f574e5f505245494d4147455f5459504560581b6044820152606490fd5b50506002905292505050565b50601f86161515612274565b905060208101809111610bf2576020830151516001600160401b0316103861226c565b6040519061287482610f2e565b8160405161288181610f2e565b6040368237815260206040519161289783610f2e565b60403684370152565b61014081019160001980845103612903576101209163ffffffff60401b60e085015160401b1663ffffffff60201b61010086015160201b16179363ffffffff80948192015116911601828111610bf257821601818111610bf25716179052600190565b50505050600090565b604080820180515160a0840180515194959091600019808714908115612ad6575b50612aca576080870193845196602095868901519960009a5b8a5180518d1015612a09579061295e8d60019361135b565b518a61296a8251613afa565b918d828201516060828401519301519151938401946b29ba30b1b590333930b6b29d60a11b8652602c850152604c84015263ffffffff60e01b918260e091821b16606c8501521b166070820152605481526129c481610fd1565b519020908b51908b8201927129ba30b1b590333930b6b29039ba30b1b59d60711b845260328301526052908183015281526129fe81610fd1565b5190209b019a612946565b50919497995091949950889295515201928351908782015191805151516000915b898b838510612a5e57505050505090606093929151525190868201525251928301525190612a5782610f49565b6060825252565b612a878584959698936000612a8d94600197519251612a7c81610f2e565b82815201525161135b565b51613afa565b908b51908d8201926b2b30b63ab29039ba30b1b59d60a11b8452602c830152604c90818301528152612abe81610fd1565b51902094019190612a2a565b50600290955250505050565b905081143861292d565b90916049926831b7ba343932b0b21d60b91b8352600983015260298201520190565b15612b0957565b60405162461bcd60e51b815260206004820152601460248201527357524f4e475f434f5448524541445f454d50545960601b6044820152606490fd5b90600092600091825b60208110612bea575090612b629291613613565b506000198303612b8c57612b768115612b02565b612b84602083015115612b02565b602082015252565b6040516020810190612ba38161167e858886612ae0565b519020602083015114612b845760405162461bcd60e51b8152602060048201526012602482015271057524f4e475f434f5448524541445f504f560741b6044820152606490fd5b9294612bfd600191611ea488868661179e565b9301612b4e565b600090805182198103612c15575052565b602082019061167e612c338351604051928391602083019586612ae0565b519020905252565b8015159081612c48575090565b60001981019150808211610bf257161590565b60005b838110612c6e5750506000910152565b8181015183820152602001612c5e565b602090612c946040959382815194859201612c5b565b0191825260208201520190565b91926000936000925b8451958651851015612d29579060019182978689841615600014612cfc5761167e9150612cdb612cee918a5161135b565b5160405192839160208301958987612c7e565b519020965b1c930192612caa565b612d0c612d209161167e9361135b565b519260405192839160208301958987612c7e565b51902096612cf3565b95509250509150612d3657565b60405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606490fd5b92939091936000925b8451958651851015612d29579060019182978689841615600014612db55761167e9150612cdb612da7918a5161135b565b519020965b1c930192612d76565b612d0c612dc59161167e9361135b565b51902096612dac565b90602091604051612dde81610f2e565b600093818580935201525191806020604051612df981610f2e565b8281520152825180516000199391848201918211612e835790612e1b9161135b565b5192845151908101908111612e6f57612e3390613579565b915b8251811015612e695780612e4c600192875161135b565b51612e57828661135b565b52612e62818561135b565b5001612e35565b50925290565b634e487b7160e01b83526011600452602483fd5b634e487b7160e01b84526011600452602484fd5b6020810151905160078110156107f457612eef57600160201b811015612ec05763ffffffff1690565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606490fd5b60405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606490fd5b6020820190818311610bf257516001600160401b031610159081612f40575090565b601f9150161590565b93909291936060604051612f5c81610f49565b526040612f82612f78612f6f86896135dc565b909687996134f5565b959095968661319e565b91015103612f8c57565b60405162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d3515357d493d3d560921b6044820152606490fd5b60006020604051612fd281610f2e565b828152015263ffffffff60405191612fe983610f2e565b6000835216602082015290565b518051519160019260018101809111610bf25761301290613579565b926000815b61303a575b50508151516130369161302f828661135b565b528361135b565b5052565b8351805182101561307057816130528493849361135b565b5161305d828961135b565b52613068818861135b565b500190613017565b5061301c565b6020810151905160078110156107f4576001036130d457600160401b8110156130a5576001600160401b031690565b60405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b6044820152606490fd5b60405162461bcd60e51b81526020600482015260076024820152661393d517d24d8d60ca1b6044820152606490fd5b6001600160401b039081166000190191908211610bf257565b600311156107f457565b909160208310156131615782601f0392601f8411610bf2578360031b93840460081490601f141715610bf25760ff809116831b921b19161790565b60405162461bcd60e51b81526020600482015260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b6044820152606490fd5b906132059260405160208101916b26b2b6b7b93c903632b0b31d60a11b8352602c820152602c81526131cf81610f80565b51902090604051926131e084610f2e565b601384527226b2b6b7b93c9036b2b935b632903a3932b29d60691b6020850152612d6d565b90565b60ff1660ff8114610bf25760010190565b90600281101561136f5760051b0190565b8051906020808351930151910151602081519101516040519260208401946c23b637b130b61039ba30ba329d60991b8652602d850152604d84015260018060c01b0319809260c01b16606d84015260c01b166075820152605d815261328e81610fd1565b51902090565b602081019163ffffffff6132b86113c26132b16118038751612dce565b9551612dce565b169160028310156112af57509060206132d2920151613219565b6001600160401b039091169052565b602081019163ffffffff6132f86113c28551612dce565b169160028310156112af575061331e61121a935192602060018060401b03930151613219565b5116906000602060405161333181610f2e565b82815201526040519161334383610f2e565b600183526020830152612ff6565b936113c294939293602081019061336b6113c28351612dce565b9161337d63ffffffff98899251612dce565b1693600285101561346e57602061339991019783168851612f1e565b156134625750906307ffffff6133c59260051c169360606040516133bc81610f49565b52848751612f49565b949190506101643561ffff811680910361006f576180108103613403575050916133f6604094926133fd9451613219565b519161319e565b91510152565b919450945061801191925014600014613423576134209151613219565b52565b60405162461bcd60e51b81526020600482015260176024820152764241445f474c4f42414c5f53544154455f4f50434f444560481b6044820152606490fd5b60029052505050505050565b5060029052505050505050565b600093918491905b6004831061349057505050565b909193946134ba60019163ffffff006134aa89868861179e565b3560f81c9160081b1617966117aa565b94019190613483565b906134cd8261108e565b6134da6040519182610fec565b82815280926134eb601f199161108e565b0190602036910137565b90929192606060405161350781610f49565b5261352061351685838561179e565b3560f81c946117aa565b9061352a856134c3565b9260005b60ff811693878510156135605760ff9161355861354e6001938787613613565b919091978961135b565b52011661352e565b959650505050506040519061357482610f49565b815291565b906135838261108e565b6040906135936040519182610fec565b83815280936135a4601f199161108e565b019160009060005b8481106135ba575050505050565b60209082516135c881610f2e565b8481528285818301528287010152016135ac565b90916000926000926000915b602083106135f557505050565b9091939461360a600191611ea488858761179e565b940191906135e8565b600093918491905b6020831061362857505050565b9091939461363d600191611ea488858761179e565b9401919061361b565b9192600091825b6008908185101561368a5760019167ffffffffffffff00613681926136738a878b61179e565b3560f81c921b1617966117aa565b9301929461364d565b95945050509050565b9081602091031261006f57516001600160401b038116810361006f5790565b919082604091031261006f576040516136ca81610f2e565b9182908035906001600160f81b03198216820361006f57908252602001359063ffffffff198216820361006f5760200152565b91909160608184031261006f576040519061371782610f80565b909283919081356001600160401b03811161006f57820181601f8201121561006f5760209181356137478161108e565b926137556040519485610fec565b818452848085019260051b82010192831161006f578401905b82821061378d5750505083528082013590830152604090810135910152565b8135815290840190840161376e565b8160081161006f57803560c01c906040908151926303f16d4b60e11b92838552602060049473c3e209eb245fd59c8586777b499d6a665df3abd29082888881855afa978815613aef57600098613ad0575b506001600160401b039788166103e801888111613abb5788168511613aad5782908785518094819382525afa908115613aa257908791600091613a75575b50168311613a69578660581161006f578387018481039360571993916101008587011261006f5760588701356001600160a01b0381160361006f5760a06077198097011261006f5782519461387f86610f80565b61388c8260788a016136b2565b865261389b8260b88a016136b2565b8587015260f8880135848701526101188801358a811161006f578260586138c4928b01016136fd565b966101388901358b811161006f5789019060808285039384011261006f5785908151936138f085610f80565b60588401358552011261006f5784519661390988610f2e565b6078820135885260988201358789015286830197885260b88201358c811161006f5761396094605861393e928c9501016136fd565b86840152600086805161395081610f80565b60608152828a8201520152613b57565b509351908382519201519415613a3657036139ff5750506038830135036139f55761398a90613b39565b509384156139f5578060101161006f5760181161006f576008601082013560c01c91013560c01c01908282116139e057506139c76139ce91613103565b92806121ab565b911610156139db57600090565b600190565b601190634e487b7160e01b6000525260246000fd5b5050505050600190565b5162461bcd60e51b81528086019190915260116024820152704d69736d6174636865644865696768747360781b6044820152606490fd5b825162461bcd60e51b8152808901859052600d60248201526c24a72b20a624a22fa82927a7a360991b6044820152606490fd5b50505050505050600290565b613a959150833d8511613a9b575b613a8d8183610fec565b810190613693565b3861382b565b503d613a83565b83513d6000823e3d90fd5b505050505050505050600190565b601188634e487b7160e01b6000525260246000fd5b613ae8919850833d8511613a9b57613a8d8183610fec565b96386137ed565b84513d6000823e3d90fd5b80519060078210156107f45760200151604051906020820192652b30b63ab29d60d11b845260f81b602683015260278201526027815261328e81610f80565b6040015160038116613b4e5760021c90600090565b50600090600890565b9082519260209081808201956040875193015182604051948593631f3302a960e01b8552600485015280516024850152015160448301526080606483015260e482019080519160606084850152825180915284610104850193019060005b818110613c54575050508381015160a48401526040015160c48301528190038173c3e209eb245fd59c8586777b499d6a665df3abd25afa908115611b8b57600091613c1e575b5015613c1257613c0e9351015191613c6d565b9091565b50505050600090600490565b8281813d8311613c4d575b613c338183610fec565b81010312611cd05751908115158203611cb9575038613bfb565b503d613c29565b8251855288968896509485019490920191600101613bb5565b90613cbd92613c7c8351613cd3565b6040613c8b6020860151613cd3565b940151906040519462ffffff19809216602087015216603d850152605a840152605a8352613cb883610fd1565b613d46565b5015613ccb57600190600090565b600090600390565b80516020918201516040516001600160f81b031990921692820192835263ffffffff19166021820152601d8152613d0981610f2e565b51905162ffffff19918282169190601d8110613d27575b5050905090565b83919250601d0360031b1b1616803880613d20565b600611156107f457565b60408201805193949360018111613e455750825151613e38575b6020830190815181511115613e2a57602060405196613dbc82890189613dab602160009c8d96878652613d9b815180928b8686019101612c5b565b8101036001810184520182610fec565b604051928392839251928391612c5b565b8101039060025afa15611f6a578551935191825115613e135790613de593949151905190613f1c565b6006811015613dff5780613df95750149190565b92915050565b634e487b7160e01b85526021600452602485fd5b5051600114159050613e2457149190565b50508190565b505050509050600090600290565b5050509050600090600190565b613e56845151916020860151613e68565b14613d60575050509050600090600190565b9060019081811115613f145760005b8183821b10613ef357610100908103908111610bf257613e969061132f565b9282613ea18561133f565b1b90613eac8261133f565b8111613eb9575050505090565b9293509091838203613ecb5750505090565b613eea935090613ede81613ee49361134e565b9261134e565b90613e68565b61320590611385565b82810180911115613e7757634e487b7160e01b600052601160045260246000fd5b505050600090565b93929381156140035760018214613fee57845115613fe35782613f3e8361400e565b613f51613f4b885161133f565b88614059565b92818110613fae5781613ede613f719693613f6b9361134e565b90613f1c565b9091613f7c82613d3c565b81613fa7575050613f9b83613f95613fa194955161133f565b9061135b565b516140fa565b90600090565b9350919050565b613fb89450613f1c565b9091613fc382613d3c565b81613fa7575050613fdc83613f95613fa194955161133f565b51906140fa565b505090915090600590565b5050909151613ffd5790600090565b90600490565b505090915090600390565b600180821061006f578180916000925b6140445750506000198101908111610bf2576001901b90811461403e5790565b60011c90565b909161404f906117aa565b91811c908161401e565b91908251811161409d5761406c816134c3565b9060005b81811061407e575090925050565b8061408b6001928761135b565b51614096828661135b565b5201614070565b60405162461bcd60e51b815260206004820152602f60248201527f496e76616c69642072616e67653a205f626567696e206f72205f656e6420617260448201526e65206f7574206f6620626f756e647360881b6064820152608490fd5b6141296000916020936040519085820192600160f81b84526021830152604182015260418152613dab81610fd1565b8101039060025afa15611b8b576000519056fea26469706673582212203d4700133e2183cc4ab46e2f66a93efd3399e5376f064a521620ad2c3bec79a264736f6c63430008190033",
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCallerSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// OneStepProverMathMetaData contains all meta data concerning the OneStepProverMath contract.
var OneStepProverMathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611ede908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c633604366f1461002757600080fd5b34610dd35736600319016101c08112610dd357604013610dd3576001600160401b0360443511610dd3576101c060443536036003190112610dd357610100366063190112610dd357604036610163190112610dd3576101a4356001600160401b038111610dd35736602382011215610dd35760048101356001600160401b038111610dd35736910160240111610dd35760006101606100c4610ee5565b8281526100cf610fa6565b60208201526100dc610f05565b83815283602082015260408201526100f2610fa6565b60608201526100ff610f05565b606081528360208201526080820152610116610f05565b83815283602082015260a08201528260c08201528260e08201528261010082015282610120820152826101408201520152600060a0610153610f24565b82815261015e610f43565b8381528360208201528360408201526020820152826040820152826060820152826080820152015261018e610ee5565b60046044358101351015610dd357604435600481013582526001600160401b0360249091013511610dd3576101ce3660443560248101350160040161100c565b60208201526101e13660448035016110d3565b60408201526001600160401b036044356084013511610dd35761020f3660443560848101350160040161100c565b60608201526001600160401b0360443560a4013511610dd357604060443560a48101350136036003190112610dd357610246610f05565b60443560a481013501600401356001600160401b038111610dd3573660443560a48101350182016023011215610dd35760443560a4810135018101600401359061029761029283610fc6565b610f81565b8281529160208301913660443560a481013501820160a084020160240111610dd35760443560a481013501810160240192905b60443560a481013501810160a08402016024018410610dd857505050508152602460a46044350135604435010135602082015260808201526103113660c4604435016110d3565b60a0820152610104604435013560c0820152610332610124604435016110f7565b60e0820152610346610144604435016110f7565b61010082015261035b610164604435016110f7565b6101208201526044356101848101356101408301526101a40135610160820152610383610f24565b6064358152906060366083190112610dd35761039d610f43565b6084356001600160401b0381169003610dd357608435815260a4356001600160401b0381169003610dd35760a435602082015260c4356040820152602083015260e435604083015261010435606083015261012435608083015263ffffffff61014435166101443503610dd3576101443560a083015261ffff61041e611108565b16604581148015610dc9575b15610c47575060015b80600b14610b1f5780600a1461096a578060091461092157806008146108f957806007146108ef57806005146108a45780600314610858578060061461084e57806004146108105780600214610733576001146104a057634e487b7160e01b600052605160045260246000fd5b6104ad60208201516115d9565b604561ffff6104ba611108565b16036106d25780519060078210156106b2576104d8602092156115a3565b01516106c8576104f860015b6104f2602084015191611831565b9061169a565b604051906101208252805160048110156106b25761012083015261052e60208201516101c06101408501526102e0840190610e7f565b90602060408201518051610160860152015161018084015261056460608201519261011f199384868303016101a0870152610e7f565b608082015192848203016101c085015282519260408252835180604084015260206060840195019060005b81811061066b575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c08601528051602086015260406020820151600180831b0381511682880152600180831b03602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b51610689848251610e67565b85810151604085015282604082015116828501520151166080820152019701910191909161058f565b634e487b7160e01b600052602160045260246000fd5b6104f860006104e4565b605061ffff6106df611108565b16036107045780519060078210156106b2576106ff60016020931461156d565b6104d8565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa2a8ad60c91b6044820152606490fd5b5061074961074460208301516115d9565b61171a565b61075961074460208401516115d9565b60451961ffff610767611108565b160161ffff81116107fa576107c5926107b79261ffff8316600281149081156107ef575b81156107e4575b81156107d9575b50156107ca576107ab6107b191611778565b91611778565b90611c7c565b6104f2602084015191611dcb565b6104f8565b63ffffffff9182169116611c7c565b600891501438610799565b600681149150610792565b60048114915061078b565b634e487b7160e01b600052601160045260246000fd5b5061082161074460208301516115d9565b60661961ffff61082f611108565b16019061ffff82116107fa576107c59163ffffffff6104e49216611b25565b506107c58161135e565b5061086e61086960208301516115d9565b6117ce565b61087e61086960208401516115d9565b60501961ffff61088c611108565b160161ffff81116107fa576107c5926107b792611c7c565b506108b561086960208301516115d9565b60781961ffff6108c3611108565b160161ffff81116107fa576108e063ffffffff916107c5936119c2565b166104f260208401519161179e565b506107c58161111a565b506107c563ffffffff61091261086960208501516115d9565b166104f2602084015191611831565b506107c561093561074460208401516115d9565b60ac61ffff610942611108565b160361095f5761095190611778565b6104f260208401519161179e565b63ffffffff16610951565b5060c061ffff610978611108565b1603610a6f57600860005b60078110156106b25780610a625763ffffffff5b6109a460208501516115d9565b91825160078110156106b25703610a2157600160ff84161b806000198101116107fa57600019810160208401511680602085015260ff60001981871601116107fa57600160ff6107c59681600019911601161b16610a09575b5050602083015161169a565b600019011916602082015117602082015238806109fd565b60405162461bcd60e51b81526020600482015260196024820152784241445f455854454e445f53414d455f545950455f5459504560381b6044820152606490fd5b6001600160401b03610997565b60c161ffff610a7c611108565b1603610a8b5760106000610983565b60c261ffff610a98611108565b1603610aa75760086001610983565b60c361ffff610ab4611108565b1603610ac35760106001610983565b60c461ffff610ad0611108565b1603610adf5760206001610983565b60405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f455854454e445f53414d455f5459504560401b6044820152606490fd5b5060bc61ffff610b2d611108565b1603610bb857600060025b610b4560208401516115d9565b90815160078110156106b25760078210156106b25703610b785760078210156106b2576107c5918152602083015161169a565b60405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f5245494e544552505245545f5459504560401b6044820152606490fd5b60bd61ffff610bc5611108565b1603610bd45760016003610b38565b60be61ffff610be1611108565b1603610bf05760026000610b38565b60bf61ffff610bfd611108565b1603610c0c5760036001610b38565b60405162461bcd60e51b81526020600482015260136024820152721253959053125117d491525395115494149155606a1b6044820152606490fd5b604681101580610dbe575b15610c5f57506002610433565b606781101580610db3575b15610c7757506004610433565b606a81101580610da8575b15610c8f57506006610433565b605181101580610d9d575b15610ca757506003610433565b607981101580610d92575b15610cbf57506005610433565b607c81101580610d87575b15610cd757506007610433565b60a78103610ce757506008610433565b60ac81148015610d7d575b15610cff57506009610433565b60c081101580610d72575b15610d175750600a610433565b60bc8110159081610d66575b5015610d3057600b610433565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b6044820152606490fd5b60bf9150111538610d23565b5060c4811115610d0a565b5060ad8114610cf2565b50608a811115610cca565b50607b811115610cb2565b50605a811115610c9a565b506078811115610c82565b506069811115610c6a565b50604f811115610c52565b506050811461042a565b600080fd5b60a084360312610dd3576040516001600160401b036080820190811190821117610e515760a06020602494836080849501604052610e16368a610fdd565b8152604089013583820152610e2d60608a016110f7565b6040820152610e3e60808a016110f7565b60608201528152019501949250506102ca565b634e487b7160e01b600052604160045260246000fd5b805160078110156106b2578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b838210610ec357505050505060208091015191015290565b9091929396838282610ed86001948c51610e67565b0198019493920190610eab565b6040519061018082016001600160401b03811183821017610e5157604052565b60408051919082016001600160401b03811183821017610e5157604052565b6040519060c082016001600160401b03811183821017610e5157604052565b60405190606082016001600160401b03811183821017610e5157604052565b60405190602082016001600160401b03811183821017610e5157604052565b6040519190601f01601f191682016001600160401b03811183821017610e5157604052565b610fae610f05565b90610fb7610f62565b60608152825260006020830152565b6001600160401b038111610e515760051b60200190565b9190826040910312610dd357610ff1610f05565b918035906007821015610dd357602091845201356020830152565b9190604092604081830312610dd357611023610f05565b936001600160401b03928235848111610dd3578301916020948584840312610dd35761104d610f62565b938035918211610dd3570182601f82011215610dd357803561107161029282610fc6565b93878086848152019260061b84010192818411610dd3579088809897969594939201915b8383106110ad57505050505081528552013590830152565b9784959697986110c1838596979495610fdd565b81520192019088979695949392611095565b9190826040910312610dd35760206110e9610f05565b928035845201356020830152565b359063ffffffff82168203610dd357565b6101643561ffff81168103610dd35790565b6020810161112b61086982516115d9565b9161113961086983516115d9565b9060006101643561ffff9081811680910361135a57607b19019181831161134657508116600381036111eb57506001600160401b03948516919050811580156111c6575b6111bc575060070b9283156111a6576111a493816104f2931660070b0516915b519161179e565b565b634e487b7160e01b600052601260045260246000fd5b6002905250505050565b5082851660070b60016001603f1b031914801561117d57506000198260070b1461117d565b6005810361122857506001600160401b0394851691905081156111bc575060070b9283156111a6576111a493816104f2931660070b07169161119d565b600a8103611251575050506104f2906111a493603f60018060401b0380931691161b169161119d565b600c810361127a575050506104f2906111a493603f60018060401b0380931691161c169161119d565b600b81036112a6575050506104f2906111a493603f60018060401b0380931660070b91161d169161119d565b600d81036112e357506111a4946104f293603f90911692506001600160401b03915081168181816112d68661196a565b161c16921b16179161119d565b600e0361131e57506111a4936104f292603f90911691506001600160401b039081168181816113118661196a565b161b16921c16179161119d565b90929361132a9261185d565b91909161133e57506104f26111a49261119d565b600290525050565b634e487b7160e01b81526011600452602490fd5b8280fd5b6020810161136f61074482516115d9565b61137c61074483516115d9565b906000936101643561ffff90818116809103611569576069190195818711611346575060039590811686810361140f57505063ffffffff80921690811580156113ef575b6113e45750840b80156111a6576111a494826104f29416900b0516915b5191611831565b600290525050505050565b50838316860b637fffffff191480156113c0575081860b600019146113c0565b6005810361144557505063ffffffff8092169081156113e45750840b80156111a6576111a494826104f29416900b0716916113dd565b929592600a810361147057505050506104f2906111a493601f63ffffffff80931691161b16916113dd565b600c810361149857505050506104f2906111a493601f63ffffffff80931691161c16916113dd565b600b81036114c257505050906111a493601f6104f29363ffffffff809416900b91161d16916113dd565b91925090600d81036114fe575050506104f290601f6111a494169063ffffffff8091168181816114f186611997565b161c16921b1617916113dd565b600e036115345750506104f290601f6111a494169063ffffffff80911681818161152786611997565b161b16921c1617916113dd565b611549919263ffffffff80809716911661185d565b91909161156057506111a4926104f29116916113dd565b60029052505050565b8680fd5b1561157457565b60405162461bcd60e51b81526020600482015260076024820152661393d517d24d8d60ca1b6044820152606490fd5b156115aa57565b60405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606490fd5b906020916115e5610f05565b6000938185809352015251918060206115fc610f05565b8281520152825180516000199391848201918211611686579061161e91611e2e565b51928451519081019081116116725761163690611e58565b915b825181101561166c578061164f6001928751611e2e565b5161165a8286611e2e565b526116658185611e2e565b5001611638565b50925290565b634e487b7160e01b83526011600452602483fd5b634e487b7160e01b84526011600452602484fd5b5180515191600192600181018091116107fa576116b690611e58565b926000815b6116de575b50508151516116da916116d38286611e2e565b5283611e2e565b5052565b8351805182101561171457816116f684938493611e2e565b516117018289611e2e565b5261170c8188611e2e565b5001906116bb565b506116c0565b6020810151905160078110156106b25761173490156115a3565b600160201b8110156117495763ffffffff1690565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606490fd5b6380000000811661178c5763ffffffff1690565b63ffffffff1663ffffffff60201b1790565b600060206117aa610f05565b82815201526117b7610f05565b600181526001600160401b03909116602082015290565b6020810151905160078110156106b25760016117ea911461156d565b600160401b811015611802576001600160401b031690565b60405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b6044820152606490fd5b6000602061183d610f05565b828152015263ffffffff61184f610f05565b916000835216602082015290565b909161ffff16806118795750016001600160401b031690600090565b600181036118925750036001600160401b031690600090565b600281036118ab5750026001600160401b031690600090565b600481036118da57506001600160401b03918216919082156118cf57160490600090565b505050600090600190565b600681036118fe57506001600160401b03918216919082156118cf57160690600090565b6007810361190e57501690600090565b6008810361191e57501790600090565b60090361192c571890600090565b60405162461bcd60e51b81526020600482015260166024820152750494e56414c49445f47454e455249435f42494e5f4f560541b6044820152606490fd5b6001600160401b03908116604003919082116107fa57565b90600163ffffffff809316019182116107fa57565b9063ffffffff8092166020039182116107fa57565b63ffffffff90811660001901919082116107fa57565b9061ffff1680611a2d575060405b63ffffffff80821615159081611a0b575b50156119f5576119f0906119ac565b6119d0565b63ffffffff908116604003915081116107fa5790565b6001600160401b039150600190611a21846119ac565b161b83161615386119e1565b6001918291808303611a8b57506000925b611a4757505090565b63ffffffff8316604081109081611a74575b5015611a6f57611a698293611982565b92611a3e565b505090565b83901b82166001600160401b031615905038611a59565b600291925014611ac65760405162461bcd60e51b815260206004820152600960248201526804241442049556e4f760bc1b6044820152606490fd5b81906000906000935b611ada575b50505090565b63ffffffff809216916040831015611b1f5783831b82166001600160401b0316611b0f575b82146107fa578280920191611acf565b93611b1990611982565b93611aff565b50611ad4565b9061ffff1680611b90575060205b63ffffffff80821615159081611b6e575b5015611b5857611b53906119ac565b611b33565b63ffffffff908116602003915081116107fa5790565b6001600160401b039150600190611b84846119ac565b161b8316161538611b44565b6001918291808303611be957506000925b611baa57505090565b63ffffffff8316602081109081611bd2575b5015611a6f57611bcc8293611982565b92611ba1565b83901b82166001600160401b031615905038611bbc565b600291925014611c245760405162461bcd60e51b815260206004820152600960248201526804241442049556e4f760bc1b6044820152606490fd5b81906000906000935b611c375750505090565b63ffffffff809216916020831015611b1f5783831b82166001600160401b0316611c6c575b82146107fa578280920191611c2d565b93611c7690611982565b93611c5c565b9161ffff1680611c9857506001600160401b0391821691161490565b60018103611cb357506001600160401b039182169116141590565b60028103611cd2575060018060401b0380911660070b911660070b1290565b60038103611cec57506001600160401b0390811691161090565b60048103611d0b575060018060401b0380911660070b911660070b1390565b60058103611d2557506001600160401b0390811691161190565b60068103611d45575060018060401b0380911660070b911660070b131590565b60078103611d6057506001600160401b039081169116111590565b60088103611d80575060018060401b0380911660070b911660070b121590565b600903611d99576001600160401b039081169116101590565b60405162461bcd60e51b815260206004820152600a6024820152690424144204952454c4f560b41b6044820152606490fd5b602090611dd6610f05565b6000808252920182905215611e0a57806020611df0610f05565b8281520152611dfd610f05565b9081526001602082015290565b806020611e15610f05565b8281520152611e22610f05565b90808252602082015290565b8051821015611e425760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b90611e6561029283610fc6565b8281528092611e76601f1991610fc6565b01906000805b838110611e895750505050565b602090611e94610f05565b838152828481830152828601015201611e7c56fea2646970667358221220dd1bd418a29c9c265d668903e03c19d127a7c2a15075ef1a4054ec03144bd2bc64736f6c63430008190033",
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverMemoryMetaData contains all meta data concerning the OneStepProverMemory contract.
var OneStepProverMemoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460155761198b908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c633604366f1461002857600080fd5b3461083e5736600319016101c0811261083e5760401361083e576001600160401b036044351161083e576101c06044353603600319011261083e5761010036606319011261083e5760403661016319011261083e576101a435906001600160401b03821161083e573660238301121561083e576001600160401b0360048301351161083e5736602483600401358401011161083e57610160816100cc600093610950565b8281526100d76109fb565b60208201526040516100e88161096c565b83815283602082015260408201526100fe6109fb565b606082015260405161010f8161096c565b60608152836020820152608082015260405161012a8161096c565b83815283602082015260a08201528260c08201528260e08201528261010082015282610120820152826101408201520152600060a060405161016b81610987565b82815260405161017a816109a2565b838152836020820152836040820152602082015282604082015282606082015282608082015201526040516101ae81610950565b6004604435810135101561083e57604435600481013582526001600160401b036024909101351161083e576101ee36604435602481013501600401610a6c565b6020820152610201366044803501610b44565b60408201526001600160401b03604435608401351161083e5761022f36604435608481013501600401610a6c565b60608201526001600160401b0360443560a401351161083e57604060443560a4810135013603600319011261083e5760405161026a8161096c565b60443560a481013501600401356001600160401b03811161083e573660443560a4810135018201602301121561083e5760443560a481013501810160040135906102b382610a23565b916102c160405193846109d8565b80835260208301913660443560a481013501820160a08402016024011161083e5760443560a481013501810160240192905b60443560a481013501810160a0840201602401841061084357505050508152602460a460443501356044350101356020820152608082015261033a3660c460443501610b44565b60a0820152610104604435013560c082015261035b61012460443501610b6c565b60e082015261036f61014460443501610b6c565b61010082015261038461016460443501610b6c565b6101208201526044356101848101356101408301526101a40135610160820152604051916103b183610987565b6064358352606036608319011261083e576040516103ce816109a2565b6084356001600160401b038116900361083e57608435815260a4356001600160401b038116900361083e5760a435602082015260c4356040820152602084015260e435604084015261010435606084015261012435608084015263ffffffff6101443516610144350361083e576101443560a08401526101643561ffff8116810361083e57602861ffff821610158061082f575b156107a1575060015b806004146106c6578060031461069e5780600214610684576001146104a057634e487b7160e01b600052605160045260246000fd5b8060246104b4926004013591018484610ee6565b6040519061012082528051600481101561066e576101208301526104ea60208201516101c06101408501526102e08401906108ea565b90602060408201518051610160860152015161018084015261052060608201519261011f199384868303016101a08701526108ea565b608082015192848203016101c085015282519260408252835180604084015260206060840195019060005b818110610627575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c08601528051602086015260406020820151600180831b0381511682880152600180831b03602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b516106458482516108d2565b85810151604085015282604082015116828501520151166080820152019701910191909161054b565b634e487b7160e01b600052602160045260246000fd5b50806024610699926004013591018484610b8a565b6104b4565b505061069963ffffffff60208401515160101c166106c06020840151916113c7565b906113fb565b505063ffffffff60208301515160101c166106fb63ffffffff6106f46106ef6020860151611277565b611340565b1682610b7d565b60208481015101519091906001600160401b0316821161076257818060101b0462010000148215171561074c576106999160018060401b039060101b166020850151526106c06020840151916113c7565b634e487b7160e01b600052601160045260246000fd5b505061069960208201516000602060405161077c8161096c565b82815201526040519061078e8261096c565b6000825263ffffffff60208301526113fb565b603661ffff8216101580610820575b156107bd5750600261046b565b61ffff8116603f036107d15750600361046b565b61ffff166040036107e357600461046b565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b6044820152606490fd5b50603e61ffff821611156107b0565b50603561ffff82161115610462565b600080fd5b60a08436031261083e576040516001600160401b0360808201908111908211176108bc5760a06020602494836080849501604052610881368a610a3a565b815260408901358382015261089860608a01610b6c565b60408201526108a960808a01610b6c565b60608201528152019501949250506102f3565b634e487b7160e01b600052604160045260246000fd5b8051600781101561066e578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b83821061092e57505050505060208091015191015290565b90919293968382826109436001948c516108d2565b0198019493920190610916565b61018081019081106001600160401b038211176108bc57604052565b604081019081106001600160401b038211176108bc57604052565b60c081019081106001600160401b038211176108bc57604052565b606081019081106001600160401b038211176108bc57604052565b602081019081106001600160401b038211176108bc57604052565b601f909101601f19168101906001600160401b038211908210176108bc57604052565b60405190610a088261096c565b6000602083604051610a19816109bd565b6060815281520152565b6001600160401b0381116108bc5760051b60200190565b919082604091031261083e57604051610a528161096c565b80928035600781101561083e578252602090810135910152565b60409291818103841361083e5760405191610a868361096c565b829481359260018060401b039384811161083e57830191602094858484031261083e5760405193610ab6856109bd565b803591821161083e570182601f8201121561083e578035610ad681610a23565b93610ae460405195866109d8565b818552878086019260061b8401019281841161083e579088809897969594939201915b838310610b1e575050505050815284520135910152565b978495969798610b32838596979495610a3a565b81520192019088979695949392610b07565b919082604091031261083e57604051610b5c8161096c565b6020808294803584520135910152565b359063ffffffff8216820361083e57565b9190820180921161074c57565b9290916000610164359061ffff8216809203610ee3575060368103610e0357506000916004905b60209384870190610bc28251611277565b908151600781101561066e57600782101561066e5703610dcd578501516001600160401b039788851694898316936008939091848810610d8a575b505063ffffffff610c146106ef610c1e9351611277565b1661018435610b7d565b9787610c2a878b610b7d565b9101998a51511610610d7c5750909291939560009160001996604093604051610c52816109bd565b6060815297819a82955b888710610c80575050505050505050505091610c7a9160409361147b565b91510152565b9091929394959697989b8d8c610c968a85610b7d565b918260051c918203610d35575b5050601f1660ff9087811015610cf957610cbc90611744565b8060031b908082048a149015171561074c5781811b19909216908e1690911b179b861c66ffffffffffffff16989796600101959493929190610c5c565b8a5162461bcd60e51b81526004810189905260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b6044820152606490fd5b81601f98939f9188939f8e90610d5b978a956000198103610d65575b50505050516115b5565b9b9c90958f610ca3565b610d6e9261147b565b90825101528d388080610d51565b600290525050505050505050565b9194509060031b6008600160401b038116906008600160431b0316810361074c576001901b8a16600019018a811161074c571689169263ffffffff610c14610bfd565b60405162461bcd60e51b815260048101879052600e60248201526d4241445f53544f52455f5459504560901b6044820152606490fd5b60378103610e175750600191600890610bb1565b60388103610e2b5750600291600490610bb1565b60398103610e3f5750600391600890610bb1565b603a8103610e535750600091600190610bb1565b603b8103610e675750600091600290610bb1565b603c8103610e7a57506001918290610bb1565b603d8103610e8e5750600191600290610bb1565b603e03610ea057600191600490610bb1565b60405162461bcd60e51b815260206004820152601b60248201527a494e56414c49445f4d454d4f52595f53544f52455f4f50434f444560281b6044820152606490fd5b80fd5b9092916000936101643561ffff95868216809203610ee35750602881036110ed5750600094600490610f3b87965b60208701958463ffffffff966020610f3389610c146106ef8d51611277565b910151611752565b5095906110e057506001600160401b0385811696909590610f86575b50505050519060405192610f6a8461096c565b600785101561066e57610f849484521660208301526113fb565b565b908593949692916001928381148080916110cb575b15610fb957505050505060ff91501660000b16915b38808080610f57565b806110b5575b15610fd7575050505060ff1660000b16929050610fb0565b92949193926002811480806110a0575b15610ffa575050505016900b1691610fb0565b929691928061108a575b156110165750505016900b1691610fb0565b91955093925060049150149081611077575b501561103a571660030b811691610fb0565b60405162461bcd60e51b815260206004820152601560248201527410905117d491505117d096551154d7d4d251d39151605a1b6044820152606490fd5b9050600786101561066e57851438611028565b509550600789101561066e578695858a14611004565b50925060078a101561066e5787928a15610fe7565b509450600789101561066e578694838a14610fbf565b50955060078a101561066e5787958a15610f9b565b6002905250505050505050565b602981036111075750600194600890610f3b600096610f14565b602a81036111215750600294600490610f3b600096610f14565b602b810361113b5750600394600890610f3b600096610f14565b602c81036111545750600094600190610f3b8296610f14565b602d810361116d5750600094600190610f3b8796610f14565b602e81036111875750600094600290610f3b600196610f14565b602f81036111a05750600094600290610f3b8796610f14565b603081036111b857506001948590610f3b8796610f14565b603181036111d157506001948590610f3b600096610f14565b603281036111ea5750600194600290610f3b8796610f14565b603381036112045750600194600290610f3b600096610f14565b6034810361121d5750600194600490610f3b8796610f14565b60350361123557600194600490610f3b600096610f14565b60405162461bcd60e51b815260206004820152601a602482015279494e56414c49445f4d454d4f52595f4c4f41445f4f50434f444560301b6044820152606490fd5b906020916040516112878161096c565b6000938185809352015251918060206040516112a28161096c565b828152015282518051600019939184820191821161132c57906112c491611878565b5192845151908101908111611318576112dc906118a2565b915b825181101561131257806112f56001928751611878565b516113008286611878565b5261130b8185611878565b50016112de565b50925290565b634e487b7160e01b83526011600452602483fd5b634e487b7160e01b84526011600452602484fd5b60208101519051600781101561066e5761139857600160201b8110156113695763ffffffff1690565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606490fd5b60405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606490fd5b600060206040516113d78161096c565b828152015263ffffffff604051916113ee8361096c565b6000835216602082015290565b51805151916001926001810180911161074c57611417906118a2565b926000815b61143f575b505081515161143b916114348286611878565b5283611878565b5052565b83518051821015611475578161145784938493611878565b516114628289611878565b5261146d8188611878565b50019061141c565b50611421565b926040916040519360209460208101916b26b2b6b7b93c903632b0b31d60a11b8352602c820152602c81526114af816109a2565b51902092604051916114c08361096c565b601383527226b2b6b7b93c9036b2b935b632903a3932b29d60691b6020840152936000945b875191825187101561156e579060019182938885841615600014611543576115279150611516611535918d51611878565b5187519283918d8301958b87611905565b03601f1981018352826109d8565b519020925b1c9501946114e5565b6115536115659161152793611878565b519287519283918d8301958b87611905565b5190209261153a565b96509450505050925061157e5790565b60405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606490fd5b919294909394604091606083516115cb816109bd565b526000969560005b602081106117135750916116149695939188959360ff9991606085516115f8816109bd565b5261160489858961193a565b359460f89560f81c9a8b9a611946565b9661161e8b610a23565b9961162b83519b8c6109d8565b8b8b526116378c610a23565b60209990601f1901368d8c013760009e8f5b169d8e10156116aa5760009060005b898d8d8310611682575050509d61166f908d611878565b5260ff6001819f01169d9c8c9d8f611649565b93611695846116a393600195969761193a565b358d1c9060081b1793611946565b9101611658565b949c50949c50965096509697508593506116d492508351976116cb896109bd565b8852978761147b565b910151036116df5750565b5162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d3515357d493d3d560921b6044820152606490fd5b9593969297611738600191611729868c8961193a565b3560f81c9060081b1794611946565b939893979496016115d3565b601f0390601f821161074c57565b94929190946000946117648388610b7d565b82516001600160401b03161061186b5791906000969396198693879888965b8588106117995750505050505050506000929190565b90919293949596996117ab8b83610b7d565b8060051c86810361184d575b50601f1660208082101561181057506117cf90611744565b9060039180831b9260089180850483149015171561074c578d901b908d8204148d15171561074c5760ff8a6001941c16901b179a0196959493929190611783565b6064906040519062461bcd60e51b8252600482015260166024820152750848288bea0aa9898be988a828cbe84b2a88abe9288b60531b6044820152fd5b98509450611860601f9a85858b8a6115b5565b509a909895906117b7565b5050505091506001918190565b805182101561188c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b906118ac82610a23565b6040906118bc60405191826109d8565b83815280936118cd601f1991610a23565b019160009060005b8481106118e3575050505050565b60209082516118f18161096c565b8481528285818301528287010152016118d5565b9392909384519060005b82811061192757500191825260208201526040019150565b806020809289010151818401520161190f565b9082101561188c570190565b600019811461074c576001019056fea26469706673582212200172c2ae9b6d26e35fb9f346ec947cd5a4c7e1157d53513614fe4f7e74ea415964736f6c63430008190033",
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
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

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemorySession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}
