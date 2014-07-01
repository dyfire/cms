package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.initialization()
}

func (this *BaseController) initialization() {
	this.Data["a"] = "ok"

	var FilterUser = func(ctx *context.Context) {
		beego.Debug(ctx.Input.Session("login"))
		if ctx.Input.Session("login") == nil {
			this.Redirect("/admin/login", 302)
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

}
