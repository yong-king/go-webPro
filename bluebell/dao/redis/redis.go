package redis

import (
	"bubble/setting"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

var rdb *redis.Client
var ErrorNoAuth = errors.New("user not exist")
var ErrorGetRedis = errors.New("get redis error")

func Init(cfg *setting.RedisConfig) (err error) {
	// 连接redis
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MaxIdleConns: cfg.MaxIdle,
	})
	// 配置redis
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("redis connect failed", zap.Error(err))
		return
	}
	return
}

// InsertAuth 存入Auth信息
func InsertAuth(userID int64, token string) (err error) {
	err = rdb.Set(context.Background(), fmt.Sprintf("%d", userID), token, time.Minute*5).Err()
	if err != nil {
		zap.L().Error("redis set failed", zap.Error(err))
		return
	}
	return
}

// FetchAuth 获取Auth信息
func FetchAuth(userID int64) (token string, err error) {
	token, err = rdb.Get(context.Background(), fmt.Sprintf("%d", userID)).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrorNoAuth
		}
		zap.L().Error("redis get failed", zap.Error(err))
		return "", ErrorGetRedis
	}
	return token, nil
}

func Close() {
	_ = rdb.Close()
}
