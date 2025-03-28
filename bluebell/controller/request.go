package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUserID(c *gin.Context) (userId int64, err error) {
	value, ok := c.Get(CtxUserIDKey)
	fmt.Printf("value: %v, ok: %v\n", value, ok)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = value.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageCof(c *gin.Context) (page int64, size int64) {
	// 获取参数信息
	// 偏移量、分页
	var err error
	pageStr := c.Query("page")
	sizestr := c.Query("size")

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizestr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
