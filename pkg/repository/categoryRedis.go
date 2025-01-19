package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CategoryRedis struct {
	db *RedisDB
}

func NewCategoryRedis(db *RedisDB) *CategoryRedis {
	return &CategoryRedis{
		db: db,
	}
}

func (r *CategoryRedis) GetUserCachedCategories(userID string) (string, error) {
	cacheKey := fmt.Sprintf("categories:%s", userID)

	data, err := r.db.redisClient.Get(context.Background(), cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrCacheNotFound
		}
		return "", err
	}
	return data, nil
}

func (r *CategoryRedis) CacheUserCategories(userID, data string) error {
	cacheKey := fmt.Sprintf("categories:%s", userID)
	err := r.db.redisClient.Set(context.Background(), cacheKey, data, time.Second*10).Err()
	if err != nil {
		return err
	}

	return nil
}
