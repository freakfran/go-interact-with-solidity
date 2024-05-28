// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interactions

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

// BasicNftMetaData contains all meta data concerning the BasicNft contract.
var BasicNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"mintNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040518060400160405280600581526020017f446f6769650000000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f444f470000000000000000000000000000000000000000000000000000000000815250815f908161008a91906102e3565b50806001908161009a91906102e3565b5050505f6006819055506103b2565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061012457607f821691505b602082108103610137576101366100e0565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101997fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261015e565b6101a3868361015e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6101e76101e26101dd846101bb565b6101c4565b6101bb565b9050919050565b5f819050919050565b610200836101cd565b61021461020c826101ee565b84845461016a565b825550505050565b5f90565b61022861021c565b6102338184846101f7565b505050565b5b818110156102565761024b5f82610220565b600181019050610239565b5050565b601f82111561029b5761026c8161013d565b6102758461014f565b81016020851015610284578190505b6102986102908561014f565b830182610238565b50505b505050565b5f82821c905092915050565b5f6102bb5f19846008026102a0565b1980831691505092915050565b5f6102d383836102ac565b9150826002028217905092915050565b6102ec826100a9565b67ffffffffffffffff811115610305576103046100b3565b5b61030f825461010d565b61031a82828561025a565b5f60209050601f83116001811461034b575f8415610339578287015190505b61034385826102c8565b8655506103aa565b601f1984166103598661013d565b5f5b828110156103805784890151825560018201915060208501945060208101905061035b565b8683101561039d5784890151610399601f8916826102ac565b8355505b6001600288020188555050505b505050505050565b611f2e806103bf5f395ff3fe608060405234801561000f575f80fd5b50600436106100f3575f3560e01c80636e02007d11610095578063b88d4fde11610064578063b88d4fde14610281578063c87b56dd1461029d578063da86f4a9146102cd578063e985e9c5146102e9576100f3565b80636e02007d146101f957806370a082311461021757806395d89b4114610247578063a22cb46514610265576100f3565b8063095ea7b3116100d1578063095ea7b31461017557806323b872dd1461019157806342842e0e146101ad5780636352211e146101c9576100f3565b806301ffc9a7146100f757806306fdde0314610127578063081812fc14610145575b5f80fd5b610111600480360381019061010c919061145e565b610319565b60405161011e91906114a3565b60405180910390f35b61012f6103fa565b60405161013c9190611546565b60405180910390f35b61015f600480360381019061015a9190611599565b610489565b60405161016c9190611603565b60405180910390f35b61018f600480360381019061018a9190611646565b6104a4565b005b6101ab60048036038101906101a69190611684565b6104ba565b005b6101c760048036038101906101c29190611684565b6105b9565b005b6101e360048036038101906101de9190611599565b6105d8565b6040516101f09190611603565b60405180910390f35b6102016105e9565b60405161020e91906116e3565b60405180910390f35b610231600480360381019061022c91906116fc565b6105f2565b60405161023e91906116e3565b60405180910390f35b61024f6106a8565b60405161025c9190611546565b60405180910390f35b61027f600480360381019061027a9190611751565b610738565b005b61029b600480360381019061029691906118bb565b61074e565b005b6102b760048036038101906102b29190611599565b610773565b6040516102c49190611546565b60405180910390f35b6102e760048036038101906102e291906119d9565b610814565b005b61030360048036038101906102fe9190611a20565b61085b565b60405161031091906114a3565b60405180910390f35b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103e357507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806103f357506103f2826108e9565b5b9050919050565b60605f805461040890611a8b565b80601f016020809104026020016040519081016040528092919081815260200182805461043490611a8b565b801561047f5780601f106104565761010080835404028352916020019161047f565b820191905f5260205f20905b81548152906001019060200180831161046257829003601f168201915b5050505050905090565b5f61049382610952565b5061049d826109d8565b9050919050565b6104b682826104b1610a11565b610a18565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361052a575f6040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016105219190611603565b60405180910390fd5b5f61053d8383610538610a11565b610a2a565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146105b3578382826040517f64283d7b0000000000000000000000000000000000000000000000000000000081526004016105aa93929190611abb565b60405180910390fd5b50505050565b6105d383838360405180602001604052805f81525061074e565b505050565b5f6105e282610952565b9050919050565b5f600654905090565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610663575f6040517f89c62b6400000000000000000000000000000000000000000000000000000000815260040161065a9190611603565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6060600180546106b790611a8b565b80601f01602080910402602001604051908101604052809291908181526020018280546106e390611a8b565b801561072e5780601f106107055761010080835404028352916020019161072e565b820191905f5260205f20905b81548152906001019060200180831161071157829003601f168201915b5050505050905090565b61074a610743610a11565b8383610c35565b5050565b6107598484846104ba565b61076d610764610a11565b85858585610d9e565b50505050565b606060075f8381526020019081526020015f20805461079190611a8b565b80601f01602080910402602001604051908101604052809291908181526020018280546107bd90611a8b565b80156108085780601f106107df57610100808354040283529160200191610808565b820191905f5260205f20905b8154815290600101906020018083116107eb57829003601f168201915b50505050509050919050565b8060075f60065481526020019081526020015f2090816108349190611c8d565b5061084133600654610f4a565b60065f81548092919061085390611d89565b919050555050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f8061095d83610f67565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109cf57826040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016109c691906116e3565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610a258383836001610fa0565b505050565b5f80610a3584610f67565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610a7657610a7581848661115f565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b0157610ab55f855f80610fa0565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610b8057600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ca557816040517f5b08ba18000000000000000000000000000000000000000000000000000000008152600401610c9c9190611603565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610d9191906114a3565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115610f43578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401610dfc9493929190611e22565b6020604051808303815f875af1925050508015610e3757506040513d601f19601f82011682018060405250810190610e349190611e80565b60015b610eb8573d805f8114610e65576040519150601f19603f3d011682016040523d82523d5f602084013e610e6a565b606091505b505f815103610eb057836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610ea79190611603565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610f4157836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610f389190611603565b60405180910390fd5b505b5050505050565b610f63828260405180602001604052805f815250611222565b5050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8080610fd857505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b1561110a575f610fe784610952565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561105157508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156110645750611062818461085b565b155b156110a657826040517fa9fbf51f00000000000000000000000000000000000000000000000000000000815260040161109d9190611603565b60405180910390fd5b811561110857838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b61116a838383611245565b61121d575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036111de57806040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016111d591906116e3565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401611214929190611eab565b60405180910390fd5b505050565b61122c8383611305565b611240611237610a11565b5f858585610d9e565b505050565b5f8073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156112fc57508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806112bd57506112bc848461085b565b5b806112fb57508273ffffffffffffffffffffffffffffffffffffffff166112e3836109d8565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611375575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161136c9190611603565b60405180910390fd5b5f61138183835f610a2a565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146113f3575f6040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081526004016113ea9190611603565b60405180910390fd5b505050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61143d81611409565b8114611447575f80fd5b50565b5f8135905061145881611434565b92915050565b5f6020828403121561147357611472611401565b5b5f6114808482850161144a565b91505092915050565b5f8115159050919050565b61149d81611489565b82525050565b5f6020820190506114b65f830184611494565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156114f35780820151818401526020810190506114d8565b5f8484015250505050565b5f601f19601f8301169050919050565b5f611518826114bc565b61152281856114c6565b93506115328185602086016114d6565b61153b816114fe565b840191505092915050565b5f6020820190508181035f83015261155e818461150e565b905092915050565b5f819050919050565b61157881611566565b8114611582575f80fd5b50565b5f813590506115938161156f565b92915050565b5f602082840312156115ae576115ad611401565b5b5f6115bb84828501611585565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6115ed826115c4565b9050919050565b6115fd816115e3565b82525050565b5f6020820190506116165f8301846115f4565b92915050565b611625816115e3565b811461162f575f80fd5b50565b5f813590506116408161161c565b92915050565b5f806040838503121561165c5761165b611401565b5b5f61166985828601611632565b925050602061167a85828601611585565b9150509250929050565b5f805f6060848603121561169b5761169a611401565b5b5f6116a886828701611632565b93505060206116b986828701611632565b92505060406116ca86828701611585565b9150509250925092565b6116dd81611566565b82525050565b5f6020820190506116f65f8301846116d4565b92915050565b5f6020828403121561171157611710611401565b5b5f61171e84828501611632565b91505092915050565b61173081611489565b811461173a575f80fd5b50565b5f8135905061174b81611727565b92915050565b5f806040838503121561176757611766611401565b5b5f61177485828601611632565b92505060206117858582860161173d565b9150509250929050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6117cd826114fe565b810181811067ffffffffffffffff821117156117ec576117eb611797565b5b80604052505050565b5f6117fe6113f8565b905061180a82826117c4565b919050565b5f67ffffffffffffffff82111561182957611828611797565b5b611832826114fe565b9050602081019050919050565b828183375f83830152505050565b5f61185f61185a8461180f565b6117f5565b90508281526020810184848401111561187b5761187a611793565b5b61188684828561183f565b509392505050565b5f82601f8301126118a2576118a161178f565b5b81356118b284826020860161184d565b91505092915050565b5f805f80608085870312156118d3576118d2611401565b5b5f6118e087828801611632565b94505060206118f187828801611632565b935050604061190287828801611585565b925050606085013567ffffffffffffffff81111561192357611922611405565b5b61192f8782880161188e565b91505092959194509250565b5f67ffffffffffffffff82111561195557611954611797565b5b61195e826114fe565b9050602081019050919050565b5f61197d6119788461193b565b6117f5565b90508281526020810184848401111561199957611998611793565b5b6119a484828561183f565b509392505050565b5f82601f8301126119c0576119bf61178f565b5b81356119d084826020860161196b565b91505092915050565b5f602082840312156119ee576119ed611401565b5b5f82013567ffffffffffffffff811115611a0b57611a0a611405565b5b611a17848285016119ac565b91505092915050565b5f8060408385031215611a3657611a35611401565b5b5f611a4385828601611632565b9250506020611a5485828601611632565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611aa257607f821691505b602082108103611ab557611ab4611a5e565b5b50919050565b5f606082019050611ace5f8301866115f4565b611adb60208301856116d4565b611ae860408301846115f4565b949350505050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611b4c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611b11565b611b568683611b11565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611b91611b8c611b8784611566565b611b6e565b611566565b9050919050565b5f819050919050565b611baa83611b77565b611bbe611bb682611b98565b848454611b1d565b825550505050565b5f90565b611bd2611bc6565b611bdd818484611ba1565b505050565b5b81811015611c0057611bf55f82611bca565b600181019050611be3565b5050565b601f821115611c4557611c1681611af0565b611c1f84611b02565b81016020851015611c2e578190505b611c42611c3a85611b02565b830182611be2565b50505b505050565b5f82821c905092915050565b5f611c655f1984600802611c4a565b1980831691505092915050565b5f611c7d8383611c56565b9150826002028217905092915050565b611c96826114bc565b67ffffffffffffffff811115611caf57611cae611797565b5b611cb98254611a8b565b611cc4828285611c04565b5f60209050601f831160018114611cf5575f8415611ce3578287015190505b611ced8582611c72565b865550611d54565b601f198416611d0386611af0565b5f5b82811015611d2a57848901518255600182019150602085019450602081019050611d05565b86831015611d475784890151611d43601f891682611c56565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611d9382611566565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611dc557611dc4611d5c565b5b600182019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f611df482611dd0565b611dfe8185611dda565b9350611e0e8185602086016114d6565b611e17816114fe565b840191505092915050565b5f608082019050611e355f8301876115f4565b611e4260208301866115f4565b611e4f60408301856116d4565b8181036060830152611e618184611dea565b905095945050505050565b5f81519050611e7a81611434565b92915050565b5f60208284031215611e9557611e94611401565b5b5f611ea284828501611e6c565b91505092915050565b5f604082019050611ebe5f8301856115f4565b611ecb60208301846116d4565b939250505056fea2646970667358221220becbe0831f35c118901cb5527bb22d5cb167d0259145909a4af8e74d2719ddb864736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// BasicNftABI is the input ABI used to generate the binding from.
// Deprecated: Use BasicNftMetaData.ABI instead.
var BasicNftABI = BasicNftMetaData.ABI

// BasicNftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BasicNftMetaData.Bin instead.
var BasicNftBin = BasicNftMetaData.Bin

// DeployBasicNft deploys a new Ethereum contract, binding an instance of BasicNft to it.
func DeployBasicNft(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicNft, error) {
	parsed, err := BasicNftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BasicNftBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicNft{BasicNftCaller: BasicNftCaller{contract: contract}, BasicNftTransactor: BasicNftTransactor{contract: contract}, BasicNftFilterer: BasicNftFilterer{contract: contract}}, nil
}

// BasicNft is an auto generated Go binding around an Ethereum contract.
type BasicNft struct {
	BasicNftCaller     // Read-only binding to the contract
	BasicNftTransactor // Write-only binding to the contract
	BasicNftFilterer   // Log filterer for contract events
}

// BasicNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicNftSession struct {
	Contract     *BasicNft         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicNftCallerSession struct {
	Contract *BasicNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BasicNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicNftTransactorSession struct {
	Contract     *BasicNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BasicNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicNftRaw struct {
	Contract *BasicNft // Generic contract binding to access the raw methods on
}

// BasicNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicNftCallerRaw struct {
	Contract *BasicNftCaller // Generic read-only contract binding to access the raw methods on
}

// BasicNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicNftTransactorRaw struct {
	Contract *BasicNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicNft creates a new instance of BasicNft, bound to a specific deployed contract.
func NewBasicNft(address common.Address, backend bind.ContractBackend) (*BasicNft, error) {
	contract, err := bindBasicNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicNft{BasicNftCaller: BasicNftCaller{contract: contract}, BasicNftTransactor: BasicNftTransactor{contract: contract}, BasicNftFilterer: BasicNftFilterer{contract: contract}}, nil
}

// NewBasicNftCaller creates a new read-only instance of BasicNft, bound to a specific deployed contract.
func NewBasicNftCaller(address common.Address, caller bind.ContractCaller) (*BasicNftCaller, error) {
	contract, err := bindBasicNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicNftCaller{contract: contract}, nil
}

// NewBasicNftTransactor creates a new write-only instance of BasicNft, bound to a specific deployed contract.
func NewBasicNftTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicNftTransactor, error) {
	contract, err := bindBasicNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicNftTransactor{contract: contract}, nil
}

// NewBasicNftFilterer creates a new log filterer instance of BasicNft, bound to a specific deployed contract.
func NewBasicNftFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicNftFilterer, error) {
	contract, err := bindBasicNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicNftFilterer{contract: contract}, nil
}

// bindBasicNft binds a generic wrapper to an already deployed contract.
func bindBasicNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BasicNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicNft *BasicNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicNft.Contract.BasicNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicNft *BasicNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicNft.Contract.BasicNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicNft *BasicNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicNft.Contract.BasicNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicNft *BasicNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicNft *BasicNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicNft *BasicNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicNft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BasicNft *BasicNftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasicNft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BasicNft *BasicNftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BasicNft.Contract.BalanceOf(&_BasicNft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BasicNft *BasicNftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BasicNft.Contract.BalanceOf(&_BasicNft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BasicNft *BasicNftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BasicNft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BasicNft *BasicNftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BasicNft.Contract.GetApproved(&_BasicNft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BasicNft *BasicNftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BasicNft.Contract.GetApproved(&_BasicNft.CallOpts, tokenId)
}

// GetTokenCounter is a free data retrieval call binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() view returns(uint256)
func (_BasicNft *BasicNftCaller) GetTokenCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasicNft.contract.Call(opts, &out, "getTokenCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenCounter is a free data retrieval call binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() view returns(uint256)
func (_BasicNft *BasicNftSession) GetTokenCounter() (*big.Int, error) {
	return _BasicNft.Contract.GetTokenCounter(&_BasicNft.CallOpts)
}

// GetTokenCounter is a free data retrieval call binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() view returns(uint256)
func (_BasicNft *BasicNftCallerSession) GetTokenCounter() (*big.Int, error) {
	return _BasicNft.Contract.GetTokenCounter(&_BasicNft.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BasicNft *BasicNftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BasicNft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BasicNft *BasicNftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BasicNft.Contract.IsApprovedForAll(&_BasicNft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BasicNft *BasicNftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BasicNft.Contract.IsApprovedForAll(&_BasicNft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BasicNft *BasicNftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BasicNft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BasicNft *BasicNftSession) Name() (string, error) {
	return _BasicNft.Contract.Name(&_BasicNft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BasicNft *BasicNftCallerSession) Name() (string, error) {
	return _BasicNft.Contract.Name(&_BasicNft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BasicNft *BasicNftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BasicNft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BasicNft *BasicNftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BasicNft.Contract.OwnerOf(&_BasicNft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BasicNft *BasicNftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BasicNft.Contract.OwnerOf(&_BasicNft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BasicNft *BasicNftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BasicNft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BasicNft *BasicNftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BasicNft.Contract.SupportsInterface(&_BasicNft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BasicNft *BasicNftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BasicNft.Contract.SupportsInterface(&_BasicNft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BasicNft *BasicNftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BasicNft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BasicNft *BasicNftSession) Symbol() (string, error) {
	return _BasicNft.Contract.Symbol(&_BasicNft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BasicNft *BasicNftCallerSession) Symbol() (string, error) {
	return _BasicNft.Contract.Symbol(&_BasicNft.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BasicNft *BasicNftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _BasicNft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BasicNft *BasicNftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BasicNft.Contract.TokenURI(&_BasicNft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BasicNft *BasicNftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BasicNft.Contract.TokenURI(&_BasicNft.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BasicNft *BasicNftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BasicNft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BasicNft *BasicNftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BasicNft.Contract.Approve(&_BasicNft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BasicNft *BasicNftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BasicNft.Contract.Approve(&_BasicNft.TransactOpts, to, tokenId)
}

// MintNft is a paid mutator transaction binding the contract method 0xda86f4a9.
//
// Solidity: function mintNft(string _tokenURI) returns()
func (_BasicNft *BasicNftTransactor) MintNft(opts *bind.TransactOpts, _tokenURI string) (*types.Transaction, error) {
	return _BasicNft.contract.Transact(opts, "mintNft", _tokenURI)
}

// MintNft is a paid mutator transaction binding the contract method 0xda86f4a9.
//
// Solidity: function mintNft(string _tokenURI) returns()
func (_BasicNft *BasicNftSession) MintNft(_tokenURI string) (*types.Transaction, error) {
	return _BasicNft.Contract.MintNft(&_BasicNft.TransactOpts, _tokenURI)
}

// MintNft is a paid mutator transaction binding the contract method 0xda86f4a9.
//
// Solidity: function mintNft(string _tokenURI) returns()
func (_BasicNft *BasicNftTransactorSession) MintNft(_tokenURI string) (*types.Transaction, error) {
	return _BasicNft.Contract.MintNft(&_BasicNft.TransactOpts, _tokenURI)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BasicNft *BasicNftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BasicNft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BasicNft *BasicNftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BasicNft.Contract.SafeTransferFrom(&_BasicNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BasicNft *BasicNftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BasicNft.Contract.SafeTransferFrom(&_BasicNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BasicNft *BasicNftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BasicNft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BasicNft *BasicNftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BasicNft.Contract.SafeTransferFrom0(&_BasicNft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_BasicNft *BasicNftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _BasicNft.Contract.SafeTransferFrom0(&_BasicNft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BasicNft *BasicNftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _BasicNft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BasicNft *BasicNftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BasicNft.Contract.SetApprovalForAll(&_BasicNft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BasicNft *BasicNftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BasicNft.Contract.SetApprovalForAll(&_BasicNft.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BasicNft *BasicNftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BasicNft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BasicNft *BasicNftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BasicNft.Contract.TransferFrom(&_BasicNft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BasicNft *BasicNftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BasicNft.Contract.TransferFrom(&_BasicNft.TransactOpts, from, to, tokenId)
}

// BasicNftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BasicNft contract.
type BasicNftApprovalIterator struct {
	Event *BasicNftApproval // Event containing the contract specifics and raw log

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
func (it *BasicNftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicNftApproval)
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
		it.Event = new(BasicNftApproval)
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
func (it *BasicNftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicNftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicNftApproval represents a Approval event raised by the BasicNft contract.
type BasicNftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BasicNft *BasicNftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BasicNftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BasicNft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BasicNftApprovalIterator{contract: _BasicNft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BasicNft *BasicNftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BasicNftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BasicNft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicNftApproval)
				if err := _BasicNft.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BasicNft *BasicNftFilterer) ParseApproval(log types.Log) (*BasicNftApproval, error) {
	event := new(BasicNftApproval)
	if err := _BasicNft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasicNftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the BasicNft contract.
type BasicNftApprovalForAllIterator struct {
	Event *BasicNftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BasicNftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicNftApprovalForAll)
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
		it.Event = new(BasicNftApprovalForAll)
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
func (it *BasicNftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicNftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicNftApprovalForAll represents a ApprovalForAll event raised by the BasicNft contract.
type BasicNftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BasicNft *BasicNftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BasicNftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BasicNft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BasicNftApprovalForAllIterator{contract: _BasicNft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BasicNft *BasicNftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BasicNftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BasicNft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicNftApprovalForAll)
				if err := _BasicNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BasicNft *BasicNftFilterer) ParseApprovalForAll(log types.Log) (*BasicNftApprovalForAll, error) {
	event := new(BasicNftApprovalForAll)
	if err := _BasicNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasicNftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicNft contract.
type BasicNftTransferIterator struct {
	Event *BasicNftTransfer // Event containing the contract specifics and raw log

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
func (it *BasicNftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicNftTransfer)
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
		it.Event = new(BasicNftTransfer)
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
func (it *BasicNftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicNftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicNftTransfer represents a Transfer event raised by the BasicNft contract.
type BasicNftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BasicNft *BasicNftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BasicNftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BasicNft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BasicNftTransferIterator{contract: _BasicNft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BasicNft *BasicNftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicNftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BasicNft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicNftTransfer)
				if err := _BasicNft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BasicNft *BasicNftFilterer) ParseTransfer(log types.Log) (*BasicNftTransfer, error) {
	event := new(BasicNftTransfer)
	if err := _BasicNft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
