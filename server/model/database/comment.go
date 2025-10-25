package database

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"server/global"
	"server/model/elasticsearch"
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

// AfterCreate 钩子，创建后调用
func (c *Comment) AfterCreate(_ *gorm.DB) error {
	source := "ctx._source.comments += 1"
	script := types.Script{Source: &source, Lang: &scriptlanguage.Painless}
	_, err := global.ESClient.Update(elasticsearch.ArticleIndex(), c.ArticleID).Script(&script).Do(context.TODO())
	return err
}

// AfterDelete 钩子，删除后调用
func (c *Comment) BeforeDelete(_ *gorm.DB) error {
	var articleID string
	if err := global.DB.Model(&c).Pluck("article_id", &articleID).Error; err != nil {
		return err
	}
	source := "ctx._source.comments -= 1"
	script := types.Script{Source: &source, Lang: &scriptlanguage.Painless}
	_, err := global.ESClient.Update(elasticsearch.ArticleIndex(), articleID).Script(&script).Do(context.TODO())
	return err
}
