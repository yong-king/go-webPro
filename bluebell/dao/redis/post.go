package redis

import (
	"bubble/models"
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func GetPostOreder(order string, page, size int64) ([]string, error) {
	// 判断要查询到key
	key := getRedisKey(KeyPostTimeZset)
	if order == models.OrederByScore {
		key = getRedisKey(KeyPostScoreZset)
	}
	return getIDsFormKey(key, page, size)
}

func GetVoted(ids []string) ([]int64, error) {
	// 根据ids获取当前post_id的帖子分数,即点赞的次数，
	// 要计算当前id中用户点为正点个数
	//data := make([]int64, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyVotedZsetPF + id)
	//	v, err := rdb.ZCount(context.Background(), key, "1", "1").Result()
	//	if err != nil {
	//		return nil, err
	//	}
	//	data = append(data, v)
	//}
	//return data, nil
	// 使用TxPipeline可以见少rtt
	data := make([]int64, 0, len(ids))
	pipeline := rdb.TxPipeline()
	for _, id := range ids {
		key := getRedisKey(KeyVotedZsetPF + id)
		pipeline.ZCount(context.Background(), key, "1", "1")
	}
	// 返回 cmders，其中包含 多个 Redis 命令的执行结果。
	cmders, err := pipeline.Exec(context.Background())
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmders {
		// cmder 是 redis.Cmder 接口，需要转换为 redis.IntCmd，否则无法访问 .Val() 方法。
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, nil
}

func GetCommunityPostList(p *models.ParmPostList) ([]string, error) {
	// 社区对应的帖子id
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityId)))
	// 按照时间或中分数排序的帖子
	orderKey := getRedisKey(KeyPostScoreZset)
	if p.Order == KeyPostTimeZset {
		orderKey = getRedisKey(KeyPostTimeZset)
	}
	// 新的键值对的key，即社区中帖子的排序
	key := orderKey + ckey
	// 不存在
	if rdb.Exists(context.Background(), key).Val() < 1 {
		pipeline := rdb.TxPipeline()
		// 获取交集，取交集中的最大值
		pipeline.ZInterStore(context.Background(), key, &redis.ZStore{
			Aggregate: "MAX",
			Keys:      []string{orderKey, ckey},
		})
		pipeline.Expire(context.Background(), key, time.Hour) // 设置超时时间
		_, err := pipeline.Exec(context.Background())
		if err != nil {
			zap.L().Error("GetCommunityPostList pipeline.Exec(context.Background()) failed", zap.Error(err))
			return nil, err
		}
	}
	return getIDsFormKey(key, p.Page, p.Size)
}

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	// 存在，直接查找
	start := (page - 1) * size
	end := start + size
	ids, err := rdb.ZRevRange(context.Background(), key, start, end).Result()
	zap.L().Debug("ids", zap.Any("ids", len(ids)))
	if err != nil {
		zap.L().Error("getIDsFormKey, rdb.ZRevRange(context.Background()) failed", zap.Error(err))
		return nil, err
	}
	return ids, nil
}
