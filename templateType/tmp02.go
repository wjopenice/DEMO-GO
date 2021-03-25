package main

import (
	"html/template"
	"net/http"
)

//创建处理器函数
func testTemplate2(w http.ResponseWriter, r *http.Request)  {
	//解析模板
	//t, _ := template.ParseFiles("index.html")
	t := template.Must(template.ParseFiles("index.html", "index2.html"))
	//执行
	//t.Execute(w, "Hello Template")
	//将响应数据在index2.html文件中显示
	t.ExecuteTemplate(w, "index2.html", "我要去index2.html中")
}

func main() {
	http.HandleFunc("/tmp", testTemplate2)
	http.ListenAndServe(":8080", nil)
}


