package database

import (
	"github.com/gofrs/uuid"
	"server/global"
	"server/model/appTypes"
)

type User struct {
	global.MODEL
	UUID      uuid.UUID         `json:"uuid" gorm:"type:char(36);unique"`
	Username  string            `json:"username"`
	Password  string            `json:"-"`
	Email     string            `json:"email"`
	Openid    string            `json:"openid"`
	Avatar    string            `json:"avatar" gorm:"size:255"` // 头像
	Address   string            `json:"address"`
	Signature string            `json:"signature" gorm:"default:'这里怎么什么都没有喵？'"`
	RoleID    appTypes.RoleID   `json:"role_id"`  // 角色ID，用于权限控制
	Register  appTypes.Register `json:"register"` //注册来源
	Freeze    bool              `json:"freeze"`   //是否冻结
}
