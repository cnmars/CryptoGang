package common

import (

	"github.com/gin-gonic/gin"
)



func InitCommonRouter(Router *gin.RouterGroup) {
	InitWalletRouter(Router)
	InitSysOperationRecordRouter(Router)
}