package controllers

import (
	"github.com/astaxie/beego"
	"github.com/psufighter/pjudge/conf"
	"github.com/psufighter/pjudge/models"
	"time"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("Register", c.Register)
}

// @router /user/register [post]
func (c *UserController) Register() {
	name := c.Input().Get("name")
	username := c.Input().Get("username")
	hashedPassword := c.Input().Get("password")
	email := c.Input().Get("email")

	if models.UserExisted(username) {
		makeJSONResponse(&c.Controller, conf.Failed, "Username is existed")
		return
	}

	var userID int
	models.LastID("user", &userID)
	user := models.User{
		Name:           name,
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
		UserID:         userID + 1,
		RegisterTime:   time.Now().UnixNano() / 1000000,
	}

	if err := models.UserInsert(&user); err == nil {
		models.IncreaseID("user")
		makeJSONResponse(&c.Controller, conf.Success, user)
	} else {
		makeJSONResponse(&c.Controller, conf.Failed, err.Error())
	}
}

// @router /user/login [post]
func (c *UserController) Login() {
	username := c.Input().Get("username")
	hashedPassword := c.Input().Get("password")

	var user models.User
	if err := models.UserByUserName(username, &user); err != nil {
		makeJSONResponse(&c.Controller, conf.Failed, err.Error())
		return
	}
	if user.HashedPassword != hashedPassword {
		makeJSONResponse(&c.Controller, conf.Failed, "Wrong password!")
		return
	}
	c.SetSession("user", user)
	makeJSONResponse(&c.Controller, conf.Success, "Login success!")
}

// @router /user/checkLoginSession [post]
func (c *UserController) CheckLoginSession() {
	c.Data["json"] = c.GetSession("user")
	c.ServeJson()
}

// @router /user/logout [post]
func (c *UserController) Logout() {
	c.DelSession("user")
	c.ServeJson()
}

// @router /user/view/:username [post]
func (c *UserController) View() {
	username := c.Ctx.Input.Param(":username")
	beego.Info(username)

	var user models.User
	if err := models.UserByUserName(username, &user); err != nil {
		c.Data["json"] = nil
	} else {
		c.Data["json"] = user
	}
	c.ServeJson()
}
