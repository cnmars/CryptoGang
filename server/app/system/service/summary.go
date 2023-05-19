package service

type SystemServiceSummary struct {
	JwtService
	BaseMenuService
	MenuService
	ApiService
	UserService
	CasbinService
	AuthorityService
	SystemConfigService
	OperationRecordService
	FeedbackService
}

var SystemService = new(SystemServiceSummary)
