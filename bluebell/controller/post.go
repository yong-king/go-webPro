package controller

import (
	"bubble/logic"
	"bubble/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 上传帖子
func CreatePostHandler(c *gin.Context) {
	data := new(models.Post)
	// 1.获取上传的信息
	err := c.ShouldBindJSON(data)
	if err != nil {
		zap.L().Error("CreatePostHandler ShouldBindJSON failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 获取用户id
	aid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("CreatePostHandler getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	data.AuthorID = aid
	// 2.将上传信息添加到数据库
	err = logic.PostData(data)
	if err != nil {
		zap.L().Error("logic.PostData(data) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 根据post_id获取帖子
func GetPostDetailHandler(c *gin.Context) {
	// 1.获取ID
	id := c.Param("id")
	if id == "" {
		ResponseError(c, CodeInvalidParam)
		return
	}
	pid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDetailHandler ParseInt failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.根据id查询帖子
	data, err := logic.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {
	page, size := getPageCof(c)
	// 到数据库中查询
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList(page, size) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}
