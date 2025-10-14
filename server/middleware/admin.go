package middleware

import (
	"github.com/gin-gonic/gin"
	"server/model/appTypes"
	"server/model/response"
	"server/utils"
)

// AdminAuth 管理员认证
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := utils.GetRoleID(c)

		if role != appTypes.Admin {
			response.Forbidden("Access denied. Admin privileges are required", c)
			c.Abort()
			return
		}

		c.Next()
	}
}
