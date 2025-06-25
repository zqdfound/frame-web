package utils

import (
	"context"
	"frame-web/config"
	"frame-web/global"
	"go.uber.org/zap"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitRedisClient(redisConfig config.Redis) (redis.UniversalClient, error) {
	var rdb redis.UniversalClient
	rdb = redis.NewClient(&redis.Options{
		Addr:         redisConfig.Addr,
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.String("addr", redisConfig.Addr), zap.Error(err))
		return nil, err
	}
	global.LOG.Info("redis connect ping response:", zap.String("addr", redisConfig.Addr), zap.String("pong", pong))
	return rdb, nil
}

func Redis() {
	redisClient, err := InitRedisClient(global.CONFIG.Redis)
	if err != nil {
		panic(err)
	}
	global.REDIS = redisClient
}
