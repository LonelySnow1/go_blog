// 适用于windows系统
//go:build windows
// +build windows

package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// initServer 函数初始化一个标准的 HTTP 服务器（适用于 Windows 系统）
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,          // 设置服务器监听的地址
		Handler:        router,           // 设置请求处理器（路由）
		ReadTimeout:    10 * time.Minute, // 设置请求的读取超时时间为 10 分钟
		WriteTimeout:   10 * time.Minute, // 设置响应的写入超时时间为 10 分钟
		MaxHeaderBytes: 1 << 20,          // 设置最大请求头的大小（1MB）
	}
}
