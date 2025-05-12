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
	Bin: "0x60808060405234601557611a61908161001b8239f35b600080fdfe608060408181526004918236101561001657600080fd5b600090813560e01c908163740085d7146103d05750806379754cba14610380578063ae364ac214610354578063b7465799146102d55763d4e5dd2b1461005b57600080fd5b346102d257816003193601126102d25767ffffffffffffffff9183358381116102c35761008b90369086016104cb565b9261009461046f565b906100a03686856119d3565b9586516020809801209581606094169485821161027c575b50508451916100c683610533565b60019060018452600189850194868652898352828b528883208884528b528883209051151560ff60ff1983541691161781550193519081519384116102695790849392916101168a9796546104f9565b8b601f8211610230575b50508a92601f85116001146101a05750928061018c9593610178937ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c989692610195575b50506000198260011b9260031b1c19161790565b90555b85519182918983528983019061048b565b0390a351908152f35b015190508c80610164565b8582528b8220939291601f198616915b8d838310610215575050509260019285927ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c989661018c9896106101fc575b505050811b01905561017b565b015160001960f88460031b161c191690558b80806101ef565b8486015187558d9a50958101959485019491909101906101b0565b86835280832061025892601f880160051c820192881061025f575b601f0160051c0190610667565b8b8b610120565b909150819061024b565b8060418c634e487b7160e01b6024945252fd5b9091935061028a8583610711565b8881116102cb575b61029c90866106ee565b918286116102c75782116102c357906102bb91858036930391016119d3565b9138806100b8565b8280fd5b8380fd5b5087610292565b80fd5b5091346102c35760206003193601126102c3573573ffffffffffffffffffffffffffffffffffffffff81168091036102c35791819281526001602052209067ffffffffffffffff82541690600961032e600185016105c1565b93015461034c8251948594855260606020860152606085019061048b565b918301520390f35b50903461037c578160031936011261037c57610379903383526001602052822061067e565b80f35b5080fd5b50346102d25760606003193601126102d25782359067ffffffffffffffff82116102d257506103b76020936103c9923691016104cb565b6103bf61046f565b906044359261075c565b9051908152f35b919050346102d257826003193601126102d2578335836103ee61046f565b928281528060205267ffffffffffffffff828220941693848252602052209460ff8654161561043e5761043a85610427600189016105c1565b905191829160208352602083019061048b565b0390f35b604494507f139647920000000000000000000000000000000000000000000000000000000084528301526024820152fd5b6024359067ffffffffffffffff8216820361048657565b600080fd5b919082519283825260005b8481106104b7575050601f19601f8460006020809697860101520116010190565b602081830181015184830182015201610496565b9181601f840112156104865782359167ffffffffffffffff8311610486576020838186019501011161048657565b90600182811c92168015610529575b602083101461051357565b634e487b7160e01b600052602260045260246000fd5b91607f1691610508565b6040810190811067ffffffffffffffff82111761054f57604052565b634e487b7160e01b600052604160045260246000fd5b610320810190811067ffffffffffffffff82111761054f57604052565b60a0810190811067ffffffffffffffff82111761054f57604052565b90601f601f19910116810190811067ffffffffffffffff82111761054f57604052565b906040519182600082546105d4816104f9565b908184526020946001916001811690816000146106445750600114610605575b5050506106039250038361059e565b565b600090815285812095935091905b81831061062c57505061060393508201013880806105f4565b85548884018501529485019487945091830191610613565b91505061060395935060ff1991501682840152151560051b8201013880806105f4565b818110610672575050565b60008155600101610667565b60026106ae600092838155836001820161069881546104f9565b806106b1575b5050506009810192839101610667565b55565b82601f82116001146106c9575050555b83388061069e565b90918082526106e7601f60208420940160051c840160018501610667565b55556106c1565b919082018092116106fb57565b634e487b7160e01b600052601160045260246000fd5b919082039182116106fb57565b9082101561072a570190565b634e487b7160e01b600052603260045260246000fd5b919091601983101561072a576018908360021c019260031b1690565b91909293600091600286166119b8575b60018616158015906119ad575b1561194f57336000526001602052604060002095600987015480156000146118d45767ffffffffffffffff83167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008954161788559594955b8596806107e28860098c01546106ee565b60098b01555b8815806118c9575b6118bb5760005b608881106117df575060405161080c81610565565b61032036823760005b8b601982106117a557505061032060405161082f81610565565b36903760405161083e81610582565b60a036823760405161084f81610582565b60a03682376040519061086182610565565b6103203683376040518061030081011067ffffffffffffffff6103008301111761054f57610300810160405260018152618082602082015267800000000000808a6040820152678000000080008000606082015261808b6080820152638000000160a082015267800000008000808160c082015267800000000000800960e0820152608a61010082015260886101208201526380008009610140820152638000000a610160820152638000808b61018082015267800000000000008b6101a08201526780000000000080896101c08201526780000000000080036101e082015267800000000000800261020082015267800000000000008061022082015261800a61024082015267800000008000000a6102608201526780000000800080816102808201526780000000000080806102a082015263800000016102c08201526780000000800080086102e082015260005b60188110610f0457505050505060005b60198110610ead57505060888910610a0957886088116104865760887fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7891019801976107e8565b5091939590929496505b67ffffffffffffffff80951692602084018085116106fb5781811180610ea0575b610d46575b505050505060011615610d3d5760005b60208110610cbe57506040516001850191610a6382610533565b60018252610a70836105c1565b956020830196875285600052600060205260406000208282541660005260205260016040600020935115159360ff199460ff86835416911617815501965196875183811161054f57610ac282546104f9565b601f8111610c8c575b506020601f8211600114610c20579080610b00928a9b60009b999a9b92610c155750506000198260011b9260031b1c19161790565b90555b54169260405191602083526000918054610b1c816104f9565b92836020870152600182169182600014610bd5575050600114610b79575b505090807ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c920390a3336000526001602052610603604060002061067e565b6000908152602081209092505b818310610bbb5750508101604001817ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c610b3a565b805460408585010152879450602090920191600101610b86565b869550604093507ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c969492501682840152151560051b8201019192610b3a565b015190503880610164565b601f198216998360005260206000209a60005b818110610c745750918a9b9184600195949c9a9b9c10610c5b575b505050811b019055610b03565b015160001960f88460031b161c19169055388080610c4e565b838301518d556001909c019b60209384019301610c33565b610cb890836000526020600020601f840160051c8101916020851061025f57601f0160051c0190610667565b38610acb565b9260039084821c60058082069182820292828404036106fb578592610cf192610ce89204906106ee565b60028901610740565b905490841b1c169185901b603881166008600788168015908304821417156106fb5787830414871517156106fb5760f89182039182116106fb5760019360ff911c16901b179301610a49565b50915050600090565b600094828111610e87575b5090610d5c91610711565b92818411610e7e575b929160018901935b83811015610a3957610d8081848461071e565b3590855491610d8e836104f9565b906801000000000000000082101561054f57601f9384831115610e0b576002610dba9101808a556104f9565b938483101561072a57602060019510600014610df3578892811690035b60ff83549160031b9260f81c831b921b19161790555b01610d6d565b886000528060206000208460051c0193169003610dd7565b90917fff00000000000000000000000000000000000000000000000000000000000000600195931660001a9281601f14610e5d57906002910360031b9260ff908116841b931b19910116178655610ded565b5050908760005260ff19602060002092169060ff1617905560418655610ded565b92508092610d65565b82610d5c93929650610e9891610711565b949091610d51565b5060098a01548510610a34565b80610efa8d610eda959e9c9583600267ffffffffffffffff610ed16001988a611a1a565b51169301610740565b90919067ffffffffffffffff8084549260031b9316831b921b1916179055565b019a91989a6109c2565b959e9c956001908651602088015118604088015118606088015118608088015118865260a087015160c08801511860e088015118610100880151186101208801511880602088015261014088015161016089015118610180890151186101a0890151186101c08901511860408801526101e08801516102008901511861022089015118610240890151186102608901511860608801526102808801516102a0890151186102c0890151186102e089015118610300890151189081608089015267ffffffffffffffff81603f1c91851b1617188085528651604088015167ffffffffffffffff81603f1c91861b16171860208601526020870151606088015167ffffffffffffffff81603f1c91861b16171860408601526040870151608088015167ffffffffffffffff81603f1c91861b16171860608601526060870151875167ffffffffffffffff81603f1c91861b16171860808601528751188088526020880151855118602089015260408801518551186040890152606088015185511860608901526080880151855118608089015260a088015160208601511860a089015260c088015160208601511860c089015260e088015160208601511860e08901526101008801516020860151186101008901526101208801516020860151186101208901526101408801516040860151186101408901526101608801516040860151186101608901526101808801516040860151186101808901526101a08801516040860151186101a08901526101c08801516040860151186101c08901526101e08801516060860151186101e08901526102008801516060860151186102008901526102208801516060860151186102208901526102408801516060860151186102408901526102608801516060860151186102608901526102808801516080860151186102808901526102a08801516080860151186102a08901526102c08801516080860151186102c08901526102e08801516080860151186102e0890152610300880151608086015118610300890152808652602088015167ffffffffffffffff81601c1c9160241b1617610100870152604088015167ffffffffffffffff81603d1c9160031b1617610160870152606088015167ffffffffffffffff8160171c9160291b1617610260870152608088015167ffffffffffffffff81602e1c9160121b16176102c087015260a088015167ffffffffffffffff81603f1c91851b1617604087015260c0880151908160141c67ffffffffffffffff83602c1b161760a088015260e089015167ffffffffffffffff8160361c91600a1b16176101a088015261010089015167ffffffffffffffff8160131c91602d1b1617610200880152610120890151603e9067ffffffffffffffff81831c9160021b161761030089015267ffffffffffffffff6101408b01518060021c921b1617608088015261016089015167ffffffffffffffff81603a1c9160061b161760e0880152610180890151918260151c67ffffffffffffffff84602b1b16176101408901526101a08a015167ffffffffffffffff8160311c91600f1b16176102408901526101c08a015167ffffffffffffffff8160031c91603d1b16176102a08901526101e08a015167ffffffffffffffff8160241c91601c1b161760208901526102008a015167ffffffffffffffff8160091c9160371b16176101208901526102208a015167ffffffffffffffff8160271c9160191b16176101808901526102408a015167ffffffffffffffff81602b1c9160151b16176101e08901526102608a015167ffffffffffffffff8160081c9160381b16176102e08901526102808a015167ffffffffffffffff8160251c91601b1b161760608901526102a08a015167ffffffffffffffff81602c1c9160141b161760c08901526102c08a015167ffffffffffffffff8160191c9160271b16176101c08901526102e08a015167ffffffffffffffff8160381c9160081b16176102208901526103008a015167ffffffffffffffff8160321c91600e1b16176102808901528260151c67ffffffffffffffff84602b1b16178160141c67ffffffffffffffff83602c1b1617191682188a52602088015160c0890151196101608a0151161860208b0152604088015160e0890151196101808a0151161860408b01526060880151610100890151196101a08a0151161860608b01526080880151610120890151196101c08a0151161860808b015260a0880151610140890151196101e08a0151161860a08b015260c0880151610160890151196102008a0151161860c08b015260e0880151610180890151196102208a0151161860e08b01526101008801516101a0890151196102408a015116186101008b01526101208801516101c0890151196102608a015116186101208b01526101408801516101e0890151196102808a015116186101408b0152610160880151610200890151196102a08a015116186101608b0152610180880151610220890151196102c08a015116186101808b01526101a0880151610240890151196102e08a015116186101a08b01526101c0880151610260890151196103008a015116186101c08b01526101e088015161028089015119895116186101e08b01526102008801516102a08901511960208a015116186102008b01526102208801516102c08901511960408a015116186102208b01526102408801516102e08901511960608a015116186102408b01526102608801516103008901511960808a015116186102608b015261028088015188511960a08a015116186102808b01526102a088015160208901511960c08a015116186102a08b01526102c088015160408901511960e08a015116186102c08b01526102e08801516060890151196101008a015116186102e08b01526103008801516080890151196101208a015116186103008b01528360051b8601519267ffffffffffffffff8160151c91602b1b16179067ffffffffffffffff8160141c91602c1b1617191618188752019e959c9e6109b2565b9067ffffffffffffffff6117c2826002969e9d9660019501610740565b90549060031b1c166117d48285611a1a565b520199919899610815565b9098979060008982101561189757506117f9818a8c61071e565b3560f81c905b60039181831c60058082069081810291818304036106fb576118229204906106ee565b9060078316603884861b1690808204600814901517156106fb576001948f60029060ff9361185f61188d9767ffffffffffffffff94859401610740565b95909616901b1691838554911b1c16189067ffffffffffffffff8084549260031b9316831b921b1916179055565b01989097986107f7565b908981146118b2575b608781036117ff5790608017906117ff565b600191506118a0565b509193959092949650610a13565b5060018416156107f0565b67ffffffffffffffff80895416908416036118f1579594956107d1565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f444946465f4f46465345540000000000000000000000000000000000000000006044820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f4e4f545f424c4f434b5f414c49474e45440000000000000000000000000000006044820152fd5b506088850615610779565b3360005260016020526119ce604060002061067e565b61076c565b92919267ffffffffffffffff821161054f57604051916119fd6020601f19601f840116018461059e565b829481845281830111610486578281602093846000960137010152565b90601981101561072a5760051b019056fea264697066735822122028faaebcb58c6a2b8e5fa2b3278b690c68e63de8ea3a37bb6bd2365c09799a1e64736f6c63430008190033",
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
	Bin: "0x60803460ae57601f612daa38819003918201601f19168301916001600160401b0383118484101760b35780849260809460405283398101031260ae5760428160c9565b90604d6020820160c9565b60616060605b6040850160c9565b930160c9565b9060018060a01b03928380928160018060a01b03199716876000541617600055168560015416176001551683600254161760025516906003541617600355604051612ccd90816100dd8239f35b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820360ae5756fe6080604052600436101561001257600080fd5b60003560e01c806304997be4146100875780631f128bc01461008257806330a5509f1461007d5780635f52fd7c1461007857806366e5d9c314610073578063b5112fd21461006e5763c39619c41461006957600080fd5b6101ca565b610151565b61012a565b610103565b6100dc565b6100b5565b346100b05760406003193601126100b05760206100a8602435600435610437565b604051908152f35b600080fd5b346100b05760006003193601126100b05760206001600160a01b0360015416604051908152f35b346100b05760006003193601126100b05760206001600160a01b0360005416604051908152f35b346100b05760006003193601126100b05760206001600160a01b0360035416604051908152f35b346100b05760006003193601126100b05760206001600160a01b0360025416604051908152f35b346100b057600319360160c081126100b0576060136100b05767ffffffffffffffff60a4358181116100b057366023820112156100b05780600401359182116100b05736602483830101116100b0576101c69160246101b69201608435606435610f3e565b6040519081529081906020820190565b0390f35b346100b05760a06003193601126100b05760206100a86115b5565b634e487b7160e01b600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761021757604052565b6101e5565b6020810190811067ffffffffffffffff82111761021757604052565b6060810190811067ffffffffffffffff82111761021757604052565b60c0810190811067ffffffffffffffff82111761021757604052565b6080810190811067ffffffffffffffff82111761021757604052565b90601f601f19910116810190811067ffffffffffffffff82111761021757604052565b604051906102bc8261021c565b565b604051906102bc826101fb565b60405190610180820182811067ffffffffffffffff82111761021757604052565b67ffffffffffffffff81116102175760051b60200190565b60405190610311826101fb565b60006020838281520152565b6040519061032a82610270565b600382528160005b6060811061033e575050565b602090610349610304565b82828501015201610332565b634e487b7160e01b600052603260045260246000fd5b8051156103785760200190565b610355565b8051600110156103785760400190565b8051600210156103785760600190565b80518210156103785760209160051b010190565b604051906103be8261021c565b60608252565b604051906103d1826101fb565b60006020836040516103e28161021c565b6060815281520152565b604051906103f9826101fb565b6000602083606081520152565b634e487b7160e01b600052602160045260246000fd5b6003111561042657565b610406565b60038210156104265752565b6105319161044361031d565b9161044c6116d9565b6104558461036b565b5261045f8361036b565b506104686116fc565b6104718461037d565b5261047b8361037d565b506104846116fc565b61048d8461038d565b526104978361038d565b506104a06102af565b9283526104ab6102be565b928352600060208401526104bd6103c4565b6104c56103ec565b6104cd610304565b600019815260006020820152916104e26102cb565b956000875260208701528260408701526060860152608085015260a084015260c0830152600060e08301526000610100830152600061012083015260001961014083015261016082015261171f565b90565b60405190610180820182811067ffffffffffffffff821117610217576040526000808352610160836105646103c4565b6020820152610571610304565b604082015261057e6103c4565b606082015261058b6103ec565b6080820152610598610304565b60a08201528260c08201528260e08201528261010082015282610120820152826101408201520152565b604051906105cf82610238565b60006040838281528260208201520152565b604051906105ee82610254565b600060a0838281526105fe6105c2565b60208201528260408201528260608201528260808201520152565b1561062057565b606460405162461bcd60e51b815260206004820152601360248201527f4d414348494e455f4245464f52455f48415348000000000000000000000000006044820152fd5b60405190610671826101fb565b6040368337565b60405190610685826101fb565b81604051610692816101fb565b604036823781526020604051916106a8836101fb565b60403684370152565b156106b857565b606460405162461bcd60e51b815260206004820152601060248201527f4241445f474c4f42414c5f5354415445000000000000000000000000000000006044820152fd5b634e487b7160e01b600052601160045260246000fd5b906001820180921161072057565b6106fc565b600e019081600e1161072057565b906002820180921161072057565b906020820180921161072057565b1561075657565b606460405162461bcd60e51b815260206004820152600c60248201527f4d4f44554c45535f524f4f5400000000000000000000000000000000000000006044820152fd5b156107a157565b606460405162461bcd60e51b815260206004820152601260248201527f4241445f46554e4354494f4e535f524f4f5400000000000000000000000000006044820152fd5b909392938483116100b05784116100b0578101920390565b90600163ffffffff8093160191821161072057565b600311156100b057565b51906102bc82610812565b91908260409103126100b05760405161083f816101fb565b8092805160078110156100b0578252602090810151910152565b6040929181810384136100b05760405191610873836101fb565b829481519267ffffffffffffffff938481116100b05783019160209485848403126100b057604051936108a58561021c565b80519182116100b0570182601f820112156100b05780516108c5816102ec565b936108d3604051958661028c565b818552878086019260061b840101928184116100b0579088809897969594939201915b83831061090d575050505050815284520151910152565b978495969798610921838596979495610827565b815201920190889796959493926108f6565b91908260409103126100b05760405161094b816101fb565b6020808294805184520151910152565b519063ffffffff821682036100b057565b919091604080828503126100b057805191610986836101fb565b8294815167ffffffffffffffff81116100b057820181601f820112156100b0578051906020946109b5836102ec565b936109c28251958661028c565b838552868501918760a0809602850101938285116100b0579088809897969594939201925b8484106109fd5750505050505084520151910152565b90919293948096979850848403126100b0578886918351610a1d81610270565b610a278688610827565b815284870151838201526060610a3e81890161095b565b86830152610a4e6080890161095b565b90820152815201930191908897969594936109e7565b67ffffffffffffffff8116036100b057565b809291039161010083126100b05760405190610a9182610254565b6060601f1983958351855201126100b057610b0b60e060a092604051610ab681610238565b6020820151610ac481610a64565b81526040820151610ad481610a64565b60208201526060820151604082015260208601526080810151604086015283810151606086015260c081015160808601520161095b565b910152565b91906101209081848203126100b05783519167ffffffffffffffff928381116100b0578501906101c0828403126100b057610b496102cb565b90610b538361081c565b825260208301518581116100b05784610b6d918501610859565b6020830152610b7f8460408501610933565b604083015260808301518581116100b05784610b9c918501610859565b606083015260a08301519485116100b0576101a083610bc286610531986020970161096c565b6080850152610bd48660c08301610933565b60a08501526101008082015160c0860152610bf084830161095b565b60e086015261014090610c0482840161095b565b9086015261016093610c1785840161095b565b90860152610180820151908501520151908201529401610a76565b9060038210156104265752565b80516007811015610426578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b838210610c9b57505050505060208091015191015290565b9091929396838282610cb06001948c51610c3f565b0198019493920190610c83565b90604091604082019281519360408452845180915260609460608501956020809201936000915b848310610cfd5750505050505060208091015191015290565b9091929394978460a06001928b51610d16828251610c3f565b80840151828801528681015163ffffffff9081168784015290860151166080820152019901959493019190610ce4565b601f8260209493601f19938186528686013760008582860101520116010190565b9093919492946101e06004358352602435966001600160a01b0388168098036100b057610db1610f2792610531996020870152604435604087015280606087015285018851610c32565b610f0f602088015193610160610e41610dd96101c097886102008b01526103a08a0190610c57565b60408c015180516102208b0152602001516102408a0152610e2b60608d0151917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe2092838c8303016102608d0152610c57565b9060808d0151908a8303016102808b0152610cbd565b60a08b015180516102a08a0152602001516102c08901529960c08101516102e089015260e081015163ffffffff1661030089015261010081015163ffffffff1661032089015261012081015163ffffffff166103408901526101408101516103608901520151610380870152608086019063ffffffff60a060e092805185526040602082015167ffffffffffffffff808251166020890152602082015116828801520151606086015260408101516080860152606081015182860152608081015160c0860152015116910152565b805161ffff16610180850152602001516101a0840152565b818503910152610d46565b6040513d6000823e3d90fd5b610f8590939193610f4d610534565b50610f566105e1565b50610f5f6103b1565b50610f68610304565b50610f738385611976565b929095610f7f8761171f565b14610619565b8451610f908161041c565b610f998161041c565b61146957610fad6508000000000091610712565b1461145957610fbd908284611cad565b610fcb908385969396611d87565b9290918160e0820195878751610fe49063ffffffff1690565b63ffffffff1690610ff59187611e83565b956101609687850151146110089061074f565b6110106103b1565b506110196103b1565b50611025908383611fcc565b611033908484979397611d87565b611041908385989398611d87565b91906101208701978289516110599063ffffffff1690565b60061c6303ffffff1663ffffffff16611071926120c8565b61010088015163ffffffff1663ffffffff1661108c9261220e565b60608c01511461109b9061079a565b86516110ac91603f9091169061039d565b51936110b7936107e5565b9096516110c79063ffffffff1690565b63ffffffff169380516110dd9063ffffffff1690565b6110e6906107fd565b63ffffffff169052815161ffff16976000936001600160a01b03908560288c10158061144e575b8015611437575b801561142d575b8015611423575b8714611248575061117661113e6001546001600160a01b031690565b945b6040519b8c97889687957fa92cb50100000000000000000000000000000000000000000000000000000000875260048701610d67565b0392165afa9081156112435761053195600095869361121a575b50618023811490811561120e575b50156111f8575b50505050600281516111b68161041c565b6111bf8161041c565b14806111e7575b1561171f576111d481612290565b6111dd81612312565b506000815261171f565b5060001961014082015114156111c6565b61120192611e83565b90820152388080806111a5565b6180249150143861119e565b90925061123991953d8091833e611231818361028c565b810190610b10565b9490949138611190565b610f32565b60458c148015611419575b8015611402575b80156113eb575b80156113d4575b80156113bd575b80156113a6575b801561138f575b8015611385575b8015611371575b801561135a575b8015611343575b87146112bb57506111766112b56002546001600160a01b031690565b94611140565b6180108c101580611337575b801561131e575b8015611305575b87146112f157506111766112b56003546001600160a01b031690565b54611176906001600160a01b031694611140565b506180308c101580156112d557506180328c11156112d5565b506180208c101580156112ce57506180248c11156112ce565b506180138c11156112c7565b5060bc8c10158015611299575060bf8c1115611299565b5060c08c10158015611292575060c48c1115611292565b5060ac8c148061128b575060ad8c1461128b565b5060a78c14611284565b50607c8c1015801561127d5750608a8c111561127d565b5060798c101580156112765750607b8c1115611276565b5060518c1015801561126f5750605a8c111561126f565b50606a8c10158015611268575060788c1115611268565b5060678c10158015611261575060698c1115611261565b5060468c1015801561125a5750604f8c111561125a565b5060508c14611253565b5060408c14611122565b50603f8c1461111b565b5060368c101580156111145750603e8c1115611114565b5060358c111561110d565b505060028252506105319061171f565b929161147d92611477610678565b50611b5f565b509061148882611c20565b9161149960c08501938451146106b1565b600184516114a68161041c565b6114af8161041c565b1491826114f7575b50816114dc575b506114cd57506105319061171f565b61053191505160443590610437565b67ffffffffffffffff915060200151511660043511386114be565b159150386114b7565b60843561053181610812565b9060806003198301126100b05760405191611526836101fb565b8281602312156100b05760405161153c816101fb565b6044818482116100b0576004905b8282106115a557505050815281606312156100b0576040519161156c836101fb565b826084916084116100b0576044905b82821061158b5750505060200152565b60208091833561159a81610a64565b81520191019061157b565b813581526020918201910161154a565b60016115bf611500565b6115c88161041c565b03611631576115de6115d93661150c565b611c20565b60405161162b8161161d6020820194856031917f4d616368696e652066696e69736865643a000000000000000000000000000000825260118201520190565b03601f19810183528261028c565b51902090565b600261163b611500565b6116448161041c565b03611694576116556115d93661150c565b60405161162b8161161d6020820194856030917f4d616368696e65206572726f7265643a00000000000000000000000000000000825260108201520190565b60405162461bcd60e51b815260206004820152601260248201527f4241445f4d414348494e455f53544154555300000000000000000000000000006044820152606490fd5b6116e1610304565b506040516116ee816101fb565b600481526000602082015290565b611704610304565b50604051611711816101fb565b600081526000602082015290565b805161172a8161041c565b6117338161041c565b61187557604081015161162b61174c6020840151612338565b9261161d61178a61176b61014084019560001997888851141591612463565b9560a08401519061177f6080860151612530565b908751141591612463565b936117986060840151612338565b9260c0810151916117b060e083015163ffffffff1690565b61010083015163ffffffff16906101606117d261012086015163ffffffff1690565b935194015194604051998a9860208a019c8d97959492909160dc999794927f4d616368696e652072756e6e696e673a000000000000000000000000000000008a5260108a01526030890152605088015260708701527fffffffff000000000000000000000000000000000000000000000000000000009283809260e01b16609088015260e01b16609486015260e01b166098840152609c83015260bc8201520190565b600181516118828161041c565b61188b8161041c565b036118d35760c0015160405161162b8161161d6020820194856031917f4d616368696e652066696e69736865643a000000000000000000000000000000825260118201520190565b600281516118e08161041c565b6118e98161041c565b036119315760c0015160405161162b8161161d6020820194856030917f4d616368696e65206572726f7265643a00000000000000000000000000000000825260108201520190565b60405162461bcd60e51b815260206004820152600f60248201527f4241445f4d4143485f53544154555300000000000000000000000000000000006044820152606490fd5b9160ff9192611983610534565b5061198e848261265c565b931680611ad857506000925b6119a26103c4565b506119ab6103c4565b506119b4610304565b506119bd6103ec565b506119c6610304565b506119d2908583612687565b6119e0908684979397612750565b6119ee908385989398612687565b6119f99084866127fb565b611a07908587949394612750565b989093611a126102cb565b98611a1d908a61042b565b602089015260408801526060870152608086015260a085015260c08401906000825260e0850160008152610100860160008152610120870191600083526101408801946000865261016089019860008a5299611a7a908689612ad2565b9152611a879085886128ee565b63ffffffff909116909152611a9d9084876128ee565b63ffffffff909116909152611ab39083866128ee565b63ffffffff909116909152611ac9908285612ad2565b9252611ad492612ad2565b9252565b60018103611ae9575060019261199a565b600203611af85760029261199a565b60405162461bcd60e51b815260206004820152601360248201527f554e4b4e4f574e5f4d4143485f535441545553000000000000000000000000006044820152606490fd5b60ff1660ff81146107205760010190565b9060028110156103785760051b0190565b909291611b6a610678565b5060405193611b78856101fb565b6040368637611b85610664565b9260005b60ff8116936002851015611bbe5790611bb3611ba9611bb9938686612ad2565b919091968a611b4e565b52611b3d565b611b89565b93505060005b60ff8116936002851015611c055790611bf767ffffffffffffffff611bed611c00948787612936565b9290929789611b4e565b91169052611b3d565b611bc4565b9596949350505050611c156102be565b918252602082015291565b8051906020808351930151910151602081519101516040519260208401947f476c6f62616c2073746174653a000000000000000000000000000000000000008652602d850152604d8401527fffffffffffffffff000000000000000000000000000000000000000000000000809260c01b16606d84015260c01b166075820152605d815261162b81610270565b611d4563ffffffff91611d4f611d57611cf995611d3b611d03611ce6611d0d9a611cd56105e1565b50611cde6105c2565b508488612ad2565b611cf19a919a6105c2565b508488612936565b8488939293612936565b84889c929c612ad2565b916040519b611d1b8d610238565b67ffffffffffffffff8092168d521660208c015260408b01528286612ad2565b8286979297612ad2565b8286959295612ad2565b9190946128ee565b9390939660405196611d6888610254565b875260208701526040860152606085015260808401521660a082015291565b909291926060604051611d998161021c565b52611db2611da8858385612641565b3560f81c9461264d565b90611dbc856102ec565b92611dca604051948561028c565b858452601f19611dd9876102ec565b0136602086013760005b60ff81169387851015611e155760ff91611e0d611e036001938787612ad2565b919091978961039d565b520116611de3565b9596505050505060405190611e298261021c565b815291565b67ffffffffffffffff811161021757601f01601f191660200190565b60405190611e57826101fb565b601382527f4d6f64756c65206d65726b6c6520747265653a000000000000000000000000006020830152565b9061053192805190611fba60208201518051604060208301519201516040519160208301937f4d656d6f72793a0000000000000000000000000000000000000000000000000085527fffffffffffffffff000000000000000000000000000000000000000000000000809260c01b16602785015260c01b16602f830152603782015260378152611f1281610238565b51902061161d604084015193606081015190611f3a60a0608083015192015163ffffffff1690565b9160405196879560208701998a94919260ab96937fffffffff0000000000000000000000000000000000000000000000000000000095927f4d6f64756c653a0000000000000000000000000000000000000000000000000088526007880152602787015260478601526067850152608784015260e01b1660a78201520190565b51902090611fc6611e4a565b92612a03565b611fdc60ff93949294858361266c565b931690611fe8826102ec565b90611ff6604051928361028c565b828252601f19612005846102ec565b0160005b818110612078575050819560005b848110612025575050505050565b600190612036612040988486612a9c565b84869a929a612ad2565b989061205761204d6102be565b61ffff9093168352565b6020820152612066828761039d565b52612071818661039d565b5001612017565b602090612083610304565b82828701015201612009565b6040519061209c826101fb565b601882527f496e737472756374696f6e206d65726b6c6520747265653a00000000000000006020830152565b9291926120e56120e06120db8651612b14565b610725565b612b25565b9160209460208401947f496e737472756374696f6e733a00000000000000000000000000000000000000865261214761211f835160ff1690565b60f81b7fff000000000000000000000000000000000000000000000000000000000000001690565b9360009460001a61215787612b56565b53600093600e5b84518610156121f4576121ec878b6121de848c6121af6001976121818e8e61039d565b51966121a761211f6121a16121988b5161ffff1690565b60081c60ff1690565b60ff1690565b901a92612b66565b538c6121d88d6121c761211f6121a1895161ffff1690565b6121d085610712565b911a92612b66565b53610733565b9101518c828c010152610741565b95019461215e565b509450949150955061053194915051902090611fc661208f565b906105319260405160208101917f46756e6374696f6e3a0000000000000000000000000000000000000000000000835260298201526029815261225081610238565b5190209060405192612261846101fb565b601584527f46756e6374696f6e206d65726b6c6520747265653a00000000000000000000006020850152612a03565b604081018051519060a0830180515190600019808314908115612308575b506122fe576020856080606095969701926122c98451612530565b90515201936122d88551612338565b905152519060208201525251906020820152604051906122f78261021c565b6060825252565b5050505060029052565b90508414386122ae565b612323610140820191825190612b77565b15612332576000199052600190565b50600090565b9060209160208101519281515151906000925b8284106123585750505050565b9091929460019061237d612377888551612370610304565b505161039d565b51612bb1565b9060405190858201927f56616c756520737461636b3a00000000000000000000000000000000000000008452602c830152604c908183015281526123c081610270565b519020950192919061234b565b156123d457565b606460405162461bcd60e51b815260206004820152601960248201527f4d554c5449535441434b5f4e4f535441434b5f414354495645000000000000006044820152fd5b1561241f57565b606460405162461bcd60e51b815260206004820152601760248201527f4d554c5449535441434b5f4e4f535441434b5f4d41494e0000000000000000006044820152fd5b9160001990612474828414156123cd565b156124db576124869083511415612418565b61162b602083519301519161161d60405193849260208401968791606b93917f6d756c7469737461636b3a0000000000000000000000000000000000000000008452600b840152602b830152604b8201520190565b5081516020928301516040517f6d756c7469737461636b3a000000000000000000000000000000000000000000948101948552602b810193909352604b830191909152606b82015261162b81608b810161161d565b602080820151926000935b83518051861015612639579061161d61262e6125598860019561039d565b51926125658451612bb1565b878501516040956060878201519101518751928b8401947f537461636b206672616d653a00000000000000000000000000000000000000008652602c850152604c8401527fffffffff00000000000000000000000000000000000000000000000000000000918260e091821b16606c8501521b166070820152605481526125eb81610270565b5190209351928391888301958690916052927f537461636b206672616d6520737461636b3a00000000000000000000000000008352601283015260328201520190565b51902094019361253b565b509350915050565b90821015610378570190565b60001981146107205760010190565b9015610378573560f81c90600190565b826105319261267d92959495612641565b3560f81c9261264d565b9061269f6126a9936126976103c4565b508284612ad2565b8284959295612ad2565b92909294601f196126d26126bc866102ec565b956126ca604051978861028c565b8087526102ec565b0160005b8181106127395750506000955b845187101561270d576126f96001918585612c07565b9790612705828861039d565b5201956126e3565b949250945050604051916127208361021c565b82526040519161272f836101fb565b8252602082015291565b602090612744610304565b828289010152016126d6565b919061276961277192612761610304565b508285612ad2565b919093612ad2565b91906040519161272f836101fb565b6040519061278d82610270565b600060608361279a610304565b81528260208201528260408201520152565b6040516127b88161021c565b6000815290565b604051906127cc826101fb565b600182528160005b602090818110156127f6576020916127ea612780565b908285010152016127d4565b505050565b91612808906126976103ec565b909290807fff00000000000000000000000000000000000000000000000000000000000000612838828686612641565b3516156128d25750612884939161289661285461288e9361264d565b9461287a6128606127bf565b96612869612780565b50612872610304565b508285612c07565b8285999299612ad2565b82859692966128ee565b9190936128ee565b92604051966128a488610270565b8752602087015263ffffffff8092166040870152166060850152926128c88361036b565b525b611c156102be565b9150506128df915061264d565b906128e86127ac565b906128ca565b600093918491905b6004831061290357505050565b9091939461292d60019163ffffff0061291d898688612641565b3560f81c9160081b16179661264d565b940191906128f6565b9192600091825b6008908185101561297a5760019167ffffffffffffff00612971926129638a878b612641565b3560f81c921b16179661264d565b9301929461293d565b95945050509050565b9392909384519060005b8281106129a557500191825260208201526040019150565b806020809289010151818401520161298d565b156129bf57565b606460405162461bcd60e51b815260206004820152600f60248201527f50524f4f465f544f4f5f53484f525400000000000000000000000000000000006044820152fd5b92939091936000925b8451958651851015612a8b579060019182978689841615600014612a5e5761161d9150612a3d612a50918a5161039d565b5160405192839160208301958987612983565b519020965b1c930192612a0c565b612a6e612a829161161d9361039d565b519260405192839160208301958987612983565b51902096612a55565b9550919350506102bc9150156129b8565b600093918491905b60028310612ab157505050565b90919394612ac960019161ff0061291d898688612641565b94019190612aa4565b600093918491905b60208310612ae757505050565b90919394612b0b600191612afc888587612641565b3560f81c9060081b179661264d565b94019190612ada565b908160220291602283040361072057565b90612b2f82611e2e565b612b3c604051918261028c565b828152601f19612b4c8294611e2e565b0190602036910137565b8051600d101561037857602d0190565b908151811015610378570160200190565b908060601c612baa5760e09063ffffffff90818116610120850152818160201c1661010085015260401c16910152600190565b5050600090565b805190600782101561042657602001516040519060208201927f56616c75653a0000000000000000000000000000000000000000000000000000845260f81b602683015260278201526027815261162b81610238565b90612c10610304565b50612c29612c1f848385612641565b3560f81c9361264d565b9060068411612c5357612c3b92612ad2565b91906007821015610426576040519161272f836101fb565b606460405162461bcd60e51b815260206004820152600e60248201527f4241445f56414c55455f545950450000000000000000000000000000000000006044820152fdfea2646970667358221220b7d7cfe38c530b5c8c80db2737b4734f477ba3fbc45bd880a81dcabf0d31eaa264736f6c63430008190033",
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
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220a3e621c44e7c4219b8d4f03be81cc4f42a976315e3a14a65854d8902d9145a1464736f6c63430008190033",
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
	Bin: "0x608080604052346015576127f6908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c63a92cb5011461002757600080fd5b3461134b5760031936016101e0811261134b5760601361134b5767ffffffffffffffff6064351161134b576101c06003196064353603011261134b576101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c36011261134b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7c36011261134b576101c43567ffffffffffffffff811161134b573660238201121561134b5767ffffffffffffffff81600401351161134b5736602482600401358301011161134b576101036080611431565b60006080526101106114fd565b60a05260405161011f8161144e565b6000808252602082015260c0526101346114fd565b60e0526040516101438161144e565b60608152600060208201526101005260405161015e8161144e565b600080825260208201819052610120919091526101408190526101608190526101808190526101a08190526101c08190526101e05261019b611544565b506040516101a881611431565b600360643560040135101561134b5760643560040135815267ffffffffffffffff602460643501351161134b576101ea366064356024810135016004016115c6565b60208201526101fe366044606435016116a0565b604082015267ffffffffffffffff608460643501351161134b5761022d366064356084810135016004016115c6565b606082015267ffffffffffffffff60a460643501351161134b57604060643560a4810135013603600319011261134b576040516102698161144e565b60643560a4810135016004013567ffffffffffffffff811161134b573660643560a4810135018201602301121561134b5760643560a481013501810160040135906102b38261157c565b916102c160405193846114da565b80835260208301913660643560a481013501820160a08402016024011161134b5760643560a481013501810160240192905b60643560a481013501810160a0840201602401841061135057505050508152602460a460643501356064350101356020820152608082015261033a3660c4606435016116a0565b60a0820152610104606435013560c082015261035b610124606435016116c8565b60e082015261036f610144606435016116c8565b610100820152610384610164606435016116c8565b6101208201526064356101848101356101408301526101a40135610160820152604051916103b183611486565b608435835260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5c36011261134b576040516103ec816114a2565b67ffffffffffffffff60a4351660a4350361134b5760a435815267ffffffffffffffff60c4351660c4350361134b5760c435602082015260e4356040820152602084015261010435604084015261012435606084015261014435608084015263ffffffff6101643516610164350361134b576101643560a084015261ffff6104726116d9565b168061119a575060015b908160151461113c578160141461108f5781600314610fc45781600514610f6e5781600414610f5f5781601314610e215781601214610df15781601114610dd25781601014610d905781600f14610d585781600e14610d255781600d14610d025781600c14610ced5781600b14610ce15781600a146108a657508060091461088457806008146108075780600714610783578060061461073957806002146107335760011461053b57634e487b7160e01b600052605160045260246000fd5b600281525b6040519061012082528051600381101561071d5761012083015261057660208201516101c06101408501526102e08401906113cb565b9060206040820151805161016086015201516101808401526105c96060820151927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee09384868303016101a08701526113cb565b608082015192848203016101c085015282519260408252835180604084015260206060840195019060005b8181106106d6575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c0860152805160208601526040602082015167ffffffffffffffff8151168288015267ffffffffffffffff602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b516106f48482516113b3565b8581015160408501528260408201511682850152015116608082015201970191019190916105f4565b634e487b7160e01b600052602160045260246000fd5b50610540565b5061077e608082015161074a61203a565b506107596001825151146121b2565b6107638151612073565b519060405190610772826114be565b600082525251826124d8565b610540565b5061079b602082015161079583612427565b90611e46565b6107de6107ab60808301516121fd565b6107c6602084015161079563ffffffff6040850151166124a4565b61079563ffffffff60606020860151930151166124a4565b6101a4356107f463ffffffff82169182146116f7565b6101008201526000610120820152610540565b50610819602082015161079583612427565b610834602082015161079563ffffffff60e0850151166124a4565b61084f602082015161079563ffffffff60a0860151166124a4565b63ffffffff6101a4356108658160401c15611dfb565b818160201c1660e0840152166101008201526000610120820152610540565b50610896602082015161079583612427565b61084f6107ab60808301516121fd565b90506108bd6108b86020840151611ebf565b611f88565b6108c5611544565b5060606040516108d4816114be565b526108dd611544565b506108e6611525565b506000916000805b60208110610cb157509061099c61098b61097a67ffffffffffffffff61094d61092b61093c9761091c611525565b5087600401356024890161223d565b87600499929901356024890161223d565b8760049392930135602489016125cc565b929091816040519961095e8b6114a2565b16895216602088015260408701528460040135602486016125cc565b8460049392930135602486016125cc565b8460049492940135602486016125cc565b9290916000939560005b60048110610c7b57509163ffffffff93916109f5979695936040519a6109cb8c611486565b8b5260208b015260408a0152606089015260808801521660a086015260248160040135910161228a565b508251907fffffffff0000000000000000000000000000000000000000000000000000000060208501518051604060208301519201516040519160208301937f4d656d6f72793a0000000000000000000000000000000000000000000000000085527fffffffffffffffff000000000000000000000000000000000000000000000000809260c01b16602785015260c01b16602f830152603782015260378152610a9e816114a2565b5190206040860151606087015160808801519160a089015193604051977f4d6f64756c653a0000000000000000000000000000000000000000000000000060208a0152602789015260478801526067870152608786015260a785015260e01b1660c783015260ab82528160e081011067ffffffffffffffff60e084011117610c65578160e0610b7c930160405260e08151602083012091610b4082820161144e565b6013828201527f4d6f64756c65206d65726b6c6520747265653a00000000000000000000000000610100820152019163ffffffff851690612643565b61016084015103610bfb5763ffffffff60a0819382610be794610ba660208901516107958a612427565b610bbd60208901516107958460e08c0151166124a4565b610bd3602089015161079584878d0151166124a4565b1660e0870152015116826101a43516611d15565b166101008201526000610120820152610540565b608460405162461bcd60e51b815260206004820152602260248201527f43524f53535f4d4f44554c455f494e5445524e414c5f4d4f44554c45535f524f60448201527f4f540000000000000000000000000000000000000000000000000000000000006064820152fd5b634e487b7160e01b600052604160045260246000fd5b9694610caa60019163ffffff00610c9a898b6004013560248d01612222565b3560f81c9160081b16179661222e565b97016109a6565b9093610cda600191610ccb87866004013560248801612222565b3560f81c9060081b179561222e565b91016108ee565b505061077e8282611d2d565b61077e91506024816004013591018484611742565b50506101a435610d1a63ffffffff82169182146116f7565b610120820152610540565b505063ffffffff610d3c6108b86020840151611ebf565b1615610540576101a435610d1a63ffffffff82169182146116f7565b61077e9150610d86906020610d7060808601516121fd565b0151602482600401359201906101a43590612154565b6020830151611e46565b9050610dcb610da26020840151611ebf565b916020610db260808601516121fd565b0192835190602483600401359301916101a435906120f5565b9052610540565b61077e9150610d86908451602482600401359201906101a43590612154565b610e1a9150610e036020840151611ebf565b845190602483600401359301916101a435906120f5565b8252610540565b505090610e316020830151611ebf565b9163ffffffff610e446020830151611ebf565b9381610e65610e5f610e596020870151611ebf565b97611f88565b92611f88565b9160405196610e738861146a565b87526101a435602088015216604086015216606084015260808101519182515191600183018311610f4957601f19610ec86001610eb181870161157c565b95610ebf60405197886114da565b0180865261157c565b0160005b818110610f3257505060005b84518051821015610f0d5790610ef081600193612096565b51610efb8287612096565b52610f068186612096565b5001610ed8565b5050929093610f2b9082515190610f248286612096565b5283612096565b5052610540565b602090610f3d61203a565b82828801015201610ecc565b634e487b7160e01b600052601160045260246000fd5b50506107336020820151611ebf565b5050610f806108b86020830151611ebf565b610f8d6020830151611ebf565b63ffffffff610f9f6020850151611ebf565b921615610fb5575061077e906020830151611e46565b61077e91506020830151611e46565b505061ffff610fd16116d9565b1660418103611014575061077e60005b6020830151610ffc60405192610ff68461144e565b836116eb565b67ffffffffffffffff6101a435166020830152611e46565b60428103611027575061077e6001610fe1565b6043810361103a575061077e6002610fe1565b60440361104b5761077e6003610fe1565b606460405162461bcd60e51b815260206004820152601960248201527f434f4e53545f505553485f494e56414c49445f4f50434f4445000000000000006044820152fd5b5050600060206040516110a18161144e565b828152015261800561ffff6110b46116d9565b16036110d45761077e6110ca6020830151611ebf565b6060830151611e46565b61800661ffff6110e26116d9565b16036110f85761077e610d866060830151611ebf565b606460405162461bcd60e51b815260206004820152601c60248201527f4d4f56455f494e5445524e414c5f494e56414c49445f4f50434f4445000000006044820152fd5b50506020810151600060206040516111538161144e565b82815201525180515180600019810111610f495761077e9160001961118f92600060206040516111828161144e565b8281520152019051612096565b516020830151611e46565b600181036111aa5750600261047c565b600f81036111ba5750600661047c565b601081036111ca5750600761047c565b61800981036111db5750600861047c565b61800b81036111ec5750600961047c565b61800c81036111fd5750600a61047c565b61800a810361120e5750600b61047c565b6011810361121e5750600c61047c565b618003810361122f5750600d61047c565b61800481036112405750600e61047c565b602081036112505750600f61047c565b602181036112605750601061047c565b602381036112705750601161047c565b602481036112805750601261047c565b61800281036112915750601361047c565b601a81036112a15750600461047c565b601b81036112b15750600561047c565b604181101580611340575b156112c95750600361047c565b61800581148015611335575b156112e25750601461047c565b618008036112f157601561047c565b606460405162461bcd60e51b815260206004820152600e60248201527f494e56414c49445f4f50434f44450000000000000000000000000000000000006044820152fd5b5061800681146112d5565b5060448111156112bc565b600080fd5b60a08436031261134b5760a080602060249460405161136e8161146a565b611378368a611594565b815260408901358382015261138f60608a016116c8565b60408201526113a060808a016116c8565b60608201528152019501949250506102f3565b8051600781101561071d578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b83821061140f57505050505060208091015191015290565b90919293968382826114246001948c516113b3565b01980194939201906113f7565b610180810190811067ffffffffffffffff821117610c6557604052565b6040810190811067ffffffffffffffff821117610c6557604052565b6080810190811067ffffffffffffffff821117610c6557604052565b60c0810190811067ffffffffffffffff821117610c6557604052565b6060810190811067ffffffffffffffff821117610c6557604052565b6020810190811067ffffffffffffffff821117610c6557604052565b90601f601f19910116810190811067ffffffffffffffff821117610c6557604052565b6040519061150a8261144e565b600060208360405161151b816114be565b6060815281520152565b60405190611532826114a2565b60006040838281528260208201520152565b6040519061155182611486565b600060a083828152611561611525565b60208201528260408201528260608201528260808201520152565b67ffffffffffffffff8111610c655760051b60200190565b919082604091031261134b576040516115ac8161144e565b80928035600781101561134b578252602090810135910152565b60409291818103841361134b57604051916115e08361144e565b829481359267ffffffffffffffff9384811161134b57830191602094858484031261134b5760405193611612856114be565b803591821161134b570182601f8201121561134b5780356116328161157c565b9361164060405195866114da565b818552878086019260061b8401019281841161134b579088809897969594939201915b83831061167a575050505050815284520135910152565b97849596979861168e838596979495611594565b81520192019088979695949392611663565b919082604091031261134b576040516116b88161144e565b6020808294803584520135910152565b359063ffffffff8216820361134b57565b6101843561ffff8116810361134b5790565b600782101561071d5752565b156116fe57565b606460405162461bcd60e51b815260206004820152600d60248201527f4241445f43414c4c5f44415441000000000000000000000000000000000000006044820152fd5b9091926000936020958684019361175c6108b88651611ebf565b60409860608a5161176c816114be565b52889989978a809c5b6008809210156117b25760019167ffffffffffffff006117a89261179a8e8e8e612222565b3560f81c921b16179a61222e565b9c01809c99611775565b88979d50899399915093611801959b61180b9399956117f7989e986117f06117e9859f6117e090888c6125cc565b9788919b612222565b359561222e565b878c61223d565b878c9a929a6125cc565b878c97929761228a565b948d51948501947f43616c6c20696e6469726563743a000000000000000000000000000000000000865267ffffffffffffffff957fffffffffffffffff00000000000000000000000000000000000000000000000091828660c01b16602e8201528b60368201526036815261187f816114a2565b5190206101a43503611cd157928f91928f94611944948d89928851917fff00000000000000000000000000000000000000000000000000000000000000888401947f5461626c653a0000000000000000000000000000000000000000000000000000865216602684015260c01b1660278201528b602f820152602f8152611905816114a2565b519020927f5461626c65206d65726b6c6520747265653a00000000000000000000000000008751956119368761144e565b601287528601521690612643565b91015103611c8e5763ffffffff809c169516851015611c7d5761198890868b8b5161196e8161144e565b828152015260608a51611980816114be565b5283886125cc565b949092868b8b516119988161144e565b82815201526119b56119ab87838b612222565b3560f81c9661222e565b60068711611c3a576119c890828a6125cc565b9190986007881015611c26578c92611a8894926119ff926119f48f519b6119ee8d61144e565b8c6116eb565b858b019c8d5261228a565b5090611a0a8861272a565b8c51848101917f5461626c6520656c656d656e743a000000000000000000000000000000000000835288602e830152604e820152604e8152611a4b8161146a565b519020917f5461626c6520656c656d656e74206d65726b6c6520747265653a0000000000008d5194611a7c8661144e565b601a8652850152612643565b03611be35703611bd65780516007811015611bc257600403611ab1575050505050505060029052565b51906007821015611bae5750600503611b6b575192848416938403611b28575091611b1d849261079561012096606060009997611af281516107958b612427565b611b13611b0260808b01516121fd565b9561079585845192890151166124a4565b51930151166124a4565b166101008201520152565b60649083519062461bcd60e51b82526004820152601560248201527f4241445f46554e435f5245465f434f4e54454e545300000000000000000000006044820152fd5b60648484519062461bcd60e51b82526004820152600d60248201527f4241445f454c454d5f54595045000000000000000000000000000000000000006044820152fd5b80634e487b7160e01b602492526021600452fd5b602483634e487b7160e01b81526021600452fd5b5050505050505060029052565b60648888519062461bcd60e51b82526004820152601160248201527f4241445f454c454d454e54535f524f4f540000000000000000000000000000006044820152fd5b602489634e487b7160e01b81526021600452fd5b60648c8c519062461bcd60e51b82526004820152600e60248201527f4241445f56414c55455f545950450000000000000000000000000000000000006044820152fd5b505050505050505050505060029052565b60648b8b519062461bcd60e51b82526004820152600f60248201527f4241445f5441424c45535f524f4f5400000000000000000000000000000000006044820152fd5b5060648f8f519062461bcd60e51b82526004820152601660248201527f4241445f43414c4c5f494e4449524543545f44415441000000000000000000006044820152fd5b91909163ffffffff80809416911601918211610f4957565b90611b1391611d686020820192611d48845161079585612427565b610795845160a060e086019663ffffffff988993610795858b51166124a4565b611d7560808201516121fd565b9260608401928184511615611df1576101a435828116908103611dad578261012095611b1d9382604060009a01511690525116611d15565b606460405162461bcd60e51b815260206004820152601d60248201527f4241445f43414c4c45525f494e5445524e414c5f43414c4c5f444154410000006044820152fd5b5050600290525050565b15611e0257565b606460405162461bcd60e51b815260206004820152601a60248201527f4241445f43524f53535f4d4f44554c455f43414c4c5f444154410000000000006044820152fd5b518051519160019260018101809111610f4957611e6290612520565b926000815b611e83575b5050815151611e7f91610f248286612096565b5052565b83518051821015611eb95781611e9b84938493612096565b51611ea68289612096565b52611eb18188612096565b500190611e67565b50611e6c565b90602091604051611ecf8161144e565b600093818580935201525191806020604051611eea8161144e565b8281520152825191825160001993848201918211611f745790611f0c91612096565b5192845151908101908111611f6057611f2490612520565b915b8251811015611f5a5780611f3d6001928751612096565b51611f488286612096565b52611f538185612096565b5001611f26565b50925290565b602483634e487b7160e01b81526011600452fd5b602484634e487b7160e01b81526011600452fd5b60208101519051600781101561071d57611ff657640100000000811015611fb25763ffffffff1690565b606460405162461bcd60e51b815260206004820152600760248201527f4241445f493332000000000000000000000000000000000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152600760248201527f4e4f545f493332000000000000000000000000000000000000000000000000006044820152fd5b604051906120478261146a565b60006060836040516120588161144e565b83815283602082015281528260208201528260408201520152565b8051156120805760200190565b634e487b7160e01b600052603260045260246000fd5b80518210156120805760209160051b010190565b156120b157565b606460405162461bcd60e51b815260206004820152601160248201527f57524f4e475f4d45524b4c455f524f4f540000000000000000000000000000006044820152fd5b61214c9061214661213d6121519796949596600060206040516121178161144e565b8281520152606060405161212a816114be565b526121358187612331565b91909661228a565b50938585612582565b146120aa565b612582565b90565b9061215192612146916000602060405161216d8161144e565b8281520152600060206040516121828161144e565b82815201526060604051612195816114be565b526121ac6121a38784612331565b9097889461228a565b50612582565b156121b957565b606460405162461bcd60e51b815260206004820152601160248201527f4241445f57494e444f575f4c454e4754480000000000000000000000000000006044820152fd5b61221e9061220961203a565b506122186001825151146121b2565b51612073565b5190565b90821015612080570190565b6000198114610f495760010190565b9192600091825b600890818510156122815760019167ffffffffffffff006122789261226a8a878b612222565b3560f81c921b16179661222e565b93019294612244565b95945050509050565b90929192606060405161229c816114be565b526122b56122ab858385612222565b3560f81c9461222e565b906122bf8561157c565b926122cd60405194856114da565b858452601f196122dc8761157c565b0136602086013760005b60ff811693878510156123185760ff9161231061230660019387876125cc565b9190919789612096565b5201166122e6565b959650505050506040519061232c826114be565b815291565b906040519061233f8261144e565b6000928383528360208094015281156124135760f891813560f81c93600685116123cf5785936001938593859392915b83861061239f57505050505050926007831015611bae575061239760405192610ff68461144e565b602082015291565b909192936123c387986123b58398998686612222565b35861c9060081b179861222e565b9601949392919061236f565b606460405162461bcd60e51b815260206004820152600e60248201527f4241445f56414c55455f545950450000000000000000000000000000000000006044820152fd5b602484634e487b7160e01b81526032600452fd5b600060206040516124378161144e565b828152015263ffffffff610120820151169060e061010082015191015191600060206040516124658161144e565b828152015267ffffffff000000006bffffffff00000000000000006040519461248d8661144e565b6006865260401b169260201b161717602082015290565b600060206040516124b48161144e565b828152015263ffffffff604051916124cb8361144e565b6000835216602082015290565b908051600781101561071d57600414612519578051600781101561071d5760060361251957602061250b91015182612786565b156125135750565b60029052565b5060029052565b9061252a8261157c565b60409061253a60405191826114da565b838152601f1961254a829561157c565b019160009060005b848110612560575050505050565b602090825161256e8161144e565b848152828581830152828701015201612552565b9061258f6121519361272a565b906040519261259d8461144e565b601284527f56616c7565206d65726b6c6520747265653a00000000000000000000000000006020850152612643565b600093918491905b602083106125e157505050565b909193946126056001916125f6888587612222565b3560f81c9060081b179661222e565b940191906125d4565b9392909384519060005b82811061263057500191825260208201526040019150565b8060208092890101518184015201612618565b92939091936000925b84519586518510156126d95790600191829786898416156000146126ac57612690915061267d61269e918a51612096565b516040519283916020830195898761260e565b03601f1981018352826114da565b519020965b1c93019261264c565b6126bc6126d09161269093612096565b51926040519283916020830195898761260e565b519020966126a3565b955092505091506126e657565b606460405162461bcd60e51b815260206004820152600f60248201527f50524f4f465f544f4f5f53484f525400000000000000000000000000000000006044820152fd5b805190600782101561071d57602001516040519060208201927f56616c75653a0000000000000000000000000000000000000000000000000000845260f81b6026830152602782015260278152612780816114a2565b51902090565b908060601c6127b95760e09063ffffffff90818116610120850152818160201c1661010085015260401c16910152600190565b505060009056fea264697066735822122097f9d9db61ac8498e0ddb7711023907cb60498baba7613d8797d9bd6cbcc67a164736f6c63430008190033",
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initialWasmModuleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576139d4908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c63a92cb5011461002857600080fd5b34610eb45760031936016101e08112610eb457606013610eb45767ffffffffffffffff60643511610eb4576101c060031960643536030112610eb4576101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c360112610eb45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7c360112610eb4576101c4359067ffffffffffffffff8211610eb45736602383011215610eb45767ffffffffffffffff826004013511610eb457366024836004013584010111610eb4576101608161010a600093610fc8565b8281526101156110b0565b602082015260405161012681610fe5565b838152836020820152604082015261013c6110b0565b606082015260405161014d81610fe5565b60608152836020820152608082015260405161016881610fe5565b83815283602082015260a08201528260c08201528260e082015282610100820152826101208201528261014082015201526101a16110f7565b506040516101ae81610fc8565b6003606435600401351015610eb45760643560040135815267ffffffffffffffff6024606435013511610eb4576101f036606435602481013501600401611179565b602082015261020436604460643501611253565b604082015267ffffffffffffffff6084606435013511610eb45761023336606435608481013501600401611179565b606082015267ffffffffffffffff60a4606435013511610eb457604060643560a48101350136036003190112610eb45760405161026f81610fe5565b60643560a4810135016004013567ffffffffffffffff8111610eb4573660643560a48101350182016023011215610eb45760643560a481013501810160040135906102b98261112f565b916102c7604051938461108d565b80835260208301913660643560a481013501820160a084020160240111610eb45760643560a481013501810160240192905b60643560a481013501810160a08402016024018410610eb957505050508152602460a46064350135606435010135602082015260808201526103403660c460643501611253565b60a0820152610104606435013560c08201526103616101246064350161127b565b60e08201526103756101446064350161127b565b61010082015261038a6101646064350161127b565b6101208201526064356101848101356101408301526101a401356101608201526040516103b681611001565b608435815260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5c360112610eb4576040516103f18161101d565b67ffffffffffffffff60a4351660a43503610eb45760a435815267ffffffffffffffff60c4351660c43503610eb45760c435602082015260e4356040820152602082015261010435604082015261012435606082015261014435608082015263ffffffff61016435166101643503610eb4576101643560a082015261047461128c565b61801061ffff8216101580610ea4575b15610dba575060085b80600b14610dae5780600a14610d995780600914610d8d5780600714610940578060061461092a578060051461091e578060041461090857806001146108ed576008146104ea57634e487b7160e01b600052605160045260246000fd5b6104f261128c565b926104fb6129e2565b506105046129e2565b506000916040519461051586610fe5565b60403687376040519161052783610fe5565b60403684376000945b600260ff8716101561056c57610551610566918660040135602488016138c8565b969061056060ff83168b6134ae565b5261349d565b94610530565b945091959490946000945b600260ff871610156105bf576105986105b9918660040135602488016138fb565b969067ffffffffffffffff6105b060ff84168c6134ae565b9116905261349d565b94610577565b9296909450949094604051946105d486610fe5565b855260208501526105e4846134bf565b60c0860151036108a95761ffff166180108114801561089e575b156108295750610630929161062b9161062291600481013591908290602401611306565b90838787613611565b6134bf565b60c08201525b604051906101208252805160038110156108135761012083015261066c60208201516101c06101408501526102e0840190610f62565b9060206040820151805161016086015201516101808401526106bf6060820151927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee09384868303016101a0870152610f62565b608082015192848203016101c085015282519260408252835180604084015260206060840195019060005b8181106107cc575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c0860152805160208601526040602082015167ffffffffffffffff8151168288015267ffffffffffffffff602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b516107ea848251610f4a565b8581015160408501528260408201511682850152015116608082015201970191019190916106ea565b634e487b7160e01b600052602160045260246000fd5b915050618012810361084457508061062b610630928461359f565b6180130361085a578061062b6106309284613552565b606460405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f474c4f42414c53544154455f4f50434f44450000000000006044820152fd5b5061801181146105fe565b606460405162461bcd60e51b815260206004820152601060248201527f4241445f474c4f42414c5f5354415445000000000000000000000000000000006044820152fd5b5091806024610903926004013591018484612357565b610636565b50918060246109039260040135910184846118c4565b50915060018152610636565b5091806024610903926004013591018484611446565b509161094a6113d6565b60405161095681611039565b6060905260405161096681611039565b606090526109726113d6565b91610160840151906109826110f7565b5061098b6110f7565b506109946110d8565b50816109a7600483013560248401613891565b6109b29691966110d8565b506109c5906004850135602486016138fb565b6109d7906004860135602487016138fb565b6109e9906004870135602488016138c8565b91604051936109f78561101d565b67ffffffffffffffff90811685521660208401526040830152610a22906004860135602487016138c8565b909190610a37906004870135602488016138c8565b610a49906004880135602489016138c8565b909190610a5e90600489013560248a01613740565b9490936040519b610a6e8d611001565b8c5260208c015260408b015260608a0152608089015263ffffffff1660a0880152610aa190600485013560248601613740565b909490610ab690600486013560248701613788565b91909780519060208101518051602082015191604001516040519160208301937f4d656d6f72793a0000000000000000000000000000000000000000000000000085527fffffffffffffffff000000000000000000000000000000000000000000000000809260c01b16602785015260c01b16602f830152603782015260378152610b408161101d565b51902090604081015190606081015160808201519160a00151926040519460208601967f4d6f64756c653a000000000000000000000000000000000000000000000000008852602787015260478601526067850152608784015260a783015260e01b7fffffffff000000000000000000000000000000000000000000000000000000001660c782015260ab8152610bd681611055565b519020610be16113d6565b610bf29163ffffffff89168b612f7f565b03610d4957600163ffffffff861601938463ffffffff871611610c6c57610c1885612e51565b15610ce2575050505060018451511b03610c9e575b610c3c63ffffffff8216612e51565b15610c825750505180519081600019810111610c6c57600019610c6092019061140f565b51610160820152610636565b634e487b7160e01b600052601160045260246000fd5b9163ffffffff610c93931690612ea6565b610160820152610636565b606460405162461bcd60e51b815260206004820152600a60248201527f57524f4e475f4c454146000000000000000000000000000000000000000000006044820152fd5b610d00939491816024610cfa93600401359101613788565b50612ea6565b14610c2d57606460405162461bcd60e51b815260206004820152601360248201527f57524f4e475f524f4f545f464f525f5a45524f000000000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152601360248201527f57524f4e475f524f4f545f464f525f4c454146000000000000000000000000006044820152fd5b5091506109038161138c565b5091806024610903926004013591018361131e565b5091506109038161129e565b61ffff811661802003610dcf5750600161048d565b61ffff811661802103610de45750600461048d565b61ffff811661802203610df95750600561048d565b61ffff811661802303610e0e5750600661048d565b61ffff811661802403610e235750600761048d565b61ffff811661803003610e385750600961048d565b61ffff811661803103610e4d5750600a61048d565b61ffff1661803203610e6057600b61048d565b606460405162461bcd60e51b815260206004820152601560248201527f494e56414c49445f4d454d4f52595f4f50434f444500000000000000000000006044820152fd5b5061801361ffff82161115610484565b600080fd5b60a084360312610eb45760405180608081011067ffffffffffffffff608083011117610f345760a06020602494836080849501604052610ef9368a611147565b8152604089013583820152610f1060608a0161127b565b6040820152610f2160808a0161127b565b60608201528152019501949250506102f9565b634e487b7160e01b600052604160045260246000fd5b80516007811015610813578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b838210610fa657505050505060208091015191015290565b9091929396838282610fbb6001948c51610f4a565b0198019493920190610f8e565b610180810190811067ffffffffffffffff821117610f3457604052565b6040810190811067ffffffffffffffff821117610f3457604052565b60c0810190811067ffffffffffffffff821117610f3457604052565b6060810190811067ffffffffffffffff821117610f3457604052565b6020810190811067ffffffffffffffff821117610f3457604052565b60e0810190811067ffffffffffffffff821117610f3457604052565b6080810190811067ffffffffffffffff821117610f3457604052565b90601f601f19910116810190811067ffffffffffffffff821117610f3457604052565b604051906110bd82610fe5565b60006020836040516110ce81611039565b6060815281520152565b604051906110e58261101d565b60006040838281528260208201520152565b6040519061110482611001565b600060a0838281526111146110d8565b60208201528260408201528260608201528260808201520152565b67ffffffffffffffff8111610f345760051b60200190565b9190826040910312610eb45760405161115f81610fe5565b809280356007811015610eb4578252602090810135910152565b604092918181038413610eb4576040519161119383610fe5565b829481359267ffffffffffffffff93848111610eb4578301916020948584840312610eb457604051936111c585611039565b8035918211610eb4570182601f82011215610eb45780356111e58161112f565b936111f3604051958661108d565b818552878086019260061b84010192818411610eb4579088809897969594939201915b83831061122d575050505050815284520135910152565b978495969798611241838596979495611147565b81520192019088979695949392611216565b9190826040910312610eb45760405161126b81610fe5565b6020808294803584520135910152565b359063ffffffff82168203610eb457565b6101843561ffff81168103610eb45790565b600019908160a082015151146112ff576101a435806112da5750610140810191808351146112d2576112d09252612a8d565b565b506002915052565b91610140820151036112ff576112f963ffffffff6112d0931682612a1b565b50612a8d565b6002915052565b90939293848311610eb4578411610eb4578101920390565b6000199081610140820151036113835760a081019182515114611383578383604061134a930151612cf2565b5182604011610eb45760407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06112d09401920190612cf2565b60029052505050565b600019610140820151036113b4576040816113ad60a06112d0940151612deb565b0151612deb565b60029052565b67ffffffffffffffff8111610f3457601f01601f191660200190565b604051906113e382610fe5565b601382527f4d6f64756c65206d65726b6c6520747265653a000000000000000000000000006020830152565b80518210156114235760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b91908201809211610c6c57565b9061144f6113d6565b9061016083015190602061146d61146882870151612fe0565b6130a9565b910161148063ffffffff8316825161315b565b1561189c57908686819493519260051c6307ffffff1661149f93613187565b50956040516114ad81611039565b606090526040516114bd81611039565b60608152966114ca6113d6565b61016088015180926114da6110f7565b506114e36110f7565b506114ec6110d8565b506114f89088866138c8565b6115006110d8565b5061150c9089876138fb565b61151a908a889793976138fb565b611528908b899893986138c8565b91604051976115368961101d565b67ffffffffffffffff8092168952166020880152604087015261155a908a886138c8565b611565908b896138c8565b611573908c8a9693966138c8565b611581908d8b959395613740565b9190986040519361159185611001565b845260208401526040830152606082019485526080820192835260a082019763ffffffff1688526115c3908c8a613740565b6115ce919c8a613788565b94909782519360208401518051602082015191604001516040519160208301937f4d656d6f72793a0000000000000000000000000000000000000000000000000085527fffffffffffffffff000000000000000000000000000000000000000000000000809260c01b16602785015260c01b16602f8301526037820152603781526116588161101d565b5190209360400151925190519151926040519460208601967f4d6f64756c653a000000000000000000000000000000000000000000000000008852602787015260478601526067850152608784015260a783015260e01b7fffffffff000000000000000000000000000000000000000000000000000000001660c782015260ab81526116e381611055565b5190206116ee6113d6565b6116ff9163ffffffff8c1688612f7f565b03610d495763ffffffff9b60019860018e8216019d8e911611610c6c578c9561172787612e51565b1561183557505050505060019051511b03610c9e575b60009561174988612e51565b15611814575085610c6c5791869060005b8183116117b7575050509061178e60209493926117806040519384928884019687612e71565b03601f19810183528261108d565b5190206101608201525b015190610c6c576117b163ffffffff6112d0931661320e565b90613242565b60405194856117cc8360208301938a85612e71565b03956117e0601f199788810183528261108d565b5190209461180860405191826117fc6020820195808c88612e71565b0390810183528261108d565b51902091811c9161175a565b955061182a925060209493915086600096612f7f565b610160820152611798565b909192955061184a939b506118539450613788565b50978989612ea6565b1461173d57606460405162461bcd60e51b815260206004820152601360248201527f57524f4e475f524f4f545f464f525f5a45524f000000000000000000000000006044820152fd5b5050505090506002915052565b90821015611423570190565b6000198114610c6c5760010190565b90929363ffffffff6118dc6114686020850151612fe0565b166118ed6114686020850151612fe0565b63ffffffff81169467ffffffffffffffff61191361190e6020880151612fe0565b6132c2565b16976101a4351596878061224d575b61223d5760208101809111610c6c5767ffffffffffffffff60208a01515116108015612231575b61222257606060405161195b81611039565b5261197581836307ffffff8660051c1660208c0151613187565b989195907fff000000000000000000000000000000000000000000000000000000000000006119a58286886118a9565b35166121de576119b4906118b5565b91156121ba576002965b6119ca84848188611306565b999098600081600314611e9057506002146119f557634e487b7160e01b600052605160045260246000fd5b60288a10611e4c5760009b60209b8d6008815b1015611a4657611a198e8e8e6118a9565b3560f81c9060081b67ffffffffffffff0016179c611a36906118b5565b60019e909e019d9c60088f611a08565b93989d50949990959a611a6192979c509d93989d36916122d5565b60208151910120600091829084611d6e575b67ffffffffffffffff8116611cbe575b5060405191602083019384526040830152606082015260608152611aa681611071565b5190206024359173ffffffffffffffffffffffffffffffffffffffff8316809303610eb4576020906024604051809581937f16bf557900000000000000000000000000000000000000000000000000000000835260048301525afa918215611cb257600092611c7e575b5003611c3a5760015b15611c2b57858710611be7578686810311610c6c576000995b602063ffffffff8c161080611bcd575b15611b9557611b7f8b63ffffffff9283611b738d8d611b6d8e611b688f878a1692611439565b611439565b916118a9565b3560f81c921690613399565b9a1663ffffffff8114610c6c5760010199611b32565b602096508692949a98506117b197506040939550611bc0916307ffffff6112d09b60051c1690613418565b920151015201519161320e565b50868803611be163ffffffff8d1688611439565b10611b42565b606460405162461bcd60e51b815260206004820152601160248201527f4241445f4d4553534147455f50524f4f460000000000000000000000000000006044820152fd5b50506002905295505050505050565b606460405162461bcd60e51b815260206004820152601460248201527f4241445f534551494e424f585f4d4553534147450000000000000000000000006044820152fd5b9091506020813d602011611caa575b81611c9a6020938361108d565b81010312610eb457519038611b10565b3d9150611c8d565b6040513d6000823e3d90fd5b602435915073ffffffffffffffffffffffffffffffffffffffff82168203610eb45767ffffffffffffffff6024611cf660209361337f565b73ffffffffffffffffffffffffffffffffffffffff6040519586948593636ab8cee160e11b8552166004840152165afa908115611cb257600091611d3c575b5038611a83565b90506020813d602011611d66575b81611d576020938361108d565b81010312610eb4575138611d35565b3d9150611d4a565b9260243573ffffffffffffffffffffffffffffffffffffffff81168103611e4857602067ffffffffffffffff6024611da58961337f565b73ffffffffffffffffffffffffffffffffffffffff60405195869485937f16bf5579000000000000000000000000000000000000000000000000000000008552166004840152165afa918215611e3c578092611e04575b505092611a73565b9091506020823d602011611e34575b81611e206020938361108d565b81010312611e315750513880611dfc565b80fd5b3d9150611e13565b604051903d90823e3d90fd5b5080fd5b606460405162461bcd60e51b815260206004820152601260248201527f4241445f534551494e424f585f50524f4f4600000000000000000000000000006044820152fd5b93989d92979c959a90506071819c959a92979c10612176578391836120bc575b816071116120b857611ee8367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8f8401607184016122d5565b6020815191012082156120a4578592600190875b602081106120795750505060506021604051937fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208601967fff00000000000000000000000000000000000000000000000000000000000000833516885260601b1682860152016035840137608582015260858152611f7b81611001565b5190206040519060208201928352604082015260408152611f9b8161101d565b519020906024359073ffffffffffffffffffffffffffffffffffffffff821680920361207557602090602460405180948193636ab8cee160e11b835260048301525afa928315611e3c57809361203e575b505003611ffa576001611b19565b606460405162461bcd60e51b815260206004820152601360248201527f4241445f44454c415945445f4d455353414745000000000000000000000000006044820152fd5b909192506020823d60201161206d575b8161205b6020938361108d565b81010312611e31575051903880611fec565b3d915061204e565b8380fd5b90919461209c60019161208d8886896118a9565b3560f81c9060081b17966118b5565b929101611efc565b602486634e487b7160e01b81526032600452fd5b8480fd5b915060243573ffffffffffffffffffffffffffffffffffffffff811681036120b857602067ffffffffffffffff60246120f48761337f565b73ffffffffffffffffffffffffffffffffffffffff6040519586948593636ab8cee160e11b8552166004840152165afa90811561216b578591612139575b5091611eb0565b90506020813d602011612163575b816121546020938361108d565b810103126120b8575138612132565b3d9150612147565b6040513d87823e3d90fd5b606460405162461bcd60e51b815260206004820152601160248201527f4241445f44454c415945445f50524f4f460000000000000000000000000000006044820152fd5b6101a4356001036121cd576003966119be565b505050505050935090506002915052565b606460405162461bcd60e51b815260206004820152601360248201527f554e4b4e4f574e5f494e424f585f50524f4f46000000000000000000000000006044820152fd5b50505050935090506002915052565b50601f83161515611949565b5050505050935090506002915052565b506004358a1015611922565b1561226057565b60405162461bcd60e51b815260206004820152601660248201527f554e4b4e4f574e5f505245494d4147455f50524f4f46000000000000000000006044820152606490fd5b3d156122d0573d906122b6826113ba565b916122c4604051938461108d565b82523d6000602084013e565b606090565b9291926122e1826113ba565b916122ef604051938461108d565b829481845281830111610eb4578281602093846000960137010152565b1561231357565b606460405162461bcd60e51b815260206004820152600c60248201527f4241445f505245494d41474500000000000000000000000000000000000000006044820152fd5b939091929361236c6114686020830151612fe0565b9261237d6114686020840151612fe0565b9463ffffffff8616601f861615908115916129be575b5080156129b2575b6129a65760606040516123ad81611039565b526123c787826307ffffff8960051c166020860151613187565b9691909281996060946123e56123de8284866118a9565b35916118b5565b956101a435806124cd57505060f81c61226057612421948161240693611306565b9390916124143686856122d5565b602081519101201461230c565b602063ffffffff831601908163ffffffff841611610c6c57836124559361245c9584116124c3575b63ffffffff1691611306565b36916122d5565b935b6000965b855188101561248a57612482600191896020818a01015160f81c91613399565b970196612462565b6117b194975091604060206124b4819563ffffffff97956307ffffff6112d09c60051c1690613418565b9201510152015192511661320e565b9092508290612449565b60018198959496981460001461255c575050928092916124f36124f89560f81c15612259565b611306565b92602060006040518685823780878101838152039060025afa15611cb257612523906000511461230c565b602063ffffffff831601908163ffffffff841611610c6c5783612455936125569584116124c35763ffffffff1691611306565b9361245e565b909a969594929391906002036129625761257e936124f3849260f81c15612259565b91909282602011610eb45783350361291e5760008060405184868237808581018381520390600a5afa6125af6122a5565b90156128da5780511561289657604081805181010312610eb45760406020820151910151907f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff000000018203612852578060051b811590828104602014821715610c6c5763ffffffff851610612627575b50505050505061245e565b6307ffffff60009b979b9460051c1660015b83811061282b5750506128155764010000000004828102928184041490151715610c6c57600091829160405190602082019260208452602060408401526020898401527f16a2a19edfe81f20d09b681922c813b4b63683508c2280b93829971f439f0d2b608084015260a083015260c082015260c081526126b981611055565b519060055afa6126c76122a5565b90156127d157602081510361278d5760208151910151906020811061277b575b5081604011610eb457602083013503612737578211610eb4576040519161270d83610fe5565b602083523690820111610eb45760406020910181830137600060408201529338808080808061261c565b606460405162461bcd60e51b815260206004820152601160248201527f4b5a475f50524f4f465f57524f4e475f5a0000000000000000000000000000006044820152fd5b6000199060200360031b1b16386126e7565b606460405162461bcd60e51b815260206004820152601360248201527f4d4f444558505f57524f4e475f4c454e475448000000000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152600d60248201527f4d4f444558505f4641494c4544000000000000000000000000000000000000006044820152fd5b634e487b7160e01b600052601260045260246000fd5b909460011b94600180821614612848575b60011c9060011b612639565b946001179461283c565b606460405162461bcd60e51b815260206004820152601360248201527f554e4b4e4f574e5f424c535f4d4f44554c5553000000000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152601660248201527f4b5a475f505245434f4d50494c455f4d495353494e47000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152601160248201527f494e56414c49445f4b5a475f50524f4f460000000000000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152601460248201527f4b5a475f50524f4f465f57524f4e475f484153480000000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152601560248201527f554e4b4e4f574e5f505245494d4147455f5459504500000000000000000000006044820152fd5b50506002905292505050565b50601f8616151561239b565b905060208101809111610c6c5767ffffffffffffffff602084015151161038612393565b604051906129ef82610fe5565b816040516129fc81610fe5565b60403682378152602060405191612a1283610fe5565b60403684370152565b61014081019160001980845103612a8457610120916bffffffff000000000000000060e085015160401b1667ffffffff0000000061010086015160201b16179363ffffffff80948192015116911601828111610c6c57821601818111610c6c5716179052600190565b50505050600090565b9060408083018051519260a085019081515194600019808714908115612c9d575b50612c91576080870193845196602095868901519960009a5b8a5180518d1015612bbf5790612adf8d60019361140f565b518a612aeb8251613948565b918d828201516060828401519301519151938401947f537461636b206672616d653a00000000000000000000000000000000000000008652602c850152604c8401527fffffffff00000000000000000000000000000000000000000000000000000000918260e091821b16606c8501521b16607082015260548152612b6f81611071565b519020908b51908b8201927f537461636b206672616d6520737461636b3a000000000000000000000000000084526032830152605290818301528152612bb481611071565b5190209b019a612ac7565b50919497995091949950889295515201928351908782015191805151516000915b898b838510612c1457505050505090606093929151525190868201525251928301525190612c0d82611039565b6060825252565b612c3d8584959698936000612c4394600197519251612c3281610fe5565b82815201525161140f565b51613948565b908b51908d8201927f56616c756520737461636b3a00000000000000000000000000000000000000008452602c830152604c90818301528152612c8581611071565b51902094019190612be0565b50600290955250505050565b9050811438612aae565b15612cae57565b606460405162461bcd60e51b815260206004820152601460248201527f57524f4e475f434f5448524541445f454d5054590000000000000000000000006044820152fd5b90600092600091825b60208110612dd1575090612d0f92916138c8565b506000198303612d3957612d238115612ca7565b612d31602083015115612ca7565b602082015252565b6040517f636f7468726561643a000000000000000000000000000000000000000000000060208201908152602982018590526049820183905290612d808160698101611780565b519020602083015114612d3157606460405162461bcd60e51b815260206004820152601260248201527f57524f4e475f434f5448524541445f504f5000000000000000000000000000006044820152fd5b9294612de460019161208d8886866118a9565b9301612cfb565b60009080516000198103612dfd575052565b602082810180516040517f636f7468726561643a00000000000000000000000000000000000000000000009381019384526029810194909452604984015291612e498160698101611780565b519020905252565b8015159081612e5e575090565b60001981019150808211610c6c57161590565b9392909384519060005b828110612e9357500191825260208201526040019150565b8060208092890101518184015201612e7b565b91926000936000925b8451958651851015612f2e579060019182978689841615600014612f01576117809150612ee0612ef3918a5161140f565b5160405192839160208301958987612e71565b519020965b1c930192612eaf565b612f11612f25916117809361140f565b519260405192839160208301958987612e71565b51902096612ef8565b95509250509150612f3b57565b606460405162461bcd60e51b815260206004820152600f60248201527f50524f4f465f544f4f5f53484f525400000000000000000000000000000000006044820152fd5b92939091936000925b8451958651851015612f2e579060019182978689841615600014612fc7576117809150612ee0612fb9918a5161140f565b519020965b1c930192612f88565b612f11612fd7916117809361140f565b51902096612fbe565b90602091604051612ff081610fe5565b60009381858093520152519180602060405161300b81610fe5565b8281520152825191825160001993848201918211613095579061302d9161140f565b5192845151908101908111613081576130459061382f565b915b825181101561307b578061305e600192875161140f565b51613069828661140f565b52613074818561140f565b5001613047565b50925290565b602483634e487b7160e01b81526011600452fd5b602484634e487b7160e01b81526011600452fd5b60208101519051600781101561081357613117576401000000008110156130d35763ffffffff1690565b606460405162461bcd60e51b815260206004820152600760248201527f4241445f493332000000000000000000000000000000000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152600760248201527f4e4f545f493332000000000000000000000000000000000000000000000000006044820152fd5b6020820190818311610c6c575167ffffffffffffffff161015908161317e575090565b601f9150161590565b9390929193606060405161319a81611039565b5260406131c06131b66131ad8689613891565b90968799613788565b9590959686613418565b910151036131ca57565b606460405162461bcd60e51b815260206004820152600e60248201527f57524f4e475f4d454d5f524f4f540000000000000000000000000000000000006044820152fd5b6000602060405161321e81610fe5565b828152015263ffffffff6040519161323583610fe5565b6000835216602082015290565b518051519160019260018101809111610c6c5761325e9061382f565b926000815b613286575b50508151516132829161327b828661140f565b528361140f565b5052565b835180518210156132bc578161329e8493849361140f565b516132a9828961140f565b526132b4818861140f565b500190613263565b50613268565b6020810151905160078110156108135760010361333b57680100000000000000008110156132f75767ffffffffffffffff1690565b606460405162461bcd60e51b815260206004820152600760248201527f4241445f493634000000000000000000000000000000000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152600760248201527f4e4f545f493634000000000000000000000000000000000000000000000000006044820152fd5b9060001967ffffffffffffffff80931601918211610c6c57565b909160208310156133d45782601f0392601f8411610c6c578360031b93840460081490601f141715610c6c5760ff809116831b921b19161790565b606460405162461bcd60e51b815260206004820152601560248201527f4241445f5345545f4c4541465f425954455f49445800000000000000000000006044820152fd5b9061349a9260405160208101917f4d656d6f7279206c6561663a00000000000000000000000000000000000000008352602c820152602c815261345a8161101d565b519020906040519261346b84610fe5565b601384527f4d656d6f7279206d65726b6c6520747265653a000000000000000000000000006020850152612f7f565b90565b60ff1660ff8114610c6c5760010190565b9060028110156114235760051b0190565b8051906020808351930151910151602081519101516040519260208401947f476c6f62616c2073746174653a000000000000000000000000000000000000008652602d850152604d8401527fffffffffffffffff000000000000000000000000000000000000000000000000809260c01b16606d84015260c01b166075820152605d815261354c81611071565b51902090565b602081019163ffffffff61357661146861356f61190e8751612fe0565b9551612fe0565b16916002831015611383575067ffffffffffffffff9160206135999201516134ae565b91169052565b906020820163ffffffff6135b66114688351612fe0565b1692600284101561138357506135de6112d093602067ffffffffffffffff93519401516134ae565b511690600060206040516135f181610fe5565b82815201526040519161360383610fe5565b600183526020830152613242565b9361146894939293602081019061362b6114688351612fe0565b9161363d63ffffffff98899251612fe0565b169360028510156137335760206136599101978316885161315b565b156137275750906307ffffff6136859260051c1693606060405161367c81611039565b52848751613187565b949190506101843561ffff8116809103610eb45761801081036136c3575050916136b6604094926136bd94516134ae565b5191613418565b91510152565b9194509450618011919250146000146136e3576136e091516134ae565b52565b606460405162461bcd60e51b815260206004820152601760248201527f4241445f474c4f42414c5f53544154455f4f50434f44450000000000000000006044820152fd5b60029052505050505050565b5060029052505050505050565b600093918491905b6004831061375557505050565b9091939461377f60019163ffffff0061376f8986886118a9565b3560f81c9160081b1617966118b5565b94019190613748565b90929192606060405161379a81611039565b526137b36137a98583856118a9565b3560f81c946118b5565b906137bd8561112f565b926137cb604051948561108d565b858452601f196137da8761112f565b0136602086013760005b60ff811693878510156138165760ff9161380e61380460019387876138c8565b919091978961140f565b5201166137e4565b959650505050506040519061382a82611039565b815291565b906138398261112f565b604090613849604051918261108d565b838152601f19613859829561112f565b019160009060005b84811061386f575050505050565b602090825161387d81610fe5565b848152828581830152828701015201613861565b90916000926000926000915b602083106138aa57505050565b909193946138bf60019161208d8885876118a9565b9401919061389d565b600093918491905b602083106138dd57505050565b909193946138f260019161208d8885876118a9565b940191906138d0565b9192600091825b6008908185101561393f5760019167ffffffffffffff00613936926139288a878b6118a9565b3560f81c921b1617966118b5565b93019294613902565b95945050509050565b805190600782101561081357602001516040519060208201927f56616c75653a0000000000000000000000000000000000000000000000000000845260f81b602683015260278201526027815261354c8161101d56fea26469706673582212206968a01e65110ef14351b31fc8cfd457d42d001b64089afae6591eb96b65189c64736f6c63430008190033",
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
	Bin: "0x60808060405234601557612143908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c63a92cb5011461002757600080fd5b34610f065760031936016101e08112610f0657606013610f065767ffffffffffffffff60643511610f06576101c060031960643536030112610f06576101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c360112610f065760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7c360112610f06576101c43567ffffffffffffffff8111610f065736602382011215610f0657806004013567ffffffffffffffff8111610f065736910160240111610f0657600061016061010261101a565b82815261010d6110e1565b602082015261011a61103b565b83815283602082015260408201526101306110e1565b606082015261013d61103b565b60608152836020820152608082015261015461103b565b83815283602082015260a08201528260c08201528260e08201528261010082015282610120820152826101408201520152600060a061019161105b565b82815261019c61107b565b838152836020820152836040820152602082015282604082015282606082015282608082015201526101cc61101a565b606435600401356003811015610f0657815267ffffffffffffffff6024606435013511610f065761020836606435602481013501600401611148565b602082015261021c36604460643501611210565b604082015267ffffffffffffffff6084606435013511610f065761024b36606435608481013501600401611148565b606082015267ffffffffffffffff60a4606435013511610f0657604060643560a48101350136036003190112610f065761028361103b565b60643560a4810135016004013567ffffffffffffffff8111610f06573660643560a48101350182016023011215610f065760643560a481013501810160040135906102d56102d083611101565b6110bb565b8281529160208301913660643560a481013501820160a084020160240111610f065760643560a481013501810160240192905b60643560a481013501810160a08402016024018410610f0b57505050508152602460a460643501356064350101356020820152608082015261034f3660c460643501611210565b60a0820152610104606435013560c082015261037061012460643501611234565b60e082015261038461014460643501611234565b61010082015261039961016460643501611234565b6101208201526064356101848101356101408301526101a401356101608201526103c161105b565b90608435825260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5c360112610f06576103f961107b565b67ffffffffffffffff60a4351660a43503610f065760a435815267ffffffffffffffff60c4351660c43503610f065760c435602082015260e4356040820152602083015261010435604083015261012435606083015261014435608083015263ffffffff61016435166101643503610f06576101643560a083015261ffff61047f611245565b16604581148015610efc575b15610d6c575060015b80600b14610c375780600a14610a7b5780600914610a325780600814610a0a5780600714610a005780600514610997578060031461092d578060061461092357806004146108c757806002146107cc5760011461050157634e487b7160e01b600052605160045260246000fd5b61050e60208201516117b3565b604561ffff61051b611245565b16036107565780519060078210156107365761053960209215611768565b015161074c5761055960015b610553602084015191611a3d565b90611874565b604051906101208252805160038110156107365761012083015261058f60208201516101c06101408501526102e0840190610fb4565b9060206040820151805161016086015201516101808401526105e26060820151927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee09384868303016101a0870152610fb4565b608082015192848203016101c085015282519260408252835180604084015260206060840195019060005b8181106106ef575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c0860152805160208601526040602082015167ffffffffffffffff8151168288015267ffffffffffffffff602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b5161070d848251610f9c565b85810151604085015282604082015116828501520151166080820152019701910191909161060d565b634e487b7160e01b600052602160045260246000fd5b6105596000610545565b605061ffff610763611245565b16036107885780519060078210156107365761078360016020931461171d565b610539565b606460405162461bcd60e51b815260206004820152600760248201527f4241445f45515a000000000000000000000000000000000000000000000000006044820152fd5b506107e26107dd60208301516117b3565b6118f4565b6107f26107dd60208401516117b3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffba61ffff61081e611245565b160161ffff81116108b15761087c9261086e9261ffff8316600281149081156108a6575b811561089b575b8115610890575b50156108815761086261086891611968565b91611968565b90611ec2565b610553602084015191612031565b610559565b63ffffffff9182169116611ec2565b600891501438610850565b600681149150610849565b600481149150610842565b634e487b7160e01b600052601160045260246000fd5b506108d86107dd60208301516117b3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9961ffff610904611245565b16019061ffff82116108b15761087c9163ffffffff6105459216611d54565b5061087c816114d5565b5061094361093e60208301516117b3565b6119bf565b61095361093e60208401516117b3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffaf61ffff61097f611245565b160161ffff81116108b15761087c9261086e92611ec2565b506109a861093e60208301516117b3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8761ffff6109d4611245565b160161ffff81116108b1576109f163ffffffff9161087c93611bda565b1661055360208401519161198f565b5061087c81611257565b5061087c63ffffffff610a2361093e60208501516117b3565b16610553602084015191611a3d565b5061087c610a466107dd60208401516117b3565b60ac61ffff610a53611245565b1603610a7057610a6290611968565b61055360208401519161198f565b63ffffffff16610a62565b5060c061ffff610a89611245565b1603610b8357600860005b60078110156107365780610b755763ffffffff5b610ab560208501516117b3565b91825160078110156107365703610b3157600160ff84161b806000198101116108b157600019810160208401511680602085015260ff60001981871601116108b157600160ff6000198161087c981601161b16610b19575b50506020830151611874565b60001901191660208201511760208201523880610b0d565b606460405162461bcd60e51b815260206004820152601960248201527f4241445f455854454e445f53414d455f545950455f54595045000000000000006044820152fd5b67ffffffffffffffff610aa8565b60c161ffff610b90611245565b1603610b9f5760106000610a94565b60c261ffff610bac611245565b1603610bbb5760086001610a94565b60c361ffff610bc8611245565b1603610bd75760106001610a94565b60c461ffff610be4611245565b1603610bf35760206001610a94565b606460405162461bcd60e51b815260206004820152601860248201527f494e56414c49445f455854454e445f53414d455f5459504500000000000000006044820152fd5b5060bc61ffff610c45611245565b1603610cd457600060025b610c5d60208401516117b3565b90815160078110156107365760078210156107365703610c905760078210156107365761087c9181526020830151611874565b606460405162461bcd60e51b815260206004820152601860248201527f494e56414c49445f5245494e544552505245545f5459504500000000000000006044820152fd5b60bd61ffff610ce1611245565b1603610cf05760016003610c50565b60be61ffff610cfd611245565b1603610d0c5760026000610c50565b60bf61ffff610d19611245565b1603610d285760036001610c50565b606460405162461bcd60e51b815260206004820152601360248201527f494e56414c49445f5245494e54455250524554000000000000000000000000006044820152fd5b604681101580610ef1575b15610d8457506002610494565b606781101580610ee6575b15610d9c57506004610494565b606a81101580610edb575b15610db457506006610494565b605181101580610ed0575b15610dcc57506003610494565b607981101580610ec5575b15610de457506005610494565b607c81101580610eba575b15610dfc57506007610494565b60a78103610e0c57506008610494565b60ac81148015610eb0575b15610e2457506009610494565b60c081101580610ea5575b15610e3c5750600a610494565b60bc8110159081610e99575b5015610e5557600b610494565b606460405162461bcd60e51b815260206004820152600e60248201527f494e56414c49445f4f50434f44450000000000000000000000000000000000006044820152fd5b60bf9150111538610e48565b5060c4811115610e2f565b5060ad8114610e17565b50608a811115610def565b50607b811115610dd7565b50605a811115610dbf565b506078811115610da7565b506069811115610d8f565b50604f811115610d77565b506050811461048b565b600080fd5b60a084360312610f065760405180608081011067ffffffffffffffff608083011117610f865760a06020602494836080849501604052610f4b368a611119565b8152604089013583820152610f6260608a01611234565b6040820152610f7360808a01611234565b6060820152815201950194925050610308565b634e487b7160e01b600052604160045260246000fd5b80516007811015610736578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b838210610ff857505050505060208091015191015290565b909192939683828261100d6001948c51610f9c565b0198019493920190610fe0565b60405190610180820182811067ffffffffffffffff821117610f8657604052565b604051906040820182811067ffffffffffffffff821117610f8657604052565b6040519060c0820182811067ffffffffffffffff821117610f8657604052565b604051906060820182811067ffffffffffffffff821117610f8657604052565b604051906020820182811067ffffffffffffffff821117610f8657604052565b90601f19601f604051930116820182811067ffffffffffffffff821117610f8657604052565b6110e961103b565b906110f261109b565b60608152825260006020830152565b67ffffffffffffffff8111610f865760051b60200190565b9190826040910312610f065761112d61103b565b918035906007821015610f0657602091845201356020830152565b9190604092604081830312610f065761115f61103b565b9367ffffffffffffffff928235848111610f06578301916020948584840312610f065761118a61109b565b938035918211610f06570182601f82011215610f065780356111ae6102d082611101565b93878086848152019260061b84010192818411610f06579088809897969594939201915b8383106111ea57505050505081528552013590830152565b9784959697986111fe838596979495611119565b815201920190889796959493926111d2565b9190826040910312610f0657602061122661103b565b928035845201356020830152565b359063ffffffff82168203610f0657565b6101843561ffff81168103610f065790565b6020810161126861093e82516117b3565b9161127661093e83516117b3565b9060006101843561ffff908181168091036114d1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401918183116114bd575081166003810361135e57505067ffffffffffffffff8094169081158015611321575b611317575060070b928315611301576112ff9381610553931660070b0516915b519161198f565b565b634e487b7160e01b600052601260045260246000fd5b6002905250505050565b507fffffffffffffffffffffffffffffffffffffffffffffffff800000000000000085841660070b1480156112d857506000198260070b146112d8565b6005810361139b57505067ffffffffffffffff809416908115611317575060070b928315611301576112ff9381610553931660070b0716916112f8565b600a81036113c657505050610553906112ff93603f67ffffffffffffffff80931691161b16916112f8565b600c81036113f157505050610553906112ff93603f67ffffffffffffffff80931691161c16916112f8565b600b810361141f57505050610553906112ff93603f67ffffffffffffffff80931660070b91161d16916112f8565b600d810361145b5750505061055390603f6112ff94169067ffffffffffffffff80911681818161144e86611b81565b161c16921b1617916112f8565b600e0361149557505061055390603f6112ff94169067ffffffffffffffff80911681818161148886611b81565b161b16921c1617916112f8565b9092936114a192611a69565b9190916114b557506105536112ff926112f8565b600290525050565b80634e487b7160e01b602492526011600452fd5b8280fd5b602081016114e66107dd82516117b3565b6114f36107dd83516117b3565b906000936101843561ffff90818116809103611719577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9601958187116114bd57506003959081168681036115bf57505063ffffffff8092169081158015611584575b6115795750840b8015611301576112ff94826105539416900b0516915b5191611a3d565b600290525050505050565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffff80000000838516870b148015611555575060001982870b14611555565b600581036115f557505063ffffffff8092169081156115795750840b8015611301576112ff94826105539416900b071691611572565b929592600a81036116205750505050610553906112ff93601f63ffffffff80931691161b1691611572565b600c81036116485750505050610553906112ff93601f63ffffffff80931691161c1691611572565b600b810361167257505050906112ff93601f6105539363ffffffff809416900b91161d1691611572565b91925090600d81036116ae5750505061055390601f6112ff94169063ffffffff8091168181816116a186611baf565b161c16921b161791611572565b600e036116e457505061055390601f6112ff94169063ffffffff8091168181816116d786611baf565b161b16921c161791611572565b6116f9919263ffffffff808097169116611a69565b91909161171057506112ff92610553911691611572565b60029052505050565b8680fd5b1561172457565b606460405162461bcd60e51b815260206004820152600760248201527f4e4f545f493634000000000000000000000000000000000000000000000000006044820152fd5b1561176f57565b606460405162461bcd60e51b815260206004820152600760248201527f4e4f545f493332000000000000000000000000000000000000000000000000006044820152fd5b906020916117bf61103b565b6000938185809352015251918060206117d661103b565b828152015282519182516000199384820191821161186057906117f891612094565b519284515190810190811161184c57611810906120be565b915b825181101561184657806118296001928751612094565b516118348286612094565b5261183f8185612094565b5001611812565b50925290565b602483634e487b7160e01b81526011600452fd5b602484634e487b7160e01b81526011600452fd5b5180515191600192600181018091116108b157611890906120be565b926000815b6118b8575b50508151516118b4916118ad8286612094565b5283612094565b5052565b835180518210156118ee57816118d084938493612094565b516118db8289612094565b526118e68188612094565b500190611895565b5061189a565b6020810151905160078110156107365761190e9015611768565b6401000000008110156119245763ffffffff1690565b606460405162461bcd60e51b815260206004820152600760248201527f4241445f493332000000000000000000000000000000000000000000000000006044820152fd5b6380000000811661197c5763ffffffff1690565b63ffffffff1667ffffffff000000001790565b6000602061199b61103b565b828152015267ffffffffffffffff6119b161103b565b916001835216602082015290565b6020810151905160078110156107365760016119db911461171d565b680100000000000000008110156119f95767ffffffffffffffff1690565b606460405162461bcd60e51b815260206004820152600760248201527f4241445f493634000000000000000000000000000000000000000000000000006044820152fd5b60006020611a4961103b565b828152015263ffffffff611a5b61103b565b916000835216602082015290565b909161ffff1680611a8657500167ffffffffffffffff1690600090565b60018103611aa057500367ffffffffffffffff1690600090565b60028103611aba57500267ffffffffffffffff1690600090565b60048103611aea57509067ffffffffffffffff809116918215611adf57160490600090565b505050600090600190565b60068103611b0f57509067ffffffffffffffff809116918215611adf57160690600090565b60078103611b1f57501690600090565b60088103611b2f57501790600090565b600903611b3d571890600090565b606460405162461bcd60e51b815260206004820152601660248201527f494e56414c49445f47454e455249435f42494e5f4f50000000000000000000006044820152fd5b9067ffffffffffffffff8092166040039182116108b157565b90600163ffffffff809316019182116108b157565b9063ffffffff8092166020039182116108b157565b9060001963ffffffff809316019182116108b157565b9061ffff1680611c46575060405b63ffffffff80821615159081611c23575b5015611c0d57611c0890611bc4565b611be8565b63ffffffff908116604003915081116108b15790565b67ffffffffffffffff9150600190611c3a84611bc4565b161b8316161538611bf9565b6001918291808303611ca657506000925b611c6057505090565b63ffffffff831682604082109182611c8f575b505015611c8a57611c848293611b9a565b92611c57565b505090565b67ffffffffffffffff92501b821616158238611c73565b600291925014611cf457606460405162461bcd60e51b815260206004820152600960248201527f4241442049556e4f7000000000000000000000000000000000000000000000006044820152fd5b81906000906000935b611d08575b50505090565b63ffffffff809216916040831015611d4e5767ffffffffffffffff84841b831616611d3e575b82146108b1578280920191611cfd565b93611d4890611b9a565b93611d2e565b50611d02565b9061ffff1680611dc0575060205b63ffffffff80821615159081611d9d575b5015611d8757611d8290611bc4565b611d62565b63ffffffff908116602003915081116108b15790565b67ffffffffffffffff9150600190611db484611bc4565b161b8316161538611d73565b6001918291808303611e1b57506000925b611dda57505090565b63ffffffff831682602082109182611e04575b505015611c8a57611dfe8293611b9a565b92611dd1565b67ffffffffffffffff92501b821616158238611ded565b600291925014611e6957606460405162461bcd60e51b815260206004820152600960248201527f4241442049556e4f7000000000000000000000000000000000000000000000006044820152fd5b81906000906000935b611e7c5750505090565b63ffffffff809216916020831015611d4e5767ffffffffffffffff84841b831616611eb2575b82146108b1578280920191611e72565b93611ebc90611b9a565b93611ea2565b9161ffff1680611edf575067ffffffffffffffff80911691161490565b60018103611efb575067ffffffffffffffff8091169116141590565b60028103611f1c575067ffffffffffffffff80911660070b911660070b1290565b60038103611f37575067ffffffffffffffff80911691161090565b60048103611f58575067ffffffffffffffff80911660070b911660070b1390565b60058103611f73575067ffffffffffffffff80911691161190565b60068103611f95575067ffffffffffffffff80911660070b911660070b131590565b60078103611fb1575067ffffffffffffffff8091169116111590565b60088103611fd3575067ffffffffffffffff80911660070b911660070b121590565b600903611fed5767ffffffffffffffff8091169116101590565b606460405162461bcd60e51b815260206004820152600a60248201527f424144204952454c4f50000000000000000000000000000000000000000000006044820152fd5b60209061203c61103b565b60008082529201829052156120705780602061205661103b565b828152015261206361103b565b9081526001602082015290565b80602061207b61103b565b828152015261208861103b565b90808252602082015290565b80518210156120a85760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b906120cb6102d083611101565b828152601f196120db8294611101565b01906000805b8381106120ee5750505050565b6020906120f961103b565b8381528284818301528286010152016120e156fea264697066735822122087e86405294a70a34908151f50593b6bc07c393a16ea84d9d180be4d44a9cd0264736f6c63430008190033",
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
	Bin: "0x60808060405234601557611aae908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c63a92cb5011461002757600080fd5b346108cc5760031936016101e081126108cc576060136108cc5767ffffffffffffffff606435116108cc576101c0600319606435360301126108cc576101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c3601126108cc5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7c3601126108cc576101c43567ffffffffffffffff81116108cc57366023820112156108cc5767ffffffffffffffff8160040135116108cc573660248260040135830101116108cc5761010360806109e0565b6000608052610110610a90565b60a05260405161011f816109fd565b6000808252602082015260c052610134610a90565b60e052604051610143816109fd565b60608152600060208201526101005260405161015e816109fd565b600080825260208201819052610120919091526101408190526101608190526101808190526101a08190526101c08190526101e081905260405160a0906101a481610a19565b8281526040516101b381610a35565b838152836020820152836040820152602082015282604082015282606082015282608082015201526040516101e7816109e0565b60036064356004013510156108cc5760643560040135815267ffffffffffffffff60246064350135116108cc5761022936606435602481013501600401610b02565b602082015261023d36604460643501610bdc565b604082015267ffffffffffffffff60846064350135116108cc5761026c36606435608481013501600401610b02565b606082015267ffffffffffffffff60a46064350135116108cc57604060643560a481013501360360031901126108cc576040516102a8816109fd565b60643560a4810135016004013567ffffffffffffffff81116108cc573660643560a481013501820160230112156108cc5760643560a481013501810160040135906102f282610ab8565b916103006040519384610a6d565b80835260208301913660643560a481013501820160a0840201602401116108cc5760643560a481013501810160240192905b60643560a481013501810160a084020160240184106108d157505050508152602460a46064350135606435010135602082015260808201526103793660c460643501610bdc565b60a0820152610104606435013560c082015261039a61012460643501610c04565b60e08201526103ae61014460643501610c04565b6101008201526103c361016460643501610c04565b6101208201526064356101848101356101408301526101a40135610160820152604051916103f083610a19565b608435835260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5c3601126108cc5760405161042b81610a35565b67ffffffffffffffff60a4351660a435036108cc5760a435815267ffffffffffffffff60c4351660c435036108cc5760c435602082015260e4356040820152602084015261010435604084015261012435606084015261014435608084015263ffffffff610164351661016435036108cc576101643560a08401526101843561ffff811681036108cc57602861ffff82161015806108bd575b15610828575060015b8060041461074b578060031461072357806002146107095760011461050257634e487b7160e01b600052605160045260246000fd5b806024610516926004013591018484610f98565b604051906101208252805160038110156106f35761012083015261054c60208201516101c06101408501526102e084019061097a565b90602060408201518051610160860152015161018084015261059f6060820151927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee09384868303016101a087015261097a565b608082015192848203016101c085015282519260408252835180604084015260206060840195019060005b8181106106ac575050509461016063ffffffff9360a093602080899a015191015260208482015180516101e08a0152015161020088015260c08101516102208801528460e082015116610240880152846101008201511661026088015284610120820151166102808801526101408101516102a088015201516102c0860152805160208601526040602082015167ffffffffffffffff8151168288015267ffffffffffffffff602082015116606088015201516080860152604081015182860152606081015160c0860152608081015160e08601520151166101008301520390f35b909195602060a060019263ffffffff60608b516106ca848251610962565b8581015160408501528260408201511682850152015116608082015201970191019190916105ca565b634e487b7160e01b600052602160045260246000fd5b5080602461071e926004013591018484610c22565b610516565b505061071e63ffffffff60208401515160101c166107456020840151916114ac565b906114e0565b505063ffffffff60208301515160101c1661078063ffffffff6107796107746020860151611331565b6113fa565b1682610c15565b602084810151015190919067ffffffffffffffff1682116107e957818060101b046201000014821517156107d35767ffffffffffffffff61071e9260101b166020850151526107456020840151916114ac565b634e487b7160e01b600052601160045260246000fd5b505061071e602082015160006020604051610803816109fd565b828152015260405190610815826109fd565b6000825263ffffffff60208301526114e0565b603661ffff82161015806108ae575b15610844575060026104cd565b61ffff8116603f03610858575060036104cd565b61ffff1660400361086a5760046104cd565b606460405162461bcd60e51b815260206004820152601560248201527f494e56414c49445f4d454d4f52595f4f50434f444500000000000000000000006044820152fd5b50603e61ffff82161115610837565b50603561ffff821611156104c4565b600080fd5b60a0843603126108cc5760405180608081011067ffffffffffffffff60808301111761094c5760a06020602494836080849501604052610911368a610ad0565b815260408901358382015261092860608a01610c04565b604082015261093960808a01610c04565b6060820152815201950194925050610332565b634e487b7160e01b600052604160045260246000fd5b805160078110156106f3578252602090810151910152565b906040918051926040835260608301935193602080604086015285518092526020608086019601926000905b8382106109be57505050505060208091015191015290565b90919293968382826109d36001948c51610962565b01980194939201906109a6565b610180810190811067ffffffffffffffff82111761094c57604052565b6040810190811067ffffffffffffffff82111761094c57604052565b60c0810190811067ffffffffffffffff82111761094c57604052565b6060810190811067ffffffffffffffff82111761094c57604052565b6020810190811067ffffffffffffffff82111761094c57604052565b90601f601f19910116810190811067ffffffffffffffff82111761094c57604052565b60405190610a9d826109fd565b6000602083604051610aae81610a51565b6060815281520152565b67ffffffffffffffff811161094c5760051b60200190565b91908260409103126108cc57604051610ae8816109fd565b8092803560078110156108cc578252602090810135910152565b6040929181810384136108cc5760405191610b1c836109fd565b829481359267ffffffffffffffff938481116108cc5783019160209485848403126108cc5760405193610b4e85610a51565b80359182116108cc570182601f820112156108cc578035610b6e81610ab8565b93610b7c6040519586610a6d565b818552878086019260061b840101928184116108cc579088809897969594939201915b838310610bb6575050505050815284520135910152565b978495969798610bca838596979495610ad0565b81520192019088979695949392610b9f565b91908260409103126108cc57604051610bf4816109fd565b6020808294803584520135910152565b359063ffffffff821682036108cc57565b919082018092116107d357565b9290916000610184359061ffff8216809203610f95575060368103610eb457506000916004905b60209384870190610c5a8251611331565b90815160078110156106f35760078210156106f35703610e70578501519267ffffffffffffffff97888516928983169560089360088810610e2a575b505063ffffffff610cad610774610cb79351611331565b166101a435610c15565b9787610cc3878b610c15565b9101998a51511610610e1c5750909291939560009160001996604093604051610ceb81610a51565b6060815297819a82955b888710610d19575050505050505050505091610d1391604093611560565b91510152565b9091929394959697989b8d8c610d2f8a85610c15565b918260051c918203610dd5575b5050601f1660ff9087811015610d9257610d5590611860565b8060031b908082048a14901517156107d35781811b19909216908e1690911b179b861c66ffffffffffffff16989796600101959493929190610cf5565b6064888c519062461bcd60e51b82526004820152601560248201527f4241445f5345545f4c4541465f425954455f49445800000000000000000000006044820152fd5b81601f98939f9188939f8e90610dfb97856000198c9703610e05575b50505050516116c2565b9b9c90958f610d3c565b610e0e92611560565b90825101528d388080610df1565b600290525050505050505050565b9091945060031b6807fffffffffffffff867fffffffffffffff88216911681036107d3578a6001600019921b16018a81116107d3571689169263ffffffff610cad610c96565b6064866040519062461bcd60e51b82526004820152600e60248201527f4241445f53544f52455f545950450000000000000000000000000000000000006044820152fd5b60378103610ec85750600191600890610c49565b60388103610edc5750600291600490610c49565b60398103610ef05750600391600890610c49565b603a8103610f045750600091600190610c49565b603b8103610f185750600091600290610c49565b603c8103610f2b57506001918290610c49565b603d8103610f3f5750600191600290610c49565b603e03610f5157600191600490610c49565b606460405162461bcd60e51b815260206004820152601b60248201527f494e56414c49445f4d454d4f52595f53544f52455f4f50434f444500000000006044820152fd5b80fd5b9092916000936101843561ffff95868216809203610f955750602881036111a55750600094600490610fed87965b60208701958463ffffffff966020610fe589610cad6107748d51611331565b91015161186e565b509590611198575067ffffffffffffffff9485811696611037575b5050505051906040519261101b846109fd565b60078510156106f3576110359484521660208301526114e0565b565b90859394969291600192838114808091611183575b1561106a57505050505060ff91501660000b16915b38808080611008565b8061116d575b15611088575050505060ff1660000b16929050611061565b9294919392600281148080611158575b156110ab575050505016900b1691611061565b9296919280611142575b156110c75750505016900b1691611061565b9195509392506004915014908161112f575b50156110eb571660030b811691611061565b606460405162461bcd60e51b815260206004820152601560248201527f4241445f524541445f42595445535f5349474e454400000000000000000000006044820152fd5b905060078610156106f3578514386110d9565b50955060078910156106f3578695858a146110b5565b50925060078a10156106f35787928a15611098565b50945060078910156106f3578694838a14611070565b50955060078a10156106f35787958a1561104c565b6002905250505050505050565b602981036111bf5750600194600890610fed600096610fc6565b602a81036111d95750600294600490610fed600096610fc6565b602b81036111f35750600394600890610fed600096610fc6565b602c810361120c5750600094600190610fed8296610fc6565b602d81036112255750600094600190610fed8796610fc6565b602e810361123f5750600094600290610fed600196610fc6565b602f81036112585750600094600290610fed8796610fc6565b6030810361127057506001948590610fed8796610fc6565b6031810361128957506001948590610fed600096610fc6565b603281036112a25750600194600290610fed8796610fc6565b603381036112bc5750600194600290610fed600096610fc6565b603481036112d55750600194600490610fed8796610fc6565b6035036112ed57600194600490610fed600096610fc6565b606460405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f4d454d4f52595f4c4f41445f4f50434f44450000000000006044820152fd5b90602091604051611341816109fd565b60009381858093520152519180602060405161135c816109fd565b82815201528251918251600019938482019182116113e6579061137e9161199c565b51928451519081019081116113d257611396906119c6565b915b82518110156113cc57806113af600192875161199c565b516113ba828661199c565b526113c5818561199c565b5001611398565b50925290565b602483634e487b7160e01b81526011600452fd5b602484634e487b7160e01b81526011600452fd5b6020810151905160078110156106f357611468576401000000008110156114245763ffffffff1690565b606460405162461bcd60e51b815260206004820152600760248201527f4241445f493332000000000000000000000000000000000000000000000000006044820152fd5b606460405162461bcd60e51b815260206004820152600760248201527f4e4f545f493332000000000000000000000000000000000000000000000000006044820152fd5b600060206040516114bc816109fd565b828152015263ffffffff604051916114d3836109fd565b6000835216602082015290565b5180515191600192600181018091116107d3576114fc906119c6565b926000815b611524575b505081515161152091611519828661199c565b528361199c565b5052565b8351805182101561155a578161153c8493849361199c565b51611547828961199c565b52611552818861199c565b500190611501565b50611506565b926040916040519360209460208101917f4d656d6f7279206c6561663a00000000000000000000000000000000000000008352602c820152602c81526115a581610a35565b51902092604051916115b6836109fd565b601383527f4d656d6f7279206d65726b6c6520747265653a000000000000000000000000006020840152936000945b875191825187101561166e579060019182938885841615600014611643576116279150611616611635918d5161199c565b5187519283918d8301958b87611a28565b03601f198101835282610a6d565b519020925b1c9501946115e5565b611653611665916116279361199c565b519287519283918d8301958b87611a28565b5190209261163a565b96509450505050925061167e5790565b606460405162461bcd60e51b815260206004820152600f60248201527f50524f4f465f544f4f5f53484f525400000000000000000000000000000000006044820152fd5b919294909394604091606083516116d881610a51565b526000969560005b6020811061182f5750916117219695939188959360ff99916060855161170581610a51565b52611711898589611a5d565b359460f89560f81c9a8b9a611a69565b9661172b8b610ab8565b9961173883519b8c610a6d565b8b8b526117448c610ab8565b60209990601f1901368d8c013760009e8f5b169d8e10156117b75760009060005b898d8d831061178f575050509d61177c908d61199c565b5260ff6001819f01169d9c8c9d8f611756565b936117a2846117b0936001959697611a5d565b358d1c9060081b1793611a69565b9101611765565b949c50949c50965096509697508593506117e192508351976117d889610a51565b88529787611560565b910151036117ec5750565b6064905162461bcd60e51b815260206004820152600e60248201527f57524f4e475f4d454d5f524f4f540000000000000000000000000000000000006044820152fd5b9593969297611854600191611845868c89611a5d565b3560f81c9060081b1794611a69565b939893979496016116e0565b601f0390601f82116107d357565b94929190946000946118808388610c15565b67ffffffffffffffff8351161061198f5791906000199693968693879888965b8588106118b65750505050505050506000929190565b90919293949596996118c88b83610c15565b8060051c868103611971575b50601f1660208082101561192d57506118ec90611860565b9060039180831b926008918085048314901517156107d3578d901b908d8204148d1517156107d35760ff8a6001941c16901b179a01969594939291906118a0565b6064906040519062461bcd60e51b82526004820152601660248201527f4241445f50554c4c5f4c4541465f425954455f494458000000000000000000006044820152fd5b98509450611984601f9a85858b8a6116c2565b509a909895906118d4565b5050505091506001918190565b80518210156119b05760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b906119d082610ab8565b6040906119e06040519182610a6d565b838152601f196119f08295610ab8565b019160009060005b848110611a06575050505050565b6020908251611a14816109fd565b8481528285818301528287010152016119f8565b9392909384519060005b828110611a4a57500191825260208201526040019150565b8060208092890101518184015201611a32565b908210156119b0570190565b60001981146107d3576001019056fea2646970667358221220e4ebe1600f5204e8cadf09370e7b0e34a7f99f036e88e6f8acd00837c9fc9b6864736f6c63430008190033",
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
