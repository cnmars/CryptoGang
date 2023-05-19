
 package etherscan

 const (
	 Mainnet Network = "api"
	 Goerli Network = "api-goerli"
 )
 
 type Network string
 
 func (n Network) SubDomain() (sub string) {
	 return string(n)
 }