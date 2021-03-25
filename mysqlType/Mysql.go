package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//DSN链接信息 【账号:密码@tcp(IP:端口)/数据库?charset=编码】
const DdConnection string = "mysql" //驱动
const DbHost string = "127.0.0.1"
const DbPort string = "3306"
const DbDatabase string = "127_tms"
const DbUsername string = "root"
const DbPassword string = "123456"
const DbCharset string = "utf8"
const Dsn string = DbUsername+":"+DbPassword+"@tcp("+DbHost+":"+DbPort+")/"+DbDatabase+"?charset="+DbCharset

func main()  {


	//链接数据库
	db, err := sql.Open(DdConnection,Dsn)
    if err != nil{
		fmt.Println("数据库链接失败",err.Error())
	} else {
		fmt.Println("链接成功")

		////增删改方案一
		//db.Exec("预定义sql语句","预定义参数")
		//
        ////方案二
        //stmt, _ := db.Prepare("预定义sql语句")
        //result, _ := stmt.Exec("预定义参数")
        //count, _ := result.RowsAffected()
        //fmt.Println(count)
		//
        ////查询
        //db.Query("预定义sql语句","预定义参数")
        ////查询单条数据
		//db.QueryRow("预定义sql语句","预定义参数")
		//
		//
		//stmt2, _ := db.Prepare("查询预定义sql语句")
		//rows, _ := stmt2.Query("预定义参数")
		////循环数据
        //for rows.Next() {
        //	err = rows.Scan("字段名")
        //	if err != nil{
        //		panic(err)
        //		continue
		//	}
		//}
		////关闭预处理
		//rows.Close()


        //关闭mysql链接
        db.Close()
	}
}
