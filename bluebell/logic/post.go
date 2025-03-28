package logic

import (
	"bubble/dao/mysql"
	"bubble/dao/redis"
	"bubble/models"
	"bubble/pkg/snowflake"
	"go.uber.org/zap"
)

// PostData 上传帖子
func PostData(post *models.Post) (err error) {
	// 生成post_id
	post.ID = snowflake.GenID()
	// 将帖子内容上传到数据库
	err = mysql.PostData(post)
	if err != nil {
		zap.L().Error("mysql.PostData failed", zap.Error(err))
		return
	}
	// 存入到redis
	err = redis.CreatePost(post.ID, post.CommunityId)
	if err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return
	}
	return
}

// GetPostDetailByID 根据post_id获取帖子
func GetPostDetailByID(pid int64) (data *models.PostDetail, err error) {
	// 根据post_id到数据库查询获取帖子信息
	post, err := mysql.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(pid)", zap.Error(err))
		return
	}
	// 获取信息中的用户名称
	user, err := mysql.GetUserNameById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserNameById(data.AuthorID)", zap.Error(err))
		return
	}
	// 获取信息中的社区名称
	community, err := mysql.CommunityDetailDataById(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameById(data.CommunityId)", zap.Error(err))
		return
	}
	data = &models.PostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList(page, size int64) (data []*models.PostDetail, err error) {
	// 到数据库中查找
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList(page, size)", zap.Error(err))
		return
	}
	data = make([]*models.PostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserNameById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetPostDetailByID(post.ID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		community, err := mysql.CommunityDetailDataById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityNameById(post.CommunityId) failed",
				zap.Int64("community_id", post.CommunityId),
				zap.Error(err))
			continue
		}
		postDetail := &models.PostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, nil
}

func GetPostList2(p *models.ParmPostList) (data []*models.PostDetail, err error) {
	// 到redis中获取排序后到帖子的post_id
	postId, err := redis.GetPostOreder(p.Order, p.Page, p.Size)
	if err != nil {
		zap.L().Error("GetPostList2Handler,  redis.GetPostOreder() failed", zap.Error(err))
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("postId", len(postId)))
	// 根据post_id到数据库中拿到详情列表
	return getAuthorCommunityInfo(postId)

}

func GetCommunityPostList(p *models.ParmPostList) (data []*models.PostDetail, err error) {
	//  根据community_id查找对于的帖子id
	postId, err := redis.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("redis.GetCommunityPostList() failed", zap.Error(err))
		return
	}
	if len(postId) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetCommunityPostList", zap.Any("postId", postId))
	// 根据帖子ID到数据库中拿到详细信息
	// 根据拿到的信息获取作者详细
	// 根据community_id获取社区信息
	// 根据post_id到数据库中拿到详情列表
	return getAuthorCommunityInfo(postId)
}

func getAuthorCommunityInfo(ids []string) (data []*models.PostDetail, err error) {
	posts, err := mysql.GetPostList2(ids)
	if err != nil {
		zap.L().Error("GetPostList2Handler, mysql.GetPostList2 failed", zap.Error(err))
		return
	}
	// 帖子的分数信息
	votes, err := redis.GetVoted(ids)
	if err != nil {
		zap.L().Error("GetPostList2Handler, mysql.GetVoted() failed", zap.Error(err))
		return
	}
	// 根据拿到的帖子信息查找作者信息和社区信息
	data = make([]*models.PostDetail, 0, len(posts))
	for idx, post := range posts {
		user, err := mysql.GetUserNameById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserNameById(data.AuthorID)", zap.Error(err))
			continue
		}
		community, err := mysql.CommunityDetailDataById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityNameById(data.CommunityId)", zap.Error(err))
			continue
		}
		postDetail := &models.PostDetail{
			AuthorName:      user.Username,
			VoteNum:         votes[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

// 根据是否有社区id获取排序后到信息
func GetPosListNew(p *models.ParmPostList) (data []*models.PostDetail, err error) {
	if p.CommunityId == 0 {
		data, err = GetPostList2(p)
	} else {
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostList2Handler, mysql.GetPostList2 failed", zap.Error(err))
		return
	}
	return
}
