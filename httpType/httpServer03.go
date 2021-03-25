package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const HTML_PREFIX string = ""
const PUBLIC_DIR string =""

func main() {
	//multiplexer 多路由转换器
	serveMux := http.NewServeMux()
	//静态文件的路由
	serveMux.HandleFunc(HTML_PREFIX,staticHandler)
	err := http.ListenAndServe(":4000",serveMux)
	helper.CheckErr(err)
	
	//动态文件路由
	serveMux.HandleFunc("/login".loginHandler)
	serveMux.HandleFunc("/reg",redHandler)
}

//静态文件路由处理器
func staticHandler(w http.ResponseWriter,r *http.Request) {
	//解析路由路径
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Path[len(HTML_PREFIX)-1:])
	//拼凑成文件的真实物路径
	file := PUBLIC_DIR + r.URL.Path[len(HTML_PREFIX)-1:]
	fmt.Println(file)
	//判断文件是否存在
	if ok := helper.IsFileExist(file); !ok {
		//http.NotFound(w, r)
		w.Write([]byte("异常：你访问的文件不存在！"))
		return
	}
	http.ServeFile(w, r, file)
}

//获取GET方式传递的参数
func loginActionHandlerGet(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	if r.Method == "GET" && r.ParseForm() == nil {
		username := r.FormValue("username")
		pwd := r.FormValue("password")
		if len(username) < 4 || len(username) > 10 {
			w.Write([]byte("用户名不符合规范"))
		}
		if len(pwd) < 6 || len(pwd) > 16 {
			w.Write([]byte("密码不符合规范"))
		}
		//页面跳转
		http.Redirect(w, r, "/list", http.StatusFound)
		return
	} else {
		w.Write([]byte("请求方法不对"))
		return
	}
	w.Write([]byte("登录失败！"))
}

//普通POSt表单请求:Content-type=application/x-www-form-urlencoded   r.PostFormValue(key)
//获取Post方式传递的参数
func loginActionHandlerPost(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	if r.Method == "POST" && r.ParseForm() == nil {
		username := r.PostFormValue("username")
		pwd := r.PostFormValue("password")
		pwd = helper.Hash(pwd, "md5", false)
		logininfo := username + ":" + pwd
		//将信息保存文件
		helper.AppendToRegInfo(USER_FILE, logininfo+"\n")
		//注册信息保存进cookie
		helper.SetCooke(w, r, COOKIE_NAME,logininfo)
		//页面跳转
		http.Redirect(w, r, "/list", 302)
		return
	} else {
		w.Write([]byte("请求方法不对"))
		return
	}
	w.Write([]byte("登录失败！"))
}
//有文件上传的表单：Content-type=multipart/form-data  r.FormFile(key)
//获取Post方式传递的参数
func uploadHander(w http.ResponseWriter, r *http.Request){
	//判断用户是否判断
	validateUser(w, r)
	if r.Method == "GET" {
		w.Write([]byte(VIEW_UPLOAD))
		return
	}
	if r.Method == "POST" {
		srcFile, fileHeader, err := r.FormFile("uploadfile")
		defer srcFile.Close();
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := fileHeader.Filename
		ext := helper.GetFileSuffix(filename)
		if ext != "jpg" && ext != "jpeg" && ext != "png" && ext != "bmp" && ext != "gif" {
			http.Redirect(w, r, "/upload", http.StatusFound)
			return
		}
		destFile, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer destFile.Close()
		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/html/upload/" + filename, http.StatusFound)
	}
}
//设置COOKIE的值
func setCookeHandler(w http.ResponseWriter, r *http.Request)  {
	str := "steven:123456"
	cookie1 := http.Cookie{
		Name: "logininfo",
		Value: str,
		MaxAge: 60,
		Expires: time.Unix(60, 0),
	}
	http.SetCookie(w, &cookie1)
	w.Write([]byte("cookie已经被设置"))
}
//获取COOKIE的值
func getCookieHandler(w http.ResponseWriter, r *http.Request)  {
	//header := r.Hander("Cookie")
	//fmt.Println(w, "获取cookie:", header)
	//var htmlStr = "获取cookie" + header[0]
	//io.WriteString(w, htmlStr)
	cookie1, _ := r.Cookie("logininfo")
	res := cookie1.Value
	w.Write([]byte(res))
	//res, _ := base64.URLEncoding.DecodeString(cookie1.Value)
	//w.Write(res)
}