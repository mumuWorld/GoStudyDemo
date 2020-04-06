package routers

import (
	"github.com/astaxie/beego"
	"../controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/register",&controllers.RegisterController{},"get:ShowRegister;post:HandlePost")


}
