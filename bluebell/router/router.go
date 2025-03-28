package router

import (
	"bubble/controller"
	_ "bubble/docs"
	"bubble/logger"
	"bubble/middlewares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.LoadHTMLFiles("./templates/index.html")
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := router.Group("/api/v1")
	// 用户注册
	v1.POST("/signup", controller.SignuUpHandler)
	// 用户登录
	v1.POST("/login", controller.LoginHandler)
	//v1.GET("/ping", middlewares.Ratelimit2(time.Second*2, 1), func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//})

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.GET("/posts", controller.GetPostListHandler)
	v1.GET("/posts2", controller.GetPostList2Handler)
	// 插入中间件
	v1.Use(middlewares.JwtAuthMiddleWare())
	{
		// 获取社区信息
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		// 帖子
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		// 按照顺序获取帖子详情

		// 投票
		v1.POST("/vote", controller.PostVoteHandler)
	}

	pprof.Register(router)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return router
}
