package controller

import (
	"bubble/dao/mysql"
	"bubble/logic"
	"bubble/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 用户注册
// @Summary 用户注册
// @Description 根据用户传入信息注册用户，包括用户名和密码
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param re_password body string true "重复输入的密码，需与密码一致"
// @Success 200 {object} _ResponseUser
// @Router /api/v1/signup [post]
func SignuUpHandler(c *gin.Context) {
	// 1.获取参数信息
	p := new(models.ParamSignUp)
	err := c.ShouldBindJSON(p)
	if err != nil {
		// 参数不正确，返回响应
		zap.L().Error("Signup with invalid param", zap.Error(err))
		// 判断是不是验证类型错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.业务处理
	err = logic.SignUp(p)
	if errors.Is(err, mysql.ErrUserExit) {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		ResponseError(c, CodeUserExit)
		return
	}
	if err != nil {
		ResponseError(c, CodeServeBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 用户登录
// @Summary 用户登录
// @Description 根据用户名和密码登录，获取token
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} _ResponseAuth "用户信息和 token"
// @Router /api/v1/login [post]
func LoginHandler(c *gin.Context) {
	// 1.获取参数信息及校验
	p := new(models.ParamLogin)
	err := c.ShouldBindJSON(p)
	if err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.处理业务逻辑
	user, err := logic.Login(p)
	if errors.Is(err, mysql.ErrUSerNotExist) {
		zap.L().Error("logic.Login failed", zap.Error(err))
		ResponseError(c, CodeUserNotExit)
		return
	}
	if err != nil {
		zap.L().Error("Login failed", zap.String("username", p.Username), zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":  fmt.Sprintf("%d", user.UserID), // 防止int64中数值过大，到js中失真
		"username": user.Username,
		"token":    user.Token,
	})
}
