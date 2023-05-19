package common

import (
	"golddigger/app/user/apis"

	"github.com/gin-gonic/gin"
)

func InitWalletRouter(Router *gin.RouterGroup) {
	commonRouter := Router.Group("wallet")
	commonApi := apis.UserApi.CommonApiGroup.WalletApi
	{
		commonRouter.POST("createCommon", commonApi.GenerateCommonWallet)
		commonRouter.POST("getbalance", commonApi.GetBalance)

	}
}
