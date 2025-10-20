package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	ImageRouter
}

var RouterGroupApp = new(RouterGroup)
