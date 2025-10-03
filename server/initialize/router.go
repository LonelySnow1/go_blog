package initialize

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/router"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()

	routerGroup := router.RouterGroupApp

	// 公共路由
	publicGroup := Router.Group(global.Config.System.RouterPrefix)
	{
		routerGroup.InitBaseRouter(publicGroup)
	}
	return Router
}
