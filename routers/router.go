package routers

import (
	admin "github.com/bakanis/simpleauth/controllers"

	"github.com/bakanis/simpleauth-sample/controllers"

	"github.com/astaxie/beego"
)

func init() {
	///// auth module
	beego.Router("/login", &admin.AdminController{}, "get:LoginShow;post:LoginDo")
	beego.Router("/logout", &admin.AdminController{}, "get:LogoutDo")

	///// protected path
	beego.Router("/admin/home", &controllers.AdminController{})

	///// public
	beego.Router("/", &controllers.MainController{})
}
