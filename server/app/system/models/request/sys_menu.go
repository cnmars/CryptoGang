package request

import (
	"golddigger/global"
	system"golddigger/app/system/models"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []system.SysBaseMenu `json:"menus"`
	AuthorityId uint                 `json:"authorityId"` // 角色ID
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		COMMON_MODEL: global.COMMON_MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "views/dashboard/index.vue",
		Sort:      1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
