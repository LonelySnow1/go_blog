package global

import (
	"gorm.io/gorm"
	"time"
)

type MODEL struct { // 自动添加时间信息
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index"` //软删除
}
