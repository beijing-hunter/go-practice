package service

import (
	"beida/dbp"
	"fmt"
)

type School struct {
	Id   int64
	Name string
	Area string
}

type SchoolService struct{}

func (s *SchoolService) FindInfos(searchKey string, area string, pageIndex int64, pageSize int64) (datas []School) {

	pageIndex = (pageIndex - 1) * pageSize
	sql := "select id,name,area from school where 1=1 "

	if len(area) > 0 {
		sql = sql + " and area='" + area + "'"
	}

	if len(searchKey) > 0 {
		sql = sql + " and name like '%" + searchKey + "%'"
	}

	sql = sql + " order by id limit ?,?"

	fmt.Println(sql)
	rows, err := dbp.Db.Query(sql, pageIndex, pageSize)
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {
		info := School{}

		err = rows.Scan(&info.Id, &info.Name, &info.Area)

		if err != nil {
			panic(err)
			return
		}

		datas = append(datas, info)
	}

	return
}

func (s *SchoolService) AddInfo(info School) (result bool) {

	result = false

	datas := s.FindInfos(info.Name, "", 1, 10)

	if len(datas) > 0 {
		return
	}

	sql := "insert into school(name,area) values(?,?)"
	stmt, err := dbp.Db.Prepare(sql)
	defer stmt.Close()

	if err != nil {
		panic(err)
		return
	}

	stmt.QueryRow(info.Name, info.Area)
	result = true

	return
}

func (s *SchoolService) Del(id int64) {

	sql := "delete from school where id=?"
	dbp.Db.Exec(sql, id)
}
