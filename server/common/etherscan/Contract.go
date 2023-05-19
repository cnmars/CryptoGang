package etherscan

import (
	"strings"
	// "fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
)
//获取合约创建者和及其交易，一次可查五个合约，地址以逗号分开
func (c *Client) ContractCreator(address string) (abi string, err error) {
	param := M{
		"address": address,
	}

	err = c.call("contractaddresses", "getcontractcreation", param, &abi)
	// fmt.Println("contract abi:",abi)
	return
}
//原始abi
func (c *Client) ContractRawABI(address string) (abi string, err error) {
	param := M{
		"address": address,
	}

	err = c.call("contract", "getabi", param, &abi)
	// fmt.Println("contract abi:",abi)
	return
}
//源码
func (c *Client) ContractSource(address string) (source []ContractSource, err error) {
	param := M{
		"address": address,
	}

	err = c.call("contract", "getsourcecode", param, &source)
	return
}
//abi对象
func (c *Client) GetContractABIFromRaw(rawAbi string) (contractABI abi.ABI, err error) {
	contractABI, err = abi.JSON(strings.NewReader(rawAbi))
	return contractABI, err
}

