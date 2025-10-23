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
	CommentService
	AdvertisementService
	FriendLinkService
}

var ServiceGroupApp = new(ServiceGroup)
