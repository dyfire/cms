package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// "fmt"
	// "strconv"
	// "strings"
	"errors"
	"time"
)

type Photo struct {
	Id         int64  `orm:"pk"`
	Title      string `orm:size(200)`
	Path       string `orm:size(100)`
	Filename   string `orm:size(50)`
	PostId     int64
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

func (m *Photo) InsertAll(path, desc []string, post_id int) error {
	if len(path) == 0 {
		return errors.New("no path")
	}

	ps := []*Photo{}

	for i := 0; i < len(path); i++ {
		ps = append(ps, &Photo{Title: desc[i], Path: path[i]})
	}

	for _, v := range ps {
		id, _ := orm.NewOrm().Insert(v)

		orm.NewOrm().Update(&Photo{Title: v.Title, Path: v.Path, PostId: int64(post_id), Id: int64(id)})
		beego.Debug(id)
	}
	// if err != nil {
	// 	return err
	// }

	return nil
}
