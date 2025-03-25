package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

// TODO Module
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {

	// 创建数据库 CREATE DATABAS bubble
	// 连接数据库
	db, err := gorm.Open("mysql", "root:youngking98@tcp(127.0.0.1:3307)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 连接表单
	db.AutoMigrate(&Todo{})

	route := gin.Default()

	// 加载html路径
	route.LoadHTMLFiles("./templates/index.html")
	// 加载静态文件
	route.Static("/css", "./static/css")
	route.Static("/js", "./static/js")
	route.Static("/static/fonts", "./static/fonts")
	// 获取html页面
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 代办事务
	v1Serve := route.Group("/v1")
	{
		// 添加事务
		v1Serve.POST("/todo", func(c *gin.Context) {
			var to Todo
			err := c.ShouldBind(&to)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}
			if err := db.Create(&to).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "fail to create"})
				return
			}
			c.JSON(http.StatusOK, to)
		})

		// 查询所有事务
		v1Serve.GET("/todo", func(c *gin.Context) {
			// 从数据库中拿到数据
			var to []Todo
			if err := db.Find(&to).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "fail to find"})
				return
			}
			// 将数据传到回页面
			c.JSON(http.StatusOK, to)
		})

		// 查询单个事务
		v1Serve.GET("/todo/:id", func(c *gin.Context) {
			// 获取url参数
			id := c.Param("id")

			// 查询记录
			var todo Todo
			err := db.First(&todo, id).Error
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err})
				return
			}
			c.JSON(http.StatusOK, todo)

		})

		// 修改事务
		v1Serve.PUT("/todo/:id", func(c *gin.Context) {
			// 获取url中id
			id := c.Param("id")

			// 从数据库中查询
			var todo Todo
			err := db.First(&todo, id).Error
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err})
				return
			}

			// 取反status
			todo.Status = !todo.Status

			// 更新数据库
			db.Save(&todo)

			// 返回更新后到数据
			c.JSON(http.StatusOK, todo)

		})

		// 删除事务
		v1Serve.DELETE("/todo/:id", func(c *gin.Context) {
			// 获取id
			id := c.Param("id")

			// 查询id的记录
			var todo Todo
			err := db.Find(&todo, id).Error
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err})
				return
			}

			// 删除记录
			db.Delete(&todo)

			// 返回响应
			c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
		})
	}

	// 运行服务
	route.Run(":9090")
}
