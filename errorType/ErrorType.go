package main

import (
	"fmt"
	"math"
	"os"
)

func main(){

	//错误类
	type error interface{
		Error() string
	}

	//案例一
    res := math.Sqrt(-100)
    fmt.Println(res)
    res, err := Sqrt(-100)
    if err != nil {
    	fmt.Println(err)
	}else{
		fmt.Println(res)
	}

	//案例二
	res1, err := Divde(100, 0)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(res1)
	}

	//案例三
	f, err := os.Open("/1.txt")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(f.Name())
	}


    //延迟 defer
    //中断 panic
    //恢复 recover
}