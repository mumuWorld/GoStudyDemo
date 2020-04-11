package routers

import (
	"github.com/astaxie/beego"
	"../controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/register",&controllers.RegisterController{},"get:ShowRegister;post:HandlePost")
	beego.Router("/showArticleList",&controllers.ArticleController{},"get:ShowList")
	beego.Router("/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArticle")


}
