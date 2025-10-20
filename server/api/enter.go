package api

import "server/service"

type ApiGroup struct {
	BaseApi
	UserApi
	ImageApi
}

var ApiGroupApp = new(ApiGroup)

var baseService = service.ServiceGroupApp.BaseService
var userService = service.ServiceGroupApp.UserService
var jwtService = service.ServiceGroupApp.JwtService
var qqService = service.ServiceGroupApp.QQService
var imageService = service.ServiceGroupApp.ImageService
