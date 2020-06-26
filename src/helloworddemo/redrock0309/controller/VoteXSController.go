package controller


import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"redrock0309/sql"
	"redrock0309/models"
	"redrock0309/token"
	"strconv"
)

//用户报名参与赛事
func PartakeVote(c *gin.Context)  {
	
	tokenValue:= c.GetHeader("token")
	vidStr:=c.PostForm("vid")
	vidInt,_:=strconv.Atoi(vidStr)  
	uid,username,_:=token.CheckToken(tokenValue)

	db:=sql.DbConn()
	defer db.Close()

	var u []models.Voteuser
	db.Where("voteid=? and xsuserid=?", vidInt,uid).Find(&u)

	if len(u)>0{
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "重复报名参与赛事"})
		return
	}

	db.AutoMigrate(&models.Voteuser{})
	db.Create(&models.Voteuser{ Voteid: vidInt, Xsuserid: uid,Xsusername:username,Votetotalcount:0})
	c.JSON(200, gin.H{"status": http.StatusOK, "message": "报名参与赛事成功"})
}

//用户退出赛事
func ExitVote(c *gin.Context)  {
	
	tokenValue:= c.GetHeader("token")
	vidStr:=c.PostForm("vid")
	vidInt,_:=strconv.Atoi(vidStr)  
	uid,_,_:=token.CheckToken(tokenValue)

	db:=sql.DbConn()
	defer db.Close()

	db.Delete(&models.Voteuser{ Voteid: vidInt, Xsuserid: uid})
	c.JSON(200, gin.H{"status": http.StatusOK, "message": "退出赛事成功"})
}

//获取参与赛事的用户信息及排行榜，根据投票个数从高到底排序
func PartakeVoteList(c *gin.Context)  {

	vidStr:=c.PostForm("vid")
	vidInt,_:=strconv.Atoi(vidStr)  

	db:=sql.DbConn()
	defer db.Close()

	var datas []models.Voteuser
	db.Where("voteid=?", vidInt).Order("votetotalcount desc").Find(&datas)

	c.JSON(200, gin.H{"status": http.StatusOK, "message": "参与赛事的用户信息及排行榜","datas":datas})
}