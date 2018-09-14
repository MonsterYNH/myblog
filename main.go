package main

import (
	"github.com/astaxie/beego"
	_ "myblog/models/mongodb"
	_ "myblog/routers"
)

func main() {
	beego.Run()
}
