package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// "fmt"
	// "strconv"
	// "errors"
	// "strings"
	// "time"
)

type User struct {
	Id       int64  `orm:"pk"`
	Username string `orm:size(100)`
	Password string `orm:size(100)`
}

func init() {
	orm.RegisterModelWithPrefix("td_", new(User))
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}

	return nil
}
