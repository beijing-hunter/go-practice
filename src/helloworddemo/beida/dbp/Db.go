package dbp

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbusername = "root"
	dbpassword = "12234"
	dbhostip   = "123.56.24.34:3306"
	dbname     = "db20190813"
	Db         *sql.DB
)

//type DbFacatory struct {

//}

func init() {
	var err error

	Db, err = sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8")

	if err != nil {
		panic(err)
	}
}
