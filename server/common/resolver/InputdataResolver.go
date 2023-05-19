package resolver

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/go-resty/resty/v2"
	"log"
	// "strings"
)

type (
	RawABIResponse struct {
		Status  *string `json:"status"`
		Message *string `json:"message"`
		Result  *string `json:"result"`
	}
)
func DecodeTransactionInputData(contractABI *abi.ABI, data []byte)(method *abi.Method,inputsMap map[string]interface{},err error) {
	if len(data)>0{
		methodSigData := data[:4]
		inputsSigData := data[4:]
		method, err = contractABI.MethodById(methodSigData)
		inputsMap = make(map[string]interface{})
		if inputsSigData !=nil{
			if method!=nil{

				err = method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); 
			}
		}
	}
	return method,inputsMap,err
	// fmt.Printf("Method Name: %s\n", method.Name)
	// fmt.Printf("Method inputs: %v\n", inputsMap)
}

func GetTransactionMessage(tx *types.Transaction) types.Message {
	msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}

func DecodeTransactionLogs(receipt *types.Receipt, contractABI *abi.ABI) {
	for _, vLog := range receipt.Logs {
		// topic[0] is the event name
		event, err := contractABI.EventByID(vLog.Topics[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Event Name: %s\n", event.Name)
		// topic[1:] is other indexed params in event
		if len(vLog.Topics) > 1 {
			for i, param := range vLog.Topics[1:] {
				fmt.Printf("Indexed params %d in hex: %s\n", i, param)
				fmt.Printf("Indexed params %d decoded %s\n", i, common.HexToAddress(param.Hex()))
			}
		}
		if len(vLog.Data) > 0 {
			fmt.Printf("Log Data in Hex: %s\n", hex.EncodeToString(vLog.Data))
			outputDataMap := make(map[string]interface{})
			err = contractABI.UnpackIntoMap(outputDataMap, event.Name, vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Event outputs: %v\n", outputDataMap)
		}
	}
}


func GetTransactionReceipt(client *ethclient.Client, txHash common.Hash) *types.Receipt {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	return receipt
}

// func test() {
// 	// get etherscanAPIKEY from https://docs.etherscan.io/getting-started/viewing-api-usage-statistics
// 	const etherscanAPIKEY = "M3SF4WTDC4NWQIIVNAZDFXBW1SW49QWDNZ"
// 	const providerUrl = "https://ropsten.infura.io/v3/28d5693e8bee4b58a61f0c627d62331e"

// 	client, err := ethclient.Dial(providerUrl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// https://ropsten.etherscan.io/tx/0x7e605f68ff30509eb2bf3238936ef65a01bfa25243488c007244aabe645d0ec9
// 	txHash := common.HexToHash("0x7e605f68ff30509eb2bf3238936ef65a01bfa25243488c007244aabe645d0ec9")
// 	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("tx isPending: %t\n", isPending)
// 	ParseTransactionBaseInfo(tx)
// 	contractABI := GetContractABI(tx.To().String(), etherscanAPIKEY)
// 	DecodeTransactionInputData(contractABI, tx.Data())
// 	receipt := GetTransactionReceipt(client, txHash)
// 	DecodeTransactionLogs(receipt, contractABI)
// }