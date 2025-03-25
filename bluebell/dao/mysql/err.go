package mysql

import "errors"

var (
	ErrUserExit        = errors.New("用户存在")
	ErrUSerNotExist    = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvalidID     = errors.New("非法ID")
)
