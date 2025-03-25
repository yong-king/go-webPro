package middlewares

import (
	"bubble/controller"
	"bubble/dao/redis"
	"bubble/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

func JwtAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization
		// Authorization: Bearer xxx.xxx.xxx
		authHead := c.Request.Header.Get("Authorization")
		if authHead == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			return
		}
		// 按空格划分
		parse := strings.SplitN(authHead, " ", 2)
		if len(parse) != 2 && parse[0] != "Bearer" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// parse[1]为token,解析验证token
		mc, err := jwt.ParseToken(parse[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 根据名称拿到用户id
		
		// 检测是否有用户登录
		token, err := redis.FetchAuth(mc.UserID)
		if err != nil && err != redis.ErrorNoAuth {
			zap.L().Error("redis error", zap.Error(err))
			c.Abort()
			return
		}

		// 有多个用户登录
		if token != "" && parse[1] != token {
			zap.L().Error("multi user login", zap.Int64("mutil login", mc.UserID))
			controller.ResponseError(c, controller.CodeMutilUser)
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
