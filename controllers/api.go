package controllers

import (
	"github.com/astaxie/beego"
)

type APIController struct {
	beego.Controller
}

func (c *APIController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}
