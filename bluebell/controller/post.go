package controller

import (
	"bubble/logic"
	"bubble/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 上传帖子
// @Summary 上传新帖子的接口
// @Description 根据社区id上传帖子到特定的社区
// @Tags 帖子相关接口
// @Accept application/json
// @produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object body models.Post true "帖子参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /api/v1/post [post]xxz
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
		ResponseError(c, CodeNeedLogin)
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
// @Summary 获取单个帖子详情
// @Description 根据帖子id获取帖子信息，包括帖子作者和社区信息
// @Tags 帖子相关接口
// @Accept application/json
// @produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param id path int true "帖子 ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /api/v1/post/{id} [get]
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

// GetPostListHandler 按分页获取数据库中获取帖子列表
// @Summary 获取帖子列表
// @Description 根据分页和偏移量到数据库中获取帖子列表，包括帖子、作者和社区信息
// @Tags 帖子相关接口
// @Accept application/json
// @produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param page query int false "分页值（默认 1）"
// @Param size query int false "偏移量值 （默认 10）"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostLists "帖子列表"
// @Router /api/v1/posts [get]
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

// GetPostList2Handler 按照指定的顺序返回排序后的帖子
// @Summary 按照指定的顺序返回排序后的帖子
// @Description 根据指定的排序方式("time", "score")以及分页和偏移量返回帖子详情
// @Tags 帖子相关接口
// @Accept application/json
// @produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param page query int false "分页值（默认 1）"
// @Param size query int false "偏移量值 （默认 10）"
// @Param order query string false "排序方式 (默认 time, 可选 time, score)"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostLists "帖子列表"
// @Router /api/v1/post2 [get]
func GetPostList2Handler(c *gin.Context) {
	// 1.获取参数信息
	// 按照给定的顺序去获取order, page, size
	//page, size := getPageCof(c)
	p := &models.ParmPostList{
		Page:  1,
		Size:  10,
		Order: models.OrederByTime,
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		zap.L().Error("GetPostList2Handler,  c.ShouldBindQuery(cof) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 处理业务
	// 到redis中获取排序后到帖子的post_id
	// 根据post_id到数据库中拿到详情列表
	// 根据拿到的帖子信息查找作者信息和社区信息
	data, err := logic.GetPosListNew(p)
	if err != nil {
		zap.L().Error("logic.GetPostList2(page, size) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

//// 在给定的社区中按照排序获取帖子
//func GetCommunityPostListHandler(c *gin.Context) {
//	// 1.获取参数信息
//	// 按照给定的顺序去获取order, page, size, community_id
//	//page, size := getPageCof(c)
//	p := &models.ParmPostList{
//		CommunityId: 1,
//		Page:  1,
//		Size:  10,
//		Order: models.OrederByTime,
//	}
//	err := c.ShouldBindQuery(p)
//	if err != nil {
//		zap.L().Error("GetPostList2Handler,  c.ShouldBindQuery(cof) failed", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	// 根据community_id查找该id中的帖子
//	// 并按给定的顺序返回
//	data, err := logic.GetCommunityPostList(p)
//	if err != nil {
//		zap.L().Error("logic.GetPostList2(page, size) failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//
//	ResponseSuccess(c, data)
//}
