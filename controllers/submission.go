package controllers

import (
	"github.com/astaxie/beego"
)

type SubmissionController struct {
	beego.Controller
}

func (c *SubmissionController) URLMapping() {
	c.Mapping("ViewLanguage", c.ViewLanguage)
}

// @router /submission/view/language [post]
func (c *SubmissionController) ViewLanguage() {
	c.Data["json"] = []string{"Java 7", "GNU C", "GNU C++11"}
	c.ServeJson()
}
