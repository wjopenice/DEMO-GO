package main

func main() {

	//结构体
	type 类型名 interface {
		方法1(【参数列表】)【返回值】
		方法2(【参数列表】)【返回值】
		方法3(【参数列表】)【返回值】
		....
	}

	//使用接口方法
	func (变量名 结构体类型名) 方法1(【参数列表】) 【返回值】{
	   //方法体
	}
	func (变量名 结构体类型名) 方法2(【参数列表】) 【返回值】{
		//方法体
	}
	func (变量名 结构体类型名) 方法3(【参数列表】) 【返回值】{
		//方法体
	}


	//空接口==任意数据类型
	interface{}


	//接口对象转型
	interface, ok := 接口对象.(实际类型)  配合if...else使用

	interface := 接口对象.(type) 配合switch...case使用
}
