package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CourseLimitedInfo struct {
	Id int
	CourseId int
	UserId int
	Number string
	CreateTime time.Time
	StartTime time.Time
	EndTime time.Time
}

type CourseLimitedInfoQueryParam struct {
	BaseQueryParam
	CourseId int
	UserId int
}

func (a *CourseLimitedInfo) TableName() string {
	return CourseLimitedInfoTBName()
}

//获取课程当前有效的限制考勤编号
func CourseLimitedInfoNewNumberByUserId(userId int,courseId int) string  {
	query := orm.NewOrm().QueryTable(CourseLimitedInfoTBName())
	data := make([]*CourseLimitedInfo, 0)
	query=query.Filter("userid",userId)
	query=query.Filter("courseid",courseId)
	query=query.Filter("endtime__gte",time.Now())
	//默认排序
	sortorder := "-Id"
	query.OrderBy(sortorder).Limit(-1, 0).All(&data)

	if len(data)>0{
		return  data[0].Number
	}

	return ""
}




