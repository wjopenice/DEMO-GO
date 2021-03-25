package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	var Website = this.GetString("Website","beego.me")
	var Email = this.GetString("Email","astaxie@gmail.com")
	//this.Ctx.WriteString(Website)
	//this.Ctx.WriteString(Email)
	this.Data["Website"] = Website
	this.Data["Email"] = Email
	this.TplName = "index.tpl"
}


//func (this *MainController) Init() {
//
//}
//
//func (this *MainController) Prepare() {
//
//}
//
func (this *MainController) Post() {
    var name = this.GetString("name")
    var age = this.GetString("age")
	//this.SetSession("username","abc")
	//x := this.GetSession("username")
    this.Ctx.WriteString(name)
    this.Ctx.WriteString(age)
	//fmt.Println(x)
    return
}
//
//func (this *MainController) Delete() {
//
//}
//
//func (this *MainController) Head() {
//
//}