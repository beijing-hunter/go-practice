package controllers

import (
	"github.com/astaxie/beego/orm"
	"sdrms/enums"
	"sdrms/models"
	"sdrms/utils"
	"time"
)

type CourseLimitedInfoController struct {
	BaseController
}

func (c *CourseLimitedInfoController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *CourseLimitedInfoController) Save() {

	userid := c.curUser.Id
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.CourseLimitedInfoQueryParam
	params.UserId=userid
	params.CourseId,_=c.GetInt("CourseId")

	query:=orm.NewOrm().QueryTable(models.CourseLimitedInfoTBName())
	query=query.Filter("userid",params.UserId)

	curTime:=time.Now()
	h,_:=time.ParseDuration("-1h")
	curTime=curTime.Add(h)//一小时之内

	query=query.Filter("createtime__gt",curTime)
	total,_:=query.Count()

	if total>0{
		c.jsonResult(enums.JRCodeFailed, "1小时之内不可重复发起限时考勤！", params.CourseId)
		return
	}

	o := orm.NewOrm()
	m := models.CourseLimitedInfo{}
	m.CourseId=params.CourseId
	m.UserId=params.UserId
	m.CreateTime=time.Now()
	m.StartTime=time.Now()

	mm,_:=time.ParseDuration("1m")//加1分钟
	m.EndTime=m.StartTime.Add(mm)
	m.Number=utils.RandomIntString(6)

	if _, err := o.Insert(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "发起失败", m.Id)
	}else{
		c.jsonResult(enums.JRCodeSucc, "发起成功", m.Number)
	}
}
