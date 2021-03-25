package routers

import (
	"myproject/controllers"
	//"github.com/astaxie/beego"
	"github.com/beego/beego/v2/server/web"
)

func init() {
    web.Router("/", &controllers.MainController{})
	web.Router("/index",&controllers.IndexController{})
	web.Router("/user",&controllers.UserController{})
}
