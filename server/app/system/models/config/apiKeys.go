package config

type ApiKeys struct{
	OpenSeaKey  string   `mapstructure:"opensea-key" json:"opensea-key" yaml:"opensea-key"`
	EtherScanKey  string   `mapstructure:"etherscan-key" json:"etherscan-key" yaml:"etherscan-key"`
	ProviderUrlHttps  string   `mapstructure:"provider-url-https" json:"provider-url-https" yaml:"provider-url-https"`
	ProviderUrlWs  string   `mapstructure:"provider-url-ws" json:"provider-url-ws" yaml:"provider-url-ws"`
	BlockNativeKey  string   `mapstructure:"blocknative-key" json:"blocknative-key" yaml:"blocknative-key"`
}