package controllers

import (
	"encoding/json"
	"sdrms/models"
)

type UserCourseController struct {
	BaseController
}

func (c *UserCourseController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *UserCourseController) Index() {
	//是否显示更多查询条件的按钮弃用，前端自动判断
	//c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "usercourse/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "usercourse/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = "true"
	c.Data["canDelete"] ="true"
}

func (c *UserCourseController) DataGrid() {
	userid := c.curUser.Id
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.UserCourseRefQueryParam
	params.UserId=userid
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.UserCourseRefPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
