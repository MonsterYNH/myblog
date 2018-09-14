package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "ynhmonster@163.com"
	c.TplName = "index.tpl"
}

type Controller struct {
	beego.Controller
}

func (c *Controller) Get() {
	c.Data["123"] = 123
	c.Data["456"] = 456
	c.TplName = "index.tpl"
}
