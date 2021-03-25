package main

import (
	"fmt"
	"net/http"
)

func main() {
	//绑定路由，去触发方法
	http.HandleFunc("/index",indexHandler)
	http.HandleFunc("/index2",indexHandler2)
	//绑定端口
	err := http.ListenAndServe(":2004",nil)
	fmt.Println(err)
}

func indexHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("/index===")
	w.Write([]byte("这是默认首页"))
}

func indexHandler2(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("/index2===")
	w.Write([]byte("这是默认首页"))
}
