package dao

import (
	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func InitRedis() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "youngking98",
		DB:       0,  //数据库
		PoolSize: 10, // 连接池大小
	})
	_, err = Rdb.Ping().Result()
	return
}
