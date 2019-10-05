package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}

func init() {
	// set default database (database's name:test,user:root,pwd:1234)
	orm.RegisterDataBase("default", "mysql", "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	// register model
	orm.RegisterModel(new(User))
	// create table
	orm.RunSyncdb("default", false, true)
}
