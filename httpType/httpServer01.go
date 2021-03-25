package main

import "net/http"

func main() {
	testFileServer()
}

func testFileServer()  {
    //如果该路径里面有index.html文件，会优先显示index.html文件，否则会看到文件目录
	http.ListenAndServe(":2003",http.FileServer(http.Dir("./httpServer/")))
}
