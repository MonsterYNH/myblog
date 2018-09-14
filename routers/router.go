package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/123", &controllers.Controller{})
}
