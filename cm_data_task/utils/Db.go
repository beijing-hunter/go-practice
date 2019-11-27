package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbusername = "124ww"
	dbpassword = "12wdw134"
	dbhostip   = "123.562.34.314:3306"
	dbname     = "12323"
	Db         *sql.DB
)

//type DbFacatory struct {

//}

func init() {
	var err error

	Db, err = sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8")
	Db.SetMaxOpenConns(100)
	Db.SetMaxIdleConns(40)

	if err != nil {
		panic(err)
	}
}

func ErrorCatch() {
	error := recover() //异常接受内置函数
	if error != nil {
		InfoLogger.Println(error)
	}
}
