package utils

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
)

//type DbFacatory struct {

//}

func init() {

	dbusername := IniParserInstance.GetString("mysql", "dbusername")
	dbpassword := IniParserInstance.GetString("mysql", "dbpassword")
	dbhostip := IniParserInstance.GetString("mysql", "dbhostip")
	dbname := IniParserInstance.GetString("mysql", "dbname")
	var err error

	Db, err = sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8")
	Db.SetMaxOpenConns(int(IniParserInstance.GetInt64("mysql", "maxOpenConns")))
	Db.SetMaxIdleConns(int(IniParserInstance.GetInt64("mysql", "maxIdleConns")))
	Db.SetConnMaxLifetime(time.Minute * time.Duration(IniParserInstance.GetInt64("mysql", "maxLifetime")))

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
