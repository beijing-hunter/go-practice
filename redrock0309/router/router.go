package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redrock0309/controller"
	"redrock0309/middleware"
)

func SetupRouter(app *gin.Engine)  {
	app.POST("/register", controller.UserRegister)
	app.POST("/login",controller.UserLogin)

	app.Use(middleware.User)//token 拦截验证

	app.POST("/vote/add", controller.AddVote)
	app.POST("/vote/update/time",controller.UpdateVoteTime)
	app.POST("/vote/list",controller.VoteList)
	app.POST("/vote/xs/partake",controller.PartakeVote)
	app.POST("/vote/xs/exit",controller.ExitVote)
	app.POST("/vote/xs/list",controller.PartakeVoteList)
	app.POST("/vote/result/tp",controller.Tp)
	fmt.Println("start")
}


