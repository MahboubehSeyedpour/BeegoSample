package routers

import (
	"beegoSample/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// user
	beego.Router("/register/", &controllers.UserController{}, "post:Register")
	beego.Router("/login/", &controllers.UserController{}, "post:Login")

	// category
	beego.Router("/create_category/", &controllers.CategoryController{}, "post:Create")
}
