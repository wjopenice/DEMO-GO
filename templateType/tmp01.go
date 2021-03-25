package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main(){
	http.HandleFunc("/tmp", testTemplate1)
	err :=http.ListenAndServe(":8080",nil)
	fmt.Println(err)
}
//创建处理器函数
func testTemplate1(w http.ResponseWriter, r *http.Request){
	//申明结构体
	type User struct{
		UserId int
		Username string
		Age uint
		Sex string
	}


	user := User{1,"openice",18,"男"}

	//申明map 方案一
	//m := map[string] string {
	//	"name" : "小明",
	//	"age" : "18",
	//	"color" : "red",
	//}

	//申明map 方案二
	m := make(map[string]interface{})
	m["name"] = "小明"
	m["age"] = 18
	m["color"] = "red"
	m["array"] = []string{"A1","B1","C1","D1"}

	//解析模板
	t, err := template.ParseFiles("./httpServer/index.gohtml")
	if err != nil {
		fmt.Println("Error " +  err.Error())
	}
	//执行
	t.Execute(w, user)
	t.Execute(w, m)
}




