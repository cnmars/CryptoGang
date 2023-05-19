package account

import (
	"golddigger/app/user/service"
)

/*
*
*ApiGroup方便外部router层调用，对外
var里的变量是方便在api层调用，对内
*/
type AccountGroup struct {
	AccountApi
	//....
}

var (
	accountService = service.UserService.AccountService.LoginService
)
