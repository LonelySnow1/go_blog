package database

import (
	"github.com/gofrs/uuid"
	"server/model/appTypes"
)

type User struct {
	UUID      uuid.UUID         `json:"uuid" gorm:"type:char(36);unique"`
	Username  string            `json:"username"`
	Password  string            `json:"-"`
	Email     string            `json:"email"`
	Openid    string            `json:"openid"`
	Avatar    string            `json:"avatar" gorm:"size:255"`
	Address   string            `json:"address"`
	Signature string            `json:"signature" gorm:"default:'这里怎么什么都没有喵？'"`
	RoleID    appTypes.RoleID   `json:"role_id"`
	Register  appTypes.Register `json:"register"`
	Freeze    bool              `json:"freeze"`
}
