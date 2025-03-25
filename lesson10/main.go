package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"name": "dlrb",
		//	"age":  20,
		//	"love": "yes",
		//}
		data := gin.H{
			"name": "dlrb",
			"age":  20,
			"love": "yes",
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/json1", func(c *gin.Context) {
		type Data struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
			Love string `json:"love"`
		}
		data := Data{
			"zrn",
			20,
			"yes",
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9090")
}
