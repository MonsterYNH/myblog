package about

import "github.com/astaxie/beego"

type AboutController struct {
	beego.Controller
}

func (self *AboutController) About() {
	self.TplName = "about.html"
}
