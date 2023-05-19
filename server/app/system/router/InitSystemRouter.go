package router

import (
	"github.com/gin-gonic/gin"
)

func InitSystemModuleRouter(Router *gin.RouterGroup) {

	r := Router.Group("system")
	initCasbinRouter(r)
	
	// initAuthorityBtnRouterRouter(r)
	initAuthorityRouter(r)
	initJwtRouter(r)
	// InitMenuRouter(r)
	initSystemRouter(r)
	initSysOperationRecordRouter(r)
	initUserRouter(r)
	initMenuRouter(r)
	initFeedbackRouter(r)
	initWebsocketRouter(r)
	// common.InitCommonRouter(r)
}
func InitSystemPublicBasicRouter(Router *gin.RouterGroup){
	r := Router.Group("system")
	initBaseRouter(r)
}
