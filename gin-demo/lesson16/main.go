package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 跳转到百度
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 跳转到其他路由
	r.GET("/a", func(c *gin.Context) {
		// 修改请求路径，套转到处理b到的处理函数
		c.Request.URL.Path = "/b"
		// 继续后续处理
		r.HandleContext(c)
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":9090")
}
