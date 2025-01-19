package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type RedisDB struct {
	redisClient *redis.Client
}

func NewRedisClient(cfg RedisConfig) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, errors.New("redis connection failed")
	}

	return &RedisDB{
		redisClient: client,
	}, nil
}
