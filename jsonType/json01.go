package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	encode_json()
}

func encode_json()  {

	//混合类型请使用 【 interface{} 】

	type User struct{
		UserId int
		Username string
		Age uint
		Sex string
	}

	type UserArr struct{
		Id int
		Name string
	}

	type JsonData struct{
		User []User
		UserArr UserArr
	}
    //二维
	user := []User{
		{1,"openice",18,"男"},
		{2,"asax",19,"男"},
	}
    //一维
	userarr := UserArr{
		Id: 1,
		Name:"小明",
	}

    jsondata := JsonData{
    	User : user,
    	UserArr : userarr,
	}

    fmt.Println("jsondata",jsondata)

    //map数据转换json
	m := make(map[string]interface{})
	m["name"] = "小明"
	m["age"] = 18
	m["color"] = "red"
	m["array"] = []string{"A1","B1","C1","D1"}
	//非格式化输出
	if data, err := json.Marshal(m); err == nil {
		fmt.Printf("%s\n",data)
	}
    //格式化输出
	if data, err := json.MarshalIndent(m," ",""); err == nil {
		fmt.Printf("%s\n",data)
	}


    //struct数据转换json
	//非格式化输出
	if data, err := json.Marshal(jsondata); err == nil {
		fmt.Printf("%s\n",data)
	}
	//格式化输出
	if data, err := json.MarshalIndent(jsondata," ",""); err == nil {
		fmt.Printf("%s\n",data)
	}


}
