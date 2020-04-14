package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

type  TUser struct {
	Id int
	Name string
	Password string
}

func init()  {

	orm.RegisterDataBase("default","mysql","root:cmt123@tcp(123.56.4.34:3306)/azkaban?charset=utf8")
	orm.RegisterModel(new(TUser))
	orm.RunSyncdb("default",false,true)
}
