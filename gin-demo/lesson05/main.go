package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2. 解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed:", err)
		return
	}
	//3. 渲染模版
	user := User{
		Name:   "dlrb",
		Gender: "female",
		Age:    18,
	}

	m := map[string]interface{}{
		"name":   "zrl",
		"age":    20,
		"gender": "female",
	}

	hobby := []string{
		"唱",
		"条",
		"rap",
	}
	err = t.Execute(w, map[string]interface{}{
		"user":  user,
		"m":     m,
		"hobby": hobby,
	})

	if err != nil {
		fmt.Println("Execute failed:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed:", err)
		return
	}
}

//func main() {
//	r := gin.Default() // 创建一个默认路由请求
//	r.GET("/hello", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "hello world",
//		})
//	})
//	r.POST("/book", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "post",
//		})
//	})
//	r.PUT("/book", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "put",
//		})
//	})
//	r.DELETE("/book", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "delete",
//		})
//	})
//	r.Run(":8080")
//}

//func sayHello(w http.ResponseWriter, r *http.Request) {
//	b, _ := ioutil.ReadFile(". /hello.txt")
//	fmt.Println(string(b))
//	_, _ = fmt.Fprintln(w, string(b))
//}
//
//func main() {
//	http.HandleFunc("/hello", sayHello)
//	err := http.ListenAndServe("127.0.0.1:8080", nil)
//	if err != nil {
//		fmt.Printf("http listen err: %v\n", err)
//		return
//	}
//}
