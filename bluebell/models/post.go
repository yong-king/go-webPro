package models

type Post struct {
	ID          int64  `json:"id" db:"post_id,string"`
	AuthorID    int64  `json:"author_id,string" db:"author_id"`
	CommunityId int64  `json:"community_id,string" db:"community_id" binding:"required"`
	Title       string `json:"title" db:"title" binding:"required"`
	Content     string `json:"content" db:"content" binding:"required"`
	Create_time string `json:"create_time" db:"create_time"`
}

type PostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum    int64  `json:"vote_num"`
	*Post
	*CommunityDetail `json:"community"`
}
