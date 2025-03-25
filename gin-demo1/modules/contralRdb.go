package modules

import (
	"fmt"
	"gin-demo/dao"
	"github.com/go-redis/redis"
)

func RedisExample() {
	err := dao.Rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Println("set score failed , err:", err)
		return
	}

	val, err := dao.Rdb.Get("score").Result()
	if err != nil {
		fmt.Println("get score failed , err:", err)
		return
	}
	fmt.Println("score:", val)

	val2, err := dao.Rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("get name failed , err:", err)
		return
	} else {
		fmt.Println("name:", val2)
	}
}

func HgetDemo() {
	v, err := dao.Rdb.HGetAll("user").Result()
	if err != nil {
		fmt.Println("get user failed , err:", err)
		return
	}
	fmt.Println("user:", v)

	val := dao.Rdb.HMGet("user", "name", "age").Val()
	fmt.Println("user:", val)

	val1 := dao.Rdb.HGet("user", "name").Val()
	fmt.Println("user:", val1)
}

func RedisExample2() {
	zsetKey := "language_rank"
	langulages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	num, err := dao.Rdb.ZAdd(zsetKey, langulages...).Result()
	if err != nil {
		fmt.Println("zadd failed , err:", err)
		return
	}
	fmt.Println("zadd:", num)

	// 把Golang的分数加10
	newScore, err := dao.Rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Println("zincrby failed , err:", err)
		return
	}
	fmt.Println("zincrby:", newScore)

	//// 取分数最高的3个
	ret, err := dao.Rdb.ZRevRangeWithScores(zsetKey, 0, 3).Result()
	if err != nil {
		fmt.Println("zrevrange failed , err:", err)
		return
	}
	for _, v := range ret {
		fmt.Println(v)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = dao.Rdb.ZRevRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Println("zrevrange failed , err:", err)
		return
	}
	for _, v := range ret {
		fmt.Println(v.Member, v.Score)
	}
}
