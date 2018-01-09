package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) MyRender(viewFile string) {
	this.TplName = viewFile
	this.Render()
}

func (this *BaseController) CheckLogin() bool {

	sess := this.StartSession()
	if sess != nil && sess.Get("name") != nil {
		fmt.Println("name: " + sess.Get("name").(string))
		return true
	} else {
		this.Redirect("/login", 302)
		return false
	}
}
