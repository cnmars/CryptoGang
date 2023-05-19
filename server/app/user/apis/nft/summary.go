package system

import (
	"golddigger/app/user/service"
	"golddigger/app/user/apis/common"
)

/**
*ApiGroup方便外部router层调用，对外
var里的变量是方便在api层调用，对内
*/
type ApiGroup struct {
	
	OperationRecordApi
	
}

var (
	operationRecordService  = service.CommonService.OperationRecordService
)
