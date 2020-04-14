package controllers

import (
	"beegodemo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {



	findUser()

	c.Data["content"] = "hello beego"
	c.TplName = "test.html"

}

func (c *MainController) Post() {
	c.Data["content"] = "beego.me"
	c.TplName = "test.html"

}

func findUser()  {
	o:=orm.NewOrm()
	user:=models.TUser{}
	user.Name="test01"

	o.Read(&user,"name")
	beego.Info(user)
}

func addUser(){
	o:=orm.NewOrm()
	user:=models.TUser{}
	user.Name="test01"
	user.Password="wwww"
	o.Insert(&user)
}
