package admin

import (
	"cms/models"
	"github.com/astaxie/beego"
	"strings"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Add() {
	if this.Ctx.Input.IsPost() {
		title := strings.TrimSpace(this.GetString("title"))
		path := strings.TrimSpace(this.GetString("path"))
		is_visible, _ := this.GetInt("is_visible")
		order, _ := this.GetInt("order")
		err := models.AddCategory(title, path, is_visible, order)
		if err != nil {
			beego.Debug(err)
		}

		this.Redirect("/admin/category/add", 302)
	}

	lists, err := models.GetCategoryList()
	if err != nil {
		beego.Error(err)
	}

	this.Data["lists"] = lists
	this.TplNames = "category_add.html"
	this.Render()
}

func (this *CategoryController) Edit() {

}

func (this *CategoryController) List() {
	lists, err := models.GetCategoryList()
	if err != nil {
		beego.Error(err)
	}

	this.Data["lists"] = lists
	this.TplNames = "category_list.html"
	this.Render()
}
