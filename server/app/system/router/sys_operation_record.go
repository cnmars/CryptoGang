package router

import (
	v1 "golddigger/app/system/apis"
	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func /* (s *OperationRecordRouter) */ initSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	authorityMenuApi := v1.SystemApis.OperationRecordApi
	{
		operationRecordRouter.POST("createSysOperationRecord", authorityMenuApi.CreateSysOperationRecord)             // 新建SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecord", authorityMenuApi.DeleteSysOperationRecord)           // 删除SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", authorityMenuApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		operationRecordRouter.GET("findSysOperationRecord", authorityMenuApi.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
		operationRecordRouter.GET("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)            // 获取SysOperationRecord列表

	}
}
