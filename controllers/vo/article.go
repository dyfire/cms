package vo

import (
	// "fmt"
	// "cms/models"
	// "cms/utils"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	// BaseController
	beego.Controller
}

func (this *ArticleController) Get() {
	m1 := make(map[string]string)
	m1["a"] = "aa"
	m1["b"] = "bb"
	this.Data["json"] = m1
	this.ServeJson()
}
