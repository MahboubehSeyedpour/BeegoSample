package routers

import (
	"beegoSample/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/create_user/", &controllers.UserController{}, "post:Register")
	beego.Router("/login/", &controllers.UserController{}, "post:Login")
}
