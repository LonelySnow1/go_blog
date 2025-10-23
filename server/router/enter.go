package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	ImageRouter
	ArticleRouter
	CommentRouter
	AdvertisementRouter
	FriendLinkRouter
	FeedbackRouter
	WebsiteRouter
}

var RouterGroupApp = new(RouterGroup)
