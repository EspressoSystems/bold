// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package celestiagen

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

// AttestationProof is an auto generated low-level Go binding around an user-defined struct.
type AttestationProof struct {
	TupleRootNonce *big.Int
	Tuple          DataRootTuple
	Proof          BinaryMerkleProof
}

// BinaryMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type BinaryMerkleProof struct {
	SideNodes [][32]byte
	Key       *big.Int
	NumLeaves *big.Int
}

// DataRootTuple is an auto generated low-level Go binding around an user-defined struct.
type DataRootTuple struct {
	Height   *big.Int
	DataRoot [32]byte
}

// Namespace is an auto generated low-level Go binding around an user-defined struct.
type Namespace struct {
	Version [1]byte
	Id      [28]byte
}

// NamespaceNode is an auto generated low-level Go binding around an user-defined struct.
type NamespaceNode struct {
	Min    Namespace
	Max    Namespace
	Digest [32]byte
}

// CelestiaBatchVerifierMetaData contains all meta data concerning the CelestiaBatchVerifier contract.
var CelestiaBatchVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedHeights\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blobstream\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"min\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes1\",\"name\":\"version\",\"type\":\"bytes1\"},{\"internalType\":\"bytes28\",\"name\":\"id\",\"type\":\"bytes28\"}],\"internalType\":\"structNamespace\",\"name\":\"max\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"internalType\":\"structNamespaceNode\",\"name\":\"_rowRoot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_rowProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structAttestationProof\",\"name\":\"_attestationProof\",\"type\":\"tuple\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"proofHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"proofDataRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"rowProof\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601957610a6e908161001f823930815050f35b600080fdfe60806040818152600436101561001457600080fd5b600090813560e01c63d5a784cd1461002b57600080fd5b61010090600319828136011261022c576004356001600160a01b0381169490859003610229576023199136830160a0811261022557849061006b89610230565b126102215783519661007c88610261565b6001600160f81b0319602435818116810361021d5789526044359863ffffffff1990818b168b036102195760209a8b82015283528660631936011261021d578651916100c783610261565b6064359081168103610219578252608435908116810361021d57818a01528189015260a435858201526001600160401b039660c43588811161021d576101119036906004016102d1565b9460e435908982116102195760808236039586011261021957879081519561013887610230565b83600401358752011261021d578651986101518a610261565b60248201358a5260448201358b8b01528a8501998a526064820135908111610219579161018a6101ac95949260048995369201016102d1565b8885015285888c8282519161019e83610230565b606083528201520152610375565b509451868151910151908451961515875287870152838601526080606086015260e085018683519560606080890152865180935287019501915b81811061020557505050839481015160a0850152015160c08301520390f35b8251865294870194918701916001016101e6565b8580fd5b8480fd5b5080fd5b8280fd5b80fd5b8380fd5b606081019081106001600160401b0382111761024b57604052565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b0382111761024b57604052565b608081019081106001600160401b0382111761024b57604052565b601f909101601f19168101906001600160401b0382119082101761024b57604052565b6001600160401b03811161024b5760051b60200190565b91909160608184031261037057604051906102eb82610230565b909283919081356001600160401b03811161037057820181601f8201121561037057602091813561031b816102ba565b926103296040519485610297565b818452848085019260051b820101928311610370578401905b8282106103615750505083528082013590830152604090810135910152565b81358152908401908401610342565b600080fd5b9291909180519060208092818301966040885194015183604051958694631f3302a960e01b8652600486015280516024860152015160448401526080606484015260e483019080519160606084860152825180915285610104860193019060005b818110610477575050508481015160a48501526040015160c48401528290039082906001600160a01b03165afa90811561046b57600091610435575b5015610429576104259351015191610490565b9091565b50505050600090600490565b8281813d8311610464575b61044a8183610297565b810103126102215751908115158203610229575038610412565b503d610440565b6040513d6000823e3d90fd5b82518552899789975094850194909201916001016103d6565b906104e09261049f83516104f6565b60406104ae60208601516104f6565b940151906040519462ffffff19809216602087015216603d850152605a840152605a83526104db8361027c565b61057f565b50156104ee57600190600090565b600090600390565b80516020918201516040516001600160f81b031990921692820192835263ffffffff19166021820152601d815261052c81610261565b51905162ffffff19918282169190601d811061054a575b5050905090565b83919250601d0360031b1b1616803880610543565b6006111561056957565b634e487b7160e01b600052602160045260246000fd5b60408201805193949360018111610678575082515161066b575b602083019081518151111561065d5760206105cc966105e66040516105da8160009b8c95868884015260218301906107a2565b03601f198101835282610297565b604051918280926107a2565b039060025afa1561065257855193519182511561063b579061060d939491519051906107f7565b600681101561062757806106215750149190565b92915050565b634e487b7160e01b85526021600452602485fd5b505160011415905061064c57149190565b50508190565b6040513d87823e3d90fd5b505050509050600090600290565b5050509050600090600190565b6106898451519160208601516106eb565b14610599575050509050600090600190565b906101009182039182116106ab57565b634e487b7160e01b600052601160045260246000fd5b6000198101919082116106ab57565b919082039182116106ab57565b60010190816001116106ab57565b906001908181111561079a5760005b8183821b10610779576101009081039081116106ab576107199061069b565b9282610724856106c1565b1b9061072f826106c1565b811161073c575050505090565b929350909183820361074e5750505090565b61076d93509061076181610767936106d0565b926106d0565b906106eb565b610776906106dd565b90565b828101809111156106fa57634e487b7160e01b600052601160045260246000fd5b505050600090565b9081519160005b8381106107ba575050016000815290565b80602080928401015181850152016107a9565b80518210156107e15760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b93929381156108de57600182146108c9578451156108be5782610819836108e9565b61082c61082688516106c1565b88610936565b92818110610889578161076161084c9693610846936106d0565b906107f7565b90916108578261055f565b816108825750506108768361087061087c9495516106c1565b906107cd565b516109fa565b90600090565b9350919050565b61089394506107f7565b909161089e8261055f565b816108825750506108b78361087061087c9495516106c1565b51906109fa565b505090915090600590565b50509091516108d85790600090565b90600490565b505090915090600390565b6001808210610370578180916000925b61091f57505060001981019081116106ab576001901b9081146109195790565b60011c90565b909160001981146106ab57810191811c90816108f9565b91908251811161099d57610949816102ba565b906109576040519283610297565b808252601f19610966826102ba565b0136602084013760005b81811061097e575090925050565b8061098b600192876107cd565b5161099682866107cd565b5201610970565b60405162461bcd60e51b815260206004820152602f60248201527f496e76616c69642072616e67653a205f626567696e206f72205f656e6420617260448201526e65206f7574206f6620626f756e647360881b6064820152608490fd5b610a2760009160209360405191600160f81b8684015260218301526041820152604181526105da8161027c565b039060025afa1561046b576000519056fea2646970667358221220b29c93d6934b3e5bb490bebbb02f14bd7b3d81bda4c5933cdd2ec5d087f25ac064736f6c63430008190033",
}

// CelestiaBatchVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use CelestiaBatchVerifierMetaData.ABI instead.
var CelestiaBatchVerifierABI = CelestiaBatchVerifierMetaData.ABI

// CelestiaBatchVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CelestiaBatchVerifierMetaData.Bin instead.
var CelestiaBatchVerifierBin = CelestiaBatchVerifierMetaData.Bin

// DeployCelestiaBatchVerifier deploys a new Ethereum contract, binding an instance of CelestiaBatchVerifier to it.
func DeployCelestiaBatchVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CelestiaBatchVerifier, error) {
	parsed, err := CelestiaBatchVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CelestiaBatchVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CelestiaBatchVerifier{CelestiaBatchVerifierCaller: CelestiaBatchVerifierCaller{contract: contract}, CelestiaBatchVerifierTransactor: CelestiaBatchVerifierTransactor{contract: contract}, CelestiaBatchVerifierFilterer: CelestiaBatchVerifierFilterer{contract: contract}}, nil
}

// CelestiaBatchVerifier is an auto generated Go binding around an Ethereum contract.
type CelestiaBatchVerifier struct {
	CelestiaBatchVerifierCaller     // Read-only binding to the contract
	CelestiaBatchVerifierTransactor // Write-only binding to the contract
	CelestiaBatchVerifierFilterer   // Log filterer for contract events
}

// CelestiaBatchVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type CelestiaBatchVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CelestiaBatchVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CelestiaBatchVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CelestiaBatchVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CelestiaBatchVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CelestiaBatchVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CelestiaBatchVerifierSession struct {
	Contract     *CelestiaBatchVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CelestiaBatchVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CelestiaBatchVerifierCallerSession struct {
	Contract *CelestiaBatchVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// CelestiaBatchVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CelestiaBatchVerifierTransactorSession struct {
	Contract     *CelestiaBatchVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// CelestiaBatchVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type CelestiaBatchVerifierRaw struct {
	Contract *CelestiaBatchVerifier // Generic contract binding to access the raw methods on
}

// CelestiaBatchVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CelestiaBatchVerifierCallerRaw struct {
	Contract *CelestiaBatchVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// CelestiaBatchVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CelestiaBatchVerifierTransactorRaw struct {
	Contract *CelestiaBatchVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCelestiaBatchVerifier creates a new instance of CelestiaBatchVerifier, bound to a specific deployed contract.
func NewCelestiaBatchVerifier(address common.Address, backend bind.ContractBackend) (*CelestiaBatchVerifier, error) {
	contract, err := bindCelestiaBatchVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CelestiaBatchVerifier{CelestiaBatchVerifierCaller: CelestiaBatchVerifierCaller{contract: contract}, CelestiaBatchVerifierTransactor: CelestiaBatchVerifierTransactor{contract: contract}, CelestiaBatchVerifierFilterer: CelestiaBatchVerifierFilterer{contract: contract}}, nil
}

// NewCelestiaBatchVerifierCaller creates a new read-only instance of CelestiaBatchVerifier, bound to a specific deployed contract.
func NewCelestiaBatchVerifierCaller(address common.Address, caller bind.ContractCaller) (*CelestiaBatchVerifierCaller, error) {
	contract, err := bindCelestiaBatchVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CelestiaBatchVerifierCaller{contract: contract}, nil
}

// NewCelestiaBatchVerifierTransactor creates a new write-only instance of CelestiaBatchVerifier, bound to a specific deployed contract.
func NewCelestiaBatchVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*CelestiaBatchVerifierTransactor, error) {
	contract, err := bindCelestiaBatchVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CelestiaBatchVerifierTransactor{contract: contract}, nil
}

// NewCelestiaBatchVerifierFilterer creates a new log filterer instance of CelestiaBatchVerifier, bound to a specific deployed contract.
func NewCelestiaBatchVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*CelestiaBatchVerifierFilterer, error) {
	contract, err := bindCelestiaBatchVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CelestiaBatchVerifierFilterer{contract: contract}, nil
}

// bindCelestiaBatchVerifier binds a generic wrapper to an already deployed contract.
func bindCelestiaBatchVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CelestiaBatchVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CelestiaBatchVerifier *CelestiaBatchVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CelestiaBatchVerifier.Contract.CelestiaBatchVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CelestiaBatchVerifier *CelestiaBatchVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CelestiaBatchVerifier.Contract.CelestiaBatchVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CelestiaBatchVerifier *CelestiaBatchVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CelestiaBatchVerifier.Contract.CelestiaBatchVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CelestiaBatchVerifier *CelestiaBatchVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CelestiaBatchVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CelestiaBatchVerifier *CelestiaBatchVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CelestiaBatchVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CelestiaBatchVerifier *CelestiaBatchVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CelestiaBatchVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf27a509b.
//
// Solidity: function verifyProof(address _blobstream, ((bytes1,bytes28),(bytes1,bytes28),bytes32) _rowRoot, (bytes32[],uint256,uint256) _rowProof, (uint256,(uint256,bytes32),(bytes32[],uint256,uint256)) _attestationProof) view returns(bool isValid, uint256 proofHeight, bytes32 proofDataRoot, (bytes32[],uint256,uint256) rowProof)
func (_CelestiaBatchVerifier *CelestiaBatchVerifierCaller) VerifyProof(opts *bind.CallOpts, _blobstream common.Address, _rowRoot NamespaceNode, _rowProof BinaryMerkleProof, _attestationProof AttestationProof) (struct {
	IsValid       bool
	ProofHeight   *big.Int
	ProofDataRoot [32]byte
	RowProof      BinaryMerkleProof
}, error) {
	var out []interface{}
	err := _CelestiaBatchVerifier.contract.Call(opts, &out, "verifyProof", _blobstream, _rowRoot, _rowProof, _attestationProof)

	outstruct := new(struct {
		IsValid       bool
		ProofHeight   *big.Int
		ProofDataRoot [32]byte
		RowProof      BinaryMerkleProof
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ProofHeight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProofDataRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.RowProof = *abi.ConvertType(out[3], new(BinaryMerkleProof)).(*BinaryMerkleProof)

	return *outstruct, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xf27a509b.
//
// Solidity: function verifyProof(address _blobstream, ((bytes1,bytes28),(bytes1,bytes28),bytes32) _rowRoot, (bytes32[],uint256,uint256) _rowProof, (uint256,(uint256,bytes32),(bytes32[],uint256,uint256)) _attestationProof) view returns(bool isValid, uint256 proofHeight, bytes32 proofDataRoot, (bytes32[],uint256,uint256) rowProof)
func (_CelestiaBatchVerifier *CelestiaBatchVerifierSession) VerifyProof(_blobstream common.Address, _rowRoot NamespaceNode, _rowProof BinaryMerkleProof, _attestationProof AttestationProof) (struct {
	IsValid       bool
	ProofHeight   *big.Int
	ProofDataRoot [32]byte
	RowProof      BinaryMerkleProof
}, error) {
	return _CelestiaBatchVerifier.Contract.VerifyProof(&_CelestiaBatchVerifier.CallOpts, _blobstream, _rowRoot, _rowProof, _attestationProof)
}

// VerifyProof is a free data retrieval call binding the contract method 0xf27a509b.
//
// Solidity: function verifyProof(address _blobstream, ((bytes1,bytes28),(bytes1,bytes28),bytes32) _rowRoot, (bytes32[],uint256,uint256) _rowProof, (uint256,(uint256,bytes32),(bytes32[],uint256,uint256)) _attestationProof) view returns(bool isValid, uint256 proofHeight, bytes32 proofDataRoot, (bytes32[],uint256,uint256) rowProof)
func (_CelestiaBatchVerifier *CelestiaBatchVerifierCallerSession) VerifyProof(_blobstream common.Address, _rowRoot NamespaceNode, _rowProof BinaryMerkleProof, _attestationProof AttestationProof) (struct {
	IsValid       bool
	ProofHeight   *big.Int
	ProofDataRoot [32]byte
	RowProof      BinaryMerkleProof
}, error) {
	return _CelestiaBatchVerifier.Contract.VerifyProof(&_CelestiaBatchVerifier.CallOpts, _blobstream, _rowRoot, _rowProof, _attestationProof)
}

// DAVerifierMetaData contains all meta data concerning the DAVerifier contract.
var DAVerifierMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea26469706673582212204742f5d163f343077f4c29e8db53481bbd85e3da46cf3455df79d7cdaf1585e164736f6c63430008190033",
}

// DAVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use DAVerifierMetaData.ABI instead.
var DAVerifierABI = DAVerifierMetaData.ABI

// DAVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DAVerifierMetaData.Bin instead.
var DAVerifierBin = DAVerifierMetaData.Bin

// DeployDAVerifier deploys a new Ethereum contract, binding an instance of DAVerifier to it.
func DeployDAVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DAVerifier, error) {
	parsed, err := DAVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DAVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DAVerifier{DAVerifierCaller: DAVerifierCaller{contract: contract}, DAVerifierTransactor: DAVerifierTransactor{contract: contract}, DAVerifierFilterer: DAVerifierFilterer{contract: contract}}, nil
}

// DAVerifier is an auto generated Go binding around an Ethereum contract.
type DAVerifier struct {
	DAVerifierCaller     // Read-only binding to the contract
	DAVerifierTransactor // Write-only binding to the contract
	DAVerifierFilterer   // Log filterer for contract events
}

// DAVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAVerifierSession struct {
	Contract     *DAVerifier       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAVerifierCallerSession struct {
	Contract *DAVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DAVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAVerifierTransactorSession struct {
	Contract     *DAVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DAVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAVerifierRaw struct {
	Contract *DAVerifier // Generic contract binding to access the raw methods on
}

// DAVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAVerifierCallerRaw struct {
	Contract *DAVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// DAVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAVerifierTransactorRaw struct {
	Contract *DAVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAVerifier creates a new instance of DAVerifier, bound to a specific deployed contract.
func NewDAVerifier(address common.Address, backend bind.ContractBackend) (*DAVerifier, error) {
	contract, err := bindDAVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAVerifier{DAVerifierCaller: DAVerifierCaller{contract: contract}, DAVerifierTransactor: DAVerifierTransactor{contract: contract}, DAVerifierFilterer: DAVerifierFilterer{contract: contract}}, nil
}

// NewDAVerifierCaller creates a new read-only instance of DAVerifier, bound to a specific deployed contract.
func NewDAVerifierCaller(address common.Address, caller bind.ContractCaller) (*DAVerifierCaller, error) {
	contract, err := bindDAVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAVerifierCaller{contract: contract}, nil
}

// NewDAVerifierTransactor creates a new write-only instance of DAVerifier, bound to a specific deployed contract.
func NewDAVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*DAVerifierTransactor, error) {
	contract, err := bindDAVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAVerifierTransactor{contract: contract}, nil
}

// NewDAVerifierFilterer creates a new log filterer instance of DAVerifier, bound to a specific deployed contract.
func NewDAVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*DAVerifierFilterer, error) {
	contract, err := bindDAVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAVerifierFilterer{contract: contract}, nil
}

// bindDAVerifier binds a generic wrapper to an already deployed contract.
func bindDAVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DAVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAVerifier *DAVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAVerifier.Contract.DAVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAVerifier *DAVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAVerifier.Contract.DAVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAVerifier *DAVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAVerifier.Contract.DAVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAVerifier *DAVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAVerifier *DAVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAVerifier *DAVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAVerifier.Contract.contract.Transact(opts, method, params...)
}

// IBlobstreamXMetaData contains all meta data concerning the IBlobstreamX contract.
var IBlobstreamXMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ContractFrozen\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proofNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"startBlock\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"endBlock\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dataCommitment\",\"type\":\"bytes32\"}],\"name\":\"DataCommitmentStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"frozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_dataCommitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_proofNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tupleRootNonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structDataRootTuple\",\"name\":\"_tuple\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"sideNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numLeaves\",\"type\":\"uint256\"}],\"internalType\":\"structBinaryMerkleProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBlobstreamXABI is the input ABI used to generate the binding from.
// Deprecated: Use IBlobstreamXMetaData.ABI instead.
var IBlobstreamXABI = IBlobstreamXMetaData.ABI

// IBlobstreamX is an auto generated Go binding around an Ethereum contract.
type IBlobstreamX struct {
	IBlobstreamXCaller     // Read-only binding to the contract
	IBlobstreamXTransactor // Write-only binding to the contract
	IBlobstreamXFilterer   // Log filterer for contract events
}

// IBlobstreamXCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlobstreamXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlobstreamXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlobstreamXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlobstreamXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlobstreamXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlobstreamXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlobstreamXSession struct {
	Contract     *IBlobstreamX     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBlobstreamXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlobstreamXCallerSession struct {
	Contract *IBlobstreamXCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IBlobstreamXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlobstreamXTransactorSession struct {
	Contract     *IBlobstreamXTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IBlobstreamXRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlobstreamXRaw struct {
	Contract *IBlobstreamX // Generic contract binding to access the raw methods on
}

// IBlobstreamXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlobstreamXCallerRaw struct {
	Contract *IBlobstreamXCaller // Generic read-only contract binding to access the raw methods on
}

// IBlobstreamXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlobstreamXTransactorRaw struct {
	Contract *IBlobstreamXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlobstreamX creates a new instance of IBlobstreamX, bound to a specific deployed contract.
func NewIBlobstreamX(address common.Address, backend bind.ContractBackend) (*IBlobstreamX, error) {
	contract, err := bindIBlobstreamX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamX{IBlobstreamXCaller: IBlobstreamXCaller{contract: contract}, IBlobstreamXTransactor: IBlobstreamXTransactor{contract: contract}, IBlobstreamXFilterer: IBlobstreamXFilterer{contract: contract}}, nil
}

// NewIBlobstreamXCaller creates a new read-only instance of IBlobstreamX, bound to a specific deployed contract.
func NewIBlobstreamXCaller(address common.Address, caller bind.ContractCaller) (*IBlobstreamXCaller, error) {
	contract, err := bindIBlobstreamX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXCaller{contract: contract}, nil
}

// NewIBlobstreamXTransactor creates a new write-only instance of IBlobstreamX, bound to a specific deployed contract.
func NewIBlobstreamXTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlobstreamXTransactor, error) {
	contract, err := bindIBlobstreamX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXTransactor{contract: contract}, nil
}

// NewIBlobstreamXFilterer creates a new log filterer instance of IBlobstreamX, bound to a specific deployed contract.
func NewIBlobstreamXFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlobstreamXFilterer, error) {
	contract, err := bindIBlobstreamX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXFilterer{contract: contract}, nil
}

// bindIBlobstreamX binds a generic wrapper to an already deployed contract.
func bindIBlobstreamX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBlobstreamXMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlobstreamX *IBlobstreamXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlobstreamX.Contract.IBlobstreamXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlobstreamX *IBlobstreamXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlobstreamX.Contract.IBlobstreamXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlobstreamX *IBlobstreamXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlobstreamX.Contract.IBlobstreamXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlobstreamX *IBlobstreamXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlobstreamX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlobstreamX *IBlobstreamXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlobstreamX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlobstreamX *IBlobstreamXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlobstreamX.Contract.contract.Transact(opts, method, params...)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_IBlobstreamX *IBlobstreamXCaller) Frozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IBlobstreamX.contract.Call(opts, &out, "frozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_IBlobstreamX *IBlobstreamXSession) Frozen() (bool, error) {
	return _IBlobstreamX.Contract.Frozen(&_IBlobstreamX.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0x054f7d9c.
//
// Solidity: function frozen() view returns(bool)
func (_IBlobstreamX *IBlobstreamXCallerSession) Frozen() (bool, error) {
	return _IBlobstreamX.Contract.Frozen(&_IBlobstreamX.CallOpts)
}

// LatestBlock is a free data retrieval call binding the contract method 0x07e2da96.
//
// Solidity: function latestBlock() view returns(uint64)
func (_IBlobstreamX *IBlobstreamXCaller) LatestBlock(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBlobstreamX.contract.Call(opts, &out, "latestBlock")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestBlock is a free data retrieval call binding the contract method 0x07e2da96.
//
// Solidity: function latestBlock() view returns(uint64)
func (_IBlobstreamX *IBlobstreamXSession) LatestBlock() (uint64, error) {
	return _IBlobstreamX.Contract.LatestBlock(&_IBlobstreamX.CallOpts)
}

// LatestBlock is a free data retrieval call binding the contract method 0x07e2da96.
//
// Solidity: function latestBlock() view returns(uint64)
func (_IBlobstreamX *IBlobstreamXCallerSession) LatestBlock() (uint64, error) {
	return _IBlobstreamX.Contract.LatestBlock(&_IBlobstreamX.CallOpts)
}

// StateDataCommitments is a free data retrieval call binding the contract method 0xaeeed33e.
//
// Solidity: function state_dataCommitments(uint256 ) view returns(bytes32)
func (_IBlobstreamX *IBlobstreamXCaller) StateDataCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IBlobstreamX.contract.Call(opts, &out, "state_dataCommitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateDataCommitments is a free data retrieval call binding the contract method 0xaeeed33e.
//
// Solidity: function state_dataCommitments(uint256 ) view returns(bytes32)
func (_IBlobstreamX *IBlobstreamXSession) StateDataCommitments(arg0 *big.Int) ([32]byte, error) {
	return _IBlobstreamX.Contract.StateDataCommitments(&_IBlobstreamX.CallOpts, arg0)
}

// StateDataCommitments is a free data retrieval call binding the contract method 0xaeeed33e.
//
// Solidity: function state_dataCommitments(uint256 ) view returns(bytes32)
func (_IBlobstreamX *IBlobstreamXCallerSession) StateDataCommitments(arg0 *big.Int) ([32]byte, error) {
	return _IBlobstreamX.Contract.StateDataCommitments(&_IBlobstreamX.CallOpts, arg0)
}

// StateProofNonce is a free data retrieval call binding the contract method 0x55ae3f22.
//
// Solidity: function state_proofNonce() view returns(uint256)
func (_IBlobstreamX *IBlobstreamXCaller) StateProofNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBlobstreamX.contract.Call(opts, &out, "state_proofNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateProofNonce is a free data retrieval call binding the contract method 0x55ae3f22.
//
// Solidity: function state_proofNonce() view returns(uint256)
func (_IBlobstreamX *IBlobstreamXSession) StateProofNonce() (*big.Int, error) {
	return _IBlobstreamX.Contract.StateProofNonce(&_IBlobstreamX.CallOpts)
}

// StateProofNonce is a free data retrieval call binding the contract method 0x55ae3f22.
//
// Solidity: function state_proofNonce() view returns(uint256)
func (_IBlobstreamX *IBlobstreamXCallerSession) StateProofNonce() (*big.Int, error) {
	return _IBlobstreamX.Contract.StateProofNonce(&_IBlobstreamX.CallOpts)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_IBlobstreamX *IBlobstreamXCaller) VerifyAttestation(opts *bind.CallOpts, _tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	var out []interface{}
	err := _IBlobstreamX.contract.Call(opts, &out, "verifyAttestation", _tupleRootNonce, _tuple, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_IBlobstreamX *IBlobstreamXSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _IBlobstreamX.Contract.VerifyAttestation(&_IBlobstreamX.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// VerifyAttestation is a free data retrieval call binding the contract method 0x1f3302a9.
//
// Solidity: function verifyAttestation(uint256 _tupleRootNonce, (uint256,bytes32) _tuple, (bytes32[],uint256,uint256) _proof) view returns(bool)
func (_IBlobstreamX *IBlobstreamXCallerSession) VerifyAttestation(_tupleRootNonce *big.Int, _tuple DataRootTuple, _proof BinaryMerkleProof) (bool, error) {
	return _IBlobstreamX.Contract.VerifyAttestation(&_IBlobstreamX.CallOpts, _tupleRootNonce, _tuple, _proof)
}

// IBlobstreamXDataCommitmentStoredIterator is returned from FilterDataCommitmentStored and is used to iterate over the raw logs and unpacked data for DataCommitmentStored events raised by the IBlobstreamX contract.
type IBlobstreamXDataCommitmentStoredIterator struct {
	Event *IBlobstreamXDataCommitmentStored // Event containing the contract specifics and raw log

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
func (it *IBlobstreamXDataCommitmentStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlobstreamXDataCommitmentStored)
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
		it.Event = new(IBlobstreamXDataCommitmentStored)
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
func (it *IBlobstreamXDataCommitmentStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlobstreamXDataCommitmentStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlobstreamXDataCommitmentStored represents a DataCommitmentStored event raised by the IBlobstreamX contract.
type IBlobstreamXDataCommitmentStored struct {
	ProofNonce     *big.Int
	StartBlock     uint64
	EndBlock       uint64
	DataCommitment [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDataCommitmentStored is a free log retrieval operation binding the contract event 0x34dd3689f5bd77a60a3ff2e09483dcab032fa2f1fd7227af3e24bed21beab1cb.
//
// Solidity: event DataCommitmentStored(uint256 proofNonce, uint64 indexed startBlock, uint64 indexed endBlock, bytes32 indexed dataCommitment)
func (_IBlobstreamX *IBlobstreamXFilterer) FilterDataCommitmentStored(opts *bind.FilterOpts, startBlock []uint64, endBlock []uint64, dataCommitment [][32]byte) (*IBlobstreamXDataCommitmentStoredIterator, error) {

	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}
	var dataCommitmentRule []interface{}
	for _, dataCommitmentItem := range dataCommitment {
		dataCommitmentRule = append(dataCommitmentRule, dataCommitmentItem)
	}

	logs, sub, err := _IBlobstreamX.contract.FilterLogs(opts, "DataCommitmentStored", startBlockRule, endBlockRule, dataCommitmentRule)
	if err != nil {
		return nil, err
	}
	return &IBlobstreamXDataCommitmentStoredIterator{contract: _IBlobstreamX.contract, event: "DataCommitmentStored", logs: logs, sub: sub}, nil
}

// WatchDataCommitmentStored is a free log subscription operation binding the contract event 0x34dd3689f5bd77a60a3ff2e09483dcab032fa2f1fd7227af3e24bed21beab1cb.
//
// Solidity: event DataCommitmentStored(uint256 proofNonce, uint64 indexed startBlock, uint64 indexed endBlock, bytes32 indexed dataCommitment)
func (_IBlobstreamX *IBlobstreamXFilterer) WatchDataCommitmentStored(opts *bind.WatchOpts, sink chan<- *IBlobstreamXDataCommitmentStored, startBlock []uint64, endBlock []uint64, dataCommitment [][32]byte) (event.Subscription, error) {

	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}
	var dataCommitmentRule []interface{}
	for _, dataCommitmentItem := range dataCommitment {
		dataCommitmentRule = append(dataCommitmentRule, dataCommitmentItem)
	}

	logs, sub, err := _IBlobstreamX.contract.WatchLogs(opts, "DataCommitmentStored", startBlockRule, endBlockRule, dataCommitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlobstreamXDataCommitmentStored)
				if err := _IBlobstreamX.contract.UnpackLog(event, "DataCommitmentStored", log); err != nil {
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

// ParseDataCommitmentStored is a log parse operation binding the contract event 0x34dd3689f5bd77a60a3ff2e09483dcab032fa2f1fd7227af3e24bed21beab1cb.
//
// Solidity: event DataCommitmentStored(uint256 proofNonce, uint64 indexed startBlock, uint64 indexed endBlock, bytes32 indexed dataCommitment)
func (_IBlobstreamX *IBlobstreamXFilterer) ParseDataCommitmentStored(log types.Log) (*IBlobstreamXDataCommitmentStored, error) {
	event := new(IBlobstreamXDataCommitmentStored)
	if err := _IBlobstreamX.contract.UnpackLog(event, "DataCommitmentStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
