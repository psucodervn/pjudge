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

func makeJSONResponse(c *beego.Controller, status string, data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"status": status,
		"data":   data,
	}
	c.ServeJson()
}

func makeJSONMessage(c *beego.Controller, data interface{}) {
	c.Data["json"] = data
	c.ServeJson()
}
