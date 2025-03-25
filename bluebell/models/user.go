package models

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserID   int64  `db:"user_id"`
	Token    string
}
