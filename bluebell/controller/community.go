package controller

import (
	"bubble/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandler 获取社区简要信息
func CommunityHandler(c *gin.Context) {
	// 获取社区id和名称
	data, err := logic.CommunityList()
	if err != nil {
		zap.L().Error("logic.CommunityIdList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDatilHandler 获取社区详细信息
func CommunityDatilHandler(c *gin.Context) {
	// 获取参数信息
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 处理业务逻辑
	data, err := logic.CommunityDetailData(id)
	if err != nil {
		zap.L().Error("logic.CommunityDetailData() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}
