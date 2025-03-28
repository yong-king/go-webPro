package redis

const (
	Prefix            = "bluebell:"   // 项目key的前缀
	KeyPostTimeZset   = "post:time"   // zset; 帖子及发帖时间
	KeyPostScoreZset  = "post:score"  // zset; 帖子及帖子投票分数
	KeyVotedZsetPF    = "post:voted:" // zset;记录用户及投票类型;参数是post id
	KeyCommunitySetPF = "comunity:"
)

func getRedisKey(key string) string {
	return Prefix + key
}
