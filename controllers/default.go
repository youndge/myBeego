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
	/*ORM Insert*/
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

	/*ORM Query*/
	/*1.Generate one ORM object*/
	o1 := orm.NewOrm()
	/*2.Generate one object needed queried*/
	user1 := models.User{}
	/*3.Set values for struct*/
	//user1.Id = 1 /*read by Id*/
	user1.Name = "Tony" /*read by Name*/
	/*4.Read data*/
	//err1 := o1.Read(&user1)/*read by Id*/
	err1 := o1.Read(&user1, "Name")/*read by Name*/
	if nil != err1 {
		beego.Info("Read Fail! ", err1)
		return
	}
	beego.Info("Read Success!")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
