package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	ImageRouter
	ArticleRouter
}

var RouterGroupApp = new(RouterGroup)
