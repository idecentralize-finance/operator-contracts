// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package operator

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

// SequencerLoanSequence is an auto generated low-level Go binding around an user-defined struct.
type SequencerLoanSequence struct {
	Assets  []common.Address
	Amounts []*big.Int
	Modes   []*big.Int
}

// OperatorMetaData contains all meta data concerning the Operator contract.
var OperatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"_addressProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Sequenced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LENDING_POOL\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"premiums\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"modes\",\"type\":\"uint256[]\"}],\"internalType\":\"structSequencer.LoanSequence\",\"name\":\"loanData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"tradeData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"useLoan\",\"type\":\"bool\"}],\"name\":\"executeSequence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620026b6380380620026b683398181016040528101906200003791906200031c565b818073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff16630261bf8b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620000b8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000de919062000363565b73ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250505080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630261bf8b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000225919062000363565b600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000395565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200029f8262000272565b9050919050565b6000620002b38262000292565b9050919050565b620002c581620002a6565b8114620002d157600080fd5b50565b600081519050620002e581620002ba565b92915050565b620002f68162000292565b81146200030257600080fd5b50565b6000815190506200031681620002eb565b92915050565b600080604083850312156200033657620003356200026d565b5b60006200034685828601620002d4565b9250506020620003598582860162000305565b9150509250929050565b6000602082840312156200037c576200037b6200026d565b5b60006200038c8482850162000305565b91505092915050565b60805160a0516122fb620003bb60003960006107e60152600061017701526122fb6000f3fe6080604052600436106100595760003560e01c80630542975c1461006557806351cff8d91461009057806371c6f4cc146100b95780638da5cb5b146100e2578063920f5c841461010d578063b4dcfc771461014a57610060565b3661006057005b600080fd5b34801561007157600080fd5b5061007a610175565b6040516100879190610f32565b60405180910390f35b34801561009c57600080fd5b506100b760048036038101906100b29190610f9f565b610199565b005b3480156100c557600080fd5b506100e060048036038101906100db91906113cd565b6103c5565b005b3480156100ee57600080fd5b506100f7610558565b6040516101049190611467565b60405180910390f35b34801561011957600080fd5b50610134600480360381019061012f9190611589565b61057e565b6040516101419190611694565b60405180910390f35b34801561015657600080fd5b5061015f6107e4565b60405161016c91906116d0565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610229576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022090611748565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036102fb576000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16476040516102a590611799565b60006040518083038185875af1925050503d80600081146102e2576040519150601f19603f3d011682016040523d82523d6000602084013e6102e7565b606091505b50509050806102f557600080fd5b506103c2565b6103c1600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161035a9190611467565b602060405180830381865afa158015610377573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061039b91906117c3565b8373ffffffffffffffffffffffffffffffffffffffff166108089092919063ffffffff16565b5b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610455576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044c90611748565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008490508215610510578173ffffffffffffffffffffffffffffffffffffffff1663ab9c4b5d30836000015184602001518560400151338a60006040518863ffffffff1660e01b81526004016104d99796959493929190611a34565b600060405180830381600087803b1580156104f357600080fd5b505af1158015610507573d6000803e3d6000fd5b50505050610551565b6105198461088e565b7f418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce336040516105489190611467565b60405180910390a15b5050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610610576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060790611b0b565b60405180910390fd5b600061065f84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061093f565b905060005b8160000151518110156106c9576106b78260000151828151811061068b5761068a611b2b565b5b6020026020010151836020015183815181106106aa576106a9611b2b565b5b6020026020010151610980565b50806106c290611b89565b9050610664565b5060005b8b8b905081101561079a57610789600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1689898481811061071157610710611b2b565b5b905060200201358c8c8581811061072b5761072a611b2b565b5b9050602002013561073c9190611bd1565b8e8e8581811061074f5761074e611b2b565b5b90506020020160208101906107649190610f9f565b73ffffffffffffffffffffffffffffffffffffffff16610acd9092919063ffffffff16565b8061079390611b89565b90506106cd565b507f418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce856040516107ca9190611467565b60405180910390a160019150509998505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6108898363a9059cbb60e01b8484604051602401610827929190611c14565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610c1c565b505050565b60006108998261093f565b905060005b816000015151811015610903576108f1826000015182815181106108c5576108c4611b2b565b5b6020026020010151836020015183815181106108e4576108e3611b2b565b5b6020026020010151610980565b50806108fc90611b89565b905061089e565b507f418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce336040516109339190611467565b60405180910390a15050565b610947610e99565b6000808380602001905181019061095e9190611e3a565b9150915060405180604001604052808381526020018281525092505050919050565b60603073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e790611efe565b60405180910390fd5b600060608473ffffffffffffffffffffffffffffffffffffffff1684604051610a199190611f4f565b6000604051808303816000865af19150503d8060008114610a56576040519150601f19603f3d011682016040523d82523d6000602084013e610a5b565b606091505b50809250819350505081158015610a73575060008151115b15610a82573d6000803e3d6000fd5b81610ac2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab990611fb2565b60405180910390fd5b809250505092915050565b6000811480610b57575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b8152600401610b14929190611fd2565b602060405180830381865afa158015610b31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5591906117c3565b145b610b96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8d9061206d565b60405180910390fd5b610c178363095ea7b360e01b8484604051602401610bb5929190611c14565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610c1c565b505050565b6000610c7e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610ce39092919063ffffffff16565b9050600081511115610cde5780806020019051810190610c9e91906120a2565b610cdd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd490612141565b60405180910390fd5b5b505050565b6060610cf28484600085610cfb565b90509392505050565b606082471015610d40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d37906121d3565b60405180910390fd5b610d4985610e0f565b610d88576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7f9061223f565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610db19190611f4f565b60006040518083038185875af1925050503d8060008114610dee576040519150601f19603f3d011682016040523d82523d6000602084013e610df3565b606091505b5091509150610e03828286610e32565b92505050949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60608315610e4257829050610e92565b600083511115610e555782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8991906122a3565b60405180910390fd5b9392505050565b604051806040016040528060608152602001606081525090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610ef8610ef3610eee84610eb3565b610ed3565b610eb3565b9050919050565b6000610f0a82610edd565b9050919050565b6000610f1c82610eff565b9050919050565b610f2c81610f11565b82525050565b6000602082019050610f476000830184610f23565b92915050565b6000604051905090565b600080fd5b600080fd5b6000610f6c82610eb3565b9050919050565b610f7c81610f61565b8114610f8757600080fd5b50565b600081359050610f9981610f73565b92915050565b600060208284031215610fb557610fb4610f57565b5b6000610fc384828501610f8a565b91505092915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61101a82610fd1565b810181811067ffffffffffffffff8211171561103957611038610fe2565b5b80604052505050565b600061104c610f4d565b90506110588282611011565b919050565b600080fd5b600080fd5b600067ffffffffffffffff82111561108257611081610fe2565b5b602082029050602081019050919050565b600080fd5b60006110ab6110a684611067565b611042565b905080838252602082019050602084028301858111156110ce576110cd611093565b5b835b818110156110f757806110e38882610f8a565b8452602084019350506020810190506110d0565b5050509392505050565b600082601f83011261111657611115611062565b5b8135611126848260208601611098565b91505092915050565b600067ffffffffffffffff82111561114a57611149610fe2565b5b602082029050602081019050919050565b6000819050919050565b61116e8161115b565b811461117957600080fd5b50565b60008135905061118b81611165565b92915050565b60006111a461119f8461112f565b611042565b905080838252602082019050602084028301858111156111c7576111c6611093565b5b835b818110156111f057806111dc888261117c565b8452602084019350506020810190506111c9565b5050509392505050565b600082601f83011261120f5761120e611062565b5b813561121f848260208601611191565b91505092915050565b60006060828403121561123e5761123d610fcc565b5b6112486060611042565b9050600082013567ffffffffffffffff8111156112685761126761105d565b5b61127484828501611101565b600083015250602082013567ffffffffffffffff8111156112985761129761105d565b5b6112a4848285016111fa565b602083015250604082013567ffffffffffffffff8111156112c8576112c761105d565b5b6112d4848285016111fa565b60408301525092915050565b600080fd5b600067ffffffffffffffff821115611300576112ff610fe2565b5b61130982610fd1565b9050602081019050919050565b82818337600083830152505050565b6000611338611333846112e5565b611042565b905082815260208101848484011115611354576113536112e0565b5b61135f848285611316565b509392505050565b600082601f83011261137c5761137b611062565b5b813561138c848260208601611325565b91505092915050565b60008115159050919050565b6113aa81611395565b81146113b557600080fd5b50565b6000813590506113c7816113a1565b92915050565b6000806000606084860312156113e6576113e5610f57565b5b600084013567ffffffffffffffff81111561140457611403610f5c565b5b61141086828701611228565b935050602084013567ffffffffffffffff81111561143157611430610f5c565b5b61143d86828701611367565b925050604061144e868287016113b8565b9150509250925092565b61146181610f61565b82525050565b600060208201905061147c6000830184611458565b92915050565b600080fd5b60008083601f84011261149d5761149c611062565b5b8235905067ffffffffffffffff8111156114ba576114b9611482565b5b6020830191508360208202830111156114d6576114d5611093565b5b9250929050565b60008083601f8401126114f3576114f2611062565b5b8235905067ffffffffffffffff8111156115105761150f611482565b5b60208301915083602082028301111561152c5761152b611093565b5b9250929050565b60008083601f84011261154957611548611062565b5b8235905067ffffffffffffffff81111561156657611565611482565b5b60208301915083600182028301111561158257611581611093565b5b9250929050565b600080600080600080600080600060a08a8c0312156115ab576115aa610f57565b5b60008a013567ffffffffffffffff8111156115c9576115c8610f5c565b5b6115d58c828d01611487565b995099505060208a013567ffffffffffffffff8111156115f8576115f7610f5c565b5b6116048c828d016114dd565b975097505060408a013567ffffffffffffffff81111561162757611626610f5c565b5b6116338c828d016114dd565b955095505060606116468c828d01610f8a565b93505060808a013567ffffffffffffffff81111561166757611666610f5c565b5b6116738c828d01611533565b92509250509295985092959850929598565b61168e81611395565b82525050565b60006020820190506116a96000830184611685565b92915050565b60006116ba82610eff565b9050919050565b6116ca816116af565b82525050565b60006020820190506116e560008301846116c1565b92915050565b600082825260208201905092915050565b7f3432000000000000000000000000000000000000000000000000000000000000600082015250565b60006117326002836116eb565b915061173d826116fc565b602082019050919050565b6000602082019050818103600083015261176181611725565b9050919050565b600081905092915050565b50565b6000611783600083611768565b915061178e82611773565b600082019050919050565b60006117a482611776565b9150819050919050565b6000815190506117bd81611165565b92915050565b6000602082840312156117d9576117d8610f57565b5b60006117e7848285016117ae565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61182581610f61565b82525050565b6000611837838361181c565b60208301905092915050565b6000602082019050919050565b600061185b826117f0565b61186581856117fb565b93506118708361180c565b8060005b838110156118a1578151611888888261182b565b975061189383611843565b925050600181019050611874565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6118e38161115b565b82525050565b60006118f583836118da565b60208301905092915050565b6000602082019050919050565b6000611919826118ae565b61192381856118b9565b935061192e836118ca565b8060005b8381101561195f57815161194688826118e9565b975061195183611901565b925050600181019050611932565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156119a657808201518184015260208101905061198b565b60008484015250505050565b60006119bd8261196c565b6119c78185611977565b93506119d7818560208601611988565b6119e081610fd1565b840191505092915050565b6000819050919050565b600061ffff82169050919050565b6000611a1e611a19611a14846119eb565b610ed3565b6119f5565b9050919050565b611a2e81611a03565b82525050565b600060e082019050611a49600083018a611458565b8181036020830152611a5b8189611850565b90508181036040830152611a6f818861190e565b90508181036060830152611a83818761190e565b9050611a926080830186611458565b81810360a0830152611aa481856119b2565b9050611ab360c0830184611a25565b98975050505050505050565b7f706f6f6c206e6f74205061756c00000000000000000000000000000000000000600082015250565b6000611af5600d836116eb565b9150611b0082611abf565b602082019050919050565b60006020820190508181036000830152611b2481611ae8565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b948261115b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611bc657611bc5611b5a565b5b600182019050919050565b6000611bdc8261115b565b9150611be78361115b565b9250828201905080821115611bff57611bfe611b5a565b5b92915050565b611c0e8161115b565b82525050565b6000604082019050611c296000830185611458565b611c366020830184611c05565b9392505050565b600081519050611c4c81610f73565b92915050565b6000611c65611c6084611067565b611042565b90508083825260208201905060208402830185811115611c8857611c87611093565b5b835b81811015611cb15780611c9d8882611c3d565b845260208401935050602081019050611c8a565b5050509392505050565b600082601f830112611cd057611ccf611062565b5b8151611ce0848260208601611c52565b91505092915050565b600067ffffffffffffffff821115611d0457611d03610fe2565b5b602082029050602081019050919050565b6000611d28611d23846112e5565b611042565b905082815260208101848484011115611d4457611d436112e0565b5b611d4f848285611988565b509392505050565b600082601f830112611d6c57611d6b611062565b5b8151611d7c848260208601611d15565b91505092915050565b6000611d98611d9384611ce9565b611042565b90508083825260208201905060208402830185811115611dbb57611dba611093565b5b835b81811015611e0257805167ffffffffffffffff811115611de057611ddf611062565b5b808601611ded8982611d57565b85526020850194505050602081019050611dbd565b5050509392505050565b600082601f830112611e2157611e20611062565b5b8151611e31848260208601611d85565b91505092915050565b60008060408385031215611e5157611e50610f57565b5b600083015167ffffffffffffffff811115611e6f57611e6e610f5c565b5b611e7b85828601611cbb565b925050602083015167ffffffffffffffff811115611e9c57611e9b610f5c565b5b611ea885828601611e0c565b9150509250929050565b7f496e7465726e616c207265656e74727900000000000000000000000000000000600082015250565b6000611ee86010836116eb565b9150611ef382611eb2565b602082019050919050565b60006020820190508181036000830152611f1781611edb565b9050919050565b6000611f298261196c565b611f338185611768565b9350611f43818560208601611988565b80840191505092915050565b6000611f5b8284611f1e565b915081905092915050565b7f564d3a206578656375746543616c6c2072657665727465640000000000000000600082015250565b6000611f9c6018836116eb565b9150611fa782611f66565b602082019050919050565b60006020820190508181036000830152611fcb81611f8f565b9050919050565b6000604082019050611fe76000830185611458565b611ff46020830184611458565b9392505050565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b60006120576036836116eb565b915061206282611ffb565b604082019050919050565b600060208201905081810360008301526120868161204a565b9050919050565b60008151905061209c816113a1565b92915050565b6000602082840312156120b8576120b7610f57565b5b60006120c68482850161208d565b91505092915050565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b600061212b602a836116eb565b9150612136826120cf565b604082019050919050565b6000602082019050818103600083015261215a8161211e565b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b60006121bd6026836116eb565b91506121c882612161565b604082019050919050565b600060208201905081810360008301526121ec816121b0565b9050919050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b6000612229601d836116eb565b9150612234826121f3565b602082019050919050565b600060208201905081810360008301526122588161221c565b9050919050565b600081519050919050565b60006122758261225f565b61227f81856116eb565b935061228f818560208601611988565b61229881610fd1565b840191505092915050565b600060208201905081810360008301526122bd818461226a565b90509291505056fea2646970667358221220cc616b86537034aacaf7b512c7cbfccf78632d7cde813309fcd48aede80588ae64736f6c63430008110033",
}

// OperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorMetaData.ABI instead.
var OperatorABI = OperatorMetaData.ABI

// OperatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OperatorMetaData.Bin instead.
var OperatorBin = OperatorMetaData.Bin

// DeployOperator deploys a new Ethereum contract, binding an instance of Operator to it.
func DeployOperator(auth *bind.TransactOpts, backend bind.ContractBackend, _addressProvider common.Address, _devAddress common.Address) (common.Address, *types.Transaction, *Operator, error) {
	parsed, err := OperatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OperatorBin), backend, _addressProvider, _devAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Operator{OperatorCaller: OperatorCaller{contract: contract}, OperatorTransactor: OperatorTransactor{contract: contract}, OperatorFilterer: OperatorFilterer{contract: contract}}, nil
}

// Operator is an auto generated Go binding around an Ethereum contract.
type Operator struct {
	OperatorCaller     // Read-only binding to the contract
	OperatorTransactor // Write-only binding to the contract
	OperatorFilterer   // Log filterer for contract events
}

// OperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorSession struct {
	Contract     *Operator         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorCallerSession struct {
	Contract *OperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorTransactorSession struct {
	Contract     *OperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorRaw struct {
	Contract *Operator // Generic contract binding to access the raw methods on
}

// OperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorCallerRaw struct {
	Contract *OperatorCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorTransactorRaw struct {
	Contract *OperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperator creates a new instance of Operator, bound to a specific deployed contract.
func NewOperator(address common.Address, backend bind.ContractBackend) (*Operator, error) {
	contract, err := bindOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Operator{OperatorCaller: OperatorCaller{contract: contract}, OperatorTransactor: OperatorTransactor{contract: contract}, OperatorFilterer: OperatorFilterer{contract: contract}}, nil
}

// NewOperatorCaller creates a new read-only instance of Operator, bound to a specific deployed contract.
func NewOperatorCaller(address common.Address, caller bind.ContractCaller) (*OperatorCaller, error) {
	contract, err := bindOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorCaller{contract: contract}, nil
}

// NewOperatorTransactor creates a new write-only instance of Operator, bound to a specific deployed contract.
func NewOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorTransactor, error) {
	contract, err := bindOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorTransactor{contract: contract}, nil
}

// NewOperatorFilterer creates a new log filterer instance of Operator, bound to a specific deployed contract.
func NewOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorFilterer, error) {
	contract, err := bindOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorFilterer{contract: contract}, nil
}

// bindOperator binds a generic wrapper to an already deployed contract.
func bindOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OperatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.OperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Operator *OperatorCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Operator *OperatorSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Operator.Contract.ADDRESSESPROVIDER(&_Operator.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Operator *OperatorCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Operator.Contract.ADDRESSESPROVIDER(&_Operator.CallOpts)
}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_Operator *OperatorCaller) LENDINGPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "LENDING_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_Operator *OperatorSession) LENDINGPOOL() (common.Address, error) {
	return _Operator.Contract.LENDINGPOOL(&_Operator.CallOpts)
}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_Operator *OperatorCallerSession) LENDINGPOOL() (common.Address, error) {
	return _Operator.Contract.LENDINGPOOL(&_Operator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorSession) Owner() (common.Address, error) {
	return _Operator.Contract.Owner(&_Operator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorCallerSession) Owner() (common.Address, error) {
	return _Operator.Contract.Owner(&_Operator.CallOpts)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_Operator *OperatorTransactor) ExecuteOperation(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "executeOperation", assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_Operator *OperatorSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _Operator.Contract.ExecuteOperation(&_Operator.TransactOpts, assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_Operator *OperatorTransactorSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _Operator.Contract.ExecuteOperation(&_Operator.TransactOpts, assets, amounts, premiums, initiator, params)
}

// ExecuteSequence is a paid mutator transaction binding the contract method 0x71c6f4cc.
//
// Solidity: function executeSequence((address[],uint256[],uint256[]) loanData, bytes tradeData, bool useLoan) returns()
func (_Operator *OperatorTransactor) ExecuteSequence(opts *bind.TransactOpts, loanData SequencerLoanSequence, tradeData []byte, useLoan bool) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "executeSequence", loanData, tradeData, useLoan)
}

// ExecuteSequence is a paid mutator transaction binding the contract method 0x71c6f4cc.
//
// Solidity: function executeSequence((address[],uint256[],uint256[]) loanData, bytes tradeData, bool useLoan) returns()
func (_Operator *OperatorSession) ExecuteSequence(loanData SequencerLoanSequence, tradeData []byte, useLoan bool) (*types.Transaction, error) {
	return _Operator.Contract.ExecuteSequence(&_Operator.TransactOpts, loanData, tradeData, useLoan)
}

// ExecuteSequence is a paid mutator transaction binding the contract method 0x71c6f4cc.
//
// Solidity: function executeSequence((address[],uint256[],uint256[]) loanData, bytes tradeData, bool useLoan) returns()
func (_Operator *OperatorTransactorSession) ExecuteSequence(loanData SequencerLoanSequence, tradeData []byte, useLoan bool) (*types.Transaction, error) {
	return _Operator.Contract.ExecuteSequence(&_Operator.TransactOpts, loanData, tradeData, useLoan)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _asset) returns()
func (_Operator *OperatorTransactor) Withdraw(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "withdraw", _asset)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _asset) returns()
func (_Operator *OperatorSession) Withdraw(_asset common.Address) (*types.Transaction, error) {
	return _Operator.Contract.Withdraw(&_Operator.TransactOpts, _asset)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _asset) returns()
func (_Operator *OperatorTransactorSession) Withdraw(_asset common.Address) (*types.Transaction, error) {
	return _Operator.Contract.Withdraw(&_Operator.TransactOpts, _asset)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Operator *OperatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Operator *OperatorSession) Receive() (*types.Transaction, error) {
	return _Operator.Contract.Receive(&_Operator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Operator *OperatorTransactorSession) Receive() (*types.Transaction, error) {
	return _Operator.Contract.Receive(&_Operator.TransactOpts)
}

// OperatorSequencedIterator is returned from FilterSequenced and is used to iterate over the raw logs and unpacked data for Sequenced events raised by the Operator contract.
type OperatorSequencedIterator struct {
	Event *OperatorSequenced // Event containing the contract specifics and raw log

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
func (it *OperatorSequencedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSequenced)
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
		it.Event = new(OperatorSequenced)
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
func (it *OperatorSequencedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSequencedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSequenced represents a Sequenced event raised by the Operator contract.
type OperatorSequenced struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSequenced is a free log retrieval operation binding the contract event 0x418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce.
//
// Solidity: event Sequenced(address user)
func (_Operator *OperatorFilterer) FilterSequenced(opts *bind.FilterOpts) (*OperatorSequencedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "Sequenced")
	if err != nil {
		return nil, err
	}
	return &OperatorSequencedIterator{contract: _Operator.contract, event: "Sequenced", logs: logs, sub: sub}, nil
}

// WatchSequenced is a free log subscription operation binding the contract event 0x418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce.
//
// Solidity: event Sequenced(address user)
func (_Operator *OperatorFilterer) WatchSequenced(opts *bind.WatchOpts, sink chan<- *OperatorSequenced) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "Sequenced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSequenced)
				if err := _Operator.contract.UnpackLog(event, "Sequenced", log); err != nil {
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

// ParseSequenced is a log parse operation binding the contract event 0x418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce.
//
// Solidity: event Sequenced(address user)
func (_Operator *OperatorFilterer) ParseSequenced(log types.Log) (*OperatorSequenced, error) {
	event := new(OperatorSequenced)
	if err := _Operator.contract.UnpackLog(event, "Sequenced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
