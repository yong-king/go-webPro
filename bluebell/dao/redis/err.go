package redis

import "errors"

var (
	ErrorNoAuth   = errors.New("user not exist")
	ErrorGetRedis = errors.New("get redis error")

	ErrorTimeOver     = errors.New("超过了投票时间")
	ErrorVoteRepeated = errors.New("不允许重复投票")
	ErrorConnRedis    = errors.New("连接redis失败")
)
