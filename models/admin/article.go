package admin

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
	"time"
)

type Img struct {
	Id    int64 `orm:"pk"`
	Title string
}

type Article struct {
	Id         int64  `orm:"pk"`
	Title      string `valid:"Required"`
	TitleColor string
	Content    string
	ShowType   int64
	From       string
	Author     string
	Keywords   string
	Imgs       string
	Hits       int64
	Order      int64
	IsPass     int64
	CreateTime time.Time
}

func (a *Article) Insert() error {

}
