package router

import (
	v1 "golddigger/app/system/apis"

	"golddigger/middleware"

	"github.com/gin-gonic/gin"
)

type FeedbackRouter struct{}

func /* (s *FeedbackRouter)  */initFeedbackRouter(Router *gin.RouterGroup) {
	fbRouter := Router.Group("feedback").Use(middleware.OperationRecord())
	fbRouterWithoutRecord := Router.Group("feedback")
	fbRouterApi := v1.SystemApis.FeedbackApi
	{
		fbRouter.POST("createFb", fbRouterApi.CreateFeedback)               // 创建
	}
	{
		fbRouterWithoutRecord.POST("deleteFb", fbRouterApi.DeleteById)      // 删除
		fbRouterWithoutRecord.POST("setRead", fbRouterApi.SetReaded)        // 设置为已读
		fbRouterWithoutRecord.POST("getFbList", fbRouterApi.GetApiList)     // 获取列表
	}
}
