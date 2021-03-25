package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type DataUser struct{
	Id int
	Name string `orm:"size(100)"`
}

type Post struct {
	Id    int    `orm:"auto"`
	Title string `orm:"size(100)"`
	DataUser  *DataUser  `orm:"rel(fk)"`
}

func itit(){
	orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8&loc=Local", 30)
	orm.RegisterModel(new(DataUser))
	orm.RunSyncdb("default",false,true)
}

func main() {
    o := orm.NewOrm()

	// insert
	user := DataUser{Name:"slenn"}
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := DataUser{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//Associated query
	var posts []*Post
	qs := o.QueryTable("post")
	num1, err1 := qs.Filter("User__Name", "slene").All(&posts)
	fmt.Printf("NUM: %d, ERR: %v\n", num1, err1)

	//Native sql
	//var maps []orm.Params
	//num2, err2 := o.Raw("SELECT * FROM user").Values(&maps)
	//for _,term := range maps{
	//	fmt.Println(term["id"],":",term["name"])
	//}

	//Transaction processing
	to,err := o.Begin()
	user2 := DataUser{Name: "slene"}
	id, err2 := o.Insert(&user2)
	if err2 == nil {
		to.Commit()
	} else {
		to.Rollback()
	}


	//print sql
	orm.Debug = true
}