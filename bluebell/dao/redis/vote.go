package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

/*
规则：
 1. 按照时间作为帖子的分数，且当帖子发布时间超过一周后，不允许在点赞。
	将到期的帖子及其票数存入到sql中
 2. 为了保持让帖子能多呆一天，则按200票多呆一天计算：86400 / 200 = 432，即投一票 +432分
 3. 用户投票(direction)为：-1， 0， 1
    当direction=1:
    用户之前没投过，现在投了赞成票 	  0 1--> + 1 * 432
    用户之前投反对票，现在投了赞成票  -1 1--> + 2 * 432
    当direction=0:
    用户之前投反对票，现在取消投票    -1 0--> + 1 * 432
    用户之前投赞成票，现在取消投票 	  1 0--> - 1 * 432
    当direction=1:
    用户之前没投过，现在投了反对票 	 0 -1--> - 1 * 432
    用户之前投赞成票，现在投了反对票  1 -1--> - 2 * 432
*/

const (
	TimeLimit    = 3600 * 24 * 7
	scorePreVote = 432 // 每一票的分数
)

func CreatePost(pid, cid int64) error {
	pipeline := rdb.TxPipeline()
	timestamp := float64(time.Now().Unix())
	// 帖子时间
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostTimeZset), redis.Z{
		Score:  timestamp,
		Member: pid,
	})
	// 帖子分数
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostScoreZset), redis.Z{
		Score:  timestamp,
		Member: pid,
	})
	// 社区set
	pipeline.SAdd(context.Background(), getRedisKey(KeyCommunitySetPF+strconv.FormatInt(cid, 10)), pid)
	_, err := pipeline.Exec(context.Background())
	return err
}

func PostVote(uid string, postId string, value float64) error {
	// 查看是否超过限制
	postTime, err := rdb.ZScore(context.Background(), getRedisKey(KeyPostTimeZset), postId).Result()
	if err != nil {
		return ErrorConnRedis
	}
	if time.Now().Unix()-int64(postTime) > TimeLimit {
		// 存到数据库
		return ErrorTimeOver
	}
	// 判断用户当前投票状态与之前的状态
	ov := rdb.ZScore(context.Background(), getRedisKey(KeyVotedZsetPF+postId), uid).Val()
	if ov == value {
		return ErrorVoteRepeated
	}
	score := (value - ov) * scorePreVote
	// 更新帖子票数
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(context.Background(), getRedisKey(KeyPostScoreZset), score, postId)
	// 更新用户状态
	if value == 0 {
		pipeline.ZRem(context.Background(), getRedisKey(KeyPostTimeZset), uid)
	} else {
		pipeline.ZAdd(context.Background(), getRedisKey(KeyVotedZsetPF+postId), redis.Z{
			Score:  value,
			Member: uid,
		})
	}
	_, err = pipeline.Exec(context.Background())
	return err
}
