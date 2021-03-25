package main
import (
   "fmt"
   "net/http"
   "io/ioutil"
)

func main() {
	testCilentGet()
	testhttpGet()
}

func testCilentGet() {
	//创建客户端
	client := http.Client{}
	//通过client去请求
	response, err := client.Get("https://www.toutiao.com/search/suggest/initial_page?id=1")
	CheckErr(err)
	fmt.Printf("响应状态码：%v\n", response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 {
		fmt.Println("网络请求成功")
		defer response.Body.Close()
		//处理
	}
	fmt.Println(string(body))
}


func testhttpGet() {
	//获取服务器的数据
	response, err := http.Get("https://www.toutiao.com/search/suggest/initial_page?id=1")
	CheckErr(err)
	fmt.Printf("响应状态码:%v\n",response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 {
		//操作响应数据
		fmt.Println("网络请求成功")
		defer response.Body.Close()
		CheckErr(err)
	}else{
		fmt.Println("网络请求失败",response.Status)
	}
	fmt.Println(string(body))
}

func CheckErr(err error)  {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常：",ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}