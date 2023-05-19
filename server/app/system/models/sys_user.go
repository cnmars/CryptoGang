package models

import (
	"golddigger/global"
	"time"
	"gorm.io/gorm"
	// "github.com/satori/go.uuid"
)

type SysUser struct {
	global.COMMON_MODEL
	Address     		string         `json:"address" gorm:"index;comment:用户钱包地址"`                                 // 用户UUID
	Ens    				string         `json:"ens" gorm:"comment:用户侧边主题"`                                          // 用户Ens域名
	HttpProvider    	string         `json:"httpProvider" gorm:"comment:用户提供的http节点"`                            // 用户Ens域名
	WssProvider    		string         `json:"wssProvider" gorm:"comment:用户提供的wss节点"`                              // 用户Ens域名
	UseDefaultHttp    	bool         	`json:"useDefaultHttp" gorm:"default:1;comment:使用系统默认"`                     // 用户Ens域名
	UseDefaultWss    	bool         	`json:"useDefaultWss" gorm:"default:1;comment:使用系统默认"`                      // 用户Ens域名
	RecentLogin 		time.Time      `json:"recentLogin" gorm:"comment:上次登录时间"`  								// 创建时间
	AuthorityId 		uint           `json:"authorityId" gorm:"default:333;comment:用户角色ID"`                       // 用户角色ID
	SiteTheme 			string         `json:"siteTheme" gorm:"default:dark;comment:用户网站主题"`                          // 用户角色ID
	UpgradeAt 			time.Time      `json:"upgradeAt" gorm:"comment:vip充值日期"`                                        // 用户角色ID
	ExpireAt 			time.Time      `json:"expireAt" gorm:"comment:vip失效日期"`                                        // 用户角色ID
	Authority   		SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities 		[]SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Enable      		int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}
//登录时查询表，查询完成之后把登录时间设置为当前
func (u *SysUser) AfterFind(tx *gorm.DB) (err error) {
	tx.Model(&SysUser{}).Where("address = ?", u.Address).Update("recent_login", time.Now())
	return
  }