package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	"strconv"
	// "time"
	"fmt"
	"strings"
)

type Category struct {
	Id        int64 `orm:"pk"`
	Pid       int
	Level     int64
	Title     string `orm:"size(100)"`
	Path      string `orm:"size(100)"`
	IsVisible int64
	Type      int64
	Order     int64
}

func init() {
	// orm.RegisterModel(new(Category))
	orm.RegisterModelWithPrefix("td_", new(Category))
}

func AddCategory(title string, path string, is_visible, order int64) error {
	o := orm.NewOrm()
	m := Category{Id: 0, Pid: 0, Level: 0, Title: title, Path: path, IsVisible: is_visible, Type: 1, Order: order}

	last_id, err := o.Insert(&m)
	if err == nil {
		m.createPath(last_id)
		if _, err := o.Update(&m); err == nil {
			return nil
		}
	}

	return err
}

func GetCategoryList() (lists []*Category, err error) {
	o := orm.NewOrm()
	ob := new(Category)
	qs := o.QueryTable(ob).Filter("is_visible", 1).OrderBy("path")
	_, err = qs.All(&lists)

	return lists, err
}

func (m *Category) createPath(category_id int64) error {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("id", category_id).One(m)

	if err != nil {

	}

	n := strings.Split(m.Path, "_")

	lens := len(n)
	m.Level = int64(lens)

	if lens == 1 {
		m.Path = "0_" + fmt.Sprintf("%d", category_id)
		m.Pid = 0
	} else {
		m.Path = m.Path + "_" + fmt.Sprintf("%d", category_id)
		lens = len(n)
		m.Pid, _ = strconv.Atoi(n[lens-1])

	}

	return nil
}
