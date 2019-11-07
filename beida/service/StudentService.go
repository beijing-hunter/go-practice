package service

import (
	"beida/dbp"
	"beida/utils"

	uuid "github.com/satori/go.uuid"
)

var (
	saltPfix = "sfeajofwefjwei"
)

type Student struct {
	Name     string
	Password string
	Salt     string
	Uuid     string
}

func FindInfo(name string) (datas []Student) {

	rows, err := dbp.Db.Query("select name,password,salt,uuid from student where name=?", name)

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		stu := Student{}
		err = rows.Scan(&stu.Name, &stu.Password, &stu.Salt, &stu.Uuid)

		if err != nil {
			return
		}

		datas = append(datas, stu)
	}

	defer rows.Close()
	return
}

func AddInfo(stu Student) (result bool) {

	result = false
	datas := FindInfo(stu.Name)

	if len(datas) > 0 {
		return
	}

	uuids := uuid.NewV4()
	stu.Uuid = uuids.String()
	stu.Password = utils.Md5V(stu.Password + saltPfix)
	stu.Salt = saltPfix

	sql := "insert into student(name,password,salt,uuid) values (?,?,?,?)"
	stmt, err := dbp.Db.Prepare(sql)

	if err != nil {
		panic(err)
		return
	}

	stmt.QueryRow(stu.Name, stu.Password, stu.Salt, stu.Uuid)
	result = true

	defer stmt.Close()
	return result
}

func Login(name string, password string) (datas []Student) {

	rows, err := dbp.Db.Query("select name,password,salt,uuid from student where name=? and password=?", name, password)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {

		stu := Student{}
		err = rows.Scan(&stu.Name, &stu.Password, &stu.Salt, &stu.Uuid)

		if err != nil {
			return
		}

		datas = append(datas, stu)
	}

	return
}

func updatePassword(uuid string, password string) {

	password = utils.Md5V(password + saltPfix)
	salt := saltPfix
	dbp.Db.Exec("update student set password=?,salt=? where uuid=?", password, salt, uuid)
}
