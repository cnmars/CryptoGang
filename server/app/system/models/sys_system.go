package models

import (
	"golddigger/app/system/models/config"
)

// 配置文件结构体
type System struct {
	Config config.ConfigCollection `json:"config"`
}
