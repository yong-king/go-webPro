package mysql

import (
	"bubble/models"
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
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
				from post
				limit ?, ?`
	data = make([]*models.Post, 0, 2)
	err = db.Select(&data, sqlStr, (page-1)*size, size)
	return
}
