package mysql

import (
	"bubble/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

// PostData 上传帖子
func PostData(post *models.Post) (err error) {
	sqlStr := `insert into post (
                  post_id, title, content, author_id, community_id) 
                  values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityId)
	return
}

// GetPostDetailByID 根据post_idh获取帖子信息
func GetPostDetailByID(pid int64) (data *models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time  
				from post 
				where post_id = ?`
	data = new(models.Post)
	err = db.Get(data, sqlStr, pid)
	return
}

func GetPostList(page, size int64) (data []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post
				order by create_time
				desc 
				limit ?, ?`
	data = make([]*models.Post, 0, 2)
	err = db.Select(&data, sqlStr, (page-1)*size, size)
	return
}

func GetPostList2(pIds []string) (post []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
				from post
				where post_id in (?)
				order by FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, pIds, strings.Join(pIds, ","))
	if err != nil {
		zap.L().Error("GetPostList2 sqlx.In(sqlStr, strings.Join failed", zap.Error(err))
		return nil, err
	}
	query = db.Rebind(query)
	post = make([]*models.Post, 0, len(pIds))
	err = db.Select(&post, query, args...)
	if err != nil {
		zap.L().Error("GetPostList2 sqlx.Scan failed", zap.Error(err))
		return nil, err
	}
	return
}
