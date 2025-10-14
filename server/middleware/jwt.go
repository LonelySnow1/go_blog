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

/*
	双 token ：
		AccessToken : 访问令牌 用于直接访问API接口或资源，有效期较短
		RefreshToken : 刷新令牌 用于获取新的访问令牌，不直接参与接口/资源的访问，有效期较长
		服务器请求时，优先验证AccessToken
		-> 若 AccessToken无效或过期，验证 RefreshToken，不在黑名单中就生成新的 AccessToken
		-> 若 RefreshToken 无效，则要求用户重新登录
*/

/*
	优点：
		1. 提升安全性，降低令牌泄露风险： 访问令牌有效时间短，而刷新令牌可添加更细粒度的安全校验
		2. 避免重复登陆，用户体验好 ： 一次登录，长期有效
		3. 令牌吊销平滑 ： 需要令牌主动失效的场景，双token表现更好，如账号登出

	双token常用场景：
		1. 企业级单点登录 SSO
		2. 金融支付等高敏操作
		3. 长期离线场景： 邮箱客户端等

	单token优点：
		1. 开发效率高
		2. 对分布式系统天然友好：水平扩展零压力，无需额外的认证中转
		3. 无感续期，采用与单 token中的accesstoken相似的续期机制，自动续期
*/

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
