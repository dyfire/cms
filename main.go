package main

import (
	_ "cms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func main() {
	loadDb()
	beego.AddFuncMap("repeat", repeat)
	beego.Run()
}

func repeat(sum int64) string {
	return strings.Repeat("---", int(sum))
}

func loadDb() {
	host := beego.AppConfig.String("dbhost")
	port := beego.AppConfig.String("dbport")
	user := beego.AppConfig.String("dbuser")
	pass := beego.AppConfig.String("dbpassword")
	name := beego.AppConfig.String("dbname")

	orm.Debug = true
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8"

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
}
