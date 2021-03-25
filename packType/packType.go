package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main(){

	//goland ide 使用go mod下载第三方包
	//1. file==>settings==>go===>go modules==>enable go modules integration 勾上
	//2. 打开Terminal运行：go mod
	//3. 打开Terminal运行：go mod init 模块名
	//3. 打开Terminal运行：go mod vendor

	var str string = ""
	//string包的字符串处理函数
	//查找字符串
    strings.Contains(str,"查找xx字符串") // bool
    strings.ContainsAny(str, "查找包含xx中的任一字符") //bool
    strings.Count(str,"返回包含xx的格式")  //int
    strings.HasPrefix(str, "判断是否有前缀xxx") //bool
	strings.HasSuffix(str, "判断是否有后缀xxx") //bool
	strings.Index(str, "判断字符串xx出现的位置") //int
	strings.IndexByte(str,"返回字符串xx首次出现的位置") // int
	strings.LastIndex(str,"返回字符串xx最后一次出现的问题")  //int
	strings.LastIndexByte(str, "返回字符串xx中的字符串中的最后一次的位置") //int
    //字符串分割
	strings.Fields(str, "以空白字符分割") //string
	strings.Split(str, "xxx分隔符") //string
	strings.SplitAfter(str, "xxx分隔符,字符串最后附上xxx") //string
	strings.SplitAfterN(str, "xxx分隔符,字符串最后附上xxx","决定返回的切片数") //string
	strings.SplitN(str, "xxx分隔符","决定返回的切片数") //string
	//大小写转换
	strings.Title(str) //string   每个单词字母大写返回
	strings.ToLower(str) //string  转换小写
	strings.ToUpper(str) //string  转换大写
	strings.ToTitle(str) //string  转换大写
	//字符串修饰
	strings.Trim(str, "去掉字符串首尾xxx字符") //string
	strings.TrimSpace(str)  //string 去掉字符串首尾空白字符
	//字符串比较
	strings.Repeat(str,"将字符xxx重复count次返回") //strings
	strings.Replace(str,"替换前","替换后") //string
	strings.Join([]str,"连接字符串") //string

	//strconv包的常用函数
	//parse类函数
	strconv.Atoi(str)  //int 字符串转换整型
	strconv.ParseBool(str) //bool 字符串转换布尔型
	strconv.ParseFloat(str,"64|32") //float 字符串转换浮点型
	strconv.ParseInt(str,"进制","64|32")   //int  字符串转换有符号数字
	strconv.ParseUint(str,"进制","64|32")  //uint 字符串转换无符号数字
	//format类函数
	var int1 = 10
	var float1 = 10.01
	strconv.Itoa(int1)  //string  int转换string
	strconv.FormatBool(true) //string   布尔型转换字符串
	strconv.FormatFloat(float1,"进制","精度","64|32") //string 浮点型转换字符串
	strconv.FormatInt(int1,"64|32")   //string  有符号数字转换字符串
	strconv.FormatUint(int1,"64|32")  //string 无符号数字转换字符串
	//字符串与切片转换
	[]byte(str) //slice  字符串转换切片
	string([]byte{}) //string  切片转换字符串

	//regexp正则表达式包
	regexp.Match("正则",[]byte(str))  //正则表达式是否与字节匹配
	regexp.MatchString("正则",str)  //正则表达式是否与字符串匹配
	regexp.Compile("正则") //正则表达式编译成正则对象
	regexp.MustCompile("正则") //正则表达式编译成正则对象
	正则对象.ReplaceAll([]byte(str),[]byte("替换内容"))  //正则替换
	正则对象.Split(str,"长度")  //正则分割

	//time包
	time.Now() //返回本地时间
	time.Now().Format("格式化")  //返回格式化时间日期
	time.Now().Unix()  //返回时间戳（秒）
	time.Now().Date()  //年月日
	time.Now().Year()  //年
	time.Now().Month()  //月
	time.Now().Day()  //日
	time.Now().Weekday()  //周
	time.Now().Hour()  //时
	time.Now().Minute()  //分
	time.Now().Second()  //秒

	//math包
	math.IsNaN(f)
	math.Ceil(f)
	math.Floor(f)
	math.Trunc(f)
	math.Abs(f)
	math.Max(x,y)
	math.Min(x,y)

	//rand随机数
	rand.Int()
	rand.Intn(x)
	rand.Float32()
	rand.Float32()

	//scanln键盘输入
    fmt.Scanln()



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


	//bee 工具的使用
	//如果是创建一个 Web 项目，在 $GOPATH/src 下执行 bee new <项目名>，可以快速生成一个目录结构
	//如果是创建一个 API 应用，同样在 $GOPATH/src 下执行 bee api <项目名>，可以快速生成一个目录结构。
	//如果是创建一个 RPC 应用，同样在 $GOPATH/src 下执行 bee hprose <项目名>，可以快速生成一个目录结构。
	//同时，该命令还支持一些自定义参数自动连接数据库创建相关 model 和 controller ：
	//bee api [appname] [-tables=""] [-driver=mysql] [-conn="root:<password>@tcp(127.0.0.1:3306)/test"]
	//如果 conn 参数为空则创建一个示例项目，否则将基于链接信息链接数据库创建项目。
	//如果开发过程中需要热编译，在 $GOPATH/src/appname 下执行 bee run，通过 fsnotify 监控文件系统
	//如果需要将项目打包压缩，在 $GOPATH/src/appname 下执行 bee pack，会在 $GOPATH/src/appname 下生成 appname.tar.gz 包
	//自动生成代码：generate
	//这个命令有些复杂，而且短期内应该是用不到
	//数据库迁移：migrate
	//这个……略
	//自动生成 Dockerfile 文件
	//可以通过 bee help dockerize 查看帮助信息



}


