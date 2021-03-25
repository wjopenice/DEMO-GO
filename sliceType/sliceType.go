package main

func main(){
	//切片
	//var 变量名 = make([]数据类型,数组长度,【容量】)
	var numbers = make([]int, 3, 5)

	//直接初始化切片
	s := []int{1, 2, 3}
    //通过数组截取来初始化切片
	arr := []int{1, 2, 3}
	s := arr[:]
	//切片中包含数组所有元素
	s := arr[startIndex:endIndex]
	//缺损endIndex表示一直待最后一个人元素
	s := arr[startIndex:]
	//缺损startIndex表示从第一个元素开始
	s := arr[:endIndex]

	//求长度 len()

	//求容量 cap()

	//追加新元素 append()

	//复制copy()

	//遍历切片  for循环   for...range循环



}
