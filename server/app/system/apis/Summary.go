package apis

import "golddigger/app/system/service"

type SystemApisSummary struct {
	CasbinApi
	UserApi
	JwtApi
	AuthorityApi
	OperationRecordApi
	SystemApi
	SystemApiApi
	AuthorityMenuApi
	FeedbackApi
}

var SystemApis = new(SystemApisSummary)
var (
	jwtService             = service.SystemService.JwtService
	baseMenuService        = service.SystemService.BaseMenuService
	menuService            = service.SystemService.MenuService
	apiService             = service.SystemService.ApiService
	userService            = service.SystemService.UserService
	casbinService          = service.SystemService.CasbinService
	authorityService       = service.SystemService.AuthorityService
	systemConfigService    = service.SystemService.SystemConfigService
	operationRecordService = service.SystemService.OperationRecordService
	feedbackService 	   = service.SystemService.FeedbackService
)
