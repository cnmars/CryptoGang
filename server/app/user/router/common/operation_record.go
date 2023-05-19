package common

import (
	"golddigger/app/user/apis"

	"github.com/gin-gonic/gin"
)

func InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	authorityMenuApi := apis.UserApi.CommonApiGroup.OperationRecordApi
	{
		operationRecordRouter.POST("createSysOperationRecord", authorityMenuApi.CreateSysOperationRecord)             // 新建SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecord", authorityMenuApi.DeleteSysOperationRecord)           // 删除SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", authorityMenuApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		operationRecordRouter.GET("findSysOperationRecord", authorityMenuApi.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
		operationRecordRouter.GET("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)            // 获取SysOperationRecord列表

	}
}
