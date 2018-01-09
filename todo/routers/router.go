package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/samples/todo/controllers"
)

func init() {
	fmt.Println("routers init...")
	beego.Router("/", &controllers.TaskController{}, "get:Home")
	beego.Router("/index", &controllers.TaskController{}, "get:Home")

	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}
