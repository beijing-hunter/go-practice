package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"sdrms/enums"
	"sdrms/models"
	"sdrms/utils"
	"strconv"
	"strings"
	"time"
)

type StudentInfoController struct {
	BaseController
}

func (c *StudentInfoController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *StudentInfoController) Index() {
	//是否显示更多查询条件的按钮弃用，前端自动判断
	//c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "studentinfo/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "studentinfo/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = "true"
	c.Data["canDelete"] ="true"
}

func (c *StudentInfoController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.StudentInfoQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.StudentInfoPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *StudentInfoController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &models.StudentInfo{}
	var err error
	if Id > 0 {
		m, err = models.StudentInfoOne(Id)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}

	c.Data["m"] = m
	c.setTpl("studentinfo/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "studentinfo/edit_footerjs.html"
}

func (c *StudentInfoController) Save() {
	m := models.StudentInfo{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
		return
	}

	if models.StudentInfoSonCount(m.Id,m.Sno){
		c.jsonResult(enums.JRCodeFailed, "学号重复", m.Id)
		return
	}

	if m.Id == 0 {

		if len(m.Password) == 0 {
			m.Password="123456"
		}

		m.CreateTime=time.Now()
		//对密码进行加密
		m.Password = utils.String2md5(m.Password)
		if _, err := o.Insert(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		if oM, err := models.StudentInfoOne(m.Id); err != nil {
			c.jsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", m.Id)
		} else {
			m.Password = strings.TrimSpace(m.Password)
			if len(m.Password) == 0 {
				//如果密码为空则不修改
				m.Password = oM.Password
			} else {
				m.Password = utils.String2md5(m.Password)
			}
			m.CreateTime=oM.CreateTime
		}
		if _, err := o.Update(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}

	c.jsonResult(enums.JRCodeSucc, "保存成功", m.Id)//返回学生id,用于页面判断高亮哪一行数据。
}

func (c *StudentInfoController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))

	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}

	query := orm.NewOrm().QueryTable(models.StudentInfoTBName())

	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
