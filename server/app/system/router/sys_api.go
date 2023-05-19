package router

import (
	v1 "golddigger/app/system/apis"

	"golddigger/middleware"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func /* (s *ApiRouter)  */initApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")
	apiRouterApi := v1.SystemApis.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除Api
		apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有api
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
