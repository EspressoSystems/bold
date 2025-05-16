// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test_helpersgen

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

// BridgeTesterMetaData contains all meta data concerning the BridgeTester contract.
var BridgeTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotDelayedInbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotOutbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotRollupOrOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"RollupUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedDelayedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"rollup_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerReportedSubMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"_rollup\",\"type\":\"address\"}],\"name\":\"updateRollupAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a08060405234601f573060805261116090816100258239608051815050f35b600080fdfe6080604081815260049182361015610022575b505050361561002057600080fd5b005b600092833560e01c91826284120c14610bef5750816316bf557914610bbf578163413b35bd14610b7d57816347fb24c514610ae05781634f61f850146109f35781635fca4a16146109d45781637a88b107146109af57816386598a56146109825781638db5993b14610777578163919cc7061461073a578163945e1147146106fa5781639e5d4c49146104fa578163ab5d8943146104cc578163ad48cb5e146104a7578163ae60bd1314610466578163c4d66de814610322578163cb23bcb5146102f9578163cee3d7281461023b578163d5719dc214610202578163e1758bd8146101d9578163e76f5c8d1461019457508063e77145f414610182578063eca067ad146101645763ee35f327146101395780610012565b3461016057816003193601126101605760075490516001600160a01b039091168152602090f35b5080fd5b50346101605781600319360112610160576020906009549051908152f35b82806003193601126101915780f35b80fd5b9050346101d55760203660031901126101d557359160035483101561019157506101bf602092610c57565b60018060a01b0391549060031b1c169051908152f35b8280fd5b50503461016057816003193601126101605760085490516001600160a01b039091168152602090f35b9050346101d55760203660031901126101d5573591600954831015610191575061022d602092610c0b565b91905490519160031b1c8152f35b839150346101605761024c36610c8d565b9160018060a01b03806006541695863303610270575b8561026d8686610f5c565b80f35b8051638da5cb5b60e01b81529160208385818b5afa9283156102ef5787936102be575b5082163303156102625751630739600760e01b81529182916102ba91889033908501610d2f565b0390fd5b6102e191935060203d6020116102e8575b6102d98183610ced565b810190610d10565b9188610293565b503d6102cf565b82513d89823e3d90fd5b50503461016057816003193601126101605760065490516001600160a01b039091168152602090f35b9050346101d55760203660031901126101d55780356001600160a01b038181169182900361046257845460ff8160081c161593848095610455575b801561043e575b156103e4575060ff1981166001178655836103d3575b5060018060a01b031990816005541617600555600654161760065561039d575080f35b60207f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989161ff001984541684555160018152a180f35b61ffff19166101011785553861037a565b608490602087519162461bcd60e51b8352820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152fd5b50303b1580156103645750600160ff831614610364565b50600160ff83161061035d565b8480fd5b5050346101605760203660031901126101605760209160ff9060019083906001600160a01b03610494610c72565b1681528286522001541690519015158152f35b50503461016057816003193601126101605760209060ff60085460a01c169051908152f35b5050346101605781600319360112610160576020906104e9610f3e565b90516001600160a01b039091168152f35b828434610191576060918260031936011261016057610517610c72565b916044356001600160401b03602480359082841161046257366023850112156104625783890135908382116106f65780850194818336920101116106f6573386526020946002865260ff6001898920015416156106e257821515806106d9575b6106bc57600580546001600160a01b0319808216331790925589516001600160a01b0394908a9081908e90898882378a818b810185815203925af19d3d15610697573d9889116106865750509184918460008051602061110b83398151915295948c51996105ee8c601f19601f840116018c610ced565b8a523d8c8c8c013e9e9b999e9c9a989c5b169060055416176005558751958652878d87015281888701528b860137848a8486010152601f19961692898188601f339601168101030190a381519687941515855282818601528551809386015281955b83871061066e5750508394508582601f949501015201168101030190f35b86810182015189880189015295810195889550610650565b634e487b7160e01b8b526041905289fd5b505095508291849160008051602061110b833981519152948d9c9a989e9b999e6105ff565b50865163b5cf5b8f60e01b81526001600160a01b038916818c0152fd5b50883b15610577565b5086516332ea82ab60e01b815233818c0152fd5b8580fd5b82843461019157602036600319011261019157823592548310156101915750610724602092610c3c565b905491519160018060a01b039160031b1c168152f35b83903461016057602036600319011261016057356001600160a01b0381169081900361016057600680546001600160a01b03191691909117905580f35b91905060603660031901126101d55781359260ff8416938481036101605760248035946001600160a01b038616908187036104625760443592338652600160205260ff60018888200154161561096d57600954875160f89690961b6001600160f81b0319166020870190815260609990991b6001600160601b03191660218701524360c090811b6001600160c01b0319908116603589015242821b16603d880152604587018290524860658801526085808801879052875290986001600160401b0396909181018781118282101761095b578952519020869089610926575b885160208101918983528a8201528981526060810181811089821117610913578a5251902092600160401b8a1015610903575050509686937f5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe19360c09360209a600188016009556108c688610c0b565b819291549060031b91821b91600019901b19161790558851933385528b850152888401526060830152486080830152421660a0820152a351908152f35b634e487b7160e01b825260419052fd5b5050634e487b7160e01b82525060418352fd5b9650600019890189811161094a5761093d90610c0b565b90549060031b1c96610856565b50634e487b7160e01b815260118352fd5b634e487b7160e01b8952604185528389fd5b9086519063b6c60ea360e01b82523390820152fd5b50503461016057608036600319011261016057816080928251928184528160208501528301526060820152f35b505034610160578060031936011261016057906020916109cd610c72565b5051908152f35b505034610160578160031936011261016057602090600b549051908152f35b9050346101d55760203660031901126101d557610a0e610c72565b6006546001600160a01b03929083169033829003610a69575b857f8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a602087878716908160018060a01b0319600754161760075551908152a180f35b8451638da5cb5b60e01b81526020818381865afa908115610ad6578791610ab7575b508481163303610a9b5750610a27565b6102ba908651938493630739600760e01b855233908501610d2f565b610ad0915060203d6020116102e8576102d98183610ced565b38610a8b565b86513d89823e3d90fd5b8391503461016057610af136610c8d565b9160018060a01b03806006541695863303610b12575b8561026d8686610d6f565b8051638da5cb5b60e01b81529160208385818b5afa9283156102ef578793610b5c575b508216330315610b075751630739600760e01b81529182916102ba91889033908501610d2f565b610b7691935060203d6020116102e8576102d98183610ced565b9188610b35565b5050346101605760203660031901126101605760209160ff9060019083906001600160a01b03610bab610c72565b168152600286522001541690519015158152f35b9050346101d55760203660031901126101d55735600a548110156101d557602083600a8295522001549051908152f35b849034610160578160031936011261016057602090600a548152f35b600954811015610c2657600960005260206000200190600090565b634e487b7160e01b600052603260045260246000fd5b600454811015610c2657600460005260206000200190600090565b600354811015610c2657600360005260206000200190600090565b600435906001600160a01b0382168203610c8857565b600080fd5b6040906003190112610c88576004356001600160a01b0381168103610c8857906024358015158103610c885790565b604081019081106001600160401b03821117610cd757604052565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b03821190821017610cd757604052565b90816020910312610c8857516001600160a01b0381168103610c885790565b6001600160a01b0391821681529181166020830152909116604082015260600190565b919060018060a01b038084549260031b9316831b921b1916179055565b9060018060a01b038083166000928184526001926020918483526040918287209160ff87840154168451877f6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf5218885159384158152a28180610f36575b8215610f1d575b5050610f125715610e565750509083918260035492825193610df385610cbc565b845280840195828752885252852090518155019051151560ff8019835416911617905560035491600160401b831015610e42575081610e3a91610e40949301600355610c57565b90610d52565b565b634e487b7160e01b81526041600452602490fd5b9091949596506003939293549160001992838101908111610efe5782610e7e610e9092610c57565b90549060031b1c16610e3a8354610c57565b5490610e9b82610c57565b90549060031b1c168752858452848720556003548015610eea578593929101610edb610ec682610c57565b81549060018060a01b039060031b1b19169055565b60035585525282208281550155565b634e487b7160e01b87526031600452602487fd5b634e487b7160e01b89526011600452602489fd5b505050505050505050565b15915081610f2e575b503880610dd2565b905038610f26565b839250610dcb565b6005546001600160a01b03908116908114610f565790565b50600090565b9060018060a01b039081831690600092828452602091600283526040918286209160ff6001840154168451877f49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa8885159384158152a28180611102575b82156110e9575b50506110df571561102b57505090600191600260045492825193610fe385610cbc565b845280840195858752875252842090518155019051151560ff8019835416911617905560045490600160401b821015610e42575090610e3a826001610e409401600455610c3c565b909194955060049392935491600019928381019081116110cb5761105161106c91610c3c565b90548461105e8554610c3c565b92909360031b1c1691610d52565b549061107782610c3c565b90549060031b1c168652600284528486205560045480156110b757916001949391600293016110a8610ec682610c3c565b60045585525282208281550155565b634e487b7160e01b86526031600452602486fd5b634e487b7160e01b88526011600452602488fd5b5050505050505050565b159150816110fa575b503880610fc0565b9050386110f2565b839250610fb956fe2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466a2646970667358221220fe148f87af6efec629af0a1edc2694f92fff68bc024018a3909fe3966e6dab9164736f6c63430008190033",
}

// BridgeTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTesterMetaData.ABI instead.
var BridgeTesterABI = BridgeTesterMetaData.ABI

// BridgeTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeTesterMetaData.Bin instead.
var BridgeTesterBin = BridgeTesterMetaData.Bin

// DeployBridgeTester deploys a new Ethereum contract, binding an instance of BridgeTester to it.
func DeployBridgeTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeTester, error) {
	parsed, err := BridgeTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeTester{BridgeTesterCaller: BridgeTesterCaller{contract: contract}, BridgeTesterTransactor: BridgeTesterTransactor{contract: contract}, BridgeTesterFilterer: BridgeTesterFilterer{contract: contract}}, nil
}

// BridgeTester is an auto generated Go binding around an Ethereum contract.
type BridgeTester struct {
	BridgeTesterCaller     // Read-only binding to the contract
	BridgeTesterTransactor // Write-only binding to the contract
	BridgeTesterFilterer   // Log filterer for contract events
}

// BridgeTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTesterSession struct {
	Contract     *BridgeTester     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTesterCallerSession struct {
	Contract *BridgeTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTesterTransactorSession struct {
	Contract     *BridgeTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTesterRaw struct {
	Contract *BridgeTester // Generic contract binding to access the raw methods on
}

// BridgeTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTesterCallerRaw struct {
	Contract *BridgeTesterCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTesterTransactorRaw struct {
	Contract *BridgeTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTester creates a new instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTester(address common.Address, backend bind.ContractBackend) (*BridgeTester, error) {
	contract, err := bindBridgeTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTester{BridgeTesterCaller: BridgeTesterCaller{contract: contract}, BridgeTesterTransactor: BridgeTesterTransactor{contract: contract}, BridgeTesterFilterer: BridgeTesterFilterer{contract: contract}}, nil
}

// NewBridgeTesterCaller creates a new read-only instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTesterCaller(address common.Address, caller bind.ContractCaller) (*BridgeTesterCaller, error) {
	contract, err := bindBridgeTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterCaller{contract: contract}, nil
}

// NewBridgeTesterTransactor creates a new write-only instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTesterTransactor, error) {
	contract, err := bindBridgeTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterTransactor{contract: contract}, nil
}

// NewBridgeTesterFilterer creates a new log filterer instance of BridgeTester, bound to a specific deployed contract.
func NewBridgeTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTesterFilterer, error) {
	contract, err := bindBridgeTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterFilterer{contract: contract}, nil
}

// bindBridgeTester binds a generic wrapper to an already deployed contract.
func bindBridgeTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTester *BridgeTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTester.Contract.BridgeTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTester *BridgeTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.Contract.BridgeTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTester *BridgeTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTester.Contract.BridgeTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTester *BridgeTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTester *BridgeTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTester *BridgeTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTester.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeTester *BridgeTesterCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeTester *BridgeTesterSession) ActiveOutbox() (common.Address, error) {
	return _BridgeTester.Contract.ActiveOutbox(&_BridgeTester.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) ActiveOutbox() (common.Address, error) {
	return _BridgeTester.Contract.ActiveOutbox(&_BridgeTester.CallOpts)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCaller) AllowedDelayedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedDelayedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeTester *BridgeTesterCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeTester *BridgeTesterSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxes(&_BridgeTester.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeTester *BridgeTesterCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedDelayedInboxes(&_BridgeTester.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedOutboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeTester.Contract.AllowedOutboxList(&_BridgeTester.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeTester *BridgeTesterCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeTester *BridgeTesterSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedOutboxes(&_BridgeTester.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeTester *BridgeTesterCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeTester.Contract.AllowedOutboxes(&_BridgeTester.CallOpts, outbox)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCaller) DelayedInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "delayedInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.DelayedInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCallerSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.DelayedInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.DelayedMessageCount(&_BridgeTester.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.DelayedMessageCount(&_BridgeTester.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_BridgeTester *BridgeTesterCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_BridgeTester *BridgeTesterSession) NativeToken() (common.Address, error) {
	return _BridgeTester.Contract.NativeToken(&_BridgeTester.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) NativeToken() (common.Address, error) {
	return _BridgeTester.Contract.NativeToken(&_BridgeTester.CallOpts)
}

// NativeTokenDecimals is a free data retrieval call binding the contract method 0xad48cb5e.
//
// Solidity: function nativeTokenDecimals() view returns(uint8)
func (_BridgeTester *BridgeTesterCaller) NativeTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "nativeTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NativeTokenDecimals is a free data retrieval call binding the contract method 0xad48cb5e.
//
// Solidity: function nativeTokenDecimals() view returns(uint8)
func (_BridgeTester *BridgeTesterSession) NativeTokenDecimals() (uint8, error) {
	return _BridgeTester.Contract.NativeTokenDecimals(&_BridgeTester.CallOpts)
}

// NativeTokenDecimals is a free data retrieval call binding the contract method 0xad48cb5e.
//
// Solidity: function nativeTokenDecimals() view returns(uint8)
func (_BridgeTester *BridgeTesterCallerSession) NativeTokenDecimals() (uint8, error) {
	return _BridgeTester.Contract.NativeTokenDecimals(&_BridgeTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeTester *BridgeTesterCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeTester *BridgeTesterSession) Rollup() (common.Address, error) {
	return _BridgeTester.Contract.Rollup(&_BridgeTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) Rollup() (common.Address, error) {
	return _BridgeTester.Contract.Rollup(&_BridgeTester.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeTester *BridgeTesterCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeTester *BridgeTesterSession) SequencerInbox() (common.Address, error) {
	return _BridgeTester.Contract.SequencerInbox(&_BridgeTester.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeTester *BridgeTesterCallerSession) SequencerInbox() (common.Address, error) {
	return _BridgeTester.Contract.SequencerInbox(&_BridgeTester.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCaller) SequencerInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.SequencerInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeTester *BridgeTesterCallerSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeTester.Contract.SequencerInboxAccs(&_BridgeTester.CallOpts, arg0)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerMessageCount(&_BridgeTester.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerMessageCount(&_BridgeTester.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCaller) SequencerReportedSubMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTester.contract.Call(opts, &out, "sequencerReportedSubMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerReportedSubMessageCount(&_BridgeTester.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeTester *BridgeTesterCallerSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeTester.Contract.SequencerReportedSubMessageCount(&_BridgeTester.CallOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeTester *BridgeTesterTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeTester *BridgeTesterSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeTester.Contract.AcceptFundsFromOldBridge(&_BridgeTester.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeTester *BridgeTesterTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeTester.Contract.AcceptFundsFromOldBridge(&_BridgeTester.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeTester *BridgeTesterTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeTester *BridgeTesterSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueDelayedMessage(&_BridgeTester.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeTester *BridgeTesterTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueDelayedMessage(&_BridgeTester.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeTester *BridgeTesterTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeTester *BridgeTesterSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueSequencerMessage(&_BridgeTester.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeTester *BridgeTesterTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeTester.Contract.EnqueueSequencerMessage(&_BridgeTester.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeTester *BridgeTesterTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "executeCall", to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeTester *BridgeTesterSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.ExecuteCall(&_BridgeTester.TransactOpts, to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeTester *BridgeTesterTransactorSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.ExecuteCall(&_BridgeTester.TransactOpts, to, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeTester *BridgeTesterTransactor) Initialize(opts *bind.TransactOpts, rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "initialize", rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeTester *BridgeTesterSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.Initialize(&_BridgeTester.TransactOpts, rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeTester *BridgeTesterTransactorSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.Initialize(&_BridgeTester.TransactOpts, rollup_)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetDelayedInbox(&_BridgeTester.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetDelayedInbox(&_BridgeTester.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactor) SetOutbox(opts *bind.TransactOpts, outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "setOutbox", outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetOutbox(&_BridgeTester.TransactOpts, outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeTester *BridgeTesterTransactorSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetOutbox(&_BridgeTester.TransactOpts, outbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeTester *BridgeTesterTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeTester *BridgeTesterSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetSequencerInbox(&_BridgeTester.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeTester *BridgeTesterTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.SetSequencerInbox(&_BridgeTester.TransactOpts, _sequencerInbox)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeTester *BridgeTesterTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "submitBatchSpendingReport", batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeTester *BridgeTesterSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.SubmitBatchSpendingReport(&_BridgeTester.TransactOpts, batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeTester *BridgeTesterTransactorSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeTester.Contract.SubmitBatchSpendingReport(&_BridgeTester.TransactOpts, batchPoster, dataHash)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address _rollup) returns()
func (_BridgeTester *BridgeTesterTransactor) UpdateRollupAddress(opts *bind.TransactOpts, _rollup common.Address) (*types.Transaction, error) {
	return _BridgeTester.contract.Transact(opts, "updateRollupAddress", _rollup)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address _rollup) returns()
func (_BridgeTester *BridgeTesterSession) UpdateRollupAddress(_rollup common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.UpdateRollupAddress(&_BridgeTester.TransactOpts, _rollup)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x919cc706.
//
// Solidity: function updateRollupAddress(address _rollup) returns()
func (_BridgeTester *BridgeTesterTransactorSession) UpdateRollupAddress(_rollup common.Address) (*types.Transaction, error) {
	return _BridgeTester.Contract.UpdateRollupAddress(&_BridgeTester.TransactOpts, _rollup)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeTester *BridgeTesterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTester.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeTester *BridgeTesterSession) Receive() (*types.Transaction, error) {
	return _BridgeTester.Contract.Receive(&_BridgeTester.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BridgeTester *BridgeTesterTransactorSession) Receive() (*types.Transaction, error) {
	return _BridgeTester.Contract.Receive(&_BridgeTester.TransactOpts)
}

// BridgeTesterBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the BridgeTester contract.
type BridgeTesterBridgeCallTriggeredIterator struct {
	Event *BridgeTesterBridgeCallTriggered // Event containing the contract specifics and raw log

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
func (it *BridgeTesterBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterBridgeCallTriggered)
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
		it.Event = new(BridgeTesterBridgeCallTriggered)
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
func (it *BridgeTesterBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterBridgeCallTriggered represents a BridgeCallTriggered event raised by the BridgeTester contract.
type BridgeTesterBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeTester *BridgeTesterFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*BridgeTesterBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterBridgeCallTriggeredIterator{contract: _BridgeTester.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeTester *BridgeTesterFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *BridgeTesterBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterBridgeCallTriggered)
				if err := _BridgeTester.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
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

// ParseBridgeCallTriggered is a log parse operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeTester *BridgeTesterFilterer) ParseBridgeCallTriggered(log types.Log) (*BridgeTesterBridgeCallTriggered, error) {
	event := new(BridgeTesterBridgeCallTriggered)
	if err := _BridgeTester.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the BridgeTester contract.
type BridgeTesterInboxToggleIterator struct {
	Event *BridgeTesterInboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeTesterInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterInboxToggle)
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
		it.Event = new(BridgeTesterInboxToggle)
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
func (it *BridgeTesterInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterInboxToggle represents a InboxToggle event raised by the BridgeTester contract.
type BridgeTesterInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*BridgeTesterInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterInboxToggleIterator{contract: _BridgeTester.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeTesterInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterInboxToggle)
				if err := _BridgeTester.contract.UnpackLog(event, "InboxToggle", log); err != nil {
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

// ParseInboxToggle is a log parse operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) ParseInboxToggle(log types.Log) (*BridgeTesterInboxToggle, error) {
	event := new(BridgeTesterInboxToggle)
	if err := _BridgeTester.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeTester contract.
type BridgeTesterInitializedIterator struct {
	Event *BridgeTesterInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeTesterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterInitialized)
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
		it.Event = new(BridgeTesterInitialized)
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
func (it *BridgeTesterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterInitialized represents a Initialized event raised by the BridgeTester contract.
type BridgeTesterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeTester *BridgeTesterFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeTesterInitializedIterator, error) {

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeTesterInitializedIterator{contract: _BridgeTester.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeTester *BridgeTesterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeTesterInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterInitialized)
				if err := _BridgeTester.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeTester *BridgeTesterFilterer) ParseInitialized(log types.Log) (*BridgeTesterInitialized, error) {
	event := new(BridgeTesterInitialized)
	if err := _BridgeTester.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the BridgeTester contract.
type BridgeTesterMessageDeliveredIterator struct {
	Event *BridgeTesterMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeTesterMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterMessageDelivered)
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
		it.Event = new(BridgeTesterMessageDelivered)
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
func (it *BridgeTesterMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterMessageDelivered represents a MessageDelivered event raised by the BridgeTester contract.
type BridgeTesterMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	BaseFeeL1       *big.Int
	Timestamp       uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeTester *BridgeTesterFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeTesterMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterMessageDeliveredIterator{contract: _BridgeTester.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeTester *BridgeTesterFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeTesterMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterMessageDelivered)
				if err := _BridgeTester.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeTester *BridgeTesterFilterer) ParseMessageDelivered(log types.Log) (*BridgeTesterMessageDelivered, error) {
	event := new(BridgeTesterMessageDelivered)
	if err := _BridgeTester.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the BridgeTester contract.
type BridgeTesterOutboxToggleIterator struct {
	Event *BridgeTesterOutboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeTesterOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterOutboxToggle)
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
		it.Event = new(BridgeTesterOutboxToggle)
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
func (it *BridgeTesterOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterOutboxToggle represents a OutboxToggle event raised by the BridgeTester contract.
type BridgeTesterOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*BridgeTesterOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTesterOutboxToggleIterator{contract: _BridgeTester.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeTesterOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterOutboxToggle)
				if err := _BridgeTester.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
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

// ParseOutboxToggle is a log parse operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeTester *BridgeTesterFilterer) ParseOutboxToggle(log types.Log) (*BridgeTesterOutboxToggle, error) {
	event := new(BridgeTesterOutboxToggle)
	if err := _BridgeTester.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterRollupUpdatedIterator is returned from FilterRollupUpdated and is used to iterate over the raw logs and unpacked data for RollupUpdated events raised by the BridgeTester contract.
type BridgeTesterRollupUpdatedIterator struct {
	Event *BridgeTesterRollupUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTesterRollupUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterRollupUpdated)
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
		it.Event = new(BridgeTesterRollupUpdated)
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
func (it *BridgeTesterRollupUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterRollupUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterRollupUpdated represents a RollupUpdated event raised by the BridgeTester contract.
type BridgeTesterRollupUpdated struct {
	Rollup common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRollupUpdated is a free log retrieval operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeTester *BridgeTesterFilterer) FilterRollupUpdated(opts *bind.FilterOpts) (*BridgeTesterRollupUpdatedIterator, error) {

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "RollupUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTesterRollupUpdatedIterator{contract: _BridgeTester.contract, event: "RollupUpdated", logs: logs, sub: sub}, nil
}

// WatchRollupUpdated is a free log subscription operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeTester *BridgeTesterFilterer) WatchRollupUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTesterRollupUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "RollupUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterRollupUpdated)
				if err := _BridgeTester.contract.UnpackLog(event, "RollupUpdated", log); err != nil {
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

// ParseRollupUpdated is a log parse operation binding the contract event 0xae1f5aa15f6ff844896347ceca2a3c24c8d3a27785efdeacd581a0a95172784a.
//
// Solidity: event RollupUpdated(address rollup)
func (_BridgeTester *BridgeTesterFilterer) ParseRollupUpdated(log types.Log) (*BridgeTesterRollupUpdated, error) {
	event := new(BridgeTesterRollupUpdated)
	if err := _BridgeTester.contract.UnpackLog(event, "RollupUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTesterSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the BridgeTester contract.
type BridgeTesterSequencerInboxUpdatedIterator struct {
	Event *BridgeTesterSequencerInboxUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTesterSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTesterSequencerInboxUpdated)
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
		it.Event = new(BridgeTesterSequencerInboxUpdated)
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
func (it *BridgeTesterSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTesterSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTesterSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the BridgeTester contract.
type BridgeTesterSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeTester *BridgeTesterFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*BridgeTesterSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _BridgeTester.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTesterSequencerInboxUpdatedIterator{contract: _BridgeTester.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeTester *BridgeTesterFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTesterSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTester.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTesterSequencerInboxUpdated)
				if err := _BridgeTester.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
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

// ParseSequencerInboxUpdated is a log parse operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeTester *BridgeTesterFilterer) ParseSequencerInboxUpdated(log types.Log) (*BridgeTesterSequencerInboxUpdated, error) {
	event := new(BridgeTesterSequencerInboxUpdated)
	if err := _BridgeTester.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptographyPrimitivesTesterMetaData contains all meta data concerning the CryptographyPrimitivesTester contract.
var CryptographyPrimitivesTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"input\",\"type\":\"uint256[25]\"}],\"name\":\"keccakF\",\"outputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"\",\"type\":\"uint256[25]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[2]\",\"name\":\"inputChunk\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32\",\"name\":\"hashState\",\"type\":\"bytes32\"}],\"name\":\"sha256Block\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608080604052346019576112ef908161001f823930815050f35b600080fdfe608060405260048036101561001357600080fd5b60003560e01c8063ac90ed46146107de5763e479f5321461003357600080fd5b60603660031901126107c45736602312156107c457604080516001600160401b0392918101838111828210176107c9576040526044813682116107c45783905b8282106107b45750505060405190604082018281108582111761079f5790602091604052805183520151906020810191825260405190610800928383018381108782111761078a5760405263428a2f9883526371374491602084015263b5c0fbcf604084015263e9b5dba56060840152633956c25b60808401526359f111f160a084015263923f82a460c084015263ab1c5ed560e08401526101009163d807aa98838501526312835b0161012085015263243185be61014085015263550c7dc36101608501526372be5d746101808501526380deb1fe6101a0850152639bdc06a76101c085015263c19bf1746101e085015263e49b69c161020085015263efbe4786610220850152630fc19dc661024085015263240ca1cc610260850152632de92c6f610280850152634a7484aa6102a0850152635cb0a9dc6102c08501526376f988da6102e085015263983e515261030085015263a831c66d61032085015263b00327c861034085015263bf597fc761036085015263c6e00bf361038085015263d5a791476103a08501526306ca63516103c085015263142929676103e08501526327b70a85610400850152632e1b2138610420850152634d2c6dfc6104408501526353380d1361046085015263650a735461048085015263766a0abb6104a08501526381c2c92e6104c08501526392722c856104e085015263a2bfe8a161050085015263a81a664b61052085015263c24b8b7061054085015263c76c51a361056085015263d192e81961058085015263d69906246105a085015263f40e35856105c085015263106aa0706105e08501526319a4c116610600850152631e376c08610620850152632748774c6106408501526334b0bcb561066085015263391c0cb3610680850152634ed8aa4a6106a0850152635b9cca4f6106c085015263682e6ff36106e085015263748f82ee6107008501526378a5636f6107208501526384c87814610740850152638cc702086107608501526390befffa61078085015263a4506ceb6107a085015263bef9a3f76107c085015263c67178f26107e085015260405194808601868110898211176107755760405236863760005b600863ffffffff828116828110156103e657916103de8263ffffffff95936001958c6103ce8a51928480808b60051b1660e0031694851c1692611281565b528951901c16928401168a611281565b520116610390565b8a888a898d60105b63ffffffff80821690604082101561050257916001916104fa8488818063ffffffff98600e198801166104218185611281565b51631fffffff61043c6104348488611281565b519387611281565b5160031c1691613fff63ffffc00082600e1b169160121c1617906301ffffff63fe0000008260191b169160071c16171818816104d8818b198b01166104818188611281565b51623fffff61049b610493848b611281565b51938a611281565b51600a1c1691611fff63ffffe00082600d1b169160131c161790617fff63ffff800082600f1b169160111c161718189582600f198c011690611281565b51160116816104ed600619890182168e611281565b5116011601169188611281565b5201166103ee565b505060405195838701929150821186831017610760575060405236843760005b63ffffffff80821691600883101561055f5763ffffffff9261055760019360443581808660051b1660e003161c1691886112a8565b520116610522565b50505060005b604063ffffffff821610156106c25760018163ffffffff80808095818960808c0151828d81806105e98163ffffff808760071b16607f8860191c161763ffe000008860151b16621fffff89600b1c161763fc00000089601a1b166303ffffff8a60061c161718188160e0870151160116968260c060a08701519601519a1690611281565b51169681851619161692161618011601168161060781861689611281565b5116011681885160208a0151828b8180841693604083015190828216938360c08201511660e08201528360a08201511660c08201528360808201511660a0820152838a8160608401511601166080820152846060820152866040820152602084891691015216169218831616189063fffffc0081600a1b166103ff8260161c16179063fff800008160131b166207ffff82600d1c161790633fffffff63c000000082601e1b169160021c161718180116011686520116610565565b8360005b63ffffffff8082169160088310156107135763ffffffff9261070b8360019460443581808760051b1660e003161c1681610700858a6112a8565b5116011691866112a8565b5201166106c6565b836000805b63ffffffff808316600881101561075557918160019261073d63ffffffff95886112a8565b511690808660051b1660e003161b1792011690610718565b602083604051908152f35b604190634e487b7160e01b6000525260246000fd5b604188634e487b7160e01b6000525260246000fd5b604186634e487b7160e01b6000525260246000fd5b604184634e487b7160e01b6000525260246000fd5b8135815260209182019101610073565b600080fd5b604183634e487b7160e01b6000525260246000fd5b506103203660031901126107c45736602312156107c4576107fd61122c565b90610324823682116107c45782905b82821061121c575050506103203661082261122c565b376103203661082f61122c565b37610838611262565b9060a0368337610846611262565b60a036823761085361122c565b91610320368437604051906103008201906001600160401b038211838310176107605750604090815260018252618082602083015267800000000000808a90820152678000000080008000606082015261808b6080820152638000000160a0820181905267800000008000808160c0830181905267800000000000800960e0840152608a61010084015260886101208401526380008009610140840152638000000a610160840152638000808b610180840152608b6001603f1b016101a08401526780000000000080896101c08401526780000000000080036101e084015267800000000000800261020084015260806001603f1b0161022084015261800a61024084015267800000008000000a6102608401526102808301526780000000000080806102a08301526102c08201526780000000800080086102e082015260005b601881106109c857604051866000825b601982106109b25761032084f35b60208060019285518152019301910190916109a4565b6001908651602088015118604088015118606088015118608088015118865260a087015160c08801511860e088015118610100880151186101208801511880602088015261014088015161016089015118610180890151186101a0890151186101c08901511860408801526101e08801516102008901511861022089015118610240890151186102608901511860608801526102808801516102a0890151186102c0890151186102e089015118610300890151189081608089015280603f1c90848060401b0390851b1617188085528651604088015180603f1c90858060401b0390861b16171860208601526020870151606088015180603f1c90858060401b0390861b16171860408601526040870151608088015180603f1c90858060401b0390861b16171860608601526060870151875180603f1c90858060401b0390861b16171860808601528751188088526020880151855118602089015260408801518551186040890152606088015185511860608901526080880151855118608089015260a088015160208601511860a089015260c088015160208601511860c089015260e088015160208601511860e08901526101008801516020860151186101008901526101208801516020860151186101208901526101408801516040860151186101408901526101608801516040860151186101608901526101808801516040860151186101808901526101a08801516040860151186101a08901526101c08801516040860151186101c08901526101e08801516060860151186101e08901526102008801516060860151186102008901526102208801516060860151186102208901526102408801516060860151186102408901526102608801516060860151186102608901526102808801516080860151186102808901526102a08801516080860151186102a08901526102c08801516080860151186102c08901526102e08801516080860151186102e0890152610300880151608086015118610300890152808652602088015180601c1c90848060401b039060241b1617610100870152604088015180603d1c90848060401b039060031b161761016087015260608801518060171c90848060401b039060291b1617610260870152608088015180602e1c90848060401b039060121b16176102c087015260a088015180603f1c90848060401b0390851b1617604087015260c0880151908160141c848060401b0383602c1b161760a088015260e08901518060361c90858060401b0390600a1b16176101a08801526101008901518060131c90858060401b0390602d1b1617610200880152610120890151600290603e9080821c90878060401b0390841b16176103008a01526101408b015180921c91868060401b03911b1617608088015261016089015180603a1c90858060401b039060061b161760e0880152610180890151918260151c858060401b0384602b1b16176101408901526101a08a01518060311c90868060401b0390600f1b16176102408901526101c08a01518060031c90868060401b0390603d1b16176102a08901526101e08a01518060241c90868060401b0390601c1b161760208901526102008a01518060091c90868060401b039060371b16176101208901526102208a01518060271c90868060401b039060191b16176101808901526102408a015180602b1c90868060401b039060151b16176101e08901526102608a01518060081c90868060401b039060381b16176102e08901526102808a01518060251c90868060401b0390601b1b161760608901526102a08a015180602c1c90868060401b039060141b161760c08901526102c08a01518060191c90868060401b039060271b16176101c08901526102e08a01518060381c90868060401b039060081b16176102208901526103008a01518060321c90868060401b0390600e1b16176102808901528260151c858060401b0384602b1b16178160141c868060401b0383602c1b1617191682188a52602088015160c0890151196101608a0151161860208b0152604088015160e0890151196101808a0151161860408b01526060880151610100890151196101a08a0151161860608b01526080880151610120890151196101c08a0151161860808b015260a0880151610140890151196101e08a0151161860a08b015260c0880151610160890151196102008a0151161860c08b015260e0880151610180890151196102208a0151161860e08b01526101008801516101a0890151196102408a015116186101008b01526101208801516101c0890151196102608a015116186101208b01526101408801516101e0890151196102808a015116186101408b0152610160880151610200890151196102a08a015116186101608b0152610180880151610220890151196102c08a015116186101808b01526101a0880151610240890151196102e08a015116186101a08b01526101c0880151610260890151196103008a015116186101c08b01526101e088015161028089015119895116186101e08b01526102008801516102a08901511960208a015116186102008b01526102208801516102c08901511960408a015116186102208b01526102408801516102e08901511960608a015116186102408b01526102608801516103008901511960808a015116186102608b015261028088015188511960a08a015116186102808b01526102a088015160208901511960c08a015116186102a08b01526102c088015160408901511960e08a015116186102c08b01526102e08801516060890151196101008a015116186102e08b01526103008801516080890151196101208a015116186103008b01528360051b860151928060151c90868060401b0390602b1b1617908060141c90868060401b0390602c1b161719161818875201610994565b813581526020918201910161080c565b6040519061032082016001600160401b0381118382101761124c57604052565b634e487b7160e01b600052604160045260246000fd5b6040519060a082016001600160401b0381118382101761124c57604052565b9060408110156112925760051b0190565b634e487b7160e01b600052603260045260246000fd5b9060088110156112925760051b019056fea26469706673582212206d0da7c314ac58d1ed084ff14a9cb0ea3d316660c387faaf0ac0d14d6db8807564736f6c63430008190033",
}

// CryptographyPrimitivesTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use CryptographyPrimitivesTesterMetaData.ABI instead.
var CryptographyPrimitivesTesterABI = CryptographyPrimitivesTesterMetaData.ABI

// CryptographyPrimitivesTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CryptographyPrimitivesTesterMetaData.Bin instead.
var CryptographyPrimitivesTesterBin = CryptographyPrimitivesTesterMetaData.Bin

// DeployCryptographyPrimitivesTester deploys a new Ethereum contract, binding an instance of CryptographyPrimitivesTester to it.
func DeployCryptographyPrimitivesTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CryptographyPrimitivesTester, error) {
	parsed, err := CryptographyPrimitivesTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CryptographyPrimitivesTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CryptographyPrimitivesTester{CryptographyPrimitivesTesterCaller: CryptographyPrimitivesTesterCaller{contract: contract}, CryptographyPrimitivesTesterTransactor: CryptographyPrimitivesTesterTransactor{contract: contract}, CryptographyPrimitivesTesterFilterer: CryptographyPrimitivesTesterFilterer{contract: contract}}, nil
}

// CryptographyPrimitivesTester is an auto generated Go binding around an Ethereum contract.
type CryptographyPrimitivesTester struct {
	CryptographyPrimitivesTesterCaller     // Read-only binding to the contract
	CryptographyPrimitivesTesterTransactor // Write-only binding to the contract
	CryptographyPrimitivesTesterFilterer   // Log filterer for contract events
}

// CryptographyPrimitivesTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptographyPrimitivesTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptographyPrimitivesTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptographyPrimitivesTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptographyPrimitivesTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptographyPrimitivesTesterSession struct {
	Contract     *CryptographyPrimitivesTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CryptographyPrimitivesTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptographyPrimitivesTesterCallerSession struct {
	Contract *CryptographyPrimitivesTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// CryptographyPrimitivesTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptographyPrimitivesTesterTransactorSession struct {
	Contract     *CryptographyPrimitivesTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// CryptographyPrimitivesTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterRaw struct {
	Contract *CryptographyPrimitivesTester // Generic contract binding to access the raw methods on
}

// CryptographyPrimitivesTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterCallerRaw struct {
	Contract *CryptographyPrimitivesTesterCaller // Generic read-only contract binding to access the raw methods on
}

// CryptographyPrimitivesTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptographyPrimitivesTesterTransactorRaw struct {
	Contract *CryptographyPrimitivesTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptographyPrimitivesTester creates a new instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTester(address common.Address, backend bind.ContractBackend) (*CryptographyPrimitivesTester, error) {
	contract, err := bindCryptographyPrimitivesTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTester{CryptographyPrimitivesTesterCaller: CryptographyPrimitivesTesterCaller{contract: contract}, CryptographyPrimitivesTesterTransactor: CryptographyPrimitivesTesterTransactor{contract: contract}, CryptographyPrimitivesTesterFilterer: CryptographyPrimitivesTesterFilterer{contract: contract}}, nil
}

// NewCryptographyPrimitivesTesterCaller creates a new read-only instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTesterCaller(address common.Address, caller bind.ContractCaller) (*CryptographyPrimitivesTesterCaller, error) {
	contract, err := bindCryptographyPrimitivesTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTesterCaller{contract: contract}, nil
}

// NewCryptographyPrimitivesTesterTransactor creates a new write-only instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptographyPrimitivesTesterTransactor, error) {
	contract, err := bindCryptographyPrimitivesTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTesterTransactor{contract: contract}, nil
}

// NewCryptographyPrimitivesTesterFilterer creates a new log filterer instance of CryptographyPrimitivesTester, bound to a specific deployed contract.
func NewCryptographyPrimitivesTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptographyPrimitivesTesterFilterer, error) {
	contract, err := bindCryptographyPrimitivesTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptographyPrimitivesTesterFilterer{contract: contract}, nil
}

// bindCryptographyPrimitivesTester binds a generic wrapper to an already deployed contract.
func bindCryptographyPrimitivesTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CryptographyPrimitivesTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptographyPrimitivesTester.Contract.CryptographyPrimitivesTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.CryptographyPrimitivesTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.CryptographyPrimitivesTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptographyPrimitivesTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptographyPrimitivesTester.Contract.contract.Transact(opts, method, params...)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCaller) KeccakF(opts *bind.CallOpts, input [25]*big.Int) ([25]*big.Int, error) {
	var out []interface{}
	err := _CryptographyPrimitivesTester.contract.Call(opts, &out, "keccakF", input)

	if err != nil {
		return *new([25]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([25]*big.Int)).(*[25]*big.Int)

	return out0, err

}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _CryptographyPrimitivesTester.Contract.KeccakF(&_CryptographyPrimitivesTester.CallOpts, input)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCallerSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _CryptographyPrimitivesTester.Contract.KeccakF(&_CryptographyPrimitivesTester.CallOpts, input)
}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCaller) Sha256Block(opts *bind.CallOpts, inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CryptographyPrimitivesTester.contract.Call(opts, &out, "sha256Block", inputChunk, hashState)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterSession) Sha256Block(inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	return _CryptographyPrimitivesTester.Contract.Sha256Block(&_CryptographyPrimitivesTester.CallOpts, inputChunk, hashState)
}

// Sha256Block is a free data retrieval call binding the contract method 0xe479f532.
//
// Solidity: function sha256Block(bytes32[2] inputChunk, bytes32 hashState) pure returns(bytes32)
func (_CryptographyPrimitivesTester *CryptographyPrimitivesTesterCallerSession) Sha256Block(inputChunk [2][32]byte, hashState [32]byte) ([32]byte, error) {
	return _CryptographyPrimitivesTester.Contract.Sha256Block(&_CryptographyPrimitivesTester.CallOpts, inputChunk, hashState)
}

// EthVaultMetaData contains all meta data concerning the EthVault contract.
var EthVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"justRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346017576000805560dd9081601d8239f35b600080fdfe6080806040526004361015601257600080fd5b600090813560e01c908163408def1e14609357816350b23fd214605b57506354fd4d5014603e57600080fd5b346058578060031936011260585760209054604051908152f35b80fd5b905081600319360112608f5762461bcd60e51b815260206004820152600360248201526262796560e81b6044820152606490fd5b5080fd5b826020366003190112605857600435815580f3fea2646970667358221220edaf9342bff8b461a8ffc9df08b776eaedd11060ec3af1d4b569907e67f0969864736f6c63430008190033",
}

// EthVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use EthVaultMetaData.ABI instead.
var EthVaultABI = EthVaultMetaData.ABI

// EthVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthVaultMetaData.Bin instead.
var EthVaultBin = EthVaultMetaData.Bin

// DeployEthVault deploys a new Ethereum contract, binding an instance of EthVault to it.
func DeployEthVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthVault, error) {
	parsed, err := EthVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthVault{EthVaultCaller: EthVaultCaller{contract: contract}, EthVaultTransactor: EthVaultTransactor{contract: contract}, EthVaultFilterer: EthVaultFilterer{contract: contract}}, nil
}

// EthVault is an auto generated Go binding around an Ethereum contract.
type EthVault struct {
	EthVaultCaller     // Read-only binding to the contract
	EthVaultTransactor // Write-only binding to the contract
	EthVaultFilterer   // Log filterer for contract events
}

// EthVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthVaultSession struct {
	Contract     *EthVault         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthVaultCallerSession struct {
	Contract *EthVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthVaultTransactorSession struct {
	Contract     *EthVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthVaultRaw struct {
	Contract *EthVault // Generic contract binding to access the raw methods on
}

// EthVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthVaultCallerRaw struct {
	Contract *EthVaultCaller // Generic read-only contract binding to access the raw methods on
}

// EthVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthVaultTransactorRaw struct {
	Contract *EthVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthVault creates a new instance of EthVault, bound to a specific deployed contract.
func NewEthVault(address common.Address, backend bind.ContractBackend) (*EthVault, error) {
	contract, err := bindEthVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthVault{EthVaultCaller: EthVaultCaller{contract: contract}, EthVaultTransactor: EthVaultTransactor{contract: contract}, EthVaultFilterer: EthVaultFilterer{contract: contract}}, nil
}

// NewEthVaultCaller creates a new read-only instance of EthVault, bound to a specific deployed contract.
func NewEthVaultCaller(address common.Address, caller bind.ContractCaller) (*EthVaultCaller, error) {
	contract, err := bindEthVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthVaultCaller{contract: contract}, nil
}

// NewEthVaultTransactor creates a new write-only instance of EthVault, bound to a specific deployed contract.
func NewEthVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EthVaultTransactor, error) {
	contract, err := bindEthVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthVaultTransactor{contract: contract}, nil
}

// NewEthVaultFilterer creates a new log filterer instance of EthVault, bound to a specific deployed contract.
func NewEthVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*EthVaultFilterer, error) {
	contract, err := bindEthVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthVaultFilterer{contract: contract}, nil
}

// bindEthVault binds a generic wrapper to an already deployed contract.
func bindEthVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthVault *EthVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthVault.Contract.EthVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthVault *EthVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthVault.Contract.EthVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthVault *EthVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthVault.Contract.EthVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthVault *EthVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthVault *EthVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthVault *EthVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthVault.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthVault *EthVaultCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthVault *EthVaultSession) Version() (*big.Int, error) {
	return _EthVault.Contract.Version(&_EthVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthVault *EthVaultCallerSession) Version() (*big.Int, error) {
	return _EthVault.Contract.Version(&_EthVault.CallOpts)
}

// JustRevert is a paid mutator transaction binding the contract method 0x50b23fd2.
//
// Solidity: function justRevert() payable returns()
func (_EthVault *EthVaultTransactor) JustRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthVault.contract.Transact(opts, "justRevert")
}

// JustRevert is a paid mutator transaction binding the contract method 0x50b23fd2.
//
// Solidity: function justRevert() payable returns()
func (_EthVault *EthVaultSession) JustRevert() (*types.Transaction, error) {
	return _EthVault.Contract.JustRevert(&_EthVault.TransactOpts)
}

// JustRevert is a paid mutator transaction binding the contract method 0x50b23fd2.
//
// Solidity: function justRevert() payable returns()
func (_EthVault *EthVaultTransactorSession) JustRevert() (*types.Transaction, error) {
	return _EthVault.Contract.JustRevert(&_EthVault.TransactOpts)
}

// SetVersion is a paid mutator transaction binding the contract method 0x408def1e.
//
// Solidity: function setVersion(uint256 _version) payable returns()
func (_EthVault *EthVaultTransactor) SetVersion(opts *bind.TransactOpts, _version *big.Int) (*types.Transaction, error) {
	return _EthVault.contract.Transact(opts, "setVersion", _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x408def1e.
//
// Solidity: function setVersion(uint256 _version) payable returns()
func (_EthVault *EthVaultSession) SetVersion(_version *big.Int) (*types.Transaction, error) {
	return _EthVault.Contract.SetVersion(&_EthVault.TransactOpts, _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x408def1e.
//
// Solidity: function setVersion(uint256 _version) payable returns()
func (_EthVault *EthVaultTransactorSession) SetVersion(_version *big.Int) (*types.Transaction, error) {
	return _EthVault.Contract.SetVersion(&_EthVault.TransactOpts, _version)
}

// MessageTesterMetaData contains all meta data concerning the MessageTester contract.
var MessageTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inbox\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"}],\"name\":\"accumulateInboxMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"messageType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"inboxSeqNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceL1\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576101b5908161001b8239f35b600080fdfe6080604090808252600436101561001557600080fd5b600090813560e01c9081638f3c79c014610132575063bf0090521461003957600080fd5b3461012f5760e036600319011261012f5760043560ff8116810361012b576024356001600160a01b0381168103610127576001600160401b0392604435848116810361012b57606435858116810361012757865160f89590951b6001600160f81b0319166020860190815260609490941b6001600160601b031916602186015260c091821b6001600160c01b0319908116603587015290821b16603d850152608435604585015260a435606585015260c4356085808601919091528452830193919291841182851017610113575082602094525190208152f35b634e487b7160e01b81526041600452602490fd5b8280fd5b5080fd5b80fd5b90503461012b578260031936011261012b57600435602082019081526024358483015283825260608201929091906001600160401b03841182851017610113575082602094525190208152f3fea2646970667358221220a25d17c9ba7dc102b724643faf33c69d1c40f7c92f5350837370cb0ce3160b5964736f6c63430008190033",
}

// MessageTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageTesterMetaData.ABI instead.
var MessageTesterABI = MessageTesterMetaData.ABI

// MessageTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageTesterMetaData.Bin instead.
var MessageTesterBin = MessageTesterMetaData.Bin

// DeployMessageTester deploys a new Ethereum contract, binding an instance of MessageTester to it.
func DeployMessageTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageTester, error) {
	parsed, err := MessageTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// MessageTester is an auto generated Go binding around an Ethereum contract.
type MessageTester struct {
	MessageTesterCaller     // Read-only binding to the contract
	MessageTesterTransactor // Write-only binding to the contract
	MessageTesterFilterer   // Log filterer for contract events
}

// MessageTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageTesterSession struct {
	Contract     *MessageTester    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageTesterCallerSession struct {
	Contract *MessageTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MessageTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTesterTransactorSession struct {
	Contract     *MessageTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MessageTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageTesterRaw struct {
	Contract *MessageTester // Generic contract binding to access the raw methods on
}

// MessageTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageTesterCallerRaw struct {
	Contract *MessageTesterCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTesterTransactorRaw struct {
	Contract *MessageTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageTester creates a new instance of MessageTester, bound to a specific deployed contract.
func NewMessageTester(address common.Address, backend bind.ContractBackend) (*MessageTester, error) {
	contract, err := bindMessageTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageTester{MessageTesterCaller: MessageTesterCaller{contract: contract}, MessageTesterTransactor: MessageTesterTransactor{contract: contract}, MessageTesterFilterer: MessageTesterFilterer{contract: contract}}, nil
}

// NewMessageTesterCaller creates a new read-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterCaller(address common.Address, caller bind.ContractCaller) (*MessageTesterCaller, error) {
	contract, err := bindMessageTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterCaller{contract: contract}, nil
}

// NewMessageTesterTransactor creates a new write-only instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTesterTransactor, error) {
	contract, err := bindMessageTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTesterTransactor{contract: contract}, nil
}

// NewMessageTesterFilterer creates a new log filterer instance of MessageTester, bound to a specific deployed contract.
func NewMessageTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageTesterFilterer, error) {
	contract, err := bindMessageTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageTesterFilterer{contract: contract}, nil
}

// bindMessageTester binds a generic wrapper to an already deployed contract.
func bindMessageTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.MessageTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.MessageTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTester *MessageTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTester *MessageTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTester *MessageTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTester.Contract.contract.Transact(opts, method, params...)
}

// AccumulateInboxMessage is a free data retrieval call binding the contract method 0x8f3c79c0.
//
// Solidity: function accumulateInboxMessage(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) AccumulateInboxMessage(opts *bind.CallOpts, inbox [32]byte, message [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "accumulateInboxMessage", inbox, message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AccumulateInboxMessage is a free data retrieval call binding the contract method 0x8f3c79c0.
//
// Solidity: function accumulateInboxMessage(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) AccumulateInboxMessage(inbox [32]byte, message [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.AccumulateInboxMessage(&_MessageTester.CallOpts, inbox, message)
}

// AccumulateInboxMessage is a free data retrieval call binding the contract method 0x8f3c79c0.
//
// Solidity: function accumulateInboxMessage(bytes32 inbox, bytes32 message) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) AccumulateInboxMessage(inbox [32]byte, message [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.AccumulateInboxMessage(&_MessageTester.CallOpts, inbox, message)
}

// MessageHash is a free data retrieval call binding the contract method 0xbf009052.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint64 blockNumber, uint64 timestamp, uint256 inboxSeqNum, uint256 gasPriceL1, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterCaller) MessageHash(opts *bind.CallOpts, messageType uint8, sender common.Address, blockNumber uint64, timestamp uint64, inboxSeqNum *big.Int, gasPriceL1 *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageTester.contract.Call(opts, &out, "messageHash", messageType, sender, blockNumber, timestamp, inboxSeqNum, gasPriceL1, messageDataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageHash is a free data retrieval call binding the contract method 0xbf009052.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint64 blockNumber, uint64 timestamp, uint256 inboxSeqNum, uint256 gasPriceL1, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterSession) MessageHash(messageType uint8, sender common.Address, blockNumber uint64, timestamp uint64, inboxSeqNum *big.Int, gasPriceL1 *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageHash(&_MessageTester.CallOpts, messageType, sender, blockNumber, timestamp, inboxSeqNum, gasPriceL1, messageDataHash)
}

// MessageHash is a free data retrieval call binding the contract method 0xbf009052.
//
// Solidity: function messageHash(uint8 messageType, address sender, uint64 blockNumber, uint64 timestamp, uint256 inboxSeqNum, uint256 gasPriceL1, bytes32 messageDataHash) pure returns(bytes32)
func (_MessageTester *MessageTesterCallerSession) MessageHash(messageType uint8, sender common.Address, blockNumber uint64, timestamp uint64, inboxSeqNum *big.Int, gasPriceL1 *big.Int, messageDataHash [32]byte) ([32]byte, error) {
	return _MessageTester.Contract.MessageHash(&_MessageTester.CallOpts, messageType, sender, blockNumber, timestamp, inboxSeqNum, gasPriceL1, messageDataHash)
}

// OutboxWithoutOptTesterMetaData contains all meta data concerning the OutboxWithoutOptTester contract.
var OutboxWithoutOptTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AlreadySpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeCallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProofLength\",\"type\":\"uint256\"}],\"name\":\"MerkleProofTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIndex\",\"type\":\"uint256\"}],\"name\":\"PathNotMinimal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proofLength\",\"type\":\"uint256\"}],\"name\":\"ProofTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"UnknownRoot\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"zero\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionIndex\",\"type\":\"uint256\"}],\"name\":\"OutBoxTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2BlockHash\",\"type\":\"bytes32\"}],\"name\":\"SendRootUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OUTBOX_VERSION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"calculateItemHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"item\",\"type\":\"bytes32\"}],\"name\":\"calculateMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeTransactionSimulation\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1BatchNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1EthBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1OutputId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRollupAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2BlockHash\",\"type\":\"bytes32\"}],\"name\":\"updateSendRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0806040523460225730608052610ef9908161002882396080518161048d0152f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80627436d31461015657806308635a9514610151578063119852711461014c578063288e5b101461014757806346547790146101425780635a129efe1461013d5780636ae71f121461013857806372f2a8c71461013357806380648b021461012e5780638515bc6a1461012957806395fcea78146101245780639f0c04bf1461011f578063a04cee601461011a578063ae6dead714610115578063b0f3053714610110578063c4d66de81461010b578063c75184df14610106578063cb23bcb514610101578063d5b5cc23146100fc5763e78cea92146100f757600080fd5b6108ab565b610864565b61083b565b61081f565b610753565b61072a565b6106fe565b6106af565b610648565b610635565b610614565b6105eb565b6105cd565b610470565b61043f565b610416565b610396565b61037a565b6102e3565b61022e565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b0382111761018c57604052565b61015b565b601f909101601f19168101906001600160401b0382119082101761018c57604052565b6040519060a082016001600160401b0381118382101761018c57604052565b9092916001600160401b03841161018c578360051b60209260206040516101fc82850182610191565b809781520191810192831161022957905b82821061021a5750505050565b8135815290830190830161020d565b600080fd5b34610229576060366003190112610229576004356001600160401b03811161022957366023820112156102295761028161027460209236906024816004013591016101d3565b60443590602435906108d4565b604051908152f35b6001600160a01b0381160361022957565b604435906102a782610289565b565b606435906102a782610289565b9181601f84011215610229578235916001600160401b038311610229576020838186019501011161022957565b3461022957610120366003190112610229576001600160401b036004358181116102295736602382011215610229578060040135828111610229573660248260051b840101116102295761033561029a565b61033d6102a9565b90610104359485116102295761035a6103789536906004016102b6565b94909360e4359360c4359360a4359360843593602480359201610a2d565b005b3461022957600036600319011261022957602060405160008152f35b3461022957610100366003190112610229576103b3602435610289565b6103be604435610289565b60e4356001600160401b038111610229576103dd9036906004016102b6565b505060405162461bcd60e51b815260206004820152600f60248201526e139bdd081a5b5c1b195b595b9d1959608a1b6044820152606490fd5b34610229576000366003190112610229576004546040516001600160801b039091168152602090f35b34610229576020366003190112610229576004356000526002602052602060ff604060002054166040519015158152f35b34610229576000806003193601126105ca576001600160a01b03307f000000000000000000000000000000000000000000000000000000000000000082161461057057807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103541680330361055257506004602082600154166040519283809263cb23bcb560e01b82525afa90811561054d57839161051e575b5082546001600160a01b031916911617815580f35b610540915060203d602011610546575b6105388183610191565b810190610c90565b38610509565b503d61052e565b610ca8565b60449060405190631194af8760e11b82523360048301526024820152fd5b60405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608490fd5b80fd5b34610229576000366003190112610229576020600654604051908152f35b34610229576000366003190112610229576007546040516001600160a01b039091168152602090f35b3461022957600036600319011261022957602060045460801c604051908152f35b3461022957600036600319011261022957005b346102295760e03660031901126102295760043561066581610289565b60243561067181610289565b60c435906001600160401b038211610229576020926106976102819336906004016102b6565b92909160a43591608435916064359160443591610cb4565b346102295760403660031901126102295760043560243560009180835260036020528160408420557fb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f67488380a380f35b346102295760203660031901126102295760043560005260036020526020604060002054604051908152f35b34610229576000366003190112610229576005546040516001600160801b039091168152602090f35b346102295760203660031901126102295760043561077081610289565b6001546001600160a01b039190821661080d57600180546001600160a01b0319166001600160a01b038316179055819060209060046040518095819363cb23bcb560e01b8352165afa801561054d57610378926000916107ee575b5060008054919092166001600160a01b03166001600160a01b0319909116179055565b610807915060203d602011610546576105388183610191565b386107cb565b604051633bcd329760e21b8152600490fd5b3461022957600036600319011261022957602060405160028152f35b34610229576000366003190112610229576000546040516001600160a01b039091168152602090f35b346102295760203660031901126102295760405162461bcd60e51b815260206004820152600e60248201526d1393d517d253541311535155115160921b6044820152606490fd5b34610229576000366003190112610229576001546040516001600160a01b039091168152602090f35b91604080519160209260208101918252602081526108f181610171565b5190209184519261010080851161095f575094939291906000945b83861061091c5750505050505090565b90919293949561092c8783610d1c565b5160019182891b881661094f57600052835283600020965b01949392919061090c565b9060005283528360002096610944565b8460449160405191637ed6198f60e11b835260048301526024820152fd5b6040519060a082016001600160401b0381118382101761018c5760409081526004546001600160801b038082168552608091821c6020860152600554169184019190915260065460608401526007546001600160a01b031690830152565b6001600160401b03811161018c57601f01601f191660200190565b929192610a02826109db565b91610a106040519384610191565b829481845281830111610229578281602093846000960137010152565b610a5490610a4c8c8c8a969f9e99989c8c9e9b9c819d8f8c908c610cb4565b9236916101d3565b8051610100811015610c76575080610a6d8d9251610d46565b821015610c4d5790610a7f92916108d4565b610a93816000526003602052604060002090565b5415610c335750610ab8610ab18b6000526002602052604060002090565b5460ff1690565b610c1a57610ba596959492610ba5610c1495938c610b90610c0d956102a79e9f610aef610afc916000526002602052604060002090565b805460ff19166001179055565b6040518381526000906001600160a01b0386811691908f16907f20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab1896490602090a4610b4361097d565b9d6001600160801b0391610b7f908390610b6e82610b5f6101b4565b9b166001600160801b03168b52565b166001600160801b03166020890152565b166001600160801b03166040860152565b60608401526001600160a01b03166080830152565b80516020820151608090811b6001600160801b03199081166001600160801b039384161760045560408401516005805490921693169290921790915560608201516006550151600780546001600160a01b0319166001600160a01b0392909216919091179055565b36916109f6565b91610dfd565b604051639715b8d360e01b8152600481018b9052602490fd5b6040516310e61af960e31b81526004810191909152602490fd5b610c579051610d46565b604051630b8a724b60e01b815260048101929092526024820152604490fd5b60405163ab6a068360e01b81526004810191909152602490fd5b908160209103126102295751610ca581610289565b90565b6040513d6000823e3d90fd5b9295610d169560c895899498999399604051998a97602089019c8d60018060601b0319809360601b16905260601b16603489015260488801526068870152608886015260a88501528484013781016000838201520360a8810184520182610191565b51902090565b8051821015610d305760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b60ff8111610d55576001901b90565b634e487b7160e01b600052601160045260246000fd5b60005b838110610d7e5750506000910152565b8181015183820152602001610d6e565b9190604083820312610229578251801515810361022957602084015190936001600160401b038211610229570181601f82011215610229578051610dd1816109db565b92610ddf6040519485610191565b8184526020828401011161022957610ca59160208085019101610d6b565b9060009160018060a01b03836080826001541693600460405198899788968794639e5d4c4960e01b86521683850152602484015260606044840152610e52815180928160648701526020608487019101610d6b565b601f01601f19168201010301925af190811561054d57600090600092610e9d575b5015610e7c5750565b805115610e8b57805190602001fd5b604051631bb7daad60e11b8152600490fd5b9050610ebc91503d806000833e610eb48183610191565b810190610d8e565b9038610e7356fea2646970667358221220e820ddde6396ab5eb14c0b4903b42900187561dbe0bdbebd80f5382972a6b67f64736f6c63430008190033",
}

// OutboxWithoutOptTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use OutboxWithoutOptTesterMetaData.ABI instead.
var OutboxWithoutOptTesterABI = OutboxWithoutOptTesterMetaData.ABI

// OutboxWithoutOptTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OutboxWithoutOptTesterMetaData.Bin instead.
var OutboxWithoutOptTesterBin = OutboxWithoutOptTesterMetaData.Bin

// DeployOutboxWithoutOptTester deploys a new Ethereum contract, binding an instance of OutboxWithoutOptTester to it.
func DeployOutboxWithoutOptTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OutboxWithoutOptTester, error) {
	parsed, err := OutboxWithoutOptTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OutboxWithoutOptTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutboxWithoutOptTester{OutboxWithoutOptTesterCaller: OutboxWithoutOptTesterCaller{contract: contract}, OutboxWithoutOptTesterTransactor: OutboxWithoutOptTesterTransactor{contract: contract}, OutboxWithoutOptTesterFilterer: OutboxWithoutOptTesterFilterer{contract: contract}}, nil
}

// OutboxWithoutOptTester is an auto generated Go binding around an Ethereum contract.
type OutboxWithoutOptTester struct {
	OutboxWithoutOptTesterCaller     // Read-only binding to the contract
	OutboxWithoutOptTesterTransactor // Write-only binding to the contract
	OutboxWithoutOptTesterFilterer   // Log filterer for contract events
}

// OutboxWithoutOptTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxWithoutOptTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxWithoutOptTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxWithoutOptTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxWithoutOptTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxWithoutOptTesterSession struct {
	Contract     *OutboxWithoutOptTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OutboxWithoutOptTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxWithoutOptTesterCallerSession struct {
	Contract *OutboxWithoutOptTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// OutboxWithoutOptTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxWithoutOptTesterTransactorSession struct {
	Contract     *OutboxWithoutOptTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// OutboxWithoutOptTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxWithoutOptTesterRaw struct {
	Contract *OutboxWithoutOptTester // Generic contract binding to access the raw methods on
}

// OutboxWithoutOptTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterCallerRaw struct {
	Contract *OutboxWithoutOptTesterCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxWithoutOptTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxWithoutOptTesterTransactorRaw struct {
	Contract *OutboxWithoutOptTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutboxWithoutOptTester creates a new instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTester(address common.Address, backend bind.ContractBackend) (*OutboxWithoutOptTester, error) {
	contract, err := bindOutboxWithoutOptTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTester{OutboxWithoutOptTesterCaller: OutboxWithoutOptTesterCaller{contract: contract}, OutboxWithoutOptTesterTransactor: OutboxWithoutOptTesterTransactor{contract: contract}, OutboxWithoutOptTesterFilterer: OutboxWithoutOptTesterFilterer{contract: contract}}, nil
}

// NewOutboxWithoutOptTesterCaller creates a new read-only instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTesterCaller(address common.Address, caller bind.ContractCaller) (*OutboxWithoutOptTesterCaller, error) {
	contract, err := bindOutboxWithoutOptTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterCaller{contract: contract}, nil
}

// NewOutboxWithoutOptTesterTransactor creates a new write-only instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxWithoutOptTesterTransactor, error) {
	contract, err := bindOutboxWithoutOptTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterTransactor{contract: contract}, nil
}

// NewOutboxWithoutOptTesterFilterer creates a new log filterer instance of OutboxWithoutOptTester, bound to a specific deployed contract.
func NewOutboxWithoutOptTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxWithoutOptTesterFilterer, error) {
	contract, err := bindOutboxWithoutOptTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterFilterer{contract: contract}, nil
}

// bindOutboxWithoutOptTester binds a generic wrapper to an already deployed contract.
func bindOutboxWithoutOptTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OutboxWithoutOptTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxWithoutOptTester.Contract.OutboxWithoutOptTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.OutboxWithoutOptTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.OutboxWithoutOptTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutboxWithoutOptTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.contract.Transact(opts, method, params...)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) OUTBOXVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "OUTBOX_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) OUTBOXVERSION() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.OUTBOXVERSION(&_OutboxWithoutOptTester.CallOpts)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) OUTBOXVERSION() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.OUTBOXVERSION(&_OutboxWithoutOptTester.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Bridge() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Bridge(&_OutboxWithoutOptTester.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Bridge() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Bridge(&_OutboxWithoutOptTester.CallOpts)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) CalculateItemHash(opts *bind.CallOpts, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "calculateItemHash", l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateItemHash(&_OutboxWithoutOptTester.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateItemHash(&_OutboxWithoutOptTester.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) CalculateMerkleRoot(opts *bind.CallOpts, proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "calculateMerkleRoot", proof, path, item)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateMerkleRoot(&_OutboxWithoutOptTester.CallOpts, proof, path, item)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.CalculateMerkleRoot(&_OutboxWithoutOptTester.CallOpts, proof, path, item)
}

// ExecuteTransactionSimulation is a free data retrieval call binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 , address , address , uint256 , uint256 , uint256 , uint256 , bytes ) pure returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) ExecuteTransactionSimulation(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte) error {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "executeTransactionSimulation", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	if err != nil {
		return err
	}

	return err

}

// ExecuteTransactionSimulation is a free data retrieval call binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 , address , address , uint256 , uint256 , uint256 , uint256 , bytes ) pure returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) ExecuteTransactionSimulation(arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte) error {
	return _OutboxWithoutOptTester.Contract.ExecuteTransactionSimulation(&_OutboxWithoutOptTester.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// ExecuteTransactionSimulation is a free data retrieval call binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 , address , address , uint256 , uint256 , uint256 , uint256 , bytes ) pure returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) ExecuteTransactionSimulation(arg0 *big.Int, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 []byte) error {
	return _OutboxWithoutOptTester.Contract.ExecuteTransactionSimulation(&_OutboxWithoutOptTester.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) IsSpent(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "isSpent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) IsSpent(arg0 *big.Int) (bool, error) {
	return _OutboxWithoutOptTester.Contract.IsSpent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) IsSpent(arg0 *big.Int) (bool, error) {
	return _OutboxWithoutOptTester.Contract.IsSpent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1BatchNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1BatchNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1BatchNum() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1BatchNum(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1BatchNum() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1BatchNum(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1Block() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Block(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1Block() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Block(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1EthBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1EthBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1EthBlock() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1EthBlock(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1EthBlock() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1EthBlock(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1OutputId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1OutputId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1OutputId() ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1OutputId(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1OutputId() ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1OutputId(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1Sender() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Sender(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1Sender() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Sender(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) L2ToL1Timestamp() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Timestamp(&_OutboxWithoutOptTester.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _OutboxWithoutOptTester.Contract.L2ToL1Timestamp(&_OutboxWithoutOptTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Rollup() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Rollup(&_OutboxWithoutOptTester.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Rollup() (common.Address, error) {
	return _OutboxWithoutOptTester.Contract.Rollup(&_OutboxWithoutOptTester.CallOpts)
}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Roots(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Roots(arg0 [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Roots(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Roots(arg0 [32]byte) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Roots(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCaller) Spent(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _OutboxWithoutOptTester.contract.Call(opts, &out, "spent", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Spent(arg0 *big.Int) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Spent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) pure returns(bytes32)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterCallerSession) Spent(arg0 *big.Int) ([32]byte, error) {
	return _OutboxWithoutOptTester.Contract.Spent(&_OutboxWithoutOptTester.CallOpts, arg0)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) ExecuteTransaction(opts *bind.TransactOpts, proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "executeTransaction", proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.ExecuteTransaction(&_OutboxWithoutOptTester.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.ExecuteTransaction(&_OutboxWithoutOptTester.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "initialize", _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.Initialize(&_OutboxWithoutOptTester.TransactOpts, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.Initialize(&_OutboxWithoutOptTester.TransactOpts, _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) PostUpgradeInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "postUpgradeInit")
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) PostUpgradeInit() (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.PostUpgradeInit(&_OutboxWithoutOptTester.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) PostUpgradeInit() (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.PostUpgradeInit(&_OutboxWithoutOptTester.TransactOpts)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) UpdateRollupAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "updateRollupAddress")
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateRollupAddress(&_OutboxWithoutOptTester.TransactOpts)
}

// UpdateRollupAddress is a paid mutator transaction binding the contract method 0x6ae71f12.
//
// Solidity: function updateRollupAddress() returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) UpdateRollupAddress() (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateRollupAddress(&_OutboxWithoutOptTester.TransactOpts)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactor) UpdateSendRoot(opts *bind.TransactOpts, root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.contract.Transact(opts, "updateSendRoot", root, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterSession) UpdateSendRoot(root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateSendRoot(&_OutboxWithoutOptTester.TransactOpts, root, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterTransactorSession) UpdateSendRoot(root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _OutboxWithoutOptTester.Contract.UpdateSendRoot(&_OutboxWithoutOptTester.TransactOpts, root, l2BlockHash)
}

// OutboxWithoutOptTesterOutBoxTransactionExecutedIterator is returned from FilterOutBoxTransactionExecuted and is used to iterate over the raw logs and unpacked data for OutBoxTransactionExecuted events raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterOutBoxTransactionExecutedIterator struct {
	Event *OutboxWithoutOptTesterOutBoxTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *OutboxWithoutOptTesterOutBoxTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
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
		it.Event = new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
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
func (it *OutboxWithoutOptTesterOutBoxTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxWithoutOptTesterOutBoxTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxWithoutOptTesterOutBoxTransactionExecuted represents a OutBoxTransactionExecuted event raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterOutBoxTransactionExecuted struct {
	To               common.Address
	L2Sender         common.Address
	Zero             *big.Int
	TransactionIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOutBoxTransactionExecuted is a free log retrieval operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) FilterOutBoxTransactionExecuted(opts *bind.FilterOpts, to []common.Address, l2Sender []common.Address, zero []*big.Int) (*OutboxWithoutOptTesterOutBoxTransactionExecutedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.FilterLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterOutBoxTransactionExecutedIterator{contract: _OutboxWithoutOptTester.contract, event: "OutBoxTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchOutBoxTransactionExecuted is a free log subscription operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) WatchOutBoxTransactionExecuted(opts *bind.WatchOpts, sink chan<- *OutboxWithoutOptTesterOutBoxTransactionExecuted, to []common.Address, l2Sender []common.Address, zero []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.WatchLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
				if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
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

// ParseOutBoxTransactionExecuted is a log parse operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) ParseOutBoxTransactionExecuted(log types.Log) (*OutboxWithoutOptTesterOutBoxTransactionExecuted, error) {
	event := new(OutboxWithoutOptTesterOutBoxTransactionExecuted)
	if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxWithoutOptTesterSendRootUpdatedIterator is returned from FilterSendRootUpdated and is used to iterate over the raw logs and unpacked data for SendRootUpdated events raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterSendRootUpdatedIterator struct {
	Event *OutboxWithoutOptTesterSendRootUpdated // Event containing the contract specifics and raw log

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
func (it *OutboxWithoutOptTesterSendRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxWithoutOptTesterSendRootUpdated)
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
		it.Event = new(OutboxWithoutOptTesterSendRootUpdated)
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
func (it *OutboxWithoutOptTesterSendRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxWithoutOptTesterSendRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxWithoutOptTesterSendRootUpdated represents a SendRootUpdated event raised by the OutboxWithoutOptTester contract.
type OutboxWithoutOptTesterSendRootUpdated struct {
	OutputRoot  [32]byte
	L2BlockHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSendRootUpdated is a free log retrieval operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed outputRoot, bytes32 indexed l2BlockHash)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) FilterSendRootUpdated(opts *bind.FilterOpts, outputRoot [][32]byte, l2BlockHash [][32]byte) (*OutboxWithoutOptTesterSendRootUpdatedIterator, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockHashRule []interface{}
	for _, l2BlockHashItem := range l2BlockHash {
		l2BlockHashRule = append(l2BlockHashRule, l2BlockHashItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.FilterLogs(opts, "SendRootUpdated", outputRootRule, l2BlockHashRule)
	if err != nil {
		return nil, err
	}
	return &OutboxWithoutOptTesterSendRootUpdatedIterator{contract: _OutboxWithoutOptTester.contract, event: "SendRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSendRootUpdated is a free log subscription operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed outputRoot, bytes32 indexed l2BlockHash)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) WatchSendRootUpdated(opts *bind.WatchOpts, sink chan<- *OutboxWithoutOptTesterSendRootUpdated, outputRoot [][32]byte, l2BlockHash [][32]byte) (event.Subscription, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2BlockHashRule []interface{}
	for _, l2BlockHashItem := range l2BlockHash {
		l2BlockHashRule = append(l2BlockHashRule, l2BlockHashItem)
	}

	logs, sub, err := _OutboxWithoutOptTester.contract.WatchLogs(opts, "SendRootUpdated", outputRootRule, l2BlockHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxWithoutOptTesterSendRootUpdated)
				if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
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

// ParseSendRootUpdated is a log parse operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed outputRoot, bytes32 indexed l2BlockHash)
func (_OutboxWithoutOptTester *OutboxWithoutOptTesterFilterer) ParseSendRootUpdated(log types.Log) (*OutboxWithoutOptTesterSendRootUpdated, error) {
	event := new(OutboxWithoutOptTesterSendRootUpdated)
	if err := _OutboxWithoutOptTester.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMockMetaData contains all meta data concerning the RollupMock contract.
var RollupMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"WithdrawTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ZombieTriggered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removeOldZombies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStakerFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608034607057601f61019d38819003918201601f19168301916001600160401b03831184841017607557808492602094604052833981010312607057516001600160a01b03811690819003607057600080546001600160a01b031916919091179055604051610111908161008c8239f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080806040526004361015601257600080fd5b600090813560e01c9081636137391914609c575080638da5cb5b1460775763edfd03ed14603e57600080fd5b3460745760203660031901126074577fb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d8180a180f35b80fd5b503460745780600319360112607457546040516001600160a01b039091168152602090f35b90503460d7578160031936011260d757817f1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f81602094a18152f35b5080fdfea26469706673582212202881fc8a414424712453f86914431a3fc28e9bdccfcff77ae7e1b06ac3c6ca0764736f6c63430008190033",
}

// RollupMockABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupMockMetaData.ABI instead.
var RollupMockABI = RollupMockMetaData.ABI

// RollupMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupMockMetaData.Bin instead.
var RollupMockBin = RollupMockMetaData.Bin

// DeployRollupMock deploys a new Ethereum contract, binding an instance of RollupMock to it.
func DeployRollupMock(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *RollupMock, error) {
	parsed, err := RollupMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupMockBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupMock{RollupMockCaller: RollupMockCaller{contract: contract}, RollupMockTransactor: RollupMockTransactor{contract: contract}, RollupMockFilterer: RollupMockFilterer{contract: contract}}, nil
}

// RollupMock is an auto generated Go binding around an Ethereum contract.
type RollupMock struct {
	RollupMockCaller     // Read-only binding to the contract
	RollupMockTransactor // Write-only binding to the contract
	RollupMockFilterer   // Log filterer for contract events
}

// RollupMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupMockSession struct {
	Contract     *RollupMock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupMockCallerSession struct {
	Contract *RollupMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RollupMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupMockTransactorSession struct {
	Contract     *RollupMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RollupMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupMockRaw struct {
	Contract *RollupMock // Generic contract binding to access the raw methods on
}

// RollupMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupMockCallerRaw struct {
	Contract *RollupMockCaller // Generic read-only contract binding to access the raw methods on
}

// RollupMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupMockTransactorRaw struct {
	Contract *RollupMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupMock creates a new instance of RollupMock, bound to a specific deployed contract.
func NewRollupMock(address common.Address, backend bind.ContractBackend) (*RollupMock, error) {
	contract, err := bindRollupMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupMock{RollupMockCaller: RollupMockCaller{contract: contract}, RollupMockTransactor: RollupMockTransactor{contract: contract}, RollupMockFilterer: RollupMockFilterer{contract: contract}}, nil
}

// NewRollupMockCaller creates a new read-only instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockCaller(address common.Address, caller bind.ContractCaller) (*RollupMockCaller, error) {
	contract, err := bindRollupMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMockCaller{contract: contract}, nil
}

// NewRollupMockTransactor creates a new write-only instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupMockTransactor, error) {
	contract, err := bindRollupMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupMockTransactor{contract: contract}, nil
}

// NewRollupMockFilterer creates a new log filterer instance of RollupMock, bound to a specific deployed contract.
func NewRollupMockFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupMockFilterer, error) {
	contract, err := bindRollupMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupMockFilterer{contract: contract}, nil
}

// bindRollupMock binds a generic wrapper to an already deployed contract.
func bindRollupMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMock *RollupMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupMock.Contract.RollupMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMock *RollupMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.Contract.RollupMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMock *RollupMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMock.Contract.RollupMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupMock *RollupMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupMock *RollupMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupMock *RollupMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupMock.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupMock *RollupMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupMock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupMock *RollupMockSession) Owner() (common.Address, error) {
	return _RollupMock.Contract.Owner(&_RollupMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupMock *RollupMockCallerSession) Owner() (common.Address, error) {
	return _RollupMock.Contract.Owner(&_RollupMock.CallOpts)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 ) returns()
func (_RollupMock *RollupMockTransactor) RemoveOldZombies(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _RollupMock.contract.Transact(opts, "removeOldZombies", arg0)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 ) returns()
func (_RollupMock *RollupMockSession) RemoveOldZombies(arg0 *big.Int) (*types.Transaction, error) {
	return _RollupMock.Contract.RemoveOldZombies(&_RollupMock.TransactOpts, arg0)
}

// RemoveOldZombies is a paid mutator transaction binding the contract method 0xedfd03ed.
//
// Solidity: function removeOldZombies(uint256 ) returns()
func (_RollupMock *RollupMockTransactorSession) RemoveOldZombies(arg0 *big.Int) (*types.Transaction, error) {
	return _RollupMock.Contract.RemoveOldZombies(&_RollupMock.TransactOpts, arg0)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x61373919.
//
// Solidity: function withdrawStakerFunds() returns(uint256)
func (_RollupMock *RollupMockTransactor) WithdrawStakerFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupMock.contract.Transact(opts, "withdrawStakerFunds")
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x61373919.
//
// Solidity: function withdrawStakerFunds() returns(uint256)
func (_RollupMock *RollupMockSession) WithdrawStakerFunds() (*types.Transaction, error) {
	return _RollupMock.Contract.WithdrawStakerFunds(&_RollupMock.TransactOpts)
}

// WithdrawStakerFunds is a paid mutator transaction binding the contract method 0x61373919.
//
// Solidity: function withdrawStakerFunds() returns(uint256)
func (_RollupMock *RollupMockTransactorSession) WithdrawStakerFunds() (*types.Transaction, error) {
	return _RollupMock.Contract.WithdrawStakerFunds(&_RollupMock.TransactOpts)
}

// RollupMockWithdrawTriggeredIterator is returned from FilterWithdrawTriggered and is used to iterate over the raw logs and unpacked data for WithdrawTriggered events raised by the RollupMock contract.
type RollupMockWithdrawTriggeredIterator struct {
	Event *RollupMockWithdrawTriggered // Event containing the contract specifics and raw log

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
func (it *RollupMockWithdrawTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMockWithdrawTriggered)
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
		it.Event = new(RollupMockWithdrawTriggered)
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
func (it *RollupMockWithdrawTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMockWithdrawTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMockWithdrawTriggered represents a WithdrawTriggered event raised by the RollupMock contract.
type RollupMockWithdrawTriggered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTriggered is a free log retrieval operation binding the contract event 0x1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f.
//
// Solidity: event WithdrawTriggered()
func (_RollupMock *RollupMockFilterer) FilterWithdrawTriggered(opts *bind.FilterOpts) (*RollupMockWithdrawTriggeredIterator, error) {

	logs, sub, err := _RollupMock.contract.FilterLogs(opts, "WithdrawTriggered")
	if err != nil {
		return nil, err
	}
	return &RollupMockWithdrawTriggeredIterator{contract: _RollupMock.contract, event: "WithdrawTriggered", logs: logs, sub: sub}, nil
}

// WatchWithdrawTriggered is a free log subscription operation binding the contract event 0x1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f.
//
// Solidity: event WithdrawTriggered()
func (_RollupMock *RollupMockFilterer) WatchWithdrawTriggered(opts *bind.WatchOpts, sink chan<- *RollupMockWithdrawTriggered) (event.Subscription, error) {

	logs, sub, err := _RollupMock.contract.WatchLogs(opts, "WithdrawTriggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMockWithdrawTriggered)
				if err := _RollupMock.contract.UnpackLog(event, "WithdrawTriggered", log); err != nil {
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

// ParseWithdrawTriggered is a log parse operation binding the contract event 0x1c09fbbf7cfd024f5e4e5472dd87afd5d67ee5db6a0ca715bf508d96abce309f.
//
// Solidity: event WithdrawTriggered()
func (_RollupMock *RollupMockFilterer) ParseWithdrawTriggered(log types.Log) (*RollupMockWithdrawTriggered, error) {
	event := new(RollupMockWithdrawTriggered)
	if err := _RollupMock.contract.UnpackLog(event, "WithdrawTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMockZombieTriggeredIterator is returned from FilterZombieTriggered and is used to iterate over the raw logs and unpacked data for ZombieTriggered events raised by the RollupMock contract.
type RollupMockZombieTriggeredIterator struct {
	Event *RollupMockZombieTriggered // Event containing the contract specifics and raw log

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
func (it *RollupMockZombieTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMockZombieTriggered)
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
		it.Event = new(RollupMockZombieTriggered)
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
func (it *RollupMockZombieTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMockZombieTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMockZombieTriggered represents a ZombieTriggered event raised by the RollupMock contract.
type RollupMockZombieTriggered struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterZombieTriggered is a free log retrieval operation binding the contract event 0xb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d.
//
// Solidity: event ZombieTriggered()
func (_RollupMock *RollupMockFilterer) FilterZombieTriggered(opts *bind.FilterOpts) (*RollupMockZombieTriggeredIterator, error) {

	logs, sub, err := _RollupMock.contract.FilterLogs(opts, "ZombieTriggered")
	if err != nil {
		return nil, err
	}
	return &RollupMockZombieTriggeredIterator{contract: _RollupMock.contract, event: "ZombieTriggered", logs: logs, sub: sub}, nil
}

// WatchZombieTriggered is a free log subscription operation binding the contract event 0xb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d.
//
// Solidity: event ZombieTriggered()
func (_RollupMock *RollupMockFilterer) WatchZombieTriggered(opts *bind.WatchOpts, sink chan<- *RollupMockZombieTriggered) (event.Subscription, error) {

	logs, sub, err := _RollupMock.contract.WatchLogs(opts, "ZombieTriggered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMockZombieTriggered)
				if err := _RollupMock.contract.UnpackLog(event, "ZombieTriggered", log); err != nil {
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

// ParseZombieTriggered is a log parse operation binding the contract event 0xb774f793432a37585a7638b9afe49e91c478887a2c0fef32877508bf2f76429d.
//
// Solidity: event ZombieTriggered()
func (_RollupMock *RollupMockFilterer) ParseZombieTriggered(log types.Log) (*RollupMockZombieTriggered, error) {
	event := new(RollupMockZombieTriggered)
	if err := _RollupMock.contract.UnpackLog(event, "ZombieTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestTokenMetaData contains all meta data concerning the TestToken contract.
var TestTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080346103a857601f19906001600160401b03601f610c3e3881900382810186168501848111868210176102b8578592829160405283396020948591810103126103a857519361004d6103ad565b9260098452682a32b9ba2a37b5b2b760b91b8585015261006b6103ad565b936002855261151560f21b868601528051908282116102b85760039182546001928382811c9216801561039e575b8a831014610388578188849311610337575b5089908883116001146102d9576000926102ce575b505060001982851b1c191690821b1782555b85519283116102b85760049586548281811c911680156102ae575b8982101461029957868111610253575b5087908685116001146101ef57849550908492916000956101e4575b50501b92600019911b1c19161782555b33156101a457506002549083820180921161018f57506000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91600255338352828152604083208481540190556040519384523393a360405161087190816103cd8239f35b601190634e487b7160e01b6000525260246000fd5b6064926040519262461bcd60e51b845283015260248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b015193503880610119565b9291948416928760005284896000209460005b8b8983831061023c5750505010610222575b50505050811b018255610129565b01519060f884600019921b161c1916905538808080610214565b868601518955909701969485019488935001610202565b87600052886000208780870160051c8201928b8810610290575b0160051c019083905b8281106102845750506100fd565b60008155018390610276565b9250819261026d565b602288634e487b7160e01b6000525260246000fd5b90607f16906100ed565b634e487b7160e01b600052604160045260246000fd5b0151905038806100c0565b908785941691866000528b6000209260005b8d8282106103215750508411610309575b505050811b0182556100d2565b015160001983871b60f8161c191690553880806102fc565b83850151865588979095019493840193016102eb565b90915084600052896000208880850160051c8201928c861061037f575b918691869594930160051c01915b8281106103705750506100ab565b60008155859450869101610362565b92508192610354565b634e487b7160e01b600052602260045260246000fd5b91607f1691610099565b600080fd5b60408051919082016001600160401b038111838210176102b85760405256fe608060408181526004918236101561001657600080fd5b600092833560e01c91826306fdde031461049157508163095ea7b31461046757816318160ddd1461044857816323b872dd1461037e578163313ce5671461036257816339509351146102fb57816370a08231146102c457816395d89b41146101c1578163a457c2d71461011957508063a9059cbb146100e95763dd62ed3e1461009e57600080fd5b346100e557806003193601126100e557806020926100ba61059a565b6100c26105b5565b6001600160a01b0391821683526001865283832091168252845220549051908152f35b5080fd5b50346100e557806003193601126100e55760209061011261010861059a565b60243590336106cd565b5160018152f35b905082346101be57826003193601126101be5761013461059a565b918360243592338152600160205281812060018060a01b038616825260205220549082821061016d5760208561011285850387336105cb565b608490602086519162461bcd60e51b8352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152fd5b80fd5b8383346100e557816003193601126100e55780519082845460018160011c90600183169283156102ba575b60209384841081146102a75783885290811561028b5750600114610253575b505050829003601f01601f19168201926001600160401b03841183851017610240575082918261023c925282610551565b0390f35b634e487b7160e01b815260418552602490fd5b919250868652828620918387935b838510610277575050505083010185808061020b565b805488860183015293019284908201610261565b60ff1916878501525050151560051b840101905085808061020b565b634e487b7160e01b895260228a52602489fd5b91607f16916101ec565b5050346100e55760203660031901126100e55760209181906001600160a01b036102ec61059a565b16815280845220549051908152f35b8284346101be57816003193601126101be5761031561059a565b338252600160209081528383206001600160a01b038316845290528282205460243581019290831061034f576020846101128585336105cb565b634e487b7160e01b815260118552602490fd5b5050346100e557816003193601126100e5576020905160128152f35b839150346100e55760603660031901126100e55761039a61059a565b6103a26105b5565b91846044359460018060a01b0384168152600160205281812033825260205220549060001982036103dc575b6020866101128787876106cd565b84821061040557509183916103fa60209695610112950333836105cb565b9193948193506103ce565b606490602087519162461bcd60e51b8352820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152fd5b5050346100e557816003193601126100e5576020906002549051908152f35b5050346100e557806003193601126100e55760209061011261048761059a565b60243590336105cb565b8490843461054d578260031936011261054d578260035460018160011c9060018316928315610543575b60209384841081146102a75783885290811561028b575060011461050a57505050829003601f01601f19168201926001600160401b03841183851017610240575082918261023c925282610551565b91925060038652828620918387935b83851061052f575050505083010185808061020b565b805488860183015293019284908201610519565b91607f16916104bb565b8280fd5b6020808252825181830181905290939260005b82811061058657505060409293506000838284010152601f8019910116010190565b818101860151848201604001528501610564565b600435906001600160a01b03821682036105b057565b600080fd5b602435906001600160a01b03821682036105b057565b6001600160a01b0390811691821561067c571691821561062c5760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b6001600160a01b039081169182156107e857169182156107975760008281528060205260408120549180831061074357604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815220818154019055604051908152a3565b60405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608490fdfea26469706673582212200ced0685814ab9f7a521a0e2d42b5bd794d72a762403b96aac97c41e55f223a764736f6c63430008190033",
}

// TestTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TestTokenMetaData.ABI instead.
var TestTokenABI = TestTokenMetaData.ABI

// TestTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestTokenMetaData.Bin instead.
var TestTokenBin = TestTokenMetaData.Bin

// DeployTestToken deploys a new Ethereum contract, binding an instance of TestToken to it.
func DeployTestToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *TestToken, error) {
	parsed, err := TestTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestTokenBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestToken{TestTokenCaller: TestTokenCaller{contract: contract}, TestTokenTransactor: TestTokenTransactor{contract: contract}, TestTokenFilterer: TestTokenFilterer{contract: contract}}, nil
}

// TestToken is an auto generated Go binding around an Ethereum contract.
type TestToken struct {
	TestTokenCaller     // Read-only binding to the contract
	TestTokenTransactor // Write-only binding to the contract
	TestTokenFilterer   // Log filterer for contract events
}

// TestTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestTokenSession struct {
	Contract     *TestToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestTokenCallerSession struct {
	Contract *TestTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTokenTransactorSession struct {
	Contract     *TestTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestTokenRaw struct {
	Contract *TestToken // Generic contract binding to access the raw methods on
}

// TestTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestTokenCallerRaw struct {
	Contract *TestTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TestTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTokenTransactorRaw struct {
	Contract *TestTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestToken creates a new instance of TestToken, bound to a specific deployed contract.
func NewTestToken(address common.Address, backend bind.ContractBackend) (*TestToken, error) {
	contract, err := bindTestToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestToken{TestTokenCaller: TestTokenCaller{contract: contract}, TestTokenTransactor: TestTokenTransactor{contract: contract}, TestTokenFilterer: TestTokenFilterer{contract: contract}}, nil
}

// NewTestTokenCaller creates a new read-only instance of TestToken, bound to a specific deployed contract.
func NewTestTokenCaller(address common.Address, caller bind.ContractCaller) (*TestTokenCaller, error) {
	contract, err := bindTestToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestTokenCaller{contract: contract}, nil
}

// NewTestTokenTransactor creates a new write-only instance of TestToken, bound to a specific deployed contract.
func NewTestTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTokenTransactor, error) {
	contract, err := bindTestToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTokenTransactor{contract: contract}, nil
}

// NewTestTokenFilterer creates a new log filterer instance of TestToken, bound to a specific deployed contract.
func NewTestTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TestTokenFilterer, error) {
	contract, err := bindTestToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestTokenFilterer{contract: contract}, nil
}

// bindTestToken binds a generic wrapper to an already deployed contract.
func bindTestToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken *TestTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken.Contract.TestTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken *TestTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken.Contract.TestTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken *TestTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken.Contract.TestTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestToken *TestTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestToken *TestTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestToken *TestTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestToken *TestTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestToken *TestTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestToken.Contract.Allowance(&_TestToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestToken *TestTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestToken.Contract.Allowance(&_TestToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestToken *TestTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestToken *TestTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestToken.Contract.BalanceOf(&_TestToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestToken *TestTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestToken.Contract.BalanceOf(&_TestToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken *TestTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken *TestTokenSession) Decimals() (uint8, error) {
	return _TestToken.Contract.Decimals(&_TestToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestToken *TestTokenCallerSession) Decimals() (uint8, error) {
	return _TestToken.Contract.Decimals(&_TestToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken *TestTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken *TestTokenSession) Name() (string, error) {
	return _TestToken.Contract.Name(&_TestToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestToken *TestTokenCallerSession) Name() (string, error) {
	return _TestToken.Contract.Name(&_TestToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken *TestTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken *TestTokenSession) Symbol() (string, error) {
	return _TestToken.Contract.Symbol(&_TestToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestToken *TestTokenCallerSession) Symbol() (string, error) {
	return _TestToken.Contract.Symbol(&_TestToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken *TestTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken *TestTokenSession) TotalSupply() (*big.Int, error) {
	return _TestToken.Contract.TotalSupply(&_TestToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestToken *TestTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TestToken.Contract.TotalSupply(&_TestToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestToken *TestTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Approve(&_TestToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Approve(&_TestToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestToken *TestTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestToken *TestTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.DecreaseAllowance(&_TestToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestToken *TestTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.DecreaseAllowance(&_TestToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestToken *TestTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestToken *TestTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.IncreaseAllowance(&_TestToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestToken *TestTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.IncreaseAllowance(&_TestToken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Transfer(&_TestToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.Transfer(&_TestToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.TransferFrom(&_TestToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestToken *TestTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestToken.Contract.TransferFrom(&_TestToken.TransactOpts, from, to, amount)
}

// TestTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestToken contract.
type TestTokenApprovalIterator struct {
	Event *TestTokenApproval // Event containing the contract specifics and raw log

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
func (it *TestTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestTokenApproval)
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
		it.Event = new(TestTokenApproval)
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
func (it *TestTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestTokenApproval represents a Approval event raised by the TestToken contract.
type TestTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken *TestTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestTokenApprovalIterator{contract: _TestToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken *TestTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestTokenApproval)
				if err := _TestToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestToken *TestTokenFilterer) ParseApproval(log types.Log) (*TestTokenApproval, error) {
	event := new(TestTokenApproval)
	if err := _TestToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestToken contract.
type TestTokenTransferIterator struct {
	Event *TestTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TestTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestTokenTransfer)
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
		it.Event = new(TestTokenTransfer)
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
func (it *TestTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestTokenTransfer represents a Transfer event raised by the TestToken contract.
type TestTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken *TestTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestTokenTransferIterator{contract: _TestToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken *TestTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestTokenTransfer)
				if err := _TestToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestToken *TestTokenFilterer) ParseTransfer(log types.Log) (*TestTokenTransfer, error) {
	event := new(TestTokenTransfer)
	if err := _TestToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValueArrayTesterMetaData contains all meta data concerning the ValueArrayTester contract.
var ValueArrayTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460155761069d908161001b8239f35b600080fdfe608060048036101561001057600080fd5b600090813560e01c63f8a8fd6d1461002757600080fd5b3461054357816003193601126105435760608301906001600160401b038083118584101761053057604092835260028552835b8381106105115750825194602091828701908111878210176104fe57845280865251600119016104d257838161008e610547565b828152015261009b610547565b928484526100bc60019485848201528751906100b682610600565b52610600565b5084826100c7610547565b82815201526100f26100d7610547565b8681526002848201528751906100ec82610623565b52610623565b5084826100fd610547565b828152015261010a610547565b8581526003838201528651518581018091116103995761012990610593565b908587815b610497575b5050906101516003928951519061014a8285610633565b5282610633565b50808852510361046b5784845b6103bf575b50848261016e610547565b828152015285518051600019918282019182116103ac579061018f91610633565b5190875151908101908111610399576101a790610593565b8587815b610360575b50508752805190600791600781101561034d5761031b57830151600219016102e6576002875151036102ba57849686975b6101e9578680f35b8051518810156102b6576101fd8882610647565b8051838110156102a3576102715784015186890190818a1161025e57036102285796850196856101e1565b50505162461bcd60e51b815291820152601060248201526f504f505f56414c5f434f4e54454e545360801b604482015260649150fd5b634e487b7160e01b895260118752602489fd5b835162461bcd60e51b8152808701869052600c60248201526b504f505f56414c5f5459504560a01b6044820152606490fd5b634e487b7160e01b895260218752602489fd5b8680f35b505162461bcd60e51b81529182015260076024820152662827a82fa622a760c91b604482015260649150fd5b505162461bcd60e51b815291820152601060248201526f504f505f5245545f434f4e54454e545360801b604482015260649150fd5b50505162461bcd60e51b815291820152600c60248201526b504f505f5245545f5459504560a01b604482015260649150fd5b634e487b7160e01b885260218652602488fd5b8251811015610394578061037683928c51610633565b516103818286610633565b5261038c8185610633565b5001816101ab565b6101b0565b634e487b7160e01b875260118552602487fd5b634e487b7160e01b885260118652602488fd5b865151811015610466576103d38188610647565b8051600781101561034d5761043357830151858201908183116103ac57036103fd5784018461015e565b505162461bcd60e51b8152918201526011602482015270505553485f56414c5f434f4e54454e545360781b604482015260649150fd5b50505162461bcd60e51b815291820152600d60248201526c505553485f56414c5f5459504560981b604482015260649150fd5b610163565b5162461bcd60e51b8152918201526008602482015267282aa9a42fa622a760c11b604482015260649150fd5b895180518210156104cc57906104ae818493610633565b516104b98287610633565b526104c48186610633565b50018161012e565b50610133565b606492519162461bcd60e51b8352820152600960248201526829aa20a92a2fa622a760b91b6044820152fd5b634e487b7160e01b865260418452602486fd5b60209061051c610547565b86815282878183015282890101520161005a565b634e487b7160e01b845260418252602484fd5b5080fd5b60408051919082016001600160401b0381118382101761056657604052565b634e487b7160e01b600052604160045260246000fd5b6001600160401b0381116105665760051b60200190565b9061059d8261057c565b60405190601f1990601f01811682016001600160401b03811183821017610566576040528382526105ce829461057c565b01906000805b8381106105e15750505050565b6020906105ec610547565b8381528284818301528286010152016105d4565b80511561060d5760200190565b634e487b7160e01b600052603260045260246000fd5b80516001101561060d5760400190565b805182101561060d5760209160051b010190565b906106639160006020610658610547565b828152015251610633565b519056fea26469706673582212206f163004ac475d23692c589eff591eb49b20f91e1373b4fe491cd509b3c1dbb664736f6c63430008190033",
}

// ValueArrayTesterABI is the input ABI used to generate the binding from.
// Deprecated: Use ValueArrayTesterMetaData.ABI instead.
var ValueArrayTesterABI = ValueArrayTesterMetaData.ABI

// ValueArrayTesterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValueArrayTesterMetaData.Bin instead.
var ValueArrayTesterBin = ValueArrayTesterMetaData.Bin

// DeployValueArrayTester deploys a new Ethereum contract, binding an instance of ValueArrayTester to it.
func DeployValueArrayTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValueArrayTester, error) {
	parsed, err := ValueArrayTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValueArrayTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValueArrayTester{ValueArrayTesterCaller: ValueArrayTesterCaller{contract: contract}, ValueArrayTesterTransactor: ValueArrayTesterTransactor{contract: contract}, ValueArrayTesterFilterer: ValueArrayTesterFilterer{contract: contract}}, nil
}

// ValueArrayTester is an auto generated Go binding around an Ethereum contract.
type ValueArrayTester struct {
	ValueArrayTesterCaller     // Read-only binding to the contract
	ValueArrayTesterTransactor // Write-only binding to the contract
	ValueArrayTesterFilterer   // Log filterer for contract events
}

// ValueArrayTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueArrayTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueArrayTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueArrayTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueArrayTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueArrayTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueArrayTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueArrayTesterSession struct {
	Contract     *ValueArrayTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueArrayTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueArrayTesterCallerSession struct {
	Contract *ValueArrayTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValueArrayTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueArrayTesterTransactorSession struct {
	Contract     *ValueArrayTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValueArrayTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueArrayTesterRaw struct {
	Contract *ValueArrayTester // Generic contract binding to access the raw methods on
}

// ValueArrayTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueArrayTesterCallerRaw struct {
	Contract *ValueArrayTesterCaller // Generic read-only contract binding to access the raw methods on
}

// ValueArrayTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueArrayTesterTransactorRaw struct {
	Contract *ValueArrayTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValueArrayTester creates a new instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTester(address common.Address, backend bind.ContractBackend) (*ValueArrayTester, error) {
	contract, err := bindValueArrayTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTester{ValueArrayTesterCaller: ValueArrayTesterCaller{contract: contract}, ValueArrayTesterTransactor: ValueArrayTesterTransactor{contract: contract}, ValueArrayTesterFilterer: ValueArrayTesterFilterer{contract: contract}}, nil
}

// NewValueArrayTesterCaller creates a new read-only instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTesterCaller(address common.Address, caller bind.ContractCaller) (*ValueArrayTesterCaller, error) {
	contract, err := bindValueArrayTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTesterCaller{contract: contract}, nil
}

// NewValueArrayTesterTransactor creates a new write-only instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueArrayTesterTransactor, error) {
	contract, err := bindValueArrayTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTesterTransactor{contract: contract}, nil
}

// NewValueArrayTesterFilterer creates a new log filterer instance of ValueArrayTester, bound to a specific deployed contract.
func NewValueArrayTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueArrayTesterFilterer, error) {
	contract, err := bindValueArrayTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueArrayTesterFilterer{contract: contract}, nil
}

// bindValueArrayTester binds a generic wrapper to an already deployed contract.
func bindValueArrayTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValueArrayTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueArrayTester *ValueArrayTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueArrayTester.Contract.ValueArrayTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueArrayTester *ValueArrayTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.ValueArrayTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueArrayTester *ValueArrayTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.ValueArrayTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueArrayTester *ValueArrayTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueArrayTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueArrayTester *ValueArrayTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueArrayTester *ValueArrayTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueArrayTester.Contract.contract.Transact(opts, method, params...)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_ValueArrayTester *ValueArrayTesterCaller) Test(opts *bind.CallOpts) error {
	var out []interface{}
	err := _ValueArrayTester.contract.Call(opts, &out, "test")

	if err != nil {
		return err
	}

	return err

}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_ValueArrayTester *ValueArrayTesterSession) Test() error {
	return _ValueArrayTester.Contract.Test(&_ValueArrayTester.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() pure returns()
func (_ValueArrayTester *ValueArrayTesterCallerSession) Test() error {
	return _ValueArrayTester.Contract.Test(&_ValueArrayTester.CallOpts)
}
