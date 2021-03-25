package main
import (
   "fmt"
   "net/http"
   "io/ioutil"
   "net/url"
)

func main() {
	testCilentPostForm()
	testHttpPostForm()
}


func testCilentPostForm() {
	//创建客户端
	client := http.Client{}
	//通过client去请求
	response, err := client.PostForm("https://www.toutiao.com/search/suggest/initial_page",url.Values{"key": {"Value"}, "id": {"123"}})
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

func testHttpPostForm() {
	response, err := http.PostForm("https://www.toutiao.com/search/suggest/initial_page",url.Values{"key": {"Value"}, "id": {"123"}})
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