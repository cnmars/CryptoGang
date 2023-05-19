package router

import (
	"golddigger/app/user/router/account"
	"golddigger/app/user/router/common"

	"github.com/gin-gonic/gin"
)

func InitUserPublicRouter(Router *gin.RouterGroup) {

	r := Router.Group("user")
	common.InitCommonRouter(r)
	account.InitAccountRouter(r)
}
func InitUserPrivateRouter(Router *gin.RouterGroup) {

	// r := Router.Group("user")
	// common.InitCommonRouter(r)
	// account.InitAccountRouter(r)
}
