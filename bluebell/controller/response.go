package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	{
		"code": "xxx",
		"msg":	"xxx",
		"data":	"xxx",
	}
*/

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// ResponseError 错误响应
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.GetMsg(),
		Data: nil,
	})
}

// ResponseSuccess 正确响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.GetMsg(),
		Data: data,
	})
}

// ResponseErrorWithMsg 带有信息的错误响应
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
