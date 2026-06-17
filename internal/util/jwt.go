package util

import (
	"BackendPOS/internal/config"
	"errors"
	"strings"
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

func ExtractBearerToken(authHeader string) (string, error) {
	if strings.TrimSpace(authHeader) == "" {
		return "", errors.New("authorization header is required")
	}

	parts := strings.SplitN(strings.TrimSpace(authHeader), " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.New("invalid authorization header format")
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", errors.New("token is empty")
	}

	return token, nil
}

func ParseUserClaims(tokenString string) (*UserClaims, error) {
	if strings.TrimSpace(tokenString) == "" {
		return nil, errors.New("token is required")
	}

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func GetUserClaimsFromAuthHeader(authHeader string) (*UserClaims, error) {
	token, err := ExtractBearerToken(authHeader)
	if err != nil {
		return nil, err
	}

	return ParseUserClaims(token)
}
