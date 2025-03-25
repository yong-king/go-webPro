package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/logger"
)

func Setup() *gin.Engine {
	router := gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))
	router.GET("/", func(c *gin.Context) {
		//time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Hello World")
	})
	return router
}
