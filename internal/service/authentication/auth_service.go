package service

import (
	"BackendPOS/internal/cache"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthorizationService struct {
	db *gorm.DB
}

func NewAuthorizationService(db *gorm.DB) *AuthorizationService {
	return &AuthorizationService{db}
}

func (s *AuthorizationService) HasPermission(userID uint64, permission string) (bool, error) {
	var count int64

	err := s.db.Table("permissions").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Joins("JOIN user_roles ON user_roles.role_id = role_permissions.role_id").
		Where("user_roles.user_id = ?", userID).
		Where("permissions.code = ?", permission).
		Count(&count).Error

	return count > 0, err
}

func (s *AuthorizationService) HasRole(userID uint64, role string) (bool, error) {
	var count int64

	err := s.db.Table("roles").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Where("roles.name = ?", role).
		Count(&count).Error

	return count > 0, err
}

func (s *AuthorizationService) GetUserPermissions(userId uint64) ([]string, error) {
	//cacheKey := fmt.Sprintf("user:permissions:%d", userId)
	cacheKey := s.GetUserPermissionCacheKey(userId)

	val, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	if err == nil {
		logrus.Info("Permissions loaded from Redis cache|", cacheKey)
		var perms []string
		err := json.Unmarshal([]byte(val), &perms)
		if err != nil {
			return []string{}, err
		}
		return perms, nil
	}

	var permissions []string
	logrus.Info("Permissions loaded from db|", cacheKey)
	err = s.db.
		Table("permissions").
		Select("permissions.code").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Joins("JOIN user_roles ON user_roles.role_id = role_permissions.role_id").
		Where("user_roles.user_id = ?", userId).
		Pluck("permissions.code", &permissions).Error

	if err != nil {
		return nil, err
	}

	bytes, _ := json.Marshal(permissions)
	cache.Rdb.Set(
		cache.Ctx,
		cacheKey,
		bytes,
		24*time.Hour,
	)

	return permissions, nil
}

func (s *AuthorizationService) DeleteUserPermissionCache(userId uint64) {
	cacheKey := s.GetUserPermissionCacheKey(userId)
	cache.Rdb.Del(cache.Ctx, cacheKey)
}

func (s *AuthorizationService) GetUserPermissionCacheKey(userId uint64) string {
	return fmt.Sprintf("user:permissions:%d", userId)
}

func (s *AuthorizationService) GetUserId(ctx *gin.Context) uint64 {
	userId, exists := ctx.Get("userId")
	if !exists {
		return 0
	}
	return userId.(uint64)
}

func (s *AuthorizationService) GetUserName(ctx *gin.Context) string {
	name, exists := ctx.Get("name")
	if !exists {
		return ""
	}
	return name.(string)
}
