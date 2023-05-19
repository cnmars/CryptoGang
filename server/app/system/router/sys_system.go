package router

import (
	v1 "golddigger/app/system/apis"
	"golddigger/middleware"
	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func /* (s *SysRouter) */ initSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	systemApi := v1.SystemApis.SystemApi
	{
		// sysRouter.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
		// sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("getServerInfo", systemApi.GetServerInfo)     // 获取服务器信息
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem)       // 重启服务
		sysRouter.POST("getEthPrice", systemApi.FetchEthPriceOnece)  // 获取比价
		sysRouter.POST("getUpgradePrice", systemApi.GetUpgradePrice)  // 获取比价
	}
}
