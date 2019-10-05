package controllers

import (
	"myBeego/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*1.Generate one ORM object*/
	o := orm.NewOrm()
	/*2.Generate one object needed inserted*/
	user := models.User{}
	/*3.Set values for struct*/
	user.Name = "Tony"
	user.Pwd = "123"
	/*4.insert data*/
	_, err := o.Insert(&user)
	if nil != err {
		beego.Info("Insert Fail! ", err)
		return
	}
	beego.Info("Insert Success!")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
