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
	this.Data["Lang"] = "zh-CN"

	var FilterUser = func(ctx *context.Context) {
		if ctx.Input.Session("username") == nil {
			if ctx.Request.RequestURI != "/admin/login" {
				ctx.Redirect(302, "/admin/login")
			}
		}
	}

	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)

}
