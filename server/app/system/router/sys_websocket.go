package router

import (
	 "golddigger/app/system/apis"
	"github.com/gin-gonic/gin"
)


func initWebsocketRouter(Router *gin.RouterGroup) {
	Router.GET("ws",apis.WsHandler)
}
