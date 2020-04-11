package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

/*展示文章列表页*/
func (c *ArticleController) ShowList() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

/*展示添加文章页面*/
func (c *ArticleController) ShowAddArticle() {
	c.TplName = "add.html"
}

/*处理添加文章*/
func (c *ArticleController) HandleAddArticle() {
	logs.Info("处理--添加文章")
	articleName := c.GetString("articleName")
	content := c.GetString("content")
	logs.Info("user-",articleName,"pw-",content)

	if articleName == "" || content == "" {
		logs.Error("数据有误")
		c.Data["errMsg"] = "请输入正确数据"
		c.TplName = "add.html"
		return
	}
	//获取上传文件
	file,fileHeader,fileErr := c.GetFile("uploadname")
	//关闭文件
	defer file.Close()
	if fileErr != nil {
		logs.Error("文件错误--",fileErr.Error())
		c.Data["errMsg"] = "文件错误--"
		c.TplName = "add.html"
		return
	}

	saveErr := c.SaveToFile("uploadname","./static/img/"+fileHeader.Filename)
	if saveErr != nil {
		logs.Error("文件保存失败--",saveErr.Error())
		c.Data["errMsg"] = "文件保存失败--"
		c.TplName = "add.html"
		return
	}
	// 1 文件大小 字节
	if fileHeader.Size > 600000 {
		logs.Error("文件太大")
		c.Data["errMsg"] = "文件太大,禁止上传"
		c.TplName = "add.html"
		return
	}
	// 2 文件格式
	//fileHeader.Filename
	// 3 防止重名

	c.Data["errMsg￿"] = "登录成功"
	c.TplName = "add.html"

}