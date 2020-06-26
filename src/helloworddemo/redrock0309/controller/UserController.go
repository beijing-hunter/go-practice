package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"redrock0309/sql"
	"redrock0309/models"
	"redrock0309/token"
)

//用户注册
func UserRegister(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	db:=sql.DbConn()
	defer db.Close()

	var u []models.User
	db.Where("username=?", username).Find(&u)

	if len(u)>0{
		c.JSON(200, gin.H{"status": http.StatusOK, "message": "用户名称重复"})
		return
	}


	db.AutoMigrate(&models.User{})
	db.Create(&models.User{ Username: username, Password: password})
	c.JSON(200, gin.H{"status": http.StatusOK, "message": "注册成功"})
}

func UserLogin(c *gin.Context){
	username := c.PostForm("username")
	//数据库连接
	password := c.PostForm("password")
	db:=sql.DbConn()
	defer db.Close()

	var u models.User
	db.Where("username=?", username).First(&u)
	tokens := token.Create(u.Username, u.Id)//创建token

	if password == u.Password {
		c.JSON(200,gin.H{"status:":http.StatusOK,"message":"登录成功","token":&tokens})
	}else {
		c.JSON(400,gin.H{"status:":http.StatusBadRequest,"message":"用户名或密码错误"})
	}
}

