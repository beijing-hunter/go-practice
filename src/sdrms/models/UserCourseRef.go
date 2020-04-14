package models

import "github.com/astaxie/beego/orm"

type UserCourseRef struct {
	Id int
	UserId int
	CourseId int
}

type UserCourseDto struct {
	CourseId int
	CourseName string
	UserId int
	LimitedNumber string
}

type UserCourseRefQueryParam struct {
	BaseQueryParam
	UserId int
}

func (a *UserCourseRef) TableName() string {
	return UserCourseRefTBName()
}

func UserCourseRefPageList(params *UserCourseRefQueryParam) ([]*UserCourseDto, int64) {
	query := orm.NewOrm()
	var data []*UserCourseDto
	sql:="SELECT uc.course_id,c.`name` as course_name,uc.user_id "+
		"FROM rms_user_course_ref uc INNER JOIN rms_course c ON uc.course_id=c.id "+
		"WHERE uc.user_id=? "+
		"ORDER BY uc.course_id desc "+
		"LIMIT ?,?"
	query.Raw(sql,params.UserId,params.Offset,params.Limit).QueryRows(&data)

	if len(data)>0{
		for _,item:=range data{
			item.LimitedNumber=CourseLimitedInfoNewNumberByUserId(params.UserId,item.CourseId)
		}
	}

	query2:=orm.NewOrm().QueryTable(UserCourseRefTBName())
	query2=query2.Filter("userid",params.UserId)
	total,_:=query2.Count()

	return data, total
}

//教师所教授的课程
func UserCourseDataListByUserId(userId int)[]*UserCourseRef  {
	query := orm.NewOrm().QueryTable(UserCourseRefTBName())
	data := make([]*UserCourseRef, 0)
	query=query.Filter("userid",userId)
	//默认排序
	sortorder := "Id"
	query.OrderBy(sortorder).Limit(-1, 0).All(&data)
	return data
}

func UserCourseDataDeleteByUserId(userId int) {
	query := orm.NewOrm().QueryTable(UserCourseRefTBName())
	query=query.Filter("userid",userId)
	query.Delete()
}
