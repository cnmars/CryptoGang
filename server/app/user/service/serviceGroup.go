package service

import(
	"golddigger/app/user/service/common"
	"golddigger/app/user/service/account"
)
type UserServiceGroup struct {
	CommonService common.CommonServiceGroup
	AccountService account.AccountServiceGroup
}

var UserService = new(UserServiceGroup)
