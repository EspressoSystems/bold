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
	Bin: "0x6080806040523461001a57610a7b9081610020823930815050f35b600080fdfe60806040818152600436101561001457600080fd5b600090813560e01c63d5a784cd1461002b57600080fd5b61010090600319828136011261022c576004356001600160a01b0381169490859003610229576023199136830160a0811261022557849061006b89610230565b126102215783519661007c88610261565b6001600160f81b0319602435818116810361021d5789526044359863ffffffff1990818b168b036102195760209a8b82015283528660631936011261021d578651916100c783610261565b6064359081168103610219578252608435908116810361021d57818a01528189015260a435858201526001600160401b039660c43588811161021d576101119036906004016102d1565b9460e435908982116102195760808236039586011261021957879081519561013887610230565b83600401358752011261021d578651986101518a610261565b60248201358a5260448201358b8b01528a8501998a526064820135908111610219579161018a6101ac95949260048995369201016102d1565b8885015285888c8282519161019e83610230565b606083528201520152610377565b509451868151910151908451961515875287870152838601526080606086015260e085018683519560606080890152865180935287019501915b81811061020557505050839481015160a0850152015160c08301520390f35b8251865294870194918701916001016101e6565b8580fd5b8480fd5b5080fd5b8280fd5b80fd5b8380fd5b606081019081106001600160401b0382111761024b57604052565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b0382111761024b57604052565b608081019081106001600160401b0382111761024b57604052565b601f909101601f19168101906001600160401b0382119082101761024b57604052565b6001600160401b03811161024b5760051b60200190565b91909160608184031261037257604051906102eb82610230565b909283919081356001600160401b0381116103725782019080601f830112156103725781359061031a826102ba565b916103286040519384610297565b808352602093848085019260051b820101928311610372578401905b8282106103635750505083528082013590830152604090810135910152565b81358152908401908401610344565b600080fd5b9291909180519060208092818301966040885194015183604051958694631f3302a960e01b8652600486015280516024860152015160448401526080606484015260e483019080519160606084860152825180915285610104860193019060005b818110610479575050508481015160a48501526040015160c48401528290039082906001600160a01b03165afa90811561046d57600091610437575b501561042b576104279351015191610492565b9091565b50505050600090600490565b8281813d8311610466575b61044c8183610297565b810103126102215751908115158203610229575038610414565b503d610442565b6040513d6000823e3d90fd5b82518552899789975094850194909201916001016103d8565b906104e2926104a183516104f8565b60406104b060208601516104f8565b940151906040519462ffffff19809216602087015216603d850152605a840152605a83526104dd8361027c565b610581565b50156104f057600190600090565b600090600390565b80516020918201516040516001600160f81b031990921692820192835263ffffffff19166021820152601d815261052e81610261565b51905162ffffff19918282169190601d811061054c575b5050905090565b83919250601d0360031b1b1616803880610545565b6006111561056b57565b634e487b7160e01b600052602160045260246000fd5b6040820180519394936001811161067a575082515161066d575b602083019081518151111561065f5760206105ce966105e86040516105dc8160009b8c95868884015260218301906109dc565b03601f198101835282610297565b604051918280926109dc565b039060025afa1561065457855193519182511561063d579061060f939491519051906107ce565b600681101561062957806106235750149190565b92915050565b634e487b7160e01b85526021600452602485fd5b505160011415905061064e57149190565b50508190565b6040513d87823e3d90fd5b505050509050600090600290565b5050509050600090600190565b61068b8451519160208601516106ed565b1461059b575050509050600090600190565b906101009182039182116106ad57565b634e487b7160e01b600052601160045260246000fd5b6000198101919082116106ad57565b919082039182116106ad57565b60010190816001116106ad57565b906001908181111561079c5760005b8183821b1061077b576101009081039081116106ad5761071b9061069d565b9282610726856106c3565b1b90610731826106c3565b811161073e575050505090565b92935090918382036107505750505090565b61076f93509061076381610769936106d2565b926106d2565b906106ed565b610778906106df565b90565b828101809111156106fc57634e487b7160e01b600052601160045260246000fd5b505050600090565b80518210156107b85760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b93929381156108b557600182146108a05784511561089557826107f0836108c0565b6108036107fd88516106c3565b88610913565b928181106108605781610763610823969361081d936106d2565b906107ce565b909161082e82610561565b8161085957505061084d836108476108539495516106c3565b906107a4565b51610a07565b90600090565b9350919050565b61086a94506107ce565b909161087582610561565b8161085957505061088e836108476108539495516106c3565b5190610a07565b505090915090600590565b50509091516108af5790600090565b90600490565b505090915090600390565b600180821061037257600082805b6108f1575060001981019081116106ad5781901b9182146108ed575090565b1c90565b906108fb90610904565b90821c806108ce565b60001981146106ad5760010190565b91908251811161097f57610926816102ba565b906109346040519283610297565b808252601f19610943826102ba565b0136602084013760005b81811061095b575090925050565b8061096961097a92876107a4565b5161097482866107a4565b52610904565b61094d565b60405162461bcd60e51b815260206004820152602f60248201527f496e76616c69642072616e67653a205f626567696e206f72205f656e6420617260448201526e65206f7574206f6620626f756e647360881b6064820152608490fd5b9081519160005b8381106109f4575050016000815290565b80602080928401015181850152016109e3565b610a3460009160209360405191600160f81b8684015260218301526041820152604181526105dc8161027c565b039060025afa1561046d576000519056fea264697066735822122089722efb8e42cf4bfe0b434f7c0dae7ab768f20d87c5a084c2ed21c74d2ce99364736f6c63430008130033",
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
	Bin: "0x60808060405234601757603a9081601d823930815050f35b600080fdfe600080fdfea2646970667358221220113c7add5ebb3a1befb6eb58a0bcf8c001c74b5c3f4e1fde944cfd47f972c50564736f6c63430008130033",
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
