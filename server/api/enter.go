package api

import "server/service"

type ApiGroup struct {
	BaseApi
	UserApi
}

var ApiGroupApp = new(ApiGroup)

var baseService = service.ServiceGroupApp.BaseService
var userService = service.ServiceGroupApp.userService
