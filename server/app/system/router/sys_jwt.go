package router

import (
	v1 "golddigger/app/system/apis"

	"github.com/gin-gonic/gin"
)

type JwtRouter struct{}

func initJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt")
	jwtApi := v1.SystemApis.JwtApi
	{
		jwtRouter.POST("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
