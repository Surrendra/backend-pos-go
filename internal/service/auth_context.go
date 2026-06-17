package service

import "BackendPOS/internal/util"

type AuthContext interface {
	GetAuthUser() *util.UserClaims
}
