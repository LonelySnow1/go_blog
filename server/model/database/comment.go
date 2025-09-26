package database

import (
	"github.com/gofrs/uuid"
	"server/global"
)

// Comment 评论表
type Comment struct {
	global.MODEL
	ArticleID string    `json:"article_id"`                                      // 文章 ID  文章是存储到ES里的
	PID       *uint     `json:"p_id"`                                            // 父评论 ID 这里用的是指针类型，代表有可能是nil
	PComment  *Comment  `json:"-" gorm:"foreignKey:PID"`                         // 1对1 关联自己的父评论
	Children  []Comment `json:"children" gorm:"foreignKey:PID"`                  // 1对多 子评论 关联自己的子评论
	UserUUID  uuid.UUID `json:"user_uuid" gorm:"type:char(36)"`                  // 用户 uuid
	User      User      `json:"user" gorm:"foreignKey:UserUUID;references:UUID"` // 关联的用户
	Content   string    `json:"content"`                                         // 内容
}

// TODO 创建和删除评论时需要更新文章评论数
