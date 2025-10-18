package other

import (
	"gorm.io/gorm"
	"server/model/request"
)

type MySQLOption struct {
	request.PageInfo
	Order   string
	Where   *gorm.DB
	Preload []string
}
