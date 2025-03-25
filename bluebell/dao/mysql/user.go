package mysql

import (
	"bubble/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "youngking"

// CheckExistByName 按姓名查看用户是否存在
func CheckExistByName(username string) (err error) {
	var count int
	sqlStr := `select count(user_id) from user where username=?`
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return
	}
	if count > 0 {
		return ErrUserExit
	}
	return
}

// InsertUser 用户注册
func InsertUser(user *models.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)
	// 执行语句
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// Login 用户登录
func Login(user *models.User) (err error) {
	// 查询用户的密码
	oPassword := user.Password
	sqlStr := `select username, password, user_id from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrUSerNotExist
	}
	if err != nil {
		return
	}
	// 传入密码是否正确
	oPassword = encryptPassword(oPassword)
	if oPassword != user.Password {
		return ErrInvalidPassword
	}
	return
}

func encryptPassword(str string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(str)))
}

// GetUserNameById 根据用户id获取用户名称
func GetUserNameById(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, id)
	return
}
