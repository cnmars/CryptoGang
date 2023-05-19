package models

import (
	"golddigger/global"
)

type JwtBlacklist struct {
	global.COMMON_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
