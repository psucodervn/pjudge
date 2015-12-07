package controllers

import (
	"github.com/astaxie/beego"
)

type ProblemController struct {
	beego.Controller
}

func (c *ProblemController) URLMapping() {
	c.Mapping("Count", c.Count)
}

// @router /problem/count [post]
func (c *ProblemController) Count() {
	c.Data["json"] = 2
	c.ServeJson()
}

// @router /problem/view/simple/:page:int [post]
func (c *ProblemController) ViewSimple() {
	beego.Info(c.Ctx.Input.Param(":page"))
	c.ServeJson()
}
