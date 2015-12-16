package main

import (
	"github.com/astaxie/beego"
	_ "github.com/psufighter/pjudge/routers"
)

func main() {
	beego.SetStaticPath("/images", "images")
	// beego.EnableGzip = true
	beego.Run()
}
