package global

import (
	"time"

	"gorm.io/gorm"
)

type COMMON_MODEL struct {
	ID        uint           `json:"id" gorm:"primarykey"`  // 主键ID
	CreatedAt time.Time      `json:"createAt"`				// 创建时间
	UpdatedAt time.Time      `json:"updateAt"`				// 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` 		// 删除时间
}
