package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/hltl/GoMall/app/user/biz/dal/redis"
)

const (
	TokenKeyPrefix = "user:token:"
	TokenExpire    = 24 * time.Hour
)

// SetUserToken 保存用户ID和token的映射
func SetUserToken(ctx context.Context, userID uint, token string) error {
	key := fmt.Sprintf("%s%d", TokenKeyPrefix, userID)
	return redis.RedisClient.Set(ctx, key, token, TokenExpire).Err()
}

// ValidateToken 验证用户token是否有效
func ValidateToken(ctx context.Context,  userID uint, token string) (bool, error) {
	key := fmt.Sprintf("%s%d", TokenKeyPrefix, userID)
	storedToken, err := redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return storedToken == token, nil
}

func InvalidateToken(ctx context.Context, userID uint) error {
	key := fmt.Sprintf("%s%d", TokenKeyPrefix, userID)
	return redis.RedisClient.Del(ctx, key).Err()
}

func GetUserToken(ctx context.Context, userID uint) (string, error) {
	key := fmt.Sprintf("%s%d", TokenKeyPrefix, userID)
	return redis.RedisClient.Get(ctx, key).Result()
}
