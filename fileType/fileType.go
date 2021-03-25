package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

//FileInfo接口信息
type FileInfo interface {
	Name() string
	Size() int64
	Mode() FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() interface{}
}

//fileStat结构体
type fileStat struct {
	name string
	size int64
	mode FileMode
	modTime time.Time
	sys syscall.Stat_t
}

//fileStat结构体的常用方法
func (fs *fileStat) Name() string { return fs.name }
func (fs *fileStat) IsDir() bool { return fs.Mode().IsDir() }
func (fs *fileStat) Size() int64 { return fs.size }
func (fs *fileStat) Mode() FileMode { return fs.modTime }
func (fs *fileStat) ModTime() time.Time { return fs.modTime }
func (fs *fileStat) Sys() interface{} { return &fs.sys}


func main() {

	//绝对路径
	path1 := "D:/GoWorks/httpServer/u1561.png"
	printMessage(path1)
	pathMessage(path1)

	//相对路径
	path2 := "./httpServer/u1561.png"
	printMessage(path2)
	pathMessage(path2)

    //ioutil包
    ioutil.ReadFile(path1) //读取文件中的所有数据，返回读者的字节数组
    ioutil.WriteFile(path1,[]byte("写入数据"),os.ModePerm)  //向指定文件写入数据，如果文件不存在，则创建文件，写入数据之前清空文件
    ioutil.ReadDir("./httpServer")  //遍历目录
    ioutil.TempDir("./","目录前缀")  //创建目录
    ioutil.TempFile("./","文件前缀") //创建文件

    //bufio包
	Reader类
	func NewReader(rd io.Reader) *Reader
	func NewReaderSize(rd io.Reader, size int) *Reader // 可以配置缓冲区的大小
	func (b *Reader) Read(p []byte) (n int, err error)
	func (b *Reader) Discard(n int) (discarded int, err error)
	func (b *Reader) Peek(n int) ([]byte, error)
	func (b *Reader) Reset(r io.Reader)
	func (b *Reader) ReadByte() (byte, error)
	func (b *Reader) ReadRune() (r rune, size int, err error)
	func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
	func (b *Reader) ReadBytes(delim byte) ([]byte, error)
	func (b *Reader) ReadString(delim byte) (string, error)
	//Writer类demo1:
	filename1 := "./a.png"
	f1, _ := os.Open(filename1)
	red1 := bufio.NewReader(f1)
	fmt.Println("%t\n",red1)
	for {
		s1, err := red1.ReadString('\n')
		fmt.Println(s1)
		if err == io.EOF {
			fmt.Println("ok")
			break
		}
	}
	f1.Close()

	Writer类
	func NewWriter(rd io.Writer) *Writer
	func NewWriterSize(rd io.Writer, size int) *Writer // 可以配置缓冲区的大小
	func (b *Writer) Write(p []byte) (nn int, err error) // 写入n byte数据
	func (b *Writer) Reset(w io.Writer) // 重置当前缓冲区
	func (b *Writer) Flush() error // 清空当前缓冲区，将数据写入输出
	func (b *Writer) WriteByte(c byte) error // 写入一个字节
	func (b *Writer) WriteRune(r rune) (size int, err error） // 写入一个字符
	func (b *Writer) WriteString(s string) (int, error) // 写入一个字符串
    //Writer类demo2:
	file1 := "./a.png"
	file2, _ := os.Open(file1)
	reader1 := bufio.NewReader(file2)
	file3 := "./b.png"
	file4, _ := os.OpenFile(file3, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	writer1 := bufio.NewWriter(file4)
	for {
		bs, err := reader1.ReadBytes(' ')
		writer1.Write(bs)
		writer1.Flush()
		if err == io.EOF {
			fmt.Println("ok")
			break
		}
	}
	file2.Close()
	file4.Close()

	Scanner类
	func (s *Scanner) Scan() bool
	func (s *Scanner) Text() string
	func (s *Scanner) Bytes() []byte
	//Scanner类 demo3:
	scanner := bufio.NewScanner(strings.NewReader("hello world !"))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}




}
//文件信息案例：
func printMessage(filePathx string){
    fileInfo, err := os.Stat(filePathx)
	if err != nil {
		fmt.Println("err:",err.Error())
	} else {
		fmt.Printf("数据类型是：%T \n",fileInfo)
		fmt.Println("文件名：",fileInfo.Name())
		fmt.Println("是否为目录：",fileInfo.IsDir())
		fmt.Println("文件大小：",fileInfo.Size())
		fmt.Println("文件权限：",fileInfo.Mode())
		fmt.Println("文件最后修改时间：",fileInfo.ModTime())
	}
}
//文件路径案例
func pathMessage(filePathx string)  {
	fmt.Println(filepath.IsAbs(filePathx))  //判断是否绝对路径
	fmt.Println(filepath.Rel("D:/GoWorks",filePathx))  //获取相对路径
	fmt.Println(filepath.Abs(filePathx))   //获取绝对路径
	fmt.Println(path.Join(filePathx,".."))  //拼接路径
	fmt.Println(path.Join(filePathx,"."))  //拼接路径
}
//创建目录案例
func pathMkdir()  {
	err := os.Mkdir("./test1",os.ModePerm) //创建单层目录
	if err != nil {
		fmt.Println("err:",err.Error())
	} else {
		fmt.Println("%s 目录创建成功! \n","./test1")
	}
	err = os.MkdirAll("./test1/abc/xyz",os.ModePerm)  //创建多层目录
	if err != nil {
		fmt.Println("err:",err.Error())
	} else {
		fmt.Println("%s 目录创建成功! \n","./test1/abc/xyz")
	}
}
//创建文件案例
func fileCreate()  {
	file := "./httpServer/u1561.png"
	file1, err := os.Create(file)  //创建文件
	if err != nil {
		fmt.Println("err:",err.Error())
	} else {
		fmt.Println("%s 创建成功! %v \n",file,file1)
	}
}
//打开文件与关闭
func fileOpenClose()  {
	file := "./httpServer/u1561.png"
	file2, err := os.Open(file)  //打开文件
	if err != nil {
		fmt.Println("err:",err.Error())
	} else {
		fmt.Println("%s 打开成功! %v \n",file,file2)
	}

	file3, err := os.OpenFile(file,os.O_RDONLY|os.O_WRONLY|os.O_APPEND|os.O_RDWR|os.O_CREATE,os.ModePerm)  //打开文件
	if err != nil {
		fmt.Println("err:",err.Error())
	} else {
		fmt.Println("%s 打开成功! %v \n",file,file3)
	}

	file2.Close()  //关闭文件
	file3.Close()
}
//删除文件和空目录
func fileRemove(){
	file := "./httpServer"
	err := os.Remove(file)  //删除文件和空目录
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%s 删除成功! %v \n",file)
	}

	err = os.RemoveAll(file)  //移除所有的路径和它包含的任何字节点
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%s 删除成功! %v \n",file)
	}
}
//读取文件
func fileRead()  {
	file := "./httpServer/u1561.png"
	file2, err := os.Open(file)  //打开文件
	if err != nil {
		fmt.Println("err:",err.Error())
	} else {
		fmt.Println("%s 打开成功! %v \n",file,file2)
        bs := make([]byte,1024*8, 1024*8)
        n := -1
        for {
        	n, err = file2.Read(bs)  //读取文件
			if n == 0 || err == io.EOF {
				fmt.Println("读取文件结束")
			}
			fmt.Println(string(bs[:n]))
		}
	}
	file2.Close()
}
//写入文件
func fileWrite()  {
	file := "./httpServer/u1561.png"
	file3, err := os.OpenFile(file,os.O_RDWR|os.O_CREATE,os.ModePerm)  //打开文件
	defer file3.Close()
	if err != nil {
		fmt.Println("err:",err.Error())
	} else {
		fmt.Println("%s 打开成功! %v \n",file,file3)
		n, err := file3.Write([]byte("写入内容"))  //写入数据
		if err != nil {
			fmt.Println("写入文件异常：",err.Error())
		} else {
			fmt.Println("写入文件ok", n)
		}

		n, err = file3.WriteString("写入内容") //写入数据
		if err != nil {
			fmt.Println("写入文件异常：",err.Error())
		} else {
			fmt.Println("写入文件ok", n)
		}
	}

	file3.Close()
}
//复制文件
func fileCopy()  {
	//源文件路径
	srcFile := "./httpServer/u1561.png"
	//新文件路径
	destFile := "./httpServer2/u1561.png"
	total, err := copyFile(srcFile,destFile)
	if err != nil {
		fmt.Println("复制异常：",err.Error())
	} else {
		fmt.Println("复制ok", total)
	}
}

func copyFile(srcFile string, destFile string) (int64, error) {
    file1, err := os.Open(srcFile)
    if err != nil {
    	return 0, err
    }
    file2, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
    return io.Copy(file2,file1)
}

