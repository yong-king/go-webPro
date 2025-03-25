package main

import "github.com/gin-gonic/gin"

// 获取请求URL参数
func main() {
	r := gin.Default()

	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 这里前面加上一个前缀，防止会起冲突导致出错
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(200, gin.H{
			"year":  year,
			"month": month,
		})
	})

	r.Run(":9090")
}
