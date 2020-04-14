package controllers

import (
	"encoding/json"
	"sdrms/models"
	"time"
)

type StuKqInfoController struct {
	BaseController
}

func (c *StuKqInfoController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *StuKqInfoController) Index() {
	//是否显示更多查询条件的按钮弃用，前端自动判断
	//c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "stukqinfo/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "stukqinfo/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = "true"
	c.Data["canDelete"] ="true"
}

func (c *StuKqInfoController) DataGrid() {
	userid := c.curUser.Id
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.StuKqInfoQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	if userid!=1{//1是超级管理员
		params.UserId=userid
	}
	//获取数据列表和总数
	data, total := models.StuKqInfoPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *StuKqInfoController) DataExport() {
	userid := c.curUser.Id
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.StuKqInfoQueryParam
	params.SnoLike=c.GetString("SnoLike")
	params.CourseNameLike=c.GetString("CourseNameLike")
	params.Limit=20000
	params.Offset=0

	if userid!=1{//1是超级管理员
		params.UserId=userid
	}

	fileName:=models.StuKqInfoDataExport(&params)
	c.Ctx.Output.Download(fileName,"学生考勤记录-"+time.Now().Format("2006-01-02")+".xlsx")
}

