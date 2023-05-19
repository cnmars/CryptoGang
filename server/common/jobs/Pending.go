package jobs

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golddigger/common/resolver"
	"golddigger/common/wsserver"
	"golddigger/global"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type PendingData struct {
	ProjectName string
	Address     string
	TxHash      string
	WebSite     string
	Number      int
	//单价
	Price   string
	Icon    string
	Twitter string
	//其他很多链接都和address相关，在前端拼接
}

//GasLimit对应Gas
//MaxFee对应GasFeeCap
//GasTip对应GasTipCap
// var(
// 	backend, err = ethclient.Dial(global.GD_CONFIG.ApiKeys.ProviderUrlHttps)
// 	etherClient = etherscan.New(etherscan.Mainnet, global.GD_CONFIG.ApiKeys.EtherScanKey)
// 	rpcCli, _ = rpc.Dial(wss)
// 	gcli      = gethclient.New(rpcCli)
// )

func TestSpiltPend() {
	var (
		TxHashCh = make(chan common.Hash, 600)
		TxCh     = make(chan types.Message, 600)
	)
	go GetPendingHash(TxHashCh)
	go GetTxByHash(TxHashCh, TxCh)
	go GetTxByHash(TxHashCh, TxCh)
	go GetTxByHash(TxHashCh, TxCh)
	go GetTxByHash(TxHashCh, TxCh)
	go GetTxByHash(TxHashCh, TxCh)
	go Consumer(TxCh)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}

func GetPendingHash(out chan<- common.Hash) {

	global.GD_PROVIDERS.ProviderWss.SubscribePendingTransactions(context.Background(), out)
	fmt.Println("1")

}

func GetTxByHash(in <-chan common.Hash, out chan<- types.Message) {
	chainID, _ := global.GD_PROVIDERS.ProviderHttps.NetworkID(context.Background())
	for {
		select {
		case txhash := <-in:
			tx, _, err := global.GD_PROVIDERS.ProviderHttps.TransactionByHash(context.Background(), txhash)
			// fmt.Println(txhash)
			// fmt.Println(err)
			if err == nil {
				if msg, err2 := tx.AsMessage(types.LatestSignerForChainID(chainID), tx.GasPrice()); err == nil {
					// if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), tx.GasPrice()); err == nil {
					// fmt.Println("2")
					if err2 == nil {

						out <- msg
					} else {
						fmt.Println("err2->", err2)
					}
					// fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
				}

			} else {
				fmt.Println("txhash:", txhash, err)
			}
		}
	}
}
func Consumer(in <-chan types.Message) {
	fmt.Println("3")
	for {
		select {
		case tx := <-in:
			fmt.Println("==========================================")
			fmt.Println("from-->", tx.From().String())
			fmt.Println("to-->", tx.To().String())
			fmt.Println("value-->", tx.Value())
		}
	}
}

func FilterMintTx(in <-chan common.Hash) {
	for {
		select {
		case txhash := <-in:
			tx, isPending, _ := global.GD_PROVIDERS.ProviderHttps.TransactionByHash(context.Background(), txhash)
			// data, _ := tx.MarshalJSON()
			//get abi to parse method and filter mint method
			abiRaw, _ := global.GD_PROVIDERS.EtherscanClient.ContractRawABI(tx.To().String())
			abi, _ := global.GD_PROVIDERS.EtherscanClient.GetContractABIFromRaw(abiRaw)
			// method,inputsMap,_:= resolver.DecodeTransactionInputData(&abi, tx.Data())
			method, _, _ := resolver.DecodeTransactionInputData(&abi, tx.Data())
			if method != nil && (strings.Contains(method.Name, "mint") || strings.Contains(method.Name, "Mint")) {

				fmt.Println("==========================================")
				fmt.Println("is pending:", isPending)
				fmt.Println("method name:", method.Name)
				// if inputsMap!=nil{

				// 	fmt.Println("inputsMap:",inputsMap)
				// }
			}
			// fmt.Println("from:",tx.To().String())
			// fmt.Println("inputsMap:",inputsMap)
		}
	}
}
func /* (ps *wsserver.PubSub)  */ FetchPending() {
	// gcli.SubscribePendingTransactions(context.Background(), TxHashCh)
	for !global.GD_WS.IsStop {
		m := wsserver.WsPendingData{
			ProjectName: "nft1",
			Icon:        "aaa",
		}
		global.GD_WS.ThirdPartData <- m.ToCommonData()
		// time.Sleep(time.Second * 2)

	}
}
// func /*  (ps *wsserver.PubSub)  */ FetchGas() {
// 	for !global.GD_WS.IsStop {
// 		m := wsserver.WsGasData{
// 			Data:    "1000",
// 		}
// 		global.GD_WS.ThirdPartData <- m.ToCommonData()
// 		// time.Sleep(time.Second * 2)
// 	}
// }
// func /*  (ps *wsserver.PubSub)  */ FetchPrice() {
// 	for !global.GD_WS.IsStop {
// 		m := wsserver.WsGasData{
// 			Data: "1000",
// 		}
// 		global.GD_WS.ThirdPartData <- m.ToCommonData()
// 		// time.Sleep(time.Second * 2)
// 	}
// }
