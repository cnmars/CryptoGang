package common

import (
	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/sha3"
)

type WalletApi struct{}

func (w *WalletApi) GetBalance(c *gin.Context) {
	add := c.DefaultQuery("address", "0x1c4E5c4c6F8AaDF5F3B90EcE84F750f7293c188a")
	if bal,err :=walletService.GetBalance(add);err==nil{
		c.JSON(200, gin.H{
			"balance":  bal,
		})
	}
}
func (w *WalletApi) GenerateCommonWallet(c *gin.Context) {
	wallet, err := walletService.GenerateCommonWallet()
	if err == nil {
		// if w, ok := wallet.(*models.Wallet); ok {
		c.JSON(200, gin.H{
			"publicKey":  wallet.PublicKey,
			"address":    wallet.Address,
			"privateKey": wallet.PrivateKey,
		})
		// }
	}
}

func ATOMCreateHDWallet(c *gin.Context) {

}
func ETHGenerateHDWallet(c *gin.Context) {

}

func CreateKs(c *gin.Context) {

}
