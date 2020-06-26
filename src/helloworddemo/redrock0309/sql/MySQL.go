package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "strconv"
)

var DB *gorm.DB

func Link() (err error){

	DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	return nil
}

func DbConn() *gorm.DB {

	userName:="rootsss"
	password:="fsefe"
	host:="143.516.4.324"
	dbName:="fewsfe" 
	port :=3306
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName,password, host, port, dbName )
	db, err := gorm.Open("mysql", connArgs)
	
	if err != nil {
		fmt.Println("err:", err)
		return nil
	}

	db.SingularTable(true)
	return db
}
//连接数据库（尽量保证在登录注册的接口里可以不用再写一遍，redis也是的）