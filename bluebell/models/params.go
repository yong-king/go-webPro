package models

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostId    string `json:"post_id"`                          // 帖子id
	Direction int8   `json:"direction" binding:"oneof=0 1 -1"` // 投票的值 -1， 0， 1
}

type ParmPostList struct {
	CommunityId int64  `json:"community_id" form:"community_id"`
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"orderBy" form:"order"`
}

const (
	OrederByTime  = "time"
	OrederByScore = "score"
)

//type ParmCommunityPostList struct {
//	CommunityId int64 `json:"community_id" form:"community_id"`
//	*ParmPostList
//}
