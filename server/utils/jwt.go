package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"server/global"
	"server/model/request"
	"time"
)

type JWT struct {
	AccessTokenSecret  []byte // 访问令牌（Access Token）的签名密钥
	RefreshTokenSecret []byte // 刷新令牌（Refresh Token）的签名密钥
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("token malformed")
	TokenInvalid     = errors.New("token invalid")
)

func NewJWT() *JWT {
	return &JWT{
		AccessTokenSecret:  []byte(global.Config.Jwt.AccessTokenSecret),
		RefreshTokenSecret: []byte(global.Config.Jwt.RefreshTokenSecret),
	}
}

/*
	Claims 是 JWT 的「数据载体」和「规则契约」，本质上是存储在JWT[payload]中的部分键值对：
	对开发者：它是 “装业务数据的容器”，能减少接口的数据库查询；		—— 可直接将业务数据存储
	对系统：它是 “安全规则的说明书” 						 		—— 定义了令牌的有效期、使用范围等；
	核心原则：不存敏感数据、依赖签名保证完整性、按需设计自定义字段。	—— 任何人都可解密，但不允许被修改
*/

// CreateAccessClaims 构建访问令牌（Access Token）的 Claims 对象
func (j *JWT) CreateAccessClaims(baseClaims request.BaseClaims) request.JwtCustomClaims {
	// 1. 解析配置中的「访问令牌过期时间」（字符串转成 time.Duration）
	ep, _ := ParseDuration(global.Config.Jwt.AccessTokenExpiryTime)

	// 2. 构建自定义 Claims（结合标准声明 + 业务声明）
	claims := request.JwtCustomClaims{
		// 嵌入业务自定义的基础声明（如用户 ID、用户名、角色等）
		BaseClaims: baseClaims,
		// 嵌入 JWT 标准注册声明（规范要求的通用字段）
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"go blog"},            // 令牌受众（谁能使用这个令牌，这里是 "go blog" 项目）
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间（当前时间 + 过期时长）
			Issuer:    global.Config.Jwt.Issuer,               // 令牌签发者（如 "go-blog-api"）
		},
	}
	return claims
}

func (j *JWT) CreateAccessToken(claims request.JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.AccessTokenSecret)
}

func (j *JWT) CreateRefreshClaims(baseClaims request.BaseClaims) request.JwtCustomRefreshClaims {
	ep, _ := ParseDuration(global.Config.Jwt.AccessTokenExpiryTime)
	claims := request.JwtCustomRefreshClaims{
		UserID: baseClaims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"go blog"},            // 令牌受众（谁能使用这个令牌，这里是 "go blog" 项目）
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间（当前时间 + 过期时长）
			Issuer:    global.Config.Jwt.Issuer,               // 令牌签发者（如 "go-blog-api"）
		},
	}
	return claims
}

func (j *JWT) CreateRefreshToken(claims request.JwtCustomRefreshClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.AccessTokenSecret)
}

// 解析函数

func (j *JWT) ParseToken(TokenString string, claims jwt.Claims, secretKey interface{}) (interface{}, error) {
	token, err := jwt.ParseWithClaims(TokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0: // 捕获的错误与格式错误相同
				return nil, TokenMalformed
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return nil, TokenExpired
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, TokenNotValidYet
			case ve.Errors&jwt.ValidationErrorSignatureInvalid != 0:
				return nil, TokenInvalid
			}
		}
		return nil, TokenInvalid
	}
	if token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
