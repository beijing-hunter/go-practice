package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"redrock0309/sql"
	"redrock0309/models"
	"strconv"
)

//添加赛事
func AddVote(c *gin.Context)  {
	
	vname:= c.PostForm("vname")
	starttime:=c.PostForm("starttime")
	endtime:=c.PostForm("endtime")

	db:=sql.DbConn()
	defer db.Close()

	db.AutoMigrate(&models.Vote{})
	db.Create(&models.Vote{ Vname: vname, Starttime: starttime,Endtime:endtime})
	c.JSON(200, gin.H{"status": http.StatusOK, "message": "赛事创建成功"})
}

//修改赛事时间
func UpdateVoteTime(c *gin.Context)  {
	
	vid:= c.PostForm("vid")
	starttime:=c.PostForm("starttime")
	endtime:=c.PostForm("endtime")
	vidInt,_:=strconv.Atoi(vid)  

	db:=sql.DbConn()
	defer db.Close()

	v:=&models.Vote{Id:int(vidInt)}
	db.Model(v).Updates(models.Vote{Starttime:starttime,Endtime:endtime})

	c.JSON(200, gin.H{"status": http.StatusOK, "message": "赛事时间更新成功"})
}

//获取赛事数据列表
func VoteList(c *gin.Context)  {
	 
	db:=sql.DbConn()
	defer db.Close()

	var datas []models.Vote
	db.Find(&datas)

	c.JSON(200, gin.H{"status": http.StatusOK, "message": "赛事数据列表","datas":&datas})
}

