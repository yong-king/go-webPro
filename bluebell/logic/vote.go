package logic

import (
	"bubble/dao/redis"
	"bubble/models"
	"strconv"
)

/*
规则：
 1. 按照时间作为帖子的分数，且当帖子发布时间超过一周后，不允许在点赞。
	将到期的帖子及其票数存入到sql中
 2. 为了保持让帖子能多呆一天，则按200票多呆一天计算：86400 / 200 = 432，即投一票 +432分
 3. 用户投票(direction)为：-1， 0， 1
    当direction=1:
    用户之前没投过，现在投了赞成票 	 --> + 1 * 432
    用户之前投反对票，现在投了赞成票  --> + 2 * 432
    当direction=0:
    用户之前投反对票，现在取消投票    --> + 1 * 432
    用户之前投赞成票，现在取消投票 	 --> - 1 * 432
    当direction=1:
    用户之前没投过，现在投了反对票 	 --> - 1 * 432
    用户之前投赞成票，现在投了反对票  --> - 2 * 432
*/

func PostVote(uid int64, vote *models.ParamVoteData) error {
	// 更新帖子票数和用户状态
	uidStr := strconv.FormatInt(uid, 10)
	postId := vote.PostId
	value := float64(vote.Direction)
	return redis.PostVote(uidStr, postId, value)
}
