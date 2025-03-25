package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由和路由组
func main() {
	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "get",
		})
	})

	router.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "post",
		})
	})

	router.PUT("/put", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "put",
		})
	})

	router.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete",
		})
	})

	router.Any("any", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"message": "get"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"message": "post"})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{"message": "put"})
			// ...
		}
	})

	// 路由租
	videoGroup := router.Group("/shop")
	{
		videoGroup.GET("/cloth", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "cloth"})
		})
		videoGroup.GET("/shores", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shores"})
		})
		videoGroup.POST("/t-shirt", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "t-shirt"})
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})
	router.Run(":9090")
}
