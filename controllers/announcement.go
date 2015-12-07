package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type AnnouncementController struct {
	beego.Controller
}

func (c *AnnouncementController) URLMapping() {
	c.Mapping("GetServerTime", c.GetServerTime)
}

// @router /announcement/get_server_time [get]
func (c *AnnouncementController) GetServerTime() {
	c.Data["json"] = time.Now().UnixNano() / 1000000
	c.ServeJson()
}
