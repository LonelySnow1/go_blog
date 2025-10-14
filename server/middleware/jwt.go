package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"server/service"
	"server/utils"
	"strconv"
)

var jwtService = service.ServiceGroupApp.JwtService

func JETAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := utils.GetAccessToken(c)
		refreshToken := utils.GetRefreshToken(c)

		// 检查refreshToken是否在黑名单中
		if jwtService.IsInBlacklist(refreshToken) {
			utils.ClearRefreshToken(c)
			response.NoAuth("Account logged in from another location or token is invalid", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()

		claims, err := j.ParseAccessToken(accessToken)
		if err != nil {
			if accessToken == "" || errors.Is(err, utils.TokenExpired) {
				refreshClaims, err := j.ParseRefreshToken(accessToken)
				if err != nil {
					utils.ClearRefreshToken(c)
					response.NoAuth("Refresh token expired or invalid", c)
					c.Abort()
					return
				}

				var user database.User
				if err := global.DB.Select("uuid", "role_id").Take(&user, refreshClaims.UserID).Error; err != nil {
					utils.ClearRefreshToken(c)
					response.NoAuth("The user does not exist", c)
					c.Abort()
					return
				}

				newAccessClaims := j.CreateAccessClaims(request.BaseClaims{
					UserID: user.ID,
					RoleID: claims.RoleID,
					UUID:   user.UUID,
				})

				newAccessToken, err := j.CreateAccessToken(newAccessClaims)
				if err != nil {
					utils.ClearRefreshToken(c)
					response.NoAuth("Failed to create new access token", c)
					c.Abort()
					return
				}

				c.Header("new-access-token", newAccessToken)
				c.Header("new-access-expires-at", strconv.FormatInt(newAccessClaims.ExpiresAt.Unix(), 10))

				c.Set("claims", &newAccessClaims)
				c.Next()
				return
			}

			utils.ClearRefreshToken(c)
			response.NoAuth("Invalid access token", c)
			return
		}

		c.Set("claims", claims) // 存在上下文里
		c.Next()
	}
}
