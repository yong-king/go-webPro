package modules

import (
	"context"
	"errors"
	"fmt"
	"gin-demo/dao"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

// doCommand go-redis基本使用示例
func DoCommand() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 获取结果
	val, err := dao.Rdb9.Get(ctx, "score").Result()
	fmt.Println(val, err)

	// 先获取到命令对象
	cmder := dao.Rdb9.Get(ctx, "score")
	fmt.Println(cmder.Val())
	fmt.Println(cmder.Err())

	// 直接执行命令获取错误
	err = dao.Rdb9.Set(ctx, "score", 10, 0).Err()
	fmt.Println(err)

	// 直接执行命令获取值
	val1 := dao.Rdb9.Get(ctx, "score").Val()
	fmt.Println(val1)
}

func DoDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := dao.Rdb9.Do(ctx, "set", "score", 10, "EX", 3600).Err()
	fmt.Println(err)

	result, err := dao.Rdb9.Do(ctx, "get", "score").Result()
	fmt.Println(result, err)
}

func ScanKEyDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//var cursor uint64
	//for {
	//	var keys []string
	//	var err error
	//	keys, cursor, err = dao.Rdb9.Scan(ctx, cursor, "my*", 0).Result()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	for _, key := range keys {
	//		fmt.Println("key:", key)
	//	}
	//
	//	if cursor == 0 {
	//		break
	//	}
	//}

	iter := dao.Rdb9.Scan(ctx, 0, "my*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

func PipelineDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	pipe := dao.Rdb9.Pipeline() // 创建 Pipeline

	// 添加命令到 Pipeline
	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Minute)

	// 执行 Pipeline
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
		return
	}

	for _, cmd := range cmds {
		fmt.Println("redis result:", cmd)
	}

	fmt.Println(incr.Val())
}

func PipelineDemo2() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var incr *redis.IntCmd
	cmds, err := dao.Rdb9.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, "pipeline_counter")
		pipe.Expire(ctx, "pipeline_counter", time.Minute)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, cmd := range cmds {
		fmt.Println("redis result:", cmd)
	}
	fmt.Println(incr.Val())
}

// 事务
func TransactionRedisDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	pipe := dao.Rdb9.TxPipeline()
	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Minute)
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println("执行事务失败!", err)
		return
	}
	for cmd := range cmds {
		fmt.Println("redis result:", cmd)
	}
	fmt.Println(incr.Val())

	var incr2 *redis.IntCmd
	cmds2, err := dao.Rdb9.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		incr2 = pipe.Incr(ctx, "pipeline_counter")
		pipe.Expire(ctx, "pipeline_counter", time.Minute)
		return nil
	})
	if err != nil {
		fmt.Println("执行事务失败!", err)
		return
	}
	for _, cmd := range cmds2 {
		fmt.Println("redis result:", cmd)
	}
	fmt.Println(incr2.Val())
}

func WatchDemo(ctx context.Context, key string) error {
	return dao.Rdb9.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		time.Sleep(5 * time.Second)

		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)
}

func Example() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	var maxcount = 10

	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			n++

			_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		// 最多执行100次
		for retries := maxcount; retries > 0; retries-- {
			err := dao.Rdb9.Watch(ctx, txf, key)
			if err != redis.TxFailedErr {
				return err // 成功或者其他错误
			}
		}

		return errors.New("increment reached maximum number of retries")
	}

	var wg sync.WaitGroup
	wg.Add(maxcount)
	for i := 0; i < maxcount; i++ {
		go func() {
			defer wg.Done()
			if err := increment("count1"); err != nil {
				fmt.Println("increment err:", err)
			}
		}()
	}
	wg.Wait()

	n, err := dao.Rdb9.Get(ctx, "count1").Int()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("count1", n)
}
