package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
	"os"
	"strconv"
	"time"
)

type StuKqInfo struct {
	Id int
	StuId int
	CourseId int
	UserId int
	CourseLimitedId int
	CreateTime time.Time
	Status int
}

type StuKqInfoDto struct {
	CourseLimitedId int
	CourseName string
	RealName string
	StartTime time.Time
	EndTime time.Time
	LimitedNumber string

	StuRealName string
	Sno string
	StatusName string
	Status int

	StartTimeStr string
	EndTimeStr string
}

type StuKqInfoCountDto struct {
	TotalCount int
}

func (a *StuKqInfo) TableName() string {
	return StuKqInfoTBName()
}

type StuKqInfoQueryParam struct {
	BaseQueryParam
	SnoLike string
	CourseNameLike string
	UserId int
	StuId int
}

//查询学生考勤记录
func StuKqInfoPageList(params *StuKqInfoQueryParam) ([]*StuKqInfoDto, int) {
	query := orm.NewOrm()
	data := make([]*StuKqInfoDto, 0)
	sql:="SELECT cli.id as course_limited_id,c.`name` as course_name,bu.real_name,cli.start_time,cli.end_time,cli.number as limited_number,stu.real_name as stu_real_name,stu.sno,sk.`status`,(case sk.`status` when 0 THEN '已打卡' ELSE '缺勤' end) as status_name "
	sqlCount:="SELECT count(1) as total_count"
	sqlCommon:=" FROM rms_course_limited_info cli "+
		"INNER JOIN rms_course c ON cli.course_id=c.id "+
		"INNER JOIN rms_backend_user bu ON cli.user_id=bu.id "+
		"INNER JOIN rms_stu_kq_info sk ON cli.id=sk.course_limited_id "+
		"INNER JOIN rms_student_info stu ON sk.stu_id=stu.id "+
		"WHERE 1=1 "
	sql=sql+sqlCommon
	sqlCount=sqlCount+sqlCommon
	var args []interface{}

	if params.UserId>0{
		sql=sql+"and cli.user_id=? "
		sqlCount=sqlCount+"and cli.user_id=? "
		args=append(args,params.UserId)
	}

	if len(params.CourseNameLike)>0{
		sql=sql+" and c.`name`=? "
		sqlCount=sqlCount+" and c.`name`=? "
		args=append(args,params.CourseNameLike)
	}

	if len(params.SnoLike)>0{
		sql=sql+"AND stu.sno=? "
		sqlCount=sqlCount+"AND stu.sno=? "
		args=append(args,params.SnoLike)
	}

	if params.StuId>0{
		sql=sql+"AND stu.id=? "
		sqlCount=sqlCount+"AND stu.id=? "
		args=append(args,params.StuId)
	}

	var totalCountData StuKqInfoCountDto
	query.Raw(sqlCount,args).QueryRow(&totalCountData)

	sql=sql+" ORDER BY cli.start_time desc LIMIT ?,?"
	args=append(args,params.Offset)
	args=append(args,params.Limit)
	query.Raw(sql,args).QueryRows(&data)

	return data, totalCountData.TotalCount
}

//导出考勤数据excel
func StuKqInfoDataExport(params *StuKqInfoQueryParam) string  {

	datas,_:=StuKqInfoPageList(params)

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "课程名"
	cell = row.AddCell()
	cell.Value = "教师姓名"
	cell = row.AddCell()
	cell.Value = "考勤开启时间"
	cell = row.AddCell()
	cell.Value = "学生姓名"
	cell = row.AddCell()
	cell.Value = "学号"
	cell = row.AddCell()
	cell.Value = "考勤状态"


	for _,kqInfoItem:=range datas{
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = kqInfoItem.CourseName
		cell = row.AddCell()
		cell.Value = kqInfoItem.RealName
		cell = row.AddCell()
		cell.Value = kqInfoItem.StartTime.Format("2006-01-02 15:04:05")
		cell = row.AddCell()
		cell.Value = kqInfoItem.StuRealName
		cell = row.AddCell()
		cell.Value = kqInfoItem.Sno
		cell = row.AddCell()
		cell.Value = kqInfoItem.StatusName
	}

	if !utils.FileExists("logs") {
		os.MkdirAll("logs", os.ModePerm)
	}

	filename := "logs/" + cast.ToString(time.Now().Unix()) + ".xlsx"
	file.Save(filename)
	return filename
}


//定时查询缺勤的学生信息并将缺勤信息补充的考勤记录表中
func StuKqInfoTimeTask()  {

	for {

		time.Sleep(time.Second*180)

		curTime := time.Now()
		h, _ := time.ParseDuration("-1h")
		curTime = curTime.Add(h) //一小时之内

		courseLimitedData := make([]*CourseLimitedInfo, 0)
		courseLimitedQuery := orm.NewOrm().QueryTable(CourseLimitedInfoTBName())
		courseLimitedQuery = courseLimitedQuery.Filter("createtime__gt", curTime)
		courseLimitedQuery.All(&courseLimitedData)

		if len(courseLimitedData) > 0 {

			for _, courseLimitedItem := range courseLimitedData {

				if courseLimitedItem.EndTime.Unix()>time.Now().Unix(){//课程限时打卡时间未结束，不做处理
					continue
				}

				stuKqQuery := orm.NewOrm().QueryTable(StuKqInfoTBName())
				stuKqQuery = stuKqQuery.Filter("courselimitedid", courseLimitedItem.Id)

				stuKqData := make([]*StuKqInfo, 0)
				stuKqQuery.All(&stuKqData)
				stuIds := make([]int, 0)

				if len(stuKqData) > 0 {

					for _, kqInfo := range stuKqData {
						stuIds = append(stuIds, kqInfo.StuId)
					}
				}

				//学生创建时间，必须小于考勤发起时间
				studentInfoSql := "SELECT s.id,s.sno,s.real_name FROM rms_student_info s where s.create_time < '"+courseLimitedItem.CreateTime.Format("2006-01-02 15:04:05")+"'"

				if len(stuIds) > 0 {
					studentInfoSql = studentInfoSql + " and s.id not in ( "
					for _, stuId := range stuIds {
						studentInfoSql = studentInfoSql + strconv.Itoa(stuId) + ","
					}

					studentInfoSql = studentInfoSql + "0)"
				}

				var studentInfoData []*StudentInfo //缺勤的学生信息
				studentInfoQuery := orm.NewOrm()
				studentInfoQuery.Raw(studentInfoSql).QueryRows(&studentInfoData)

				if len(studentInfoData) > 0 {

					stuKqRecords := make([]*StuKqInfo, 0)

					for _, queQinStudentInfoItem := range studentInfoData {

						stuKqInfo := new(StuKqInfo)
						stuKqInfo.CourseId = courseLimitedItem.CourseId
						stuKqInfo.StuId = queQinStudentInfoItem.Id
						stuKqInfo.CreateTime = time.Now()
						stuKqInfo.UserId = courseLimitedItem.UserId
						stuKqInfo.CourseLimitedId = courseLimitedItem.Id
						stuKqInfo.Status = 1

						stuKqRecords = append(stuKqRecords, stuKqInfo)
					}

					stuKqOrm := orm.NewOrm()
					stuKqOrm.InsertMulti(len(stuKqRecords), stuKqRecords) //补充缺勤的学科考勤记录
				}
			}

		}
	}
}
