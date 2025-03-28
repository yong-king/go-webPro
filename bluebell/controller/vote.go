package controller

import (
	"bubble/logic"
	"bubble/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PostVoteHandler 用户给帖子投票
// @Summary 用户给帖子投票
// @Description 根据用户名给指定的帖子投票，包括帖子id和投票类型(-1,0,1)
// @Tags 投票相关接口
// @Accept application/json
// @Produce application/json
// @Param post_id body string true "帖子ID"
// @Param direction body int8 true "投票方向（1：支持，0：取消，-1：反对）"
// @Success 200 {object} _ResponseVote
// @Router /api/v1/vote [post]
func PostVoteHandler(c *gin.Context) {
	// 获取校验参数
	vote := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(vote); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok { // 如果不是validator.ValidationErrors，ok为false
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	userId, err := getCurrentUserID(c) // 获取当前用户id
	if err != nil {
		zap.L().Error("getCurrentUserID(c)", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 处理业务逻辑
	// 更新帖子票数和用户的投票状态
	err = logic.PostVote(userId, vote)
	if err != nil {
		zap.L().Error("redis.PostVote(c) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
