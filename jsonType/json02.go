package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	decode_json()
}

func decode_json()  {

	//混合类型请使用 【 interface{} 】
	type Struct_data1 struct{
		Name string
		Age int
		Color string
	}

	type Struct_data2 struct{
		Level string
		Msg string
	}

	data1 := `[{
      "Name":"小明",
      "Age":20,
      "Color":"red"
	},{
      "Name":"小王",
      "Age":18,
      "Color":"cyan"
	}]`

	data2 := `[{"Level":"debug","Msg":"File:\"test.txt\" Not Found"},{"Level":"debug","Msg":"Logic error"}]`

	//json转换map
	var map_data1 []map[string]string
	json.Unmarshal([]byte(data2),&map_data1)
	fmt.Println("map_data", map_data1)

	var map_data2 []map[string]interface{}
	json.Unmarshal([]byte(data1),&map_data2)
	fmt.Println("map_data", map_data2)

	//json转struct
	var struct_data1 []Struct_data1
	json.Unmarshal([]byte(data1),&struct_data1)
	fmt.Println("struct_data1", struct_data1)

	var struct_data2 []Struct_data2
	json.Unmarshal([]byte(data2),&struct_data2)
	fmt.Println("struct_data1", struct_data2)
}