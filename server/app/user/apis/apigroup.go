package apis
import (
	"golddigger/app/user/apis/common"
	"golddigger/app/user/apis/account"
)

type UserApiGroup struct {
	CommonApiGroup common.CommonApiGroup
	AccountApiGroup account.AccountGroup
	// .....
}

var UserApi = new(UserApiGroup)
