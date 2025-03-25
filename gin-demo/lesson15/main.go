package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 单个文件上传
	//r.POST("upload", func(c *gin.Context) {
	//	// 从请求中读取文件
	//	f, err := c.FormFile("f1")
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	} else {
	//		// 将读取到的文件存在本地
	//		//src := fmt.Sprintf("./%s", f.Filename)
	//		src := path.Join("./", f.Filename)
	//		err := c.SaveUploadedFile(f, src)
	//		if err != nil {
	//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		}
	//		c.JSON(http.StatusOK, gin.H{"message": "success"})
	//	}
	//})

	// 多个文件上传
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中获取上传的文件
		form, _ := c.MultipartForm()
		files := form.File["f1"]

		// 将文件保存到服务器
		for _, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf(".upload/%s", file.Filename)
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d files uploaded", len(files))})
	})

	r.Run(":9090")
}
