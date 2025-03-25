package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayIndex(w http.ResponseWriter, r *http.Request) {
	// 定义模版

	// 解析模版
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
	}

	name := "dlrb"
	// 渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("execute template failed, err:", err)
	}
}

func sayHome(w http.ResponseWriter, r *http.Request) {
	// 定义模版

	// 解析模版
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
	}

	name := "zrl"
	// 渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("execute template failed, err:", err)
	}
}

func sayHome1(w http.ResponseWriter, r *http.Request) {
	// 定义模版

	// 解析模版
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home1.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
	}

	name := "章若楠"
	// 渲染模板
	err = t.ExecuteTemplate(w, "home1.tmpl", name)
	if err != nil {
		fmt.Println("execute template failed, err:", err)
	}
}

func sayIndex1(w http.ResponseWriter, r *http.Request) {
	// 定义模版

	// 解析模版
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index1.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
	}

	name := "迪丽热巴"
	// 渲染模板
	err = t.ExecuteTemplate(w, "index1.tmpl", name)
	if err != nil {
		fmt.Println("execute template failed, err:", err)
	}
}
func main() {
	http.HandleFunc("/index", sayIndex)
	http.HandleFunc("/home", sayHome)
	http.HandleFunc("/home1", sayHome1)
	http.HandleFunc("/index1", sayIndex1)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("listen failed, err:", err)
	}
}
