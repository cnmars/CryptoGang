package models

import (
	"golddigger/global"
)

type FeedBack struct {
	global.COMMON_MODEL
	Address string `json:"address" gorm:"comment:用户钱包地址"`
	Phone   string `json:"phone" gorm:"comment:用户电话"`
	Topic   string `json:"topic" gorm:"comment:反馈主题"`
	Type    uint16 `json:"type" gorm:"default:1;comment:反馈类型"` //1：意见建议、2：系统bug
	Content string `json:"content" gorm:"type:text;comment:反馈内容"`
	Status  uint `json:"status" gorm:"default:2;comment:状态"` //2：未读，1：已读
}

func (FeedBack) TableName() string {
	return "sys_feedback"
}
