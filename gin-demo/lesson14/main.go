package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func main() {
	r := gin.Default()

	// query string : http://localhost:9090/query?name=123&password=123
	r.GET("/query", func(c *gin.Context) {
		name := c.Query("name")
		password := c.Query("password")
		c.JSON(http.StatusOK, gin.H{
			"name":     name,
			"password": password,
		})
	})

	// bondform query string: http://localhost:9090/BQ?name=123&password=123
	r.GET("/BQ", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"password": user.Password,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
	})

	// 拿到index.html页面，提交后的动作是post到form，这里可以直接用postman来做测试
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// postman body from-data user=q1mi password=123456
	r.POST("/form", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"password": user.Password,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
	})

	//绑定JSON的示例 GET - body -raw -json {
	//    "user": "ykk",
	//    "pwd": "123456"
	//}
	r.GET("/json", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBindJSON(&user)
		if err == nil {
			c.JSON(200, gin.H{
				"user": user.Username,
				"pwd":  user.Password,
			})
		}
	})

	r.Run(":9090")
}
