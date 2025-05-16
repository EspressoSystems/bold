// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chaingen

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

// CacheManagerEntry is an auto generated low-level Go binding around an user-defined struct.
type CacheManagerEntry struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}

// CacheManagerMetaData contains all meta data concerning the CacheManager contract.
var CacheManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"}],\"name\":\"AlreadyCached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queueSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cacheSize\",\"type\":\"uint256\"}],\"name\":\"AsmTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"name\":\"BidTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidsArePaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"MakeSpaceTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotChainOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"DeleteBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"InsertBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"SetCacheSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"decay\",\"type\":\"uint64\"}],\"name\":\"SetDecayRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cacheSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entries\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"evictAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"evictPrograms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"}],\"internalType\":\"structCacheManager.Entry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"}],\"name\":\"getMinBid\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"}],\"name\":\"getMinBid\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"getMinBid\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"getSmallestEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"}],\"internalType\":\"structCacheManager.Entry[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initCacheSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initDecay\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"makeSpace\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"space\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSize\",\"type\":\"uint64\"}],\"name\":\"setCacheSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDecay\",\"type\":\"uint64\"}],\"name\":\"setDecayRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweepFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0806040523460225730608052611bdf9081610028823960805181610be60152f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806317be85c31461015757806320f2f345146101525780632dd4f5661461014d57806332052a9b146101485780633f4ba83a1461014357806354fac9191461013e5780635c32e943146101395780635c975abb14610134578063674a64e01461012f578063a8d6fe041461012a578063b187bd2614610125578063b30906d414610120578063bae6c2ad1461011b578063c1c013c414610116578063c565a20814610111578063c77ed13e1461010c578063cadb43e214610107578063d29b303e14610102578063e4940157146100fd5763e9c1bc0f146100f857600080fd5b610a97565b6109e8565b6109ca565b610905565b61088b565b61086a565b6107c5565b61079a565b610751565b6106df565b61063e565b610615565b6105a4565b610533565b610506565b610492565b610450565b61038f565b61029e565b6101de565b600091031261016757565b600080fd5b6001600160401b031690565b60208082019080835283518092528060408094019401926000905b8382106101a257505050505090565b845180518752808401516001600160401b0316878501528101516001600160c01b03168682015260609095019493820193600190910190610193565b346101675760008060031936011261026f5760028054906101fe82610b37565b9261020c6040519485610b05565b8284526002815260207f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace8186015b858410610253576040518061024f8982610178565b0390f35b848360019261026185610b4e565b81520192019301929061023a565b80fd5b600435906001600160401b038216820361016757565b602435906001600160401b038216820361016757565b34610167576040366003190112610167576102b7610272565b6103036102c2610288565b600054926102e760ff8560081c161580958196610381575b8115610361575b50610b7f565b836102fa600160ff196000541617600055565b61034857610be2565b61030957005b61031961ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b61035c61010061ff00196000541617600055565b610be2565b303b15915081610373575b50386102e1565b6001915060ff16143861036c565b600160ff82161091506102da565b34610167576020366003190112610167576103a8610272565b6040516304ddefed60e31b8152602081806103c63360048301610cd9565b0381606b5afa90811561043a5760009161040b575b50156103ec576103ea90610cf8565b005b604051639531eff160e01b8152806104073360048301610cd9565b0390fd5b61042d915060203d602011610433575b6104258183610b05565b810190610cc1565b386103db565b503d61041b565b610cec565b6001600160a01b0381160361016757565b3461016757602036600319011261016757602061048061047b6004356104758161043f565b3f61145f565b610fe9565b6040516001600160c01b039091168152f35b346101675760008060031936011261026f576040516304ddefed60e31b8152602081806104c23360048301610cd9565b0381606b5afa90811561043a5782916104e7575b50156103ec576104e4610d44565b80f35b610500915060203d602011610433576104258183610b05565b386104d6565b346101675760003660031901126101675760035460405160809190911c6001600160401b03168152602090f35b346101675760008060031936011261026f576040516304ddefed60e31b8152602081806105633360048301610cd9565b0381606b5afa90811561043a578291610585575b50156103ec576104e4610e10565b61059e915060203d602011610433576104258183610b05565b38610577565b346101675760008060031936011261026f576040516304ddefed60e31b8152602081806105d43360048301610cd9565b0381606b5afa90811561043a5782916105f6575b50156103ec576104e4610ea6565b61060f915060203d602011610433576104258183610b05565b386105e8565b34610167576000366003190112610167576003546040516001600160401b039091168152602090f35b346101675760008060031936011261026f57604051632d9125e960e01b8152602081600481606b5afa90811561043a5782918291829161069b575b5081809147905af1610689610ee0565b9015610693575080f35b602081519101fd5b9150506020813d6020116106d7575b816106b760209383610b05565b810103126106d45781808092516106cd8161043f565b9150610679565b50fd5b3d91506106aa565b3461016757600036600319011261016757602060ff60035460c01c166040519015158152f35b60209060031901126101675760043590565b634e487b7160e01b600052603260045260246000fd5b60025481101561074c57600260005260206000209060011b0190600090565b610717565b346101675761075f36610705565b6002548110156101675761077460609161072d565b506001815491015460405191825260018060401b038116602083015260401c6040820152f35b34610167576000366003190112610167576003546040805191901c6001600160401b03168152602090f35b6020366003190112610167576107d9610272565b60ff60035460c01c16610858576001600160401b03908082166250000080821161083a5761024f6108208561080d86611354565b505060035490808260401c169116610f1f565b6040516001600160401b0390911681529081906020820190565b604492506040519163e6b801f360e01b835260048301526024820152fd5b6040516323d5725b60e21b8152600490fd5b3461016757602036600319011261016757602061048061047b60043561145f565b34610167576020366003190112610167576108a4610272565b6040516304ddefed60e31b8152602081806108c23360048301610cd9565b0381606b5afa90811561043a576000916108e6575b50156103ec576103ea90610f38565b6108ff915060203d602011610433576104258183610b05565b386108d7565b346101675761091336610705565b6040906040516304ddefed60e31b8152336004820152602081602481606b5afa90811561043a576000916109ab575b5015610993575b60015415158061098a575b156103ea57610977610964611a05565b6001600160401b0381169150841c6115e2565b600019810190811115610949575b610d78565b50801515610954565b604051639531eff160e01b8152336004820152602490fd5b6109c4915060203d602011610433576104258183610b05565b38610942565b3461016757602036600319011261016757602061048061047b610272565b602036600319011261016757600435610a008161043f565b60ff60035460c01c1661085857803f9060405163a72f179b60e01b815282600482015260208160248160725afa90811561043a57600091610a78575b50610a5f5781610a4e6103ea9361145f565b90610a5882611354565b93906117a6565b60405163c7e2d8e560e01b815260048101839052602490fd5b610a91915060203d602011610433576104258183610b05565b38610a3c565b346101675761024f610ab0610aab36610705565b6111bc565b60405191829182610178565b634e487b7160e01b600052604160045260246000fd5b606081019081106001600160401b03821117610aed57604052565b610abc565b6001600160401b038111610aed57604052565b601f909101601f19168101906001600160401b03821190821017610aed57604052565b60405190610b3582610ad2565b565b6001600160401b038111610aed5760051b60200190565b90604051610b5b81610ad2565b825481526001909201546001600160401b0381166020840152604090811c90830152565b15610b8657565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b90307f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614610c3c57600380546001600160401b0319166001600160401b0390931692909217909155610b3590610c96565b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b60038054600160801b600160c01b03191660809290921b600160801b600160c01b0316919091179055565b90816020910312610167575180151581036101675790565b6001600160a01b03909116815260200190565b6040513d6000823e3d90fd5b600380546001600160401b0319166001600160401b0390921691821790556040519081527fca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c183990602090a1565b6003805460ff60c01b191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33600080a1565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052600060045260246000fd5b90610db55760018160008093550155565b610d8e565b6002805460008060025581610dce57505050565b60016001600160ff1b038316830361098557600260005260206000209260011b8301925b838110610e00575050505050565b8083869255838382015501610df2565b604080516304ddefed60e31b8152336004820152602081602481606b5afa90811561043a57600091610e87575b5015610993576000195b600154151580610e7e575b15610e7457610e62610964611a05565b600019810190811115610e4757610d78565b5050610b35610dba565b50801515610e52565b610ea0915060203d602011610433576104258183610b05565b38610e3d565b6003805460ff60c01b1916600160c01b1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625600080a1565b3d15610f1a573d906001600160401b038211610aed5760405191610f0e601f8201601f191660200184610b05565b82523d6000602084013e565b606090565b6001600160401b03918216908216039190821161098557565b60207fd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c91610f6581610c96565b6040516001600160401b039091168152a1565b60001981019190821161098557565b9190820391821161098557565b6001600160401b03918216908216019190821161098557565b90611000820180921161098557565b805182101561074c5760209160051b010190565b6001600160c01b03918216908216039190821161098557565b6003546000916001600160401b03808316918181168381116111245761102d9261102892909161100090811161111c5750905b60409560401c16610f94565b61016c565b90808211156111135761103f91610f87565b9261105d610aab61105761105287610fad565b610f78565b600c1c90565b906000945b825186101561110957602061108d6110288261107e8a88610fbc565b5101516001600160401b031690565b8211156110b9576001916110ab6110286110b19361107e8b89610fbc565b90610f87565b950194611062565b5050936110d993506110ca91610fbc565b5101516001600160c01b031690565b6110e16116f9565b906001600160c01b038082168311611101576110fe921690610fd0565b90565b505050600090565b50935050506110d9565b50505050600090565b90509061101c565b60405163bcc27c3760e01b81526001600160401b0383811660048301526000602483015285166044820152606490fd5b9061115e82610b37565b60409061116e6040519182610b05565b838152809361117f601f1991610b37565b019160009160005b848110611195575050505050565b60209083516111a381610ad2565b8581528286818301528686830152828501015201611187565b90600180549280841061134d575b604080519483600052602090602060002092602088019482811090831802821860051b8801916020830191611206865484600091815260200152565b60408401958784141595865b6112aa575b5050505050505050601f198482030160051c8452604052806112398451611154565b936000915b61124757505050565b80518210156112a557828261128861128261127b611266859787610fbc565b51604081901c916001600160401b0390911690565b905061072d565b50610b4e565b6112928289610fbc565b5261129d8188610fbc565b50019161123e565b505050565b90919293949596978551815281019785891461134857898151811b8181019087821015611321576112e48187019284840154908d8c6119a8565b600280910191888303611305575b50505050875b9695949392919097611212565b01549299919261131691838a611961565b0196893880806112f2565b505050966000190196876113438160061b89018685820151910151838a6119a8565b6112f8565b611217565b50826111ca565b9061135d6116f9565b340190813411610985576001600160c01b03918281116114475791809392169261138860025461016c565b90600060039261139c61102860035461016c565b915b845483906113c290611028908a9060401c6001600160401b0316610f94565b610f94565b111561140c5750506113c26110286113bd966113f36113df6114ec565b604081901c916001600160401b0390911690565b9890966114008a896115e2565b9692505096915061139e565b93509450508116841061141c5750565b604051631be6e1c960e31b81526001600160c01b038581166004830152919091166024820152604490fd5b60249060405190631edd0da560e31b82526004820152fd5b60405190634089267f60e01b8252600482015260208160248160715afa90811561043a576000916114a5575b5063ffffffff166110008082106114a0575090565b905090565b6020813d6020116114e4575b816114be60209383610b05565b810103126114e057519063ffffffff8216820361026f575063ffffffff61148b565b5080fd5b3d91506114b1565b6001805460008281529182918291600080516020611b8a83398151915291600019906115a0565b84905b80831061156357505b61153a575b01611530575b50505090565b015538808061152a565b926000198101821c9081840154908187101561155a57840155928361151f565b94915050611524565b9180915084019485549184808201970154928381108589101115611594575b50508401558184811b01908490611516565b90965091503880611582565b955093508480156115d557600019908101808355958301015493849086156115ce5750508154948190611513565b9095611513565b63a6ca772e84526004601cfd5b906115ef6112828261072d565b9081519160723b156101675760405163ce97201360e01b815260048101939093526000836024818360725af1801561043a57610b35946116da947f65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7926116e0575b506116b060208401936116a861167d611669875161016c565b60035460401c6001600160401b0316610f1f565b60038054600160401b600160801b03191660409290921b600160401b600160801b0316919091179055565b51935161016c565b604080516001600160c01b039390931683526001600160401b0391909116602083015290a261072d565b90610da4565b806116ed6116f392610af2565b8061015c565b38611650565b60018060401b0360035460801c168042029042820414421517156109855790565b9190610db55780518255602081015160409182015190911b6001600160401b0319166001600160401b0390911617600190910155565b600254600160401b811015610aed57600181018060025581101561074c576001906002600052602060002090821b019180518355818060401b036020820151169060408380821b031991015160401b1617910155565b9193926003549060018060401b039182808260401c1691816117c88585610f94565b9116918291161161191f5750506117dd610b28565b8681526001600160401b0382166020820152946001600160c01b038516604087015260723b156101675760405163739d64f960e11b815295600087806118268860048301610cd9565b03818360725af193841561043a577fb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec976118f09561190c575b50611882604088901b6001600160401b0319166001600160401b03851617611af2565b6003546118a09061167d90869060401c6001600160401b0316610f94565b600254908316036118f5576118b59150611750565b604080516001600160a01b0390941684526001600160c01b0390941660208401526001600160401b0316928201929092529081906060820190565b0390a2565b6119016119079261072d565b9061171a565b6118b5565b806116ed61191992610af2565b3861185f565b60405163bcc27c3760e01b81526001600160401b038481166004830152928316602482015291166044820152606490fd5b9092919260061b0190815260200152565b905b6000198101600581901b603f1916830180519192908287108202156119995760200151611991929185611950565b60011c611963565b50915050610b35939291611950565b6001949392916000919086905b8082106119c9575050610b35949550611961565b90926119fa90600694838a8201108101861b85015181871b860151110180951b840190815191602001519085611950565b82871b8701906119b5565b600090600090818060018054908083528060208420928460001992611aab575b5084905b808310611a6e57505b611a45575b01611a4157505050565b0155565b926000198101821c90818401549081871015611a65578401559283611a32565b94915050611a37565b9180915084019485549184808201970154928381108589101115611a9f575b50508401558184811b01908490611a29565b90965091503880611a8d565b975080919550968115611ae5575060001990810180835596830101549384908715611add575050815495819038611a25565b909638611a25565b63a6ca772e90526004601cfd5b600180546000829052808201825590600080516020611b8a833981519152908183805b808310611b4c57505b611b2c5701611a4157505050565b926000198101821c90818401549081871015611a65578401559283611b1e565b9180915084019485549184808201970154928381108589101115611b7d575b50508401558184811b01908490611b15565b90965091503880611b6b56feb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6a2646970667358221220427e5f5e40c901415f48893526358789ae49c2580fa06cba28a0f8c1b1d5044064736f6c63430008190033",
}

// CacheManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CacheManagerMetaData.ABI instead.
var CacheManagerABI = CacheManagerMetaData.ABI

// CacheManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CacheManagerMetaData.Bin instead.
var CacheManagerBin = CacheManagerMetaData.Bin

// DeployCacheManager deploys a new Ethereum contract, binding an instance of CacheManager to it.
func DeployCacheManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CacheManager, error) {
	parsed, err := CacheManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CacheManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CacheManager{CacheManagerCaller: CacheManagerCaller{contract: contract}, CacheManagerTransactor: CacheManagerTransactor{contract: contract}, CacheManagerFilterer: CacheManagerFilterer{contract: contract}}, nil
}

// CacheManager is an auto generated Go binding around an Ethereum contract.
type CacheManager struct {
	CacheManagerCaller     // Read-only binding to the contract
	CacheManagerTransactor // Write-only binding to the contract
	CacheManagerFilterer   // Log filterer for contract events
}

// CacheManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CacheManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CacheManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CacheManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CacheManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CacheManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CacheManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CacheManagerSession struct {
	Contract     *CacheManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CacheManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CacheManagerCallerSession struct {
	Contract *CacheManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CacheManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CacheManagerTransactorSession struct {
	Contract     *CacheManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CacheManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CacheManagerRaw struct {
	Contract *CacheManager // Generic contract binding to access the raw methods on
}

// CacheManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CacheManagerCallerRaw struct {
	Contract *CacheManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CacheManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CacheManagerTransactorRaw struct {
	Contract *CacheManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCacheManager creates a new instance of CacheManager, bound to a specific deployed contract.
func NewCacheManager(address common.Address, backend bind.ContractBackend) (*CacheManager, error) {
	contract, err := bindCacheManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CacheManager{CacheManagerCaller: CacheManagerCaller{contract: contract}, CacheManagerTransactor: CacheManagerTransactor{contract: contract}, CacheManagerFilterer: CacheManagerFilterer{contract: contract}}, nil
}

// NewCacheManagerCaller creates a new read-only instance of CacheManager, bound to a specific deployed contract.
func NewCacheManagerCaller(address common.Address, caller bind.ContractCaller) (*CacheManagerCaller, error) {
	contract, err := bindCacheManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CacheManagerCaller{contract: contract}, nil
}

// NewCacheManagerTransactor creates a new write-only instance of CacheManager, bound to a specific deployed contract.
func NewCacheManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CacheManagerTransactor, error) {
	contract, err := bindCacheManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CacheManagerTransactor{contract: contract}, nil
}

// NewCacheManagerFilterer creates a new log filterer instance of CacheManager, bound to a specific deployed contract.
func NewCacheManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CacheManagerFilterer, error) {
	contract, err := bindCacheManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CacheManagerFilterer{contract: contract}, nil
}

// bindCacheManager binds a generic wrapper to an already deployed contract.
func bindCacheManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CacheManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CacheManager *CacheManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CacheManager.Contract.CacheManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CacheManager *CacheManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.Contract.CacheManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CacheManager *CacheManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CacheManager.Contract.CacheManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CacheManager *CacheManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CacheManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CacheManager *CacheManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CacheManager *CacheManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CacheManager.Contract.contract.Transact(opts, method, params...)
}

// CacheSize is a free data retrieval call binding the contract method 0x674a64e0.
//
// Solidity: function cacheSize() view returns(uint64)
func (_CacheManager *CacheManagerCaller) CacheSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "cacheSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CacheSize is a free data retrieval call binding the contract method 0x674a64e0.
//
// Solidity: function cacheSize() view returns(uint64)
func (_CacheManager *CacheManagerSession) CacheSize() (uint64, error) {
	return _CacheManager.Contract.CacheSize(&_CacheManager.CallOpts)
}

// CacheSize is a free data retrieval call binding the contract method 0x674a64e0.
//
// Solidity: function cacheSize() view returns(uint64)
func (_CacheManager *CacheManagerCallerSession) CacheSize() (uint64, error) {
	return _CacheManager.Contract.CacheSize(&_CacheManager.CallOpts)
}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint64)
func (_CacheManager *CacheManagerCaller) Decay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "decay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint64)
func (_CacheManager *CacheManagerSession) Decay() (uint64, error) {
	return _CacheManager.Contract.Decay(&_CacheManager.CallOpts)
}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint64)
func (_CacheManager *CacheManagerCallerSession) Decay() (uint64, error) {
	return _CacheManager.Contract.Decay(&_CacheManager.CallOpts)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size, uint192 bid)
func (_CacheManager *CacheManagerCaller) Entries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "entries", arg0)

	outstruct := new(struct {
		Code [32]byte
		Size uint64
		Bid  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Code = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Size = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Bid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size, uint192 bid)
func (_CacheManager *CacheManagerSession) Entries(arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}, error) {
	return _CacheManager.Contract.Entries(&_CacheManager.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size, uint192 bid)
func (_CacheManager *CacheManagerCallerSession) Entries(arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}, error) {
	return _CacheManager.Contract.Entries(&_CacheManager.CallOpts, arg0)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes32,uint64,uint192)[])
func (_CacheManager *CacheManagerCaller) GetEntries(opts *bind.CallOpts) ([]CacheManagerEntry, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getEntries")

	if err != nil {
		return *new([]CacheManagerEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]CacheManagerEntry)).(*[]CacheManagerEntry)

	return out0, err

}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes32,uint64,uint192)[])
func (_CacheManager *CacheManagerSession) GetEntries() ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetEntries(&_CacheManager.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes32,uint64,uint192)[])
func (_CacheManager *CacheManagerCallerSession) GetEntries() ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetEntries(&_CacheManager.CallOpts)
}

// GetMinBid32052a9b is a free data retrieval call binding the contract method 0x32052a9b.
//
// Solidity: function getMinBid(address program) view returns(uint192 min)
func (_CacheManager *CacheManagerCaller) GetMinBid32052a9b(opts *bind.CallOpts, program common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getMinBid", program)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBid32052a9b is a free data retrieval call binding the contract method 0x32052a9b.
//
// Solidity: function getMinBid(address program) view returns(uint192 min)
func (_CacheManager *CacheManagerSession) GetMinBid32052a9b(program common.Address) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBid32052a9b(&_CacheManager.CallOpts, program)
}

// GetMinBid32052a9b is a free data retrieval call binding the contract method 0x32052a9b.
//
// Solidity: function getMinBid(address program) view returns(uint192 min)
func (_CacheManager *CacheManagerCallerSession) GetMinBid32052a9b(program common.Address) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBid32052a9b(&_CacheManager.CallOpts, program)
}

// GetMinBidc565a208 is a free data retrieval call binding the contract method 0xc565a208.
//
// Solidity: function getMinBid(bytes32 codehash) view returns(uint192 min)
func (_CacheManager *CacheManagerCaller) GetMinBidc565a208(opts *bind.CallOpts, codehash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getMinBid0", codehash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBidc565a208 is a free data retrieval call binding the contract method 0xc565a208.
//
// Solidity: function getMinBid(bytes32 codehash) view returns(uint192 min)
func (_CacheManager *CacheManagerSession) GetMinBidc565a208(codehash [32]byte) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidc565a208(&_CacheManager.CallOpts, codehash)
}

// GetMinBidc565a208 is a free data retrieval call binding the contract method 0xc565a208.
//
// Solidity: function getMinBid(bytes32 codehash) view returns(uint192 min)
func (_CacheManager *CacheManagerCallerSession) GetMinBidc565a208(codehash [32]byte) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidc565a208(&_CacheManager.CallOpts, codehash)
}

// GetMinBidd29b303e is a free data retrieval call binding the contract method 0xd29b303e.
//
// Solidity: function getMinBid(uint64 size) view returns(uint192 min)
func (_CacheManager *CacheManagerCaller) GetMinBidd29b303e(opts *bind.CallOpts, size uint64) (*big.Int, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getMinBid1", size)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBidd29b303e is a free data retrieval call binding the contract method 0xd29b303e.
//
// Solidity: function getMinBid(uint64 size) view returns(uint192 min)
func (_CacheManager *CacheManagerSession) GetMinBidd29b303e(size uint64) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidd29b303e(&_CacheManager.CallOpts, size)
}

// GetMinBidd29b303e is a free data retrieval call binding the contract method 0xd29b303e.
//
// Solidity: function getMinBid(uint64 size) view returns(uint192 min)
func (_CacheManager *CacheManagerCallerSession) GetMinBidd29b303e(size uint64) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidd29b303e(&_CacheManager.CallOpts, size)
}

// GetSmallestEntries is a free data retrieval call binding the contract method 0xe9c1bc0f.
//
// Solidity: function getSmallestEntries(uint256 k) view returns((bytes32,uint64,uint192)[] result)
func (_CacheManager *CacheManagerCaller) GetSmallestEntries(opts *bind.CallOpts, k *big.Int) ([]CacheManagerEntry, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getSmallestEntries", k)

	if err != nil {
		return *new([]CacheManagerEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]CacheManagerEntry)).(*[]CacheManagerEntry)

	return out0, err

}

// GetSmallestEntries is a free data retrieval call binding the contract method 0xe9c1bc0f.
//
// Solidity: function getSmallestEntries(uint256 k) view returns((bytes32,uint64,uint192)[] result)
func (_CacheManager *CacheManagerSession) GetSmallestEntries(k *big.Int) ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetSmallestEntries(&_CacheManager.CallOpts, k)
}

// GetSmallestEntries is a free data retrieval call binding the contract method 0xe9c1bc0f.
//
// Solidity: function getSmallestEntries(uint256 k) view returns((bytes32,uint64,uint192)[] result)
func (_CacheManager *CacheManagerCallerSession) GetSmallestEntries(k *big.Int) ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetSmallestEntries(&_CacheManager.CallOpts, k)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CacheManager *CacheManagerCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CacheManager *CacheManagerSession) IsPaused() (bool, error) {
	return _CacheManager.Contract.IsPaused(&_CacheManager.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CacheManager *CacheManagerCallerSession) IsPaused() (bool, error) {
	return _CacheManager.Contract.IsPaused(&_CacheManager.CallOpts)
}

// QueueSize is a free data retrieval call binding the contract method 0xbae6c2ad.
//
// Solidity: function queueSize() view returns(uint64)
func (_CacheManager *CacheManagerCaller) QueueSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "queueSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// QueueSize is a free data retrieval call binding the contract method 0xbae6c2ad.
//
// Solidity: function queueSize() view returns(uint64)
func (_CacheManager *CacheManagerSession) QueueSize() (uint64, error) {
	return _CacheManager.Contract.QueueSize(&_CacheManager.CallOpts)
}

// QueueSize is a free data retrieval call binding the contract method 0xbae6c2ad.
//
// Solidity: function queueSize() view returns(uint64)
func (_CacheManager *CacheManagerCallerSession) QueueSize() (uint64, error) {
	return _CacheManager.Contract.QueueSize(&_CacheManager.CallOpts)
}

// EvictAll is a paid mutator transaction binding the contract method 0x5c32e943.
//
// Solidity: function evictAll() returns()
func (_CacheManager *CacheManagerTransactor) EvictAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "evictAll")
}

// EvictAll is a paid mutator transaction binding the contract method 0x5c32e943.
//
// Solidity: function evictAll() returns()
func (_CacheManager *CacheManagerSession) EvictAll() (*types.Transaction, error) {
	return _CacheManager.Contract.EvictAll(&_CacheManager.TransactOpts)
}

// EvictAll is a paid mutator transaction binding the contract method 0x5c32e943.
//
// Solidity: function evictAll() returns()
func (_CacheManager *CacheManagerTransactorSession) EvictAll() (*types.Transaction, error) {
	return _CacheManager.Contract.EvictAll(&_CacheManager.TransactOpts)
}

// EvictPrograms is a paid mutator transaction binding the contract method 0xcadb43e2.
//
// Solidity: function evictPrograms(uint256 count) returns()
func (_CacheManager *CacheManagerTransactor) EvictPrograms(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "evictPrograms", count)
}

// EvictPrograms is a paid mutator transaction binding the contract method 0xcadb43e2.
//
// Solidity: function evictPrograms(uint256 count) returns()
func (_CacheManager *CacheManagerSession) EvictPrograms(count *big.Int) (*types.Transaction, error) {
	return _CacheManager.Contract.EvictPrograms(&_CacheManager.TransactOpts, count)
}

// EvictPrograms is a paid mutator transaction binding the contract method 0xcadb43e2.
//
// Solidity: function evictPrograms(uint256 count) returns()
func (_CacheManager *CacheManagerTransactorSession) EvictPrograms(count *big.Int) (*types.Transaction, error) {
	return _CacheManager.Contract.EvictPrograms(&_CacheManager.TransactOpts, count)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 initCacheSize, uint64 initDecay) returns()
func (_CacheManager *CacheManagerTransactor) Initialize(opts *bind.TransactOpts, initCacheSize uint64, initDecay uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "initialize", initCacheSize, initDecay)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 initCacheSize, uint64 initDecay) returns()
func (_CacheManager *CacheManagerSession) Initialize(initCacheSize uint64, initDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.Initialize(&_CacheManager.TransactOpts, initCacheSize, initDecay)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 initCacheSize, uint64 initDecay) returns()
func (_CacheManager *CacheManagerTransactorSession) Initialize(initCacheSize uint64, initDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.Initialize(&_CacheManager.TransactOpts, initCacheSize, initDecay)
}

// MakeSpace is a paid mutator transaction binding the contract method 0xc1c013c4.
//
// Solidity: function makeSpace(uint64 size) payable returns(uint64 space)
func (_CacheManager *CacheManagerTransactor) MakeSpace(opts *bind.TransactOpts, size uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "makeSpace", size)
}

// MakeSpace is a paid mutator transaction binding the contract method 0xc1c013c4.
//
// Solidity: function makeSpace(uint64 size) payable returns(uint64 space)
func (_CacheManager *CacheManagerSession) MakeSpace(size uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.MakeSpace(&_CacheManager.TransactOpts, size)
}

// MakeSpace is a paid mutator transaction binding the contract method 0xc1c013c4.
//
// Solidity: function makeSpace(uint64 size) payable returns(uint64 space)
func (_CacheManager *CacheManagerTransactorSession) MakeSpace(size uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.MakeSpace(&_CacheManager.TransactOpts, size)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns()
func (_CacheManager *CacheManagerTransactor) Paused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "paused")
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns()
func (_CacheManager *CacheManagerSession) Paused() (*types.Transaction, error) {
	return _CacheManager.Contract.Paused(&_CacheManager.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns()
func (_CacheManager *CacheManagerTransactorSession) Paused() (*types.Transaction, error) {
	return _CacheManager.Contract.Paused(&_CacheManager.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xe4940157.
//
// Solidity: function placeBid(address program) payable returns()
func (_CacheManager *CacheManagerTransactor) PlaceBid(opts *bind.TransactOpts, program common.Address) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "placeBid", program)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xe4940157.
//
// Solidity: function placeBid(address program) payable returns()
func (_CacheManager *CacheManagerSession) PlaceBid(program common.Address) (*types.Transaction, error) {
	return _CacheManager.Contract.PlaceBid(&_CacheManager.TransactOpts, program)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xe4940157.
//
// Solidity: function placeBid(address program) payable returns()
func (_CacheManager *CacheManagerTransactorSession) PlaceBid(program common.Address) (*types.Transaction, error) {
	return _CacheManager.Contract.PlaceBid(&_CacheManager.TransactOpts, program)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x2dd4f566.
//
// Solidity: function setCacheSize(uint64 newSize) returns()
func (_CacheManager *CacheManagerTransactor) SetCacheSize(opts *bind.TransactOpts, newSize uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "setCacheSize", newSize)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x2dd4f566.
//
// Solidity: function setCacheSize(uint64 newSize) returns()
func (_CacheManager *CacheManagerSession) SetCacheSize(newSize uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetCacheSize(&_CacheManager.TransactOpts, newSize)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x2dd4f566.
//
// Solidity: function setCacheSize(uint64 newSize) returns()
func (_CacheManager *CacheManagerTransactorSession) SetCacheSize(newSize uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetCacheSize(&_CacheManager.TransactOpts, newSize)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xc77ed13e.
//
// Solidity: function setDecayRate(uint64 newDecay) returns()
func (_CacheManager *CacheManagerTransactor) SetDecayRate(opts *bind.TransactOpts, newDecay uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "setDecayRate", newDecay)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xc77ed13e.
//
// Solidity: function setDecayRate(uint64 newDecay) returns()
func (_CacheManager *CacheManagerSession) SetDecayRate(newDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetDecayRate(&_CacheManager.TransactOpts, newDecay)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xc77ed13e.
//
// Solidity: function setDecayRate(uint64 newDecay) returns()
func (_CacheManager *CacheManagerTransactorSession) SetDecayRate(newDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetDecayRate(&_CacheManager.TransactOpts, newDecay)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns()
func (_CacheManager *CacheManagerTransactor) SweepFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "sweepFunds")
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns()
func (_CacheManager *CacheManagerSession) SweepFunds() (*types.Transaction, error) {
	return _CacheManager.Contract.SweepFunds(&_CacheManager.TransactOpts)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns()
func (_CacheManager *CacheManagerTransactorSession) SweepFunds() (*types.Transaction, error) {
	return _CacheManager.Contract.SweepFunds(&_CacheManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CacheManager *CacheManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CacheManager *CacheManagerSession) Unpause() (*types.Transaction, error) {
	return _CacheManager.Contract.Unpause(&_CacheManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CacheManager *CacheManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _CacheManager.Contract.Unpause(&_CacheManager.TransactOpts)
}

// CacheManagerDeleteBidIterator is returned from FilterDeleteBid and is used to iterate over the raw logs and unpacked data for DeleteBid events raised by the CacheManager contract.
type CacheManagerDeleteBidIterator struct {
	Event *CacheManagerDeleteBid // Event containing the contract specifics and raw log

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
func (it *CacheManagerDeleteBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerDeleteBid)
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
		it.Event = new(CacheManagerDeleteBid)
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
func (it *CacheManagerDeleteBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerDeleteBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerDeleteBid represents a DeleteBid event raised by the CacheManager contract.
type CacheManagerDeleteBid struct {
	Codehash [32]byte
	Bid      *big.Int
	Size     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteBid is a free log retrieval operation binding the contract event 0x65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7.
//
// Solidity: event DeleteBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) FilterDeleteBid(opts *bind.FilterOpts, codehash [][32]byte) (*CacheManagerDeleteBidIterator, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "DeleteBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return &CacheManagerDeleteBidIterator{contract: _CacheManager.contract, event: "DeleteBid", logs: logs, sub: sub}, nil
}

// WatchDeleteBid is a free log subscription operation binding the contract event 0x65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7.
//
// Solidity: event DeleteBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) WatchDeleteBid(opts *bind.WatchOpts, sink chan<- *CacheManagerDeleteBid, codehash [][32]byte) (event.Subscription, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "DeleteBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerDeleteBid)
				if err := _CacheManager.contract.UnpackLog(event, "DeleteBid", log); err != nil {
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

// ParseDeleteBid is a log parse operation binding the contract event 0x65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7.
//
// Solidity: event DeleteBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) ParseDeleteBid(log types.Log) (*CacheManagerDeleteBid, error) {
	event := new(CacheManagerDeleteBid)
	if err := _CacheManager.contract.UnpackLog(event, "DeleteBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CacheManager contract.
type CacheManagerInitializedIterator struct {
	Event *CacheManagerInitialized // Event containing the contract specifics and raw log

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
func (it *CacheManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerInitialized)
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
		it.Event = new(CacheManagerInitialized)
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
func (it *CacheManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerInitialized represents a Initialized event raised by the CacheManager contract.
type CacheManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CacheManager *CacheManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*CacheManagerInitializedIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CacheManagerInitializedIterator{contract: _CacheManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CacheManager *CacheManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CacheManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerInitialized)
				if err := _CacheManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CacheManager *CacheManagerFilterer) ParseInitialized(log types.Log) (*CacheManagerInitialized, error) {
	event := new(CacheManagerInitialized)
	if err := _CacheManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerInsertBidIterator is returned from FilterInsertBid and is used to iterate over the raw logs and unpacked data for InsertBid events raised by the CacheManager contract.
type CacheManagerInsertBidIterator struct {
	Event *CacheManagerInsertBid // Event containing the contract specifics and raw log

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
func (it *CacheManagerInsertBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerInsertBid)
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
		it.Event = new(CacheManagerInsertBid)
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
func (it *CacheManagerInsertBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerInsertBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerInsertBid represents a InsertBid event raised by the CacheManager contract.
type CacheManagerInsertBid struct {
	Codehash [32]byte
	Program  common.Address
	Bid      *big.Int
	Size     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInsertBid is a free log retrieval operation binding the contract event 0xb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec.
//
// Solidity: event InsertBid(bytes32 indexed codehash, address program, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) FilterInsertBid(opts *bind.FilterOpts, codehash [][32]byte) (*CacheManagerInsertBidIterator, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "InsertBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return &CacheManagerInsertBidIterator{contract: _CacheManager.contract, event: "InsertBid", logs: logs, sub: sub}, nil
}

// WatchInsertBid is a free log subscription operation binding the contract event 0xb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec.
//
// Solidity: event InsertBid(bytes32 indexed codehash, address program, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) WatchInsertBid(opts *bind.WatchOpts, sink chan<- *CacheManagerInsertBid, codehash [][32]byte) (event.Subscription, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "InsertBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerInsertBid)
				if err := _CacheManager.contract.UnpackLog(event, "InsertBid", log); err != nil {
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

// ParseInsertBid is a log parse operation binding the contract event 0xb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec.
//
// Solidity: event InsertBid(bytes32 indexed codehash, address program, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) ParseInsertBid(log types.Log) (*CacheManagerInsertBid, error) {
	event := new(CacheManagerInsertBid)
	if err := _CacheManager.contract.UnpackLog(event, "InsertBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the CacheManager contract.
type CacheManagerPauseIterator struct {
	Event *CacheManagerPause // Event containing the contract specifics and raw log

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
func (it *CacheManagerPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerPause)
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
		it.Event = new(CacheManagerPause)
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
func (it *CacheManagerPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerPause represents a Pause event raised by the CacheManager contract.
type CacheManagerPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CacheManager *CacheManagerFilterer) FilterPause(opts *bind.FilterOpts) (*CacheManagerPauseIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &CacheManagerPauseIterator{contract: _CacheManager.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CacheManager *CacheManagerFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *CacheManagerPause) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerPause)
				if err := _CacheManager.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CacheManager *CacheManagerFilterer) ParsePause(log types.Log) (*CacheManagerPause, error) {
	event := new(CacheManagerPause)
	if err := _CacheManager.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerSetCacheSizeIterator is returned from FilterSetCacheSize and is used to iterate over the raw logs and unpacked data for SetCacheSize events raised by the CacheManager contract.
type CacheManagerSetCacheSizeIterator struct {
	Event *CacheManagerSetCacheSize // Event containing the contract specifics and raw log

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
func (it *CacheManagerSetCacheSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerSetCacheSize)
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
		it.Event = new(CacheManagerSetCacheSize)
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
func (it *CacheManagerSetCacheSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerSetCacheSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerSetCacheSize represents a SetCacheSize event raised by the CacheManager contract.
type CacheManagerSetCacheSize struct {
	Size uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetCacheSize is a free log retrieval operation binding the contract event 0xca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839.
//
// Solidity: event SetCacheSize(uint64 size)
func (_CacheManager *CacheManagerFilterer) FilterSetCacheSize(opts *bind.FilterOpts) (*CacheManagerSetCacheSizeIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "SetCacheSize")
	if err != nil {
		return nil, err
	}
	return &CacheManagerSetCacheSizeIterator{contract: _CacheManager.contract, event: "SetCacheSize", logs: logs, sub: sub}, nil
}

// WatchSetCacheSize is a free log subscription operation binding the contract event 0xca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839.
//
// Solidity: event SetCacheSize(uint64 size)
func (_CacheManager *CacheManagerFilterer) WatchSetCacheSize(opts *bind.WatchOpts, sink chan<- *CacheManagerSetCacheSize) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "SetCacheSize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerSetCacheSize)
				if err := _CacheManager.contract.UnpackLog(event, "SetCacheSize", log); err != nil {
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

// ParseSetCacheSize is a log parse operation binding the contract event 0xca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839.
//
// Solidity: event SetCacheSize(uint64 size)
func (_CacheManager *CacheManagerFilterer) ParseSetCacheSize(log types.Log) (*CacheManagerSetCacheSize, error) {
	event := new(CacheManagerSetCacheSize)
	if err := _CacheManager.contract.UnpackLog(event, "SetCacheSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerSetDecayRateIterator is returned from FilterSetDecayRate and is used to iterate over the raw logs and unpacked data for SetDecayRate events raised by the CacheManager contract.
type CacheManagerSetDecayRateIterator struct {
	Event *CacheManagerSetDecayRate // Event containing the contract specifics and raw log

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
func (it *CacheManagerSetDecayRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerSetDecayRate)
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
		it.Event = new(CacheManagerSetDecayRate)
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
func (it *CacheManagerSetDecayRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerSetDecayRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerSetDecayRate represents a SetDecayRate event raised by the CacheManager contract.
type CacheManagerSetDecayRate struct {
	Decay uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetDecayRate is a free log retrieval operation binding the contract event 0xd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c.
//
// Solidity: event SetDecayRate(uint64 decay)
func (_CacheManager *CacheManagerFilterer) FilterSetDecayRate(opts *bind.FilterOpts) (*CacheManagerSetDecayRateIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "SetDecayRate")
	if err != nil {
		return nil, err
	}
	return &CacheManagerSetDecayRateIterator{contract: _CacheManager.contract, event: "SetDecayRate", logs: logs, sub: sub}, nil
}

// WatchSetDecayRate is a free log subscription operation binding the contract event 0xd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c.
//
// Solidity: event SetDecayRate(uint64 decay)
func (_CacheManager *CacheManagerFilterer) WatchSetDecayRate(opts *bind.WatchOpts, sink chan<- *CacheManagerSetDecayRate) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "SetDecayRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerSetDecayRate)
				if err := _CacheManager.contract.UnpackLog(event, "SetDecayRate", log); err != nil {
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

// ParseSetDecayRate is a log parse operation binding the contract event 0xd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c.
//
// Solidity: event SetDecayRate(uint64 decay)
func (_CacheManager *CacheManagerFilterer) ParseSetDecayRate(log types.Log) (*CacheManagerSetDecayRate, error) {
	event := new(CacheManagerSetDecayRate)
	if err := _CacheManager.contract.UnpackLog(event, "SetDecayRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the CacheManager contract.
type CacheManagerUnpauseIterator struct {
	Event *CacheManagerUnpause // Event containing the contract specifics and raw log

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
func (it *CacheManagerUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerUnpause)
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
		it.Event = new(CacheManagerUnpause)
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
func (it *CacheManagerUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerUnpause represents a Unpause event raised by the CacheManager contract.
type CacheManagerUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CacheManager *CacheManagerFilterer) FilterUnpause(opts *bind.FilterOpts) (*CacheManagerUnpauseIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &CacheManagerUnpauseIterator{contract: _CacheManager.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CacheManager *CacheManagerFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *CacheManagerUnpause) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerUnpause)
				if err := _CacheManager.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CacheManager *CacheManagerFilterer) ParseUnpause(log types.Log) (*CacheManagerUnpause, error) {
	event := new(CacheManagerUnpause)
	if err := _CacheManager.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
