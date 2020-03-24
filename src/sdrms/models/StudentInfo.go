package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func (a *StudentInfo) TableName() string {
	return StudentInfoTBName()
}

type StudentInfoQueryParam struct {
	BaseQueryParam
	SnoLike string
	RealNameLike string
}

type StudentInfo struct {
	Id int
	Sno string
	RealName string
	Password string
	CreateTime time.Time
}

func StudentInfoPageList(params *StudentInfoQueryParam) ([]*StudentInfo, int64) {
	query := orm.NewOrm().QueryTable(StudentInfoTBName())
	data := make([]*StudentInfo, 0)
	//默认排序
	sortorder := "Id"

	if len(params.SnoLike)>0{
		query=query.Filter("sno__contains",params.SnoLike)
	}

	if len(params.RealNameLike)>0{
		query=query.Filter("realname__contains",params.RealNameLike)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func StudentInfoOne(id int) (*StudentInfo, error) {
	o := orm.NewOrm()
	m := StudentInfo{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func StudentInfoSonCount(id int,sno string) (bool) {
	query := orm.NewOrm().QueryTable(StudentInfoTBName())
	query=query.Filter("sno",sno)
	c,_:= query.Count()

	if c==0{
		return  false
	}

	if id>0{

		query=query.Filter("id",id)
		c,_:= query.Count()

		if c==1{
			return  false
		}

		return true
	}

	return c>0
}

