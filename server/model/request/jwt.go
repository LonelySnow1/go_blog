package request

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"server/model/appTypes"
)

// JwtCustomClaims : access token
type JwtCustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

// JwtCustomRefreshClaims : refresh token
type JwtCustomRefreshClaims struct {
	UserID uint
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UserID uint            `json:"user_id"`
	UUID   uuid.UUID       `json:"uuid"`
	RoleID appTypes.RoleID `json:"role_id"`
}
