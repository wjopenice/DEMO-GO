package main

import "fmt"

func main(){

	//var 指针变量名 *指针类型

	var a int = 10
	var ip *int
	ip = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(ip)
	fmt.Println(*ip)
	fmt.Println(*&a)


	//空指针
	// nil
    //if(ptr == nil)


    //指针数组
	//var 指针变量名 [长度]*类型
	//var ptr [3]*string


	//指针的指针
	//var 指针变量名 **类型


	//指针函数
	//func 函数名(参数名称 *类型) (返回参数列表){函数体}

}