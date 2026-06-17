package util

import (
	"BackendPOS/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var cfg = config.Load()
var jwtSecret = []byte(cfg.JwtSecret)

type UserClaims struct {
	Code       string `json:"code"`
	UserId     uint64 `json:"user_id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	MerchantId uint64 `json:"merchant_id"`
	RoleId     uint64 `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateToken(
	Code string,
	UserId uint64,
	Username string,
	Name string,
	MerchantId uint64,
	RoleId uint64,

) (string, error) {

	claims := UserClaims{
		Code:       Code,
		UserId:     UserId,
		Username:   Username,
		Name:       Name,
		MerchantId: MerchantId,
		RoleId:     RoleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
