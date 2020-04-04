package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) ShowLogin() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}