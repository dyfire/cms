package admin

import (
	"cms/models"
	"github.com/astaxie/beego"
	"strings"
)

type UserController struct {
	BaseController
}

func (this *UserController) Login() {
	if this.Ctx.Input.IsPost() {
		username := strings.Trim(this.GetString("username"), "")
		password := strings.Trim(this.GetString("password"), "")

		m := models.User{Username: username}
		err := m.Read("username")
		if err != nil {
			this.Redirect("/admin/login", 302)
		}
		beego.Debug(m)
		if m.Password == password {
			this.SetSession("username", username)
			this.Redirect("/admin/article/list", 302)
		}
	}

	this.TplNames = "user_login.html"
	this.Render()
}

func (this *UserController) Logout() {
	this.SetSession("username", "")
	this.DestroySession()

	this.Redirect("/admin/login", 302)
}
