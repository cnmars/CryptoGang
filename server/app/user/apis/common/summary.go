package common

import (
	"golddigger/app/user/service"
)

/*
*
*ApiGroup方便外部router层调用，对外
var里的变量是方便在api层调用，对内
*/
type CommonApiGroup struct {
	OperationRecordApi
	WalletApi
	//....
}

var (
	operationRecordService = service.UserService.CommonService.OperationRecordService
	walletService = service.UserService.CommonService.WalletService
	accountService = service.UserService.AccountService.LoginService
)
