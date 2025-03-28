package controller

import (
	"bubble/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandler 获取社区简要信息
// @Summary 获取社区的列表
// @Description 获取社区列表，包括社区id和名称
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityLists "社区列表“
// @Router /api/v1/community [get]
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

// CommunityDetailHandler 获取社区详细信息
// @Summary 获取单个社区的详细信息
// @Description 根据社区id获取单个社区的详细信息，包括id,名称,介绍,创建时间和修改时间
// @Tags 社区相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param id path int true "社区ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityDetailLists "社区详细信息“
// @Router /api/v1/community/{id} [get]
func CommunityDetailHandler(c *gin.Context) {
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
