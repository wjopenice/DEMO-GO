package main

import (
	_ "myproject/routers"
	//"github.com/astaxie/beego"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	//web.SetStaticPath("/static","static")
	web.BConfig.WebConfig.Session.SessionOn = true
	web.Run()
}



//beego 安装
//第一步：新建目录  $ mkdir beego
//第二步：生成go.mod文件 $ go mod
//第三步：安装依赖包
//$ go get github.com/astaxie/beego
//$ go get github.com/beego/bee
//第四步：安装后  go get github.com/beego/bee 后，在gopath下面多了一个bin文件夹，下面就有bee.exe，需要将这个命令加入系统环境变量
//第五步：生成项目 $ bee new myproject
//第六步：运行 $ bee run
//第七步：安装远程依赖 $ go mod vendor
//第八步：运行项目 $ go run hello
//第九步：浏览器访问 http://127.0.0.1:8080/