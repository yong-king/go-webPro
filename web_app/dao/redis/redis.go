package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"web_app/settings"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	// 连接和配置redis
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			//viper.GetString("redis.host"),
			//viper.GetInt("redis.port"),
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // viper.GetString("redis.password"),
		DB:       cfg.DB,       // viper.GetInt("redis.db"),
		PoolSize: cfg.PoolSize, // viper.GetInt("redis.pool_size"),
	})

	_, err = rdb.Ping(context.Background()).Result()
	return
}

// 关闭redis
func Close() {
	rdb.Close()
}
