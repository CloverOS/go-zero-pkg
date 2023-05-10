package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func NewRedisService(config Config) *redis.Client {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := redisDb.Ping(ctx).Result()
	if err != nil {
		panic("redis connect ping failed, err:" + err.Error())
	}
	return redisDb
}
