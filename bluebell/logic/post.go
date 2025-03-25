package logic

import (
	"bubble/dao/mysql"
	"bubble/models"
	"bubble/pkg/snowflake"
	"go.uber.org/zap"
)

// PostData 上传帖子
func PostData(post *models.Post) (err error) {
	// 生成post_id
	post.ID = snowflake.GenID()
	// 将帖子内容上传到数据库
	return mysql.PostData(post)
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
	community, err := mysql.GetCommunityNameById(post.CommunityId)
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
	data = make([]*models.PostDetail, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserNameById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetPostDetailByID(post.ID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		community, err := mysql.GetCommunityNameById(post.CommunityId)
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
