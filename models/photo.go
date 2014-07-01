package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// "fmt"
	// "strconv"
	"errors"
	"strings"
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
	create_time := time.Now().Format("2006-01-02 15:04:05")
	ct, _ := time.Parse("2006-01-02 15:04:05", create_time)

	for i := 0; i < len(path); i++ {
		filename := strings.Split(path[i], "/")
		f := filename[len(filename)-1]
		ps = append(ps, &Photo{Title: desc[i], Path: strings.TrimLeft(path[i], "."), Filename: f, CreateTime: ct})
	}

	for _, v := range ps {
		id, _ := orm.NewOrm().Insert(v)
		orm.NewOrm().Update(&Photo{PostId: int64(post_id), Id: int64(id)})

	}

	return nil
}

func (m *Photo) GetLists(post_id, num, offset int) ([]*Photo, error) {
	var p []*Photo
	o := orm.NewOrm()
	_, err := o.QueryTable(m).Filter("post_id", post_id).Offset(offset).Limit(num).All(&p)

	if err != nil {
		return nil, err
	}

	return p, nil
}
