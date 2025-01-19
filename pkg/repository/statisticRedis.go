package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type StatisticRedis struct {
	db *RedisDB
}

func NewStatisticRedis(db *RedisDB) *StatisticRedis {
	return &StatisticRedis{
		db: db,
	}
}

func (r *StatisticRedis) CacheUserStatistic(userID, month, stats string) error {
	cacheKey := fmt.Sprintf("statistic:monthly:%s:%s", userID, month)
	err := r.db.redisClient.Set(context.Background(), cacheKey, stats, time.Second*20).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *StatisticRedis) GetUserCachedStatistic(userID, month string) (string, error) {
	cacheKey := fmt.Sprintf("statistic:monthly:%s:%s", userID, month)

	data, err := r.db.redisClient.Get(context.Background(), cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrCacheNotFound
		}
		return "", err
	}

	return data, err
}

func (r *StatisticRedis) DeleteCachedStatisticByMonth(userID, month string) error {
	cacheKey := fmt.Sprintf("statistic:monthly:%s:%s", userID, month)

	return r.db.redisClient.Del(context.Background(), cacheKey).Err()
}

func (r *StatisticRedis) DeleteAllUserCachedStatistics(userID string) error {
	pattern := fmt.Sprintf("statistic:monthly:%s:*", userID)

	keys, err := r.db.redisClient.Keys(context.Background(), pattern).Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		if err := r.db.redisClient.Del(context.Background(), keys...).Err(); err != nil {
			return err
		}
	}

	return nil
}
