package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "18")
		gender, ok := c.GetQuery("gender")
		if !ok {
			gender = "w"
		}
		c.JSON(200, gin.H{
			"name":   name,
			"age":    age,
			"gender": gender,
		})
	})

	r.Run(":9090")
}
