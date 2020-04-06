package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"../models"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) ShowRegister() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "register.html"
}

func (c *RegisterController) HandlePost() {
	userName := c.GetString("userName")
	password := c.GetString("password")

	logs.Info("user-",userName,"pw-",password)
	if userName == "" || password == "" {
		logs.Error("数据不对")
		c.Data["errMsg"] = "请输入正确用户名和密码"
		c.TplName = "register.html"
		return
	}
	o := orm.NewOrm()
	var userN models.User
	userN.Name = userName
	userN.PassWord = password
	//插入数据
	con, err := o.Insert(&userN)
	if err != nil {
		c.Data["errMsg￿"] = "注册失败"
		c.TplName = "register.html"
		logs.Error("插入数据失败")
		return
	}
	c.Data["errMsg￿"] = "注册成功"
	c.TplName = "register.html"
	logs.Info("插入数据成功-con-",con)
	//c.Ctx.WriteString("插入数据成功")
	/*重定向 跳转到登录
	url: 路径
	code：状态码*/
	c.Redirect("login", 302)
}