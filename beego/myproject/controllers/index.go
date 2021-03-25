package controllers

import (
	//"github.com/astaxie/beego"
	"github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	web.Controller
}

func (this *IndexController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index2.tpl"
}
