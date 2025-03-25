package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 渲染模版
	str := "迪丽热巴"
	err = t.Execute(w, str)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	str := `<script>alert('嘿嘿嘿')</script>`
	err = t.Execute(w, str)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
