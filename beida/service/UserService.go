package service

import (
	"beida/dbp"
	"beida/utils"

	uuid "github.com/satori/go.uuid"
)

var (
	saltPfix = "sfeajofwefjwei"
)

type User struct {
	Name     string
	Password string
	Salt     string
	Uuid     string
}

type UserService struct{}

func (s *UserService) FindInfo(name string) (datas []User) {

	rows, err := dbp.Db.Query("select name,password,salt,uuid from suser where name=?", name)

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		stu := User{}
		err = rows.Scan(&stu.Name, &stu.Password, &stu.Salt, &stu.Uuid)

		if err != nil {
			return
		}

		datas = append(datas, stu)
	}

	defer rows.Close()
	return
}

func (s *UserService) AddInfo(stu User) (result bool) {

	result = false
	datas := s.FindInfo(stu.Name)

	if len(datas) > 0 {
		return
	}

	uuids := uuid.NewV4()
	stu.Uuid = uuids.String()
	stu.Password = utils.Md5V(stu.Password + saltPfix)
	stu.Salt = saltPfix

	sql := "insert into suser(name,password,salt,uuid) values (?,?,?,?)"
	stmt, err := dbp.Db.Prepare(sql)
	defer stmt.Close()

	if err != nil {
		panic(err)
		return
	}

	stmt.QueryRow(stu.Name, stu.Password, stu.Salt, stu.Uuid)
	result = true

	return
}

func (s *UserService) Login(name string, password string) (datas []User) {

	password = utils.Md5V(password + saltPfix)
	rows, err := dbp.Db.Query("select name,password,salt,uuid from suser where name=? and password=?", name, password)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {

		stu := User{}
		err = rows.Scan(&stu.Name, &stu.Password, &stu.Salt, &stu.Uuid)

		if err != nil {
			return
		}

		datas = append(datas, stu)
	}

	return
}

func (s *UserService) UpdatePassword(uuid string, password string) {

	password = utils.Md5V(password + saltPfix)
	salt := saltPfix
	dbp.Db.Exec("update suser set password=?,salt=? where uuid=?", password, salt, uuid)
}
