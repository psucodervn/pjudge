package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/psufighter/pjudge/controllers"
	"io/ioutil"
)

func init() {
	beego.Include(&controllers.AnnouncementController{})
	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.ProblemController{})
	beego.Include(&controllers.SubmissionController{})
	beego.Include(&controllers.UserController{})
	beego.Any("*", func(ctx *context.Context) {
		if !ctx.Input.IsGet() {
			ctx.Output.Body([]byte(""))
			return
		}
		data, _ := ioutil.ReadFile("views/index.html")
		ctx.Output.Body(data)
	})
}
