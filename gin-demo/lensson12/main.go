package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// login POST
	r.POST("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Username": c.DefaultPostForm("username", "sb"),
			"Password": c.DefaultPostForm("password", "***"),
		})
	})

	r.Run(":9090")
}
