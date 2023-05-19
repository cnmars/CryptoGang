package request

import (
	"golddigger/app/system/models"
	"golddigger/app/system/models/request/common"
	
)

type SysOperationRecordSearch struct {
	models.SysOperationRecord
	common.PageInfo
}
