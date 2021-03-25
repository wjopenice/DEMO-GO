package main

func main(){
	//映射
	//var 变量名 = map[key类型]value类型
	//var numbers = map[string]string{}

	m := map[string] string {
		"name" : "小明",
		"age" : "18",
		"color" : "red",
	}


	//var 变量名 = make(map[key类型]value类型)
	//var numbers = make(map[string]string)

	m := make(map[string]interface{})
	m["name"] = "小明"
	m["age"] = 18
	m["color"] = "red"


    //删除 delete(变量名, key)

}
