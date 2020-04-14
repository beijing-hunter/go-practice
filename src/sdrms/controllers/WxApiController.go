package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"sdrms/enums"
	"sdrms/models"
	"sdrms/utils"
	"strconv"
	"time"
)

type WxApiController struct {
	BaseController
}

func (c *WxApiController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//学生登录
func (c *WxApiController) Login()  {

	var params models.WxApiParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	params.StuPassword= utils.String2md5(params.StuPassword)

	query:=orm.NewOrm().QueryTable(models.StudentInfoTBName())
	query=query.Filter("sno",params.Sno).Filter("password",params.StuPassword)
	data := make([]*models.StudentInfo, 0)
	n,_:=query.All(&data)

	if n==0{
		c.jsonResult(enums.JRCodeFailed, "学号或密码错误",params.Sno)
		return
	}

	c.jsonResult(enums.JRCodeSucc, "登录成功",data[0])
}

//考勤打卡
func (c *WxApiController) KqDaKa()  {

	var params models.WxApiParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	courseLimitedQuery:=orm.NewOrm().QueryTable(models.CourseLimitedInfoTBName())
	courseLimitedQuery=courseLimitedQuery.Filter("id",params.CourseLimitedId).Filter("number",params.LimitedNumber)

	data := make([]*models.CourseLimitedInfo, 0)
	n,_:=courseLimitedQuery.All(&data)

	if n==0{
		c.jsonResult(enums.JRCodeFailed, "课程考勤码输入错误",params.Sno)
		return
	}

	courseLimitedInfo:=data[0]

	if courseLimitedInfo.EndTime.Unix()<time.Now().Unix(){
		c.jsonResult(enums.JRCodeFailed, "课程考勤码失效",params.Sno)
		return
	}

	stuKqInfoQuery:=orm.NewOrm().QueryTable(models.StuKqInfoTBName())
	stuKqInfoQuery=stuKqInfoQuery.Filter("stuid",params.StuId).Filter("courselimitedid",params.CourseLimitedId)
	n,_=stuKqInfoQuery.Count()

	if n>0{
		c.jsonResult(enums.JRCodeFailed, "重复打卡",params.Sno)
		return
	}

	 stuKqInfo :=models.StuKqInfo{}
	stuKqInfo.CourseLimitedId=courseLimitedInfo.Id
	stuKqInfo.StuId=params.StuId
	stuKqInfo.UserId=courseLimitedInfo.UserId
	stuKqInfo.CreateTime=time.Now()
	stuKqInfo.Status=0
	stuKqInfo.CourseId=courseLimitedInfo.CourseId
	stuKqInfoOrm:=orm.NewOrm()
	stuKqInfoOrm.Insert(&stuKqInfo)
	c.jsonResult(enums.JRCodeSucc, "打卡成功",params.Sno)
}

//今日限时考勤课程
func (c *WxApiController) TodayCourseLimited()  {

	var params models.WxApiParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	curTime:=time.Now()
	h,_:=time.ParseDuration(strconv.Itoa(curTime.Hour())+"h")
	m,_:=time.ParseDuration(strconv.Itoa(curTime.Minute())+"m")
	curTime=curTime.Add(-h).Add(-m)

	sql:="SELECT cli.id as course_limited_id,c.`name` as course_name,bu.real_name,cli.start_time,cli.end_time,cli.number as limited_number "+
		"FROM rms_course_limited_info cli "+
		"INNER JOIN rms_course c ON cli.course_id=c.id "+
		"INNER JOIN rms_backend_user bu ON cli.user_id=bu.id "+
		"WHERE cli.create_time>'"+curTime.Format("2006-01-02 15:04:05")+"' "+
		"ORDER BY cli.create_time desc"

	query := orm.NewOrm()
	data := make([]*models.StuKqInfoDto, 0)
	query.Raw(sql).QueryRows(&data)

	for _,item:=range data{

		item.StartTimeStr=item.StartTime.Format("2006-01-02 15:04:05")
		item.EndTimeStr=item.EndTime.Format("2006-01-02 15:04:05")

		stuKqInfoQuery:=orm.NewOrm().QueryTable(models.StuKqInfoTBName())
		stuKqInfoQuery=stuKqInfoQuery.Filter("stuid",params.StuId).Filter("courselimitedid",item.CourseLimitedId)

		stuKqInfoDatas := make([]*models.StuKqInfo, 0)
		n,_:=stuKqInfoQuery.All(&stuKqInfoDatas)

		if n==0 {
			if item.EndTime.Unix() < time.Now().Unix() {
				item.Status = 1
				item.StatusName = "缺勤"
			} else {
				item.Status = -1
				item.StatusName = "未打卡"
			}
		}else{
			item.Status=stuKqInfoDatas[0].Status

			if stuKqInfoDatas[0].Status==0{
				item.StatusName="已打卡"
			}else{
				item.StatusName = "缺勤"
			}
		}
	}

	c.jsonResult(enums.JRCodeSucc, "今日限时考勤课程",data)
}

//学生考勤记录
func (c *WxApiController) StuKqRecord() {

	var params models.WxApiParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	 stuKqParams :=new(models.StuKqInfoQueryParam)
	stuKqParams.StuId=params.StuId
	stuKqParams.Limit=10000
	stuKqParams.Offset=0
	stuKqInfoDatas,_:=models.StuKqInfoPageList(stuKqParams)

	for _,item:=range stuKqInfoDatas{
		item.StartTimeStr=item.StartTime.Format("2006-01-02 15:04:05")
		item.EndTimeStr=item.EndTime.Format("2006-01-02 15:04:05")
	}

	c.jsonResult(enums.JRCodeSucc, "学生考勤记录",stuKqInfoDatas)
}


