package router

import (
	v1 "golddigger/app/system/apis"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func /* (s *BaseRouter)  */ initBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.SystemApis.UserApi
	{
		baseRouter.POST("getNonce", baseApi.GetRandomSignMsg)
		baseRouter.POST("login", baseApi.SignAndLogin)
		// baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
