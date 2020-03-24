package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"redrock0309/sql"
	"redrock0309/models"
	"redrock0309/token"
	"strconv"
	"fmt"
	"encoding/json"
	"time"
	"strings"
)

//用户给参赛选手投票
func Tp(c *gin.Context)  {
	
	tokenValue:= c.GetHeader("token")
	vidStr:=c.PostForm("vid")
	vidInt,_:=strconv.Atoi(vidStr) 

	xsUserIdStr:=c.PostForm("xsuserid")
	xsUserIdInt,_:=strconv.Atoi(xsUserIdStr) 

	uid,username,_:=token.CheckToken(tokenValue)

	redisClient:=sql.RedisConnect();//1.链接redis
	voteInfoStr, _ := redisClient.Get(vidStr).Result()//2.从redis中获取赛事信息，从redis获取赛事信息比从数据库中读取信息更快，尤其是在高并发的情况下。
	var voteInfo models.Vote

	if len(voteInfoStr)>0 {//3.redis中缓存了赛事信息
		
		err:=json.Unmarshal([]byte(voteInfoStr),&voteInfo)//4.将redis中的缓存的赛事信息，转换成models.Vote结构体

		if err!=nil{
			fmt.Println("反序列化失败3",err)
			return
		}
	}else{//5.redis中没有缓存赛事信息

		db:=sql.DbConn()
		defer db.Close()
		db.Where("id=?", vidInt).First(&voteInfo)//6.从数据库中读取赛事信息

		dataByte,err:=json.Marshal(voteInfo)//7.将赛事信息转换成json string

		if err!=nil{
			fmt.Println("序列化失败",err)
			return
		}

		dataStr:=string(dataByte)
		redisClient.Set(vidStr, dataStr, 60*time.Second).Err()//8.将从数据库中读取到的赛事信息缓存到redis，下次直接从redis中获取赛事信息， 缓存过期时间60s
	}

	voteInfo.Starttime=strings.ReplaceAll(voteInfo.Starttime, "T", " ")
	voteInfo.Starttime=strings.ReplaceAll(voteInfo.Starttime, "+08:00", "")
	voteInfo.Endtime=strings.ReplaceAll(voteInfo.Endtime, "T", " ")
	voteInfo.Endtime=strings.ReplaceAll(voteInfo.Endtime, "+08:00", "")
	nowTime:=time.Now() //获取当前时间
	sTime,_:=time.Parse("2006-01-02 15:04:05", voteInfo.Starttime)
	eTime,_:=time.Parse("2006-01-02 15:04:05", voteInfo.Endtime)
	fmt.Println(voteInfo)
	

	if nowTime.Unix()<sTime.Unix(){
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "投票时间未开始"})
		return
	}

	if nowTime.Unix()>eTime.Unix(){
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "投票已结束"})
		return
	}

	nowTimeStr:=nowTime.Format("2006-01-02")
	voteCountKey:=vidStr+strconv.Itoa(uid)+nowTimeStr
	result, _ := redisClient.Incr(voteCountKey).Result()

	if result>3{
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "今日3次投票机会已用完"})
		return
	}

	go saveResult(vidInt,xsUserIdInt,uid,username)
	c.JSON(200, gin.H{"status": http.StatusOK, "message": "投票成功"})
}

func saveResult(voteId int,xsuserid int ,uid int,tpusername string){

	db:=sql.DbConn()
	defer db.Close()

	var voteUser models.Voteuser
	db.Where("voteid = ? AND xsuserid >= ?", voteId, xsuserid).First(&voteUser)
	voteUser.Votetotalcount=voteUser.Votetotalcount+1

	db.Model(voteUser).Updates(models.Voteuser{Votetotalcount:voteUser.Votetotalcount})

	var u models.User
	db.Where("id=?", xsuserid).First(&u)
	

	db.AutoMigrate(&models.Voteresult{})
	db.Create(&models.Voteresult{ Voteid: voteId, Xsuserid: xsuserid,Tpuserid:uid,Xsusername:u.Username,Tpusername:tpusername})
}

