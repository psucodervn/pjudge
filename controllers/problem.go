package controllers

import (
	"github.com/astaxie/beego"
	"github.com/psufighter/pjudge/conf"
	"github.com/psufighter/pjudge/models"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type ProblemController struct {
	beego.Controller
}

func (c *ProblemController) URLMapping() {
	c.Mapping("Count", c.Count)
	c.Mapping("ViewSimple", c.ViewSimple)
	c.Mapping("View", c.View)
}

// @router /problem/count [post]
func (c *ProblemController) Count() {
	count, _ := models.Count(conf.CStat, nil)
	c.Data["json"] = count
	c.ServeJson()
}

// @router /problem/view/simple/:page:int [post]
func (c *ProblemController) ViewSimple() {
	var stats []models.PStat
	query := bson.M{}
	models.PStatList(&query, &stats)
	c.Data["json"] = stats
	c.ServeJson()
}

// @router /problem/view/:pid [post]
func (c *ProblemController) View() {
	pid, _ := strconv.Atoi(c.Ctx.Input.Params[":pid"])
	query := bson.M{"problemId": pid}

	var problem models.PDetail
	models.PDetailFind(&query, &problem)
	c.Data["json"] = problem
	c.ServeJson()
}

// @router /problem/check/hidden [post]
func (c *ProblemController) CheckHidden() {
	c.Data["json"] = false
	c.ServeJson()
}
