package routers

import (
	// "cms/controllers"
	"cms/controllers/admin"
	"cms/controllers/vo"
	// "fmt"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/category/list", &admin.CategoryController{}, "GET:List")
	beego.Router("/admin/category/add", &admin.CategoryController{}, "*:Add")
	beego.Router("/admin/category/edit", &admin.CategoryController{}, "*:Edit")
	beego.Router("/admin/category/delete", &admin.CategoryController{}, "*:Delete")

	beego.Router("/admin/article/list", &admin.ArticleController{}, "GET:List")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/article/upload", &admin.ArticleController{}, "GET:Upload")
	beego.Router("/admin/article/img", &admin.ArticleController{}, "*:Img")
	beego.Router("/admin/article", &admin.ArticleController{}, "GET:Upload")

	beego.Router("/admin/upload/upload", &admin.UploadController{}, "*:Upload")

	beego.Router("/admin/login", &admin.UserController{}, "*:Login")
	beego.Router("/admin/logout", &admin.UserController{}, "*:Logout")

	beego.Router("/article/list", &vo.ArticleController{})
}
