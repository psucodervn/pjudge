package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @router / [get]
func (c *MainController) Get() {
	c.TplNames = "index.html"
}
