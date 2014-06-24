package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// "fmt"
	// "strconv"
	// "strings"
	"time"
)

type Photo struct {
	ID         int64  `orm:"pk"`
	Title      string `orm:size(200)`
	Path       string `orm:size(100)`
	Filename   string `orm:size(50)`
	CreateTime time.Time
}

func init() {
	orm.RegisterModelWithPrefix("td_", new(Photo))
}

func (m *Photo) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}

	return nil
}

func (m *Photo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}

	return nil
}

func (m *Photo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}

	return nil
}

func (m *Photo) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}

	return nil
}

func (m *Photo) InsertAll(path, desc []string) error {
	return nil
}
