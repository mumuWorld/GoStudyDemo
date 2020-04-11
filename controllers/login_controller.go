package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"../models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) ShowLogin() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}

func (c *LoginController) HandleLogin()  {
	username := c.GetString("userName")
	password := c.GetString("password")
	logs.Info("user-",username,"pw-",password)
	if username == "" || password == "" {
		logs.Error("数据不对")
		c.Data["errMsg"] = "请输入正确用户名和密码"
		c.TplName = "login.html"
		return
	}
	o := orm.NewOrm()
	//查询数据
	var user models.User
	user.Name = username
	err := o.Read(&user,"Name")
	if err != nil {
		c.Data["errMsg"] = "用户未注册"
		c.TplName = "login.html"
		logs.Error("用户为注册")
		return
	}
	if user.PassWord != password {
		c.Data["errMsg"] = "密码不正确"
		c.TplName = "login.html"
		logs.Error("密码不正确")
		return
	}
	c.Data["errMsg￿"] = "登录成功"
	//c.TplName = "login.html"
	//c.Ctx.WriteString("登录成功")
	c.Redirect("/showArticleList",302)
}
