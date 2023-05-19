package account


import (
	"golddigger/app/user/apis"

	"github.com/gin-gonic/gin"
)
func InitAccountRouter(Router *gin.RouterGroup) {
	accountRouter := Router.Group("account")
	accountApi := apis.UserApi.AccountApiGroup.AccountApi
	{
		accountRouter.POST("getSignMsg", accountApi.GetRandomSignMsg)

	}
}
