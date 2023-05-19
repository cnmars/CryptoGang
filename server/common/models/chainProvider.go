package models

import(
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"golddigger/common/etherscan"
	// "github.com/ethereum/go-ethereum/rpc"
)
type ChainProviders struct{
	ProviderHttps   *ethclient.Client
	ProviderWss		*gethclient.Client
	EtherscanClient *etherscan.Client
}