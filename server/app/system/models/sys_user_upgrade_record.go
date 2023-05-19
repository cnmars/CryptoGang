package models

import (
	"golddigger/global"
	"time"
)

type UpgradeRecord struct {
	global.COMMON_MODEL
	Address 		string `json:"address" gorm:"comment:用户钱包地址"`
	TranscationHash string `json:"hash" gorm:"comment:充值记录"`
	ExpireAt 		time.Time  `json:"expireAt" gorm:"comment:失效日期"`
	Type    		uint16 `json:"type" gorm:"default:1;comment:充值类型"` //1：意见建议、2：系统bug
}

func (UpgradeRecord) TableName() string {
	return "sys_upgrade_record"
}
