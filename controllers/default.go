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
	// /*ORM Insert*/
	// /*1.Generate one ORM object*/
	// o := orm.NewOrm()
	// /*2.Generate one object needed inserted*/
	// user := models.User{}
	// /*3.Set the value you want to insert*/
	// user.Name = "Tony"
	// user.Pwd = "123"
	// /*4.insert data*/
	// _, err := o.Insert(&user)
	// if nil != err {
	// 	beego.Info("Insert Fail! ", err)
	// 	return
	// }
	// beego.Info("Insert Success!")

	/*ORM Query*/
	/*1.Generate one ORM object*/
	o1 := orm.NewOrm()
	/*2.Generate one object needed queried*/
	user1 := models.User{}
	/*3.Set the value you want to find*/
	//user1.Id = 1 /*read by Id*/
	user1.Name = "Tony" /*read by Name*/
	/*4.Read data*/
	//err1 := o1.Read(&user1)/*read by Id*/
	err1 := o1.Read(&user1, "Name") /*read by Name*/
	if nil != err1 {
		beego.Info("Read Fail! ", err1)
		return
	}
	beego.Info("Read Success!")

	/*ORM Update*/
	/*1.Generate one ORM object*/
	o2 := orm.NewOrm()
	/*2.Generate one object needed updated*/
	user2 := models.User{}
	/*3.Set the value you want to uodate*/
	user2.Id = 1 /*update by Id*/
	/*4.Find data & Update it*/
	err2 := o2.Read(&user2) /*read by Id*/
	if nil != err2 {
		beego.Info("Update Fail! ", err2)
		return
	} else {
		user2.Name = "Tom"
		user2.Pwd = "222"
		_, err2 = o2.Update(&user2)
		if nil != err2 {
			beego.Info("Update Fail! ", err2)
		}
	}
	beego.Info("Update Success!")

	/*ORM Delete*/
	/*1.Generate one ORM object*/
	o3 := orm.NewOrm()
	/*2.Generate one object need to delete*/
	user3 := models.User{}
	/*3.Set the value you want to delete*/
	user3.Id = 1 /*update by Id*/
	//user2.Name = "Tony" /*read by Name*/
	/*4.Find data & Update it*/
	_, err3 := o3.Delete(&user3) /*read by Id*/
	if nil != err3 {
		beego.Info("Delete Fail! ", err3)
		return
	}
	beego.Info("Delete Success!")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
