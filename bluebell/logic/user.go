package logic

import (
	"bubble/dao/mysql"
	"bubble/models"
	"bubble/pkg/jwt"
	"bubble/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 查询是否存在
	err = mysql.CheckExistByName(p.Username)
	if err != nil {
		return
	}
	// 生成id
	userID := snowflake.GenID()
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
		UserID:   userID,
	}
	// 不存在，添加到数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	// 查询是否存在
	// 查询密码是否正确
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return user, nil
	//token, err = jwt.GenToken(user.UserID, user.Username)
	//if err != nil {
	//	return "", err
	//}
	//// 将生成的token和userid存入redis中
	//err = redis.InsertAuth(user.UserID, token)
	//if err != nil {
	//	return "", err
	//}
	//return token, mysql.Login(user)
}
