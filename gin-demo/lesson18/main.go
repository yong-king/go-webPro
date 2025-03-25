package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AdminUser(c *gin.Context) {
	start := time.Now()
	c.Next() // 处理后续逻辑
	//c.Abort()	// 终止处理后续逻辑
	lost := time.Since(start)
	c.JSON(http.StatusOK, gin.H{"message": "AdminUser", "lost": lost})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login"})
	// c.Set 传递值
	c.Set("name", "yykk")
	// go func(c.Copy)  //开启go携程并行时只能用copy，不能直接使用，不然造成并行错误
	c.Next()
}

func TimeSpend(c *gin.Context) {
	start := time.Now()
	c.Next()
	lost := time.Since(start)
	c.JSON(http.StatusOK, gin.H{"message": "TimeSpend", "lost": lost})
}

func Logout(c *gin.Context) {
	// c.Get获取传递值
	name, _ := c.Get("name")
	c.JSON(http.StatusOK, gin.H{"message": "Logout", "name": name})
}

func Subbmit(start bool) gin.HandlerFunc {
	// 一些逻辑
	return func(c *gin.Context) {
		if start {
			c.JSON(http.StatusOK, gin.H{"message": "Subbmit"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "No_Subbmit"})
		}
	}
}
func main() {

	router := gin.Default() // 莫仍使用了Logger和Recover中间件
	// gin.New()  // 初始化一个为空的

	// 全局中间件
	router.Use(TimeSpend)

	// 一般情况
	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "no middle_ware",
		})
	})

	// 插入中间件  type HandlerFunc func(*Context)
	router.GET("/user", AdminUser)

	// 多个中间件
	router.GET("/user/login", AdminUser, Login)

	// 组中使用
	PersonGroup := router.Group("/person")
	PersonGroup.Use(Login) // 中间件全局使用
	{
		PersonGroup.GET("/xxx", Logout)
	}

	// 闭包的形式
	router.GET("/bbb", Subbmit(true))

	router.Run(":9988")
}
