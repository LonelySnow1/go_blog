package other

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"gorm.io/gorm"
	"server/model/request"
)

type MySQLOption struct {
	request.PageInfo
	Order   string
	Where   *gorm.DB
	Preload []string
}

type EsOption struct {
	request.PageInfo
	Index          string
	Request        *search.Request
	SourceIncludes []string
}
