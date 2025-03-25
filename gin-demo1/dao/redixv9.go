package dao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Rdb9 *redis.Client

func InitRedis9() error {
	Rdb9 = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "youngking98",
		DB:       0,
		PoolSize: 20,
	})
	_, err := Rdb9.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v\n", err)
		return err
	}
	return nil
}
