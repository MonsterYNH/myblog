package main

import (
	"github.com/astaxie/beego"
	_ "myblog/routers"
	_ "myblog/models/subscribe"
)

func main() {
	beego.Run()
}
