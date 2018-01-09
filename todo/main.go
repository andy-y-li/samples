package main

import (
	"github.com/astaxie/beego"
	//"github.com/beego/samples/todo/controllers"
	_ "github.com/beego/samples/todo/routers"
)

func main() {
	//beego.Router("/", &controllers.BaseController{})
	//beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	//beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
	beego.Run()
}
