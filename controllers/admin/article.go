package admin

import (
	// "fmt"
	"bytes"
	"cms/models"
	"cms/utils"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Add() {
	if this.Ctx.Input.IsPost() {
		title := strings.TrimSpace(this.GetString("title"))
		color := strings.TrimSpace(this.GetString("color"))
		category_id, _ := this.GetInt("category_id")
		content := strings.TrimSpace(this.GetString("content"))
		is_visible, _ := this.GetInt("is_visible")
		order, _ := this.GetInt("order")
		pic, _ := this.GetInt("pic")
		create_time := strings.TrimSpace(this.GetString("create_time"))
		ct, _ := time.Parse("2006-01-02 15:04:05", create_time)
		photos := this.GetStrings("picarr[]")
		descs := this.GetStrings("picarr_txt[]")

		m := models.Article{Title: title, Color: color, CategoryId: category_id, Content: content, IsVisible: is_visible, Order: order, Pic: pic, CreateTime: ct}
		err := m.Insert()
		if err != nil {
			pm := models.Photo{}
			pm.InsertAll(photos, descs)

		}
		this.Redirect("/admin/article/add", 302)
	}

	lists, err := models.GetCategoryList()
	if err != nil {
		beego.Error(err)
	}

	this.Data["lists"] = lists
	this.TplNames = "article_add.html"
	this.Render()
}

func (this *ArticleController) Edit() {
	id, _ := this.GetInt("id")
	if this.Ctx.Input.IsPost() {
		id, _ := this.GetInt("id")
		title := strings.TrimSpace(this.GetString("title"))
		color := strings.TrimSpace(this.GetString("color"))
		category_id, _ := this.GetInt("category_id")
		content := strings.TrimSpace(this.GetString("content"))
		is_visible, _ := this.GetInt("is_visible")
		create_time := strings.TrimSpace(this.GetString("create_time"))
		ct, _ := time.Parse("2006-01-02 15:04:05", create_time)
		order, _ := this.GetInt("order")
		pic, _ := this.GetInt("pic")

		beego.Debug(ct)
		m := models.Article{Id: id, Title: title, Color: color, CategoryId: category_id, Content: content, IsVisible: is_visible, Order: order, Pic: pic, CreateTime: ct}
		m.Update()

		i := strconv.Itoa(int(id))
		this.Redirect("/admin/article/edit?id="+i, 302)
	}

	m := models.Article{Id: id}
	if m.Read() != nil {
		this.Abort("404")
	}

	lists, err := models.GetCategoryList()
	if err != nil {
		beego.Error(err)
	}

	this.Data["rs"] = m
	this.Data["lists"] = lists
	this.TplNames = "article_edit.html"
	this.Render()
}

func (this *ArticleController) Delete() {
	id, _ := this.GetInt("id")
	m := models.Article{Id: id}
	if m.Read() == nil {
		m.Delete()
	}

	this.Redirect("/admin/article/list", 302)
}

func (this *ArticleController) List() {
	var (
		page   int64
		num    int64 = 20
		offset int64
	)

	m := models.Article{}

	if page, _ = this.GetInt("p"); page < 1 {
		page = 1
	}
	offset = (page - 1) * num
	lists := m.List(num, offset)

	total := m.Count()
	p := utils.NewPaginator(this.Ctx.Request, int(num), total)

	this.Data["lists"] = lists
	this.Data["paginator"] = p
	// this.Data["Lang"] = "zh-CN"
	this.TplNames = "article_list.html"
	this.Render()
}

func (this *ArticleController) Upload() {
	if this.Ctx.Input.IsPost() {
		_, header, err := this.GetFile("uploadify")
		if err != nil {
			beego.Debug("upload failtrue")
		}

		suffix := strings.ToLower(header.Filename[strings.LastIndex(header.Filename, "."):])

		file_name := time.Now().Format("20060102235959") + suffix
		save_path := "./static/attachments/" + file_name
		size := "1000"

		out := fmt.Sprintf("%s,%s,%s,%s", file_name, size, save_path, save_path)

		this.SaveToFile("uploadify", save_path)

		// this.Data["out"] = out
		// o := []byte{}
		// for k, v := range out {
		// 	o[k] = v
		// }
		ibytes := bytes.NewBufferString(out)
		icontent, _ := ioutil.ReadAll(ibytes)
		this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
		this.Ctx.Output.Body(icontent)

	} else {
		sess := this.StartSession()

		this.Data["session_id"] = sess.SessionID()
		this.Data["num"] = this.GetString("num")
		this.Data["size"] = this.GetString("size")
		this.Data["input"] = this.GetString("input")
		this.Data["area"] = this.GetString("area")
		this.Data["frame"] = this.GetString("frame")
		this.Data["title"] = this.GetString("title")
		this.Data["type"] = this.GetString("type")
		this.TplNames = "upload.html"
		this.Render()
	}
}
