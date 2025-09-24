package initialize

import (
	"github.com/gin-gonic/gin"
	"server/global"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	// TODO
	return Router
}
