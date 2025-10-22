package service

type ServiceGroup struct {
	EsService
	BaseService
	JwtService
	GaodeService
	UserService
	QQService
	ImageService
	ArticleService
}

var ServiceGroupApp = new(ServiceGroup)
