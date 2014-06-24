package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	"fmt"
	// "strconv"
	// "strings"
	"time"
)

type Article struct {
	Id         int64  `orm:"pk"`
	Title      string `orm:"size(450)"`
	Color      string `orm:"size(20)"`
	CategoryId int64
	Content    string `orm:"type(text)"`
	IsVisible  int64
	Order      int64
	Pic        int64 `orm:"-"`
	CreateTime time.Time
}

func init() {
	orm.RegisterModelWithPrefix("td_", new(Article))
}

func (m *Article) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}

	return nil
}

func (m *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}

	return nil
}

func (m *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}

	return nil
}

func (m *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}

	return nil
}

func (m *Article) List(num, offset int64) []*Article {
	o := orm.NewOrm()
	lists := []*Article{}
	o.QueryTable(m).OrderBy("-id").Limit(num).Offset(offset).All(&lists)

	return lists
}

func (m *Article) Count() int64 {
	o := orm.NewOrm()
	total, _ := o.QueryTable(m).Count()
	return total
}

func (m *Article) ColorTitle() string {
	if m.Color != "" {
		return fmt.Sprintf("<span style=\"color:%s\">%s</span>", m.Color, m.Title)
	}

	return m.Title
}
