package router

import (
	v1 "golddigger/app/system/apis"
	"golddigger/middleware"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct{}

func/*  (s *CasbinRouter)  */initCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouterWithoutRecord := Router.Group("casbin")
	casbinApi := v1.SystemApis.CasbinApi
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
