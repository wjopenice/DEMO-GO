package main
import (
	"fmt"
	"net/http"
)

func main() {
	testHttpNewRequest()
}

func testHttpNewRequest() {
	//创建客户端
	client := http.Client{}
	//创建一个请求，请求方式既可以是get，可以是是post
	request, err := http.NewRequest("GET","https://www.toutiao.com/search/suggest/initial_page",nil)
	CheckErr(err)
	//客户端发送请求
	cookName := &http.Cookie{Name:"username",Value:"Stevent"}
	//添加Cookie
	request.AddCookie(cookName)
	response, err := client.Do(request)
	CheckErr(err)
	//设置请求头
	request.Header.Set("Accept-Lanauage","zh-cn")
	defer response.Body.Close()
	//查看请求头数据
	fmt.Printf("Header:%+v\n",request.Header)
	fmt.Printf("响应状态码：%v\n",response.StatusCode)
	//操作数据
	if(response.StatusCode == 200){
		//data ,err := ioutil.ReadAll(response.Body)
		fmt.Println("网络请求成功")
		CheckErr(err)
		//fmt.Println(string(data))
	}else{
        fmt.Println("网络请求失败",response.Status)
	}
}


//检查错误
func CheckErr(err error){
	//fmt.Println("--------------")
	defer func() {
		if ins,ok := recover().(error); ok {
			fmt.Println("程序出现异常：",ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}