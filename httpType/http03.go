package main
import (
   "fmt"
   "net/http"
   "io/ioutil"
   "strings"
)

func main() {
	testCilentPost()
	testHttpPost()
}

func testCilentPost() {
	//创建客户端
	client := http.Client{}
	//通过client去请求
	response, err := client.Post("https://www.toutiao.com/search/suggest/initial_page","application/x-www-form-urlencoded",strings.NewReader("name=cjb"))
	CheckErr(err)
	fmt.Printf("响应状态码：%v\n", response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	if response.StatusCode == 200 {
		fmt.Println("网络请求成功")
		defer response.Body.Close()
		//处理
	}

	fmt.Println(string(body))
}

func testHttpPost() {
	response, err := http.Post("https://www.toutiao.com/search/suggest/initial_page","application/x-www-form-urlencoded",strings.NewReader("name=cjb"))
	CheckErr(err)
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