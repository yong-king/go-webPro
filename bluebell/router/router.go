package router

import (
	"bubble/controller"
	"bubble/logger"
	"bubble/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化
	router := gin.New()
	// 插入中间件
	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	v1 := router.Group("/api/v1")
	// 用户注册
	v1.POST("/signup", controller.SignuUpHandler)
	// 用户登录
	v1.POST("/login", controller.LoginHandler)

	// 插入中间件
	v1.Use(middlewares.JwtAuthMiddleWare())
	{
		// 获取社区信息
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDatilHandler)

		// 帖子
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return router
}
